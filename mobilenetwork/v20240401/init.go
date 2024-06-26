// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401

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
	case "azure-native:mobilenetwork/v20240401:AttachedDataNetwork":
		r = &AttachedDataNetwork{}
	case "azure-native:mobilenetwork/v20240401:DataNetwork":
		r = &DataNetwork{}
	case "azure-native:mobilenetwork/v20240401:DiagnosticsPackage":
		r = &DiagnosticsPackage{}
	case "azure-native:mobilenetwork/v20240401:MobileNetwork":
		r = &MobileNetwork{}
	case "azure-native:mobilenetwork/v20240401:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:mobilenetwork/v20240401:PacketCoreControlPlane":
		r = &PacketCoreControlPlane{}
	case "azure-native:mobilenetwork/v20240401:PacketCoreDataPlane":
		r = &PacketCoreDataPlane{}
	case "azure-native:mobilenetwork/v20240401:Service":
		r = &Service{}
	case "azure-native:mobilenetwork/v20240401:Sim":
		r = &Sim{}
	case "azure-native:mobilenetwork/v20240401:SimGroup":
		r = &SimGroup{}
	case "azure-native:mobilenetwork/v20240401:SimPolicy":
		r = &SimPolicy{}
	case "azure-native:mobilenetwork/v20240401:Site":
		r = &Site{}
	case "azure-native:mobilenetwork/v20240401:Slice":
		r = &Slice{}
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
		"mobilenetwork/v20240401",
		&module{version},
	)
}
