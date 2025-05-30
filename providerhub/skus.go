// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2021-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-09-01-preview.
type Skus struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSkus registers a new resource with the given unique name, arguments, and options.
func NewSkus(ctx *pulumi.Context,
	name string, args *SkusArgs, opts ...pulumi.ResourceOption) (*Skus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:Skus"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Skus
	err := ctx.RegisterResource("azure-native:providerhub:Skus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSkus gets an existing Skus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusState, opts ...pulumi.ResourceOption) (*Skus, error) {
	var resource Skus
	err := ctx.ReadResource("azure-native:providerhub:Skus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Skus resources.
type skusState struct {
}

type SkusState struct {
}

func (SkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusState)(nil)).Elem()
}

type skusArgs struct {
	Properties *SkuResourceProperties `pulumi:"properties"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku *string `pulumi:"sku"`
}

// The set of arguments for constructing a Skus resource.
type SkusArgs struct {
	Properties SkuResourcePropertiesPtrInput
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput
	// The resource type.
	ResourceType pulumi.StringInput
	// The SKU.
	Sku pulumi.StringPtrInput
}

func (SkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusArgs)(nil)).Elem()
}

type SkusInput interface {
	pulumi.Input

	ToSkusOutput() SkusOutput
	ToSkusOutputWithContext(ctx context.Context) SkusOutput
}

func (*Skus) ElementType() reflect.Type {
	return reflect.TypeOf((**Skus)(nil)).Elem()
}

func (i *Skus) ToSkusOutput() SkusOutput {
	return i.ToSkusOutputWithContext(context.Background())
}

func (i *Skus) ToSkusOutputWithContext(ctx context.Context) SkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusOutput)
}

type SkusOutput struct{ *pulumi.OutputState }

func (SkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Skus)(nil)).Elem()
}

func (o SkusOutput) ToSkusOutput() SkusOutput {
	return o
}

func (o SkusOutput) ToSkusOutputWithContext(ctx context.Context) SkusOutput {
	return o
}

// The Azure API version of the resource.
func (o SkusOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Skus) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o SkusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Skus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SkusOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *Skus) SkuResourceResponsePropertiesOutput { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o SkusOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Skus) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SkusOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Skus) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SkusOutput{})
}
