package logic

import (
	"context"

	"zero-demo/user-rpc-model/internal/svc"
	"zero-demo/user-rpc-model/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDataByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDataByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDataByIdLogic {
	return &GetUserDataByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDataByIdLogic) GetUserDataById(in *pb.GetUserDataByIdReq) (*pb.GetUserDataByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserDataByIdResp{}, nil
}
