## 1 在api中通过拦截器,把要传递的数据放入metadata中,塞入context中
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

func MyClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 拦截前
	fmt.Println("[拦截前] 执行相关逻辑:")

	md := metadata.New(map[string]string{"username": "jerry"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	// 拦截后
	fmt.Println("[拦截后] 执行相关逻辑:")
	return nil
}

```

## 2 在rpc中,从context中取到metadata
```go
func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		fmt.Println("[从metadata中获取username的数据]", tmp)
	}

	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	fmt.Printf("[+++++++++++++++++++++]user:%v \n", user)

	return &pb.GetUserInfoResp{
		Id:       user.Id,
		Nickname: user.Nickname,
	}, nil
}
```