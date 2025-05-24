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
	"sync"
	"time"

	"github.com/go-the-way/events"
	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/deps/lock"
	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/events/logevent"
	"github.com/go-the-way/unipay/models"
	"github.com/go-the-way/unipay/services/order"
)

func Paid(order *models.Order)    { paid.Fire(order) }
func Expired(order *models.Order) { expired.Fire(order) }

type (
	evt            struct{}
	PaidHandler    func(order *models.Order)
	ExpiredHandler func(order *models.Order)
)

var (
	paid    = events.NewHandler[evt, *models.Order]()
	expired = events.NewHandler[evt, *models.Order]()

	paidHandler    PaidHandler
	expiredHandler ExpiredHandler

	orderValidMinute = 10 // 订单有效期，默认10分钟
	orderTaskDur     = time.Minute
	orderTaskMu      = &sync.Mutex{}
)

func SetPaidHandler(handler PaidHandler)       { paidHandler = handler }
func SetExpiredHandler(handler ExpiredHandler) { expiredHandler = handler }

func init() {
	bindAll()
	tasks()
}

func bindAll() {
	paid.Bind(bindPaid)
	paid.Bind(bindDeleteLock)
	expired.Bind(bindExpired)
	expired.Bind(bindDeleteLock)
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
	if expiredHandler != nil {
		expiredHandler(o)
	}
	if err := order.Service.Cancel(order.CancelReq{
		IdReq:      order.IdReq{Id: o.Id},
		Message:    o.Message,
		CancelTime: pkg.TimeNowStr(),
	}); err != nil {
		logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]保存错误：%s", o.Id, o.PayChannelType, err.Error())))
	}
}

func bindDeleteLock(o *models.Order) {
	if o.PayChannelType == models.OrderTypeErc20 || o.PayChannelType == models.OrderTypeTrc20 {
		// only release lock erc20 & trc20
		lock.DeleteWithLock(o.LockKey())
	}
}

func tasks() {
	time.AfterFunc(time.Second*5, func() { e20Cancelled() }) // e20订单全部失效
	go orderTask()
}

func e20Cancelled() {
	var m = map[string]any{
		"message":     "服务重载强制取消",
		"state":       models.OrderStateCancelled,
		"cancel_time": pkg.FormatTime(time.Now()),
	}
	err := db.GetDb().Model(new(models.Order)).Where("pay_channel_type in ('erc20','trc20') and state=1").Updates(m).Error
	if err != nil {
		fmt.Println(err)
	}
}

func orderTask() {
	ticker := time.NewTicker(orderTaskDur)
	defer ticker.Stop()
	for range ticker.C {
		orderCancelledTask()
	}
}

func orderCancelledTask() {
	if !orderTaskMu.TryLock() {
		return
	}
	defer orderTaskMu.Unlock()

	var cols = "id"
	if expiredHandler != nil {
		cols = "*"
	}

	var orders []*models.Order
	if err := db.GetDb().Model(new(models.Order)).Where("state = ? and pay_channel_type = ? and adddate(create_time, interval ? minute) < NOW()", models.OrderStateWaitPay, models.OrderTypeNormal, orderValidMinute).Select(cols).Find(&orders).Error; err != nil {
		return
	}
	for _, order := range orders {
		expired.Fire(order)
	}
}
