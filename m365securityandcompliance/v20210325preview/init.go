// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210325preview

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
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsAdtAPI":
		r = &PrivateEndpointConnectionsAdtAPI{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsComp":
		r = &PrivateEndpointConnectionsComp{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForEDM":
		r = &PrivateEndpointConnectionsForEDM{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForMIPPolicySync":
		r = &PrivateEndpointConnectionsForMIPPolicySync{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForSCCPowershell":
		r = &PrivateEndpointConnectionsForSCCPowershell{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsSec":
		r = &PrivateEndpointConnectionsSec{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForEDMUpload":
		r = &PrivateLinkServicesForEDMUpload{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForM365ComplianceCenter":
		r = &PrivateLinkServicesForM365ComplianceCenter{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForM365SecurityCenter":
		r = &PrivateLinkServicesForM365SecurityCenter{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForMIPPolicySync":
		r = &PrivateLinkServicesForMIPPolicySync{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForO365ManagementActivityAPI":
		r = &PrivateLinkServicesForO365ManagementActivityAPI{}
	case "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForSCCPowershell":
		r = &PrivateLinkServicesForSCCPowershell{}
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
		"m365securityandcompliance/v20210325preview",
		&module{version},
	)
}
