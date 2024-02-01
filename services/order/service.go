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

	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/models"
)

type service struct{}

func (s *service) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDb().Model(new(models.Order))
	pkg.IfNotEmptyFunc(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfNotEmptyFunc(req.BusinessId1, func() { q.Where("business_id1=?", req.BusinessId1) })
	pkg.IfNotEmptyFunc(req.BusinessId2, func() { q.Where("business_id2=?", req.BusinessId2) })
	pkg.IfNotEmptyFunc(req.BusinessId3, func() { q.Where("business_id3=?", req.BusinessId3) })
	pkg.IfNotEmptyFunc(req.TradeId, func() { q.Where("trade_id=?", req.TradeId) })
	pkg.IfGt0Func(req.PayChannelId, func() { q.Where("pay_channel_id=?", req.PayChannelId) })
	pkg.IfGt0Func(req.Amount, func() { q.Where("amount=?", req.Amount) })
	pkg.IfGt0Func(req.Amount1, func() { q.Where("amount>=?", req.Amount1) })
	pkg.IfGt0Func(req.Amount2, func() { q.Where("amount>=?", req.Amount2) })
	pkg.IfNotEmptyFunc(req.Message, func() { q.Where("message=?", req.Message) })
	pkg.IfGt0Func(req.State, func() { q.Where("state=?", req.State) })
	pkg.IfNotEmptyFunc(req.Remark1, func() { q.Where("remark1 like concat('%',?,'%')", req.Remark1) })
	pkg.IfNotEmptyFunc(req.Remark2, func() { q.Where("remark2 like concat('%',?,'%')", req.Remark2) })
	pkg.IfNotEmptyFunc(req.Remark3, func() { q.Where("remark3 like concat('%',?,'%')", req.Remark3) })
	pkg.IfNotEmptyFunc(req.CreateTime1, func() { q.Where("create_time>=concat(?,' 00:00:00')", req.CreateTime1) })
	pkg.IfNotEmptyFunc(req.CreateTime2, func() { q.Where("create_time<=concat(?,' 23:59:59')", req.CreateTime2) })
	pkg.IfNotEmptyFunc(req.PayTime1, func() { q.Where("pay_time>=concat(?,' 00:00:00')", req.PayTime1) })
	pkg.IfNotEmptyFunc(req.PayTime2, func() { q.Where("pay_time<=concat(?,' 23:59:59')", req.PayTime2) })
	pkg.IfNotEmptyFunc(req.UpdateTime1, func() { q.Where("update_time>=concat(?,' 00:00:00')", req.UpdateTime1) })
	pkg.IfNotEmptyFunc(req.UpdateTime2, func() { q.Where("update_time<=concat(?,' 23:59:59')", req.UpdateTime2) })
	if req.OrderBy != "" {
		q.Order(req.OrderBy)
	}
	err = db.GetPagination()(q, req.Page, req.Limit, &resp.Total, &resp.List)
	return
}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
	var list []models.Order
	if err = db.GetDb().Model(new(models.Order)).Where("id=?", req.Id).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付订单[%s]不存在", req.Id))
		return
	}
	resp.Order = list[0]
	return
}

func (s *service) GetBusinessId(req GetBusinessIdReq) (resp GetBusinessIdResp, err error) {
	var list []models.Order
	if err = db.GetDb().Model(new(models.Order)).Where("id=? and business_id=?", req.Id, req.BusinessId).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付订单[%s]不存在", req.Id))
		return
	}
	resp.Order = list[0]
	return
}

func (s *service) Add(req AddReq) (err error) {
	return db.GetDb().Create(req.Transform()).Error
}

func (s *service) Update(req UpdateReq) (err error) {
	return db.GetDb().Model(&models.Order{Id: req.Id}).Omit("create_time", "pay_time").Updates(req.Transform()).Error
}

func (s *service) Del(req DelReq) (err error) {
	return db.GetDb().Delete(&models.Order{Id: req.Id}).Error
}

func (s *service) PaySuccess(req PaySuccessReq) (err error) {
	return db.GetDb().Model(&models.Order{Id: req.Id}).Updates(models.Order{Message: req.Message, State: models.OrderStatePaySuccess, PayTime: pkg.TimeNowStr()}).Error
}

func (s *service) PayFailure(req PayFailureReq) (err error) {
	return db.GetDb().Model(&models.Order{Id: req.Id}).Updates(models.Order{Message: req.Message, State: models.OrderStatePayFailure, UpdateTime: pkg.TimeNowStr()}).Error
}

func (s *service) GetPayState(req GetPayStateReq) (resp GetPayStateResp, err error) {
	err = db.GetDb().Model(&models.Order{Id: req.Id}).Select("state", "message").Scan(&resp).Error
	return
}
