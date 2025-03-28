// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a Namespace AuthorizationRules.
type NamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	// A string that describes the claim type
	ClaimType pulumi.StringOutput `pulumi:"claimType"`
	// A string that describes the claim value
	ClaimValue pulumi.StringOutput `pulumi:"claimValue"`
	// The created time for this rule
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// A string that describes the authorization rule.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The last modified time for this rule
	ModifiedTime pulumi.StringOutput `pulumi:"modifiedTime"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The revision number for the rule
	Revision pulumi.IntOutput `pulumi:"revision"`
	// The rights associated with the rule.
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// The sku of the created namespace
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNamespaceAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20230101preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20230901:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20231001preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs:NamespaceAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure-native:notificationhubs/v20170401:NamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceAuthorizationRule gets an existing NamespaceAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure-native:notificationhubs/v20170401:NamespaceAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceAuthorizationRule resources.
type namespaceAuthorizationRuleState struct {
}

type NamespaceAuthorizationRuleState struct {
}

func (NamespaceAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleState)(nil)).Elem()
}

type namespaceAuthorizationRuleArgs struct {
	// Authorization Rule Name.
	AuthorizationRuleName *string `pulumi:"authorizationRuleName"`
	// The namespace name.
	NamespaceName string `pulumi:"namespaceName"`
	// Properties of the Namespace AuthorizationRules.
	Properties SharedAccessAuthorizationRuleProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NamespaceAuthorizationRule resource.
type NamespaceAuthorizationRuleArgs struct {
	// Authorization Rule Name.
	AuthorizationRuleName pulumi.StringPtrInput
	// The namespace name.
	NamespaceName pulumi.StringInput
	// Properties of the Namespace AuthorizationRules.
	Properties SharedAccessAuthorizationRulePropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (NamespaceAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleArgs)(nil)).Elem()
}

type NamespaceAuthorizationRuleInput interface {
	pulumi.Input

	ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput
	ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput
}

func (*NamespaceAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return i.ToNamespaceAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceAuthorizationRuleOutput)
}

type NamespaceAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NamespaceAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return o
}

// A string that describes the claim type
func (o NamespaceAuthorizationRuleOutput) ClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.ClaimType }).(pulumi.StringOutput)
}

// A string that describes the claim value
func (o NamespaceAuthorizationRuleOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.ClaimValue }).(pulumi.StringOutput)
}

// The created time for this rule
func (o NamespaceAuthorizationRuleOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o NamespaceAuthorizationRuleOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// Resource location
func (o NamespaceAuthorizationRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The last modified time for this rule
func (o NamespaceAuthorizationRuleOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.ModifiedTime }).(pulumi.StringOutput)
}

// Resource name
func (o NamespaceAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o NamespaceAuthorizationRuleOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The revision number for the rule
func (o NamespaceAuthorizationRuleOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.IntOutput { return v.Revision }).(pulumi.IntOutput)
}

// The rights associated with the rule.
func (o NamespaceAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o NamespaceAuthorizationRuleOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.SecondaryKey }).(pulumi.StringOutput)
}

// The sku of the created namespace
func (o NamespaceAuthorizationRuleOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o NamespaceAuthorizationRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o NamespaceAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceAuthorizationRuleOutput{})
}
