// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package service_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.service_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.service_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.service_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.service_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.service_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.service_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.service_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.service_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.service_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.service_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.service_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.service_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.service_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.service_policy.ServicePolicyHits"] = ServicePolicyHitsValidator()
	vr["ves.io.schema.service_policy.ServicePolicyHitsId"] = ServicePolicyHitsIdValidator()
	vr["ves.io.schema.service_policy.ServicePolicyHitsRequest"] = ServicePolicyHitsRequestValidator()
	vr["ves.io.schema.service_policy.ServicePolicyHitsResponse"] = ServicePolicyHitsResponseValidator()
	vr["ves.io.schema.service_policy.ServicePolicyMetricLabelFilter"] = ServicePolicyMetricLabelFilterValidator()

	vr["ves.io.schema.service_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.service_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.service_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.service_policy.LegacyRuleList"] = LegacyRuleListValidator()
	vr["ves.io.schema.service_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.service_policy.Rule"] = RuleValidator()
	vr["ves.io.schema.service_policy.RuleList"] = RuleListValidator()
	vr["ves.io.schema.service_policy.SimpleRule"] = SimpleRuleValidator()
	vr["ves.io.schema.service_policy.SourceList"] = SourceListValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.service_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.service_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.service_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.service_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.service_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.service_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.service_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.service_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.service_policy.API.Create"] = []string{
		"spec.algo",
		"spec.allow_list.prefix_list.ipv6_prefixes.#",
		"spec.deny_list.prefix_list.ipv6_prefixes.#",
		"spec.internally_generated",
		"spec.legacy_rule_list",
		"spec.port_matcher",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.spec.additional_api_group_matchers.#",
		"spec.rule_list.rules.#.spec.any_dst_asn",
		"spec.rule_list.rules.#.spec.any_dst_ip",
		"spec.rule_list.rules.#.spec.challenge_action",
		"spec.rule_list.rules.#.spec.client_role",
		"spec.rule_list.rules.#.spec.content_rewrite_action",
		"spec.rule_list.rules.#.spec.dst_asn_list",
		"spec.rule_list.rules.#.spec.dst_asn_matcher",
		"spec.rule_list.rules.#.spec.dst_ip_matcher",
		"spec.rule_list.rules.#.spec.dst_ip_prefix_list",
		"spec.rule_list.rules.#.spec.forwarding_class.#",
		"spec.rule_list.rules.#.spec.goto_policy.#",
		"spec.rule_list.rules.#.spec.graphql_settings",
		"spec.rule_list.rules.#.spec.ip_prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.spec.ip_reputation_action",
		"spec.rule_list.rules.#.spec.l4_dest_matcher",
		"spec.rule_list.rules.#.spec.openapi_validation_action",
		"spec.rule_list.rules.#.spec.origin_server_subsets_action",
		"spec.rule_list.rules.#.spec.rate_limiter.#",
		"spec.rule_list.rules.#.spec.scheme.#",
		"spec.rule_list.rules.#.spec.server_selector",
		"spec.rule_list.rules.#.spec.shape_protected_endpoint_action",
		"spec.rule_list.rules.#.spec.url_matcher",
		"spec.rule_list.rules.#.spec.virtual_host_matcher",
		"spec.rule_list.rules.#.spec.waf_action.data_guard_control",
		"spec.rule_list.rules.#.spec.waf_action.waf_in_monitoring_mode",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.service_policy.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.service_policy.API.Replace"] = []string{
		"spec.algo",
		"spec.allow_list.prefix_list.ipv6_prefixes.#",
		"spec.deny_list.prefix_list.ipv6_prefixes.#",
		"spec.internally_generated",
		"spec.port_matcher",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.spec.additional_api_group_matchers.#",
		"spec.rule_list.rules.#.spec.any_dst_asn",
		"spec.rule_list.rules.#.spec.any_dst_ip",
		"spec.rule_list.rules.#.spec.challenge_action",
		"spec.rule_list.rules.#.spec.client_role",
		"spec.rule_list.rules.#.spec.content_rewrite_action",
		"spec.rule_list.rules.#.spec.dst_asn_list",
		"spec.rule_list.rules.#.spec.dst_asn_matcher",
		"spec.rule_list.rules.#.spec.dst_ip_matcher",
		"spec.rule_list.rules.#.spec.dst_ip_prefix_list",
		"spec.rule_list.rules.#.spec.forwarding_class.#",
		"spec.rule_list.rules.#.spec.goto_policy.#",
		"spec.rule_list.rules.#.spec.graphql_settings",
		"spec.rule_list.rules.#.spec.ip_prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.spec.ip_reputation_action",
		"spec.rule_list.rules.#.spec.l4_dest_matcher",
		"spec.rule_list.rules.#.spec.openapi_validation_action",
		"spec.rule_list.rules.#.spec.origin_server_subsets_action",
		"spec.rule_list.rules.#.spec.rate_limiter.#",
		"spec.rule_list.rules.#.spec.scheme.#",
		"spec.rule_list.rules.#.spec.server_selector",
		"spec.rule_list.rules.#.spec.shape_protected_endpoint_action",
		"spec.rule_list.rules.#.spec.url_matcher",
		"spec.rule_list.rules.#.spec.virtual_host_matcher",
		"spec.rule_list.rules.#.spec.waf_action.data_guard_control",
		"spec.rule_list.rules.#.spec.waf_action.waf_in_monitoring_mode",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.service_policy.API"] = "config"
	sm["ves.io.schema.service_policy.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.service_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.service_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.service_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.service_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.service_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.service_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.service_policy.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.service_policy.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.service_policy.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.service_policy.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.service_policy.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.service_policy.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.service_policy.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
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
