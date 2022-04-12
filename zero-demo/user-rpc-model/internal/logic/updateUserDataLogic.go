package logic

import (
	"context"

	"zero-demo/user-rpc-model/internal/svc"
	"zero-demo/user-rpc-model/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserDataLogic {
	return &UpdateUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserDataLogic) UpdateUserData(in *pb.UpdateUserDataReq) (*pb.UpdateUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserDataResp{}, nil
}
