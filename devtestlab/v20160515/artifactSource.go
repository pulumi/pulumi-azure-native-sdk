// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties of an artifact source.
//
// Deprecated: Version 2016-05-15 will be removed in v2 of the provider.
type ArtifactSource struct {
	pulumi.CustomResourceState

	// The folder containing Azure Resource Manager templates.
	ArmTemplateFolderPath pulumi.StringPtrOutput `pulumi:"armTemplateFolderPath"`
	// The artifact source's branch reference.
	BranchRef pulumi.StringPtrOutput `pulumi:"branchRef"`
	// The artifact source's creation date.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The artifact source's display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The folder containing artifacts.
	FolderPath pulumi.StringPtrOutput `pulumi:"folderPath"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The security token to authenticate to the artifact source.
	SecurityToken pulumi.StringPtrOutput `pulumi:"securityToken"`
	// The artifact source's type.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// Indicates if the artifact source is enabled (values: Enabled, Disabled).
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
	// The artifact source's URI.
	Uri pulumi.StringPtrOutput `pulumi:"uri"`
}

// NewArtifactSource registers a new resource with the given unique name, arguments, and options.
func NewArtifactSource(ctx *pulumi.Context,
	name string, args *ArtifactSourceArgs, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
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
			Type: pulumi.String("azure-native:devtestlab:ArtifactSource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:ArtifactSource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ArtifactSource"),
		},
	})
	opts = append(opts, aliases)
	var resource ArtifactSource
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:ArtifactSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactSource gets an existing ArtifactSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactSourceState, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	var resource ArtifactSource
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:ArtifactSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactSource resources.
type artifactSourceState struct {
}

type ArtifactSourceState struct {
}

func (ArtifactSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceState)(nil)).Elem()
}

type artifactSourceArgs struct {
	// The folder containing Azure Resource Manager templates.
	ArmTemplateFolderPath *string `pulumi:"armTemplateFolderPath"`
	// The artifact source's branch reference.
	BranchRef *string `pulumi:"branchRef"`
	// The artifact source's display name.
	DisplayName *string `pulumi:"displayName"`
	// The folder containing artifacts.
	FolderPath *string `pulumi:"folderPath"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the artifact source.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security token to authenticate to the artifact source.
	SecurityToken *string `pulumi:"securityToken"`
	// The artifact source's type.
	SourceType *string `pulumi:"sourceType"`
	// Indicates if the artifact source is enabled (values: Enabled, Disabled).
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// The artifact source's URI.
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a ArtifactSource resource.
type ArtifactSourceArgs struct {
	// The folder containing Azure Resource Manager templates.
	ArmTemplateFolderPath pulumi.StringPtrInput
	// The artifact source's branch reference.
	BranchRef pulumi.StringPtrInput
	// The artifact source's display name.
	DisplayName pulumi.StringPtrInput
	// The folder containing artifacts.
	FolderPath pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the artifact source.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The security token to authenticate to the artifact source.
	SecurityToken pulumi.StringPtrInput
	// The artifact source's type.
	SourceType pulumi.StringPtrInput
	// Indicates if the artifact source is enabled (values: Enabled, Disabled).
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
	// The artifact source's URI.
	Uri pulumi.StringPtrInput
}

func (ArtifactSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceArgs)(nil)).Elem()
}

type ArtifactSourceInput interface {
	pulumi.Input

	ToArtifactSourceOutput() ArtifactSourceOutput
	ToArtifactSourceOutputWithContext(ctx context.Context) ArtifactSourceOutput
}

func (*ArtifactSource) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactSource)(nil)).Elem()
}

func (i *ArtifactSource) ToArtifactSourceOutput() ArtifactSourceOutput {
	return i.ToArtifactSourceOutputWithContext(context.Background())
}

func (i *ArtifactSource) ToArtifactSourceOutputWithContext(ctx context.Context) ArtifactSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactSourceOutput)
}

type ArtifactSourceOutput struct{ *pulumi.OutputState }

func (ArtifactSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactSource)(nil)).Elem()
}

func (o ArtifactSourceOutput) ToArtifactSourceOutput() ArtifactSourceOutput {
	return o
}

func (o ArtifactSourceOutput) ToArtifactSourceOutputWithContext(ctx context.Context) ArtifactSourceOutput {
	return o
}

// The folder containing Azure Resource Manager templates.
func (o ArtifactSourceOutput) ArmTemplateFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.ArmTemplateFolderPath }).(pulumi.StringPtrOutput)
}

// The artifact source's branch reference.
func (o ArtifactSourceOutput) BranchRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.BranchRef }).(pulumi.StringPtrOutput)
}

// The artifact source's creation date.
func (o ArtifactSourceOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The artifact source's display name.
func (o ArtifactSourceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The folder containing artifacts.
func (o ArtifactSourceOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.FolderPath }).(pulumi.StringPtrOutput)
}

// The location of the resource.
func (o ArtifactSourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ArtifactSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o ArtifactSourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The security token to authenticate to the artifact source.
func (o ArtifactSourceOutput) SecurityToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.SecurityToken }).(pulumi.StringPtrOutput)
}

// The artifact source's type.
func (o ArtifactSourceOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

// Indicates if the artifact source is enabled (values: Enabled, Disabled).
func (o ArtifactSourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o ArtifactSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o ArtifactSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o ArtifactSourceOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

// The artifact source's URI.
func (o ArtifactSourceOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ArtifactSourceOutput{})
}