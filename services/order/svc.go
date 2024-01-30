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

var Service = Impl()

type SVC interface {
	GetPage(req GetPageReq, order ...string) (resp GetPageResp, err error)
	Get(req GetReq) (resp GetResp, err error)
	GetBusinessId(req GetBusinessIdReq) (resp GetBusinessIdResp, err error)
	Add(req AddReq) (err error)
	Update(req UpdateReq) (err error)
	Del(req DelReq) (err error)
	PaySuccess(req PaySuccessReq) (err error)
	PayFailure(req PayFailureReq) (err error)
	GetPayState(req GetPayStateReq) (resp GetPayStateResp, err error)
}
