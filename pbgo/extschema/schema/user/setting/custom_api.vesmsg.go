// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package setting

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

func (m *Empty) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Empty) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Empty) DeepCopy() *Empty {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Empty{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Empty) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Empty) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EmptyValidator().Validate(ctx, m, opts...)
}

type ValidateEmpty struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEmpty) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Empty)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Empty got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEmptyValidator = func() *ValidateEmpty {
	v := &ValidateEmpty{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func EmptyValidator() db.Validator {
	return DefaultEmptyValidator
}

// augmented methods on protoc/std generated struct

func (m *InitialAccess) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *InitialAccess) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *InitialAccess) DeepCopy() *InitialAccess {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &InitialAccess{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *InitialAccess) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *InitialAccess) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return InitialAccessValidator().Validate(ctx, m, opts...)
}

type ValidateInitialAccess struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateInitialAccess) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*InitialAccess)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *InitialAccess got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["email_sent"]; exists {

		vOpts := append(opts, db.WithValidateField("email_sent"))
		if err := fv(ctx, m.GetEmailSent(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_requesting_time"]; exists {

		vOpts := append(opts, db.WithValidateField("last_requesting_time"))
		if err := fv(ctx, m.GetLastRequestingTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["requested"]; exists {

		vOpts := append(opts, db.WithValidateField("requested"))
		if err := fv(ctx, m.GetRequested(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultInitialAccessValidator = func() *ValidateInitialAccess {
	v := &ValidateInitialAccess{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func InitialAccessValidator() db.Validator {
	return DefaultInitialAccessValidator
}

// augmented methods on protoc/std generated struct

func (m *Notification) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Notification) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Notification) DeepCopy() *Notification {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Notification{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Notification) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Notification) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NotificationValidator().Validate(ctx, m, opts...)
}

type ValidateNotification struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNotification) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Notification)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Notification got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["code"]; exists {

		vOpts := append(opts, db.WithValidateField("code"))
		if err := fv(ctx, m.GetCode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["enabled"]; exists {

		vOpts := append(opts, db.WithValidateField("enabled"))
		if err := fv(ctx, m.GetEnabled(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["label"]; exists {

		vOpts := append(opts, db.WithValidateField("label"))
		if err := fv(ctx, m.GetLabel(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNotificationValidator = func() *ValidateNotification {
	v := &ValidateNotification{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NotificationValidator() db.Validator {
	return DefaultNotificationValidator
}

// augmented methods on protoc/std generated struct

func (m *NotificationList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NotificationList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NotificationList) DeepCopy() *NotificationList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NotificationList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NotificationList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NotificationList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NotificationListValidator().Validate(ctx, m, opts...)
}

type ValidateNotificationList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNotificationList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NotificationList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NotificationList got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["notifications"]; exists {

		vOpts := append(opts, db.WithValidateField("notifications"))
		for idx, item := range m.GetNotifications() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNotificationListValidator = func() *ValidateNotificationList {
	v := &ValidateNotificationList{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NotificationListValidator() db.Validator {
	return DefaultNotificationListValidator
}

// augmented methods on protoc/std generated struct

func (m *NtfnPreferencesMap) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NtfnPreferencesMap) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NtfnPreferencesMap) DeepCopy() *NtfnPreferencesMap {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NtfnPreferencesMap{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NtfnPreferencesMap) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NtfnPreferencesMap) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NtfnPreferencesMapValidator().Validate(ctx, m, opts...)
}

type ValidateNtfnPreferencesMap struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNtfnPreferencesMap) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NtfnPreferencesMap)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NtfnPreferencesMap got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ntfn_preferences_map"]; exists {

		vOpts := append(opts, db.WithValidateField("ntfn_preferences_map"))
		for key, value := range m.GetNtfnPreferencesMap() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNtfnPreferencesMapValidator = func() *ValidateNtfnPreferencesMap {
	v := &ValidateNtfnPreferencesMap{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NtfnPreferencesMapValidator() db.Validator {
	return DefaultNtfnPreferencesMapValidator
}

// augmented methods on protoc/std generated struct

func (m *NtfnToUnset) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NtfnToUnset) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NtfnToUnset) DeepCopy() *NtfnToUnset {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NtfnToUnset{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NtfnToUnset) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NtfnToUnset) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NtfnToUnsetValidator().Validate(ctx, m, opts...)
}

type ValidateNtfnToUnset struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNtfnToUnset) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NtfnToUnset)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NtfnToUnset got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ntfn_to_unset"]; exists {

		vOpts := append(opts, db.WithValidateField("ntfn_to_unset"))
		if err := fv(ctx, m.GetNtfnToUnset(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNtfnToUnsetValidator = func() *ValidateNtfnToUnset {
	v := &ValidateNtfnToUnset{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func NtfnToUnsetValidator() db.Validator {
	return DefaultNtfnToUnsetValidator
}

// augmented methods on protoc/std generated struct

func (m *PersonaPreferences) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *PersonaPreferences) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *PersonaPreferences) DeepCopy() *PersonaPreferences {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &PersonaPreferences{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *PersonaPreferences) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *PersonaPreferences) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PersonaPreferencesValidator().Validate(ctx, m, opts...)
}

type ValidatePersonaPreferences struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePersonaPreferences) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*PersonaPreferences)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *PersonaPreferences got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["billing"]; exists {

		vOpts := append(opts, db.WithValidateField("billing"))
		if err := fv(ctx, m.GetBilling(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["dev_ops"]; exists {

		vOpts := append(opts, db.WithValidateField("dev_ops"))
		if err := fv(ctx, m.GetDevOps(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["developer"]; exists {

		vOpts := append(opts, db.WithValidateField("developer"))
		if err := fv(ctx, m.GetDeveloper(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["master"]; exists {

		vOpts := append(opts, db.WithValidateField("master"))
		if err := fv(ctx, m.GetMaster(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["net_ops"]; exists {

		vOpts := append(opts, db.WithValidateField("net_ops"))
		if err := fv(ctx, m.GetNetOps(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sec_ops"]; exists {

		vOpts := append(opts, db.WithValidateField("sec_ops"))
		if err := fv(ctx, m.GetSecOps(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPersonaPreferencesValidator = func() *ValidatePersonaPreferences {
	v := &ValidatePersonaPreferences{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func PersonaPreferencesValidator() db.Validator {
	return DefaultPersonaPreferencesValidator
}

// augmented methods on protoc/std generated struct

func (m *SetViewPreferenceRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetViewPreferenceRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetViewPreferenceRequest) DeepCopy() *SetViewPreferenceRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetViewPreferenceRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetViewPreferenceRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetViewPreferenceRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetViewPreferenceRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetViewPreferenceRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetViewPreferenceRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetViewPreferenceRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetViewPreferenceRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["advanced_view"]; exists {

		vOpts := append(opts, db.WithValidateField("advanced_view"))
		if err := fv(ctx, m.GetAdvancedView(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["persona_preferences"]; exists {

		vOpts := append(opts, db.WithValidateField("persona_preferences"))
		if err := fv(ctx, m.GetPersonaPreferences(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetViewPreferenceRequestValidator = func() *ValidateSetViewPreferenceRequest {
	v := &ValidateSetViewPreferenceRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetViewPreferenceRequestValidator() db.Validator {
	return DefaultSetViewPreferenceRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *UpdateImageRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateImageRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateImageRequest) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.Image = []byte{}

	return copy.string()
}

func (m *UpdateImageRequest) GoString() string {
	copy := m.DeepCopy()
	copy.Image = []byte{}

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *UpdateImageRequest) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.Image = []byte{}

	return nil
}

func (m *UpdateImageRequest) DeepCopy() *UpdateImageRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateImageRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateImageRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateImageRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateImageRequestValidator().Validate(ctx, m, opts...)
}

type ValidateUpdateImageRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateImageRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateImageRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateImageRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["content_type"]; exists {

		vOpts := append(opts, db.WithValidateField("content_type"))
		if err := fv(ctx, m.GetContentType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["image"]; exists {

		vOpts := append(opts, db.WithValidateField("image"))
		if err := fv(ctx, m.GetImage(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpdateImageRequestValidator = func() *ValidateUpdateImageRequest {
	v := &ValidateUpdateImageRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UpdateImageRequestValidator() db.Validator {
	return DefaultUpdateImageRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *UserEmail) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UserEmail) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UserEmail) DeepCopy() *UserEmail {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UserEmail{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UserEmail) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UserEmail) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UserEmailValidator().Validate(ctx, m, opts...)
}

type ValidateUserEmail struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUserEmail) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UserEmail)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UserEmail got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["user_email"]; exists {

		vOpts := append(opts, db.WithValidateField("user_email"))
		if err := fv(ctx, m.GetUserEmail(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUserEmailValidator = func() *ValidateUserEmail {
	v := &ValidateUserEmail{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UserEmailValidator() db.Validator {
	return DefaultUserEmailValidator
}

// augmented methods on protoc/std generated struct

func (m *UserSession) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UserSession) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UserSession) DeepCopy() *UserSession {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UserSession{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UserSession) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UserSession) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UserSessionValidator().Validate(ctx, m, opts...)
}

type ValidateUserSession struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUserSession) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UserSession)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UserSession got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ip_address"]; exists {

		vOpts := append(opts, db.WithValidateField("ip_address"))
		if err := fv(ctx, m.GetIpAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_access"]; exists {

		vOpts := append(opts, db.WithValidateField("last_access"))
		if err := fv(ctx, m.GetLastAccess(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start"]; exists {

		vOpts := append(opts, db.WithValidateField("start"))
		if err := fv(ctx, m.GetStart(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUserSessionValidator = func() *ValidateUserSession {
	v := &ValidateUserSession{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UserSessionValidator() db.Validator {
	return DefaultUserSessionValidator
}

// augmented methods on protoc/std generated struct

func (m *UserSessionList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UserSessionList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UserSessionList) DeepCopy() *UserSessionList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UserSessionList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UserSessionList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UserSessionList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UserSessionListValidator().Validate(ctx, m, opts...)
}

type ValidateUserSessionList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUserSessionList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UserSessionList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UserSessionList got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["user_sessions"]; exists {

		vOpts := append(opts, db.WithValidateField("user_sessions"))
		for idx, item := range m.GetUserSessions() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUserSessionListValidator = func() *ValidateUserSessionList {
	v := &ValidateUserSessionList{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UserSessionListValidator() db.Validator {
	return DefaultUserSessionListValidator
}

// augmented methods on protoc/std generated struct

func (m *UserSettingsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UserSettingsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UserSettingsRequest) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.Image = []byte{}

	return copy.string()
}

func (m *UserSettingsRequest) GoString() string {
	copy := m.DeepCopy()
	copy.Image = []byte{}

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *UserSettingsRequest) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.Image = []byte{}

	return nil
}

func (m *UserSettingsRequest) DeepCopy() *UserSettingsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UserSettingsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UserSettingsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UserSettingsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UserSettingsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateUserSettingsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUserSettingsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UserSettingsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UserSettingsRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["enabled_notifications"]; exists {

		vOpts := append(opts, db.WithValidateField("enabled_notifications"))
		for idx, item := range m.GetEnabledNotifications() {
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

	if fv, exists := v.FldValidators["image"]; exists {

		vOpts := append(opts, db.WithValidateField("image"))
		if err := fv(ctx, m.GetImage(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["otp_enabled"]; exists {

		vOpts := append(opts, db.WithValidateField("otp_enabled"))
		if err := fv(ctx, m.GetOtpEnabled(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUserSettingsRequestValidator = func() *ValidateUserSettingsRequest {
	v := &ValidateUserSettingsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UserSettingsRequestValidator() db.Validator {
	return DefaultUserSettingsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *UserSettingsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UserSettingsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UserSettingsResponse) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.Image = []byte{}

	return copy.string()
}

func (m *UserSettingsResponse) GoString() string {
	copy := m.DeepCopy()
	copy.Image = []byte{}

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *UserSettingsResponse) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.Image = []byte{}

	return nil
}

func (m *UserSettingsResponse) DeepCopy() *UserSettingsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UserSettingsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UserSettingsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UserSettingsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UserSettingsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateUserSettingsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUserSettingsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UserSettingsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UserSettingsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["image"]; exists {

		vOpts := append(opts, db.WithValidateField("image"))
		if err := fv(ctx, m.GetImage(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["initial_access"]; exists {

		vOpts := append(opts, db.WithValidateField("initial_access"))
		if err := fv(ctx, m.GetInitialAccess(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["is_next_request_allowed"]; exists {

		vOpts := append(opts, db.WithValidateField("is_next_request_allowed"))
		if err := fv(ctx, m.GetIsNextRequestAllowed(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["notifications"]; exists {

		vOpts := append(opts, db.WithValidateField("notifications"))
		for idx, item := range m.GetNotifications() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["otp_enabled"]; exists {

		vOpts := append(opts, db.WithValidateField("otp_enabled"))
		if err := fv(ctx, m.GetOtpEnabled(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["otp_status"]; exists {

		vOpts := append(opts, db.WithValidateField("otp_status"))
		if err := fv(ctx, m.GetOtpStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUserSettingsResponseValidator = func() *ValidateUserSettingsResponse {
	v := &ValidateUserSettingsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UserSettingsResponseValidator() db.Validator {
	return DefaultUserSettingsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ViewPreference) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ViewPreference) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ViewPreference) DeepCopy() *ViewPreference {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ViewPreference{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ViewPreference) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ViewPreference) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ViewPreferenceValidator().Validate(ctx, m, opts...)
}

type ValidateViewPreference struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateViewPreference) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ViewPreference)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ViewPreference got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["advanced_view"]; exists {

		vOpts := append(opts, db.WithValidateField("advanced_view"))
		if err := fv(ctx, m.GetAdvancedView(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["initialized"]; exists {

		vOpts := append(opts, db.WithValidateField("initialized"))
		if err := fv(ctx, m.GetInitialized(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["persona_preferences"]; exists {

		vOpts := append(opts, db.WithValidateField("persona_preferences"))
		if err := fv(ctx, m.GetPersonaPreferences(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultViewPreferenceValidator = func() *ValidateViewPreference {
	v := &ValidateViewPreference{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ViewPreferenceValidator() db.Validator {
	return DefaultViewPreferenceValidator
}
