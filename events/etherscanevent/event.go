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

package etherscanevent

import (
	"github.com/go-the-way/events"
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/models"
)

func Run(order models.Order, address Address) { event.Fire(order, address) }

type (
	evt     struct{}
	Address = string
	arg     struct {
		models.Order
		Address
	}
)

var (
	event = events.NewBIHandler[evt, models.Order, Address]()
	ch    = make(chan arg)
)

func init() {
	event.Bind(bind)
	go task()
}

func bind(order models.Order, address Address) { go func() { ch <- arg{order, address} }() }

func task() {
	for {
		select {
		case a := <-ch:
			apikey := getApikey()
			if apikey == "" {
				// TODO: save api log
				// apilogevent.Send(models.ApiLog{})
			} else {
				startReq(a.Order, a.Address, apikey)
			}
		}
	}
}

func getApikey() (apikey string) {
	_ = db.GetDb().Model(new(models.ApiConfig)).Where("id=1").Scan(&apikey).Error
	return
}
