// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package role

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.role.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.role.Object"] = ObjectValidator()
	vr["ves.io.schema.role.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.role.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.role.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.role.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.role.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.role.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.role.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.role.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.role.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.role.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.role.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.role.CustomCreateRequest"] = CustomCreateRequestValidator()
	vr["ves.io.schema.role.CustomGetRequest"] = CustomGetRequestValidator()
	vr["ves.io.schema.role.CustomGetResponse"] = CustomGetResponseValidator()
	vr["ves.io.schema.role.CustomListRequest"] = CustomListRequestValidator()
	vr["ves.io.schema.role.CustomListResponse"] = CustomListResponseValidator()
	vr["ves.io.schema.role.CustomReplaceRequest"] = CustomReplaceRequestValidator()
	vr["ves.io.schema.role.Role"] = RoleValidator()

	vr["ves.io.schema.role.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.role.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.role.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.role.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.role.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.role.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.role.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.role.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.role.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.role.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.role.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.role.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.role.API.Get"] = []string{
		"object",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.role.API"] = "web"
	sm["ves.io.schema.role.CustomAPI"] = "web"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.role.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.role.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.role.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.role.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.role.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.role.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.role.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.role.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.role.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.role.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.role.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.role.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.role.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomAPIServer(svc)
		}

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
