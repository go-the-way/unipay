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

package log

import (
	"github.com/rwscode/unipay/services/base"
)

type (
	GetPageReq struct {
		base.PageReq

		OrderBy string `form:"order_by"` // 排序

		Id          uint   `form:"id"`           // id
		Text        string `form:"text"`         // 日志内容
		CreateTime1 string `form:"create_time1"` // 创建时间
		CreateTime2 string `form:"create_time2"` // 创建时间
	}
)
