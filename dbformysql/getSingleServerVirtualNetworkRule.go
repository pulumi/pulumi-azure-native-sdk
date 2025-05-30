// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbformysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a virtual network rule.
//
// Uses Azure REST API version 2017-12-01.
func LookupSingleServerVirtualNetworkRule(ctx *pulumi.Context, args *LookupSingleServerVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupSingleServerVirtualNetworkRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSingleServerVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:dbformysql:getSingleServerVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSingleServerVirtualNetworkRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}

// A virtual network rule.
type LookupSingleServerVirtualNetworkRuleResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Virtual Network Rule State
	State string `pulumi:"state"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId string `pulumi:"virtualNetworkSubnetId"`
}

func LookupSingleServerVirtualNetworkRuleOutput(ctx *pulumi.Context, args LookupSingleServerVirtualNetworkRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSingleServerVirtualNetworkRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSingleServerVirtualNetworkRuleResultOutput, error) {
			args := v.(LookupSingleServerVirtualNetworkRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:dbformysql:getSingleServerVirtualNetworkRule", args, LookupSingleServerVirtualNetworkRuleResultOutput{}, options).(LookupSingleServerVirtualNetworkRuleResultOutput), nil
		}).(LookupSingleServerVirtualNetworkRuleResultOutput)
}

type LookupSingleServerVirtualNetworkRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName pulumi.StringInput `pulumi:"virtualNetworkRuleName"`
}

func (LookupSingleServerVirtualNetworkRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSingleServerVirtualNetworkRuleArgs)(nil)).Elem()
}

// A virtual network rule.
type LookupSingleServerVirtualNetworkRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSingleServerVirtualNetworkRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSingleServerVirtualNetworkRuleResult)(nil)).Elem()
}

func (o LookupSingleServerVirtualNetworkRuleResultOutput) ToLookupSingleServerVirtualNetworkRuleResultOutput() LookupSingleServerVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupSingleServerVirtualNetworkRuleResultOutput) ToLookupSingleServerVirtualNetworkRuleResultOutputWithContext(ctx context.Context) LookupSingleServerVirtualNetworkRuleResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSingleServerVirtualNetworkRuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSingleServerVirtualNetworkRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Create firewall rule before the virtual network has vnet service endpoint enabled.
func (o LookupSingleServerVirtualNetworkRuleResultOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o LookupSingleServerVirtualNetworkRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Virtual Network Rule State
func (o LookupSingleServerVirtualNetworkRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSingleServerVirtualNetworkRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

// The ARM resource id of the virtual network subnet.
func (o LookupSingleServerVirtualNetworkRuleResultOutput) VirtualNetworkSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSingleServerVirtualNetworkRuleResult) string { return v.VirtualNetworkSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSingleServerVirtualNetworkRuleResultOutput{})
}
