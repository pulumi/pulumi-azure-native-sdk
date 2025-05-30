// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Table.
//
// Uses Azure REST API version 2016-03-31.
//
// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type DatabaseAccountTable struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountTable registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountTable(ctx *pulumi.Context,
	name string, args *DatabaseAccountTableArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cosmosdb/v20150401:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20150408:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20151106:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20160319:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20160331:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20190801:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20191212:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200301:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200401:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200601preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200901:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210115:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210301preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210315:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210401preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210415:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210515:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210615:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210701preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211015:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211015preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211115preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220215preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230301preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230415:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240215preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240815:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240901preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241115:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241201preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250415:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250501preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240215preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240815:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240901preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241115:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241201preview:TableResourceTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:TableResourceTable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseAccountTable
	err := ctx.RegisterResource("azure-native:cosmosdb:DatabaseAccountTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountTable gets an existing DatabaseAccountTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountTableState, opts ...pulumi.ResourceOption) (*DatabaseAccountTable, error) {
	var resource DatabaseAccountTable
	err := ctx.ReadResource("azure-native:cosmosdb:DatabaseAccountTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountTable resources.
type databaseAccountTableState struct {
}

type DatabaseAccountTableState struct {
}

func (DatabaseAccountTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountTableState)(nil)).Elem()
}

type databaseAccountTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a Table
	Resource TableResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a DatabaseAccountTable resource.
type DatabaseAccountTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a Table
	Resource TableResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB table name.
	TableName pulumi.StringPtrInput
}

func (DatabaseAccountTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountTableArgs)(nil)).Elem()
}

type DatabaseAccountTableInput interface {
	pulumi.Input

	ToDatabaseAccountTableOutput() DatabaseAccountTableOutput
	ToDatabaseAccountTableOutputWithContext(ctx context.Context) DatabaseAccountTableOutput
}

func (*DatabaseAccountTable) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountTable)(nil)).Elem()
}

func (i *DatabaseAccountTable) ToDatabaseAccountTableOutput() DatabaseAccountTableOutput {
	return i.ToDatabaseAccountTableOutputWithContext(context.Background())
}

func (i *DatabaseAccountTable) ToDatabaseAccountTableOutputWithContext(ctx context.Context) DatabaseAccountTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountTableOutput)
}

type DatabaseAccountTableOutput struct{ *pulumi.OutputState }

func (DatabaseAccountTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountTable)(nil)).Elem()
}

func (o DatabaseAccountTableOutput) ToDatabaseAccountTableOutput() DatabaseAccountTableOutput {
	return o
}

func (o DatabaseAccountTableOutput) ToDatabaseAccountTableOutputWithContext(ctx context.Context) DatabaseAccountTableOutput {
	return o
}

// The Azure API version of the resource.
func (o DatabaseAccountTableOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountTable) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o DatabaseAccountTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o DatabaseAccountTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o DatabaseAccountTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o DatabaseAccountTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountTableOutput{})
}
