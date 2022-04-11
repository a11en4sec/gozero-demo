package logic

import (
	"context"

	"zero-demo/user-rpc-sql2pb/internal/svc"
	"zero-demo/user-rpc-sql2pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAuthLogic {
	return &UpdateUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserAuthLogic) UpdateUserAuth(in *pb.UpdateUserAuthReq) (*pb.UpdateUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserAuthResp{}, nil
}
