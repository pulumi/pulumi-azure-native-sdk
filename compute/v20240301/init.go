// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

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
	case "azure-native:compute/v20240301:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20240301:CapacityReservation":
		r = &CapacityReservation{}
	case "azure-native:compute/v20240301:CapacityReservationGroup":
		r = &CapacityReservationGroup{}
	case "azure-native:compute/v20240301:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20240301:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20240301:Image":
		r = &Image{}
	case "azure-native:compute/v20240301:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20240301:RestorePoint":
		r = &RestorePoint{}
	case "azure-native:compute/v20240301:RestorePointCollection":
		r = &RestorePointCollection{}
	case "azure-native:compute/v20240301:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute/v20240301:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20240301:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20240301:VirtualMachineRunCommandByVirtualMachine":
		r = &VirtualMachineRunCommandByVirtualMachine{}
	case "azure-native:compute/v20240301:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20240301:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20240301:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute/v20240301:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
	case "azure-native:compute/v20240301:VirtualMachineScaleSetVMRunCommand":
		r = &VirtualMachineScaleSetVMRunCommand{}
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
		"compute/v20240301",
		&module{version},
	)
}
