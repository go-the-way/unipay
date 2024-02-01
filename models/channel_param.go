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
	ChannelParam       = UnipayChannelParam
	UnipayChannelParam struct {
		Id        uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                // id
		ChannelId uint   `gorm:"column:channel_id;type:uint;not null;default:0;comment:支付通道Id;index;uniqueIndex:idx_channel_id_name" json:"channel_id"` // 支付通道Id
		Name      string `gorm:"column:name;type:varchar(200);not null;default:'';comment:参数名称;uniqueIndex:idx_channel_id_name" json:"name"`            // 参数名称
		Value     string `gorm:"column:value;type:varchar(1000);not null;default:'';comment:参数值" json:"value"`                                          // 参数值
		Remark    string `gorm:"column:remark;type:varchar(500);not null;default:'';comment:备注" json:"remark"`                                          // 备注
		Pass      byte   `gorm:"column:pass;type:tinyint;not null;default:2;comment:1传递2不传递" json:"pass"`                                               // 1传递2不传递
	}
)

const (
	_ byte = iota
	// ChannelParamPassYes 传递参数
	ChannelParamPassYes
	// ChannelParamPassNo 不传递参数
	ChannelParamPassNo
)
