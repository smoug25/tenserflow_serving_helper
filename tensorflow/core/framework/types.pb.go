// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework/types.proto

package framework

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

// (== suppress_warning documentation-presence ==)
// LINT.IfChange
type DataType int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	DataType_DT_INVALID DataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	DataType_DT_FLOAT      DataType = 1
	DataType_DT_DOUBLE     DataType = 2
	DataType_DT_INT32      DataType = 3
	DataType_DT_UINT8      DataType = 4
	DataType_DT_INT16      DataType = 5
	DataType_DT_INT8       DataType = 6
	DataType_DT_STRING     DataType = 7
	DataType_DT_COMPLEX64  DataType = 8
	DataType_DT_INT64      DataType = 9
	DataType_DT_BOOL       DataType = 10
	DataType_DT_QINT8      DataType = 11
	DataType_DT_QUINT8     DataType = 12
	DataType_DT_QINT32     DataType = 13
	DataType_DT_BFLOAT16   DataType = 14
	DataType_DT_QINT16     DataType = 15
	DataType_DT_QUINT16    DataType = 16
	DataType_DT_UINT16     DataType = 17
	DataType_DT_COMPLEX128 DataType = 18
	DataType_DT_HALF       DataType = 19
	DataType_DT_RESOURCE   DataType = 20
	DataType_DT_VARIANT    DataType = 21
	DataType_DT_UINT32     DataType = 22
	DataType_DT_UINT64     DataType = 23
	// Do not use!  These are only for parameters.  Every enum above
	// should have a corresponding value below (verified by types_test).
	DataType_DT_FLOAT_REF      DataType = 101
	DataType_DT_DOUBLE_REF     DataType = 102
	DataType_DT_INT32_REF      DataType = 103
	DataType_DT_UINT8_REF      DataType = 104
	DataType_DT_INT16_REF      DataType = 105
	DataType_DT_INT8_REF       DataType = 106
	DataType_DT_STRING_REF     DataType = 107
	DataType_DT_COMPLEX64_REF  DataType = 108
	DataType_DT_INT64_REF      DataType = 109
	DataType_DT_BOOL_REF       DataType = 110
	DataType_DT_QINT8_REF      DataType = 111
	DataType_DT_QUINT8_REF     DataType = 112
	DataType_DT_QINT32_REF     DataType = 113
	DataType_DT_BFLOAT16_REF   DataType = 114
	DataType_DT_QINT16_REF     DataType = 115
	DataType_DT_QUINT16_REF    DataType = 116
	DataType_DT_UINT16_REF     DataType = 117
	DataType_DT_COMPLEX128_REF DataType = 118
	DataType_DT_HALF_REF       DataType = 119
	DataType_DT_RESOURCE_REF   DataType = 120
	DataType_DT_VARIANT_REF    DataType = 121
	DataType_DT_UINT32_REF     DataType = 122
	DataType_DT_UINT64_REF     DataType = 123
)

var DataType_name = map[int32]string{
	0:   "DT_INVALID",
	1:   "DT_FLOAT",
	2:   "DT_DOUBLE",
	3:   "DT_INT32",
	4:   "DT_UINT8",
	5:   "DT_INT16",
	6:   "DT_INT8",
	7:   "DT_STRING",
	8:   "DT_COMPLEX64",
	9:   "DT_INT64",
	10:  "DT_BOOL",
	11:  "DT_QINT8",
	12:  "DT_QUINT8",
	13:  "DT_QINT32",
	14:  "DT_BFLOAT16",
	15:  "DT_QINT16",
	16:  "DT_QUINT16",
	17:  "DT_UINT16",
	18:  "DT_COMPLEX128",
	19:  "DT_HALF",
	20:  "DT_RESOURCE",
	21:  "DT_VARIANT",
	22:  "DT_UINT32",
	23:  "DT_UINT64",
	101: "DT_FLOAT_REF",
	102: "DT_DOUBLE_REF",
	103: "DT_INT32_REF",
	104: "DT_UINT8_REF",
	105: "DT_INT16_REF",
	106: "DT_INT8_REF",
	107: "DT_STRING_REF",
	108: "DT_COMPLEX64_REF",
	109: "DT_INT64_REF",
	110: "DT_BOOL_REF",
	111: "DT_QINT8_REF",
	112: "DT_QUINT8_REF",
	113: "DT_QINT32_REF",
	114: "DT_BFLOAT16_REF",
	115: "DT_QINT16_REF",
	116: "DT_QUINT16_REF",
	117: "DT_UINT16_REF",
	118: "DT_COMPLEX128_REF",
	119: "DT_HALF_REF",
	120: "DT_RESOURCE_REF",
	121: "DT_VARIANT_REF",
	122: "DT_UINT32_REF",
	123: "DT_UINT64_REF",
}

