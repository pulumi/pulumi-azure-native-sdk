// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a Kusto kusto pool.
// API Version: 2021-04-01-preview.
type KustoPool struct {
	pulumi.CustomResourceState

	// The Kusto Pool data ingestion URI.
	DataIngestionUri pulumi.StringOutput `pulumi:"dataIngestionUri"`
	// The engine type
	EngineType pulumi.StringPtrOutput `pulumi:"engineType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU of the kusto pool.
	Sku AzureSkuResponseOutput `pulumi:"sku"`
	// The state of the resource.
	State pulumi.StringOutput `pulumi:"state"`
	// The reason for the Kusto Pool's current state.
	StateReason pulumi.StringOutput `pulumi:"stateReason"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The Kusto Pool URI.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// The workspace unique identifier.
	WorkspaceUid pulumi.StringPtrOutput `pulumi:"workspaceUid"`
}

// NewKustoPool registers a new resource with the given unique name, arguments, and options.
func NewKustoPool(ctx *pulumi.Context,
	name string, args *KustoPoolArgs, opts ...pulumi.ResourceOption) (*KustoPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:kustoPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:kustoPool"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPool
	err := ctx.RegisterResource("azure-native:synapse:kustoPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKustoPool gets an existing KustoPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKustoPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolState, opts ...pulumi.ResourceOption) (*KustoPool, error) {
	var resource KustoPool
	err := ctx.ReadResource("azure-native:synapse:kustoPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KustoPool resources.
type kustoPoolState struct {
}

type KustoPoolState struct {
}

func (KustoPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolState)(nil)).Elem()
}

type kustoPoolArgs struct {
	// The engine type
	EngineType *string `pulumi:"engineType"`
	// The name of the Kusto pool.
	KustoPoolName *string `pulumi:"kustoPoolName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the kusto pool.
	Sku AzureSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
	// The workspace unique identifier.
	WorkspaceUid *string `pulumi:"workspaceUid"`
}

// The set of arguments for constructing a KustoPool resource.
type KustoPoolArgs struct {
	// The engine type
	EngineType pulumi.StringPtrInput
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of the kusto pool.
	Sku AzureSkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
	// The workspace unique identifier.
	WorkspaceUid pulumi.StringPtrInput
}

func (KustoPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolArgs)(nil)).Elem()
}

type KustoPoolInput interface {
	pulumi.Input

	ToKustoPoolOutput() KustoPoolOutput
	ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput
}

func (*KustoPool) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPool)(nil)).Elem()
}

func (i *KustoPool) ToKustoPoolOutput() KustoPoolOutput {
	return i.ToKustoPoolOutputWithContext(context.Background())
}

func (i *KustoPool) ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolOutput)
}

type KustoPoolOutput struct{ *pulumi.OutputState }

func (KustoPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPool)(nil)).Elem()
}

func (o KustoPoolOutput) ToKustoPoolOutput() KustoPoolOutput {
	return o
}

func (o KustoPoolOutput) ToKustoPoolOutputWithContext(ctx context.Context) KustoPoolOutput {
	return o
}

// The Kusto Pool data ingestion URI.
func (o KustoPoolOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.DataIngestionUri }).(pulumi.StringOutput)
}

// The engine type
func (o KustoPoolOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringPtrOutput { return v.EngineType }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o KustoPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o KustoPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o KustoPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o KustoPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the kusto pool.
func (o KustoPoolOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v *KustoPool) AzureSkuResponseOutput { return v.Sku }).(AzureSkuResponseOutput)
}

// The state of the resource.
func (o KustoPoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The reason for the Kusto Pool's current state.
func (o KustoPoolOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o KustoPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o KustoPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o KustoPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The Kusto Pool URI.
func (o KustoPoolOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

// The workspace unique identifier.
func (o KustoPoolOutput) WorkspaceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPool) pulumi.StringPtrOutput { return v.WorkspaceUid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolOutput{})
}