// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private Endpoint Connection ARM resource.
type WebAppPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// PrivateEndpoint of a remote private endpoint connection
	PrivateEndpoint ArmIdWrapperResponsePtrOutput `pulumi:"privateEndpoint"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewWebAppPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *WebAppPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppPrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:web/v20190801:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppPrivateEndpointConnection gets an existing WebAppPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*WebAppPrivateEndpointConnection, error) {
	var resource WebAppPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppPrivateEndpointConnection resources.
type webAppPrivateEndpointConnectionState struct {
}

type WebAppPrivateEndpointConnectionState struct {
}

func (WebAppPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPrivateEndpointConnectionState)(nil)).Elem()
}

type webAppPrivateEndpointConnectionArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the site.
	Name                          string  `pulumi:"name"`
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppPrivateEndpointConnection resource.
type WebAppPrivateEndpointConnectionArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the site.
	Name                          pulumi.StringInput
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// The state of a private link connection
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPrivateEndpointConnectionArgs)(nil)).Elem()
}

type WebAppPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToWebAppPrivateEndpointConnectionOutput() WebAppPrivateEndpointConnectionOutput
	ToWebAppPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionOutput
}

func (*WebAppPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPrivateEndpointConnection)(nil)).Elem()
}

func (i *WebAppPrivateEndpointConnection) ToWebAppPrivateEndpointConnectionOutput() WebAppPrivateEndpointConnectionOutput {
	return i.ToWebAppPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *WebAppPrivateEndpointConnection) ToWebAppPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPrivateEndpointConnectionOutput)
}

type WebAppPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (WebAppPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPrivateEndpointConnection)(nil)).Elem()
}

func (o WebAppPrivateEndpointConnectionOutput) ToWebAppPrivateEndpointConnectionOutput() WebAppPrivateEndpointConnectionOutput {
	return o
}

func (o WebAppPrivateEndpointConnectionOutput) ToWebAppPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionOutput {
	return o
}

// Kind of resource.
func (o WebAppPrivateEndpointConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// PrivateEndpoint of a remote private endpoint connection
func (o WebAppPrivateEndpointConnectionOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) ArmIdWrapperResponsePtrOutput { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

// The state of a private link connection
func (o WebAppPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) PrivateLinkConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o WebAppPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o WebAppPrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o WebAppPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppPrivateEndpointConnectionOutput{})
}
