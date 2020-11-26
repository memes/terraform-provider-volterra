// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/terraform-provider-volterra/dummy.proto

/*
	Package pbgo is a generated protocol buffer package.

	It is generated from these files:
		ves.io/terraform-provider-volterra/dummy.proto

	It has these top-level messages:
		Conf
*/
package pbgo

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	google_protobuf1 "github.com/gogo/protobuf/types"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema5 "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Service configuration for ExampleSvc
type Conf struct {
	// Objects for service to start with, maybe updated dynamically
	Bootstrap []*google_protobuf1.Any `protobuf:"bytes,1,rep,name=Bootstrap" json:"Bootstrap,omitempty"`
	// Objects for service to always have, may NOT be updated dynamically
	Overrides           []*google_protobuf1.Any             `protobuf:"bytes,2,rep,name=Overrides" json:"Overrides,omitempty"`
	GrpcPort            int32                               `protobuf:"varint,4,opt,name=GrpcPort,proto3" json:"GrpcPort,omitempty"`
	RestPort            int32                               `protobuf:"varint,5,opt,name=RestPort,proto3" json:"RestPort,omitempty"`
	GrpcTLSPort         int32                               `protobuf:"varint,12,opt,name=GrpcTLSPort,proto3" json:"GrpcTLSPort,omitempty"`
	RestTLSPort         int32                               `protobuf:"varint,13,opt,name=RestTLSPort,proto3" json:"RestTLSPort,omitempty"`
	Tls                 *ves_io_schema5.DaemonTLSParamsType `protobuf:"bytes,14,opt,name=tls" json:"tls,omitempty"`
	EmbedEtcdDataDir    string                              `protobuf:"bytes,6,opt,name=EmbedEtcdDataDir,proto3" json:"EmbedEtcdDataDir,omitempty"`
	EmbedEtcdClientPort int32                               `protobuf:"varint,7,opt,name=EmbedEtcdClientPort,proto3" json:"EmbedEtcdClientPort,omitempty"`
	EmbedEtcdPeerPort   int32                               `protobuf:"varint,8,opt,name=EmbedEtcdPeerPort,proto3" json:"EmbedEtcdPeerPort,omitempty"`
	EtcdServerURLs      []string                            `protobuf:"bytes,10,rep,name=EtcdServerURLs" json:"EtcdServerURLs,omitempty"`
	EtcdKeyPrefix       string                              `protobuf:"bytes,11,opt,name=EtcdKeyPrefix,proto3" json:"EtcdKeyPrefix,omitempty"`
}

func (m *Conf) Reset()                    { *m = Conf{} }
func (*Conf) ProtoMessage()               {}
func (*Conf) Descriptor() ([]byte, []int) { return fileDescriptorDummy, []int{0} }

func (m *Conf) GetBootstrap() []*google_protobuf1.Any {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Conf) GetOverrides() []*google_protobuf1.Any {
	if m != nil {
		return m.Overrides
	}
	return nil
}

func (m *Conf) GetGrpcPort() int32 {
	if m != nil {
		return m.GrpcPort
	}
	return 0
}

func (m *Conf) GetRestPort() int32 {
	if m != nil {
		return m.RestPort
	}
	return 0
}

func (m *Conf) GetGrpcTLSPort() int32 {
	if m != nil {
		return m.GrpcTLSPort
	}
	return 0
}

func (m *Conf) GetRestTLSPort() int32 {
	if m != nil {
		return m.RestTLSPort
	}
	return 0
}

func (m *Conf) GetTls() *ves_io_schema5.DaemonTLSParamsType {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *Conf) GetEmbedEtcdDataDir() string {
	if m != nil {
		return m.EmbedEtcdDataDir
	}
	return ""
}

func (m *Conf) GetEmbedEtcdClientPort() int32 {
	if m != nil {
		return m.EmbedEtcdClientPort
	}
	return 0
}

func (m *Conf) GetEmbedEtcdPeerPort() int32 {
	if m != nil {
		return m.EmbedEtcdPeerPort
	}
	return 0
}

func (m *Conf) GetEtcdServerURLs() []string {
	if m != nil {
		return m.EtcdServerURLs
	}
	return nil
}

