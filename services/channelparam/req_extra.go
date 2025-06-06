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
	"github.com/go-the-way/unipay/models"
	"github.com/go-the-way/unipay/services/base"
)

func (r *AddReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckChannelExist(r.ChannelId) },
		func() (err error) { return base.CheckChannelParamNameExist(r.ChannelId, 0, r.Name) },
	)
}

func (r *UpdateReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckChannelParamExist(r.Id) },
		func() (err error) { return base.CheckChannelExist(r.ChannelId) },
		func() (err error) { return base.CheckChannelParamNameExist(r.ChannelId, r.Id, r.Name) },
	)
}

func (r *DelReq) Check() (err error) { return base.CheckChannelParamExist(r.Id) }

func (r *AddReq) transform() *models.ChannelParam {
	return &models.ChannelParam{
		ChannelId: r.ChannelId,
		Name:      r.Name,
		Value:     r.Value,
		Remark:    r.Remark,
		Pass:      r.Pass,
	}
}

func (r *UpdateReq) transform() map[string]any {
	return map[string]any{
		"channel_id": r.ChannelId,
		"name":       r.Name,
		"value":      r.Value,
		"remark":     r.Remark,
		"pass":       r.Pass,
	}
}
