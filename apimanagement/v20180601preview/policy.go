// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy Contract details.
type Policy struct {
	pulumi.CustomResourceState

	// Format of the policyContent.
	ContentFormat pulumi.StringPtrOutput `pulumi:"contentFormat"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringOutput `pulumi:"policyContent"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyContent == nil {
		return nil, errors.New("invalid value for required argument 'PolicyContent'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if isZero(args.ContentFormat) {
		args.ContentFormat = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Policy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Format of the policyContent.
	ContentFormat *string `pulumi:"contentFormat"`
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent string `pulumi:"policyContent"`
	// The identifier of the Policy.
	PolicyId *string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Format of the policyContent.
	ContentFormat pulumi.StringPtrInput
	// Json escaped Xml Encoded contents of the Policy.
	PolicyContent pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Format of the policyContent.
func (o PolicyOutput) ContentFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.ContentFormat }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Json escaped Xml Encoded contents of the Policy.
func (o PolicyOutput) PolicyContent() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PolicyContent }).(pulumi.StringOutput)
}

// Resource type for API Management resource.
func (o PolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}