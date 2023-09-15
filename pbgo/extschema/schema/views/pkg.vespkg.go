// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package views

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.AWSNATGatewaychoiceType"] = AWSNATGatewaychoiceTypeValidator()
	vr["ves.io.schema.views.AWSSubnetIdsType"] = AWSSubnetIdsTypeValidator()
	vr["ves.io.schema.views.AWSSubnetInfoType"] = AWSSubnetInfoTypeValidator()
	vr["ves.io.schema.views.AWSVPCOneInterfaceNodeType"] = AWSVPCOneInterfaceNodeTypeValidator()
	vr["ves.io.schema.views.AWSVPCParamsType"] = AWSVPCParamsTypeValidator()
	vr["ves.io.schema.views.AWSVPCTwoInterfaceNodeType"] = AWSVPCTwoInterfaceNodeTypeValidator()
	vr["ves.io.schema.views.AWSVPCchoiceType"] = AWSVPCchoiceTypeValidator()
	vr["ves.io.schema.views.AWSVirtualPrivateGatewaychoiceType"] = AWSVirtualPrivateGatewaychoiceTypeValidator()
	vr["ves.io.schema.views.AllowedVIPPorts"] = AllowedVIPPortsValidator()
	vr["ves.io.schema.views.AzureSpecialSubnetType"] = AzureSpecialSubnetTypeValidator()
	vr["ves.io.schema.views.AzureSubnetChoiceType"] = AzureSubnetChoiceTypeValidator()
	vr["ves.io.schema.views.AzureSubnetChoiceWithAutoType"] = AzureSubnetChoiceWithAutoTypeValidator()
	vr["ves.io.schema.views.AzureSubnetType"] = AzureSubnetTypeValidator()
	vr["ves.io.schema.views.AzureVnetChoiceType"] = AzureVnetChoiceTypeValidator()
	vr["ves.io.schema.views.AzureVnetOneInterfaceNodeARType"] = AzureVnetOneInterfaceNodeARTypeValidator()
	vr["ves.io.schema.views.AzureVnetOneInterfaceNodeType"] = AzureVnetOneInterfaceNodeTypeValidator()
	vr["ves.io.schema.views.AzureVnetParamsType"] = AzureVnetParamsTypeValidator()
	vr["ves.io.schema.views.AzureVnetTwoInterfaceNodeARType"] = AzureVnetTwoInterfaceNodeARTypeValidator()
	vr["ves.io.schema.views.AzureVnetTwoInterfaceNodeType"] = AzureVnetTwoInterfaceNodeTypeValidator()
	vr["ves.io.schema.views.AzureVnetType"] = AzureVnetTypeValidator()
	vr["ves.io.schema.views.CloudLinkADNType"] = CloudLinkADNTypeValidator()
	vr["ves.io.schema.views.CloudSubnetParamType"] = CloudSubnetParamTypeValidator()
	vr["ves.io.schema.views.CloudSubnetType"] = CloudSubnetTypeValidator()
	vr["ves.io.schema.views.CustomPorts"] = CustomPortsValidator()
	vr["ves.io.schema.views.DCGWBGPPeerType"] = DCGWBGPPeerTypeValidator()
	vr["ves.io.schema.views.DirectConnectConfigType"] = DirectConnectConfigTypeValidator()
	vr["ves.io.schema.views.DirectConnectInfo"] = DirectConnectInfoValidator()
	vr["ves.io.schema.views.GCPSubnetParamsType"] = GCPSubnetParamsTypeValidator()
	vr["ves.io.schema.views.GCPSubnetType"] = GCPSubnetTypeValidator()
	vr["ves.io.schema.views.GCPVPCNetworkAutogenerateParamsType"] = GCPVPCNetworkAutogenerateParamsTypeValidator()
	vr["ves.io.schema.views.GCPVPCNetworkChoiceType"] = GCPVPCNetworkChoiceTypeValidator()
	vr["ves.io.schema.views.GCPVPCNetworkParamsType"] = GCPVPCNetworkParamsTypeValidator()
	vr["ves.io.schema.views.GCPVPCNetworkType"] = GCPVPCNetworkTypeValidator()
	vr["ves.io.schema.views.GCPVPCSubnetChoiceType"] = GCPVPCSubnetChoiceTypeValidator()
	vr["ves.io.schema.views.GlobalNetworkConnectionListType"] = GlobalNetworkConnectionListTypeValidator()
	vr["ves.io.schema.views.GlobalNetworkConnectionType"] = GlobalNetworkConnectionTypeValidator()
	vr["ves.io.schema.views.HostedVIFConfigType"] = HostedVIFConfigTypeValidator()
	vr["ves.io.schema.views.OfflineSurvivabilityModeType"] = OfflineSurvivabilityModeTypeValidator()
	vr["ves.io.schema.views.SecurityGroupType"] = SecurityGroupTypeValidator()
	vr["ves.io.schema.views.SiteStaticRoutesListType"] = SiteStaticRoutesListTypeValidator()
	vr["ves.io.schema.views.SiteStaticRoutesType"] = SiteStaticRoutesTypeValidator()
	vr["ves.io.schema.views.VifRegionConfig"] = VifRegionConfigValidator()

	vr["ves.io.schema.views.StorageClassListType"] = StorageClassListTypeValidator()
	vr["ves.io.schema.views.StorageClassOpenebsEnterpriseType"] = StorageClassOpenebsEnterpriseTypeValidator()
	vr["ves.io.schema.views.StorageClassType"] = StorageClassTypeValidator()

	vr["ves.io.schema.views.AdvertiseCustom"] = AdvertiseCustomValidator()
	vr["ves.io.schema.views.AdvertisePublic"] = AdvertisePublicValidator()
	vr["ves.io.schema.views.AdvertiseSiteVsite"] = AdvertiseSiteVsiteValidator()
	vr["ves.io.schema.views.CustomCiphers"] = CustomCiphersValidator()
	vr["ves.io.schema.views.DownstreamTLSCertsParams"] = DownstreamTLSCertsParamsValidator()
	vr["ves.io.schema.views.DownstreamTlsParamsType"] = DownstreamTlsParamsTypeValidator()
	vr["ves.io.schema.views.DownstreamTlsValidationContext"] = DownstreamTlsValidationContextValidator()
	vr["ves.io.schema.views.GlobalConnectorType"] = GlobalConnectorTypeValidator()
	vr["ves.io.schema.views.InternetVIPInfo"] = InternetVIPInfoValidator()
	vr["ves.io.schema.views.InternetVIPListenerStatusType"] = InternetVIPListenerStatusTypeValidator()
	vr["ves.io.schema.views.InternetVIPStatus"] = InternetVIPStatusValidator()
	vr["ves.io.schema.views.InternetVIPTargetGroupStatusType"] = InternetVIPTargetGroupStatusTypeValidator()
	vr["ves.io.schema.views.L3PerformanceEnhancementType"] = L3PerformanceEnhancementTypeValidator()
	vr["ves.io.schema.views.LinkRefType"] = LinkRefTypeValidator()
	vr["ves.io.schema.views.MasterNode"] = MasterNodeValidator()
	vr["ves.io.schema.views.ObjectRefType"] = ObjectRefTypeValidator()
	vr["ves.io.schema.views.OperatingSystemType"] = OperatingSystemTypeValidator()
	vr["ves.io.schema.views.OriginPoolListType"] = OriginPoolListTypeValidator()
	vr["ves.io.schema.views.OriginPoolWithWeight"] = OriginPoolWithWeightValidator()
	vr["ves.io.schema.views.PerformanceEnhancementModeType"] = PerformanceEnhancementModeTypeValidator()
	vr["ves.io.schema.views.PrefixStringListType"] = PrefixStringListTypeValidator()
	vr["ves.io.schema.views.SiteLocator"] = SiteLocatorValidator()
	vr["ves.io.schema.views.SiteReferenceListType"] = SiteReferenceListTypeValidator()
	vr["ves.io.schema.views.TlsConfig"] = TlsConfigValidator()
	vr["ves.io.schema.views.VolterraSoftwareType"] = VolterraSoftwareTypeValidator()
	vr["ves.io.schema.views.WhereSite"] = WhereSiteValidator()
	vr["ves.io.schema.views.WhereType"] = WhereTypeValidator()
	vr["ves.io.schema.views.WhereTypeSiteVsite"] = WhereTypeSiteVsiteValidator()
	vr["ves.io.schema.views.WhereVK8SService"] = WhereVK8SServiceValidator()
	vr["ves.io.schema.views.WhereVirtualNetwork"] = WhereVirtualNetworkValidator()
	vr["ves.io.schema.views.WhereVirtualSite"] = WhereVirtualSiteValidator()
	vr["ves.io.schema.views.XfccHeaderKeys"] = XfccHeaderKeysValidator()

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
