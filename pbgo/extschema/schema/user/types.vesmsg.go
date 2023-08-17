// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package user

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

func (m *AddonServiceStatus) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AddonServiceStatus) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AddonServiceStatus) DeepCopy() *AddonServiceStatus {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AddonServiceStatus{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AddonServiceStatus) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AddonServiceStatus) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AddonServiceStatusValidator().Validate(ctx, m, opts...)
}

type ValidateAddonServiceStatus struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAddonServiceStatus) DisplayNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for display_name")
	}

	return validatorFn, nil
}

func (v *ValidateAddonServiceStatus) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AddonServiceStatus)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AddonServiceStatus got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["access"]; exists {

		vOpts := append(opts, db.WithValidateField("access"))
		if err := fv(ctx, m.GetAccess(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["display_name"]; exists {

		vOpts := append(opts, db.WithValidateField("display_name"))
		if err := fv(ctx, m.GetDisplayName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAddonServiceStatusValidator = func() *ValidateAddonServiceStatus {
	v := &ValidateAddonServiceStatus{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhDisplayName := v.DisplayNameValidationRuleHandler
	rulesDisplayName := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhDisplayName(rulesDisplayName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AddonServiceStatus.display_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["display_name"] = vFn

	return v
}()

func AddonServiceStatusValidator() db.Validator {
	return DefaultAddonServiceStatusValidator
}

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

	return m.GetContactsDRefInfo()

}

func (m *CreateSpecType) GetContactsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetContacts()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("CreateSpecType.contacts[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "contact.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "contacts",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetContactsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CreateSpecType) GetContactsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "contact.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: contact")
	}
	for _, ref := range m.GetContacts() {
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

	if fv, exists := v.FldValidators["contacts"]; exists {

		vOpts := append(opts, db.WithValidateField("contacts"))
		for idx, item := range m.GetContacts() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locale"]; exists {

		vOpts := append(opts, db.WithValidateField("locale"))
		if err := fv(ctx, m.GetLocale(), vOpts...); err != nil {
			return err
		}

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
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

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

	return m.GetContactsDRefInfo()

}

func (m *GetSpecType) GetContactsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetContacts()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GetSpecType.contacts[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "contact.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "contacts",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetContactsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetSpecType) GetContactsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "contact.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: contact")
	}
	for _, ref := range m.GetContacts() {
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

	if fv, exists := v.FldValidators["contacts"]; exists {

		vOpts := append(opts, db.WithValidateField("contacts"))
		for idx, item := range m.GetContacts() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locale"]; exists {

		vOpts := append(opts, db.WithValidateField("locale"))
		if err := fv(ctx, m.GetLocale(), vOpts...); err != nil {
			return err
		}

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

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetContactsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetContactsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetGroupsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetGroupsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *GlobalSpecType) GetContactsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetContacts()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.contacts[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "contact.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "contacts",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetContactsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetContactsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "contact.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: contact")
	}
	for _, ref := range m.GetContacts() {
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

func (m *GlobalSpecType) GetGroupsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetGroups()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.groups[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "user_group.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "groups",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetGroupsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetGroupsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "user_group.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: user_group")
	}
	for _, ref := range m.GetGroups() {
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

	if fv, exists := v.FldValidators["contacts"]; exists {

		vOpts := append(opts, db.WithValidateField("contacts"))
		for idx, item := range m.GetContacts() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["domain_owner"]; exists {

		vOpts := append(opts, db.WithValidateField("domain_owner"))
		if err := fv(ctx, m.GetDomainOwner(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["groups"]; exists {

		vOpts := append(opts, db.WithValidateField("groups"))
		for idx, item := range m.GetGroups() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["idm_type"]; exists {

		vOpts := append(opts, db.WithValidateField("idm_type"))
		if err := fv(ctx, m.GetIdmType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_login_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("last_login_timestamp"))
		if err := fv(ctx, m.GetLastLoginTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locale"]; exists {

		vOpts := append(opts, db.WithValidateField("locale"))
		if err := fv(ctx, m.GetLocale(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sync_mode"]; exists {

		vOpts := append(opts, db.WithValidateField("sync_mode"))
		if err := fv(ctx, m.GetSyncMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_accepted"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_accepted"))
		if err := fv(ctx, m.GetTosAccepted(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_accepted_at"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_accepted_at"))
		if err := fv(ctx, m.GetTosAcceptedAt(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_version"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_version"))
		if err := fv(ctx, m.GetTosVersion(), vOpts...); err != nil {
			return err
		}

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

func (m *ReplaceSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetContactsDRefInfo()

}

func (m *ReplaceSpecType) GetContactsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetContacts()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("ReplaceSpecType.contacts[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "contact.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "contacts",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetContactsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *ReplaceSpecType) GetContactsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "contact.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: contact")
	}
	for _, ref := range m.GetContacts() {
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

	if fv, exists := v.FldValidators["contacts"]; exists {

		vOpts := append(opts, db.WithValidateField("contacts"))
		for idx, item := range m.GetContacts() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locale"]; exists {

		vOpts := append(opts, db.WithValidateField("locale"))
		if err := fv(ctx, m.GetLocale(), vOpts...); err != nil {
			return err
		}

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

// augmented methods on protoc/std generated struct

func (m *TileAccess) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *TileAccess) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *TileAccess) DeepCopy() *TileAccess {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &TileAccess{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *TileAccess) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *TileAccess) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return TileAccessValidator().Validate(ctx, m, opts...)
}

type ValidateTileAccess struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateTileAccess) DisplayNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for display_name")
	}

	return validatorFn, nil
}

func (v *ValidateTileAccess) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*TileAccess)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *TileAccess got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["display_name"]; exists {

		vOpts := append(opts, db.WithValidateField("display_name"))
		if err := fv(ctx, m.GetDisplayName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultTileAccessValidator = func() *ValidateTileAccess {
	v := &ValidateTileAccess{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhDisplayName := v.DisplayNameValidationRuleHandler
	rulesDisplayName := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhDisplayName(rulesDisplayName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for TileAccess.display_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["display_name"] = vFn

	return v
}()

func TileAccessValidator() db.Validator {
	return DefaultTileAccessValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.Contacts = f.GetContacts()
	m.Email = f.GetEmail()
	m.FirstName = f.GetFirstName()
	m.LastName = f.GetLastName()
	m.Locale = f.GetLocale()
	m.Type = f.GetType()
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

	f.Contacts = m1.Contacts
	f.Email = m1.Email
	f.FirstName = m1.FirstName
	f.LastName = m1.LastName
	f.Locale = m1.Locale
	f.Type = m1.Type
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
	m.Contacts = f.GetContacts()
	m.Email = f.GetEmail()
	m.FirstName = f.GetFirstName()
	m.LastName = f.GetLastName()
	m.Locale = f.GetLocale()
	m.Type = f.GetType()
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

	f.Contacts = m1.Contacts
	f.Email = m1.Email
	f.FirstName = m1.FirstName
	f.LastName = m1.LastName
	f.Locale = m1.Locale
	f.Type = m1.Type
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
	m.Contacts = f.GetContacts()
	m.FirstName = f.GetFirstName()
	m.LastName = f.GetLastName()
	m.Locale = f.GetLocale()
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

	f.Contacts = m1.Contacts
	f.FirstName = m1.FirstName
	f.LastName = m1.LastName
	f.Locale = m1.Locale
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
