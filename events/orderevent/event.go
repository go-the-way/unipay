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
	"time"

	"github.com/go-the-way/events"
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/lock"
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/events/logevent"
	"github.com/rwscode/unipay/models"
	"github.com/rwscode/unipay/services/order"
)

func Paid(order *models.Order)    { paid.Fire(order) }
func Expired(order *models.Order) { expired.Fire(order) }

type (
	evt         struct{}
	PaidHandler func(order *models.Order)
)

var (
	paid        = events.NewHandler[evt, *models.Order]()
	expired     = events.NewHandler[evt, *models.Order]()
	paidHandler PaidHandler
)

func SetPaidHandler(handler PaidHandler) { paidHandler = handler }

func init() {
	paid.Bind(bindPaid)
	paid.Bind(bindDeleteLock)
	expired.Bind(bindExpired)
	expired.Bind(bindDeleteLock)
	// E20订单全部失效
	e20OrderCancelled()
}

func bindPaid(o *models.Order) {
	if paidHandler != nil {
		paidHandler(o)
	}
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

func e20OrderCancelled() {
	err := db.GetDb().Model(new(models.Order)).Where("pay_channel_type in ('erc20','trc20') and state=1").Updates(map[string]any{
		"state":       models.OrderStateCancelled,
		"message":     "服务重载强制取消",
		"cancel_time": pkg.FormatTime(time.Now()),
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}
