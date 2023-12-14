// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cminstance

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

// Redact squashes sensitive info in m (in-place)
func (m *CreateSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting CreateSpecType.password")
	}

	if err := m.GetApiToken().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting CreateSpecType.api_token")
	}

	return nil
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

func (v *ValidateCreateSpecType) IpValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.Ipv4AddressTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) UsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for username")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) ApiTokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for api_token")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["api_token"]; exists {

		vOpts := append(opts, db.WithValidateField("api_token"))
		if err := fv(ctx, m.GetApiToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ip"]; exists {

		vOpts := append(opts, db.WithValidateField("ip"))
		if err := fv(ctx, m.GetIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["username"]; exists {

		vOpts := append(opts, db.WithValidateField("username"))
		if err := fv(ctx, m.GetUsername(), vOpts...); err != nil {
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

	vrhIp := v.IpValidationRuleHandler
	rulesIp := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIp(rulesIp)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.ip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gte":       "1",
		"ves.io.schema.rules.uint32.lte":       "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	vrhUsername := v.UsernameValidationRuleHandler
	rulesUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
		"ves.io.schema.rules.string.min_len":   "4",
	}
	vFn, err = vrhUsername(rulesUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["username"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhApiToken := v.ApiTokenValidationRuleHandler
	rulesApiToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhApiToken(rulesApiToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.api_token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_token"] = vFn

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

// Redact squashes sensitive info in m (in-place)
func (m *GetSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GetSpecType.password")
	}

	if err := m.GetApiToken().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GetSpecType.api_token")
	}

	return nil
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

func (v *ValidateGetSpecType) IpValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.Ipv4AddressTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) UsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for username")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) ApiTokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for api_token")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["api_token"]; exists {

		vOpts := append(opts, db.WithValidateField("api_token"))
		if err := fv(ctx, m.GetApiToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ip"]; exists {

		vOpts := append(opts, db.WithValidateField("ip"))
		if err := fv(ctx, m.GetIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["username"]; exists {

		vOpts := append(opts, db.WithValidateField("username"))
		if err := fv(ctx, m.GetUsername(), vOpts...); err != nil {
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

	vrhIp := v.IpValidationRuleHandler
	rulesIp := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIp(rulesIp)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.ip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gte":       "1",
		"ves.io.schema.rules.uint32.lte":       "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	vrhUsername := v.UsernameValidationRuleHandler
	rulesUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
		"ves.io.schema.rules.string.min_len":   "4",
	}
	vFn, err = vrhUsername(rulesUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["username"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhApiToken := v.ApiTokenValidationRuleHandler
	rulesApiToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhApiToken(rulesApiToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.api_token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_token"] = vFn

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

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.password")
	}

	if err := m.GetApiToken().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.api_token")
	}

	return nil
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

func (v *ValidateGlobalSpecType) IpValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.Ipv4AddressTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) UsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for username")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ApiTokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for api_token")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["api_token"]; exists {

		vOpts := append(opts, db.WithValidateField("api_token"))
		if err := fv(ctx, m.GetApiToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ip"]; exists {

		vOpts := append(opts, db.WithValidateField("ip"))
		if err := fv(ctx, m.GetIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["username"]; exists {

		vOpts := append(opts, db.WithValidateField("username"))
		if err := fv(ctx, m.GetUsername(), vOpts...); err != nil {
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

	vrhIp := v.IpValidationRuleHandler
	rulesIp := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIp(rulesIp)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.ip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gte":       "1",
		"ves.io.schema.rules.uint32.lte":       "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	vrhUsername := v.UsernameValidationRuleHandler
	rulesUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
		"ves.io.schema.rules.string.min_len":   "4",
	}
	vFn, err = vrhUsername(rulesUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["username"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhApiToken := v.ApiTokenValidationRuleHandler
	rulesApiToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhApiToken(rulesApiToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.api_token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_token"] = vFn

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

// Redact squashes sensitive info in m (in-place)
func (m *ReplaceSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting ReplaceSpecType.password")
	}

	if err := m.GetApiToken().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting ReplaceSpecType.api_token")
	}

	return nil
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

func (v *ValidateReplaceSpecType) IpValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for ip")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.Ipv4AddressTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) UsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for username")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) ApiTokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for api_token")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
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

	if fv, exists := v.FldValidators["api_token"]; exists {

		vOpts := append(opts, db.WithValidateField("api_token"))
		if err := fv(ctx, m.GetApiToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ip"]; exists {

		vOpts := append(opts, db.WithValidateField("ip"))
		if err := fv(ctx, m.GetIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["username"]; exists {

		vOpts := append(opts, db.WithValidateField("username"))
		if err := fv(ctx, m.GetUsername(), vOpts...); err != nil {
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

	vrhIp := v.IpValidationRuleHandler
	rulesIp := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIp(rulesIp)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.ip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ip"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gte":       "1",
		"ves.io.schema.rules.uint32.lte":       "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	vrhUsername := v.UsernameValidationRuleHandler
	rulesUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
		"ves.io.schema.rules.string.min_len":   "4",
	}
	vFn, err = vrhUsername(rulesUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["username"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhApiToken := v.ApiTokenValidationRuleHandler
	rulesApiToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhApiToken(rulesApiToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.api_token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_token"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.ApiToken = f.GetApiToken()
	m.Ip = f.GetIp()
	m.Password = f.GetPassword()
	m.Port = f.GetPort()
	m.Username = f.GetUsername()
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

	f.ApiToken = m1.ApiToken
	f.Ip = m1.Ip
	f.Password = m1.Password
	f.Port = m1.Port
	f.Username = m1.Username
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
	m.ApiToken = f.GetApiToken()
	m.Ip = f.GetIp()
	m.Password = f.GetPassword()
	m.Port = f.GetPort()
	m.Username = f.GetUsername()
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

	f.ApiToken = m1.ApiToken
	f.Ip = m1.Ip
	f.Password = m1.Password
	f.Port = m1.Port
	f.Username = m1.Username
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
	m.ApiToken = f.GetApiToken()
	m.Ip = f.GetIp()
	m.Password = f.GetPassword()
	m.Port = f.GetPort()
	m.Username = f.GetUsername()
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

	f.ApiToken = m1.ApiToken
	f.Ip = m1.Ip
	f.Password = m1.Password
	f.Port = m1.Port
	f.Username = m1.Username
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
