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

	"github.com/go-the-way/unipay/deps/db"
	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/models"
	"gorm.io/gorm"
)

type service struct{}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
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

func (s *service) GetChannelId(req GetChannelIdReq) (resp GetChannelIdResp, err error) {
	err = db.GetDb().Model(new(models.ChannelParam)).Where("channel_id=?", req.ChannelId).Find(&resp.List).Error
	return
}

func (s *service) GetName(req GetNameReq) (resp GetNameResp, err error) {
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

func (s *service) Add(req AddReq) (err error) {
	tx := db.GetDb().Begin()
	if err = tx.Create(req.transform()).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = s.disableChannel(tx, req.ChannelId, 0); err != nil {
		_ = tx.Rollback().Error
		return
	}
	_ = tx.Commit().Error
	return
}

func (s *service) Update(req UpdateReq) (err error) {
	tx := db.GetDb().Begin()
	if err = tx.Model(&models.ChannelParam{Id: req.Id}).Updates(req.transform()).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = s.disableChannel(tx, 0, req.Id); err != nil {
		_ = tx.Rollback().Error
		return
	}
	_ = tx.Commit().Error
	return
}

func (s *service) Del(req DelReq) (err error) {
	tx := db.GetDb().Begin()
	if err = tx.Delete(&models.ChannelParam{Id: req.Id}).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = s.disableChannel(tx, 0, req.Id); err != nil {
		_ = tx.Rollback().Error
		return
	}
	_ = tx.Commit().Error
	return
}

func (s *service) disableChannel(tx *gorm.DB, channelId, channelParamId uint) (err error) {
	if channelId == 0 {
		if err = db.GetDb().Model(new(models.ChannelParam)).Select("channel_id").Where("id=?", channelParamId).Scan(&channelId).Error; err != nil {
			return
		}
	}
	if channelId <= 0 {
		return
	}
	return tx.Model(&models.Channel{Id: channelId}).Updates(&models.Channel{State: models.ChannelStateDisable, UpdateTime: pkg.TimeNowStr()}).Error
}
