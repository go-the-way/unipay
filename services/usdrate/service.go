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

package usdrate

import (
	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/models"
	"sync"
)

type service struct{ *sync.Once }

func (s *service) syncRate() { s.Do(s.syncRate0) }

func (s *service) syncRate0() {
	// TODO: sync usd rate from https://www.waihui999.com/usdcny
}

func (s *service) Get() (resp GetResp, err error) {
	err = db.GetDb().Model(new(models.UsdRate)).Where("id=1").Select("rate").Scan(&resp.Rate).Error
	return
}

func (s *service) Update(req UpdateReq) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.UsdRate)).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		if err = db.GetDb().Updates(req.Transform()).Error; err != nil {
			return
		}
	} else {
		if err = db.GetDb().Create(req.Transform()).Error; err != nil {
			return
		}
	}
	if fn := req.Callback; fn != nil {
		return
	}
	return
}
