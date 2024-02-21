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

type (
	ReqPayReq  = pay.Req
	ReqPayResp = pay.Resp
	NotifyReq  = pay.NotifyReq
)

type (
	ChannelGetPageReq    = channel.GetPageReq
	ChannelIdReq         = channel.IdReq
	ChannelGetReq        = channel.GetReq
	ChannelAddReq        = channel.AddReq
	ChannelUpdateReq     = channel.UpdateReq
	ChannelDelReq        = channel.DelReq
	ChannelEnableReq     = channel.EnableReq
	ChannelDisableReq    = channel.DisableReq
	ChannelGetMatchesReq = channel.GetMatchesReq

	ChannelGetPageResp    = channel.GetPageResp
	ChannelGetResp        = channel.GetResp
	ChannelGetMatchesResp = channel.GetMatchesResp
)

type (
	ChannelParamGetReq          = channelparam.GetReq
	ChannelParamGetChannelIdReq = channelparam.GetChannelIdReq
	ChannelParamGetNameReq      = channelparam.GetNameReq
	ChannelParamAddReq          = channelparam.AddReq
	ChannelParamUpdateReq       = channelparam.UpdateReq
	ChannelParamDelReq          = channelparam.DelReq

	ChannelParamGetResp          = channelparam.GetResp
	ChannelParamGetChannelIdResp = channelparam.GetChannelIdResp
	ChannelParamGetNameResp      = channelparam.GetNameResp
)

type (
	OrderGetPageReq            = order.GetPageReq
	OrderIdReq                 = order.IdReq
	OrderGetReq                = order.GetReq
	OrderGetBusinessIdReq      = order.GetBusinessIdReq
	OrderGetIdAndBusinessIdReq = order.GetIdAndBusinessIdReq
	OrderAddReq                = order.AddReq
	OrderUpdateReq             = order.UpdateReq
	OrderDelReq                = order.DelReq
	OrderPaidReq               = order.PaidReq
	OrderCancelReq             = order.CancelReq
	OrderGetStateReq           = order.GetStateReq

	OrderGetPageResp            = order.GetPageResp
	OrderGetResp                = order.GetResp
	OrderGetBusinessIdResp      = order.GetBusinessIdResp
	OrderGetIdAndBusinessIdResp = order.GetIdAndBusinessIdResp
	OrderGetStateResp           = order.GetStateResp
)
