// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package virtual_network

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

func (m *SIDCounterData) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SIDCounterData) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SIDCounterData) DeepCopy() *SIDCounterData {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SIDCounterData{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SIDCounterData) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SIDCounterData) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SIDCounterDataValidator().Validate(ctx, m, opts...)
}

type ValidateSIDCounterData struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSIDCounterData) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SIDCounterData)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SIDCounterData got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSIDCounterDataValidator = func() *ValidateSIDCounterData {
	v := &ValidateSIDCounterData{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SIDCounterDataValidator() db.Validator {
	return DefaultSIDCounterDataValidator
}

// augmented methods on protoc/std generated struct

func (m *SIDCounterRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SIDCounterRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SIDCounterRequest) DeepCopy() *SIDCounterRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SIDCounterRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SIDCounterRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SIDCounterRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SIDCounterRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSIDCounterRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSIDCounterRequest) FieldSelectorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(SIDCounterType)
		return int32(i)
	}
	// SIDCounterType_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, SIDCounterType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for field_selector")
	}
	itemsValidatorFn := func(ctx context.Context, elems []SIDCounterType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for field_selector")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]SIDCounterType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []SIDCounterType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated field_selector")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items field_selector")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSIDCounterRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateSIDCounterRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateSIDCounterRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateSIDCounterRequest) RangeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for range")
	}

	return validatorFn, nil
}

func (v *ValidateSIDCounterRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SIDCounterRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SIDCounterRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["field_selector"]; exists {
		vOpts := append(opts, db.WithValidateField("field_selector"))
		if err := fv(ctx, m.GetFieldSelector(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["filter"]; exists {

		vOpts := append(opts, db.WithValidateField("filter"))
		if err := fv(ctx, m.GetFilter(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["group_by"]; exists {

		vOpts := append(opts, db.WithValidateField("group_by"))
		for idx, item := range m.GetGroupBy() {
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

	if fv, exists := v.FldValidators["range"]; exists {

		vOpts := append(opts, db.WithValidateField("range"))
		if err := fv(ctx, m.GetRange(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSIDCounterRequestValidator = func() *ValidateSIDCounterRequest {
	v := &ValidateSIDCounterRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhFieldSelector := v.FieldSelectorValidationRuleHandler
	rulesFieldSelector := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "8",
		"ves.io.schema.rules.repeated.min_items": "1",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhFieldSelector(rulesFieldSelector)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SIDCounterRequest.field_selector: %s", err)
		panic(errMsg)
	}
	v.FldValidators["field_selector"] = vFn

	vrhStartTime := v.StartTimeValidationRuleHandler
	rulesStartTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhStartTime(rulesStartTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SIDCounterRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SIDCounterRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.query_step": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SIDCounterRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	vrhRange := v.RangeValidationRuleHandler
	rulesRange := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhRange(rulesRange)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SIDCounterRequest.range: %s", err)
		panic(errMsg)
	}
	v.FldValidators["range"] = vFn

	return v
}()

func SIDCounterRequestValidator() db.Validator {
	return DefaultSIDCounterRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SIDCounterResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SIDCounterResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SIDCounterResponse) DeepCopy() *SIDCounterResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SIDCounterResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SIDCounterResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SIDCounterResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SIDCounterResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSIDCounterResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSIDCounterResponse) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateSIDCounterResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SIDCounterResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SIDCounterResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSIDCounterResponseValidator = func() *ValidateSIDCounterResponse {
	v := &ValidateSIDCounterResponse{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SIDCounterResponse.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func SIDCounterResponseValidator() db.Validator {
	return DefaultSIDCounterResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SIDCounterTypeData) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SIDCounterTypeData) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SIDCounterTypeData) DeepCopy() *SIDCounterTypeData {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SIDCounterTypeData{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SIDCounterTypeData) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SIDCounterTypeData) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SIDCounterTypeDataValidator().Validate(ctx, m, opts...)
}

type ValidateSIDCounterTypeData struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSIDCounterTypeData) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SIDCounterTypeData)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SIDCounterTypeData got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		for key, value := range m.GetKey() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		for idx, item := range m.GetValue() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSIDCounterTypeDataValidator = func() *ValidateSIDCounterTypeData {
	v := &ValidateSIDCounterTypeData{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SIDCounterTypeDataValidator() db.Validator {
	return DefaultSIDCounterTypeDataValidator
}
