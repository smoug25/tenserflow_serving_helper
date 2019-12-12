// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/prediction_service.proto

package apis

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/prediction_service.proto", fileDescriptor_3d3f0e8f893ea594)
}

var fileDescriptor_3d3f0e8f893ea594 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x19, 0x82, 0x93, 0x1c, 0x44, 0x73, 0xec, 0x51, 0x51, 0xf0, 0x0f, 0x2d, 0x4c, 0xfc,
	0x02, 0x4e, 0x90, 0x1d, 0x06, 0x52, 0x15, 0x54, 0x90, 0x92, 0xb5, 0x6f, 0xb3, 0x40, 0x9b, 0xd4,
	0xbc, 0xa9, 0xe2, 0x37, 0xf7, 0xe0, 0x41, 0x66, 0xd2, 0xca, 0x66, 0x98, 0x3b, 0xe4, 0x9a, 0x3c,
	0xcf, 0xef, 0x49, 0xde, 0xe7, 0x25, 0x2f, 0x5c, 0x98, 0x79, 0x3b, 0x8b, 0x73, 0x55, 0x27, 0x58,
	0xab, 0x96, 0x8f, 0x2e, 0x13, 0x03, 0x12, 0x95, 0x2e, 0x2b, 0xf5, 0x9e, 0x21, 0xe8, 0x37, 0x21,
	0x79, 0x36, 0x87, 0xaa, 0x01, 0xed, 0xb9, 0x49, 0x58, 0x23, 0x30, 0x69, 0x34, 0x14, 0x22, 0x37,
	0x42, 0x49, 0x7b, 0x9e, 0x43, 0xdc, 0x68, 0x65, 0x14, 0xa5, 0xbf, 0x8e, 0xd8, 0x39, 0xa2, 0xa7,
	0x20, 0x91, 0x79, 0xc5, 0x10, 0x45, 0x29, 0x72, 0xb6, 0x88, 0xb5, 0x71, 0x51, 0x98, 0xdf, 0x70,
	0x30, 0x59, 0xad, 0x0a, 0xa8, 0xb2, 0x1a, 0x0c, 0x2b, 0x98, 0x61, 0x0e, 0x7f, 0x1f, 0x04, 0x2f,
	0x64, 0x09, 0x1a, 0x64, 0x37, 0xa3, 0x28, 0x0d, 0x59, 0x81, 0x63, 0x3e, 0x04, 0x61, 0x6a, 0xe0,
	0x1a, 0x10, 0xfb, 0xf9, 0x8e, 0xbe, 0xb6, 0xc8, 0xfe, 0x6d, 0xdf, 0xf5, 0x9d, 0xad, 0x9a, 0x32,
	0xb2, 0x33, 0xb6, 0x6d, 0x7c, 0xd0, 0x93, 0xf8, 0x6f, 0xe3, 0xf1, 0x78, 0xa9, 0xab, 0x14, 0x5e,
	0x5b, 0x40, 0x13, 0x9d, 0x6e, 0x22, 0xc5, 0x46, 0x49, 0x04, 0xfa, 0x48, 0x86, 0xa9, 0x7d, 0x0c,
	0x3d, 0xf2, 0xd9, 0xd2, 0xfe, 0xa5, 0x1d, 0xfd, 0xf8, 0x3f, 0x99, 0x23, 0xa7, 0x64, 0xe8, 0x7e,
	0x44, 0x0f, 0x7c, 0x16, 0x77, 0xd9, 0x61, 0x0f, 0xd7, 0x6a, 0x1c, 0x93, 0x93, 0xdd, 0x69, 0x5b,
	0x19, 0x31, 0xe9, 0x9a, 0xf6, 0x8f, 0x65, 0x59, 0xb3, 0x76, 0x2c, 0xab, 0x52, 0x17, 0x54, 0x93,
	0xbd, 0x1b, 0x30, 0xd3, 0xc5, 0xae, 0x4e, 0xdd, 0xaa, 0xd2, 0x33, 0x9f, 0x7f, 0x55, 0xd5, 0x85,
	0x9d, 0x6f, 0x26, 0xb6, 0x71, 0x57, 0x93, 0xe7, 0xeb, 0x10, 0x7b, 0xf5, 0x39, 0x18, 0xcc, 0xb6,
	0x7f, 0x16, 0xea, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x59, 0xd6, 0x4b, 0xe9, 0x80, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Classify.
	Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error)
	// Regress.
	Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error) {
	out := new(ClassificationResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Classify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error) {
	out := new(RegressionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Regress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error) {
	out := new(MultiInferenceResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/MultiInference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error) {
	out := new(GetModelMetadataResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/GetModelMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Classify.
	Classify(context.Context, *ClassificationRequest) (*ClassificationResponse, error)
	// Regress.
	Regress(context.Context, *RegressionRequest) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(context.Context, *MultiInferenceRequest) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(context.Context, *GetModelMetadataRequest) (*GetModelMetadataResponse, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Classify(ctx context.Context, req *ClassificationRequest) (*ClassificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Classify not implemented")
}
func (*UnimplementedPredictionServiceServer) Regress(ctx context.Context, req *RegressionRequest) (*RegressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Regress not implemented")
}
func (*UnimplementedPredictionServiceServer) Predict(ctx context.Context, req *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (*UnimplementedPredictionServiceServer) MultiInference(ctx context.Context, req *MultiInferenceRequest) (*MultiInferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiInference not implemented")
}
func (*UnimplementedPredictionServiceServer) GetModelMetadata(ctx context.Context, req *GetModelMetadataRequest) (*GetModelMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelMetadata not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Classify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Classify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Classify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Classify(ctx, req.(*ClassificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Regress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Regress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Regress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Regress(ctx, req.(*RegressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_MultiInference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).MultiInference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/MultiInference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).MultiInference(ctx, req.(*MultiInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetModelMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/GetModelMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, req.(*GetModelMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classify",
			Handler:    _PredictionService_Classify_Handler,
		},
		{
			MethodName: "Regress",
			Handler:    _PredictionService_Regress_Handler,
		},
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "MultiInference",
			Handler:    _PredictionService_MultiInference_Handler,
		},
		{
			MethodName: "GetModelMetadata",
			Handler:    _PredictionService_GetModelMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/prediction_service.proto",
}