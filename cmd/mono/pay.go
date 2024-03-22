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
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getPayOrderUrl(orderId string) string {
	return fmt.Sprintf("%s/api/payOrder?order_id=%s", domainUrl, orderId)
}

func payOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("order_id")
	if orderId == "" {
		writeNon200(w, htmlH1("订单id为空", ""))
		return
	}
	mu.Lock()
	order, ok := orderMap[orderId]
	mu.Unlock()
	if !ok {
		writeNon200(w, htmlH1("订单id不存在", order.RedirectUrl))
		return
	}
	if order.Paid {
		writeNon200(w, htmlH1("订单已经支付", order.RedirectUrl))
		return
	}
	write200(w, htmlH1("订单支付成功", order.RedirectUrl))
	order.Paid = true
	// wait 3 seconds, notify pay
	go notify(order)
}

// req url: $notify_url
// req method: POST
// req content type: application/json
// req body: {"data":{"paid":true,"message":"已支付"}}
// resp content type: text/plain
// resp body: success
func notify(order *orderInfo) {
	fmt.Println("after 3 seconds, start notify pay...")
	time.Sleep(time.Second * 3)
	for {
		result := notifyReq(order)
		if result == "success" {
			break
		}
		fmt.Println(fmt.Sprintf("notify[%s] result:%v", order.OrderId, result))
		time.Sleep(time.Second)
	}
	fmt.Println(fmt.Sprintf("notify[%s] finished", order.OrderId))
}

func notifyReq(order *orderInfo) (result string) {
	resp, err := http.Post(order.NotifyUrl, "application/json", strings.NewReader(order.paidJson()))
	if err != nil {
		fmt.Println(fmt.Sprintf("notify[%s] request err:%s", order.OrderId, err.Error()))
		return
	}
	if resp == nil {
		fmt.Println(fmt.Sprintf("notify[%s] request resp nil", order.OrderId))
		return
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("notify[%s] read all err:%s", order.OrderId, err.Error()))
		return
	}
	return strings.TrimSpace(string(buf))
}
