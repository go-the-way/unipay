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

package trongridevent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-the-way/unipay/deps/pkg"
	"github.com/go-the-way/unipay/events/apilogevent"
	"github.com/go-the-way/unipay/events/orderevent"
	"github.com/go-the-way/unipay/models"
)

var (
	apiUrl               = "https://api.shasta.trongrid.io/v1/accounts/%s/transactions/trc20?contract_address=%s&limit=%s&min_timestamp=%d"
	tokenContractAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	getReqUrl            = func(address, limit string) string {
		minTimestamp := time.Now().Add(-time.Hour).UnixMilli()
		return fmt.Sprintf(apiUrl, address, tokenContractAddress, limit, minTimestamp)
	}
)

func startReq(order *models.Order, apikey string) {
	var (
		page        = 1
		limit       = "200"
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
		reqUrl := getReqUrl(order.Other1, limit)
		req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
		req.Header = make(http.Header)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("TRON-PRO-API-KEY", apikey)
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
				if statusCode != http.StatusOK {
					apilogevent.Save(errLog(reqUrl, errors.New("状态码异常:"+string(buf)), statusCode))
				} else {
					var rm respModel
					if err = json.Unmarshal(buf, &rm); err != nil {
						apilogevent.Save(errLog(reqUrl, errors.New("反序列化响应错误："+err.Error()), statusCode))
					} else if !rm.Success {
						apilogevent.Save(errLog(reqUrl, errors.New("请求错误"), statusCode))
					} else {
						if len(rm.Data) <= 0 {
							page = 1
						} else {
							transactionTime := rm.Data[0].BlockTimestamp
							timeStamp := time.UnixMilli(transactionTime).UnixMicro()
							if timeStamp < pkg.ParseTimeUTC(order.CreateTime).UnixMicro() {
								page = 1
							} else {
								page++
							}
							if totalPage := rm.Meta.PageSize; page >= totalPage {
								page = 1
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
			// 失败次数达到最大次数，则订单失效
			order.Message = fmt.Sprintf("调用接口错误已达到%d次，订单已被取消", maxErrCount)
			orderevent.Expired(order)
			break
		}

		time.Sleep(sleepDur)
	}
}

func txnFind(order *models.Order, rm respModel) (matched bool) {
	for _, tx := range rm.Data {
		var timeStamp int64
		timeStamp = time.UnixMilli(tx.BlockTimestamp).UnixMilli() // 1739443548000
		txTime := pkg.ParseTime(pkg.FormatTime(time.UnixMilli(timeStamp)))

		decimals := tx.TokenInfo.Decimals
		orderAmount, _ := strconv.ParseFloat(order.Other2, 64)
		amount := int(orderAmount * math.Pow10(decimals)) // USD

		if tx.To == order.Other1 && order.Other2 == fmt.Sprintf("%d", amount) && order.CreateTimeBeforeTime(txTime) {
			matched = true
			order.TradeId = tx.TransactionId
			order.PayTime = pkg.FormatTime(txTime)
			break
		}
	}
	return
}

type respModel struct {
	Data []struct {
		TransactionId string `json:"transaction_id"`
		TokenInfo     struct {
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
		} `json:"token_info"`
		BlockTimestamp int64  `json:"block_timestamp"`
		From           string `json:"from"`
		To             string `json:"to"`
		Type           string `json:"type"`
		Value          string `json:"value"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At          int64  `json:"at"`
		Fingerprint string `json:"fingerprint"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
		PageSize int `json:"page_size"`
	} `json:"meta"`
}
