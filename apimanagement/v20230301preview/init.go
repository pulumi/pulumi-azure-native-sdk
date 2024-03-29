// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

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
	case "azure-native:apimanagement/v20230301preview:Api":
		r = &Api{}
	case "azure-native:apimanagement/v20230301preview:ApiDiagnostic":
		r = &ApiDiagnostic{}
	case "azure-native:apimanagement/v20230301preview:ApiIssue":
		r = &ApiIssue{}
	case "azure-native:apimanagement/v20230301preview:ApiIssueAttachment":
		r = &ApiIssueAttachment{}
	case "azure-native:apimanagement/v20230301preview:ApiIssueComment":
		r = &ApiIssueComment{}
	case "azure-native:apimanagement/v20230301preview:ApiManagementService":
		r = &ApiManagementService{}
	case "azure-native:apimanagement/v20230301preview:ApiOperation":
		r = &ApiOperation{}
	case "azure-native:apimanagement/v20230301preview:ApiOperationPolicy":
		r = &ApiOperationPolicy{}
	case "azure-native:apimanagement/v20230301preview:ApiPolicy":
		r = &ApiPolicy{}
	case "azure-native:apimanagement/v20230301preview:ApiRelease":
		r = &ApiRelease{}
	case "azure-native:apimanagement/v20230301preview:ApiSchema":
		r = &ApiSchema{}
	case "azure-native:apimanagement/v20230301preview:ApiTagDescription":
		r = &ApiTagDescription{}
	case "azure-native:apimanagement/v20230301preview:ApiVersionSet":
		r = &ApiVersionSet{}
	case "azure-native:apimanagement/v20230301preview:ApiWiki":
		r = &ApiWiki{}
	case "azure-native:apimanagement/v20230301preview:Authorization":
		r = &Authorization{}
	case "azure-native:apimanagement/v20230301preview:AuthorizationAccessPolicy":
		r = &AuthorizationAccessPolicy{}
	case "azure-native:apimanagement/v20230301preview:AuthorizationProvider":
		r = &AuthorizationProvider{}
	case "azure-native:apimanagement/v20230301preview:AuthorizationServer":
		r = &AuthorizationServer{}
	case "azure-native:apimanagement/v20230301preview:Backend":
		r = &Backend{}
	case "azure-native:apimanagement/v20230301preview:Cache":
		r = &Cache{}
	case "azure-native:apimanagement/v20230301preview:Certificate":
		r = &Certificate{}
	case "azure-native:apimanagement/v20230301preview:ContentItem":
		r = &ContentItem{}
	case "azure-native:apimanagement/v20230301preview:ContentType":
		r = &ContentType{}
	case "azure-native:apimanagement/v20230301preview:Diagnostic":
		r = &Diagnostic{}
	case "azure-native:apimanagement/v20230301preview:Documentation":
		r = &Documentation{}
	case "azure-native:apimanagement/v20230301preview:EmailTemplate":
		r = &EmailTemplate{}
	case "azure-native:apimanagement/v20230301preview:Gateway":
		r = &Gateway{}
	case "azure-native:apimanagement/v20230301preview:GatewayApiEntityTag":
		r = &GatewayApiEntityTag{}
	case "azure-native:apimanagement/v20230301preview:GatewayCertificateAuthority":
		r = &GatewayCertificateAuthority{}
	case "azure-native:apimanagement/v20230301preview:GatewayHostnameConfiguration":
		r = &GatewayHostnameConfiguration{}
	case "azure-native:apimanagement/v20230301preview:GlobalSchema":
		r = &GlobalSchema{}
	case "azure-native:apimanagement/v20230301preview:GraphQLApiResolver":
		r = &GraphQLApiResolver{}
	case "azure-native:apimanagement/v20230301preview:GraphQLApiResolverPolicy":
		r = &GraphQLApiResolverPolicy{}
	case "azure-native:apimanagement/v20230301preview:Group":
		r = &Group{}
	case "azure-native:apimanagement/v20230301preview:GroupUser":
		r = &GroupUser{}
	case "azure-native:apimanagement/v20230301preview:IdentityProvider":
		r = &IdentityProvider{}
	case "azure-native:apimanagement/v20230301preview:Logger":
		r = &Logger{}
	case "azure-native:apimanagement/v20230301preview:NamedValue":
		r = &NamedValue{}
	case "azure-native:apimanagement/v20230301preview:NotificationRecipientEmail":
		r = &NotificationRecipientEmail{}
	case "azure-native:apimanagement/v20230301preview:NotificationRecipientUser":
		r = &NotificationRecipientUser{}
	case "azure-native:apimanagement/v20230301preview:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "azure-native:apimanagement/v20230301preview:Policy":
		r = &Policy{}
	case "azure-native:apimanagement/v20230301preview:PolicyFragment":
		r = &PolicyFragment{}
	case "azure-native:apimanagement/v20230301preview:PrivateEndpointConnectionByName":
		r = &PrivateEndpointConnectionByName{}
	case "azure-native:apimanagement/v20230301preview:Product":
		r = &Product{}
	case "azure-native:apimanagement/v20230301preview:ProductApi":
		r = &ProductApi{}
	case "azure-native:apimanagement/v20230301preview:ProductApiLink":
		r = &ProductApiLink{}
	case "azure-native:apimanagement/v20230301preview:ProductGroup":
		r = &ProductGroup{}
	case "azure-native:apimanagement/v20230301preview:ProductGroupLink":
		r = &ProductGroupLink{}
	case "azure-native:apimanagement/v20230301preview:ProductPolicy":
		r = &ProductPolicy{}
	case "azure-native:apimanagement/v20230301preview:ProductWiki":
		r = &ProductWiki{}
	case "azure-native:apimanagement/v20230301preview:Subscription":
		r = &Subscription{}
	case "azure-native:apimanagement/v20230301preview:Tag":
		r = &Tag{}
	case "azure-native:apimanagement/v20230301preview:TagApiLink":
		r = &TagApiLink{}
	case "azure-native:apimanagement/v20230301preview:TagByApi":
		r = &TagByApi{}
	case "azure-native:apimanagement/v20230301preview:TagByOperation":
		r = &TagByOperation{}
	case "azure-native:apimanagement/v20230301preview:TagByProduct":
		r = &TagByProduct{}
	case "azure-native:apimanagement/v20230301preview:TagOperationLink":
		r = &TagOperationLink{}
	case "azure-native:apimanagement/v20230301preview:TagProductLink":
		r = &TagProductLink{}
	case "azure-native:apimanagement/v20230301preview:User":
		r = &User{}
	case "azure-native:apimanagement/v20230301preview:Workspace":
		r = &Workspace{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApi":
		r = &WorkspaceApi{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApiOperation":
		r = &WorkspaceApiOperation{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApiOperationPolicy":
		r = &WorkspaceApiOperationPolicy{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApiPolicy":
		r = &WorkspaceApiPolicy{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApiRelease":
		r = &WorkspaceApiRelease{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApiSchema":
		r = &WorkspaceApiSchema{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceApiVersionSet":
		r = &WorkspaceApiVersionSet{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceGlobalSchema":
		r = &WorkspaceGlobalSchema{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceGroup":
		r = &WorkspaceGroup{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceGroupUser":
		r = &WorkspaceGroupUser{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceNamedValue":
		r = &WorkspaceNamedValue{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceNotificationRecipientEmail":
		r = &WorkspaceNotificationRecipientEmail{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceNotificationRecipientUser":
		r = &WorkspaceNotificationRecipientUser{}
	case "azure-native:apimanagement/v20230301preview:WorkspacePolicy":
		r = &WorkspacePolicy{}
	case "azure-native:apimanagement/v20230301preview:WorkspacePolicyFragment":
		r = &WorkspacePolicyFragment{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceProduct":
		r = &WorkspaceProduct{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceProductApiLink":
		r = &WorkspaceProductApiLink{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceProductGroupLink":
		r = &WorkspaceProductGroupLink{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceProductPolicy":
		r = &WorkspaceProductPolicy{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceSubscription":
		r = &WorkspaceSubscription{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceTag":
		r = &WorkspaceTag{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceTagApiLink":
		r = &WorkspaceTagApiLink{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceTagOperationLink":
		r = &WorkspaceTagOperationLink{}
	case "azure-native:apimanagement/v20230301preview:WorkspaceTagProductLink":
		r = &WorkspaceTagProductLink{}
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
		"apimanagement/v20230301preview",
		&module{version},
	)
}
