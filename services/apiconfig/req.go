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

import "github.com/go-the-way/unipay/models"

type (
	UpdateReq struct {
		Erc20Apikey       string `validate:"minlength(1,erc20_apikey不能为空) maxlength(100,erc20_apikey长度不能超过100)" json:"erc20_apikey"`                         // erc20_apikey
		OkLinkTrc20Apikey string `validate:"minlength(1,ok_link_trc20_apikey不能为空) maxlength(100,ok_link_trc20_apikey长度不能超过100)" json:"ok_link_trc20_apikey"` // ok_link_trc20_apikey
		OkLinkErc20Apikey string `validate:"minlength(1,ok_link_erc20_apikey不能为空) maxlength(100,ok_link_erc20_apikey长度不能超过100)" json:"ok_link_erc20_apikey"` // ok_link_erc20_apikey
		ValidPeriodMinute uint   `validate:"min(1,valid_period_minute不能为空)" json:"valid_period_minute"`                                                      // valid_period_minute

		Callback func(config models.ApiConfig) `json:"callback"`
	}
)
