package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	caCert     string = "./ca/ca.crt"
	serverCert string = "./server/cert/server.crt"
	serverKey  string = "./server/cert/server.key"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	// 客户端CA证书
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(caCert)
	if err != nil {
		fmt.Printf("load ca err: %s", err)
		return
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		fmt.Printf("certpool append ca fail.")
		return
	}

	// 可以直接用注释的代码代替最后两行
	server := &http.Server{
		Addr:    ":8999",
		Handler: router,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			//这里一定要注意，服务端设置ClientCAs，用于服务端验证客户端证书，客户端设置RootCAs，用户客户端验证服务端证书。设置错误或者设置反了都会造成认证不通过。
			//RootCAs:    certPool,
			ClientCAs: certPool,
		},
	}

	_ = server.ListenAndServeTLS(serverCert, serverKey)
}
