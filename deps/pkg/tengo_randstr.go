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
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

var randModule = make(map[string]tengo.Object, 2)

func init() {
	randModule["rand_str"] = &tengo.UserFunction{Name: "rand_str", Value: getRandFunc(false)}
	randModule["rand_num"] = &tengo.UserFunction{Name: "rand_num", Value: getRandFunc(true)}
	stdlib.BuiltinModules["rand"] = randModule
}

func getRandFunc(numOnly bool) func(args ...tengo.Object) (ret tengo.Object, err error) {
	return func(args ...tengo.Object) (ret tengo.Object, err error) {
		if len(args) != 1 {
			return nil, tengo.ErrWrongNumArguments
		}
		i, ok := tengo.ToInt(args[0])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &tengo.String{Value: RandStr(i, numOnly)}, nil
	}
}
