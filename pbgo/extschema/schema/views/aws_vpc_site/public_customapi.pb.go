// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/aws_vpc_site/public_customapi.proto

package aws_vpc_site

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request to configure VPC k8s node hostname set
//
// x-displayName: "Request to configure VPC k8s node hostname set"
// Request to configure VPC k8s node hostname set
type SetVPCK8SHostnamesRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "default"
	// Namespace for the object to be configured
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "aws-vpc-site-1"
	// Name of the object to be configured
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Hostname of K8S nodes
	//
	// x-displayName: "Hostname of K8S nodes"
	// x-example: "master-0"
	// Hostname of K8S nodes for deinition of Mayastore pools
	NodeNames []string `protobuf:"bytes,3,rep,name=node_names,json=nodeNames" json:"node_names,omitempty"`
}

func (m *SetVPCK8SHostnamesRequest) Reset()      { *m = SetVPCK8SHostnamesRequest{} }
func (*SetVPCK8SHostnamesRequest) ProtoMessage() {}
func (*SetVPCK8SHostnamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{0}
}

func (m *SetVPCK8SHostnamesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SetVPCK8SHostnamesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetVPCK8SHostnamesRequest) GetNodeNames() []string {
	if m != nil {
		return m.NodeNames
	}
	return nil
}

// Response to configure VPC k8s node hostname set
//
// x-displayName: "Response to configure VPC k8s node hostname set"
type SetVPCK8SHostnamesResponse struct {
}

func (m *SetVPCK8SHostnamesResponse) Reset()      { *m = SetVPCK8SHostnamesResponse{} }
func (*SetVPCK8SHostnamesResponse) ProtoMessage() {}
func (*SetVPCK8SHostnamesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{1}
}

// Request to configure AWS VPC Site VIP information
//
// x-displayName: "Request to configure AWS VPC Site VIP information"
// Request to configure AWS VPC Site VIP information
type SetVIPInfoRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "default"
	// Namespace for the object to be configured
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "aws-vpc-site-1"
	// Name of the object to be configured
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// VIP Params Per AZ
	//
	// x-displayName: "VIP Params Per AZ"
	// x-example: "master-0"
	// VIP Parameters per AZ
	VipParamsPerAz []*ves_io_schema_site.PublishVIPParamsPerAz `protobuf:"bytes,3,rep,name=vip_params_per_az,json=vipParamsPerAz" json:"vip_params_per_az,omitempty"`
}

func (m *SetVIPInfoRequest) Reset()      { *m = SetVIPInfoRequest{} }
func (*SetVIPInfoRequest) ProtoMessage() {}
func (*SetVIPInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{2}
}

func (m *SetVIPInfoRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SetVIPInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetVIPInfoRequest) GetVipParamsPerAz() []*ves_io_schema_site.PublishVIPParamsPerAz {
	if m != nil {
		return m.VipParamsPerAz
	}
	return nil
}

// Response to configure AWS VPC Site VIP Information
//
// x-displayName: "Response to configure AWS VPC Site VIP Information"
type SetVIPInfoResponse struct {
}

func (m *SetVIPInfoResponse) Reset()      { *m = SetVIPInfoResponse{} }
func (*SetVIPInfoResponse) ProtoMessage() {}
func (*SetVIPInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{3}
}

