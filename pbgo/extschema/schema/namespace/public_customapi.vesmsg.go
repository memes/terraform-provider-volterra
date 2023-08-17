// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package namespace

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

func (m *APIItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIItem) DeepCopy() *APIItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIItemValidator().Validate(ctx, m, opts...)
}

type ValidateAPIItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["result"]; exists {

		vOpts := append(opts, db.WithValidateField("result"))
		if err := fv(ctx, m.GetResult(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIItemValidator = func() *ValidateAPIItem {
	v := &ValidateAPIItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIItemValidator() db.Validator {
	return DefaultAPIItemValidator
}

// augmented methods on protoc/std generated struct

func (m *APIItemList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIItemList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIItemList) DeepCopy() *APIItemList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIItemList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIItemList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIItemList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIItemListValidator().Validate(ctx, m, opts...)
}

type ValidateAPIItemList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIItemList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIItemList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIItemList got type %s", t)
		}
	}
	if m == nil {
		return nil
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

	if fv, exists := v.FldValidators["list_id"]; exists {

		vOpts := append(opts, db.WithValidateField("list_id"))
		if err := fv(ctx, m.GetListId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["result"]; exists {

		vOpts := append(opts, db.WithValidateField("result"))
		if err := fv(ctx, m.GetResult(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIItemListValidator = func() *ValidateAPIItemList {
	v := &ValidateAPIItemList{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIItemListValidator() db.Validator {
	return DefaultAPIItemListValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateAPIAccessReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateAPIAccessReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateAPIAccessReq) DeepCopy() *EvaluateAPIAccessReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateAPIAccessReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateAPIAccessReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateAPIAccessReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateAPIAccessReqValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateAPIAccessReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateAPIAccessReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateAPIAccessReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateAPIAccessReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item_lists"]; exists {

		vOpts := append(opts, db.WithValidateField("item_lists"))
		for idx, item := range m.GetItemLists() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
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
var DefaultEvaluateAPIAccessReqValidator = func() *ValidateEvaluateAPIAccessReq {
	v := &ValidateEvaluateAPIAccessReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EvaluateAPIAccessReqValidator() db.Validator {
	return DefaultEvaluateAPIAccessReqValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateAPIAccessResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateAPIAccessResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateAPIAccessResp) DeepCopy() *EvaluateAPIAccessResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateAPIAccessResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateAPIAccessResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateAPIAccessResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateAPIAccessRespValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateAPIAccessResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateAPIAccessResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateAPIAccessResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateAPIAccessResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item_lists"]; exists {

		vOpts := append(opts, db.WithValidateField("item_lists"))
		for idx, item := range m.GetItemLists() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateAPIAccessRespValidator = func() *ValidateEvaluateAPIAccessResp {
	v := &ValidateEvaluateAPIAccessResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EvaluateAPIAccessRespValidator() db.Validator {
	return DefaultEvaluateAPIAccessRespValidator
}

// augmented methods on protoc/std generated struct

func (m *LookupUserRolesReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *LookupUserRolesReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *LookupUserRolesReq) DeepCopy() *LookupUserRolesReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &LookupUserRolesReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *LookupUserRolesReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *LookupUserRolesReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LookupUserRolesReqValidator().Validate(ctx, m, opts...)
}

type ValidateLookupUserRolesReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLookupUserRolesReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*LookupUserRolesReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *LookupUserRolesReq got type %s", t)
		}
	}
	if m == nil {
		return nil
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
var DefaultLookupUserRolesReqValidator = func() *ValidateLookupUserRolesReq {
	v := &ValidateLookupUserRolesReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func LookupUserRolesReqValidator() db.Validator {
	return DefaultLookupUserRolesReqValidator
}

// augmented methods on protoc/std generated struct

func (m *LookupUserRolesResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *LookupUserRolesResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *LookupUserRolesResp) DeepCopy() *LookupUserRolesResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &LookupUserRolesResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *LookupUserRolesResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *LookupUserRolesResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LookupUserRolesRespValidator().Validate(ctx, m, opts...)
}

type ValidateLookupUserRolesResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLookupUserRolesResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*LookupUserRolesResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *LookupUserRolesResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["roles"]; exists {

		vOpts := append(opts, db.WithValidateField("roles"))
		for idx, item := range m.GetRoles() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLookupUserRolesRespValidator = func() *ValidateLookupUserRolesResp {
	v := &ValidateLookupUserRolesResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func LookupUserRolesRespValidator() db.Validator {
	return DefaultLookupUserRolesRespValidator
}
