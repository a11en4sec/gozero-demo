package svc

import (
	"context"
	"fmt"
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
	"zero-demo/user-rpc/usercenter"

	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
	TestMiddleware rest.Middleware // 包含中间件
	UserRpcClient  usercenter.Usercenter
}

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
