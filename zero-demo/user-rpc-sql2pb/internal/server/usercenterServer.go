// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package server

import (
	"context"

	"zero-demo/user-rpc-sql2pb/internal/logic"
	"zero-demo/user-rpc-sql2pb/internal/svc"
	"zero-demo/user-rpc-sql2pb/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

// -----------------------用户表-----------------------
func (s *UsercenterServer) AddUser(ctx context.Context, in *pb.AddUserReq) (*pb.AddUserResp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UsercenterServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UsercenterServer) DelUser(ctx context.Context, in *pb.DelUserReq) (*pb.DelUserResp, error) {
	l := logic.NewDelUserLogic(ctx, s.svcCtx)
	return l.DelUser(in)
}

func (s *UsercenterServer) GetUserById(ctx context.Context, in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UsercenterServer) SearchUser(ctx context.Context, in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	l := logic.NewSearchUserLogic(ctx, s.svcCtx)
	return l.SearchUser(in)
}

// -----------------------用户授权表-----------------------
func (s *UsercenterServer) AddUserAuth(ctx context.Context, in *pb.AddUserAuthReq) (*pb.AddUserAuthResp, error) {
	l := logic.NewAddUserAuthLogic(ctx, s.svcCtx)
	return l.AddUserAuth(in)
}

func (s *UsercenterServer) UpdateUserAuth(ctx context.Context, in *pb.UpdateUserAuthReq) (*pb.UpdateUserAuthResp, error) {
	l := logic.NewUpdateUserAuthLogic(ctx, s.svcCtx)
	return l.UpdateUserAuth(in)
}

func (s *UsercenterServer) DelUserAuth(ctx context.Context, in *pb.DelUserAuthReq) (*pb.DelUserAuthResp, error) {
	l := logic.NewDelUserAuthLogic(ctx, s.svcCtx)
	return l.DelUserAuth(in)
}

func (s *UsercenterServer) GetUserAuthById(ctx context.Context, in *pb.GetUserAuthByIdReq) (*pb.GetUserAuthByIdResp, error) {
	l := logic.NewGetUserAuthByIdLogic(ctx, s.svcCtx)
	return l.GetUserAuthById(in)
}

func (s *UsercenterServer) SearchUserAuth(ctx context.Context, in *pb.SearchUserAuthReq) (*pb.SearchUserAuthResp, error) {
	l := logic.NewSearchUserAuthLogic(ctx, s.svcCtx)
	return l.SearchUserAuth(in)
}