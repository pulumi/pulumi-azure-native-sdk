// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The galleryimages resource definition.
type GalleryimageRetrieve struct {
	pulumi.CustomResourceState

	// Container Name for storage container
	ContainerName pulumi.StringPtrOutput `pulumi:"containerName"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// location of the image the gallery image should be created from
	ImagePath pulumi.StringPtrOutput `pulumi:"imagePath"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// Provisioning state of the gallery image.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// name of the object to be used in moc
	ResourceName pulumi.StringPtrOutput `pulumi:"resourceName"`
	// GalleryImageStatus defines the observed state of MOCGalleryImage
	Status GalleryImageStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryimageRetrieve registers a new resource with the given unique name, arguments, and options.
func NewGalleryimageRetrieve(ctx *pulumi.Context,
	name string, args *GalleryimageRetrieveArgs, opts ...pulumi.ResourceOption) (*GalleryimageRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource GalleryimageRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210701preview:galleryimageRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryimageRetrieve gets an existing GalleryimageRetrieve resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryimageRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryimageRetrieveState, opts ...pulumi.ResourceOption) (*GalleryimageRetrieve, error) {
	var resource GalleryimageRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210701preview:galleryimageRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryimageRetrieve resources.
type galleryimageRetrieveState struct {
}

type GalleryimageRetrieveState struct {
}

func (GalleryimageRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryimageRetrieveState)(nil)).Elem()
}

type galleryimageRetrieveArgs struct {
	// Container Name for storage container
	ContainerName *string `pulumi:"containerName"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Name of the gallery image
	GalleryimagesName *string `pulumi:"galleryimagesName"`
	// location of the image the gallery image should be created from
	ImagePath *string `pulumi:"imagePath"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType *OperatingSystemTypes `pulumi:"osType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// name of the object to be used in moc
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryimageRetrieve resource.
type GalleryimageRetrieveArgs struct {
	// Container Name for storage container
	ContainerName pulumi.StringPtrInput
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// Name of the gallery image
	GalleryimagesName pulumi.StringPtrInput
	// location of the image the gallery image should be created from
	ImagePath pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// operating system type that the gallery image uses. Expected to be linux or windows
	OsType OperatingSystemTypesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// name of the object to be used in moc
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (GalleryimageRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryimageRetrieveArgs)(nil)).Elem()
}

type GalleryimageRetrieveInput interface {
	pulumi.Input

	ToGalleryimageRetrieveOutput() GalleryimageRetrieveOutput
	ToGalleryimageRetrieveOutputWithContext(ctx context.Context) GalleryimageRetrieveOutput
}

func (*GalleryimageRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryimageRetrieve)(nil)).Elem()
}

func (i *GalleryimageRetrieve) ToGalleryimageRetrieveOutput() GalleryimageRetrieveOutput {
	return i.ToGalleryimageRetrieveOutputWithContext(context.Background())
}

func (i *GalleryimageRetrieve) ToGalleryimageRetrieveOutputWithContext(ctx context.Context) GalleryimageRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryimageRetrieveOutput)
}

type GalleryimageRetrieveOutput struct{ *pulumi.OutputState }

func (GalleryimageRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryimageRetrieve)(nil)).Elem()
}

func (o GalleryimageRetrieveOutput) ToGalleryimageRetrieveOutput() GalleryimageRetrieveOutput {
	return o
}

func (o GalleryimageRetrieveOutput) ToGalleryimageRetrieveOutputWithContext(ctx context.Context) GalleryimageRetrieveOutput {
	return o
}

// Container Name for storage container
func (o GalleryimageRetrieveOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.ContainerName }).(pulumi.StringPtrOutput)
}

// The extendedLocation of the resource.
func (o GalleryimageRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// location of the image the gallery image should be created from
func (o GalleryimageRetrieveOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o GalleryimageRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o GalleryimageRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// operating system type that the gallery image uses. Expected to be linux or windows
func (o GalleryimageRetrieveOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the gallery image.
func (o GalleryimageRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// name of the object to be used in moc
func (o GalleryimageRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// GalleryImageStatus defines the observed state of MOCGalleryImage
func (o GalleryimageRetrieveOutput) Status() GalleryImageStatusResponseOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) GalleryImageStatusResponseOutput { return v.Status }).(GalleryImageStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GalleryimageRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o GalleryimageRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GalleryimageRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryimageRetrieveOutput{})
}