


## 1 将protobuf文件分开定义
- a.proto  (没有servic定义，有message定义)
- b.proto  (没有servic定义，有message定义)
- c.proto  (有service定义)

使用goctl的步骤：
1. 先将没有service定义的proto文件,生成代码
```shell
protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. a.proto
protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. b.proto

```
2. 在将包含service的proto文件，生成代码
```shell
goctl rpc protoc c.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero
```

## 2 编写业务逻辑

## 3 测试
```
grpcui --plaintext 127.0.0.1:8080
```
