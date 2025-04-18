// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a SAP Migration discovery site resource.
//
// Uses Azure REST API version 2023-10-01-preview.
func LookupSapDiscoverySite(ctx *pulumi.Context, args *LookupSapDiscoverySiteArgs, opts ...pulumi.InvokeOption) (*LookupSapDiscoverySiteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSapDiscoverySiteResult
	err := ctx.Invoke("azure-native:workloads:getSapDiscoverySite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapDiscoverySiteArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the discovery site resource for SAP Migration.
	SapDiscoverySiteName string `pulumi:"sapDiscoverySiteName"`
}

// Define the SAP Migration discovery site resource.
type LookupSapDiscoverySiteResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Indicates any errors on the SAP Migration discovery site resource.
	Errors SAPMigrateErrorResponse `pulumi:"errors"`
	// The extended location definition.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The master site ID from Azure Migrate.
	MasterSiteId *string `pulumi:"masterSiteId"`
	// The migrate project ID from Azure Migrate.
	MigrateProjectId *string `pulumi:"migrateProjectId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSapDiscoverySiteOutput(ctx *pulumi.Context, args LookupSapDiscoverySiteOutputArgs, opts ...pulumi.InvokeOption) LookupSapDiscoverySiteResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSapDiscoverySiteResultOutput, error) {
			args := v.(LookupSapDiscoverySiteArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:workloads:getSapDiscoverySite", args, LookupSapDiscoverySiteResultOutput{}, options).(LookupSapDiscoverySiteResultOutput), nil
		}).(LookupSapDiscoverySiteResultOutput)
}

type LookupSapDiscoverySiteOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the discovery site resource for SAP Migration.
	SapDiscoverySiteName pulumi.StringInput `pulumi:"sapDiscoverySiteName"`
}

func (LookupSapDiscoverySiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapDiscoverySiteArgs)(nil)).Elem()
}

// Define the SAP Migration discovery site resource.
type LookupSapDiscoverySiteResultOutput struct{ *pulumi.OutputState }

func (LookupSapDiscoverySiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapDiscoverySiteResult)(nil)).Elem()
}

func (o LookupSapDiscoverySiteResultOutput) ToLookupSapDiscoverySiteResultOutput() LookupSapDiscoverySiteResultOutput {
	return o
}

func (o LookupSapDiscoverySiteResultOutput) ToLookupSapDiscoverySiteResultOutputWithContext(ctx context.Context) LookupSapDiscoverySiteResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSapDiscoverySiteResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Indicates any errors on the SAP Migration discovery site resource.
func (o LookupSapDiscoverySiteResultOutput) Errors() SAPMigrateErrorResponseOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) SAPMigrateErrorResponse { return v.Errors }).(SAPMigrateErrorResponseOutput)
}

// The extended location definition.
func (o LookupSapDiscoverySiteResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSapDiscoverySiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSapDiscoverySiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) string { return v.Location }).(pulumi.StringOutput)
}

// The master site ID from Azure Migrate.
func (o LookupSapDiscoverySiteResultOutput) MasterSiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) *string { return v.MasterSiteId }).(pulumi.StringPtrOutput)
}

// The migrate project ID from Azure Migrate.
func (o LookupSapDiscoverySiteResultOutput) MigrateProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) *string { return v.MigrateProjectId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupSapDiscoverySiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o LookupSapDiscoverySiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSapDiscoverySiteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSapDiscoverySiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSapDiscoverySiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapDiscoverySiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSapDiscoverySiteResultOutput{})
}
