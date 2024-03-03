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

package walletaddress

import (
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/models"
)

type service struct{}

func (s *service) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDb().Model(new(models.WalletAddress))
	pkg.IfGt0Func(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfNotEmptyFunc(req.Address, func() { q.Where("address like concat('%',?,'%')", req.Address) })
	pkg.IfNotEmptyFunc(req.Protocol, func() { q.Where("protocol=?", req.Protocol) })
	pkg.IfGt0Func(req.State, func() { q.Where("state=?", req.State) })
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

func (s *service) Add(req AddReq) (err error) {
	return db.GetDb().Create(req.Transform()).Error
}

func (s *service) Update(req UpdateReq) (err error) {
	return db.GetDb().Model(&models.Channel{Id: req.Id}).Omit("create_time").Updates(req.Transform()).Error
}

func (s *service) Del(req DelReq) (err error) {
	err = db.GetDb().Delete(&models.WalletAddress{Id: req.Id}).Error
	return
}

func (s *service) buildPayMap() map[string]any {
	return map[string]any{
		"ChannelId":   "100",
		"Amount":      "100",
		"AmountYuan":  "1",
		"AmountFen":   "100",
		"Subject":     "subject",
		"ClientIp":    "127.0.0.1",
		"NotifyUrl":   "http://example.com",
		"BusinessId1": "",
		"BusinessId2": "",
		"BusinessId3": "",
		"Remark1":     "",
		"Remark2":     "",
		"Remark3":     "",
	}
}

func (s *service) Enable(req EnableReq) (err error) {
	return s.updateState(req.Id, models.WalletAddressStateEnable)
}

func (s *service) Disable(req DisableReq) (err error) {
	return s.updateState(req.Id, models.WalletAddressStateDisable)
}

func (s *service) updateState(id uint, state byte) (err error) {
	return db.GetDb().Model(&models.WalletAddress{Id: id}).Updates(models.WalletAddress{State: state, UpdateTime: pkg.TimeNowStr()}).Error
}
