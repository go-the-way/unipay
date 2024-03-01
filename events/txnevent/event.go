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

package txnevent

import (
	"github.com/go-the-way/events"
	"github.com/rwscode/unipay/models"
)

type event struct{}

// 虚拟支付订单创建时
var created = events.NewHandler[event, models.Order]()

func BindCreated(fn func(order models.Order)) { created.Bind(fn) }
func FireCreated(order models.Order)          { created.Fire(order) }
