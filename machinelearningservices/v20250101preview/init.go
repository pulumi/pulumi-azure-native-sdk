// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250101preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
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
	case "azure-native:machinelearningservices/v20250101preview:BatchDeployment":
		r = &BatchDeployment{}
	case "azure-native:machinelearningservices/v20250101preview:BatchEndpoint":
		r = &BatchEndpoint{}
	case "azure-native:machinelearningservices/v20250101preview:CapabilityHost":
		r = &CapabilityHost{}
	case "azure-native:machinelearningservices/v20250101preview:CodeContainer":
		r = &CodeContainer{}
	case "azure-native:machinelearningservices/v20250101preview:CodeVersion":
		r = &CodeVersion{}
	case "azure-native:machinelearningservices/v20250101preview:ComponentContainer":
		r = &ComponentContainer{}
	case "azure-native:machinelearningservices/v20250101preview:ComponentVersion":
		r = &ComponentVersion{}
	case "azure-native:machinelearningservices/v20250101preview:Compute":
		r = &Compute{}
	case "azure-native:machinelearningservices/v20250101preview:ConnectionDeployment":
		r = &ConnectionDeployment{}
	case "azure-native:machinelearningservices/v20250101preview:ConnectionRaiBlocklist":
		r = &ConnectionRaiBlocklist{}
	case "azure-native:machinelearningservices/v20250101preview:ConnectionRaiBlocklistItem":
		r = &ConnectionRaiBlocklistItem{}
	case "azure-native:machinelearningservices/v20250101preview:ConnectionRaiPolicy":
		r = &ConnectionRaiPolicy{}
	case "azure-native:machinelearningservices/v20250101preview:DataContainer":
		r = &DataContainer{}
	case "azure-native:machinelearningservices/v20250101preview:DataVersion":
		r = &DataVersion{}
	case "azure-native:machinelearningservices/v20250101preview:Datastore":
		r = &Datastore{}
	case "azure-native:machinelearningservices/v20250101preview:EndpointDeployment":
		r = &EndpointDeployment{}
	case "azure-native:machinelearningservices/v20250101preview:EnvironmentContainer":
		r = &EnvironmentContainer{}
	case "azure-native:machinelearningservices/v20250101preview:EnvironmentVersion":
		r = &EnvironmentVersion{}
	case "azure-native:machinelearningservices/v20250101preview:FeaturesetContainerEntity":
		r = &FeaturesetContainerEntity{}
	case "azure-native:machinelearningservices/v20250101preview:FeaturesetVersion":
		r = &FeaturesetVersion{}
	case "azure-native:machinelearningservices/v20250101preview:FeaturestoreEntityContainerEntity":
		r = &FeaturestoreEntityContainerEntity{}
	case "azure-native:machinelearningservices/v20250101preview:FeaturestoreEntityVersion":
		r = &FeaturestoreEntityVersion{}
	case "azure-native:machinelearningservices/v20250101preview:InferenceEndpoint":
		r = &InferenceEndpoint{}
	case "azure-native:machinelearningservices/v20250101preview:InferenceGroup":
		r = &InferenceGroup{}
	case "azure-native:machinelearningservices/v20250101preview:InferencePool":
		r = &InferencePool{}
	case "azure-native:machinelearningservices/v20250101preview:Job":
		r = &Job{}
	case "azure-native:machinelearningservices/v20250101preview:ManagedNetworkSettingsRule":
		r = &ManagedNetworkSettingsRule{}
	case "azure-native:machinelearningservices/v20250101preview:MarketplaceSubscription":
		r = &MarketplaceSubscription{}
	case "azure-native:machinelearningservices/v20250101preview:ModelContainer":
		r = &ModelContainer{}
	case "azure-native:machinelearningservices/v20250101preview:ModelVersion":
		r = &ModelVersion{}
	case "azure-native:machinelearningservices/v20250101preview:OnlineDeployment":
		r = &OnlineDeployment{}
	case "azure-native:machinelearningservices/v20250101preview:OnlineEndpoint":
		r = &OnlineEndpoint{}
	case "azure-native:machinelearningservices/v20250101preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:machinelearningservices/v20250101preview:RaiPolicy":
		r = &RaiPolicy{}
	case "azure-native:machinelearningservices/v20250101preview:Registry":
		r = &Registry{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryCodeContainer":
		r = &RegistryCodeContainer{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryCodeVersion":
		r = &RegistryCodeVersion{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryComponentContainer":
		r = &RegistryComponentContainer{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryComponentVersion":
		r = &RegistryComponentVersion{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryDataContainer":
		r = &RegistryDataContainer{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryDataVersion":
		r = &RegistryDataVersion{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryEnvironmentContainer":
		r = &RegistryEnvironmentContainer{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryEnvironmentVersion":
		r = &RegistryEnvironmentVersion{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryModelContainer":
		r = &RegistryModelContainer{}
	case "azure-native:machinelearningservices/v20250101preview:RegistryModelVersion":
		r = &RegistryModelVersion{}
	case "azure-native:machinelearningservices/v20250101preview:Schedule":
		r = &Schedule{}
	case "azure-native:machinelearningservices/v20250101preview:ServerlessEndpoint":
		r = &ServerlessEndpoint{}
	case "azure-native:machinelearningservices/v20250101preview:Workspace":
		r = &Workspace{}
	case "azure-native:machinelearningservices/v20250101preview:WorkspaceConnection":
		r = &WorkspaceConnection{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"machinelearningservices/v20250101preview",
		&module{version},
	)
}
