// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cminstance

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.cminstance.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.cminstance.Object"] = ObjectValidator()
	vr["ves.io.schema.cminstance.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.cminstance.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.cminstance.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.cminstance.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.cminstance.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.cminstance.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.cminstance.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.cminstance.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.cminstance.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.cminstance.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.cminstance.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.cminstance.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.cminstance.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.cminstance.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.cminstance.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.cminstance.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.cminstance.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cminstance.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.cminstance.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.cminstance.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.cminstance.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cminstance.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.cminstance.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.cminstance.API.Create"] = []string{
		"spec.api_token.blindfold_secret_info_internal",
		"spec.api_token.secret_encoding_type",
		"spec.api_token.vault_secret_info",
		"spec.api_token.wingman_secret_info",
		"spec.password.blindfold_secret_info_internal",
		"spec.password.secret_encoding_type",
		"spec.password.vault_secret_info",
		"spec.password.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.cminstance.API.Create"] = "ves.io.schema.cminstance.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.cminstance.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.cminstance.API.Replace"] = []string{
		"spec.api_token.blindfold_secret_info_internal",
		"spec.api_token.secret_encoding_type",
		"spec.api_token.vault_secret_info",
		"spec.api_token.wingman_secret_info",
		"spec.password.blindfold_secret_info_internal",
		"spec.password.secret_encoding_type",
		"spec.password.vault_secret_info",
		"spec.password.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.cminstance.API.Replace"] = "ves.io.schema.cminstance.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.cminstance.API"] = "config"

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

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.cminstance.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.cminstance.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.cminstance.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.cminstance.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.cminstance.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.cminstance.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.cminstance.Object"] = NewCRUDAPIServer

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
