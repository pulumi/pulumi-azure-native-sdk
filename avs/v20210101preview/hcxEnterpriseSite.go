// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An HCX Enterprise Site resource
type HcxEnterpriseSite struct {
	pulumi.CustomResourceState

	// The activation key
	ActivationKey pulumi.StringOutput `pulumi:"activationKey"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the HCX Enterprise Site
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHcxEnterpriseSite registers a new resource with the given unique name, arguments, and options.
func NewHcxEnterpriseSite(ctx *pulumi.Context,
	name string, args *HcxEnterpriseSiteArgs, opts ...pulumi.ResourceOption) (*HcxEnterpriseSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200320:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:HcxEnterpriseSite"),
		},
	})
	opts = append(opts, aliases)
	var resource HcxEnterpriseSite
	err := ctx.RegisterResource("azure-native:avs/v20210101preview:HcxEnterpriseSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHcxEnterpriseSite gets an existing HcxEnterpriseSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHcxEnterpriseSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HcxEnterpriseSiteState, opts ...pulumi.ResourceOption) (*HcxEnterpriseSite, error) {
	var resource HcxEnterpriseSite
	err := ctx.ReadResource("azure-native:avs/v20210101preview:HcxEnterpriseSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HcxEnterpriseSite resources.
type hcxEnterpriseSiteState struct {
}

type HcxEnterpriseSiteState struct {
}

func (HcxEnterpriseSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*hcxEnterpriseSiteState)(nil)).Elem()
}

type hcxEnterpriseSiteArgs struct {
	// Name of the HCX Enterprise Site in the private cloud
	HcxEnterpriseSiteName *string `pulumi:"hcxEnterpriseSiteName"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a HcxEnterpriseSite resource.
type HcxEnterpriseSiteArgs struct {
	// Name of the HCX Enterprise Site in the private cloud
	HcxEnterpriseSiteName pulumi.StringPtrInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (HcxEnterpriseSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hcxEnterpriseSiteArgs)(nil)).Elem()
}

type HcxEnterpriseSiteInput interface {
	pulumi.Input

	ToHcxEnterpriseSiteOutput() HcxEnterpriseSiteOutput
	ToHcxEnterpriseSiteOutputWithContext(ctx context.Context) HcxEnterpriseSiteOutput
}

func (*HcxEnterpriseSite) ElementType() reflect.Type {
	return reflect.TypeOf((**HcxEnterpriseSite)(nil)).Elem()
}

func (i *HcxEnterpriseSite) ToHcxEnterpriseSiteOutput() HcxEnterpriseSiteOutput {
	return i.ToHcxEnterpriseSiteOutputWithContext(context.Background())
}

func (i *HcxEnterpriseSite) ToHcxEnterpriseSiteOutputWithContext(ctx context.Context) HcxEnterpriseSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HcxEnterpriseSiteOutput)
}

type HcxEnterpriseSiteOutput struct{ *pulumi.OutputState }

func (HcxEnterpriseSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HcxEnterpriseSite)(nil)).Elem()
}

func (o HcxEnterpriseSiteOutput) ToHcxEnterpriseSiteOutput() HcxEnterpriseSiteOutput {
	return o
}

func (o HcxEnterpriseSiteOutput) ToHcxEnterpriseSiteOutputWithContext(ctx context.Context) HcxEnterpriseSiteOutput {
	return o
}

// The activation key
func (o HcxEnterpriseSiteOutput) ActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.ActivationKey }).(pulumi.StringOutput)
}

// Resource name.
func (o HcxEnterpriseSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the HCX Enterprise Site
func (o HcxEnterpriseSiteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o HcxEnterpriseSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HcxEnterpriseSiteOutput{})
}