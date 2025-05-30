// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a RdsDBSnapshotAttributesResult
//
// Uses Azure REST API version 2024-12-01.
func LookupRdsDbSnapshotAttributesResult(ctx *pulumi.Context, args *LookupRdsDbSnapshotAttributesResultArgs, opts ...pulumi.InvokeOption) (*LookupRdsDbSnapshotAttributesResultResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRdsDbSnapshotAttributesResultResult
	err := ctx.Invoke("azure-native:awsconnector:getRdsDbSnapshotAttributesResult", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRdsDbSnapshotAttributesResultArgs struct {
	// Name of RdsDBSnapshotAttributesResult
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupRdsDbSnapshotAttributesResultResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties RdsDBSnapshotAttributesResultPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupRdsDbSnapshotAttributesResultOutput(ctx *pulumi.Context, args LookupRdsDbSnapshotAttributesResultOutputArgs, opts ...pulumi.InvokeOption) LookupRdsDbSnapshotAttributesResultResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRdsDbSnapshotAttributesResultResultOutput, error) {
			args := v.(LookupRdsDbSnapshotAttributesResultArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getRdsDbSnapshotAttributesResult", args, LookupRdsDbSnapshotAttributesResultResultOutput{}, options).(LookupRdsDbSnapshotAttributesResultResultOutput), nil
		}).(LookupRdsDbSnapshotAttributesResultResultOutput)
}

type LookupRdsDbSnapshotAttributesResultOutputArgs struct {
	// Name of RdsDBSnapshotAttributesResult
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRdsDbSnapshotAttributesResultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdsDbSnapshotAttributesResultArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupRdsDbSnapshotAttributesResultResultOutput struct{ *pulumi.OutputState }

func (LookupRdsDbSnapshotAttributesResultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdsDbSnapshotAttributesResultResult)(nil)).Elem()
}

func (o LookupRdsDbSnapshotAttributesResultResultOutput) ToLookupRdsDbSnapshotAttributesResultResultOutput() LookupRdsDbSnapshotAttributesResultResultOutput {
	return o
}

func (o LookupRdsDbSnapshotAttributesResultResultOutput) ToLookupRdsDbSnapshotAttributesResultResultOutputWithContext(ctx context.Context) LookupRdsDbSnapshotAttributesResultResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupRdsDbSnapshotAttributesResultResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupRdsDbSnapshotAttributesResultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupRdsDbSnapshotAttributesResultResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRdsDbSnapshotAttributesResultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupRdsDbSnapshotAttributesResultResultOutput) Properties() RdsDBSnapshotAttributesResultPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) RdsDBSnapshotAttributesResultPropertiesResponse {
		return v.Properties
	}).(RdsDBSnapshotAttributesResultPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRdsDbSnapshotAttributesResultResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupRdsDbSnapshotAttributesResultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRdsDbSnapshotAttributesResultResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdsDbSnapshotAttributesResultResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRdsDbSnapshotAttributesResultResultOutput{})
}
