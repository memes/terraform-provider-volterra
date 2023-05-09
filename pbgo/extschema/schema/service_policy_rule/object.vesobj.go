// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package service_policy_rule

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	google_protobuf "github.com/gogo/protobuf/types"
	multierror "github.com/hashicorp/go-multierror"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/store"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	"github.com/google/uuid"
	"gopkg.volterra.us/stdlib/db/sro"
)

const (
	ObjectDefTblName = "ves.io.schema.service_policy_rule.Object.default"
	ObjectType       = "ves.io.schema.service_policy_rule.Object"

	StatusObjectDefTblName = "ves.io.schema.service_policy_rule.StatusObject.default"
	StatusObjectType       = "ves.io.schema.service_policy_rule.StatusObject"
)

// augmented methods on protoc/std generated struct
func (e *Object) Type() string {
	return "ves.io.schema.service_policy_rule.Object"
}

func (e *Object) ToEntry() db.Entry {
	return NewDBObject(e, db.OpWithNoCopy())
}

func LocateObject(ctx context.Context, locator db.EntryLocator, uid, tenant, namespace, name string, opts ...db.FindEntryOpt) (*DBObject, error) {
	timestamp, err := google_protobuf.TimestampProto(time.Now())
	if err != nil {
		return nil, errors.Wrapf(err, "%s: LocateObject", uid)
	}
	if uid != "" {
		obj, exist, err := FindObject(ctx, locator, uid, opts...)
		if err != nil {
			return nil, errors.Wrapf(err, "%s: LocateObject", uid)
		}
		if exist && obj != nil {
			obj.SystemMetadata.ModificationTimestamp = timestamp
			return obj, nil
		}
	} else {
		uid = uuid.New().String()
	}

	sysMD := &ves_io_schema.SystemObjectMetaType{
		Uid:                   uid,
		Tenant:                tenant,
		CreatorClass:          locator.GetCreatorClass(),
		CreatorId:             locator.GetCreatorID(),
		CreationTimestamp:     timestamp,
		ModificationTimestamp: timestamp,
	}
	obj := NewDBObject(nil)
	obj.SetObjUid(uid)
	obj.SetObjName(name)
	obj.SetObjNamespace(namespace)
	obj.SetObjSystemMetadata(sysMD)
	obj.Spec = &SpecType{}
	return obj, nil
}

func FindObject(ctx context.Context, finder db.EntryFinder, key string, opts ...db.FindEntryOpt) (*DBObject, bool, error) {
	e, exist, err := finder.FindEntry(ctx, ObjectDefTblName, key, opts...)
	if !exist || err != nil {
		return nil, exist, err
	}
	obj, ok := e.(*DBObject)
	if !ok {
		return nil, false, fmt.Errorf("Cannot convert entry to object")
	}
	return obj, exist, err
}

func ListObject(ctx context.Context, lister db.EntryLister, opts ...db.ListEntriesOpt) ([]*DBObject, error) {
	var (
		oList []*DBObject
		merr  *multierror.Error
	)
	eList, err := lister.ListEntries(ctx, ObjectDefTblName, opts...)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	for _, e := range eList {
		obj, ok := e.(*DBObject)
		if ok {
			oList = append(oList, obj)
		} else {
			merr = multierror.Append(merr, fmt.Errorf("Cannot convert entry to %s object", ObjectType))
		}
	}
	return oList, errors.ErrOrNil(merr)
}

