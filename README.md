## 核心包
https://github.com/gorilla/websocket

## 扩展安装
```shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
```
## Docker 安装mongoDB
```shell
docker run -d --network some-network --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
```