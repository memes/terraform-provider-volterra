//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package virtual_site

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *SpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SpecType) DeepCopy() *SpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *SpecType) GetSRefInfo() ([]db.SelrFldInfo, error) {
	if m == nil {
		return nil, nil
	}
	return m.GetGcSpecSRefInfo()

}

// GetGcSpecSRefInfo returns the selector info (fld-name/val, selectee-type) of this field
func (m *SpecType) GetGcSpecSRefInfo() ([]db.SelrFldInfo, error) {
	if m.GetGcSpec() == nil {
		return nil, nil
	}

	srInfos, err := m.GetGcSpec().GetSRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetGcSpec().GetSRefInfo() FAILED")
	}

	for i := range srInfos {
		sri := &srInfos[i]
		// Convert GlobalSpecType.vnRefs to LcSpec.vnRefs i.e. convert field-type to field-name
		sl := strings.Split(sri.Name, ".")
		sri.Name = "gc_spec." + strings.Join(sl[1:], ".")
	}
	return srInfos, nil

}

type ValidateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["gc_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("gc_spec"))
		if err := fv(ctx, m.GetGcSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSpecTypeValidator = func() *ValidateSpecType {
	v := &ValidateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["gc_spec"] = GlobalSpecTypeValidator().Validate

	return v
}()

func SpecTypeValidator() db.Validator {
	return DefaultSpecTypeValidator
}
