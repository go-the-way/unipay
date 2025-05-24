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

package apiconfig

import "github.com/go-the-way/unipay/models"

func (r *UpdateReq) transformCreate() models.ApiConfig {
	return models.ApiConfig{
		Erc20Apikey:       r.Erc20Apikey,
		BackupPlan:        r.BackupPlan,
		BackupVar1:        r.BackupVar1,
		ValidPeriodMinute: r.ValidPeriodMinute,
	}
}

func (r *UpdateReq) transformUpdate() map[string]any {
	return map[string]any{
		"erc20_apikey":        r.Erc20Apikey,
		"backup_plan":         r.BackupPlan,
		"backup_var1":         r.BackupVar1,
		"valid_period_minute": r.ValidPeriodMinute,
	}
}
