package logic

import (
	"context"

	"zero-demo/user-rpc-model/internal/svc"
	"zero-demo/user-rpc-model/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserDataLogic {
	return &AddUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------newTable-----------------------
func (l *AddUserDataLogic) AddUserData(in *pb.AddUserDataReq) (*pb.AddUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserDataResp{}, nil
}
