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
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/events/apilogevent"
	"github.com/rwscode/unipay/events/oklinkevent"
	"github.com/rwscode/unipay/events/orderevent"
	"github.com/rwscode/unipay/models"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// curl -i "https://apilist.tronscanapi.com/api/transfer/trc20?sort=-timestamp&direction=2&db_version=1&trc20Id=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&address=TU8fjcJFpgGd2q9roMBmv5c9wo7q2Pwt2d&start=0&limit=1"

var (
	apiUrl    = "https://apilist.tronscanapi.com/api/transfer/trc20?sort=-timestamp&direction=2&db_version=1&trc20Id=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&address=%s&start=%s&limit=%s"
	getReqUrl = func(address, start, limit string) string {
		return fmt.Sprintf(apiUrl, address, start, limit)
	}
)

func startReq(order *models.Order) {
	var (
		start       = 1
		limit       = "500"
		errCount    = 0
		maxErrCount = 3
		sleepDur    = time.Millisecond * 200
		reqTimeout  = time.Second * 30
		client      = &http.Client{Timeout: reqTimeout}
	)
	errLog := func(reqUrl string, err error, statusCode int) *models.ApiLog {
		return models.NewApiLogGetNoParam(reqUrl, err.Error(), fmt.Sprintf("%d", statusCode))
	}
	for {
		reqUrl := getReqUrl(order.Other1, fmt.Sprintf("%d", start), limit)
		req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
		resp, err := client.Do(req)
		if err != nil {
			var urlError *url.Error
			switch {
			case errors.As(err, &urlError):
				// 服务器挂了
				errCount++
			}
			apilogevent.Save(errLog(reqUrl, err, resp.StatusCode))
		} else {
			buf, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				apilogevent.Save(errLog(reqUrl, errors.New("读取相应错误："+err.Error()), resp.StatusCode))
			} else {
				var rm respModel
				if err = json.Unmarshal(buf, &rm); err != nil {
					apilogevent.Save(errLog(reqUrl, errors.New("反序列化相应错误："+err.Error()), resp.StatusCode))
				} else {
					if matched := txnFind(order, rm); matched {
						// 找到该订单
						break
					}
				}
			}
		}
		if order.CancelTimeBeforeNow() {
			orderevent.Expired(order)
			break
		}
		if errCount >= maxErrCount {
			oklinkevent.Run(order)
			break
		}
		time.Sleep(sleepDur)
	}
}

func txnFind(order *models.Order, rm respModel) (matched bool) {
	for _, tx := range rm.Data {
		// "decimals": 6,
		decimals := tx.Decimals
		orderAmount, _ := strconv.Atoi(order.AmountYuan)
		// USD
		amount := orderAmount * decimals
		// "block_timestamp": 1709278716000,
		txTimeStamp := tx.BlockTimestamp
		txTime := time.UnixMicro(txTimeStamp)
		if tx.To == order.Other1 && fmt.Sprintf("%d", amount) == tx.Amount && tx.ContractRet == "SUCCESS" && order.CreateTimeBeforeTime(txTime) {
			matched = true
			order.PayTime = pkg.FormatTime(txTime)
			break
		}
	}
	return
}

type respModel struct {
	ContractMap struct {
		// THPvaUhoh2Qn2Y9THCZML3H815HhFhn5YC bool `json:"THPvaUhoh2Qn2y9THCZML3H815hhFhn5YC"`
		// TSaRZDiBPD8Rd5VrvX8A4ZgunHczM9Mj8S bool `json:"TSaRZDiBPD8Rd5vrvX8a4zgunHczM9mj8S"`
		// TU8FjcJFpgGd2Q9RoMBmv5C9Wo7Q2Pwt2D bool `json:"TU8fjcJFpgGd2q9roMBmv5c9wo7q2Pwt2d"`
		// } `json:"contractMap"`
	} `json:"-"`
	TokenInfo struct {
		TokenId      string `json:"tokenId"`
		TokenAbbr    string `json:"tokenAbbr"`
		TokenName    string `json:"tokenName"`
		TokenDecimal int    `json:"tokenDecimal"`
		TokenCanShow int    `json:"tokenCanShow"`
		TokenType    string `json:"tokenType"`
		TokenLogo    string `json:"tokenLogo"`
		TokenLevel   string `json:"tokenLevel"`
		IssuerAddr   string `json:"issuerAddr"`
		Vip          bool   `json:"vip"`
	} `json:"tokenInfo"`
	PageSize int `json:"page_size"`
	Code     int `json:"code"`
	Data     []struct {
		Amount         string `json:"amount"`
		Status         int    `json:"status"`
		ApprovalAmount string `json:"approval_amount"`
		BlockTimestamp int64  `json:"block_timestamp"`
		Block          int    `json:"block"`
		From           string `json:"from"`
		To             string `json:"to"`
		Hash           string `json:"hash"`
		Confirmed      int    `json:"confirmed"`
		ContractType   string `json:"contract_type"`
		ContractType1  int    `json:"contractType"`
		Revert         int    `json:"revert"`
		ContractRet    string `json:"contract_ret"`
		EventType      string `json:"event_type"`
		IssueAddress   string `json:"issue_address"`
		Decimals       int    `json:"decimals"`
		TokenName      string `json:"token_name"`
		Id             string `json:"id"`
		Direction      int    `json:"direction"`
	} `json:"data"`
}
