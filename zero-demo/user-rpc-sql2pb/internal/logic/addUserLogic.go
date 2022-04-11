package logic

import (
	"context"

	"zero-demo/user-rpc-sql2pb/internal/svc"
	"zero-demo/user-rpc-sql2pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户表-----------------------
func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserResp{
	}, nil
}
