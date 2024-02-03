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

package models

import "fmt"

type (
	Channel       = UnipayChannel
	UnipayChannel struct {
		Id                         uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:支付通道id" json:"id"`                                                                 // 支付通道id
		Name                       string `gorm:"column:name;type:varchar(200);not null;default:'';comment:支付通道名称" json:"name"`                                                               // 支付通道名称
		AdminUrl                   string `gorm:"column:admin_url;type:varchar(500);not null;default:'';comment:后台登录Url" json:"admin_user"`                                                   // 后台登录Url
		AdminUser                  string `gorm:"column:admin_user;type:varchar(200);not null;default:'';comment:后台登录用户名" json:"admin_user"`                                                  // 后台登录用户名
		AdminPasswd                string `gorm:"column:admin_passwd;type:varchar(200);not null;default:'';comment:后台登录密码" json:"admin_passwd"`                                               // 后台登录密码
		LogoUrl                    string `gorm:"column:logo_url;type:varchar(500);not null;default:'';comment:支付通道LogoUrl" json:"logo_url"`                                                  // 支付通道LogoUrl
		AmountType                 byte   `gorm:"column:amount_type;type:tinyint;not null;default:1;comment:金额类型：1元2分" json:"amount_type"`                                                    // 金额类型：1元2分
		AmountValidateCond         string `gorm:"column:amount_validate_cond;type:varchar(500);not null;default:'';comment:支付金额验证条件" json:"amount_validate_cond"`                             // 支付金额验证条件
		ReqUrl                     string `gorm:"column:req_url;type:varchar(500);not null;default:'';comment:请求url" json:"req_url"`                                                          // 请求url
		ReqMethod                  string `gorm:"column:req_method;type:varchar(20);not null;default:'POST';comment:请求方式" json:"req_method"`                                                  // 请求方式
		ReqContentType             string `gorm:"column:req_content_type;type:varchar(50);not null;default:'json';comment:请求数据类型" json:"req_content_type"`                                    // 请求数据类型
		ReqSuccessExpr             string `gorm:"column:req_success_expr;type:varchar(500);not null;default:'';comment:请求成功计算表达式" json:"req_success_expr"`                                    // 请求成功计算表达式
		ReqPayPageUrlExpr          string `gorm:"column:req_pay_page_url_expr;type:varchar(500);not null;default:'';comment:请求支付页面Url获取表达式" json:"req_pay_page_url_expr"`                     // 请求支付页面Url获取表达式
		ReqPayQrUrlExpr            string `gorm:"column:req_pay_qr_url_expr;type:varchar(500);not null;default:'';comment:请求支付二维码Url获取表达式" json:"req_pay_qr_url_expr"`                        // 请求支付二维码Url获取表达式
		ReqPayMessageExpr          string `gorm:"column:req_pay_message_expr;type:varchar(500);not null;default:'';comment:请求支付获取消息表达式" json:"req_pay_message_expr"`                          // 请求支付获取消息表达式
		NotifyPayContentType       string `gorm:"column:notify_pay_content_type;type:varchar(500);not null;default:'json';comment:回调支付数据类型" json:"notify_pay_content_type"`                   // 回调支付数据类型
		NotifyPaySuccessExpr       string `gorm:"column:notify_pay_success_expr;type:varchar(500);not null;default:'';comment:回调支付成功计算表达式" json:"notify_pay_success_expr"`                    // 回调支付成功计算表达式
		NotifyPayIdExpr            string `gorm:"column:notify_pay_id_expr;type:varchar(500);not null;default:'';comment:回调支付成功获取Id表达式" json:"notify_pay_id_expr"`                            // 回调支付成功获取Id表达式
		NotifyPayReturnContent     string `gorm:"column:notify_pay_return_content;type:varchar(500);not null;default:'success';comment:回调支付成功返回内容" json:"notify_pay_return_content"`          // 回调支付成功返回内容
		NotifyPayReturnContentType string `gorm:"column:notify_pay_return_content_type;type:varchar(500);not null;default:'text';comment:回调支付成功返回数据类型" json:"notify_pay_return_content_type"` // 回调支付成功返回数据类型
		State                      byte   `gorm:"column:state;type:tinyint;not null;default:1;comment:状态：1启用2禁用;index" json:"state"`                                                          // 状态：1启用2禁用
		Sort                       byte   `gorm:"column:sort;type:tinyint;not null;default:0;comment:升序排序值" json:"sort"`                                                                      // 升序排序值
		Remark                     string `gorm:"column:remark;type:varchar(2000);not null;default:'';comment:备注" json:"remark"`                                                              // 备注
		CreateTime                 string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`                                                    // 创建时间
		UpdateTime                 string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"`                                                    // 修改时间
	}
)

func (c *Channel) ToMap() map[string]any {
	return map[string]any{
		"Id":                         fmt.Sprintf("%d", c.Id),
		"Name":                       c.Name,
		"AdminUrl":                   c.AdminUrl,
		"AdminUser":                  c.AdminUser,
		"AdminPasswd":                c.AdminPasswd,
		"LogoUrl":                    c.LogoUrl,
		"AmountType":                 c.AmountType,
		"AmountValidateCond":         c.AmountValidateCond,
		"ReqUrl":                     c.ReqUrl,
		"ReqMethod":                  c.ReqMethod,
		"ReqContentType":             c.ReqContentType,
		"ReqSuccessExpr":             c.ReqSuccessExpr,
		"ReqPayPageUrlExpr":          c.ReqPayPageUrlExpr,
		"ReqPayQrUrlExpr":            c.ReqPayQrUrlExpr,
		"ReqPayMessageExpr":          c.ReqPayMessageExpr,
		"NotifyPayContentType":       c.NotifyPayContentType,
		"NotifyPaySuccessExpr":       c.NotifyPaySuccessExpr,
		"NotifyPayIdExpr":            c.NotifyPayIdExpr,
		"NotifyPayReturnContent":     c.NotifyPayReturnContent,
		"NotifyPayReturnContentType": c.NotifyPayReturnContentType,
		"State":                      fmt.Sprintf("%d", c.State),
		"Sort":                       fmt.Sprintf("%d", c.Sort),
		"Remark":                     c.Remark,
		"CreateTime":                 c.CreateTime,
		"UpdateTime":                 c.UpdateTime,
	}
}

const (
	_ byte = iota
	// ChannelAmountTypeYuan 元
	ChannelAmountTypeYuan
	// ChannelAmountTypeFen 分
	ChannelAmountTypeFen
)

const (
	_ byte = iota
	// ChannelStateEnable 启用
	ChannelStateEnable
	// ChannelStateDisable 禁用
	ChannelStateDisable
)
