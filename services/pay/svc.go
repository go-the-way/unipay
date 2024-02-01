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
	// ReqPay 请求支付接口
	ReqPay(req Req) (resp Resp, err error)
	// NotifyPay 回调支付接口
	NotifyPay(req *http.Request, resp http.ResponseWriter, r NotifyReq, paidCallback func()) (err error)
}
