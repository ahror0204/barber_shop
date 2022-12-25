// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: customer_auth.proto

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

// CustomerAuthServiceClient is the client API for CustomerAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerAuthServiceClient interface {
	CustomerRegister(ctx context.Context, in *CustomerRegisterRequest, opts ...grpc.CallOption) (*Empty, error)
	CustomerVerify(ctx context.Context, in *VerifyCustomerRegisterRequest, opts ...grpc.CallOption) (*CustomerAuthResponse, error)
	CustomerLogin(ctx context.Context, in *CustomerLoginRequest, opts ...grpc.CallOption) (*CustomerAuthResponse, error)
	CustomerForgotPassword(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Empty, error)
	VerifyCustomerForgotPassword(ctx context.Context, in *VerifyCustomerRegisterRequest, opts ...grpc.CallOption) (*CustomerAuthResponse, error)
	UpdateCustomerPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Empty, error)
}

type customerAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerAuthServiceClient(cc grpc.ClientConnInterface) CustomerAuthServiceClient {
	return &customerAuthServiceClient{cc}
}

func (c *customerAuthServiceClient) CustomerRegister(ctx context.Context, in *CustomerRegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CustomerAuthService/CustomerRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) CustomerVerify(ctx context.Context, in *VerifyCustomerRegisterRequest, opts ...grpc.CallOption) (*CustomerAuthResponse, error) {
	out := new(CustomerAuthResponse)
	err := c.cc.Invoke(ctx, "/protos.CustomerAuthService/CustomerVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) CustomerLogin(ctx context.Context, in *CustomerLoginRequest, opts ...grpc.CallOption) (*CustomerAuthResponse, error) {
	out := new(CustomerAuthResponse)
	err := c.cc.Invoke(ctx, "/protos.CustomerAuthService/CustomerLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) CustomerForgotPassword(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CustomerAuthService/CustomerForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) VerifyCustomerForgotPassword(ctx context.Context, in *VerifyCustomerRegisterRequest, opts ...grpc.CallOption) (*CustomerAuthResponse, error) {
	out := new(CustomerAuthResponse)
	err := c.cc.Invoke(ctx, "/protos.CustomerAuthService/VerifyCustomerForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) UpdateCustomerPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.CustomerAuthService/UpdateCustomerPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerAuthServiceServer is the server API for CustomerAuthService service.
// All implementations must embed UnimplementedCustomerAuthServiceServer
// for forward compatibility
type CustomerAuthServiceServer interface {
	CustomerRegister(context.Context, *CustomerRegisterRequest) (*Empty, error)
	CustomerVerify(context.Context, *VerifyCustomerRegisterRequest) (*CustomerAuthResponse, error)
	CustomerLogin(context.Context, *CustomerLoginRequest) (*CustomerAuthResponse, error)
	CustomerForgotPassword(context.Context, *Email) (*Empty, error)
	VerifyCustomerForgotPassword(context.Context, *VerifyCustomerRegisterRequest) (*CustomerAuthResponse, error)
	UpdateCustomerPassword(context.Context, *UpdatePasswordRequest) (*Empty, error)
	mustEmbedUnimplementedCustomerAuthServiceServer()
}

// UnimplementedCustomerAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerAuthServiceServer struct {
}

func (UnimplementedCustomerAuthServiceServer) CustomerRegister(context.Context, *CustomerRegisterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerRegister not implemented")
}
func (UnimplementedCustomerAuthServiceServer) CustomerVerify(context.Context, *VerifyCustomerRegisterRequest) (*CustomerAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerVerify not implemented")
}
func (UnimplementedCustomerAuthServiceServer) CustomerLogin(context.Context, *CustomerLoginRequest) (*CustomerAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerLogin not implemented")
}
func (UnimplementedCustomerAuthServiceServer) CustomerForgotPassword(context.Context, *Email) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerForgotPassword not implemented")
}
func (UnimplementedCustomerAuthServiceServer) VerifyCustomerForgotPassword(context.Context, *VerifyCustomerRegisterRequest) (*CustomerAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCustomerForgotPassword not implemented")
}
func (UnimplementedCustomerAuthServiceServer) UpdateCustomerPassword(context.Context, *UpdatePasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerPassword not implemented")
}
func (UnimplementedCustomerAuthServiceServer) mustEmbedUnimplementedCustomerAuthServiceServer() {}

// UnsafeCustomerAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerAuthServiceServer will
// result in compilation errors.
type UnsafeCustomerAuthServiceServer interface {
	mustEmbedUnimplementedCustomerAuthServiceServer()
}

func RegisterCustomerAuthServiceServer(s grpc.ServiceRegistrar, srv CustomerAuthServiceServer) {
	s.RegisterService(&CustomerAuthService_ServiceDesc, srv)
}

func _CustomerAuthService_CustomerRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).CustomerRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerAuthService/CustomerRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).CustomerRegister(ctx, req.(*CustomerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_CustomerVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCustomerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).CustomerVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerAuthService/CustomerVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).CustomerVerify(ctx, req.(*VerifyCustomerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_CustomerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).CustomerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerAuthService/CustomerLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).CustomerLogin(ctx, req.(*CustomerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_CustomerForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).CustomerForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerAuthService/CustomerForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).CustomerForgotPassword(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_VerifyCustomerForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCustomerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).VerifyCustomerForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerAuthService/VerifyCustomerForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).VerifyCustomerForgotPassword(ctx, req.(*VerifyCustomerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_UpdateCustomerPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).UpdateCustomerPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CustomerAuthService/UpdateCustomerPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).UpdateCustomerPassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerAuthService_ServiceDesc is the grpc.ServiceDesc for CustomerAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CustomerAuthService",
	HandlerType: (*CustomerAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CustomerRegister",
			Handler:    _CustomerAuthService_CustomerRegister_Handler,
		},
		{
			MethodName: "CustomerVerify",
			Handler:    _CustomerAuthService_CustomerVerify_Handler,
		},
		{
			MethodName: "CustomerLogin",
			Handler:    _CustomerAuthService_CustomerLogin_Handler,
		},
		{
			MethodName: "CustomerForgotPassword",
			Handler:    _CustomerAuthService_CustomerForgotPassword_Handler,
		},
		{
			MethodName: "VerifyCustomerForgotPassword",
			Handler:    _CustomerAuthService_VerifyCustomerForgotPassword_Handler,
		},
		{
			MethodName: "UpdateCustomerPassword",
			Handler:    _CustomerAuthService_UpdateCustomerPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer_auth.proto",
}
