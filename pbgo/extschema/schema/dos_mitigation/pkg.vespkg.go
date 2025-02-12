// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package dos_mitigation

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.dos_mitigation.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.dos_mitigation.Object"] = ObjectValidator()
	vr["ves.io.schema.dos_mitigation.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.dos_mitigation.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.dos_mitigation.Destination"] = DestinationValidator()
	vr["ves.io.schema.dos_mitigation.DoSMitigationRuleInfo"] = DoSMitigationRuleInfoValidator()
	vr["ves.io.schema.dos_mitigation.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.dos_mitigation.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.dos_mitigation.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.dos_mitigation.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.dos_mitigation.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dos_mitigation.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.dos_mitigation.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.dos_mitigation.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.dos_mitigation.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dos_mitigation.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.dos_mitigation.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

	sm["config"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-config",
		ServiceSelector: "akar\\.gc.*\\",
	}

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
