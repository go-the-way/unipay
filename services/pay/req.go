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

type (
	Req struct {
		ChannelId   uint   `validate:"min(1,支付通道id不能为空)" form:"channel_id" json:"channel_id"`                                      // 支付通道id
		Amount      uint   `validate:"min(1,支付金额不能少于1)" form:"amount" json:"amount"`                                               // 支付金额（单位：分）
		Subject     string `validate:"minlength(1,支付主题不能为空) maxlength(200,支付主题长度不能超过200)" form:"subject" json:"subject"`           // 支付主题
		ClientIp    string `validate:"minlength(1,客户端Ip不能为空) maxlength(20,客户端Ip长度不能超过50)" form:"client_ip" json:"client_ip"`       // 客户端Ip
		NotifyUrl   string `validate:"minlength(1,回调Url不能为空) maxlength(500,回调Url长度不能超过500)" form:"notify_url" json:"notify_url"`   // 回调Url
		BusinessId1 string `validate:"minlength(1,业务id1不能为空) maxlength(50,业务id1长度不能超过50)" form:"business_id1" json:"business_id1"` // 业务id1
		BusinessId2 string `validate:"maxlength(50,业务id2长度不能超过50)" form:"business_id2" json:"business_id2"`                        // 业务id2
		BusinessId3 string `validate:"maxlength(50,业务id3长度不能超过50)" form:"business_id3" json:"business_id3"`                        // 业务id3
		Remark1     string `validate:"maxlength(500,备注1长度不能超过500)" form:"remark1" json:"remark1"`                                  // 备注1
		Remark2     string `validate:"maxlength(500,备注2长度不能超过500)" form:"remark2" json:"remark2"`                                  // 备注2
		Remark3     string `validate:"maxlength(500,备注3长度不能超过500)" form:"remark3" json:"remark3"`                                  // 备注3
	}
	NotifyReq struct {
		ChannelId  uint   `validate:"min(1,支付通道id不能为空)" form:"channel_id" json:"channel_id"`                                  // 支付通道id
		OrderId    string `validate:"minlength(1,订单Id不能为空) maxlength(50,订单Id长度不能超过50)" form:"order_id" json:"order_id"`       // 订单Id
		BusinessId string `validate:"minlength(1,业务Id不能为空) maxlength(50,业务Id长度不能超过50)" form:"business_id" json:"business_id"` // 业务Id
	}
)
