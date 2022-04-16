
- 拦截器是对中间件的一个补充.
- 中间件和
api:  中间件(路由) -> rpc客户端拦截器
                                |
                                |
rpc:                            rpc服务端拦截器
## 1 在svc中定义rpc客户端拦截函数
```go
func MyClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 拦截前
	fmt.Println("[拦截前] 执行相关逻辑:")

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	// 拦截后
	fmt.Println("[拦截后] 执行相关逻辑:")
	return nil
}
```

## 2 在svc中使用rpc客户端拦截器
```go
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TestMiddleware: middleware.NewTestMiddleware().Handle, // 初始化
		UserRpcClient:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(MyClientInterceptor))),
	}
}
```

## 3 效果
```
➜  user-api git:(master) ✗ go run user.go 
Starting server at 0.0.0.0:8888...
----------开始使用中间件test-----------------
进入API函数.
[拦截前] 执行相关逻辑:
[拦截后] 执行相关逻辑:
----------结束使用中间件test-----------------
```