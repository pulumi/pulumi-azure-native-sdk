// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250101

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
	case "azure-native:authorization/v20250101:PolicyAssignment":
		r = &PolicyAssignment{}
	case "azure-native:authorization/v20250101:PolicyDefinition":
		r = &PolicyDefinition{}
	case "azure-native:authorization/v20250101:PolicyDefinitionAtManagementGroup":
		r = &PolicyDefinitionAtManagementGroup{}
	case "azure-native:authorization/v20250101:PolicyDefinitionVersion":
		r = &PolicyDefinitionVersion{}
	case "azure-native:authorization/v20250101:PolicyDefinitionVersionAtManagementGroup":
		r = &PolicyDefinitionVersionAtManagementGroup{}
	case "azure-native:authorization/v20250101:PolicySetDefinition":
		r = &PolicySetDefinition{}
	case "azure-native:authorization/v20250101:PolicySetDefinitionAtManagementGroup":
		r = &PolicySetDefinitionAtManagementGroup{}
	case "azure-native:authorization/v20250101:PolicySetDefinitionVersion":
		r = &PolicySetDefinitionVersion{}
	case "azure-native:authorization/v20250101:PolicySetDefinitionVersionAtManagementGroup":
		r = &PolicySetDefinitionVersionAtManagementGroup{}
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
		"authorization/v20250101",
		&module{version},
	)
}
