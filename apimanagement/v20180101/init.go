// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

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
	case "azure-native:apimanagement/v20180101:Api":
		r = &Api{}
	case "azure-native:apimanagement/v20180101:ApiDiagnostic":
		r = &ApiDiagnostic{}
	case "azure-native:apimanagement/v20180101:ApiDiagnosticLogger":
		r = &ApiDiagnosticLogger{}
	case "azure-native:apimanagement/v20180101:ApiIssue":
		r = &ApiIssue{}
	case "azure-native:apimanagement/v20180101:ApiIssueAttachment":
		r = &ApiIssueAttachment{}
	case "azure-native:apimanagement/v20180101:ApiIssueComment":
		r = &ApiIssueComment{}
	case "azure-native:apimanagement/v20180101:ApiManagementService":
		r = &ApiManagementService{}
	case "azure-native:apimanagement/v20180101:ApiOperation":
		r = &ApiOperation{}
	case "azure-native:apimanagement/v20180101:ApiOperationPolicy":
		r = &ApiOperationPolicy{}
	case "azure-native:apimanagement/v20180101:ApiPolicy":
		r = &ApiPolicy{}
	case "azure-native:apimanagement/v20180101:ApiRelease":
		r = &ApiRelease{}
	case "azure-native:apimanagement/v20180101:ApiSchema":
		r = &ApiSchema{}
	case "azure-native:apimanagement/v20180101:ApiVersionSet":
		r = &ApiVersionSet{}
	case "azure-native:apimanagement/v20180101:AuthorizationServer":
		r = &AuthorizationServer{}
	case "azure-native:apimanagement/v20180101:Backend":
		r = &Backend{}
	case "azure-native:apimanagement/v20180101:Certificate":
		r = &Certificate{}
	case "azure-native:apimanagement/v20180101:Diagnostic":
		r = &Diagnostic{}
	case "azure-native:apimanagement/v20180101:DiagnosticLogger":
		r = &DiagnosticLogger{}
	case "azure-native:apimanagement/v20180101:EmailTemplate":
		r = &EmailTemplate{}
	case "azure-native:apimanagement/v20180101:Group":
		r = &Group{}
	case "azure-native:apimanagement/v20180101:GroupUser":
		r = &GroupUser{}
	case "azure-native:apimanagement/v20180101:IdentityProvider":
		r = &IdentityProvider{}
	case "azure-native:apimanagement/v20180101:Logger":
		r = &Logger{}
	case "azure-native:apimanagement/v20180101:NotificationRecipientEmail":
		r = &NotificationRecipientEmail{}
	case "azure-native:apimanagement/v20180101:NotificationRecipientUser":
		r = &NotificationRecipientUser{}
	case "azure-native:apimanagement/v20180101:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "azure-native:apimanagement/v20180101:Policy":
		r = &Policy{}
	case "azure-native:apimanagement/v20180101:Product":
		r = &Product{}
	case "azure-native:apimanagement/v20180101:ProductApi":
		r = &ProductApi{}
	case "azure-native:apimanagement/v20180101:ProductGroup":
		r = &ProductGroup{}
	case "azure-native:apimanagement/v20180101:ProductPolicy":
		r = &ProductPolicy{}
	case "azure-native:apimanagement/v20180101:Property":
		r = &Property{}
	case "azure-native:apimanagement/v20180101:Subscription":
		r = &Subscription{}
	case "azure-native:apimanagement/v20180101:Tag":
		r = &Tag{}
	case "azure-native:apimanagement/v20180101:TagByApi":
		r = &TagByApi{}
	case "azure-native:apimanagement/v20180101:TagByOperation":
		r = &TagByOperation{}
	case "azure-native:apimanagement/v20180101:TagByProduct":
		r = &TagByProduct{}
	case "azure-native:apimanagement/v20180101:TagDescription":
		r = &TagDescription{}
	case "azure-native:apimanagement/v20180101:User":
		r = &User{}
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
		"apimanagement/v20180101",
		&module{version},
	)
}