// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tag Contract details.
func LookupTagByProduct(ctx *pulumi.Context, args *LookupTagByProductArgs, opts ...pulumi.InvokeOption) (*LookupTagByProductResult, error) {
	var rv LookupTagByProductResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getTagByProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// Tag Contract details.
type LookupTagByProductResult struct {
	// Tag name.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTagByProductOutput(ctx *pulumi.Context, args LookupTagByProductOutputArgs, opts ...pulumi.InvokeOption) LookupTagByProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagByProductResult, error) {
			args := v.(LookupTagByProductArgs)
			r, err := LookupTagByProduct(ctx, &args, opts...)
			var s LookupTagByProductResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagByProductResultOutput)
}

type LookupTagByProductOutputArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagByProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByProductArgs)(nil)).Elem()
}

// Tag Contract details.
type LookupTagByProductResultOutput struct{ *pulumi.OutputState }

func (LookupTagByProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByProductResult)(nil)).Elem()
}

func (o LookupTagByProductResultOutput) ToLookupTagByProductResultOutput() LookupTagByProductResultOutput {
	return o
}

func (o LookupTagByProductResultOutput) ToLookupTagByProductResultOutputWithContext(ctx context.Context) LookupTagByProductResultOutput {
	return o
}

// Tag name.
func (o LookupTagByProductResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTagByProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTagByProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTagByProductResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagByProductResultOutput{})
}