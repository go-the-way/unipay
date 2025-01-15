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
	"errors"
	"fmt"

	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/models"
	"github.com/go-the-way/unipay/services/base"
)

func (r *AddReq) Check() (err error) {
	if r.PayChannelId > 0 {
		return base.CheckChannelExist(r.PayChannelId)
	}
	return nil
}

func (r *UpdateReq) Check() (err error) { return base.CheckOrderExist(r.Id) }

func (r *DelReq) Check() (err error) { return base.CheckOrderExist(r.Id) }

func (r *PaidReq) Check() (err error) {
	type paid struct {
		Id    string
		State byte
	}
	var pd paid
	if err = db.GetDb().Model(new(models.Order)).Select("id", "state").Where("id=?", r.Id).Scan(&pd).Error; err != nil {
		return
	}
	if pd.Id == "" {
		return errors.New(fmt.Sprintf("支付订单[%s]不存在", r.Id))
	}
	if pd.State != models.OrderStateWaitPay {
		return errors.New(fmt.Sprintf("支付订单[%s]当前不支持此操作", r.Id))
	}
	return
}

func (r *CancelReq) Check() (err error) { return (&PaidReq{IdReq: IdReq{Id: r.Id}}).Check() }

func (r *AddReq) Transform() *models.Order {
	if r.PayChannelId > 0 && r.PayChannelName == "" {
		_ = db.GetDb().Model(new(models.Channel)).Select("name").Where("id=?", r.PayChannelId).Scan(&r.PayChannelName).Error
	}
	if r.OrderId == "" {
		r.OrderId = pkg.RandStr(30)
	}
	if r.Upgrade == 0 {
		r.Upgrade = models.OrderNoUpgrade
	}
	return &models.Order{
		Id:             r.OrderId,
		BusinessId1:    r.BusinessId1,
		BusinessId2:    r.BusinessId2,
		BusinessId3:    r.BusinessId3,
		PayChannelId:   r.PayChannelId,
		PayChannelName: r.PayChannelName,
		PayChannelType: r.PayChannelType,
		AmountYuan:     r.AmountYuan,
		AmountFen:      r.AmountFen,
		Message:        r.Message,
		PayPageUrl:     r.PayPageUrl,
		PayQrUrl:       r.PayQrUrl,
		NotifyUrl:      r.NotifyUrl,
		State:          models.OrderStateWaitPay,
		Other1:         r.Other1,
		Other2:         r.Other2,
		Other3:         r.Other3,
		Remark1:        r.Remark1,
		Remark2:        r.Remark2,
		Remark3:        r.Remark3,
		Upgrade:        r.Upgrade,
		CreateTime:     pkg.TimeNowStr(),
		UpdateTime:     pkg.TimeNowStr(),
	}
}

func (r *UpdateReq) Transform() *models.Order {
	return &models.Order{
		BusinessId1: r.BusinessId1,
		BusinessId2: r.BusinessId2,
		BusinessId3: r.BusinessId3,
		Message:     r.Message,
		Remark1:     r.Remark1,
		Remark2:     r.Remark2,
		Remark3:     r.Remark3,
		UpdateTime:  pkg.TimeNowStr(),
	}
}
