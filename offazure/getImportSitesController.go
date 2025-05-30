// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a ImportSite
//
// Uses Azure REST API version 2023-10-01-preview.
//
// Other available API versions: 2023-06-06, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native offazure [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupImportSitesController(ctx *pulumi.Context, args *LookupImportSitesControllerArgs, opts ...pulumi.InvokeOption) (*LookupImportSitesControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupImportSitesControllerResult
	err := ctx.Invoke("azure-native:offazure:getImportSitesController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImportSitesControllerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// A ImportSite
type LookupImportSitesControllerResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets or sets the ARM ID of migration hub solution for SDS.
	DiscoverySolutionId *string `pulumi:"discoverySolutionId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Gets the Master Site this site is linked to.
	MasterSiteId string `pulumi:"masterSiteId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the service endpoint.
	ServiceEndpoint string `pulumi:"serviceEndpoint"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupImportSitesControllerOutput(ctx *pulumi.Context, args LookupImportSitesControllerOutputArgs, opts ...pulumi.InvokeOption) LookupImportSitesControllerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupImportSitesControllerResultOutput, error) {
			args := v.(LookupImportSitesControllerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:offazure:getImportSitesController", args, LookupImportSitesControllerResultOutput{}, options).(LookupImportSitesControllerResultOutput), nil
		}).(LookupImportSitesControllerResultOutput)
}

type LookupImportSitesControllerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupImportSitesControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportSitesControllerArgs)(nil)).Elem()
}

// A ImportSite
type LookupImportSitesControllerResultOutput struct{ *pulumi.OutputState }

func (LookupImportSitesControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportSitesControllerResult)(nil)).Elem()
}

func (o LookupImportSitesControllerResultOutput) ToLookupImportSitesControllerResultOutput() LookupImportSitesControllerResultOutput {
	return o
}

func (o LookupImportSitesControllerResultOutput) ToLookupImportSitesControllerResultOutputWithContext(ctx context.Context) LookupImportSitesControllerResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupImportSitesControllerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the ARM ID of migration hub solution for SDS.
func (o LookupImportSitesControllerResultOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) *string { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupImportSitesControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupImportSitesControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets the Master Site this site is linked to.
func (o LookupImportSitesControllerResultOutput) MasterSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.MasterSiteId }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupImportSitesControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupImportSitesControllerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets the service endpoint.
func (o LookupImportSitesControllerResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupImportSitesControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupImportSitesControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupImportSitesControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportSitesControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportSitesControllerResultOutput{})
}
