// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the connection strings for the specified Azure Cosmos DB database account.
//
// Uses Azure REST API version 2024-11-15.
//
// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19, 2016-03-31, 2019-08-01, 2019-12-12, 2020-03-01, 2020-04-01, 2020-06-01-preview, 2020-09-01, 2021-01-15, 2021-03-01-preview, 2021-03-15, 2021-04-01-preview, 2021-04-15, 2021-05-15, 2021-06-15, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListDatabaseAccountConnectionStrings(ctx *pulumi.Context, args *ListDatabaseAccountConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountConnectionStringsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatabaseAccountConnectionStringsResult
	err := ctx.Invoke("azure-native:cosmosdb:listDatabaseAccountConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountConnectionStringsArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connection strings for the given database account.
type ListDatabaseAccountConnectionStringsResult struct {
	// An array that contains the connection strings for the Cosmos DB account.
	ConnectionStrings []DatabaseAccountConnectionStringResponse `pulumi:"connectionStrings"`
}

func ListDatabaseAccountConnectionStringsOutput(ctx *pulumi.Context, args ListDatabaseAccountConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseAccountConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabaseAccountConnectionStringsResultOutput, error) {
			args := v.(ListDatabaseAccountConnectionStringsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cosmosdb:listDatabaseAccountConnectionStrings", args, ListDatabaseAccountConnectionStringsResultOutput{}, options).(ListDatabaseAccountConnectionStringsResultOutput), nil
		}).(ListDatabaseAccountConnectionStringsResultOutput)
}

type ListDatabaseAccountConnectionStringsOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseAccountConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountConnectionStringsArgs)(nil)).Elem()
}

// The connection strings for the given database account.
type ListDatabaseAccountConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseAccountConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountConnectionStringsResult)(nil)).Elem()
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ToListDatabaseAccountConnectionStringsResultOutput() ListDatabaseAccountConnectionStringsResultOutput {
	return o
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ToListDatabaseAccountConnectionStringsResultOutputWithContext(ctx context.Context) ListDatabaseAccountConnectionStringsResultOutput {
	return o
}

// An array that contains the connection strings for the Cosmos DB account.
func (o ListDatabaseAccountConnectionStringsResultOutput) ConnectionStrings() DatabaseAccountConnectionStringResponseArrayOutput {
	return o.ApplyT(func(v ListDatabaseAccountConnectionStringsResult) []DatabaseAccountConnectionStringResponse {
		return v.ConnectionStrings
	}).(DatabaseAccountConnectionStringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseAccountConnectionStringsResultOutput{})
}
