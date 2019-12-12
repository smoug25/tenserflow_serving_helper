// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/config/logging_config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SamplingConfig struct {
	// Requests will be logged uniformly at random with this probability. Valid
	// range: [0, 1.0].
	SamplingRate         float64  `protobuf:"fixed64,1,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SamplingConfig) Reset()         { *m = SamplingConfig{} }
func (m *SamplingConfig) String() string { return proto.CompactTextString(m) }
func (*SamplingConfig) ProtoMessage()    {}
func (*SamplingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7295619f292c8c, []int{0}
}

func (m *SamplingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SamplingConfig.Unmarshal(m, b)
}
func (m *SamplingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SamplingConfig.Marshal(b, m, deterministic)
}
func (m *SamplingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SamplingConfig.Merge(m, src)
}
func (m *SamplingConfig) XXX_Size() int {
	return xxx_messageInfo_SamplingConfig.Size(m)
}
func (m *SamplingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SamplingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SamplingConfig proto.InternalMessageInfo

func (m *SamplingConfig) GetSamplingRate() float64 {
	if m != nil {
		return m.SamplingRate
	}
	return 0
}

// Configuration for logging query/responses.
type LoggingConfig struct {
	LogCollectorConfig   *LogCollectorConfig `protobuf:"bytes,1,opt,name=log_collector_config,json=logCollectorConfig,proto3" json:"log_collector_config,omitempty"`
	SamplingConfig       *SamplingConfig     `protobuf:"bytes,2,opt,name=sampling_config,json=samplingConfig,proto3" json:"sampling_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoggingConfig) Reset()         { *m = LoggingConfig{} }
func (m *LoggingConfig) String() string { return proto.CompactTextString(m) }
func (*LoggingConfig) ProtoMessage()    {}
func (*LoggingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7295619f292c8c, []int{1}
}

func (m *LoggingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoggingConfig.Unmarshal(m, b)
}
func (m *LoggingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoggingConfig.Marshal(b, m, deterministic)
}
func (m *LoggingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggingConfig.Merge(m, src)
}
func (m *LoggingConfig) XXX_Size() int {
	return xxx_messageInfo_LoggingConfig.Size(m)
}
func (m *LoggingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoggingConfig proto.InternalMessageInfo

func (m *LoggingConfig) GetLogCollectorConfig() *LogCollectorConfig {
	if m != nil {
		return m.LogCollectorConfig
	}
	return nil
}

func (m *LoggingConfig) GetSamplingConfig() *SamplingConfig {
	if m != nil {
		return m.SamplingConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*SamplingConfig)(nil), "tensorflow.serving.SamplingConfig")
	proto.RegisterType((*LoggingConfig)(nil), "tensorflow.serving.LoggingConfig")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/config/logging_config.proto", fileDescriptor_0f7295619f292c8c)
}

var fileDescriptor_0f7295619f292c8c = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x8a, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcd, 0x2f, 0x4d, 0x37, 0x32, 0xd5, 0x2f,
	0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc,
	0x4b, 0x8f, 0xcf, 0x48, 0xcd, 0x29, 0x48, 0x2d, 0xc2, 0x22, 0xa3, 0x9f, 0x9c, 0x9f, 0x97, 0x96,
	0x99, 0xae, 0x9f, 0x93, 0x9f, 0x9e, 0x0e, 0x52, 0x08, 0xe1, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0x09, 0x21, 0xd4, 0xeb, 0x41, 0xd5, 0x4b, 0x25, 0x52, 0xcf, 0xc2, 0xf8, 0xe4, 0xfc, 0x9c,
	0x9c, 0xd4, 0xe4, 0x92, 0xfc, 0x22, 0x14, 0x6b, 0x95, 0x4c, 0xb9, 0xf8, 0x82, 0x13, 0x73, 0x0b,
	0x72, 0x32, 0xf3, 0xd2, 0x9d, 0xc1, 0xe2, 0x42, 0xca, 0x5c, 0xbc, 0xc5, 0x50, 0x91, 0xf8, 0xa2,
	0xc4, 0x92, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xc6, 0x20, 0x1e, 0x98, 0x60, 0x50, 0x62, 0x49,
	0xaa, 0xd2, 0x36, 0x46, 0x2e, 0x5e, 0x1f, 0x88, 0x37, 0xa0, 0xda, 0x22, 0xb8, 0x44, 0xb0, 0x59,
	0x03, 0xd6, 0xcd, 0x6d, 0xa4, 0xa6, 0x87, 0xe9, 0x3d, 0x3d, 0x9f, 0xfc, 0x74, 0x67, 0x98, 0x72,
	0x88, 0x29, 0x41, 0x42, 0x39, 0x18, 0x62, 0x42, 0xde, 0x5c, 0xfc, 0x70, 0x07, 0x41, 0x0d, 0x65,
	0x02, 0x1b, 0xaa, 0x84, 0xcd, 0x50, 0x54, 0xdf, 0x04, 0xf1, 0x15, 0xa3, 0xf0, 0x9d, 0xbc, 0xa3,
	0xdc, 0xa8, 0x13, 0xa8, 0x3f, 0x18, 0x19, 0x93, 0xd8, 0xc0, 0x61, 0x68, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0x07, 0x9d, 0x85, 0x19, 0x02, 0x00, 0x00,
}
