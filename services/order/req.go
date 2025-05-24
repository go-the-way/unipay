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

import (
	"github.com/go-the-way/unipay/services/base"
	"gorm.io/gorm"
)

type (
	GetPageReq struct {
		base.PageReq

		OrderBy string `form:"order_by"` // 排序

		Id             string `form:"id"`               // id
		BusinessId1    string `form:"business_id1"`     // 业务id1
		BusinessId2    string `form:"business_id2"`     // 业务id2
		BusinessId3    string `form:"business_id3"`     // 业务id3
		TradeId        string `form:"trade_id"`         // 第三方支付平台交易id
		PayChannelId   uint   `form:"pay_channel_id"`   // 支付通道Id
		PayChannelType string `form:"pay_channel_type"` // 支付通道类型 normal/trc20/erc20
		AmountYuan     string `form:"amount_yuan"`      // 交易金额（元）
		AmountFen      string `form:"amount_fen"`       // 交易金额（分）
		Message        string `form:"message"`          // 支付结果信息
		State          byte   `form:"state"`            // 支付状态：1待支付2已支付3已取消
		Other1         string `form:"other1"`           // 其他1
		Other2         string `form:"other2"`           // 其他2
		Other3         string `form:"other3"`           // 其他3
		Remark1        string `form:"remark1"`          // 备注1
		Remark2        string `form:"remark2"`          // 备注2
		Remark3        string `form:"remark3"`          // 备注3
		Upgrade        byte   `form:"upgrade"`          // 是否升级订单1:是2不是
		CreateTime1    string `form:"create_time1"`     // 创建时间
		CreateTime2    string `form:"create_time2"`     // 创建时间
		PayTime1       string `form:"pay_time1"`        // 支付时间
		PayTime2       string `form:"pay_time2"`        // 支付时间
		UpdateTime1    string `form:"update_time1"`     // 修改时间
		UpdateTime2    string `form:"update_time2"`     // 修改时间
		CancelTime1    string `form:"cancel_time1"`     // 取消时间
		CancelTime2    string `form:"cancel_time2"`     // 取消时间

		ExtraCallback func(q *gorm.DB) `form:"-"`
	}
	IdReq struct {
		Id string `validate:"minlength(1,订单id不能为空) maxlength(50,订单id长度不能超过50)" form:"id" json:"id"` // 订单id
	}
	GetReq           IdReq
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
		PayChannelId   uint   `json:"pay_channel_id"` // 支付通道id
		PayChannelName string `json:"-"`
		PayChannelType string `json:"-"`
		BusinessId1    string `validate:"minlength(1,业务id1不能为空) maxlength(50,业务id1长度不能超过50)" json:"business_id1"` // 业务id1
		BusinessId2    string `validate:"maxlength(50,业务id2长度不能超过50)" json:"business_id2"`                        // 业务id2
		BusinessId3    string `validate:"maxlength(50,业务id3长度不能超过50)" json:"business_id3"`                        // 业务id3
		AmountFen      string `validate:"minlength(1,交易金额分不能为空)" json:"amount_fen"`                               // 交易金额（单位：分）
		AmountYuan     string `validate:"minlength(1,交易金额元不能为空)" json:"amount_yuan"`                              // 交易金额（单位：元）
		Message        string `validate:"maxlength(500,支付结果信息长度不能超过500)" json:"message"`                          // 支付结果信息
		Other1         string `validate:"maxlength(500,其他1长度不能超过500)" json:"other1"`                              // 其他1
		Other2         string `validate:"maxlength(500,其他2长度不能超过500)" json:"other2"`                              // 其他2
		Other3         string `validate:"maxlength(500,其他3长度不能超过500)" json:"other3"`                              // 其他3
		Remark1        string `validate:"maxlength(500,备注1长度不能超过500)" json:"remark1"`                             // 备注1
		Remark2        string `validate:"maxlength(500,备注2长度不能超过500)" json:"remark2"`                             // 备注2
		Remark3        string `validate:"maxlength(500,备注3长度不能超过500)" json:"remark3"`                             // 备注3
		Upgrade        byte   `json:"upgrade"`                                                                    // 是否升级订单1:是2不是
		State          byte   `json:"state"`                                                                      // 支付状态：1待支付2已支付3已取消

		OrderId    string `json:"-"`
		PayPageUrl string `json:"-"`
		PayQrUrl   string `json:"-"`
		NotifyUrl  string `json:"-"`
	}
	UpdateReq struct {
		Id          string `validate:"minlength(1,订单id不能为空) maxlength(50,订单id长度不能超过50)" form:"id" json:"id"`   // 订单id
		BusinessId1 string `validate:"minlength(1,业务id1不能为空) maxlength(50,业务id1长度不能超过50)" json:"business_id1"` // 业务id1
		BusinessId2 string `validate:"maxlength(50,业务id2长度不能超过50)" json:"business_id2"`                        // 业务id2
		BusinessId3 string `validate:"maxlength(50,业务id3长度不能超过50)" json:"business_id3"`                        // 业务id3
		Other1      string `validate:"maxlength(500,其他1长度不能超过500)" json:"other1"`                              // 其他1
		Other2      string `validate:"maxlength(500,其他2长度不能超过500)" json:"other2"`                              // 其他2
		Other3      string `validate:"maxlength(500,其他3长度不能超过500)" json:"other3"`                              // 其他3
		Remark1     string `validate:"maxlength(500,备注1长度不能超过500)" json:"remark1"`                             // 备注1
		Remark2     string `validate:"maxlength(500,备注2长度不能超过500)" json:"remark2"`                             // 备注2
		Remark3     string `validate:"maxlength(500,备注3长度不能超过500)" json:"remark3"`                             // 备注3
		Message     string `validate:"maxlength(500,支付结果信息长度不能超过500)" json:"message"`                          // 支付结果信息
	}
	DelReq  IdReq
	PaidReq struct {
		IdReq   `validate:"valid(T)"`
		TradeId string `validate:"minlength(1,第三方支付平台id不能为空) maxlength(100,第三方支付平台id长度不能超过100)" json:"trade_id"` // 第三方支付平台id
		Message string `json:"message"`
	}
	CancelReq struct {
		IdReq      `validate:"valid(T)"`
		Message    string `json:"message"`
		CancelTime string `json:"-"`
	}
	GetStateReq IdReq
)
