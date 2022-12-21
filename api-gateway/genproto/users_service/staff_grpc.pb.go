// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: staff.proto

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

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	CreateStaff(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error)
	UpdateStaff(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error)
	GetStaffByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Staff, error)
	GetListStaff(ctx context.Context, in *GetListParams, opts ...grpc.CallOption) (*ListStaff, error)
	DeleteStaff(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	GetStaffByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Customer, error)
	UpdateStaffPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Empty, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) CreateStaff(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/protos.StaffService/CreateStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) UpdateStaff(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/protos.StaffService/UpdateStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetStaffByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/protos.StaffService/GetStaffByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetListStaff(ctx context.Context, in *GetListParams, opts ...grpc.CallOption) (*ListStaff, error) {
	out := new(ListStaff)
	err := c.cc.Invoke(ctx, "/protos.StaffService/GetListStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) DeleteStaff(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.StaffService/DeleteStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetStaffByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/protos.StaffService/GetStaffByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) UpdateStaffPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.StaffService/UpdateStaffPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	CreateStaff(context.Context, *Staff) (*Staff, error)
	UpdateStaff(context.Context, *Staff) (*Staff, error)
	GetStaffByID(context.Context, *ID) (*Staff, error)
	GetListStaff(context.Context, *GetListParams) (*ListStaff, error)
	DeleteStaff(context.Context, *ID) (*Empty, error)
	GetStaffByEmail(context.Context, *Email) (*Customer, error)
	UpdateStaffPassword(context.Context, *UpdatePasswordRequest) (*Empty, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) CreateStaff(context.Context, *Staff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaff not implemented")
}
func (UnimplementedStaffServiceServer) UpdateStaff(context.Context, *Staff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaff not implemented")
}
func (UnimplementedStaffServiceServer) GetStaffByID(context.Context, *ID) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaffByID not implemented")
}
func (UnimplementedStaffServiceServer) GetListStaff(context.Context, *GetListParams) (*ListStaff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListStaff not implemented")
}
func (UnimplementedStaffServiceServer) DeleteStaff(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStaff not implemented")
}
func (UnimplementedStaffServiceServer) GetStaffByEmail(context.Context, *Email) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaffByEmail not implemented")
}
func (UnimplementedStaffServiceServer) UpdateStaffPassword(context.Context, *UpdatePasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaffPassword not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_CreateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Staff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).CreateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/CreateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).CreateStaff(ctx, req.(*Staff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_UpdateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Staff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).UpdateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/UpdateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).UpdateStaff(ctx, req.(*Staff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetStaffByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetStaffByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/GetStaffByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetStaffByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetListStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetListStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/GetListStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetListStaff(ctx, req.(*GetListParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_DeleteStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).DeleteStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/DeleteStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).DeleteStaff(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetStaffByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetStaffByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/GetStaffByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetStaffByEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_UpdateStaffPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).UpdateStaffPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.StaffService/UpdateStaffPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).UpdateStaffPassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStaff",
			Handler:    _StaffService_CreateStaff_Handler,
		},
		{
			MethodName: "UpdateStaff",
			Handler:    _StaffService_UpdateStaff_Handler,
		},
		{
			MethodName: "GetStaffByID",
			Handler:    _StaffService_GetStaffByID_Handler,
		},
		{
			MethodName: "GetListStaff",
			Handler:    _StaffService_GetListStaff_Handler,
		},
		{
			MethodName: "DeleteStaff",
			Handler:    _StaffService_DeleteStaff_Handler,
		},
		{
			MethodName: "GetStaffByEmail",
			Handler:    _StaffService_GetStaffByEmail_Handler,
		},
		{
			MethodName: "UpdateStaffPassword",
			Handler:    _StaffService_UpdateStaffPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff.proto",
}
