// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an associated project catalog.
//
// Uses Azure REST API version 2024-02-01.
//
// Other available API versions: 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupProjectCatalog(ctx *pulumi.Context, args *LookupProjectCatalogArgs, opts ...pulumi.InvokeOption) (*LookupProjectCatalogResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectCatalogResult
	err := ctx.Invoke("azure-native:devcenter:getProjectCatalog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectCatalogArgs struct {
	// The name of the Catalog.
	CatalogName string `pulumi:"catalogName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a catalog.
type LookupProjectCatalogResult struct {
	// Properties for an Azure DevOps catalog type.
	AdoGit *GitCatalogResponse `pulumi:"adoGit"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The connection state of the catalog.
	ConnectionState string `pulumi:"connectionState"`
	// Properties for a GitHub catalog type.
	GitHub *GitCatalogResponse `pulumi:"gitHub"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// When the catalog was last connected.
	LastConnectionTime string `pulumi:"lastConnectionTime"`
	// Stats of the latest synchronization.
	LastSyncStats SyncStatsResponse `pulumi:"lastSyncStats"`
	// When the catalog was last synced.
	LastSyncTime string `pulumi:"lastSyncTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The synchronization state of the catalog.
	SyncState string `pulumi:"syncState"`
	// Indicates the type of sync that is configured for the catalog.
	SyncType *string `pulumi:"syncType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupProjectCatalogOutput(ctx *pulumi.Context, args LookupProjectCatalogOutputArgs, opts ...pulumi.InvokeOption) LookupProjectCatalogResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectCatalogResultOutput, error) {
			args := v.(LookupProjectCatalogArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter:getProjectCatalog", args, LookupProjectCatalogResultOutput{}, options).(LookupProjectCatalogResultOutput), nil
		}).(LookupProjectCatalogResultOutput)
}

type LookupProjectCatalogOutputArgs struct {
	// The name of the Catalog.
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// The name of the project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProjectCatalogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectCatalogArgs)(nil)).Elem()
}

// Represents a catalog.
type LookupProjectCatalogResultOutput struct{ *pulumi.OutputState }

func (LookupProjectCatalogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectCatalogResult)(nil)).Elem()
}

func (o LookupProjectCatalogResultOutput) ToLookupProjectCatalogResultOutput() LookupProjectCatalogResultOutput {
	return o
}

func (o LookupProjectCatalogResultOutput) ToLookupProjectCatalogResultOutputWithContext(ctx context.Context) LookupProjectCatalogResultOutput {
	return o
}

// Properties for an Azure DevOps catalog type.
func (o LookupProjectCatalogResultOutput) AdoGit() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) *GitCatalogResponse { return v.AdoGit }).(GitCatalogResponsePtrOutput)
}

// The Azure API version of the resource.
func (o LookupProjectCatalogResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The connection state of the catalog.
func (o LookupProjectCatalogResultOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.ConnectionState }).(pulumi.StringOutput)
}

// Properties for a GitHub catalog type.
func (o LookupProjectCatalogResultOutput) GitHub() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) *GitCatalogResponse { return v.GitHub }).(GitCatalogResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupProjectCatalogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.Id }).(pulumi.StringOutput)
}

// When the catalog was last connected.
func (o LookupProjectCatalogResultOutput) LastConnectionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.LastConnectionTime }).(pulumi.StringOutput)
}

// Stats of the latest synchronization.
func (o LookupProjectCatalogResultOutput) LastSyncStats() SyncStatsResponseOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) SyncStatsResponse { return v.LastSyncStats }).(SyncStatsResponseOutput)
}

// When the catalog was last synced.
func (o LookupProjectCatalogResultOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupProjectCatalogResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupProjectCatalogResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The synchronization state of the catalog.
func (o LookupProjectCatalogResultOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.SyncState }).(pulumi.StringOutput)
}

// Indicates the type of sync that is configured for the catalog.
func (o LookupProjectCatalogResultOutput) SyncType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) *string { return v.SyncType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupProjectCatalogResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupProjectCatalogResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupProjectCatalogResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectCatalogResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectCatalogResultOutput{})
}
