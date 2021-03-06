package logic

import (
	"context"

	"zero-demo/user-rpc-sql2pb/internal/svc"
	"zero-demo/user-rpc-sql2pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserAuthLogic {
	return &SearchUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserAuthLogic) SearchUserAuth(in *pb.SearchUserAuthReq) (*pb.SearchUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserAuthResp{}, nil
}
