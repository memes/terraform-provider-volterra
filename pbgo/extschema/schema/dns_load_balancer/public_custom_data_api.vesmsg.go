//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package dns_load_balancer

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

func (m *DNSLBHealthStatusListRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBHealthStatusListRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBHealthStatusListRequest) DeepCopy() *DNSLBHealthStatusListRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBHealthStatusListRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBHealthStatusListRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBHealthStatusListRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBHealthStatusListRequestValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBHealthStatusListRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBHealthStatusListRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBHealthStatusListRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBHealthStatusListRequest got type %s", t)
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
var DefaultDNSLBHealthStatusListRequestValidator = func() *ValidateDNSLBHealthStatusListRequest {
	v := &ValidateDNSLBHealthStatusListRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBHealthStatusListRequestValidator() db.Validator {
	return DefaultDNSLBHealthStatusListRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBHealthStatusListResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBHealthStatusListResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBHealthStatusListResponse) DeepCopy() *DNSLBHealthStatusListResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBHealthStatusListResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBHealthStatusListResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBHealthStatusListResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBHealthStatusListResponseValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBHealthStatusListResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBHealthStatusListResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBHealthStatusListResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBHealthStatusListResponse got type %s", t)
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
var DefaultDNSLBHealthStatusListResponseValidator = func() *ValidateDNSLBHealthStatusListResponse {
	v := &ValidateDNSLBHealthStatusListResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBHealthStatusListResponseValidator() db.Validator {
	return DefaultDNSLBHealthStatusListResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBHealthStatusListResponseItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBHealthStatusListResponseItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBHealthStatusListResponseItem) DeepCopy() *DNSLBHealthStatusListResponseItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBHealthStatusListResponseItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBHealthStatusListResponseItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBHealthStatusListResponseItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBHealthStatusListResponseItemValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBHealthStatusListResponseItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBHealthStatusListResponseItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBHealthStatusListResponseItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBHealthStatusListResponseItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_lb_pools_status_summary"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_lb_pools_status_summary"))
		for idx, item := range m.GetDnsLbPoolsStatusSummary() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDNSLBHealthStatusListResponseItemValidator = func() *ValidateDNSLBHealthStatusListResponseItem {
	v := &ValidateDNSLBHealthStatusListResponseItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBHealthStatusListResponseItemValidator() db.Validator {
	return DefaultDNSLBHealthStatusListResponseItemValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBHealthStatusRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBHealthStatusRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBHealthStatusRequest) DeepCopy() *DNSLBHealthStatusRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBHealthStatusRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBHealthStatusRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBHealthStatusRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBHealthStatusRequestValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBHealthStatusRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBHealthStatusRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBHealthStatusRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBHealthStatusRequest got type %s", t)
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
var DefaultDNSLBHealthStatusRequestValidator = func() *ValidateDNSLBHealthStatusRequest {
	v := &ValidateDNSLBHealthStatusRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBHealthStatusRequestValidator() db.Validator {
	return DefaultDNSLBHealthStatusRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBHealthStatusResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBHealthStatusResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBHealthStatusResponse) DeepCopy() *DNSLBHealthStatusResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBHealthStatusResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBHealthStatusResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBHealthStatusResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBHealthStatusResponseValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBHealthStatusResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBHealthStatusResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBHealthStatusResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBHealthStatusResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_lb_pool_items"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_lb_pool_items"))
		for idx, item := range m.GetDnsLbPoolItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDNSLBHealthStatusResponseValidator = func() *ValidateDNSLBHealthStatusResponse {
	v := &ValidateDNSLBHealthStatusResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBHealthStatusResponseValidator() db.Validator {
	return DefaultDNSLBHealthStatusResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBPoolHealthStatusListResponseItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBPoolHealthStatusListResponseItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBPoolHealthStatusListResponseItem) DeepCopy() *DNSLBPoolHealthStatusListResponseItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBPoolHealthStatusListResponseItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBPoolHealthStatusListResponseItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBPoolHealthStatusListResponseItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBPoolHealthStatusListResponseItemValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBPoolHealthStatusListResponseItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBPoolHealthStatusListResponseItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBPoolHealthStatusListResponseItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBPoolHealthStatusListResponseItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_lb_pool_members_status_summary"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_lb_pool_members_status_summary"))
		for idx, item := range m.GetDnsLbPoolMembersStatusSummary() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDNSLBPoolHealthStatusListResponseItemValidator = func() *ValidateDNSLBPoolHealthStatusListResponseItem {
	v := &ValidateDNSLBPoolHealthStatusListResponseItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBPoolHealthStatusListResponseItemValidator() db.Validator {
	return DefaultDNSLBPoolHealthStatusListResponseItemValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBPoolHealthStatusRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBPoolHealthStatusRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBPoolHealthStatusRequest) DeepCopy() *DNSLBPoolHealthStatusRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBPoolHealthStatusRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBPoolHealthStatusRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBPoolHealthStatusRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBPoolHealthStatusRequestValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBPoolHealthStatusRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBPoolHealthStatusRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBPoolHealthStatusRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBPoolHealthStatusRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_lb_name"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_lb_name"))
		if err := fv(ctx, m.GetDnsLbName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["dns_lb_pool_name"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_lb_pool_name"))
		if err := fv(ctx, m.GetDnsLbPoolName(), vOpts...); err != nil {
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
var DefaultDNSLBPoolHealthStatusRequestValidator = func() *ValidateDNSLBPoolHealthStatusRequest {
	v := &ValidateDNSLBPoolHealthStatusRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBPoolHealthStatusRequestValidator() db.Validator {
	return DefaultDNSLBPoolHealthStatusRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBPoolHealthStatusResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBPoolHealthStatusResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBPoolHealthStatusResponse) DeepCopy() *DNSLBPoolHealthStatusResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBPoolHealthStatusResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBPoolHealthStatusResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBPoolHealthStatusResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBPoolHealthStatusResponseValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBPoolHealthStatusResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBPoolHealthStatusResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBPoolHealthStatusResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBPoolHealthStatusResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_lb_pool_member_items"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_lb_pool_member_items"))
		for idx, item := range m.GetDnsLbPoolMemberItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDNSLBPoolHealthStatusResponseValidator = func() *ValidateDNSLBPoolHealthStatusResponse {
	v := &ValidateDNSLBPoolHealthStatusResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBPoolHealthStatusResponseValidator() db.Validator {
	return DefaultDNSLBPoolHealthStatusResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *DNSLBPoolMemberHealthStatusListResponseItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DNSLBPoolMemberHealthStatusListResponseItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DNSLBPoolMemberHealthStatusListResponseItem) DeepCopy() *DNSLBPoolMemberHealthStatusListResponseItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DNSLBPoolMemberHealthStatusListResponseItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DNSLBPoolMemberHealthStatusListResponseItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DNSLBPoolMemberHealthStatusListResponseItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DNSLBPoolMemberHealthStatusListResponseItemValidator().Validate(ctx, m, opts...)
}

type ValidateDNSLBPoolMemberHealthStatusListResponseItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDNSLBPoolMemberHealthStatusListResponseItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DNSLBPoolMemberHealthStatusListResponseItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DNSLBPoolMemberHealthStatusListResponseItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["health_check_type"]; exists {

		vOpts := append(opts, db.WithValidateField("health_check_type"))
		if err := fv(ctx, m.GetHealthCheckType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		for idx, item := range m.GetStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDNSLBPoolMemberHealthStatusListResponseItemValidator = func() *ValidateDNSLBPoolMemberHealthStatusListResponseItem {
	v := &ValidateDNSLBPoolMemberHealthStatusListResponseItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DNSLBPoolMemberHealthStatusListResponseItemValidator() db.Validator {
	return DefaultDNSLBPoolMemberHealthStatusListResponseItemValidator
}

// augmented methods on protoc/std generated struct

func (m *HealthStatusSummary) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *HealthStatusSummary) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *HealthStatusSummary) DeepCopy() *HealthStatusSummary {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &HealthStatusSummary{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *HealthStatusSummary) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *HealthStatusSummary) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return HealthStatusSummaryValidator().Validate(ctx, m, opts...)
}

type ValidateHealthStatusSummary struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateHealthStatusSummary) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*HealthStatusSummary)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *HealthStatusSummary got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["count"]; exists {

		vOpts := append(opts, db.WithValidateField("count"))
		for idx, item := range m.GetCount() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultHealthStatusSummaryValidator = func() *ValidateHealthStatusSummary {
	v := &ValidateHealthStatusSummary{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func HealthStatusSummaryValidator() db.Validator {
	return DefaultHealthStatusSummaryValidator
}
