// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataFlowResource
//
// Uses Azure REST API version 2024-07-01-preview.
func LookupDataFlow(ctx *pulumi.Context, args *LookupDataFlowArgs, opts ...pulumi.InvokeOption) (*LookupDataFlowResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataFlowResult
	err := ctx.Invoke("azure-native:iotoperations:getDataFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDataFlowArgs struct {
	// Name of Instance dataflowProfile dataflow resource
	DataflowName string `pulumi:"dataflowName"`
	// Name of Instance dataflowProfile resource
	DataflowProfileName string `pulumi:"dataflowProfileName"`
	// Name of instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Instance dataflowProfile dataflow resource
type LookupDataFlowResult struct {
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties DataFlowPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDataFlowResult
func (val *LookupDataFlowResult) Defaults() *LookupDataFlowResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupDataFlowOutput(ctx *pulumi.Context, args LookupDataFlowOutputArgs, opts ...pulumi.InvokeOption) LookupDataFlowResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataFlowResultOutput, error) {
			args := v.(LookupDataFlowArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:iotoperations:getDataFlow", args, LookupDataFlowResultOutput{}, options).(LookupDataFlowResultOutput), nil
		}).(LookupDataFlowResultOutput)
}

type LookupDataFlowOutputArgs struct {
	// Name of Instance dataflowProfile dataflow resource
	DataflowName pulumi.StringInput `pulumi:"dataflowName"`
	// Name of Instance dataflowProfile resource
	DataflowProfileName pulumi.StringInput `pulumi:"dataflowProfileName"`
	// Name of instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataFlowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataFlowArgs)(nil)).Elem()
}

// Instance dataflowProfile dataflow resource
type LookupDataFlowResultOutput struct{ *pulumi.OutputState }

func (LookupDataFlowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataFlowResult)(nil)).Elem()
}

func (o LookupDataFlowResultOutput) ToLookupDataFlowResultOutput() LookupDataFlowResultOutput {
	return o
}

func (o LookupDataFlowResultOutput) ToLookupDataFlowResultOutputWithContext(ctx context.Context) LookupDataFlowResultOutput {
	return o
}

// Edge location of the resource.
func (o LookupDataFlowResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupDataFlowResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDataFlowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataFlowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupDataFlowResultOutput) Properties() DataFlowPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDataFlowResult) DataFlowPropertiesResponse { return v.Properties }).(DataFlowPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataFlowResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataFlowResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataFlowResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataFlowResultOutput{})
}
