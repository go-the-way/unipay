# unipay
The unified payment package

# Services

## Pay 支付和回调
- pay.ReqPay 请求支付接口
- pay.NotifyPay 回调支付接口

## Channel 支付通道
- channel.GetPage 分页查询
- channel.Get 查询详情
- channel.Add 新增
- channel.Update 修改
- channel.Del 删除
- channel.Enable 启用
- channel.Disable 禁用
- channel.GetMatches 查询匹配的支付通道列表

## Channel Param 支付通道参数
- channelparam.Get 查询详情
- channelparam.GetChannelId 查询支付通道参数列表
- channelparam.GetName 查询支付通道+名称的支付参数
- channelparam.Add 新增
- channelparam.Update 修改
- channelparam.Del 删除

## Order 支付订单
- order.GetPage 分页查询
- order.Get 查询详情
- order.GetBusinessId 查询业务id
- order.GetIdAndBusinessId 查询订单id+业务id
- order.Add 新增
- order.Update 修改
- order.Del 删除
- order.Paid 支付成功
- order.Cancel 取消
- order.GetState 查询状态

# Models
- models.Channel/models.UnipayChannel 支付通道
- models.ChannelParam/models.UnipayChannelParam 支付通道参数
- models.Order/models.UnipayOrder 支付订单

# Mono 模拟支付通道

## install
```
go install github.com/rwscode/unipay/cmd/mono@latest
```

## run
```
mono
```

## envs
- SERVER_ADDR mono server addr (default: :9988)
- APP_KEY mono app key (default: BmnXsm843uA9WjWh22CWIXbrASo)
- APP_SECRET mono app secret (default: Ne4WZgphE1GicyYgQAYn0ZqhwvA)
- DOMAIN_URL mono server domain url (default: http://publicIp:9988)


Tether USD(USDT) TRC20 on [TRON](https://www.oklink.com/cn/trx)
---
TRON 是 Tron 基金会于 2017 年推出的基于区块链技术的去中心化操作系统，旨在建立一个免费的全球数字内容娱乐系统。TRON 支持智能合约且兼容 EVM 的特性使开发者可以方便快捷地在 TRON 上部署智能合约和构建 DApp。TRON 依靠代理权益证明 (DPoS) 共识机制为协议上的 DApp 提供高吞吐、易扩展、可靠的底层公链支持。
```
合约地址：TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t
https://www.oklink.com/cn/trx/token/TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t
```

Tether USD(USDT) ERC20 on [Ethereum](https://www.oklink.com/cn/eth)
---
以太坊 (Ethereum) 是一个开源、去中心化，支持智能合约的区块链网络。以太坊最初采用工作量证明 (PoW) 的共识机制，2022 年 9 月 15 日于区块高度 15537394 切换为权益证明机制 (PoS)。ETH 是以太坊网络的原生代币，是用户在以太坊网络交互的必需品，亦可以质押 ETH 成为验证者并通过验证区块等行为来维护网络安全。
```
合约地址：0xdac17f958d2ee523a2206206994597c13d831ec7
https://www.oklink.com/cn/eth/token/0xdac17f958d2ee523a2206206994597c13d831ec7
```
