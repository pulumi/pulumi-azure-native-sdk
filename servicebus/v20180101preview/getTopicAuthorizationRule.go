// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a namespace authorization rule.
func LookupTopicAuthorizationRule(ctx *pulumi.Context, args *LookupTopicAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupTopicAuthorizationRuleResult, error) {
	var rv LookupTopicAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20180101preview:getTopicAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}

// Description of a namespace authorization rule.
type LookupTopicAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}

func LookupTopicAuthorizationRuleOutput(ctx *pulumi.Context, args LookupTopicAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupTopicAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicAuthorizationRuleResult, error) {
			args := v.(LookupTopicAuthorizationRuleArgs)
			r, err := LookupTopicAuthorizationRule(ctx, &args, opts...)
			var s LookupTopicAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicAuthorizationRuleResultOutput)
}

type LookupTopicAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicAuthorizationRuleArgs)(nil)).Elem()
}

// Description of a namespace authorization rule.
type LookupTopicAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupTopicAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupTopicAuthorizationRuleResultOutput) ToLookupTopicAuthorizationRuleResultOutput() LookupTopicAuthorizationRuleResultOutput {
	return o
}

func (o LookupTopicAuthorizationRuleResultOutput) ToLookupTopicAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupTopicAuthorizationRuleResultOutput {
	return o
}

func (o LookupTopicAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTopicAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTopicAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTopicAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupTopicAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicAuthorizationRuleResultOutput{})
}