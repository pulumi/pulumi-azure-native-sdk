// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

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
	case "azure-native:workloads/v20211201preview:PhpWorkload":
		r = &PhpWorkload{}
	case "azure-native:workloads/v20211201preview:ProviderInstance":
		r = &ProviderInstance{}
	case "azure-native:workloads/v20211201preview:SAPApplicationServerInstance":
		r = &SAPApplicationServerInstance{}
	case "azure-native:workloads/v20211201preview:SAPCentralInstance":
		r = &SAPCentralInstance{}
	case "azure-native:workloads/v20211201preview:SAPDatabaseInstance":
		r = &SAPDatabaseInstance{}
	case "azure-native:workloads/v20211201preview:SAPVirtualInstance":
		r = &SAPVirtualInstance{}
	case "azure-native:workloads/v20211201preview:WordpressInstance":
		r = &WordpressInstance{}
	case "azure-native:workloads/v20211201preview:monitor":
		r = &Monitor{}
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
		"workloads/v20211201preview",
		&module{version},
	)
}