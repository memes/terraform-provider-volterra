// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package app_type

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.app_type.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.app_type.Object"] = ObjectValidator()
	vr["ves.io.schema.app_type.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.app_type.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.app_type.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.app_type.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.app_type.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.app_type.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.app_type.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.app_type.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.app_type.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.app_type.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.app_type.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.app_type.APIEndpointLearntSchemaReq"] = APIEndpointLearntSchemaReqValidator()
	vr["ves.io.schema.app_type.APIEndpointLearntSchemaRsp"] = APIEndpointLearntSchemaRspValidator()
	vr["ves.io.schema.app_type.APIEndpointPDFReq"] = APIEndpointPDFReqValidator()
	vr["ves.io.schema.app_type.APIEndpointPDFRsp"] = APIEndpointPDFRspValidator()
	vr["ves.io.schema.app_type.APIEndpointsReq"] = APIEndpointsReqValidator()
	vr["ves.io.schema.app_type.APIEndpointsRsp"] = APIEndpointsRspValidator()
	vr["ves.io.schema.app_type.OverridePopReq"] = OverridePopReqValidator()
	vr["ves.io.schema.app_type.OverridePopRsp"] = OverridePopRspValidator()
	vr["ves.io.schema.app_type.OverridePushReq"] = OverridePushReqValidator()
	vr["ves.io.schema.app_type.OverridePushRsp"] = OverridePushRspValidator()
	vr["ves.io.schema.app_type.OverridesReq"] = OverridesReqValidator()
	vr["ves.io.schema.app_type.OverridesRsp"] = OverridesRspValidator()
	vr["ves.io.schema.app_type.ServiceAPIEndpointPDFReq"] = ServiceAPIEndpointPDFReqValidator()
	vr["ves.io.schema.app_type.ServiceAPIEndpointsReq"] = ServiceAPIEndpointsReqValidator()
	vr["ves.io.schema.app_type.SwaggerSpecReq"] = SwaggerSpecReqValidator()
	vr["ves.io.schema.app_type.SwaggerSpecRsp"] = SwaggerSpecRspValidator()

	vr["ves.io.schema.app_type.APIEPDynExample"] = APIEPDynExampleValidator()
	vr["ves.io.schema.app_type.APIEPInfo"] = APIEPInfoValidator()
	vr["ves.io.schema.app_type.APIEPPDFInfo"] = APIEPPDFInfoValidator()
	vr["ves.io.schema.app_type.APIEndpoint"] = APIEndpointValidator()
	vr["ves.io.schema.app_type.AuthData"] = AuthDataValidator()
	vr["ves.io.schema.app_type.Authentication"] = AuthenticationValidator()
	vr["ves.io.schema.app_type.AuthenticationTypeLocPair"] = AuthenticationTypeLocPairValidator()
	vr["ves.io.schema.app_type.BuiltInSensitiveDataType"] = BuiltInSensitiveDataTypeValidator()
	vr["ves.io.schema.app_type.BusinessLogicMarkupSetting"] = BusinessLogicMarkupSettingValidator()
	vr["ves.io.schema.app_type.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.app_type.CustomDataDetectionConfig"] = CustomDataDetectionConfigValidator()
	vr["ves.io.schema.app_type.CustomSections"] = CustomSectionsValidator()
	vr["ves.io.schema.app_type.CustomSensitiveDataDetectionRule"] = CustomSensitiveDataDetectionRuleValidator()
	vr["ves.io.schema.app_type.CustomSensitiveDataType"] = CustomSensitiveDataTypeValidator()
	vr["ves.io.schema.app_type.DiscoveredAPISettings"] = DiscoveredAPISettingsValidator()
	vr["ves.io.schema.app_type.DiscoveredSchema"] = DiscoveredSchemaValidator()
	vr["ves.io.schema.app_type.Feature"] = FeatureValidator()
	vr["ves.io.schema.app_type.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.app_type.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.app_type.KeyPattern"] = KeyPatternValidator()
	vr["ves.io.schema.app_type.KeyValuePattern"] = KeyValuePatternValidator()
	vr["ves.io.schema.app_type.OverrideInfo"] = OverrideInfoValidator()
	vr["ves.io.schema.app_type.PDFSpec"] = PDFSpecValidator()
	vr["ves.io.schema.app_type.PDFStat"] = PDFStatValidator()
	vr["ves.io.schema.app_type.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.app_type.RequestSchema"] = RequestSchemaValidator()
	vr["ves.io.schema.app_type.ResponseSchema"] = ResponseSchemaValidator()
	vr["ves.io.schema.app_type.RiskScore"] = RiskScoreValidator()
	vr["ves.io.schema.app_type.SchemaStruct"] = SchemaStructValidator()
	vr["ves.io.schema.app_type.SensitiveData"] = SensitiveDataValidator()
	vr["ves.io.schema.app_type.SensitiveDataDetectionRules"] = SensitiveDataDetectionRulesValidator()
	vr["ves.io.schema.app_type.ValuePattern"] = ValuePatternValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.app_type.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.app_type.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.app_type.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.app_type.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.app_type.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.app_type.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.app_type.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.app_type.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.app_type.API.Create"] = []string{
		"spec.business_logic_markup_setting.sensitive_data_detection_rules.custom_sensitive_data_detection_rules.#.metadata.disable",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.app_type.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.app_type.API.Replace"] = []string{
		"spec.business_logic_markup_setting.sensitive_data_detection_rules.custom_sensitive_data_detection_rules.#.metadata.disable",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.app_type.CustomAPI.GetAPIEndpointLearntSchema"] = []string{
		"discovered_openapi_spec",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.app_type.API"] = "config"
	sm["ves.io.schema.app_type.CustomAPI"] = "ml/data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.app_type.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.app_type.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.app_type.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.app_type.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.app_type.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.app_type.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.app_type.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.app_type.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.app_type.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.app_type.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.app_type.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.app_type.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.app_type.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
