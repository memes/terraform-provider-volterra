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

func (m *CertInfoType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CertInfoType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CertInfoType) DeepCopy() *CertInfoType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CertInfoType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CertInfoType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CertInfoType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CertInfoTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCertInfoType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCertInfoType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CertInfoType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CertInfoType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["common_name"]; exists {

		vOpts := append(opts, db.WithValidateField("common_name"))
		if err := fv(ctx, m.GetCommonName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiry"]; exists {

		vOpts := append(opts, db.WithValidateField("expiry"))
		if err := fv(ctx, m.GetExpiry(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["issuer"]; exists {

		vOpts := append(opts, db.WithValidateField("issuer"))
		if err := fv(ctx, m.GetIssuer(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["organization"]; exists {

		vOpts := append(opts, db.WithValidateField("organization"))
		if err := fv(ctx, m.GetOrganization(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_key_algorithm"]; exists {

		vOpts := append(opts, db.WithValidateField("public_key_algorithm"))
		if err := fv(ctx, m.GetPublicKeyAlgorithm(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subject_alternative_names"]; exists {

		vOpts := append(opts, db.WithValidateField("subject_alternative_names"))
		for idx, item := range m.GetSubjectAlternativeNames() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCertInfoTypeValidator = func() *ValidateCertInfoType {
	v := &ValidateCertInfoType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CertInfoTypeValidator() db.Validator {
	return DefaultCertInfoTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *CertificateParamsType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CertificateParamsType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CertificateParamsType) DeepCopy() *CertificateParamsType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CertificateParamsType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CertificateParamsType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CertificateParamsType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CertificateParamsTypeValidator().Validate(ctx, m, opts...)
}

func (m *CertificateParamsType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetCertificatesDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCertificatesDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetCrlDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCrlDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *CertificateParamsType) GetCertificatesDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetCertificates()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("CertificateParamsType.certificates[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "certificate.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "certificates",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetCertificatesDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CertificateParamsType) GetCertificatesDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "certificate.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: certificate")
	}
	for _, ref := range m.GetCertificates() {
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

func (m *CertificateParamsType) GetCrlDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetCrl()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("CertificateParamsType.crl[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "crl.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "crl",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetCrlDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CertificateParamsType) GetCrlDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "crl.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: crl")
	}
	for _, ref := range m.GetCrl() {
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

type ValidateCertificateParamsType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCertificateParamsType) CertificatesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for certificates")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for certificates")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ObjectRefType, got %T", val)
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
			return errors.Wrap(err, "repeated certificates")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items certificates")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCertificateParamsType) CipherSuitesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for cipher_suites")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for cipher_suites")
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
			return errors.Wrap(err, "repeated cipher_suites")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items cipher_suites")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCertificateParamsType) CrlValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for crl")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for crl")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ObjectRefType, got %T", val)
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
			return errors.Wrap(err, "repeated crl")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items crl")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCertificateParamsType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CertificateParamsType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CertificateParamsType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["certificates"]; exists {
		vOpts := append(opts, db.WithValidateField("certificates"))
		if err := fv(ctx, m.GetCertificates(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["cipher_suites"]; exists {
		vOpts := append(opts, db.WithValidateField("cipher_suites"))
		if err := fv(ctx, m.GetCipherSuites(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["crl"]; exists {
		vOpts := append(opts, db.WithValidateField("crl"))
		if err := fv(ctx, m.GetCrl(), vOpts...); err != nil {
			return err
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

	if fv, exists := v.FldValidators["require_client_certificate"]; exists {

		vOpts := append(opts, db.WithValidateField("require_client_certificate"))
		if err := fv(ctx, m.GetRequireClientCertificate(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["validation_params"]; exists {

		vOpts := append(opts, db.WithValidateField("validation_params"))
		if err := fv(ctx, m.GetValidationParams(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCertificateParamsTypeValidator = func() *ValidateCertificateParamsType {
	v := &ValidateCertificateParamsType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhCertificates := v.CertificatesValidationRuleHandler
	rulesCertificates := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "32",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhCertificates(rulesCertificates)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CertificateParamsType.certificates: %s", err)
		panic(errMsg)
	}
	v.FldValidators["certificates"] = vFn

	vrhCipherSuites := v.CipherSuitesValidationRuleHandler
	rulesCipherSuites := map[string]string{
		"ves.io.schema.rules.repeated.items.string.in": "[\"TLS_AES_128_GCM_SHA256\",\"TLS_AES_256_GCM_SHA384\",\"TLS_CHACHA20_POLY1305_SHA256\",\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\",\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\",\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\",\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\",\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\",\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\",\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\",\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\",\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\",\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\",\"TLS_RSA_WITH_AES_128_CBC_SHA\",\"TLS_RSA_WITH_AES_128_GCM_SHA256\",\"TLS_RSA_WITH_AES_256_CBC_SHA\",\"TLS_RSA_WITH_AES_256_GCM_SHA384\"]",
		"ves.io.schema.rules.repeated.unique":          "true",
	}
	vFn, err = vrhCipherSuites(rulesCipherSuites)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CertificateParamsType.cipher_suites: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cipher_suites"] = vFn

	vrhCrl := v.CrlValidationRuleHandler
	rulesCrl := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhCrl(rulesCrl)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CertificateParamsType.crl: %s", err)
		panic(errMsg)
	}
	v.FldValidators["crl"] = vFn

	v.FldValidators["validation_params"] = TlsValidationParamsTypeValidator().Validate

	return v
}()

func CertificateParamsTypeValidator() db.Validator {
	return DefaultCertificateParamsTypeValidator
}