func (o *Object) DeepCopy() *Object {
	if o == nil {
		return nil
	}
	ser, err := o.Marshal()
	if err != nil {
		return nil
	}
	c := &Object{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (e *Object) ToJSON() (string, error) {
	return codec.ToJSON(e)
}

func (e *Object) ToYAML() (string, error) {
	return codec.ToYAML(e)
}

func (e *Object) GetTraceInfo() string {
	sysMD := e.GetSystemMetadata()
	if sysMD == nil {
		return ""
	}
	return sysMD.GetTraceInfo()
}

// A struct wrapping protoc/std generated struct with additional capabilities
// forming a db.Entry
type DBObject struct {
	// Anonymous embed of standard protobuf generated struct
	*Object

	tbl db.Table
}

// GetObjectIndexers returns the associated store.Indexers for Object
func GetObjectIndexers() store.Indexers {

	return nil

}

func (e *DBObject) GetDB() (*db.DB, error) {
	if e.tbl == nil {
		return nil, fmt.Errorf("Entry has no table")
	}
	return e.tbl.GetDB(), nil
}

// Implement ves.io/stdlib/db.Entry interface
func (e *DBObject) Key(opts ...db.KeyOpt) (string, error) {
	ko := db.NewKeyOpts(opts...)
	if ko.Public {
		md := e.GetMetadata()
		if md == nil {
			return "", fmt.Errorf("Metadata is nil")
		}
		return fmt.Sprintf("%s/%s", md.GetNamespace(), md.GetName()), nil
	} else {
		if e.GetSystemMetadata() == nil {
			return "", fmt.Errorf("SystemMetadata is nil")
		}
		return e.GetSystemMetadata().GetUid(), nil
	}
}

func (e *DBObject) Type() string {
	return "ves.io.schema.service_policy_rule.Object"
}

func (e *DBObject) DeepCopy() db.Entry {
	if e == nil {
		return nil
	}
	n := NewDBObject(e.Object)
	n.tbl = e.tbl
	return n
}

func (e *DBObject) MarshalBytes() ([]byte, error) {
	return e.Marshal()
}

func (e *DBObject) UnmarshalBytes(b []byte) error {
	return e.Unmarshal(b)
}

func (e *DBObject) Sample(r *rand.Rand) (db.Entry, error) {
	uid := uuid.New().String()
	o := &Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Name:      uuid.New().String(),
			Namespace: uuid.New().String(),
			Uid:       uid,
		},
		SystemMetadata: &ves_io_schema.SystemObjectMetaType{
			Uid:    uid,
			Tenant: uuid.New().String(),
		},
		Spec: &SpecType{},
	}

	return &DBObject{o, e.tbl}, nil
}

func (e *DBObject) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectValidator().Validate(ctx, e.Object, opts...)
}

func (e *DBObject) SetBlob(ctx context.Context, bID string, bVal interface{}, opts ...db.BlobOpt) error {
	db, err := e.GetDB()
	if err != nil {
		return errors.Wrap(err, "SetBlob")
	}
	key, err := e.Key()
	if err != nil {
		return errors.Wrap(err, "SetBlob accessing key")
	}
	err = db.SetEntryBlob(ctx, key, e.Type(), bID, bVal, opts...)
	if err != nil {
		return errors.Wrap(err, "SetBlob setting in db")
	}
	return nil
}

func (e *DBObject) ClrBlob(ctx context.Context, bID string, opts ...db.BlobOpt) error {
	db, err := e.GetDB()
	if err != nil {
		return errors.Wrap(err, "ClrBlob")
	}
	key, err := e.Key()
	if err != nil {
		return errors.Wrap(err, "ClrBlob accessing key")
	}
	err = db.ClrEntryBlob(ctx, key, e.Type(), bID, opts...)
	if err != nil {
		return errors.Wrap(err, "ClrBlob clearing in db")
	}
	return nil
}

func (e *DBObject) GetBlob(ctx context.Context, bID string, opts ...db.BlobOpt) (interface{}, error) {
	db, err := e.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlob")
	}
	key, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlob accessing key")
	}
	return db.GetEntryBlob(ctx, key, e.Type(), bID, opts...)
}

func (e *DBObject) GetBlobs(ctx context.Context, opts ...db.BlobOpt) (map[string]interface{}, error) {
	db, err := e.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlobs")
	}
	key, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlobs accessing key")
	}
	return db.GetEntryBlobs(ctx, key, e.Type(), opts...)
}

func (e *DBObject) IsDeleted() (bool, error) {
	db, err := e.GetDB()
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted")
	}
	key, err := e.Key()
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted accessing key")
	}
	isDel, err := db.IsEntryDeleted(key, e.Type())
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted accessing db")
	}
	return isDel, nil
}

// Implement ves.io/stdlib/db.EntryPvt interface
func (e *DBObject) SetTable(tbl db.Table) {
	e.tbl = tbl
}

