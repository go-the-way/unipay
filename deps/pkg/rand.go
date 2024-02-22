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
	"math/rand"
	"strings"
)

func RandStr(length int, numOnly ...bool) string {
	no := false
	if numOnly != nil && len(numOnly) > 0 {
		no = numOnly[0]
	}
	results := make([]string, 0)
	symbol := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-*/_"
	for i := 0; i < length; i++ {
		randLen := len(symbol)
		if no {
			randLen = 10
		}
		cur := symbol[rand.Intn(10000)%randLen]
		results = append(results, fmt.Sprintf("%c", cur))
	}
	return strings.Join(results, "")
}
