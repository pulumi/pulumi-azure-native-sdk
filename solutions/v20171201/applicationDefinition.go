// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about managed application definition.
//
// Deprecated: Version 2017-12-01 will be removed in v2 of the provider.
type ApplicationDefinition struct {
	pulumi.CustomResourceState

	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts ApplicationArtifactResponseArrayOutput `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations ApplicationProviderAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition pulumi.AnyOutput `pulumi:"createUiDefinition"`
	// The managed application definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The managed application definition display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// A value indicating whether the package is enabled or not.
	IsEnabled pulumi.StringPtrOutput `pulumi:"isEnabled"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The managed application lock level.
	LockLevel pulumi.StringOutput `pulumi:"lockLevel"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate pulumi.AnyOutput `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri pulumi.StringPtrOutput `pulumi:"packageFileUri"`
	// The SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationDefinition registers a new resource with the given unique name, arguments, and options.
func NewApplicationDefinition(ctx *pulumi.Context,
	name string, args *ApplicationDefinitionArgs, opts ...pulumi.ResourceOption) (*ApplicationDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorizations == nil {
		return nil, errors.New("invalid value for required argument 'Authorizations'")
	}
	if args.LockLevel == nil {
		return nil, errors.New("invalid value for required argument 'LockLevel'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:solutions:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20160901preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20170901:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180201:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180301:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180601:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20180901preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20190701:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20200821preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210201preview:ApplicationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:solutions/v20210701:ApplicationDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationDefinition
	err := ctx.RegisterResource("azure-native:solutions/v20171201:ApplicationDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationDefinition gets an existing ApplicationDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationDefinitionState, opts ...pulumi.ResourceOption) (*ApplicationDefinition, error) {
	var resource ApplicationDefinition
	err := ctx.ReadResource("azure-native:solutions/v20171201:ApplicationDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationDefinition resources.
type applicationDefinitionState struct {
}

type ApplicationDefinitionState struct {
}

func (ApplicationDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationDefinitionState)(nil)).Elem()
}

type applicationDefinitionArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName *string `pulumi:"applicationDefinitionName"`
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts []ApplicationArtifact `pulumi:"artifacts"`
	// The managed application provider authorizations.
	Authorizations []ApplicationProviderAuthorization `pulumi:"authorizations"`
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition interface{} `pulumi:"createUiDefinition"`
	// The managed application definition description.
	Description *string `pulumi:"description"`
	// The managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// A value indicating whether the package is enabled or not.
	IsEnabled *string `pulumi:"isEnabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// The managed application lock level.
	LockLevel ApplicationLockLevel `pulumi:"lockLevel"`
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The managed application definition package file Uri. Use this element
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationDefinition resource.
type ApplicationDefinitionArgs struct {
	// The name of the managed application definition.
	ApplicationDefinitionName pulumi.StringPtrInput
	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
	Artifacts ApplicationArtifactArrayInput
	// The managed application provider authorizations.
	Authorizations ApplicationProviderAuthorizationArrayInput
	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUiDefinition pulumi.Input
	// The managed application definition description.
	Description pulumi.StringPtrInput
	// The managed application definition display name.
	DisplayName pulumi.StringPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// A value indicating whether the package is enabled or not.
	IsEnabled pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The managed application lock level.
	LockLevel ApplicationLockLevelInput
	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate pulumi.Input
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrInput
	// The managed application definition package file Uri. Use this element
	PackageFileUri pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of the resource.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ApplicationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationDefinitionArgs)(nil)).Elem()
}

type ApplicationDefinitionInput interface {
	pulumi.Input

	ToApplicationDefinitionOutput() ApplicationDefinitionOutput
	ToApplicationDefinitionOutputWithContext(ctx context.Context) ApplicationDefinitionOutput
}

func (*ApplicationDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationDefinition)(nil)).Elem()
}

func (i *ApplicationDefinition) ToApplicationDefinitionOutput() ApplicationDefinitionOutput {
	return i.ToApplicationDefinitionOutputWithContext(context.Background())
}

func (i *ApplicationDefinition) ToApplicationDefinitionOutputWithContext(ctx context.Context) ApplicationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionOutput)
}

type ApplicationDefinitionOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationDefinition)(nil)).Elem()
}

func (o ApplicationDefinitionOutput) ToApplicationDefinitionOutput() ApplicationDefinitionOutput {
	return o
}

func (o ApplicationDefinitionOutput) ToApplicationDefinitionOutputWithContext(ctx context.Context) ApplicationDefinitionOutput {
	return o
}

// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
func (o ApplicationDefinitionOutput) Artifacts() ApplicationArtifactResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationDefinition) ApplicationArtifactResponseArrayOutput { return v.Artifacts }).(ApplicationArtifactResponseArrayOutput)
}

// The managed application provider authorizations.
func (o ApplicationDefinitionOutput) Authorizations() ApplicationProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationDefinition) ApplicationProviderAuthorizationResponseArrayOutput {
		return v.Authorizations
	}).(ApplicationProviderAuthorizationResponseArrayOutput)
}

// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
func (o ApplicationDefinitionOutput) CreateUiDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.AnyOutput { return v.CreateUiDefinition }).(pulumi.AnyOutput)
}

// The managed application definition description.
func (o ApplicationDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The managed application definition display name.
func (o ApplicationDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The identity of the resource.
func (o ApplicationDefinitionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// A value indicating whether the package is enabled or not.
func (o ApplicationDefinitionOutput) IsEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringPtrOutput { return v.IsEnabled }).(pulumi.StringPtrOutput)
}

// Resource location
func (o ApplicationDefinitionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The managed application lock level.
func (o ApplicationDefinitionOutput) LockLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringOutput { return v.LockLevel }).(pulumi.StringOutput)
}

// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
func (o ApplicationDefinitionOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.AnyOutput { return v.MainTemplate }).(pulumi.AnyOutput)
}

// ID of the resource that manages this resource.
func (o ApplicationDefinitionOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// Resource name
func (o ApplicationDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The managed application definition package file Uri. Use this element
func (o ApplicationDefinitionOutput) PackageFileUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringPtrOutput { return v.PackageFileUri }).(pulumi.StringPtrOutput)
}

// The SKU of the resource.
func (o ApplicationDefinitionOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationDefinition) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o ApplicationDefinitionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ApplicationDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationDefinitionOutput{})
}