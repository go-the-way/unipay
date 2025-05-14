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

package trongridevent

import (
	"fmt"
	"github.com/go-the-way/events"
	"github.com/go-the-way/unipay/events/logevent"
	"github.com/go-the-way/unipay/models"
)

func Run(arg *Arg) { event.Fire(arg) }

type (
	evt struct{}
	Arg struct {
		Order *models.Order
		Ac    models.ApiConfig
	}
)

var (
	event = events.NewHandler[evt, *Arg]()
	ch    = make(chan *Arg, 1)
)

func init() {
	go task()
	event.Bind(bind)
}

func bind(arg *Arg) { go func() { ch <- arg }() }

func task() {
	for {
		select {
		case arg := <-ch:
			// trongrid只支持trc20，如果不支持记录一下然后return
			ord := arg.Order
			if ord.PayChannelType == models.OrderTypeErc20 {
				logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]查询交易记录退出，trongrid只支持trc20", ord.Id, ord.PayChannelType)))
				return
			}
			backupApiKey := arg.Ac.BackupVar1
			if backupApiKey == "" {
				logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]查询交易记录退出，apikey为空", ord.Id, ord.PayChannelType)))
				continue
			}
			km := map[string]string{"trc20": backupApiKey}
			go startReq(ord, km[ord.PayChannelType])
		}
	}
}
