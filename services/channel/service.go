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

package channel

import (
	"errors"
	"fmt"

	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/models"
	"github.com/rwscode/unipay/deps/pkg"
)

type impl struct{}

func Impl() SVC { return &impl{} }

func (s *impl) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDb().Model(new(models.Channel))
	pkg.IfGt0Func(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfNotEmptyFunc(req.Name, func() { q.Where("name like concat('%',?,'%')", req.Name) })
	pkg.IfNotEmptyFunc(req.AdminUrl, func() { q.Where("admin_url like concat('%',?,'%')", req.AdminUrl) })
	pkg.IfNotEmptyFunc(req.AdminUser, func() { q.Where("admin_user like concat('%',?,'%')", req.AdminUser) })
	pkg.IfNotEmptyFunc(req.AdminPasswd, func() { q.Where("admin_passwd like concat('%',?,'%')", req.AdminPasswd) })
	pkg.IfGt0Func(req.AmountType, func() { q.Where("amount_type=?", req.AmountType) })
	pkg.IfNotEmptyFunc(req.AmountValidateCond, func() { q.Where("amount_validate_cond like concat('%',?,'%')", req.AmountValidateCond) })
	pkg.IfNotEmptyFunc(req.ReqUrl, func() { q.Where("req_url like concat('%',?,'%')", req.ReqUrl) })
	pkg.IfNotEmptyFunc(req.ReqMethod, func() { q.Where("req_method = ?", req.ReqMethod) })
	pkg.IfNotEmptyFunc(req.ReqContentType, func() { q.Where("req_content_type = ?", req.ReqContentType) })
	pkg.IfNotEmptyFunc(req.NotifyPayContentType, func() { q.Where("notify_pay_content_type = ?", req.NotifyPayContentType) })
	pkg.IfNotEmptyFunc(req.NotifyPayReturnContent, func() { q.Where("notify_pay_return_content like concat('%',?,'%')", req.NotifyPayReturnContent) })
	pkg.IfNotEmptyFunc(req.NotifyPayReturnContentType, func() { q.Where("notify_pay_return_content_type = ?", req.NotifyPayReturnContentType) })
	pkg.IfGt0Func(req.State, func() { q.Where("state=?", req.State) })
	pkg.IfGt0Func(req.Sort, func() { q.Where("sort=?", req.Sort) })
	pkg.IfGt0Func(req.Sort1, func() { q.Where("sort>=?", req.Sort1) })
	pkg.IfGt0Func(req.Sort2, func() { q.Where("sort<=?", req.Sort2) })
	pkg.IfNotEmptyFunc(req.Remark, func() { q.Where("remark like concat('%',?,'%')", req.Remark) })
	pkg.IfNotEmptyFunc(req.CreateTime1, func() { q.Where("create_time>=concat(?,' 00:00:00')", req.CreateTime1) })
	pkg.IfNotEmptyFunc(req.CreateTime2, func() { q.Where("create_time<=concat(?,' 23:59:59')", req.CreateTime2) })
	pkg.IfNotEmptyFunc(req.UpdateTime1, func() { q.Where("update_time>=concat(?,' 00:00:00')", req.UpdateTime1) })
	pkg.IfNotEmptyFunc(req.UpdateTime2, func() { q.Where("update_time<=concat(?,' 23:59:59')", req.UpdateTime2) })
	if req.OrderBy != "" {
		q.Order(req.OrderBy)
	}
	err = db.GetPagination()(q, req.Page, req.Limit, &resp.Total, &resp.List)
	return
}

func (s *impl) Get(req GetReq) (resp GetResp, err error) {
	var list []models.Channel
	if err = db.GetDb().Model(new(models.Channel)).Where("id=?", req.Id).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付通道[%d]不存在", req.Id))
		return
	}
	resp.Channel = list[0]
	return
}

func (s *impl) Add(req AddReq) (err error) {
	return db.GetDb().Create(req.Transform()).Error
}

func (s *impl) Update(req UpdateReq) (err error) {
	return db.GetDb().Model(&models.Channel{Id: req.Id}).Omit("create_time").Updates(req.Transform()).Error
}

func (s *impl) Del(req DelReq) (err error) {
	tx := db.GetDb().Begin()
	if err = tx.Delete(&models.Channel{Id: req.Id}).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = tx.Where("channel_id=?", req.Id).Delete(new(models.ChannelParam)).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	_ = tx.Commit().Error
	return
}

func (s *impl) Enable(req EnableReq) (err error) {
	return s.updateState(req.Id, models.ChannelStateEnable)
}

func (s *impl) Disable(req DisableReq) (err error) {
	return s.updateState(req.Id, models.ChannelStateDisable)
}

func (s *impl) updateState(id uint, state byte) (err error) {
	return db.GetDb().Model(&models.Channel{Id: id}).Updates(models.Channel{State: state, UpdateTime: pkg.TimeNowStr()}).Error
}

func (s *impl) GetMatches(req GetMatchesReq) (resp GetMatchesResp, err error) {
	q := db.GetDb().Model(new(models.Channel))
	if req.Order != "" {
		q.Order(req.Order)
	}
	var list []models.Channel
	if err = q.Find(&list).Error; err != nil {
		return
	}
	var result []models.Channel
	if len(list) > 0 && req.Amount > 0 {
		for _, c := range list {
			cond := c.AmountValidateCond
			if ok := cond == "" || pkg.ValidAmount(req.Amount, c.AmountType == models.ChannelAmountTypeYuan, cond); ok {
				result = append(result, c)
			}
			if req.Limit > 0 && uint(len(result)) >= req.Limit {
				break
			}
		}
	} else {
		// FIXME: find all limit ?
		result = append(result, list...)
	}
	resp.List = result
	return
}
