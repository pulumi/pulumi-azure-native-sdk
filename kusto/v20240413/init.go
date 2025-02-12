// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240413

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
	case "azure-native:kusto/v20240413:AttachedDatabaseConfiguration":
		r = &AttachedDatabaseConfiguration{}
	case "azure-native:kusto/v20240413:Cluster":
		r = &Cluster{}
	case "azure-native:kusto/v20240413:ClusterPrincipalAssignment":
		r = &ClusterPrincipalAssignment{}
	case "azure-native:kusto/v20240413:CosmosDbDataConnection":
		r = &CosmosDbDataConnection{}
	case "azure-native:kusto/v20240413:DatabasePrincipalAssignment":
		r = &DatabasePrincipalAssignment{}
	case "azure-native:kusto/v20240413:EventGridDataConnection":
		r = &EventGridDataConnection{}
	case "azure-native:kusto/v20240413:EventHubDataConnection":
		r = &EventHubDataConnection{}
	case "azure-native:kusto/v20240413:IotHubDataConnection":
		r = &IotHubDataConnection{}
	case "azure-native:kusto/v20240413:ManagedPrivateEndpoint":
		r = &ManagedPrivateEndpoint{}
	case "azure-native:kusto/v20240413:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:kusto/v20240413:ReadOnlyFollowingDatabase":
		r = &ReadOnlyFollowingDatabase{}
	case "azure-native:kusto/v20240413:ReadWriteDatabase":
		r = &ReadWriteDatabase{}
	case "azure-native:kusto/v20240413:SandboxCustomImage":
		r = &SandboxCustomImage{}
	case "azure-native:kusto/v20240413:Script":
		r = &Script{}
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
		"kusto/v20240413",
		&module{version},
	)
}
