package logic

import (
	"context"

	"zero-demo/user-rpc-model/internal/svc"
	"zero-demo/user-rpc-model/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserDataLogic {
	return &SearchUserDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserDataLogic) SearchUserData(in *pb.SearchUserDataReq) (*pb.SearchUserDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserDataResp{}, nil
}
