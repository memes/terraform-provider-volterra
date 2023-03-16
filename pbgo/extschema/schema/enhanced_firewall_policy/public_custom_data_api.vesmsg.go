//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package enhanced_firewall_policy

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

func (m *EnhancedFirewallPolicyHits) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnhancedFirewallPolicyHits) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnhancedFirewallPolicyHits) DeepCopy() *EnhancedFirewallPolicyHits {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnhancedFirewallPolicyHits{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnhancedFirewallPolicyHits) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnhancedFirewallPolicyHits) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnhancedFirewallPolicyHitsValidator().Validate(ctx, m, opts...)
}

type ValidateEnhancedFirewallPolicyHits struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnhancedFirewallPolicyHits) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnhancedFirewallPolicyHits)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnhancedFirewallPolicyHits got type %s", t)
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
var DefaultEnhancedFirewallPolicyHitsValidator = func() *ValidateEnhancedFirewallPolicyHits {
	v := &ValidateEnhancedFirewallPolicyHits{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EnhancedFirewallPolicyHitsValidator() db.Validator {
	return DefaultEnhancedFirewallPolicyHitsValidator
}

// augmented methods on protoc/std generated struct

func (m *EnhancedFirewallPolicyHitsId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnhancedFirewallPolicyHitsId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnhancedFirewallPolicyHitsId) DeepCopy() *EnhancedFirewallPolicyHitsId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnhancedFirewallPolicyHitsId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnhancedFirewallPolicyHitsId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnhancedFirewallPolicyHitsId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnhancedFirewallPolicyHitsIdValidator().Validate(ctx, m, opts...)
}

type ValidateEnhancedFirewallPolicyHitsId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnhancedFirewallPolicyHitsId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnhancedFirewallPolicyHitsId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnhancedFirewallPolicyHitsId got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["action"]; exists {

		vOpts := append(opts, db.WithValidateField("action"))
		if err := fv(ctx, m.GetAction(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policy"]; exists {

		vOpts := append(opts, db.WithValidateField("policy"))
		if err := fv(ctx, m.GetPolicy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policy_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("policy_rule"))
		if err := fv(ctx, m.GetPolicyRule(), vOpts...); err != nil {
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
var DefaultEnhancedFirewallPolicyHitsIdValidator = func() *ValidateEnhancedFirewallPolicyHitsId {
	v := &ValidateEnhancedFirewallPolicyHitsId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EnhancedFirewallPolicyHitsIdValidator() db.Validator {
	return DefaultEnhancedFirewallPolicyHitsIdValidator
}

// augmented methods on protoc/std generated struct

func (m *EnhancedFirewallPolicyHitsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnhancedFirewallPolicyHitsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnhancedFirewallPolicyHitsRequest) DeepCopy() *EnhancedFirewallPolicyHitsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnhancedFirewallPolicyHitsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnhancedFirewallPolicyHitsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnhancedFirewallPolicyHitsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnhancedFirewallPolicyHitsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateEnhancedFirewallPolicyHitsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnhancedFirewallPolicyHitsRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateEnhancedFirewallPolicyHitsRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateEnhancedFirewallPolicyHitsRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateEnhancedFirewallPolicyHitsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnhancedFirewallPolicyHitsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnhancedFirewallPolicyHitsRequest got type %s", t)
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
var DefaultEnhancedFirewallPolicyHitsRequestValidator = func() *ValidateEnhancedFirewallPolicyHitsRequest {
	v := &ValidateEnhancedFirewallPolicyHitsRequest{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for EnhancedFirewallPolicyHitsRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EnhancedFirewallPolicyHitsRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.query_step": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EnhancedFirewallPolicyHitsRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func EnhancedFirewallPolicyHitsRequestValidator() db.Validator {
	return DefaultEnhancedFirewallPolicyHitsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *EnhancedFirewallPolicyHitsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnhancedFirewallPolicyHitsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnhancedFirewallPolicyHitsResponse) DeepCopy() *EnhancedFirewallPolicyHitsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnhancedFirewallPolicyHitsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnhancedFirewallPolicyHitsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnhancedFirewallPolicyHitsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnhancedFirewallPolicyHitsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateEnhancedFirewallPolicyHitsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnhancedFirewallPolicyHitsResponse) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateEnhancedFirewallPolicyHitsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnhancedFirewallPolicyHitsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnhancedFirewallPolicyHitsResponse got type %s", t)
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
var DefaultEnhancedFirewallPolicyHitsResponseValidator = func() *ValidateEnhancedFirewallPolicyHitsResponse {
	v := &ValidateEnhancedFirewallPolicyHitsResponse{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for EnhancedFirewallPolicyHitsResponse.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func EnhancedFirewallPolicyHitsResponseValidator() db.Validator {
	return DefaultEnhancedFirewallPolicyHitsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *EnhancedFirewallPolicyMetricLabelFilter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnhancedFirewallPolicyMetricLabelFilter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnhancedFirewallPolicyMetricLabelFilter) DeepCopy() *EnhancedFirewallPolicyMetricLabelFilter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnhancedFirewallPolicyMetricLabelFilter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnhancedFirewallPolicyMetricLabelFilter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnhancedFirewallPolicyMetricLabelFilter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnhancedFirewallPolicyMetricLabelFilterValidator().Validate(ctx, m, opts...)
}

type ValidateEnhancedFirewallPolicyMetricLabelFilter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnhancedFirewallPolicyMetricLabelFilter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnhancedFirewallPolicyMetricLabelFilter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnhancedFirewallPolicyMetricLabelFilter got type %s", t)
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
var DefaultEnhancedFirewallPolicyMetricLabelFilterValidator = func() *ValidateEnhancedFirewallPolicyMetricLabelFilter {
	v := &ValidateEnhancedFirewallPolicyMetricLabelFilter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EnhancedFirewallPolicyMetricLabelFilterValidator() db.Validator {
	return DefaultEnhancedFirewallPolicyMetricLabelFilterValidator
}
