// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
//
// Deprecated: Version 2019-05-01-preview will be removed in v2 of the provider.
type VirtualMachineImageTemplate struct {
	pulumi.CustomResourceState

	// Maximum duration to wait while building the image template. Omit or specify 0 to use the default (4 hours).
	BuildTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"buildTimeoutInMinutes"`
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize pulumi.ArrayOutput `pulumi:"customize"`
	// The distribution targets where the image output needs to go to.
	Distribute pulumi.ArrayOutput `pulumi:"distribute"`
	// The identity of the image template, if configured.
	Identity ImageTemplateIdentityResponsePtrOutput `pulumi:"identity"`
	// State of 'run' that is currently executing or was last executed.
	LastRunStatus ImageTemplateLastRunStatusResponseOutput `pulumi:"lastRunStatus"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning error, if any
	ProvisioningError ProvisioningErrorResponseOutput `pulumi:"provisioningError"`
	// Provisioning state of the resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies the properties used to describe the source image.
	Source pulumi.AnyOutput `pulumi:"source"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Describes how virtual machine is set up to build images
	VmProfile ImageTemplateVmProfileResponsePtrOutput `pulumi:"vmProfile"`
}

// NewVirtualMachineImageTemplate registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineImageTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distribute == nil {
		return nil, errors.New("invalid value for required argument 'Distribute'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:virtualmachineimages:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20190201preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20200214:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20211001:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20220214:VirtualMachineImageTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineImageTemplate
	err := ctx.RegisterResource("azure-native:virtualmachineimages/v20190501preview:VirtualMachineImageTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineImageTemplate gets an existing VirtualMachineImageTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineImageTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	var resource VirtualMachineImageTemplate
	err := ctx.ReadResource("azure-native:virtualmachineimages/v20190501preview:VirtualMachineImageTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineImageTemplate resources.
type virtualMachineImageTemplateState struct {
}

type VirtualMachineImageTemplateState struct {
}

func (VirtualMachineImageTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateState)(nil)).Elem()
}

type virtualMachineImageTemplateArgs struct {
	// Maximum duration to wait while building the image template. Omit or specify 0 to use the default (4 hours).
	BuildTimeoutInMinutes *int `pulumi:"buildTimeoutInMinutes"`
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize []interface{} `pulumi:"customize"`
	// The distribution targets where the image output needs to go to.
	Distribute []interface{} `pulumi:"distribute"`
	// The identity of the image template, if configured.
	Identity *ImageTemplateIdentity `pulumi:"identity"`
	// The name of the image Template
	ImageTemplateName *string `pulumi:"imageTemplateName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the properties used to describe the source image.
	Source interface{} `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Describes how virtual machine is set up to build images
	VmProfile *ImageTemplateVmProfile `pulumi:"vmProfile"`
}

// The set of arguments for constructing a VirtualMachineImageTemplate resource.
type VirtualMachineImageTemplateArgs struct {
	// Maximum duration to wait while building the image template. Omit or specify 0 to use the default (4 hours).
	BuildTimeoutInMinutes pulumi.IntPtrInput
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize pulumi.ArrayInput
	// The distribution targets where the image output needs to go to.
	Distribute pulumi.ArrayInput
	// The identity of the image template, if configured.
	Identity ImageTemplateIdentityPtrInput
	// The name of the image Template
	ImageTemplateName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the properties used to describe the source image.
	Source pulumi.Input
	// Resource tags
	Tags pulumi.StringMapInput
	// Describes how virtual machine is set up to build images
	VmProfile ImageTemplateVmProfilePtrInput
}

func (VirtualMachineImageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateArgs)(nil)).Elem()
}

type VirtualMachineImageTemplateInput interface {
	pulumi.Input

	ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput
	ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput
}

func (*VirtualMachineImageTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageTemplate)(nil)).Elem()
}

func (i *VirtualMachineImageTemplate) ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput {
	return i.ToVirtualMachineImageTemplateOutputWithContext(context.Background())
}

func (i *VirtualMachineImageTemplate) ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageTemplateOutput)
}

type VirtualMachineImageTemplateOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageTemplate)(nil)).Elem()
}

func (o VirtualMachineImageTemplateOutput) ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput {
	return o
}

func (o VirtualMachineImageTemplateOutput) ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput {
	return o
}

// Maximum duration to wait while building the image template. Omit or specify 0 to use the default (4 hours).
func (o VirtualMachineImageTemplateOutput) BuildTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.IntPtrOutput { return v.BuildTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// Specifies the properties used to describe the customization steps of the image, like Image source etc
func (o VirtualMachineImageTemplateOutput) Customize() pulumi.ArrayOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.ArrayOutput { return v.Customize }).(pulumi.ArrayOutput)
}

// The distribution targets where the image output needs to go to.
func (o VirtualMachineImageTemplateOutput) Distribute() pulumi.ArrayOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.ArrayOutput { return v.Distribute }).(pulumi.ArrayOutput)
}

// The identity of the image template, if configured.
func (o VirtualMachineImageTemplateOutput) Identity() ImageTemplateIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ImageTemplateIdentityResponsePtrOutput { return v.Identity }).(ImageTemplateIdentityResponsePtrOutput)
}

// State of 'run' that is currently executing or was last executed.
func (o VirtualMachineImageTemplateOutput) LastRunStatus() ImageTemplateLastRunStatusResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ImageTemplateLastRunStatusResponseOutput { return v.LastRunStatus }).(ImageTemplateLastRunStatusResponseOutput)
}

// Resource location
func (o VirtualMachineImageTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o VirtualMachineImageTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning error, if any
func (o VirtualMachineImageTemplateOutput) ProvisioningError() ProvisioningErrorResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ProvisioningErrorResponseOutput { return v.ProvisioningError }).(ProvisioningErrorResponseOutput)
}

// Provisioning state of the resource
func (o VirtualMachineImageTemplateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies the properties used to describe the source image.
func (o VirtualMachineImageTemplateOutput) Source() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.AnyOutput { return v.Source }).(pulumi.AnyOutput)
}

// Resource tags
func (o VirtualMachineImageTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o VirtualMachineImageTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Describes how virtual machine is set up to build images
func (o VirtualMachineImageTemplateOutput) VmProfile() ImageTemplateVmProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ImageTemplateVmProfileResponsePtrOutput { return v.VmProfile }).(ImageTemplateVmProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineImageTemplateOutput{})
}