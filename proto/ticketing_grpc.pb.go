// ticketing.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/ticketing.proto

package proto

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

const (
	TicketService_PurchaseTicket_FullMethodName    = "/ticketing.TicketService/PurchaseTicket"
	TicketService_GetReceiptDetails_FullMethodName = "/ticketing.TicketService/GetReceiptDetails"
	TicketService_ViewUsersAndSeats_FullMethodName = "/ticketing.TicketService/ViewUsersAndSeats"
	TicketService_RemoveUser_FullMethodName        = "/ticketing.TicketService/RemoveUser"
	TicketService_ModifyUserSeat_FullMethodName    = "/ticketing.TicketService/ModifyUserSeat"
)

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*Receipt, error)
	GetReceiptDetails(ctx context.Context, in *User, opts ...grpc.CallOption) (*Receipt, error)
	ViewUsersAndSeats(ctx context.Context, in *Section, opts ...grpc.CallOption) (*UserSeatList, error)
	RemoveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	ModifyUserSeat(ctx context.Context, in *UserSeatModification, opts ...grpc.CallOption) (*Empty, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TicketService_PurchaseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetReceiptDetails(ctx context.Context, in *User, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TicketService_GetReceiptDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ViewUsersAndSeats(ctx context.Context, in *Section, opts ...grpc.CallOption) (*UserSeatList, error) {
	out := new(UserSeatList)
	err := c.cc.Invoke(ctx, TicketService_ViewUsersAndSeats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TicketService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifyUserSeat(ctx context.Context, in *UserSeatModification, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TicketService_ModifyUserSeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	PurchaseTicket(context.Context, *TicketRequest) (*Receipt, error)
	GetReceiptDetails(context.Context, *User) (*Receipt, error)
	ViewUsersAndSeats(context.Context, *Section) (*UserSeatList, error)
	RemoveUser(context.Context, *User) (*Empty, error)
	ModifyUserSeat(context.Context, *UserSeatModification) (*Empty, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) PurchaseTicket(context.Context, *TicketRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTicketServiceServer) GetReceiptDetails(context.Context, *User) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceiptDetails not implemented")
}
func (UnimplementedTicketServiceServer) ViewUsersAndSeats(context.Context, *Section) (*UserSeatList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewUsersAndSeats not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *User) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ModifyUserSeat(context.Context, *UserSeatModification) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserSeat not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetReceiptDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetReceiptDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_GetReceiptDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetReceiptDetails(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ViewUsersAndSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Section)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ViewUsersAndSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ViewUsersAndSeats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ViewUsersAndSeats(ctx, req.(*Section))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifyUserSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSeatModification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifyUserSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ModifyUserSeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifyUserSeat(ctx, req.(*UserSeatModification))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticketing.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TicketService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceiptDetails",
			Handler:    _TicketService_GetReceiptDetails_Handler,
		},
		{
			MethodName: "ViewUsersAndSeats",
			Handler:    _TicketService_ViewUsersAndSeats_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifyUserSeat",
			Handler:    _TicketService_ModifyUserSeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ticketing.proto",
}
