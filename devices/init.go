// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devices

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
	case "azure-native:devices:Certificate":
		r = &Certificate{}
	case "azure-native:devices:DpsCertificate":
		r = &DpsCertificate{}
	case "azure-native:devices:IotDpsResource":
		r = &IotDpsResource{}
	case "azure-native:devices:IotDpsResourcePrivateEndpointConnection":
		r = &IotDpsResourcePrivateEndpointConnection{}
	case "azure-native:devices:IotHubResource":
		r = &IotHubResource{}
	case "azure-native:devices:IotHubResourceEventHubConsumerGroup":
		r = &IotHubResourceEventHubConsumerGroup{}
	case "azure-native:devices:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
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
		"devices",
		&module{version},
	)
}