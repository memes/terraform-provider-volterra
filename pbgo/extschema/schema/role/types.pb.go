// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/role/types.proto

package role

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/gogo/protobuf/types"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema4 "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Role
//
// x-displayName: "Role"
// role internally uses name (like admin) and namespace (shared) and tenant (ves-io) for global roles
// roles can also be tenant specific living under their respective tenant
type GlobalSpecType struct {
	// RBAC Policy
	//
	// x-displayName: "RBAC Policy"
	// x-required
	// reference to rabc_policy object which has rules defined for access on list of API Groups.
	RbacPolicy []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=rbac_policy,json=rbacPolicy" json:"rbac_policy,omitempty"`
}

func (m *GlobalSpecType) Reset()                    { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage()               {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *GlobalSpecType) GetRbacPolicy() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.RbacPolicy
	}
	return nil
}

// Role
//
// x-displayName: "Role"
// Creates a role
type CreateSpecType struct {
}

func (m *CreateSpecType) Reset()                    { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage()               {}
func (*CreateSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

// Role
//
// x-displayName: "Role"
// Amends a role
type ReplaceSpecType struct {
}

func (m *ReplaceSpecType) Reset()                    { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage()               {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

// Role
//
// x-displayName: "Role"
// Reads a role
type GetSpecType struct {
	// RBAC Policy
	//
	// x-displayName: "RBAC Policy"
	// x-required
	// reference to rabc_policy object which has rules defined for access on list of API Groups.
	RbacPolicy []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=rbac_policy,json=rbacPolicy" json:"rbac_policy,omitempty"`
}

func (m *GetSpecType) Reset()                    { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage()               {}
func (*GetSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *GetSpecType) GetRbacPolicy() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.RbacPolicy
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.role.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.role.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.role.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.role.GetSpecType")
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
	if len(this.RbacPolicy) != len(that1.RbacPolicy) {
		return false
	}
	for i := range this.RbacPolicy {
		if !this.RbacPolicy[i].Equal(that1.RbacPolicy[i]) {
			return false
		}
	}
	return true
}
func (this *CreateSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSpecType)
	if !ok {
		that2, ok := that.(CreateSpecType)
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
func (this *ReplaceSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceSpecType)
	if !ok {
		that2, ok := that.(ReplaceSpecType)
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
func (this *GetSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType)
	if !ok {
		that2, ok := that.(GetSpecType)
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
	if len(this.RbacPolicy) != len(that1.RbacPolicy) {
		return false
	}
	for i := range this.RbacPolicy {
		if !this.RbacPolicy[i].Equal(that1.RbacPolicy[i]) {
			return false
		}
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&role.GlobalSpecType{")
	if this.RbacPolicy != nil {
		s = append(s, "RbacPolicy: "+fmt.Sprintf("%#v", this.RbacPolicy)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&role.CreateSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&role.ReplaceSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&role.GetSpecType{")
	if this.RbacPolicy != nil {
		s = append(s, "RbacPolicy: "+fmt.Sprintf("%#v", this.RbacPolicy)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
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
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RbacPolicy) > 0 {
		for _, msg := range m.RbacPolicy {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CreateSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ReplaceSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RbacPolicy) > 0 {
		for _, msg := range m.RbacPolicy {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGlobalSpecType(r randyTypes, easy bool) *GlobalSpecType {
	this := &GlobalSpecType{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.RbacPolicy = make([]*ves_io_schema4.ObjectRefType, v1)
		for i := 0; i < v1; i++ {
			this.RbacPolicy[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCreateSpecType(r randyTypes, easy bool) *CreateSpecType {
	this := &CreateSpecType{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedReplaceSpecType(r randyTypes, easy bool) *ReplaceSpecType {
	this := &ReplaceSpecType{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGetSpecType(r randyTypes, easy bool) *GetSpecType {
	this := &GetSpecType{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.RbacPolicy = make([]*ves_io_schema4.ObjectRefType, v2)
		for i := 0; i < v2; i++ {
			this.RbacPolicy[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypes(r randyTypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypes(r randyTypes) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTypes(r)
	}
	return string(tmps)
}
func randUnrecognizedTypes(r randyTypes, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTypes(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTypes(dAtA []byte, r randyTypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTypes(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GlobalSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.RbacPolicy) > 0 {
		for _, e := range m.RbacPolicy {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *CreateSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetSpecType) Size() (n int) {
	var l int
	_ = l
	if len(m.RbacPolicy) > 0 {
		for _, e := range m.RbacPolicy {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`RbacPolicy:` + strings.Replace(fmt.Sprintf("%v", this.RbacPolicy), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType{`,
		`RbacPolicy:` + strings.Replace(fmt.Sprintf("%v", this.RbacPolicy), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
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
			wire |= (uint64(b) & 0x7F) << shift
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
				return fmt.Errorf("proto: wrong wireType = %d for field RbacPolicy", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RbacPolicy = append(m.RbacPolicy, &ves_io_schema4.ObjectRefType{})
			if err := m.RbacPolicy[len(m.RbacPolicy)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *CreateSpecType) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *ReplaceSpecType) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplaceSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GetSpecType) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RbacPolicy", wireType)
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RbacPolicy = append(m.RbacPolicy, &ves_io_schema4.ObjectRefType{})
			if err := m.RbacPolicy[len(m.RbacPolicy)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
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
				next, err := skipTypes(dAtA[start:])
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
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/role/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0xab, 0x13, 0x41,
	0x10, 0xc6, 0x6f, 0x0d, 0x5a, 0x5c, 0x20, 0x6a, 0xaa, 0x18, 0xc3, 0x20, 0xd1, 0xc2, 0x26, 0xbb,
	0xa0, 0x9d, 0x85, 0x45, 0x2c, 0x62, 0xa7, 0x44, 0x41, 0xb0, 0x91, 0xbd, 0xcb, 0x64, 0x73, 0xba,
	0xc9, 0x2c, 0x7b, 0x9b, 0xc3, 0x2b, 0x04, 0x2b, 0x6b, 0xf1, 0xaf, 0x10, 0xff, 0x02, 0xb1, 0x4a,
	0x29, 0x56, 0x29, 0x53, 0x7a, 0x6b, 0xa3, 0x5d, 0x4a, 0xcb, 0x47, 0xee, 0x12, 0x5e, 0xee, 0xf1,
	0x78, 0xdd, 0xeb, 0x66, 0xe6, 0xf7, 0x7d, 0xf7, 0xcd, 0x31, 0x1b, 0x42, 0x86, 0x29, 0x4f, 0x48,
	0xa4, 0xf1, 0x0c, 0xe7, 0x52, 0x58, 0xd2, 0x28, 0x5c, 0x6e, 0x30, 0xe5, 0xc6, 0x92, 0xa3, 0x76,
	0xbb, 0xe2, 0xbc, 0xe2, 0x7c, 0xc7, 0xbb, 0x03, 0x95, 0xb8, 0xd9, 0x32, 0xe2, 0x31, 0xcd, 0x85,
	0x22, 0x45, 0xa2, 0x94, 0x46, 0xcb, 0x69, 0xd9, 0x95, 0x4d, 0x59, 0x55, 0x9f, 0xe8, 0xde, 0x52,
	0x44, 0x4a, 0xe3, 0xa9, 0x4a, 0x2e, 0xf2, 0x3d, 0xba, 0x5d, 0x4f, 0x27, 0xe3, 0x12, 0x5a, 0xa4,
	0x07, 0x5f, 0x1d, 0x1e, 0x6d, 0xd5, 0xed, 0xd5, 0x51, 0x26, 0x75, 0x32, 0x91, 0x0e, 0x2b, 0xda,
	0x4f, 0xc2, 0xd6, 0x48, 0x53, 0x24, 0xf5, 0x0b, 0x83, 0xf1, 0xcb, 0xdc, 0x60, 0xfb, 0x55, 0xd8,
	0xb4, 0x91, 0x8c, 0xdf, 0x18, 0xd2, 0x49, 0x9c, 0x77, 0xd8, 0x9d, 0xc6, 0xfd, 0xe6, 0x83, 0x1e,
	0xaf, 0xff, 0xdb, 0xb3, 0xe8, 0x2d, 0xc6, 0x6e, 0x8c, 0xd3, 0x9d, 0x65, 0xd8, 0xf9, 0xf6, 0xe1,
	0xd8, 0xf3, 0xe3, 0xdf, 0xaa, 0x71, 0xf5, 0x0b, 0xbb, 0x72, 0x83, 0x8d, 0xc3, 0xdd, 0xf8, 0x79,
	0x39, 0xed, 0xdf, 0x0d, 0x5b, 0x4f, 0x2c, 0x4a, 0x87, 0x87, 0xa8, 0x47, 0x37, 0x7f, 0x3d, 0x3e,
	0x93, 0xde, 0xbf, 0x17, 0x5e, 0x1f, 0xa3, 0xd1, 0x32, 0xbe, 0x50, 0x95, 0x87, 0xcd, 0x11, 0xba,
	0x4b, 0x5f, 0xf9, 0x9c, 0xe8, 0xe1, 0x27, 0xb6, 0x2e, 0x20, 0xd8, 0x14, 0x10, 0x6c, 0x0b, 0x60,
	0xff, 0x0b, 0x60, 0x1f, 0x3d, 0xb0, 0xaf, 0x1e, 0xd8, 0x77, 0x0f, 0x6c, 0xe5, 0x81, 0xfd, 0xf4,
	0xc0, 0xd6, 0x1e, 0xd8, 0xc6, 0x03, 0xfb, 0xed, 0x81, 0xfd, 0xf5, 0x10, 0x6c, 0x3d, 0xb0, 0xcf,
	0x7f, 0x20, 0x78, 0xfd, 0x54, 0x91, 0x79, 0xa7, 0x78, 0x46, 0xda, 0xa1, 0xb5, 0x92, 0x2f, 0x53,
	0x51, 0x16, 0x53, 0xb2, 0xf3, 0x81, 0xb1, 0x94, 0x25, 0x13, 0xb4, 0x83, 0x03, 0x16, 0x26, 0x52,
	0x24, 0xf0, 0xbd, 0xdb, 0x1f, 0xef, 0xe8, 0xe5, 0x45, 0xd7, 0xca, 0x03, 0x3e, 0x3c, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0x00, 0x77, 0x92, 0x96, 0x02, 0x00, 0x00,
}
