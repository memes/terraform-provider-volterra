// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package topology

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

func (m *GetAWSDCConnectionRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAWSDCConnectionRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAWSDCConnectionRequest) DeepCopy() *GetAWSDCConnectionRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAWSDCConnectionRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAWSDCConnectionRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAWSDCConnectionRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAWSDCConnectionRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetAWSDCConnectionRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAWSDCConnectionRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAWSDCConnectionRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAWSDCConnectionRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["connect_id"]; exists {

		vOpts := append(opts, db.WithValidateField("connect_id"))
		if err := fv(ctx, m.GetConnectId(), vOpts...); err != nil {
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
var DefaultGetAWSDCConnectionRequestValidator = func() *ValidateGetAWSDCConnectionRequest {
	v := &ValidateGetAWSDCConnectionRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAWSDCConnectionRequestValidator() db.Validator {
	return DefaultGetAWSDCConnectionRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAWSDCConnectionResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAWSDCConnectionResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAWSDCConnectionResponse) DeepCopy() *GetAWSDCConnectionResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAWSDCConnectionResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAWSDCConnectionResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAWSDCConnectionResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAWSDCConnectionResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetAWSDCConnectionResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAWSDCConnectionResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAWSDCConnectionResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAWSDCConnectionResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["connection_status"]; exists {

		vOpts := append(opts, db.WithValidateField("connection_status"))
		if err := fv(ctx, m.GetConnectionStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAWSDCConnectionResponseValidator = func() *ValidateGetAWSDCConnectionResponse {
	v := &ValidateGetAWSDCConnectionResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAWSDCConnectionResponseValidator() db.Validator {
	return DefaultGetAWSDCConnectionResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAWSDCGatewayRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAWSDCGatewayRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAWSDCGatewayRequest) DeepCopy() *GetAWSDCGatewayRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAWSDCGatewayRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAWSDCGatewayRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAWSDCGatewayRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAWSDCGatewayRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetAWSDCGatewayRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAWSDCGatewayRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAWSDCGatewayRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAWSDCGatewayRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["gateway_id"]; exists {

		vOpts := append(opts, db.WithValidateField("gateway_id"))
		if err := fv(ctx, m.GetGatewayId(), vOpts...); err != nil {
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
var DefaultGetAWSDCGatewayRequestValidator = func() *ValidateGetAWSDCGatewayRequest {
	v := &ValidateGetAWSDCGatewayRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAWSDCGatewayRequestValidator() db.Validator {
	return DefaultGetAWSDCGatewayRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAWSDCGatewayResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAWSDCGatewayResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAWSDCGatewayResponse) DeepCopy() *GetAWSDCGatewayResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAWSDCGatewayResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAWSDCGatewayResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAWSDCGatewayResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAWSDCGatewayResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetAWSDCGatewayResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAWSDCGatewayResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAWSDCGatewayResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAWSDCGatewayResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["gateway_status"]; exists {

		vOpts := append(opts, db.WithValidateField("gateway_status"))
		if err := fv(ctx, m.GetGatewayStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAWSDCGatewayResponseValidator = func() *ValidateGetAWSDCGatewayResponse {
	v := &ValidateGetAWSDCGatewayResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAWSDCGatewayResponseValidator() db.Validator {
	return DefaultGetAWSDCGatewayResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAWSDCVIFRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAWSDCVIFRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAWSDCVIFRequest) DeepCopy() *GetAWSDCVIFRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAWSDCVIFRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAWSDCVIFRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAWSDCVIFRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAWSDCVIFRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetAWSDCVIFRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAWSDCVIFRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAWSDCVIFRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAWSDCVIFRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vif_id"]; exists {

		vOpts := append(opts, db.WithValidateField("vif_id"))
		if err := fv(ctx, m.GetVifId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAWSDCVIFRequestValidator = func() *ValidateGetAWSDCVIFRequest {
	v := &ValidateGetAWSDCVIFRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAWSDCVIFRequestValidator() db.Validator {
	return DefaultGetAWSDCVIFRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetAWSDCVIFResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetAWSDCVIFResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetAWSDCVIFResponse) DeepCopy() *GetAWSDCVIFResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetAWSDCVIFResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetAWSDCVIFResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetAWSDCVIFResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetAWSDCVIFResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetAWSDCVIFResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetAWSDCVIFResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetAWSDCVIFResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetAWSDCVIFResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["vif_status"]; exists {

		vOpts := append(opts, db.WithValidateField("vif_status"))
		if err := fv(ctx, m.GetVifStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetAWSDCVIFResponseValidator = func() *ValidateGetAWSDCVIFResponse {
	v := &ValidateGetAWSDCVIFResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetAWSDCVIFResponseValidator() db.Validator {
	return DefaultGetAWSDCVIFResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ListCloudNetworkTagKeysRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListCloudNetworkTagKeysRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListCloudNetworkTagKeysRequest) DeepCopy() *ListCloudNetworkTagKeysRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListCloudNetworkTagKeysRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListCloudNetworkTagKeysRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListCloudNetworkTagKeysRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListCloudNetworkTagKeysRequestValidator().Validate(ctx, m, opts...)
}

type ValidateListCloudNetworkTagKeysRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListCloudNetworkTagKeysRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListCloudNetworkTagKeysRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListCloudNetworkTagKeysRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cloud_type"]; exists {

		vOpts := append(opts, db.WithValidateField("cloud_type"))
		if err := fv(ctx, m.GetCloudType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query_key"]; exists {

		vOpts := append(opts, db.WithValidateField("query_key"))
		if err := fv(ctx, m.GetQueryKey(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListCloudNetworkTagKeysRequestValidator = func() *ValidateListCloudNetworkTagKeysRequest {
	v := &ValidateListCloudNetworkTagKeysRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListCloudNetworkTagKeysRequestValidator() db.Validator {
	return DefaultListCloudNetworkTagKeysRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ListCloudNetworkTagKeysResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListCloudNetworkTagKeysResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListCloudNetworkTagKeysResponse) DeepCopy() *ListCloudNetworkTagKeysResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListCloudNetworkTagKeysResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListCloudNetworkTagKeysResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListCloudNetworkTagKeysResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListCloudNetworkTagKeysResponseValidator().Validate(ctx, m, opts...)
}

type ValidateListCloudNetworkTagKeysResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListCloudNetworkTagKeysResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListCloudNetworkTagKeysResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListCloudNetworkTagKeysResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["keys"]; exists {

		vOpts := append(opts, db.WithValidateField("keys"))
		for idx, item := range m.GetKeys() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListCloudNetworkTagKeysResponseValidator = func() *ValidateListCloudNetworkTagKeysResponse {
	v := &ValidateListCloudNetworkTagKeysResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListCloudNetworkTagKeysResponseValidator() db.Validator {
	return DefaultListCloudNetworkTagKeysResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ListCloudNetworkTagValuesRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListCloudNetworkTagValuesRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListCloudNetworkTagValuesRequest) DeepCopy() *ListCloudNetworkTagValuesRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListCloudNetworkTagValuesRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListCloudNetworkTagValuesRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListCloudNetworkTagValuesRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListCloudNetworkTagValuesRequestValidator().Validate(ctx, m, opts...)
}

type ValidateListCloudNetworkTagValuesRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListCloudNetworkTagValuesRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListCloudNetworkTagValuesRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListCloudNetworkTagValuesRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cloud_type"]; exists {

		vOpts := append(opts, db.WithValidateField("cloud_type"))
		if err := fv(ctx, m.GetCloudType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query_key"]; exists {

		vOpts := append(opts, db.WithValidateField("query_key"))
		if err := fv(ctx, m.GetQueryKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query_value"]; exists {

		vOpts := append(opts, db.WithValidateField("query_value"))
		if err := fv(ctx, m.GetQueryValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListCloudNetworkTagValuesRequestValidator = func() *ValidateListCloudNetworkTagValuesRequest {
	v := &ValidateListCloudNetworkTagValuesRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListCloudNetworkTagValuesRequestValidator() db.Validator {
	return DefaultListCloudNetworkTagValuesRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ListCloudNetworkTagValuesResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListCloudNetworkTagValuesResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListCloudNetworkTagValuesResponse) DeepCopy() *ListCloudNetworkTagValuesResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListCloudNetworkTagValuesResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListCloudNetworkTagValuesResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListCloudNetworkTagValuesResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListCloudNetworkTagValuesResponseValidator().Validate(ctx, m, opts...)
}

type ValidateListCloudNetworkTagValuesResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListCloudNetworkTagValuesResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListCloudNetworkTagValuesResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListCloudNetworkTagValuesResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["values"]; exists {

		vOpts := append(opts, db.WithValidateField("values"))
		for idx, item := range m.GetValues() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListCloudNetworkTagValuesResponseValidator = func() *ValidateListCloudNetworkTagValuesResponse {
	v := &ValidateListCloudNetworkTagValuesResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListCloudNetworkTagValuesResponseValidator() db.Validator {
	return DefaultListCloudNetworkTagValuesResponseValidator
}
