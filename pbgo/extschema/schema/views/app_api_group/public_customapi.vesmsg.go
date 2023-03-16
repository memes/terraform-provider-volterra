//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package app_api_group

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

func (m *ApiEndpoint) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ApiEndpoint) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ApiEndpoint) DeepCopy() *ApiEndpoint {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ApiEndpoint{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ApiEndpoint) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ApiEndpoint) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApiEndpointValidator().Validate(ctx, m, opts...)
}

type ValidateApiEndpoint struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApiEndpoint) MethodValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for method")
	}

	return validatorFn, nil
}

func (v *ValidateApiEndpoint) PathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path")
	}

	return validatorFn, nil
}

func (v *ValidateApiEndpoint) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ApiEndpoint)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ApiEndpoint got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultApiEndpointValidator = func() *ValidateApiEndpoint {
	v := &ValidateApiEndpoint{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMethod := v.MethodValidationRuleHandler
	rulesMethod := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMethod(rulesMethod)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApiEndpoint.method: %s", err)
		panic(errMsg)
	}
	v.FldValidators["method"] = vFn

	vrhPath := v.PathValidationRuleHandler
	rulesPath := map[string]string{
		"ves.io.schema.rules.message.required":           "true",
		"ves.io.schema.rules.string.max_bytes":           "1024",
		"ves.io.schema.rules.string.min_bytes":           "1",
		"ves.io.schema.rules.string.templated_http_path": "true",
	}
	vFn, err = vrhPath(rulesPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApiEndpoint.path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path"] = vFn

	return v
}()

func ApiEndpointValidator() db.Validator {
	return DefaultApiEndpointValidator
}

// augmented methods on protoc/std generated struct

func (m *ApiGroupId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ApiGroupId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ApiGroupId) DeepCopy() *ApiGroupId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ApiGroupId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ApiGroupId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ApiGroupId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApiGroupIdValidator().Validate(ctx, m, opts...)
}

type ValidateApiGroupId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApiGroupId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ApiGroupId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ApiGroupId got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultApiGroupIdValidator = func() *ValidateApiGroupId {
	v := &ValidateApiGroupId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ApiGroupIdValidator() db.Validator {
	return DefaultApiGroupIdValidator
}

// augmented methods on protoc/std generated struct

func (m *ApiGroupStats) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ApiGroupStats) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ApiGroupStats) DeepCopy() *ApiGroupStats {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ApiGroupStats{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ApiGroupStats) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ApiGroupStats) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApiGroupStatsValidator().Validate(ctx, m, opts...)
}

