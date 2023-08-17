// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package protocol_inspection

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *CreateSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetEnableDisableComplianceChecksDRefInfo()

}

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetEnableDisableComplianceChecksDRefInfo() ([]db.DRefInfo, error) {
	if m.GetEnableDisableComplianceChecks() == nil {
		return nil, nil
	}

	drInfos, err := m.GetEnableDisableComplianceChecks().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetEnableDisableComplianceChecks().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "enable_disable_compliance_checks." + dri.DRField
	}
	return drInfos, err

}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) EnableDisableComplianceChecksValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_compliance_checks")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableComplianceChecksValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) EnableDisableSignaturesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_signatures")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableSignaturesValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
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

	if fv, exists := v.FldValidators["enable_disable_compliance_checks"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_compliance_checks"))
		if err := fv(ctx, m.GetEnableDisableComplianceChecks(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["enable_disable_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_signatures"))
		if err := fv(ctx, m.GetEnableDisableSignatures(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEnableDisableComplianceChecks := v.EnableDisableComplianceChecksValidationRuleHandler
	rulesEnableDisableComplianceChecks := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableComplianceChecks(rulesEnableDisableComplianceChecks)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.enable_disable_compliance_checks: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_compliance_checks"] = vFn

	vrhEnableDisableSignatures := v.EnableDisableSignaturesValidationRuleHandler
	rulesEnableDisableSignatures := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableSignatures(rulesEnableDisableSignatures)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.enable_disable_signatures: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_signatures"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *EnableDisableComplianceChecks) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnableDisableComplianceChecks) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnableDisableComplianceChecks) DeepCopy() *EnableDisableComplianceChecks {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnableDisableComplianceChecks{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnableDisableComplianceChecks) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnableDisableComplianceChecks) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnableDisableComplianceChecksValidator().Validate(ctx, m, opts...)
}

func (m *EnableDisableComplianceChecks) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetComplianceCheckChoiceDRefInfo()

}

func (m *EnableDisableComplianceChecks) GetComplianceCheckChoiceDRefInfo() ([]db.DRefInfo, error) {
	switch m.GetComplianceCheckChoice().(type) {
	case *EnableDisableComplianceChecks_EnableComplianceChecks:

		vref := m.GetEnableComplianceChecks()
		if vref == nil {
			return nil, nil
		}
		vdRef := db.NewDirectRefForView(vref)
		vdRef.SetKind("dns_compliance_checks.Object")
		dri := db.DRefInfo{
			RefdType:   "dns_compliance_checks.Object",
			RefdTenant: vref.Tenant,
			RefdNS:     vref.Namespace,
			RefdName:   vref.Name,
			DRField:    "enable_compliance_checks",
			Ref:        vdRef,
		}
		return []db.DRefInfo{dri}, nil

	case *EnableDisableComplianceChecks_DisableComplianceChecks:

		return nil, nil

	default:
		return nil, nil
	}
}

// GetComplianceCheckChoiceDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *EnableDisableComplianceChecks) GetComplianceCheckChoiceDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry

	switch m.GetComplianceCheckChoice().(type) {
	case *EnableDisableComplianceChecks_EnableComplianceChecks:
		refdType, err := d.TypeForEntryKind("", "", "dns_compliance_checks.Object")
		if err != nil {
			return nil, errors.Wrap(err, "Cannot find type for kind: dns_compliance_checks")
		}

		vref := m.GetEnableComplianceChecks()
		if vref == nil {
			return nil, nil
		}
		ref := &ves_io_schema.ObjectRefType{
			Kind:      "dns_compliance_checks.Object",
			Tenant:    vref.Tenant,
			Namespace: vref.Namespace,
			Name:      vref.Name,
		}
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}

	case *EnableDisableComplianceChecks_DisableComplianceChecks:

	}

	return entries, nil
}

type ValidateEnableDisableComplianceChecks struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnableDisableComplianceChecks) ComplianceCheckChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for compliance_check_choice")
	}
	return validatorFn, nil
}

