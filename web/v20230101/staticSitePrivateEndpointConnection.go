// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Remote Private Endpoint Connection ARM resource.
type StaticSitePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Private IPAddresses mapped to the remote private endpoint
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// PrivateEndpoint of a remote private endpoint connection
	PrivateEndpoint ArmIdWrapperResponsePtrOutput `pulumi:"privateEndpoint"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStaticSitePrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewStaticSitePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *StaticSitePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*StaticSitePrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web:StaticSitePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StaticSitePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:web/v20230101:StaticSitePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSitePrivateEndpointConnection gets an existing StaticSitePrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSitePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSitePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*StaticSitePrivateEndpointConnection, error) {
	var resource StaticSitePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:web/v20230101:StaticSitePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSitePrivateEndpointConnection resources.
type staticSitePrivateEndpointConnectionState struct {
}

type StaticSitePrivateEndpointConnectionState struct {
}

func (StaticSitePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSitePrivateEndpointConnectionState)(nil)).Elem()
}

type staticSitePrivateEndpointConnectionArgs struct {
	// Private IPAddresses mapped to the remote private endpoint
	IpAddresses []string `pulumi:"ipAddresses"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StaticSitePrivateEndpointConnection resource.
type StaticSitePrivateEndpointConnectionArgs struct {
	// Private IPAddresses mapped to the remote private endpoint
	IpAddresses pulumi.StringArrayInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the static site.
	Name pulumi.StringInput
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// The state of a private link connection
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (StaticSitePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSitePrivateEndpointConnectionArgs)(nil)).Elem()
}

type StaticSitePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToStaticSitePrivateEndpointConnectionOutput() StaticSitePrivateEndpointConnectionOutput
	ToStaticSitePrivateEndpointConnectionOutputWithContext(ctx context.Context) StaticSitePrivateEndpointConnectionOutput
}

func (*StaticSitePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSitePrivateEndpointConnection)(nil)).Elem()
}

func (i *StaticSitePrivateEndpointConnection) ToStaticSitePrivateEndpointConnectionOutput() StaticSitePrivateEndpointConnectionOutput {
	return i.ToStaticSitePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *StaticSitePrivateEndpointConnection) ToStaticSitePrivateEndpointConnectionOutputWithContext(ctx context.Context) StaticSitePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSitePrivateEndpointConnectionOutput)
}

type StaticSitePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (StaticSitePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSitePrivateEndpointConnection)(nil)).Elem()
}

func (o StaticSitePrivateEndpointConnectionOutput) ToStaticSitePrivateEndpointConnectionOutput() StaticSitePrivateEndpointConnectionOutput {
	return o
}

func (o StaticSitePrivateEndpointConnectionOutput) ToStaticSitePrivateEndpointConnectionOutputWithContext(ctx context.Context) StaticSitePrivateEndpointConnectionOutput {
	return o
}

// Private IPAddresses mapped to the remote private endpoint
func (o StaticSitePrivateEndpointConnectionOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Kind of resource.
func (o StaticSitePrivateEndpointConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o StaticSitePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// PrivateEndpoint of a remote private endpoint connection
func (o StaticSitePrivateEndpointConnectionOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) ArmIdWrapperResponsePtrOutput { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

// The state of a private link connection
func (o StaticSitePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) PrivateLinkConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o StaticSitePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSitePrivateEndpointConnectionOutput{})
}
