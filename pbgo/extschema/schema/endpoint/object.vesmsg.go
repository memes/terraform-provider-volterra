// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package endpoint

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

func (m *SpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetGcSpecDRefInfo()

}

// GetDRefInfo for the field's type
func (m *SpecType) GetGcSpecDRefInfo() ([]db.DRefInfo, error) {
	if m.GetGcSpec() == nil {
		return nil, nil
	}

	drInfos, err := m.GetGcSpec().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetGcSpec().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "gc_spec." + dri.DRField
	}
	return drInfos, err

}

type ValidateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSpecType) GcSpecValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for gc_spec")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := GlobalSpecTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
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

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhGcSpec := v.GcSpecValidationRuleHandler
	rulesGcSpec := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhGcSpec(rulesGcSpec)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SpecType.gc_spec: %s", err)
		panic(errMsg)
	}
	v.FldValidators["gc_spec"] = vFn

	return v
}()

func SpecTypeValidator() db.Validator {
	return DefaultSpecTypeValidator
}
