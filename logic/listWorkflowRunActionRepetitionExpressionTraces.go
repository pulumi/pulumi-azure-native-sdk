// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists a workflow run expression trace.
//
// Uses Azure REST API version 2019-05-01.
//
// Other available API versions: 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListWorkflowRunActionRepetitionExpressionTraces(ctx *pulumi.Context, args *ListWorkflowRunActionRepetitionExpressionTracesArgs, opts ...pulumi.InvokeOption) (*ListWorkflowRunActionRepetitionExpressionTracesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkflowRunActionRepetitionExpressionTracesResult
	err := ctx.Invoke("azure-native:logic:listWorkflowRunActionRepetitionExpressionTraces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowRunActionRepetitionExpressionTracesArgs struct {
	// The workflow action name.
	ActionName string `pulumi:"actionName"`
	// The workflow repetition.
	RepetitionName string `pulumi:"repetitionName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow run name.
	RunName string `pulumi:"runName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The expression traces.
type ListWorkflowRunActionRepetitionExpressionTracesResult struct {
	Inputs []ExpressionRootResponse `pulumi:"inputs"`
}

func ListWorkflowRunActionRepetitionExpressionTracesOutput(ctx *pulumi.Context, args ListWorkflowRunActionRepetitionExpressionTracesOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWorkflowRunActionRepetitionExpressionTracesResultOutput, error) {
			args := v.(ListWorkflowRunActionRepetitionExpressionTracesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:logic:listWorkflowRunActionRepetitionExpressionTraces", args, ListWorkflowRunActionRepetitionExpressionTracesResultOutput{}, options).(ListWorkflowRunActionRepetitionExpressionTracesResultOutput), nil
		}).(ListWorkflowRunActionRepetitionExpressionTracesResultOutput)
}

type ListWorkflowRunActionRepetitionExpressionTracesOutputArgs struct {
	// The workflow action name.
	ActionName pulumi.StringInput `pulumi:"actionName"`
	// The workflow repetition.
	RepetitionName pulumi.StringInput `pulumi:"repetitionName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workflow run name.
	RunName pulumi.StringInput `pulumi:"runName"`
	// The workflow name.
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowRunActionRepetitionExpressionTracesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionRepetitionExpressionTracesArgs)(nil)).Elem()
}

// The expression traces.
type ListWorkflowRunActionRepetitionExpressionTracesResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionRepetitionExpressionTracesResult)(nil)).Elem()
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToListWorkflowRunActionRepetitionExpressionTracesResultOutput() ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToListWorkflowRunActionRepetitionExpressionTracesResultOutputWithContext(ctx context.Context) ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) Inputs() ExpressionRootResponseArrayOutput {
	return o.ApplyT(func(v ListWorkflowRunActionRepetitionExpressionTracesResult) []ExpressionRootResponse {
		return v.Inputs
	}).(ExpressionRootResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowRunActionRepetitionExpressionTracesResultOutput{})
}
