// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240201

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
	case "azure-native:devcenter/v20240201:AttachedNetworkByDevCenter":
		r = &AttachedNetworkByDevCenter{}
	case "azure-native:devcenter/v20240201:Catalog":
		r = &Catalog{}
	case "azure-native:devcenter/v20240201:DevBoxDefinition":
		r = &DevBoxDefinition{}
	case "azure-native:devcenter/v20240201:DevCenter":
		r = &DevCenter{}
	case "azure-native:devcenter/v20240201:EnvironmentType":
		r = &EnvironmentType{}
	case "azure-native:devcenter/v20240201:Gallery":
		r = &Gallery{}
	case "azure-native:devcenter/v20240201:NetworkConnection":
		r = &NetworkConnection{}
	case "azure-native:devcenter/v20240201:Pool":
		r = &Pool{}
	case "azure-native:devcenter/v20240201:Project":
		r = &Project{}
	case "azure-native:devcenter/v20240201:ProjectCatalog":
		r = &ProjectCatalog{}
	case "azure-native:devcenter/v20240201:ProjectEnvironmentType":
		r = &ProjectEnvironmentType{}
	case "azure-native:devcenter/v20240201:Schedule":
		r = &Schedule{}
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
		"devcenter/v20240201",
		&module{version},
	)
}
