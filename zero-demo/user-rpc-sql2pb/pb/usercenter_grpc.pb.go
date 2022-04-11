// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: usercenter.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	//-----------------------用户表-----------------------
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
	SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error)
	//-----------------------用户授权表-----------------------
	AddUserAuth(ctx context.Context, in *AddUserAuthReq, opts ...grpc.CallOption) (*AddUserAuthResp, error)
	UpdateUserAuth(ctx context.Context, in *UpdateUserAuthReq, opts ...grpc.CallOption) (*UpdateUserAuthResp, error)
	DelUserAuth(ctx context.Context, in *DelUserAuthReq, opts ...grpc.CallOption) (*DelUserAuthResp, error)
	GetUserAuthById(ctx context.Context, in *GetUserAuthByIdReq, opts ...grpc.CallOption) (*GetUserAuthByIdResp, error)
	SearchUserAuth(ctx context.Context, in *SearchUserAuthReq, opts ...grpc.CallOption) (*SearchUserAuthResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	out := new(AddUserResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	out := new(DelUserResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/DelUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	out := new(GetUserByIdResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error) {
	out := new(SearchUserResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) AddUserAuth(ctx context.Context, in *AddUserAuthReq, opts ...grpc.CallOption) (*AddUserAuthResp, error) {
	out := new(AddUserAuthResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/AddUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserAuth(ctx context.Context, in *UpdateUserAuthReq, opts ...grpc.CallOption) (*UpdateUserAuthResp, error) {
	out := new(UpdateUserAuthResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/UpdateUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DelUserAuth(ctx context.Context, in *DelUserAuthReq, opts ...grpc.CallOption) (*DelUserAuthResp, error) {
	out := new(DelUserAuthResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/DelUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserAuthById(ctx context.Context, in *GetUserAuthByIdReq, opts ...grpc.CallOption) (*GetUserAuthByIdResp, error) {
	out := new(GetUserAuthByIdResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/GetUserAuthById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SearchUserAuth(ctx context.Context, in *SearchUserAuthReq, opts ...grpc.CallOption) (*SearchUserAuthResp, error) {
	out := new(SearchUserAuthResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/SearchUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	//-----------------------用户表-----------------------
	AddUser(context.Context, *AddUserReq) (*AddUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	DelUser(context.Context, *DelUserReq) (*DelUserResp, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	SearchUser(context.Context, *SearchUserReq) (*SearchUserResp, error)
	//-----------------------用户授权表-----------------------
	AddUserAuth(context.Context, *AddUserAuthReq) (*AddUserAuthResp, error)
	UpdateUserAuth(context.Context, *UpdateUserAuthReq) (*UpdateUserAuthResp, error)
	DelUserAuth(context.Context, *DelUserAuthReq) (*DelUserAuthResp, error)
	GetUserAuthById(context.Context, *GetUserAuthByIdReq) (*GetUserAuthByIdResp, error)
	SearchUserAuth(context.Context, *SearchUserAuthReq) (*SearchUserAuthResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) AddUser(context.Context, *AddUserReq) (*AddUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUsercenterServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsercenterServer) DelUser(context.Context, *DelUserReq) (*DelUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUser not implemented")
}
func (UnimplementedUsercenterServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsercenterServer) SearchUser(context.Context, *SearchUserReq) (*SearchUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedUsercenterServer) AddUserAuth(context.Context, *AddUserAuthReq) (*AddUserAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAuth not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserAuth(context.Context, *UpdateUserAuthReq) (*UpdateUserAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAuth not implemented")
}
func (UnimplementedUsercenterServer) DelUserAuth(context.Context, *DelUserAuthReq) (*DelUserAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserAuth not implemented")
}
func (UnimplementedUsercenterServer) GetUserAuthById(context.Context, *GetUserAuthByIdReq) (*GetUserAuthByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthById not implemented")
}
func (UnimplementedUsercenterServer) SearchUserAuth(context.Context, *SearchUserAuthReq) (*SearchUserAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserAuth not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DelUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DelUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/DelUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DelUser(ctx, req.(*DelUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SearchUser(ctx, req.(*SearchUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_AddUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/AddUserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUserAuth(ctx, req.(*AddUserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/UpdateUserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserAuth(ctx, req.(*UpdateUserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DelUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DelUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/DelUserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DelUserAuth(ctx, req.(*DelUserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserAuthById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserAuthById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/GetUserAuthById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserAuthById(ctx, req.(*GetUserAuthByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SearchUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SearchUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/SearchUserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SearchUserAuth(ctx, req.(*SearchUserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Usercenter_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Usercenter_UpdateUser_Handler,
		},
		{
			MethodName: "DelUser",
			Handler:    _Usercenter_DelUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Usercenter_GetUserById_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _Usercenter_SearchUser_Handler,
		},
		{
			MethodName: "AddUserAuth",
			Handler:    _Usercenter_AddUserAuth_Handler,
		},
		{
			MethodName: "UpdateUserAuth",
			Handler:    _Usercenter_UpdateUserAuth_Handler,
		},
		{
			MethodName: "DelUserAuth",
			Handler:    _Usercenter_DelUserAuth_Handler,
		},
		{
			MethodName: "GetUserAuthById",
			Handler:    _Usercenter_GetUserAuthById_Handler,
		},
		{
			MethodName: "SearchUserAuth",
			Handler:    _Usercenter_SearchUserAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}