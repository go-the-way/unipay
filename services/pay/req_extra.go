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
	"fmt"
	"strings"
)

func (r *Req) ToMap(orderId string) map[string]any {
	notifyUrl := r.NotifyUrl
	if strings.Index(r.NotifyUrl, "?") == -1 {
		notifyUrl += "?"
	} else {
		notifyUrl += "&"
	}
	notifyUrl += fmt.Sprintf("channel_id=%d", r.ChannelId)
	notifyUrl += fmt.Sprintf("&order_id=%s", orderId)
	notifyUrl += fmt.Sprintf("&business_id1=%s", r.BusinessId1)
	notifyUrl += fmt.Sprintf("&business_id2=%s", r.BusinessId2)
	notifyUrl += fmt.Sprintf("&business_id3=%s", r.BusinessId3)
	return map[string]any{
		"ChannelId":   fmt.Sprintf("%d", r.ChannelId),
		"AmountYuan":  r.AmountYuan,
		"AmountFen":   r.AmountFen,
		"Subject":     r.Subject,
		"ClientIp":    r.ClientIp,
		"NotifyUrl":   notifyUrl,
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