func (e *DBObject) GetDRefInfo() ([]db.DRefInfo, error) {
	if e == nil {
		return nil, nil
	}
	refrUID, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetDRefInfo, error in key")
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := e.GetSpecDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetSpecDRefInfo() FAILED")
	} else {
		for i := range fdrInfos {
			dri := &fdrInfos[i]
			// Convert Spec.LcSpec.vnRefs to ves.io.examplesvc.objectone.Object.Spec.LcSpec.vnRefs
			dri.DRField = "ves.io.schema.service_policy_rule.Object." + dri.DRField
			dri.RefrType = e.Type()
			dri.RefrUID = refrUID

			// convert any ref_to schema annotation specified by kind value to type value
			if !strings.HasPrefix(dri.RefdType, "ves.io") {
				d, err := e.GetDB()
				if err != nil {
					return nil, errors.Wrap(err, "Cannot find db for entry to resolve kind to type")
				}
				refdType, err := d.TypeForEntryKind(dri.RefrType, dri.RefrUID, dri.RefdType)
				if err != nil {
					return nil, errors.Wrap(err, fmt.Sprintf("Cannot convert kind %s to type", dri.RefdType))
				}
				dri.RefdType = refdType
			}
		}
		drInfos = append(drInfos, fdrInfos...)
	}
	if fdrInfos, err := e.GetSystemMetadataDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetSystemMetadataDRefInfo() FAILED")
	} else {
		for i := range fdrInfos {
			dri := &fdrInfos[i]
			// Convert Spec.LcSpec.vnRefs to ves.io.examplesvc.objectone.Object.Spec.LcSpec.vnRefs
			dri.DRField = "ves.io.schema.service_policy_rule.Object." + dri.DRField
			dri.RefrType = e.Type()
			dri.RefrUID = refrUID

			// convert any ref_to schema annotation specified by kind value to type value
			if !strings.HasPrefix(dri.RefdType, "ves.io") {
				d, err := e.GetDB()
				if err != nil {
					return nil, errors.Wrap(err, "Cannot find db for entry to resolve kind to type")
				}
				refdType, err := d.TypeForEntryKind(dri.RefrType, dri.RefrUID, dri.RefdType)
				if err != nil {
					return nil, errors.Wrap(err, fmt.Sprintf("Cannot convert kind %s to type", dri.RefdType))
				}
				dri.RefdType = refdType
			}
		}
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (e *DBObject) ToStore() store.Entry {
	return e.Object
}

func (e *DBObject) ToJSON() (string, error) {
	return e.ToStore().ToJSON()
}

func (e *DBObject) ToYAML() (string, error) {
	return e.ToStore().ToYAML()
}

func (e *DBObject) GetTable() db.Table {
	return e.tbl
}

func NewDBObject(o *Object, opts ...db.OpOption) *DBObject {
	op := db.NewOpFrom(opts...)
	if o == nil {
		return &DBObject{Object: &Object{}}
	}
	obj := o
	if !op.NoCopy() {
		obj = o.DeepCopy()
	}
	return &DBObject{Object: obj}
}

func NewEntryObject(opts ...db.OpOption) db.Entry {
	op := db.NewOpFrom(opts...)
	s := op.StoreEntry()
	switch v := s.(type) {
	case nil:
		return NewDBObject(nil, opts...)
	case *Object:
		return NewDBObject(v, opts...)
	}
	return nil
}

// GetDRefInfo for the field's type
func (e *DBObject) GetSpecDRefInfo() ([]db.DRefInfo, error) {
	if e.GetSpec() == nil {
		return nil, nil
	}

	drInfos, err := e.GetSpec().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetSpec().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "spec." + dri.DRField
	}
	return drInfos, err

}

// GetDRefInfo for the field's type
func (e *DBObject) GetSystemMetadataDRefInfo() ([]db.DRefInfo, error) {
	if e.GetSystemMetadata() == nil {
		return nil, nil
	}

	drInfos, err := e.GetSystemMetadata().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetSystemMetadata().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "system_metadata." + dri.DRField
	}
	return drInfos, err

}

// Implement sro.SRO interface
func (o *DBObject) GetObjMetadata() sro.ObjectMetadata {
	return o.GetMetadata()
}

