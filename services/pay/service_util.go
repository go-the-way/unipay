// Copyright 2024 unipay Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pay

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/models"
	"github.com/rwscode/unipay/services/channel"
	"github.com/rwscode/unipay/services/channelparam"
	"github.com/rwscode/unipay/services/order"
)

func channelValid(c channel.GetResp, req Req) (err error) {
	if err = channelStateValid(c); err != nil {
		return
	}
	return channelAmountValid(c, req)
}

func channelStateValid(c channel.GetResp) (err error) {
	if c.State == models.ChannelStateDisable {
		err = errors.New("该支付通道已禁用")
		return
	}
	return
}

func channelAmountValid(c channel.GetResp, req Req) (err error) {
	avCond := c.AmountValidateCond
	if avCond == "" {
		return
	}
	if valid := pkg.ValidAmount(req.Amount, c.AmountType == models.ChannelAmountTypeYuan, avCond); !valid {
		err = errors.New(fmt.Sprintf("支付金额不符合该通道验证条件"))
		return
	}
	return
}

func reqDo(c channel.GetResp, cp channelparam.GetChannelIdResp, params map[string]any) (respMap map[string]any, err error) {
	parsedUrl, pErr := url.Parse(c.ReqUrl)
	if pErr != nil {
		err = errors.New("解析支付请求Url错误：" + pErr.Error())
		return
	}
	skipInsecure := parsedUrl.Scheme == "https"
	reqBody, postForm, contentType := getReqCT(c, cp, params)
	req, _ := http.NewRequest(c.ReqMethod, c.ReqUrl, strings.NewReader(reqBody))
	req.Header = make(http.Header)
	req.Header.Set("Content-Type", contentType)
	if len(postForm) > 0 {
		req.PostForm = postForm
	}
	client := &http.Client{Timeout: time.Minute}
	if skipInsecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	resp, rErr := client.Do(req)
	if rErr != nil {
		err = errors.New("支付请求错误：" + rErr.Error())
		return
	}
	buf, bErr := io.ReadAll(resp.Body)
	if bErr != nil {
		err = errors.New(fmt.Sprintf("支付请求不成功，HTTP状态码：%d，读取响应错误：%s", resp.StatusCode, bErr.Error()))
		return
	}
	result := string(buf)
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("支付请求不成功，HTTP状态码：%d，响应结果：%s", resp.StatusCode, result))
		return
	}
	respMap = map[string]any{}
	if err = json.Unmarshal(buf, &respMap); err != nil {
		err = errors.New(fmt.Sprintf("请求不成功，HTTP状态码：%d，原始数据：%s，解析JSON错误：%s", resp.StatusCode, result, err.Error()))
		return
	}
	return
}

func reqCallback(req Req, c channel.GetResp, respMap map[string]any, orderId string) (resp Resp, err error) {
	reqSuccess, pErr := pkg.EvalBool(c.ReqSuccessExpr, respMap)
	if pErr != nil {
		err = errors.New(fmt.Sprintf("支付请求成功，但是解析请求成功计算表达式：%s，错误：%s", c.ReqSuccessExpr, pErr.Error()))
		return
	}
	if !reqSuccess {
		err = errors.New(fmt.Sprintf("支付请求成功，但是解析请求成功计算表达式：%s，计算结果：%v,为不成功", c.ReqSuccessExpr, reqSuccess))
		return
	}

	var (
		pageUrl, qrUrl, message string
	)

	if expr := c.ReqPayPageUrlExpr; expr != "" {
		pageUrl, err = pkg.EvalString(expr, respMap)
		if err != nil {
			err = errors.New(fmt.Sprintf("支付请求成功，但是解析请求支付页面Url获取表达式：%s，错误：%s", expr, pErr.Error()))
			return
		}
	}

	if expr := c.ReqPayQrUrlExpr; expr != "" {
		qrUrl, err = pkg.EvalString(expr, respMap)
		if err != nil {
			err = errors.New(fmt.Sprintf("支付请求成功，但是解析请求支付页面Url获取表达式：%s，错误：%s", expr, pErr.Error()))
			return
		}
	}

	if expr := c.ReqPayMessageExpr; expr != "" {
		message, err = pkg.EvalString(expr, respMap)
		if err != nil {
			err = errors.New(fmt.Sprintf("支付请求成功，但是解析请求支付获取消息表达式：%s，错误：%s", expr, pErr.Error()))
			return
		}
	}

	resp.OrderId = orderId
	resp.PayPageUrl = pageUrl
	resp.PayQrUrl = qrUrl
	resp.Message = message

	if err = order.Service.Add(buildOrderAddReq(c, req, resp)); err != nil {
		return
	}

	if fn := req.Callback; fn != nil {
		go fn(req)
	}

	return
}

func buildOrderAddReq(c channel.GetResp, req Req, resp Resp) order.AddReq {
	return order.AddReq{
		PayChannelId:   c.Id,
		PayChannelName: c.Name,
		BusinessId1:    req.BusinessId1,
		BusinessId2:    req.BusinessId2,
		BusinessId3:    req.BusinessId3,
		Amount:         req.Amount,
		AmountYuan:     req.AmountYuan,
		Message:        resp.Message,
		Remark1:        req.Remark1,
		Remark2:        req.Remark2,
		Remark3:        req.Remark3,
		OrderId:        resp.OrderId,
		PayPageUrl:     resp.PayPageUrl,
		PayQrUrl:       resp.PayQrUrl,
	}
}
