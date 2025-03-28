// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupDataContainer(ctx *pulumi.Context, args *LookupDataContainerArgs, opts ...pulumi.InvokeOption) (*LookupDataContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20250101preview:getDataContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDataContainerArgs struct {
	// Container name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupDataContainerResult struct {
	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerResponse `pulumi:"dataContainerProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDataContainerResult
func (val *LookupDataContainerResult) Defaults() *LookupDataContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DataContainerProperties = *tmp.DataContainerProperties.Defaults()

	return &tmp
}
func LookupDataContainerOutput(ctx *pulumi.Context, args LookupDataContainerOutputArgs, opts ...pulumi.InvokeOption) LookupDataContainerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataContainerResultOutput, error) {
			args := v.(LookupDataContainerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices/v20250101preview:getDataContainer", args, LookupDataContainerResultOutput{}, options).(LookupDataContainerResultOutput), nil
		}).(LookupDataContainerResultOutput)
}

type LookupDataContainerOutputArgs struct {
	// Container name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupDataContainerResultOutput struct{ *pulumi.OutputState }

func (LookupDataContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataContainerResult)(nil)).Elem()
}

func (o LookupDataContainerResultOutput) ToLookupDataContainerResultOutput() LookupDataContainerResultOutput {
	return o
}

func (o LookupDataContainerResultOutput) ToLookupDataContainerResultOutputWithContext(ctx context.Context) LookupDataContainerResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupDataContainerResultOutput) DataContainerProperties() DataContainerResponseOutput {
	return o.ApplyT(func(v LookupDataContainerResult) DataContainerResponse { return v.DataContainerProperties }).(DataContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDataContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataContainerResultOutput{})
}
