// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The view resource format.
// API Version: 2017-04-26.
type View struct {
	pulumi.CustomResourceState

	// Date time when view was last modified.
	Changed pulumi.StringOutput `pulumi:"changed"`
	// Date time when view was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// View definition.
	Definition pulumi.StringOutput `pulumi:"definition"`
	// Localized display name for the view.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// the hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// the user ID.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
	// Name of the view.
	ViewName pulumi.StringOutput `pulumi:"viewName"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:View"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:View"),
		},
	})
	opts = append(opts, aliases)
	var resource View
	err := ctx.RegisterResource("azure-native:customerinsights:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetView gets an existing View resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("azure-native:customerinsights:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering View resources.
type viewState struct {
}

type ViewState struct {
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	// View definition.
	Definition string `pulumi:"definition"`
	// Localized display name for the view.
	DisplayName map[string]string `pulumi:"displayName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// the user ID.
	UserId *string `pulumi:"userId"`
	// The name of the view.
	ViewName *string `pulumi:"viewName"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	// View definition.
	Definition pulumi.StringInput
	// Localized display name for the view.
	DisplayName pulumi.StringMapInput
	// The name of the hub.
	HubName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// the user ID.
	UserId pulumi.StringPtrInput
	// The name of the view.
	ViewName pulumi.StringPtrInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

// Date time when view was last modified.
func (o ViewOutput) Changed() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Changed }).(pulumi.StringOutput)
}

// Date time when view was created.
func (o ViewOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// View definition.
func (o ViewOutput) Definition() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Definition }).(pulumi.StringOutput)
}

// Localized display name for the view.
func (o ViewOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *View) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

// Resource name.
func (o ViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the hub name.
func (o ViewOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o ViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// the user ID.
func (o ViewOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

// Name of the view.
func (o ViewOutput) ViewName() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ViewName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ViewOutput{})
}