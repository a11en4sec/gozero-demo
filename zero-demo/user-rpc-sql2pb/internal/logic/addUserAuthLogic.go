package logic

import (
	"context"

	"zero-demo/user-rpc-sql2pb/internal/svc"
	"zero-demo/user-rpc-sql2pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAuthLogic {
	return &AddUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户授权表-----------------------
func (l *AddUserAuthLogic) AddUserAuth(in *pb.AddUserAuthReq) (*pb.AddUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserAuthResp{}, nil
}
