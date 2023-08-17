// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

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

func (m *SelecteeItemType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SelecteeItemType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SelecteeItemType) DeepCopy() *SelecteeItemType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SelecteeItemType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SelecteeItemType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SelecteeItemType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SelecteeItemTypeValidator().Validate(ctx, m, opts...)
}

type ValidateSelecteeItemType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSelecteeItemType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SelecteeItemType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SelecteeItemType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["kind"]; exists {

		vOpts := append(opts, db.WithValidateField("kind"))
		if err := fv(ctx, m.GetKind(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSelecteeItemTypeValidator = func() *ValidateSelecteeItemType {
	v := &ValidateSelecteeItemType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SelecteeItemTypeValidator() db.Validator {
	return DefaultSelecteeItemTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *SelecteeRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SelecteeRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SelecteeRequest) DeepCopy() *SelecteeRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SelecteeRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SelecteeRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SelecteeRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SelecteeRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSelecteeRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSelecteeRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SelecteeRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SelecteeRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSelecteeRequestValidator = func() *ValidateSelecteeRequest {
	v := &ValidateSelecteeRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SelecteeRequestValidator() db.Validator {
	return DefaultSelecteeRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SelecteeResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SelecteeResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SelecteeResponse) DeepCopy() *SelecteeResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SelecteeResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SelecteeResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SelecteeResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SelecteeResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSelecteeResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSelecteeResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SelecteeResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SelecteeResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["error"]; exists {

		vOpts := append(opts, db.WithValidateField("error"))
		if err := fv(ctx, m.GetError(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSelecteeResponseValidator = func() *ValidateSelecteeResponse {
	v := &ValidateSelecteeResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SelecteeResponseValidator() db.Validator {
	return DefaultSelecteeResponseValidator
}
