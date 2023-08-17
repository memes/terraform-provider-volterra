// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package fast_acl_rule

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.fast_acl_rule.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.fast_acl_rule.Object"] = ObjectValidator()
	vr["ves.io.schema.fast_acl_rule.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.fast_acl_rule.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.fast_acl_rule.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.fast_acl_rule.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.fast_acl_rule.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.fast_acl_rule.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.fast_acl_rule.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.fast_acl_rule.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.fast_acl_rule.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.fast_acl_rule.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.fast_acl_rule.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.fast_acl_rule.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.fast_acl_rule.FastAclRuleAction"] = FastAclRuleActionValidator()
	vr["ves.io.schema.fast_acl_rule.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.fast_acl_rule.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.fast_acl_rule.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.fast_acl_rule.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.fast_acl_rule.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.fast_acl_rule.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.fast_acl_rule.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.fast_acl_rule.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.fast_acl_rule.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.fast_acl_rule.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.fast_acl_rule.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.fast_acl_rule.API.Get"] = []string{
		"object",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.fast_acl_rule.API"] = "config"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.fast_acl_rule.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.fast_acl_rule.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.fast_acl_rule.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.fast_acl_rule.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.fast_acl_rule.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.fast_acl_rule.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.fast_acl_rule.Object"] = NewCRUDAPIServer

	}()

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
