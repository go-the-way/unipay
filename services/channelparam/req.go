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

type (
	IdReq struct {
		Id uint `validate:"min(1,支付通道参数id不能为空)" form:"id" json:"id"`
	}
	GetReq          IdReq
	GetChannelIdReq struct {
		ChannelId uint `validate:"min(1,支付通道id不能为空)" form:"channel_id"`
	}
	GetNameReq struct {
		ChannelId uint   `validate:"min(1,支付通道id不能为空)" form:"channel_id"`
		Name      string `validate:"minlength(1,参数名称不能为空)" form:"name"`
	}
	AddReq struct {
		ChannelId uint   `validate:"min(1,支付通道id不能为空)" json:"channel_id"`                             // 支付通道Id
		Name      string `validate:"minlength(1,参数名称不能为空) maxlength(200,参数名称长度不能超过200)" json:"name"`  // 参数名称
		Value     string `validate:"minlength(1,参数值不能为空) maxlength(1000,参数值长度不能超过1000)" json:"value"` // 参数值
		Remark    string `validate:"maxlength(500,备注长度不能超过500)" json:"remark"`                        // 备注
		Pass      byte   `validate:"enum(1|2,是否传递不合法)" json:"pass"`                                   // 1传递2不传递
	}
	UpdateReq struct {
		IdReq  `validate:"valid(T)"`
		AddReq `validate:"valid(T)"`
	}
	DelReq IdReq
)
