// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cognitive Services Account is an Azure resource representing the provisioned account, its type, location and SKU.
//
// Deprecated: Version 2016-02-01-preview will be removed in v2 of the provider.
type CognitiveServicesAccount struct {
	pulumi.CustomResourceState

	// Endpoint of the created account
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// Entity Tag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Type of cognitive service account.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the created account
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets the status of the cognitive services account at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU of the cognitive services account.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewCognitiveServicesAccount registers a new resource with the given unique name, arguments, and options.
func NewCognitiveServicesAccount(ctx *pulumi.Context,
	name string, args *CognitiveServicesAccountArgs, opts ...pulumi.ResourceOption) (*CognitiveServicesAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20170418:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20210430:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20211001:CognitiveServicesAccount"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:CognitiveServicesAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource CognitiveServicesAccount
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20160201preview:CognitiveServicesAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCognitiveServicesAccount gets an existing CognitiveServicesAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCognitiveServicesAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CognitiveServicesAccountState, opts ...pulumi.ResourceOption) (*CognitiveServicesAccount, error) {
	var resource CognitiveServicesAccount
	err := ctx.ReadResource("azure-native:cognitiveservices/v20160201preview:CognitiveServicesAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CognitiveServicesAccount resources.
type cognitiveServicesAccountState struct {
}

type CognitiveServicesAccountState struct {
}

func (CognitiveServicesAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*cognitiveServicesAccountState)(nil)).Elem()
}

type cognitiveServicesAccountArgs struct {
	// The name of the cognitive services account within the specified resource group. Cognitive Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName *string `pulumi:"accountName"`
	// Required. Indicates the type of cognitive service account.
	Kind string `pulumi:"kind"`
	// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update the request will succeed.
	Location *string `pulumi:"location"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the cognitive services account.
	Sku Sku `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CognitiveServicesAccount resource.
type CognitiveServicesAccountArgs struct {
	// The name of the cognitive services account within the specified resource group. Cognitive Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringPtrInput
	// Required. Indicates the type of cognitive service account.
	Kind pulumi.StringInput
	// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update the request will succeed.
	Location pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The SKU of the cognitive services account.
	Sku SkuInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (CognitiveServicesAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cognitiveServicesAccountArgs)(nil)).Elem()
}

type CognitiveServicesAccountInput interface {
	pulumi.Input

	ToCognitiveServicesAccountOutput() CognitiveServicesAccountOutput
	ToCognitiveServicesAccountOutputWithContext(ctx context.Context) CognitiveServicesAccountOutput
}

func (*CognitiveServicesAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccount)(nil)).Elem()
}

func (i *CognitiveServicesAccount) ToCognitiveServicesAccountOutput() CognitiveServicesAccountOutput {
	return i.ToCognitiveServicesAccountOutputWithContext(context.Background())
}

func (i *CognitiveServicesAccount) ToCognitiveServicesAccountOutputWithContext(ctx context.Context) CognitiveServicesAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CognitiveServicesAccountOutput)
}

type CognitiveServicesAccountOutput struct{ *pulumi.OutputState }

func (CognitiveServicesAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CognitiveServicesAccount)(nil)).Elem()
}

func (o CognitiveServicesAccountOutput) ToCognitiveServicesAccountOutput() CognitiveServicesAccountOutput {
	return o
}

func (o CognitiveServicesAccountOutput) ToCognitiveServicesAccountOutputWithContext(ctx context.Context) CognitiveServicesAccountOutput {
	return o
}

// Endpoint of the created account
func (o CognitiveServicesAccountOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// Entity Tag
func (o CognitiveServicesAccountOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Type of cognitive service account.
func (o CognitiveServicesAccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The location of the resource
func (o CognitiveServicesAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the created account
func (o CognitiveServicesAccountOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Gets the status of the cognitive services account at the time the operation was called.
func (o CognitiveServicesAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the cognitive services account.
func (o CognitiveServicesAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
func (o CognitiveServicesAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o CognitiveServicesAccountOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CognitiveServicesAccount) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CognitiveServicesAccountOutput{})
}