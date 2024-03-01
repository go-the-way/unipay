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

package etherscanevent

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rwscode/unipay/events/oklinkerc20event"
	"github.com/rwscode/unipay/models"
	"io"
	"net/http"
	"net/url"
	"time"
)

// curl -i "https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=0xdac17f958d2ee523a2206206994597c13d831ec7
// &address=0xdac17f958d2ee523a2206206994597c13d831ec7&startblock=0&endblock=99999999&page=1&offset=1&sort=desc
// &apikey=VD5PCBEH24K5MYIMATY91XBINHGI2YDSDD"

var (
	apiUrl = "https://xwwww.io/api?module=account&action=tokentx&contractaddress=0xdac17f958d2ee523a2206206994597c13d831ec7&address=%s&startblock=0&endblock=99999999&apikey=%s&page=%s&offset=%s&sort=desc"
	// apiUrl       = "https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=0xdac17f958d2ee523a2206206994597c13d831ec7&address=%s&startblock=0&endblock=99999999&apikey=%s&page=%s&offset=%s&sort=desc"
	getReqUrl = func(address, apikey, page, offset string) string {
		return fmt.Sprintf(apiUrl, address, apikey, page, offset)
	}
)

func startReq(order models.Order, address, apikey string) {
	var (
		page        = 1
		offset      = "500"
		errCount    = 0
		maxErrCount = 3
		sleepDur    = time.Millisecond * 200
		reqTimeout  = time.Second * 30
	)
	for {
		reqUrl := getReqUrl(address, apikey, fmt.Sprintf("%d", page), offset)
		req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
		client := &http.Client{Timeout: reqTimeout}
		resp, err := client.Do(req)
		if err != nil {
			var urlError *url.Error
			switch {
			case errors.As(err, &urlError):
				// 服务器挂了
				errCount++
			}
			// TODO: save apilog
			// apilogevent.Send(models.ApiLog{})
		} else {
			buf, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				// TODO: save apilog
				// apilogevent.Send(models.ApiLog{})
			} else {
				var rm respModel
				if err = json.Unmarshal(buf, &rm); err != nil {
					// TODO: save apilog
					// apilogevent.Send(models.ApiLog{})
				} else {
					if matched := txnQuery(order, rm); matched {
						// 找到该订单
						break
					}
				}
			}
		}
		if errCount >= maxErrCount {
			oklinkerc20event.Run()
			break
		}
		time.Sleep(sleepDur)
	}
}

func txnQuery(order models.Order, rm respModel) (matched bool) {
	fmt.Println(order)
	fmt.Println(rm)
	return
}

type respModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		BlockNumber       string `json:"blockNumber"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		From              string `json:"from"`
		ContractAddress   string `json:"contractAddress"`
		To                string `json:"to"`
		Value             string `json:"value"`
		TokenName         string `json:"tokenName"`
		TokenSymbol       string `json:"tokenSymbol"`
		TokenDecimal      string `json:"tokenDecimal"`
		TransactionIndex  string `json:"transactionIndex"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		GasUsed           string `json:"gasUsed"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		Input             string `json:"input"`
		Confirmations     string `json:"confirmations"`
	} `json:"result"`
}
