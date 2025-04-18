// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package addons

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The status of the Canonical support plan.
//
// Uses Azure REST API version 2018-03-01. In version 2.x of the Azure Native provider, it used API version 2018-03-01.
type SupportPlanType struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the Canonical support plan, i.e. "essential", "standard" or "advanced".
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Microsoft.Addons/supportProvider
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSupportPlanType registers a new resource with the given unique name, arguments, and options.
func NewSupportPlanType(ctx *pulumi.Context,
	name string, args *SupportPlanTypeArgs, opts ...pulumi.ResourceOption) (*SupportPlanType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:addons/v20170515:SupportPlanType"),
		},
		{
			Type: pulumi.String("azure-native:addons/v20180301:SupportPlanType"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SupportPlanType
	err := ctx.RegisterResource("azure-native:addons:SupportPlanType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSupportPlanType gets an existing SupportPlanType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSupportPlanType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SupportPlanTypeState, opts ...pulumi.ResourceOption) (*SupportPlanType, error) {
	var resource SupportPlanType
	err := ctx.ReadResource("azure-native:addons:SupportPlanType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SupportPlanType resources.
type supportPlanTypeState struct {
}

type SupportPlanTypeState struct {
}

func (SupportPlanTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*supportPlanTypeState)(nil)).Elem()
}

type supportPlanTypeArgs struct {
	// The Canonical support plan type.
	PlanTypeName *string `pulumi:"planTypeName"`
	// The support plan type. For now the only valid type is "canonical".
	ProviderName string `pulumi:"providerName"`
}

// The set of arguments for constructing a SupportPlanType resource.
type SupportPlanTypeArgs struct {
	// The Canonical support plan type.
	PlanTypeName pulumi.StringPtrInput
	// The support plan type. For now the only valid type is "canonical".
	ProviderName pulumi.StringInput
}

func (SupportPlanTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*supportPlanTypeArgs)(nil)).Elem()
}

type SupportPlanTypeInput interface {
	pulumi.Input

	ToSupportPlanTypeOutput() SupportPlanTypeOutput
	ToSupportPlanTypeOutputWithContext(ctx context.Context) SupportPlanTypeOutput
}

func (*SupportPlanType) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportPlanType)(nil)).Elem()
}

func (i *SupportPlanType) ToSupportPlanTypeOutput() SupportPlanTypeOutput {
	return i.ToSupportPlanTypeOutputWithContext(context.Background())
}

func (i *SupportPlanType) ToSupportPlanTypeOutputWithContext(ctx context.Context) SupportPlanTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportPlanTypeOutput)
}

type SupportPlanTypeOutput struct{ *pulumi.OutputState }

func (SupportPlanTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportPlanType)(nil)).Elem()
}

func (o SupportPlanTypeOutput) ToSupportPlanTypeOutput() SupportPlanTypeOutput {
	return o
}

func (o SupportPlanTypeOutput) ToSupportPlanTypeOutputWithContext(ctx context.Context) SupportPlanTypeOutput {
	return o
}

// The Azure API version of the resource.
func (o SupportPlanTypeOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SupportPlanType) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the Canonical support plan, i.e. "essential", "standard" or "advanced".
func (o SupportPlanTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SupportPlanType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o SupportPlanTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SupportPlanType) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Microsoft.Addons/supportProvider
func (o SupportPlanTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SupportPlanType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SupportPlanTypeOutput{})
}
