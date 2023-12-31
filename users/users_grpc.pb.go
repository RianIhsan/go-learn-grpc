// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: users/users.proto

package users

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

const (
	DataUsers_FindUserByEmail_FullMethodName = "/users.DataUsers/FindUserByEmail"
)

// DataUsersClient is the client API for DataUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataUsersClient interface {
	FindUserByEmail(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type dataUsersClient struct {
	cc grpc.ClientConnInterface
}

func NewDataUsersClient(cc grpc.ClientConnInterface) DataUsersClient {
	return &dataUsersClient{cc}
}

func (c *dataUsersClient) FindUserByEmail(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, DataUsers_FindUserByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataUsersServer is the server API for DataUsers service.
// All implementations must embed UnimplementedDataUsersServer
// for forward compatibility
type DataUsersServer interface {
	FindUserByEmail(context.Context, *User) (*User, error)
	mustEmbedUnimplementedDataUsersServer()
}

// UnimplementedDataUsersServer must be embedded to have forward compatible implementations.
type UnimplementedDataUsersServer struct {
}

func (UnimplementedDataUsersServer) FindUserByEmail(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmail not implemented")
}
func (UnimplementedDataUsersServer) mustEmbedUnimplementedDataUsersServer() {}

// UnsafeDataUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataUsersServer will
// result in compilation errors.
type UnsafeDataUsersServer interface {
	mustEmbedUnimplementedDataUsersServer()
}

func RegisterDataUsersServer(s grpc.ServiceRegistrar, srv DataUsersServer) {
	s.RegisterService(&DataUsers_ServiceDesc, srv)
}

func _DataUsers_FindUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUsersServer).FindUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataUsers_FindUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUsersServer).FindUserByEmail(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// DataUsers_ServiceDesc is the grpc.ServiceDesc for DataUsers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataUsers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.DataUsers",
	HandlerType: (*DataUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserByEmail",
			Handler:    _DataUsers_FindUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/users.proto",
}
