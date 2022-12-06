// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package schema

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

func (m *AwsVpcList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AwsVpcList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AwsVpcList) DeepCopy() *AwsVpcList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AwsVpcList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AwsVpcList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AwsVpcList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AwsVpcListValidator().Validate(ctx, m, opts...)
}

type ValidateAwsVpcList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAwsVpcList) VpcIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for vpc_id")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vpc_id")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated vpc_id")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vpc_id")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateAwsVpcList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AwsVpcList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AwsVpcList got type %s", t)
		}
	}
	if m == nil {
		return nil
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
var DefaultAwsVpcListValidator = func() *ValidateAwsVpcList {
	v := &ValidateAwsVpcList{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhVpcId := v.VpcIdValidationRuleHandler
	rulesVpcId := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "256",
		"ves.io.schema.rules.repeated.unique":    "true",
		"ves.io.schema.rules.string.max_len":     "64",
		"ves.io.schema.rules.string.pattern":     "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
	}
	vFn, err = vrhVpcId(rulesVpcId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AwsVpcList.vpc_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vpc_id"] = vFn

	return v
}()

func AwsVpcListValidator() db.Validator {
	return DefaultAwsVpcListValidator
}

// augmented methods on protoc/std generated struct

func (m *IpAddressType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IpAddressType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IpAddressType) DeepCopy() *IpAddressType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IpAddressType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IpAddressType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IpAddressType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IpAddressTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIpAddressType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpAddressType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IpAddressType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IpAddressType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetVer().(type) {
	case *IpAddressType_Ipv4:
		if fv, exists := v.FldValidators["ver.ipv4"]; exists {
			val := m.GetVer().(*IpAddressType_Ipv4).Ipv4
			vOpts := append(opts,
				db.WithValidateField("ver"),
				db.WithValidateField("ipv4"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *IpAddressType_Ipv6:
		if fv, exists := v.FldValidators["ver.ipv6"]; exists {
			val := m.GetVer().(*IpAddressType_Ipv6).Ipv6
			vOpts := append(opts,
				db.WithValidateField("ver"),
				db.WithValidateField("ipv6"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpAddressTypeValidator = func() *ValidateIpAddressType {
	v := &ValidateIpAddressType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["ver.ipv4"] = Ipv4AddressTypeValidator().Validate
	v.FldValidators["ver.ipv6"] = Ipv6AddressTypeValidator().Validate

	return v
}()

func IpAddressTypeValidator() db.Validator {
	return DefaultIpAddressTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *IpSubnetType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IpSubnetType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IpSubnetType) DeepCopy() *IpSubnetType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IpSubnetType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IpSubnetType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IpSubnetType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IpSubnetTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIpSubnetType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpSubnetType) VerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ver")
	}
	return validatorFn, nil
}

func (v *ValidateIpSubnetType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IpSubnetType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IpSubnetType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ver"]; exists {
		val := m.GetVer()
		vOpts := append(opts,
			db.WithValidateField("ver"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetVer().(type) {
	case *IpSubnetType_Ipv4:
		if fv, exists := v.FldValidators["ver.ipv4"]; exists {
			val := m.GetVer().(*IpSubnetType_Ipv4).Ipv4
			vOpts := append(opts,
				db.WithValidateField("ver"),
				db.WithValidateField("ipv4"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *IpSubnetType_Ipv6:
		if fv, exists := v.FldValidators["ver.ipv6"]; exists {
			val := m.GetVer().(*IpSubnetType_Ipv6).Ipv6
			vOpts := append(opts,
				db.WithValidateField("ver"),
				db.WithValidateField("ipv6"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpSubnetTypeValidator = func() *ValidateIpSubnetType {
	v := &ValidateIpSubnetType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhVer := v.VerValidationRuleHandler
	rulesVer := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhVer(rulesVer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for IpSubnetType.ver: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ver"] = vFn

	v.FldValidators["ver.ipv4"] = Ipv4SubnetTypeValidator().Validate
	v.FldValidators["ver.ipv6"] = Ipv6SubnetTypeValidator().Validate

	return v
}()

func IpSubnetTypeValidator() db.Validator {
	return DefaultIpSubnetTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Ipv4AddressType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Ipv4AddressType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Ipv4AddressType) DeepCopy() *Ipv4AddressType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Ipv4AddressType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Ipv4AddressType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Ipv4AddressType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return Ipv4AddressTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIpv4AddressType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpv4AddressType) AddrValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for addr")
	}

	return validatorFn, nil
}

func (v *ValidateIpv4AddressType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Ipv4AddressType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Ipv4AddressType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["addr"]; exists {

		vOpts := append(opts, db.WithValidateField("addr"))
		if err := fv(ctx, m.GetAddr(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpv4AddressTypeValidator = func() *ValidateIpv4AddressType {
	v := &ValidateIpv4AddressType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAddr := v.AddrValidationRuleHandler
	rulesAddr := map[string]string{
		"ves.io.schema.rules.string.ipv4": "true",
	}
	vFn, err = vrhAddr(rulesAddr)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Ipv4AddressType.addr: %s", err)
		panic(errMsg)
	}
	v.FldValidators["addr"] = vFn

	return v
}()

func Ipv4AddressTypeValidator() db.Validator {
	return DefaultIpv4AddressTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Ipv4SubnetType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Ipv4SubnetType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Ipv4SubnetType) DeepCopy() *Ipv4SubnetType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Ipv4SubnetType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Ipv4SubnetType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Ipv4SubnetType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return Ipv4SubnetTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIpv4SubnetType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpv4SubnetType) PrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for prefix")
	}

	return validatorFn, nil
}

func (v *ValidateIpv4SubnetType) PlenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for plen")
	}

	return validatorFn, nil
}

func (v *ValidateIpv4SubnetType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Ipv4SubnetType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Ipv4SubnetType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["plen"]; exists {

		vOpts := append(opts, db.WithValidateField("plen"))
		if err := fv(ctx, m.GetPlen(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["prefix"]; exists {

		vOpts := append(opts, db.WithValidateField("prefix"))
		if err := fv(ctx, m.GetPrefix(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpv4SubnetTypeValidator = func() *ValidateIpv4SubnetType {
	v := &ValidateIpv4SubnetType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPrefix := v.PrefixValidationRuleHandler
	rulesPrefix := map[string]string{
		"ves.io.schema.rules.string.ipv4": "true",
	}
	vFn, err = vrhPrefix(rulesPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Ipv4SubnetType.prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["prefix"] = vFn

	vrhPlen := v.PlenValidationRuleHandler
	rulesPlen := map[string]string{
		"ves.io.schema.rules.uint32.lte": "32",
	}
	vFn, err = vrhPlen(rulesPlen)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Ipv4SubnetType.plen: %s", err)
		panic(errMsg)
	}
	v.FldValidators["plen"] = vFn

	return v
}()

func Ipv4SubnetTypeValidator() db.Validator {
	return DefaultIpv4SubnetTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Ipv6AddressType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Ipv6AddressType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Ipv6AddressType) DeepCopy() *Ipv6AddressType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Ipv6AddressType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Ipv6AddressType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Ipv6AddressType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return Ipv6AddressTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIpv6AddressType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpv6AddressType) AddrValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for addr")
	}

	return validatorFn, nil
}

func (v *ValidateIpv6AddressType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Ipv6AddressType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Ipv6AddressType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["addr"]; exists {

		vOpts := append(opts, db.WithValidateField("addr"))
		if err := fv(ctx, m.GetAddr(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpv6AddressTypeValidator = func() *ValidateIpv6AddressType {
	v := &ValidateIpv6AddressType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAddr := v.AddrValidationRuleHandler
	rulesAddr := map[string]string{
		"ves.io.schema.rules.string.ipv6": "true",
	}
	vFn, err = vrhAddr(rulesAddr)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Ipv6AddressType.addr: %s", err)
		panic(errMsg)
	}
	v.FldValidators["addr"] = vFn

	return v
}()

func Ipv6AddressTypeValidator() db.Validator {
	return DefaultIpv6AddressTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Ipv6SubnetType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Ipv6SubnetType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Ipv6SubnetType) DeepCopy() *Ipv6SubnetType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Ipv6SubnetType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Ipv6SubnetType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Ipv6SubnetType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return Ipv6SubnetTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIpv6SubnetType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpv6SubnetType) PrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for prefix")
	}

	return validatorFn, nil
}

func (v *ValidateIpv6SubnetType) PlenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for plen")
	}

	return validatorFn, nil
}

func (v *ValidateIpv6SubnetType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Ipv6SubnetType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Ipv6SubnetType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["plen"]; exists {

		vOpts := append(opts, db.WithValidateField("plen"))
		if err := fv(ctx, m.GetPlen(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["prefix"]; exists {

		vOpts := append(opts, db.WithValidateField("prefix"))
		if err := fv(ctx, m.GetPrefix(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpv6SubnetTypeValidator = func() *ValidateIpv6SubnetType {
	v := &ValidateIpv6SubnetType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPrefix := v.PrefixValidationRuleHandler
	rulesPrefix := map[string]string{
		"ves.io.schema.rules.string.ipv6": "true",
	}
	vFn, err = vrhPrefix(rulesPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Ipv6SubnetType.prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["prefix"] = vFn

	vrhPlen := v.PlenValidationRuleHandler
	rulesPlen := map[string]string{
		"ves.io.schema.rules.uint32.lte": "128",
	}
	vFn, err = vrhPlen(rulesPlen)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Ipv6SubnetType.plen: %s", err)
		panic(errMsg)
	}
	v.FldValidators["plen"] = vFn

	return v
}()

func Ipv6SubnetTypeValidator() db.Validator {
	return DefaultIpv6SubnetTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *MacAddressType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *MacAddressType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *MacAddressType) DeepCopy() *MacAddressType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &MacAddressType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *MacAddressType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *MacAddressType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return MacAddressTypeValidator().Validate(ctx, m, opts...)
}

type ValidateMacAddressType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateMacAddressType) MacValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mac")
	}

	return validatorFn, nil
}

func (v *ValidateMacAddressType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*MacAddressType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *MacAddressType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["mac"]; exists {

		vOpts := append(opts, db.WithValidateField("mac"))
		if err := fv(ctx, m.GetMac(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultMacAddressTypeValidator = func() *ValidateMacAddressType {
	v := &ValidateMacAddressType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMac := v.MacValidationRuleHandler
	rulesMac := map[string]string{
		"ves.io.schema.rules.string.mac": "true",
	}
	vFn, err = vrhMac(rulesMac)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for MacAddressType.mac: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mac"] = vFn

	return v
}()

func MacAddressTypeValidator() db.Validator {
	return DefaultMacAddressTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *PrefixListType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *PrefixListType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *PrefixListType) DeepCopy() *PrefixListType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &PrefixListType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *PrefixListType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *PrefixListType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PrefixListTypeValidator().Validate(ctx, m, opts...)
}

type ValidatePrefixListType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePrefixListType) PrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for prefix")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for prefix")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated prefix")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items prefix")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidatePrefixListType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*PrefixListType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *PrefixListType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["prefix"]; exists {
		vOpts := append(opts, db.WithValidateField("prefix"))
		if err := fv(ctx, m.GetPrefix(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPrefixListTypeValidator = func() *ValidatePrefixListType {
	v := &ValidatePrefixListType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPrefix := v.PrefixValidationRuleHandler
	rulesPrefix := map[string]string{
		"ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
		"ves.io.schema.rules.repeated.max_items":                "256",
	}
	vFn, err = vrhPrefix(rulesPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PrefixListType.prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["prefix"] = vFn

	return v
}()

func PrefixListTypeValidator() db.Validator {
	return DefaultPrefixListTypeValidator
}
