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

package channel

var (
	Service SVC = &service{}

	GetPage    = Service.GetPage
	Get        = Service.Get
	Add        = Service.Add
	Update     = Service.Update
	Del        = Service.Del
	Enable     = Service.Enable
	Disable    = Service.Disable
	GetMatches = Service.GetMatches
	E20Html    = Service.E20Html
)

type SVC interface {
	GetPage(req GetPageReq) (resp GetPageResp, err error)
	Get(req GetReq) (resp GetResp, err error)
	Add(req AddReq) (err error)
	Update(req UpdateReq) (err error)
	Del(req DelReq) (err error)
	Enable(req EnableReq) (err error)
	Disable(req DisableReq) (err error)
	GetMatches(req GetMatchesReq) (resp GetMatchesResp, err error)
	E20Html(req E20HtmlReq) (resp E20HtmlResp, err error)
}
