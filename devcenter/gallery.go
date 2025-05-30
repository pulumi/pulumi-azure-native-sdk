// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a gallery.
//
// Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
//
// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Gallery struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The resource ID of the backing Azure Compute Gallery.
	GalleryResourceId pulumi.StringOutput `pulumi:"galleryResourceId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGallery registers a new resource with the given unique name, arguments, and options.
func NewGallery(ctx *pulumi.Context,
	name string, args *GalleryArgs, opts ...pulumi.ResourceOption) (*Gallery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.GalleryResourceId == nil {
		return nil, errors.New("invalid value for required argument 'GalleryResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230401:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230801preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240201:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240501preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240601preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240701preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240801preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20241001preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20250201:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20250401preview:Gallery"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Gallery
	err := ctx.RegisterResource("azure-native:devcenter:Gallery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGallery gets an existing Gallery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGallery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryState, opts ...pulumi.ResourceOption) (*Gallery, error) {
	var resource Gallery
	err := ctx.ReadResource("azure-native:devcenter:Gallery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gallery resources.
type galleryState struct {
}

type GalleryState struct {
}

func (GalleryState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryState)(nil)).Elem()
}

type galleryArgs struct {
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the gallery.
	GalleryName *string `pulumi:"galleryName"`
	// The resource ID of the backing Azure Compute Gallery.
	GalleryResourceId string `pulumi:"galleryResourceId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Gallery resource.
type GalleryArgs struct {
	// The name of the devcenter.
	DevCenterName pulumi.StringInput
	// The name of the gallery.
	GalleryName pulumi.StringPtrInput
	// The resource ID of the backing Azure Compute Gallery.
	GalleryResourceId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (GalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryArgs)(nil)).Elem()
}

type GalleryInput interface {
	pulumi.Input

	ToGalleryOutput() GalleryOutput
	ToGalleryOutputWithContext(ctx context.Context) GalleryOutput
}

func (*Gallery) ElementType() reflect.Type {
	return reflect.TypeOf((**Gallery)(nil)).Elem()
}

func (i *Gallery) ToGalleryOutput() GalleryOutput {
	return i.ToGalleryOutputWithContext(context.Background())
}

func (i *Gallery) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOutput)
}

type GalleryOutput struct{ *pulumi.OutputState }

func (GalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gallery)(nil)).Elem()
}

func (o GalleryOutput) ToGalleryOutput() GalleryOutput {
	return o
}

func (o GalleryOutput) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return o
}

// The Azure API version of the resource.
func (o GalleryOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The resource ID of the backing Azure Compute Gallery.
func (o GalleryOutput) GalleryResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.GalleryResourceId }).(pulumi.StringOutput)
}

// The name of the resource
func (o GalleryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o GalleryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GalleryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Gallery) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GalleryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryOutput{})
}
