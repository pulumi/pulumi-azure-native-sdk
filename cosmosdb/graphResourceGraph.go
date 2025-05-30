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

// An Azure Cosmos DB Graph resource.
//
// Uses Azure REST API version 2024-12-01-preview.
//
// Other available API versions: 2021-07-01-preview, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15-preview, 2022-08-15-preview, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15-preview, 2023-09-15-preview, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15-preview, 2024-09-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type GraphResourceGraph struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                 `pulumi:"name"`
	Options  GraphResourceGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource GraphResourceGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGraphResourceGraph registers a new resource with the given unique name, arguments, and options.
func NewGraphResourceGraph(ctx *pulumi.Context,
	name string, args *GraphResourceGraphArgs, opts ...pulumi.ResourceOption) (*GraphResourceGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210701preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211015preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211115preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220215preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230301preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240215preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240901preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241201preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250501preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240215preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240901preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241201preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:GraphResourceGraph"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GraphResourceGraph
	err := ctx.RegisterResource("azure-native:cosmosdb:GraphResourceGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphResourceGraph gets an existing GraphResourceGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphResourceGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphResourceGraphState, opts ...pulumi.ResourceOption) (*GraphResourceGraph, error) {
	var resource GraphResourceGraph
	err := ctx.ReadResource("azure-native:cosmosdb:GraphResourceGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphResourceGraph resources.
type graphResourceGraphState struct {
}

type GraphResourceGraphState struct {
}

func (GraphResourceGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphResourceGraphState)(nil)).Elem()
}

type graphResourceGraphArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB graph resource name.
	GraphName *string `pulumi:"graphName"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Graph resource
	Resource GraphResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GraphResourceGraph resource.
type GraphResourceGraphArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB graph resource name.
	GraphName pulumi.StringPtrInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a Graph resource
	Resource GraphResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (GraphResourceGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphResourceGraphArgs)(nil)).Elem()
}

type GraphResourceGraphInput interface {
	pulumi.Input

	ToGraphResourceGraphOutput() GraphResourceGraphOutput
	ToGraphResourceGraphOutputWithContext(ctx context.Context) GraphResourceGraphOutput
}

func (*GraphResourceGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphResourceGraph)(nil)).Elem()
}

func (i *GraphResourceGraph) ToGraphResourceGraphOutput() GraphResourceGraphOutput {
	return i.ToGraphResourceGraphOutputWithContext(context.Background())
}

func (i *GraphResourceGraph) ToGraphResourceGraphOutputWithContext(ctx context.Context) GraphResourceGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphResourceGraphOutput)
}

type GraphResourceGraphOutput struct{ *pulumi.OutputState }

func (GraphResourceGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphResourceGraph)(nil)).Elem()
}

func (o GraphResourceGraphOutput) ToGraphResourceGraphOutput() GraphResourceGraphOutput {
	return o
}

func (o GraphResourceGraphOutput) ToGraphResourceGraphOutputWithContext(ctx context.Context) GraphResourceGraphOutput {
	return o
}

// The Azure API version of the resource.
func (o GraphResourceGraphOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o GraphResourceGraphOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o GraphResourceGraphOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o GraphResourceGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GraphResourceGraphOutput) Options() GraphResourceGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) GraphResourceGetPropertiesResponseOptionsPtrOutput { return v.Options }).(GraphResourceGetPropertiesResponseOptionsPtrOutput)
}

func (o GraphResourceGraphOutput) Resource() GraphResourceGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) GraphResourceGetPropertiesResponseResourcePtrOutput { return v.Resource }).(GraphResourceGetPropertiesResponseResourcePtrOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o GraphResourceGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o GraphResourceGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphResourceGraphOutput{})
}
