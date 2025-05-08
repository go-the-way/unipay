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

package channel

import (
	"errors"
	"net/http"

	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/models"
	"github.com/go-the-way/unipay/services/base"
)

var typeMap = map[string]struct{}{"normal": {}, "erc20": {}, "trc20": {}}

func (r *AddReq) Check() (err error) {
	if cond := r.AmountValidateCond; cond != "" {
		if !pkg.ValidAmountCond(cond) {
			return errors.New("支付金额验证条件格式不合法")
		}
	}

	if _, ok := typeMap[r.Type]; !ok {
		return errors.New("类型不合法")
	}

	switch r.Type {
	case "normal":
		if len(r.ReqUrl) == 0 {
			return errors.New("请求url不能为空")
		}
		if !(r.ReqMethod == "GET" || r.ReqMethod == "POST") {
			return errors.New("请求方式不合法")
		}
		if !(r.ReqContentType == "json" || r.ReqContentType == "form" || r.ReqContentType == "urlencoded") {
			return errors.New("请求数据类型不合法")
		}
		if len(r.ReqSuccessExpr) == 0 {
			return errors.New("请求成功计算表达式不能为空")
		}
		if len(r.ReqPayMessageExpr) == 0 {
			return errors.New("请求支付获取消息表达式不能为空")
		}
		if !(r.NotifyPayContentType == "json" || r.NotifyPayContentType == "form" || r.NotifyPayContentType == "urlencoded") {
			return errors.New("回调支付数据类型不合法")
		}
		if len(r.NotifyPaySuccessExpr) == 0 {
			return errors.New("回调支付成功计算表达式不能为空")
		}
		if len(r.NotifyPayReturnContent) == 0 {
			return errors.New("回调支付成功返回内容不能为空")
		}
		if !(r.NotifyPayReturnContentType == "text" || r.NotifyPayReturnContentType == "json") {
			return errors.New("回调支付成功返回数据类型不合法")
		}
		if r.ReqPayPageUrlExpr == "" && r.ReqPayQrUrlExpr == "" {
			return errors.New("支付页面Url获取表达式和支付二维码Url获取表达式不能同时为空")
		}
		if r.ReqMethod == http.MethodGet && r.ReqContentType != "urlencoded" {
			return errors.New("当请求类型为GET时，请求数据类型仅支持urlencoded")
		}

	case "erc20", "trc20":
	}

	return
}

func (r *UpdateReq) Check() (err error) {
	return base.CheckChannelExist(r.Id)
}

func (r *DelReq) Check() (err error) { return base.CheckChannelExist(r.Id) }

func (r *AddReq) Transform() *models.Channel {
	return &models.Channel{
		BusinessId1:                r.BusinessId1,
		BusinessId2:                r.BusinessId2,
		BusinessId3:                r.BusinessId3,
		Name:                       r.Name,
		ProductName:                r.ProductName,
		Currency:                   r.Currency,
		Type:                       r.Type,
		AdminUrl:                   r.AdminUrl,
		AdminUser:                  r.AdminUser,
		AdminPasswd:                r.AdminPasswd,
		LogoUrl:                    r.LogoUrl,
		PcLogoUrl:                  r.PcLogoUrl,
		AmountType:                 r.AmountType,
		KeepDecimal:                r.KeepDecimal,
		AmountValidateCond:         r.AmountValidateCond,
		ReqUrl:                     r.ReqUrl,
		ReqMethod:                  r.ReqMethod,
		ReqContentType:             r.ReqContentType,
		ReqSuccessExpr:             r.ReqSuccessExpr,
		ReqPayPageUrlExpr:          r.ReqPayPageUrlExpr,
		ReqPayQrUrlExpr:            r.ReqPayQrUrlExpr,
		ReqPayMessageExpr:          r.ReqPayMessageExpr,
		NotifyPayContentType:       r.NotifyPayContentType,
		NotifyPaySuccessExpr:       r.NotifyPaySuccessExpr,
		NotifyPayIdExpr:            r.NotifyPayIdExpr,
		NotifyPayReturnContent:     r.NotifyPayReturnContent,
		NotifyPayReturnContentType: r.NotifyPayReturnContentType,
		State:                      models.ChannelStateDisable,
		Sort:                       r.Sort,
		Remark:                     r.Remark,
		Remark1:                    r.Remark1,
		Remark2:                    r.Remark2,
		Remark3:                    r.Remark3,
		CreateTime:                 pkg.TimeNowStr(),
		UpdateTime:                 pkg.TimeNowStr(),
	}
}

func (r *UpdateReq) Transform() map[string]any {
	return map[string]any{
		"business_id1":                   r.BusinessId1,
		"business_id2":                   r.BusinessId2,
		"business_id3":                   r.BusinessId3,
		"name":                           r.Name,
		"product_name":                   r.ProductName,
		"currency":                       r.Currency,
		"admin_url":                      r.AdminUrl,
		"admin_user":                     r.AdminUser,
		"admin_passwd":                   r.AdminPasswd,
		"logo_url":                       r.LogoUrl,
		"pc_logo_url":                    r.PcLogoUrl,
		"amount_type":                    r.AmountType,
		"keep_decimal":                   r.KeepDecimal,
		"amount_validate_cond":           r.AmountValidateCond,
		"req_url":                        r.ReqUrl,
		"req_method":                     r.ReqMethod,
		"req_content_type":               r.ReqContentType,
		"req_success_expr":               r.ReqSuccessExpr,
		"req_pay_page_url_expr":          r.ReqPayPageUrlExpr,
		"req_pay_qr_url_expr":            r.ReqPayQrUrlExpr,
		"req_pay_message_expr":           r.ReqPayMessageExpr,
		"notify_pay_content_type":        r.NotifyPayContentType,
		"notify_pay_success_expr":        r.NotifyPaySuccessExpr,
		"notify_pay_id_expr":             r.NotifyPayIdExpr,
		"notify_pay_return_content":      r.NotifyPayReturnContent,
		"notify_pay_return_content_type": r.NotifyPayReturnContentType,
		"remark":                         r.Remark,
		"remark1":                        r.Remark1,
		"remark2":                        r.Remark2,
		"remark3":                        r.Remark3,
		"sort":                           r.Sort,
		"update_time":                    pkg.TimeNowStr(),
	}
}
