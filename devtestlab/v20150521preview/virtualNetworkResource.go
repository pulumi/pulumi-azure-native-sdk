// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A virtual network.
//
// Deprecated: Version 2015-05-21-preview will be removed in v2 of the provider.
type VirtualNetworkResource struct {
	pulumi.CustomResourceState

	// The allowed subnets of the virtual network.
	AllowedSubnets SubnetResponseArrayOutput `pulumi:"allowedSubnets"`
	// The description of the virtual network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId pulumi.StringPtrOutput `pulumi:"externalProviderResourceId"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The subnet overrides of the virtual network.
	SubnetOverrides SubnetOverrideResponseArrayOutput `pulumi:"subnetOverrides"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewVirtualNetworkResource registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkResource(ctx *pulumi.Context,
	name string, args *VirtualNetworkResourceArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualNetworkResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualNetworkResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualNetworkResource"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:VirtualNetworkResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkResource gets an existing VirtualNetworkResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkResourceState, opts ...pulumi.ResourceOption) (*VirtualNetworkResource, error) {
	var resource VirtualNetworkResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:VirtualNetworkResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkResource resources.
type virtualNetworkResourceState struct {
}

type VirtualNetworkResourceState struct {
}

func (VirtualNetworkResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkResourceState)(nil)).Elem()
}

type virtualNetworkResourceArgs struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets []Subnet `pulumi:"allowedSubnets"`
	// The description of the virtual network.
	Description *string `pulumi:"description"`
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId *string `pulumi:"externalProviderResourceId"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subnet overrides of the virtual network.
	SubnetOverrides []SubnetOverride `pulumi:"subnetOverrides"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a VirtualNetworkResource resource.
type VirtualNetworkResourceArgs struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets SubnetArrayInput
	// The description of the virtual network.
	Description pulumi.StringPtrInput
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId pulumi.StringPtrInput
	// The identifier of the resource.
	Id pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The subnet overrides of the virtual network.
	SubnetOverrides SubnetOverrideArrayInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (VirtualNetworkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkResourceArgs)(nil)).Elem()
}

type VirtualNetworkResourceInput interface {
	pulumi.Input

	ToVirtualNetworkResourceOutput() VirtualNetworkResourceOutput
	ToVirtualNetworkResourceOutputWithContext(ctx context.Context) VirtualNetworkResourceOutput
}

func (*VirtualNetworkResource) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkResource)(nil)).Elem()
}

func (i *VirtualNetworkResource) ToVirtualNetworkResourceOutput() VirtualNetworkResourceOutput {
	return i.ToVirtualNetworkResourceOutputWithContext(context.Background())
}

func (i *VirtualNetworkResource) ToVirtualNetworkResourceOutputWithContext(ctx context.Context) VirtualNetworkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResourceOutput)
}

type VirtualNetworkResourceOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkResource)(nil)).Elem()
}

func (o VirtualNetworkResourceOutput) ToVirtualNetworkResourceOutput() VirtualNetworkResourceOutput {
	return o
}

func (o VirtualNetworkResourceOutput) ToVirtualNetworkResourceOutputWithContext(ctx context.Context) VirtualNetworkResourceOutput {
	return o
}

// The allowed subnets of the virtual network.
func (o VirtualNetworkResourceOutput) AllowedSubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) SubnetResponseArrayOutput { return v.AllowedSubnets }).(SubnetResponseArrayOutput)
}

// The description of the virtual network.
func (o VirtualNetworkResourceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Microsoft.Network resource identifier of the virtual network.
func (o VirtualNetworkResourceOutput) ExternalProviderResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.ExternalProviderResourceId }).(pulumi.StringPtrOutput)
}

// The location of the resource.
func (o VirtualNetworkResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o VirtualNetworkResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o VirtualNetworkResourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The subnet overrides of the virtual network.
func (o VirtualNetworkResourceOutput) SubnetOverrides() SubnetOverrideResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) SubnetOverrideResponseArrayOutput { return v.SubnetOverrides }).(SubnetOverrideResponseArrayOutput)
}

// The tags of the resource.
func (o VirtualNetworkResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o VirtualNetworkResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkResourceOutput{})
}