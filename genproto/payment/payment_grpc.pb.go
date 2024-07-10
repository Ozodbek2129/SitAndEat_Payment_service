// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
<<<<<<< HEAD
// - protoc             v3.12.4
// source: payment.proto
=======
// - protoc             v3.18.0
// source: SitAndEat_protos/payment/payment.proto
>>>>>>> origin/ozodbek

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Payment_CreatePayments_FullMethodName       = "/payment.Payment/CreatePayments"
	Payment_GetPaymentStatusById_FullMethodName = "/payment.Payment/GetPaymentStatusById"
	Payment_UpdatePayments_FullMethodName       = "/payment.Payment/UpdatePayments"
)

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	CreatePayments(ctx context.Context, in *CreatePayment, opts ...grpc.CallOption) (*Status, error)
	GetPaymentStatusById(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*GetByIdResponse, error)
	UpdatePayments(ctx context.Context, in *UpdatePayment, opts ...grpc.CallOption) (*Status, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) CreatePayments(ctx context.Context, in *CreatePayment, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, Payment_CreatePayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetPaymentStatusById(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, Payment_GetPaymentStatusById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) UpdatePayments(ctx context.Context, in *UpdatePayment, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, Payment_UpdatePayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility
type PaymentServer interface {
	CreatePayments(context.Context, *CreatePayment) (*Status, error)
	GetPaymentStatusById(context.Context, *GetById) (*GetByIdResponse, error)
	UpdatePayments(context.Context, *UpdatePayment) (*Status, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (UnimplementedPaymentServer) CreatePayments(context.Context, *CreatePayment) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayments not implemented")
}
func (UnimplementedPaymentServer) GetPaymentStatusById(context.Context, *GetById) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentStatusById not implemented")
}
func (UnimplementedPaymentServer) UpdatePayments(context.Context, *UpdatePayment) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayments not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_CreatePayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).CreatePayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_CreatePayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).CreatePayments(ctx, req.(*CreatePayment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetPaymentStatusById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetPaymentStatusById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetPaymentStatusById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetPaymentStatusById(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_UpdatePayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).UpdatePayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_UpdatePayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).UpdatePayments(ctx, req.(*UpdatePayment))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayments",
			Handler:    _Payment_CreatePayments_Handler,
		},
		{
			MethodName: "GetPaymentStatusById",
			Handler:    _Payment_GetPaymentStatusById_Handler,
		},
		{
			MethodName: "UpdatePayments",
			Handler:    _Payment_UpdatePayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SitAndEat_protos/payment/payment.proto",
}
