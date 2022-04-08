// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	foo "zero-demo/user-api/internal/handler/foo"
	"zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: foo.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update",
				Handler: foo.UserUpdateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)
}
