// Code generated by protoc-gen-gogo.
// source: client/pkg/config/config.proto
// DO NOT EDIT!

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	client/pkg/config/config.proto

It has these top-level messages:
	Config
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	UserID string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "Config")
}

func init() { proto.RegisterFile("client/pkg/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x84, 0x51, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11,
	0x55, 0xd2, 0xe5, 0x62, 0x73, 0x06, 0xab, 0x12, 0x52, 0xe6, 0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a,
	0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xe2, 0x7a, 0x74, 0x4f, 0x9e, 0x2d, 0xb4,
	0x38, 0xb5, 0xc8, 0xd3, 0x25, 0x88, 0x0d, 0x24, 0xe5, 0x99, 0x92, 0xc4, 0x06, 0xd6, 0x65, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x04, 0xa5, 0x4c, 0x6d, 0x00, 0x00, 0x00,
}