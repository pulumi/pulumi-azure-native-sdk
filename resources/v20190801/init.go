// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

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
	case "azure-native:resources/v20190801:Deployment":
		r = &Deployment{}
	case "azure-native:resources/v20190801:DeploymentAtManagementGroupScope":
		r = &DeploymentAtManagementGroupScope{}
	case "azure-native:resources/v20190801:DeploymentAtScope":
		r = &DeploymentAtScope{}
	case "azure-native:resources/v20190801:DeploymentAtSubscriptionScope":
		r = &DeploymentAtSubscriptionScope{}
	case "azure-native:resources/v20190801:DeploymentAtTenantScope":
		r = &DeploymentAtTenantScope{}
	case "azure-native:resources/v20190801:Resource":
		r = &Resource{}
	case "azure-native:resources/v20190801:ResourceGroup":
		r = &ResourceGroup{}
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
		"resources/v20190801",
		&module{version},
	)
}