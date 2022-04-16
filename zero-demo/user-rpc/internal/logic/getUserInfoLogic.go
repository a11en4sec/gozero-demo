package logic

import (
	"context"
	"fmt"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
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

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		fmt.Println("[从metadata中获取username的数据]", tmp)
	}

	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	fmt.Printf("[+++++++++++++++++++++]user:%v \n", user)

	return &pb.GetUserInfoResp{
		Id:       user.Id,
		Nickname: user.Nickname,
	}, nil
}
