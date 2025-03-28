// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve a Redis cache's access keys. This operation requires write permission to the cache resource.
//
// Uses Azure REST API version 2023-04-01.
//
// Other available API versions: 2020-06-01, 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview, 2024-11-01.
func ListRedisKeys(ctx *pulumi.Context, args *ListRedisKeysArgs, opts ...pulumi.InvokeOption) (*ListRedisKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListRedisKeysResult
	err := ctx.Invoke("azure-native:cache:listRedisKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRedisKeysArgs struct {
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Redis cache access keys.
type ListRedisKeysResult struct {
	// The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey string `pulumi:"primaryKey"`
	// The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListRedisKeysOutput(ctx *pulumi.Context, args ListRedisKeysOutputArgs, opts ...pulumi.InvokeOption) ListRedisKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListRedisKeysResultOutput, error) {
			args := v.(ListRedisKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cache:listRedisKeys", args, ListRedisKeysResultOutput{}, options).(ListRedisKeysResultOutput), nil
		}).(ListRedisKeysResultOutput)
}

type ListRedisKeysOutputArgs struct {
	// The name of the Redis cache.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRedisKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRedisKeysArgs)(nil)).Elem()
}

// Redis cache access keys.
type ListRedisKeysResultOutput struct{ *pulumi.OutputState }

func (ListRedisKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRedisKeysResult)(nil)).Elem()
}

func (o ListRedisKeysResultOutput) ToListRedisKeysResultOutput() ListRedisKeysResultOutput {
	return o
}

func (o ListRedisKeysResultOutput) ToListRedisKeysResultOutputWithContext(ctx context.Context) ListRedisKeysResultOutput {
	return o
}

// The current primary key that clients can use to authenticate with Redis cache.
func (o ListRedisKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRedisKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The current secondary key that clients can use to authenticate with Redis cache.
func (o ListRedisKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRedisKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRedisKeysResultOutput{})
}
