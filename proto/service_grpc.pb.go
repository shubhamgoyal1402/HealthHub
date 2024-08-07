// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/service.proto

package proto

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
	HospitalService_Appointment_FullMethodName    = "/hospital.HospitalService/Appointment"
	HospitalService_GetBookedSlots_FullMethodName = "/hospital.HospitalService/GetBookedSlots"
)

// HospitalServiceClient is the client API for HospitalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalServiceClient interface {
	Appointment(ctx context.Context, in *AppointmentRequest, opts ...grpc.CallOption) (*AppointmentResponse, error)
	GetBookedSlots(ctx context.Context, in *GetBookedSlotsRequest, opts ...grpc.CallOption) (*GetBookedSlotsResponse, error)
}

type hospitalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalServiceClient(cc grpc.ClientConnInterface) HospitalServiceClient {
	return &hospitalServiceClient{cc}
}

func (c *hospitalServiceClient) Appointment(ctx context.Context, in *AppointmentRequest, opts ...grpc.CallOption) (*AppointmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppointmentResponse)
	err := c.cc.Invoke(ctx, HospitalService_Appointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) GetBookedSlots(ctx context.Context, in *GetBookedSlotsRequest, opts ...grpc.CallOption) (*GetBookedSlotsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookedSlotsResponse)
	err := c.cc.Invoke(ctx, HospitalService_GetBookedSlots_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServiceServer is the server API for HospitalService service.
// All implementations must embed UnimplementedHospitalServiceServer
// for forward compatibility
type HospitalServiceServer interface {
	Appointment(context.Context, *AppointmentRequest) (*AppointmentResponse, error)
	GetBookedSlots(context.Context, *GetBookedSlotsRequest) (*GetBookedSlotsResponse, error)
	mustEmbedUnimplementedHospitalServiceServer()
}

// UnimplementedHospitalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServiceServer struct {
}

func (UnimplementedHospitalServiceServer) Appointment(context.Context, *AppointmentRequest) (*AppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Appointment not implemented")
}
func (UnimplementedHospitalServiceServer) GetBookedSlots(context.Context, *GetBookedSlotsRequest) (*GetBookedSlotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookedSlots not implemented")
}
func (UnimplementedHospitalServiceServer) mustEmbedUnimplementedHospitalServiceServer() {}

// UnsafeHospitalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServiceServer will
// result in compilation errors.
type UnsafeHospitalServiceServer interface {
	mustEmbedUnimplementedHospitalServiceServer()
}

func RegisterHospitalServiceServer(s grpc.ServiceRegistrar, srv HospitalServiceServer) {
	s.RegisterService(&HospitalService_ServiceDesc, srv)
}

func _HospitalService_Appointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).Appointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_Appointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).Appointment(ctx, req.(*AppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_GetBookedSlots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookedSlotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).GetBookedSlots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_GetBookedSlots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).GetBookedSlots(ctx, req.(*GetBookedSlotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HospitalService_ServiceDesc is the grpc.ServiceDesc for HospitalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HospitalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.HospitalService",
	HandlerType: (*HospitalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Appointment",
			Handler:    _HospitalService_Appointment_Handler,
		},
		{
			MethodName: "GetBookedSlots",
			Handler:    _HospitalService_GetBookedSlots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
