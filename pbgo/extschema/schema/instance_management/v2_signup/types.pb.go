// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/instance_management/v2_signup/types.proto

package v2_signup

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/infraprotect_information"
	signup "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/signup"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Signup specification
//
// x-displayName: "Signup"
// desired state of Signup
type GlobalSpecType struct {
	// Company Details
	//
	// x-displayName: "Company Details"
	// x-required
	// details of the company
	CompanyDetails *signup.CompanyMeta `protobuf:"bytes,1,opt,name=company_details,json=companyDetails,proto3" json:"company_details,omitempty"`
	// User Details
	//
	// x-displayName: "User Details"
	// x-required
	// details of the user
	UserDetails *signup.UserMeta `protobuf:"bytes,2,opt,name=user_details,json=userDetails,proto3" json:"user_details,omitempty"`
	// Account Details
	//
	// x-displayName: "Account Details"
	// x-required
	// details of the new F5XC account to be created
	AccountDetails *signup.AccountMeta `protobuf:"bytes,3,opt,name=account_details,json=accountDetails,proto3" json:"account_details,omitempty"`
	// Billing Details
	//
	// x-displayName: "Billing Details"
	// x-required
	// details about the billing of the account
	BillingDetails *signup.BillingMeta `protobuf:"bytes,4,opt,name=billing_details,json=billingDetails,proto3" json:"billing_details,omitempty"`
	// Source Choice
	//
	// x-displayName: "Source Choice"
	// x-required
	// origin of the request
	//
	// Types that are valid to be assigned to SourceChoice:
	//	*GlobalSpecType_SourcePublic
	//	*GlobalSpecType_SourceInternalSre
	//	*GlobalSpecType_SourceInternalScaling
	//	*GlobalSpecType_SourcePlanTransition
	//	*GlobalSpecType_SourceMsp
	//	*GlobalSpecType_SourceInternalSso
	SourceChoice isGlobalSpecType_SourceChoice `protobuf_oneof:"source_choice"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_142ba11d320bb60e, []int{0}
}
func (m *GlobalSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GlobalSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSpecType.Merge(m, src)
}
func (m *GlobalSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSpecType proto.InternalMessageInfo

type isGlobalSpecType_SourceChoice interface {
	isGlobalSpecType_SourceChoice()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type GlobalSpecType_SourcePublic struct {
	SourcePublic *signup.SourcePublic `protobuf:"bytes,6,opt,name=source_public,json=sourcePublic,proto3,oneof" json:"source_public,omitempty"`
}
type GlobalSpecType_SourceInternalSre struct {
	SourceInternalSre *signup.SourceInternalSre `protobuf:"bytes,7,opt,name=source_internal_sre,json=sourceInternalSre,proto3,oneof" json:"source_internal_sre,omitempty"`
}
type GlobalSpecType_SourceInternalScaling struct {
	SourceInternalScaling *signup.SourceInternalScaling `protobuf:"bytes,8,opt,name=source_internal_scaling,json=sourceInternalScaling,proto3,oneof" json:"source_internal_scaling,omitempty"`
}
type GlobalSpecType_SourcePlanTransition struct {
	SourcePlanTransition *signup.SourcePlanTransition `protobuf:"bytes,9,opt,name=source_plan_transition,json=sourcePlanTransition,proto3,oneof" json:"source_plan_transition,omitempty"`
}
type GlobalSpecType_SourceMsp struct {
	SourceMsp *signup.SourceMsp `protobuf:"bytes,10,opt,name=source_msp,json=sourceMsp,proto3,oneof" json:"source_msp,omitempty"`
}
type GlobalSpecType_SourceInternalSso struct {
	SourceInternalSso *signup.SourceInternalSso `protobuf:"bytes,11,opt,name=source_internal_sso,json=sourceInternalSso,proto3,oneof" json:"source_internal_sso,omitempty"`
}

func (*GlobalSpecType_SourcePublic) isGlobalSpecType_SourceChoice()          {}
func (*GlobalSpecType_SourceInternalSre) isGlobalSpecType_SourceChoice()     {}
func (*GlobalSpecType_SourceInternalScaling) isGlobalSpecType_SourceChoice() {}
func (*GlobalSpecType_SourcePlanTransition) isGlobalSpecType_SourceChoice()  {}
func (*GlobalSpecType_SourceMsp) isGlobalSpecType_SourceChoice()             {}
func (*GlobalSpecType_SourceInternalSso) isGlobalSpecType_SourceChoice()     {}

func (m *GlobalSpecType) GetSourceChoice() isGlobalSpecType_SourceChoice {
	if m != nil {
		return m.SourceChoice
	}
	return nil
}

func (m *GlobalSpecType) GetCompanyDetails() *signup.CompanyMeta {
	if m != nil {
		return m.CompanyDetails
	}
	return nil
}

func (m *GlobalSpecType) GetUserDetails() *signup.UserMeta {
	if m != nil {
		return m.UserDetails
	}
	return nil
}

func (m *GlobalSpecType) GetAccountDetails() *signup.AccountMeta {
	if m != nil {
		return m.AccountDetails
	}
	return nil
}

func (m *GlobalSpecType) GetBillingDetails() *signup.BillingMeta {
	if m != nil {
		return m.BillingDetails
	}
	return nil
}

func (m *GlobalSpecType) GetSourcePublic() *signup.SourcePublic {
	if x, ok := m.GetSourceChoice().(*GlobalSpecType_SourcePublic); ok {
		return x.SourcePublic
	}
	return nil
}

func (m *GlobalSpecType) GetSourceInternalSre() *signup.SourceInternalSre {
	if x, ok := m.GetSourceChoice().(*GlobalSpecType_SourceInternalSre); ok {
		return x.SourceInternalSre
	}
	return nil
}

func (m *GlobalSpecType) GetSourceInternalScaling() *signup.SourceInternalScaling {
	if x, ok := m.GetSourceChoice().(*GlobalSpecType_SourceInternalScaling); ok {
		return x.SourceInternalScaling
	}
	return nil
}

func (m *GlobalSpecType) GetSourcePlanTransition() *signup.SourcePlanTransition {
	if x, ok := m.GetSourceChoice().(*GlobalSpecType_SourcePlanTransition); ok {
		return x.SourcePlanTransition
	}
	return nil
}

func (m *GlobalSpecType) GetSourceMsp() *signup.SourceMsp {
	if x, ok := m.GetSourceChoice().(*GlobalSpecType_SourceMsp); ok {
		return x.SourceMsp
	}
	return nil
}

func (m *GlobalSpecType) GetSourceInternalSso() *signup.SourceInternalSso {
	if x, ok := m.GetSourceChoice().(*GlobalSpecType_SourceInternalSso); ok {
		return x.SourceInternalSso
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GlobalSpecType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GlobalSpecType_SourcePublic)(nil),
		(*GlobalSpecType_SourceInternalSre)(nil),
		(*GlobalSpecType_SourceInternalScaling)(nil),
		(*GlobalSpecType_SourcePlanTransition)(nil),
		(*GlobalSpecType_SourceMsp)(nil),
		(*GlobalSpecType_SourceInternalSso)(nil),
	}
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.instance_management.v2_signup.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/instance_management/v2_signup/types.proto", fileDescriptor_142ba11d320bb60e)
}

var fileDescriptor_142ba11d320bb60e = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd4, 0x3a,
	0x14, 0xc6, 0xe3, 0x7b, 0xdb, 0xb9, 0x6d, 0xda, 0xdb, 0x42, 0xf8, 0x97, 0x16, 0xe4, 0x96, 0xb2,
	0x00, 0x51, 0x35, 0x91, 0x86, 0x05, 0x5b, 0x5a, 0x90, 0x98, 0x56, 0xaa, 0x84, 0xa6, 0x65, 0x01,
	0x9b, 0xc8, 0x71, 0x4f, 0x53, 0x8b, 0xc4, 0xb6, 0x6c, 0x67, 0xa0, 0x3b, 0x1e, 0x81, 0x3d, 0x2f,
	0xc0, 0x3b, 0x74, 0xc3, 0x92, 0xe5, 0x2c, 0xbb, 0xa4, 0xa9, 0x84, 0x58, 0xf6, 0x11, 0xd0, 0x38,
	0xc9, 0x30, 0x53, 0xa2, 0x51, 0xd9, 0xe5, 0xe4, 0xfb, 0xbe, 0x9f, 0x8f, 0x8f, 0xac, 0xe3, 0x3e,
	0xed, 0x81, 0x0e, 0x98, 0x08, 0x35, 0x3d, 0x82, 0x8c, 0x84, 0x8c, 0x6b, 0x43, 0x38, 0x85, 0x28,
	0x23, 0x9c, 0x24, 0x90, 0x01, 0x37, 0x61, 0xaf, 0x1d, 0x69, 0x96, 0xf0, 0x5c, 0x86, 0xe6, 0x58,
	0x82, 0x0e, 0xa4, 0x12, 0x46, 0x78, 0xeb, 0x65, 0x30, 0x28, 0x83, 0x41, 0x43, 0x30, 0x18, 0x06,
	0x97, 0x37, 0x12, 0x66, 0x8e, 0xf2, 0x38, 0xa0, 0x22, 0x0b, 0x13, 0x91, 0x88, 0xd0, 0x32, 0xe2,
	0xfc, 0xd0, 0x56, 0xb6, 0xb0, 0x5f, 0x25, 0x7b, 0x79, 0x25, 0x11, 0x22, 0x49, 0xe1, 0xb7, 0xcb,
	0xb0, 0x0c, 0xb4, 0x21, 0x99, 0xac, 0x0c, 0xed, 0xcb, 0x5d, 0x1f, 0x2a, 0x32, 0x50, 0x80, 0x9a,
	0x88, 0xf1, 0x43, 0xa1, 0x32, 0x62, 0x98, 0xe0, 0xa3, 0x0d, 0x2f, 0xdf, 0x1d, 0xcf, 0x08, 0x39,
	0x70, 0xd4, 0xe2, 0x83, 0x71, 0x71, 0xf4, 0xbe, 0x51, 0xaf, 0x5d, 0x99, 0x96, 0xc6, 0x4d, 0xa3,
	0xf0, 0x7b, 0xe3, 0x52, 0x8f, 0xa4, 0xec, 0x80, 0x18, 0xa8, 0xd4, 0xd5, 0x4b, 0x2a, 0x83, 0xf7,
	0xd1, 0xd8, 0xf9, 0x6b, 0x3f, 0x5a, 0xee, 0xc2, 0xcb, 0x54, 0xc4, 0x24, 0xdd, 0x93, 0x40, 0xf7,
	0x8f, 0x25, 0x78, 0x3b, 0xee, 0x22, 0x15, 0x99, 0x24, 0xfc, 0x38, 0x3a, 0x00, 0x43, 0x58, 0xaa,
	0x7d, 0xb4, 0x8a, 0x1e, 0xcd, 0xb5, 0xef, 0x07, 0xe3, 0xa3, 0x2f, 0x9b, 0x0d, 0x9e, 0x97, 0xe6,
	0x5d, 0x30, 0xa4, 0xbb, 0x50, 0x25, 0x5f, 0x94, 0x41, 0x6f, 0xd3, 0x9d, 0xcf, 0x35, 0xa8, 0x21,
	0xe8, 0x1f, 0x0b, 0xc2, 0xcd, 0xa0, 0xd7, 0x1a, 0x94, 0xa5, 0xcc, 0x0d, 0x32, 0x35, 0x62, 0xc7,
	0x5d, 0x24, 0x94, 0x8a, 0x9c, 0x9b, 0x21, 0xe5, 0xdf, 0x49, 0xed, 0x6c, 0x96, 0xe6, 0xb2, 0x9d,
	0x2a, 0x39, 0xc2, 0x8a, 0x59, 0x9a, 0x32, 0x9e, 0x0c, 0x59, 0x53, 0x93, 0x58, 0x5b, 0xa5, 0xb9,
	0x64, 0x55, 0xc9, 0x9a, 0xb5, 0xed, 0xfe, 0xaf, 0x45, 0xae, 0x28, 0x44, 0x32, 0x8f, 0x53, 0x46,
	0xfd, 0x96, 0x25, 0xad, 0x35, 0x93, 0xf6, 0xac, 0xf5, 0x95, 0x75, 0x76, 0x9c, 0xee, 0xbc, 0x1e,
	0xa9, 0xbd, 0x37, 0xee, 0x8d, 0x0a, 0xc5, 0xb8, 0x01, 0xc5, 0x49, 0x1a, 0x69, 0x05, 0xfe, 0x7f,
	0x16, 0xf8, 0x70, 0x12, 0x70, 0xbb, 0xf2, 0xef, 0x29, 0xe8, 0x38, 0xdd, 0xeb, 0xfa, 0xf2, 0x4f,
	0x0f, 0xdc, 0x3b, 0x7f, 0xa0, 0x29, 0x19, 0xdc, 0xc3, 0x9f, 0xb1, 0xf8, 0xf5, 0x2b, 0xe1, 0xcb,
	0x48, 0xc7, 0xe9, 0xde, 0xd2, 0x4d, 0x82, 0x17, 0xbb, 0xb7, 0xeb, 0x61, 0xa4, 0x84, 0x47, 0x46,
	0x11, 0xae, 0xd9, 0xe0, 0x9d, 0xf9, 0xb3, 0xf6, 0x94, 0xc7, 0x13, 0xa7, 0x92, 0x12, 0xbe, 0x3f,
	0x4c, 0x74, 0x9c, 0xee, 0x4d, 0xdd, 0xf0, 0xdf, 0x7b, 0xe6, 0xba, 0xd5, 0x19, 0x99, 0x96, 0xbe,
	0x6b, 0xb9, 0x2b, 0x93, 0xb8, 0xbb, 0x5a, 0x76, 0x9c, 0xee, 0xac, 0xae, 0x8b, 0xc6, 0x39, 0x6b,
	0xe1, 0xcf, 0xfd, 0xc5, 0x9c, 0xb5, 0x68, 0x98, 0xb3, 0x16, 0x5b, 0x4b, 0xc3, 0xd7, 0x40, 0x8f,
	0x04, 0xa3, 0xe0, 0xcd, 0x7c, 0x3d, 0x41, 0xad, 0xfe, 0x09, 0x9a, 0xde, 0x99, 0x9a, 0x99, 0xbe,
	0xd6, 0xda, 0xfa, 0x8c, 0xfa, 0x67, 0xd8, 0x39, 0x3d, 0xc3, 0xce, 0xc5, 0x19, 0x46, 0x1f, 0x0b,
	0x8c, 0xbe, 0x14, 0x18, 0x7d, 0x2b, 0x30, 0xea, 0x17, 0x18, 0x9d, 0x16, 0x18, 0x7d, 0x2f, 0x30,
	0xfa, 0x59, 0x60, 0xe7, 0xa2, 0xc0, 0xe8, 0xd3, 0x39, 0x76, 0xfa, 0xe7, 0xd8, 0x39, 0x3d, 0xc7,
	0xce, 0x5b, 0x92, 0x08, 0xf9, 0x2e, 0x09, 0x7a, 0x22, 0x35, 0xa0, 0x14, 0x09, 0x72, 0x1d, 0xda,
	0x8f, 0xc1, 0xba, 0xd9, 0x90, 0x4a, 0xf4, 0xd8, 0x01, 0xa8, 0x8d, 0x5a, 0x0e, 0x65, 0x9c, 0x88,
	0x10, 0x3e, 0x98, 0x7a, 0xa5, 0x5c, 0x61, 0xc1, 0xc6, 0x2d, 0xbb, 0x0d, 0x9e, 0xfc, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0x24, 0x4b, 0xda, 0x96, 0x05, 0x00, 0x00,
}

func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
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
	if !this.CompanyDetails.Equal(that1.CompanyDetails) {
		return false
	}
	if !this.UserDetails.Equal(that1.UserDetails) {
		return false
	}
	if !this.AccountDetails.Equal(that1.AccountDetails) {
		return false
	}
	if !this.BillingDetails.Equal(that1.BillingDetails) {
		return false
	}
	if that1.SourceChoice == nil {
		if this.SourceChoice != nil {
			return false
		}
	} else if this.SourceChoice == nil {
		return false
	} else if !this.SourceChoice.Equal(that1.SourceChoice) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SourcePublic) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SourcePublic)
	if !ok {
		that2, ok := that.(GlobalSpecType_SourcePublic)
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
	if !this.SourcePublic.Equal(that1.SourcePublic) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SourceInternalSre) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SourceInternalSre)
	if !ok {
		that2, ok := that.(GlobalSpecType_SourceInternalSre)
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
	if !this.SourceInternalSre.Equal(that1.SourceInternalSre) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SourceInternalScaling) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SourceInternalScaling)
	if !ok {
		that2, ok := that.(GlobalSpecType_SourceInternalScaling)
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
	if !this.SourceInternalScaling.Equal(that1.SourceInternalScaling) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SourcePlanTransition) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SourcePlanTransition)
	if !ok {
		that2, ok := that.(GlobalSpecType_SourcePlanTransition)
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
	if !this.SourcePlanTransition.Equal(that1.SourcePlanTransition) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SourceMsp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SourceMsp)
	if !ok {
		that2, ok := that.(GlobalSpecType_SourceMsp)
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
	if !this.SourceMsp.Equal(that1.SourceMsp) {
		return false
	}
	return true
}
func (this *GlobalSpecType_SourceInternalSso) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType_SourceInternalSso)
	if !ok {
		that2, ok := that.(GlobalSpecType_SourceInternalSso)
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
	if !this.SourceInternalSso.Equal(that1.SourceInternalSso) {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&v2_signup.GlobalSpecType{")
	if this.CompanyDetails != nil {
		s = append(s, "CompanyDetails: "+fmt.Sprintf("%#v", this.CompanyDetails)+",\n")
	}
	if this.UserDetails != nil {
		s = append(s, "UserDetails: "+fmt.Sprintf("%#v", this.UserDetails)+",\n")
	}
	if this.AccountDetails != nil {
		s = append(s, "AccountDetails: "+fmt.Sprintf("%#v", this.AccountDetails)+",\n")
	}
	if this.BillingDetails != nil {
		s = append(s, "BillingDetails: "+fmt.Sprintf("%#v", this.BillingDetails)+",\n")
	}
	if this.SourceChoice != nil {
		s = append(s, "SourceChoice: "+fmt.Sprintf("%#v", this.SourceChoice)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GlobalSpecType_SourcePublic) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&v2_signup.GlobalSpecType_SourcePublic{` +
		`SourcePublic:` + fmt.Sprintf("%#v", this.SourcePublic) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_SourceInternalSre) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&v2_signup.GlobalSpecType_SourceInternalSre{` +
		`SourceInternalSre:` + fmt.Sprintf("%#v", this.SourceInternalSre) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_SourceInternalScaling) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&v2_signup.GlobalSpecType_SourceInternalScaling{` +
		`SourceInternalScaling:` + fmt.Sprintf("%#v", this.SourceInternalScaling) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_SourcePlanTransition) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&v2_signup.GlobalSpecType_SourcePlanTransition{` +
		`SourcePlanTransition:` + fmt.Sprintf("%#v", this.SourcePlanTransition) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_SourceMsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&v2_signup.GlobalSpecType_SourceMsp{` +
		`SourceMsp:` + fmt.Sprintf("%#v", this.SourceMsp) + `}`}, ", ")
	return s
}
func (this *GlobalSpecType_SourceInternalSso) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&v2_signup.GlobalSpecType_SourceInternalSso{` +
		`SourceInternalSso:` + fmt.Sprintf("%#v", this.SourceInternalSso) + `}`}, ", ")
	return s
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SourceChoice != nil {
		{
			size := m.SourceChoice.Size()
			i -= size
			if _, err := m.SourceChoice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.BillingDetails != nil {
		{
			size, err := m.BillingDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.AccountDetails != nil {
		{
			size, err := m.AccountDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.UserDetails != nil {
		{
			size, err := m.UserDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CompanyDetails != nil {
		{
			size, err := m.CompanyDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GlobalSpecType_SourcePublic) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SourcePublic) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourcePublic != nil {
		{
			size, err := m.SourcePublic.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_SourceInternalSre) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SourceInternalSre) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourceInternalSre != nil {
		{
			size, err := m.SourceInternalSre.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_SourceInternalScaling) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SourceInternalScaling) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourceInternalScaling != nil {
		{
			size, err := m.SourceInternalScaling.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_SourcePlanTransition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SourcePlanTransition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourcePlanTransition != nil {
		{
			size, err := m.SourcePlanTransition.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_SourceMsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SourceMsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourceMsp != nil {
		{
			size, err := m.SourceMsp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *GlobalSpecType_SourceInternalSso) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType_SourceInternalSso) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourceInternalSso != nil {
		{
			size, err := m.SourceInternalSso.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CompanyDetails != nil {
		l = m.CompanyDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.UserDetails != nil {
		l = m.UserDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.AccountDetails != nil {
		l = m.AccountDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.BillingDetails != nil {
		l = m.BillingDetails.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SourceChoice != nil {
		n += m.SourceChoice.Size()
	}
	return n
}

func (m *GlobalSpecType_SourcePublic) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourcePublic != nil {
		l = m.SourcePublic.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_SourceInternalSre) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceInternalSre != nil {
		l = m.SourceInternalSre.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_SourceInternalScaling) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceInternalScaling != nil {
		l = m.SourceInternalScaling.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_SourcePlanTransition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourcePlanTransition != nil {
		l = m.SourcePlanTransition.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_SourceMsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceMsp != nil {
		l = m.SourceMsp.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *GlobalSpecType_SourceInternalSso) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceInternalSso != nil {
		l = m.SourceInternalSso.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`CompanyDetails:` + strings.Replace(fmt.Sprintf("%v", this.CompanyDetails), "CompanyMeta", "signup.CompanyMeta", 1) + `,`,
		`UserDetails:` + strings.Replace(fmt.Sprintf("%v", this.UserDetails), "UserMeta", "signup.UserMeta", 1) + `,`,
		`AccountDetails:` + strings.Replace(fmt.Sprintf("%v", this.AccountDetails), "AccountMeta", "signup.AccountMeta", 1) + `,`,
		`BillingDetails:` + strings.Replace(fmt.Sprintf("%v", this.BillingDetails), "BillingMeta", "signup.BillingMeta", 1) + `,`,
		`SourceChoice:` + fmt.Sprintf("%v", this.SourceChoice) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SourcePublic) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SourcePublic{`,
		`SourcePublic:` + strings.Replace(fmt.Sprintf("%v", this.SourcePublic), "SourcePublic", "signup.SourcePublic", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SourceInternalSre) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SourceInternalSre{`,
		`SourceInternalSre:` + strings.Replace(fmt.Sprintf("%v", this.SourceInternalSre), "SourceInternalSre", "signup.SourceInternalSre", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SourceInternalScaling) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SourceInternalScaling{`,
		`SourceInternalScaling:` + strings.Replace(fmt.Sprintf("%v", this.SourceInternalScaling), "SourceInternalScaling", "signup.SourceInternalScaling", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SourcePlanTransition) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SourcePlanTransition{`,
		`SourcePlanTransition:` + strings.Replace(fmt.Sprintf("%v", this.SourcePlanTransition), "SourcePlanTransition", "signup.SourcePlanTransition", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SourceMsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SourceMsp{`,
		`SourceMsp:` + strings.Replace(fmt.Sprintf("%v", this.SourceMsp), "SourceMsp", "signup.SourceMsp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GlobalSpecType_SourceInternalSso) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType_SourceInternalSso{`,
		`SourceInternalSso:` + strings.Replace(fmt.Sprintf("%v", this.SourceInternalSso), "SourceInternalSso", "signup.SourceInternalSso", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CompanyDetails == nil {
				m.CompanyDetails = &signup.CompanyMeta{}
			}
			if err := m.CompanyDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserDetails == nil {
				m.UserDetails = &signup.UserMeta{}
			}
			if err := m.UserDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccountDetails == nil {
				m.AccountDetails = &signup.AccountMeta{}
			}
			if err := m.AccountDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillingDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BillingDetails == nil {
				m.BillingDetails = &signup.BillingMeta{}
			}
			if err := m.BillingDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePublic", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &signup.SourcePublic{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SourceChoice = &GlobalSpecType_SourcePublic{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceInternalSre", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &signup.SourceInternalSre{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SourceChoice = &GlobalSpecType_SourceInternalSre{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceInternalScaling", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &signup.SourceInternalScaling{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SourceChoice = &GlobalSpecType_SourceInternalScaling{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePlanTransition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &signup.SourcePlanTransition{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SourceChoice = &GlobalSpecType_SourcePlanTransition{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceMsp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &signup.SourceMsp{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SourceChoice = &GlobalSpecType_SourceMsp{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceInternalSso", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &signup.SourceInternalSso{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SourceChoice = &GlobalSpecType_SourceInternalSso{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
