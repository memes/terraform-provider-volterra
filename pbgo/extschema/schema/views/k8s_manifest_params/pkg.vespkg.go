// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package k8s_manifest_params

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.k8s_manifest_params.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.k8s_manifest_params.Object"] = ObjectValidator()
	vr["ves.io.schema.views.k8s_manifest_params.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.k8s_manifest_params.BigIpBareMetalDeviceK8sParams"] = BigIpBareMetalDeviceK8SParamsValidator()
	vr["ves.io.schema.views.k8s_manifest_params.BigIpBareMetalK8sType"] = BigIpBareMetalK8STypeValidator()
	vr["ves.io.schema.views.k8s_manifest_params.DeploymentStatusType"] = DeploymentStatusTypeValidator()
	vr["ves.io.schema.views.k8s_manifest_params.GlobalSpecType"] = GlobalSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.k8s_manifest_params.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.k8s_manifest_params.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.k8s_manifest_params.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.k8s_manifest_params.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.k8s_manifest_params.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.k8s_manifest_params.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.k8s_manifest_params.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.k8s_manifest_params.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

}

func InitializeMDRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	initializeEntryRegistry(mdr)
	initializeValidatorRegistry(mdr.ValidatorRegistry)

	initializeCRUDServiceRegistry(mdr, isExternal)
	if isExternal {
		return
	}

	initializeRPCRegistry(mdr)
	initializeAPIGwServiceSlugsRegistry(mdr.APIGwServiceSlugs)
	initializeP0PolicyRegistry(mdr.P0PolicyRegistry)

}
