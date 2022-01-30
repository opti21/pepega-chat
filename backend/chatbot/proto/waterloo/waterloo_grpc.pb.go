// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package waterloo

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

// PublicQueueCacheServiceClient is the client API for PublicQueueCacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicQueueCacheServiceClient interface {
	Lookup(ctx context.Context, in *CacheSearch, opts ...grpc.CallOption) (*CacheResult, error)
	Store(ctx context.Context, in *CacheVideoDetails, opts ...grpc.CallOption) (*StoreStatus, error)
}

type publicQueueCacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicQueueCacheServiceClient(cc grpc.ClientConnInterface) PublicQueueCacheServiceClient {
	return &publicQueueCacheServiceClient{cc}
}

func (c *publicQueueCacheServiceClient) Lookup(ctx context.Context, in *CacheSearch, opts ...grpc.CallOption) (*CacheResult, error) {
	out := new(CacheResult)
	err := c.cc.Invoke(ctx, "/onethirtyone.PublicQueueCacheService/lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicQueueCacheServiceClient) Store(ctx context.Context, in *CacheVideoDetails, opts ...grpc.CallOption) (*StoreStatus, error) {
	out := new(StoreStatus)
	err := c.cc.Invoke(ctx, "/onethirtyone.PublicQueueCacheService/store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicQueueCacheServiceServer is the server API for PublicQueueCacheService service.
// All implementations must embed UnimplementedPublicQueueCacheServiceServer
// for forward compatibility
type PublicQueueCacheServiceServer interface {
	Lookup(context.Context, *CacheSearch) (*CacheResult, error)
	Store(context.Context, *CacheVideoDetails) (*StoreStatus, error)
	mustEmbedUnimplementedPublicQueueCacheServiceServer()
}

// UnimplementedPublicQueueCacheServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicQueueCacheServiceServer struct {
}

func (UnimplementedPublicQueueCacheServiceServer) Lookup(context.Context, *CacheSearch) (*CacheResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedPublicQueueCacheServiceServer) Store(context.Context, *CacheVideoDetails) (*StoreStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedPublicQueueCacheServiceServer) mustEmbedUnimplementedPublicQueueCacheServiceServer() {
}

// UnsafePublicQueueCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicQueueCacheServiceServer will
// result in compilation errors.
type UnsafePublicQueueCacheServiceServer interface {
	mustEmbedUnimplementedPublicQueueCacheServiceServer()
}

func RegisterPublicQueueCacheServiceServer(s grpc.ServiceRegistrar, srv PublicQueueCacheServiceServer) {
	s.RegisterService(&PublicQueueCacheService_ServiceDesc, srv)
}

func _PublicQueueCacheService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicQueueCacheServiceServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onethirtyone.PublicQueueCacheService/lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicQueueCacheServiceServer).Lookup(ctx, req.(*CacheSearch))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicQueueCacheService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheVideoDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicQueueCacheServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onethirtyone.PublicQueueCacheService/store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicQueueCacheServiceServer).Store(ctx, req.(*CacheVideoDetails))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicQueueCacheService_ServiceDesc is the grpc.ServiceDesc for PublicQueueCacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicQueueCacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "onethirtyone.PublicQueueCacheService",
	HandlerType: (*PublicQueueCacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "lookup",
			Handler:    _PublicQueueCacheService_Lookup_Handler,
		},
		{
			MethodName: "store",
			Handler:    _PublicQueueCacheService_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "waterloo.proto",
}
