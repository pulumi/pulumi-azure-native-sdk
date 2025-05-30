// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the information about the volume resource with the given name. The information include the description and other properties of the volume.
//
// Uses Azure REST API version 2018-09-01-preview.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The identity of the volume.
	VolumeResourceName string `pulumi:"volumeResourceName"`
}

// This type describes a volume resource.
type LookupVolumeResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters *VolumeProviderParametersAzureFileResponse `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description *string `pulumi:"description"`
	// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provider of the volume.
	Provider string `pulumi:"provider"`
	// State of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Status of the volume.
	Status string `pulumi:"status"`
	// Gives additional information about the current status of the volume.
	StatusDetails string `pulumi:"statusDetails"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVolumeResultOutput, error) {
			args := v.(LookupVolumeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:servicefabricmesh:getVolume", args, LookupVolumeResultOutput{}, options).(LookupVolumeResultOutput), nil
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	// Azure resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The identity of the volume.
	VolumeResourceName pulumi.StringInput `pulumi:"volumeResourceName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

// This type describes a volume resource.
type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupVolumeResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// This type describes a volume provided by an Azure Files file share.
func (o LookupVolumeResultOutput) AzureFileParameters() VolumeProviderParametersAzureFileResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumeProviderParametersAzureFileResponse { return v.AzureFileParameters }).(VolumeProviderParametersAzureFileResponsePtrOutput)
}

// User readable description of the volume.
func (o LookupVolumeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provider of the volume.
func (o LookupVolumeResultOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Provider }).(pulumi.StringOutput)
}

// State of the resource.
func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Status of the volume.
func (o LookupVolumeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Status }).(pulumi.StringOutput)
}

// Gives additional information about the current status of the volume.
func (o LookupVolumeResultOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.StatusDetails }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
