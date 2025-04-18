// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Managed Network resource
//
// Uses Azure REST API version 2019-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2019-06-01-preview.
type ManagedNetwork struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The collection of groups and policies concerned with connectivity
	Connectivity ConnectivityCollectionResponseOutput `pulumi:"connectivity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope ScopeResponsePtrOutput `pulumi:"scope"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedNetwork registers a new resource with the given unique name, arguments, and options.
func NewManagedNetwork(ctx *pulumi.Context,
	name string, args *ManagedNetworkArgs, opts ...pulumi.ResourceOption) (*ManagedNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetwork/v20190601preview:ManagedNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedNetwork
	err := ctx.RegisterResource("azure-native:managednetwork:ManagedNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedNetwork gets an existing ManagedNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkState, opts ...pulumi.ResourceOption) (*ManagedNetwork, error) {
	var resource ManagedNetwork
	err := ctx.ReadResource("azure-native:managednetwork:ManagedNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedNetwork resources.
type managedNetworkState struct {
}

type ManagedNetworkState struct {
}

func (ManagedNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkState)(nil)).Elem()
}

type managedNetworkArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Managed Network.
	ManagedNetworkName *string `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope *Scope `pulumi:"scope"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedNetwork resource.
type ManagedNetworkArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Managed Network.
	ManagedNetworkName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
	Scope ScopePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ManagedNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkArgs)(nil)).Elem()
}

type ManagedNetworkInput interface {
	pulumi.Input

	ToManagedNetworkOutput() ManagedNetworkOutput
	ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput
}

func (*ManagedNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetwork)(nil)).Elem()
}

func (i *ManagedNetwork) ToManagedNetworkOutput() ManagedNetworkOutput {
	return i.ToManagedNetworkOutputWithContext(context.Background())
}

func (i *ManagedNetwork) ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkOutput)
}

type ManagedNetworkOutput struct{ *pulumi.OutputState }

func (ManagedNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetwork)(nil)).Elem()
}

func (o ManagedNetworkOutput) ToManagedNetworkOutput() ManagedNetworkOutput {
	return o
}

func (o ManagedNetworkOutput) ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput {
	return o
}

// The Azure API version of the resource.
func (o ManagedNetworkOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The collection of groups and policies concerned with connectivity
func (o ManagedNetworkOutput) Connectivity() ConnectivityCollectionResponseOutput {
	return o.ApplyT(func(v *ManagedNetwork) ConnectivityCollectionResponseOutput { return v.Connectivity }).(ConnectivityCollectionResponseOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ManagedNetworkOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ManagedNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ManagedNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the ManagedNetwork resource.
func (o ManagedNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only property that is reflective of all ScopeAssignments for this Managed Network
func (o ManagedNetworkOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v *ManagedNetwork) ScopeResponsePtrOutput { return v.Scope }).(ScopeResponsePtrOutput)
}

// Resource tags
func (o ManagedNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o ManagedNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkOutput{})
}
