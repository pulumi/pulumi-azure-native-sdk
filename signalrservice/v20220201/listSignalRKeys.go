// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A class represents the access keys of the resource.
func ListSignalRKeys(ctx *pulumi.Context, args *ListSignalRKeysArgs, opts ...pulumi.InvokeOption) (*ListSignalRKeysResult, error) {
	var rv ListSignalRKeysResult
	err := ctx.Invoke("azure-native:signalrservice/v20220201:listSignalRKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSignalRKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}

// A class represents the access keys of the resource.
type ListSignalRKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListSignalRKeysOutput(ctx *pulumi.Context, args ListSignalRKeysOutputArgs, opts ...pulumi.InvokeOption) ListSignalRKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSignalRKeysResult, error) {
			args := v.(ListSignalRKeysArgs)
			r, err := ListSignalRKeys(ctx, &args, opts...)
			var s ListSignalRKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSignalRKeysResultOutput)
}

type ListSignalRKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListSignalRKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSignalRKeysArgs)(nil)).Elem()
}

// A class represents the access keys of the resource.
type ListSignalRKeysResultOutput struct{ *pulumi.OutputState }

func (ListSignalRKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSignalRKeysResult)(nil)).Elem()
}

func (o ListSignalRKeysResultOutput) ToListSignalRKeysResultOutput() ListSignalRKeysResultOutput {
	return o
}

func (o ListSignalRKeysResultOutput) ToListSignalRKeysResultOutputWithContext(ctx context.Context) ListSignalRKeysResultOutput {
	return o
}

func (o ListSignalRKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListSignalRKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListSignalRKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListSignalRKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSignalRKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSignalRKeysResultOutput{})
}