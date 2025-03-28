// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an image from the Azure Marketplace
//
// Uses Azure REST API version 2018-10-15. In version 1.x of the Azure Native provider, it used API version 2018-10-15.
type GalleryImage struct {
	pulumi.CustomResourceState

	// The author of the gallery image.
	Author pulumi.StringOutput `pulumi:"author"`
	// The creation date of the gallery image.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the gallery image.
	Description pulumi.StringOutput `pulumi:"description"`
	// The icon of the gallery image.
	Icon pulumi.StringOutput `pulumi:"icon"`
	// The image reference of the gallery image.
	ImageReference GalleryImageReferenceResponseOutput `pulumi:"imageReference"`
	// Indicates whether this gallery image is enabled.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Indicates whether this gallery has been overridden for this lab account
	IsOverride pulumi.BoolPtrOutput `pulumi:"isOverride"`
	// Indicates if the plan has been authorized for programmatic deployment.
	IsPlanAuthorized pulumi.BoolPtrOutput `pulumi:"isPlanAuthorized"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The third party plan that applies to this image
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
}

// NewGalleryImage registers a new resource with the given unique name, arguments, and options.
func NewGalleryImage(ctx *pulumi.Context,
	name string, args *GalleryImageArgs, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20181015:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:labservices:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryImage gets an existing GalleryImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:labservices:GalleryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryImage resources.
type galleryImageState struct {
}

type GalleryImageState struct {
}

func (GalleryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageState)(nil)).Elem()
}

type galleryImageArgs struct {
	// The name of the gallery Image.
	GalleryImageName *string `pulumi:"galleryImageName"`
	// Indicates whether this gallery image is enabled.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Indicates whether this gallery has been overridden for this lab account
	IsOverride *bool `pulumi:"isOverride"`
	// Indicates if the plan has been authorized for programmatic deployment.
	IsPlanAuthorized *bool `pulumi:"isPlanAuthorized"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

// The set of arguments for constructing a GalleryImage resource.
type GalleryImageArgs struct {
	// The name of the gallery Image.
	GalleryImageName pulumi.StringPtrInput
	// Indicates whether this gallery image is enabled.
	IsEnabled pulumi.BoolPtrInput
	// Indicates whether this gallery has been overridden for this lab account
	IsOverride pulumi.BoolPtrInput
	// Indicates if the plan has been authorized for programmatic deployment.
	IsPlanAuthorized pulumi.BoolPtrInput
	// The name of the lab Account.
	LabAccountName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (GalleryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageArgs)(nil)).Elem()
}

type GalleryImageInput interface {
	pulumi.Input

	ToGalleryImageOutput() GalleryImageOutput
	ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput
}

func (*GalleryImage) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (i *GalleryImage) ToGalleryImageOutput() GalleryImageOutput {
	return i.ToGalleryImageOutputWithContext(context.Background())
}

func (i *GalleryImage) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageOutput)
}

type GalleryImageOutput struct{ *pulumi.OutputState }

func (GalleryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (o GalleryImageOutput) ToGalleryImageOutput() GalleryImageOutput {
	return o
}

func (o GalleryImageOutput) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return o
}

// The author of the gallery image.
func (o GalleryImageOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

// The creation date of the gallery image.
func (o GalleryImageOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The description of the gallery image.
func (o GalleryImageOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The icon of the gallery image.
func (o GalleryImageOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Icon }).(pulumi.StringOutput)
}

// The image reference of the gallery image.
func (o GalleryImageOutput) ImageReference() GalleryImageReferenceResponseOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageReferenceResponseOutput { return v.ImageReference }).(GalleryImageReferenceResponseOutput)
}

// Indicates whether this gallery image is enabled.
func (o GalleryImageOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Indicates whether this gallery has been overridden for this lab account
func (o GalleryImageOutput) IsOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.BoolPtrOutput { return v.IsOverride }).(pulumi.BoolPtrOutput)
}

// Indicates if the plan has been authorized for programmatic deployment.
func (o GalleryImageOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.BoolPtrOutput { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

// The details of the latest operation. ex: status, error
func (o GalleryImageOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *GalleryImage) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

// The location of the resource.
func (o GalleryImageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o GalleryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The third party plan that applies to this image
func (o GalleryImageOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o GalleryImageOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o GalleryImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o GalleryImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o GalleryImageOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryImageOutput{})
}
