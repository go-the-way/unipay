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

package order

import "github.com/go-the-way/unipay/models"

var (
	Service SVC = &service{}

	GetPage            = Service.GetPage
	Get                = Service.Get
	GetBusinessId      = Service.GetBusinessId
	GetIdAndBusinessId = Service.GetIdAndBusinessId
	Add                = Service.Add
	AddReturn          = Service.AddReturn
	Update             = Service.Update
	Del                = Service.Del
	Paid               = Service.Paid
	Cancel             = Service.Cancel
	GetState           = Service.GetState
)

type SVC interface {
	GetPage(req GetPageReq) (resp GetPageResp, err error)
	Get(req GetReq) (resp GetResp, err error)
	GetBusinessId(req GetBusinessIdReq) (resp GetBusinessIdResp, err error)
	GetIdAndBusinessId(req GetIdAndBusinessIdReq) (resp GetIdAndBusinessIdResp, err error)
	Add(req AddReq) (err error)
	AddReturn(req AddReq) (order *models.Order, err error)
	Update(req UpdateReq) (err error)
	Del(req DelReq) (err error)
	Paid(req PaidReq, callback ...CallbackFunc) (err error)
	Cancel(req CancelReq, callback ...CallbackFunc) (err error)
	GetState(req GetStateReq) (resp GetStateResp, err error)
}
