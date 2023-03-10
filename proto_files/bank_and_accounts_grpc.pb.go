// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto_files/bank_and_accounts.proto

package proto_files

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountParams, opts ...grpc.CallOption) (*Account, error)
	GetAccounts(ctx context.Context, in *GetAccountsParams, opts ...grpc.CallOption) (AccountService_GetAccountsClient, error)
	GenerateAddress(ctx context.Context, in *GenerateAddressParams, opts ...grpc.CallOption) (*Account, error)
	Deposit(ctx context.Context, in *DepositParams, opts ...grpc.CallOption) (*Account, error)
	Withdrawal(ctx context.Context, in *WithdrawalParams, opts ...grpc.CallOption) (*Account, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountParams, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/bank_and_accounts.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *GetAccountsParams, opts ...grpc.CallOption) (AccountService_GetAccountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccountService_ServiceDesc.Streams[0], "/bank_and_accounts.AccountService/GetAccounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountServiceGetAccountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountService_GetAccountsClient interface {
	Recv() (*Account, error)
	grpc.ClientStream
}

type accountServiceGetAccountsClient struct {
	grpc.ClientStream
}

func (x *accountServiceGetAccountsClient) Recv() (*Account, error) {
	m := new(Account)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountServiceClient) GenerateAddress(ctx context.Context, in *GenerateAddressParams, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/bank_and_accounts.AccountService/GenerateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Deposit(ctx context.Context, in *DepositParams, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/bank_and_accounts.AccountService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Withdrawal(ctx context.Context, in *WithdrawalParams, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/bank_and_accounts.AccountService/Withdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountParams) (*Account, error)
	GetAccounts(*GetAccountsParams, AccountService_GetAccountsServer) error
	GenerateAddress(context.Context, *GenerateAddressParams) (*Account, error)
	Deposit(context.Context, *DepositParams) (*Account, error)
	Withdrawal(context.Context, *WithdrawalParams) (*Account, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountParams) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) GetAccounts(*GetAccountsParams, AccountService_GetAccountsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedAccountServiceServer) GenerateAddress(context.Context, *GenerateAddressParams) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAddress not implemented")
}
func (UnimplementedAccountServiceServer) Deposit(context.Context, *DepositParams) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedAccountServiceServer) Withdrawal(context.Context, *WithdrawalParams) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdrawal not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_and_accounts.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAccountsParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountServiceServer).GetAccounts(m, &accountServiceGetAccountsServer{stream})
}

type AccountService_GetAccountsServer interface {
	Send(*Account) error
	grpc.ServerStream
}

type accountServiceGetAccountsServer struct {
	grpc.ServerStream
}

func (x *accountServiceGetAccountsServer) Send(m *Account) error {
	return x.ServerStream.SendMsg(m)
}

func _AccountService_GenerateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAddressParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GenerateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_and_accounts.AccountService/GenerateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GenerateAddress(ctx, req.(*GenerateAddressParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_and_accounts.AccountService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Deposit(ctx, req.(*DepositParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Withdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawalParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Withdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_and_accounts.AccountService/Withdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Withdrawal(ctx, req.(*WithdrawalParams))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank_and_accounts.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "GenerateAddress",
			Handler:    _AccountService_GenerateAddress_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _AccountService_Deposit_Handler,
		},
		{
			MethodName: "Withdrawal",
			Handler:    _AccountService_Withdrawal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAccounts",
			Handler:       _AccountService_GetAccounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto_files/bank_and_accounts.proto",
}