var DataType_value = map[string]int32{
	"DT_INVALID":        0,
	"DT_FLOAT":          1,
	"DT_DOUBLE":         2,
	"DT_INT32":          3,
	"DT_UINT8":          4,
	"DT_INT16":          5,
	"DT_INT8":           6,
	"DT_STRING":         7,
	"DT_COMPLEX64":      8,
	"DT_INT64":          9,
	"DT_BOOL":           10,
	"DT_QINT8":          11,
	"DT_QUINT8":         12,
	"DT_QINT32":         13,
	"DT_BFLOAT16":       14,
	"DT_QINT16":         15,
	"DT_QUINT16":        16,
	"DT_UINT16":         17,
	"DT_COMPLEX128":     18,
	"DT_HALF":           19,
	"DT_RESOURCE":       20,
	"DT_VARIANT":        21,
	"DT_UINT32":         22,
	"DT_UINT64":         23,
	"DT_FLOAT_REF":      101,
	"DT_DOUBLE_REF":     102,
	"DT_INT32_REF":      103,
	"DT_UINT8_REF":      104,
	"DT_INT16_REF":      105,
	"DT_INT8_REF":       106,
	"DT_STRING_REF":     107,
	"DT_COMPLEX64_REF":  108,
	"DT_INT64_REF":      109,
	"DT_BOOL_REF":       110,
	"DT_QINT8_REF":      111,
	"DT_QUINT8_REF":     112,
	"DT_QINT32_REF":     113,
	"DT_BFLOAT16_REF":   114,
	"DT_QINT16_REF":     115,
	"DT_QUINT16_REF":    116,
	"DT_UINT16_REF":     117,
	"DT_COMPLEX128_REF": 118,
	"DT_HALF_REF":       119,
	"DT_RESOURCE_REF":   120,
	"DT_VARIANT_REF":    121,
	"DT_UINT32_REF":     122,
	"DT_UINT64_REF":     123,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fb14d58823f6c5f, []int{0}
}

func init() {
	proto.RegisterEnum("tensorflow.DataType", DataType_name, DataType_value)
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework/types.proto", fileDescriptor_1fb14d58823f6c5f)
}

