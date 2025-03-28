// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2024-04-01-preview.
//
// Other available API versions: 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview.
func LookupConnectionRaiBlocklist(ctx *pulumi.Context, args *LookupConnectionRaiBlocklistArgs, opts ...pulumi.InvokeOption) (*LookupConnectionRaiBlocklistResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionRaiBlocklistResult
	err := ctx.Invoke("azure-native:machinelearningservices:getConnectionRaiBlocklist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionRaiBlocklistArgs struct {
	// Friendly name of the workspace connection
	ConnectionName string `pulumi:"connectionName"`
	// Name of the RaiBlocklist Item
	RaiBlocklistItemName string `pulumi:"raiBlocklistItemName"`
	// The name of the RaiBlocklist.
	RaiBlocklistName string `pulumi:"raiBlocklistName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupConnectionRaiBlocklistResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// RAI Custom Blocklist Item properties.
	Properties RaiBlocklistItemPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupConnectionRaiBlocklistOutput(ctx *pulumi.Context, args LookupConnectionRaiBlocklistOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionRaiBlocklistResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupConnectionRaiBlocklistResultOutput, error) {
			args := v.(LookupConnectionRaiBlocklistArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices:getConnectionRaiBlocklist", args, LookupConnectionRaiBlocklistResultOutput{}, options).(LookupConnectionRaiBlocklistResultOutput), nil
		}).(LookupConnectionRaiBlocklistResultOutput)
}

type LookupConnectionRaiBlocklistOutputArgs struct {
	// Friendly name of the workspace connection
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// Name of the RaiBlocklist Item
	RaiBlocklistItemName pulumi.StringInput `pulumi:"raiBlocklistItemName"`
	// The name of the RaiBlocklist.
	RaiBlocklistName pulumi.StringInput `pulumi:"raiBlocklistName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupConnectionRaiBlocklistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionRaiBlocklistArgs)(nil)).Elem()
}

type LookupConnectionRaiBlocklistResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionRaiBlocklistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionRaiBlocklistResult)(nil)).Elem()
}

func (o LookupConnectionRaiBlocklistResultOutput) ToLookupConnectionRaiBlocklistResultOutput() LookupConnectionRaiBlocklistResultOutput {
	return o
}

func (o LookupConnectionRaiBlocklistResultOutput) ToLookupConnectionRaiBlocklistResultOutputWithContext(ctx context.Context) LookupConnectionRaiBlocklistResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectionRaiBlocklistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionRaiBlocklistResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConnectionRaiBlocklistResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionRaiBlocklistResult) string { return v.Name }).(pulumi.StringOutput)
}

// RAI Custom Blocklist Item properties.
func (o LookupConnectionRaiBlocklistResultOutput) Properties() RaiBlocklistItemPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectionRaiBlocklistResult) RaiBlocklistItemPropertiesResponse { return v.Properties }).(RaiBlocklistItemPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConnectionRaiBlocklistResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectionRaiBlocklistResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectionRaiBlocklistResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionRaiBlocklistResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionRaiBlocklistResultOutput{})
}
