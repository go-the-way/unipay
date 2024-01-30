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
	ChannelGetReq         = channel.GetReq
	ChannelDelReq         = channel.DelReq
	ChannelUpdateReq      = channel.UpdateReq
	ChannelAddReq         = channel.AddReq
	ChannelDisableReq     = channel.DisableReq
	ChannelEnableReq      = channel.EnableReq
	ChannelIdReq          = channel.IdReq
	ChannelGetPageReq     = channel.GetPageReq
	ChannelGetMatchesReq  = channel.GetMatchesReq
	ChannelGetPageResp    = channel.GetPageResp
	ChannelGetResp        = channel.GetResp
	ChannelGetMatchesResp = channel.GetMatchesResp
)

type (
	ChannelParamGetNameResp      = channelparam.GetNameResp
	ChannelParamGetResp          = channelparam.GetResp
	ChannelParamGetChannelIdResp = channelparam.GetChannelIdResp
	ChannelParamDelReq           = channelparam.DelReq
	ChannelParamUpdateReq        = channelparam.UpdateReq
	ChannelParamIdReq            = channelparam.IdReq
	ChannelParamAddReq           = channelparam.AddReq
	ChannelParamGetNameReq       = channelparam.GetNameReq
	ChannelParamGetChannelIdReq  = channelparam.GetChannelIdReq
	ChannelParamGetReq           = channelparam.GetReq
)

type (
	OrderGetPageReq        = order.GetPageReq
	OrderGetReq            = order.GetReq
	OrderGetBusinessIdReq  = order.GetBusinessIdReq
	OrderAddReq            = order.AddReq
	OrderIdReq             = order.IdReq
	OrderUpdateReq         = order.UpdateReq
	OrderDelReq            = order.DelReq
	OrderPaySuccessReq     = order.PaySuccessReq
	OrderPayFailureReq     = order.PayFailureReq
	OrderGetPayStateReq    = order.GetPayStateReq
	OrderGetPageResp       = order.GetPageResp
	OrderGetResp           = order.GetResp
	OrderGetBusinessIdResp = order.GetBusinessIdResp
	OrderGetPayStateResp   = order.GetPayStateResp
)
