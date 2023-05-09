// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package rate_limiter

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
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

	return m.GetUserIdentificationDRefInfo()

}

func (m *CreateSpecType) GetUserIdentificationDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetUserIdentification()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("CreateSpecType.user_identification[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "user_identification.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "user_identification",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetUserIdentificationDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CreateSpecType) GetUserIdentificationDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "user_identification.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: user_identification")
	}
	for _, ref := range m.GetUserIdentification() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) LimitsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for limits")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*RateLimitValue, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := RateLimitValueValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for limits")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*RateLimitValue)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*RateLimitValue, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated limits")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items limits")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) UserIdentificationValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for user_identification")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for user_identification")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated user_identification")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items user_identification")
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

	if fv, exists := v.FldValidators["limits"]; exists {
		vOpts := append(opts, db.WithValidateField("limits"))
		if err := fv(ctx, m.GetLimits(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_identification"]; exists {
		vOpts := append(opts, db.WithValidateField("user_identification"))
		if err := fv(ctx, m.GetUserIdentification(), vOpts...); err != nil {
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

	vrhLimits := v.LimitsValidationRuleHandler
	rulesLimits := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhLimits(rulesLimits)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.limits: %s", err)
		panic(errMsg)
	}
	v.FldValidators["limits"] = vFn

	vrhUserIdentification := v.UserIdentificationValidationRuleHandler
	rulesUserIdentification := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhUserIdentification(rulesUserIdentification)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.user_identification: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_identification"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
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

	return m.GetUserIdentificationDRefInfo()

}

func (m *GetSpecType) GetUserIdentificationDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetUserIdentification()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GetSpecType.user_identification[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "user_identification.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "user_identification",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetUserIdentificationDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetSpecType) GetUserIdentificationDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "user_identification.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: user_identification")
	}
	for _, ref := range m.GetUserIdentification() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) LimitsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for limits")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*RateLimitValue, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := RateLimitValueValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for limits")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*RateLimitValue)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*RateLimitValue, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated limits")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items limits")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) UserIdentificationValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for user_identification")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for user_identification")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated user_identification")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items user_identification")
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

	if fv, exists := v.FldValidators["limits"]; exists {
		vOpts := append(opts, db.WithValidateField("limits"))
		if err := fv(ctx, m.GetLimits(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_identification"]; exists {
		vOpts := append(opts, db.WithValidateField("user_identification"))
		if err := fv(ctx, m.GetUserIdentification(), vOpts...); err != nil {
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

	vrhLimits := v.LimitsValidationRuleHandler
	rulesLimits := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhLimits(rulesLimits)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.limits: %s", err)
		panic(errMsg)
	}
	v.FldValidators["limits"] = vFn

	vrhUserIdentification := v.UserIdentificationValidationRuleHandler
	rulesUserIdentification := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhUserIdentification(rulesUserIdentification)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.user_identification: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_identification"] = vFn

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

	return m.GetUserIdentificationDRefInfo()

}

func (m *GlobalSpecType) GetUserIdentificationDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetUserIdentification()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.user_identification[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "user_identification.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "user_identification",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetUserIdentificationDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetUserIdentificationDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "user_identification.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: user_identification")
	}
	for _, ref := range m.GetUserIdentification() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) LimitsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for limits")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*RateLimitValue, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := RateLimitValueValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for limits")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*RateLimitValue)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*RateLimitValue, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated limits")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items limits")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) UserIdentificationValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for user_identification")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for user_identification")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated user_identification")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items user_identification")
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

	if fv, exists := v.FldValidators["limits"]; exists {
		vOpts := append(opts, db.WithValidateField("limits"))
		if err := fv(ctx, m.GetLimits(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_identification"]; exists {
		vOpts := append(opts, db.WithValidateField("user_identification"))
		if err := fv(ctx, m.GetUserIdentification(), vOpts...); err != nil {
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

	vrhLimits := v.LimitsValidationRuleHandler
	rulesLimits := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhLimits(rulesLimits)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.limits: %s", err)
		panic(errMsg)
	}
	v.FldValidators["limits"] = vFn

	vrhUserIdentification := v.UserIdentificationValidationRuleHandler
	rulesUserIdentification := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhUserIdentification(rulesUserIdentification)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.user_identification: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_identification"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *RateLimitValue) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RateLimitValue) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RateLimitValue) DeepCopy() *RateLimitValue {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RateLimitValue{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RateLimitValue) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RateLimitValue) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RateLimitValueValidator().Validate(ctx, m, opts...)
}

type ValidateRateLimitValue struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRateLimitValue) UnitValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(RateLimitPeriodUnit)
		return int32(i)
	}
	// RateLimitPeriodUnit_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, RateLimitPeriodUnit_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for unit")
	}

	return validatorFn, nil
}

