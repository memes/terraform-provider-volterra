// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package alert_receiver

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.alert_receiver.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.alert_receiver.Object"] = ObjectValidator()
	vr["ves.io.schema.alert_receiver.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.alert_receiver.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.alert_receiver.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.alert_receiver.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.alert_receiver.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.alert_receiver.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.alert_receiver.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.alert_receiver.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.alert_receiver.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.alert_receiver.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.alert_receiver.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.alert_receiver.ConfirmAlertReceiverRequest"] = ConfirmAlertReceiverRequestValidator()
	vr["ves.io.schema.alert_receiver.ConfirmAlertReceiverResponse"] = ConfirmAlertReceiverResponseValidator()
	vr["ves.io.schema.alert_receiver.TestAlertReceiverRequest"] = TestAlertReceiverRequestValidator()
	vr["ves.io.schema.alert_receiver.TestAlertReceiverResponse"] = TestAlertReceiverResponseValidator()
	vr["ves.io.schema.alert_receiver.VerifyAlertReceiverRequest"] = VerifyAlertReceiverRequestValidator()
	vr["ves.io.schema.alert_receiver.VerifyAlertReceiverResponse"] = VerifyAlertReceiverResponseValidator()

	vr["ves.io.schema.alert_receiver.AuthToken"] = AuthTokenValidator()
	vr["ves.io.schema.alert_receiver.CACertificateObj"] = CACertificateObjValidator()
	vr["ves.io.schema.alert_receiver.ClientCertificateObj"] = ClientCertificateObjValidator()
	vr["ves.io.schema.alert_receiver.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.alert_receiver.EmailConfig"] = EmailConfigValidator()
	vr["ves.io.schema.alert_receiver.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.alert_receiver.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.alert_receiver.HTTPConfig"] = HTTPConfigValidator()
	vr["ves.io.schema.alert_receiver.HttpBasicAuth"] = HttpBasicAuthValidator()
	vr["ves.io.schema.alert_receiver.OpsGenieConfig"] = OpsGenieConfigValidator()
	vr["ves.io.schema.alert_receiver.PagerDutyConfig"] = PagerDutyConfigValidator()
	vr["ves.io.schema.alert_receiver.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.alert_receiver.SMSConfig"] = SMSConfigValidator()
	vr["ves.io.schema.alert_receiver.SlackConfig"] = SlackConfigValidator()
	vr["ves.io.schema.alert_receiver.TLSConfig"] = TLSConfigValidator()
	vr["ves.io.schema.alert_receiver.UpstreamTlsValidationContext"] = UpstreamTlsValidationContextValidator()
	vr["ves.io.schema.alert_receiver.WebhookConfig"] = WebhookConfigValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.alert_receiver.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.alert_receiver.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.alert_receiver.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.alert_receiver.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.alert_receiver.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.alert_receiver.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.alert_receiver.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.alert_receiver.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.alert_receiver.API.Create"] = []string{
		"spec.opsgenie.api_key.blindfold_secret_info_internal",
		"spec.opsgenie.api_key.secret_encoding_type",
		"spec.opsgenie.api_key.vault_secret_info",
		"spec.opsgenie.api_key.wingman_secret_info",
		"spec.pagerduty.routing_key.blindfold_secret_info_internal",
		"spec.pagerduty.routing_key.secret_encoding_type",
		"spec.pagerduty.routing_key.vault_secret_info",
		"spec.pagerduty.routing_key.wingman_secret_info",
		"spec.slack.url.blindfold_secret_info_internal",
		"spec.slack.url.secret_encoding_type",
		"spec.slack.url.vault_secret_info",
		"spec.slack.url.wingman_secret_info",
		"spec.webhook.http_config.auth_token.token.blindfold_secret_info_internal",
		"spec.webhook.http_config.auth_token.token.secret_encoding_type",
		"spec.webhook.http_config.auth_token.token.vault_secret_info",
		"spec.webhook.http_config.auth_token.token.wingman_secret_info",
		"spec.webhook.http_config.basic_auth.password.blindfold_secret_info_internal",
		"spec.webhook.http_config.basic_auth.password.secret_encoding_type",
		"spec.webhook.http_config.basic_auth.password.vault_secret_info",
		"spec.webhook.http_config.basic_auth.password.wingman_secret_info",
		"spec.webhook.url.blindfold_secret_info_internal",
		"spec.webhook.url.secret_encoding_type",
		"spec.webhook.url.vault_secret_info",
		"spec.webhook.url.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.alert_receiver.API.Create"] = "ves.io.schema.alert_receiver.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.alert_receiver.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.alert_receiver.API.Replace"] = []string{
		"spec.opsgenie.api_key.blindfold_secret_info_internal",
		"spec.opsgenie.api_key.secret_encoding_type",
		"spec.opsgenie.api_key.vault_secret_info",
		"spec.opsgenie.api_key.wingman_secret_info",
		"spec.pagerduty.routing_key.blindfold_secret_info_internal",
		"spec.pagerduty.routing_key.secret_encoding_type",
		"spec.pagerduty.routing_key.vault_secret_info",
		"spec.pagerduty.routing_key.wingman_secret_info",
		"spec.slack.url.blindfold_secret_info_internal",
		"spec.slack.url.secret_encoding_type",
		"spec.slack.url.vault_secret_info",
		"spec.slack.url.wingman_secret_info",
		"spec.webhook.http_config.auth_token.token.blindfold_secret_info_internal",
		"spec.webhook.http_config.auth_token.token.secret_encoding_type",
		"spec.webhook.http_config.auth_token.token.vault_secret_info",
		"spec.webhook.http_config.auth_token.token.wingman_secret_info",
		"spec.webhook.http_config.basic_auth.password.blindfold_secret_info_internal",
		"spec.webhook.http_config.basic_auth.password.secret_encoding_type",
		"spec.webhook.http_config.basic_auth.password.vault_secret_info",
		"spec.webhook.http_config.basic_auth.password.wingman_secret_info",
		"spec.webhook.url.blindfold_secret_info_internal",
		"spec.webhook.url.secret_encoding_type",
		"spec.webhook.url.vault_secret_info",
		"spec.webhook.url.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.alert_receiver.API.Replace"] = "ves.io.schema.alert_receiver.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.alert_receiver.API"] = "config"
	sm["ves.io.schema.alert_receiver.CustomAPI"] = "alert"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.alert_receiver.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.alert_receiver.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.alert_receiver.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.alert_receiver.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.alert_receiver.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.alert_receiver.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.alert_receiver.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.alert_receiver.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.alert_receiver.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.alert_receiver.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.alert_receiver.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.alert_receiver.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.alert_receiver.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
