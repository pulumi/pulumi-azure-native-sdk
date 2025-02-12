// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Template Spec Version object.
type TemplateSpecVersion struct {
	pulumi.CustomResourceState

	// An array of Template Spec artifacts.
	Artifacts TemplateSpecTemplateArtifactResponseArrayOutput `pulumi:"artifacts"`
	// Template Spec version description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Azure Resource Manager template content.
	Template pulumi.AnyOutput `pulumi:"template"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTemplateSpecVersion registers a new resource with the given unique name, arguments, and options.
func NewTemplateSpecVersion(ctx *pulumi.Context,
	name string, args *TemplateSpecVersionArgs, opts ...pulumi.ResourceOption) (*TemplateSpecVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateSpecName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateSpecName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20210301preview:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210501:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220201:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources:TemplateSpecVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TemplateSpecVersion
	err := ctx.RegisterResource("azure-native:resources/v20190601preview:TemplateSpecVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateSpecVersion gets an existing TemplateSpecVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateSpecVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateSpecVersionState, opts ...pulumi.ResourceOption) (*TemplateSpecVersion, error) {
	var resource TemplateSpecVersion
	err := ctx.ReadResource("azure-native:resources/v20190601preview:TemplateSpecVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateSpecVersion resources.
type templateSpecVersionState struct {
}

type TemplateSpecVersionState struct {
}

func (TemplateSpecVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecVersionState)(nil)).Elem()
}

type templateSpecVersionArgs struct {
	// An array of Template Spec artifacts.
	Artifacts []TemplateSpecTemplateArtifact `pulumi:"artifacts"`
	// Template Spec version description.
	Description *string `pulumi:"description"`
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Resource Manager template content.
	Template interface{} `pulumi:"template"`
	// Name of the Template Spec.
	TemplateSpecName string `pulumi:"templateSpecName"`
	// The version of the Template Spec.
	TemplateSpecVersion *string `pulumi:"templateSpecVersion"`
}

// The set of arguments for constructing a TemplateSpecVersion resource.
type TemplateSpecVersionArgs struct {
	// An array of Template Spec artifacts.
	Artifacts TemplateSpecTemplateArtifactArrayInput
	// Template Spec version description.
	Description pulumi.StringPtrInput
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Azure Resource Manager template content.
	Template pulumi.Input
	// Name of the Template Spec.
	TemplateSpecName pulumi.StringInput
	// The version of the Template Spec.
	TemplateSpecVersion pulumi.StringPtrInput
}

func (TemplateSpecVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecVersionArgs)(nil)).Elem()
}

type TemplateSpecVersionInput interface {
	pulumi.Input

	ToTemplateSpecVersionOutput() TemplateSpecVersionOutput
	ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput
}

func (*TemplateSpecVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecVersion)(nil)).Elem()
}

func (i *TemplateSpecVersion) ToTemplateSpecVersionOutput() TemplateSpecVersionOutput {
	return i.ToTemplateSpecVersionOutputWithContext(context.Background())
}

func (i *TemplateSpecVersion) ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecVersionOutput)
}

type TemplateSpecVersionOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecVersion)(nil)).Elem()
}

func (o TemplateSpecVersionOutput) ToTemplateSpecVersionOutput() TemplateSpecVersionOutput {
	return o
}

func (o TemplateSpecVersionOutput) ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput {
	return o
}

// An array of Template Spec artifacts.
func (o TemplateSpecVersionOutput) Artifacts() TemplateSpecTemplateArtifactResponseArrayOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) TemplateSpecTemplateArtifactResponseArrayOutput { return v.Artifacts }).(TemplateSpecTemplateArtifactResponseArrayOutput)
}

// Template Spec version description.
func (o TemplateSpecVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The location of the Template Spec Version. It must match the location of the parent Template Spec.
func (o TemplateSpecVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of this resource.
func (o TemplateSpecVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TemplateSpecVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o TemplateSpecVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The Azure Resource Manager template content.
func (o TemplateSpecVersionOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.Template }).(pulumi.AnyOutput)
}

// Type of this resource.
func (o TemplateSpecVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateSpecVersionOutput{})
}
