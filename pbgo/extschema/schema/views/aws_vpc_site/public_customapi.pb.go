// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/aws_vpc_site/public_customapi.proto

// AWS VPC site
//
// x-displayName: "Configure AWS VPC Site"
// AWS VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC.
// It can be used to either automatically create or Manually assisted site creation in AWS VPC.
//
// View will create following child objects.
//
// * Site
//

package aws_vpc_site

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
	site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
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
	NodeNames []string `protobuf:"bytes,3,rep,name=node_names,json=nodeNames,proto3" json:"node_names,omitempty"`
}

func (m *SetVPCK8SHostnamesRequest) Reset()      { *m = SetVPCK8SHostnamesRequest{} }
func (*SetVPCK8SHostnamesRequest) ProtoMessage() {}
func (*SetVPCK8SHostnamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66fea4ba44356a2b, []int{0}
}
func (m *SetVPCK8SHostnamesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetVPCK8SHostnamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetVPCK8SHostnamesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetVPCK8SHostnamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVPCK8SHostnamesRequest.Merge(m, src)
}
func (m *SetVPCK8SHostnamesRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetVPCK8SHostnamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVPCK8SHostnamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetVPCK8SHostnamesRequest proto.InternalMessageInfo

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
	return fileDescriptor_66fea4ba44356a2b, []int{1}
}
func (m *SetVPCK8SHostnamesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetVPCK8SHostnamesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetVPCK8SHostnamesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetVPCK8SHostnamesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVPCK8SHostnamesResponse.Merge(m, src)
}
func (m *SetVPCK8SHostnamesResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetVPCK8SHostnamesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVPCK8SHostnamesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetVPCK8SHostnamesResponse proto.InternalMessageInfo

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
	VipParamsPerAz []*site.PublishVIPParamsPerAz `protobuf:"bytes,3,rep,name=vip_params_per_az,json=vipParamsPerAz,proto3" json:"vip_params_per_az,omitempty"`
}

