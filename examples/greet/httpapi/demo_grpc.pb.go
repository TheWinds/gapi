// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package httpapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DemoAPIClient is the client API for DemoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoAPIClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReply, error)
	GreetCount(ctx context.Context, in *GreetCountRequest, opts ...grpc.CallOption) (*GreetCountReply, error)
}

type demoAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoAPIClient(cc grpc.ClientConnInterface) DemoAPIClient {
	return &demoAPIClient{cc}
}

func (c *demoAPIClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReply, error) {
	out := new(GreetReply)
	err := c.cc.Invoke(ctx, "/demo.DemoAPI/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoAPIClient) GreetCount(ctx context.Context, in *GreetCountRequest, opts ...grpc.CallOption) (*GreetCountReply, error) {
	out := new(GreetCountReply)
	err := c.cc.Invoke(ctx, "/demo.DemoAPI/GreetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoAPIServer is the server API for DemoAPI service.
// All implementations must embed UnimplementedDemoAPIServer
// for forward compatibility
type DemoAPIServer interface {
	Greet(context.Context, *GreetRequest) (*GreetReply, error)
	GreetCount(context.Context, *GreetCountRequest) (*GreetCountReply, error)
	mustEmbedUnimplementedDemoAPIServer()
}

// UnimplementedDemoAPIServer must be embedded to have forward compatible implementations.
type UnimplementedDemoAPIServer struct {
}

func (*UnimplementedDemoAPIServer) Greet(context.Context, *GreetRequest) (*GreetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedDemoAPIServer) GreetCount(context.Context, *GreetCountRequest) (*GreetCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetCount not implemented")
}
func (*UnimplementedDemoAPIServer) mustEmbedUnimplementedDemoAPIServer() {}

func RegisterDemoAPIServer(s *grpc.Server, srv DemoAPIServer) {
	s.RegisterService(&_DemoAPI_serviceDesc, srv)
}

func _DemoAPI_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoAPIServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoAPI/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoAPIServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoAPI_GreetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoAPIServer).GreetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoAPI/GreetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoAPIServer).GreetCount(ctx, req.(*GreetCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.DemoAPI",
	HandlerType: (*DemoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _DemoAPI_Greet_Handler,
		},
		{
			MethodName: "GreetCount",
			Handler:    _DemoAPI_GreetCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
