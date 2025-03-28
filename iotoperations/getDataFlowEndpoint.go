// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataFlowEndpointResource
//
// Uses Azure REST API version 2024-07-01-preview.
func LookupDataFlowEndpoint(ctx *pulumi.Context, args *LookupDataFlowEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDataFlowEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataFlowEndpointResult
	err := ctx.Invoke("azure-native:iotoperations:getDataFlowEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDataFlowEndpointArgs struct {
	// Name of Instance dataflowEndpoint resource
	DataflowEndpointName string `pulumi:"dataflowEndpointName"`
	// Name of instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Instance dataflowEndpoint resource
type LookupDataFlowEndpointResult struct {
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties DataFlowEndpointPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDataFlowEndpointResult
func (val *LookupDataFlowEndpointResult) Defaults() *LookupDataFlowEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupDataFlowEndpointOutput(ctx *pulumi.Context, args LookupDataFlowEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDataFlowEndpointResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataFlowEndpointResultOutput, error) {
			args := v.(LookupDataFlowEndpointArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:iotoperations:getDataFlowEndpoint", args, LookupDataFlowEndpointResultOutput{}, options).(LookupDataFlowEndpointResultOutput), nil
		}).(LookupDataFlowEndpointResultOutput)
}

type LookupDataFlowEndpointOutputArgs struct {
	// Name of Instance dataflowEndpoint resource
	DataflowEndpointName pulumi.StringInput `pulumi:"dataflowEndpointName"`
	// Name of instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataFlowEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataFlowEndpointArgs)(nil)).Elem()
}

// Instance dataflowEndpoint resource
type LookupDataFlowEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDataFlowEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataFlowEndpointResult)(nil)).Elem()
}

func (o LookupDataFlowEndpointResultOutput) ToLookupDataFlowEndpointResultOutput() LookupDataFlowEndpointResultOutput {
	return o
}

func (o LookupDataFlowEndpointResultOutput) ToLookupDataFlowEndpointResultOutputWithContext(ctx context.Context) LookupDataFlowEndpointResultOutput {
	return o
}

// Edge location of the resource.
func (o LookupDataFlowEndpointResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupDataFlowEndpointResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDataFlowEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataFlowEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupDataFlowEndpointResultOutput) Properties() DataFlowEndpointPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDataFlowEndpointResult) DataFlowEndpointPropertiesResponse { return v.Properties }).(DataFlowEndpointPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataFlowEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataFlowEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataFlowEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataFlowEndpointResultOutput{})
}
