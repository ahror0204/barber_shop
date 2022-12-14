// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: salon.proto

package users_service

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

// SalonServiceClient is the client API for SalonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SalonServiceClient interface {
	CreateSalon(ctx context.Context, in *Salon, opts ...grpc.CallOption) (*Salon, error)
	UpdateSalon(ctx context.Context, in *Salon, opts ...grpc.CallOption) (*Salon, error)
	GetSalonByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Salon, error)
	GetListSalons(ctx context.Context, in *GetListParams, opts ...grpc.CallOption) (*AllSalons, error)
	DeleteSalon(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type salonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSalonServiceClient(cc grpc.ClientConnInterface) SalonServiceClient {
	return &salonServiceClient{cc}
}

func (c *salonServiceClient) CreateSalon(ctx context.Context, in *Salon, opts ...grpc.CallOption) (*Salon, error) {
	out := new(Salon)
	err := c.cc.Invoke(ctx, "/protos.SalonService/CreateSalon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salonServiceClient) UpdateSalon(ctx context.Context, in *Salon, opts ...grpc.CallOption) (*Salon, error) {
	out := new(Salon)
	err := c.cc.Invoke(ctx, "/protos.SalonService/UpdateSalon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salonServiceClient) GetSalonByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Salon, error) {
	out := new(Salon)
	err := c.cc.Invoke(ctx, "/protos.SalonService/GetSalonByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salonServiceClient) GetListSalons(ctx context.Context, in *GetListParams, opts ...grpc.CallOption) (*AllSalons, error) {
	out := new(AllSalons)
	err := c.cc.Invoke(ctx, "/protos.SalonService/GetListSalons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salonServiceClient) DeleteSalon(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.SalonService/DeleteSalon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SalonServiceServer is the server API for SalonService service.
// All implementations must embed UnimplementedSalonServiceServer
// for forward compatibility
type SalonServiceServer interface {
	CreateSalon(context.Context, *Salon) (*Salon, error)
	UpdateSalon(context.Context, *Salon) (*Salon, error)
	GetSalonByID(context.Context, *ID) (*Salon, error)
	GetListSalons(context.Context, *GetListParams) (*AllSalons, error)
	DeleteSalon(context.Context, *ID) (*Empty, error)
	mustEmbedUnimplementedSalonServiceServer()
}

// UnimplementedSalonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSalonServiceServer struct {
}

func (UnimplementedSalonServiceServer) CreateSalon(context.Context, *Salon) (*Salon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSalon not implemented")
}
func (UnimplementedSalonServiceServer) UpdateSalon(context.Context, *Salon) (*Salon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSalon not implemented")
}
func (UnimplementedSalonServiceServer) GetSalonByID(context.Context, *ID) (*Salon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSalonByID not implemented")
}
func (UnimplementedSalonServiceServer) GetListSalons(context.Context, *GetListParams) (*AllSalons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListSalons not implemented")
}
func (UnimplementedSalonServiceServer) DeleteSalon(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSalon not implemented")
}
func (UnimplementedSalonServiceServer) mustEmbedUnimplementedSalonServiceServer() {}

// UnsafeSalonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SalonServiceServer will
// result in compilation errors.
type UnsafeSalonServiceServer interface {
	mustEmbedUnimplementedSalonServiceServer()
}

func RegisterSalonServiceServer(s grpc.ServiceRegistrar, srv SalonServiceServer) {
	s.RegisterService(&SalonService_ServiceDesc, srv)
}

func _SalonService_CreateSalon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Salon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalonServiceServer).CreateSalon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalonService/CreateSalon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalonServiceServer).CreateSalon(ctx, req.(*Salon))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalonService_UpdateSalon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Salon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalonServiceServer).UpdateSalon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalonService/UpdateSalon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalonServiceServer).UpdateSalon(ctx, req.(*Salon))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalonService_GetSalonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalonServiceServer).GetSalonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalonService/GetSalonByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalonServiceServer).GetSalonByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalonService_GetListSalons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalonServiceServer).GetListSalons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalonService/GetListSalons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalonServiceServer).GetListSalons(ctx, req.(*GetListParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalonService_DeleteSalon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalonServiceServer).DeleteSalon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalonService/DeleteSalon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalonServiceServer).DeleteSalon(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// SalonService_ServiceDesc is the grpc.ServiceDesc for SalonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SalonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.SalonService",
	HandlerType: (*SalonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSalon",
			Handler:    _SalonService_CreateSalon_Handler,
		},
		{
			MethodName: "UpdateSalon",
			Handler:    _SalonService_UpdateSalon_Handler,
		},
		{
			MethodName: "GetSalonByID",
			Handler:    _SalonService_GetSalonByID_Handler,
		},
		{
			MethodName: "GetListSalons",
			Handler:    _SalonService_GetListSalons_Handler,
		},
		{
			MethodName: "DeleteSalon",
			Handler:    _SalonService_DeleteSalon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "salon.proto",
}
