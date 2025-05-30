// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tunnel Policy maps domains to target endpoints to process traffic over the tunnelling protocol.
//
// Uses Azure REST API version 2024-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-06-01-preview.
type TunnelPolicy struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion  pulumi.StringOutput `pulumi:"azureApiVersion"`
	DeploymentStatus pulumi.StringOutput `pulumi:"deploymentStatus"`
	// Domains referenced by this tunnel policy.
	Domains ActivatedResourceReferenceResponseArrayOutput `pulumi:"domains"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Target Groups referenced by this tunnel policy.
	TargetGroups ResourceReferenceResponseArrayOutput `pulumi:"targetGroups"`
	// Protocol this tunnel will use for allowing traffic to backends.
	TunnelType pulumi.StringPtrOutput `pulumi:"tunnelType"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTunnelPolicy registers a new resource with the given unique name, arguments, and options.
func NewTunnelPolicy(ctx *pulumi.Context,
	name string, args *TunnelPolicyArgs, opts ...pulumi.ResourceOption) (*TunnelPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domains == nil {
		return nil, errors.New("invalid value for required argument 'Domains'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20240601preview:TunnelPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TunnelPolicy
	err := ctx.RegisterResource("azure-native:cdn:TunnelPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTunnelPolicy gets an existing TunnelPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunnelPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TunnelPolicyState, opts ...pulumi.ResourceOption) (*TunnelPolicy, error) {
	var resource TunnelPolicy
	err := ctx.ReadResource("azure-native:cdn:TunnelPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TunnelPolicy resources.
type tunnelPolicyState struct {
}

type TunnelPolicyState struct {
}

func (TunnelPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelPolicyState)(nil)).Elem()
}

type tunnelPolicyArgs struct {
	// Domains referenced by this tunnel policy.
	Domains []ActivatedResourceReference `pulumi:"domains"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Target Groups referenced by this tunnel policy.
	TargetGroups []ResourceReference `pulumi:"targetGroups"`
	// Name of the Tunnel Policy under the profile.
	TunnelPolicyName *string `pulumi:"tunnelPolicyName"`
	// Protocol this tunnel will use for allowing traffic to backends.
	TunnelType *string `pulumi:"tunnelType"`
}

// The set of arguments for constructing a TunnelPolicy resource.
type TunnelPolicyArgs struct {
	// Domains referenced by this tunnel policy.
	Domains ActivatedResourceReferenceArrayInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Target Groups referenced by this tunnel policy.
	TargetGroups ResourceReferenceArrayInput
	// Name of the Tunnel Policy under the profile.
	TunnelPolicyName pulumi.StringPtrInput
	// Protocol this tunnel will use for allowing traffic to backends.
	TunnelType pulumi.StringPtrInput
}

func (TunnelPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelPolicyArgs)(nil)).Elem()
}

type TunnelPolicyInput interface {
	pulumi.Input

	ToTunnelPolicyOutput() TunnelPolicyOutput
	ToTunnelPolicyOutputWithContext(ctx context.Context) TunnelPolicyOutput
}

func (*TunnelPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelPolicy)(nil)).Elem()
}

func (i *TunnelPolicy) ToTunnelPolicyOutput() TunnelPolicyOutput {
	return i.ToTunnelPolicyOutputWithContext(context.Background())
}

func (i *TunnelPolicy) ToTunnelPolicyOutputWithContext(ctx context.Context) TunnelPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelPolicyOutput)
}

type TunnelPolicyOutput struct{ *pulumi.OutputState }

func (TunnelPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TunnelPolicy)(nil)).Elem()
}

func (o TunnelPolicyOutput) ToTunnelPolicyOutput() TunnelPolicyOutput {
	return o
}

func (o TunnelPolicyOutput) ToTunnelPolicyOutputWithContext(ctx context.Context) TunnelPolicyOutput {
	return o
}

// The Azure API version of the resource.
func (o TunnelPolicyOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelPolicy) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

func (o TunnelPolicyOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelPolicy) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Domains referenced by this tunnel policy.
func (o TunnelPolicyOutput) Domains() ActivatedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *TunnelPolicy) ActivatedResourceReferenceResponseArrayOutput { return v.Domains }).(ActivatedResourceReferenceResponseArrayOutput)
}

// Resource name.
func (o TunnelPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning status
func (o TunnelPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o TunnelPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TunnelPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Target Groups referenced by this tunnel policy.
func (o TunnelPolicyOutput) TargetGroups() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *TunnelPolicy) ResourceReferenceResponseArrayOutput { return v.TargetGroups }).(ResourceReferenceResponseArrayOutput)
}

// Protocol this tunnel will use for allowing traffic to backends.
func (o TunnelPolicyOutput) TunnelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TunnelPolicy) pulumi.StringPtrOutput { return v.TunnelType }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o TunnelPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TunnelPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TunnelPolicyOutput{})
}
