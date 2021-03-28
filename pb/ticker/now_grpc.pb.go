// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ticker

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

// TickClient is the client API for Tick service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TickClient interface {
	Now(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tick_NowClient, error)
}

type tickClient struct {
	cc grpc.ClientConnInterface
}

func NewTickClient(cc grpc.ClientConnInterface) TickClient {
	return &tickClient{cc}
}

func (c *tickClient) Now(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tick_NowClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tick_serviceDesc.Streams[0], "/now.Tick/Now", opts...)
	if err != nil {
		return nil, err
	}
	x := &tickNowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tick_NowClient interface {
	Recv() (*timestamppb.Timestamp, error)
	grpc.ClientStream
}

type tickNowClient struct {
	grpc.ClientStream
}

func (x *tickNowClient) Recv() (*timestamppb.Timestamp, error) {
	m := new(timestamppb.Timestamp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TickServer is the server API for Tick service.
// All implementations must embed UnimplementedTickServer
// for forward compatibility
type TickServer interface {
	Now(*emptypb.Empty, Tick_NowServer) error
	mustEmbedUnimplementedTickServer()
}

// UnimplementedTickServer must be embedded to have forward compatible implementations.
type UnimplementedTickServer struct {
}

func (UnimplementedTickServer) Now(*emptypb.Empty, Tick_NowServer) error {
	return status.Errorf(codes.Unimplemented, "method Now not implemented")
}
func (UnimplementedTickServer) mustEmbedUnimplementedTickServer() {}

// UnsafeTickServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TickServer will
// result in compilation errors.
type UnsafeTickServer interface {
	mustEmbedUnimplementedTickServer()
}

func RegisterTickServer(s grpc.ServiceRegistrar, srv TickServer) {
	s.RegisterService(&_Tick_serviceDesc, srv)
}

func _Tick_Now_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TickServer).Now(m, &tickNowServer{stream})
}

type Tick_NowServer interface {
	Send(*timestamppb.Timestamp) error
	grpc.ServerStream
}

type tickNowServer struct {
	grpc.ServerStream
}

func (x *tickNowServer) Send(m *timestamppb.Timestamp) error {
	return x.ServerStream.SendMsg(m)
}

var _Tick_serviceDesc = grpc.ServiceDesc{
	ServiceName: "now.Tick",
	HandlerType: (*TickServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Now",
			Handler:       _Tick_Now_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/now.proto",
}