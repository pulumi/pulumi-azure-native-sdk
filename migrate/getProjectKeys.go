// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Log Analytics Workspace ID and Primary Key for the specified project.
//
// Uses Azure REST API version 2018-02-02.
func GetProjectKeys(ctx *pulumi.Context, args *GetProjectKeysArgs, opts ...pulumi.InvokeOption) (*GetProjectKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetProjectKeysResult
	err := ctx.Invoke("azure-native:migrate:getProjectKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProjectKeysArgs struct {
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ID and Key for Migration Project.
type GetProjectKeysResult struct {
	// ID of Migration Project.
	WorkspaceId string `pulumi:"workspaceId"`
	// Key of Migration Project.
	WorkspaceKey string `pulumi:"workspaceKey"`
}

func GetProjectKeysOutput(ctx *pulumi.Context, args GetProjectKeysOutputArgs, opts ...pulumi.InvokeOption) GetProjectKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectKeysResultOutput, error) {
			args := v.(GetProjectKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:migrate:getProjectKeys", args, GetProjectKeysResultOutput{}, options).(GetProjectKeysResultOutput), nil
		}).(GetProjectKeysResultOutput)
}

type GetProjectKeysOutputArgs struct {
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetProjectKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectKeysArgs)(nil)).Elem()
}

// ID and Key for Migration Project.
type GetProjectKeysResultOutput struct{ *pulumi.OutputState }

func (GetProjectKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectKeysResult)(nil)).Elem()
}

func (o GetProjectKeysResultOutput) ToGetProjectKeysResultOutput() GetProjectKeysResultOutput {
	return o
}

func (o GetProjectKeysResultOutput) ToGetProjectKeysResultOutputWithContext(ctx context.Context) GetProjectKeysResultOutput {
	return o
}

// ID of Migration Project.
func (o GetProjectKeysResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectKeysResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

// Key of Migration Project.
func (o GetProjectKeysResultOutput) WorkspaceKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectKeysResult) string { return v.WorkspaceKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectKeysResultOutput{})
}
