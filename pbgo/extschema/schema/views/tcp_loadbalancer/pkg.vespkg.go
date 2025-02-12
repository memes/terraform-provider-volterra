// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package tcp_loadbalancer

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.tcp_loadbalancer.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.tcp_loadbalancer.Object"] = ObjectValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.tcp_loadbalancer.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.tcp_loadbalancer.GetDnsInfoRequest"] = GetDnsInfoRequestValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.GetDnsInfoResponse"] = GetDnsInfoResponseValidator()

	vr["ves.io.schema.views.tcp_loadbalancer.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ProxyTypeTLSTCP"] = ProxyTypeTLSTCPValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ProxyTypeTLSTCPAutoCerts"] = ProxyTypeTLSTCPAutoCertsValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.views.tcp_loadbalancer.ServicePolicyList"] = ServicePolicyListValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.tcp_loadbalancer.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.tcp_loadbalancer.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.tcp_loadbalancer.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.tcp_loadbalancer.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.tcp_loadbalancer.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.tcp_loadbalancer.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.tcp_loadbalancer.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.tcp_loadbalancer.API.Create"] = []string{
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.secret_encoding_type",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.vault_secret_info",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.tcp_loadbalancer.API.Create"] = "ves.io.schema.views.tcp_loadbalancer.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.tcp_loadbalancer.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.tcp_loadbalancer.API.Replace"] = []string{
		"spec.origin_pools.#",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.secret_encoding_type",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.vault_secret_info",
		"spec.tls_tcp.tls_parameters.tls_certificates.#.private_key.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.tcp_loadbalancer.API.Replace"] = "ves.io.schema.views.tcp_loadbalancer.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.tcp_loadbalancer.API"] = "config"
	sm["ves.io.schema.views.tcp_loadbalancer.CustomAPI"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.tcp_loadbalancer.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.tcp_loadbalancer.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.tcp_loadbalancer.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.tcp_loadbalancer.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.tcp_loadbalancer.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.tcp_loadbalancer.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.tcp_loadbalancer.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.tcp_loadbalancer.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
