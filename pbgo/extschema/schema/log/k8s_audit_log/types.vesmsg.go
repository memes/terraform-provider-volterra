//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package k8s_audit_log

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

func (m *AggregationRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AggregationRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AggregationRequest) DeepCopy() *AggregationRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AggregationRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AggregationRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AggregationRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AggregationRequestValidator().Validate(ctx, m, opts...)
}

type ValidateAggregationRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAggregationRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AggregationRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AggregationRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetAggregationType().(type) {
	case *AggregationRequest_DateAggregation:
		if fv, exists := v.FldValidators["aggregation_type.date_aggregation"]; exists {
			val := m.GetAggregationType().(*AggregationRequest_DateAggregation).DateAggregation
			vOpts := append(opts,
				db.WithValidateField("aggregation_type"),
				db.WithValidateField("date_aggregation"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAggregationRequestValidator = func() *ValidateAggregationRequest {
	v := &ValidateAggregationRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["aggregation_type.date_aggregation"] = DateAggregationValidator().Validate

	return v
}()

func AggregationRequestValidator() db.Validator {
	return DefaultAggregationRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *DateAggregation) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DateAggregation) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DateAggregation) DeepCopy() *DateAggregation {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DateAggregation{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DateAggregation) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DateAggregation) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DateAggregationValidator().Validate(ctx, m, opts...)
}

type ValidateDateAggregation struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDateAggregation) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateDateAggregation) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DateAggregation)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DateAggregation got type %s", t)
		}
	}
	if m == nil {
		return nil
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
var DefaultDateAggregationValidator = func() *ValidateDateAggregation {
	v := &ValidateDateAggregation{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.message.required":  "true",
		"ves.io.schema.rules.string.query_step": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for DateAggregation.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func DateAggregationValidator() db.Validator {
	return DefaultDateAggregationValidator
}