func (o *DBObject) SetObjMetadata(in sro.ObjectMetadata) error {
	if in == nil {
		o.Metadata = nil
		return nil
	}

	m, ok := in.(*ves_io_schema.ObjectMetaType)
	if !ok {
		return fmt.Errorf("Error: SetObjMetadata expected *ObjectMetaType, got %T", in)
	}
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjSystemMetadata() sro.SystemMetadata {
	if o.GetSystemMetadata() == nil {
		return nil
	}
	return o.GetSystemMetadata()
}

func (o *DBObject) SetObjSystemMetadata(in sro.SystemMetadata) error {
	if in == nil {
		o.SystemMetadata = nil
		return nil
	}

	m, ok := in.(*ves_io_schema.SystemObjectMetaType)
	if !ok {
		return fmt.Errorf("Error: SetObjSystemMetadata expected *SystemObjectMetaType, got %T", in)
	}
	o.SystemMetadata = m
	return nil
}

func (o *DBObject) GetObjSpec() sro.Spec {
	if o.GetSpec() == nil {
		return nil
	}
	return o.GetSpec()
}

func (o *DBObject) SetObjSpec(in sro.Spec) error {
	if in == nil {
		o.Spec = nil
		return nil
	}

	m, ok := in.(*SpecType)
	if !ok {
		return fmt.Errorf("Error: SetObjSpec expected *SpecType, got %T", in)
	}
	o.Spec = m
	return nil
}

func FindObjectStatus(ctx context.Context, d db.Interface, objUid string) ([]*StatusObject, error) {
	statusDBEntries, err := d.GetEntryBackrefs(ctx, objUid, ObjectType, db.WithBackrefTypes([]string{"ves.io.schema.service_policy_rule.StatusObject"}))
	if err != nil {
		return nil, err
	}
	var merr *multierror.Error
	var statusObjs []*StatusObject
	for _, statusDBEntry := range statusDBEntries {
		statusEntry := statusDBEntry.ToStore()
		statusObj, ok := statusEntry.(*StatusObject)
		if !ok {
			merr = multierror.Append(merr, fmt.Errorf("Status Backref expected *StatusObject, got %T: %v", statusEntry, statusEntry))
			continue
		}
		statusObjs = append(statusObjs, statusObj)
	}
	return statusObjs, errors.ErrOrNil(merr)
}

// SetObjectRef sets reference to a configuration object
func (o *StatusObject) SetObjectRef(objKind, objUid string) error {
	if len(o.GetObjectRefs()) != 0 {
		return fmt.Errorf("StatusObject already has a reference to %v", o.GetObjectRefs())
	}
	o.ObjectRefs = append(o.ObjectRefs, &ves_io_schema.ObjectRefType{Kind: objKind, Uid: objUid})
	return nil
}

func (o *StatusObject) GetStatusObjMetadata() sro.StatusObjectMetadata {
	return o.GetMetadata()
}

func (o *StatusObject) SetStatusObjMetadata(md sro.StatusObjectMetadata) {
	if o == nil {
		return
	}
	if o.Metadata == nil {
		o.Metadata = &ves_io_schema.StatusMetaType{}
	}
	o.Metadata = md.(*ves_io_schema.StatusMetaType)
}

// GenerateUuidv5() returns a deterministic UUIDv5 based on the unique semantic key of status object
func (o *StatusObject) GenerateUuidV5() (string, error) {
	statusObjectMetaData := o.GetStatusObjMetadata()
	creatorClass := statusObjectMetaData.GetCreatorClass()
	creatorId := statusObjectMetaData.GetCreatorId()
	statusId := statusObjectMetaData.GetStatusId()
	objectRefArray := o.GetObjectRefs()
	if len(objectRefArray) == 0 {
		return "", fmt.Errorf("StatusObject does not have a reference to config object.")
	}
	configObjectUuid := objectRefArray[0].Uid
	configObjectKind := objectRefArray[0].Kind
	keyFields := []string{creatorClass, creatorId, statusId, configObjectKind, configObjectUuid}
	secKey := strings.Join(keyFields, "::")
	newUuid := uuid.NewSHA1(uuid.NameSpaceOID, []byte(secKey)).String()
	return newUuid, nil
}

// SetUuidV5 sets deterministic uuid for a status object.
func (o *StatusObject) SetUuidV5() error {
	if o == nil {
		return fmt.Errorf("Status object is nil")
	}
	uuidV5, err := o.GenerateUuidV5()
	if err != nil {
		return err
	}
	o.GetMetadata().SetUid(uuidV5)
	return nil
}

// GetVtrpId returns vtrpId of the status object.
func (o *StatusObject) GetVtrpId() string {
	return o.GetMetadata().GetVtrpId()
}

// SetVtrpId sets vtrpId of the status object.
func (o *StatusObject) SetVtrpId(id string) {
	o.GetMetadata().SetVtrpId(id)
}

// GetVtrpStale returns true if the object is stale in Mars
func (o *StatusObject) GetVtrpStale() bool {
	return o.GetMetadata().GetVtrpStale()
}

// SetVtrpStale sets vtrpStale on the status object
func (o *StatusObject) SetVtrpStale(isStale bool) {
	o.GetMetadata().SetVtrpStale(isStale)
}

func (o *StatusObject) GetStatusObjConditions() []sro.StatusObjectCondition {
	if o == nil {
		return nil
	}
	return ves_io_schema.ToStatusObjectConditions(o.GetConditions())
}

func (o *StatusObject) SetStatusObjConditions(socSet []sro.StatusObjectCondition) {
	if o == nil {
		return
	}
	o.Conditions = ves_io_schema.FromStatusObjectConditions(socSet)
}

func (o *DBObject) GetObjType() string {
	return o.Type()
}

// GetObjUid returns uuid from source-of-truth, in systemMetadata
func (o *DBObject) GetObjUid() string {
	return o.GetSystemMetadata().GetUid()
}

// GetObjTenant returns tenant from source-of-truth, in systemMetadata
func (o *DBObject) GetObjTenant() string {
	return o.GetSystemMetadata().GetTenant()
}

// GetObjCreatorClass returns creator-class from systemMetadata
func (o *DBObject) GetObjCreatorClass() string {
	return o.GetSystemMetadata().GetCreatorClass()
}

// GetObjectIndex returns object-index from systemMetadata
func (o *DBObject) GetObjectIndex() uint32 {
	return o.GetSystemMetadata().GetObjectIndex()
}

// SetObjUid sets uuid as a hint, in Metadata
func (o *DBObject) SetObjUid(u string) error {
	// TODO: make sure 'u' is of uuid form
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Uid = u
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjName() string {
	return o.GetMetadata().GetName()
}

func (o *DBObject) SetObjName(n string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Name = n
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjNamespace() string {
	return o.GetMetadata().GetNamespace()
}

func (o *DBObject) SetObjNamespace(ns string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Namespace = ns
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjLabels() map[string]string {
	return o.GetMetadata().GetLabels()
}

func (o *DBObject) SetObjLabels(l map[string]string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Labels = l
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjAnnotations() map[string]string {
	return o.GetMetadata().GetAnnotations()
}

func (o *DBObject) SetObjAnnotations(a map[string]string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Annotations = a
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjDescription() string {
	return o.GetMetadata().GetDescription()
}

func (o *DBObject) SetObjDescription(d string) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Description = d
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjDisable() bool {
	return o.GetMetadata().GetDisable()
}

func (o *DBObject) SetObjDisable(d bool) error {
	m := o.GetMetadata()
	if m == nil {
		m = &ves_io_schema.ObjectMetaType{}
	}
	m.Disable = d
	o.Metadata = m
	return nil
}

func (o *DBObject) GetObjSREDisable() bool {
	return o.GetSystemMetadata().GetSreDisable()
}

func (o *DBObject) SetObjSREDisable(d bool) error {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.SreDisable = d
	return nil
}

func (o *DBObject) SetObjCreator(cls, inst string) error {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.CreatorClass = cls
	m.CreatorId = inst
	o.SystemMetadata = m
	return nil
}

func (o *DBObject) SetObjectIndex(idx uint32) error {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
	}
	m.ObjectIndex = idx
	o.SystemMetadata = m
	return nil
}

func (o *DBObject) GetObjFinalizers() []string {
	return o.GetSystemMetadata().GetFinalizers()
}

func (o *DBObject) SetObjFinalizers(values ...string) error {
	m := o.GetSystemMetadata()
	if m == nil {
		return fmt.Errorf("Object has nil system_metadata")
	}
	m.Finalizers = values
	return nil
}

func (o *DBObject) GetObjPendingInitializers() []string {
	initializers := o.GetSystemMetadata().GetInitializers()
	var pending []string
	for _, p := range initializers.GetPending() {
		pending = append(pending, p.GetName())
	}
	return pending
}

func (o *DBObject) SetObjPendingInitializers(pending ...string) {
	m := o.GetSystemMetadata()
	if m == nil {
		m = &ves_io_schema.SystemObjectMetaType{}
		o.SystemMetadata = m
	}
	initializers := m.GetInitializers()
	if initializers == nil {
		initializers = &ves_io_schema.InitializersType{}
		m.Initializers = initializers
	}
	var pendingInitializers []*ves_io_schema.InitializerType
	for _, p := range pending {
		pendingInitializers = append(pendingInitializers, &ves_io_schema.InitializerType{Name: p})
	}
	initializers.Pending = pendingInitializers
}

func (o *DBObject) IsSpecEqual(other sro.SRO) bool {
	otherObjSpec := other.GetObjSpec()
	otherSpec, ok := otherObjSpec.(*SpecType)
	if !ok {
		return false
	}

	return o.GetSpec().Equal(otherSpec)
}

// GetVtrpId returns vtrpId of the object.
func (o *DBObject) GetVtrpId() string {
	return o.GetSystemMetadata().GetVtrpId()
}

// SetVtrpId sets vtrpId of the object.
func (o *DBObject) SetVtrpId(id string) {
	o.GetSystemMetadata().SetVtrpId(id)
}

// GetVtrpStale returns true if the object is stale in Mars
func (o *DBObject) GetVtrpStale() bool {
	return o.GetSystemMetadata().GetVtrpStale()
}

// SetVtrpStale sets vtrpStale on the object
func (o *DBObject) SetVtrpStale(isStale bool) {
	o.GetSystemMetadata().SetVtrpStale(isStale)
}

type ValidateObject struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObject) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	e, ok := pm.(*Object)
	if !ok {
		switch t := pm.(type) {
		default:
			return fmt.Errorf("Expected type *Object got type %s", t)
		}
	}
	if e == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, e.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, e.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["system_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("system_metadata"))
		if err := fv(ctx, e.GetSystemMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectValidator = func() *ValidateObject {
	v := &ValidateObject{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectMetaTypeValidator().Validate

	v.FldValidators["system_metadata"] = ves_io_schema.SystemObjectMetaTypeValidator().Validate

	v.FldValidators["spec"] = SpecTypeValidator().Validate

	return v
}()

func ObjectValidator() db.Validator {
	return DefaultObjectValidator
}

// augmented methods on protoc/std generated struct
func (e *StatusObject) Type() string {
	return "ves.io.schema.service_policy_rule.StatusObject"
}

func (e *StatusObject) ToEntry() db.Entry {
	return NewDBStatusObject(e, db.OpWithNoCopy())
}

func FindStatusObject(ctx context.Context, finder db.EntryFinder, key string, opts ...db.FindEntryOpt) (*DBStatusObject, bool, error) {
	e, exist, err := finder.FindEntry(ctx, StatusObjectDefTblName, key, opts...)
	if !exist || err != nil {
		return nil, exist, err
	}
	obj, ok := e.(*DBStatusObject)
	if !ok {
		return nil, false, fmt.Errorf("Cannot convert entry to object")
	}
	return obj, exist, err
}

func ListStatusObject(ctx context.Context, lister db.EntryLister, opts ...db.ListEntriesOpt) ([]*DBStatusObject, error) {
	var (
		oList []*DBStatusObject
		merr  *multierror.Error
	)
	eList, err := lister.ListEntries(ctx, StatusObjectDefTblName, opts...)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	for _, e := range eList {
		obj, ok := e.(*DBStatusObject)
		if ok {
			oList = append(oList, obj)
		} else {
			merr = multierror.Append(merr, fmt.Errorf("Cannot convert entry to %s object", StatusObjectType))
		}
	}
	return oList, errors.ErrOrNil(merr)
}

func (o *StatusObject) DeepCopy() *StatusObject {
	if o == nil {
		return nil
	}
	ser, err := o.Marshal()
	if err != nil {
		return nil
	}
	c := &StatusObject{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (e *StatusObject) ToJSON() (string, error) {
	return codec.ToJSON(e)
}

func (e *StatusObject) ToYAML() (string, error) {
	return codec.ToYAML(e)
}

// A struct wrapping protoc/std generated struct with additional capabilities
// forming a db.Entry
type DBStatusObject struct {
	// Anonymous embed of standard protobuf generated struct
	*StatusObject

	tbl db.Table
}

// GetStatusObjectIndexers returns the associated store.Indexers for StatusObject
func GetStatusObjectIndexers() store.Indexers {

	return nil

}

func (e *DBStatusObject) GetDB() (*db.DB, error) {
	if e.tbl == nil {
		return nil, fmt.Errorf("Entry has no table")
	}
	return e.tbl.GetDB(), nil
}

// Implement ves.io/stdlib/db.Entry interface
func (e *DBStatusObject) Key(opts ...db.KeyOpt) (string, error) {
	return e.GetMetadata().GetUid(), nil
}

func (e *DBStatusObject) Type() string {
	return "ves.io.schema.service_policy_rule.StatusObject"
}

func (e *DBStatusObject) DeepCopy() db.Entry {
	if e == nil {
		return nil
	}
	n := NewDBStatusObject(e.StatusObject)
	n.tbl = e.tbl
	return n
}

func (e *DBStatusObject) MarshalBytes() ([]byte, error) {
	return e.Marshal()
}

func (e *DBStatusObject) UnmarshalBytes(b []byte) error {
	return e.Unmarshal(b)
}

func (e *DBStatusObject) Sample(r *rand.Rand) (db.Entry, error) {

	o := &StatusObject{}

	return &DBStatusObject{o, e.tbl}, nil
}

func (e *DBStatusObject) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return StatusObjectValidator().Validate(ctx, e.StatusObject, opts...)
}

func (e *DBStatusObject) SetBlob(ctx context.Context, bID string, bVal interface{}, opts ...db.BlobOpt) error {
	db, err := e.GetDB()
	if err != nil {
		return errors.Wrap(err, "SetBlob")
	}
	key, err := e.Key()
	if err != nil {
		return errors.Wrap(err, "SetBlob accessing key")
	}
	err = db.SetEntryBlob(ctx, key, e.Type(), bID, bVal, opts...)
	if err != nil {
		return errors.Wrap(err, "SetBlob setting in db")
	}
	return nil
}

func (e *DBStatusObject) ClrBlob(ctx context.Context, bID string, opts ...db.BlobOpt) error {
	db, err := e.GetDB()
	if err != nil {
		return errors.Wrap(err, "ClrBlob")
	}
	key, err := e.Key()
	if err != nil {
		return errors.Wrap(err, "ClrBlob accessing key")
	}
	err = db.ClrEntryBlob(ctx, key, e.Type(), bID, opts...)
	if err != nil {
		return errors.Wrap(err, "ClrBlob clearing in db")
	}
	return nil
}

func (e *DBStatusObject) GetBlob(ctx context.Context, bID string, opts ...db.BlobOpt) (interface{}, error) {
	db, err := e.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlob")
	}
	key, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlob accessing key")
	}
	return db.GetEntryBlob(ctx, key, e.Type(), bID, opts...)
}

func (e *DBStatusObject) GetBlobs(ctx context.Context, opts ...db.BlobOpt) (map[string]interface{}, error) {
	db, err := e.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlobs")
	}
	key, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetBlobs accessing key")
	}
	return db.GetEntryBlobs(ctx, key, e.Type(), opts...)
}

func (e *DBStatusObject) IsDeleted() (bool, error) {
	db, err := e.GetDB()
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted")
	}
	key, err := e.Key()
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted accessing key")
	}
	isDel, err := db.IsEntryDeleted(key, e.Type())
	if err != nil {
		return false, errors.Wrap(err, "IsDeleted accessing db")
	}
	return isDel, nil
}

// Implement ves.io/stdlib/db.EntryPvt interface
func (e *DBStatusObject) SetTable(tbl db.Table) {
	e.tbl = tbl
}

func (e *DBStatusObject) GetDRefInfo() ([]db.DRefInfo, error) {
	if e == nil {
		return nil, nil
	}
	refrUID, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetDRefInfo, error in key")
	}

	drInfos, err := e.GetObjectRefsDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetObjectRefsDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		// Convert Spec.LcSpec.vnRefs to ves.io.examplesvc.objectone.Object.Spec.LcSpec.vnRefs
		dri.DRField = "ves.io.schema.service_policy_rule.StatusObject." + dri.DRField
		dri.RefrType = e.Type()
		dri.RefrUID = refrUID

		// convert any ref_to schema annotation specified by kind value to type value
		if !strings.HasPrefix(dri.RefdType, "ves.io") {
			d, err := e.GetDB()
			if err != nil {
				return nil, errors.Wrap(err, "Cannot find db for entry to resolve kind to type")
			}
			refdType, err := d.TypeForEntryKind(dri.RefrType, dri.RefrUID, dri.RefdType)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Cannot convert kind %s to type", dri.RefdType))
			}
			dri.RefdType = refdType
		}
	}
	return drInfos, nil

}

