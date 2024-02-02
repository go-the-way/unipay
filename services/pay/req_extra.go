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
	"net/url"
)

func (r *Req) ToMap(orderId string) map[string]any {
	nu, _ := url.Parse(r.NotifyUrl)
	nu.Query().Set("channel_id", fmt.Sprintf("%d", r.ChannelId))
	nu.Query().Set("order_id", orderId)
	nu.Query().Set("business_id1", r.BusinessId1)
	nu.Query().Set("business_id2", r.BusinessId2)
	nu.Query().Set("business_id3", r.BusinessId3)
	notifyUrl := nu.String()
	return map[string]any{
		"ChannelId":   fmt.Sprintf("%d", r.ChannelId),
		"Amount":      fmt.Sprintf("%d", r.Amount),
		"AmountYuan":  fmt.Sprintf("%d", r.Amount*100),
		"AmountFen":   fmt.Sprintf("%d", r.Amount),
		"Subject":     r.Subject,
		"ClientIp":    r.ClientIp,
		"NotifyUrl":   notifyUrl,
		"BusinessId1": r.BusinessId1,
		"BusinessId2": r.BusinessId2,
		"BusinessId3": r.BusinessId3,
		"Remark1":     r.Remark1,
		"Remark2":     r.Remark2,
		"Remark3":     r.Remark3,
	}
}
