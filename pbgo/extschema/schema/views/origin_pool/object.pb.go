// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/origin_pool/object.proto

/*
	Package origin_pool is a generated protocol buffer package.

	It is generated from these files:
		ves.io/schema/views/origin_pool/object.proto
		ves.io/schema/views/origin_pool/public_crudapi.proto
		ves.io/schema/views/origin_pool/types.proto

	It has these top-level messages:
		Object
		SpecType
		CreateRequest
		CreateResponse
		ReplaceRequest
		ReplaceResponse
		GetRequest
		GetResponse
		ListRequest
		ListResponseItem
		ListResponse
		DeleteRequest
		OriginServerPublicIP
		OriginServerPublicName
		OriginServerPrivateIP
		OriginServerPrivateName
		OriginServerK8SService
		OriginServerConsulService
		OriginServerCustomEndpoint
		OriginServerType
		UpstreamTlsValidationContext
		TlsCertificatesType
		UpstreamTlsParameters
		OriginPoolDefaultSubset
		OriginPoolSubsets
		OriginPoolAdvancedOptions
		GlobalSpecType
		CreateSpecType
		ReplaceSpecType
		GetSpecType
*/
package origin_pool

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema4 "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/views"

	strings "strings"

	reflect "reflect"

	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Object
//
// x-displayName: "Object"
// Origin pool view object
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *ves_io_schema4.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *ves_io_schema4.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the desired behavior of the tenant
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{0} }

func (m *Object) GetMetadata() *ves_io_schema4.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *ves_io_schema4.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

