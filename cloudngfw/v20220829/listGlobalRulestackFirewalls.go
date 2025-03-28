// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220829

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of Firewalls associated with Rulestack
func ListGlobalRulestackFirewalls(ctx *pulumi.Context, args *ListGlobalRulestackFirewallsArgs, opts ...pulumi.InvokeOption) (*ListGlobalRulestackFirewallsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListGlobalRulestackFirewallsResult
	err := ctx.Invoke("azure-native:cloudngfw/v20220829:listGlobalRulestackFirewalls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGlobalRulestackFirewallsArgs struct {
	// GlobalRulestack resource name
	GlobalRulestackName string `pulumi:"globalRulestackName"`
}

// List firewalls response
type ListGlobalRulestackFirewallsResult struct {
	// next link
	NextLink *string `pulumi:"nextLink"`
	// firewalls list
	Value []string `pulumi:"value"`
}

func ListGlobalRulestackFirewallsOutput(ctx *pulumi.Context, args ListGlobalRulestackFirewallsOutputArgs, opts ...pulumi.InvokeOption) ListGlobalRulestackFirewallsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListGlobalRulestackFirewallsResultOutput, error) {
			args := v.(ListGlobalRulestackFirewallsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cloudngfw/v20220829:listGlobalRulestackFirewalls", args, ListGlobalRulestackFirewallsResultOutput{}, options).(ListGlobalRulestackFirewallsResultOutput), nil
		}).(ListGlobalRulestackFirewallsResultOutput)
}

type ListGlobalRulestackFirewallsOutputArgs struct {
	// GlobalRulestack resource name
	GlobalRulestackName pulumi.StringInput `pulumi:"globalRulestackName"`
}

func (ListGlobalRulestackFirewallsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGlobalRulestackFirewallsArgs)(nil)).Elem()
}

// List firewalls response
type ListGlobalRulestackFirewallsResultOutput struct{ *pulumi.OutputState }

func (ListGlobalRulestackFirewallsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGlobalRulestackFirewallsResult)(nil)).Elem()
}

func (o ListGlobalRulestackFirewallsResultOutput) ToListGlobalRulestackFirewallsResultOutput() ListGlobalRulestackFirewallsResultOutput {
	return o
}

func (o ListGlobalRulestackFirewallsResultOutput) ToListGlobalRulestackFirewallsResultOutputWithContext(ctx context.Context) ListGlobalRulestackFirewallsResultOutput {
	return o
}

// next link
func (o ListGlobalRulestackFirewallsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGlobalRulestackFirewallsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// firewalls list
func (o ListGlobalRulestackFirewallsResultOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListGlobalRulestackFirewallsResult) []string { return v.Value }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGlobalRulestackFirewallsResultOutput{})
}
