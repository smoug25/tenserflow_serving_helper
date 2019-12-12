// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/remote_tensor_handle.proto

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

type ResourceDtypeAndShape struct {
	Dtype                DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResourceDtypeAndShape) Reset()         { *m = ResourceDtypeAndShape{} }
func (m *ResourceDtypeAndShape) String() string { return proto.CompactTextString(m) }
func (*ResourceDtypeAndShape) ProtoMessage()    {}
func (*ResourceDtypeAndShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_617410f375f42e29, []int{0}
}

func (m *ResourceDtypeAndShape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceDtypeAndShape.Unmarshal(m, b)
}
func (m *ResourceDtypeAndShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceDtypeAndShape.Marshal(b, m, deterministic)
}
func (m *ResourceDtypeAndShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDtypeAndShape.Merge(m, src)
}
func (m *ResourceDtypeAndShape) XXX_Size() int {
	return xxx_messageInfo_ResourceDtypeAndShape.Size(m)
}
func (m *ResourceDtypeAndShape) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDtypeAndShape.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDtypeAndShape proto.InternalMessageInfo

func (m *ResourceDtypeAndShape) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *ResourceDtypeAndShape) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

type RemoteTensorHandle struct {
	// The ID of the operation that produced this tensor.
	OpId int64 `protobuf:"varint,1,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	// The index into the outputs of the operation that produced this tensor.
	OutputNum int32 `protobuf:"varint,2,opt,name=output_num,json=outputNum,proto3" json:"output_num,omitempty"`
	// Device of the operation that produced this tensor. Cannot be empty.
	// For multi-device functions, it's the default device passed to placer.
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	// Device where the tensor is located. Can be empty if the operation producing
	// this tensor is a multi-device function.
	OpDevice string `protobuf:"bytes,4,opt,name=op_device,json=opDevice,proto3" json:"op_device,omitempty"`
	// Tensor type.
	Dtype DataType `protobuf:"varint,5,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Optional data types and shapes of a remote resource variable.
	ResourceDtypesAndShapes []*ResourceDtypeAndShape `protobuf:"bytes,6,rep,name=resource_dtypes_and_shapes,json=resourceDtypesAndShapes,proto3" json:"resource_dtypes_and_shapes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *RemoteTensorHandle) Reset()         { *m = RemoteTensorHandle{} }
func (m *RemoteTensorHandle) String() string { return proto.CompactTextString(m) }
func (*RemoteTensorHandle) ProtoMessage()    {}
func (*RemoteTensorHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_617410f375f42e29, []int{1}
}

func (m *RemoteTensorHandle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteTensorHandle.Unmarshal(m, b)
}
func (m *RemoteTensorHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteTensorHandle.Marshal(b, m, deterministic)
}
func (m *RemoteTensorHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteTensorHandle.Merge(m, src)
}
func (m *RemoteTensorHandle) XXX_Size() int {
	return xxx_messageInfo_RemoteTensorHandle.Size(m)
}
func (m *RemoteTensorHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteTensorHandle.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteTensorHandle proto.InternalMessageInfo

func (m *RemoteTensorHandle) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *RemoteTensorHandle) GetOutputNum() int32 {
	if m != nil {
		return m.OutputNum
	}
	return 0
}

func (m *RemoteTensorHandle) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *RemoteTensorHandle) GetOpDevice() string {
	if m != nil {
		return m.OpDevice
	}
	return ""
}

func (m *RemoteTensorHandle) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *RemoteTensorHandle) GetResourceDtypesAndShapes() []*ResourceDtypeAndShape {
	if m != nil {
		return m.ResourceDtypesAndShapes
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceDtypeAndShape)(nil), "tensorflow.eager.ResourceDtypeAndShape")
	proto.RegisterType((*RemoteTensorHandle)(nil), "tensorflow.eager.RemoteTensorHandle")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/remote_tensor_handle.proto", fileDescriptor_617410f375f42e29)
}

var fileDescriptor_617410f375f42e29 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4f, 0xeb, 0xd3, 0x40,
	0x14, 0x64, 0x7f, 0xfd, 0xa5, 0xd8, 0x2d, 0x88, 0xac, 0xff, 0x42, 0x55, 0x08, 0xbd, 0x18, 0x3c,
	0x24, 0x10, 0xf1, 0x03, 0x58, 0x8a, 0xe8, 0x45, 0xca, 0xda, 0x8b, 0x5e, 0xd6, 0x34, 0xfb, 0x9a,
	0x04, 0x9b, 0xbc, 0x65, 0x37, 0xdb, 0xd2, 0xa3, 0x9f, 0xc0, 0xaf, 0xeb, 0x51, 0xb2, 0x1b, 0x6b,
	0x50, 0xc1, 0x83, 0xbd, 0x25, 0x33, 0x93, 0x99, 0xbc, 0x79, 0x8f, 0x7e, 0x2e, 0xeb, 0xae, 0xb2,
	0xbb, 0xa4, 0xc0, 0x26, 0x35, 0x0d, 0xda, 0x32, 0x7b, 0x95, 0x76, 0xd0, 0x1a, 0xd4, 0xfb, 0x03,
	0x9e, 0x84, 0x01, 0x7d, 0xac, 0xdb, 0x52, 0x54, 0x70, 0x50, 0xa0, 0x47, 0x4c, 0x5a, 0xa0, 0x86,
	0x54, 0x69, 0xec, 0x70, 0x67, 0xf7, 0xa9, 0x86, 0x06, 0x3b, 0x10, 0x9e, 0x17, 0x55, 0xde, 0xca,
	0x03, 0x24, 0x8e, 0x65, 0xf7, 0x7e, 0x7d, 0x94, 0x40, 0x5e, 0x82, 0x5e, 0x7c, 0xfc, 0xaf, 0xcc,
	0xbd, 0xce, 0x1b, 0x38, 0xa1, 0xfe, 0x32, 0x30, 0xc2, 0x54, 0xb9, 0x1a, 0xc2, 0x16, 0xfc, 0x5a,
	0xd6, 0x67, 0x05, 0xc6, 0x7b, 0x2e, 0x4f, 0xf4, 0x21, 0x07, 0x83, 0x56, 0x17, 0xb0, 0xee, 0xf1,
	0xd7, 0xad, 0xfc, 0xd0, 0x47, 0xb2, 0x17, 0x34, 0x90, 0x3d, 0x10, 0x92, 0x88, 0xc4, 0x77, 0xb3,
	0x07, 0xc9, 0x68, 0xd2, 0x75, 0xde, 0xe5, 0xdb, 0xb3, 0x02, 0xee, 0x25, 0x2c, 0xa3, 0x81, 0xfb,
	0xcf, 0xf0, 0x26, 0x22, 0xf1, 0x3c, 0x7b, 0x3a, 0xd6, 0x6e, 0xdd, 0xa3, 0xf3, 0xdc, 0xf4, 0x89,
	0xdc, 0x4b, 0x97, 0xdf, 0x6e, 0x28, 0xe3, 0xae, 0x58, 0xaf, 0x78, 0xeb, 0x6a, 0x65, 0xf7, 0x69,
	0x80, 0x4a, 0xd4, 0xd2, 0xc5, 0x4e, 0xf8, 0x2d, 0xaa, 0x77, 0x92, 0x3d, 0xa3, 0x14, 0x6d, 0xa7,
	0x6c, 0x27, 0x5a, 0xdb, 0xb8, 0x90, 0x80, 0xcf, 0x3c, 0xf2, 0xde, 0x36, 0xec, 0x11, 0x9d, 0x4a,
	0x38, 0xd6, 0x05, 0x84, 0x93, 0x88, 0xc4, 0x33, 0x3e, 0xbc, 0xb1, 0x27, 0x74, 0x86, 0x4a, 0x0c,
	0xd4, 0xad, 0xa3, 0xee, 0xa0, 0x5a, 0x7b, 0xf2, 0x32, 0x5f, 0xf0, 0xef, 0xf9, 0x24, 0x5d, 0xe8,
	0xa1, 0x24, 0xe1, 0x10, 0x23, 0xf2, 0x56, 0xfa, 0xdd, 0x98, 0x70, 0x1a, 0x4d, 0xe2, 0x79, 0xf6,
	0x3c, 0xf9, 0xfd, 0x14, 0x92, 0xbf, 0x16, 0xcb, 0x1f, 0xeb, 0x31, 0x6c, 0x7e, 0xe2, 0x66, 0xf5,
	0x95, 0xd0, 0x10, 0x75, 0x39, 0xf6, 0xb9, 0xec, 0x6c, 0x15, 0xfe, 0xd9, 0x95, 0xab, 0xd3, 0x6c,
	0xc8, 0xa7, 0x37, 0xd7, 0xb9, 0x8b, 0xef, 0x84, 0xec, 0xa6, 0xee, 0x2a, 0x5e, 0xfe, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0x83, 0x31, 0xe8, 0x3a, 0x03, 0x00, 0x00,
}
