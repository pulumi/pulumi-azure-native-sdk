// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Managed private endpoint resource type.
type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.ManagedVirtualNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedVirtualNetworkName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:ManagedPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure-native:datafactory/v20180601:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrivateEndpoint gets an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure-native:datafactory/v20180601:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
type managedPrivateEndpointState struct {
}

type ManagedPrivateEndpointState struct {
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Managed private endpoint name
	ManagedPrivateEndpointName *string `pulumi:"managedPrivateEndpointName"`
	// Managed virtual network name
	ManagedVirtualNetworkName string `pulumi:"managedVirtualNetworkName"`
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointType `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedPrivateEndpoint resource.
type ManagedPrivateEndpointArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput
	// Managed private endpoint name
	ManagedPrivateEndpointName pulumi.StringPtrInput
	// Managed virtual network name
	ManagedVirtualNetworkName pulumi.StringInput
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointTypeInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

// Etag identifies change in the resource.
func (o ManagedPrivateEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The resource name.
func (o ManagedPrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Managed private endpoint properties.
func (o ManagedPrivateEndpointOutput) Properties() ManagedPrivateEndpointResponseOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) ManagedPrivateEndpointResponseOutput { return v.Properties }).(ManagedPrivateEndpointResponseOutput)
}

// The resource type.
func (o ManagedPrivateEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
}