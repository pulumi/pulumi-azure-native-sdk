// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170418

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cognitive Services Account is an Azure resource representing the provisioned account, its type, location and SKU.
type Account struct {
	pulumi.CustomResourceState

	// Entity Tag
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The identity of Cognitive Services account.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the created account
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Cognitive Services account.
	Properties CognitiveServicesAccountPropertiesResponseOutput `pulumi:"properties"`
	// The SKU of Cognitive Services account.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToCognitiveServicesAccountPropertiesPtrOutput().ApplyT(func(v *CognitiveServicesAccountProperties) *CognitiveServicesAccountProperties { return v.Defaults() }).(CognitiveServicesAccountPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20160201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20210430:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20211001:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20170418:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:cognitiveservices/v20170418:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of Cognitive Services account.
	AccountName *string `pulumi:"accountName"`
	// The identity of Cognitive Services account.
	Identity *Identity `pulumi:"identity"`
	// The Kind of the resource.
	Kind *string `pulumi:"kind"`
	// The location of the resource
	Location *string `pulumi:"location"`
	// Properties of Cognitive Services account.
	Properties *CognitiveServicesAccountProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of Cognitive Services account.
	Sku *Sku `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of Cognitive Services account.
	AccountName pulumi.StringPtrInput
	// The identity of Cognitive Services account.
	Identity IdentityPtrInput
	// The Kind of the resource.
	Kind pulumi.StringPtrInput
	// The location of the resource
	Location pulumi.StringPtrInput
	// Properties of Cognitive Services account.
	Properties CognitiveServicesAccountPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of Cognitive Services account.
	Sku SkuPtrInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// Entity Tag
func (o AccountOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The identity of Cognitive Services account.
func (o AccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The Kind of the resource.
func (o AccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The location of the resource
func (o AccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the created account
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Cognitive Services account.
func (o AccountOutput) Properties() CognitiveServicesAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *Account) CognitiveServicesAccountPropertiesResponseOutput { return v.Properties }).(CognitiveServicesAccountPropertiesResponseOutput)
}

// The SKU of Cognitive Services account.
func (o AccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Account) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}