// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceGuard struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Input Managed Identity Details
	Identity DppIdentityDetailsResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ResourceGuardResource properties
	Properties ResourceGuardResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceGuard registers a new resource with the given unique name, arguments, and options.
func NewResourceGuard(ctx *pulumi.Context,
	name string, args *ResourceGuardArgs, opts ...pulumi.ResourceOption) (*ResourceGuard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dataprotection:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210701:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220501:ResourceGuard"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGuard
	err := ctx.RegisterResource("azure-native:dataprotection/v20220101:ResourceGuard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGuard gets an existing ResourceGuard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGuard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGuardState, opts ...pulumi.ResourceOption) (*ResourceGuard, error) {
	var resource ResourceGuard
	err := ctx.ReadResource("azure-native:dataprotection/v20220101:ResourceGuard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGuard resources.
type resourceGuardState struct {
}

type ResourceGuardState struct {
}

func (ResourceGuardState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardState)(nil)).Elem()
}

type resourceGuardArgs struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Input Managed Identity Details
	Identity *DppIdentityDetails `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group where the backup vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of ResourceGuard
	ResourceGuardsName *string `pulumi:"resourceGuardsName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceGuard resource.
type ResourceGuardArgs struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Input Managed Identity Details
	Identity DppIdentityDetailsPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group where the backup vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of ResourceGuard
	ResourceGuardsName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ResourceGuardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardArgs)(nil)).Elem()
}

type ResourceGuardInput interface {
	pulumi.Input

	ToResourceGuardOutput() ResourceGuardOutput
	ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput
}

func (*ResourceGuard) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuard)(nil)).Elem()
}

func (i *ResourceGuard) ToResourceGuardOutput() ResourceGuardOutput {
	return i.ToResourceGuardOutputWithContext(context.Background())
}

func (i *ResourceGuard) ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOutput)
}

type ResourceGuardOutput struct{ *pulumi.OutputState }

func (ResourceGuardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuard)(nil)).Elem()
}

func (o ResourceGuardOutput) ToResourceGuardOutput() ResourceGuardOutput {
	return o
}

func (o ResourceGuardOutput) ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput {
	return o
}

// Optional ETag.
func (o ResourceGuardOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Input Managed Identity Details
func (o ResourceGuardOutput) Identity() DppIdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ResourceGuard) DppIdentityDetailsResponsePtrOutput { return v.Identity }).(DppIdentityDetailsResponsePtrOutput)
}

// Resource location.
func (o ResourceGuardOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o ResourceGuardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ResourceGuardResource properties
func (o ResourceGuardOutput) Properties() ResourceGuardResponseOutput {
	return o.ApplyT(func(v *ResourceGuard) ResourceGuardResponseOutput { return v.Properties }).(ResourceGuardResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ResourceGuardOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ResourceGuard) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ResourceGuardOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o ResourceGuardOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGuardOutput{})
}