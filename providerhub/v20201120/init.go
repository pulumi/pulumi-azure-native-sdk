// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201120

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
	case "azure-native:providerhub/v20201120:DefaultRollout":
		r = &DefaultRollout{}
	case "azure-native:providerhub/v20201120:NotificationRegistration":
		r = &NotificationRegistration{}
	case "azure-native:providerhub/v20201120:OperationByProviderRegistration":
		r = &OperationByProviderRegistration{}
	case "azure-native:providerhub/v20201120:ProviderRegistration":
		r = &ProviderRegistration{}
	case "azure-native:providerhub/v20201120:ResourceTypeRegistration":
		r = &ResourceTypeRegistration{}
	case "azure-native:providerhub/v20201120:Skus":
		r = &Skus{}
	case "azure-native:providerhub/v20201120:SkusNestedResourceTypeFirst":
		r = &SkusNestedResourceTypeFirst{}
	case "azure-native:providerhub/v20201120:SkusNestedResourceTypeSecond":
		r = &SkusNestedResourceTypeSecond{}
	case "azure-native:providerhub/v20201120:SkusNestedResourceTypeThird":
		r = &SkusNestedResourceTypeThird{}
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
		"providerhub/v20201120",
		&module{version},
	)
}