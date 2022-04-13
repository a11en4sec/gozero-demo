package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
	"zero-demo/user-rpc/usercenter"

	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
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
		UserRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
