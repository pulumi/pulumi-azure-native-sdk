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

// An Azure Cosmos DB Throughputpool.
//
// Uses Azure REST API version 2024-12-01-preview.
//
// Other available API versions: 2023-11-15-preview, 2024-02-15-preview, 2024-05-15-preview, 2024-09-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ThroughputPool struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Value for throughput to be shared among CosmosDB resources in the pool.
	MaxThroughput pulumi.IntPtrOutput `pulumi:"maxThroughput"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A provisioning state of the ThroughputPool.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewThroughputPool registers a new resource with the given unique name, arguments, and options.
func NewThroughputPool(ctx *pulumi.Context,
	name string, args *ThroughputPoolArgs, opts ...pulumi.ResourceOption) (*ThroughputPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240215preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240901preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241201preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250501preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240215preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240901preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241201preview:ThroughputPool"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:ThroughputPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ThroughputPool
	err := ctx.RegisterResource("azure-native:cosmosdb:ThroughputPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThroughputPool gets an existing ThroughputPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThroughputPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThroughputPoolState, opts ...pulumi.ResourceOption) (*ThroughputPool, error) {
	var resource ThroughputPool
	err := ctx.ReadResource("azure-native:cosmosdb:ThroughputPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThroughputPool resources.
type throughputPoolState struct {
}

type ThroughputPoolState struct {
}

func (ThroughputPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*throughputPoolState)(nil)).Elem()
}

type throughputPoolArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Value for throughput to be shared among CosmosDB resources in the pool.
	MaxThroughput *int `pulumi:"maxThroughput"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Cosmos DB Throughput Pool name.
	ThroughputPoolName *string `pulumi:"throughputPoolName"`
}

// The set of arguments for constructing a ThroughputPool resource.
type ThroughputPoolArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Value for throughput to be shared among CosmosDB resources in the pool.
	MaxThroughput pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Cosmos DB Throughput Pool name.
	ThroughputPoolName pulumi.StringPtrInput
}

func (ThroughputPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*throughputPoolArgs)(nil)).Elem()
}

type ThroughputPoolInput interface {
	pulumi.Input

	ToThroughputPoolOutput() ThroughputPoolOutput
	ToThroughputPoolOutputWithContext(ctx context.Context) ThroughputPoolOutput
}

func (*ThroughputPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ThroughputPool)(nil)).Elem()
}

func (i *ThroughputPool) ToThroughputPoolOutput() ThroughputPoolOutput {
	return i.ToThroughputPoolOutputWithContext(context.Background())
}

func (i *ThroughputPool) ToThroughputPoolOutputWithContext(ctx context.Context) ThroughputPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThroughputPoolOutput)
}

type ThroughputPoolOutput struct{ *pulumi.OutputState }

func (ThroughputPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThroughputPool)(nil)).Elem()
}

func (o ThroughputPoolOutput) ToThroughputPoolOutput() ThroughputPoolOutput {
	return o
}

func (o ThroughputPoolOutput) ToThroughputPoolOutputWithContext(ctx context.Context) ThroughputPoolOutput {
	return o
}

// The Azure API version of the resource.
func (o ThroughputPoolOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ThroughputPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Value for throughput to be shared among CosmosDB resources in the pool.
func (o ThroughputPoolOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.IntPtrOutput { return v.MaxThroughput }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o ThroughputPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A provisioning state of the ThroughputPool.
func (o ThroughputPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ThroughputPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ThroughputPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ThroughputPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ThroughputPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ThroughputPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ThroughputPoolOutput{})
}
