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

package e20svc

import (
	_ "embed"

	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/models"
)

var (
	//go:embed e20.html
	e20Html0 string
	//go:embed error.html
	errorHtml0 string

	e20Html, errorHtml string
)

func SetE20Html(html0 string)   { e20Html = html0 }
func SetErrorHtml(html0 string) { errorHtml = html0 }

func getE20Html() string {
	if e20Html == "" {
		return e20Html0
	}
	return e20Html
}

func getErrorHtml() string {
	if errorHtml == "" {
		return errorHtml0
	}
	return errorHtml
}

type service struct{}

func (s *service) OrderPayHtml(req OrderPayHtmlReq) (resp OrderPayHtmlResp, err error) {
	var order models.Order
	_ = db.GetDb().Model(new(models.Order)).Where("id=?", req.OrderId).First(&order).Error
	if order.Id == "" {
		resp.Html = fmt.Sprintf(getErrorHtml(), fmt.Sprintf("订单[%s]不存在", req.OrderId), req.Platform, req.RedirectUrl)
		return
	}
	if order.State == models.OrderStatePaid {
		resp.Html = fmt.Sprintf(getErrorHtml(), fmt.Sprintf("订单[%s]已支付", order.Id), req.Platform, req.RedirectUrl)
		return
	}
	if order.State == models.OrderStateCancelled {
		resp.Html = fmt.Sprintf(getErrorHtml(), fmt.Sprintf("订单[%s]已失效", order.Id), req.Platform, req.RedirectUrl)
		return
	}
	var validPeriodMinute uint
	if err = db.GetDb().Model(new(models.ApiConfig)).Where("id=1").Select("valid_period_minute").Scan(&validPeriodMinute).Error; err != nil {
		return
	}
	if validPeriodMinute == 0 {
		validPeriodMinute = 5
	}
	dur, _ := time.ParseDuration(fmt.Sprintf("%dm", validPeriodMinute))
	expireTime := pkg.ParseTime(order.CreateTime).Add(dur)
	nowTime := pkg.ParseTime(pkg.FormatTime(time.Now()))
	if nowTime.After(expireTime) {
		// 订单已失效
		err = errors.New(fmt.Sprintf("订单[%s]已失效", order.Id))
		return
	}
	expireTimeUnixMilli := fmt.Sprintf("%d", expireTime.UnixMilli())
	return s.E20Html(E20HtmlReq{
		OrderId:            order.Id,
		Protocol:           order.PayChannelType,
		Amount:             order.Other2,
		Address:            order.Other1,
		ExpirationTime:     expireTimeUnixMilli,
		CheckOrderStateUrl: req.CheckOrderStateUrl,
		RedirectUrl:        req.RedirectUrl,
		Platform:           req.Platform,
	})
}

func (s *service) E20Html(req E20HtmlReq) (resp E20HtmlResp, err error) {
	buf := &bytes.Buffer{}
	tpl := template.Must(template.New("").Parse(getE20Html()))
	if err = tpl.Execute(buf, req); err != nil {
		err = errors.New("解析错误：" + err.Error())
		return
	}
	resp.Html = buf.String()
	buf.Reset()
	return
}
