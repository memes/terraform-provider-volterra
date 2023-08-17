// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package fleet

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.fleet.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.fleet.Object"] = ObjectValidator()
	vr["ves.io.schema.fleet.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.fleet.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.fleet.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.fleet.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.fleet.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.fleet.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.fleet.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.fleet.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.fleet.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.fleet.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.fleet.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.fleet.BGPConfiguration"] = BGPConfigurationValidator()
	vr["ves.io.schema.fleet.BlockedServices"] = BlockedServicesValidator()
	vr["ves.io.schema.fleet.BlockedServicesListType"] = BlockedServicesListTypeValidator()
	vr["ves.io.schema.fleet.BondLacpType"] = BondLacpTypeValidator()
	vr["ves.io.schema.fleet.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.fleet.CustomAllowedServiceList"] = CustomAllowedServiceListValidator()
	vr["ves.io.schema.fleet.CustomServiceItem"] = CustomServiceItemValidator()
	vr["ves.io.schema.fleet.DeviceInstanceType"] = DeviceInstanceTypeValidator()
	vr["ves.io.schema.fleet.DeviceNetappBackendOntapSanChapType"] = DeviceNetappBackendOntapSanChapTypeValidator()
	vr["ves.io.schema.fleet.FlashArrayEndpoint"] = FlashArrayEndpointValidator()
	vr["ves.io.schema.fleet.FlashArrayType"] = FlashArrayTypeValidator()
	vr["ves.io.schema.fleet.FlashBladeEndpoint"] = FlashBladeEndpointValidator()
	vr["ves.io.schema.fleet.FlashBladeType"] = FlashBladeTypeValidator()
	vr["ves.io.schema.fleet.FleetBondDeviceType"] = FleetBondDeviceTypeValidator()
	vr["ves.io.schema.fleet.FleetBondDevicesListType"] = FleetBondDevicesListTypeValidator()
	vr["ves.io.schema.fleet.FleetDeviceListType"] = FleetDeviceListTypeValidator()
	vr["ves.io.schema.fleet.FleetInterfaceListType"] = FleetInterfaceListTypeValidator()
	vr["ves.io.schema.fleet.FleetStatus"] = FleetStatusValidator()
	vr["ves.io.schema.fleet.FleetStorageClassListType"] = FleetStorageClassListTypeValidator()
	vr["ves.io.schema.fleet.FleetStorageClassType"] = FleetStorageClassTypeValidator()
	vr["ves.io.schema.fleet.FleetStorageDeviceListType"] = FleetStorageDeviceListTypeValidator()
	vr["ves.io.schema.fleet.FleetStorageDeviceType"] = FleetStorageDeviceTypeValidator()
	vr["ves.io.schema.fleet.FleetStorageStaticRoutesListType"] = FleetStorageStaticRoutesListTypeValidator()
	vr["ves.io.schema.fleet.GenericDeviceInstanceType"] = GenericDeviceInstanceTypeValidator()
	vr["ves.io.schema.fleet.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.fleet.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.fleet.LocalControlPlaneType"] = LocalControlPlaneTypeValidator()
	vr["ves.io.schema.fleet.NetworkingDeviceInstanceType"] = NetworkingDeviceInstanceTypeValidator()
	vr["ves.io.schema.fleet.OntapVirtualStoragePoolType"] = OntapVirtualStoragePoolTypeValidator()
	vr["ves.io.schema.fleet.OntapVolumeDefaults"] = OntapVolumeDefaultsValidator()
	vr["ves.io.schema.fleet.PsoArrayConfiguration"] = PsoArrayConfigurationValidator()
	vr["ves.io.schema.fleet.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.fleet.SriovInterface"] = SriovInterfaceValidator()
	vr["ves.io.schema.fleet.SriovInterfacesListType"] = SriovInterfacesListTypeValidator()
	vr["ves.io.schema.fleet.StorageClassCustomType"] = StorageClassCustomTypeValidator()
	vr["ves.io.schema.fleet.StorageClassDellIsilonF800Type"] = StorageClassDellIsilonF800TypeValidator()
	vr["ves.io.schema.fleet.StorageClassHPENimbusStorageAf40Type"] = StorageClassHPENimbusStorageAf40TypeValidator()
	vr["ves.io.schema.fleet.StorageClassHpeStorageType"] = StorageClassHpeStorageTypeValidator()
	vr["ves.io.schema.fleet.StorageClassNetappTridentType"] = StorageClassNetappTridentTypeValidator()
	vr["ves.io.schema.fleet.StorageClassOpenebsEnterpriseType"] = StorageClassOpenebsEnterpriseTypeValidator()
	vr["ves.io.schema.fleet.StorageClassPureServiceOrchestratorType"] = StorageClassPureServiceOrchestratorTypeValidator()
	vr["ves.io.schema.fleet.StorageDeviceDellIsilonF800Type"] = StorageDeviceDellIsilonF800TypeValidator()
	vr["ves.io.schema.fleet.StorageDeviceHPENimbusStorageAf40Type"] = StorageDeviceHPENimbusStorageAf40TypeValidator()
	vr["ves.io.schema.fleet.StorageDeviceHpeStorageType"] = StorageDeviceHpeStorageTypeValidator()
	vr["ves.io.schema.fleet.StorageDeviceNetappBackendOntapNasType"] = StorageDeviceNetappBackendOntapNasTypeValidator()
	vr["ves.io.schema.fleet.StorageDeviceNetappBackendOntapSanType"] = StorageDeviceNetappBackendOntapSanTypeValidator()
	vr["ves.io.schema.fleet.StorageDeviceNetappTridentType"] = StorageDeviceNetappTridentTypeValidator()
	vr["ves.io.schema.fleet.StorageDevicePureStorageServiceOrchestratorType"] = StorageDevicePureStorageServiceOrchestratorTypeValidator()
	vr["ves.io.schema.fleet.VGPUConfiguration"] = VGPUConfigurationValidator()
	vr["ves.io.schema.fleet.VMConfiguration"] = VMConfigurationValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.fleet.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.fleet.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.fleet.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.fleet.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.fleet.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.fleet.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.fleet.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.fleet.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.fleet.API.Create"] = []string{
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.ipv6_prefixes.#",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.fleet.API.Create"] = "ves.io.schema.fleet.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.fleet.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.fleet.API.Replace"] = []string{
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.hpe_storage.iscsi_chap_password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.hpe_storage.password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.ipv6_prefixes.#",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.client_private_key.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_nas.password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.client_private_key.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.password.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_array.flash_arrays.#.api_token.wingman_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.blindfold_secret_info_internal",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.secret_encoding_type",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.vault_secret_info",
		"spec.storage_device_list.storage_devices.#.pure_service_orchestrator.arrays.flash_blade.flash_blades.#.api_token.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.fleet.API.Replace"] = "ves.io.schema.fleet.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.fleet.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.fleet.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.fleet.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.fleet.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.fleet.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.fleet.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.fleet.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.fleet.Object"] = NewCRUDAPIServer

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
