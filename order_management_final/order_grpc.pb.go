// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: order.proto

package __

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
	OrderManagementService_CreateOrder_FullMethodName = "/OrderManagementService/CreateOrder"
	OrderManagementService_UpdateOrder_FullMethodName = "/OrderManagementService/UpdateOrder"
	OrderManagementService_ReadOrder_FullMethodName   = "/OrderManagementService/ReadOrder"
	OrderManagementService_DeleteOrder_FullMethodName = "/OrderManagementService/DeleteOrder"
)

// OrderManagementServiceClient is the client API for OrderManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementServiceClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	ReadOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	DeleteOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
}

type orderManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementServiceClient(cc grpc.ClientConnInterface) OrderManagementServiceClient {
	return &orderManagementServiceClient{cc}
}

func (c *orderManagementServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderManagementService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementServiceClient) UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderManagementService_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementServiceClient) ReadOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderManagementService_ReadOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementServiceClient) DeleteOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderManagementService_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagementServiceServer is the server API for OrderManagementService service.
// All implementations must embed UnimplementedOrderManagementServiceServer
// for forward compatibility
type OrderManagementServiceServer interface {
	CreateOrder(context.Context, *Order) (*Order, error)
	UpdateOrder(context.Context, *Order) (*Order, error)
	ReadOrder(context.Context, *Order) (*Order, error)
	DeleteOrder(context.Context, *Order) (*Order, error)
	mustEmbedUnimplementedOrderManagementServiceServer()
}

// UnimplementedOrderManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServiceServer struct {
}

func (UnimplementedOrderManagementServiceServer) CreateOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) UpdateOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) ReadOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) DeleteOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) mustEmbedUnimplementedOrderManagementServiceServer() {
}

// UnsafeOrderManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServiceServer will
// result in compilation errors.
type UnsafeOrderManagementServiceServer interface {
	mustEmbedUnimplementedOrderManagementServiceServer()
}

func RegisterOrderManagementServiceServer(s grpc.ServiceRegistrar, srv OrderManagementServiceServer) {
	s.RegisterService(&OrderManagementService_ServiceDesc, srv)
}

func _OrderManagementService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagementService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagementService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagementService_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).UpdateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagementService_ReadOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).ReadOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagementService_ReadOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).ReadOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagementService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagementService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServiceServer).DeleteOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderManagementService_ServiceDesc is the grpc.ServiceDesc for OrderManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderManagementService",
	HandlerType: (*OrderManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderManagementService_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderManagementService_UpdateOrder_Handler,
		},
		{
			MethodName: "ReadOrder",
			Handler:    _OrderManagementService_ReadOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderManagementService_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
