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
	reqUrl := fmt.Sprintf("%s/api/createOrder", domainUrl)
	fmt.Println(`
The unipay channel mono served on ` + serverAddr + `

channel info
---
app_key: ` + appKey + `
app_secret: ` + appSecret + `
signature: md5(join(sort(["app_key="+$Param.app_key,"rand="+$Param.rand,"subject="+$Param.subject,"price="+$Param.price,"notify_url="+$Param.notify_url,"redirect_url="+$Param.redirect_url]),"&")+"&app_secret="+$Param.app_secret)

insert sql info
---
INSERT INTO unipay_channels (id, name,type, admin_url, admin_user, admin_passwd, logo_url, amount_type, amount_validate_cond, req_url, req_method, req_content_type, req_success_expr, req_pay_page_url_expr, req_pay_qr_url_expr, req_pay_message_expr, notify_pay_content_type, notify_pay_success_expr, notify_pay_id_expr, notify_pay_return_content, notify_pay_return_content_type, state, sort, remark, create_time, update_time) VALUES (1, 'Mono Payment','normal', 'N/A', 'N/A', 'N/A', 'https://cdn.dribbble.com/users/12504221/screenshots/19846170/media/cbaf9486d45481a92d88d37960f04014.png', 1, '', '` + reqUrl + `', 'POST', 'json', '$code=="200"', '$data.pay_url', '', '$message', 'json', '$data.paid', '$data.order_id', 'success', 'text', 1, 0, 'Mono Payment', '2024-02-03 10:49:00', '2024-02-03 10:49:00');
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (1, 1, 'app_key', '` + appKey + `', 'app_key', 1);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (2, 1, 'app_secret', '` + appSecret + `', 'app_secret', 2);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (3, 1, 'rand', '$rand_str(30)', 'rand str 30', 1);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (4, 1, 'subject', '$Pay.Subject', 'subject', 1);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (5, 1, 'price', '$Pay.AmountYuan', 'price', 1);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (6, 1, 'notify_url', '$Pay.NotifyUrl', 'notify_url', 1);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (7, 1, 'redirect_url', '$Pay.AppWakeUri', 'redirect_url', 1);
INSERT INTO unipay_channel_params (id, channel_id, name, value, remark, pass) VALUES (8, 1, 'sign', 'md5(join(sort(["app_key="+$Param.app_key,"rand="+$Param.rand,"subject="+$Param.subject,"price="+$Param.price,"notify_url="+$Param.notify_url,"redirect_url="+$Param.redirect_url]),"&")+"&app_secret="+$Param.app_secret)', 'md5 signature', 1);

create order info
---
req url: ` + reqUrl + `
req method: POST
req content type: application/json
req body: {"app_key":"app key","rand":"rand str 30 len","subject":"pay subject","price":"100","notify_url":"notify url","redirect_url":"redirect url","sign": "signature md5 str"}
resp content type: application/json
resp body: {"code":"200","data":{"order_id":"order id","pay_url":"http://example.com/api/payOrder?order_id=wow"},"message":"ok"}

notify order info
---
req url: $notify_url
req method: POST
req content type: application/json
req body: {"data":{"order_id":"order id","paid":true,"message":"paid"}}
resp content type: text/plain
resp body: success


envs
---
SERVER_ADDR mono server addr (default: :9988)
APP_KEY mono app key (default: BmnXsm843uA9WjWh22CWIXbrASo)
APP_SECRET mono app secret (default: Ne4WZgphE1GicyYgQAYn0ZqhwvA)
DOMAIN_URL mono server domain url (default: http://publicIp:9988)
`)
}

func serve() { fmt.Println(http.ListenAndServe(serverAddr, nil)) }
