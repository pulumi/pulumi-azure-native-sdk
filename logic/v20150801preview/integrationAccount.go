// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccount struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The sku.
	Sku IntegrationAccountSkuResponsePtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewIntegrationAccount registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccount(ctx *pulumi.Context,
	name string, args *IntegrationAccountArgs, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationAccount
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccount gets an existing IntegrationAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountState, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	var resource IntegrationAccount
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccount resources.
type integrationAccountState struct {
}

type IntegrationAccountState struct {
}

func (IntegrationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountState)(nil)).Elem()
}

type integrationAccountArgs struct {
	// The resource id.
	Id *string `pulumi:"id"`
	// The integration account name.
	IntegrationAccountName *string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku.
	Sku *IntegrationAccountSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IntegrationAccount resource.
type IntegrationAccountArgs struct {
	// The resource id.
	Id pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The sku.
	Sku IntegrationAccountSkuPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountArgs)(nil)).Elem()
}

type IntegrationAccountInput interface {
	pulumi.Input

	ToIntegrationAccountOutput() IntegrationAccountOutput
	ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput
}

func (*IntegrationAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccount)(nil)).Elem()
}

func (i *IntegrationAccount) ToIntegrationAccountOutput() IntegrationAccountOutput {
	return i.ToIntegrationAccountOutputWithContext(context.Background())
}

func (i *IntegrationAccount) ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountOutput)
}

type IntegrationAccountOutput struct{ *pulumi.OutputState }

func (IntegrationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccount)(nil)).Elem()
}

func (o IntegrationAccountOutput) ToIntegrationAccountOutput() IntegrationAccountOutput {
	return o
}

func (o IntegrationAccountOutput) ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput {
	return o
}

// The resource location.
func (o IntegrationAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o IntegrationAccountOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccount) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The sku.
func (o IntegrationAccountOutput) Sku() IntegrationAccountSkuResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationAccount) IntegrationAccountSkuResponsePtrOutput { return v.Sku }).(IntegrationAccountSkuResponsePtrOutput)
}

// The resource tags.
func (o IntegrationAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o IntegrationAccountOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccount) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountOutput{})
}
