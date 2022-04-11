package logic

import (
	"context"

	"zero-demo/user-rpc-01/internal/svc"
	"zero-demo/user-rpc-01/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserRequest) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	m := map[int64]string{
		1: "user1",
		2: "user2",
	}
	nickName := "unknown"
	if name, ok := m[in.Id] ; ok {
		nickName = name
	}

	return &pb.GetUserInfoResp{
		UserModel: &pb.UserModel{
			Id:       in.Id,
			NickName: nickName,
		},
	}, nil
}