func (v *ValidateRateLimitValue) TotalNumberValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for total_number")
	}

	return validatorFn, nil
}

func (v *ValidateRateLimitValue) BurstMultiplierValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for burst_multiplier")
	}

	return validatorFn, nil
}

func (v *ValidateRateLimitValue) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RateLimitValue)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RateLimitValue got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["burst_multiplier"]; exists {

		vOpts := append(opts, db.WithValidateField("burst_multiplier"))
		if err := fv(ctx, m.GetBurstMultiplier(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["total_number"]; exists {

		vOpts := append(opts, db.WithValidateField("total_number"))
		if err := fv(ctx, m.GetTotalNumber(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["unit"]; exists {

		vOpts := append(opts, db.WithValidateField("unit"))
		if err := fv(ctx, m.GetUnit(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRateLimitValueValidator = func() *ValidateRateLimitValue {
	v := &ValidateRateLimitValue{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhUnit := v.UnitValidationRuleHandler
	rulesUnit := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhUnit(rulesUnit)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RateLimitValue.unit: %s", err)
		panic(errMsg)
	}
	v.FldValidators["unit"] = vFn

	vrhTotalNumber := v.TotalNumberValidationRuleHandler
	rulesTotalNumber := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
		"ves.io.schema.rules.uint32.lte":       "8192",
	}
	vFn, err = vrhTotalNumber(rulesTotalNumber)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RateLimitValue.total_number: %s", err)
		panic(errMsg)
	}
	v.FldValidators["total_number"] = vFn

	vrhBurstMultiplier := v.BurstMultiplierValidationRuleHandler
	rulesBurstMultiplier := map[string]string{
		"ves.io.schema.rules.uint32.gt":  "0",
		"ves.io.schema.rules.uint32.lte": "100",
	}
	vFn, err = vrhBurstMultiplier(rulesBurstMultiplier)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RateLimitValue.burst_multiplier: %s", err)
		panic(errMsg)
	}
	v.FldValidators["burst_multiplier"] = vFn

	return v
}()

func RateLimitValueValidator() db.Validator {
	return DefaultRateLimitValueValidator
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

	return m.GetUserIdentificationDRefInfo()

}

func (m *ReplaceSpecType) GetUserIdentificationDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetUserIdentification()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("ReplaceSpecType.user_identification[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "user_identification.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "user_identification",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetUserIdentificationDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *ReplaceSpecType) GetUserIdentificationDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "user_identification.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: user_identification")
	}
	for _, ref := range m.GetUserIdentification() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) LimitsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for limits")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*RateLimitValue, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := RateLimitValueValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for limits")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*RateLimitValue)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*RateLimitValue, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated limits")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items limits")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) UserIdentificationValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for user_identification")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for user_identification")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated user_identification")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items user_identification")
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

	if fv, exists := v.FldValidators["limits"]; exists {
		vOpts := append(opts, db.WithValidateField("limits"))
		if err := fv(ctx, m.GetLimits(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_identification"]; exists {
		vOpts := append(opts, db.WithValidateField("user_identification"))
		if err := fv(ctx, m.GetUserIdentification(), vOpts...); err != nil {
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

	vrhLimits := v.LimitsValidationRuleHandler
	rulesLimits := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhLimits(rulesLimits)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.limits: %s", err)
		panic(errMsg)
	}
	v.FldValidators["limits"] = vFn

	vrhUserIdentification := v.UserIdentificationValidationRuleHandler
	rulesUserIdentification := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhUserIdentification(rulesUserIdentification)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.user_identification: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_identification"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.Limits = f.GetLimits()
	m.UserIdentification = f.GetUserIdentification()
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

	f.Limits = m1.Limits
	f.UserIdentification = m1.UserIdentification
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
	m.Limits = f.GetLimits()
	m.UserIdentification = f.GetUserIdentification()
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

	f.Limits = m1.Limits
	f.UserIdentification = m1.UserIdentification
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
	m.Limits = f.GetLimits()
	m.UserIdentification = f.GetUserIdentification()
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

	f.Limits = m1.Limits
	f.UserIdentification = m1.UserIdentification
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
