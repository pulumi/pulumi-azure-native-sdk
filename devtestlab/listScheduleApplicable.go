// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists all applicable schedules
//
// Uses Azure REST API version 2018-09-15.
func ListScheduleApplicable(ctx *pulumi.Context, args *ListScheduleApplicableArgs, opts ...pulumi.InvokeOption) (*ListScheduleApplicableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListScheduleApplicableResult
	err := ctx.Invoke("azure-native:devtestlab:listScheduleApplicable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListScheduleApplicableArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the schedule.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The response of a list operation.
type ListScheduleApplicableResult struct {
	// Link for next set of results.
	NextLink *string `pulumi:"nextLink"`
	// Results of the list operation.
	Value []ScheduleResponse `pulumi:"value"`
}

func ListScheduleApplicableOutput(ctx *pulumi.Context, args ListScheduleApplicableOutputArgs, opts ...pulumi.InvokeOption) ListScheduleApplicableResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListScheduleApplicableResultOutput, error) {
			args := v.(ListScheduleApplicableArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devtestlab:listScheduleApplicable", args, ListScheduleApplicableResultOutput{}, options).(ListScheduleApplicableResultOutput), nil
		}).(ListScheduleApplicableResultOutput)
}

type ListScheduleApplicableOutputArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the schedule.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListScheduleApplicableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListScheduleApplicableArgs)(nil)).Elem()
}

// The response of a list operation.
type ListScheduleApplicableResultOutput struct{ *pulumi.OutputState }

func (ListScheduleApplicableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListScheduleApplicableResult)(nil)).Elem()
}

func (o ListScheduleApplicableResultOutput) ToListScheduleApplicableResultOutput() ListScheduleApplicableResultOutput {
	return o
}

func (o ListScheduleApplicableResultOutput) ToListScheduleApplicableResultOutputWithContext(ctx context.Context) ListScheduleApplicableResultOutput {
	return o
}

// Link for next set of results.
func (o ListScheduleApplicableResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListScheduleApplicableResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Results of the list operation.
func (o ListScheduleApplicableResultOutput) Value() ScheduleResponseArrayOutput {
	return o.ApplyT(func(v ListScheduleApplicableResult) []ScheduleResponse { return v.Value }).(ScheduleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListScheduleApplicableResultOutput{})
}
