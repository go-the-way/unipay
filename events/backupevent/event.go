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

package backupevent

import (
	"github.com/go-the-way/events"
	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/events/backupevent/trongridevent"
	"github.com/go-the-way/unipay/models"
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

const (
	backupPlanTrongrid = "trongrid"
)

func getApiConfig() (ac models.ApiConfig) {
	_ = db.GetDb().Model(new(models.ApiConfig)).Where("id=1").Scan(&ac).Error
	return
}

func task() {
	for {
		select {
		case order := <-ch:
			ac := getApiConfig()
			switch ac.BackupPlan {
			case backupPlanTrongrid:
				trongridevent.Run(&trongridevent.Arg{Order: order, Ac: ac})
			}
		}
	}
}
