// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies information about the gallery Application Definition that you want to create or update.
//
// Deprecated: Version 2019-12-01 will be removed in v2 of the provider.
type GalleryApplication struct {
	pulumi.CustomResourceState

	// The description of this gallery Application Definition resource. This property is updatable.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrOutput `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery Application Definition.
	Eula pulumi.StringPtrOutput `pulumi:"eula"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrOutput `pulumi:"privacyStatementUri"`
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrOutput `pulumi:"releaseNoteUri"`
	// This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	SupportedOSType pulumi.StringOutput `pulumi:"supportedOSType"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryApplication registers a new resource with the given unique name, arguments, and options.
func NewGalleryApplication(ctx *pulumi.Context,
	name string, args *GalleryApplicationArgs, opts ...pulumi.ResourceOption) (*GalleryApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SupportedOSType == nil {
		return nil, errors.New("invalid value for required argument 'SupportedOSType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211001:GalleryApplication"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220103:GalleryApplication"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryApplication
	err := ctx.RegisterResource("azure-native:compute/v20191201:GalleryApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryApplication gets an existing GalleryApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryApplicationState, opts ...pulumi.ResourceOption) (*GalleryApplication, error) {
	var resource GalleryApplication
	err := ctx.ReadResource("azure-native:compute/v20191201:GalleryApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryApplication resources.
type galleryApplicationState struct {
}

type GalleryApplicationState struct {
}

func (GalleryApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationState)(nil)).Elem()
}

type galleryApplicationArgs struct {
	// The description of this gallery Application Definition resource. This property is updatable.
	Description *string `pulumi:"description"`
	// The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate *string `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery Application Definition.
	Eula *string `pulumi:"eula"`
	// The name of the gallery Application Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
	GalleryApplicationName *string `pulumi:"galleryApplicationName"`
	// The name of the Shared Application Gallery in which the Application Definition is to be created.
	GalleryName string `pulumi:"galleryName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The privacy statement uri.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// The release note uri.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	SupportedOSType OperatingSystemTypes `pulumi:"supportedOSType"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryApplication resource.
type GalleryApplicationArgs struct {
	// The description of this gallery Application Definition resource. This property is updatable.
	Description pulumi.StringPtrInput
	// The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrInput
	// The Eula agreement for the gallery Application Definition.
	Eula pulumi.StringPtrInput
	// The name of the gallery Application Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
	GalleryApplicationName pulumi.StringPtrInput
	// The name of the Shared Application Gallery in which the Application Definition is to be created.
	GalleryName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrInput
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	SupportedOSType OperatingSystemTypesInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationArgs)(nil)).Elem()
}

type GalleryApplicationInput interface {
	pulumi.Input

	ToGalleryApplicationOutput() GalleryApplicationOutput
	ToGalleryApplicationOutputWithContext(ctx context.Context) GalleryApplicationOutput
}

func (*GalleryApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryApplication)(nil)).Elem()
}

func (i *GalleryApplication) ToGalleryApplicationOutput() GalleryApplicationOutput {
	return i.ToGalleryApplicationOutputWithContext(context.Background())
}

func (i *GalleryApplication) ToGalleryApplicationOutputWithContext(ctx context.Context) GalleryApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryApplicationOutput)
}

type GalleryApplicationOutput struct{ *pulumi.OutputState }

func (GalleryApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryApplication)(nil)).Elem()
}

func (o GalleryApplicationOutput) ToGalleryApplicationOutput() GalleryApplicationOutput {
	return o
}

func (o GalleryApplicationOutput) ToGalleryApplicationOutputWithContext(ctx context.Context) GalleryApplicationOutput {
	return o
}

// The description of this gallery Application Definition resource. This property is updatable.
func (o GalleryApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
func (o GalleryApplicationOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

// The Eula agreement for the gallery Application Definition.
func (o GalleryApplicationOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.Eula }).(pulumi.StringPtrOutput)
}

// Resource location
func (o GalleryApplicationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o GalleryApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The privacy statement uri.
func (o GalleryApplicationOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

// The release note uri.
func (o GalleryApplicationOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringPtrOutput { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

// This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
func (o GalleryApplicationOutput) SupportedOSType() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.SupportedOSType }).(pulumi.StringOutput)
}

// Resource tags
func (o GalleryApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o GalleryApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplication) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryApplicationOutput{})
}