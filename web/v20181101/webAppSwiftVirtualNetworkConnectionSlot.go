// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Swift Virtual Network Contract. This is used to enable the new Swift way of doing virtual network integration.
type WebAppSwiftVirtualNetworkConnectionSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
	SubnetResourceId pulumi.StringPtrOutput `pulumi:"subnetResourceId"`
	// A flag that specifies if the scale unit this Web App is on supports Swift integration.
	SwiftSupported pulumi.BoolPtrOutput `pulumi:"swiftSupported"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSwiftVirtualNetworkConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppSwiftVirtualNetworkConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppSwiftVirtualNetworkConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSwiftVirtualNetworkConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSwiftVirtualNetworkConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSwiftVirtualNetworkConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppSwiftVirtualNetworkConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSwiftVirtualNetworkConnectionSlot gets an existing WebAppSwiftVirtualNetworkConnectionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSwiftVirtualNetworkConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSwiftVirtualNetworkConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppSwiftVirtualNetworkConnectionSlot, error) {
	var resource WebAppSwiftVirtualNetworkConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppSwiftVirtualNetworkConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSwiftVirtualNetworkConnectionSlot resources.
type webAppSwiftVirtualNetworkConnectionSlotState struct {
}

type WebAppSwiftVirtualNetworkConnectionSlotState struct {
}

func (WebAppSwiftVirtualNetworkConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSwiftVirtualNetworkConnectionSlotState)(nil)).Elem()
}

type webAppSwiftVirtualNetworkConnectionSlotArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will add or update connections for the production slot.
	Slot string `pulumi:"slot"`
	// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	// A flag that specifies if the scale unit this Web App is on supports Swift integration.
	SwiftSupported *bool `pulumi:"swiftSupported"`
}

// The set of arguments for constructing a WebAppSwiftVirtualNetworkConnectionSlot resource.
type WebAppSwiftVirtualNetworkConnectionSlotArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will add or update connections for the production slot.
	Slot pulumi.StringInput
	// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
	SubnetResourceId pulumi.StringPtrInput
	// A flag that specifies if the scale unit this Web App is on supports Swift integration.
	SwiftSupported pulumi.BoolPtrInput
}

func (WebAppSwiftVirtualNetworkConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSwiftVirtualNetworkConnectionSlotArgs)(nil)).Elem()
}

type WebAppSwiftVirtualNetworkConnectionSlotInput interface {
	pulumi.Input

	ToWebAppSwiftVirtualNetworkConnectionSlotOutput() WebAppSwiftVirtualNetworkConnectionSlotOutput
	ToWebAppSwiftVirtualNetworkConnectionSlotOutputWithContext(ctx context.Context) WebAppSwiftVirtualNetworkConnectionSlotOutput
}

func (*WebAppSwiftVirtualNetworkConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSwiftVirtualNetworkConnectionSlot)(nil)).Elem()
}

func (i *WebAppSwiftVirtualNetworkConnectionSlot) ToWebAppSwiftVirtualNetworkConnectionSlotOutput() WebAppSwiftVirtualNetworkConnectionSlotOutput {
	return i.ToWebAppSwiftVirtualNetworkConnectionSlotOutputWithContext(context.Background())
}

func (i *WebAppSwiftVirtualNetworkConnectionSlot) ToWebAppSwiftVirtualNetworkConnectionSlotOutputWithContext(ctx context.Context) WebAppSwiftVirtualNetworkConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSwiftVirtualNetworkConnectionSlotOutput)
}

type WebAppSwiftVirtualNetworkConnectionSlotOutput struct{ *pulumi.OutputState }

func (WebAppSwiftVirtualNetworkConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSwiftVirtualNetworkConnectionSlot)(nil)).Elem()
}

func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) ToWebAppSwiftVirtualNetworkConnectionSlotOutput() WebAppSwiftVirtualNetworkConnectionSlotOutput {
	return o
}

func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) ToWebAppSwiftVirtualNetworkConnectionSlotOutputWithContext(ctx context.Context) WebAppSwiftVirtualNetworkConnectionSlotOutput {
	return o
}

// Kind of resource.
func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnectionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnectionSlot) pulumi.StringPtrOutput { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

// A flag that specifies if the scale unit this Web App is on supports Swift integration.
func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) SwiftSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnectionSlot) pulumi.BoolPtrOutput { return v.SwiftSupported }).(pulumi.BoolPtrOutput)
}

// Resource type.
func (o WebAppSwiftVirtualNetworkConnectionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnectionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSwiftVirtualNetworkConnectionSlotOutput{})
}