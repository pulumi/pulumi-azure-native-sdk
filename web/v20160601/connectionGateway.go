// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The gateway definition
type ConnectionGateway struct {
	pulumi.CustomResourceState

	// Resource ETag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name       pulumi.StringOutput                                 `pulumi:"name"`
	Properties ConnectionGatewayDefinitionResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectionGateway registers a new resource with the given unique name, arguments, and options.
func NewConnectionGateway(ctx *pulumi.Context,
	name string, args *ConnectionGatewayArgs, opts ...pulumi.ResourceOption) (*ConnectionGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:ConnectionGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectionGateway
	err := ctx.RegisterResource("azure-native:web/v20160601:ConnectionGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionGateway gets an existing ConnectionGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionGatewayState, opts ...pulumi.ResourceOption) (*ConnectionGateway, error) {
	var resource ConnectionGateway
	err := ctx.ReadResource("azure-native:web/v20160601:ConnectionGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionGateway resources.
type connectionGatewayState struct {
}

type ConnectionGatewayState struct {
}

func (ConnectionGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionGatewayState)(nil)).Elem()
}

type connectionGatewayArgs struct {
	// The connection gateway name
	ConnectionGatewayName *string `pulumi:"connectionGatewayName"`
	// Resource location
	Location   *string                                `pulumi:"location"`
	Properties *ConnectionGatewayDefinitionProperties `pulumi:"properties"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConnectionGateway resource.
type ConnectionGatewayArgs struct {
	// The connection gateway name
	ConnectionGatewayName pulumi.StringPtrInput
	// Resource location
	Location   pulumi.StringPtrInput
	Properties ConnectionGatewayDefinitionPropertiesPtrInput
	// The resource group
	ResourceGroupName pulumi.StringInput
	// Subscription Id
	SubscriptionId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ConnectionGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionGatewayArgs)(nil)).Elem()
}

type ConnectionGatewayInput interface {
	pulumi.Input

	ToConnectionGatewayOutput() ConnectionGatewayOutput
	ToConnectionGatewayOutputWithContext(ctx context.Context) ConnectionGatewayOutput
}

func (*ConnectionGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGateway)(nil)).Elem()
}

func (i *ConnectionGateway) ToConnectionGatewayOutput() ConnectionGatewayOutput {
	return i.ToConnectionGatewayOutputWithContext(context.Background())
}

func (i *ConnectionGateway) ToConnectionGatewayOutputWithContext(ctx context.Context) ConnectionGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayOutput)
}

type ConnectionGatewayOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionGateway)(nil)).Elem()
}

func (o ConnectionGatewayOutput) ToConnectionGatewayOutput() ConnectionGatewayOutput {
	return o
}

func (o ConnectionGatewayOutput) ToConnectionGatewayOutputWithContext(ctx context.Context) ConnectionGatewayOutput {
	return o
}

// Resource ETag
func (o ConnectionGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location
func (o ConnectionGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o ConnectionGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionGatewayOutput) Properties() ConnectionGatewayDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v *ConnectionGateway) ConnectionGatewayDefinitionResponsePropertiesOutput { return v.Properties }).(ConnectionGatewayDefinitionResponsePropertiesOutput)
}

// Resource tags
func (o ConnectionGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ConnectionGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionGatewayOutput{})
}