func init() {
	proto.RegisterType((*SetVPCK8SHostnamesRequest)(nil), "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesRequest")
	golang_proto.RegisterType((*SetVPCK8SHostnamesRequest)(nil), "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesRequest")
	proto.RegisterType((*SetVPCK8SHostnamesResponse)(nil), "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesResponse")
	golang_proto.RegisterType((*SetVPCK8SHostnamesResponse)(nil), "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesResponse")
	proto.RegisterType((*SetVIPInfoRequest)(nil), "ves.io.schema.views.aws_vpc_site.SetVIPInfoRequest")
	golang_proto.RegisterType((*SetVIPInfoRequest)(nil), "ves.io.schema.views.aws_vpc_site.SetVIPInfoRequest")
	proto.RegisterType((*SetVIPInfoResponse)(nil), "ves.io.schema.views.aws_vpc_site.SetVIPInfoResponse")
	golang_proto.RegisterType((*SetVIPInfoResponse)(nil), "ves.io.schema.views.aws_vpc_site.SetVIPInfoResponse")
}
func (this *SetVPCK8SHostnamesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetVPCK8SHostnamesRequest)
	if !ok {
		that2, ok := that.(SetVPCK8SHostnamesRequest)
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
	if len(this.NodeNames) != len(that1.NodeNames) {
		return false
	}
	for i := range this.NodeNames {
		if this.NodeNames[i] != that1.NodeNames[i] {
			return false
		}
	}
	return true
}
func (this *SetVPCK8SHostnamesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetVPCK8SHostnamesResponse)
	if !ok {
		that2, ok := that.(SetVPCK8SHostnamesResponse)
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
func (this *SetVIPInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetVIPInfoRequest)
	if !ok {
		that2, ok := that.(SetVIPInfoRequest)
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
	if len(this.VipParamsPerAz) != len(that1.VipParamsPerAz) {
		return false
	}
	for i := range this.VipParamsPerAz {
		if !this.VipParamsPerAz[i].Equal(that1.VipParamsPerAz[i]) {
			return false
		}
	}
	return true
}
func (this *SetVIPInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetVIPInfoResponse)
	if !ok {
		that2, ok := that.(SetVIPInfoResponse)
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
func (this *SetVPCK8SHostnamesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&aws_vpc_site.SetVPCK8SHostnamesRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "NodeNames: "+fmt.Sprintf("%#v", this.NodeNames)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetVPCK8SHostnamesResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&aws_vpc_site.SetVPCK8SHostnamesResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetVIPInfoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&aws_vpc_site.SetVIPInfoRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.VipParamsPerAz != nil {
		s = append(s, "VipParamsPerAz: "+fmt.Sprintf("%#v", this.VipParamsPerAz)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetVIPInfoResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&aws_vpc_site.SetVIPInfoResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
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

// Client API for CustomAPI service

type CustomAPIClient interface {
	// Configure VPC k8s hostnames
	//
	// x-displayName: "Configure VPC k8s hostnames"
	// Configure VPC k8s node hostname set
	SetVPCK8SHostnames(ctx context.Context, in *SetVPCK8SHostnamesRequest, opts ...grpc.CallOption) (*SetVPCK8SHostnamesResponse, error)
	// Configure AWS VPC Site VIP Information
	//
	// x-displayName: "Configure AWS VPC Site VIP Information"
	// Configure AWS VPC Site VIP Information
	SetVIPInfo(ctx context.Context, in *SetVIPInfoRequest, opts ...grpc.CallOption) (*SetVIPInfoResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) SetVPCK8SHostnames(ctx context.Context, in *SetVPCK8SHostnamesRequest, opts ...grpc.CallOption) (*SetVPCK8SHostnamesResponse, error) {
	out := new(SetVPCK8SHostnamesResponse)
	err := grpc.Invoke(ctx, "/ves.io.schema.views.aws_vpc_site.CustomAPI/SetVPCK8SHostnames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) SetVIPInfo(ctx context.Context, in *SetVIPInfoRequest, opts ...grpc.CallOption) (*SetVIPInfoResponse, error) {
	out := new(SetVIPInfoResponse)
	err := grpc.Invoke(ctx, "/ves.io.schema.views.aws_vpc_site.CustomAPI/SetVIPInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomAPI service

type CustomAPIServer interface {
	// Configure VPC k8s hostnames
	//
	// x-displayName: "Configure VPC k8s hostnames"
	// Configure VPC k8s node hostname set
	SetVPCK8SHostnames(context.Context, *SetVPCK8SHostnamesRequest) (*SetVPCK8SHostnamesResponse, error)
	// Configure AWS VPC Site VIP Information
	//
	// x-displayName: "Configure AWS VPC Site VIP Information"
	// Configure AWS VPC Site VIP Information
	SetVIPInfo(context.Context, *SetVIPInfoRequest) (*SetVIPInfoResponse, error)
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_SetVPCK8SHostnames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVPCK8SHostnamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).SetVPCK8SHostnames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.aws_vpc_site.CustomAPI/SetVPCK8SHostnames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).SetVPCK8SHostnames(ctx, req.(*SetVPCK8SHostnamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_SetVIPInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVIPInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).SetVIPInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.aws_vpc_site.CustomAPI/SetVIPInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).SetVIPInfo(ctx, req.(*SetVIPInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.aws_vpc_site.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetVPCK8SHostnames",
			Handler:    _CustomAPI_SetVPCK8SHostnames_Handler,
		},
		{
			MethodName: "SetVIPInfo",
			Handler:    _CustomAPI_SetVIPInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/aws_vpc_site/public_customapi.proto",
}

func (m *SetVPCK8SHostnamesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVPCK8SHostnamesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.NodeNames) > 0 {
		for _, s := range m.NodeNames {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *SetVPCK8SHostnamesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVPCK8SHostnamesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SetVIPInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVIPInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.VipParamsPerAz) > 0 {
		for _, msg := range m.VipParamsPerAz {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SetVIPInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVIPInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SetVPCK8SHostnamesRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if len(m.NodeNames) > 0 {
		for _, s := range m.NodeNames {
			l = len(s)
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func (m *SetVPCK8SHostnamesResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SetVIPInfoRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if len(m.VipParamsPerAz) > 0 {
		for _, e := range m.VipParamsPerAz {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func (m *SetVIPInfoResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SetVPCK8SHostnamesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetVPCK8SHostnamesRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`NodeNames:` + fmt.Sprintf("%v", this.NodeNames) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetVPCK8SHostnamesResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetVPCK8SHostnamesResponse{`,
		`}`,
	}, "")
	return s
}
func (this *SetVIPInfoRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetVIPInfoRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`VipParamsPerAz:` + strings.Replace(fmt.Sprintf("%v", this.VipParamsPerAz), "PublishVIPParamsPerAz", "ves_io_schema_site.PublishVIPParamsPerAz", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetVIPInfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetVIPInfoResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SetVPCK8SHostnamesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetVPCK8SHostnamesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetVPCK8SHostnamesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
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
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeNames = append(m.NodeNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *SetVPCK8SHostnamesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetVPCK8SHostnamesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetVPCK8SHostnamesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *SetVIPInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetVIPInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetVIPInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
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
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VipParamsPerAz", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VipParamsPerAz = append(m.VipParamsPerAz, &ves_io_schema_site.PublishVIPParamsPerAz{})
			if err := m.VipParamsPerAz[len(m.VipParamsPerAz)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *SetVIPInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetVIPInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetVIPInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPublicCustomapi
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPublicCustomapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPublicCustomapi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPublicCustomapi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPublicCustomapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/views/aws_vpc_site/public_customapi.proto", fileDescriptorPublicCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/aws_vpc_site/public_customapi.proto", fileDescriptorPublicCustomapi)
}

var fileDescriptorPublicCustomapi = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xce, 0x25, 0xfd, 0xfd, 0x44, 0x8d, 0x40, 0xd4, 0x62, 0x08, 0x21, 0x58, 0x51, 0x16, 0xa0,
	0xaa, 0x7d, 0xa2, 0x05, 0x51, 0x15, 0x24, 0xd4, 0x96, 0x81, 0x08, 0x09, 0xac, 0x56, 0x74, 0x60,
	0xb1, 0x2e, 0xce, 0x8b, 0x73, 0x34, 0xf1, 0x1d, 0x77, 0x67, 0x17, 0x8a, 0x2a, 0xa1, 0x4e, 0x48,
	0x2c, 0x08, 0xfe, 0x09, 0xc4, 0xce, 0x42, 0x97, 0x6e, 0x30, 0xa1, 0x0a, 0x16, 0x46, 0xea, 0x32,
	0xc0, 0xd6, 0x3f, 0x01, 0xf9, 0x92, 0xb4, 0x71, 0xa1, 0x4a, 0xe9, 0xf6, 0x9e, 0xbf, 0xf7, 0xe5,
	0xbe, 0xef, 0xe5, 0xbb, 0x33, 0xae, 0xc7, 0x20, 0x1d, 0xca, 0xb0, 0xf4, 0x5b, 0xd0, 0x21, 0x38,
	0xa6, 0xb0, 0x22, 0x31, 0x59, 0x91, 0x5e, 0xcc, 0x7d, 0x4f, 0x52, 0x05, 0x98, 0x47, 0xf5, 0x36,
	0xf5, 0x3d, 0x3f, 0x92, 0x8a, 0x75, 0x08, 0xa7, 0x0e, 0x17, 0x4c, 0x31, 0xb3, 0xd2, 0x25, 0x3a,
	0x5d, 0xa2, 0xa3, 0x89, 0xce, 0x20, 0xb1, 0x64, 0x07, 0x54, 0xb5, 0xa2, 0xba, 0xe3, 0xb3, 0x0e,
	0x0e, 0x58, 0xc0, 0xb0, 0x26, 0xd6, 0xa3, 0xa6, 0xee, 0x74, 0xa3, 0xab, 0xee, 0x0f, 0x96, 0xca,
	0x01, 0x63, 0x41, 0x1b, 0x30, 0xe1, 0x14, 0x93, 0x30, 0x64, 0x8a, 0x28, 0xca, 0x42, 0xd9, 0x43,
	0xcf, 0x67, 0x75, 0x32, 0x3e, 0x08, 0x5a, 0x59, 0x50, 0xcb, 0x56, 0x4f, 0x39, 0xf4, 0xf1, 0xf2,
	0x01, 0x93, 0xa4, 0x4d, 0x1b, 0x44, 0x41, 0x0f, 0xad, 0x1e, 0x40, 0x41, 0x42, 0x18, 0x1f, 0x38,
	0xa1, 0xf2, 0xe7, 0x9a, 0xbc, 0xec, 0x84, 0x3d, 0x74, 0x91, 0xac, 0xfe, 0x08, 0x7c, 0xd5, 0x1b,
	0x9f, 0x18, 0x3a, 0x3e, 0x60, 0xa0, 0xda, 0x36, 0xce, 0x2d, 0x82, 0x5a, 0x72, 0xe7, 0xef, 0x4e,
	0x2f, 0xde, 0x61, 0x52, 0x85, 0xa4, 0x03, 0x72, 0x01, 0x1e, 0x47, 0x20, 0x95, 0x59, 0x36, 0x46,
	0x75, 0xcf, 0x89, 0x0f, 0x45, 0x54, 0x41, 0x97, 0x46, 0x17, 0xf6, 0x3f, 0x98, 0xa6, 0x31, 0x92,
	0x36, 0xc5, 0xbc, 0x06, 0x74, 0x6d, 0x5e, 0x30, 0x8c, 0x90, 0x35, 0xc0, 0xd3, 0x53, 0xc5, 0x42,
	0xa5, 0xa0, 0x29, 0xac, 0x01, 0xf7, 0xd2, 0x0f, 0xd5, 0xb2, 0x51, 0xfa, 0xdb, 0x69, 0x92, 0xb3,
	0x50, 0x42, 0xf5, 0x1d, 0x32, 0xc6, 0x52, 0xb8, 0xe6, 0xd6, 0xc2, 0x26, 0x3b, 0xbe, 0x08, 0xdf,
	0x18, 0x8b, 0x29, 0xf7, 0x38, 0x11, 0xa4, 0x23, 0x3d, 0x0e, 0xc2, 0x23, 0xab, 0x5a, 0xcb, 0xc9,
	0xc9, 0xcb, 0x4e, 0x36, 0x5c, 0xe9, 0x3e, 0x1c, 0x37, 0xcd, 0xa1, 0x6c, 0x2d, 0xd5, 0x5c, 0x57,
	0x53, 0x5c, 0x10, 0xb3, 0xab, 0x73, 0xa7, 0x3e, 0xfc, 0xda, 0x2c, 0x9c, 0x78, 0x8d, 0xfe, 0x1b,
	0x2f, 0x5c, 0x99, 0x98, 0x5a, 0x38, 0x1d, 0x53, 0x3e, 0x00, 0x57, 0xcf, 0x1a, 0xe6, 0xa0, 0xd6,
	0xae, 0x85, 0xc9, 0xf7, 0x23, 0xc6, 0xe8, 0xbc, 0xce, 0xf3, 0xac, 0x5b, 0x33, 0x5f, 0xe6, 0xbb,
	0x43, 0x59, 0xbf, 0xe6, 0x0d, 0x67, 0x58, 0xc2, 0x9d, 0x43, 0xff, 0x93, 0xd2, 0xcd, 0xe3, 0x91,
	0x7b, 0x2b, 0x5e, 0x4d, 0x3e, 0x16, 0xd3, 0xc0, 0xdb, 0x94, 0xd9, 0x34, 0x6c, 0x0a, 0x22, 0x95,
	0x88, 0x7c, 0x15, 0x09, 0xb0, 0x57, 0x04, 0x55, 0xb0, 0xfe, 0xf5, 0xc7, 0x9b, 0xfc, 0x83, 0xaa,
	0xdb, 0xbb, 0x9b, 0x78, 0x6f, 0xd1, 0x12, 0x3f, 0xdb, 0xab, 0xd7, 0xb2, 0x71, 0xd2, 0xc0, 0x1a,
	0x96, 0x8a, 0x09, 0x12, 0x00, 0x96, 0xa0, 0x34, 0xb6, 0x3c, 0x2d, 0xbd, 0x56, 0x5f, 0xc0, 0x0c,
	0x1a, 0x37, 0x13, 0x64, 0x18, 0xfb, 0x2b, 0x33, 0xa7, 0x8e, 0x66, 0x24, 0x13, 0x86, 0xd2, 0xd5,
	0x7f, 0x23, 0xf5, 0x5c, 0xd3, 0xa3, 0xb8, 0xbe, 0x5d, 0xbd, 0x75, 0x1c, 0xd7, 0xa9, 0x5b, 0xca,
	0x3d, 0x1a, 0x36, 0xd9, 0x0c, 0x1a, 0x2f, 0x5d, 0xdb, 0xdc, 0x40, 0x85, 0x2f, 0x1b, 0xe8, 0xe2,
	0x50, 0x9d, 0xf7, 0xf5, 0xa5, 0x5d, 0xff, 0x5c, 0xcc, 0x9f, 0x41, 0x73, 0x2f, 0xd0, 0xd6, 0xb6,
	0x95, 0xfb, 0xb6, 0x6d, 0xe5, 0x76, 0xb7, 0x2d, 0xf4, 0x3c, 0xb1, 0xd0, 0xdb, 0xc4, 0x42, 0x9f,
	0x12, 0x0b, 0x6d, 0x25, 0x16, 0xfa, 0x9e, 0x58, 0xe8, 0x67, 0x62, 0xe5, 0x76, 0x13, 0x0b, 0xbd,
	0xda, 0xb1, 0x72, 0x9b, 0x3b, 0x16, 0x7a, 0xb8, 0x14, 0x30, 0xbe, 0x1c, 0x38, 0x31, 0x6b, 0x2b,
	0x10, 0x82, 0x38, 0x91, 0xc4, 0xba, 0x68, 0x32, 0xd1, 0xb1, 0xb9, 0x60, 0x31, 0x6d, 0x80, 0xb0,
	0xfb, 0x30, 0xe6, 0xf5, 0x80, 0x61, 0x78, 0xa2, 0xfa, 0x2f, 0xd9, 0x61, 0xaf, 0x43, 0xfd, 0x7f,
	0xfd, 0x30, 0x4c, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x60, 0xdf, 0xd7, 0xc0, 0x05, 0x00,
	0x00,
}
