// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databasefleetmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets fleetspace resource.
//
// Uses Azure REST API version 2025-02-01-preview.
func LookupFleetspace(ctx *pulumi.Context, args *LookupFleetspaceArgs, opts ...pulumi.InvokeOption) (*LookupFleetspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFleetspaceResult
	err := ctx.Invoke("azure-native:databasefleetmanager:getFleetspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetspaceArgs struct {
	// Name of the database fleet.
	FleetName string `pulumi:"fleetName"`
	// Name of the fleetspace.
	FleetspaceName string `pulumi:"fleetspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A fleetspace.
type LookupFleetspaceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// A Fleetspace properties.
	Properties FleetspacePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFleetspaceOutput(ctx *pulumi.Context, args LookupFleetspaceOutputArgs, opts ...pulumi.InvokeOption) LookupFleetspaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFleetspaceResultOutput, error) {
			args := v.(LookupFleetspaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:databasefleetmanager:getFleetspace", args, LookupFleetspaceResultOutput{}, options).(LookupFleetspaceResultOutput), nil
		}).(LookupFleetspaceResultOutput)
}

type LookupFleetspaceOutputArgs struct {
	// Name of the database fleet.
	FleetName pulumi.StringInput `pulumi:"fleetName"`
	// Name of the fleetspace.
	FleetspaceName pulumi.StringInput `pulumi:"fleetspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFleetspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetspaceArgs)(nil)).Elem()
}

// A fleetspace.
type LookupFleetspaceResultOutput struct{ *pulumi.OutputState }

func (LookupFleetspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetspaceResult)(nil)).Elem()
}

func (o LookupFleetspaceResultOutput) ToLookupFleetspaceResultOutput() LookupFleetspaceResultOutput {
	return o
}

func (o LookupFleetspaceResultOutput) ToLookupFleetspaceResultOutputWithContext(ctx context.Context) LookupFleetspaceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupFleetspaceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupFleetspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFleetspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// A Fleetspace properties.
func (o LookupFleetspaceResultOutput) Properties() FleetspacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) FleetspacePropertiesResponse { return v.Properties }).(FleetspacePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFleetspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFleetspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetspaceResultOutput{})
}
