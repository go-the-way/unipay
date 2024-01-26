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

package order

import "github.com/rwscode/unipay/services/base"

type (
	GetPageReq struct {
		base.PageReq

		Id           string `form:"id"`             // id
		BusinessId   string `form:"business_id"`    // 业务id
		TradeId      string `form:"trade_id"`       // 第三方支付平台交易id
		PayChannelId uint   `form:"pay_channel_id"` // 支付通道Id
		Amount       uint   `form:"amount"`         // 交易金额
		Amount1      uint   `form:"amount1"`        // 交易金额
		Amount2      uint   `form:"amount2"`        // 交易金额
		Message      string `form:"message"`        // 支付结果信息
		State        byte   `form:"state"`          // 支付状态：1待支付2支付成功3支付失败
		Remark       string `form:"remark"`         // 备注
		CreateTime1  string `form:"create_time1"`   // 创建时间
		CreateTime2  string `form:"create_time2"`   // 创建时间
		PayTime1     string `form:"pay_time1"`      // 支付时间
		PayTime2     string `form:"pay_time2"`      // 支付时间
		UpdateTime1  string `form:"update_time1"`   // 修改时间
		UpdateTime2  string `form:"update_time2"`   // 修改时间
	}
	GetReq struct {
		Id string `validate:"minlength(1,订单id不能为空)" form:"id"`
	}
	AddReq struct {
		PayChannelId   uint   `validate:"min(1,支付通道id不能为空)" json:"pay_channel_id"` // 支付通道Id
		PayChannelName string `json:"-"`
		BusinessId     string `validate:"minlength(1,业务id不能为空) maxlength(50,业务id长度不能超过50)" json:"business_id"` // 业务id
		Amount         uint   `validate:"min(1,交易金额不能少于1)" json:"amount"`                                      // 交易金额（单位：分）
		Message        string `validate:"maxlength(500,支付结果信息长度不能超过500)" json:"message"`                       // 支付结果信息
		Remark         string `validate:"maxlength(500,备注长度不能超过500)" json:"remark"`                            // 备注

		OrderId string `json:"-"`
	}
	idReq struct {
		Id string `validate:"minlength(1,订单id不能为空) maxlength(50,订单id长度不能超过50)" json:"id"` // 订单id
	}
	UpdateReq struct {
		idReq  `validate:"valid(T)"`
		AddReq `validate:"valid(T)"`
	}
	DelReq        idReq
	PaySuccessReq struct {
		idReq   `validate:"valid(T)"`
		TradeId string `validate:"minlength(1,第三方支付平台id不能为空) maxlength(50,第三方支付平台id长度不能超过50)" json:"trade_id"` // 第三方支付平台id
		Message string `json:"message"`                                                                        // 支付信息
	}
	PayFailureReq struct {
		idReq   `validate:"valid(T)"`
		Message string `json:"message"` // 支付信息
	}
	GetPayStateReq idReq
)
