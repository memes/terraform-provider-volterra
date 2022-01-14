//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package terraform_parameters

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *AvailabilitySetsInfoType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AvailabilitySetsInfoType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AvailabilitySetsInfoType) DeepCopy() *AvailabilitySetsInfoType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AvailabilitySetsInfoType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AvailabilitySetsInfoType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AvailabilitySetsInfoType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AvailabilitySetsInfoTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAvailabilitySetsInfoType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAvailabilitySetsInfoType) EnableValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for enable")
	}

	return validatorFn, nil
}

func (v *ValidateAvailabilitySetsInfoType) FaultDomainsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for fault_domains")
	}

	return validatorFn, nil
}

func (v *ValidateAvailabilitySetsInfoType) UpdateDomainsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for update_domains")
	}

	return validatorFn, nil
}

func (v *ValidateAvailabilitySetsInfoType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AvailabilitySetsInfoType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AvailabilitySetsInfoType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["enable"]; exists {

		vOpts := append(opts, db.WithValidateField("enable"))
		if err := fv(ctx, m.GetEnable(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["fault_domains"]; exists {

		vOpts := append(opts, db.WithValidateField("fault_domains"))
		if err := fv(ctx, m.GetFaultDomains(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["update_domains"]; exists {

		vOpts := append(opts, db.WithValidateField("update_domains"))
		if err := fv(ctx, m.GetUpdateDomains(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAvailabilitySetsInfoTypeValidator = func() *ValidateAvailabilitySetsInfoType {
	v := &ValidateAvailabilitySetsInfoType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEnable := v.EnableValidationRuleHandler
	rulesEnable := map[string]string{
		"ves.io.schema.rules.string.in": "[\"true\",\"false\"]",
	}
	vFn, err = vrhEnable(rulesEnable)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AvailabilitySetsInfoType.enable: %s", err)
		panic(errMsg)
	}
	v.FldValidators["enable"] = vFn

	vrhFaultDomains := v.FaultDomainsValidationRuleHandler
	rulesFaultDomains := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "3",
	}
	vFn, err = vrhFaultDomains(rulesFaultDomains)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AvailabilitySetsInfoType.fault_domains: %s", err)
		panic(errMsg)
	}
	v.FldValidators["fault_domains"] = vFn

	vrhUpdateDomains := v.UpdateDomainsValidationRuleHandler
	rulesUpdateDomains := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "20",
	}
	vFn, err = vrhUpdateDomains(rulesUpdateDomains)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AvailabilitySetsInfoType.update_domains: %s", err)
		panic(errMsg)
	}
	v.FldValidators["update_domains"] = vFn

	return v
}()

func AvailabilitySetsInfoTypeValidator() db.Validator {
	return DefaultAvailabilitySetsInfoTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureExistingSubnetParamType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureExistingSubnetParamType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureExistingSubnetParamType) DeepCopy() *AzureExistingSubnetParamType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureExistingSubnetParamType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureExistingSubnetParamType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureExistingSubnetParamType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureExistingSubnetParamTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureExistingSubnetParamType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureExistingSubnetParamType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureExistingSubnetParamType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureExistingSubnetParamType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["subnet_name"]; exists {

		vOpts := append(opts, db.WithValidateField("subnet_name"))
		if err := fv(ctx, m.GetSubnetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subnet_resource_grp"]; exists {

		vOpts := append(opts, db.WithValidateField("subnet_resource_grp"))
		if err := fv(ctx, m.GetSubnetResourceGrp(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureExistingSubnetParamTypeValidator = func() *ValidateAzureExistingSubnetParamType {
	v := &ValidateAzureExistingSubnetParamType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func AzureExistingSubnetParamTypeValidator() db.Validator {
	return DefaultAzureExistingSubnetParamTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureInstanceType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureInstanceType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureInstanceType) DeepCopy() *AzureInstanceType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureInstanceType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureInstanceType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureInstanceType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureInstanceTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureInstanceType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureInstanceType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureInstanceType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureInstanceType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["az"]; exists {

		vOpts := append(opts, db.WithValidateField("az"))
		if err := fv(ctx, m.GetAz(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["disk_size"]; exists {

		vOpts := append(opts, db.WithValidateField("disk_size"))
		if err := fv(ctx, m.GetDiskSize(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["machine_type"]; exists {

		vOpts := append(opts, db.WithValidateField("machine_type"))
		if err := fv(ctx, m.GetMachineType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["marketplace"]; exists {

		vOpts := append(opts, db.WithValidateField("marketplace"))
		if err := fv(ctx, m.GetMarketplace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["node_count"]; exists {

		vOpts := append(opts, db.WithValidateField("node_count"))
		if err := fv(ctx, m.GetNodeCount(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["private_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("private_subnet_id"))
		if err := fv(ctx, m.GetPrivateSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("public_subnet_id"))
		if err := fv(ctx, m.GetPublicSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volt_node_id"]; exists {

		vOpts := append(opts, db.WithValidateField("volt_node_id"))
		if err := fv(ctx, m.GetVoltNodeId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volt_vnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("volt_vnet_id"))
		if err := fv(ctx, m.GetVoltVnetId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureInstanceTypeValidator = func() *ValidateAzureInstanceType {
	v := &ValidateAzureInstanceType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func AzureInstanceTypeValidator() db.Validator {
	return DefaultAzureInstanceTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureSubnetChoice) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureSubnetChoice) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureSubnetChoice) DeepCopy() *AzureSubnetChoice {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureSubnetChoice{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureSubnetChoice) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureSubnetChoice) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureSubnetChoiceValidator().Validate(ctx, m, opts...)
}

type ValidateAzureSubnetChoice struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureSubnetChoice) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureSubnetChoice)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureSubnetChoice got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["existing_subnet"]; exists {

		vOpts := append(opts, db.WithValidateField("existing_subnet"))
		if err := fv(ctx, m.GetExistingSubnet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subnet_param"]; exists {

		vOpts := append(opts, db.WithValidateField("subnet_param"))
		if err := fv(ctx, m.GetSubnetParam(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureSubnetChoiceValidator = func() *ValidateAzureSubnetChoice {
	v := &ValidateAzureSubnetChoice{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["subnet_param"] = AzureSubnetParamTypeValidator().Validate

	return v
}()

func AzureSubnetChoiceValidator() db.Validator {
	return DefaultAzureSubnetChoiceValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureSubnetParamType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureSubnetParamType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureSubnetParamType) DeepCopy() *AzureSubnetParamType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureSubnetParamType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureSubnetParamType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureSubnetParamType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureSubnetParamTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureSubnetParamType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureSubnetParamType) ResourceGroupValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for resource_group")
	}

	return validatorFn, nil
}

func (v *ValidateAzureSubnetParamType) Ipv4ValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ipv4")
	}

	return validatorFn, nil
}

func (v *ValidateAzureSubnetParamType) Ipv6ValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ipv6")
	}

	return validatorFn, nil
}

func (v *ValidateAzureSubnetParamType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureSubnetParamType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureSubnetParamType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ipv4"]; exists {

		vOpts := append(opts, db.WithValidateField("ipv4"))
		if err := fv(ctx, m.GetIpv4(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ipv6"]; exists {

		vOpts := append(opts, db.WithValidateField("ipv6"))
		if err := fv(ctx, m.GetIpv6(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["resource_group"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_group"))
		if err := fv(ctx, m.GetResourceGroup(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureSubnetParamTypeValidator = func() *ValidateAzureSubnetParamType {
	v := &ValidateAzureSubnetParamType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhResourceGroup := v.ResourceGroupValidationRuleHandler
	rulesResourceGroup := map[string]string{
		"ves.io.schema.rules.string.max_len": "64",
	}
	vFn, err = vrhResourceGroup(rulesResourceGroup)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AzureSubnetParamType.resource_group: %s", err)
		panic(errMsg)
	}
	v.FldValidators["resource_group"] = vFn

	vrhIpv4 := v.Ipv4ValidationRuleHandler
	rulesIpv4 := map[string]string{
		"ves.io.schema.rules.string.ipv4_prefix": "true",
	}
	vFn, err = vrhIpv4(rulesIpv4)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AzureSubnetParamType.ipv4: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ipv4"] = vFn

	vrhIpv6 := v.Ipv6ValidationRuleHandler
	rulesIpv6 := map[string]string{
		"ves.io.schema.rules.string.ipv6_prefix": "true",
	}
	vFn, err = vrhIpv6(rulesIpv6)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AzureSubnetParamType.ipv6: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ipv6"] = vFn

	return v
}()

func AzureSubnetParamTypeValidator() db.Validator {
	return DefaultAzureSubnetParamTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureSubnetType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureSubnetType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureSubnetType) DeepCopy() *AzureSubnetType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureSubnetType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureSubnetType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureSubnetType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureSubnetTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureSubnetType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureSubnetType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureSubnetType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureSubnetType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["interface_type"]; exists {

		vOpts := append(opts, db.WithValidateField("interface_type"))
		if err := fv(ctx, m.GetInterfaceType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subnet"]; exists {

		vOpts := append(opts, db.WithValidateField("subnet"))
		if err := fv(ctx, m.GetSubnet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volt_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("volt_subnet_id"))
		if err := fv(ctx, m.GetVoltSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volt_vnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("volt_vnet_id"))
		if err := fv(ctx, m.GetVoltVnetId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureSubnetTypeValidator = func() *ValidateAzureSubnetType {
	v := &ValidateAzureSubnetType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["subnet"] = AzureSubnetChoiceValidator().Validate

	return v
}()

func AzureSubnetTypeValidator() db.Validator {
	return DefaultAzureSubnetTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureVnetInfoType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureVnetInfoType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureVnetInfoType) DeepCopy() *AzureVnetInfoType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureVnetInfoType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureVnetInfoType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureVnetInfoType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureVnetInfoTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureVnetInfoType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureVnetInfoType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureVnetInfoType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureVnetInfoType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["new_vnet"]; exists {

		vOpts := append(opts, db.WithValidateField("new_vnet"))
		if err := fv(ctx, m.GetNewVnet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vnet"]; exists {

		vOpts := append(opts, db.WithValidateField("vnet"))
		if err := fv(ctx, m.GetVnet(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureVnetInfoTypeValidator = func() *ValidateAzureVnetInfoType {
	v := &ValidateAzureVnetInfoType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["new_vnet"] = ves_io_schema_views.AzureVnetParamsTypeValidator().Validate

	v.FldValidators["vnet"] = ves_io_schema_views.AzureVnetTypeValidator().Validate

	return v
}()

func AzureVnetInfoTypeValidator() db.Validator {
	return DefaultAzureVnetInfoTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureVnetParamsType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureVnetParamsType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureVnetParamsType) DeepCopy() *AzureVnetParamsType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureVnetParamsType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureVnetParamsType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureVnetParamsType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureVnetParamsTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureVnetParamsType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureVnetParamsType) ResourceGroupValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for resource_group")
	}

	return validatorFn, nil
}

func (v *ValidateAzureVnetParamsType) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateAzureVnetParamsType) PrimaryIpv4ValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for primary_ipv4")
	}

	return validatorFn, nil
}

func (v *ValidateAzureVnetParamsType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureVnetParamsType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureVnetParamsType got type %s", t)
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

	if fv, exists := v.FldValidators["primary_ipv4"]; exists {

		vOpts := append(opts, db.WithValidateField("primary_ipv4"))
		if err := fv(ctx, m.GetPrimaryIpv4(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["resource_group"]; exists {

		vOpts := append(opts, db.WithValidateField("resource_group"))
		if err := fv(ctx, m.GetResourceGroup(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureVnetParamsTypeValidator = func() *ValidateAzureVnetParamsType {
	v := &ValidateAzureVnetParamsType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhResourceGroup := v.ResourceGroupValidationRuleHandler
	rulesResourceGroup := map[string]string{
		"ves.io.schema.rules.string.max_len": "64",
	}
	vFn, err = vrhResourceGroup(rulesResourceGroup)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AzureVnetParamsType.resource_group: %s", err)
		panic(errMsg)
	}
	v.FldValidators["resource_group"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.string.max_len": "64",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AzureVnetParamsType.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhPrimaryIpv4 := v.PrimaryIpv4ValidationRuleHandler
	rulesPrimaryIpv4 := map[string]string{
		"ves.io.schema.rules.string.ipv4_prefix": "true",
	}
	vFn, err = vrhPrimaryIpv4(rulesPrimaryIpv4)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AzureVnetParamsType.primary_ipv4: %s", err)
		panic(errMsg)
	}
	v.FldValidators["primary_ipv4"] = vFn

	return v
}()

func AzureVnetParamsTypeValidator() db.Validator {
	return DefaultAzureVnetParamsTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *AzureVnetSiteType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AzureVnetSiteType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AzureVnetSiteType) DeepCopy() *AzureVnetSiteType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AzureVnetSiteType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AzureVnetSiteType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AzureVnetSiteType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AzureVnetSiteTypeValidator().Validate(ctx, m, opts...)
}

type ValidateAzureVnetSiteType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAzureVnetSiteType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AzureVnetSiteType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AzureVnetSiteType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["availability_set"]; exists {

		vOpts := append(opts, db.WithValidateField("availability_set"))
		if err := fv(ctx, m.GetAvailabilitySet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["azure_name"]; exists {

		vOpts := append(opts, db.WithValidateField("azure_name"))
		if err := fv(ctx, m.GetAzureName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["azure_region"]; exists {

		vOpts := append(opts, db.WithValidateField("azure_region"))
		if err := fv(ctx, m.GetAzureRegion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["azure_resource_grp"]; exists {

		vOpts := append(opts, db.WithValidateField("azure_resource_grp"))
		if err := fv(ctx, m.GetAzureResourceGrp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["certified_hw"]; exists {

		vOpts := append(opts, db.WithValidateField("certified_hw"))
		if err := fv(ctx, m.GetCertifiedHw(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["fleet_label"]; exists {

		vOpts := append(opts, db.WithValidateField("fleet_label"))
		if err := fv(ctx, m.GetFleetLabel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["gateway_type"]; exists {

		vOpts := append(opts, db.WithValidateField("gateway_type"))
		if err := fv(ctx, m.GetGatewayType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["inside_vip_port_config"]; exists {

		vOpts := append(opts, db.WithValidateField("inside_vip_port_config"))
		for idx, item := range m.GetInsideVipPortConfig() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["master_nodes"]; exists {

		vOpts := append(opts, db.WithValidateField("master_nodes"))
		for idx, item := range m.GetMasterNodes() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["outside_vip_port_config"]; exists {

		vOpts := append(opts, db.WithValidateField("outside_vip_port_config"))
		for idx, item := range m.GetOutsideVipPortConfig() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("site_name"))
		if err := fv(ctx, m.GetSiteName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ssh_key"]; exists {

		vOpts := append(opts, db.WithValidateField("ssh_key"))
		if err := fv(ctx, m.GetSshKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subnets"]; exists {

		vOpts := append(opts, db.WithValidateField("subnets"))
		for idx, item := range m.GetSubnets() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["tags"]; exists {

		vOpts := append(opts, db.WithValidateField("tags"))
		for key, value := range m.GetTags() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["vnet"]; exists {

		vOpts := append(opts, db.WithValidateField("vnet"))
		if err := fv(ctx, m.GetVnet(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volt_vnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("volt_vnet_id"))
		if err := fv(ctx, m.GetVoltVnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["worker_nodes"]; exists {

		vOpts := append(opts, db.WithValidateField("worker_nodes"))
		if err := fv(ctx, m.GetWorkerNodes(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAzureVnetSiteTypeValidator = func() *ValidateAzureVnetSiteType {
	v := &ValidateAzureVnetSiteType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["vnet"] = AzureVnetInfoTypeValidator().Validate

	v.FldValidators["subnets"] = AzureSubnetTypeValidator().Validate

	v.FldValidators["inside_vip_port_config"] = VIPPortConfigValidator().Validate

	v.FldValidators["outside_vip_port_config"] = VIPPortConfigValidator().Validate

	v.FldValidators["availability_set"] = AvailabilitySetsInfoTypeValidator().Validate

	return v
}()

func AzureVnetSiteTypeValidator() db.Validator {
	return DefaultAzureVnetSiteTypeValidator
}
