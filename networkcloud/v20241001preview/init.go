// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20241001preview

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
	case "azure-native:networkcloud/v20241001preview:AgentPool":
		r = &AgentPool{}
	case "azure-native:networkcloud/v20241001preview:BareMetalMachine":
		r = &BareMetalMachine{}
	case "azure-native:networkcloud/v20241001preview:BareMetalMachineKeySet":
		r = &BareMetalMachineKeySet{}
	case "azure-native:networkcloud/v20241001preview:BmcKeySet":
		r = &BmcKeySet{}
	case "azure-native:networkcloud/v20241001preview:CloudServicesNetwork":
		r = &CloudServicesNetwork{}
	case "azure-native:networkcloud/v20241001preview:Cluster":
		r = &Cluster{}
	case "azure-native:networkcloud/v20241001preview:ClusterManager":
		r = &ClusterManager{}
	case "azure-native:networkcloud/v20241001preview:Console":
		r = &Console{}
	case "azure-native:networkcloud/v20241001preview:KubernetesCluster":
		r = &KubernetesCluster{}
	case "azure-native:networkcloud/v20241001preview:KubernetesClusterFeature":
		r = &KubernetesClusterFeature{}
	case "azure-native:networkcloud/v20241001preview:L2Network":
		r = &L2Network{}
	case "azure-native:networkcloud/v20241001preview:L3Network":
		r = &L3Network{}
	case "azure-native:networkcloud/v20241001preview:MetricsConfiguration":
		r = &MetricsConfiguration{}
	case "azure-native:networkcloud/v20241001preview:Rack":
		r = &Rack{}
	case "azure-native:networkcloud/v20241001preview:StorageAppliance":
		r = &StorageAppliance{}
	case "azure-native:networkcloud/v20241001preview:TrunkedNetwork":
		r = &TrunkedNetwork{}
	case "azure-native:networkcloud/v20241001preview:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:networkcloud/v20241001preview:Volume":
		r = &Volume{}
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
		"networkcloud/v20241001preview",
		&module{version},
	)
}
