// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupRegistryComponentVersion(ctx *pulumi.Context, args *LookupRegistryComponentVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryComponentVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryComponentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20231001:getRegistryComponentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryComponentVersionArgs struct {
	// Container name.
	ComponentName string `pulumi:"componentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version string `pulumi:"version"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryComponentVersionResult struct {
	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionResponse `pulumi:"componentVersionProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryComponentVersionResult
func (val *LookupRegistryComponentVersionResult) Defaults() *LookupRegistryComponentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentVersionProperties = *tmp.ComponentVersionProperties.Defaults()

	return &tmp
}
func LookupRegistryComponentVersionOutput(ctx *pulumi.Context, args LookupRegistryComponentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryComponentVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRegistryComponentVersionResultOutput, error) {
			args := v.(LookupRegistryComponentVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices/v20231001:getRegistryComponentVersion", args, LookupRegistryComponentVersionResultOutput{}, options).(LookupRegistryComponentVersionResultOutput), nil
		}).(LookupRegistryComponentVersionResultOutput)
}

type LookupRegistryComponentVersionOutputArgs struct {
	// Container name.
	ComponentName pulumi.StringInput `pulumi:"componentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier.
	Version pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryComponentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryComponentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryComponentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentVersionResult)(nil)).Elem()
}

func (o LookupRegistryComponentVersionResultOutput) ToLookupRegistryComponentVersionResultOutput() LookupRegistryComponentVersionResultOutput {
	return o
}

func (o LookupRegistryComponentVersionResultOutput) ToLookupRegistryComponentVersionResultOutputWithContext(ctx context.Context) LookupRegistryComponentVersionResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryComponentVersionResultOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) ComponentVersionResponse {
		return v.ComponentVersionProperties
	}).(ComponentVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryComponentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryComponentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryComponentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryComponentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryComponentVersionResultOutput{})
}
