// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A domain specific resource identifier.
type WebAppDomainOwnershipIdentifierSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// String representation of the identity.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewWebAppDomainOwnershipIdentifierSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, args *WebAppDomainOwnershipIdentifierSlotArgs, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppDomainOwnershipIdentifierSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.RegisterResource("azure-native:web/v20201201:WebAppDomainOwnershipIdentifierSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppDomainOwnershipIdentifierSlot gets an existing WebAppDomainOwnershipIdentifierSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDomainOwnershipIdentifierSlotState, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.ReadResource("azure-native:web/v20201201:WebAppDomainOwnershipIdentifierSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppDomainOwnershipIdentifierSlot resources.
type webAppDomainOwnershipIdentifierSlotState struct {
}

type WebAppDomainOwnershipIdentifierSlotState struct {
}

func (WebAppDomainOwnershipIdentifierSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierSlotState)(nil)).Elem()
}

type webAppDomainOwnershipIdentifierSlotArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName *string `pulumi:"domainOwnershipIdentifierName"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
	Slot string `pulumi:"slot"`
	// String representation of the identity.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a WebAppDomainOwnershipIdentifierSlot resource.
type WebAppDomainOwnershipIdentifierSlotArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
	Slot pulumi.StringInput
	// String representation of the identity.
	Value pulumi.StringPtrInput
}

func (WebAppDomainOwnershipIdentifierSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierSlotArgs)(nil)).Elem()
}

type WebAppDomainOwnershipIdentifierSlotInput interface {
	pulumi.Input

	ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput
	ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput
}

func (*WebAppDomainOwnershipIdentifierSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifierSlot)(nil)).Elem()
}

func (i *WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return i.ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(context.Background())
}

func (i *WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDomainOwnershipIdentifierSlotOutput)
}

type WebAppDomainOwnershipIdentifierSlotOutput struct{ *pulumi.OutputState }

func (WebAppDomainOwnershipIdentifierSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifierSlot)(nil)).Elem()
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

// Kind of resource.
func (o WebAppDomainOwnershipIdentifierSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppDomainOwnershipIdentifierSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppDomainOwnershipIdentifierSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// String representation of the identity.
func (o WebAppDomainOwnershipIdentifierSlotOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifierSlot) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppDomainOwnershipIdentifierSlotOutput{})
}