type SpecType struct {
	GcSpec *GlobalSpecType `protobuf:"bytes,1,opt,name=gc_spec,json=gcSpec" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()                    { *m = SpecType{} }
func (*SpecType) ProtoMessage()               {}
func (*SpecType) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{1} }

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.views.origin_pool.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.views.origin_pool.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.views.origin_pool.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.views.origin_pool.SpecType")
}
func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
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
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&origin_pool.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&origin_pool.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SystemMetadata != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.SystemMetadata.Size()))
		n2, err := m.SystemMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Spec != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Spec.Size()))
		n3, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.GcSpec.Size()))
		n4, err := m.GcSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedObject(r randyObject, easy bool) *Object {
	this := &Object{}
	if r.Intn(10) != 0 {
		this.Metadata = ves_io_schema4.NewPopulatedObjectMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		this.SystemMetadata = ves_io_schema4.NewPopulatedSystemObjectMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Spec = NewPopulatedSpecType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSpecType(r randyObject, easy bool) *SpecType {
	this := &SpecType{}
	if r.Intn(10) != 0 {
		this.GcSpec = NewPopulatedGlobalSpecType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyObject interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneObject(r randyObject) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringObject(r randyObject) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneObject(r)
	}
	return string(tmps)
}
func randUnrecognizedObject(r randyObject, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldObject(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldObject(dAtA []byte, r randyObject, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateObject(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateObject(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateObject(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Object) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func sovObject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "ves_io_schema4.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "ves_io_schema4.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(fmt.Sprintf("%v", this.Spec), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &ves_io_schema4.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowObject
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
				next, err := skipObject(dAtA[start:])
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
	ErrInvalidLengthObject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/views/origin_pool/object.proto", fileDescriptorObject) }
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/origin_pool/object.proto", fileDescriptorObject)
}

var fileDescriptorObject = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x6e, 0x13, 0x41,
	0x10, 0xc6, 0x3d, 0x01, 0x99, 0xe8, 0x40, 0x80, 0x5c, 0x20, 0x13, 0x60, 0x13, 0x99, 0x06, 0x04,
	0xde, 0x95, 0xa0, 0x02, 0x89, 0x86, 0x06, 0x0a, 0x22, 0x24, 0x12, 0x1a, 0x1a, 0x6b, 0xef, 0x3c,
	0xd9, 0x2c, 0xdc, 0x79, 0x56, 0xb7, 0xeb, 0x23, 0xee, 0x78, 0x04, 0x3a, 0x5e, 0x81, 0x47, 0xa0,
	0x4c, 0xa8, 0x10, 0x55, 0xca, 0x94, 0x78, 0xd3, 0xa4, 0x4c, 0x49, 0x89, 0xbc, 0xe7, 0xb3, 0x7c,
	0x07, 0xc8, 0xdd, 0xac, 0xbe, 0xdf, 0x7c, 0xf3, 0x6f, 0xa3, 0x87, 0x05, 0x5a, 0xae, 0x49, 0xd8,
	0x64, 0x1f, 0x33, 0x29, 0x0a, 0x8d, 0x1f, 0xad, 0xa0, 0x5c, 0x2b, 0x3d, 0x1a, 0x18, 0xa2, 0x54,
	0x50, 0xfc, 0x1e, 0x13, 0xc7, 0x4d, 0x4e, 0x8e, 0x3a, 0x9b, 0x25, 0xcd, 0x4b, 0x9a, 0x07, 0x9a,
	0x2f, 0xd1, 0x1b, 0x7d, 0xa5, 0xdd, 0xfe, 0x38, 0xe6, 0x09, 0x65, 0x42, 0x91, 0x22, 0x11, 0xf2,
	0xe2, 0xf1, 0x5e, 0x78, 0x85, 0x47, 0x88, 0x4a, 0xbf, 0x8d, 0x5b, 0xf5, 0xea, 0x64, 0x9c, 0xa6,
	0x91, 0x9d, 0x8b, 0x37, 0xeb, 0xa2, 0x9b, 0x18, 0xac, 0xa4, 0xdb, 0x8d, 0xae, 0x65, 0xaa, 0x87,
	0xd2, 0xe1, 0x5c, 0xed, 0x35, 0x54, 0xb4, 0x38, 0x2a, 0x1a, 0xe6, 0x5b, 0x7f, 0xcf, 0x3d, 0xa8,
	0x13, 0x0f, 0x56, 0x6d, 0x66, 0xb9, 0xa1, 0xcd, 0x7f, 0xc1, 0x4b, 0x40, 0xef, 0x0c, 0xa2, 0xf6,
	0xeb, 0xb0, 0xca, 0xce, 0x93, 0x68, 0x3d, 0x43, 0x27, 0x87, 0xd2, 0xc9, 0x2e, 0x6c, 0xc1, 0xbd,
	0xcb, 0x8f, 0xee, 0xf0, 0xfa, 0x5e, 0x4b, 0x70, 0x1b, 0x9d, 0xdc, 0x9d, 0x18, 0x7c, 0xb3, 0xc0,
	0x3b, 0xaf, 0xa2, 0x6b, 0x76, 0x62, 0x1d, 0x66, 0x83, 0x85, 0xc3, 0x5a, 0x70, 0xb8, 0xdb, 0x70,
	0xd8, 0x09, 0x54, 0xc3, 0xe7, 0x6a, 0x99, 0xbb, 0x5d, 0xb9, 0x3d, 0x8b, 0x2e, 0x5a, 0x83, 0x49,
	0xf7, 0x42, 0xb0, 0xb8, 0xcf, 0x57, 0x1c, 0x97, 0xef, 0x18, 0x4c, 0x82, 0x51, 0x48, 0x7b, 0x7a,
	0xe3, 0xfb, 0x51, 0x77, 0xed, 0x3a, 0xfc, 0x3c, 0xea, 0x5e, 0x29, 0xd0, 0xf6, 0x35, 0xf5, 0x4d,
	0x4e, 0x07, 0x93, 0xde, 0x6e, 0xb4, 0x5e, 0x91, 0x9d, 0x97, 0xd1, 0x25, 0x95, 0x0c, 0x42, 0x95,
	0x72, 0x54, 0xb1, 0xb2, 0xca, 0x8b, 0x94, 0x62, 0x99, 0x2e, 0x6a, 0xb5, 0x55, 0x32, 0x8b, 0x9f,
	0x7f, 0x81, 0xe3, 0x29, 0x6b, 0x9d, 0x4c, 0x59, 0xeb, 0x7c, 0xca, 0xe0, 0xf7, 0x94, 0xc1, 0x27,
	0xcf, 0xe0, 0xab, 0x67, 0xf0, 0xcd, 0x33, 0x38, 0xf4, 0x0c, 0x7e, 0x78, 0x06, 0xc7, 0x9e, 0xc1,
	0x89, 0x67, 0xf0, 0xcb, 0x33, 0x38, 0xf3, 0xac, 0x75, 0xee, 0x19, 0x7c, 0x3e, 0x65, 0xad, 0xc3,
	0x53, 0x06, 0xef, 0xde, 0x2a, 0x32, 0x1f, 0x14, 0x2f, 0x28, 0x75, 0x98, 0xe7, 0x92, 0x8f, 0xad,
	0x08, 0xc1, 0x1e, 0xe5, 0xd9, 0xac, 0xf7, 0x42, 0x0f, 0x31, 0xef, 0x57, 0xb2, 0x30, 0xb1, 0x22,
	0x81, 0x07, 0x6e, 0x7e, 0xd3, 0xff, 0xfd, 0x83, 0xb8, 0x1d, 0x2e, 0xfc, 0xf8, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x05, 0xf9, 0x5e, 0x23, 0x4b, 0x03, 0x00, 0x00,
}
