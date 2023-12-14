// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_re_region

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

func (m *AWSType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AWSType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AWSType) DeepCopy() *AWSType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AWSType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AWSType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AWSType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AWSTypeValidator().Validate(ctx, m, opts...)
}

func (m *AWSType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNodesDRefInfo()

}

// GetDRefInfo for the field's type
func (m *AWSType) GetNodesDRefInfo() ([]db.DRefInfo, error) {
	if m.GetNodes() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetNodes() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetNodes() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("nodes[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateAWSType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAWSType) TgwSubnetValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for tgw_subnet")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.IpSubnetTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateAWSType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AWSType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AWSType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["inside_gre_subnet"]; exists {

		vOpts := append(opts, db.WithValidateField("inside_gre_subnet"))
		if err := fv(ctx, m.GetInsideGreSubnet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["nodes"]; exists {

		vOpts := append(opts, db.WithValidateField("nodes"))
		for idx, item := range m.GetNodes() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["tgw_subnet"]; exists {

		vOpts := append(opts, db.WithValidateField("tgw_subnet"))
		if err := fv(ctx, m.GetTgwSubnet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tgws"]; exists {

		vOpts := append(opts, db.WithValidateField("tgws"))
		for idx, item := range m.GetTgws() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAWSTypeValidator = func() *ValidateAWSType {
	v := &ValidateAWSType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhTgwSubnet := v.TgwSubnetValidationRuleHandler
	rulesTgwSubnet := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTgwSubnet(rulesTgwSubnet)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AWSType.tgw_subnet: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tgw_subnet"] = vFn

	v.FldValidators["tgws"] = TGWTypeValidator().Validate

	v.FldValidators["nodes"] = NodeTypeValidator().Validate

	v.FldValidators["inside_gre_subnet"] = ves_io_schema.IpSubnetTypeValidator().Validate

	return v
}()

func AWSTypeValidator() db.Validator {
	return DefaultAWSTypeValidator
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

	return m.GetCloudDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetCloudDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCloud() == nil {
		return nil, nil
	}
	switch m.GetCloud().(type) {
	case *GlobalSpecType_Aws:
		drInfos, err := m.GetAws().GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAws().GetDRefInfo() FAILED")
		}
		for i := range drInfos {
			dri := &drInfos[i]
			dri.DRField = "aws." + dri.DRField
		}
		return drInfos, err

	default:
		return nil, nil
	}

}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) CloudValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cloud")
	}
	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) RegionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for region")
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

	if fv, exists := v.FldValidators["cloud"]; exists {
		val := m.GetCloud()
		vOpts := append(opts,
			db.WithValidateField("cloud"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetCloud().(type) {
	case *GlobalSpecType_Aws:
		if fv, exists := v.FldValidators["cloud.aws"]; exists {
			val := m.GetCloud().(*GlobalSpecType_Aws).Aws
			vOpts := append(opts,
				db.WithValidateField("cloud"),
				db.WithValidateField("aws"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["region"]; exists {

		vOpts := append(opts, db.WithValidateField("region"))
		if err := fv(ctx, m.GetRegion(), vOpts...); err != nil {
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

	vrhCloud := v.CloudValidationRuleHandler
	rulesCloud := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhCloud(rulesCloud)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.cloud: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cloud"] = vFn

	vrhRegion := v.RegionValidationRuleHandler
	rulesRegion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhRegion(rulesRegion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.region: %s", err)
		panic(errMsg)
	}
	v.FldValidators["region"] = vFn

	v.FldValidators["cloud.aws"] = AWSTypeValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *NodeType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NodeType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NodeType) DeepCopy() *NodeType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NodeType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NodeType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NodeType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NodeTypeValidator().Validate(ctx, m, opts...)
}

func (m *NodeType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetSiteDRefInfo()

}

func (m *NodeType) GetSiteDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetSite()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("site.Object")
	dri := db.DRefInfo{
		RefdType:   "site.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "site",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetSiteDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *NodeType) GetSiteDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "site.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: site")
	}

	vref := m.GetSite()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "site.Object",
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

	return entries, nil
}

type ValidateNodeType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNodeType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NodeType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NodeType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address"]; exists {

		vOpts := append(opts, db.WithValidateField("address"))
		if err := fv(ctx, m.GetAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
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
var DefaultNodeTypeValidator = func() *ValidateNodeType {
	v := &ValidateNodeType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["address"] = ves_io_schema.IpAddressTypeValidator().Validate

	v.FldValidators["site"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func NodeTypeValidator() db.Validator {
	return DefaultNodeTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *TGWType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *TGWType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *TGWType) DeepCopy() *TGWType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &TGWType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *TGWType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *TGWType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return TGWTypeValidator().Validate(ctx, m, opts...)
}

type ValidateTGWType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateTGWType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*TGWType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *TGWType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["asn"]; exists {

		vOpts := append(opts, db.WithValidateField("asn"))
		if err := fv(ctx, m.GetAsn(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tgw_subnet"]; exists {

		vOpts := append(opts, db.WithValidateField("tgw_subnet"))
		if err := fv(ctx, m.GetTgwSubnet(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultTGWTypeValidator = func() *ValidateTGWType {
	v := &ValidateTGWType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tgw_subnet"] = ves_io_schema.IpSubnetTypeValidator().Validate

	return v
}()

func TGWTypeValidator() db.Validator {
	return DefaultTGWTypeValidator
}
