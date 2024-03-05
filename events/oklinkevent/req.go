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

package oklinkevent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/events/apilogevent"
	"github.com/rwscode/unipay/events/orderevent"
	"github.com/rwscode/unipay/models"
)

const (
	eth = "ETH"
	trx = "TRON"
)

var (
	tm        = map[string]string{"erc20": eth, "trc20": trx}
	apiUrl    = "https://www.oklink.com/api/v5/explorer/address/transaction-list?chainShortName=%s&tokenContractAddress=%s&address=%s&protocolType=token_20&page=%s&limit=%s"
	getReqUrl = func(chainShortName, address, page, limit string) string {
		var tokenContractAddress string
		switch chainShortName {
		case eth:
			tokenContractAddress = "0xdac17f958d2ee523a2206206994597c13d831ec7"
		case trx:
			tokenContractAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
		}
		return fmt.Sprintf(apiUrl, chainShortName, tokenContractAddress, address, page, limit)
	}
)

func startReq(order *models.Order, apikey string, chainShortName string) {
	var (
		page       = 1
		limit      = "100"
		sleepDur   = time.Millisecond * 200
		reqTimeout = time.Second * 30
		client     = &http.Client{Timeout: reqTimeout}
	)
	errLog := func(reqUrl string, err error, statusCode int) *models.ApiLog {
		return models.NewApiLogGetNoParam(reqUrl, err.Error(), fmt.Sprintf("%d", statusCode))
	}
	for {
		reqUrl := getReqUrl(chainShortName, order.Other1, fmt.Sprintf("%d", page), limit)
		req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
		req.Header = make(http.Header)
		req.Header.Set("Ok-Access-Key", apikey)
		resp, err := client.Do(req)
		if err != nil {
			var urlError *url.Error
			switch {
			case errors.As(err, &urlError):
				// 服务器挂了
			}
			apilogevent.Save(errLog(reqUrl, err, resp.StatusCode))
		} else {
			if resp == nil {
				apilogevent.Save(errLog(reqUrl, errors.New("读取相应为空"), resp.StatusCode))
			} else {
				buf, readErr := io.ReadAll(resp.Body)
				if readErr != nil {
					apilogevent.Save(errLog(reqUrl, errors.New("读取响应错误："+err.Error()), resp.StatusCode))
				} else {
					var rm respModel
					if err = json.Unmarshal(buf, &rm); err != nil {
						apilogevent.Save(errLog(reqUrl, errors.New("反序列化响应错误："+err.Error()), resp.StatusCode))
					} else if rm.Data != nil && len(rm.Data) > 0 {
						d := rm.Data[0]

						if len(d.TransactionLists) > 0 {
							page++
						}

						if len(d.TransactionLists) <= 0 {
							page = 1
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

		if time.Now().Before(pkg.ParseTime(order.CancelTime)) {
			// 订单已失效
			orderevent.Expired(order)
			break
		}

		time.Sleep(sleepDur)
	}
}

func txnFind(order *models.Order, rm respModel) (matched bool) {
	for _, tx := range rm.Data[0].TransactionLists {
		// "transactionTime": "1709277731000",  eth
		// "transactionTime": "1709277731",     trx
		txTime := pkg.FromUnix(tx.TransactionTime)
		if tx.To == order.Other1 && tx.State == "success" && order.Other2 == tx.Amount && order.CreateTimeBeforeTime(txTime) {
			matched = true
			order.PayTime = pkg.FormatTime(txTime)
			break
		}
	}
	return
}

type respModel struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Page             string `json:"page"`
		Limit            string `json:"limit"`
		TotalPage        string `json:"totalPage"`
		ChainFullName    string `json:"chainFullName"`
		ChainShortName   string `json:"chainShortName"`
		TransactionLists []struct {
			TxId                 string `json:"txId"`
			MethodId             string `json:"methodId"`
			BlockHash            string `json:"blockHash"`
			Height               string `json:"height"`
			TransactionTime      string `json:"transactionTime"`
			From                 string `json:"from"`
			To                   string `json:"to"`
			IsFromContract       bool   `json:"isFromContract"`
			IsToContract         bool   `json:"isToContract"`
			Amount               string `json:"amount"`
			TransactionSymbol    string `json:"transactionSymbol"`
			TxFee                string `json:"txFee"`
			State                string `json:"state"`
			TokenId              string `json:"tokenId"`
			TokenContractAddress string `json:"tokenContractAddress"`
			ChallengeStatus      string `json:"challengeStatus"`
			L1OriginHash         string `json:"l1OriginHash"`
		} `json:"transactionLists"`
	} `json:"data"`
}
