// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:web/v20181101:Certificate":
		r = &Certificate{}
	case "azure-native:web/v20181101:WebApp":
		r = &WebApp{}
	case "azure-native:web/v20181101:WebAppApplicationSettings":
		r = &WebAppApplicationSettings{}
	case "azure-native:web/v20181101:WebAppApplicationSettingsSlot":
		r = &WebAppApplicationSettingsSlot{}
	case "azure-native:web/v20181101:WebAppAuthSettings":
		r = &WebAppAuthSettings{}
	case "azure-native:web/v20181101:WebAppAuthSettingsSlot":
		r = &WebAppAuthSettingsSlot{}
	case "azure-native:web/v20181101:WebAppAzureStorageAccounts":
		r = &WebAppAzureStorageAccounts{}
	case "azure-native:web/v20181101:WebAppAzureStorageAccountsSlot":
		r = &WebAppAzureStorageAccountsSlot{}
	case "azure-native:web/v20181101:WebAppBackupConfiguration":
		r = &WebAppBackupConfiguration{}
	case "azure-native:web/v20181101:WebAppBackupConfigurationSlot":
		r = &WebAppBackupConfigurationSlot{}
	case "azure-native:web/v20181101:WebAppConnectionStrings":
		r = &WebAppConnectionStrings{}
	case "azure-native:web/v20181101:WebAppConnectionStringsSlot":
		r = &WebAppConnectionStringsSlot{}
	case "azure-native:web/v20181101:WebAppDeployment":
		r = &WebAppDeployment{}
	case "azure-native:web/v20181101:WebAppDeploymentSlot":
		r = &WebAppDeploymentSlot{}
	case "azure-native:web/v20181101:WebAppDiagnosticLogsConfiguration":
		r = &WebAppDiagnosticLogsConfiguration{}
	case "azure-native:web/v20181101:WebAppDomainOwnershipIdentifier":
		r = &WebAppDomainOwnershipIdentifier{}
	case "azure-native:web/v20181101:WebAppDomainOwnershipIdentifierSlot":
		r = &WebAppDomainOwnershipIdentifierSlot{}
	case "azure-native:web/v20181101:WebAppFunction":
		r = &WebAppFunction{}
	case "azure-native:web/v20181101:WebAppHostNameBinding":
		r = &WebAppHostNameBinding{}
	case "azure-native:web/v20181101:WebAppHostNameBindingSlot":
		r = &WebAppHostNameBindingSlot{}
	case "azure-native:web/v20181101:WebAppHybridConnection":
		r = &WebAppHybridConnection{}
	case "azure-native:web/v20181101:WebAppHybridConnectionSlot":
		r = &WebAppHybridConnectionSlot{}
	case "azure-native:web/v20181101:WebAppInstanceFunctionSlot":
		r = &WebAppInstanceFunctionSlot{}
	case "azure-native:web/v20181101:WebAppMetadata":
		r = &WebAppMetadata{}
	case "azure-native:web/v20181101:WebAppMetadataSlot":
		r = &WebAppMetadataSlot{}
	case "azure-native:web/v20181101:WebAppPremierAddOn":
		r = &WebAppPremierAddOn{}
	case "azure-native:web/v20181101:WebAppPremierAddOnSlot":
		r = &WebAppPremierAddOnSlot{}
	case "azure-native:web/v20181101:WebAppPublicCertificate":
		r = &WebAppPublicCertificate{}
	case "azure-native:web/v20181101:WebAppPublicCertificateSlot":
		r = &WebAppPublicCertificateSlot{}
	case "azure-native:web/v20181101:WebAppRelayServiceConnection":
		r = &WebAppRelayServiceConnection{}
	case "azure-native:web/v20181101:WebAppRelayServiceConnectionSlot":
		r = &WebAppRelayServiceConnectionSlot{}
	case "azure-native:web/v20181101:WebAppSiteExtension":
		r = &WebAppSiteExtension{}
	case "azure-native:web/v20181101:WebAppSiteExtensionSlot":
		r = &WebAppSiteExtensionSlot{}
	case "azure-native:web/v20181101:WebAppSitePushSettings":
		r = &WebAppSitePushSettings{}
	case "azure-native:web/v20181101:WebAppSitePushSettingsSlot":
		r = &WebAppSitePushSettingsSlot{}
	case "azure-native:web/v20181101:WebAppSlot":
		r = &WebAppSlot{}
	case "azure-native:web/v20181101:WebAppSlotConfigurationNames":
		r = &WebAppSlotConfigurationNames{}
	case "azure-native:web/v20181101:WebAppSourceControl":
		r = &WebAppSourceControl{}
	case "azure-native:web/v20181101:WebAppSourceControlSlot":
		r = &WebAppSourceControlSlot{}
	case "azure-native:web/v20181101:WebAppSwiftVirtualNetworkConnection":
		r = &WebAppSwiftVirtualNetworkConnection{}
	case "azure-native:web/v20181101:WebAppSwiftVirtualNetworkConnectionSlot":
		r = &WebAppSwiftVirtualNetworkConnectionSlot{}
	case "azure-native:web/v20181101:WebAppVnetConnection":
		r = &WebAppVnetConnection{}
	case "azure-native:web/v20181101:WebAppVnetConnectionSlot":
		r = &WebAppVnetConnectionSlot{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := pulumiazurenativesdk.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"web/v20181101",
		&module{version},
	)
}