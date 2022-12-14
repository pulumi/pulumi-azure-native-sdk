// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20161002

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The result of the GetSupportedOptimizationTypes API
func GetProfileSupportedOptimizationTypes(ctx *pulumi.Context, args *GetProfileSupportedOptimizationTypesArgs, opts ...pulumi.InvokeOption) (*GetProfileSupportedOptimizationTypesResult, error) {
	var rv GetProfileSupportedOptimizationTypesResult
	err := ctx.Invoke("azure-native:cdn/v20161002:getProfileSupportedOptimizationTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProfileSupportedOptimizationTypesArgs struct {
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The result of the GetSupportedOptimizationTypes API
type GetProfileSupportedOptimizationTypesResult struct {
	// Supported optimization types for a profile.
	SupportedOptimizationTypes []string `pulumi:"supportedOptimizationTypes"`
}

func GetProfileSupportedOptimizationTypesOutput(ctx *pulumi.Context, args GetProfileSupportedOptimizationTypesOutputArgs, opts ...pulumi.InvokeOption) GetProfileSupportedOptimizationTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProfileSupportedOptimizationTypesResult, error) {
			args := v.(GetProfileSupportedOptimizationTypesArgs)
			r, err := GetProfileSupportedOptimizationTypes(ctx, &args, opts...)
			var s GetProfileSupportedOptimizationTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProfileSupportedOptimizationTypesResultOutput)
}

type GetProfileSupportedOptimizationTypesOutputArgs struct {
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetProfileSupportedOptimizationTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileSupportedOptimizationTypesArgs)(nil)).Elem()
}

// The result of the GetSupportedOptimizationTypes API
type GetProfileSupportedOptimizationTypesResultOutput struct{ *pulumi.OutputState }

func (GetProfileSupportedOptimizationTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileSupportedOptimizationTypesResult)(nil)).Elem()
}

func (o GetProfileSupportedOptimizationTypesResultOutput) ToGetProfileSupportedOptimizationTypesResultOutput() GetProfileSupportedOptimizationTypesResultOutput {
	return o
}

func (o GetProfileSupportedOptimizationTypesResultOutput) ToGetProfileSupportedOptimizationTypesResultOutputWithContext(ctx context.Context) GetProfileSupportedOptimizationTypesResultOutput {
	return o
}

// Supported optimization types for a profile.
func (o GetProfileSupportedOptimizationTypesResultOutput) SupportedOptimizationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProfileSupportedOptimizationTypesResult) []string { return v.SupportedOptimizationTypes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProfileSupportedOptimizationTypesResultOutput{})
}
