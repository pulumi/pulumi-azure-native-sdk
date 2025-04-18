// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a markup rule by its rule name.
//
// Uses Azure REST API version 2022-10-05-preview.
func LookupMarkupRule(ctx *pulumi.Context, args *LookupMarkupRuleArgs, opts ...pulumi.InvokeOption) (*LookupMarkupRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMarkupRuleResult
	err := ctx.Invoke("azure-native:costmanagement:getMarkupRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMarkupRuleArgs struct {
	// BillingAccount ID
	BillingAccountId string `pulumi:"billingAccountId"`
	// BillingProfile ID
	BillingProfileId string `pulumi:"billingProfileId"`
	// Markup rule name.
	Name string `pulumi:"name"`
}

// Markup rule
type LookupMarkupRuleResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Customer information for the markup rule.
	CustomerDetails CustomerMetadataResponse `pulumi:"customerDetails"`
	// The description of the markup rule.
	Description *string `pulumi:"description"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// Ending date of the markup rule.
	EndDate *string `pulumi:"endDate"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The markup percentage of the rule.
	Percentage float64 `pulumi:"percentage"`
	// Starting date of the markup rule.
	StartDate string `pulumi:"startDate"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupMarkupRuleOutput(ctx *pulumi.Context, args LookupMarkupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMarkupRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMarkupRuleResultOutput, error) {
			args := v.(LookupMarkupRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:costmanagement:getMarkupRule", args, LookupMarkupRuleResultOutput{}, options).(LookupMarkupRuleResultOutput), nil
		}).(LookupMarkupRuleResultOutput)
}

type LookupMarkupRuleOutputArgs struct {
	// BillingAccount ID
	BillingAccountId pulumi.StringInput `pulumi:"billingAccountId"`
	// BillingProfile ID
	BillingProfileId pulumi.StringInput `pulumi:"billingProfileId"`
	// Markup rule name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupMarkupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMarkupRuleArgs)(nil)).Elem()
}

// Markup rule
type LookupMarkupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMarkupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMarkupRuleResult)(nil)).Elem()
}

func (o LookupMarkupRuleResultOutput) ToLookupMarkupRuleResultOutput() LookupMarkupRuleResultOutput {
	return o
}

func (o LookupMarkupRuleResultOutput) ToLookupMarkupRuleResultOutputWithContext(ctx context.Context) LookupMarkupRuleResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMarkupRuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Customer information for the markup rule.
func (o LookupMarkupRuleResultOutput) CustomerDetails() CustomerMetadataResponseOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) CustomerMetadataResponse { return v.CustomerDetails }).(CustomerMetadataResponseOutput)
}

// The description of the markup rule.
func (o LookupMarkupRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupMarkupRuleResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Ending date of the markup rule.
func (o LookupMarkupRuleResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupMarkupRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupMarkupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The markup percentage of the rule.
func (o LookupMarkupRuleResultOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v LookupMarkupRuleResult) float64 { return v.Percentage }).(pulumi.Float64Output)
}

// Starting date of the markup rule.
func (o LookupMarkupRuleResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.StartDate }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupMarkupRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMarkupRuleResultOutput{})
}
