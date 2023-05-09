// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package site

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

func (m *UpgradeOSRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpgradeOSRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpgradeOSRequest) DeepCopy() *UpgradeOSRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpgradeOSRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpgradeOSRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpgradeOSRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpgradeOSRequestValidator().Validate(ctx, m, opts...)
}

type ValidateUpgradeOSRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpgradeOSRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateUpgradeOSRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateUpgradeOSRequest) VersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for version")
	}

	return validatorFn, nil
}

func (v *ValidateUpgradeOSRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpgradeOSRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpgradeOSRequest got type %s", t)
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

	if fv, exists := v.FldValidators["version"]; exists {

		vOpts := append(opts, db.WithValidateField("version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpgradeOSRequestValidator = func() *ValidateUpgradeOSRequest {
	v := &ValidateUpgradeOSRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpgradeOSRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpgradeOSRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhVersion := v.VersionValidationRuleHandler
	rulesVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVersion(rulesVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpgradeOSRequest.version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["version"] = vFn

	return v
}()

func UpgradeOSRequestValidator() db.Validator {
	return DefaultUpgradeOSRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *UpgradeOSResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpgradeOSResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpgradeOSResponse) DeepCopy() *UpgradeOSResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpgradeOSResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpgradeOSResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpgradeOSResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpgradeOSResponseValidator().Validate(ctx, m, opts...)
}

type ValidateUpgradeOSResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpgradeOSResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpgradeOSResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpgradeOSResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpgradeOSResponseValidator = func() *ValidateUpgradeOSResponse {
	v := &ValidateUpgradeOSResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UpgradeOSResponseValidator() db.Validator {
	return DefaultUpgradeOSResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *UpgradeSWRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpgradeSWRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpgradeSWRequest) DeepCopy() *UpgradeSWRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpgradeSWRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpgradeSWRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpgradeSWRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpgradeSWRequestValidator().Validate(ctx, m, opts...)
}

type ValidateUpgradeSWRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpgradeSWRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateUpgradeSWRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateUpgradeSWRequest) VersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for version")
	}

	return validatorFn, nil
}

func (v *ValidateUpgradeSWRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpgradeSWRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpgradeSWRequest got type %s", t)
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

	if fv, exists := v.FldValidators["version"]; exists {

		vOpts := append(opts, db.WithValidateField("version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpgradeSWRequestValidator = func() *ValidateUpgradeSWRequest {
	v := &ValidateUpgradeSWRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpgradeSWRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpgradeSWRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhVersion := v.VersionValidationRuleHandler
	rulesVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVersion(rulesVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpgradeSWRequest.version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["version"] = vFn

	return v
}()

func UpgradeSWRequestValidator() db.Validator {
	return DefaultUpgradeSWRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *UpgradeSWResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpgradeSWResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpgradeSWResponse) DeepCopy() *UpgradeSWResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpgradeSWResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpgradeSWResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpgradeSWResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpgradeSWResponseValidator().Validate(ctx, m, opts...)
}

type ValidateUpgradeSWResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpgradeSWResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpgradeSWResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpgradeSWResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpgradeSWResponseValidator = func() *ValidateUpgradeSWResponse {
	v := &ValidateUpgradeSWResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UpgradeSWResponseValidator() db.Validator {
	return DefaultUpgradeSWResponseValidator
}
