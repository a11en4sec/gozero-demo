# 1 goctl 安装
```
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
```
# 2 安装protobuf
# 【注】 需要安装下面3个插件
       protoc >= 3.13.0 ， 如果没安装请先安装 https://github.com/protocolbuffers/protobuf，下载解压到$GOPATH/bin下即可，前提是$GOPATH/bin已经加入$PATH中

       protoc-gen-go ，如果没有安装请先安装 
       ```
            go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
       ```
       protoc-gen-go-grpc  ，如果没有安装请先安装
       ```  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
       ```
#       如果有要使用grpc-gateway，也请安装如下两个插件 , 没有使用就忽略下面2个插件
```
       go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
       go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
``` 
