// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200815preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Arc PrivateLinkScope definition.
type PrivateLinkScope struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Current state of this PrivateLinkScope: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Provisioning ,Succeeded, Canceled and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkScope registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkScope(ctx *pulumi.Context,
	name string, args *PrivateLinkScopeArgs, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220811preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221110:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221227:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221227preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20230315preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20230620preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20231003preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240331preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240520preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240710:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240731preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240910preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20241110preview:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20250113:PrivateLinkScope"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute:PrivateLinkScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkScope
	err := ctx.RegisterResource("azure-native:hybridcompute/v20200815preview:PrivateLinkScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkScope gets an existing PrivateLinkScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkScopeState, opts ...pulumi.ResourceOption) (*PrivateLinkScope, error) {
	var resource PrivateLinkScope
	err := ctx.ReadResource("azure-native:hybridcompute/v20200815preview:PrivateLinkScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkScope resources.
type privateLinkScopeState struct {
}

type PrivateLinkScopeState struct {
}

func (PrivateLinkScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeState)(nil)).Elem()
}

type privateLinkScopeArgs struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName *string `pulumi:"scopeName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkScope resource.
type PrivateLinkScopeArgs struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (PrivateLinkScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopeArgs)(nil)).Elem()
}

type PrivateLinkScopeInput interface {
	pulumi.Input

	ToPrivateLinkScopeOutput() PrivateLinkScopeOutput
	ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput
}

func (*PrivateLinkScope) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return i.ToPrivateLinkScopeOutputWithContext(context.Background())
}

func (i *PrivateLinkScope) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopeOutput)
}

type PrivateLinkScopeOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScope)(nil)).Elem()
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutput() PrivateLinkScopeOutput {
	return o
}

func (o PrivateLinkScopeOutput) ToPrivateLinkScopeOutputWithContext(ctx context.Context) PrivateLinkScopeOutput {
	return o
}

// Resource location
func (o PrivateLinkScopeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o PrivateLinkScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections.
func (o PrivateLinkScopeOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkScope) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Current state of this PrivateLinkScope: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Provisioning ,Succeeded, Canceled and Failed.
func (o PrivateLinkScopeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
func (o PrivateLinkScopeOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o PrivateLinkScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o PrivateLinkScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopeOutput{})
}
