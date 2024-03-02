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

package models

type (
	ApiConfig       = UnipayApiConfig
	UnipayApiConfig struct {
		Id                uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                     // id
		Erc20Apikey       string `gorm:"column:erc20_apikey;type:varchar(200);not null;default:'';comment:erc20_apikey" json:"erc20_apikey"`                         // erc20_apikey
		OkLinkTrc20Apikey string `gorm:"column:ok_link_trc20_apikey;type:varchar(200);not null;default:'';comment:ok_link_trc20_apikey" json:"ok_link_trc20_apikey"` // ok_link_trc20_apikey
		OkLinkErc20Apikey string `gorm:"column:ok_link_erc20_apikey;type:varchar(200);not null;default:'';comment:ok_link_erc20_apikey" json:"ok_link_erc20_apikey"` // ok_link_erc20_apikey
		ValidPeriodMinute uint   `gorm:"column:valid_period_minute;type:uint;not null;default:15;comment:订单有效期（分钟）" json:"valid_period_minute"`                      // 订单有效期（分钟）
	}
)