type ValidateApiGroupStats struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApiGroupStats) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ApiGroupStats)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ApiGroupStats got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["outdated_api_endpoints_count"]; exists {

		vOpts := append(opts, db.WithValidateField("outdated_api_endpoints_count"))
		if err := fv(ctx, m.GetOutdatedApiEndpointsCount(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultApiGroupStatsValidator = func() *ValidateApiGroupStats {
	v := &ValidateApiGroupStats{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ApiGroupStatsValidator() db.Validator {
	return DefaultApiGroupStatsValidator
}

// augmented methods on protoc/std generated struct

func (m *ApiGroupsStatsItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ApiGroupsStatsItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ApiGroupsStatsItem) DeepCopy() *ApiGroupsStatsItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ApiGroupsStatsItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ApiGroupsStatsItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ApiGroupsStatsItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApiGroupsStatsItemValidator().Validate(ctx, m, opts...)
}

type ValidateApiGroupsStatsItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApiGroupsStatsItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ApiGroupsStatsItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ApiGroupsStatsItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["stats"]; exists {

		vOpts := append(opts, db.WithValidateField("stats"))
		if err := fv(ctx, m.GetStats(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultApiGroupsStatsItemValidator = func() *ValidateApiGroupsStatsItem {
	v := &ValidateApiGroupsStatsItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ApiGroupsStatsItemValidator() db.Validator {
	return DefaultApiGroupsStatsItemValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateApiGroupReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateApiGroupReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateApiGroupReq) DeepCopy() *EvaluateApiGroupReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateApiGroupReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateApiGroupReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateApiGroupReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateApiGroupReqValidator().Validate(ctx, m, opts...)
}

func (m *EvaluateApiGroupReq) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetApiGroupDRefInfo()

}

// GetDRefInfo for the field's type
func (m *EvaluateApiGroupReq) GetApiGroupDRefInfo() ([]db.DRefInfo, error) {
	if m.GetApiGroup() == nil {
		return nil, nil
	}

	drInfos, err := m.GetApiGroup().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetApiGroup().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "api_group." + dri.DRField
	}
	return drInfos, err

}

type ValidateEvaluateApiGroupReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateApiGroupReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateApiGroupReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateApiGroupReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_group"]; exists {

		vOpts := append(opts, db.WithValidateField("api_group"))
		if err := fv(ctx, m.GetApiGroup(), vOpts...); err != nil {
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
var DefaultEvaluateApiGroupReqValidator = func() *ValidateEvaluateApiGroupReq {
	v := &ValidateEvaluateApiGroupReq{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_group"] = GlobalSpecTypeValidator().Validate

	return v
}()

func EvaluateApiGroupReqValidator() db.Validator {
	return DefaultEvaluateApiGroupReqValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateApiGroupRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateApiGroupRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateApiGroupRsp) DeepCopy() *EvaluateApiGroupRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateApiGroupRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateApiGroupRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateApiGroupRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateApiGroupRspValidator().Validate(ctx, m, opts...)
}

func (m *EvaluateApiGroupRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetApiGroupDRefInfo()

}

// GetDRefInfo for the field's type
func (m *EvaluateApiGroupRsp) GetApiGroupDRefInfo() ([]db.DRefInfo, error) {
	if m.GetApiGroup() == nil {
		return nil, nil
	}

	drInfos, err := m.GetApiGroup().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetApiGroup().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "api_group." + dri.DRField
	}
	return drInfos, err

}

type ValidateEvaluateApiGroupRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateApiGroupRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateApiGroupRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateApiGroupRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_group"]; exists {

		vOpts := append(opts, db.WithValidateField("api_group"))
		if err := fv(ctx, m.GetApiGroup(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["apieps_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("apieps_timestamp"))
		if err := fv(ctx, m.GetApiepsTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["matching_api_endpoints"]; exists {

		vOpts := append(opts, db.WithValidateField("matching_api_endpoints"))
		for idx, item := range m.GetMatchingApiEndpoints() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateApiGroupRspValidator = func() *ValidateEvaluateApiGroupRsp {
	v := &ValidateEvaluateApiGroupRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_group"] = GlobalSpecTypeValidator().Validate

	v.FldValidators["matching_api_endpoints"] = ApiEndpointValidator().Validate

	return v
}()

func EvaluateApiGroupRspValidator() db.Validator {
	return DefaultEvaluateApiGroupRspValidator
}

// augmented methods on protoc/std generated struct

func (m *GetApiGroupsStatsReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetApiGroupsStatsReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetApiGroupsStatsReq) DeepCopy() *GetApiGroupsStatsReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetApiGroupsStatsReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetApiGroupsStatsReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetApiGroupsStatsReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetApiGroupsStatsReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetApiGroupsStatsReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetApiGroupsStatsReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetApiGroupsStatsReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetApiGroupsStatsReq got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetApiGroupsStatsReqValidator = func() *ValidateGetApiGroupsStatsReq {
	v := &ValidateGetApiGroupsStatsReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetApiGroupsStatsReqValidator() db.Validator {
	return DefaultGetApiGroupsStatsReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetApiGroupsStatsRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetApiGroupsStatsRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetApiGroupsStatsRsp) DeepCopy() *GetApiGroupsStatsRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetApiGroupsStatsRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetApiGroupsStatsRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetApiGroupsStatsRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetApiGroupsStatsRspValidator().Validate(ctx, m, opts...)
}

type ValidateGetApiGroupsStatsRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetApiGroupsStatsRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetApiGroupsStatsRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetApiGroupsStatsRsp got type %s", t)
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
var DefaultGetApiGroupsStatsRspValidator = func() *ValidateGetApiGroupsStatsRsp {
	v := &ValidateGetApiGroupsStatsRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetApiGroupsStatsRspValidator() db.Validator {
	return DefaultGetApiGroupsStatsRspValidator
}
