package logic

import (
	"context"

	"zero-demo/user-rpc-model/internal/svc"
	"zero-demo/user-rpc-model/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	
	user, err := l.svcCtx.UserModel.FindOne(l.ctx,in.Id)

	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIdResp{
		User: &pb.User{
			Id:         user.Id,
			Nickname:   user.Nickname,
		},
	}, nil
}
