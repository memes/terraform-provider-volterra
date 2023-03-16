// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/public_config_kubeconfig.proto

// Site KubeConfig custom API
//
// x-displayName: "Site"
// API for manage kube configs.

package site

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("ves.io/schema/site/public_config_kubeconfig.proto", fileDescriptor_d1ea8d623951f759)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/public_config_kubeconfig.proto", fileDescriptor_d1ea8d623951f759)
}

var fileDescriptor_d1ea8d623951f759 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6a, 0x14, 0x4d,
	0x14, 0xc5, 0xbb, 0x26, 0xf0, 0x2d, 0x66, 0xf5, 0xd1, 0x51, 0x98, 0x8c, 0x52, 0xca, 0x80, 0x04,
	0x24, 0x5d, 0x85, 0xba, 0x11, 0x51, 0x24, 0xc9, 0x26, 0x62, 0x40, 0x71, 0xe9, 0x26, 0x54, 0xf7,
	0xdc, 0xe9, 0x29, 0xa7, 0xbb, 0x6f, 0xd9, 0x55, 0xdd, 0x66, 0x90, 0x80, 0xcc, 0x13, 0x08, 0xbe,
	0x84, 0x8f, 0x20, 0x64, 0x93, 0x5d, 0xb2, 0x92, 0x41, 0x37, 0x59, 0x89, 0xd3, 0xe3, 0xc2, 0x65,
	0xf0, 0x09, 0x64, 0xaa, 0x7b, 0xfe, 0xe9, 0x80, 0x82, 0xbb, 0x5b, 0xf7, 0x9c, 0x7b, 0xba, 0x7f,
	0x55, 0xb7, 0x7e, 0x2b, 0x07, 0xcd, 0x24, 0x72, 0x1d, 0x74, 0x21, 0x16, 0x5c, 0x4b, 0x03, 0x5c,
	0x65, 0x7e, 0x24, 0x83, 0x83, 0x00, 0x93, 0x8e, 0x0c, 0x0f, 0x7a, 0x99, 0x0f, 0x65, 0xc9, 0x54,
	0x8a, 0x06, 0x5d, 0xb7, 0x1c, 0x61, 0xe5, 0x08, 0x9b, 0x8c, 0x34, 0xbd, 0x50, 0x9a, 0x6e, 0xe6,
	0xb3, 0x00, 0x63, 0x1e, 0x62, 0x88, 0xdc, 0x5a, 0xfd, 0xac, 0x63, 0x4f, 0xf6, 0x60, 0xab, 0x32,
	0xa2, 0x79, 0x35, 0x44, 0x0c, 0x23, 0xe0, 0x42, 0x49, 0x2e, 0x92, 0x04, 0x8d, 0x30, 0x12, 0x13,
	0x5d, 0xa9, 0x1b, 0x0b, 0x6a, 0xd7, 0x18, 0xe5, 0x63, 0xbb, 0x5f, 0x49, 0xd7, 0x2a, 0x69, 0x16,
	0x6f, 0x64, 0x0c, 0xda, 0x88, 0x58, 0x55, 0x86, 0x2b, 0xcb, 0x3c, 0xa8, 0x16, 0x83, 0xe9, 0x0a,
	0x58, 0xd3, 0x57, 0x30, 0xd5, 0x5b, 0xcb, 0x7a, 0x0e, 0x1a, 0x92, 0x7c, 0x39, 0xe3, 0xf6, 0x8f,
	0xb5, 0xfa, 0xfa, 0xae, 0xbd, 0x8e, 0xc7, 0x99, 0x0f, 0x65, 0xb5, 0xfd, 0xf4, 0x91, 0x3b, 0x24,
	0xf5, 0xcb, 0xbb, 0x29, 0x08, 0x03, 0xfb, 0x18, 0x88, 0x68, 0x2e, 0xba, 0x9b, 0xec, 0xf7, 0x0b,
	0x63, 0xa5, 0x75, 0xee, 0x7a, 0x06, 0x2f, 0x9b, 0x97, 0x58, 0x49, 0xc7, 0x84, 0x92, 0x6c, 0xcf,
	0x18, 0xb5, 0x83, 0xed, 0x7e, 0xeb, 0xb0, 0x38, 0x6d, 0x78, 0x39, 0x68, 0x4f, 0xa2, 0xd7, 0xbb,
	0xab, 0xbd, 0x57, 0xa9, 0x34, 0xb0, 0x75, 0xbd, 0xea, 0x44, 0x93, 0x8f, 0xd9, 0xfe, 0xfc, 0x9d,
	0x06, 0x9f, 0xbf, 0xbd, 0xab, 0x6d, 0xb7, 0xee, 0x57, 0x0f, 0xc9, 0x13, 0x11, 0x83, 0x56, 0x22,
	0x00, 0xcd, 0x5f, 0xcf, 0xea, 0x23, 0x8b, 0x5f, 0x75, 0x8e, 0x78, 0x95, 0x33, 0xcb, 0xb8, 0x47,
	0x6e, 0xba, 0x5f, 0x48, 0x7d, 0x7d, 0x5f, 0x6a, 0xf3, 0x2b, 0xd0, 0x8d, 0x55, 0x40, 0x13, 0xe3,
	0x32, 0xce, 0xdf, 0xd8, 0xb4, 0x6a, 0xf5, 0xcf, 0x3e, 0xd4, 0x48, 0x71, 0xda, 0xd8, 0x5a, 0x60,
	0x4c, 0x41, 0xb4, 0xff, 0x8c, 0xf8, 0xd0, 0x7d, 0xf0, 0x2f, 0x88, 0xba, 0xb9, 0x79, 0x72, 0x4c,
	0xd6, 0x3e, 0x1d, 0x93, 0x8d, 0x15, 0x3f, 0xfa, 0xc4, 0x7f, 0x01, 0x81, 0x19, 0x7c, 0x6c, 0xd4,
	0xfe, 0x27, 0x3b, 0x03, 0x32, 0x1c, 0x51, 0xe7, 0x7c, 0x44, 0x9d, 0x8b, 0x11, 0x25, 0x6f, 0x0a,
	0x4a, 0xde, 0x17, 0x94, 0x9c, 0x15, 0x94, 0x0c, 0x0b, 0x4a, 0xbe, 0x16, 0x94, 0x7c, 0x2f, 0xa8,
	0x73, 0x51, 0x50, 0xf2, 0x76, 0x4c, 0x9d, 0x93, 0x31, 0x25, 0xc3, 0x31, 0x75, 0xce, 0xc7, 0xd4,
	0x79, 0xbe, 0x17, 0xa2, 0xea, 0x85, 0x2c, 0xc7, 0xc8, 0x40, 0x9a, 0x0a, 0x96, 0x69, 0x6e, 0x8b,
	0x0e, 0xa6, 0xb1, 0xa7, 0x52, 0xcc, 0x65, 0x1b, 0x52, 0x6f, 0x2a, 0x73, 0xe5, 0x87, 0xc8, 0xe1,
	0xd0, 0x4c, 0x57, 0x74, 0xbe, 0xa9, 0xfe, 0x7f, 0x76, 0x01, 0xef, 0xfc, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xc0, 0x6c, 0x90, 0xfe, 0xb3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigKubeConfigAPIClient is the client API for ConfigKubeConfigAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigKubeConfigAPIClient interface {
	// Create K8s Cluster Local Kube Config
	//
	// x-displayName: "Create K8s Cluster Local Kube Config"
	// Down load kube config for local k8s cluster access
	CreateLocalKubeConfig(ctx context.Context, in *CreateKubeConfigReq, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// List Local Kube Configs
	//
	// x-displayName: "List Local Kube Configs"
	// Returns list of all local active kubeconfig minted for this site
	ListLocalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error)
}

