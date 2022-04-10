package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"

	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
	TestMiddleware rest.Middleware // 包含中间件
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel: model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TestMiddleware:    middleware.NewTestMiddleware().Handle, // 初始化
	}
}
