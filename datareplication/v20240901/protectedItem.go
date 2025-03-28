// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Protected item model.
type ProtectedItem struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties ProtectedItemModelPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProtectedItem registers a new resource with the given unique name, arguments, and options.
func NewProtectedItem(ctx *pulumi.Context,
	name string, args *ProtectedItemArgs, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:datareplication:ProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProtectedItem
	err := ctx.RegisterResource("azure-native:datareplication/v20240901:ProtectedItem", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:datareplication/v20240901:ProtectedItem", name, id, state, &resource, opts...)
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
	// The resource-specific properties for this resource.
	Properties *ProtectedItemModelProperties `pulumi:"properties"`
	// The protected item name.
	ProtectedItemName *string `pulumi:"protectedItemName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ProtectedItem resource.
type ProtectedItemArgs struct {
	// The resource-specific properties for this resource.
	Properties ProtectedItemModelPropertiesPtrInput
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

// The name of the resource
func (o ProtectedItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o ProtectedItemOutput) Properties() ProtectedItemModelPropertiesResponseOutput {
	return o.ApplyT(func(v *ProtectedItem) ProtectedItemModelPropertiesResponseOutput { return v.Properties }).(ProtectedItemModelPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ProtectedItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProtectedItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProtectedItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectedItemOutput{})
}
