// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an curation profile resource.
//
// Uses Azure REST API version 2024-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-08-01-preview.
//
// Other available API versions: 2024-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type CurationProfile struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource policies that are a part of this curation profile.
	ResourcePolicies ResourcePolicyResponseArrayOutput `pulumi:"resourcePolicies"`
	// Resources that have access to the shared resources that are a part of this curation profile.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCurationProfile registers a new resource with the given unique name, arguments, and options.
func NewCurationProfile(ctx *pulumi.Context,
	name string, args *CurationProfileArgs, opts ...pulumi.ResourceOption) (*CurationProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20240801preview:CurationProfile"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20241001preview:CurationProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CurationProfile
	err := ctx.RegisterResource("azure-native:devcenter:CurationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCurationProfile gets an existing CurationProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCurationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CurationProfileState, opts ...pulumi.ResourceOption) (*CurationProfile, error) {
	var resource CurationProfile
	err := ctx.ReadResource("azure-native:devcenter:CurationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CurationProfile resources.
type curationProfileState struct {
}

type CurationProfileState struct {
}

func (CurationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*curationProfileState)(nil)).Elem()
}

type curationProfileArgs struct {
	// The name of the curation profile.
	CurationProfileName *string `pulumi:"curationProfileName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource policies that are a part of this curation profile.
	ResourcePolicies []ResourcePolicy `pulumi:"resourcePolicies"`
	// Resources that have access to the shared resources that are a part of this curation profile.
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a CurationProfile resource.
type CurationProfileArgs struct {
	// The name of the curation profile.
	CurationProfileName pulumi.StringPtrInput
	// The name of the devcenter.
	DevCenterName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource policies that are a part of this curation profile.
	ResourcePolicies ResourcePolicyArrayInput
	// Resources that have access to the shared resources that are a part of this curation profile.
	Scopes pulumi.StringArrayInput
}

func (CurationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*curationProfileArgs)(nil)).Elem()
}

type CurationProfileInput interface {
	pulumi.Input

	ToCurationProfileOutput() CurationProfileOutput
	ToCurationProfileOutputWithContext(ctx context.Context) CurationProfileOutput
}

func (*CurationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**CurationProfile)(nil)).Elem()
}

func (i *CurationProfile) ToCurationProfileOutput() CurationProfileOutput {
	return i.ToCurationProfileOutputWithContext(context.Background())
}

func (i *CurationProfile) ToCurationProfileOutputWithContext(ctx context.Context) CurationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurationProfileOutput)
}

type CurationProfileOutput struct{ *pulumi.OutputState }

func (CurationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CurationProfile)(nil)).Elem()
}

func (o CurationProfileOutput) ToCurationProfileOutput() CurationProfileOutput {
	return o
}

func (o CurationProfileOutput) ToCurationProfileOutputWithContext(ctx context.Context) CurationProfileOutput {
	return o
}

// The Azure API version of the resource.
func (o CurationProfileOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CurationProfile) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o CurationProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CurationProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o CurationProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CurationProfile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource policies that are a part of this curation profile.
func (o CurationProfileOutput) ResourcePolicies() ResourcePolicyResponseArrayOutput {
	return o.ApplyT(func(v *CurationProfile) ResourcePolicyResponseArrayOutput { return v.ResourcePolicies }).(ResourcePolicyResponseArrayOutput)
}

// Resources that have access to the shared resources that are a part of this curation profile.
func (o CurationProfileOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CurationProfile) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CurationProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CurationProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CurationProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CurationProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CurationProfileOutput{})
}
