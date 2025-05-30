// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Customer subscription which can use a sku.
//
// Uses Azure REST API version 2022-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-01-01-preview.
type VendorSkuPreview struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The preview subscription ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the PreviewSubscription resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVendorSkuPreview registers a new resource with the given unique name, arguments, and options.
func NewVendorSkuPreview(ctx *pulumi.Context,
	name string, args *VendorSkuPreviewArgs, opts ...pulumi.ResourceOption) (*VendorSkuPreview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.VendorName == nil {
		return nil, errors.New("invalid value for required argument 'VendorName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:VendorSkuPreview"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:VendorSkuPreview"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20220101preview:VendorSkuPreview"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VendorSkuPreview
	err := ctx.RegisterResource("azure-native:hybridnetwork:VendorSkuPreview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVendorSkuPreview gets an existing VendorSkuPreview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVendorSkuPreview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkuPreviewState, opts ...pulumi.ResourceOption) (*VendorSkuPreview, error) {
	var resource VendorSkuPreview
	err := ctx.ReadResource("azure-native:hybridnetwork:VendorSkuPreview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VendorSkuPreview resources.
type vendorSkuPreviewState struct {
}

type VendorSkuPreviewState struct {
}

func (VendorSkuPreviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkuPreviewState)(nil)).Elem()
}

type vendorSkuPreviewArgs struct {
	// Preview subscription ID.
	PreviewSubscription *string `pulumi:"previewSubscription"`
	// The name of the vendor sku.
	SkuName string `pulumi:"skuName"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// The set of arguments for constructing a VendorSkuPreview resource.
type VendorSkuPreviewArgs struct {
	// Preview subscription ID.
	PreviewSubscription pulumi.StringPtrInput
	// The name of the vendor sku.
	SkuName pulumi.StringInput
	// The name of the vendor.
	VendorName pulumi.StringInput
}

func (VendorSkuPreviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkuPreviewArgs)(nil)).Elem()
}

type VendorSkuPreviewInput interface {
	pulumi.Input

	ToVendorSkuPreviewOutput() VendorSkuPreviewOutput
	ToVendorSkuPreviewOutputWithContext(ctx context.Context) VendorSkuPreviewOutput
}

func (*VendorSkuPreview) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkuPreview)(nil)).Elem()
}

func (i *VendorSkuPreview) ToVendorSkuPreviewOutput() VendorSkuPreviewOutput {
	return i.ToVendorSkuPreviewOutputWithContext(context.Background())
}

func (i *VendorSkuPreview) ToVendorSkuPreviewOutputWithContext(ctx context.Context) VendorSkuPreviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VendorSkuPreviewOutput)
}

type VendorSkuPreviewOutput struct{ *pulumi.OutputState }

func (VendorSkuPreviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkuPreview)(nil)).Elem()
}

func (o VendorSkuPreviewOutput) ToVendorSkuPreviewOutput() VendorSkuPreviewOutput {
	return o
}

func (o VendorSkuPreviewOutput) ToVendorSkuPreviewOutputWithContext(ctx context.Context) VendorSkuPreviewOutput {
	return o
}

// The Azure API version of the resource.
func (o VendorSkuPreviewOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkuPreview) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The preview subscription ID.
func (o VendorSkuPreviewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkuPreview) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the PreviewSubscription resource.
func (o VendorSkuPreviewOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkuPreview) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o VendorSkuPreviewOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VendorSkuPreview) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o VendorSkuPreviewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkuPreview) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VendorSkuPreviewOutput{})
}
