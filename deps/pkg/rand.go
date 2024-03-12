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
	"time"
)

func RandStr(length int, numOnly ...bool) string {
	num := false
	if numOnly != nil && len(numOnly) > 0 {
		num = numOnly[0]
	}
	results := make([]string, 0)
	symbolNum := "0123456789"
	symbol := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		var cur byte
		if !num {
			cur = symbol[rand.Intn(len(symbol))]
		} else {
			cur = symbolNum[rand.Intn(len(symbolNum))]
		}
		results = append(results, fmt.Sprintf("%c", cur))
	}
	return strings.Join(results, "")
}
