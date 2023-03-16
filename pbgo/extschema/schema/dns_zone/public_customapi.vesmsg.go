//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package dns_zone

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

func (m *F5CSDNSZoneConfiguration) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *F5CSDNSZoneConfiguration) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *F5CSDNSZoneConfiguration) DeepCopy() *F5CSDNSZoneConfiguration {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &F5CSDNSZoneConfiguration{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *F5CSDNSZoneConfiguration) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *F5CSDNSZoneConfiguration) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return F5CSDNSZoneConfigurationValidator().Validate(ctx, m, opts...)
}

type ValidateF5CSDNSZoneConfiguration struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateF5CSDNSZoneConfiguration) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*F5CSDNSZoneConfiguration)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *F5CSDNSZoneConfiguration got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetDnsType().(type) {
	case *F5CSDNSZoneConfiguration_DnsService:
		if fv, exists := v.FldValidators["dns_type.dns_service"]; exists {
			val := m.GetDnsType().(*F5CSDNSZoneConfiguration_DnsService).DnsService
			vOpts := append(opts,
				db.WithValidateField("dns_type"),
				db.WithValidateField("dns_service"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *F5CSDNSZoneConfiguration_AdnsService:
		if fv, exists := v.FldValidators["dns_type.adns_service"]; exists {
			val := m.GetDnsType().(*F5CSDNSZoneConfiguration_AdnsService).AdnsService
			vOpts := append(opts,
				db.WithValidateField("dns_type"),
				db.WithValidateField("adns_service"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultF5CSDNSZoneConfigurationValidator = func() *ValidateF5CSDNSZoneConfiguration {
	v := &ValidateF5CSDNSZoneConfiguration{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func F5CSDNSZoneConfigurationValidator() db.Validator {
	return DefaultF5CSDNSZoneConfigurationValidator
}

// augmented methods on protoc/std generated struct

func (m *GetLocalZoneFileRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetLocalZoneFileRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetLocalZoneFileRequest) DeepCopy() *GetLocalZoneFileRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetLocalZoneFileRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetLocalZoneFileRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetLocalZoneFileRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetLocalZoneFileRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetLocalZoneFileRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetLocalZoneFileRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetLocalZoneFileRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetLocalZoneFileRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_zone_name"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_zone_name"))
		if err := fv(ctx, m.GetDnsZoneName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetLocalZoneFileRequestValidator = func() *ValidateGetLocalZoneFileRequest {
	v := &ValidateGetLocalZoneFileRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetLocalZoneFileRequestValidator() db.Validator {
	return DefaultGetLocalZoneFileRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetLocalZoneFileResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetLocalZoneFileResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetLocalZoneFileResponse) DeepCopy() *GetLocalZoneFileResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetLocalZoneFileResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetLocalZoneFileResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetLocalZoneFileResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetLocalZoneFileResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetLocalZoneFileResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetLocalZoneFileResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetLocalZoneFileResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetLocalZoneFileResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["last_axfr_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("last_axfr_timestamp"))
		if err := fv(ctx, m.GetLastAxfrTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["zone_file"]; exists {

		vOpts := append(opts, db.WithValidateField("zone_file"))
		if err := fv(ctx, m.GetZoneFile(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetLocalZoneFileResponseValidator = func() *ValidateGetLocalZoneFileResponse {
	v := &ValidateGetLocalZoneFileResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetLocalZoneFileResponseValidator() db.Validator {
	return DefaultGetLocalZoneFileResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetRemoteZoneFileRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetRemoteZoneFileRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetRemoteZoneFileRequest) DeepCopy() *GetRemoteZoneFileRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetRemoteZoneFileRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetRemoteZoneFileRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetRemoteZoneFileRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetRemoteZoneFileRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetRemoteZoneFileRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetRemoteZoneFileRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetRemoteZoneFileRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetRemoteZoneFileRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_zone_name"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_zone_name"))
		if err := fv(ctx, m.GetDnsZoneName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetRemoteZoneFileRequestValidator = func() *ValidateGetRemoteZoneFileRequest {
	v := &ValidateGetRemoteZoneFileRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetRemoteZoneFileRequestValidator() db.Validator {
	return DefaultGetRemoteZoneFileRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetRemoteZoneFileResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetRemoteZoneFileResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetRemoteZoneFileResponse) DeepCopy() *GetRemoteZoneFileResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetRemoteZoneFileResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetRemoteZoneFileResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetRemoteZoneFileResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetRemoteZoneFileResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetRemoteZoneFileResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetRemoteZoneFileResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetRemoteZoneFileResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetRemoteZoneFileResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["zone_file"]; exists {

		vOpts := append(opts, db.WithValidateField("zone_file"))
		if err := fv(ctx, m.GetZoneFile(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetRemoteZoneFileResponseValidator = func() *ValidateGetRemoteZoneFileResponse {
	v := &ValidateGetRemoteZoneFileResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetRemoteZoneFileResponseValidator() db.Validator {
	return DefaultGetRemoteZoneFileResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ImportAXFRRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ImportAXFRRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *ImportAXFRRequest) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetTsigKeyValue().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting ImportAXFRRequest.tsig_key_value")
	}

	return nil
}

func (m *ImportAXFRRequest) DeepCopy() *ImportAXFRRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ImportAXFRRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ImportAXFRRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ImportAXFRRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ImportAXFRRequestValidator().Validate(ctx, m, opts...)
}

type ValidateImportAXFRRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateImportAXFRRequest) PrimaryServerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for primary_server")
	}

	return validatorFn, nil
}

func (v *ValidateImportAXFRRequest) TsigKeyNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tsig_key_name")
	}

	return validatorFn, nil
}

func (v *ValidateImportAXFRRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ImportAXFRRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ImportAXFRRequest got type %s", t)
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

	if fv, exists := v.FldValidators["primary_server"]; exists {

		vOpts := append(opts, db.WithValidateField("primary_server"))
		if err := fv(ctx, m.GetPrimaryServer(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tsig_key_algorithm"]; exists {

		vOpts := append(opts, db.WithValidateField("tsig_key_algorithm"))
		if err := fv(ctx, m.GetTsigKeyAlgorithm(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tsig_key_name"]; exists {

		vOpts := append(opts, db.WithValidateField("tsig_key_name"))
		if err := fv(ctx, m.GetTsigKeyName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tsig_key_value"]; exists {

		vOpts := append(opts, db.WithValidateField("tsig_key_value"))
		if err := fv(ctx, m.GetTsigKeyValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultImportAXFRRequestValidator = func() *ValidateImportAXFRRequest {
	v := &ValidateImportAXFRRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPrimaryServer := v.PrimaryServerValidationRuleHandler
	rulesPrimaryServer := map[string]string{
		"ves.io.schema.rules.string.ipv4": "true",
	}
	vFn, err = vrhPrimaryServer(rulesPrimaryServer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ImportAXFRRequest.primary_server: %s", err)
		panic(errMsg)
	}
	v.FldValidators["primary_server"] = vFn

	vrhTsigKeyName := v.TsigKeyNameValidationRuleHandler
	rulesTsigKeyName := map[string]string{
		"ves.io.schema.rules.string.hostname": "true",
	}
	vFn, err = vrhTsigKeyName(rulesTsigKeyName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ImportAXFRRequest.tsig_key_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tsig_key_name"] = vFn

	v.FldValidators["tsig_key_value"] = ves_io_schema.SecretTypeValidator().Validate

	return v
}()

func ImportAXFRRequestValidator() db.Validator {
	return DefaultImportAXFRRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ImportAXFRResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ImportAXFRResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ImportAXFRResponse) DeepCopy() *ImportAXFRResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ImportAXFRResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ImportAXFRResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ImportAXFRResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ImportAXFRResponseValidator().Validate(ctx, m, opts...)
}

func (m *ImportAXFRResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetConfigurationDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ImportAXFRResponse) GetConfigurationDRefInfo() ([]db.DRefInfo, error) {
	if m.GetConfiguration() == nil {
		return nil, nil
	}

	drInfos, err := m.GetConfiguration().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetConfiguration().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "configuration." + dri.DRField
	}
	return drInfos, err

}

type ValidateImportAXFRResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateImportAXFRResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ImportAXFRResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ImportAXFRResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["configuration"]; exists {

		vOpts := append(opts, db.WithValidateField("configuration"))
		if err := fv(ctx, m.GetConfiguration(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultImportAXFRResponseValidator = func() *ValidateImportAXFRResponse {
	v := &ValidateImportAXFRResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["configuration"] = PrimaryDNSGetSpecTypeValidator().Validate

	return v
}()

func ImportAXFRResponseValidator() db.Validator {
	return DefaultImportAXFRResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ImportF5CSZoneRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ImportF5CSZoneRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ImportF5CSZoneRequest) DeepCopy() *ImportF5CSZoneRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ImportF5CSZoneRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ImportF5CSZoneRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ImportF5CSZoneRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ImportF5CSZoneRequestValidator().Validate(ctx, m, opts...)
}

type ValidateImportF5CSZoneRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateImportF5CSZoneRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ImportF5CSZoneRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ImportF5CSZoneRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["configuration"]; exists {

		vOpts := append(opts, db.WithValidateField("configuration"))
		if err := fv(ctx, m.GetConfiguration(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultImportF5CSZoneRequestValidator = func() *ValidateImportF5CSZoneRequest {
	v := &ValidateImportF5CSZoneRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ImportF5CSZoneRequestValidator() db.Validator {
	return DefaultImportF5CSZoneRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ImportF5CSZoneResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ImportF5CSZoneResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *ImportF5CSZoneResponse) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetSpec().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting ImportF5CSZoneResponse.spec")
	}

	return nil
}

func (m *ImportF5CSZoneResponse) DeepCopy() *ImportF5CSZoneResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ImportF5CSZoneResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ImportF5CSZoneResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ImportF5CSZoneResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ImportF5CSZoneResponseValidator().Validate(ctx, m, opts...)
}

func (m *ImportF5CSZoneResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetSpecDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ImportF5CSZoneResponse) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	if m.GetSpec() == nil {
		return nil, nil
	}

	drInfos, err := m.GetSpec().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetSpec().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "spec." + dri.DRField
	}
	return drInfos, err

}

type ValidateImportF5CSZoneResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateImportF5CSZoneResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ImportF5CSZoneResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ImportF5CSZoneResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, m.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultImportF5CSZoneResponseValidator = func() *ValidateImportF5CSZoneResponse {
	v := &ValidateImportF5CSZoneResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectGetMetaTypeValidator().Validate

	v.FldValidators["spec"] = GetSpecTypeValidator().Validate

	return v
}()

func ImportF5CSZoneResponseValidator() db.Validator {
	return DefaultImportF5CSZoneResponseValidator
}

func (m *ImportF5CSZoneResponse) fromObject(e db.Entry, withDeepCopy bool) {
	f := e.(*DBObject)
	if withDeepCopy {
		f = e.DeepCopy().(*DBObject)
	}
	_ = f

	if m.Metadata == nil {
		m.Metadata = &ves_io_schema.ObjectGetMetaType{}
	}
	m.Metadata.FromObjectMetaTypeWithoutDeepCopy(f.GetMetadata())

	if m.Spec == nil {
		m.Spec = &GetSpecType{}
	}
	m.Spec.FromGlobalSpecTypeWithoutDeepCopy(f.GetSpec().GetGcSpec())

	if m.SystemMetadata == nil {
		m.SystemMetadata = &ves_io_schema.SystemObjectGetMetaType{}
	}
	m.SystemMetadata.FromSystemObjectMetaTypeWithoutDeepCopy(f.GetSystemMetadata())

}

func (m *ImportF5CSZoneResponse) FromObject(e db.Entry) {
	m.fromObject(e, true)
}

func (m *ImportF5CSZoneResponse) FromObjectWithoutDeepCopy(e db.Entry) {
	m.fromObject(e, false)
}

func (m *ImportF5CSZoneResponse) toObject(e db.Entry, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1
	f := e.(*DBObject)
	_ = f

	if m1.Metadata != nil {
		if f.Metadata == nil {
			f.Metadata = &ves_io_schema.ObjectMetaType{}
		}
	} else if f.Metadata != nil {
		f.Metadata = nil
	}

	if m1.Metadata != nil {
		m1.Metadata.ToObjectMetaTypeWithoutDeepCopy(f.Metadata)
	}

	if m1.Spec != nil {
		if f.Spec == nil {
			f.Spec = &SpecType{}
		}
	} else if f.Spec != nil {
		f.Spec = nil
	}

	if m1.Spec != nil {
		if f.Spec.GcSpec == nil {
			f.Spec.GcSpec = &GlobalSpecType{}
		}
	} else if f.Spec != nil {
		f.Spec.GcSpec = nil
	}

	if m1.Spec != nil {
		m1.Spec.ToGlobalSpecTypeWithoutDeepCopy(f.Spec.GcSpec)
	}

	if m1.SystemMetadata != nil {
		if f.SystemMetadata == nil {
			f.SystemMetadata = &ves_io_schema.SystemObjectMetaType{}
		}
	} else if f.SystemMetadata != nil {
		f.SystemMetadata = nil
	}

	if m1.SystemMetadata != nil {
		m1.SystemMetadata.ToSystemObjectMetaTypeWithoutDeepCopy(f.SystemMetadata)
	}

}

func (m *ImportF5CSZoneResponse) ToObject(e db.Entry) {
	m.toObject(e, true)
}

func (m *ImportF5CSZoneResponse) ToObjectWithoutDeepCopy(e db.Entry) {
	m.toObject(e, false)
}
