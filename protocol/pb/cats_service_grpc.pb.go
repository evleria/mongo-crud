// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatsServiceClient is the client API for CatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatsServiceClient interface {
	GetAllCats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CatsService_GetAllCatsClient, error)
	GetCat(ctx context.Context, in *GetCatRequest, opts ...grpc.CallOption) (*GetCatResponse, error)
	AddNewCat(ctx context.Context, in *AddNewCatRequest, opts ...grpc.CallOption) (*AddNewCatResponse, error)
	DeleteCat(ctx context.Context, in *DeleteCatRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdatePrice(ctx context.Context, in *UpdatePriceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type catsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatsServiceClient(cc grpc.ClientConnInterface) CatsServiceClient {
	return &catsServiceClient{cc}
}

func (c *catsServiceClient) GetAllCats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CatsService_GetAllCatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CatsService_ServiceDesc.Streams[0], "/CatsService/GetAllCats", opts...)
	if err != nil {
		return nil, err
	}
	x := &catsServiceGetAllCatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatsService_GetAllCatsClient interface {
	Recv() (*GetAllCatsResponse, error)
	grpc.ClientStream
}

type catsServiceGetAllCatsClient struct {
	grpc.ClientStream
}

func (x *catsServiceGetAllCatsClient) Recv() (*GetAllCatsResponse, error) {
	m := new(GetAllCatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catsServiceClient) GetCat(ctx context.Context, in *GetCatRequest, opts ...grpc.CallOption) (*GetCatResponse, error) {
	out := new(GetCatResponse)
	err := c.cc.Invoke(ctx, "/CatsService/GetCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsServiceClient) AddNewCat(ctx context.Context, in *AddNewCatRequest, opts ...grpc.CallOption) (*AddNewCatResponse, error) {
	out := new(AddNewCatResponse)
	err := c.cc.Invoke(ctx, "/CatsService/AddNewCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsServiceClient) DeleteCat(ctx context.Context, in *DeleteCatRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/CatsService/DeleteCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsServiceClient) UpdatePrice(ctx context.Context, in *UpdatePriceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/CatsService/UpdatePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatsServiceServer is the server API for CatsService service.
// All implementations must embed UnimplementedCatsServiceServer
// for forward compatibility
type CatsServiceServer interface {
	GetAllCats(*empty.Empty, CatsService_GetAllCatsServer) error
	GetCat(context.Context, *GetCatRequest) (*GetCatResponse, error)
	AddNewCat(context.Context, *AddNewCatRequest) (*AddNewCatResponse, error)
	DeleteCat(context.Context, *DeleteCatRequest) (*empty.Empty, error)
	UpdatePrice(context.Context, *UpdatePriceRequest) (*empty.Empty, error)
	mustEmbedUnimplementedCatsServiceServer()
}

// UnimplementedCatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatsServiceServer struct {
}

func (UnimplementedCatsServiceServer) GetAllCats(*empty.Empty, CatsService_GetAllCatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllCats not implemented")
}
func (UnimplementedCatsServiceServer) GetCat(context.Context, *GetCatRequest) (*GetCatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCat not implemented")
}
func (UnimplementedCatsServiceServer) AddNewCat(context.Context, *AddNewCatRequest) (*AddNewCatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewCat not implemented")
}
func (UnimplementedCatsServiceServer) DeleteCat(context.Context, *DeleteCatRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCat not implemented")
}
func (UnimplementedCatsServiceServer) UpdatePrice(context.Context, *UpdatePriceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrice not implemented")
}
func (UnimplementedCatsServiceServer) mustEmbedUnimplementedCatsServiceServer() {}

// UnsafeCatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatsServiceServer will
// result in compilation errors.
type UnsafeCatsServiceServer interface {
	mustEmbedUnimplementedCatsServiceServer()
}

func RegisterCatsServiceServer(s grpc.ServiceRegistrar, srv CatsServiceServer) {
	s.RegisterService(&CatsService_ServiceDesc, srv)
}

func _CatsService_GetAllCats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatsServiceServer).GetAllCats(m, &catsServiceGetAllCatsServer{stream})
}

type CatsService_GetAllCatsServer interface {
	Send(*GetAllCatsResponse) error
	grpc.ServerStream
}

type catsServiceGetAllCatsServer struct {
	grpc.ServerStream
}

func (x *catsServiceGetAllCatsServer) Send(m *GetAllCatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CatsService_GetCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).GetCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatsService/GetCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).GetCat(ctx, req.(*GetCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsService_AddNewCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).AddNewCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatsService/AddNewCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).AddNewCat(ctx, req.(*AddNewCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsService_DeleteCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).DeleteCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatsService/DeleteCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).DeleteCat(ctx, req.(*DeleteCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatsService_UpdatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServiceServer).UpdatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatsService/UpdatePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServiceServer).UpdatePrice(ctx, req.(*UpdatePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatsService_ServiceDesc is the grpc.ServiceDesc for CatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CatsService",
	HandlerType: (*CatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCat",
			Handler:    _CatsService_GetCat_Handler,
		},
		{
			MethodName: "AddNewCat",
			Handler:    _CatsService_AddNewCat_Handler,
		},
		{
			MethodName: "DeleteCat",
			Handler:    _CatsService_DeleteCat_Handler,
		},
		{
			MethodName: "UpdatePrice",
			Handler:    _CatsService_UpdatePrice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllCats",
			Handler:       _CatsService_GetAllCats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cats_service.proto",
}
