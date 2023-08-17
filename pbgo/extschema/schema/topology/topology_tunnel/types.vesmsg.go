// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package topology_tunnel

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
	if fdrInfos, err := m.GetTopologyMetadataDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetTopologyMetadataDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetTopologySpecDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetTopologySpecDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetTopologyMetadataDRefInfo() ([]db.DRefInfo, error) {
	if m.GetTopologyMetadata() == nil {
		return nil, nil
	}

	drInfos, err := m.GetTopologyMetadata().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetTopologyMetadata().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "topology_metadata." + dri.DRField
	}
	return drInfos, err

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetTopologySpecDRefInfo() ([]db.DRefInfo, error) {
	if m.GetTopologySpec() == nil {
		return nil, nil
	}

	drInfos, err := m.GetTopologySpec().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetTopologySpec().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "topology_spec." + dri.DRField
	}
	return drInfos, err

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

	if fv, exists := v.FldValidators["topology_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("topology_metadata"))
		if err := fv(ctx, m.GetTopologyMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["topology_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("topology_spec"))
		if err := fv(ctx, m.GetTopologySpec(), vOpts...); err != nil {
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
