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

package walletaddress

import (
	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/models"
	"github.com/go-the-way/unipay/services/base"
)

func (r *AddReq) Check() (err error) { return base.CheckAddressProtocolExists(r.Address, r.Protocol) }

func (r *UpdateReq) Check() (err error) { return base.CheckWalletAddressExist(r.Id) }

func (r *DelReq) Check() (err error) { return base.CheckWalletAddressExist(r.Id) }

func (r *AddReq) transform() *models.WalletAddress {
	return &models.WalletAddress{
		Address:     r.Address,
		Protocol:    r.Protocol,
		State:       models.WalletAddressStateEnable,
		BusinessId1: r.BusinessId1,
		BusinessId2: r.BusinessId2,
		BusinessId3: r.BusinessId3,
		Remark:      r.Remark,
		CreateTime:  pkg.TimeNowStr(),
		UpdateTime:  pkg.TimeNowStr(),
	}
}

func (r *UpdateReq) transform() map[string]any {
	return map[string]any{
		"address":      r.Address,
		"protocol":     r.Protocol,
		"business_id1": r.BusinessId1,
		"business_id2": r.BusinessId2,
		"business_id3": r.BusinessId3,
		"remark":       r.Remark,
		"update_time":  pkg.TimeNowStr(),
	}
}
