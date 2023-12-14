// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package http_loadbalancer

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.http_loadbalancer.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.http_loadbalancer.Object"] = ObjectValidator()
	vr["ves.io.schema.views.http_loadbalancer.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.http_loadbalancer.APIGroupsApiep"] = APIGroupsApiepValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetAPIEndpointsForGroupsReq"] = GetAPIEndpointsForGroupsReqValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetAPIEndpointsForGroupsRsp"] = GetAPIEndpointsForGroupsRspValidator()
	vr["ves.io.schema.views.http_loadbalancer.SwaggerSpecReq"] = SwaggerSpecReqValidator()
	vr["ves.io.schema.views.http_loadbalancer.SwaggerSpecRsp"] = SwaggerSpecRspValidator()

	vr["ves.io.schema.views.http_loadbalancer.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.http_loadbalancer.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.http_loadbalancer.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.http_loadbalancer.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.http_loadbalancer.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.http_loadbalancer.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.http_loadbalancer.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.http_loadbalancer.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.http_loadbalancer.DeleteDoSAutoMitigationRuleReq"] = DeleteDoSAutoMitigationRuleReqValidator()
	vr["ves.io.schema.views.http_loadbalancer.DeleteDoSAutoMitigationRuleRsp"] = DeleteDoSAutoMitigationRuleRspValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetDnsInfoRequest"] = GetDnsInfoRequestValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetDnsInfoResponse"] = GetDnsInfoResponseValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetDoSAutoMitigationRulesReq"] = GetDoSAutoMitigationRulesReqValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetDoSAutoMitigationRulesRsp"] = GetDoSAutoMitigationRulesRspValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetSecurityConfigReq"] = GetSecurityConfigReqValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetSecurityConfigRsp"] = GetSecurityConfigRspValidator()
	vr["ves.io.schema.views.http_loadbalancer.HTTPLoadBalancerList"] = HTTPLoadBalancerListValidator()

	vr["ves.io.schema.views.http_loadbalancer.APIEndpointProtectionRule"] = APIEndpointProtectionRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.APIGroupProtectionRule"] = APIGroupProtectionRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.APIGroups"] = APIGroupsValidator()
	vr["ves.io.schema.views.http_loadbalancer.APIProtectionRuleAction"] = APIProtectionRuleActionValidator()
	vr["ves.io.schema.views.http_loadbalancer.APIProtectionRules"] = APIProtectionRulesValidator()
	vr["ves.io.schema.views.http_loadbalancer.APIRateLimit"] = APIRateLimitValidator()
	vr["ves.io.schema.views.http_loadbalancer.APISpecificationSettings"] = APISpecificationSettingsValidator()
	vr["ves.io.schema.views.http_loadbalancer.AdvancedOptionsType"] = AdvancedOptionsTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ApiDefinitionList"] = ApiDefinitionListValidator()
	vr["ves.io.schema.views.http_loadbalancer.ApiDiscoverySetting"] = ApiDiscoverySettingValidator()
	vr["ves.io.schema.views.http_loadbalancer.ApiEndpointDetails"] = ApiEndpointDetailsValidator()
	vr["ves.io.schema.views.http_loadbalancer.ApiEndpointRule"] = ApiEndpointRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.AppEndpointType"] = AppEndpointTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.Audiences"] = AudiencesValidator()
	vr["ves.io.schema.views.http_loadbalancer.AutoMitigationAction"] = AutoMitigationActionValidator()
	vr["ves.io.schema.views.http_loadbalancer.BasePathsType"] = BasePathsTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.BotAdvancedMobileSDKConfigType"] = BotAdvancedMobileSDKConfigTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.BotDefenseAdvancedPolicyType"] = BotDefenseAdvancedPolicyTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.BotDefenseAdvancedType"] = BotDefenseAdvancedTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.CSDJavaScriptInsertAllWithExceptionsType"] = CSDJavaScriptInsertAllWithExceptionsTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.CSDJavaScriptInsertType"] = CSDJavaScriptInsertTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.CSDJavaScriptInsertionRule"] = CSDJavaScriptInsertionRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.ChallengeRule"] = ChallengeRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.ChallengeRuleList"] = ChallengeRuleListValidator()
	vr["ves.io.schema.views.http_loadbalancer.ClientSideDefensePolicyType"] = ClientSideDefensePolicyTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ClientSideDefenseType"] = ClientSideDefenseTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.CustomFallThroughMode"] = CustomFallThroughModeValidator()
	vr["ves.io.schema.views.http_loadbalancer.CustomIpAllowedList"] = CustomIpAllowedListValidator()
	vr["ves.io.schema.views.http_loadbalancer.DDoSClientSource"] = DDoSClientSourceValidator()
	vr["ves.io.schema.views.http_loadbalancer.DDoSMitigationRule"] = DDoSMitigationRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.EnableChallenge"] = EnableChallengeValidator()
	vr["ves.io.schema.views.http_loadbalancer.EnableDDoSDetectionSetting"] = EnableDDoSDetectionSettingValidator()
	vr["ves.io.schema.views.http_loadbalancer.FallThroughRule"] = FallThroughRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.HashPolicyListType"] = HashPolicyListTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.HttpHeaderMatcherList"] = HttpHeaderMatcherListValidator()
	vr["ves.io.schema.views.http_loadbalancer.IPThreatCategoryListType"] = IPThreatCategoryListTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.InlineRateLimiter"] = InlineRateLimiterValidator()
	vr["ves.io.schema.views.http_loadbalancer.JWKS"] = JWKSValidator()
	vr["ves.io.schema.views.http_loadbalancer.JWTValidation"] = JWTValidationValidator()
	vr["ves.io.schema.views.http_loadbalancer.MirrorPolicyType"] = MirrorPolicyTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.MobileSDKConfigType"] = MobileSDKConfigTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.MobileTrafficIdentifierType"] = MobileTrafficIdentifierTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiFallThroughMode"] = OpenApiFallThroughModeValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiValidationAllSpecEndpointsSettings"] = OpenApiValidationAllSpecEndpointsSettingsValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiValidationCommonSettings"] = OpenApiValidationCommonSettingsValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiValidationMode"] = OpenApiValidationModeValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiValidationModeActive"] = OpenApiValidationModeActiveValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiValidationModeActiveResponse"] = OpenApiValidationModeActiveResponseValidator()
	vr["ves.io.schema.views.http_loadbalancer.OpenApiValidationRule"] = OpenApiValidationRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.OriginServerSubsetRuleListType"] = OriginServerSubsetRuleListTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.PolicyBasedChallenge"] = PolicyBasedChallengeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ProtectedAppEndpointType"] = ProtectedAppEndpointTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ProxyTypeHttp"] = ProxyTypeHttpValidator()
	vr["ves.io.schema.views.http_loadbalancer.ProxyTypeHttps"] = ProxyTypeHttpsValidator()
	vr["ves.io.schema.views.http_loadbalancer.ProxyTypeHttpsAutoCerts"] = ProxyTypeHttpsAutoCertsValidator()
	vr["ves.io.schema.views.http_loadbalancer.Query"] = QueryValidator()
	vr["ves.io.schema.views.http_loadbalancer.RateLimitConfigType"] = RateLimitConfigTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.RequestBody"] = RequestBodyValidator()
	vr["ves.io.schema.views.http_loadbalancer.ReservedClaims"] = ReservedClaimsValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteSimpleAdvancedOptions"] = RouteSimpleAdvancedOptionsValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteType"] = RouteTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteTypeCustomRoute"] = RouteTypeCustomRouteValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteTypeDirectResponse"] = RouteTypeDirectResponseValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteTypeRedirect"] = RouteTypeRedirectValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteTypeSimple"] = RouteTypeSimpleValidator()
	vr["ves.io.schema.views.http_loadbalancer.RouteTypeSimpleWithDefaultOriginPool"] = RouteTypeSimpleWithDefaultOriginPoolValidator()
	vr["ves.io.schema.views.http_loadbalancer.ServerUrlRule"] = ServerUrlRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.ServicePolicyList"] = ServicePolicyListValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeBotDefensePolicyType"] = ShapeBotDefensePolicyTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeBotDefenseType"] = ShapeBotDefenseTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeJavaScriptExclusionRule"] = ShapeJavaScriptExclusionRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeJavaScriptInsertAllType"] = ShapeJavaScriptInsertAllTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeJavaScriptInsertAllWithExceptionsType"] = ShapeJavaScriptInsertAllWithExceptionsTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeJavaScriptInsertType"] = ShapeJavaScriptInsertTypeValidator()
	vr["ves.io.schema.views.http_loadbalancer.ShapeJavaScriptInsertionRule"] = ShapeJavaScriptInsertionRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.SimpleClientSrcRule"] = SimpleClientSrcRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.SingleLoadBalancerAppSetting"] = SingleLoadBalancerAppSettingValidator()
	vr["ves.io.schema.views.http_loadbalancer.Target"] = TargetValidator()
	vr["ves.io.schema.views.http_loadbalancer.TokenLocation"] = TokenLocationValidator()
	vr["ves.io.schema.views.http_loadbalancer.ValidateApiBySpecRule"] = ValidateApiBySpecRuleValidator()
	vr["ves.io.schema.views.http_loadbalancer.ValidationPropertySetting"] = ValidationPropertySettingValidator()
	vr["ves.io.schema.views.http_loadbalancer.ValidationSettingForHeaders"] = ValidationSettingForHeadersValidator()
	vr["ves.io.schema.views.http_loadbalancer.ValidationSettingForQueryParameters"] = ValidationSettingForQueryParametersValidator()
	vr["ves.io.schema.views.http_loadbalancer.WebMobileTrafficType"] = WebMobileTrafficTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.http_loadbalancer.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.http_loadbalancer.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.http_loadbalancer.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.http_loadbalancer.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.http_loadbalancer.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.http_loadbalancer.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.http_loadbalancer.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.http_loadbalancer.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.views.http_loadbalancer.CreateRequest.spec.bot_defense_choice.bot_defense",
			AddonServices: []string{"shape-bot,f5xc-bot-defense-advanced"},
		},
		{
			FieldPath:     "ves.io.schema.views.http_loadbalancer.CreateRequest.spec.bot_defense_choice.bot_defense_advanced",
			AddonServices: []string{"shape-bot,f5xc-bot-defense-advanced"},
		},
		{
			FieldPath:     "ves.io.schema.views.http_loadbalancer.CreateRequest.spec.client_side_defense_choice.client_side_defense",
			AddonServices: []string{"client-side-defense"},
		},
	}

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = []string{
		"spec.jwt_validation.jwks",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = []string{
		"spec.jwt_validation.jwks",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = []string{
		"spec.api_definition",
		"spec.api_definitions",
		"spec.api_protection_rules.api_endpoint_rules.#.metadata.disable",
		"spec.api_protection_rules.api_groups_rules.#.metadata.disable",
		"spec.api_rate_limit.api_endpoint_rules.#.base_path",
		"spec.api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.#.metadata.disable",
		"spec.api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.#.metadata.disable",
		"spec.api_specification.validation_custom_list.open_api_validation_rules.#.metadata.disable",
		"spec.blocked_clients.#.metadata.disable",
		"spec.bot_defense",
		"spec.bot_defense_advanced",
		"spec.client_side_defense.policy.js_insert_all_pages_except.exclude_list.#.metadata.disable",
		"spec.client_side_defense.policy.js_insertion_rules.exclude_list.#.metadata.disable",
		"spec.client_side_defense.policy.js_insertion_rules.rules.#.metadata.disable",
		"spec.cors_policy.max_age",
		"spec.data_guard_rules.#.metadata.disable",
		"spec.ddos_mitigation_rules.#.metadata.disable",
		"spec.default_pool.origin_servers.#.k8s_service.service_selector",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.secret_encoding_type",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.vault_secret_info",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.wingman_secret_info",
		"spec.disable_bot_defense",
		"spec.enable_api_discovery.sensitive_data_detection_rules.custom_sensitive_data_detection_rules.#.metadata.disable",
		"spec.graphql_rules.#.graphql_settings.max_value_length",
		"spec.graphql_rules.#.graphql_settings.policy_name",
		"spec.graphql_rules.#.metadata.disable",
		"spec.https.http_protocol_options",
		"spec.https.tls_parameters.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https.tls_parameters.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https.tls_parameters.tls_certificates.#.private_key.vault_secret_info",
		"spec.https.tls_parameters.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_auto_cert.http_protocol_options",
		"spec.jwt_validation.auth_server_uri",
		"spec.jwt_validation.jwks",
		"spec.jwt_validation.token_location.cookie",
		"spec.jwt_validation.token_location.header",
		"spec.jwt_validation.token_location.query_param",
		"spec.malicious_user_mitigation",
		"spec.more_option.buffer_policy.max_request_time",
		"spec.more_option.cookies_to_modify.#",
		"spec.more_option.javascript_info",
		"spec.more_option.jwt.#",
		"spec.more_option.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.more_option.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.more_option.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.more_option.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.more_option.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.more_option.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.more_option.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.more_option.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.origin_server_subset_rule_list.origin_server_subset_rules.#.body_matcher",
		"spec.origin_server_subset_rule_list.origin_server_subset_rules.#.metadata.disable",
		"spec.policy_based_challenge.rule_list.rules.#.metadata.disable",
		"spec.policy_based_challenge.rule_list.rules.#.spec.client_name",
		"spec.policy_based_challenge.rule_list.rules.#.spec.client_name_matcher",
		"spec.routes.#.redirect_route.route_redirect.all_params",
		"spec.routes.#.redirect_route.route_redirect.port_redirect",
		"spec.routes.#.redirect_route.route_redirect.strip_query_params",
		"spec.routes.#.simple_route.advanced_options.buffer_policy.max_request_time",
		"spec.routes.#.simple_route.advanced_options.cors_policy.max_age",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.simple_route.advanced_options.retry_policy.retry_on",
		"spec.routes.#.simple_route.advanced_options.web_socket_config.idle_timeout",
		"spec.routes.#.simple_route.advanced_options.web_socket_config.max_connect_attempts",
		"spec.single_lb_app.enable_discovery.sensitive_data_detection_rules.custom_sensitive_data_detection_rules.#.metadata.disable",
		"spec.trusted_clients.#.metadata.disable",
		"spec.waf_exclusion_rules.#.metadata.disable",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.http_loadbalancer.API.Create"] = "ves.io.schema.views.http_loadbalancer.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Get"] = []string{
		"create_form.spec.jwt_validation.jwks",
		"object",
		"replace_form.spec.jwt_validation.jwks",
		"spec.jwt_validation.jwks",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.http_loadbalancer.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "create_form.spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "object.spec.gc_spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "object.spec.gc_spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "replace_form.spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "replace_form.spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.List"] = []string{
		"items.#.get_spec.jwt_validation.jwks",
		"items.#.object.spec.gc_spec.jwt_validation.jwks",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.http_loadbalancer.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "items.#.get_spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "items.#.object.spec.gc_spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "items.#.object.spec.gc_spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.views.http_loadbalancer.ReplaceRequest.spec.bot_defense_choice.bot_defense",
			AddonServices: []string{"shape-bot,f5xc-bot-defense-advanced"},
		},
		{
			FieldPath:     "ves.io.schema.views.http_loadbalancer.ReplaceRequest.spec.bot_defense_choice.bot_defense_advanced",
			AddonServices: []string{"shape-bot,f5xc-bot-defense-advanced"},
		},
		{
			FieldPath:     "ves.io.schema.views.http_loadbalancer.ReplaceRequest.spec.client_side_defense_choice.client_side_defense",
			AddonServices: []string{"client-side-defense"},
		},
	}

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Replace"] = []string{
		"spec.jwt_validation.jwks",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.http_loadbalancer.API.Replace"] = []string{
		"spec.api_definition",
		"spec.api_definitions",
		"spec.api_protection_rules.api_endpoint_rules.#.metadata.disable",
		"spec.api_protection_rules.api_groups_rules.#.metadata.disable",
		"spec.api_rate_limit.api_endpoint_rules.#.base_path",
		"spec.api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.#.metadata.disable",
		"spec.api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.#.metadata.disable",
		"spec.api_specification.validation_custom_list.open_api_validation_rules.#.metadata.disable",
		"spec.blocked_clients.#.metadata.disable",
		"spec.bot_defense",
		"spec.bot_defense_advanced",
		"spec.client_side_defense.policy.js_insert_all_pages_except.exclude_list.#.metadata.disable",
		"spec.client_side_defense.policy.js_insertion_rules.exclude_list.#.metadata.disable",
		"spec.client_side_defense.policy.js_insertion_rules.rules.#.metadata.disable",
		"spec.cors_policy.max_age",
		"spec.data_guard_rules.#.metadata.disable",
		"spec.ddos_mitigation_rules.#.metadata.disable",
		"spec.default_pool.origin_servers.#.k8s_service.service_selector",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.secret_encoding_type",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.vault_secret_info",
		"spec.default_pool.use_tls.use_mtls.tls_certificates.#.private_key.wingman_secret_info",
		"spec.disable_bot_defense",
		"spec.enable_api_discovery.sensitive_data_detection_rules.custom_sensitive_data_detection_rules.#.metadata.disable",
		"spec.graphql_rules.#.graphql_settings.max_value_length",
		"spec.graphql_rules.#.graphql_settings.policy_name",
		"spec.graphql_rules.#.metadata.disable",
		"spec.https.http_protocol_options",
		"spec.https.tls_parameters.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https.tls_parameters.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https.tls_parameters.tls_certificates.#.private_key.vault_secret_info",
		"spec.https.tls_parameters.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_auto_cert.http_protocol_options",
		"spec.jwt_validation.auth_server_uri",
		"spec.jwt_validation.jwks",
		"spec.jwt_validation.token_location.cookie",
		"spec.jwt_validation.token_location.header",
		"spec.jwt_validation.token_location.query_param",
		"spec.malicious_user_mitigation",
		"spec.more_option.buffer_policy.max_request_time",
		"spec.more_option.cookies_to_modify.#",
		"spec.more_option.javascript_info",
		"spec.more_option.jwt.#",
		"spec.more_option.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.more_option.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.more_option.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.more_option.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.more_option.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.more_option.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.more_option.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.more_option.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.origin_server_subset_rule_list.origin_server_subset_rules.#.body_matcher",
		"spec.origin_server_subset_rule_list.origin_server_subset_rules.#.metadata.disable",
		"spec.policy_based_challenge.rule_list.rules.#.metadata.disable",
		"spec.policy_based_challenge.rule_list.rules.#.spec.client_name",
		"spec.policy_based_challenge.rule_list.rules.#.spec.client_name_matcher",
		"spec.routes.#.redirect_route.route_redirect.all_params",
		"spec.routes.#.redirect_route.route_redirect.port_redirect",
		"spec.routes.#.redirect_route.route_redirect.strip_query_params",
		"spec.routes.#.simple_route.advanced_options.buffer_policy.max_request_time",
		"spec.routes.#.simple_route.advanced_options.cors_policy.max_age",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.simple_route.advanced_options.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.simple_route.advanced_options.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.simple_route.advanced_options.retry_policy.retry_on",
		"spec.routes.#.simple_route.advanced_options.web_socket_config.idle_timeout",
		"spec.routes.#.simple_route.advanced_options.web_socket_config.max_connect_attempts",
		"spec.single_lb_app.enable_discovery.sensitive_data_detection_rules.custom_sensitive_data_detection_rules.#.metadata.disable",
		"spec.trusted_clients.#.metadata.disable",
		"spec.waf_exclusion_rules.#.metadata.disable",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.http_loadbalancer.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
		{
			FieldPath:           "spec.single_lb_app.enable_ddos_detection.enable_auto_mitigation.js_challenge",
			AllowedEnvironments: []string{"demo1", "devtest"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.http_loadbalancer.API.Replace"] = "ves.io.schema.views.http_loadbalancer.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI"] = "ml/data"
	sm["ves.io.schema.views.http_loadbalancer.API"] = "config"
	sm["ves.io.schema.views.http_loadbalancer.CustomAPI"] = "config"

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

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.http_loadbalancer.Object"] = ApiepLBCustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI"] = NewApiepLBCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI"] = NewApiepLBCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI"] = RegisterApiepLBCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI"] = RegisterGwApiepLBCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewApiepLBCustomAPIServer(svc)
		}

	}()

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.views.http_loadbalancer.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.http_loadbalancer.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.http_loadbalancer.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.http_loadbalancer.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.http_loadbalancer.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.http_loadbalancer.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.http_loadbalancer.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.http_loadbalancer.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.http_loadbalancer.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.http_loadbalancer.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.http_loadbalancer.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.http_loadbalancer.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.http_loadbalancer.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
