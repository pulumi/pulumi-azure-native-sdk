// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240601

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
	case "azure-native:scvmm/v20240601:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:scvmm/v20240601:Cloud":
		r = &Cloud{}
	case "azure-native:scvmm/v20240601:GuestAgent":
		r = &GuestAgent{}
	case "azure-native:scvmm/v20240601:InventoryItem":
		r = &InventoryItem{}
	case "azure-native:scvmm/v20240601:VirtualMachineInstance":
		r = &VirtualMachineInstance{}
	case "azure-native:scvmm/v20240601:VirtualMachineTemplate":
		r = &VirtualMachineTemplate{}
	case "azure-native:scvmm/v20240601:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:scvmm/v20240601:VmmServer":
		r = &VmmServer{}
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
		"scvmm/v20240601",
		&module{version},
	)
}
