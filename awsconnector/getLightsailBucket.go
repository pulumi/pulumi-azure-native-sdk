// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a LightsailBucket
//
// Uses Azure REST API version 2024-12-01.
func LookupLightsailBucket(ctx *pulumi.Context, args *LookupLightsailBucketArgs, opts ...pulumi.InvokeOption) (*LookupLightsailBucketResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLightsailBucketResult
	err := ctx.Invoke("azure-native:awsconnector:getLightsailBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLightsailBucketArgs struct {
	// Name of LightsailBucket
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupLightsailBucketResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties LightsailBucketPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLightsailBucketOutput(ctx *pulumi.Context, args LookupLightsailBucketOutputArgs, opts ...pulumi.InvokeOption) LookupLightsailBucketResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLightsailBucketResultOutput, error) {
			args := v.(LookupLightsailBucketArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getLightsailBucket", args, LookupLightsailBucketResultOutput{}, options).(LookupLightsailBucketResultOutput), nil
		}).(LookupLightsailBucketResultOutput)
}

type LookupLightsailBucketOutputArgs struct {
	// Name of LightsailBucket
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLightsailBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLightsailBucketArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupLightsailBucketResultOutput struct{ *pulumi.OutputState }

func (LookupLightsailBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLightsailBucketResult)(nil)).Elem()
}

func (o LookupLightsailBucketResultOutput) ToLookupLightsailBucketResultOutput() LookupLightsailBucketResultOutput {
	return o
}

func (o LookupLightsailBucketResultOutput) ToLookupLightsailBucketResultOutputWithContext(ctx context.Context) LookupLightsailBucketResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupLightsailBucketResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupLightsailBucketResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupLightsailBucketResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLightsailBucketResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupLightsailBucketResultOutput) Properties() LightsailBucketPropertiesResponseOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) LightsailBucketPropertiesResponse { return v.Properties }).(LightsailBucketPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLightsailBucketResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupLightsailBucketResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLightsailBucketResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLightsailBucketResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLightsailBucketResultOutput{})
}
