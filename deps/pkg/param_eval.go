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

package pkg

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/rwscode/unipay/models"
)

type (
	paramValue struct {
		Name   string
		Value  string
		Result string
		Ref    map[string]struct{}
	}
	pV struct {
		RefCount int
		Params   []paramValue
	}
)

func EvalParams(payMap, channelMap map[string]any, ps []models.ChannelParam) (map[string]any, error) {
	params := getParams(ps)
	sortParams(params)
	paramMap := map[string]any{}
	data := map[string]any{"Time": GetTimeMap(), "Channel": channelMap, "Pay": payMap, "Param": paramMap}
	for i, p := range params {
		if !strings.Contains(p.Value, "$") {
			data["__self__"] = p.Value
			p.Value = "$__self__"
		}
		output, err := Eval(p.Value, data)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("解析参数[%s]表达式[%s]错误：%s", p.Name, p.Value, err.Error()))
		}
		params[i].Result = fmt.Sprintf("%v", output)
		paramMap[p.Name] = params[i].Result
	}
	resultMap := map[string]any{}
	for _, p := range params {
		resultMap[p.Name] = p.Result
	}
	return resultMap, nil
}

func getParams(ps []models.ChannelParam) []paramValue {
	var params []paramValue
	for _, a := range ps {
		params = append(params, paramValue{Name: a.Name, Value: a.Value})
	}
	return params
}

func sortParams(params []paramValue) {
	pv := pV{RefCount: calcRefs(params), Params: params}
	for ; pv.RefCount > 0; sort.Sort(&pv) {
	}
}

func calcRefs(params []paramValue) (refCount int) {
	re := regexp.MustCompile(`\$Param.(\w+)`)
	for i, p := range params {
		if strings.Contains(p.Value, "$Param.") {
			params[i].Ref = make(map[string]struct{})
			subs := re.FindAllStringSubmatch(p.Value, -1)
			for _, sub := range subs {
				if len(sub) > 1 {
					params[i].Ref[sub[1]] = struct{}{}
					refCount++
				}
			}
		}
	}
	return
}

func (p *pV) Len() int { return len(p.Params) }

func (p *pV) Less(i, j int) bool {
	iRef := p.Params[i].Ref
	jRef := p.Params[j].Ref
	iC, jC := len(iRef), len(jRef)
	if iC > 0 || jC > 0 {
		p.RefCount--
	}
	if iC <= 0 || jC <= 0 {
		return iC < jC
	}
	_, ok := p.Params[j].Ref[p.Params[i].Name]
	p.RefCount--
	return ok
}

func (p *pV) Swap(i, j int) { p.Params[i], p.Params[j] = p.Params[j], p.Params[i] }
