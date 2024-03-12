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

package oklinkevent

import (
	"fmt"
	"github.com/go-the-way/events"
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/events/logevent"
	"github.com/rwscode/unipay/models"
)

func Run(order *models.Order) { event.Fire(order) }

type evt struct{}

var (
	event = events.NewHandler[evt, *models.Order]()
	ch    = make(chan *models.Order, 1)
)

func init() {
	go task()
	event.Bind(bind)
}

func bind(order *models.Order) { go func() { ch <- order }() }

func task() {
	for {
		select {
		case order := <-ch:
			conf := getApiConfig()
			if conf.Trc20Apikey == "" {
				logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]查询交易记录退出，ok_link_trc20_apikey为空", order.Id, order.PayChannelType)))
				continue
			}
			if conf.Erc20Apikey == "" {
				logevent.Save(models.NewLog(fmt.Sprintf("订单号[%s]类型[%s]查询交易记录退出，ok_link_erc20_apikey为空", order.Id, order.PayChannelType)))
				continue
			}
			km := map[string]string{"erc20": conf.Erc20Apikey, "trc20": conf.Trc20Apikey}
			startReq(order, km[order.PayChannelType], tm[order.PayChannelType])
		}
	}
}

type apiConfig struct {
	Trc20Apikey string `gorm:"column:ok_link_trc20_apikey"`
	Erc20Apikey string `gorm:"column:ok_link_erc20_apikey"`
}

func getApiConfig() (conf apiConfig) {
	_ = db.GetDb().Model(new(models.ApiConfig)).Where("id=1").Select("ok_link_trc20_apikey", "ok_link_erc20_apikey").Scan(&conf).Error
	return
}
