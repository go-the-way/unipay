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

import (
	_ "embed"

	"bytes"
	"errors"
	"html/template"
)

var (
	//go:embed e20.html
	e20Html string
)

func (s *service) E20Html(req E20HtmlReq) (resp E20HtmlResp, err error) {
	buf := &bytes.Buffer{}
	tpl := template.Must(template.New("").Parse(e20Html))
	if err = tpl.Execute(buf, req); err != nil {
		err = errors.New("解析错误：" + err.Error())
		return
	}
	resp.Html = buf.String()
	buf.Reset()
	return
}
