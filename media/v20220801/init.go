// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

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
	case "azure-native:media/v20220801:AccountFilter":
		r = &AccountFilter{}
	case "azure-native:media/v20220801:Asset":
		r = &Asset{}
	case "azure-native:media/v20220801:AssetFilter":
		r = &AssetFilter{}
	case "azure-native:media/v20220801:ContentKeyPolicy":
		r = &ContentKeyPolicy{}
	case "azure-native:media/v20220801:LiveEvent":
		r = &LiveEvent{}
	case "azure-native:media/v20220801:LiveOutput":
		r = &LiveOutput{}
	case "azure-native:media/v20220801:StreamingEndpoint":
		r = &StreamingEndpoint{}
	case "azure-native:media/v20220801:StreamingLocator":
		r = &StreamingLocator{}
	case "azure-native:media/v20220801:StreamingPolicy":
		r = &StreamingPolicy{}
	case "azure-native:media/v20220801:Track":
		r = &Track{}
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
		"media/v20220801",
		&module{version},
	)
}