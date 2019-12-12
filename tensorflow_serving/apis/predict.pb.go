// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/predict.proto

package apis

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework"
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

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'inputs' field.
	Inputs map[string]*framework.TensorProto `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'outputs' field.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter         []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter,proto3" json:"output_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feae9fafe7dbff83, []int{0}
}

func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictRequest) GetInputs() map[string]*framework.TensorProto {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PredictRequest) GetOutputFilter() []string {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	// Effective Model Specification used to process PredictRequest.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Output tensors.
	Outputs              map[string]*framework.TensorProto `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_feae9fafe7dbff83, []int{1}
}

func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (m *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(m, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictResponse) GetOutputs() map[string]*framework.TensorProto {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "tensorflow.serving.PredictRequest")
	proto.RegisterMapType((map[string]*framework.TensorProto)(nil), "tensorflow.serving.PredictRequest.InputsEntry")
	proto.RegisterType((*PredictResponse)(nil), "tensorflow.serving.PredictResponse")
	proto.RegisterMapType((map[string]*framework.TensorProto)(nil), "tensorflow.serving.PredictResponse.OutputsEntry")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/predict.proto", fileDescriptor_feae9fafe7dbff83)
}

var fileDescriptor_feae9fafe7dbff83 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4b, 0x3a, 0x41,
	0x18, 0x66, 0x76, 0xf9, 0xf9, 0xc3, 0xd1, 0x3e, 0x98, 0x4b, 0xcb, 0x42, 0x20, 0x76, 0xf1, 0xd2,
	0x6c, 0x18, 0x41, 0x44, 0xa7, 0x28, 0xc1, 0x20, 0x92, 0xb1, 0x53, 0x97, 0x45, 0xd7, 0x57, 0x5d,
	0xdc, 0xdd, 0x99, 0xe6, 0x43, 0xf1, 0xaf, 0xe8, 0x5f, 0xed, 0xd8, 0x31, 0x9c, 0xd1, 0xb2, 0x92,
	0xa0, 0xf0, 0xf6, 0x32, 0xef, 0xf3, 0xcc, 0xf3, 0xc1, 0x8b, 0xd9, 0x28, 0xd5, 0x63, 0xd3, 0xa7,
	0x09, 0xcf, 0x23, 0x95, 0x73, 0x33, 0x6a, 0x9e, 0x45, 0x1a, 0x0a, 0xc5, 0xe5, 0x30, 0xe3, 0xb3,
	0x58, 0x81, 0x9c, 0xa6, 0xc5, 0x28, 0x1e, 0x43, 0x26, 0x40, 0x6e, 0xd8, 0x44, 0x3d, 0x91, 0xaa,
	0x48, 0x48, 0x18, 0xa4, 0x89, 0xa6, 0x42, 0x72, 0xcd, 0x09, 0xf9, 0x80, 0xd1, 0x25, 0x2c, 0xec,
	0xfe, 0x51, 0x27, 0x4a, 0xb8, 0x84, 0x68, 0x28, 0x7b, 0x39, 0xcc, 0xb8, 0x9c, 0x2c, 0x37, 0x4e,
	0x28, 0xec, 0x6c, 0xc5, 0x7c, 0xce, 0x07, 0x90, 0xb9, 0x1f, 0xeb, 0xcf, 0x1e, 0xde, 0xed, 0xb8,
	0x30, 0x0c, 0x9e, 0x0c, 0x28, 0x4d, 0x2e, 0x31, 0xb6, 0x88, 0x58, 0x09, 0x48, 0x02, 0x54, 0x43,
	0x8d, 0x4a, 0xf3, 0x90, 0x7e, 0x8f, 0x48, 0xef, 0x16, 0xa8, 0xae, 0x80, 0x84, 0x95, 0xf3, 0xd5,
	0x48, 0x5a, 0xb8, 0x94, 0x16, 0xc2, 0x68, 0x15, 0x78, 0x35, 0xbf, 0x51, 0x69, 0xd2, 0x4d, 0xcc,
	0xcf, 0x8a, 0xb4, 0x6d, 0x09, 0x37, 0x85, 0x96, 0x73, 0xb6, 0x64, 0x93, 0x23, 0xbc, 0xc3, 0x8d,
	0x16, 0x46, 0xc7, 0xc3, 0x34, 0xd3, 0x20, 0x03, 0xbf, 0xe6, 0x37, 0xca, 0xac, 0xea, 0x1e, 0x5b,
	0xf6, 0x2d, 0x64, 0xb8, 0xb2, 0xc6, 0x25, 0xfb, 0xd8, 0x9f, 0xc0, 0xdc, 0x5a, 0x2e, 0xb3, 0xc5,
	0x48, 0x8e, 0xf1, 0xbf, 0x69, 0x2f, 0x33, 0x10, 0x78, 0x36, 0xc6, 0xc1, 0xba, 0x99, 0x07, 0x3b,
	0x76, 0x16, 0x35, 0x30, 0x87, 0xba, 0xf0, 0xce, 0x51, 0xfd, 0x05, 0xe1, 0xbd, 0x77, 0x7f, 0x4a,
	0xf0, 0x42, 0xc1, 0x97, 0x4a, 0xbc, 0x5f, 0x56, 0x72, 0x8b, 0xff, 0x3b, 0xd7, 0x2a, 0x40, 0xb6,
	0x93, 0x93, 0x1f, 0x3b, 0x71, 0x9a, 0xf4, 0xde, 0x51, 0x5c, 0x2b, 0xab, 0x0f, 0xc2, 0x2e, 0xae,
	0xae, 0x2f, 0xb6, 0x12, 0xf9, 0xaa, 0xfd, 0x78, 0xbd, 0x8d, 0xc3, 0x7a, 0x45, 0xa8, 0x5f, 0xb2,
	0x67, 0x75, 0xfa, 0x16, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xfe, 0x6f, 0x76, 0x67, 0x03, 0x00, 0x00,
}
