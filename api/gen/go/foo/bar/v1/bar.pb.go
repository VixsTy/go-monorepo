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
// source: foo/bar/v1/bar.proto

package barv1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Bar is an entity of foo.bar package.
type Bar struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd7fd7a9b68caf1, []int{0}
}

func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}

func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}

func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}

func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}

func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Bar)(nil), "foo.bar.v1.Bar")
}

func init() { proto.RegisterFile("foo/bar/v1/bar.proto", fileDescriptor_8bd7fd7a9b68caf1) }

var fileDescriptor_8bd7fd7a9b68caf1 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcb, 0xcf, 0xd7,
	0x4f, 0x4a, 0x2c, 0xd2, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c,
	0x69, 0xf9, 0xf9, 0x7a, 0x20, 0x6e, 0x99, 0xa1, 0x92, 0x3c, 0x17, 0xb3, 0x53, 0x62, 0x91, 0x90,
	0x04, 0x17, 0x7b, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b,
	0x10, 0x8c, 0xeb, 0x14, 0xc3, 0xa5, 0x9e, 0x9c, 0x9f, 0xab, 0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x57, 0x95, 0x9a, 0x97, 0x59, 0x92, 0x91, 0x58, 0xa4, 0x97, 0x9b, 0x9f, 0x97, 0x5f, 0x94,
	0x5a, 0x90, 0xaf, 0x87, 0x30, 0xcb, 0x89, 0xc3, 0x29, 0xb1, 0x28, 0x00, 0x64, 0x43, 0x00, 0x63,
	0x14, 0x6b, 0x52, 0x62, 0x51, 0x99, 0xe1, 0x22, 0x26, 0x66, 0x37, 0xa7, 0x88, 0x55, 0x4c, 0x5c,
	0x6e, 0xf9, 0xf9, 0x7a, 0x4e, 0x89, 0x45, 0x7a, 0x61, 0x86, 0xa7, 0xc0, 0x9c, 0x18, 0xa7, 0xc4,
	0xa2, 0x98, 0x30, 0xc3, 0x24, 0x36, 0xb0, 0x8b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11,
	0x49, 0x93, 0xfb, 0xa9, 0x00, 0x00, 0x00,
}
