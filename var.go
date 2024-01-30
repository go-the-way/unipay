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

package unipay

import "github.com/rwscode/unipay/services"

var (
	PayService          = services.PayService
	ChannelService      = services.ChannelService
	ChannelParamService = services.ChannelParamService
	OrderService        = services.OrderService
)

var (
	ReqPay    = PayService.ReqPay
	NotifyPay = PayService.NotifyPay
)

var (
	ChannelGetPage    = ChannelService.GetPage
	ChannelGet        = ChannelService.Get
	ChannelAdd        = ChannelService.Add
	ChannelUpdate     = ChannelService.Update
	ChannelDel        = ChannelService.Del
	ChannelEnable     = ChannelService.Enable
	ChannelDisable    = ChannelService.Disable
	ChannelGetMatches = ChannelService.GetMatches
)

var (
	ChannelParamGet          = ChannelParamService.Get
	ChannelParamGetChannelId = ChannelParamService.GetChannelId
	ChannelParamGetName      = ChannelParamService.GetName
	ChannelParamAdd          = ChannelParamService.Add
	ChannelParamUpdate       = ChannelParamService.Update
	ChannelParamDel          = ChannelParamService.Del
)

var (
	OrderGetPage       = OrderService.GetPage
	OrderGet           = OrderService.Get
	OrderGetBusinessId = OrderService.GetBusinessId
	OrderAdd           = OrderService.Add
	OrderUpdate        = OrderService.Update
	OrderDel           = OrderService.Del
	OrderPaySuccess    = OrderService.PaySuccess
	OrderPayFailure    = OrderService.PayFailure
	OrderGetPayState   = OrderService.GetPayState
)
