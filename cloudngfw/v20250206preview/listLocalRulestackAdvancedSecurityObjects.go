// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250206preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the list of advanced security objects
func ListLocalRulestackAdvancedSecurityObjects(ctx *pulumi.Context, args *ListLocalRulestackAdvancedSecurityObjectsArgs, opts ...pulumi.InvokeOption) (*ListLocalRulestackAdvancedSecurityObjectsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListLocalRulestackAdvancedSecurityObjectsResult
	err := ctx.Invoke("azure-native:cloudngfw/v20250206preview:listLocalRulestackAdvancedSecurityObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLocalRulestackAdvancedSecurityObjectsArgs struct {
	// LocalRulestack resource name
	LocalRulestackName string `pulumi:"localRulestackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Skip              *string `pulumi:"skip"`
	Top               *int    `pulumi:"top"`
	Type              string  `pulumi:"type"`
}

// advanced security object
type ListLocalRulestackAdvancedSecurityObjectsResult struct {
	// next link
	NextLink *string `pulumi:"nextLink"`
	// response value
	Value AdvSecurityObjectModelResponse `pulumi:"value"`
}

func ListLocalRulestackAdvancedSecurityObjectsOutput(ctx *pulumi.Context, args ListLocalRulestackAdvancedSecurityObjectsOutputArgs, opts ...pulumi.InvokeOption) ListLocalRulestackAdvancedSecurityObjectsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListLocalRulestackAdvancedSecurityObjectsResultOutput, error) {
			args := v.(ListLocalRulestackAdvancedSecurityObjectsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cloudngfw/v20250206preview:listLocalRulestackAdvancedSecurityObjects", args, ListLocalRulestackAdvancedSecurityObjectsResultOutput{}, options).(ListLocalRulestackAdvancedSecurityObjectsResultOutput), nil
		}).(ListLocalRulestackAdvancedSecurityObjectsResultOutput)
}

type ListLocalRulestackAdvancedSecurityObjectsOutputArgs struct {
	// LocalRulestack resource name
	LocalRulestackName pulumi.StringInput `pulumi:"localRulestackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Skip              pulumi.StringPtrInput `pulumi:"skip"`
	Top               pulumi.IntPtrInput    `pulumi:"top"`
	Type              pulumi.StringInput    `pulumi:"type"`
}

func (ListLocalRulestackAdvancedSecurityObjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocalRulestackAdvancedSecurityObjectsArgs)(nil)).Elem()
}

// advanced security object
type ListLocalRulestackAdvancedSecurityObjectsResultOutput struct{ *pulumi.OutputState }

func (ListLocalRulestackAdvancedSecurityObjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocalRulestackAdvancedSecurityObjectsResult)(nil)).Elem()
}

func (o ListLocalRulestackAdvancedSecurityObjectsResultOutput) ToListLocalRulestackAdvancedSecurityObjectsResultOutput() ListLocalRulestackAdvancedSecurityObjectsResultOutput {
	return o
}

func (o ListLocalRulestackAdvancedSecurityObjectsResultOutput) ToListLocalRulestackAdvancedSecurityObjectsResultOutputWithContext(ctx context.Context) ListLocalRulestackAdvancedSecurityObjectsResultOutput {
	return o
}

// next link
func (o ListLocalRulestackAdvancedSecurityObjectsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListLocalRulestackAdvancedSecurityObjectsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// response value
func (o ListLocalRulestackAdvancedSecurityObjectsResultOutput) Value() AdvSecurityObjectModelResponseOutput {
	return o.ApplyT(func(v ListLocalRulestackAdvancedSecurityObjectsResult) AdvSecurityObjectModelResponse { return v.Value }).(AdvSecurityObjectModelResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLocalRulestackAdvancedSecurityObjectsResultOutput{})
}
