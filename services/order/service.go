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
)

type (
	service      struct{}
	CallbackFunc func(order *models.Order)
)

func (s *service) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDb().Model(new(models.Order))
	pkg.IfNotEmptyFunc(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfNotEmptyFunc(req.BusinessId1, func() { q.Where("business_id1=?", req.BusinessId1) })
	pkg.IfNotEmptyFunc(req.BusinessId2, func() { q.Where("business_id2=?", req.BusinessId2) })
	pkg.IfNotEmptyFunc(req.BusinessId3, func() { q.Where("business_id3=?", req.BusinessId3) })
	pkg.IfNotEmptyFunc(req.TradeId, func() { q.Where("trade_id=?", req.TradeId) })
	pkg.IfGt0Func(req.PayChannelId, func() { q.Where("pay_channel_id=?", req.PayChannelId) })
	pkg.IfNotEmptyFunc(req.PayChannelType, func() { q.Where("pay_channel_type=?", req.PayChannelType) })
	pkg.IfNotEmptyFunc(req.AmountYuan, func() { q.Where("amount_yuan=?", req.AmountYuan) })
	pkg.IfNotEmptyFunc(req.AmountFen, func() { q.Where("amount_fen=?", req.AmountFen) })
	pkg.IfNotEmptyFunc(req.Message, func() { q.Where("message=?", req.Message) })
	pkg.IfGt0Func(req.State, func() { q.Where("state=?", req.State) })
	pkg.IfGt0Func(req.Upgrade, func() { q.Where("upgrade=?", req.Upgrade) })
	pkg.IfNotEmptyFunc(req.Other1, func() { q.Where("other1 like concat('%',?,'%')", req.Other1) })
	pkg.IfNotEmptyFunc(req.Other2, func() { q.Where("other2 like concat('%',?,'%')", req.Other2) })
	pkg.IfNotEmptyFunc(req.Other3, func() { q.Where("other3 like concat('%',?,'%')", req.Other3) })
	pkg.IfNotEmptyFunc(req.Remark1, func() { q.Where("remark1 like concat('%',?,'%')", req.Remark1) })
	pkg.IfNotEmptyFunc(req.Remark2, func() { q.Where("remark2 like concat('%',?,'%')", req.Remark2) })
	pkg.IfNotEmptyFunc(req.Remark3, func() { q.Where("remark3 like concat('%',?,'%')", req.Remark3) })
	pkg.IfNotEmptyFunc(req.CreateTime1, func() { q.Where("create_time>=concat(?,' 00:00:00')", req.CreateTime1) })
	pkg.IfNotEmptyFunc(req.CreateTime2, func() { q.Where("create_time<=concat(?,' 23:59:59')", req.CreateTime2) })
	pkg.IfNotEmptyFunc(req.PayTime1, func() { q.Where("pay_time>=concat(?,' 00:00:00')", req.PayTime1) })
	pkg.IfNotEmptyFunc(req.PayTime2, func() { q.Where("pay_time<=concat(?,' 23:59:59')", req.PayTime2) })
	pkg.IfNotEmptyFunc(req.UpdateTime1, func() { q.Where("update_time>=concat(?,' 00:00:00')", req.UpdateTime1) })
	pkg.IfNotEmptyFunc(req.UpdateTime2, func() { q.Where("update_time<=concat(?,' 23:59:59')", req.UpdateTime2) })
	pkg.IfNotEmptyFunc(req.CancelTime1, func() { q.Where("cancel_time>=concat(?,' 00:00:00')", req.CancelTime1) })
	pkg.IfNotEmptyFunc(req.CancelTime1, func() { q.Where("cancel_time<=concat(?,' 23:59:59')", req.CancelTime2) })
	if fn := req.ExtraCallback; fn != nil {
		fn(q)
	}
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
	q := db.GetDb().Model(new(models.Order)).Where("business_id1=?", req.BusinessId1)
	pkg.IfNotEmptyFunc(req.BusinessId2, func() { q.Where("business_id2=?", req.BusinessId2) })
	pkg.IfNotEmptyFunc(req.BusinessId3, func() { q.Where("business_id3=?", req.BusinessId3) })
	if err = q.Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付订单业务id1[%s]业务id2[%s]业务id3[%s]不存在", req.BusinessId1, req.BusinessId2, req.BusinessId3))
		return
	}
	resp.Order = list[0]
	return
}

func (s *service) GetIdAndBusinessId(req GetIdAndBusinessIdReq) (resp GetIdAndBusinessIdResp, err error) {
	var list []models.Order
	q := db.GetDb().Model(new(models.Order)).Where("id=? and business_id1=?", req.Id, req.BusinessId1)
	pkg.IfNotEmptyFunc(req.BusinessId2, func() { q.Where("business_id2=?", req.BusinessId2) })
	pkg.IfNotEmptyFunc(req.BusinessId3, func() { q.Where("business_id3=?", req.BusinessId3) })
	if err = q.Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付订单id[%s]业务id1[%s]业务id2[%s]业务id3[%s]不存在", req.Id, req.BusinessId1, req.BusinessId2, req.BusinessId3))
		return
	}
	resp.Order = list[0]
	return
}

func (s *service) Add(req AddReq) (err error) { return db.GetDb().Create(req.transform()).Error }

func (s *service) AddReturn(req AddReq) (order *models.Order, err error) {
	odr := req.transform()
	if err = db.GetDb().Create(odr).Error; err != nil {
		return
	}
	order = odr
	return
}

func (s *service) Update(req UpdateReq) (err error) {
	return db.GetDb().Model(&models.Order{Id: req.Id}).Updates(req.transform()).Error
}

func (s *service) Del(req DelReq) (err error) {
	return db.GetDb().Delete(&models.Order{Id: req.Id}).Error
}

func (s *service) Paid(req PaidReq, callback ...CallbackFunc) (err error) {
	if err = db.GetDb().Model(&models.Order{Id: req.Id}).Updates(models.Order{TradeId: req.TradeId, Message: req.Message, State: models.OrderStatePaid, PayTime: pkg.TimeNowStr()}).Error; err != nil {
		return
	}
	s.callback(req.Id, callback...)
	return
}

func (s *service) Cancel(req CancelReq, callback ...CallbackFunc) (err error) {
	if req.CancelTime == "" {
		req.CancelTime = pkg.TimeNowStr()
	}
	if err = db.GetDb().Model(&models.Order{Id: req.Id}).Updates(req.transform()).Error; err != nil {
		return
	}
	s.callback(req.Id, callback...)
	return
}

func (s *service) GetState(req GetStateReq) (resp GetStateResp, err error) {
	err = db.GetDb().Model(&models.Order{Id: req.Id}).Select("state", "message").Scan(&resp).Error
	return
}

func (s *service) callback(id string, callback ...CallbackFunc) {
	if callback != nil && len(callback) > 0 {
		if fn := callback[0]; fn != nil {
			go func() { resp, _ := s.Get(GetReq{Id: id}); fn(&resp.Order) }()
		}
	}
}
