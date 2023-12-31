# gin-mutual-trust
gin-mutual-trust

## 说明
基于 Gin 实现双向证书认证的示例

## 证书部份
bash init.sh

生成证书


## 启动 server

```
gin-mutual-trust git:(main) ✗ go run server_main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /test                     --> main.main.func1 (2 handlers)
2023/12/30 21:05:01 http: TLS handshake error from 127.0.0.1:54106: remote error: tls: bad certificate
[GIN] 2023/12/30 - 21:05:18 | 200 |      56.375µs |             ::1 | GET      "/test"


```

## 测试 client 

```
➜  gin-mutual-trust git:(main) ✗ go run client_main.go
{"message":"success"}
```