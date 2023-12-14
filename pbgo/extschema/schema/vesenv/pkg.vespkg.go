// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package vesenv

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {

	vr["ves.io.schema.vesenv.AddonServiceChoice"] = AddonServiceChoiceValidator()

	vr["ves.io.schema.vesenv.APIGroupChoice"] = APIGroupChoiceValidator()

	vr["ves.io.schema.vesenv.RouteTargetChoice"] = RouteTargetChoiceValidator()

	vr["ves.io.schema.vesenv.ServiceChoice"] = ServiceChoiceValidator()

	vr["ves.io.schema.vesenv.APIGroupElementInfo"] = APIGroupElementInfoValidator()
	vr["ves.io.schema.vesenv.APIGroupElementItem"] = APIGroupElementItemValidator()
	vr["ves.io.schema.vesenv.APIGroupNameMap"] = APIGroupNameMapValidator()
	vr["ves.io.schema.vesenv.APIGroupNameMapItem"] = APIGroupNameMapItemValidator()
	vr["ves.io.schema.vesenv.AddonServiceInfo"] = AddonServiceInfoValidator()
	vr["ves.io.schema.vesenv.BFSecretChoice"] = BFSecretChoiceValidator()
	vr["ves.io.schema.vesenv.BFSecretInfo"] = BFSecretInfoValidator()
	vr["ves.io.schema.vesenv.NameToUid"] = NameToUidValidator()
	vr["ves.io.schema.vesenv.QuotaResourceKeyInfo"] = QuotaResourceKeyInfoValidator()
	vr["ves.io.schema.vesenv.ReEncryptSecretItemType"] = ReEncryptSecretItemTypeValidator()
	vr["ves.io.schema.vesenv.ReEncryptSecretsType"] = ReEncryptSecretsTypeValidator()
	vr["ves.io.schema.vesenv.RouteTargetInfo"] = RouteTargetInfoValidator()
	vr["ves.io.schema.vesenv.ServiceInfo"] = ServiceInfoValidator()
	vr["ves.io.schema.vesenv.ServiceSlugChoice"] = ServiceSlugChoiceValidator()
	vr["ves.io.schema.vesenv.ServiceSlugInfo"] = ServiceSlugInfoValidator()

	vr["ves.io.schema.vesenv.QuotaResourceKeyChoice"] = QuotaResourceKeyChoiceValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

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
