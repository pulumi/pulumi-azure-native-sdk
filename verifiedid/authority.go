// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A VerifiedId authority resource
//
// Uses Azure REST API version 2024-01-26-preview. In version 2.x of the Azure Native provider, it used API version 2024-01-26-preview.
type Authority struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthority registers a new resource with the given unique name, arguments, and options.
func NewAuthority(ctx *pulumi.Context,
	name string, args *AuthorityArgs, opts ...pulumi.ResourceOption) (*Authority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:verifiedid/v20240126preview:Authority"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Authority
	err := ctx.RegisterResource("azure-native:verifiedid:Authority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthority gets an existing Authority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorityState, opts ...pulumi.ResourceOption) (*Authority, error) {
	var resource Authority
	err := ctx.ReadResource("azure-native:verifiedid:Authority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Authority resources.
type authorityState struct {
}

type AuthorityState struct {
}

func (AuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorityState)(nil)).Elem()
}

type authorityArgs struct {
	// The ID of the authority
	AuthorityName *string `pulumi:"authorityName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Authority resource.
type AuthorityArgs struct {
	// The ID of the authority
	AuthorityName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorityArgs)(nil)).Elem()
}

type AuthorityInput interface {
	pulumi.Input

	ToAuthorityOutput() AuthorityOutput
	ToAuthorityOutputWithContext(ctx context.Context) AuthorityOutput
}

func (*Authority) ElementType() reflect.Type {
	return reflect.TypeOf((**Authority)(nil)).Elem()
}

func (i *Authority) ToAuthorityOutput() AuthorityOutput {
	return i.ToAuthorityOutputWithContext(context.Background())
}

func (i *Authority) ToAuthorityOutputWithContext(ctx context.Context) AuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorityOutput)
}

type AuthorityOutput struct{ *pulumi.OutputState }

func (AuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authority)(nil)).Elem()
}

func (o AuthorityOutput) ToAuthorityOutput() AuthorityOutput {
	return o
}

func (o AuthorityOutput) ToAuthorityOutputWithContext(ctx context.Context) AuthorityOutput {
	return o
}

// The Azure API version of the resource.
func (o AuthorityOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Authority) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o AuthorityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Authority) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AuthorityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Authority) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o AuthorityOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Authority) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AuthorityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Authority) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AuthorityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Authority) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AuthorityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Authority) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorityOutput{})
}
