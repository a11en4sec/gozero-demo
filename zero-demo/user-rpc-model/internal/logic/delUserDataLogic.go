package logic

import (
	"context"

	"zero-demo/user-rpc-model/internal/svc"
	"zero-demo/user-rpc-model/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserDataLogic {
	return &DelUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserDataLogic) DelUserData(in *pb.DelUserDataReq) (*pb.DelUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserDataResp{}, nil
}
