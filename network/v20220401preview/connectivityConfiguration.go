// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The network manager connectivity configuration resource
type ConnectivityConfiguration struct {
	pulumi.CustomResourceState

	// Groups for configuration
	AppliesToGroups ConnectivityGroupItemResponseArrayOutput `pulumi:"appliesToGroups"`
	// Connectivity topology type.
	ConnectivityTopology pulumi.StringOutput `pulumi:"connectivityTopology"`
	// Flag if need to remove current existing peerings.
	DeleteExistingPeering pulumi.StringPtrOutput `pulumi:"deleteExistingPeering"`
	// A description of the connectivity configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of hubItems
	Hubs HubResponseArrayOutput `pulumi:"hubs"`
	// Flag if global mesh is supported.
	IsGlobal pulumi.StringPtrOutput `pulumi:"isGlobal"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the connectivity configuration resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectivityConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConnectivityConfiguration(ctx *pulumi.Context,
	name string, args *ConnectivityConfigurationArgs, opts ...pulumi.ResourceOption) (*ConnectivityConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppliesToGroups == nil {
		return nil, errors.New("invalid value for required argument 'AppliesToGroups'")
	}
	if args.ConnectivityTopology == nil {
		return nil, errors.New("invalid value for required argument 'ConnectivityTopology'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ConnectivityConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:ConnectivityConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectivityConfiguration
	err := ctx.RegisterResource("azure-native:network/v20220401preview:ConnectivityConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectivityConfiguration gets an existing ConnectivityConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectivityConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectivityConfigurationState, opts ...pulumi.ResourceOption) (*ConnectivityConfiguration, error) {
	var resource ConnectivityConfiguration
	err := ctx.ReadResource("azure-native:network/v20220401preview:ConnectivityConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectivityConfiguration resources.
type connectivityConfigurationState struct {
}

type ConnectivityConfigurationState struct {
}

func (ConnectivityConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityConfigurationState)(nil)).Elem()
}

type connectivityConfigurationArgs struct {
	// Groups for configuration
	AppliesToGroups []ConnectivityGroupItem `pulumi:"appliesToGroups"`
	// The name of the network manager connectivity configuration.
	ConfigurationName *string `pulumi:"configurationName"`
	// Connectivity topology type.
	ConnectivityTopology string `pulumi:"connectivityTopology"`
	// Flag if need to remove current existing peerings.
	DeleteExistingPeering *string `pulumi:"deleteExistingPeering"`
	// A description of the connectivity configuration.
	Description *string `pulumi:"description"`
	// List of hubItems
	Hubs []Hub `pulumi:"hubs"`
	// Flag if global mesh is supported.
	IsGlobal *string `pulumi:"isGlobal"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConnectivityConfiguration resource.
type ConnectivityConfigurationArgs struct {
	// Groups for configuration
	AppliesToGroups ConnectivityGroupItemArrayInput
	// The name of the network manager connectivity configuration.
	ConfigurationName pulumi.StringPtrInput
	// Connectivity topology type.
	ConnectivityTopology pulumi.StringInput
	// Flag if need to remove current existing peerings.
	DeleteExistingPeering pulumi.StringPtrInput
	// A description of the connectivity configuration.
	Description pulumi.StringPtrInput
	// List of hubItems
	Hubs HubArrayInput
	// Flag if global mesh is supported.
	IsGlobal pulumi.StringPtrInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectivityConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityConfigurationArgs)(nil)).Elem()
}

type ConnectivityConfigurationInput interface {
	pulumi.Input

	ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput
	ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput
}

func (*ConnectivityConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectivityConfiguration)(nil)).Elem()
}

func (i *ConnectivityConfiguration) ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput {
	return i.ToConnectivityConfigurationOutputWithContext(context.Background())
}

func (i *ConnectivityConfiguration) ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityConfigurationOutput)
}

type ConnectivityConfigurationOutput struct{ *pulumi.OutputState }

func (ConnectivityConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectivityConfiguration)(nil)).Elem()
}

func (o ConnectivityConfigurationOutput) ToConnectivityConfigurationOutput() ConnectivityConfigurationOutput {
	return o
}

func (o ConnectivityConfigurationOutput) ToConnectivityConfigurationOutputWithContext(ctx context.Context) ConnectivityConfigurationOutput {
	return o
}

// Groups for configuration
func (o ConnectivityConfigurationOutput) AppliesToGroups() ConnectivityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) ConnectivityGroupItemResponseArrayOutput { return v.AppliesToGroups }).(ConnectivityGroupItemResponseArrayOutput)
}

// Connectivity topology type.
func (o ConnectivityConfigurationOutput) ConnectivityTopology() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.ConnectivityTopology }).(pulumi.StringOutput)
}

// Flag if need to remove current existing peerings.
func (o ConnectivityConfigurationOutput) DeleteExistingPeering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringPtrOutput { return v.DeleteExistingPeering }).(pulumi.StringPtrOutput)
}

// A description of the connectivity configuration.
func (o ConnectivityConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ConnectivityConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of hubItems
func (o ConnectivityConfigurationOutput) Hubs() HubResponseArrayOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) HubResponseArrayOutput { return v.Hubs }).(HubResponseArrayOutput)
}

// Flag if global mesh is supported.
func (o ConnectivityConfigurationOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringPtrOutput { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ConnectivityConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the connectivity configuration resource.
func (o ConnectivityConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o ConnectivityConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ConnectivityConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectivityConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectivityConfigurationOutput{})
}