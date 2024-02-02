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
- order.PaySuccess 支付成功
- order.PayFailure 支付失败
- order.GetPayState 查询支付状态

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
- SERVER_ADDR 服务运行绑定地址 (default: :9988)
- APP_KEY mono app key (default: BmnXsm843uA9WjWh22CWIXbrASo)
- APP_SECRET mono app secret (default: Ne4WZgphE1GicyYgQAYn0ZqhwvA)
- DOMAIN_URL mono service domain url (default: http://publicIp:9988)