//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package network_policy

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

func (m *NetworkPolicyHits) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NetworkPolicyHits) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NetworkPolicyHits) DeepCopy() *NetworkPolicyHits {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NetworkPolicyHits{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NetworkPolicyHits) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NetworkPolicyHits) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NetworkPolicyHitsValidator().Validate(ctx, m, opts...)
}

type ValidateNetworkPolicyHits struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNetworkPolicyHits) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NetworkPolicyHits)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NetworkPolicyHits got type %s", t)
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
var DefaultNetworkPolicyHitsValidator = func() *ValidateNetworkPolicyHits {
	v := &ValidateNetworkPolicyHits{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NetworkPolicyHitsValidator() db.Validator {
	return DefaultNetworkPolicyHitsValidator
}

// augmented methods on protoc/std generated struct

func (m *NetworkPolicyHitsId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NetworkPolicyHitsId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NetworkPolicyHitsId) DeepCopy() *NetworkPolicyHitsId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NetworkPolicyHitsId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NetworkPolicyHitsId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NetworkPolicyHitsId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NetworkPolicyHitsIdValidator().Validate(ctx, m, opts...)
}

type ValidateNetworkPolicyHitsId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNetworkPolicyHitsId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NetworkPolicyHitsId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NetworkPolicyHitsId got type %s", t)
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

	if fv, exists := v.FldValidators["policy_set"]; exists {

		vOpts := append(opts, db.WithValidateField("policy_set"))
		if err := fv(ctx, m.GetPolicySet(), vOpts...); err != nil {
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
var DefaultNetworkPolicyHitsIdValidator = func() *ValidateNetworkPolicyHitsId {
	v := &ValidateNetworkPolicyHitsId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NetworkPolicyHitsIdValidator() db.Validator {
	return DefaultNetworkPolicyHitsIdValidator
}

// augmented methods on protoc/std generated struct

func (m *NetworkPolicyHitsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NetworkPolicyHitsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NetworkPolicyHitsRequest) DeepCopy() *NetworkPolicyHitsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NetworkPolicyHitsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NetworkPolicyHitsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NetworkPolicyHitsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NetworkPolicyHitsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateNetworkPolicyHitsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNetworkPolicyHitsRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateNetworkPolicyHitsRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateNetworkPolicyHitsRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateNetworkPolicyHitsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NetworkPolicyHitsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NetworkPolicyHitsRequest got type %s", t)
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
var DefaultNetworkPolicyHitsRequestValidator = func() *ValidateNetworkPolicyHitsRequest {
	v := &ValidateNetworkPolicyHitsRequest{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for NetworkPolicyHitsRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for NetworkPolicyHitsRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.query_step": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for NetworkPolicyHitsRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func NetworkPolicyHitsRequestValidator() db.Validator {
	return DefaultNetworkPolicyHitsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *NetworkPolicyHitsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NetworkPolicyHitsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NetworkPolicyHitsResponse) DeepCopy() *NetworkPolicyHitsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NetworkPolicyHitsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NetworkPolicyHitsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NetworkPolicyHitsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NetworkPolicyHitsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateNetworkPolicyHitsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNetworkPolicyHitsResponse) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateNetworkPolicyHitsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NetworkPolicyHitsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NetworkPolicyHitsResponse got type %s", t)
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
var DefaultNetworkPolicyHitsResponseValidator = func() *ValidateNetworkPolicyHitsResponse {
	v := &ValidateNetworkPolicyHitsResponse{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for NetworkPolicyHitsResponse.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func NetworkPolicyHitsResponseValidator() db.Validator {
	return DefaultNetworkPolicyHitsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *NetworkPolicyMetricLabelFilter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NetworkPolicyMetricLabelFilter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NetworkPolicyMetricLabelFilter) DeepCopy() *NetworkPolicyMetricLabelFilter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NetworkPolicyMetricLabelFilter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NetworkPolicyMetricLabelFilter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NetworkPolicyMetricLabelFilter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NetworkPolicyMetricLabelFilterValidator().Validate(ctx, m, opts...)
}

type ValidateNetworkPolicyMetricLabelFilter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNetworkPolicyMetricLabelFilter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NetworkPolicyMetricLabelFilter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NetworkPolicyMetricLabelFilter got type %s", t)
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
var DefaultNetworkPolicyMetricLabelFilterValidator = func() *ValidateNetworkPolicyMetricLabelFilter {
	v := &ValidateNetworkPolicyMetricLabelFilter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NetworkPolicyMetricLabelFilterValidator() db.Validator {
	return DefaultNetworkPolicyMetricLabelFilterValidator
}
