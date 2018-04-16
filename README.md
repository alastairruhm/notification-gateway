# notification-gateway

a gateway for multi-channel notification

## 项目背景

在目前业务系统中，消息通知有两个问题

1. 每个业务系统都自己实现了一套，未能沉淀成类库；另外有些实现甚至是阻塞的。
2. 未能正确处理发送失败、超时等问题


## 功能

1. 实现多个渠道的统一的消息推送 sms/email/bearychat/slack/企业微信 等
2. 超时、重试机制
3. 提供简单界面查看发送成功/失败的通知

## 流程

client -(1)-> server -(2)-> broker(redis/kafka) -(3)-> worker -> notify

1. 客户端通过 server 提供的 http 协议的接口发送一个 post 请求，根据不同的渠道，提供不同的参数
2. server 接收请求，验证数据，然后将通知的相关参数发送到 broker 
3. worker 通过轮询 poll 获取需要发送的通知，执行发送，结果保存在 ResultBackend