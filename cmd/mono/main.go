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
	"net/http"
)

func main() {
	routers()
	showMessage()
	serve()
}

func routers() {
	http.HandleFunc("/api/createOrder", createOrder)
	http.HandleFunc("/api/payOrder", payOrder)
}

func showMessage() {
	reqUrl := fmt.Sprintf("req url: %s/api/createOrder", domainUrl)
	fmt.Println(`
The unipay channel mono served on ` + serverAddr + `

channel info
---
appKey: ` + appKey + `
appSecret: ` + appSecret + `
signature: md5(concat(join(sort["app_key="+$appKey,"rand="+$rand,"subject="+$subject,"price="+$price,"notify_url="+$notify_url],"&"),"&appSecret=",$appSecret)

create order info
---
req url: ` + reqUrl + `
req method: POST
req content type: application/json
req body: {"app_key":"app key","rand":"rand str 30 len","subject":"pay subject","price":"100","notify_url":"notify url","sign": "signature md5 str"}
resp content type: application/json
resp body: {"code":"200",data":{"order_id":"order id",pay_url":"http://example.com/pay/1"},"message":"ok"}

notify order info
---
req url: $notify_url
req method: POST
req content type: application/json
req body: {"data":{"order_id":"order id",paid":true,"message":"paid"}}
resp content type: text/plain
resp body: success
`)
}

func serve() { fmt.Println(http.ListenAndServe(serverAddr, nil)) }
