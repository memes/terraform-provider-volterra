// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package fast_acl

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

func (m *FastACLHits) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FastACLHits) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FastACLHits) DeepCopy() *FastACLHits {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FastACLHits{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FastACLHits) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FastACLHits) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FastACLHitsValidator().Validate(ctx, m, opts...)
}

type ValidateFastACLHits struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFastACLHits) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FastACLHits)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FastACLHits got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["metric"]; exists {

		vOpts := append(opts, db.WithValidateField("metric"))
		for idx, item := range m.GetMetric() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFastACLHitsValidator = func() *ValidateFastACLHits {
	v := &ValidateFastACLHits{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func FastACLHitsValidator() db.Validator {
	return DefaultFastACLHitsValidator
}

// augmented methods on protoc/std generated struct

func (m *FastACLHitsId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FastACLHitsId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FastACLHitsId) DeepCopy() *FastACLHitsId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FastACLHitsId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FastACLHitsId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FastACLHitsId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FastACLHitsIdValidator().Validate(ctx, m, opts...)
}

type ValidateFastACLHitsId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFastACLHitsId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FastACLHitsId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FastACLHitsId got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["fast_acl"]; exists {

		vOpts := append(opts, db.WithValidateField("fast_acl"))
		if err := fv(ctx, m.GetFastAcl(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["fast_acl_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("fast_acl_rule"))
		if err := fv(ctx, m.GetFastAclRule(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site"]; exists {

		vOpts := append(opts, db.WithValidateField("site"))
		if err := fv(ctx, m.GetSite(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFastACLHitsIdValidator = func() *ValidateFastACLHitsId {
	v := &ValidateFastACLHitsId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func FastACLHitsIdValidator() db.Validator {
	return DefaultFastACLHitsIdValidator
}

// augmented methods on protoc/std generated struct

func (m *FastACLHitsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FastACLHitsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FastACLHitsRequest) DeepCopy() *FastACLHitsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FastACLHitsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FastACLHitsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FastACLHitsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FastACLHitsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateFastACLHitsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFastACLHitsRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateFastACLHitsRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateFastACLHitsRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateFastACLHitsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FastACLHitsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FastACLHitsRequest got type %s", t)
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

	if fv, exists := v.FldValidators["group_by"]; exists {

		vOpts := append(opts, db.WithValidateField("group_by"))
		for idx, item := range m.GetGroupBy() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		for idx, item := range m.GetLabelFilter() {
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
var DefaultFastACLHitsRequestValidator = func() *ValidateFastACLHitsRequest {
	v := &ValidateFastACLHitsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStartTime := v.StartTimeValidationRuleHandler
	rulesStartTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhStartTime(rulesStartTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FastACLHitsRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FastACLHitsRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.query_step": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FastACLHitsRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func FastACLHitsRequestValidator() db.Validator {
	return DefaultFastACLHitsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *FastACLHitsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FastACLHitsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FastACLHitsResponse) DeepCopy() *FastACLHitsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FastACLHitsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FastACLHitsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FastACLHitsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FastACLHitsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateFastACLHitsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFastACLHitsResponse) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateFastACLHitsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FastACLHitsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FastACLHitsResponse got type %s", t)
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
var DefaultFastACLHitsResponseValidator = func() *ValidateFastACLHitsResponse {
	v := &ValidateFastACLHitsResponse{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for FastACLHitsResponse.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func FastACLHitsResponseValidator() db.Validator {
	return DefaultFastACLHitsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *FastACLMetricLabelFilter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FastACLMetricLabelFilter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FastACLMetricLabelFilter) DeepCopy() *FastACLMetricLabelFilter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FastACLMetricLabelFilter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FastACLMetricLabelFilter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FastACLMetricLabelFilter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FastACLMetricLabelFilterValidator().Validate(ctx, m, opts...)
}

type ValidateFastACLMetricLabelFilter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFastACLMetricLabelFilter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FastACLMetricLabelFilter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FastACLMetricLabelFilter got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["label"]; exists {

		vOpts := append(opts, db.WithValidateField("label"))
		if err := fv(ctx, m.GetLabel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["op"]; exists {

		vOpts := append(opts, db.WithValidateField("op"))
		if err := fv(ctx, m.GetOp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		if err := fv(ctx, m.GetValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFastACLMetricLabelFilterValidator = func() *ValidateFastACLMetricLabelFilter {
	v := &ValidateFastACLMetricLabelFilter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func FastACLMetricLabelFilterValidator() db.Validator {
	return DefaultFastACLMetricLabelFilterValidator
}
