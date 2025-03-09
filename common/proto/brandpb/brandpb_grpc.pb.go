// Code generated by template. DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: brandpb.proto

package brandpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ModelService_GetList_FullMethodName    = "/brandpb.ModelService/GetList"
	ModelService_GetItem_FullMethodName    = "/brandpb.ModelService/GetItem"
	ModelService_CreateItem_FullMethodName = "/brandpb.ModelService/CreateItem"
	ModelService_UpdateItem_FullMethodName = "/brandpb.ModelService/UpdateItem"
	ModelService_DeleteItem_FullMethodName = "/brandpb.ModelService/DeleteItem"
)

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServiceClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, ModelService_GetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, ModelService_GetItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, ModelService_CreateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, ModelService_UpdateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, ModelService_DeleteItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServiceServer is the server API for ModelService service.
// All implementations must embed UnimplementedModelServiceServer
// for forward compatibility.
type ModelServiceServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	mustEmbedUnimplementedModelServiceServer()
}

// UnimplementedModelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModelServiceServer struct{}

func (UnimplementedModelServiceServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedModelServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedModelServiceServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedModelServiceServer) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedModelServiceServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedModelServiceServer) mustEmbedUnimplementedModelServiceServer() {}
func (UnimplementedModelServiceServer) testEmbeddedByValue()                      {}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	// If the following call pancis, it indicates UnimplementedModelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_GetItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_UpdateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_DeleteItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brandpb.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _ModelService_GetList_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _ModelService_GetItem_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _ModelService_CreateItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ModelService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ModelService_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brandpb.proto",
}
