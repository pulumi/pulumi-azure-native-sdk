// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190228preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SpatialAnchorsAccount Response.
//
// Deprecated: Version 2019-02-28-preview will be removed in v2 of the provider.
type SpatialAnchorsAccount struct {
	pulumi.CustomResourceState

	// Correspond domain name of certain Spatial Anchors Account
	AccountDomain pulumi.StringOutput `pulumi:"accountDomain"`
	// unique id of certain Spatial Anchors Account data contract.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The identity associated with this account
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSpatialAnchorsAccount registers a new resource with the given unique name, arguments, and options.
func NewSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, args *SpatialAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mixedreality:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20191202preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20200501:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210101:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:SpatialAnchorsAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azure-native:mixedreality/v20190228preview:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpatialAnchorsAccount gets an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azure-native:mixedreality/v20190228preview:SpatialAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpatialAnchorsAccount resources.
type spatialAnchorsAccountState struct {
}

type SpatialAnchorsAccountState struct {
}

func (SpatialAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountState)(nil)).Elem()
}

type spatialAnchorsAccountArgs struct {
	// The identity associated with this account
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of an Mixed Reality Spatial Anchors Account.
	SpatialAnchorsAccountName *string `pulumi:"spatialAnchorsAccountName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SpatialAnchorsAccount resource.
type SpatialAnchorsAccountArgs struct {
	// The identity associated with this account
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Name of an Mixed Reality Spatial Anchors Account.
	SpatialAnchorsAccountName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountArgs)(nil)).Elem()
}

type SpatialAnchorsAccountInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput
	ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput
}

func (*SpatialAnchorsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil)).Elem()
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return i.ToSpatialAnchorsAccountOutputWithContext(context.Background())
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountOutput)
}

type SpatialAnchorsAccountOutput struct{ *pulumi.OutputState }

func (SpatialAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil)).Elem()
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return o
}

// Correspond domain name of certain Spatial Anchors Account
func (o SpatialAnchorsAccountOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.AccountDomain }).(pulumi.StringOutput)
}

// unique id of certain Spatial Anchors Account data contract.
func (o SpatialAnchorsAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The identity associated with this account
func (o SpatialAnchorsAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SpatialAnchorsAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SpatialAnchorsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource tags.
func (o SpatialAnchorsAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SpatialAnchorsAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SpatialAnchorsAccountOutput{})
}