// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Secrets related to a Machine Learning compute. Might differ for every type of compute.
func ListMachineLearningComputeKeys(ctx *pulumi.Context, args *ListMachineLearningComputeKeysArgs, opts ...pulumi.InvokeOption) (*ListMachineLearningComputeKeysResult, error) {
	var rv ListMachineLearningComputeKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200401:listMachineLearningComputeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMachineLearningComputeKeysArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName string `pulumi:"computeName"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Secrets related to a Machine Learning compute. Might differ for every type of compute.
type ListMachineLearningComputeKeysResult struct {
	// The type of compute
	ComputeType string `pulumi:"computeType"`
}

func ListMachineLearningComputeKeysOutput(ctx *pulumi.Context, args ListMachineLearningComputeKeysOutputArgs, opts ...pulumi.InvokeOption) ListMachineLearningComputeKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMachineLearningComputeKeysResult, error) {
			args := v.(ListMachineLearningComputeKeysArgs)
			r, err := ListMachineLearningComputeKeys(ctx, &args, opts...)
			var s ListMachineLearningComputeKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMachineLearningComputeKeysResultOutput)
}

type ListMachineLearningComputeKeysOutputArgs struct {
	// Name of the Azure Machine Learning compute.
	ComputeName pulumi.StringInput `pulumi:"computeName"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListMachineLearningComputeKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMachineLearningComputeKeysArgs)(nil)).Elem()
}

// Secrets related to a Machine Learning compute. Might differ for every type of compute.
type ListMachineLearningComputeKeysResultOutput struct{ *pulumi.OutputState }

func (ListMachineLearningComputeKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMachineLearningComputeKeysResult)(nil)).Elem()
}

func (o ListMachineLearningComputeKeysResultOutput) ToListMachineLearningComputeKeysResultOutput() ListMachineLearningComputeKeysResultOutput {
	return o
}

func (o ListMachineLearningComputeKeysResultOutput) ToListMachineLearningComputeKeysResultOutputWithContext(ctx context.Context) ListMachineLearningComputeKeysResultOutput {
	return o
}

// The type of compute
func (o ListMachineLearningComputeKeysResultOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ListMachineLearningComputeKeysResult) string { return v.ComputeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMachineLearningComputeKeysResultOutput{})
}
