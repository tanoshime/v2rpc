## 这是什么？

一个把 v2ray gRPC API 转成 HTTP 接口的程序

## 目前已实现

- [x] 添加用户
- [x] 删除用户
- [x] 查询流量

## 使用方法

- 跑起来就行。本程序监听本地 12580 端口

## 接口列表

| Path | Method | Params  |
| --- | --- | --- |
| `/api/:target/stats` | `GET` | query: { pattern: string, reset: boolean } |
| `/api/:target/stats/inbound/:tag/:traffic-type` | `GET` | query: { reset: boolean } |
| `/api/:target/stats/user/:email/:traffic-type` | `GET` | query: { reset: boolean } |
| `/api/:target/:tag/user` | `POST` | body: { email, alterId, id, level } |
| `/api/:target/:tag/user/:email` | `DELETE` |

- `target`: v2ray gRPC 地址，如：`127.0.0.1:10086`
- `tag`: inbound 设置的 tag 名称
- `traffic-type`: `uplink` 或者 `downlink`