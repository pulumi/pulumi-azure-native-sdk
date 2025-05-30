// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedcache

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
	case "azure-native:connectedcache:CacheNodesOperation":
		r = &CacheNodesOperation{}
	case "azure-native:connectedcache:EnterpriseCustomerOperation":
		r = &EnterpriseCustomerOperation{}
	case "azure-native:connectedcache:EnterpriseMccCacheNodesOperation":
		r = &EnterpriseMccCacheNodesOperation{}
	case "azure-native:connectedcache:EnterpriseMccCustomer":
		r = &EnterpriseMccCustomer{}
	case "azure-native:connectedcache:IspCacheNodesOperation":
		r = &IspCacheNodesOperation{}
	case "azure-native:connectedcache:IspCustomer":
		r = &IspCustomer{}
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
		"connectedcache",
		&module{version},
	)
}
