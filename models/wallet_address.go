// Copyright 2024 unipay Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES 2OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

type (
	WalletAddress       = UnipayWalletAddress
	UnipayWalletAddress struct {
		Id         uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                           // id
		Address    string `gorm:"column:address;type:varchar(200);not null;default:'';comment:钱包地址" json:"address"`                 // 钱包地址
		protocol   string `gorm:"column:protocol;type:varchar(50);not null;default:'trc20';comment:协议 trc20/erc20" json:"protocol"` // 协议 trc20/erc20
		State      byte   `gorm:"column:state;type:tinyint;not null;default:1;comment:状态：1启用2禁用;index" json:"state"`                // 状态：1启用2禁用
		Remark     string `gorm:"column:remark;type:varchar(500);not null;default:'';comment:备注" json:"remark"`                     // 备注
		CreateTime string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`          // 创建时间
		UpdateTime string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"`          // 修改时间
	}
)
