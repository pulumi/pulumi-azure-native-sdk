// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
// API Version: 2020-08-01.
func LookupSharedPrivateLinkResource(ctx *pulumi.Context, args *LookupSharedPrivateLinkResourceArgs, opts ...pulumi.InvokeOption) (*LookupSharedPrivateLinkResourceResult, error) {
	var rv LookupSharedPrivateLinkResourceResult
	err := ctx.Invoke("azure-native:search:getSharedPrivateLinkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSharedPrivateLinkResourceArgs struct {
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	SearchServiceName             string `pulumi:"searchServiceName"`
	SharedPrivateLinkResourceName string `pulumi:"sharedPrivateLinkResourceName"`
}

// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type LookupSharedPrivateLinkResourceResult struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties SharedPrivateLinkResourcePropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}

func LookupSharedPrivateLinkResourceOutput(ctx *pulumi.Context, args LookupSharedPrivateLinkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupSharedPrivateLinkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSharedPrivateLinkResourceResult, error) {
			args := v.(LookupSharedPrivateLinkResourceArgs)
			r, err := LookupSharedPrivateLinkResource(ctx, &args, opts...)
			var s LookupSharedPrivateLinkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSharedPrivateLinkResourceResultOutput)
}

type LookupSharedPrivateLinkResourceOutputArgs struct {
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	SearchServiceName             pulumi.StringInput `pulumi:"searchServiceName"`
	SharedPrivateLinkResourceName pulumi.StringInput `pulumi:"sharedPrivateLinkResourceName"`
}

func (LookupSharedPrivateLinkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSharedPrivateLinkResourceArgs)(nil)).Elem()
}

// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type LookupSharedPrivateLinkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupSharedPrivateLinkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSharedPrivateLinkResourceResult)(nil)).Elem()
}

func (o LookupSharedPrivateLinkResourceResultOutput) ToLookupSharedPrivateLinkResourceResultOutput() LookupSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupSharedPrivateLinkResourceResultOutput) ToLookupSharedPrivateLinkResourceResultOutputWithContext(ctx context.Context) LookupSharedPrivateLinkResourceResultOutput {
	return o
}

func (o LookupSharedPrivateLinkResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSharedPrivateLinkResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSharedPrivateLinkResourceResultOutput) Properties() SharedPrivateLinkResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) SharedPrivateLinkResourcePropertiesResponse {
		return v.Properties
	}).(SharedPrivateLinkResourcePropertiesResponseOutput)
}

func (o LookupSharedPrivateLinkResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSharedPrivateLinkResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSharedPrivateLinkResourceResultOutput{})
}