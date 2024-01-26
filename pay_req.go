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

package unipay

import (
	"fmt"
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/services/channel"
	"github.com/rwscode/unipay/services/channelparam"
	"net/url"
)

// ReqPay 请求支付接口
func ReqPay(req ReqPayReq) (resp ReqPayResp, err error) {
	pm, err := ChannelService.Get(channel.GetReq{Id: req.ChannelId})
	if err != nil {
		return
	}

	if err = channelValid(pm, req); err != nil {
		return
	}

	pmm, err := ChannelParamService.GetChannelId(channelparam.GetChannelIdReq{ChannelId: req.ChannelId})
	if err != nil {
		return
	}

	// 订单id
	orderId := pkg.RandStr(30)

	evalEdParams, err := evalParams(req, pm, pmm, orderId)
	if err != nil {
		return
	}

	respMap, err := reqDo(pm, pmm, evalEdParams)

	return reqCallback(req, pm, respMap, orderId)
}

type (
	ReqPayReq struct {
		ChannelId  uint   `validate:"min(1,支付通道id不能为空)" form:"channel_id" json:"channel_id"`                                    // 支付通道id
		Amount     uint   `validate:"min(1,支付金额不能少于1)" form:"amount" json:"amount"`                                             // 支付金额（单位：分）
		Subject    string `validate:"minlength(1,支付主题不能为空) maxlength(200,支付主题长度不能超过200)" form:"subject" json:"subject"`         // 支付主题
		ClientIp   string `validate:"minlength(1,客户端Ip不能为空) maxlength(20,客户端Ip长度不能超过50)" form:"client_ip" json:"client_ip"`     // 客户端Ip
		NotifyUrl  string `validate:"minlength(1,回调Url不能为空) maxlength(500,回调Url长度不能超过500)" form:"notify_url" json:"notify_url"` // 回调Url
		BusinessId string `validate:"minlength(1,业务Id不能为空) maxlength(50,业务Id长度不能超过50)" form:"business_id" json:"business_id"`   // 业务Id
	}
	ReqPayResp struct {
		OrderId string `json:"order_id"` // 订单Id
		PageUrl string `json:"page_url"` // 页面url
		QrUrl   string `json:"qr_url"`   // 二维码url
		Message string `json:"message"`  // 支付信息
	}
	NotifyPayReq struct {
		ChannelId  uint   `validate:"min(1,支付通道id不能为空)" form:"channel_id" json:"channel_id"`                                  // 支付通道id
		OrderId    string `validate:"minlength(1,订单Id不能为空) maxlength(50,订单Id长度不能超过50)" form:"order_id" json:"order_id"`       // 订单Id
		BusinessId string `validate:"minlength(1,业务Id不能为空) maxlength(50,业务Id长度不能超过50)" form:"business_id" json:"business_id"` // 业务Id
	}
)

func (r *ReqPayReq) ToMap(orderId string) map[string]any {
	nu, _ := url.Parse(r.NotifyUrl)
	nu.Query().Set("unipay_channel_id", fmt.Sprintf("%d", r.ChannelId))
	nu.Query().Set("unipay_business_id", r.BusinessId)
	nu.Query().Set("unipay_order_id", orderId)
	notifyUrl := nu.String()
	return map[string]any{
		"ChannelId":  r.ChannelId,
		"Amount":     r.Amount,
		"AmountYuan": r.Amount * 100,
		"AmountFen":  r.Amount,
		"Subject":    r.Subject,
		"ClientIp":   r.ClientIp,
		"NotifyUrl":  notifyUrl,
		"BusinessId": r.BusinessId,
	}
}
