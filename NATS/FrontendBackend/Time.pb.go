// Code generated by protoc-gen-go.
// source: Time.proto
// DO NOT EDIT!

/*
Package Transport is a generated protocol buffer package.

It is generated from these files:
	Time.proto
	User.proto

It has these top-level messages:
	Time
	User
*/
package Transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Time struct {
	Time string `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Time)(nil), "Transport.Time")
}

var fileDescriptor0 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0c, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f,
	0x2a, 0x51, 0x92, 0xe2, 0x62, 0x01, 0x49, 0x08, 0x09, 0x71, 0xb1, 0x94, 0x00, 0x69, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x3b, 0x89, 0x0d, 0xac, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x30, 0x7b, 0x71, 0xdb, 0x3b, 0x00, 0x00, 0x00,
}
