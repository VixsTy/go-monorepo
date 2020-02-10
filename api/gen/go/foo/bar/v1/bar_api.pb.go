// Licensed to go-monorepo under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. go-monorepo licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo/bar/v1/bar_api.proto

package barv1

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PushRequest struct {
	Bar                  *Bar     `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cb6ac3dc7142747, []int{0}
}

func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}

func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}

func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}

func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}

func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

type PushResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushResponse) Reset()         { *m = PushResponse{} }
func (m *PushResponse) String() string { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()    {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cb6ac3dc7142747, []int{1}
}

func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushResponse.Unmarshal(m, b)
}

func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
}

func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}

func (m *PushResponse) XXX_Size() int {
	return xxx_messageInfo_PushResponse.Size(m)
}

func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

func (m *PushResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *PushResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "foo.bar.v1.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "foo.bar.v1.PushResponse")
}

func init() { proto.RegisterFile("foo/bar/v1/bar_api.proto", fileDescriptor_5cb6ac3dc7142747) }

var fileDescriptor_5cb6ac3dc7142747 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x24, 0xad, 0x56, 0xdd, 0x0a, 0xc2, 0x22, 0x1a, 0x7a, 0xaa, 0xbd, 0xd8, 0x4b, 0x37, 0xa6,
	0x1e, 0xbd, 0xd8, 0x05, 0x0b, 0xde, 0x42, 0x0e, 0x45, 0xa4, 0x50, 0xde, 0xd6, 0x6d, 0x12, 0x30,
	0x79, 0xeb, 0xee, 0x26, 0x58, 0x3f, 0xc7, 0xa3, 0x9f, 0xe2, 0x57, 0xc9, 0x26, 0x15, 0x23, 0x3d,
	0xed, 0x0e, 0xf3, 0xe6, 0xcd, 0x9b, 0x21, 0xfe, 0x06, 0x31, 0x10, 0xa0, 0x83, 0x2a, 0x74, 0xcf,
	0x0a, 0x54, 0xc6, 0x94, 0x46, 0x8b, 0x94, 0x6c, 0x10, 0x99, 0x00, 0xcd, 0xaa, 0x70, 0x70, 0xfe,
	0x7f, 0xaa, 0x99, 0x18, 0xcc, 0x92, 0xcc, 0xa6, 0xa5, 0x60, 0x6b, 0xcc, 0x03, 0x59, 0x54, 0xb8,
	0x55, 0x1a, 0xdf, 0xb7, 0x41, 0x4d, 0xae, 0x27, 0x89, 0x2c, 0x26, 0x15, 0xbc, 0x66, 0x2f, 0x60,
	0x65, 0xb0, 0xf7, 0x69, 0x56, 0x8c, 0x6e, 0x48, 0x3f, 0x2a, 0x4d, 0x1a, 0xcb, 0xb7, 0x52, 0x1a,
	0x4b, 0xaf, 0x48, 0x57, 0x80, 0xf6, 0xbd, 0xa1, 0x37, 0xee, 0x4f, 0xcf, 0xd8, 0xdf, 0x05, 0x8c,
	0x83, 0x8e, 0x1d, 0x37, 0xba, 0x27, 0xa7, 0x8d, 0xc2, 0x28, 0x2c, 0x8c, 0xa4, 0x17, 0xa4, 0x67,
	0x2c, 0xd8, 0xd2, 0xd4, 0xaa, 0xe3, 0x78, 0x87, 0xa8, 0x4f, 0x8e, 0x72, 0x69, 0x0c, 0x24, 0xd2,
	0xef, 0x0c, 0xbd, 0xf1, 0x49, 0xfc, 0x0b, 0xa7, 0x0f, 0xa4, 0xc7, 0x41, 0xcf, 0xa2, 0x47, 0x7a,
	0x47, 0x0e, 0xdc, 0x2e, 0x7a, 0xd9, 0x76, 0x6a, 0xdd, 0x33, 0xf0, 0xf7, 0x89, 0xc6, 0x76, 0xec,
	0xf1, 0x15, 0xb9, 0x5e, 0x63, 0xce, 0x76, 0x1d, 0x7c, 0xc8, 0x22, 0xb3, 0x29, 0x68, 0x96, 0x63,
	0x81, 0x5a, 0x2a, 0x6c, 0x29, 0x79, 0xdf, 0xf9, 0xa9, 0x2c, 0x72, 0x91, 0x23, 0xef, 0xf9, 0x50,
	0x80, 0xae, 0xc2, 0xcf, 0x4e, 0x77, 0xce, 0x9f, 0xbe, 0x3a, 0x64, 0x8e, 0xe8, 0xe2, 0xb1, 0x45,
	0xf8, 0x5d, 0x83, 0x25, 0x07, 0xbd, 0x5c, 0x84, 0xa2, 0x57, 0x57, 0x74, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x09, 0x4e, 0xf4, 0xeb, 0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BarAPIClient is the client API for BarAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BarAPIClient interface {
	// Push all requests.
	Push(ctx context.Context, opts ...grpc.CallOption) (BarAPI_PushClient, error)
}

type barAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewBarAPIClient(cc grpc.ClientConnInterface) BarAPIClient {
	return &barAPIClient{cc}
}

func (c *barAPIClient) Push(ctx context.Context, opts ...grpc.CallOption) (BarAPI_PushClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BarAPI_serviceDesc.Streams[0], "/foo.bar.v1.BarAPI/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &barAPIPushClient{stream}
	return x, nil
}

type BarAPI_PushClient interface {
	Send(*PushRequest) error
	CloseAndRecv() (*PushResponse, error)
	grpc.ClientStream
}

type barAPIPushClient struct {
	grpc.ClientStream
}

func (x *barAPIPushClient) Send(m *PushRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *barAPIPushClient) CloseAndRecv() (*PushResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PushResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BarAPIServer is the server API for BarAPI service.
type BarAPIServer interface {
	// Push all requests.
	Push(BarAPI_PushServer) error
}

// UnimplementedBarAPIServer can be embedded to have forward compatible implementations.
type UnimplementedBarAPIServer struct {
}

func (*UnimplementedBarAPIServer) Push(srv BarAPI_PushServer) error {
	return status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterBarAPIServer(s *grpc.Server, srv BarAPIServer) {
	s.RegisterService(&_BarAPI_serviceDesc, srv)
}

func _BarAPI_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarAPIServer).Push(&barAPIPushServer{stream})
}

type BarAPI_PushServer interface {
	SendAndClose(*PushResponse) error
	Recv() (*PushRequest, error)
	grpc.ServerStream
}

type barAPIPushServer struct {
	grpc.ServerStream
}

func (x *barAPIPushServer) SendAndClose(m *PushResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *barAPIPushServer) Recv() (*PushRequest, error) {
	m := new(PushRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BarAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.bar.v1.BarAPI",
	HandlerType: (*BarAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Push",
			Handler:       _BarAPI_Push_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "foo/bar/v1/bar_api.proto",
}