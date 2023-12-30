################### CA ###############
# 生成根证书的私钥
openssl genrsa -out ./ca/ca.key 4096

# 生成根证书签发申请文件(csr文件)
openssl req -new -sha256 -out ./ca/ca.csr -key ./ca/ca.key -config ./ca/ca.conf

# 自签发根证书(cer文件)
openssl x509 -req -days 3650 -in ./ca/ca.csr -signkey ./ca/ca.key -out ./ca/ca.crt


############### server ###############
# 生成服务端私钥
openssl genrsa -out ./server/cert/server.key 2048

# 生成服务端证书申请文件
openssl req -new -sha256 -out ./server/cert/server.csr -key ./server/cert/server.key -config ./server/server.conf

# 用CA证书签发服务端证书
openssl x509 -req -days 3650 -CA ./ca/ca.crt -CAkey ./ca/ca.key -CAcreateserial -in ./server/cert/server.csr -out ./server/cert/server.crt -extensions req_ext -extfile ./server/server.conf


############### client ###############

# 生成客户端私钥
openssl genrsa -out ./client/cert/client.key 2048

# 生成客户端证书申请文件
openssl req -new -sha256 -out ./client/cert/client.csr -key ./client/cert/client.key -config ./client/client.conf

# 用跟证书签发客户端证书
openssl x509 -req -days 3650 -CA ./ca/ca.crt -CAkey ./ca/ca.key -in ./client/cert/client.csr -out ./client/cert/client.crt -extensions req_ext -extfile ./client/client.conf