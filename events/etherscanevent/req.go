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
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/events/apilogevent"
	"github.com/rwscode/unipay/events/oklinkevent"
	"github.com/rwscode/unipay/events/orderevent"
	"github.com/rwscode/unipay/models"
)

// curl -i "https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=0xdac17f958d2ee523a2206206994597c13d831ec7&address=0xdac17f958d2ee523a2206206994597c13d831ec7&startblock=0&endblock=99999999&page=1&offset=2&sort=desc&apikey=VD5PCBEH24K5MYIMATY91XBINHGI2YDSDD"

var (
	apiUrl    = "https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=0xdac17f958d2ee523a2206206994597c13d831ec7&address=%s&startblock=0&endblock=99999999&apikey=%s&page=%s&offset=%s&sort=desc"
	getReqUrl = func(address, apikey, page, offset string) string {
		return fmt.Sprintf(apiUrl, address, apikey, page, offset)
	}
)

func startReq(order *models.Order, apikey string) {
	var (
		page        = 1
		offset      = "500"
		errCount    = 0
		maxErrCount = 3
		sleepDur    = time.Second
		reqTimeout  = time.Second * 3
		client      = &http.Client{Timeout: reqTimeout}
	)
	errLog := func(reqUrl string, err error, statusCode int) *models.ApiLog {
		errCount++
		return models.NewApiLogGetNoParam(reqUrl, err.Error(), fmt.Sprintf("%d", statusCode))
	}
	for {
		reqUrl := getReqUrl(order.Other1, apikey, fmt.Sprintf("%d", page), offset)
		req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
		resp, err := client.Do(req)
		var statusCode int
		if err != nil {
			apilogevent.Save(errLog(reqUrl, err, statusCode))
		} else if resp == nil {
			apilogevent.Save(errLog(reqUrl, errors.New("无响应"), statusCode))
		} else {
			statusCode = resp.StatusCode
			buf, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				apilogevent.Save(errLog(reqUrl, errors.New("读取响应错误"+err.Error()), statusCode))
			} else if len(buf) <= 0 {
				apilogevent.Save(errLog(reqUrl, errors.New("读取响应为空"), statusCode))
			} else {
				if resp.StatusCode != http.StatusOK {
					apilogevent.Save(errLog(reqUrl, errors.New("状态码异常:"+string(buf)), statusCode))
				} else {
					var rm respModel
					if err = json.Unmarshal(buf, &rm); err != nil {
						apilogevent.Save(errLog(reqUrl, errors.New("反序列化响应错误："+err.Error()), statusCode))
					} else {
						if len(rm.Result) <= 0 {
							page = 1
						} else {
							timeStamp, _ := strconv.ParseInt(rm.Result[0].TimeStamp, 10, 64)
							if timeStamp < pkg.ParseTimeUTC(order.CreateTime).UnixMilli() {
								page = 1
							} else {
								page++
							}
						}
						if matched := txnFind(order, rm); matched {
							// 找到该订单
							orderevent.Paid(order)
							break
						}
					}
				}
			}
		}

		if order.CancelTimeBeforeNow() {
			// 订单失效
			order.Message = "订单超时已被取消"
			orderevent.Expired(order)
			break
		}

		if errCount >= maxErrCount {
			// 切换到oklink
			oklinkevent.Run(order)
			break
		}

		time.Sleep(sleepDur)
	}
}

func txnFind(order *models.Order, rm respModel) (matched bool) {
	for _, tx := range rm.Result {
		// "tokenDecimal": "6",
		tokenDecimal, _ := strconv.Atoi(tx.TokenDecimal)
		orderAmount, _ := strconv.ParseFloat(order.Other2, 64)
		// USD
		amount := int(orderAmount * math.Pow10(tokenDecimal))
		timeStamp, _ := strconv.ParseInt(tx.TimeStamp, 10, 64)
		// "timeStamp": "1535035994",
		txTime := pkg.ParseTime(pkg.FormatTime(time.Unix(timeStamp, 0)))
		if tx.To == order.Other1 && fmt.Sprintf("%d", amount) == tx.Value && order.CreateTimeBeforeTime(txTime) {
			matched = true
			order.PayTime = pkg.FormatTime(txTime)
			break
		}
	}
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
