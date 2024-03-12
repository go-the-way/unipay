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
	"strconv"
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

// curl "https://www.oklink.com/api/v5/explorer/address/transaction-list?chainShortName=ETH&tokenContractAddress=0xdac17f958d2ee523a2206206994597c13d831ec7&address=0xEf8801eaf234ff82801821FFe2d78D60a0237F97&protocolType=token_20&page=1&limit=2" -H 'Ok-Access-Key: 523ded74-82c2-420b-b00e-687bc7e3a139'

func startReq(order *models.Order, apikey string, chainShortName string) {
	var (
		page        = 1
		limit       = "100"
		errCount    = 0
		maxErrCount = 3
		sleepDur    = time.Millisecond * 200
		reqTimeout  = time.Second * 3
		client      = &http.Client{Timeout: reqTimeout}
	)
	errLog := func(reqUrl string, err error, statusCode int) *models.ApiLog {
		errCount++
		return models.NewApiLogGetNoParam(reqUrl, err.Error(), fmt.Sprintf("%d", statusCode))
	}
	for {
		reqUrl := getReqUrl(chainShortName, order.Other1, fmt.Sprintf("%d", page), limit)
		req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
		req.Header = make(http.Header)
		req.Header.Set("Ok-Access-Key", apikey)
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
				var rm respModel
				if err = json.Unmarshal(buf, &rm); err != nil {
					apilogevent.Save(errLog(reqUrl, errors.New("反序列化响应错误："+err.Error()), statusCode))
				} else if rm.Data == nil || len(rm.Data) <= 0 {
					apilogevent.Save(errLog(reqUrl, errors.New("请求错误："+rm.Msg), statusCode))
				} else {
					d := rm.Data[0]
					if len(d.TransactionLists) <= 0 {
						page = 1
					} else {
						timeStamp := pkg.FromUnix(d.TransactionLists[0].TransactionTime).UnixMicro()
						if timeStamp < pkg.ParseTime(order.CreateTime).UnixMicro() {
							// FIXME
							// 尝试获取时间，和订单时间进行比较，在订单创建时间之前，则进行下一轮
							page = 1
						} else {
							page++
							if totalPage, _ := strconv.Atoi(d.TotalPage); page >= totalPage {
								page = 1
							}
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

		if order.CancelTimeBeforeNow() {
			// 订单失效
			order.Message = "订单超时已被取消"
			orderevent.Expired(order)
			break
		}

		if errCount >= maxErrCount {
			// 失败次数达到最大次数，则订单失效
			order.Message = fmt.Sprintf("调用接口错误已达到%d次，订单已被取消", maxErrCount)
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
