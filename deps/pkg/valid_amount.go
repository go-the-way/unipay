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
	"regexp"
	"strconv"
	"strings"
)

func ValidAmountCond(cond string) (ok bool) {
	// X1-Y1,X2-Y2,...
	// N1,N2,N3,...
	re := regexp.MustCompile(`^(,?(\d+-\d+|\d+),?)+$`)
	return re.MatchString(cond)
}

func ValidAmount(amount uint, cond string) (valid bool) {
	condS := strings.Split(strings.TrimSpace(cond), ",")
	if len(condS) <= 0 {
		return
	}
	var (
		condExprS []condExpr
		vs        []uint
	)
	for _, cnd := range condS {
		// X1-Y1,X2-Y2,...
		if strings.Contains(cnd, "-") {
			cs := strings.Split(cnd, "-")
			c1, c2 := cs[0], cs[1]
			r1, _ := strconv.ParseUint(c1, 10, 64)
			r2, _ := strconv.ParseUint(c2, 10, 64)
			condExprS = append(condExprS, newRangeCondExpr(uint(r1), uint(r2)))
		} else {
			// N1,N2,N3,...
			v, _ := strconv.ParseUint(cnd, 10, 64)
			vs = append(vs, uint(v))
			valueExpr := newValuesCondExpr(vs...)
			condExprS = append(condExprS, valueExpr)
		}
	}
	for _, cnd := range condExprS {
		if cnd.Valid(amount) {
			valid = true
			break
		}
	}
	return
}

type (
	condExpr       interface{ Valid(value uint) (valid bool) }
	rangeCondExpr  struct{ R1, R2 uint }
	valuesCondExpr struct{ Vs map[uint]struct{} }
)

func newValuesCondExpr(values ...uint) *valuesCondExpr {
	var m map[uint]struct{}
	if values != nil {
		for _, v := range values {
			m[v] = struct{}{}
		}
	}
	return &valuesCondExpr{m}
}
func newRangeCondExpr(r1 uint, r2 uint) *rangeCondExpr  { return &rangeCondExpr{R1: r1, R2: r2} }
func (c *rangeCondExpr) v1(value uint) (valid bool)     { return c.R2 == 0 && value >= c.R1 }
func (c *rangeCondExpr) v2(value uint) (valid bool)     { return value >= c.R1 && value <= c.R2 }
func (c *rangeCondExpr) Valid(value uint) (valid bool)  { return c.v1(value) || c.v2(value) }
func (c *valuesCondExpr) Add(value uint)                { c.Vs[value] = struct{}{} }
func (c *valuesCondExpr) Valid(value uint) (valid bool) { _, ok := c.Vs[value]; return ok }
