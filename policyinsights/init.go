// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

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
	case "azure-native:policyinsights:AttestationAtResource":
		r = &AttestationAtResource{}
	case "azure-native:policyinsights:AttestationAtResourceGroup":
		r = &AttestationAtResourceGroup{}
	case "azure-native:policyinsights:AttestationAtSubscription":
		r = &AttestationAtSubscription{}
	case "azure-native:policyinsights:RemediationAtManagementGroup":
		r = &RemediationAtManagementGroup{}
	case "azure-native:policyinsights:RemediationAtResource":
		r = &RemediationAtResource{}
	case "azure-native:policyinsights:RemediationAtResourceGroup":
		r = &RemediationAtResourceGroup{}
	case "azure-native:policyinsights:RemediationAtSubscription":
		r = &RemediationAtSubscription{}
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
		"policyinsights",
		&module{version},
	)
}
