// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a IamMFADevice
//
// Uses Azure REST API version 2024-12-01.
func LookupIamMfaDevice(ctx *pulumi.Context, args *LookupIamMfaDeviceArgs, opts ...pulumi.InvokeOption) (*LookupIamMfaDeviceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIamMfaDeviceResult
	err := ctx.Invoke("azure-native:awsconnector:getIamMfaDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIamMfaDeviceArgs struct {
	// Name of IamMFADevice
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupIamMfaDeviceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties IamMFADevicePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIamMfaDeviceOutput(ctx *pulumi.Context, args LookupIamMfaDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupIamMfaDeviceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIamMfaDeviceResultOutput, error) {
			args := v.(LookupIamMfaDeviceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getIamMfaDevice", args, LookupIamMfaDeviceResultOutput{}, options).(LookupIamMfaDeviceResultOutput), nil
		}).(LookupIamMfaDeviceResultOutput)
}

type LookupIamMfaDeviceOutputArgs struct {
	// Name of IamMFADevice
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIamMfaDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamMfaDeviceArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupIamMfaDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupIamMfaDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamMfaDeviceResult)(nil)).Elem()
}

func (o LookupIamMfaDeviceResultOutput) ToLookupIamMfaDeviceResultOutput() LookupIamMfaDeviceResultOutput {
	return o
}

func (o LookupIamMfaDeviceResultOutput) ToLookupIamMfaDeviceResultOutputWithContext(ctx context.Context) LookupIamMfaDeviceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupIamMfaDeviceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupIamMfaDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupIamMfaDeviceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIamMfaDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupIamMfaDeviceResultOutput) Properties() IamMFADevicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) IamMFADevicePropertiesResponse { return v.Properties }).(IamMFADevicePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIamMfaDeviceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupIamMfaDeviceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIamMfaDeviceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamMfaDeviceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIamMfaDeviceResultOutput{})
}
