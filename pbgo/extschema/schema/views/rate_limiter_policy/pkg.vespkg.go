// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package rate_limiter_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.rate_limiter_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.rate_limiter_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.views.rate_limiter_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.rate_limiter_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.rate_limiter_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.rate_limiter_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.rate_limiter_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.rate_limiter_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.rate_limiter_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.rate_limiter_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.rate_limiter_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.rate_limiter_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.rate_limiter_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.rate_limiter_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.rate_limiter_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.rate_limiter_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.rate_limiter_policy.PolicyList"] = PolicyListValidator()
	vr["ves.io.schema.views.rate_limiter_policy.RateLimiterRule"] = RateLimiterRuleValidator()
	vr["ves.io.schema.views.rate_limiter_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.rate_limiter_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.rate_limiter_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.rate_limiter_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.rate_limiter_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.rate_limiter_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.rate_limiter_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.rate_limiter_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.rate_limiter_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.rate_limiter_policy.API.Create"] = []string{
		"spec.rules.#.metadata.disable",
		"spec.rules.#.spec.ip_prefix_list.ipv6_prefixes.#",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.rate_limiter_policy.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.rate_limiter_policy.API.Replace"] = []string{
		"spec.rules.#.metadata.disable",
		"spec.rules.#.spec.ip_prefix_list.ipv6_prefixes.#",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.rate_limiter_policy.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.rate_limiter_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.rate_limiter_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.rate_limiter_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.rate_limiter_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.rate_limiter_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.rate_limiter_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.rate_limiter_policy.Object"] = NewCRUDAPIServer

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
