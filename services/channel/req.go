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
	"github.com/go-the-way/unipay/services/base"
)

type (
	GetPageReq struct {
		base.PageReq

		OrderBy string `form:"order_by"` // 排序

		Id                         uint   `form:"id"`                             // 支付通道id
		Name                       string `form:"name"`                           // 支付通道名称
		Currency                   string `form:"currency"`                       // 货币类型CNY人民币USD美元
		Type                       string `form:"type"`                           // 支付通道类型
		AdminUrl                   string `form:"admin_url"`                      // 后台登录Url
		AdminUser                  string `form:"admin_user"`                     // 后台登录用户名
		AdminPasswd                string `form:"admin_passwd"`                   // 后台登录密码
		KeepDecimal                byte   `form:"keep_decimal"`                   // 保留小数：1保留2不保留
		AmountType                 byte   `form:"amount_type"`                    // 金额类型：1元2分
		AmountValidateCond         string `form:"amount_validate_cond"`           // 支付金额验证条件
		ReqUrl                     string `form:"req_url"`                        // 请求url
		ReqMethod                  string `form:"req_method"`                     // 请求方式
		ReqContentType             string `form:"req_content_type"`               // 请求数据类型
		NotifyPayContentType       string `form:"notify_pay_content_type"`        // 回调支付数据类型
		NotifyPayReturnContent     string `form:"notify_pay_return_content"`      // 回调支付成功返回内容
		NotifyPayReturnContentType string `form:"notify_pay_return_content_type"` // 回调支付成功返回数据类型
		State                      byte   `form:"state"`                          // 状态：1启用2禁用
		Sort                       byte   `form:"sort"`                           // 升序排序值
		Sort1                      byte   `form:"sort1"`                          // 升序排序值
		Sort2                      byte   `form:"sort2"`                          // 升序排序值
		Remark                     string `form:"remark"`                         // 备注
		CreateTime1                string `form:"create_time1"`                   // 创建时间
		CreateTime2                string `form:"create_time2"`                   // 创建时间
		UpdateTime1                string `form:"update_time1"`                   // 修改时间
		UpdateTime2                string `form:"update_time2"`                   // 修改时间
	}
	IdReq struct {
		Id uint `validate:"min(1,支付通道id不能为空)" json:"id"`
	}
	GetReq IdReq
	AddReq struct {
		Name                       string `validate:"minlength(1,支付通道名称不能为空) maxlength(50,支付通道名称长度不能超过50)" json:"name"`                  // 支付通道名称
		Currency                   string `validate:"enum(CNY|USD,货币类型不合法)" json:"currency"`                                             // 货币类型:CNY人民币 USD美元
		Type                       string `json:"type"`                                                                                  // 类型
		AdminUrl                   string `validate:"maxlength(500,后台登录Url长度不能超过500)" json:"admin_url"`                                  // 后台登录Url
		AdminUser                  string `validate:"maxlength(200,后台登录用户名长度不能超过200)" json:"admin_user"`                                 // 后台登录用户名
		AdminPasswd                string `validate:"maxlength(200,后台登录密码长度不能超过200)" json:"admin_passwd"`                                // 后台登录密码
		LogoUrl                    string `validate:"minlength(1,支付通道Logo不能为空) maxlength(500,支付通道Logo长度不能超过500)" json:"logo_url"`        // 支付通道LogoUrl
		PcLogoUrl                  string `validate:"minlength(1,支付通道PCLogo不能为空) maxlength(500,支付通道PCLogo长度不能超过500)" json:"pc_logo_url"` // 支付通道PCLogoUrl
		AmountType                 byte   `validate:"enum(1|2,金额类型不合法)" json:"amount_type"`                                              // 金额类型：1元2分
		KeepDecimal                byte   `validate:"enum(1|2,保留小数不合法)" json:"keep_decimal"`                                             // 保留小数：1保留2不保留
		AmountValidateCond         string `json:"amount_validate_cond"`                                                                  // 支付金额验证条件
		ReqUrl                     string `validate:"maxlength(500,请求url长度不能超过500)" json:"req_url"`                                      // 请求url
		ReqMethod                  string `json:"req_method"`                                                                            // 请求方式
		ReqContentType             string `json:"req_content_type"`                                                                      // 请求数据类型
		ReqSuccessExpr             string `validate:"maxlength(500,请求成功计算表达式长度不能超过500)" json:"req_success_expr"`                         // 请求成功计算表达式
		ReqPayPageUrlExpr          string `validate:"maxlength(500,请求支付页面Url获取表达式长度不能超过500)" json:"req_pay_page_url_expr"`               // 请求支付页面Url获取表达式
		ReqPayQrUrlExpr            string `validate:"maxlength(500,请求支付二维码Url获取表达式长度不能超过500)" json:"req_pay_qr_url_expr"`                // 请求支付二维码Url获取表达式
		ReqPayMessageExpr          string `validate:"maxlength(500,请求支付获取消息表达式长度不能超过500)" json:"req_pay_message_expr"`                   // 请求支付获取消息表达式
		NotifyPayContentType       string `json:"notify_pay_content_type"`                                                               // 回调支付数据类型
		NotifyPaySuccessExpr       string `validate:"maxlength(500,回调支付成功计算表达式长度不能超过500)"  json:"notify_pay_success_expr"`               // 回调支付成功计算表达式
		NotifyPayIdExpr            string `validate:"maxlength(500,回调支付成功获取Id表达式长度不能超过500)" json:"notify_pay_id_expr"`                   // 回调支付成功获取Id表达式
		NotifyPayReturnContent     string `validate:"maxlength(500,回调支付成功返回内容长度不能超过500)" json:"notify_pay_return_content"`               // 回调支付成功返回内容
		NotifyPayReturnContentType string `json:"notify_pay_return_content_type"`                                                        // 回调支付成功返回数据类型
		Remark                     string `validate:"maxlength(500,备注长度不能超过500)" json:"remark"`                                          // 备注
		Sort                       byte   `json:"sort"`                                                                                  // 升序排序值
	}
	UpdateReq struct {
		IdReq  `validate:"valid(T)"`
		AddReq `validate:"valid(T)"`
	}
	DelReq        IdReq
	EnableReq     IdReq
	DisableReq    IdReq
	GetMatchesReq struct {
		Amount uint   `form:"amount" json:"amount"` // 支付金额：单位分
		Order  string `form:"order" json:"order"`   // 排序规则
		Limit  uint   `form:"limit" json:"limit"`   // 返回限制数量，0不限制返回所有
	}
)
