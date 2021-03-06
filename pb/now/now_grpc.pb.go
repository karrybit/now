// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package now

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NowClient is the client API for Now service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NowClient interface {
	Tick(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Now_TickClient, error)
}

type nowClient struct {
	cc grpc.ClientConnInterface
}

func NewNowClient(cc grpc.ClientConnInterface) NowClient {
	return &nowClient{cc}
}

func (c *nowClient) Tick(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Now_TickClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Now_serviceDesc.Streams[0], "/now.Now/Tick", opts...)
	if err != nil {
		return nil, err
	}
	x := &nowTickClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Now_TickClient interface {
	Recv() (*timestamppb.Timestamp, error)
	grpc.ClientStream
}

type nowTickClient struct {
	grpc.ClientStream
}

func (x *nowTickClient) Recv() (*timestamppb.Timestamp, error) {
	m := new(timestamppb.Timestamp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NowServer is the server API for Now service.
// All implementations must embed UnimplementedNowServer
// for forward compatibility
type NowServer interface {
	Tick(*emptypb.Empty, Now_TickServer) error
	mustEmbedUnimplementedNowServer()
}

// UnimplementedNowServer must be embedded to have forward compatible implementations.
type UnimplementedNowServer struct {
}

func (UnimplementedNowServer) Tick(*emptypb.Empty, Now_TickServer) error {
	return status.Errorf(codes.Unimplemented, "method Tick not implemented")
}
func (UnimplementedNowServer) mustEmbedUnimplementedNowServer() {}

// UnsafeNowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NowServer will
// result in compilation errors.
type UnsafeNowServer interface {
	mustEmbedUnimplementedNowServer()
}

func RegisterNowServer(s grpc.ServiceRegistrar, srv NowServer) {
	s.RegisterService(&_Now_serviceDesc, srv)
}

func _Now_Tick_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NowServer).Tick(m, &nowTickServer{stream})
}

type Now_TickServer interface {
	Send(*timestamppb.Timestamp) error
	grpc.ServerStream
}

type nowTickServer struct {
	grpc.ServerStream
}

func (x *nowTickServer) Send(m *timestamppb.Timestamp) error {
	return x.ServerStream.SendMsg(m)
}

var _Now_serviceDesc = grpc.ServiceDesc{
	ServiceName: "now.Now",
	HandlerType: (*NowServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tick",
			Handler:       _Now_Tick_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/now.proto",
}
