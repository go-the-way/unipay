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

package apiconfig

import (
	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/models"
)

type service struct{}

func (s *service) Get() (resp GetResp, err error) {
	err = db.GetDb().Model(new(models.ApiConfig)).Where("id=1").Scan(&resp).Error
	return
}

func (s *service) Update(req UpdateReq) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.ApiConfig)).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		if err = db.GetDb().Model(&models.ApiConfig{Id: 1}).Updates(req.transformUpdate()).Error; err != nil {
			return
		}
	} else {
		if err = db.GetDb().Create(req.transformCreate()).Error; err != nil {
			return
		}
	}
	if fn := req.Callback; fn != nil {
		return
	}
	return
}
