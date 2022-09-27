//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package dc_cluster_group

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

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DCClusterGroupMeshType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DCClusterGroupMeshType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DCClusterGroupMeshType) DeepCopy() *DCClusterGroupMeshType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DCClusterGroupMeshType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DCClusterGroupMeshType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DCClusterGroupMeshType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DCClusterGroupMeshTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDCClusterGroupMeshType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDCClusterGroupMeshType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DCClusterGroupMeshType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DCClusterGroupMeshType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetDcClusterGroupMeshChoice().(type) {
	case *DCClusterGroupMeshType_DataPlaneMesh:
		if fv, exists := v.FldValidators["dc_cluster_group_mesh_choice.data_plane_mesh"]; exists {
			val := m.GetDcClusterGroupMeshChoice().(*DCClusterGroupMeshType_DataPlaneMesh).DataPlaneMesh
			vOpts := append(opts,
				db.WithValidateField("dc_cluster_group_mesh_choice"),
				db.WithValidateField("data_plane_mesh"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *DCClusterGroupMeshType_ControlAndDataPlaneMesh:
		if fv, exists := v.FldValidators["dc_cluster_group_mesh_choice.control_and_data_plane_mesh"]; exists {
			val := m.GetDcClusterGroupMeshChoice().(*DCClusterGroupMeshType_ControlAndDataPlaneMesh).ControlAndDataPlaneMesh
			vOpts := append(opts,
				db.WithValidateField("dc_cluster_group_mesh_choice"),
				db.WithValidateField("control_and_data_plane_mesh"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDCClusterGroupMeshTypeValidator = func() *ValidateDCClusterGroupMeshType {
	v := &ValidateDCClusterGroupMeshType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DCClusterGroupMeshTypeValidator() db.Validator {
	return DefaultDCClusterGroupMeshTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DcClusterGroupStatus) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DcClusterGroupStatus) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DcClusterGroupStatus) DeepCopy() *DcClusterGroupStatus {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DcClusterGroupStatus{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DcClusterGroupStatus) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DcClusterGroupStatus) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DcClusterGroupStatusValidator().Validate(ctx, m, opts...)
}

func (m *DcClusterGroupStatus) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetSiteInfoDRefInfo()

}

// GetDRefInfo for the field's type
func (m *DcClusterGroupStatus) GetSiteInfoDRefInfo() ([]db.DRefInfo, error) {
	if m.GetSiteInfo() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetSiteInfo() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetSiteInfo() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("site_info[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateDcClusterGroupStatus struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDcClusterGroupStatus) SiteInfoValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for site_info")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.SiteInfo, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.SiteInfoValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for site_info")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.SiteInfo)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.SiteInfo, got %T", val)
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
			return errors.Wrap(err, "repeated site_info")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items site_info")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateDcClusterGroupStatus) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DcClusterGroupStatus)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DcClusterGroupStatus got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["site_info"]; exists {
		vOpts := append(opts, db.WithValidateField("site_info"))
		if err := fv(ctx, m.GetSiteInfo(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDcClusterGroupStatusValidator = func() *ValidateDcClusterGroupStatus {
	v := &ValidateDcClusterGroupStatus{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhSiteInfo := v.SiteInfoValidationRuleHandler
	rulesSiteInfo := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteInfo(rulesSiteInfo)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for DcClusterGroupStatus.site_info: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_info"] = vFn

	return v
}()

func DcClusterGroupStatusValidator() db.Validator {
	return DefaultDcClusterGroupStatusValidator
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

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

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

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
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

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

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

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
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

}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
