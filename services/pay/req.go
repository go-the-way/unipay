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

import "github.com/go-the-way/unipay/models"

type (
	Req struct {
		ChannelId        uint   `validate:"min(1,支付通道id不能为空)" form:"channel_id" json:"channel_id"`                                      // 支付通道id
		AmountFen        string `validate:"minlength(1,支付金额分不能为空)" form:"amount_fen" json:"amount_fen"`                                 // 支付金额（单位：分）
		AmountYuan       string `validate:"minlength(1,支付金额元不能为空)" form:"amount_yuan" json:"amount_yuan"`                               // 支付金额（单位：元）
		AmountCurrency   string `validate:"enum(USD|CNY,支付金额货币不合法)" form:"amount_currency" json:"amount_currency"`                      // 支付金额货币 美元USD 人民币CNY
		CurrencyRateType byte   `validate:"enum(1|2,货币汇率类型不合法)" form:"currency_rate_type" json:"currency_rate_type"`                    // 货币汇率类型 1美元兑人民币（人民币=支付金额*汇率） 2人民币兑美元（美元=支付金额/汇率）
		Subject          string `validate:"minlength(1,支付主题不能为空) maxlength(200,支付主题长度不能超过200)" form:"subject" json:"subject"`           // 支付主题
		ClientIp         string `validate:"minlength(1,客户端Ip不能为空) maxlength(20,客户端Ip长度不能超过50)" form:"client_ip" json:"client_ip"`       // 客户端Ip
		NotifyUrl        string `validate:"minlength(1,回调Url不能为空) maxlength(500,回调Url长度不能超过500)" form:"notify_url" json:"notify_url"`   // 回调Url
		ReturnUrl        string `validate:"minlength(1,页面Url不能为空) maxlength(500,页面Url长度不能超过500)" form:"return_url" json:"return_url"`   // 页面Url
		BusinessId1      string `validate:"minlength(1,业务id1不能为空) maxlength(50,业务id1长度不能超过50)" form:"business_id1" json:"business_id1"` // 业务id1
		BusinessId2      string `validate:"maxlength(50,业务id2长度不能超过50)" form:"business_id2" json:"business_id2"`                        // 业务id2
		BusinessId3      string `validate:"maxlength(50,业务id3长度不能超过50)" form:"business_id3" json:"business_id3"`                        // 业务id3
		Other1           string `validate:"maxlength(500,其他1长度不能超过500)" form:"other1" json:"other1"`                                    // 其他1
		Other2           string `validate:"maxlength(500,其他2长度不能超过500)" form:"other2" json:"other2"`                                    // 其他2
		Other3           string `validate:"maxlength(500,其他3长度不能超过500)" form:"other3" json:"other3"`                                    // 其他3
		Remark1          string `validate:"maxlength(500,备注1长度不能超过500)" form:"remark1" json:"remark1"`                                  // 备注1
		Remark2          string `validate:"maxlength(500,备注2长度不能超过500)" form:"remark2" json:"remark2"`                                  // 备注2
		Remark3          string `validate:"maxlength(500,备注3长度不能超过500)" form:"remark3" json:"remark3"`                                  // 备注3
		Upgrade          byte   `json:"upgrade"`                                                                                        // 是否升级订单1:是2不是

		Platform      byte   `validate:"enum(1|2|3|4,平台不合法)" form:"platform" json:"platform"`                                       // 平台 1Android 2iOS 3Web
		AppWakeUri    string `gorm:"column:app_wake_uri;type:varchar(50);not null;default:'';comment:App唤醒URI" json:"app_wake_uri"` // App唤醒URI
		E20PayPageUrl string `form:"-" json:"-"`                                                                                    // erc20/trc20支付页面Url

		Callback func(req Req)
	}
	NotifyReq struct {
		ChannelId   uint   `validate:"min(1,支付通道id不能为空)" uri:"channel_id" form:"channel_id" json:"channel_id"`                                        // 支付通道id
		OrderId     string `validate:"minlength(1,订单id不能为空) maxlength(50,订单id长度不能超过50)" uri:"order_id" form:"order_id" json:"order_id"`               // 订单id
		BusinessId1 string `validate:"minlength(1,业务id1不能为空) maxlength(50,业务id1长度不能超过50)" uri:"business_id1" form:"business_id1" json:"business_id1"` // 业务id1
		BusinessId2 string `validate:"maxlength(50,业务id2长度不能超过50)" uri:"business_id2" form:"business_id2" json:"business_id2"`                        // 业务id2
		BusinessId3 string `validate:"maxlength(50,业务id3长度不能超过50)" uri:"business_id3" form:"business_id3" json:"business_id3"`                        // 业务id3

		Callback func(req NotifyReq, order *models.Order)
	}
)
