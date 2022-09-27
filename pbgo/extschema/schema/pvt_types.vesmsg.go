//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
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

func (m *DaemonEnvironmentType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DaemonEnvironmentType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DaemonEnvironmentType) DeepCopy() *DaemonEnvironmentType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DaemonEnvironmentType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DaemonEnvironmentType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DaemonEnvironmentType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DaemonEnvironmentTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDaemonEnvironmentType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDaemonEnvironmentType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DaemonEnvironmentType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DaemonEnvironmentType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["internal_dns_suffix"]; exists {

		vOpts := append(opts, db.WithValidateField("internal_dns_suffix"))
		if err := fv(ctx, m.GetInternalDnsSuffix(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDaemonEnvironmentTypeValidator = func() *ValidateDaemonEnvironmentType {
	v := &ValidateDaemonEnvironmentType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DaemonEnvironmentTypeValidator() db.Validator {
	return DefaultDaemonEnvironmentTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DaemonTLSParamsType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DaemonTLSParamsType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DaemonTLSParamsType) DeepCopy() *DaemonTLSParamsType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DaemonTLSParamsType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DaemonTLSParamsType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DaemonTLSParamsType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DaemonTLSParamsTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDaemonTLSParamsType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDaemonTLSParamsType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DaemonTLSParamsType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DaemonTLSParamsType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["client_params"]; exists {

		vOpts := append(opts, db.WithValidateField("client_params"))
		if err := fv(ctx, m.GetClientParams(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_params"]; exists {

		vOpts := append(opts, db.WithValidateField("server_params"))
		if err := fv(ctx, m.GetServerParams(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDaemonTLSParamsTypeValidator = func() *ValidateDaemonTLSParamsType {
	v := &ValidateDaemonTLSParamsType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DaemonTLSParamsTypeValidator() db.Validator {
	return DefaultDaemonTLSParamsTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DaemonTlsCertificateType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DaemonTlsCertificateType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DaemonTlsCertificateType) DeepCopy() *DaemonTlsCertificateType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DaemonTlsCertificateType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DaemonTlsCertificateType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DaemonTlsCertificateType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DaemonTlsCertificateTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDaemonTlsCertificateType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDaemonTlsCertificateType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DaemonTlsCertificateType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DaemonTlsCertificateType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["certificate_url"]; exists {

		vOpts := append(opts, db.WithValidateField("certificate_url"))
		if err := fv(ctx, m.GetCertificateUrl(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["private_key_url"]; exists {

		vOpts := append(opts, db.WithValidateField("private_key_url"))
		if err := fv(ctx, m.GetPrivateKeyUrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDaemonTlsCertificateTypeValidator = func() *ValidateDaemonTlsCertificateType {
	v := &ValidateDaemonTlsCertificateType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DaemonTlsCertificateTypeValidator() db.Validator {
	return DefaultDaemonTlsCertificateTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DaemonTlsParametersType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DaemonTlsParametersType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DaemonTlsParametersType) DeepCopy() *DaemonTlsParametersType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DaemonTlsParametersType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DaemonTlsParametersType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DaemonTlsParametersType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DaemonTlsParametersTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDaemonTlsParametersType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDaemonTlsParametersType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DaemonTlsParametersType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DaemonTlsParametersType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cipher_suites"]; exists {

		vOpts := append(opts, db.WithValidateField("cipher_suites"))
		for idx, item := range m.GetCipherSuites() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["maximum_protocol_version"]; exists {

		vOpts := append(opts, db.WithValidateField("maximum_protocol_version"))
		if err := fv(ctx, m.GetMaximumProtocolVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["minimum_protocol_version"]; exists {

		vOpts := append(opts, db.WithValidateField("minimum_protocol_version"))
		if err := fv(ctx, m.GetMinimumProtocolVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tls_certificates"]; exists {

		vOpts := append(opts, db.WithValidateField("tls_certificates"))
		for idx, item := range m.GetTlsCertificates() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["trusted_ca_url"]; exists {

		vOpts := append(opts, db.WithValidateField("trusted_ca_url"))
		if err := fv(ctx, m.GetTrustedCaUrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDaemonTlsParametersTypeValidator = func() *ValidateDaemonTlsParametersType {
	v := &ValidateDaemonTlsParametersType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DaemonTlsParametersTypeValidator() db.Validator {
	return DefaultDaemonTlsParametersTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *OperMetaType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OperMetaType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OperMetaType) DeepCopy() *OperMetaType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OperMetaType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OperMetaType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OperMetaType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OperMetaTypeValidator().Validate(ctx, m, opts...)
}

type ValidateOperMetaType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOperMetaType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OperMetaType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OperMetaType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["annotations"]; exists {

		vOpts := append(opts, db.WithValidateField("annotations"))
		for key, value := range m.GetAnnotations() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["creation_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("creation_timestamp"))
		if err := fv(ctx, m.GetCreationTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["modification_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("modification_timestamp"))
		if err := fv(ctx, m.GetModificationTimestamp(), vOpts...); err != nil {
			return err
		}

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

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["uid"]; exists {

		vOpts := append(opts, db.WithValidateField("uid"))
		if err := fv(ctx, m.GetUid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOperMetaTypeValidator = func() *ValidateOperMetaType {
	v := &ValidateOperMetaType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OperMetaTypeValidator() db.Validator {
	return DefaultOperMetaTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ServiceParameters) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ServiceParameters) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ServiceParameters) DeepCopy() *ServiceParameters {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ServiceParameters{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ServiceParameters) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ServiceParameters) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ServiceParametersValidator().Validate(ctx, m, opts...)
}

type ValidateServiceParameters struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateServiceParameters) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ServiceParameters)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ServiceParameters got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ctype"]; exists {

		vOpts := append(opts, db.WithValidateField("ctype"))
		if err := fv(ctx, m.GetCtype(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["url"]; exists {

		vOpts := append(opts, db.WithValidateField("url"))
		if err := fv(ctx, m.GetUrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultServiceParametersValidator = func() *ValidateServiceParameters {
	v := &ValidateServiceParameters{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ServiceParametersValidator() db.Validator {
	return DefaultServiceParametersValidator
}

// augmented methods on protoc/std generated struct

func (m *SuggestValuesResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuggestValuesResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuggestValuesResp) DeepCopy() *SuggestValuesResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuggestValuesResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuggestValuesResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuggestValuesResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuggestValuesRespValidator().Validate(ctx, m, opts...)
}

type ValidateSuggestValuesResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuggestValuesResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuggestValuesResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuggestValuesResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuggestValuesRespValidator = func() *ValidateSuggestValuesResp {
	v := &ValidateSuggestValuesResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SuggestValuesRespValidator() db.Validator {
	return DefaultSuggestValuesRespValidator
}

// augmented methods on protoc/std generated struct

func (m *SuggestedItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuggestedItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuggestedItem) DeepCopy() *SuggestedItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuggestedItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuggestedItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuggestedItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuggestedItemValidator().Validate(ctx, m, opts...)
}

type ValidateSuggestedItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuggestedItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuggestedItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuggestedItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		if err := fv(ctx, m.GetValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuggestedItemValidator = func() *ValidateSuggestedItem {
	v := &ValidateSuggestedItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SuggestedItemValidator() db.Validator {
	return DefaultSuggestedItemValidator
}

// augmented methods on protoc/std generated struct

func (m *SyncServerParamsType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SyncServerParamsType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SyncServerParamsType) DeepCopy() *SyncServerParamsType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SyncServerParamsType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SyncServerParamsType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SyncServerParamsType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SyncServerParamsTypeValidator().Validate(ctx, m, opts...)
}

type ValidateSyncServerParamsType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSyncServerParamsType) SyncServerListValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for sync_server_list")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ServiceParameters, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ServiceParametersValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for sync_server_list")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ServiceParameters)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ServiceParameters, got %T", val)
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
			return errors.Wrap(err, "repeated sync_server_list")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items sync_server_list")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSyncServerParamsType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SyncServerParamsType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SyncServerParamsType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["sync_server_list"]; exists {
		vOpts := append(opts, db.WithValidateField("sync_server_list"))
		if err := fv(ctx, m.GetSyncServerList(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vtrp_graceful_restart_params"]; exists {

		vOpts := append(opts, db.WithValidateField("vtrp_graceful_restart_params"))
		if err := fv(ctx, m.GetVtrpGracefulRestartParams(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSyncServerParamsTypeValidator = func() *ValidateSyncServerParamsType {
	v := &ValidateSyncServerParamsType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhSyncServerList := v.SyncServerListValidationRuleHandler
	rulesSyncServerList := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "2",
		"ves.io.schema.rules.repeated.min_items": "2",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhSyncServerList(rulesSyncServerList)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SyncServerParamsType.sync_server_list: %s", err)
		panic(errMsg)
	}
	v.FldValidators["sync_server_list"] = vFn

	return v
}()

func SyncServerParamsTypeValidator() db.Validator {
	return DefaultSyncServerParamsTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *UseragentType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UseragentType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UseragentType) DeepCopy() *UseragentType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UseragentType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UseragentType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UseragentType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UseragentTypeValidator().Validate(ctx, m, opts...)
}

type ValidateUseragentType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUseragentType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UseragentType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UseragentType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["host_name"]; exists {

		vOpts := append(opts, db.WithValidateField("host_name"))
		if err := fv(ctx, m.GetHostName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["program_name"]; exists {

		vOpts := append(opts, db.WithValidateField("program_name"))
		if err := fv(ctx, m.GetProgramName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service_name"]; exists {

		vOpts := append(opts, db.WithValidateField("service_name"))
		if err := fv(ctx, m.GetServiceName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("site_name"))
		if err := fv(ctx, m.GetSiteName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUseragentTypeValidator = func() *ValidateUseragentType {
	v := &ValidateUseragentType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UseragentTypeValidator() db.Validator {
	return DefaultUseragentTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *VTRPGracefulRestartParamsType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *VTRPGracefulRestartParamsType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *VTRPGracefulRestartParamsType) DeepCopy() *VTRPGracefulRestartParamsType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &VTRPGracefulRestartParamsType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *VTRPGracefulRestartParamsType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *VTRPGracefulRestartParamsType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return VTRPGracefulRestartParamsTypeValidator().Validate(ctx, m, opts...)
}

type ValidateVTRPGracefulRestartParamsType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVTRPGracefulRestartParamsType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*VTRPGracefulRestartParamsType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *VTRPGracefulRestartParamsType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["gr_enabled"]; exists {

		vOpts := append(opts, db.WithValidateField("gr_enabled"))
		if err := fv(ctx, m.GetGrEnabled(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["gr_time"]; exists {

		vOpts := append(opts, db.WithValidateField("gr_time"))
		if err := fv(ctx, m.GetGrTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["llgr_enabled"]; exists {

		vOpts := append(opts, db.WithValidateField("llgr_enabled"))
		if err := fv(ctx, m.GetLlgrEnabled(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["llgr_time"]; exists {

		vOpts := append(opts, db.WithValidateField("llgr_time"))
		if err := fv(ctx, m.GetLlgrTime(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultVTRPGracefulRestartParamsTypeValidator = func() *ValidateVTRPGracefulRestartParamsType {
	v := &ValidateVTRPGracefulRestartParamsType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func VTRPGracefulRestartParamsTypeValidator() db.Validator {
	return DefaultVTRPGracefulRestartParamsTypeValidator
}
