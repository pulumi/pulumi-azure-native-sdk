// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a RdsDBSnapshot
//
// Uses Azure REST API version 2024-12-01.
func LookupRdsDbSnapshot(ctx *pulumi.Context, args *LookupRdsDbSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupRdsDbSnapshotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRdsDbSnapshotResult
	err := ctx.Invoke("azure-native:awsconnector:getRdsDbSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRdsDbSnapshotArgs struct {
	// Name of RdsDBSnapshot
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupRdsDbSnapshotResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties RdsDBSnapshotPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupRdsDbSnapshotOutput(ctx *pulumi.Context, args LookupRdsDbSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupRdsDbSnapshotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRdsDbSnapshotResultOutput, error) {
			args := v.(LookupRdsDbSnapshotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getRdsDbSnapshot", args, LookupRdsDbSnapshotResultOutput{}, options).(LookupRdsDbSnapshotResultOutput), nil
		}).(LookupRdsDbSnapshotResultOutput)
}

type LookupRdsDbSnapshotOutputArgs struct {
	// Name of RdsDBSnapshot
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRdsDbSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdsDbSnapshotArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupRdsDbSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupRdsDbSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdsDbSnapshotResult)(nil)).Elem()
}

func (o LookupRdsDbSnapshotResultOutput) ToLookupRdsDbSnapshotResultOutput() LookupRdsDbSnapshotResultOutput {
	return o
}

func (o LookupRdsDbSnapshotResultOutput) ToLookupRdsDbSnapshotResultOutputWithContext(ctx context.Context) LookupRdsDbSnapshotResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupRdsDbSnapshotResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupRdsDbSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupRdsDbSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRdsDbSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupRdsDbSnapshotResultOutput) Properties() RdsDBSnapshotPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) RdsDBSnapshotPropertiesResponse { return v.Properties }).(RdsDBSnapshotPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRdsDbSnapshotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupRdsDbSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRdsDbSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRdsDbSnapshotResultOutput{})
}
