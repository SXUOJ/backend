// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pb/service.proto

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

// JudgerClient is the client API for Judger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error)
	Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeReply, error)
}

type judgerClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgerClient(cc grpc.ClientConnInterface) JudgerClient {
	return &judgerClient{cc}
}

func (c *judgerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/judger.Judger/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgerClient) Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeReply, error) {
	out := new(JudgeReply)
	err := c.cc.Invoke(ctx, "/judger.Judger/Judge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgerServer is the server API for Judger service.
// All implementations must embed UnimplementedJudgerServer
// for forward compatibility
type JudgerServer interface {
	Ping(context.Context, *PingRequest) (*PongReply, error)
	Judge(context.Context, *JudgeRequest) (*JudgeReply, error)
	mustEmbedUnimplementedJudgerServer()
}

// UnimplementedJudgerServer must be embedded to have forward compatible implementations.
type UnimplementedJudgerServer struct {
}

func (UnimplementedJudgerServer) Ping(context.Context, *PingRequest) (*PongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedJudgerServer) Judge(context.Context, *JudgeRequest) (*JudgeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Judge not implemented")
}
func (UnimplementedJudgerServer) mustEmbedUnimplementedJudgerServer() {}

// UnsafeJudgerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgerServer will
// result in compilation errors.
type UnsafeJudgerServer interface {
	mustEmbedUnimplementedJudgerServer()
}

func RegisterJudgerServer(s grpc.ServiceRegistrar, srv JudgerServer) {
	s.RegisterService(&Judger_ServiceDesc, srv)
}

func _Judger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judger.Judger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judger_Judge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgerServer).Judge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judger.Judger/Judge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgerServer).Judge(ctx, req.(*JudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Judger_ServiceDesc is the grpc.ServiceDesc for Judger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Judger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judger.Judger",
	HandlerType: (*JudgerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Judger_Ping_Handler,
		},
		{
			MethodName: "Judge",
			Handler:    _Judger_Judge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}
