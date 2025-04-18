// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurearcdata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an Arc Sql Server database.
//
// Uses Azure REST API version 2024-01-01.
//
// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupSqlServerDatabase(ctx *pulumi.Context, args *LookupSqlServerDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlServerDatabaseResult
	err := ctx.Invoke("azure-native:azurearcdata:getSqlServerDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerDatabaseArgs struct {
	// Name of the database
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of SQL Server Instance
	SqlServerInstanceName string `pulumi:"sqlServerInstanceName"`
}

// Arc Sql Server database
type LookupSqlServerDatabaseResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of Arc Sql Server database
	Properties SqlServerDatabaseResourcePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSqlServerDatabaseOutput(ctx *pulumi.Context, args LookupSqlServerDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerDatabaseResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSqlServerDatabaseResultOutput, error) {
			args := v.(LookupSqlServerDatabaseArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azurearcdata:getSqlServerDatabase", args, LookupSqlServerDatabaseResultOutput{}, options).(LookupSqlServerDatabaseResultOutput), nil
		}).(LookupSqlServerDatabaseResultOutput)
}

type LookupSqlServerDatabaseOutputArgs struct {
	// Name of the database
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of SQL Server Instance
	SqlServerInstanceName pulumi.StringInput `pulumi:"sqlServerInstanceName"`
}

func (LookupSqlServerDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerDatabaseArgs)(nil)).Elem()
}

// Arc Sql Server database
type LookupSqlServerDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerDatabaseResult)(nil)).Elem()
}

func (o LookupSqlServerDatabaseResultOutput) ToLookupSqlServerDatabaseResultOutput() LookupSqlServerDatabaseResultOutput {
	return o
}

func (o LookupSqlServerDatabaseResultOutput) ToLookupSqlServerDatabaseResultOutputWithContext(ctx context.Context) LookupSqlServerDatabaseResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSqlServerDatabaseResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlServerDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSqlServerDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlServerDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of Arc Sql Server database
func (o LookupSqlServerDatabaseResultOutput) Properties() SqlServerDatabaseResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) SqlServerDatabaseResourcePropertiesResponse { return v.Properties }).(SqlServerDatabaseResourcePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSqlServerDatabaseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSqlServerDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlServerDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerDatabaseResultOutput{})
}
