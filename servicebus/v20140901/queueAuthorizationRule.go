// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a namespace authorization rule.
//
// Deprecated: Version 2014-09-01 will be removed in v2 of the provider.
type QueueAuthorizationRule struct {
	pulumi.CustomResourceState

	// A string that describes Claim Type for authorization rule.
	ClaimType pulumi.StringPtrOutput `pulumi:"claimType"`
	// A string that describes Claim Value of authorization rule.
	ClaimValue pulumi.StringPtrOutput `pulumi:"claimValue"`
	// The time the namespace was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// A string that describes the Key Name of authorization rule.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The time the namespace was updated.
	ModifiedTime pulumi.StringOutput `pulumi:"modifiedTime"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringPtrOutput `pulumi:"primaryKey"`
	// The rights associated with the rule.
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringPtrOutput `pulumi:"secondaryKey"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQueueAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewQueueAuthorizationRule(ctx *pulumi.Context,
	name string, args *QueueAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.QueueName == nil {
		return nil, errors.New("invalid value for required argument 'QueueName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:QueueAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource QueueAuthorizationRule
	err := ctx.RegisterResource("azure-native:servicebus/v20140901:QueueAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueAuthorizationRule gets an existing QueueAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueAuthorizationRuleState, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	var resource QueueAuthorizationRule
	err := ctx.ReadResource("azure-native:servicebus/v20140901:QueueAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueAuthorizationRule resources.
type queueAuthorizationRuleState struct {
}

type QueueAuthorizationRuleState struct {
}

func (QueueAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleState)(nil)).Elem()
}

type queueAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName *string `pulumi:"authorizationRuleName"`
	// A string that describes Claim Type for authorization rule.
	ClaimType *string `pulumi:"claimType"`
	// A string that describes Claim Value of authorization rule.
	ClaimValue *string `pulumi:"claimValue"`
	// A string that describes the Key Name of authorization rule.
	KeyName *string `pulumi:"keyName"`
	// data center location.
	Location *string `pulumi:"location"`
	// Name of the authorization rule.
	Name *string `pulumi:"name"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The queue name.
	QueueName string `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []AccessRights `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

// The set of arguments for constructing a QueueAuthorizationRule resource.
type QueueAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringPtrInput
	// A string that describes Claim Type for authorization rule.
	ClaimType pulumi.StringPtrInput
	// A string that describes Claim Value of authorization rule.
	ClaimValue pulumi.StringPtrInput
	// A string that describes the Key Name of authorization rule.
	KeyName pulumi.StringPtrInput
	// data center location.
	Location pulumi.StringPtrInput
	// Name of the authorization rule.
	Name pulumi.StringPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringPtrInput
	// The queue name.
	QueueName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights AccessRightsArrayInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringPtrInput
}

func (QueueAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleArgs)(nil)).Elem()
}

type QueueAuthorizationRuleInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput
	ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput
}

func (*QueueAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil)).Elem()
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return i.ToQueueAuthorizationRuleOutputWithContext(context.Background())
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleOutput)
}

type QueueAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil)).Elem()
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return o
}

// A string that describes Claim Type for authorization rule.
func (o QueueAuthorizationRuleOutput) ClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.ClaimType }).(pulumi.StringPtrOutput)
}

// A string that describes Claim Value of authorization rule.
func (o QueueAuthorizationRuleOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

// The time the namespace was created.
func (o QueueAuthorizationRuleOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// A string that describes the Key Name of authorization rule.
func (o QueueAuthorizationRuleOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.KeyName }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o QueueAuthorizationRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The time the namespace was updated.
func (o QueueAuthorizationRuleOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.ModifiedTime }).(pulumi.StringOutput)
}

// Resource name
func (o QueueAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o QueueAuthorizationRuleOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// The rights associated with the rule.
func (o QueueAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o QueueAuthorizationRuleOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringPtrOutput { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

// Resource type
func (o QueueAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueAuthorizationRuleOutput{})
}