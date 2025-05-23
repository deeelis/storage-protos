// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: storage/storage.proto

package storagev1

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
	StorageService_GetContent_FullMethodName      = "/storage.StorageService/GetContent"
	StorageService_RegisterContent_FullMethodName = "/storage.StorageService/RegisterContent"
)

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	GetContent(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentResponse, error)
	RegisterContent(ctx context.Context, in *RegisterContentRequest, opts ...grpc.CallOption) (*RegisterContentResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) GetContent(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContentResponse)
	err := c.cc.Invoke(ctx, StorageService_GetContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) RegisterContent(ctx context.Context, in *RegisterContentRequest, opts ...grpc.CallOption) (*RegisterContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterContentResponse)
	err := c.cc.Invoke(ctx, StorageService_RegisterContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility.
type StorageServiceServer interface {
	GetContent(context.Context, *ContentRequest) (*ContentResponse, error)
	RegisterContent(context.Context, *RegisterContentRequest) (*RegisterContentResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageServiceServer struct{}

func (UnimplementedStorageServiceServer) GetContent(context.Context, *ContentRequest) (*ContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedStorageServiceServer) RegisterContent(context.Context, *RegisterContentRequest) (*RegisterContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterContent not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}
func (UnimplementedStorageServiceServer) testEmbeddedByValue()                        {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	// If the following call pancis, it indicates UnimplementedStorageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_GetContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetContent(ctx, req.(*ContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_RegisterContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).RegisterContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageService_RegisterContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).RegisterContent(ctx, req.(*RegisterContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContent",
			Handler:    _StorageService_GetContent_Handler,
		},
		{
			MethodName: "RegisterContent",
			Handler:    _StorageService_RegisterContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage/storage.proto",
}
