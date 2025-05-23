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

package walletaddress

import (
	"github.com/go-the-way/unipay/services/base"
)

type (
	GetPageReq struct {
		base.PageReq

		OrderBy string `form:"order_by"` // 排序

		Id          uint   `form:"id"`       // id
		Address     string `form:"address"`  // 钱包地址
		Protocol    string `form:"protocol"` // 协议 trc20/erc20
		State       byte   `form:"state"`    // 状态：1启用2禁用
		BusinessId1 string `form:"business_id1"`
		BusinessId2 string `form:"business_id2"`
		BusinessId3 string `form:"business_id3"`
		Remark      string `form:"remark"`       // 备注
		CreateTime1 string `form:"create_time1"` // 创建时间
		CreateTime2 string `form:"create_time2"` // 创建时间
		UpdateTime1 string `form:"update_time1"` // 修改时间
		UpdateTime2 string `form:"update_time2"` // 修改时间
	}
	IdReq struct {
		Id uint `validate:"min(1,id不能为空)" json:"id"`
	}
	GetReq IdReq
	AddReq struct {
		Address     string `validate:"minlength(1,钱包地址不能为空) maxlength(100,钱包地址长度不能超过100)" json:"address"` // 钱包地址
		Protocol    string `validate:"enum(erc20|trc20,协议不合法)" json:"protocol"`                           // 协议
		Remark      string `validate:"maxlength(500,备注长度不能超过500)" json:"remark"`                          // 备注
		BusinessId1 string `json:"business_id1"`
		BusinessId2 string `json:"business_id2"`
		BusinessId3 string `json:"business_id3"`
	}
	UpdateReq struct {
		IdReq  `validate:"valid(T)"`
		AddReq `validate:"valid(T)"`
	}
	DelReq     IdReq
	EnableReq  IdReq
	DisableReq IdReq
)