func (m *SetVIPInfoRequest) Reset()      { *m = SetVIPInfoRequest{} }
func (*SetVIPInfoRequest) ProtoMessage() {}
func (*SetVIPInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66fea4ba44356a2b, []int{2}
}
func (m *SetVIPInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetVIPInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetVIPInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetVIPInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVIPInfoRequest.Merge(m, src)
}
func (m *SetVIPInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetVIPInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVIPInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetVIPInfoRequest proto.InternalMessageInfo

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

func (m *SetVIPInfoRequest) GetVipParamsPerAz() []*site.PublishVIPParamsPerAz {
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
	return fileDescriptor_66fea4ba44356a2b, []int{3}
}
func (m *SetVIPInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetVIPInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetVIPInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetVIPInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVIPInfoResponse.Merge(m, src)
}
func (m *SetVIPInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetVIPInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVIPInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetVIPInfoResponse proto.InternalMessageInfo

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

func init() {
	proto.RegisterFile("ves.io/schema/views/aws_vpc_site/public_customapi.proto", fileDescriptor_66fea4ba44356a2b)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/aws_vpc_site/public_customapi.proto", fileDescriptor_66fea4ba44356a2b)
}

var fileDescriptor_66fea4ba44356a2b = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x4f, 0x13, 0x4f,
	0x18, 0xed, 0xb4, 0xfc, 0x08, 0xec, 0x2f, 0x31, 0xb2, 0xf1, 0x50, 0x6b, 0x9d, 0x34, 0xbd, 0xa8,
	0x84, 0xdd, 0x89, 0xa0, 0x91, 0xa0, 0x89, 0x01, 0x3c, 0xd8, 0x98, 0xe8, 0x06, 0x22, 0x07, 0x2f,
	0x9b, 0xe9, 0x76, 0xba, 0x1d, 0x69, 0x77, 0xc6, 0x99, 0xd9, 0x45, 0x31, 0x24, 0x86, 0xab, 0x17,
	0x82, 0xff, 0x84, 0xff, 0x00, 0x17, 0xf1, 0xc0, 0xc5, 0xe8, 0xc9, 0x10, 0xbd, 0x70, 0x84, 0xad,
	0x07, 0xbd, 0xf1, 0x27, 0x98, 0x9d, 0xb6, 0xd0, 0x45, 0x49, 0x91, 0xdb, 0xf7, 0xed, 0xfb, 0x5e,
	0xe7, 0xbd, 0xaf, 0x6f, 0xc6, 0xb8, 0x13, 0x11, 0x69, 0x53, 0x86, 0xa4, 0xd7, 0x20, 0x2d, 0x8c,
	0x22, 0x4a, 0x56, 0x24, 0xc2, 0x2b, 0xd2, 0x8d, 0xb8, 0xe7, 0x4a, 0xaa, 0x08, 0xe2, 0x61, 0xb5,
	0x49, 0x3d, 0xd7, 0x0b, 0xa5, 0x62, 0x2d, 0xcc, 0xa9, 0xcd, 0x05, 0x53, 0xcc, 0x2c, 0x75, 0x88,
	0x76, 0x87, 0x68, 0x6b, 0xa2, 0xdd, 0x4f, 0x2c, 0x58, 0x3e, 0x55, 0x8d, 0xb0, 0x6a, 0x7b, 0xac,
	0x85, 0x7c, 0xe6, 0x33, 0xa4, 0x89, 0xd5, 0xb0, 0xae, 0x3b, 0xdd, 0xe8, 0xaa, 0xf3, 0x83, 0x85,
	0xa2, 0xcf, 0x98, 0xdf, 0x24, 0x08, 0x73, 0x8a, 0x70, 0x10, 0x30, 0x85, 0x15, 0x65, 0x81, 0xec,
	0xa2, 0x57, 0xd2, 0x3a, 0x19, 0xef, 0x07, 0x61, 0x1a, 0xd4, 0xb2, 0xd5, 0x2b, 0x4e, 0x7a, 0x78,
	0xf1, 0x84, 0x49, 0xdc, 0xa4, 0x35, 0xac, 0x48, 0x17, 0x2d, 0x9f, 0x40, 0x89, 0x24, 0x41, 0x74,
	0xe2, 0x84, 0xd2, 0x9f, 0x6b, 0x72, 0xd3, 0x13, 0xd6, 0xc0, 0x45, 0xb2, 0xea, 0x73, 0xe2, 0xa9,
	0xee, 0xf8, 0xc4, 0xc0, 0xf1, 0x3e, 0x03, 0xe5, 0xa6, 0x71, 0x79, 0x91, 0xa8, 0x25, 0x67, 0xfe,
	0xd1, 0xf4, 0xe2, 0x43, 0x26, 0x55, 0x80, 0x5b, 0x44, 0x2e, 0x90, 0x17, 0x21, 0x91, 0xca, 0x2c,
	0x1a, 0xa3, 0xba, 0xe7, 0xd8, 0x23, 0x79, 0x50, 0x02, 0xd7, 0x47, 0x17, 0x8e, 0x3f, 0x98, 0xa6,
	0x31, 0x94, 0x34, 0xf9, 0xac, 0x06, 0x74, 0x6d, 0x5e, 0x35, 0x8c, 0x80, 0xd5, 0x88, 0xab, 0xa7,
	0xf2, 0xb9, 0x52, 0x4e, 0x53, 0x58, 0x8d, 0x3c, 0x4e, 0x3e, 0x94, 0x8b, 0x46, 0xe1, 0x6f, 0xa7,
	0x49, 0xce, 0x02, 0x49, 0xca, 0x5b, 0xc0, 0x18, 0x4b, 0xe0, 0x8a, 0x53, 0x09, 0xea, 0xec, 0xfc,
	0x22, 0x02, 0x63, 0x2c, 0xa2, 0xdc, 0xe5, 0x58, 0xe0, 0x96, 0x74, 0x39, 0x11, 0x2e, 0x5e, 0xd5,
	0x5a, 0xfe, 0x9f, 0xbc, 0x61, 0xa7, 0xc3, 0x95, 0xec, 0xc3, 0x76, 0x92, 0x1c, 0xca, 0xc6, 0x52,
	0xc5, 0x71, 0x34, 0xc5, 0x21, 0x62, 0x76, 0x75, 0x2e, 0xff, 0xe1, 0xd7, 0x4e, 0x6e, 0x64, 0x13,
	0xfc, 0x37, 0x9e, 0xbb, 0x39, 0x31, 0x15, 0xef, 0x7f, 0xca, 0x0d, 0x6f, 0x7e, 0x04, 0xd9, 0x11,
	0xb0, 0x70, 0x21, 0xa2, 0xbc, 0x6f, 0xb2, 0x7c, 0xc9, 0x30, 0xfb, 0x65, 0x77, 0xdc, 0x4c, 0x6e,
	0x0d, 0x19, 0xa3, 0xf3, 0x3a, 0xda, 0xb3, 0x4e, 0xc5, 0x7c, 0x9b, 0xed, 0x0c, 0xa5, 0xad, 0x9b,
	0x77, 0xed, 0x41, 0x61, 0xb7, 0x4f, 0xfd, 0x7b, 0x0a, 0xf7, 0xce, 0x47, 0xee, 0x6e, 0x7b, 0x35,
	0xfe, 0x9c, 0x4f, 0xb2, 0x6f, 0x51, 0x66, 0xd1, 0xa0, 0x2e, 0xb0, 0x54, 0x22, 0xf4, 0x54, 0x28,
	0x88, 0xb5, 0x22, 0xa8, 0x22, 0xeb, 0xdf, 0x7f, 0xbc, 0xcb, 0x3e, 0x2d, 0x3b, 0xdd, 0x6b, 0x8a,
	0x8e, 0x76, 0x2e, 0xd1, 0xeb, 0xa3, 0x7a, 0x2d, 0x9d, 0x2c, 0x0d, 0xac, 0x21, 0xa9, 0x98, 0xc0,
	0x3e, 0x41, 0x92, 0x28, 0x8d, 0x2d, 0x4f, 0x4b, 0xb7, 0xd1, 0x13, 0x30, 0x03, 0xc6, 0xcd, 0x18,
	0x18, 0xc6, 0xf1, 0xca, 0xcc, 0xa9, 0xb3, 0x19, 0x49, 0xe5, 0xa2, 0x70, 0xeb, 0xdf, 0x48, 0x5d,
	0xd7, 0xf4, 0x2c, 0xae, 0x1f, 0x94, 0xef, 0x9f, 0xc7, 0x75, 0xe2, 0x96, 0x72, 0x97, 0x06, 0x75,
	0x36, 0x03, 0xc6, 0x0b, 0xb7, 0x77, 0xb6, 0x41, 0xee, 0xdb, 0x36, 0xb8, 0x36, 0x50, 0xe7, 0x13,
	0x7d, 0x7f, 0xd7, 0xbf, 0xe6, 0xb3, 0x17, 0xc1, 0xdc, 0x06, 0xd8, 0x3d, 0x80, 0x99, 0xbd, 0x03,
	0x98, 0x39, 0x3c, 0x80, 0xe0, 0x4d, 0x0c, 0xc1, 0xfb, 0x18, 0x82, 0x2f, 0x31, 0x04, 0xbb, 0x31,
	0x04, 0xfb, 0x31, 0x04, 0x3f, 0x63, 0x98, 0x39, 0x8c, 0x21, 0xd8, 0x68, 0xc3, 0xcc, 0x4e, 0x1b,
	0x82, 0xdd, 0x36, 0xcc, 0xec, 0xb5, 0x61, 0xe6, 0xd9, 0x92, 0xcf, 0xf8, 0xb2, 0x6f, 0x47, 0xac,
	0xa9, 0x88, 0x10, 0xd8, 0x0e, 0x25, 0xd2, 0x45, 0x9d, 0x89, 0x96, 0xc5, 0x05, 0x8b, 0x68, 0x8d,
	0x08, 0xab, 0x07, 0x23, 0x5e, 0xf5, 0x19, 0x22, 0x2f, 0x55, 0xef, 0x71, 0x3b, 0xed, 0xc1, 0xa8,
	0x0e, 0xeb, 0xb7, 0x62, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x84, 0x64, 0x48, 0xd3,
	0x05, 0x00, 0x00,
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

// CustomAPIClient is the client API for CustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.aws_vpc_site.CustomAPI/SetVPCK8SHostnames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) SetVIPInfo(ctx context.Context, in *SetVIPInfoRequest, opts ...grpc.CallOption) (*SetVIPInfoResponse, error) {
	out := new(SetVIPInfoResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.aws_vpc_site.CustomAPI/SetVIPInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
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

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) SetVPCK8SHostnames(ctx context.Context, req *SetVPCK8SHostnamesRequest) (*SetVPCK8SHostnamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVPCK8SHostnames not implemented")
}
func (*UnimplementedCustomAPIServer) SetVIPInfo(ctx context.Context, req *SetVIPInfoRequest) (*SetVIPInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVIPInfo not implemented")
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
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVPCK8SHostnamesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetVPCK8SHostnamesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeNames) > 0 {
		for iNdEx := len(m.NodeNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.NodeNames[iNdEx])
			copy(dAtA[i:], m.NodeNames[iNdEx])
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.NodeNames[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetVPCK8SHostnamesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVPCK8SHostnamesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetVPCK8SHostnamesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SetVIPInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVIPInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetVIPInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VipParamsPerAz) > 0 {
		for iNdEx := len(m.VipParamsPerAz) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VipParamsPerAz[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetVIPInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetVIPInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetVIPInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetVPCK8SHostnamesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
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
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SetVIPInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
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
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
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
	repeatedStringForVipParamsPerAz := "[]*PublishVIPParamsPerAz{"
	for _, f := range this.VipParamsPerAz {
		repeatedStringForVipParamsPerAz += strings.Replace(fmt.Sprintf("%v", f), "PublishVIPParamsPerAz", "site.PublishVIPParamsPerAz", 1) + ","
	}
	repeatedStringForVipParamsPerAz += "}"
	s := strings.Join([]string{`&SetVIPInfoRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`VipParamsPerAz:` + repeatedStringForVipParamsPerAz + `,`,
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
			wire |= uint64(b&0x7F) << shift
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
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
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
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
			if (iNdEx + skippy) < 0 {
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
			wire |= uint64(b&0x7F) << shift
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
			if (iNdEx + skippy) < 0 {
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
			wire |= uint64(b&0x7F) << shift
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
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
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
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
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VipParamsPerAz = append(m.VipParamsPerAz, &site.PublishVIPParamsPerAz{})
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
			if (iNdEx + skippy) < 0 {
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
			wire |= uint64(b&0x7F) << shift
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
			if (iNdEx + skippy) < 0 {
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
	depth := 0
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
		case 1:
			iNdEx += 8
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
			if length < 0 {
				return 0, ErrInvalidLengthPublicCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapi = fmt.Errorf("proto: unexpected end of group")
)
