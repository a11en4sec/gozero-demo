package logic

import (
	"context"

	"zero-demo/user-rpc-sql2pb/internal/svc"
	"zero-demo/user-rpc-sql2pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByIdLogic {
	return &GetUserAuthByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByIdLogic) GetUserAuthById(in *pb.GetUserAuthByIdReq) (*pb.GetUserAuthByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAuthByIdResp{}, nil
}
