// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: auth.proto

package sso

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Register(ctx context.Context, in *RegisterSchema, opts ...grpc.CallOption) (*TimeoutSchema, error)
	GetCurrentUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserSchema, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ForgotPassword(ctx context.Context, in *EmailSchema, opts ...grpc.CallOption) (*TimeoutSchema, error)
	ConfirmResetPassword(ctx context.Context, in *EmailCodeSchema, opts ...grpc.CallOption) (*AccessToken, error)
	ResetPassword(ctx context.Context, in *PasswordSchema, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckEmail(ctx context.Context, in *EmailSchema, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConfirmEmail(ctx context.Context, in *EmailCodeSchema, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EmailLogin(ctx context.Context, in *EmailPasswordSchema, opts ...grpc.CallOption) (*AccessToken, error)
	SendEmailConfirmation(ctx context.Context, in *EmailMailSchema, opts ...grpc.CallOption) (*TimeoutSchema, error)
	PhoneLogin(ctx context.Context, in *PhoneCodeSchema, opts ...grpc.CallOption) (*AccessToken, error)
	CheckPhoneNumber(ctx context.Context, in *PhoneSchema, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConfirmPhoneNumber(ctx context.Context, in *PhoneCodeSchema, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendSMSConfirmation(ctx context.Context, in *PhoneSMSSchema, opts ...grpc.CallOption) (*TimeoutSchema, error)
	LogOut(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeactivateAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterSchema, opts ...grpc.CallOption) (*TimeoutSchema, error) {
	out := new(TimeoutSchema)
	err := c.cc.Invoke(ctx, "/sso.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetCurrentUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/sso.AuthService/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUser(ctx context.Context, in *UpdateUserSchema, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *EmailSchema, opts ...grpc.CallOption) (*TimeoutSchema, error) {
	out := new(TimeoutSchema)
	err := c.cc.Invoke(ctx, "/sso.AuthService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ConfirmResetPassword(ctx context.Context, in *EmailCodeSchema, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/sso.AuthService/ConfirmResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *PasswordSchema, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckEmail(ctx context.Context, in *EmailSchema, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/CheckEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ConfirmEmail(ctx context.Context, in *EmailCodeSchema, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/ConfirmEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EmailLogin(ctx context.Context, in *EmailPasswordSchema, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/sso.AuthService/EmailLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SendEmailConfirmation(ctx context.Context, in *EmailMailSchema, opts ...grpc.CallOption) (*TimeoutSchema, error) {
	out := new(TimeoutSchema)
	err := c.cc.Invoke(ctx, "/sso.AuthService/SendEmailConfirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PhoneLogin(ctx context.Context, in *PhoneCodeSchema, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/sso.AuthService/PhoneLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckPhoneNumber(ctx context.Context, in *PhoneSchema, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/CheckPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ConfirmPhoneNumber(ctx context.Context, in *PhoneCodeSchema, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/ConfirmPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SendSMSConfirmation(ctx context.Context, in *PhoneSMSSchema, opts ...grpc.CallOption) (*TimeoutSchema, error) {
	out := new(TimeoutSchema)
	err := c.cc.Invoke(ctx, "/sso.AuthService/SendSMSConfirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LogOut(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeactivateAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/DeactivateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sso.AuthService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Register(context.Context, *RegisterSchema) (*TimeoutSchema, error)
	GetCurrentUser(context.Context, *emptypb.Empty) (*User, error)
	UpdateUser(context.Context, *UpdateUserSchema) (*emptypb.Empty, error)
	ForgotPassword(context.Context, *EmailSchema) (*TimeoutSchema, error)
	ConfirmResetPassword(context.Context, *EmailCodeSchema) (*AccessToken, error)
	ResetPassword(context.Context, *PasswordSchema) (*emptypb.Empty, error)
	CheckEmail(context.Context, *EmailSchema) (*emptypb.Empty, error)
	ConfirmEmail(context.Context, *EmailCodeSchema) (*emptypb.Empty, error)
	EmailLogin(context.Context, *EmailPasswordSchema) (*AccessToken, error)
	SendEmailConfirmation(context.Context, *EmailMailSchema) (*TimeoutSchema, error)
	PhoneLogin(context.Context, *PhoneCodeSchema) (*AccessToken, error)
	CheckPhoneNumber(context.Context, *PhoneSchema) (*emptypb.Empty, error)
	ConfirmPhoneNumber(context.Context, *PhoneCodeSchema) (*emptypb.Empty, error)
	SendSMSConfirmation(context.Context, *PhoneSMSSchema) (*TimeoutSchema, error)
	LogOut(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	DeactivateAccount(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	DeleteAccount(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterSchema) (*TimeoutSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) GetCurrentUser(context.Context, *emptypb.Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUser(context.Context, *UpdateUserSchema) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPassword(context.Context, *EmailSchema) (*TimeoutSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) ConfirmResetPassword(context.Context, *EmailCodeSchema) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *PasswordSchema) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) CheckEmail(context.Context, *EmailSchema) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmail not implemented")
}
func (UnimplementedAuthServiceServer) ConfirmEmail(context.Context, *EmailCodeSchema) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmEmail not implemented")
}
func (UnimplementedAuthServiceServer) EmailLogin(context.Context, *EmailPasswordSchema) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailLogin not implemented")
}
func (UnimplementedAuthServiceServer) SendEmailConfirmation(context.Context, *EmailMailSchema) (*TimeoutSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailConfirmation not implemented")
}
func (UnimplementedAuthServiceServer) PhoneLogin(context.Context, *PhoneCodeSchema) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneLogin not implemented")
}
func (UnimplementedAuthServiceServer) CheckPhoneNumber(context.Context, *PhoneSchema) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPhoneNumber not implemented")
}
func (UnimplementedAuthServiceServer) ConfirmPhoneNumber(context.Context, *PhoneCodeSchema) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPhoneNumber not implemented")
}
func (UnimplementedAuthServiceServer) SendSMSConfirmation(context.Context, *PhoneSMSSchema) (*TimeoutSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMSConfirmation not implemented")
}
func (UnimplementedAuthServiceServer) LogOut(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedAuthServiceServer) DeactivateAccount(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateAccount not implemented")
}
func (UnimplementedAuthServiceServer) DeleteAccount(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetCurrentUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUser(ctx, req.(*UpdateUserSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPassword(ctx, req.(*EmailSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ConfirmResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailCodeSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ConfirmResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/ConfirmResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ConfirmResetPassword(ctx, req.(*EmailCodeSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*PasswordSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/CheckEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckEmail(ctx, req.(*EmailSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ConfirmEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailCodeSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ConfirmEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/ConfirmEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ConfirmEmail(ctx, req.(*EmailCodeSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EmailLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailPasswordSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EmailLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/EmailLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EmailLogin(ctx, req.(*EmailPasswordSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SendEmailConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailMailSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SendEmailConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/SendEmailConfirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SendEmailConfirmation(ctx, req.(*EmailMailSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PhoneLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneCodeSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PhoneLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/PhoneLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PhoneLogin(ctx, req.(*PhoneCodeSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/CheckPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckPhoneNumber(ctx, req.(*PhoneSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ConfirmPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneCodeSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ConfirmPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/ConfirmPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ConfirmPhoneNumber(ctx, req.(*PhoneCodeSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SendSMSConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneSMSSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SendSMSConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/SendSMSConfirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SendSMSConfirmation(ctx, req.(*PhoneSMSSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LogOut(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeactivateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeactivateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/DeactivateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeactivateAccount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sso.AuthService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteAccount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sso.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _AuthService_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthService_UpdateUser_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _AuthService_ForgotPassword_Handler,
		},
		{
			MethodName: "ConfirmResetPassword",
			Handler:    _AuthService_ConfirmResetPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "CheckEmail",
			Handler:    _AuthService_CheckEmail_Handler,
		},
		{
			MethodName: "ConfirmEmail",
			Handler:    _AuthService_ConfirmEmail_Handler,
		},
		{
			MethodName: "EmailLogin",
			Handler:    _AuthService_EmailLogin_Handler,
		},
		{
			MethodName: "SendEmailConfirmation",
			Handler:    _AuthService_SendEmailConfirmation_Handler,
		},
		{
			MethodName: "PhoneLogin",
			Handler:    _AuthService_PhoneLogin_Handler,
		},
		{
			MethodName: "CheckPhoneNumber",
			Handler:    _AuthService_CheckPhoneNumber_Handler,
		},
		{
			MethodName: "ConfirmPhoneNumber",
			Handler:    _AuthService_ConfirmPhoneNumber_Handler,
		},
		{
			MethodName: "SendSMSConfirmation",
			Handler:    _AuthService_SendSMSConfirmation_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _AuthService_LogOut_Handler,
		},
		{
			MethodName: "DeactivateAccount",
			Handler:    _AuthService_DeactivateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AuthService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
