// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/cdn_loadbalancer/private_customapi.proto

// CDN
//
// x-displayName: "CDN Loadbalancer"
// CDN package

package cdn_loadbalancer

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	virtual_host "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// UpdateServiceDomainsRequest
//
// x-displayName: "UpdateServiceDomainsRequest"
// UpdateServiceDomainsRequest
type UpdateServiceDomainsRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "default"
	// x-required
	// Namespace scope of the request
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "default"
	// x-required
	// Name of the CDN loadbalancer
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Tenant
	//
	// x-displayName: "Tenant"
	// x-example: "default"
	// x-required
	// Tenant name
	Tenant string `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// Service Domains
	//
	// x-displayName: "Service Domains"
	// CNAME provided from service per domain
	ServiceDomains []*virtual_host.ServiceDomain `protobuf:"bytes,4,rep,name=service_domains,json=serviceDomains,proto3" json:"service_domains,omitempty"`
}

func (m *UpdateServiceDomainsRequest) Reset()      { *m = UpdateServiceDomainsRequest{} }
func (*UpdateServiceDomainsRequest) ProtoMessage() {}
func (*UpdateServiceDomainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a7635dbb4fc69a, []int{0}
}
func (m *UpdateServiceDomainsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateServiceDomainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateServiceDomainsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateServiceDomainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceDomainsRequest.Merge(m, src)
}
func (m *UpdateServiceDomainsRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateServiceDomainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceDomainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceDomainsRequest proto.InternalMessageInfo

func (m *UpdateServiceDomainsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateServiceDomainsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateServiceDomainsRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *UpdateServiceDomainsRequest) GetServiceDomains() []*virtual_host.ServiceDomain {
	if m != nil {
		return m.ServiceDomains
	}
	return nil
}

// UpdateServiceDomainsResponse
//
// x-displayName: "UpdateServiceDomainsResponse"
// UpdateServiceDomainsResponse
type UpdateServiceDomainsResponse struct {
}

func (m *UpdateServiceDomainsResponse) Reset()      { *m = UpdateServiceDomainsResponse{} }
func (*UpdateServiceDomainsResponse) ProtoMessage() {}
func (*UpdateServiceDomainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a7635dbb4fc69a, []int{1}
}
func (m *UpdateServiceDomainsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateServiceDomainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateServiceDomainsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateServiceDomainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceDomainsResponse.Merge(m, src)
}
func (m *UpdateServiceDomainsResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateServiceDomainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceDomainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceDomainsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateServiceDomainsRequest)(nil), "ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsRequest")
	golang_proto.RegisterType((*UpdateServiceDomainsRequest)(nil), "ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsRequest")
	proto.RegisterType((*UpdateServiceDomainsResponse)(nil), "ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsResponse")
	golang_proto.RegisterType((*UpdateServiceDomainsResponse)(nil), "ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/views/cdn_loadbalancer/private_customapi.proto", fileDescriptor_b6a7635dbb4fc69a)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/cdn_loadbalancer/private_customapi.proto", fileDescriptor_b6a7635dbb4fc69a)
}

var fileDescriptor_b6a7635dbb4fc69a = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6a, 0x14, 0x4d,
	0x14, 0x9d, 0xca, 0x84, 0x40, 0xfa, 0x83, 0x4f, 0x69, 0x44, 0xc6, 0xc9, 0x50, 0x0c, 0x83, 0x48,
	0xfc, 0xe9, 0x2e, 0x89, 0xb8, 0x09, 0x6e, 0x12, 0xdd, 0xb8, 0x32, 0x8c, 0x64, 0x93, 0xcd, 0x50,
	0xdd, 0x7d, 0xd3, 0x29, 0xed, 0xae, 0xaa, 0x54, 0x55, 0x77, 0x12, 0x24, 0x10, 0xf2, 0x04, 0x82,
	0x3e, 0x80, 0x4b, 0xdf, 0x21, 0x9b, 0xec, 0x74, 0xa3, 0x04, 0xdd, 0x64, 0x69, 0x7a, 0x5c, 0xb8,
	0x8c, 0x6f, 0x20, 0xa9, 0x9e, 0xfc, 0x74, 0x27, 0x84, 0x80, 0xee, 0xee, 0xad, 0x53, 0xa7, 0xee,
	0xa9, 0x73, 0xef, 0x75, 0x9e, 0xe4, 0xa0, 0x7d, 0x26, 0x88, 0x0e, 0x57, 0x20, 0xa5, 0x24, 0x67,
	0xb0, 0xa6, 0x49, 0x18, 0xf1, 0x41, 0x22, 0x68, 0x14, 0xd0, 0x84, 0xf2, 0x10, 0x14, 0x91, 0x8a,
	0xe5, 0xd4, 0xc0, 0x20, 0xcc, 0xb4, 0x11, 0x29, 0x95, 0xcc, 0x97, 0x4a, 0x18, 0xe1, 0xde, 0x2e,
	0xd9, 0x7e, 0xc9, 0xf6, 0x2d, 0xdb, 0xaf, 0xb3, 0xdb, 0x5e, 0xcc, 0xcc, 0x4a, 0x16, 0xf8, 0xa1,
	0x48, 0x49, 0x2c, 0x62, 0x41, 0x2c, 0x39, 0xc8, 0x96, 0x6d, 0x66, 0x13, 0x1b, 0x95, 0x8f, 0xb6,
	0x3b, 0xb1, 0x10, 0x71, 0x02, 0x84, 0x4a, 0x46, 0x28, 0xe7, 0xc2, 0x50, 0xc3, 0x04, 0xd7, 0x23,
	0x74, 0xaa, 0x2a, 0x58, 0xc8, 0xb3, 0xe0, 0xad, 0x2a, 0x68, 0x36, 0x24, 0x1c, 0x43, 0x9d, 0xda,
	0x47, 0x69, 0xc2, 0x22, 0x6a, 0x60, 0x84, 0xf6, 0x6a, 0x28, 0x68, 0xe0, 0x79, 0xed, 0xf1, 0xee,
	0x79, 0xab, 0x06, 0xd5, 0x1b, 0x0f, 0xaf, 0x64, 0xe6, 0x59, 0x55, 0x77, 0xea, 0x0c, 0x65, 0x32,
	0x9a, 0x0c, 0x56, 0x84, 0x36, 0x67, 0xef, 0xf5, 0x76, 0x90, 0x33, 0xb5, 0x28, 0x8f, 0x04, 0xbf,
	0x04, 0x95, 0xb3, 0x10, 0x9e, 0x89, 0x94, 0x32, 0xae, 0xfb, 0xb0, 0x9a, 0x81, 0x36, 0x6e, 0xc7,
	0x99, 0xe4, 0x34, 0x05, 0x2d, 0x69, 0x08, 0x2d, 0xd4, 0x45, 0xd3, 0x93, 0xfd, 0xd3, 0x03, 0xd7,
	0x75, 0xc6, 0x8f, 0x92, 0xd6, 0x98, 0x05, 0x6c, 0xec, 0xde, 0x74, 0x26, 0x0c, 0x70, 0xca, 0x4d,
	0xab, 0x69, 0x4f, 0x47, 0x99, 0xdb, 0x77, 0xae, 0xe9, 0xb2, 0xc4, 0x20, 0x2a, 0x6b, 0xb4, 0xc6,
	0xbb, 0xcd, 0xe9, 0xff, 0x66, 0xee, 0xfa, 0xf5, 0x66, 0x9f, 0x6a, 0xf5, 0x2b, 0xaa, 0xfa, 0xff,
	0xeb, 0x8a, 0xc8, 0x1e, 0x76, 0x3a, 0x17, 0x8b, 0xd7, 0x52, 0x70, 0x0d, 0x33, 0xbf, 0x9b, 0xce,
	0xf5, 0x85, 0x72, 0xc4, 0x9e, 0xda, 0x09, 0x9b, 0x5b, 0x78, 0xee, 0x7e, 0x68, 0x3a, 0x37, 0x2e,
	0x62, 0xb9, 0x73, 0xfe, 0x55, 0xa6, 0xce, 0xbf, 0xc4, 0xae, 0xf6, 0xfc, 0xdf, 0x3c, 0x51, 0x8a,
	0xee, 0x6d, 0x8d, 0x15, 0x9f, 0x5a, 0xd3, 0x39, 0x68, 0x8f, 0x09, 0x4f, 0x2a, 0xb1, 0xbe, 0xe1,
	0xad, 0x29, 0x66, 0xe0, 0x41, 0x77, 0xf9, 0xf1, 0x7a, 0xe8, 0x85, 0x11, 0xf7, 0x02, 0xaa, 0x59,
	0xe8, 0xd1, 0x28, 0x65, 0x7c, 0xfb, 0xfb, 0xcf, 0x77, 0x63, 0x5f, 0x50, 0x6f, 0xf1, 0x78, 0x97,
	0x48, 0xb9, 0x4b, 0xe4, 0xa4, 0x4b, 0x9a, 0xbc, 0x39, 0x89, 0x37, 0xcf, 0x4f, 0x8d, 0x05, 0x37,
	0x49, 0x66, 0x25, 0x0d, 0x6a, 0x2d, 0x9a, 0x45, 0xf7, 0x96, 0x56, 0x7b, 0x09, 0xa9, 0xfc, 0x8a,
	0x30, 0x6e, 0x94, 0xd0, 0x12, 0x42, 0x43, 0xac, 0xb2, 0x7f, 0x5d, 0xb2, 0x3d, 0xbb, 0xbb, 0x83,
	0xc6, 0xbf, 0xed, 0xa0, 0xfb, 0x57, 0x72, 0xf3, 0x45, 0xf0, 0x0a, 0x42, 0xb3, 0xfd, 0xb5, 0xd5,
	0xdc, 0x45, 0x68, 0xfe, 0x3d, 0xda, 0x3b, 0xc0, 0x8d, 0xfd, 0x03, 0xdc, 0x38, 0x3c, 0xc0, 0x68,
	0xab, 0xc0, 0xe8, 0x63, 0x81, 0xd1, 0xe7, 0x02, 0xa3, 0xbd, 0x02, 0xa3, 0x1f, 0x05, 0x46, 0xbf,
	0x0a, 0xdc, 0x38, 0x2c, 0x30, 0x7a, 0x3b, 0xc4, 0x8d, 0xdd, 0x21, 0x46, 0x7b, 0x43, 0xdc, 0xd8,
	0x1f, 0xe2, 0xc6, 0xd2, 0x52, 0x2c, 0xe4, 0xeb, 0xd8, 0xcf, 0x45, 0x62, 0x40, 0x29, 0xea, 0x67,
	0x9a, 0xd8, 0x60, 0x59, 0xa8, 0xf4, 0xa8, 0x13, 0x39, 0x8b, 0x40, 0x79, 0xc7, 0x30, 0x91, 0x41,
	0x2c, 0x08, 0xac, 0x9b, 0x91, 0x2b, 0x97, 0x2e, 0x67, 0x30, 0x61, 0xf7, 0xed, 0xd1, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x25, 0x40, 0x70, 0x18, 0x05, 0x00, 0x00,
}

func (this *UpdateServiceDomainsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateServiceDomainsRequest)
	if !ok {
		that2, ok := that.(UpdateServiceDomainsRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Tenant != that1.Tenant {
		return false
	}
	if len(this.ServiceDomains) != len(that1.ServiceDomains) {
		return false
	}
	for i := range this.ServiceDomains {
		if !this.ServiceDomains[i].Equal(that1.ServiceDomains[i]) {
			return false
		}
	}
	return true
}
func (this *UpdateServiceDomainsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateServiceDomainsResponse)
	if !ok {
		that2, ok := that.(UpdateServiceDomainsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *UpdateServiceDomainsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&cdn_loadbalancer.UpdateServiceDomainsRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Tenant: "+fmt.Sprintf("%#v", this.Tenant)+",\n")
	if this.ServiceDomains != nil {
		s = append(s, "ServiceDomains: "+fmt.Sprintf("%#v", this.ServiceDomains)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpdateServiceDomainsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&cdn_loadbalancer.UpdateServiceDomainsResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPrivateCustomapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrivateCustomAPIClient is the client API for PrivateCustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivateCustomAPIClient interface {
	// UpdateServiceDomains
	//
	// x-displayName: "UpdateServiceDomains"
	// Private API to update service domains from service.
	UpdateServiceDomains(ctx context.Context, in *UpdateServiceDomainsRequest, opts ...grpc.CallOption) (*UpdateServiceDomainsResponse, error)
}

type privateCustomAPIClient struct {
	cc *grpc.ClientConn
}

func NewPrivateCustomAPIClient(cc *grpc.ClientConn) PrivateCustomAPIClient {
	return &privateCustomAPIClient{cc}
}

func (c *privateCustomAPIClient) UpdateServiceDomains(ctx context.Context, in *UpdateServiceDomainsRequest, opts ...grpc.CallOption) (*UpdateServiceDomainsResponse, error) {
	out := new(UpdateServiceDomainsResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI/UpdateServiceDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateCustomAPIServer is the server API for PrivateCustomAPI service.
type PrivateCustomAPIServer interface {
	// UpdateServiceDomains
	//
	// x-displayName: "UpdateServiceDomains"
	// Private API to update service domains from service.
	UpdateServiceDomains(context.Context, *UpdateServiceDomainsRequest) (*UpdateServiceDomainsResponse, error)
}

// UnimplementedPrivateCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPrivateCustomAPIServer struct {
}

func (*UnimplementedPrivateCustomAPIServer) UpdateServiceDomains(ctx context.Context, req *UpdateServiceDomainsRequest) (*UpdateServiceDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServiceDomains not implemented")
}

func RegisterPrivateCustomAPIServer(s *grpc.Server, srv PrivateCustomAPIServer) {
	s.RegisterService(&_PrivateCustomAPI_serviceDesc, srv)
}

func _PrivateCustomAPI_UpdateServiceDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateCustomAPIServer).UpdateServiceDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI/UpdateServiceDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateCustomAPIServer).UpdateServiceDomains(ctx, req.(*UpdateServiceDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateCustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI",
	HandlerType: (*PrivateCustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateServiceDomains",
			Handler:    _PrivateCustomAPI_UpdateServiceDomains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/cdn_loadbalancer/private_customapi.proto",
}

func (m *UpdateServiceDomainsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateServiceDomainsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateServiceDomainsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceDomains) > 0 {
		for iNdEx := len(m.ServiceDomains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceDomains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPrivateCustomapi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateServiceDomainsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateServiceDomainsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateServiceDomainsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPrivateCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrivateCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateServiceDomainsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	if len(m.ServiceDomains) > 0 {
		for _, e := range m.ServiceDomains {
			l = e.Size()
			n += 1 + l + sovPrivateCustomapi(uint64(l))
		}
	}
	return n
}

func (m *UpdateServiceDomainsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPrivateCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrivateCustomapi(x uint64) (n int) {
	return sovPrivateCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UpdateServiceDomainsRequest) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForServiceDomains := "[]*ServiceDomain{"
	for _, f := range this.ServiceDomains {
		repeatedStringForServiceDomains += strings.Replace(fmt.Sprintf("%v", f), "ServiceDomain", "virtual_host.ServiceDomain", 1) + ","
	}
	repeatedStringForServiceDomains += "}"
	s := strings.Join([]string{`&UpdateServiceDomainsRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Tenant:` + fmt.Sprintf("%v", this.Tenant) + `,`,
		`ServiceDomains:` + repeatedStringForServiceDomains + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdateServiceDomainsResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateServiceDomainsResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringPrivateCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UpdateServiceDomainsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateServiceDomainsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateServiceDomainsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceDomains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceDomains = append(m.ServiceDomains, &virtual_host.ServiceDomain{})
			if err := m.ServiceDomains[len(m.ServiceDomains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdateServiceDomainsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateServiceDomainsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateServiceDomainsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPrivateCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrivateCustomapi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPrivateCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrivateCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrivateCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrivateCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrivateCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrivateCustomapi = fmt.Errorf("proto: unexpected end of group")
)
