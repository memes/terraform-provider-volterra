// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package network_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.network_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.network_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.network_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.network_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.network_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.network_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.network_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.network_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.network_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.network_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.network_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.network_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.network_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.network_policy.NetworkPolicyHits"] = NetworkPolicyHitsValidator()
	vr["ves.io.schema.network_policy.NetworkPolicyHitsId"] = NetworkPolicyHitsIdValidator()
	vr["ves.io.schema.network_policy.NetworkPolicyHitsRequest"] = NetworkPolicyHitsRequestValidator()
	vr["ves.io.schema.network_policy.NetworkPolicyHitsResponse"] = NetworkPolicyHitsResponseValidator()
	vr["ves.io.schema.network_policy.NetworkPolicyMetricLabelFilter"] = NetworkPolicyMetricLabelFilterValidator()

	vr["ves.io.schema.network_policy.ApplicationsType"] = ApplicationsTypeValidator()
	vr["ves.io.schema.network_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.network_policy.EndpointChoiceType"] = EndpointChoiceTypeValidator()
	vr["ves.io.schema.network_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.network_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.network_policy.LegacyNetworkPolicyRuleChoice"] = LegacyNetworkPolicyRuleChoiceValidator()
	vr["ves.io.schema.network_policy.NetworkPolicyRuleChoice"] = NetworkPolicyRuleChoiceValidator()
	vr["ves.io.schema.network_policy.NetworkPolicyRuleType"] = NetworkPolicyRuleTypeValidator()
	vr["ves.io.schema.network_policy.ProtocolPortType"] = ProtocolPortTypeValidator()
	vr["ves.io.schema.network_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.network_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.network_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.network_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.network_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.network_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.network_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.network_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.network_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.network_policy.API.Create"] = []string{
		"spec.endpoint.prefix_list.ipv6_prefixes.#",
		"spec.rules.egress_rules.#.keys.#",
		"spec.rules.egress_rules.#.metadata.disable",
		"spec.rules.egress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.rules.egress_rules.#.rule_description",
		"spec.rules.egress_rules.#.rule_name",
		"spec.rules.ingress_rules.#.keys.#",
		"spec.rules.ingress_rules.#.metadata.disable",
		"spec.rules.ingress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.rules.ingress_rules.#.rule_description",
		"spec.rules.ingress_rules.#.rule_name",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.network_policy.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.network_policy.API.Replace"] = []string{
		"spec.endpoint.prefix_list.ipv6_prefixes.#",
		"spec.rules.egress_rules.#.keys.#",
		"spec.rules.egress_rules.#.metadata.disable",
		"spec.rules.egress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.rules.egress_rules.#.rule_description",
		"spec.rules.egress_rules.#.rule_name",
		"spec.rules.ingress_rules.#.keys.#",
		"spec.rules.ingress_rules.#.metadata.disable",
		"spec.rules.ingress_rules.#.prefix_list.ipv6_prefixes.#",
		"spec.rules.ingress_rules.#.rule_description",
		"spec.rules.ingress_rules.#.rule_name",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.network_policy.API"] = "config"
	sm["ves.io.schema.network_policy.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.network_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.network_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.network_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.network_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.network_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.network_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.network_policy.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.network_policy.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.network_policy.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.network_policy.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.network_policy.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.network_policy.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.network_policy.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
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
