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

package orderevent

import (
	"fmt"
	"github.com/go-the-way/events"
	"github.com/rwscode/unipay/deps/lock"
	"github.com/rwscode/unipay/events/logevent"
	"github.com/rwscode/unipay/models"
	"github.com/rwscode/unipay/services/order"
)

func Paid(order *models.Order)    { paid.Fire(order) }
func Expired(order *models.Order) { expired.Fire(order) }

type evt struct{}

var (
	paid    = events.NewHandler[evt, *models.Order]()
	expired = events.NewHandler[evt, *models.Order]()
)

func init() {
	paid.Bind(bindPaid)
	paid.Bind(bindDeleteLock)
	expired.Bind(bindExpired)
	expired.Bind(bindDeleteLock)
}

func bindPaid(o *models.Order) {
	if err := order.Service.Paid(order.PaidReq{
		IdReq:   order.IdReq{Id: o.Id},
		TradeId: o.TradeId,
	}); err != nil {
		logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]保存错误：%s", o.Id, o.PayChannelType, err.Error())))
	}
}

func bindExpired(o *models.Order) {
	if err := order.Service.Cancel(order.CancelReq{
		IdReq:      order.IdReq{Id: o.Id},
		Message:    o.Message,
		CancelTime: o.CancelTime,
	}); err != nil {
		logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]保存错误：%s", o.Id, o.PayChannelType, err.Error())))
	}
}

func bindDeleteLock(o *models.Order) { lock.DeleteWithLock(o.LockKey()) }
