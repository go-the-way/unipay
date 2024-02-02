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

import (
	"github.com/rwscode/unipay/services/channel"
	"github.com/rwscode/unipay/services/channelparam"
	"github.com/rwscode/unipay/services/order"
	"github.com/rwscode/unipay/services/pay"
)

var (
	PayService          = pay.Service
	ChannelService      = channel.Service
	ChannelParamService = channelparam.Service
	OrderService        = order.Service
)

var (
	ReqPay    = pay.ReqPay
	NotifyPay = pay.NotifyPay
)

var (
	ChannelGetPage    = channel.GetPage
	ChannelGet        = channel.Get
	ChannelAdd        = channel.Add
	ChannelUpdate     = channel.Update
	ChannelDel        = channel.Del
	ChannelEnable     = channel.Enable
	ChannelDisable    = channel.Disable
	ChannelGetMatches = channel.GetMatches
)

var (
	ChannelParamGet          = channelparam.Get
	ChannelParamGetChannelId = channelparam.GetChannelId
	ChannelParamGetName      = channelparam.GetName
	ChannelParamAdd          = channelparam.Add
	ChannelParamUpdate       = channelparam.Update
	ChannelParamDel          = channelparam.Del
)

var (
	OrderGetPage            = order.GetPage
	OrderGet                = order.Get
	OrderGetBusinessId      = order.GetBusinessId
	OrderGetIdAndBusinessId = order.GetIdAndBusinessId
	OrderAdd                = order.Add
	OrderUpdate             = order.Update
	OrderDel                = order.Del
	OrderPaySuccess         = order.PaySuccess
	OrderPayFailure         = order.PayFailure
	OrderGetPayState        = order.GetPayState
)
