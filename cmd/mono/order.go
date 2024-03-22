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

package main

import (
	_ "embed"

	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	//go:embed index.html
	indexHtml string
	orderMap  = map[string]*orderInfo{}
	mu        = &sync.Mutex{}
)

type (
	createOrderReq struct {
		AppKey      string `json:"app_key"`
		Rand        string `json:"rand"`
		Subject     string `json:"subject"`
		Price       string `json:"price"`
		NotifyUrl   string `json:"notify_url"`
		RedirectUrl string `json:"redirect_url"`
		Sign        string `json:"sign"`
	}
	orderInfo struct {
		OrderId string
		Paid    bool
		createOrderReq
	}
)

func (r *orderInfo) paidJson() string {
	buf, _ := json.Marshal(map[string]any{"data": map[string]any{"paid": true, "order_id": r.OrderId, "message": "success"}})
	return string(buf)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	req, err := parseParam(r)
	if err != nil {
		writeJsonNon200(w, err.Error())
		return
	}
	if err = checkParam(req); err != nil {
		writeJsonNon200(w, err.Error())
		return
	}
	if err = checkSign(req); err != nil {
		writeJsonNon200(w, err.Error())
		return
	}
	orderId := createOrderId()
	mu.Lock()
	orderMap[orderId] = &orderInfo{OrderId: orderId, createOrderReq: req}
	mu.Unlock()
	writeJson200(w, map[string]any{"order_id": orderId, "pay_url": getPayOrderUrl(orderId)})
}

func writeJson200(w http.ResponseWriter, dataMap map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	buf, _ := json.Marshal(map[string]any{"code": "200", "message": "ok", "data": dataMap})
	_, _ = w.Write(buf)
}

func write200(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	_, _ = w.Write([]byte(message))
}

func writeJsonNon200(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	buf, _ := json.Marshal(map[string]any{"code": "400", "message": message})
	_, _ = w.Write(buf)
}

func writeNon200(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(message))
}

func htmlH1(str, redirectUrl string) string {
	return fmt.Sprintf(indexHtml, str, redirectUrl)
}

func parseParam(r *http.Request) (cor createOrderReq, err error) {
	var req createOrderReq
	buf, bErr := io.ReadAll(r.Body)
	if bErr != nil {
		err = errors.New("请求体读取错误")
		return
	}
	if err = json.Unmarshal(buf, &req); err != nil {
		err = errors.New("请求体解析json错误")
		return
	}
	cor = req
	return
}

func checkParam(req createOrderReq) (err error) {
	if req.AppKey == "" {
		return errors.New("参数app_key为空")
	}

	if len(req.Rand) != 30 {
		return errors.New("参数rand不合法")
	}

	if req.Subject == "" {
		return errors.New("参数subject为空")
	}

	if _, err = strconv.Atoi(req.Price); err != nil {
		return errors.New("参数price不合法")
	}

	if req.RedirectUrl == "" {
		return errors.New("参数redirect_url为空")
	}

	if req.NotifyUrl == "" {
		return errors.New("参数notify_url为空")
	}

	if len(req.Sign) != 32 {
		return errors.New("参数sign不合法")
	}
	return
}

func checkSign(req createOrderReq) (err error) {
	var arr []string
	arr = append(arr, "app_key="+req.AppKey)
	arr = append(arr, "redirect_url="+req.RedirectUrl)
	arr = append(arr, "rand="+req.Rand)
	arr = append(arr, "subject="+req.Subject)
	arr = append(arr, "price="+req.Price)
	arr = append(arr, "notify_url="+req.NotifyUrl)
	sort.Strings(arr)
	calcSign := md5Str(strings.Join(arr, "&") + "&app_secret=" + appSecret)
	if calcSign != req.Sign {
		return errors.New("signature incorrect")
	}
	return
}

func md5Str(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func createOrderId() string {
	return md5Str(time.Now().Format(time.RFC3339Nano))
}
