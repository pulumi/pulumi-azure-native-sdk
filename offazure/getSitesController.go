// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a VmwareSite
//
// Uses Azure REST API version 2023-10-01-preview.
//
// Other available API versions: 2023-06-06, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native offazure [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupSitesController(ctx *pulumi.Context, args *LookupSitesControllerArgs, opts ...pulumi.InvokeOption) (*LookupSitesControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSitesControllerResult
	err := ctx.Invoke("azure-native:offazure:getSitesController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSitesControllerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// A VmwareSite
type LookupSitesControllerResult struct {
	// Gets or sets the on-premises agent details.
	AgentDetails *SiteAgentPropertiesResponse `pulumi:"agentDetails"`
	// Gets or sets the Appliance Name.
	ApplianceName *string `pulumi:"applianceName"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId *string `pulumi:"discoverySolutionId"`
	// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	ETag string `pulumi:"eTag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Gets the Master Site this site is linked to.
	MasterSiteId string `pulumi:"masterSiteId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets the service endpoint.
	ServiceEndpoint string `pulumi:"serviceEndpoint"`
	// Gets or sets the service principal identity details used by agent for
	// communication
	//             to the service.
	ServicePrincipalIdentityDetails *SiteSpnPropertiesResponse `pulumi:"servicePrincipalIdentityDetails"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSitesControllerOutput(ctx *pulumi.Context, args LookupSitesControllerOutputArgs, opts ...pulumi.InvokeOption) LookupSitesControllerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSitesControllerResultOutput, error) {
			args := v.(LookupSitesControllerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:offazure:getSitesController", args, LookupSitesControllerResultOutput{}, options).(LookupSitesControllerResultOutput), nil
		}).(LookupSitesControllerResultOutput)
}

type LookupSitesControllerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupSitesControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSitesControllerArgs)(nil)).Elem()
}

// A VmwareSite
type LookupSitesControllerResultOutput struct{ *pulumi.OutputState }

func (LookupSitesControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSitesControllerResult)(nil)).Elem()
}

func (o LookupSitesControllerResultOutput) ToLookupSitesControllerResultOutput() LookupSitesControllerResultOutput {
	return o
}

func (o LookupSitesControllerResultOutput) ToLookupSitesControllerResultOutputWithContext(ctx context.Context) LookupSitesControllerResultOutput {
	return o
}

// Gets or sets the on-premises agent details.
func (o LookupSitesControllerResultOutput) AgentDetails() SiteAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) *SiteAgentPropertiesResponse { return v.AgentDetails }).(SiteAgentPropertiesResponsePtrOutput)
}

// Gets or sets the Appliance Name.
func (o LookupSitesControllerResultOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) *string { return v.ApplianceName }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupSitesControllerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the ARM ID of migration hub solution for SDS.
func (o LookupSitesControllerResultOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) *string { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupSitesControllerResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSitesControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSitesControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets the Master Site this site is linked to.
func (o LookupSitesControllerResultOutput) MasterSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.MasterSiteId }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSitesControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupSitesControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o LookupSitesControllerResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the service principal identity details used by agent for
// communication
//
//	to the service.
func (o LookupSitesControllerResultOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) *SiteSpnPropertiesResponse {
		return v.ServicePrincipalIdentityDetails
	}).(SiteSpnPropertiesResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSitesControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSitesControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSitesControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSitesControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSitesControllerResultOutput{})
}
