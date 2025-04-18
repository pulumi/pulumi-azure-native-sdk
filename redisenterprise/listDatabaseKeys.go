// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redisenterprise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the access keys for the RedisEnterprise database.
//
// Uses Azure REST API version 2024-03-01-preview.
//
// Other available API versions: 2020-10-01-preview, 2021-02-01-preview, 2021-03-01, 2021-08-01, 2022-01-01, 2022-11-01-preview, 2023-03-01-preview, 2023-07-01, 2023-08-01-preview, 2023-10-01-preview, 2023-11-01, 2024-02-01, 2024-06-01-preview, 2024-09-01-preview, 2024-10-01, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redisenterprise [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListDatabaseKeys(ctx *pulumi.Context, args *ListDatabaseKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatabaseKeysResult
	err := ctx.Invoke("azure-native:redisenterprise:listDatabaseKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseKeysArgs struct {
	// The name of the Redis Enterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the Redis Enterprise database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The secret access keys used for authenticating connections to redis
type ListDatabaseKeysResult struct {
	// The current primary key that clients can use to authenticate
	PrimaryKey string `pulumi:"primaryKey"`
	// The current secondary key that clients can use to authenticate
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListDatabaseKeysOutput(ctx *pulumi.Context, args ListDatabaseKeysOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabaseKeysResultOutput, error) {
			args := v.(ListDatabaseKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:redisenterprise:listDatabaseKeys", args, ListDatabaseKeysResultOutput{}, options).(ListDatabaseKeysResultOutput), nil
		}).(ListDatabaseKeysResultOutput)
}

type ListDatabaseKeysOutputArgs struct {
	// The name of the Redis Enterprise cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the Redis Enterprise database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseKeysArgs)(nil)).Elem()
}

// The secret access keys used for authenticating connections to redis
type ListDatabaseKeysResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseKeysResult)(nil)).Elem()
}

func (o ListDatabaseKeysResultOutput) ToListDatabaseKeysResultOutput() ListDatabaseKeysResultOutput {
	return o
}

func (o ListDatabaseKeysResultOutput) ToListDatabaseKeysResultOutputWithContext(ctx context.Context) ListDatabaseKeysResultOutput {
	return o
}

// The current primary key that clients can use to authenticate
func (o ListDatabaseKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The current secondary key that clients can use to authenticate
func (o ListDatabaseKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseKeysResultOutput{})
}
