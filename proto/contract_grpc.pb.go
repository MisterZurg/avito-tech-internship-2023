// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/contract.proto

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

const (
	SegmentsService_CreateSegment_FullMethodName = "/v1.SegmentsService/CreateSegment"
	SegmentsService_DeleteSegment_FullMethodName = "/v1.SegmentsService/DeleteSegment"
)

// SegmentsServiceClient is the client API for SegmentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SegmentsServiceClient interface {
	CreateSegment(ctx context.Context, in *CreateSegmentRequest, opts ...grpc.CallOption) (*CreateSegmentResponse, error)
	DeleteSegment(ctx context.Context, in *DeleteSegmentRequest, opts ...grpc.CallOption) (*DeleteSegmentResponse, error)
}

type segmentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSegmentsServiceClient(cc grpc.ClientConnInterface) SegmentsServiceClient {
	return &segmentsServiceClient{cc}
}

func (c *segmentsServiceClient) CreateSegment(ctx context.Context, in *CreateSegmentRequest, opts ...grpc.CallOption) (*CreateSegmentResponse, error) {
	out := new(CreateSegmentResponse)
	err := c.cc.Invoke(ctx, SegmentsService_CreateSegment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentsServiceClient) DeleteSegment(ctx context.Context, in *DeleteSegmentRequest, opts ...grpc.CallOption) (*DeleteSegmentResponse, error) {
	out := new(DeleteSegmentResponse)
	err := c.cc.Invoke(ctx, SegmentsService_DeleteSegment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SegmentsServiceServer is the server API for SegmentsService service.
// All implementations must embed UnimplementedSegmentsServiceServer
// for forward compatibility
type SegmentsServiceServer interface {
	CreateSegment(context.Context, *CreateSegmentRequest) (*CreateSegmentResponse, error)
	DeleteSegment(context.Context, *DeleteSegmentRequest) (*DeleteSegmentResponse, error)
	mustEmbedUnimplementedSegmentsServiceServer()
}

// UnimplementedSegmentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSegmentsServiceServer struct {
}

func (UnimplementedSegmentsServiceServer) CreateSegment(context.Context, *CreateSegmentRequest) (*CreateSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSegment not implemented")
}
func (UnimplementedSegmentsServiceServer) DeleteSegment(context.Context, *DeleteSegmentRequest) (*DeleteSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSegment not implemented")
}
func (UnimplementedSegmentsServiceServer) mustEmbedUnimplementedSegmentsServiceServer() {}

// UnsafeSegmentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SegmentsServiceServer will
// result in compilation errors.
type UnsafeSegmentsServiceServer interface {
	mustEmbedUnimplementedSegmentsServiceServer()
}

func RegisterSegmentsServiceServer(s grpc.ServiceRegistrar, srv SegmentsServiceServer) {
	s.RegisterService(&SegmentsService_ServiceDesc, srv)
}

func _SegmentsService_CreateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentsServiceServer).CreateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SegmentsService_CreateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentsServiceServer).CreateSegment(ctx, req.(*CreateSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentsService_DeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentsServiceServer).DeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SegmentsService_DeleteSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentsServiceServer).DeleteSegment(ctx, req.(*DeleteSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SegmentsService_ServiceDesc is the grpc.ServiceDesc for SegmentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SegmentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SegmentsService",
	HandlerType: (*SegmentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSegment",
			Handler:    _SegmentsService_CreateSegment_Handler,
		},
		{
			MethodName: "DeleteSegment",
			Handler:    _SegmentsService_DeleteSegment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/contract.proto",
}

const (
	UsersService_AddUserToSegment_FullMethodName  = "/v1.UsersService/AddUserToSegment"
	UsersService_GetActiveSegments_FullMethodName = "/v1.UsersService/GetActiveSegments"
)

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	AddUserToSegment(ctx context.Context, in *AddUserToSegmentRequest, opts ...grpc.CallOption) (*AddUserToSegmentResponse, error)
	GetActiveSegments(ctx context.Context, in *GetActiveSegmentsRequest, opts ...grpc.CallOption) (*GetActiveSegmentsResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) AddUserToSegment(ctx context.Context, in *AddUserToSegmentRequest, opts ...grpc.CallOption) (*AddUserToSegmentResponse, error) {
	out := new(AddUserToSegmentResponse)
	err := c.cc.Invoke(ctx, UsersService_AddUserToSegment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetActiveSegments(ctx context.Context, in *GetActiveSegmentsRequest, opts ...grpc.CallOption) (*GetActiveSegmentsResponse, error) {
	out := new(GetActiveSegmentsResponse)
	err := c.cc.Invoke(ctx, UsersService_GetActiveSegments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	AddUserToSegment(context.Context, *AddUserToSegmentRequest) (*AddUserToSegmentResponse, error)
	GetActiveSegments(context.Context, *GetActiveSegmentsRequest) (*GetActiveSegmentsResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) AddUserToSegment(context.Context, *AddUserToSegmentRequest) (*AddUserToSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToSegment not implemented")
}
func (UnimplementedUsersServiceServer) GetActiveSegments(context.Context, *GetActiveSegmentsRequest) (*GetActiveSegmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveSegments not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_AddUserToSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).AddUserToSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_AddUserToSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).AddUserToSegment(ctx, req.(*AddUserToSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetActiveSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveSegmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetActiveSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetActiveSegments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetActiveSegments(ctx, req.(*GetActiveSegmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserToSegment",
			Handler:    _UsersService_AddUserToSegment_Handler,
		},
		{
			MethodName: "GetActiveSegments",
			Handler:    _UsersService_GetActiveSegments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/contract.proto",
}