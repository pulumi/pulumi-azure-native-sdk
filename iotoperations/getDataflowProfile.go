// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataflowProfileResource
//
// Uses Azure REST API version 2024-11-01.
//
// Other available API versions: 2024-08-15-preview, 2024-09-15-preview, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotoperations [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupDataflowProfile(ctx *pulumi.Context, args *LookupDataflowProfileArgs, opts ...pulumi.InvokeOption) (*LookupDataflowProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataflowProfileResult
	err := ctx.Invoke("azure-native:iotoperations:getDataflowProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDataflowProfileArgs struct {
	// Name of Instance dataflowProfile resource
	DataflowProfileName string `pulumi:"dataflowProfileName"`
	// Name of instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Instance dataflowProfile resource
type LookupDataflowProfileResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties DataflowProfilePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDataflowProfileResult
func (val *LookupDataflowProfileResult) Defaults() *LookupDataflowProfileResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupDataflowProfileOutput(ctx *pulumi.Context, args LookupDataflowProfileOutputArgs, opts ...pulumi.InvokeOption) LookupDataflowProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataflowProfileResultOutput, error) {
			args := v.(LookupDataflowProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:iotoperations:getDataflowProfile", args, LookupDataflowProfileResultOutput{}, options).(LookupDataflowProfileResultOutput), nil
		}).(LookupDataflowProfileResultOutput)
}

type LookupDataflowProfileOutputArgs struct {
	// Name of Instance dataflowProfile resource
	DataflowProfileName pulumi.StringInput `pulumi:"dataflowProfileName"`
	// Name of instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataflowProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataflowProfileArgs)(nil)).Elem()
}

// Instance dataflowProfile resource
type LookupDataflowProfileResultOutput struct{ *pulumi.OutputState }

func (LookupDataflowProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataflowProfileResult)(nil)).Elem()
}

func (o LookupDataflowProfileResultOutput) ToLookupDataflowProfileResultOutput() LookupDataflowProfileResultOutput {
	return o
}

func (o LookupDataflowProfileResultOutput) ToLookupDataflowProfileResultOutputWithContext(ctx context.Context) LookupDataflowProfileResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDataflowProfileResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Edge location of the resource.
func (o LookupDataflowProfileResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDataflowProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataflowProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupDataflowProfileResultOutput) Properties() DataflowProfilePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) DataflowProfilePropertiesResponse { return v.Properties }).(DataflowProfilePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataflowProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataflowProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataflowProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataflowProfileResultOutput{})
}
