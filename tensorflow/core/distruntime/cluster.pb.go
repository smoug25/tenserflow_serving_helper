// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/cluster.proto

package distruntime

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

// Defines a single job in a TensorFlow cluster.
type JobDef struct {
	// The name of this job.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mapping from task ID to "hostname:port" string.
	//
	// If the `name` field contains "worker", and the `tasks` map contains a
	// mapping from 7 to "example.org:2222", then the device prefix
	// "/job:worker/task:7" will be assigned to "example.org:2222".
	Tasks                map[int32]string `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JobDef) Reset()         { *m = JobDef{} }
func (m *JobDef) String() string { return proto.CompactTextString(m) }
func (*JobDef) ProtoMessage()    {}
func (*JobDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ca946f62f52587, []int{0}
}

func (m *JobDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobDef.Unmarshal(m, b)
}
func (m *JobDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobDef.Marshal(b, m, deterministic)
}
func (m *JobDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobDef.Merge(m, src)
}
func (m *JobDef) XXX_Size() int {
	return xxx_messageInfo_JobDef.Size(m)
}
func (m *JobDef) XXX_DiscardUnknown() {
	xxx_messageInfo_JobDef.DiscardUnknown(m)
}

var xxx_messageInfo_JobDef proto.InternalMessageInfo

func (m *JobDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobDef) GetTasks() map[int32]string {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// Defines a TensorFlow cluster as a set of jobs.
type ClusterDef struct {
	// The jobs that comprise the cluster.
	Job                  []*JobDef `protobuf:"bytes,1,rep,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ClusterDef) Reset()         { *m = ClusterDef{} }
func (m *ClusterDef) String() string { return proto.CompactTextString(m) }
func (*ClusterDef) ProtoMessage()    {}
func (*ClusterDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ca946f62f52587, []int{1}
}

func (m *ClusterDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterDef.Unmarshal(m, b)
}
func (m *ClusterDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterDef.Marshal(b, m, deterministic)
}
func (m *ClusterDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDef.Merge(m, src)
}
func (m *ClusterDef) XXX_Size() int {
	return xxx_messageInfo_ClusterDef.Size(m)
}
func (m *ClusterDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDef.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDef proto.InternalMessageInfo

func (m *ClusterDef) GetJob() []*JobDef {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*JobDef)(nil), "tensorflow.JobDef")
	proto.RegisterMapType((map[int32]string)(nil), "tensorflow.JobDef.TasksEntry")
	proto.RegisterType((*ClusterDef)(nil), "tensorflow.ClusterDef")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/cluster.proto", fileDescriptor_e2ca946f62f52587)
}

var fileDescriptor_e2ca946f62f52587 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0xc6, 0x14, 0x1c, 0x11, 0x64, 0xf1, 0x10, 0x0a, 0x42, 0x29, 0x1e, 0x7a, 0x4a,
	0x20, 0x45, 0x28, 0x1e, 0xab, 0x82, 0x78, 0x2a, 0xd1, 0x93, 0x97, 0x92, 0x8d, 0x93, 0x34, 0x36,
	0xd9, 0x29, 0xfb, 0xa7, 0x52, 0xbf, 0x82, 0x5f, 0xd8, 0xa3, 0xec, 0x46, 0x48, 0xc0, 0xa3, 0xb7,
	0xd9, 0xf9, 0xcd, 0xbc, 0x37, 0xfb, 0xe0, 0xb9, 0xaa, 0xcd, 0xd6, 0x8a, 0xb8, 0xa0, 0x36, 0xd1,
	0x2d, 0xd9, 0x2a, 0xbd, 0x49, 0x0c, 0x4a, 0x4d, 0xaa, 0x6c, 0xe8, 0x63, 0xa3, 0x51, 0x1d, 0x6a,
	0x59, 0x6d, 0xb6, 0xd8, 0xec, 0x51, 0x0d, 0x48, 0x52, 0x90, 0xc2, 0x64, 0xaf, 0xc8, 0x90, 0xb0,
	0x65, 0x52, 0x34, 0x56, 0x1b, 0x54, 0xb1, 0x6f, 0x70, 0xe8, 0xe7, 0x66, 0x5f, 0x0c, 0xc6, 0x4f,
	0x24, 0xee, 0xb1, 0xe4, 0x1c, 0x4e, 0x64, 0xde, 0x62, 0xc4, 0xa6, 0x6c, 0x7e, 0x9a, 0xf9, 0x9a,
	0x2f, 0x20, 0x34, 0xb9, 0xde, 0xe9, 0x68, 0x34, 0x0d, 0xe6, 0x67, 0xe9, 0x55, 0xdc, 0xaf, 0xc6,
	0xdd, 0x5a, 0xfc, 0xe2, 0xf8, 0x83, 0x34, 0xea, 0x98, 0x75, 0xb3, 0x93, 0x25, 0x40, 0xdf, 0xe4,
	0x17, 0x10, 0xec, 0xf0, 0xe8, 0x55, 0xc3, 0xcc, 0x95, 0xfc, 0x12, 0xc2, 0x43, 0xde, 0x58, 0x8c,
	0x46, 0xde, 0xa9, 0x7b, 0xdc, 0x8e, 0x96, 0x6c, 0x96, 0x02, 0xdc, 0x75, 0xa7, 0xba, 0x83, 0xae,
	0x21, 0x78, 0x27, 0x11, 0x31, 0x6f, 0xcd, 0xff, 0x5a, 0x67, 0x0e, 0xaf, 0x3e, 0x61, 0x42, 0xaa,
	0x1a, 0xd2, 0xb7, 0x5a, 0x1b, 0x65, 0xa5, 0xa9, 0x5b, 0x5c, 0x9d, 0xff, 0xea, 0xad, 0xdd, 0xcf,
	0xf5, 0x9a, 0xbd, 0x3e, 0xfe, 0x2b, 0xd1, 0x81, 0xf4, 0x37, 0x63, 0x62, 0xec, 0x03, 0x5d, 0xfc,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xad, 0xc6, 0x77, 0xfa, 0xa7, 0x01, 0x00, 0x00,
}
