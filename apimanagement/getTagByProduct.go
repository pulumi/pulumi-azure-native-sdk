// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tag Contract details.
// API Version: 2020-12-01.
func LookupTagByProduct(ctx *pulumi.Context, args *LookupTagByProductArgs, opts ...pulumi.InvokeOption) (*LookupTagByProductResult, error) {
	var rv LookupTagByProductResult
	err := ctx.Invoke("azure-native:apimanagement:getTagByProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByProductArgs struct {
	ProductId         string `pulumi:"productId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}

// Tag Contract details.
type LookupTagByProductResult struct {
	DisplayName string `pulumi:"displayName"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
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
	ProductId         pulumi.StringInput `pulumi:"productId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TagId             pulumi.StringInput `pulumi:"tagId"`
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

func (o LookupTagByProductResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupTagByProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagByProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagByProductResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByProductResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagByProductResultOutput{})
}