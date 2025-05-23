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

package usdrate

import (
	"github.com/go-the-way/unipay/models"
	"github.com/go-the-way/unipay/services/base"
)

func (r *UpdateReq) Check() (err error) { return base.CheckRateValid(r.Rate) }

func (r *UpdateReq) transformCreate() *models.UsdRate { return &models.UsdRate{Rate: r.Rate} }
func (r *UpdateReq) transformUpdate() map[string]any  { return map[string]any{"rate": r.Rate} }
