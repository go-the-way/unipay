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
	"fmt"
	"sort"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

var sortModule = make(map[string]tengo.Object, 1)

func init() {
	sortModule["sort"] = &tengo.UserFunction{Name: "sort", Value: sortFunc}
	stdlib.BuiltinModules["sort"] = sortModule
}

func sortFunc(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	var srcArr = make([]string, 0)
	switch arg0 := args[0].(type) {
	case *tengo.Array:
		for idx, a := range arg0.Value {
			as, ok := tengo.ToString(a)
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			srcArr = append(srcArr, as)
		}
	case *tengo.ImmutableArray:
		for idx, a := range arg0.Value {
			as, ok := tengo.ToString(a)
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("first[%d]", idx),
					Expected: "string(compatible)",
					Found:    a.TypeName(),
				}
			}
			srcArr = append(srcArr, as)
		}
	default:
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}
	sort.Strings(srcArr)
	tArr := make([]tengo.Object, len(srcArr), len(srcArr))
	for i, src := range srcArr {
		tArr[i] = &tengo.String{Value: src}
	}
	return &tengo.Array{Value: tArr}, nil
}
