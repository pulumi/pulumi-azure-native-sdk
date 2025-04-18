// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the primary and secondary access keys for the specified Power BI Workspace Collection.
//
// Uses Azure REST API version 2016-01-29.
func ListWorkspaceCollectionAccessKeys(ctx *pulumi.Context, args *ListWorkspaceCollectionAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceCollectionAccessKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceCollectionAccessKeysResult
	err := ctx.Invoke("azure-native:powerbi:listWorkspaceCollectionAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceCollectionAccessKeysArgs struct {
	// Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type ListWorkspaceCollectionAccessKeysResult struct {
	// Access key 1
	Key1 *string `pulumi:"key1"`
	// Access key 2
	Key2 *string `pulumi:"key2"`
}

func ListWorkspaceCollectionAccessKeysOutput(ctx *pulumi.Context, args ListWorkspaceCollectionAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceCollectionAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWorkspaceCollectionAccessKeysResultOutput, error) {
			args := v.(ListWorkspaceCollectionAccessKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:powerbi:listWorkspaceCollectionAccessKeys", args, ListWorkspaceCollectionAccessKeysResultOutput{}, options).(ListWorkspaceCollectionAccessKeysResultOutput), nil
		}).(ListWorkspaceCollectionAccessKeysResultOutput)
}

type ListWorkspaceCollectionAccessKeysOutputArgs struct {
	// Azure resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName pulumi.StringInput `pulumi:"workspaceCollectionName"`
}

func (ListWorkspaceCollectionAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceCollectionAccessKeysArgs)(nil)).Elem()
}

type ListWorkspaceCollectionAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceCollectionAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceCollectionAccessKeysResult)(nil)).Elem()
}

func (o ListWorkspaceCollectionAccessKeysResultOutput) ToListWorkspaceCollectionAccessKeysResultOutput() ListWorkspaceCollectionAccessKeysResultOutput {
	return o
}

func (o ListWorkspaceCollectionAccessKeysResultOutput) ToListWorkspaceCollectionAccessKeysResultOutputWithContext(ctx context.Context) ListWorkspaceCollectionAccessKeysResultOutput {
	return o
}

// Access key 1
func (o ListWorkspaceCollectionAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceCollectionAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

// Access key 2
func (o ListWorkspaceCollectionAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceCollectionAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceCollectionAccessKeysResultOutput{})
}
