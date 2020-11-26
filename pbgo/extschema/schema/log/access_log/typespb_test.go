// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/log/access_log/types.proto

/*
Package access_log is a generated protocol buffer package.

It is generated from these files:
	ves.io/schema/log/access_log/types.proto

It has these top-level messages:
	DateSubAggregation
	DateAggregation
	FieldAggregation
	CardinalityAggregation
	AggregationRequest
*/
package access_log

import (
	rand "math/rand"
	testing "testing"

	time "time"

	proto "github.com/gogo/protobuf/proto"

	jsonpb "github.com/gogo/protobuf/jsonpb"

	fmt "fmt"

	parser "go/parser"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestDateSubAggregationProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateSubAggregation(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &DateSubAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestDateSubAggregationMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateSubAggregation(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &DateSubAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkDateSubAggregationProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*DateSubAggregation, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedDateSubAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDateSubAggregationProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedDateSubAggregation(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &DateSubAggregation{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestDateAggregationProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateAggregation(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &DateAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestDateAggregationMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateAggregation(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &DateAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkDateAggregationProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*DateAggregation, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedDateAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDateAggregationProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedDateAggregation(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &DateAggregation{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestFieldAggregationProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFieldAggregation(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &FieldAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestFieldAggregationMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFieldAggregation(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &FieldAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkFieldAggregationProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*FieldAggregation, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedFieldAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkFieldAggregationProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedFieldAggregation(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &FieldAggregation{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestCardinalityAggregationProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCardinalityAggregation(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CardinalityAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestCardinalityAggregationMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCardinalityAggregation(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CardinalityAggregation{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkCardinalityAggregationProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*CardinalityAggregation, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedCardinalityAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkCardinalityAggregationProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedCardinalityAggregation(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &CardinalityAggregation{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestAggregationRequestProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAggregationRequest(popr, false)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AggregationRequest{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = proto.Unmarshal(littlefuzz, msg)
	}
}

func TestAggregationRequestMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAggregationRequest(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AggregationRequest{}
	if err := proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkAggregationRequestProtoMarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*AggregationRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedAggregationRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAggregationRequestProtoUnmarshal(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := proto.Marshal(NewPopulatedAggregationRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &AggregationRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestDateSubAggregationJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateSubAggregation(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &DateSubAggregation{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestDateAggregationJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateAggregation(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &DateAggregation{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestFieldAggregationJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFieldAggregation(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &FieldAggregation{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestCardinalityAggregationJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCardinalityAggregation(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &CardinalityAggregation{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestAggregationRequestJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAggregationRequest(popr, true)
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &AggregationRequest{}
	err = jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestDateSubAggregationProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateSubAggregation(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &DateSubAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestDateSubAggregationProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateSubAggregation(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &DateSubAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestDateAggregationProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateAggregation(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &DateAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestDateAggregationProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateAggregation(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &DateAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFieldAggregationProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFieldAggregation(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &FieldAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFieldAggregationProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFieldAggregation(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &FieldAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestCardinalityAggregationProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCardinalityAggregation(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &CardinalityAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestCardinalityAggregationProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCardinalityAggregation(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &CardinalityAggregation{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestAggregationRequestProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAggregationRequest(popr, true)
	dAtA := proto.MarshalTextString(p)
	msg := &AggregationRequest{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestAggregationRequestProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAggregationRequest(popr, true)
	dAtA := proto.CompactTextString(p)
	msg := &AggregationRequest{}
	if err := proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestDateSubAggregationGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDateSubAggregation(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestDateAggregationGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDateAggregation(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestFieldAggregationGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFieldAggregation(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestCardinalityAggregationGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCardinalityAggregation(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestAggregationRequestGoString(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAggregationRequest(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestDateSubAggregationSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateSubAggregation(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkDateSubAggregationSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*DateSubAggregation, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedDateSubAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestDateAggregationSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedDateAggregation(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkDateAggregationSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*DateAggregation, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedDateAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestFieldAggregationSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedFieldAggregation(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkFieldAggregationSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*FieldAggregation, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedFieldAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestCardinalityAggregationSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedCardinalityAggregation(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkCardinalityAggregationSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*CardinalityAggregation, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedCardinalityAggregation(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestAggregationRequestSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := rand.New(rand.NewSource(seed))
	p := NewPopulatedAggregationRequest(popr, true)
	size2 := proto.Size(p)
	dAtA, err := proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkAggregationRequestSize(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*AggregationRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedAggregationRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestDateSubAggregationStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDateSubAggregation(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestDateAggregationStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDateAggregation(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestFieldAggregationStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFieldAggregation(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestCardinalityAggregationStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedCardinalityAggregation(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestAggregationRequestStringer(t *testing.T) {
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedAggregationRequest(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
