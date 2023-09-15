// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package network_policy_view

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.network_policy_view.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.network_policy_view.Object"] = ObjectValidator()
	vr["ves.io.schema.views.network_policy_view.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.network_policy_view.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.network_policy_view.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.network_policy_view.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.network_policy_view.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.network_policy_view.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.network_policy_view.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.network_policy_view.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.network_policy_view.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.network_policy_view.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.network_policy_view.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.network_policy_view.NetworkPolicyHits"] = NetworkPolicyHitsValidator()
	vr["ves.io.schema.views.network_policy_view.NetworkPolicyHitsId"] = NetworkPolicyHitsIdValidator()
	vr["ves.io.schema.views.network_policy_view.NetworkPolicyHitsRequest"] = NetworkPolicyHitsRequestValidator()
	vr["ves.io.schema.views.network_policy_view.NetworkPolicyHitsResponse"] = NetworkPolicyHitsResponseValidator()
	vr["ves.io.schema.views.network_policy_view.NetworkPolicyMetricLabelFilter"] = NetworkPolicyMetricLabelFilterValidator()

	vr["ves.io.schema.views.network_policy_view.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.network_policy_view.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.network_policy_view.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.network_policy_view.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.network_policy_view.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.network_policy_view.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.network_policy_view.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.network_policy_view.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.network_policy_view.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.network_policy_view.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.network_policy_view.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.network_policy_view.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.network_policy_view.API.Create"] = []string{
		"spec.egress_rules.#.keys.#",
		"spec.egress_rules.#.metadata.disable",
		"spec.egress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.egress_rules.#.rule_description",
		"spec.egress_rules.#.rule_name",
		"spec.endpoint.prefix_list.ipv6_prefixes.#",
		"spec.ingress_rules.#.keys.#",
		"spec.ingress_rules.#.metadata.disable",
		"spec.ingress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.ingress_rules.#.rule_description",
		"spec.ingress_rules.#.rule_name",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.network_policy_view.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.network_policy_view.API.Replace"] = []string{
		"spec.egress_rules.#.keys.#",
		"spec.egress_rules.#.metadata.disable",
		"spec.egress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.egress_rules.#.rule_description",
		"spec.egress_rules.#.rule_name",
		"spec.endpoint.prefix_list.ipv6_prefixes.#",
		"spec.ingress_rules.#.keys.#",
		"spec.ingress_rules.#.metadata.disable",
		"spec.ingress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.ingress_rules.#.rule_description",
		"spec.ingress_rules.#.rule_name",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.network_policy_view.API"] = "config"
	sm["ves.io.schema.views.network_policy_view.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.network_policy_view.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.network_policy_view.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.network_policy_view.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.network_policy_view.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.network_policy_view.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.network_policy_view.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.network_policy_view.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.network_policy_view.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.network_policy_view.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.network_policy_view.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.network_policy_view.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.network_policy_view.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.network_policy_view.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
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
