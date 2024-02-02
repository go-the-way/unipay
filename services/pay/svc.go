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

package pay

import "net/http"

var (
	Service SVC = &service{}

	ReqPay    = Service.ReqPay
	NotifyPay = Service.NotifyPay
)

type SVC interface {
	// ReqPay
	//
	// # 请求支付接口
	//
	// # 时间变量 => $Time.
	//
	// 当前时间`2006-01-02 15:04:05` => NowTime
	//
	// 当前时间`20060102150405` => NowTimeNum
	//
	// 当前时间戳`1705976043` => NowTimestamp
	//
	// 当前时间戳`1705976043000` => NowTimestampLong
	//
	// # 支付变量 => $Pay.
	//
	// 支付金额 => Amount
	//
	// 支付金额（单位：元）=> AmountYuan
	//
	// 支付金额（单位：分）=> AmountFen
	//
	// 支付主题 => Subject
	//
	// 客户端Ip => ClientIp
	//
	// 回调Url => NotifyUrl
	//
	// 业务id1 => BusinessId1
	//
	// 业务id2 => BusinessId2
	//
	// 业务id3 => BusinessId3
	//
	// 备注1 => Remark1
	//
	// 备注2 => Remark2
	//
	// 备注3 => Remark3
	//
	// # 支付通道变量 => $Channel.
	//
	// ref => models.Channel
	//
	// # 支付参数变量 => $Param.
	//
	// ref => models.ChannelParam
	ReqPay(req Req) (resp Resp, err error)
	// NotifyPay
	//
	// # 回调支付接口
	NotifyPay(req *http.Request, resp http.ResponseWriter, r NotifyReq, paidCallback func(req NotifyReq)) (err error)
}
