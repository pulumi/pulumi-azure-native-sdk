// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version Set Contract details.
func LookupApiVersionSet(ctx *pulumi.Context, args *LookupApiVersionSetArgs, opts ...pulumi.InvokeOption) (*LookupApiVersionSetResult, error) {
	var rv LookupApiVersionSetResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getApiVersionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiVersionSetArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Api Version Set identifier. Must be unique in the current API Management service instance.
	VersionSetId string `pulumi:"versionSetId"`
}

// API Version Set Contract details.
type LookupApiVersionSetResult struct {
	// Description of API Version Set.
	Description *string `pulumi:"description"`
	// Name of API Version Set
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	// Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `pulumi:"versionQueryName"`
	// An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme string `pulumi:"versioningScheme"`
}

func LookupApiVersionSetOutput(ctx *pulumi.Context, args LookupApiVersionSetOutputArgs, opts ...pulumi.InvokeOption) LookupApiVersionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiVersionSetResult, error) {
			args := v.(LookupApiVersionSetArgs)
			r, err := LookupApiVersionSet(ctx, &args, opts...)
			var s LookupApiVersionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiVersionSetResultOutput)
}

type LookupApiVersionSetOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Api Version Set identifier. Must be unique in the current API Management service instance.
	VersionSetId pulumi.StringInput `pulumi:"versionSetId"`
}

func (LookupApiVersionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiVersionSetArgs)(nil)).Elem()
}

// API Version Set Contract details.
type LookupApiVersionSetResultOutput struct{ *pulumi.OutputState }

func (LookupApiVersionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiVersionSetResult)(nil)).Elem()
}

func (o LookupApiVersionSetResultOutput) ToLookupApiVersionSetResultOutput() LookupApiVersionSetResultOutput {
	return o
}

func (o LookupApiVersionSetResultOutput) ToLookupApiVersionSetResultOutputWithContext(ctx context.Context) LookupApiVersionSetResultOutput {
	return o
}

// Description of API Version Set.
func (o LookupApiVersionSetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of API Version Set
func (o LookupApiVersionSetResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiVersionSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApiVersionSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiVersionSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.Type }).(pulumi.StringOutput)
}

// Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
func (o LookupApiVersionSetResultOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

// Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
func (o LookupApiVersionSetResultOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

// An value that determines where the API Version identifier will be located in a HTTP request.
func (o LookupApiVersionSetResultOutput) VersioningScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.VersioningScheme }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiVersionSetResultOutput{})
}