func (v *ValidateEnableDisableComplianceChecks) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnableDisableComplianceChecks)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnableDisableComplianceChecks got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["compliance_check_choice"]; exists {
		val := m.GetComplianceCheckChoice()
		vOpts := append(opts,
			db.WithValidateField("compliance_check_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetComplianceCheckChoice().(type) {
	case *EnableDisableComplianceChecks_EnableComplianceChecks:
		if fv, exists := v.FldValidators["compliance_check_choice.enable_compliance_checks"]; exists {
			val := m.GetComplianceCheckChoice().(*EnableDisableComplianceChecks_EnableComplianceChecks).EnableComplianceChecks
			vOpts := append(opts,
				db.WithValidateField("compliance_check_choice"),
				db.WithValidateField("enable_compliance_checks"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *EnableDisableComplianceChecks_DisableComplianceChecks:
		if fv, exists := v.FldValidators["compliance_check_choice.disable_compliance_checks"]; exists {
			val := m.GetComplianceCheckChoice().(*EnableDisableComplianceChecks_DisableComplianceChecks).DisableComplianceChecks
			vOpts := append(opts,
				db.WithValidateField("compliance_check_choice"),
				db.WithValidateField("disable_compliance_checks"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEnableDisableComplianceChecksValidator = func() *ValidateEnableDisableComplianceChecks {
	v := &ValidateEnableDisableComplianceChecks{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhComplianceCheckChoice := v.ComplianceCheckChoiceValidationRuleHandler
	rulesComplianceCheckChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhComplianceCheckChoice(rulesComplianceCheckChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EnableDisableComplianceChecks.compliance_check_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["compliance_check_choice"] = vFn

	v.FldValidators["compliance_check_choice.enable_compliance_checks"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func EnableDisableComplianceChecksValidator() db.Validator {
	return DefaultEnableDisableComplianceChecksValidator
}

// augmented methods on protoc/std generated struct

func (m *EnableDisableSignatures) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EnableDisableSignatures) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EnableDisableSignatures) DeepCopy() *EnableDisableSignatures {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EnableDisableSignatures{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EnableDisableSignatures) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EnableDisableSignatures) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EnableDisableSignaturesValidator().Validate(ctx, m, opts...)
}

type ValidateEnableDisableSignatures struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEnableDisableSignatures) SignatureChoiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for signature_choice")
	}
	return validatorFn, nil
}

func (v *ValidateEnableDisableSignatures) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EnableDisableSignatures)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EnableDisableSignatures got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["signature_choice"]; exists {
		val := m.GetSignatureChoice()
		vOpts := append(opts,
			db.WithValidateField("signature_choice"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetSignatureChoice().(type) {
	case *EnableDisableSignatures_EnableSignature:
		if fv, exists := v.FldValidators["signature_choice.enable_signature"]; exists {
			val := m.GetSignatureChoice().(*EnableDisableSignatures_EnableSignature).EnableSignature
			vOpts := append(opts,
				db.WithValidateField("signature_choice"),
				db.WithValidateField("enable_signature"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *EnableDisableSignatures_DisableSignature:
		if fv, exists := v.FldValidators["signature_choice.disable_signature"]; exists {
			val := m.GetSignatureChoice().(*EnableDisableSignatures_DisableSignature).DisableSignature
			vOpts := append(opts,
				db.WithValidateField("signature_choice"),
				db.WithValidateField("disable_signature"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEnableDisableSignaturesValidator = func() *ValidateEnableDisableSignatures {
	v := &ValidateEnableDisableSignatures{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhSignatureChoice := v.SignatureChoiceValidationRuleHandler
	rulesSignatureChoice := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhSignatureChoice(rulesSignatureChoice)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EnableDisableSignatures.signature_choice: %s", err)
		panic(errMsg)
	}
	v.FldValidators["signature_choice"] = vFn

	return v
}()

func EnableDisableSignaturesValidator() db.Validator {
	return DefaultEnableDisableSignaturesValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *GetSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetEnableDisableComplianceChecksDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetSpecType) GetEnableDisableComplianceChecksDRefInfo() ([]db.DRefInfo, error) {
	if m.GetEnableDisableComplianceChecks() == nil {
		return nil, nil
	}

	drInfos, err := m.GetEnableDisableComplianceChecks().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetEnableDisableComplianceChecks().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "enable_disable_compliance_checks." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) EnableDisableComplianceChecksValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_compliance_checks")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableComplianceChecksValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) EnableDisableSignaturesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_signatures")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableSignaturesValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
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

	if fv, exists := v.FldValidators["enable_disable_compliance_checks"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_compliance_checks"))
		if err := fv(ctx, m.GetEnableDisableComplianceChecks(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["enable_disable_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_signatures"))
		if err := fv(ctx, m.GetEnableDisableSignatures(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEnableDisableComplianceChecks := v.EnableDisableComplianceChecksValidationRuleHandler
	rulesEnableDisableComplianceChecks := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableComplianceChecks(rulesEnableDisableComplianceChecks)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.enable_disable_compliance_checks: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_compliance_checks"] = vFn

	vrhEnableDisableSignatures := v.EnableDisableSignaturesValidationRuleHandler
	rulesEnableDisableSignatures := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableSignatures(rulesEnableDisableSignatures)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.enable_disable_signatures: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_signatures"] = vFn

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetEnableDisableComplianceChecksDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetEnableDisableComplianceChecksDRefInfo() ([]db.DRefInfo, error) {
	if m.GetEnableDisableComplianceChecks() == nil {
		return nil, nil
	}

	drInfos, err := m.GetEnableDisableComplianceChecks().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetEnableDisableComplianceChecks().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "enable_disable_compliance_checks." + dri.DRField
	}
	return drInfos, err

}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) EnableDisableComplianceChecksValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_compliance_checks")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableComplianceChecksValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) EnableDisableSignaturesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_signatures")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableSignaturesValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
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

	if fv, exists := v.FldValidators["enable_disable_compliance_checks"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_compliance_checks"))
		if err := fv(ctx, m.GetEnableDisableComplianceChecks(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["enable_disable_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_signatures"))
		if err := fv(ctx, m.GetEnableDisableSignatures(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEnableDisableComplianceChecks := v.EnableDisableComplianceChecksValidationRuleHandler
	rulesEnableDisableComplianceChecks := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableComplianceChecks(rulesEnableDisableComplianceChecks)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.enable_disable_compliance_checks: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_compliance_checks"] = vFn

	vrhEnableDisableSignatures := v.EnableDisableSignaturesValidationRuleHandler
	rulesEnableDisableSignatures := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableSignatures(rulesEnableDisableSignatures)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.enable_disable_signatures: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_signatures"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *ReplaceSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetEnableDisableComplianceChecksDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ReplaceSpecType) GetEnableDisableComplianceChecksDRefInfo() ([]db.DRefInfo, error) {
	if m.GetEnableDisableComplianceChecks() == nil {
		return nil, nil
	}

	drInfos, err := m.GetEnableDisableComplianceChecks().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetEnableDisableComplianceChecks().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "enable_disable_compliance_checks." + dri.DRField
	}
	return drInfos, err

}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) EnableDisableComplianceChecksValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_compliance_checks")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableComplianceChecksValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) EnableDisableSignaturesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for enable_disable_signatures")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EnableDisableSignaturesValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
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

	if fv, exists := v.FldValidators["enable_disable_compliance_checks"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_compliance_checks"))
		if err := fv(ctx, m.GetEnableDisableComplianceChecks(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["enable_disable_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("enable_disable_signatures"))
		if err := fv(ctx, m.GetEnableDisableSignatures(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEnableDisableComplianceChecks := v.EnableDisableComplianceChecksValidationRuleHandler
	rulesEnableDisableComplianceChecks := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableComplianceChecks(rulesEnableDisableComplianceChecks)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.enable_disable_compliance_checks: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_compliance_checks"] = vFn

	vrhEnableDisableSignatures := v.EnableDisableSignaturesValidationRuleHandler
	rulesEnableDisableSignatures := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEnableDisableSignatures(rulesEnableDisableSignatures)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.enable_disable_signatures: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable_disable_signatures"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.Action = f.GetAction()
	m.EnableDisableComplianceChecks = f.GetEnableDisableComplianceChecks()
	m.EnableDisableSignatures = f.GetEnableDisableSignatures()
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *CreateSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *CreateSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.Action = m1.Action
	f.EnableDisableComplianceChecks = m1.EnableDisableComplianceChecks
	f.EnableDisableSignatures = m1.EnableDisableSignatures
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *CreateSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.Action = f.GetAction()
	m.EnableDisableComplianceChecks = f.GetEnableDisableComplianceChecks()
	m.EnableDisableSignatures = f.GetEnableDisableSignatures()
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *GetSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *GetSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.Action = m1.Action
	f.EnableDisableComplianceChecks = m1.EnableDisableComplianceChecks
	f.EnableDisableSignatures = m1.EnableDisableSignatures
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.Action = f.GetAction()
	m.EnableDisableComplianceChecks = f.GetEnableDisableComplianceChecks()
	m.EnableDisableSignatures = f.GetEnableDisableSignatures()
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.Action = m1.Action
	f.EnableDisableComplianceChecks = m1.EnableDisableComplianceChecks
	f.EnableDisableSignatures = m1.EnableDisableSignatures
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
