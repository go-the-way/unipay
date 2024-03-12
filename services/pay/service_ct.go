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

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rwscode/unipay/models"
	"github.com/rwscode/unipay/services/channel"
	"github.com/rwscode/unipay/services/channelparam"
)

func getReqCT(c channel.GetResp, cp channelparam.GetChannelIdResp, params map[string]any) (body string, form url.Values, contentType string) {
	body, form = ctFuncMap[c.ReqContentType](getPassMap(cp), params)
	contentType = ctMap[c.ReqContentType]
	return
}

const (
	reqContentTypeText       = "text/plain; charset=utf-8"
	reqContentTypeJSON       = "application/json; charset=utf-8"
	reqContentTypeForm       = "multipart/form-data"
	reqContentTypeUrlencoded = "application/x-www-form-urlencoded"
)

var (
	ctMap = map[string]string{
		"text":       reqContentTypeText,
		"json":       reqContentTypeJSON,
		"form":       reqContentTypeForm,
		"urlencoded": reqContentTypeUrlencoded,
	}
	ctFuncMap = map[string]ctFunc{
		"json":       buildJsonCT,
		"form":       buildFormCT,
		"urlencoded": buildUrlencodedCT,
	}
	ctRespFuncMap = map[string]ctRespFunc{
		"json":       parseJsonCTResp,
		"form":       parseFormCTResp,
		"urlencoded": parseUrlencodedCTResp,
	}
)

type (
	ctFunc     func(passMap map[string]struct{}, params map[string]any) (body string, form url.Values)
	ctRespFunc func(req *http.Request) map[string]any
)

func getPassMap(cp channelparam.GetChannelIdResp) (passMap map[string]struct{}) {
	passMap = map[string]struct{}{}
	for _, p := range cp.List {
		switch p.Pass {
		case models.ChannelParamPassYes:
			passMap[p.Name] = struct{}{}
		case models.ChannelParamPassNo:
		}
	}
	return
}

func buildJsonCT(passMap map[string]struct{}, params map[string]any) (body string, form url.Values) {
	toPassMap := map[string]any{}
	for k, v := range params {
		if _, ok := passMap[k]; ok {
			toPassMap[k] = v
		}
	}
	buf, _ := json.Marshal(toPassMap)
	body = string(buf)
	return
}

func buildFormCT(passMap map[string]struct{}, params map[string]any) (body string, form url.Values) {
	values := url.Values{}
	for k, v := range params {
		if _, ok := passMap[k]; ok {
			values.Set(k, fmt.Sprintf("%v", v))
		}
	}
	form = values
	return
}

func buildUrlencodedCT(passMap map[string]struct{}, params map[string]any) (body string, form url.Values) {
	values := &url.Values{}
	for k, v := range params {
		if _, ok := passMap[k]; ok {
			values.Set(k, fmt.Sprintf("%v", v))
		}
	}
	body = values.Encode()
	return
}

func parseJsonCTResp(req *http.Request) map[string]any {
	var m map[string]any
	buf, _ := io.ReadAll(req.Body)
	_ = json.Unmarshal(buf, &m)
	return m
}

func parseFormCTResp(req *http.Request) map[string]any {
	var m map[string]any
	for k := range req.PostForm {
		m[k] = req.PostFormValue(k)
	}
	return m
}

func parseUrlencodedCTResp(req *http.Request) map[string]any {
	var m map[string]any
	for k := range req.Form {
		m[k] = req.FormValue(k)
	}
	return m
}
