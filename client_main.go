package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	pool := x509.NewCertPool()
	caCrt, err := os.ReadFile("./ca/ca.crt")
	if err != nil {
		log.Fatal("read ca.crt file error:", err.Error())
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair("./client/cert/client.crt", "./client/cert/client.key")
	if err != nil {
		log.Fatalln("LoadX509KeyPair error:", err.Error())
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			//这里一定要注意，服务端设置ClientCAs，用于服务端验证客户端证书，客户端设置RootCAs，用户客户端验证服务端证书。设置错误或者设置反了都会造成认证不通过。
			RootCAs: pool,
			//ClientCAs:    pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8999/test")
	if err != nil {
		fmt.Printf("get failed. | err: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
