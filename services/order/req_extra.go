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
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/models"
	"github.com/rwscode/unipay/services/base"
)

func (r *AddReq) Check() (err error) { return base.CheckChannelExist(r.PayChannelId) }

func (r *UpdateReq) Check() (err error) {
	if err = base.CheckOrderExist(r.Id); err != nil {
		return
	}
	return base.CheckChannelExist(r.PayChannelId)
}

func (r *DelReq) Check() (err error) { return base.CheckOrderExist(r.Id) }

func (r *AddReq) Transform() *models.Order {
	if r.PayChannelId > 0 && r.PayChannelName == "" {
		_ = db.GetDb().Model(new(models.Channel)).Select("name").Where("id=?", r.PayChannelId).Scan(&r.PayChannelName).Error
	}
	if r.OrderId == "" {
		r.OrderId = pkg.RandStr(30)
	}
	return &models.Order{
		Id:             r.OrderId,
		BusinessId1:    r.BusinessId1,
		BusinessId2:    r.BusinessId2,
		BusinessId3:    r.BusinessId3,
		PayChannelId:   r.PayChannelId,
		PayChannelName: r.PayChannelName,
		Amount:         r.Amount,
		AmountYuan:     r.Amount * 100,
		AmountFen:      r.Amount,
		Message:        r.Message,
		PayPageUrl:     r.PayPageUrl,
		PayQrUrl:       r.PayQrUrl,
		State:          models.OrderStateWaitPay,
		Remark1:        r.Remark1,
		Remark2:        r.Remark2,
		Remark3:        r.Remark3,
		CreateTime:     pkg.TimeNowStr(),
		UpdateTime:     pkg.TimeNowStr(),
	}
}

func (r *UpdateReq) Transform() *models.Order {
	m := r.AddReq.Transform()
	m.Id = r.Id
	return m
}
