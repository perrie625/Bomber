// Code generated by protoc-gen-go.
// source: protodata/chat.proto
// DO NOT EDIT!

/*
Package protodata is a generated protocol buffer package.

It is generated from these files:
	protodata/chat.proto
	protodata/flags.proto
	protodata/login.proto

It has these top-level messages:
	SayMessage
	FlagNum
	LoginRequest
	LoginResponse
*/
package protodata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SayMessage struct {
	Words string `protobuf:"bytes,1,opt,name=words" json:"words,omitempty"`
}

func (m *SayMessage) Reset()                    { *m = SayMessage{} }
func (m *SayMessage) String() string            { return proto.CompactTextString(m) }
func (*SayMessage) ProtoMessage()               {}
func (*SayMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SayMessage) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

func init() {
	proto.RegisterType((*SayMessage)(nil), "protodata.SayMessage")
}

func init() { proto.RegisterFile("protodata/chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0xce, 0x48, 0x2c, 0xd1, 0x03, 0x73, 0x85, 0x38, 0xe1,
	0xa2, 0x4a, 0x4a, 0x5c, 0x5c, 0xc1, 0x89, 0x95, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42,
	0x22, 0x5c, 0xac, 0xe5, 0xf9, 0x45, 0x29, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10,
	0x4e, 0x12, 0x1b, 0x58, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x11, 0xbd, 0x54, 0x77, 0x4d,
	0x00, 0x00, 0x00,
}
