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

import (
	"net/http"

	"github.com/rwscode/unipay/deps/pkg"
)

type (
	Log       = UnipayLog
	ApiLog    = UnipayApiLog
	UnipayLog struct {
		Id         uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                  // id
		Text       string `gorm:"column:text;type:varchar(2000);not null;default:'';comment:日志内容" json:"text"`             // 日志内容
		CreateTime string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"` // 创建时间
	}
	UnipayApiLog struct {
		Id          uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                      // id
		ReqUrl      string `gorm:"column:req_url;type:varchar(500);not null;default:'';comment:请求Url" json:"req_url"`           // 请求Url
		ReqMethod   string `gorm:"column:req_method;type:varchar(10);not null;default:'GET';comment:请求方法" json:"req_method"`    // 请求方法
		ReqParam    string `gorm:"column:req_param;type:varchar(2000);not null;default:'';comment:请求参数" json:"req_param"`       // 请求参数
		RespContent string `gorm:"column:resp_content;type:varchar(2000);not null;default:'';comment:影响内容" json:"resp_content"` // 响应内容
		RespCode    string `gorm:"column:resp_code;type:varchar(20);not null;default:'';comment:响应Code码" json:"resp_code"`      // 响应Code码
		CreateTime  string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`     // 创建时间
	}
)

func NewLog(text string) *Log { return &Log{Text: text, CreateTime: pkg.TimeNowStr()} }

func NewApiLog(reqUrl, reqMethod, reqParam, respContent, respCode string) *ApiLog {
	return &ApiLog{
		ReqUrl:      reqUrl,
		ReqMethod:   reqMethod,
		ReqParam:    reqParam,
		RespContent: respContent,
		RespCode:    respCode,
		CreateTime:  pkg.TimeNowStr(),
	}
}

func NewApiLogGetNoParam(reqUrl, respContent, respCode string) *ApiLog {
	return NewApiLog(reqUrl, http.MethodGet, "", respContent, respContent)
}
