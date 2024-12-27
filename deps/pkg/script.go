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
	"context"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

var (
	importModules = []string{
		"text", "crypto", "rand", "sort", "util",
	}
	replaceMap = map[string]string{
		"upper": "text.to_upper",
		"lower": "text.to_lower",
		"join":  "text.join",

		"rand_num": "rand.rand_num",
		"rand_str": "rand.rand_str",

		"md5": "crypto.md5_hex",

		"rsa": "crypto.sha256WithRSA",

		"sort": "sort.sort",

		"to_int": "util.to_int",
	}
)

func importAndReplace(content string) (script string) {
	var results []string

	for k, v := range replaceMap {
		content = strings.ReplaceAll(content, k, v)
	}
	for _, mod := range importModules {
		results = append(results, mod+` := import("`+mod+`")`)
	}
	results = append(results, content)
	return strings.Join(results, "\n")
}

func EvalString(content string, in map[string]any) (out string, err error) {
	evalEd, eErr := Eval(content, in)
	if eErr != nil {
		err = eErr
		return
	}
	out = evalEd.(string)
	return
}

func EvalInt(content string, in map[string]any) (out int, err error) {
	evalEd, eErr := Eval(content, in)
	if eErr != nil {
		err = eErr
		return
	}
	out = evalEd.(int)
	return
}

func EvalBool(content string, in map[string]any) (out bool, err error) {
	evalEd, eErr := Eval(content, in)
	if eErr != nil {
		err = eErr
		return
	}
	out = evalEd.(bool)
	return
}

func Eval(content string, in map[string]any) (out any, err error) {
	content = strings.ReplaceAll(strings.TrimSpace(content), "\n", "")
	content = `var := func(){return ` + content + `}()`
	input := importAndReplace(content)
	input = strings.ReplaceAll(input, "$", "")
	script := tengo.NewScript([]byte(input))
	script.SetImports(stdlib.GetModuleMap(importModules...))

	for k, v := range in {
		if err = script.Add(k, v); err != nil {
			return
		}
	}

	compiled, err := script.RunContext(context.Background())
	if err != nil {
		return
	}

	variable := compiled.Get("var")
	out = variable.Value()

	return
}
