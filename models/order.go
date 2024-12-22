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

import (
	"fmt"
	"github.com/go-the-way/unipay/deps/pkg"
	"time"
)

type (
	Order       = UnipayOrder
	UnipayOrder struct {
		Id             string `gorm:"column:id;type:varchar(50);primaryKey;comment:id" json:"id"`                                                                     // 订单id
		BusinessId1    string `gorm:"column:business_id1;type:varchar(50);not null;default:'';comment:业务id1;index" json:"business_id1"`                               // 业务id1
		BusinessId2    string `gorm:"column:business_id2;type:varchar(50);not null;default:'';comment:业务id2;index" json:"business_id2"`                               // 业务id2
		BusinessId3    string `gorm:"column:business_id3;type:varchar(50);not null;default:'';comment:业务id3;index" json:"business_id3"`                               // 业务id3
		TradeId        string `gorm:"column:trade_id;type:varchar(100);not null;default:'';comment:第三方支付平台id;index" json:"trade_id"`                                  // 第三方支付平台id
		PayChannelId   uint   `gorm:"column:pay_channel_id;type:uint;not null;default:0;comment:支付通道Id;index" json:"pay_channel_id"`                                  // 支付通道Id
		PayChannelName string `gorm:"column:pay_channel_name;type:varchar(50);not null;default:'';comment:支付通道名称" json:"pay_channel_name"`                            // 支付通道名称
		PayChannelType string `gorm:"column:pay_channel_type;type:varchar(10);not null;default:'normal';comment:类型 normal/trc20/erc20;index" json:"pay_channel_type"` // 类型 normal/trc20/erc20
		AmountYuan     string `gorm:"column:amount_yuan;type:varchar(20);not null;default:'';comment:交易金额（元）" json:"amount_yuan"`                                     // 交易金额（元）
		AmountFen      string `gorm:"column:amount_fen;type:varchar(20);not null;default:'';comment:交易金额（分）" json:"amount_fen"`                                       // 交易金额（分）
		Message        string `gorm:"column:message;type:varchar(500);not null;default:'';comment:支付结果信息" json:"message"`                                             // 支付结果信息
		PayPageUrl     string `gorm:"column:pay_page_url;type:varchar(500);not null;default:'';comment:支付页面Url" json:"pay_page_url"`                                  // 支付页面Url
		PayQrUrl       string `gorm:"column:pay_qr_url;type:varchar(500);not null;default:'';comment:支付二维码Url" json:"pay_qr_url"`                                     // 支付二维码Url
		NotifyUrl      string `gorm:"column:notify_url;type:varchar(1000);not null;default:'';comment:回调地址" json:"notify_url"`                                        // 回调地址
		State          byte   `gorm:"column:state;type:tinyint;not null;default:1;comment:支付状态：1待支付2已支付3已取消;index" json:"state"`                                      // 支付状态：1待支付2已支付3已取消
		// 当类型为 trc20/erc20时 这个字段保存的是 钱包地址
		Other1 string `gorm:"column:other1;type:varchar(100);not null;default:'';comment:其他字段1;index" json:"other1"` // 其他字段1
		// 当类型为 trc20/erc20时 这个字段保存的是 金额元
		Other2 string `gorm:"column:other2;type:varchar(100);not null;default:'';comment:其他字段2;index" json:"other2"` // 其他字段2
		// 当类型为 trc20/erc20时 这个字段保存的是 金额分
		Other3     string `gorm:"column:other3;type:varchar(100);not null;default:'';comment:其他字段3;index" json:"other3"`   // 其他字段3
		Remark1    string `gorm:"column:remark1;type:varchar(500);not null;default:'';comment:备注1" json:"remark1"`         // 备注1
		Remark2    string `gorm:"column:remark2;type:varchar(500);not null;default:'';comment:备注2" json:"remark2"`         // 备注2
		Remark3    string `gorm:"column:remark3;type:varchar(500);not null;default:'';comment:备注3" json:"remark3"`         // 备注3
		CreateTime string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"` // 创建时间
		UpdateTime string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"` // 修改时间
		PayTime    string `gorm:"column:pay_time;type:varchar(20);not null;default:'';comment:支付时间" json:"pay_time"`       // 支付时间
		CancelTime string `gorm:"column:cancel_time;type:varchar(20);not null;default:'';comment:取消时间" json:"cancel_time"` // 取消时间
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

const (
	OrderTypeNormal = "normal"
	OrderTypeErc20  = "erc20"
	OrderTypeTrc20  = "trc20"
)

func (o *Order) CancelTimeBeforeNow() (yes bool) {
	if o.CancelTime == "" {
		return false
	}
	return o.CancelTimeBeforeTime(pkg.ParseTime(pkg.TimeNowStr()))
}

func (o *Order) CancelTimeBeforeTimeStr(str string) (yes bool) {
	return o.CancelTimeBeforeTime(pkg.ParseTime(str))
}

func (o *Order) CancelTimeBeforeTime(t time.Time) (yes bool) {
	if o.CancelTime == "" {
		return false
	}
	return pkg.ParseTime(o.CancelTime).Before(t)
}

func (o *Order) CreateTimeBeforeNow() (yes bool) {
	return o.CreateTimeBeforeTime(pkg.ParseTime(pkg.TimeNowStr()))
}

func (o *Order) CreateTimeBeforeTimeStr(str string) (yes bool) {
	return o.CreateTimeBeforeTime(pkg.ParseTime(str))
}

func (o *Order) CreateTimeBeforeTime(t time.Time) (yes bool) {
	return pkg.ParseTime(o.CreateTime).Before(t)
}

func (o *Order) LockKey() string {
	return fmt.Sprintf("%s-%s", o.Other1, o.Other2)
}
