// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The top level Linked service resource container.
func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	LinkedServiceName string `pulumi:"linkedServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

// The top level Linked service resource container.
type LookupLinkedServiceResult struct {
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	ResourceId string `pulumi:"resourceId"`
	Type       string `pulumi:"type"`
}

func LookupLinkedServiceOutput(ctx *pulumi.Context, args LookupLinkedServiceOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedServiceResult, error) {
			args := v.(LookupLinkedServiceArgs)
			r, err := LookupLinkedService(ctx, &args, opts...)
			var s LookupLinkedServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedServiceResultOutput)
}

type LookupLinkedServiceOutputArgs struct {
	LinkedServiceName pulumi.StringInput `pulumi:"linkedServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupLinkedServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceArgs)(nil)).Elem()
}

// The top level Linked service resource container.
type LookupLinkedServiceResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceResult)(nil)).Elem()
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutput() LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutputWithContext(ctx context.Context) LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServiceResultOutput{})
}