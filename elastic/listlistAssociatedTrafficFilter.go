// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of elastic traffic filters in the account
//
// Uses Azure REST API version 2023-06-01.
//
// Other available API versions: 2023-06-15-preview, 2023-07-01-preview, 2023-10-01-preview, 2023-11-01-preview, 2024-01-01-preview, 2024-03-01, 2024-05-01-preview, 2024-06-15-preview, 2024-10-01-preview, 2025-01-15-preview.
func ListlistAssociatedTrafficFilter(ctx *pulumi.Context, args *ListlistAssociatedTrafficFilterArgs, opts ...pulumi.InvokeOption) (*ListlistAssociatedTrafficFilterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListlistAssociatedTrafficFilterResult
	err := ctx.Invoke("azure-native:elastic:listlistAssociatedTrafficFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListlistAssociatedTrafficFilterArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of elastic traffic filters in the account
type ListlistAssociatedTrafficFilterResult struct {
	// List of elastic traffic filters in the account
	Rulesets []ElasticTrafficFilterResponse `pulumi:"rulesets"`
}

func ListlistAssociatedTrafficFilterOutput(ctx *pulumi.Context, args ListlistAssociatedTrafficFilterOutputArgs, opts ...pulumi.InvokeOption) ListlistAssociatedTrafficFilterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListlistAssociatedTrafficFilterResultOutput, error) {
			args := v.(ListlistAssociatedTrafficFilterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:elastic:listlistAssociatedTrafficFilter", args, ListlistAssociatedTrafficFilterResultOutput{}, options).(ListlistAssociatedTrafficFilterResultOutput), nil
		}).(ListlistAssociatedTrafficFilterResultOutput)
}

type ListlistAssociatedTrafficFilterOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListlistAssociatedTrafficFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListlistAssociatedTrafficFilterArgs)(nil)).Elem()
}

// List of elastic traffic filters in the account
type ListlistAssociatedTrafficFilterResultOutput struct{ *pulumi.OutputState }

func (ListlistAssociatedTrafficFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListlistAssociatedTrafficFilterResult)(nil)).Elem()
}

func (o ListlistAssociatedTrafficFilterResultOutput) ToListlistAssociatedTrafficFilterResultOutput() ListlistAssociatedTrafficFilterResultOutput {
	return o
}

func (o ListlistAssociatedTrafficFilterResultOutput) ToListlistAssociatedTrafficFilterResultOutputWithContext(ctx context.Context) ListlistAssociatedTrafficFilterResultOutput {
	return o
}

// List of elastic traffic filters in the account
func (o ListlistAssociatedTrafficFilterResultOutput) Rulesets() ElasticTrafficFilterResponseArrayOutput {
	return o.ApplyT(func(v ListlistAssociatedTrafficFilterResult) []ElasticTrafficFilterResponse { return v.Rulesets }).(ElasticTrafficFilterResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListlistAssociatedTrafficFilterResultOutput{})
}
