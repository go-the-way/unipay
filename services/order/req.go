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

		OrderBy string `form:"order_by"` // 排序

		Id           string `form:"id"`             // id
		BusinessId1  string `form:"business_id1"`   // 业务id1
		BusinessId2  string `form:"business_id2"`   // 业务id2
		BusinessId3  string `form:"business_id3"`   // 业务id3
		TradeId      string `form:"trade_id"`       // 第三方支付平台交易id
		PayChannelId uint   `form:"pay_channel_id"` // 支付通道Id
		Amount       uint   `form:"amount"`         // 交易金额
		Amount1      uint   `form:"amount1"`        // 交易金额
		Amount2      uint   `form:"amount2"`        // 交易金额
		Message      string `form:"message"`        // 支付结果信息
		State        byte   `form:"state"`          // 支付状态：1待支付2支付成功3支付失败
		Remark1      string `form:"remark1"`        // 备注1
		Remark2      string `form:"remark2"`        // 备注2
		Remark3      string `form:"remark3"`        // 备注3
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
	GetBusinessIdReq struct {
		BusinessId1 string `validate:"minlength(1,业务id1不能为空)" form:"business_id1"`
		BusinessId2 string `form:"business_id2"`
		BusinessId3 string `form:"business_id3"`
	}
	GetIdAndBusinessIdReq struct {
		Id          string `validate:"minlength(1,订单id不能为空)" form:"id"`
		BusinessId1 string `validate:"minlength(1,业务id1不能为空)" form:"business_id1"`
		BusinessId2 string `form:"business_id2"`
		BusinessId3 string `form:"business_id3"`
	}
	AddReq struct {
		PayChannelId   uint   `validate:"min(1,支付通道id不能为空)" json:"pay_channel_id"` // 支付通道Id
		PayChannelName string `json:"-"`
		BusinessId1    string `validate:"minlength(1,业务id1不能为空) maxlength(50,业务id1长度不能超过50)" json:"business_id1"` // 业务id1
		BusinessId2    string `validate:"maxlength(50,业务id2长度不能超过50)" json:"business_id2"`                        // 业务id2
		BusinessId3    string `validate:"maxlength(50,业务id3长度不能超过50)" json:"business_id3"`                        // 业务id3
		Amount         uint   `validate:"min(1,交易金额不能少于1)" json:"amount"`                                         // 交易金额（单位：分）
		Message        string `validate:"maxlength(500,支付结果信息长度不能超过500)" json:"message"`                          // 支付结果信息
		Remark1        string `validate:"maxlength(500,备注1长度不能超过500)" json:"remark1"`                             // 备注1
		Remark2        string `validate:"maxlength(500,备注2长度不能超过500)" json:"remark2"`                             // 备注2
		Remark3        string `validate:"maxlength(500,备注3长度不能超过500)" json:"remark3"`                             // 备注3

		OrderId    string `json:"-"`
		PayPageUrl string `json:"-"`
		PayQrUrl   string `json:"-"`
	}
	IdReq struct {
		Id string `validate:"minlength(1,订单id不能为空) maxlength(50,订单id长度不能超过50)" json:"id"` // 订单id
	}
	UpdateReq struct {
		IdReq  `validate:"valid(T)"`
		AddReq `validate:"valid(T)"`
	}
	DelReq        IdReq
	PaySuccessReq struct {
		IdReq   `validate:"valid(T)"`
		TradeId string `validate:"minlength(1,第三方支付平台id不能为空) maxlength(50,第三方支付平台id长度不能超过50)" json:"trade_id"` // 第三方支付平台id
		Message string `json:"message"`                                                                        // 支付信息
	}
	PayFailureReq struct {
		IdReq   `validate:"valid(T)"`
		Message string `json:"message"` // 支付信息
	}
	GetPayStateReq IdReq
)
