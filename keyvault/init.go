// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

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
	case "azure-native:keyvault:AccessPolicy":
		r = &AccessPolicy{}
	case "azure-native:keyvault:Key":
		r = &Key{}
	case "azure-native:keyvault:MHSMPrivateEndpointConnection":
		r = &MHSMPrivateEndpointConnection{}
	case "azure-native:keyvault:ManagedHsm":
		r = &ManagedHsm{}
	case "azure-native:keyvault:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:keyvault:Secret":
		r = &Secret{}
	case "azure-native:keyvault:Vault":
		r = &Vault{}
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
		"keyvault",
		&module{version},
	)
}
