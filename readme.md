## 这是什么？

一个把 v2ray gRPC API 转成 HTTP 接口的程序

## 为什么做这个？

还是 HTTP 好用啊， gRPC 啥的，很不好对接诶。

## 目前已实现

- [x] 添加用户
- [x] 删除用户
- [x] 查询流量

## 使用方法

- 跑起来就行。本程序运行于 http://127.0.0.1:12580

## 接口列表

| Path | Method | Params  |
| --- | --- | --- |
| `/api/:target/stats` | `GET` | query: { pattern: string, reset: boolean } |
| `/api/:target/stats/inbound/:tag/:traffic-type` | `GET` | query: { reset: boolean } |
| `/api/:target/stats/user/:email/:traffic-type` | `GET` | query: { reset: boolean } |
| `/api/:target/:tag/user` | `POST` | json body: { email, alterId, id, level } |
| `/api/:target/:tag/user/:email` | `DELETE` |

### path 参数说明

- `target`: v2ray gRPC 地址，如：`127.0.0.1:10086`
- `tag`: inbound 设置的 tag 名称
- `traffic-type`: `uplink` 或者 `downlink`

