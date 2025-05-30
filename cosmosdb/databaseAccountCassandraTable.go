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

// An Azure Cosmos DB Cassandra table.
//
// Uses Azure REST API version 2016-03-31.
//
// Other available API versions: 2015-04-01, 2015-04-08, 2015-11-06, 2016-03-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type DatabaseAccountCassandraTable struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Time to live of the Cosmos DB Cassandra table
	DefaultTtl pulumi.IntPtrOutput `pulumi:"defaultTtl"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schema of the Cosmos DB Cassandra table
	Schema CassandraSchemaResponsePtrOutput `pulumi:"schema"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountCassandraTable registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountCassandraTable(ctx *pulumi.Context,
	name string, args *DatabaseAccountCassandraTableArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.KeyspaceName == nil {
		return nil, errors.New("invalid value for required argument 'KeyspaceName'")
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
			Type: pulumi.String("azure-native:cosmosdb/v20150401:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20150408:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20151106:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20160319:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20160331:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20190801:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20191212:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200301:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200401:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200601preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20200901:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210115:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210301preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210315:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210401preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210415:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210515:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210615:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210701preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211015:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211015preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211115preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220215preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230301preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230415:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240215preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240815:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240901preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241115:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241201preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250415:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250501preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240215preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240815:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240901preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241115:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241201preview:CassandraResourceCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:CassandraResourceCassandraTable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseAccountCassandraTable
	err := ctx.RegisterResource("azure-native:cosmosdb:DatabaseAccountCassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountCassandraTable gets an existing DatabaseAccountCassandraTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountCassandraTableState, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraTable, error) {
	var resource DatabaseAccountCassandraTable
	err := ctx.ReadResource("azure-native:cosmosdb:DatabaseAccountCassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountCassandraTable resources.
type databaseAccountCassandraTableState struct {
}

type DatabaseAccountCassandraTableState struct {
}

func (DatabaseAccountCassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraTableState)(nil)).Elem()
}

type databaseAccountCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName string `pulumi:"keyspaceName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a DatabaseAccountCassandraTable resource.
type DatabaseAccountCassandraTableArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a Cassandra table
	Resource CassandraTableResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB table name.
	TableName pulumi.StringPtrInput
}

func (DatabaseAccountCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraTableArgs)(nil)).Elem()
}

type DatabaseAccountCassandraTableInput interface {
	pulumi.Input

	ToDatabaseAccountCassandraTableOutput() DatabaseAccountCassandraTableOutput
	ToDatabaseAccountCassandraTableOutputWithContext(ctx context.Context) DatabaseAccountCassandraTableOutput
}

func (*DatabaseAccountCassandraTable) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountCassandraTable)(nil)).Elem()
}

func (i *DatabaseAccountCassandraTable) ToDatabaseAccountCassandraTableOutput() DatabaseAccountCassandraTableOutput {
	return i.ToDatabaseAccountCassandraTableOutputWithContext(context.Background())
}

func (i *DatabaseAccountCassandraTable) ToDatabaseAccountCassandraTableOutputWithContext(ctx context.Context) DatabaseAccountCassandraTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountCassandraTableOutput)
}

type DatabaseAccountCassandraTableOutput struct{ *pulumi.OutputState }

func (DatabaseAccountCassandraTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountCassandraTable)(nil)).Elem()
}

func (o DatabaseAccountCassandraTableOutput) ToDatabaseAccountCassandraTableOutput() DatabaseAccountCassandraTableOutput {
	return o
}

func (o DatabaseAccountCassandraTableOutput) ToDatabaseAccountCassandraTableOutputWithContext(ctx context.Context) DatabaseAccountCassandraTableOutput {
	return o
}

// The Azure API version of the resource.
func (o DatabaseAccountCassandraTableOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Time to live of the Cosmos DB Cassandra table
func (o DatabaseAccountCassandraTableOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o DatabaseAccountCassandraTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o DatabaseAccountCassandraTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schema of the Cosmos DB Cassandra table
func (o DatabaseAccountCassandraTableOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) CassandraSchemaResponsePtrOutput { return v.Schema }).(CassandraSchemaResponsePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o DatabaseAccountCassandraTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o DatabaseAccountCassandraTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountCassandraTableOutput{})
}
