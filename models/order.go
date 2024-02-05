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

package models

type (
	Order       = UnipayOrder
	UnipayOrder struct {
		Id             string `gorm:"column:id;type:varchar(50);primaryKey;comment:id" json:"id"`                                          // 订单id
		BusinessId1    string `gorm:"column:business_id1;type:varchar(50);not null;default:'';comment:业务id1;index" json:"business_id1"`    // 业务id1
		BusinessId2    string `gorm:"column:business_id2;type:varchar(50);not null;default:'';comment:业务id2;index" json:"business_id2"`    // 业务id2
		BusinessId3    string `gorm:"column:business_id3;type:varchar(50);not null;default:'';comment:业务id3;index" json:"business_id3"`    // 业务id3
		TradeId        string `gorm:"column:trade_id;type:varchar(50);not null;default:'';comment:第三方支付平台id;index" json:"trade_id"`        // 第三方支付平台id
		PayChannelId   uint   `gorm:"column:pay_channel_id;type:uint;not null;default:0;comment:支付通道Id;index" json:"pay_channel_id"`       // 支付通道Id
		PayChannelName string `gorm:"column:pay_channel_name;type:varchar(50);not null;default:'';comment:支付通道名称" json:"pay_channel_name"` // 支付通道名称
		Amount         uint   `gorm:"column:amount;type:uint;not null;default:0;comment:交易金额" json:"amount"`                               // 交易金额
		AmountYuan     uint   `gorm:"column:amount_yuan;type:uint;not null;default:0;comment:交易金额（元）" json:"amount_yuan"`                  // 交易金额（元）
		AmountFen      uint   `gorm:"column:amount_fen;type:uint;not null;default:0;comment:交易金额（分）" json:"amount_fen"`                    // 交易金额（分）
		Message        string `gorm:"column:message;type:varchar(500);not null;default:'';comment:支付结果信息" json:"message"`                  // 支付结果信息
		PayPageUrl     string `gorm:"column:pay_page_url;type:varchar(500);not null;default:'';comment:支付页面Url" json:"pay_page_url"`       // 支付页面Url
		PayQrUrl       string `gorm:"column:pay_qr_url;type:varchar(500);not null;default:'';comment:支付二维码Url" json:"pay_qr_url"`          // 支付二维码Url
		State          byte   `gorm:"column:state;type:tinyint;not null;default:1;comment:支付状态：1待支付2已支付3已取消;index" json:"state"`           // 支付状态：1待支付2已支付3已取消
		Remark1        string `gorm:"column:remark1;type:varchar(500);not null;default:'';comment:备注1" json:"remark1"`                     // 备注1
		Remark2        string `gorm:"column:remark2;type:varchar(500);not null;default:'';comment:备注2" json:"remark2"`                     // 备注2
		Remark3        string `gorm:"column:remark3;type:varchar(500);not null;default:'';comment:备注3" json:"remark3"`                     // 备注3
		CreateTime     string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`             // 创建时间
		PayTime        string `gorm:"column:pay_time;type:varchar(20);not null;default:'';comment:支付时间" json:"pay_time"`                   // 支付时间
		UpdateTime     string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"`             // 修改时间
	}
)

const (
	_ byte = iota
	// OrderStateWaitPay 待支付
	OrderStateWaitPay
	// OrderStatePaid 已支付
	OrderStatePaid
	// OrderStateCancelled 已取消
	OrderStateCancelled
)
