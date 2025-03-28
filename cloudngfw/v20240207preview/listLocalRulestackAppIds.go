// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240207preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of AppIds for LocalRulestack ApiVersion
func ListLocalRulestackAppIds(ctx *pulumi.Context, args *ListLocalRulestackAppIdsArgs, opts ...pulumi.InvokeOption) (*ListLocalRulestackAppIdsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListLocalRulestackAppIdsResult
	err := ctx.Invoke("azure-native:cloudngfw/v20240207preview:listLocalRulestackAppIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLocalRulestackAppIdsArgs struct {
	AppIdVersion *string `pulumi:"appIdVersion"`
	AppPrefix    *string `pulumi:"appPrefix"`
	// LocalRulestack resource name
	LocalRulestackName string `pulumi:"localRulestackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Skip              *string `pulumi:"skip"`
	Top               *int    `pulumi:"top"`
}

type ListLocalRulestackAppIdsResult struct {
	// next Link
	NextLink *string `pulumi:"nextLink"`
	// List of AppIds
	Value []string `pulumi:"value"`
}

func ListLocalRulestackAppIdsOutput(ctx *pulumi.Context, args ListLocalRulestackAppIdsOutputArgs, opts ...pulumi.InvokeOption) ListLocalRulestackAppIdsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListLocalRulestackAppIdsResultOutput, error) {
			args := v.(ListLocalRulestackAppIdsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cloudngfw/v20240207preview:listLocalRulestackAppIds", args, ListLocalRulestackAppIdsResultOutput{}, options).(ListLocalRulestackAppIdsResultOutput), nil
		}).(ListLocalRulestackAppIdsResultOutput)
}

type ListLocalRulestackAppIdsOutputArgs struct {
	AppIdVersion pulumi.StringPtrInput `pulumi:"appIdVersion"`
	AppPrefix    pulumi.StringPtrInput `pulumi:"appPrefix"`
	// LocalRulestack resource name
	LocalRulestackName pulumi.StringInput `pulumi:"localRulestackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Skip              pulumi.StringPtrInput `pulumi:"skip"`
	Top               pulumi.IntPtrInput    `pulumi:"top"`
}

func (ListLocalRulestackAppIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocalRulestackAppIdsArgs)(nil)).Elem()
}

type ListLocalRulestackAppIdsResultOutput struct{ *pulumi.OutputState }

func (ListLocalRulestackAppIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocalRulestackAppIdsResult)(nil)).Elem()
}

func (o ListLocalRulestackAppIdsResultOutput) ToListLocalRulestackAppIdsResultOutput() ListLocalRulestackAppIdsResultOutput {
	return o
}

func (o ListLocalRulestackAppIdsResultOutput) ToListLocalRulestackAppIdsResultOutputWithContext(ctx context.Context) ListLocalRulestackAppIdsResultOutput {
	return o
}

// next Link
func (o ListLocalRulestackAppIdsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListLocalRulestackAppIdsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of AppIds
func (o ListLocalRulestackAppIdsResultOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListLocalRulestackAppIdsResult) []string { return v.Value }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLocalRulestackAppIdsResultOutput{})
}
