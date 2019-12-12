// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/lib/core/error_codes.proto

package error_codes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf"
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

// Code from public import github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/error_codes.proto
type Code = protobuf.Code

var Code_name = protobuf.Code_name
var Code_value = protobuf.Code_value

const Code_OK = Code(protobuf.Code_OK)
const Code_CANCELLED = Code(protobuf.Code_CANCELLED)
const Code_UNKNOWN = Code(protobuf.Code_UNKNOWN)
const Code_INVALID_ARGUMENT = Code(protobuf.Code_INVALID_ARGUMENT)
const Code_DEADLINE_EXCEEDED = Code(protobuf.Code_DEADLINE_EXCEEDED)
const Code_NOT_FOUND = Code(protobuf.Code_NOT_FOUND)
const Code_ALREADY_EXISTS = Code(protobuf.Code_ALREADY_EXISTS)
const Code_PERMISSION_DENIED = Code(protobuf.Code_PERMISSION_DENIED)
const Code_UNAUTHENTICATED = Code(protobuf.Code_UNAUTHENTICATED)
const Code_RESOURCE_EXHAUSTED = Code(protobuf.Code_RESOURCE_EXHAUSTED)
const Code_FAILED_PRECONDITION = Code(protobuf.Code_FAILED_PRECONDITION)
const Code_ABORTED = Code(protobuf.Code_ABORTED)
const Code_OUT_OF_RANGE = Code(protobuf.Code_OUT_OF_RANGE)
const Code_UNIMPLEMENTED = Code(protobuf.Code_UNIMPLEMENTED)
const Code_INTERNAL = Code(protobuf.Code_INTERNAL)
const Code_UNAVAILABLE = Code(protobuf.Code_UNAVAILABLE)
const Code_DATA_LOSS = Code(protobuf.Code_DATA_LOSS)
const Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_ = Code(protobuf.Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_)

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/lib/core/error_codes.proto", fileDescriptor_63234f20e8884c2e)
}

var fileDescriptor_63234f20e8884c2e = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcd, 0x2f, 0x4d, 0x37, 0x32, 0xd5, 0x2f,
	0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc,
	0x4b, 0x8f, 0xcf, 0x48, 0xcd, 0x29, 0x48, 0x2d, 0x42, 0x92, 0xd1, 0x4f, 0xce, 0x2f, 0x4a, 0xd5,
	0xcf, 0xc9, 0x4c, 0x82, 0x30, 0x52, 0x8b, 0x8a, 0xf2, 0x8b, 0xe2, 0x93, 0xf3, 0x53, 0x52, 0x8b,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0xa5, 0x28, 0x33, 0x18, 0x6c, 0x46, 0x52, 0x69, 0x1a, 0xa6,
	0xc1, 0x01, 0x0c, 0x49, 0x6c, 0x60, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xa7, 0x05,
	0xa2, 0xbc, 0x00, 0x00, 0x00,
}
