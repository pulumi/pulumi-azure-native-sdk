// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the HybridIdentityMetadata.
//
// Uses Azure REST API version 2022-12-15-preview. In version 2.x of the Azure Native provider, it used API version 2022-12-15-preview.
type HybridIdentityMetadatum struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Identity for the resource.
	Identity IdentityResponseOutput `pulumi:"identity"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Public Key.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// The unique identifier for the resource.
	ResourceUid pulumi.StringPtrOutput `pulumi:"resourceUid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHybridIdentityMetadatum registers a new resource with the given unique name, arguments, and options.
func NewHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, args *HybridIdentityMetadatumArgs, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:HybridIdentityMetadatum"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:HybridIdentityMetadatum"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HybridIdentityMetadatum
	err := ctx.RegisterResource("azure-native:azurestackhci:HybridIdentityMetadatum", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridIdentityMetadatum gets an existing HybridIdentityMetadatum resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridIdentityMetadatumState, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	var resource HybridIdentityMetadatum
	err := ctx.ReadResource("azure-native:azurestackhci:HybridIdentityMetadatum", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridIdentityMetadatum resources.
type hybridIdentityMetadatumState struct {
}

type HybridIdentityMetadatumState struct {
}

func (HybridIdentityMetadatumState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadatumState)(nil)).Elem()
}

type hybridIdentityMetadatumArgs struct {
	// Name of the hybridIdentityMetadata.
	MetadataName *string `pulumi:"metadataName"`
	// The Public Key.
	PublicKey *string `pulumi:"publicKey"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The unique identifier for the resource.
	ResourceUid *string `pulumi:"resourceUid"`
	// Name of the vm.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a HybridIdentityMetadatum resource.
type HybridIdentityMetadatumArgs struct {
	// Name of the hybridIdentityMetadata.
	MetadataName pulumi.StringPtrInput
	// The Public Key.
	PublicKey pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The unique identifier for the resource.
	ResourceUid pulumi.StringPtrInput
	// Name of the vm.
	VirtualMachineName pulumi.StringInput
}

func (HybridIdentityMetadatumArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadatumArgs)(nil)).Elem()
}

type HybridIdentityMetadatumInput interface {
	pulumi.Input

	ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput
	ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput
}

func (*HybridIdentityMetadatum) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadatum)(nil)).Elem()
}

func (i *HybridIdentityMetadatum) ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput {
	return i.ToHybridIdentityMetadatumOutputWithContext(context.Background())
}

func (i *HybridIdentityMetadatum) ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridIdentityMetadatumOutput)
}

type HybridIdentityMetadatumOutput struct{ *pulumi.OutputState }

func (HybridIdentityMetadatumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadatum)(nil)).Elem()
}

func (o HybridIdentityMetadatumOutput) ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput {
	return o
}

func (o HybridIdentityMetadatumOutput) ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput {
	return o
}

// The Azure API version of the resource.
func (o HybridIdentityMetadatumOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o HybridIdentityMetadatumOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) IdentityResponseOutput { return v.Identity }).(IdentityResponseOutput)
}

// The name of the resource
func (o HybridIdentityMetadatumOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o HybridIdentityMetadatumOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Public Key.
func (o HybridIdentityMetadatumOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource.
func (o HybridIdentityMetadatumOutput) ResourceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.ResourceUid }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HybridIdentityMetadatumOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HybridIdentityMetadatumOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridIdentityMetadatumOutput{})
}
