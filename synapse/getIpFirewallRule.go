// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a firewall rule
//
// Uses Azure REST API version 2021-06-01.
//
// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupIpFirewallRule(ctx *pulumi.Context, args *LookupIpFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupIpFirewallRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIpFirewallRuleResult
	err := ctx.Invoke("azure-native:synapse:getIpFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpFirewallRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The IP firewall rule name
	RuleName string `pulumi:"ruleName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// IP firewall rule
type LookupIpFirewallRuleResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Resource provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// The start IP address of the firewall rule. Must be IPv4 format
	StartIpAddress *string `pulumi:"startIpAddress"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIpFirewallRuleOutput(ctx *pulumi.Context, args LookupIpFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupIpFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIpFirewallRuleResultOutput, error) {
			args := v.(LookupIpFirewallRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:synapse:getIpFirewallRule", args, LookupIpFirewallRuleResultOutput{}, options).(LookupIpFirewallRuleResultOutput), nil
		}).(LookupIpFirewallRuleResultOutput)
}

type LookupIpFirewallRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The IP firewall rule name
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIpFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpFirewallRuleArgs)(nil)).Elem()
}

// IP firewall rule
type LookupIpFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupIpFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpFirewallRuleResult)(nil)).Elem()
}

func (o LookupIpFirewallRuleResultOutput) ToLookupIpFirewallRuleResultOutput() LookupIpFirewallRuleResultOutput {
	return o
}

func (o LookupIpFirewallRuleResultOutput) ToLookupIpFirewallRuleResultOutputWithContext(ctx context.Context) LookupIpFirewallRuleResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupIpFirewallRuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
func (o LookupIpFirewallRuleResultOutput) EndIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) *string { return v.EndIpAddress }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIpFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIpFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource provisioning state
func (o LookupIpFirewallRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The start IP address of the firewall rule. Must be IPv4 format
func (o LookupIpFirewallRuleResultOutput) StartIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) *string { return v.StartIpAddress }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIpFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpFirewallRuleResultOutput{})
}