type configKubeConfigAPIClient struct {
	cc *grpc.ClientConn
}

func NewConfigKubeConfigAPIClient(cc *grpc.ClientConn) ConfigKubeConfigAPIClient {
	return &configKubeConfigAPIClient{cc}
}

func (c *configKubeConfigAPIClient) CreateLocalKubeConfig(ctx context.Context, in *CreateKubeConfigReq, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/ves.io.schema.site.ConfigKubeConfigAPI/CreateLocalKubeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configKubeConfigAPIClient) ListLocalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error) {
	out := new(ListKubeConfigRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.site.ConfigKubeConfigAPI/ListLocalKubeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigKubeConfigAPIServer is the server API for ConfigKubeConfigAPI service.
type ConfigKubeConfigAPIServer interface {
	// Create K8s Cluster Local Kube Config
	//
	// x-displayName: "Create K8s Cluster Local Kube Config"
	// Down load kube config for local k8s cluster access
	CreateLocalKubeConfig(context.Context, *CreateKubeConfigReq) (*httpbody.HttpBody, error)
	// List Local Kube Configs
	//
	// x-displayName: "List Local Kube Configs"
	// Returns list of all local active kubeconfig minted for this site
	ListLocalKubeConfig(context.Context, *ListKubeConfigReq) (*ListKubeConfigRsp, error)
}

// UnimplementedConfigKubeConfigAPIServer can be embedded to have forward compatible implementations.
type UnimplementedConfigKubeConfigAPIServer struct {
}

func (*UnimplementedConfigKubeConfigAPIServer) CreateLocalKubeConfig(ctx context.Context, req *CreateKubeConfigReq) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocalKubeConfig not implemented")
}
func (*UnimplementedConfigKubeConfigAPIServer) ListLocalKubeConfig(ctx context.Context, req *ListKubeConfigReq) (*ListKubeConfigRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocalKubeConfig not implemented")
}

func RegisterConfigKubeConfigAPIServer(s *grpc.Server, srv ConfigKubeConfigAPIServer) {
	s.RegisterService(&_ConfigKubeConfigAPI_serviceDesc, srv)
}

func _ConfigKubeConfigAPI_CreateLocalKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKubeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigKubeConfigAPIServer).CreateLocalKubeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.ConfigKubeConfigAPI/CreateLocalKubeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigKubeConfigAPIServer).CreateLocalKubeConfig(ctx, req.(*CreateKubeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigKubeConfigAPI_ListLocalKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigKubeConfigAPIServer).ListLocalKubeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.ConfigKubeConfigAPI/ListLocalKubeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigKubeConfigAPIServer).ListLocalKubeConfig(ctx, req.(*ListKubeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigKubeConfigAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.ConfigKubeConfigAPI",
	HandlerType: (*ConfigKubeConfigAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocalKubeConfig",
			Handler:    _ConfigKubeConfigAPI_CreateLocalKubeConfig_Handler,
		},
		{
			MethodName: "ListLocalKubeConfig",
			Handler:    _ConfigKubeConfigAPI_ListLocalKubeConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/public_config_kubeconfig.proto",
}