var fileDescriptor_1fb14d58823f6c5f = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x93, 0xd2, 0x4c,
	0x10, 0xc6, 0x5f, 0x5e, 0x95, 0x65, 0x87, 0x7f, 0xcd, 0xec, 0xae, 0xfa, 0x19, 0x3c, 0x40, 0x05,
	0x30, 0xc5, 0x15, 0x36, 0x41, 0x53, 0x15, 0x13, 0xc8, 0x0e, 0x5b, 0x96, 0x97, 0x14, 0xbb, 0x15,
	0x02, 0x2e, 0x30, 0x31, 0x09, 0x8b, 0xe8, 0x37, 0xf2, 0x13, 0x7a, 0xb4, 0xe8, 0x34, 0xcc, 0x78,
	0xf7, 0xd8, 0xbf, 0xee, 0x79, 0xfa, 0x99, 0xee, 0x6a, 0x16, 0xc4, 0xab, 0x7c, 0xb9, 0x7b, 0x68,
	0x3f, 0xca, 0x4d, 0x27, 0xdb, 0xc8, 0x5d, 0xdc, 0x7d, 0xdf, 0xc9, 0xa3, 0x6d, 0x26, 0xd3, 0xc5,
	0x5a, 0xee, 0xc3, 0x2c, 0x4a, 0x9f, 0x57, 0xdb, 0x38, 0x5c, 0x46, 0xeb, 0x24, 0x4a, 0xb5, 0x4c,
	0xe7, 0x51, 0xa6, 0x51, 0x67, 0x91, 0xce, 0x37, 0xd1, 0x5e, 0xa6, 0x4f, 0x9d, 0xfc, 0x90, 0x44,
	0x59, 0x3b, 0x49, 0x65, 0x2e, 0x39, 0x53, 0x65, 0xef, 0x7e, 0x95, 0x59, 0xc5, 0x9a, 0xe7, 0x73,
	0x71, 0x48, 0x22, 0xde, 0x60, 0xcc, 0x12, 0xa1, 0xe3, 0xdd, 0x0f, 0x5d, 0xc7, 0x82, 0xff, 0x78,
	0x8d, 0x55, 0x2c, 0x11, 0x8e, 0x5d, 0x7f, 0x28, 0xa0, 0xc4, 0xeb, 0xec, 0xd2, 0x12, 0xa1, 0xe5,
	0xcf, 0x46, 0xae, 0x0d, 0xff, 0x53, 0xd2, 0xf1, 0x44, 0xaf, 0x0b, 0x2f, 0x28, 0x9a, 0x39, 0x9e,
	0x18, 0xc0, 0x4b, 0x95, 0x33, 0x4c, 0x78, 0xc5, 0xab, 0xec, 0xa2, 0x88, 0x06, 0x50, 0x26, 0x95,
	0x3b, 0x11, 0x38, 0xde, 0x07, 0xb8, 0xe0, 0xc0, 0x6a, 0x96, 0x08, 0x6f, 0xfd, 0x4f, 0x13, 0xd7,
	0xfe, 0x6c, 0xf6, 0xa1, 0xa2, 0xde, 0x9a, 0x7d, 0xb8, 0xa4, 0xb7, 0x23, 0xdf, 0x77, 0x81, 0x51,
	0x6a, 0x8a, 0x4a, 0x55, 0x52, 0x9a, 0x16, 0x3d, 0x6b, 0xa7, 0xb0, 0x30, 0x54, 0xe7, 0x4d, 0x56,
	0x3d, 0x3e, 0x44, 0xf3, 0x86, 0x09, 0x0d, 0x2d, 0x6f, 0x98, 0xd0, 0xa4, 0xbf, 0xe2, 0x6b, 0xc3,
	0x04, 0xa0, 0x34, 0x85, 0x2d, 0xde, 0x62, 0x75, 0xe5, 0xcb, 0xe8, 0x0e, 0x80, 0x93, 0x95, 0x8f,
	0x43, 0x77, 0x0c, 0x57, 0x24, 0x1f, 0xd8, 0x77, 0xfe, 0x2c, 0xb8, 0xb5, 0xe1, 0x9a, 0xf4, 0xee,
	0x87, 0x81, 0x33, 0xf4, 0x04, 0xdc, 0x68, 0x7a, 0xbd, 0x2e, 0xbc, 0xd6, 0x42, 0xb3, 0x0f, 0x6f,
	0xe8, 0xdb, 0x68, 0x2e, 0x0c, 0xec, 0x31, 0x44, 0xd4, 0xb0, 0x98, 0x2e, 0xa2, 0x05, 0x15, 0xa1,
	0x02, 0x92, 0x98, 0x08, 0xfe, 0x18, 0xc9, 0x52, 0xd5, 0x18, 0x26, 0x92, 0x15, 0x39, 0x3b, 0x97,
	0x7c, 0x25, 0xe5, 0x62, 0xe2, 0x88, 0x9e, 0xf8, 0x35, 0x03, 0x7d, 0xea, 0x48, 0xd7, 0x4a, 0x8b,
	0xc8, 0xe6, 0x34, 0x44, 0xdf, 0x77, 0x11, 0x6c, 0xa9, 0x64, 0x7a, 0x56, 0x97, 0xa4, 0x3e, 0x55,
	0x9e, 0x92, 0x13, 0x52, 0xc6, 0xbf, 0xf1, 0x2b, 0xd6, 0xd4, 0xb6, 0x81, 0x30, 0xd5, 0xea, 0x08,
	0x65, 0x9c, 0xb3, 0x86, 0xda, 0x0a, 0xb2, 0x9c, 0xca, 0x34, 0xb4, 0xe3, 0x37, 0xac, 0xf5, 0xd7,
	0x76, 0x10, 0x3f, 0x93, 0xdd, 0xe3, 0x86, 0x10, 0xec, 0xa9, 0xed, 0x69, 0x4b, 0x08, 0xbf, 0x53,
	0x0f, 0xda, 0x14, 0xb2, 0x83, 0xd6, 0x83, 0x2c, 0xff, 0xd0, 0x10, 0x8d, 0xe3, 0xe7, 0x28, 0x67,
	0x6f, 0x65, 0x1a, 0xb7, 0xd5, 0xf9, 0xb4, 0xcf, 0x07, 0x36, 0xaa, 0x1e, 0x2f, 0x28, 0x9b, 0x1c,
	0x0f, 0x2c, 0x9b, 0x94, 0xbe, 0x8c, 0xff, 0xcd, 0xdd, 0xfe, 0x2e, 0x95, 0x1e, 0xca, 0x78, 0xb5,
	0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0xe3, 0x0e, 0x6e, 0x0b, 0x04, 0x00, 0x00,
}