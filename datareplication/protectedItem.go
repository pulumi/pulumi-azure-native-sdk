// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datareplication

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Protected item model.
//
// Uses Azure REST API version 2021-02-16-preview. In version 2.x of the Azure Native provider, it used API version 2021-02-16-preview.
//
// Other available API versions: 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datareplication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ProtectedItem struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Gets or sets the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Protected item model properties.
	Properties ProtectedItemModelPropertiesResponseOutput `pulumi:"properties"`
	SystemData ProtectedItemModelResponseSystemDataOutput `pulumi:"systemData"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProtectedItem registers a new resource with the given unique name, arguments, and options.
func NewProtectedItem(ctx *pulumi.Context,
	name string, args *ProtectedItemArgs, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datareplication/v20210216preview:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:datareplication/v20240901:ProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProtectedItem
	err := ctx.RegisterResource("azure-native:datareplication:ProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectedItem gets an existing ProtectedItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectedItemState, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	var resource ProtectedItem
	err := ctx.ReadResource("azure-native:datareplication:ProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectedItem resources.
type protectedItemState struct {
}

type ProtectedItemState struct {
}

func (ProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemState)(nil)).Elem()
}

type protectedItemArgs struct {
	// Protected item model properties.
	Properties ProtectedItemModelProperties `pulumi:"properties"`
	// The protected item name.
	ProtectedItemName *string `pulumi:"protectedItemName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ProtectedItem resource.
type ProtectedItemArgs struct {
	// Protected item model properties.
	Properties ProtectedItemModelPropertiesInput
	// The protected item name.
	ProtectedItemName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The vault name.
	VaultName pulumi.StringInput
}

func (ProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemArgs)(nil)).Elem()
}

type ProtectedItemInput interface {
	pulumi.Input

	ToProtectedItemOutput() ProtectedItemOutput
	ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput
}

func (*ProtectedItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedItem)(nil)).Elem()
}

func (i *ProtectedItem) ToProtectedItemOutput() ProtectedItemOutput {
	return i.ToProtectedItemOutputWithContext(context.Background())
}

func (i *ProtectedItem) ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedItemOutput)
}

type ProtectedItemOutput struct{ *pulumi.OutputState }

func (ProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedItem)(nil)).Elem()
}

func (o ProtectedItemOutput) ToProtectedItemOutput() ProtectedItemOutput {
	return o
}

func (o ProtectedItemOutput) ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput {
	return o
}

// The Azure API version of the resource.
func (o ProtectedItemOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the name of the resource.
func (o ProtectedItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Protected item model properties.
func (o ProtectedItemOutput) Properties() ProtectedItemModelPropertiesResponseOutput {
	return o.ApplyT(func(v *ProtectedItem) ProtectedItemModelPropertiesResponseOutput { return v.Properties }).(ProtectedItemModelPropertiesResponseOutput)
}

func (o ProtectedItemOutput) SystemData() ProtectedItemModelResponseSystemDataOutput {
	return o.ApplyT(func(v *ProtectedItem) ProtectedItemModelResponseSystemDataOutput { return v.SystemData }).(ProtectedItemModelResponseSystemDataOutput)
}

// Gets or sets the type of the resource.
func (o ProtectedItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectedItemOutput{})
}
