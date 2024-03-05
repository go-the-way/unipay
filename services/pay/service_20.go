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

package pay

import (
	"errors"
	"fmt"
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/lock"
	"github.com/rwscode/unipay/deps/pkg"
	"github.com/rwscode/unipay/events/etherscanevent"
	"github.com/rwscode/unipay/events/tronscanevent"
	"github.com/rwscode/unipay/models"
	"github.com/rwscode/unipay/services/channel"
	"github.com/rwscode/unipay/services/order"
	"strconv"
	"time"
)

func e20Run(req Req, pm channel.GetResp, orderId string) (resp Resp, err error) {
	// 0.获取可用钱包+可用金额
	address, amountYuan, amountFen, getErr := getUsableWalletAddress(pm.Type, req.AmountYuan)
	if getErr != nil {
		err = getErr
		return
	}

	var curOrder *models.Order

	// 1. 创建订单
	// 重写other1，存放钱包地址
	req.Other1 = address
	// 重写other2，存放锁定金额元
	req.Other2 = amountYuan
	// 重写other3，存放锁定金额分
	req.Other3 = amountFen
	if curOrder, err = order.Service.AddReturn(buildOrderAddReq20(pm, req, orderId)); err != nil {
		return
	}

	resp.OrderId = curOrder.Id

	// 2. 返回自定义二维码
	resp.PayPageUrl = fmt.Sprintf("%s?order_id=%s", req.E20PayPageUrl, curOrder.Id)

	// 3. 订单有效期
	orderCancelTime(curOrder)

	// 4.发送订单 开始监听状态状态
	switch curOrder.PayChannelType {
	default:
	case "erc20":
		etherscanevent.Run(curOrder)
	case "trc20":
		tronscanevent.Run(curOrder)
	}

	return
}

func orderCancelTime(order *models.Order) {
	orderValidMinute := getOrderValidMinute()
	dur, _ := time.ParseDuration(fmt.Sprintf("%dm", orderValidMinute))
	order.CancelTime = pkg.FormatTime(pkg.ParseTime(order.CreateTime).Add(dur))
}

func getOrderValidMinute() (m uint) {
	_ = db.GetDb().Model(new(models.ApiConfig)).Where("id=1").Select("valid_period_minute").Scan(&m).Error
	if m == 0 {
		m = 15
	}
	return
}

func getEnableWalletAddress(payChannelType string) (addresses []string, err error) {
	err = db.GetDb().Model(new(models.WalletAddress)).Where("state=? and protocol=?", models.WalletAddressStateEnable, payChannelType).Select("address").Find(&addresses).Error
	return
}

func getUsableWalletAddress(payChannelType string, orderAmount string) (address string, orderAmountYuan, orderAmountFen string, err error) {
	// 1. 查询启用的钱包地址
	addresses, addErr := getEnableWalletAddress(payChannelType)
	if addErr != nil {
		err = addErr
		return
	}

	if len(addresses) <= 0 {
		err = errors.New("没有可用的钱包地址")
		return
	}

	lock.RLock()
	defer lock.RUnlock()

	var (
		usable       = false
		usableAddr   = ""
		curAmount    = 0.0
		curAmountStr = ""
		decimalStr   = ""
		curLockKey   = ""
	)

loop:
	for i := 1; i <= 99; i++ {
		curAmount, _ = strconv.ParseFloat(orderAmount, 32)
		for _, addr := range addresses {
			curLockKey = fmt.Sprintf("%s-%.2f%s", addr, curAmount, decimalStr)
			if locked := lock.Have(curLockKey); !locked {
				usable = true
				usableAddr = addr
				break loop
			}
		}
		if i < 10 {
			decimalStr = fmt.Sprintf(".0%d", i)
		} else {
			decimalStr = fmt.Sprintf(".%d", i)
		}
		fv, _ := strconv.ParseFloat(decimalStr, 32)
		curAmount += fv
		curAmountStr = fmt.Sprintf("%.2f", curAmount)
	}

	if !usable {
		err = errors.New("目前USDT支付通道已满，请稍后支付")
		return
	}

	lock.SetWithLock(curLockKey)

	address = usableAddr
	orderAmountYuan = curAmountStr
	orderAmountFen = fmt.Sprintf("%d", int(curAmount*float64(100)))

	return
}
