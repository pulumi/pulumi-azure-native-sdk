// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The response to an enumeration operation on Virtual Network DNS Forwarding Ruleset.
func ListDnsForwardingRulesetByVirtualNetwork(ctx *pulumi.Context, args *ListDnsForwardingRulesetByVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*ListDnsForwardingRulesetByVirtualNetworkResult, error) {
	var rv ListDnsForwardingRulesetByVirtualNetworkResult
	err := ctx.Invoke("azure-native:network/v20220701:listDnsForwardingRulesetByVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDnsForwardingRulesetByVirtualNetworkArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	Top                *int   `pulumi:"top"`
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The response to an enumeration operation on Virtual Network DNS Forwarding Ruleset.
type ListDnsForwardingRulesetByVirtualNetworkResult struct {
	NextLink string                                       `pulumi:"nextLink"`
	Value    []VirtualNetworkDnsForwardingRulesetResponse `pulumi:"value"`
}

func ListDnsForwardingRulesetByVirtualNetworkOutput(ctx *pulumi.Context, args ListDnsForwardingRulesetByVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) ListDnsForwardingRulesetByVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDnsForwardingRulesetByVirtualNetworkResult, error) {
			args := v.(ListDnsForwardingRulesetByVirtualNetworkArgs)
			r, err := ListDnsForwardingRulesetByVirtualNetwork(ctx, &args, opts...)
			var s ListDnsForwardingRulesetByVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDnsForwardingRulesetByVirtualNetworkResultOutput)
}

type ListDnsForwardingRulesetByVirtualNetworkOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	Top                pulumi.IntPtrInput `pulumi:"top"`
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (ListDnsForwardingRulesetByVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDnsForwardingRulesetByVirtualNetworkArgs)(nil)).Elem()
}

// The response to an enumeration operation on Virtual Network DNS Forwarding Ruleset.
type ListDnsForwardingRulesetByVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (ListDnsForwardingRulesetByVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDnsForwardingRulesetByVirtualNetworkResult)(nil)).Elem()
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) ToListDnsForwardingRulesetByVirtualNetworkResultOutput() ListDnsForwardingRulesetByVirtualNetworkResultOutput {
	return o
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) ToListDnsForwardingRulesetByVirtualNetworkResultOutputWithContext(ctx context.Context) ListDnsForwardingRulesetByVirtualNetworkResultOutput {
	return o
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListDnsForwardingRulesetByVirtualNetworkResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) Value() VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o.ApplyT(func(v ListDnsForwardingRulesetByVirtualNetworkResult) []VirtualNetworkDnsForwardingRulesetResponse {
		return v.Value
	}).(VirtualNetworkDnsForwardingRulesetResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDnsForwardingRulesetByVirtualNetworkResultOutput{})
}