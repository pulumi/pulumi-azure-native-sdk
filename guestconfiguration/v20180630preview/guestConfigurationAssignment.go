// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180630preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Guest configuration assignment is an association between a VM and guest configuration.
//
// Deprecated: Version 2018-06-30-preview will be removed in v2 of the provider.
type GuestConfigurationAssignment struct {
	pulumi.CustomResourceState

	// Region where the VM is located.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestConfigurationAssignment registers a new resource with the given unique name, arguments, and options.
func NewGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToGuestConfigurationAssignmentPropertiesPtrOutput().ApplyT(func(v *GuestConfigurationAssignmentProperties) *GuestConfigurationAssignmentProperties {
			return v.Defaults()
		}).(GuestConfigurationAssignmentPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:guestconfiguration:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20181120:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20200625:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20210125:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20220125:GuestConfigurationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestConfigurationAssignment
	err := ctx.RegisterResource("azure-native:guestconfiguration/v20180630preview:GuestConfigurationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestConfigurationAssignment gets an existing GuestConfigurationAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	var resource GuestConfigurationAssignment
	err := ctx.ReadResource("azure-native:guestconfiguration/v20180630preview:GuestConfigurationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestConfigurationAssignment resources.
type guestConfigurationAssignmentState struct {
}

type GuestConfigurationAssignmentState struct {
}

func (GuestConfigurationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentState)(nil)).Elem()
}

type guestConfigurationAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName *string `pulumi:"guestConfigurationAssignmentName"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a GuestConfigurationAssignment resource.
type GuestConfigurationAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName pulumi.StringPtrInput
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the virtual machine.
	VmName pulumi.StringInput
}

func (GuestConfigurationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentArgs)(nil)).Elem()
}

type GuestConfigurationAssignmentInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput
	ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput
}

func (*GuestConfigurationAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignment)(nil)).Elem()
}

func (i *GuestConfigurationAssignment) ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput {
	return i.ToGuestConfigurationAssignmentOutputWithContext(context.Background())
}

func (i *GuestConfigurationAssignment) ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentOutput)
}

type GuestConfigurationAssignmentOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignment)(nil)).Elem()
}

func (o GuestConfigurationAssignmentOutput) ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput {
	return o
}

func (o GuestConfigurationAssignmentOutput) ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput {
	return o
}

// Region where the VM is located.
func (o GuestConfigurationAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the guest configuration assignment.
func (o GuestConfigurationAssignmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Properties of the Guest configuration assignment.
func (o GuestConfigurationAssignmentOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignment) GuestConfigurationAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

// The type of the resource.
func (o GuestConfigurationAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestConfigurationAssignmentOutput{})
}