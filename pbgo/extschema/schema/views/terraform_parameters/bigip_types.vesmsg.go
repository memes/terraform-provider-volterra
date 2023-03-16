//
// Copyright (c) 2022 F5, Inc. All rights reserved.
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

func (m *BigIPAWSType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *BigIPAWSType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *BigIPAWSType) DeepCopy() *BigIPAWSType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &BigIPAWSType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *BigIPAWSType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *BigIPAWSType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return BigIPAWSTypeValidator().Validate(ctx, m, opts...)
}

type ValidateBigIPAWSType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateBigIPAWSType) AwsRegionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for aws_region")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) AwsNamePrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for aws_name_prefix")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) VpcIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for vpc_id")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) AdminUsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for admin_username")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) VipAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for vip_address")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) DevicesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for devices")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*BigIPDeviceType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := BigIPDeviceTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for devices")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*BigIPDeviceType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*BigIPDeviceType, got %T", val)
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
			return errors.Wrap(err, "repeated devices")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items devices")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) SshKeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ssh_key")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) VolterraSubnetIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for volterra_subnet_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_views.AWSSubnetIdsType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema_views.AWSSubnetIdsTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for volterra_subnet_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_views.AWSSubnetIdsType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_views.AWSSubnetIdsType, got %T", val)
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
			return errors.Wrap(err, "repeated volterra_subnet_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items volterra_subnet_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateBigIPAWSType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*BigIPAWSType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *BigIPAWSType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["admin_username"]; exists {

		vOpts := append(opts, db.WithValidateField("admin_username"))
		if err := fv(ctx, m.GetAdminUsername(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["automatic_vip"]; exists {

		vOpts := append(opts, db.WithValidateField("automatic_vip"))
		if err := fv(ctx, m.GetAutomaticVip(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["aws_name_prefix"]; exists {

		vOpts := append(opts, db.WithValidateField("aws_name_prefix"))
		if err := fv(ctx, m.GetAwsNamePrefix(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["aws_region"]; exists {

		vOpts := append(opts, db.WithValidateField("aws_region"))
		if err := fv(ctx, m.GetAwsRegion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["devices"]; exists {
		vOpts := append(opts, db.WithValidateField("devices"))
		if err := fv(ctx, m.GetDevices(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mp_ami_image"]; exists {

		vOpts := append(opts, db.WithValidateField("mp_ami_image"))
		if err := fv(ctx, m.GetMpAmiImage(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ssh_key"]; exists {

		vOpts := append(opts, db.WithValidateField("ssh_key"))
		if err := fv(ctx, m.GetSshKey(), vOpts...); err != nil {
			return err
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

	if fv, exists := v.FldValidators["vip_address"]; exists {

		vOpts := append(opts, db.WithValidateField("vip_address"))
		if err := fv(ctx, m.GetVipAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volterra_subnet_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("volterra_subnet_ids"))
		if err := fv(ctx, m.GetVolterraSubnetIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vpc_id"]; exists {

		vOpts := append(opts, db.WithValidateField("vpc_id"))
		if err := fv(ctx, m.GetVpcId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultBigIPAWSTypeValidator = func() *ValidateBigIPAWSType {
	v := &ValidateBigIPAWSType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAwsRegion := v.AwsRegionValidationRuleHandler
	rulesAwsRegion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAwsRegion(rulesAwsRegion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.aws_region: %s", err)
		panic(errMsg)
	}
	v.FldValidators["aws_region"] = vFn

	vrhAwsNamePrefix := v.AwsNamePrefixValidationRuleHandler
	rulesAwsNamePrefix := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAwsNamePrefix(rulesAwsNamePrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.aws_name_prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["aws_name_prefix"] = vFn

	vrhVpcId := v.VpcIdValidationRuleHandler
	rulesVpcId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVpcId(rulesVpcId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.vpc_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vpc_id"] = vFn

	vrhAdminUsername := v.AdminUsernameValidationRuleHandler
	rulesAdminUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAdminUsername(rulesAdminUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.admin_username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["admin_username"] = vFn

	vrhVipAddress := v.VipAddressValidationRuleHandler
	rulesVipAddress := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVipAddress(rulesVipAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.vip_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vip_address"] = vFn

	vrhDevices := v.DevicesValidationRuleHandler
	rulesDevices := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "2",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhDevices(rulesDevices)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.devices: %s", err)
		panic(errMsg)
	}
	v.FldValidators["devices"] = vFn

	vrhSshKey := v.SshKeyValidationRuleHandler
	rulesSshKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSshKey(rulesSshKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.ssh_key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ssh_key"] = vFn

	vrhVolterraSubnetIds := v.VolterraSubnetIdsValidationRuleHandler
	rulesVolterraSubnetIds := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVolterraSubnetIds(rulesVolterraSubnetIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPAWSType.volterra_subnet_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["volterra_subnet_ids"] = vFn

	return v
}()

func BigIPAWSTypeValidator() db.Validator {
	return DefaultBigIPAWSTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *BigIPDeviceType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *BigIPDeviceType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *BigIPDeviceType) DeepCopy() *BigIPDeviceType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &BigIPDeviceType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *BigIPDeviceType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *BigIPDeviceType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return BigIPDeviceTypeValidator().Validate(ctx, m, opts...)
}

type ValidateBigIPDeviceType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateBigIPDeviceType) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPDeviceType) InstanceTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for instance_type")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPDeviceType) PrivateSubnetIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for private_subnet_id")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPDeviceType) WorkloadSubnetIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for workload_subnet_id")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPDeviceType) MgmtSubnetCidrValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mgmt_subnet_cidr")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPDeviceType) MgmtSubnetIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mgmt_subnet_id")
	}

	return validatorFn, nil
}

func (v *ValidateBigIPDeviceType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*BigIPDeviceType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *BigIPDeviceType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["instance_type"]; exists {

		vOpts := append(opts, db.WithValidateField("instance_type"))
		if err := fv(ctx, m.GetInstanceType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mgmt_subnet_cidr"]; exists {

		vOpts := append(opts, db.WithValidateField("mgmt_subnet_cidr"))
		if err := fv(ctx, m.GetMgmtSubnetCidr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mgmt_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("mgmt_subnet_id"))
		if err := fv(ctx, m.GetMgmtSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["private_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("private_subnet_id"))
		if err := fv(ctx, m.GetPrivateSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["workload_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("workload_subnet_id"))
		if err := fv(ctx, m.GetWorkloadSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultBigIPDeviceTypeValidator = func() *ValidateBigIPDeviceType {
	v := &ValidateBigIPDeviceType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPDeviceType.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhInstanceType := v.InstanceTypeValidationRuleHandler
	rulesInstanceType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInstanceType(rulesInstanceType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPDeviceType.instance_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["instance_type"] = vFn

	vrhPrivateSubnetId := v.PrivateSubnetIdValidationRuleHandler
	rulesPrivateSubnetId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPrivateSubnetId(rulesPrivateSubnetId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPDeviceType.private_subnet_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["private_subnet_id"] = vFn

	vrhWorkloadSubnetId := v.WorkloadSubnetIdValidationRuleHandler
	rulesWorkloadSubnetId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhWorkloadSubnetId(rulesWorkloadSubnetId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPDeviceType.workload_subnet_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["workload_subnet_id"] = vFn

	vrhMgmtSubnetCidr := v.MgmtSubnetCidrValidationRuleHandler
	rulesMgmtSubnetCidr := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMgmtSubnetCidr(rulesMgmtSubnetCidr)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPDeviceType.mgmt_subnet_cidr: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mgmt_subnet_cidr"] = vFn

	vrhMgmtSubnetId := v.MgmtSubnetIdValidationRuleHandler
	rulesMgmtSubnetId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMgmtSubnetId(rulesMgmtSubnetId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIPDeviceType.mgmt_subnet_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mgmt_subnet_id"] = vFn

	return v
}()

func BigIPDeviceTypeValidator() db.Validator {
	return DefaultBigIPDeviceTypeValidator
}
