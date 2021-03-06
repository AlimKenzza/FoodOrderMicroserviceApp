// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package restarauntpb

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

// FoodOrderServiceClient is the client API for FoodOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodOrderServiceClient interface {
	GetOFoodsByOrder(ctx context.Context, in *FoodsByOrderRequest, opts ...grpc.CallOption) (FoodOrderService_GetOFoodsByOrderClient, error)
}

type foodOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodOrderServiceClient(cc grpc.ClientConnInterface) FoodOrderServiceClient {
	return &foodOrderServiceClient{cc}
}

func (c *foodOrderServiceClient) GetOFoodsByOrder(ctx context.Context, in *FoodsByOrderRequest, opts ...grpc.CallOption) (FoodOrderService_GetOFoodsByOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &FoodOrderService_ServiceDesc.Streams[0], "/grpc.FoodOrderService/GetOFoodsByOrder", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodOrderServiceGetOFoodsByOrderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FoodOrderService_GetOFoodsByOrderClient interface {
	Recv() (*FoodsByOrderResponse, error)
	grpc.ClientStream
}

type foodOrderServiceGetOFoodsByOrderClient struct {
	grpc.ClientStream
}

func (x *foodOrderServiceGetOFoodsByOrderClient) Recv() (*FoodsByOrderResponse, error) {
	m := new(FoodsByOrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoodOrderServiceServer is the server API for FoodOrderService service.
// All implementations must embed UnimplementedFoodOrderServiceServer
// for forward compatibility
type FoodOrderServiceServer interface {
	GetOFoodsByOrder(*FoodsByOrderRequest, FoodOrderService_GetOFoodsByOrderServer) error
	mustEmbedUnimplementedFoodOrderServiceServer()
}

// UnimplementedFoodOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFoodOrderServiceServer struct {
}

func (UnimplementedFoodOrderServiceServer) GetOFoodsByOrder(*FoodsByOrderRequest, FoodOrderService_GetOFoodsByOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOFoodsByOrder not implemented")
}
func (UnimplementedFoodOrderServiceServer) mustEmbedUnimplementedFoodOrderServiceServer() {}

// UnsafeFoodOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodOrderServiceServer will
// result in compilation errors.
type UnsafeFoodOrderServiceServer interface {
	mustEmbedUnimplementedFoodOrderServiceServer()
}

func RegisterFoodOrderServiceServer(s grpc.ServiceRegistrar, srv FoodOrderServiceServer) {
	s.RegisterService(&FoodOrderService_ServiceDesc, srv)
}

func _FoodOrderService_GetOFoodsByOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FoodsByOrderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FoodOrderServiceServer).GetOFoodsByOrder(m, &foodOrderServiceGetOFoodsByOrderServer{stream})
}

type FoodOrderService_GetOFoodsByOrderServer interface {
	Send(*FoodsByOrderResponse) error
	grpc.ServerStream
}

type foodOrderServiceGetOFoodsByOrderServer struct {
	grpc.ServerStream
}

func (x *foodOrderServiceGetOFoodsByOrderServer) Send(m *FoodsByOrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FoodOrderService_ServiceDesc is the grpc.ServiceDesc for FoodOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoodOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.FoodOrderService",
	HandlerType: (*FoodOrderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOFoodsByOrder",
			Handler:       _FoodOrderService_GetOFoodsByOrder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/models.proto",
}
