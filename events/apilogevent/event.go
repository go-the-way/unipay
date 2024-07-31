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

package apilogevent

import (
	"github.com/go-the-way/events"
	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/events/logevent"
	"github.com/go-the-way/unipay/models"
)

func Save(log *models.ApiLog) { event.Fire(log) }

type evt struct{}

var event = events.NewHandler[evt, *models.ApiLog]()

func init() { event.Bind(bind) }

func bind(log *models.ApiLog) {
	if err := db.GetDb().Model(new(models.ApiLog)).Create(log).Error; err != nil {
		logevent.Save(models.NewLog("保存api log错误：" + err.Error()))
	}
}
