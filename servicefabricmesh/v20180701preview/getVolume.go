// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This type describes a volume resource.
//
// Deprecated: Version 2018-07-01-preview will be removed in v2 of the provider.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180701preview:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}

// This type describes a volume resource.
type LookupVolumeResult struct {
	AzureFileParameters *VolumeProviderParametersAzureFileResponse `pulumi:"azureFileParameters"`
	Description         *string                                    `pulumi:"description"`
	Id                  string                                     `pulumi:"id"`
	Location            string                                     `pulumi:"location"`
	Name                string                                     `pulumi:"name"`
	Provider            string                                     `pulumi:"provider"`
	ProvisioningState   string                                     `pulumi:"provisioningState"`
	Tags                map[string]string                          `pulumi:"tags"`
	Type                string                                     `pulumi:"type"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
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

func (o LookupVolumeResultOutput) AzureFileParameters() VolumeProviderParametersAzureFileResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumeProviderParametersAzureFileResponse { return v.AzureFileParameters }).(VolumeProviderParametersAzureFileResponsePtrOutput)
}

func (o LookupVolumeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Provider }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}