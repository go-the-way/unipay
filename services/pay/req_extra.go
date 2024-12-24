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

package pay

import (
	"encoding/base64"
	"fmt"
)

func (r *Req) ToMap(orderId, realAmountYuan, realAmountFen string) map[string]any {
	notifyUrl := r.NotifyUrl
	notifyUrl += fmt.Sprintf("/%d", r.ChannelId)
	notifyUrl += fmt.Sprintf("/%s", orderId)
	businessId := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s-%s-%s", r.BusinessId1, r.BusinessId2, r.BusinessId3)))
	notifyUrl += fmt.Sprintf("/%s", businessId)
	r.NotifyUrl = notifyUrl

	returnUrl := r.ReturnUrl
	returnUrl += fmt.Sprintf("/%s", orderId)

	return map[string]any{
		"ChannelId":   fmt.Sprintf("%d", r.ChannelId),
		"AmountYuan":  realAmountYuan,
		"AmountFen":   realAmountFen,
		"Subject":     r.Subject,
		"ClientIp":    r.ClientIp,
		"NotifyUrl":   notifyUrl,
		"ReturnUrl":   returnUrl,
		"OrderId":     orderId,
		"BusinessId1": r.BusinessId1,
		"BusinessId2": r.BusinessId2,
		"BusinessId3": r.BusinessId3,
		"Other1":      r.Other1,
		"Other2":      r.Other2,
		"Other3":      r.Other3,
		"Remark1":     r.Remark1,
		"Remark2":     r.Remark2,
		"Remark3":     r.Remark3,
		"AppWakeUri":  r.AppWakeUri,
	}
}
