// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Tables under an existing Azure Cosmos DB database account with the provided name.
//
// Uses Azure REST API version 2024-11-15.
//
// Other available API versions: 2019-08-01, 2019-12-12, 2020-03-01, 2020-04-01, 2020-06-01-preview, 2020-09-01, 2021-01-15, 2021-03-01-preview, 2021-03-15, 2021-04-01-preview, 2021-04-15, 2021-05-15, 2021-06-15, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupTableResourceTable(ctx *pulumi.Context, args *LookupTableResourceTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResourceTableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTableResourceTableResult
	err := ctx.Invoke("azure-native:cosmosdb:getTableResourceTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableResourceTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName string `pulumi:"tableName"`
}

// An Azure Cosmos DB Table.
type LookupTableResourceTableResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     string                              `pulumi:"name"`
	Options  *TableGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *TableGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupTableResourceTableOutput(ctx *pulumi.Context, args LookupTableResourceTableOutputArgs, opts ...pulumi.InvokeOption) LookupTableResourceTableResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTableResourceTableResultOutput, error) {
			args := v.(LookupTableResourceTableArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cosmosdb:getTableResourceTable", args, LookupTableResourceTableResultOutput{}, options).(LookupTableResourceTableResultOutput), nil
		}).(LookupTableResourceTableResultOutput)
}

type LookupTableResourceTableOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupTableResourceTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResourceTableArgs)(nil)).Elem()
}

// An Azure Cosmos DB Table.
type LookupTableResourceTableResultOutput struct{ *pulumi.OutputState }

func (LookupTableResourceTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResourceTableResult)(nil)).Elem()
}

func (o LookupTableResourceTableResultOutput) ToLookupTableResourceTableResultOutput() LookupTableResourceTableResultOutput {
	return o
}

func (o LookupTableResourceTableResultOutput) ToLookupTableResourceTableResultOutputWithContext(ctx context.Context) LookupTableResourceTableResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupTableResourceTableResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The unique resource identifier of the ARM resource.
func (o LookupTableResourceTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupTableResourceTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupTableResourceTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTableResourceTableResultOutput) Options() TableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *TableGetPropertiesResponseOptions { return v.Options }).(TableGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupTableResourceTableResultOutput) Resource() TableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) *TableGetPropertiesResponseResource { return v.Resource }).(TableGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupTableResourceTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupTableResourceTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResourceTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableResourceTableResultOutput{})
}
