// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security operator under a given subscription and pricing
//
// Uses Azure REST API version 2023-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-01-01-preview.
type SecurityOperator struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityOperator registers a new resource with the given unique name, arguments, and options.
func NewSecurityOperator(ctx *pulumi.Context,
	name string, args *SecurityOperatorArgs, opts ...pulumi.ResourceOption) (*SecurityOperator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PricingName == nil {
		return nil, errors.New("invalid value for required argument 'PricingName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20230101preview:SecurityOperator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecurityOperator
	err := ctx.RegisterResource("azure-native:security:SecurityOperator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityOperator gets an existing SecurityOperator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityOperator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityOperatorState, opts ...pulumi.ResourceOption) (*SecurityOperator, error) {
	var resource SecurityOperator
	err := ctx.ReadResource("azure-native:security:SecurityOperator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityOperator resources.
type securityOperatorState struct {
}

type SecurityOperatorState struct {
}

func (SecurityOperatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityOperatorState)(nil)).Elem()
}

type securityOperatorArgs struct {
	// name of the pricing configuration
	PricingName string `pulumi:"pricingName"`
	// name of the securityOperator
	SecurityOperatorName *string `pulumi:"securityOperatorName"`
}

// The set of arguments for constructing a SecurityOperator resource.
type SecurityOperatorArgs struct {
	// name of the pricing configuration
	PricingName pulumi.StringInput
	// name of the securityOperator
	SecurityOperatorName pulumi.StringPtrInput
}

func (SecurityOperatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityOperatorArgs)(nil)).Elem()
}

type SecurityOperatorInput interface {
	pulumi.Input

	ToSecurityOperatorOutput() SecurityOperatorOutput
	ToSecurityOperatorOutputWithContext(ctx context.Context) SecurityOperatorOutput
}

func (*SecurityOperator) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityOperator)(nil)).Elem()
}

func (i *SecurityOperator) ToSecurityOperatorOutput() SecurityOperatorOutput {
	return i.ToSecurityOperatorOutputWithContext(context.Background())
}

func (i *SecurityOperator) ToSecurityOperatorOutputWithContext(ctx context.Context) SecurityOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityOperatorOutput)
}

type SecurityOperatorOutput struct{ *pulumi.OutputState }

func (SecurityOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityOperator)(nil)).Elem()
}

func (o SecurityOperatorOutput) ToSecurityOperatorOutput() SecurityOperatorOutput {
	return o
}

func (o SecurityOperatorOutput) ToSecurityOperatorOutputWithContext(ctx context.Context) SecurityOperatorOutput {
	return o
}

// The Azure API version of the resource.
func (o SecurityOperatorOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityOperator) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o SecurityOperatorOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *SecurityOperator) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Resource name
func (o SecurityOperatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityOperator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o SecurityOperatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityOperator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityOperatorOutput{})
}
