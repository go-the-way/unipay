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

package channelparam

import (
	"errors"
	"fmt"

	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/models"
)

func Impl() *impl { return &impl{} }

type impl struct{}

func (s *impl) Get(req GetReq) (resp GetResp, err error) {
	var list []models.ChannelParam
	if err = db.GetDb().Model(new(models.ChannelParam)).Where("id=?", req.Id).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付通道参数[%d]不存在", req.Id))
		return
	}
	resp.ChannelParam = list[0]
	return
}

func (s *impl) GetChannelId(req GetChannelIdReq) (resp GetChannelIdResp, err error) {
	err = db.GetDb().Model(new(models.ChannelParam)).Where("channel_id=?", req.ChannelId).Find(&resp.List).Error
	return
}

func (s *impl) GetName(req GetNameReq) (resp GetNameResp, err error) {
	var list []models.ChannelParam
	if err = db.GetDb().Model(new(models.ChannelParam)).Where("channel_id=? and name=?", req.ChannelId, req.Name).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("支付通道参数支付通道[%d]名称[%s]不存在", req.ChannelId, req.Name))
		return
	}
	resp.ChannelParam = list[0]
	return
}

func (s *impl) Add(req AddReq) (err error) {
	return db.GetDb().Create(req.Transform()).Error
}

func (s *impl) Update(req UpdateReq) (err error) {
	return db.GetDb().Model(&models.ChannelParam{Id: req.Id}).Updates(req.Transform()).Error
}

func (s *impl) Del(req DelReq) (err error) {
	return db.GetDb().Delete(&models.ChannelParam{Id: req.Id}).Error
}
