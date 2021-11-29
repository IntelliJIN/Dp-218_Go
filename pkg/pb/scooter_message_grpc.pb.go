// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ScooterServiceClient is the client API for ScooterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScooterServiceClient interface {
	Register(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (ScooterService_RegisterClient, error)
	Receive(ctx context.Context, opts ...grpc.CallOption) (ScooterService_ReceiveClient, error)
}

type scooterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScooterServiceClient(cc grpc.ClientConnInterface) ScooterServiceClient {
	return &scooterServiceClient{cc}
}

func (c *scooterServiceClient) Register(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (ScooterService_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScooterService_ServiceDesc.Streams[0], "/pb.ScooterService/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &scooterServiceRegisterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScooterService_RegisterClient interface {
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type scooterServiceRegisterClient struct {
	grpc.ClientStream
}

func (x *scooterServiceRegisterClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scooterServiceClient) Receive(ctx context.Context, opts ...grpc.CallOption) (ScooterService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScooterService_ServiceDesc.Streams[1], "/pb.ScooterService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &scooterServiceReceiveClient{stream}
	return x, nil
}

type ScooterService_ReceiveClient interface {
	Send(*ClientMessage) error
	CloseAndRecv() (*ServerMessage, error)
	grpc.ClientStream
}

type scooterServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *scooterServiceReceiveClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scooterServiceReceiveClient) CloseAndRecv() (*ServerMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ScooterServiceServer is the server API for ScooterService service.
// All implementations must embed UnimplementedScooterServiceServer
// for forward compatibility
type ScooterServiceServer interface {
	Register(*ClientRequest, ScooterService_RegisterServer) error
	Receive(ScooterService_ReceiveServer) error
	mustEmbedUnimplementedScooterServiceServer()
}

// UnimplementedScooterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScooterServiceServer struct {
}

func (UnimplementedScooterServiceServer) Register(*ClientRequest, ScooterService_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedScooterServiceServer) Receive(ScooterService_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedScooterServiceServer) mustEmbedUnimplementedScooterServiceServer() {}

// UnsafeScooterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScooterServiceServer will
// result in compilation errors.
type UnsafeScooterServiceServer interface {
	mustEmbedUnimplementedScooterServiceServer()
}

func RegisterScooterServiceServer(s grpc.ServiceRegistrar, srv ScooterServiceServer) {
	s.RegisterService(&ScooterService_ServiceDesc, srv)
}

func _ScooterService_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScooterServiceServer).Register(m, &scooterServiceRegisterServer{stream})
}

type ScooterService_RegisterServer interface {
	Send(*ServerMessage) error
	grpc.ServerStream
}

type scooterServiceRegisterServer struct {
	grpc.ServerStream
}

func (x *scooterServiceRegisterServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ScooterService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScooterServiceServer).Receive(&scooterServiceReceiveServer{stream})
}

type ScooterService_ReceiveServer interface {
	SendAndClose(*ServerMessage) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type scooterServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *scooterServiceReceiveServer) SendAndClose(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scooterServiceReceiveServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ScooterService_ServiceDesc is the grpc.ServiceDesc for ScooterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScooterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ScooterService",
	HandlerType: (*ScooterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ScooterService_Register_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Receive",
			Handler:       _ScooterService_Receive_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "scooter_message.proto",
}