func (m *Conf) GetEtcdKeyPrefix() string {
	if m != nil {
		return m.EtcdKeyPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*Conf)(nil), "ves.io.terraform_provider_volterra.Conf")
}
func (this *Conf) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Conf)
	if !ok {
		that2, ok := that.(Conf)
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
	if len(this.Bootstrap) != len(that1.Bootstrap) {
		return false
	}
	for i := range this.Bootstrap {
		if !this.Bootstrap[i].Equal(that1.Bootstrap[i]) {
			return false
		}
	}
	if len(this.Overrides) != len(that1.Overrides) {
		return false
	}
	for i := range this.Overrides {
		if !this.Overrides[i].Equal(that1.Overrides[i]) {
			return false
		}
	}
	if this.GrpcPort != that1.GrpcPort {
		return false
	}
	if this.RestPort != that1.RestPort {
		return false
	}
	if this.GrpcTLSPort != that1.GrpcTLSPort {
		return false
	}
	if this.RestTLSPort != that1.RestTLSPort {
		return false
	}
	if !this.Tls.Equal(that1.Tls) {
		return false
	}
	if this.EmbedEtcdDataDir != that1.EmbedEtcdDataDir {
		return false
	}
	if this.EmbedEtcdClientPort != that1.EmbedEtcdClientPort {
		return false
	}
	if this.EmbedEtcdPeerPort != that1.EmbedEtcdPeerPort {
		return false
	}
	if len(this.EtcdServerURLs) != len(that1.EtcdServerURLs) {
		return false
	}
	for i := range this.EtcdServerURLs {
		if this.EtcdServerURLs[i] != that1.EtcdServerURLs[i] {
			return false
		}
	}
	if this.EtcdKeyPrefix != that1.EtcdKeyPrefix {
		return false
	}
	return true
}
func (this *Conf) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 16)
	s = append(s, "&pbgo.Conf{")
	if this.Bootstrap != nil {
		s = append(s, "Bootstrap: "+fmt.Sprintf("%#v", this.Bootstrap)+",\n")
	}
	if this.Overrides != nil {
		s = append(s, "Overrides: "+fmt.Sprintf("%#v", this.Overrides)+",\n")
	}
	s = append(s, "GrpcPort: "+fmt.Sprintf("%#v", this.GrpcPort)+",\n")
	s = append(s, "RestPort: "+fmt.Sprintf("%#v", this.RestPort)+",\n")
	s = append(s, "GrpcTLSPort: "+fmt.Sprintf("%#v", this.GrpcTLSPort)+",\n")
	s = append(s, "RestTLSPort: "+fmt.Sprintf("%#v", this.RestTLSPort)+",\n")
	if this.Tls != nil {
		s = append(s, "Tls: "+fmt.Sprintf("%#v", this.Tls)+",\n")
	}
	s = append(s, "EmbedEtcdDataDir: "+fmt.Sprintf("%#v", this.EmbedEtcdDataDir)+",\n")
	s = append(s, "EmbedEtcdClientPort: "+fmt.Sprintf("%#v", this.EmbedEtcdClientPort)+",\n")
	s = append(s, "EmbedEtcdPeerPort: "+fmt.Sprintf("%#v", this.EmbedEtcdPeerPort)+",\n")
	s = append(s, "EtcdServerURLs: "+fmt.Sprintf("%#v", this.EtcdServerURLs)+",\n")
	s = append(s, "EtcdKeyPrefix: "+fmt.Sprintf("%#v", this.EtcdKeyPrefix)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDummy(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Conf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Conf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Bootstrap) > 0 {
		for _, msg := range m.Bootstrap {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDummy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Overrides) > 0 {
		for _, msg := range m.Overrides {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDummy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.GrpcPort != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.GrpcPort))
	}
	if m.RestPort != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.RestPort))
	}
	if len(m.EmbedEtcdDataDir) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDummy(dAtA, i, uint64(len(m.EmbedEtcdDataDir)))
		i += copy(dAtA[i:], m.EmbedEtcdDataDir)
	}
	if m.EmbedEtcdClientPort != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.EmbedEtcdClientPort))
	}
	if m.EmbedEtcdPeerPort != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.EmbedEtcdPeerPort))
	}
	if len(m.EtcdServerURLs) > 0 {
		for _, s := range m.EtcdServerURLs {
			dAtA[i] = 0x52
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
	if len(m.EtcdKeyPrefix) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintDummy(dAtA, i, uint64(len(m.EtcdKeyPrefix)))
		i += copy(dAtA[i:], m.EtcdKeyPrefix)
	}
	if m.GrpcTLSPort != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.GrpcTLSPort))
	}
	if m.RestTLSPort != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.RestTLSPort))
	}
	if m.Tls != nil {
		dAtA[i] = 0x72
		i++
		i = encodeVarintDummy(dAtA, i, uint64(m.Tls.Size()))
		n1, err := m.Tls.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintDummy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedConf(r randyDummy, easy bool) *Conf {
	this := &Conf{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Bootstrap = make([]*google_protobuf1.Any, v1)
		for i := 0; i < v1; i++ {
			this.Bootstrap[i] = google_protobuf1.NewPopulatedAny(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Overrides = make([]*google_protobuf1.Any, v2)
		for i := 0; i < v2; i++ {
			this.Overrides[i] = google_protobuf1.NewPopulatedAny(r, easy)
		}
	}
	this.GrpcPort = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.GrpcPort *= -1
	}
	this.RestPort = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.RestPort *= -1
	}
	this.EmbedEtcdDataDir = string(randStringDummy(r))
	this.EmbedEtcdClientPort = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.EmbedEtcdClientPort *= -1
	}
	this.EmbedEtcdPeerPort = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.EmbedEtcdPeerPort *= -1
	}
	v3 := r.Intn(10)
	this.EtcdServerURLs = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.EtcdServerURLs[i] = string(randStringDummy(r))
	}
	this.EtcdKeyPrefix = string(randStringDummy(r))
	this.GrpcTLSPort = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.GrpcTLSPort *= -1
	}
	this.RestTLSPort = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.RestTLSPort *= -1
	}
	if r.Intn(10) != 0 {
		this.Tls = ves_io_schema5.NewPopulatedDaemonTLSParamsType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyDummy interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneDummy(r randyDummy) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringDummy(r randyDummy) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneDummy(r)
	}
	return string(tmps)
}
func randUnrecognizedDummy(r randyDummy, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldDummy(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldDummy(dAtA []byte, r randyDummy, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateDummy(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateDummy(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateDummy(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateDummy(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateDummy(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateDummy(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateDummy(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Conf) Size() (n int) {
	var l int
	_ = l
	if len(m.Bootstrap) > 0 {
		for _, e := range m.Bootstrap {
			l = e.Size()
			n += 1 + l + sovDummy(uint64(l))
		}
	}
	if len(m.Overrides) > 0 {
		for _, e := range m.Overrides {
			l = e.Size()
			n += 1 + l + sovDummy(uint64(l))
		}
	}
	if m.GrpcPort != 0 {
		n += 1 + sovDummy(uint64(m.GrpcPort))
	}
	if m.RestPort != 0 {
		n += 1 + sovDummy(uint64(m.RestPort))
	}
	l = len(m.EmbedEtcdDataDir)
	if l > 0 {
		n += 1 + l + sovDummy(uint64(l))
	}
	if m.EmbedEtcdClientPort != 0 {
		n += 1 + sovDummy(uint64(m.EmbedEtcdClientPort))
	}
	if m.EmbedEtcdPeerPort != 0 {
		n += 1 + sovDummy(uint64(m.EmbedEtcdPeerPort))
	}
	if len(m.EtcdServerURLs) > 0 {
		for _, s := range m.EtcdServerURLs {
			l = len(s)
			n += 1 + l + sovDummy(uint64(l))
		}
	}
	l = len(m.EtcdKeyPrefix)
	if l > 0 {
		n += 1 + l + sovDummy(uint64(l))
	}
	if m.GrpcTLSPort != 0 {
		n += 1 + sovDummy(uint64(m.GrpcTLSPort))
	}
	if m.RestTLSPort != 0 {
		n += 1 + sovDummy(uint64(m.RestTLSPort))
	}
	if m.Tls != nil {
		l = m.Tls.Size()
		n += 1 + l + sovDummy(uint64(l))
	}
	return n
}

func sovDummy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDummy(x uint64) (n int) {
	return sovDummy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Conf) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Conf{`,
		`Bootstrap:` + strings.Replace(fmt.Sprintf("%v", this.Bootstrap), "Any", "google_protobuf1.Any", 1) + `,`,
		`Overrides:` + strings.Replace(fmt.Sprintf("%v", this.Overrides), "Any", "google_protobuf1.Any", 1) + `,`,
		`GrpcPort:` + fmt.Sprintf("%v", this.GrpcPort) + `,`,
		`RestPort:` + fmt.Sprintf("%v", this.RestPort) + `,`,
		`EmbedEtcdDataDir:` + fmt.Sprintf("%v", this.EmbedEtcdDataDir) + `,`,
		`EmbedEtcdClientPort:` + fmt.Sprintf("%v", this.EmbedEtcdClientPort) + `,`,
		`EmbedEtcdPeerPort:` + fmt.Sprintf("%v", this.EmbedEtcdPeerPort) + `,`,
		`EtcdServerURLs:` + fmt.Sprintf("%v", this.EtcdServerURLs) + `,`,
		`EtcdKeyPrefix:` + fmt.Sprintf("%v", this.EtcdKeyPrefix) + `,`,
		`GrpcTLSPort:` + fmt.Sprintf("%v", this.GrpcTLSPort) + `,`,
		`RestTLSPort:` + fmt.Sprintf("%v", this.RestTLSPort) + `,`,
		`Tls:` + strings.Replace(fmt.Sprintf("%v", this.Tls), "DaemonTLSParamsType", "ves_io_schema5.DaemonTLSParamsType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDummy(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Conf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDummy
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
			return fmt.Errorf("proto: Conf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Conf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bootstrap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
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
				return ErrInvalidLengthDummy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bootstrap = append(m.Bootstrap, &google_protobuf1.Any{})
			if err := m.Bootstrap[len(m.Bootstrap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
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
				return ErrInvalidLengthDummy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Overrides = append(m.Overrides, &google_protobuf1.Any{})
			if err := m.Overrides[len(m.Overrides)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcPort", wireType)
			}
			m.GrpcPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GrpcPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestPort", wireType)
			}
			m.RestPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RestPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmbedEtcdDataDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
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
				return ErrInvalidLengthDummy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmbedEtcdDataDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmbedEtcdClientPort", wireType)
			}
			m.EmbedEtcdClientPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EmbedEtcdClientPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmbedEtcdPeerPort", wireType)
			}
			m.EmbedEtcdPeerPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EmbedEtcdPeerPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdServerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
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
				return ErrInvalidLengthDummy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EtcdServerURLs = append(m.EtcdServerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdKeyPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
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
				return ErrInvalidLengthDummy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EtcdKeyPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcTLSPort", wireType)
			}
			m.GrpcTLSPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GrpcTLSPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestTLSPort", wireType)
			}
			m.RestTLSPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RestTLSPort |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tls", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDummy
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
				return ErrInvalidLengthDummy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tls == nil {
				m.Tls = &ves_io_schema5.DaemonTLSParamsType{}
			}
			if err := m.Tls.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDummy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDummy
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
func skipDummy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDummy
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
					return 0, ErrIntOverflowDummy
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
					return 0, ErrIntOverflowDummy
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
				return 0, ErrInvalidLengthDummy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDummy
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
				next, err := skipDummy(dAtA[start:])
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
	ErrInvalidLengthDummy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDummy   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/terraform-provider-volterra/dummy.proto", fileDescriptorDummy)
}

var fileDescriptorDummy = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x24, 0x2d, 0xc9, 0x85, 0x56, 0x60, 0x18, 0x4c, 0x10, 0x57, 0x2b, 0x42, 0x28,
	0x42, 0xe4, 0x82, 0x02, 0x13, 0x1b, 0x6d, 0x2a, 0x06, 0x2a, 0x11, 0xb9, 0x65, 0x61, 0x89, 0x9c,
	0xe4, 0xe2, 0x5a, 0xc4, 0x7e, 0xa7, 0xbb, 0x8b, 0x85, 0xb7, 0x0e, 0x7c, 0x00, 0x3e, 0x02, 0x6c,
	0x7c, 0x04, 0x46, 0x46, 0xc6, 0x8e, 0x1d, 0xc9, 0xb1, 0x30, 0x76, 0xec, 0x88, 0x7c, 0x8e, 0xdd,
	0x84, 0xa2, 0x6e, 0x79, 0xff, 0xdf, 0xef, 0xde, 0x3d, 0x5d, 0x9e, 0x31, 0x4d, 0x98, 0xa4, 0x21,
	0x74, 0x15, 0x13, 0xc2, 0x9f, 0x82, 0x88, 0x3a, 0x5c, 0x40, 0x12, 0x4e, 0x98, 0xe8, 0x24, 0x30,
	0x33, 0x69, 0x77, 0x32, 0x8f, 0xa2, 0x94, 0x72, 0x01, 0x0a, 0xec, 0x56, 0xee, 0xd3, 0xd2, 0x1f,
	0x16, 0xfe, 0xb0, 0xf0, 0x9b, 0x9d, 0x20, 0x54, 0xc7, 0xf3, 0x11, 0x1d, 0x43, 0xd4, 0x0d, 0x20,
	0x80, 0xae, 0x39, 0x3a, 0x9a, 0x4f, 0x4d, 0x65, 0x0a, 0xf3, 0x2b, 0x6f, 0xd9, 0xbc, 0x1f, 0x00,
	0x04, 0x33, 0x76, 0x69, 0xf9, 0x71, 0x5a, 0xa0, 0xe5, 0x74, 0x72, 0x7c, 0xcc, 0x22, 0xbf, 0xab,
	0x52, 0xce, 0xe4, 0x12, 0x3d, 0x5c, 0x47, 0x3c, 0x51, 0xc3, 0x55, 0xfc, 0x60, 0x1d, 0x03, 0x57,
	0x21, 0xc4, 0x4b, 0xd8, 0xfa, 0x54, 0xc5, 0xd5, 0x3d, 0x88, 0xa7, 0x76, 0x0f, 0xd7, 0x77, 0x01,
	0x94, 0x54, 0xc2, 0xe7, 0x0e, 0x72, 0x2b, 0xed, 0x46, 0xef, 0x1e, 0xcd, 0xc7, 0xa1, 0xc5, 0x38,
	0xf4, 0x55, 0x9c, 0x7a, 0x97, 0x5a, 0x76, 0xe6, 0x6d, 0xc2, 0x84, 0x08, 0x27, 0x4c, 0x3a, 0x37,
	0xae, 0x3b, 0x53, 0x6a, 0x76, 0x13, 0xd7, 0x5e, 0x0b, 0x3e, 0x1e, 0x80, 0x50, 0x4e, 0xd5, 0x45,
	0xed, 0x0d, 0xaf, 0xac, 0x33, 0xe6, 0x31, 0xa9, 0x0c, 0xdb, 0xc8, 0x59, 0x51, 0xdb, 0x4f, 0xf0,
	0xed, 0xfd, 0x68, 0xc4, 0x26, 0xfb, 0x6a, 0x3c, 0xe9, 0xfb, 0xca, 0xef, 0x87, 0xc2, 0xd9, 0x74,
	0x51, 0xbb, 0xee, 0x5d, 0xc9, 0xed, 0x67, 0xf8, 0x6e, 0x99, 0xed, 0xcd, 0x42, 0x16, 0xe7, 0x2d,
	0x6f, 0x9a, 0x96, 0xff, 0x43, 0xf6, 0x53, 0x7c, 0xa7, 0x8c, 0x07, 0x8c, 0x09, 0xe3, 0xd7, 0x8c,
	0x7f, 0x15, 0xd8, 0x8f, 0xf1, 0x76, 0x56, 0x1f, 0x32, 0x91, 0x30, 0xf1, 0xce, 0x3b, 0x90, 0x0e,
	0x76, 0x2b, 0xed, 0xba, 0xf7, 0x4f, 0x6a, 0x3f, 0xc2, 0x5b, 0x59, 0xf2, 0x86, 0xa5, 0x03, 0xc1,
	0xa6, 0xe1, 0x47, 0xa7, 0x61, 0x06, 0x5e, 0x0f, 0x6d, 0x17, 0x37, 0xb2, 0x17, 0x38, 0x3a, 0x38,
	0x34, 0xb7, 0xde, 0x32, 0xb7, 0xae, 0x46, 0x99, 0x91, 0xbd, 0x43, 0x61, 0x6c, 0xe5, 0xc6, 0x4a,
	0x64, 0xbf, 0xc0, 0x15, 0x35, 0x93, 0xce, 0xb6, 0x8b, 0xda, 0x8d, 0x5e, 0x6b, 0xb9, 0xc9, 0x34,
	0xff, 0xc7, 0x69, 0xdf, 0x67, 0x11, 0xc4, 0x99, 0xea, 0x0b, 0x3f, 0x92, 0x47, 0x29, 0x67, 0x5e,
	0xa6, 0xbf, 0xac, 0x5d, 0x7c, 0xd9, 0xb1, 0x4e, 0xbe, 0xee, 0x58, 0xbb, 0xfc, 0x74, 0x41, 0xac,
	0xb3, 0x05, 0xb1, 0xce, 0x17, 0x04, 0x5d, 0x2c, 0x08, 0x3a, 0xd1, 0x04, 0x7d, 0xd3, 0x04, 0x7d,
	0xd7, 0x04, 0xfd, 0xd0, 0x04, 0xfd, 0xd4, 0x04, 0x9d, 0x6a, 0x82, 0xce, 0x34, 0x41, 0xbf, 0x34,
	0x41, 0x7f, 0x34, 0xb1, 0xce, 0x35, 0x41, 0x9f, 0x7f, 0x13, 0xeb, 0x7d, 0x2f, 0x00, 0xfe, 0x21,
	0xa0, 0xc5, 0xd6, 0xd3, 0xb9, 0xbc, 0xf6, 0x23, 0xe2, 0xa3, 0x00, 0x46, 0x9b, 0x66, 0x41, 0x9e,
	0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x14, 0x7a, 0x88, 0xe6, 0x76, 0x03, 0x00, 0x00,
}
