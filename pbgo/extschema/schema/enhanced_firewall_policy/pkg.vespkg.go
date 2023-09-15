// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package enhanced_firewall_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.enhanced_firewall_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.enhanced_firewall_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.enhanced_firewall_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.enhanced_firewall_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.enhanced_firewall_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.enhanced_firewall_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.enhanced_firewall_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.enhanced_firewall_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyHits"] = EnhancedFirewallPolicyHitsValidator()
	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyHitsId"] = EnhancedFirewallPolicyHitsIdValidator()
	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyHitsRequest"] = EnhancedFirewallPolicyHitsRequestValidator()
	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyHitsResponse"] = EnhancedFirewallPolicyHitsResponseValidator()
	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyMetricLabelFilter"] = EnhancedFirewallPolicyMetricLabelFilterValidator()

	vr["ves.io.schema.enhanced_firewall_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyRuleListType"] = EnhancedFirewallPolicyRuleListTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.EnhancedFirewallPolicyRuleType"] = EnhancedFirewallPolicyRuleTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.OperGlobalSpecType"] = OperGlobalSpecTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.enhanced_firewall_policy.ServiceActionType"] = ServiceActionTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.enhanced_firewall_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.enhanced_firewall_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.enhanced_firewall_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.enhanced_firewall_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.enhanced_firewall_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.enhanced_firewall_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.enhanced_firewall_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.enhanced_firewall_policy.API.Create"] = []string{
		"spec.rule_list.rules.#.destination_namespace",
		"spec.rule_list.rules.#.destination_prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.source_namespace",
		"spec.rule_list.rules.#.source_prefix_list.ipv6_prefixes.#",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.enhanced_firewall_policy.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.enhanced_firewall_policy.API.Replace"] = []string{
		"spec.rule_list.rules.#.destination_namespace",
		"spec.rule_list.rules.#.destination_prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.source_namespace",
		"spec.rule_list.rules.#.source_prefix_list.ipv6_prefixes.#",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.enhanced_firewall_policy.API"] = "config"
	sm["ves.io.schema.enhanced_firewall_policy.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.enhanced_firewall_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.enhanced_firewall_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.enhanced_firewall_policy.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.enhanced_firewall_policy.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.enhanced_firewall_policy.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.enhanced_firewall_policy.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.enhanced_firewall_policy.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.enhanced_firewall_policy.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
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
