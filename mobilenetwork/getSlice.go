// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilenetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified network slice.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupSlice(ctx *pulumi.Context, args *LookupSliceArgs, opts ...pulumi.InvokeOption) (*LookupSliceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSliceResult
	err := ctx.Invoke("azure-native:mobilenetwork:getSlice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSliceArgs struct {
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network slice.
	SliceName string `pulumi:"sliceName"`
}

// Network slice resource. Must be created in the same location as its parent mobile network.
type LookupSliceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// An optional description for this network slice.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the network slice resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
	Snssai SnssaiResponse `pulumi:"snssai"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSliceOutput(ctx *pulumi.Context, args LookupSliceOutputArgs, opts ...pulumi.InvokeOption) LookupSliceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSliceResultOutput, error) {
			args := v.(LookupSliceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:mobilenetwork:getSlice", args, LookupSliceResultOutput{}, options).(LookupSliceResultOutput), nil
		}).(LookupSliceResultOutput)
}

type LookupSliceOutputArgs struct {
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network slice.
	SliceName pulumi.StringInput `pulumi:"sliceName"`
}

func (LookupSliceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSliceArgs)(nil)).Elem()
}

// Network slice resource. Must be created in the same location as its parent mobile network.
type LookupSliceResultOutput struct{ *pulumi.OutputState }

func (LookupSliceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSliceResult)(nil)).Elem()
}

func (o LookupSliceResultOutput) ToLookupSliceResultOutput() LookupSliceResultOutput {
	return o
}

func (o LookupSliceResultOutput) ToLookupSliceResultOutputWithContext(ctx context.Context) LookupSliceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSliceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// An optional description for this network slice.
func (o LookupSliceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSliceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSliceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSliceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network slice resource.
func (o LookupSliceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
func (o LookupSliceResultOutput) Snssai() SnssaiResponseOutput {
	return o.ApplyT(func(v LookupSliceResult) SnssaiResponse { return v.Snssai }).(SnssaiResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSliceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSliceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSliceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSliceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSliceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSliceResultOutput{})
}