func (e *DBStatusObject) ToStore() store.Entry {
	return e.StatusObject
}

func (e *DBStatusObject) ToJSON() (string, error) {
	return e.ToStore().ToJSON()
}

func (e *DBStatusObject) ToYAML() (string, error) {
	return e.ToStore().ToYAML()
}

func (e *DBStatusObject) GetTable() db.Table {
	return e.tbl
}

func NewDBStatusObject(o *StatusObject, opts ...db.OpOption) *DBStatusObject {
	op := db.NewOpFrom(opts...)
	if o == nil {
		return &DBStatusObject{StatusObject: &StatusObject{}}
	}
	obj := o
	if !op.NoCopy() {
		obj = o.DeepCopy()
	}
	return &DBStatusObject{StatusObject: obj}
}

func NewEntryStatusObject(opts ...db.OpOption) db.Entry {
	op := db.NewOpFrom(opts...)
	s := op.StoreEntry()
	switch v := s.(type) {
	case nil:
		return NewDBStatusObject(nil, opts...)
	case *StatusObject:
		return NewDBStatusObject(v, opts...)
	}
	return nil
}

func (e *DBStatusObject) GetObjectRefsDRefInfo() ([]db.DRefInfo, error) {
	refrUID, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "GetDRefInfo, error in key")
	}
	refs := e.GetObjectRefs()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("StatusObject.object_refs[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "service_policy_rule.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			RefrType:   e.Type(),
			RefrUID:    refrUID,
			DRField:    "object_refs",
			Ref:        ref,
		})
	}
	return drInfos, nil
}

// GetObjectRefsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (e *DBStatusObject) GetObjectRefsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refrUID, err := e.Key()
	if err != nil {
		return nil, errors.Wrap(err, "Get<fld>DBEntries, error in key")
	}
	refdType, err := d.TypeForEntryKind(e.Type(), refrUID, "service_policy_rule.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: service_policy_rule")
	}
	tblName := db.DefaultTableName(refdType)
	if intTbl, err := d.GetTable(ctx, db.InternalTableName(refdType)); err == nil {
		tblName = intTbl.Name()
	}
	for _, ref := range e.GetObjectRefs() {
		e, exist, err := d.FindEntry(ctx, tblName, ref.Uid)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Tbl: %s, Key: %s", tblName, ref.Uid))
		}
		if !exist {
			continue
		}
		entries = append(entries, e)
	}
	return entries, nil
}

type ValidateStatusObject struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateStatusObject) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	e, ok := pm.(*StatusObject)
	if !ok {
		switch t := pm.(type) {
		default:
			return fmt.Errorf("Expected type *StatusObject got type %s", t)
		}
	}
	if e == nil {
		return nil
	}

	if fv, exists := v.FldValidators["conditions"]; exists {

		vOpts := append(opts, db.WithValidateField("conditions"))
		for idx, item := range e.GetConditions() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, e.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["object_refs"]; exists {

		vOpts := append(opts, db.WithValidateField("object_refs"))
		for idx, item := range e.GetObjectRefs() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultStatusObjectValidator = func() *ValidateStatusObject {
	v := &ValidateStatusObject{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["conditions"] = ves_io_schema.ConditionTypeValidator().Validate

	return v
}()

func StatusObjectValidator() db.Validator {
	return DefaultStatusObjectValidator
}
