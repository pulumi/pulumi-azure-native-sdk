// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package peering

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
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
	case "azure-native:peering:ConnectionMonitorTest":
		r = &ConnectionMonitorTest{}
	case "azure-native:peering:PeerAsn":
		r = &PeerAsn{}
	case "azure-native:peering:Peering":
		r = &Peering{}
	case "azure-native:peering:PeeringService":
		r = &PeeringService{}
	case "azure-native:peering:Prefix":
		r = &Prefix{}
	case "azure-native:peering:RegisteredAsn":
		r = &RegisteredAsn{}
	case "azure-native:peering:RegisteredPrefix":
		r = &RegisteredPrefix{}
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
		"peering",
		&module{version},
	)
}
