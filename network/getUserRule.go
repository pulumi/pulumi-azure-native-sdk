// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network base rule.
// API Version: 2021-02-01-preview.
//
// Deprecated: Please use one of the variants: DefaultUserRule, UserRule.
func LookupUserRule(ctx *pulumi.Context, args *LookupUserRuleArgs, opts ...pulumi.InvokeOption) (*LookupUserRuleResult, error) {
	var rv LookupUserRuleResult
	err := ctx.Invoke("azure-native:network:getUserRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserRuleArgs struct {
	// The name of the network manager security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// Network base rule.
type LookupUserRuleResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Whether the rule is custom or default.
	Kind string `pulumi:"kind"`
	// Resource name.
	Name string `pulumi:"name"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupUserRuleOutput(ctx *pulumi.Context, args LookupUserRuleOutputArgs, opts ...pulumi.InvokeOption) LookupUserRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserRuleResult, error) {
			args := v.(LookupUserRuleArgs)
			r, err := LookupUserRule(ctx, &args, opts...)
			var s LookupUserRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserRuleResultOutput)
}

type LookupUserRuleOutputArgs struct {
	// The name of the network manager security Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupUserRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserRuleArgs)(nil)).Elem()
}

// Network base rule.
type LookupUserRuleResultOutput struct{ *pulumi.OutputState }

func (LookupUserRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserRuleResult)(nil)).Elem()
}

func (o LookupUserRuleResultOutput) ToLookupUserRuleResultOutput() LookupUserRuleResultOutput {
	return o
}

func (o LookupUserRuleResultOutput) ToLookupUserRuleResultOutputWithContext(ctx context.Context) LookupUserRuleResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupUserRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupUserRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the rule is custom or default.
func (o LookupUserRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupUserRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupUserRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupUserRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserRuleResultOutput{})
}
