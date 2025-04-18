// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A paginated list of String entities.
//
// Uses Azure REST API version 2025-01-01-preview.
//
// Other available API versions: 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListInferenceGroupDeltaModelsAsync(ctx *pulumi.Context, args *ListInferenceGroupDeltaModelsAsyncArgs, opts ...pulumi.InvokeOption) (*ListInferenceGroupDeltaModelsAsyncResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListInferenceGroupDeltaModelsAsyncResult
	err := ctx.Invoke("azure-native:machinelearningservices:listInferenceGroupDeltaModelsAsync", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListInferenceGroupDeltaModelsAsyncArgs struct {
	// Gets or sets number of delta models to return. Default: -1, means that all will be returned.
	Count *int `pulumi:"count"`
	// InferenceGroup name.
	GroupName string `pulumi:"groupName"`
	// InferencePool name.
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets skip token for paginated response.
	SkipToken *string `pulumi:"skipToken"`
	// Gets or sets target base model.
	TargetBaseModel *string `pulumi:"targetBaseModel"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Defaults sets the appropriate defaults for ListInferenceGroupDeltaModelsAsyncArgs
func (val *ListInferenceGroupDeltaModelsAsyncArgs) Defaults() *ListInferenceGroupDeltaModelsAsyncArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Count == nil {
		count_ := -1
		tmp.Count = &count_
	}
	return &tmp
}

// A paginated list of String entities.
type ListInferenceGroupDeltaModelsAsyncResult struct {
	// The link to the next page of String objects. If null, there are no additional pages.
	NextLink *string `pulumi:"nextLink"`
	// An array of objects of type String.
	Value []string `pulumi:"value"`
}

func ListInferenceGroupDeltaModelsAsyncOutput(ctx *pulumi.Context, args ListInferenceGroupDeltaModelsAsyncOutputArgs, opts ...pulumi.InvokeOption) ListInferenceGroupDeltaModelsAsyncResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListInferenceGroupDeltaModelsAsyncResultOutput, error) {
			args := v.(ListInferenceGroupDeltaModelsAsyncArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices:listInferenceGroupDeltaModelsAsync", args.Defaults(), ListInferenceGroupDeltaModelsAsyncResultOutput{}, options).(ListInferenceGroupDeltaModelsAsyncResultOutput), nil
		}).(ListInferenceGroupDeltaModelsAsyncResultOutput)
}

type ListInferenceGroupDeltaModelsAsyncOutputArgs struct {
	// Gets or sets number of delta models to return. Default: -1, means that all will be returned.
	Count pulumi.IntPtrInput `pulumi:"count"`
	// InferenceGroup name.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// InferencePool name.
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Gets or sets skip token for paginated response.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
	// Gets or sets target base model.
	TargetBaseModel pulumi.StringPtrInput `pulumi:"targetBaseModel"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListInferenceGroupDeltaModelsAsyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListInferenceGroupDeltaModelsAsyncArgs)(nil)).Elem()
}

// A paginated list of String entities.
type ListInferenceGroupDeltaModelsAsyncResultOutput struct{ *pulumi.OutputState }

func (ListInferenceGroupDeltaModelsAsyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListInferenceGroupDeltaModelsAsyncResult)(nil)).Elem()
}

func (o ListInferenceGroupDeltaModelsAsyncResultOutput) ToListInferenceGroupDeltaModelsAsyncResultOutput() ListInferenceGroupDeltaModelsAsyncResultOutput {
	return o
}

func (o ListInferenceGroupDeltaModelsAsyncResultOutput) ToListInferenceGroupDeltaModelsAsyncResultOutputWithContext(ctx context.Context) ListInferenceGroupDeltaModelsAsyncResultOutput {
	return o
}

// The link to the next page of String objects. If null, there are no additional pages.
func (o ListInferenceGroupDeltaModelsAsyncResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListInferenceGroupDeltaModelsAsyncResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// An array of objects of type String.
func (o ListInferenceGroupDeltaModelsAsyncResultOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListInferenceGroupDeltaModelsAsyncResult) []string { return v.Value }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListInferenceGroupDeltaModelsAsyncResultOutput{})
}
