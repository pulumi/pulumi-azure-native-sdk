// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of subscription resource.
//
// Deprecated: Version 2014-09-01 will be removed in v2 of the provider.
type Subscription struct {
	pulumi.CustomResourceState

	// Last time there was a receive request to this subscription.
	AccessedAt pulumi.StringOutput `pulumi:"accessedAt"`
	// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrOutput `pulumi:"autoDeleteOnIdle"`
	// Message Count Details.
	CountDetails MessageCountDetailsResponseOutput `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions pulumi.BoolPtrOutput `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	// Value that indicates whether a subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrOutput `pulumi:"deadLetteringOnMessageExpiration"`
	// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrOutput `pulumi:"defaultMessageTimeToLive"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Entity availability status for the topic.
	EntityAvailabilityStatus pulumi.StringPtrOutput `pulumi:"entityAvailabilityStatus"`
	// Value that indicates whether the entity description is read-only.
	IsReadOnly pulumi.BoolPtrOutput `pulumi:"isReadOnly"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The lock duration time span for the subscription.
	LockDuration pulumi.StringPtrOutput `pulumi:"lockDuration"`
	// Number of maximum deliveries.
	MaxDeliveryCount pulumi.IntPtrOutput `pulumi:"maxDeliveryCount"`
	// Number of messages.
	MessageCount pulumi.Float64Output `pulumi:"messageCount"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Value indicating if a subscription supports the concept of sessions.
	RequiresSession pulumi.BoolPtrOutput `pulumi:"requiresSession"`
	// Enumerates the possible values for the status of a messaging entity.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:Subscription"),
		},
	})
	opts = append(opts, aliases)
	var resource Subscription
	err := ctx.RegisterResource("azure-native:servicebus/v20140901:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-native:servicebus/v20140901:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
}

type SubscriptionState struct {
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions *bool `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	// Value that indicates whether a subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration *bool `pulumi:"deadLetteringOnMessageExpiration"`
	// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive *string `pulumi:"defaultMessageTimeToLive"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Entity availability status for the topic.
	EntityAvailabilityStatus *EntityAvailabilityStatus `pulumi:"entityAvailabilityStatus"`
	// Value that indicates whether the entity description is read-only.
	IsReadOnly *bool `pulumi:"isReadOnly"`
	// Subscription data center location.
	Location *string `pulumi:"location"`
	// The lock duration time span for the subscription.
	LockDuration *string `pulumi:"lockDuration"`
	// Number of maximum deliveries.
	MaxDeliveryCount *int `pulumi:"maxDeliveryCount"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Value indicating if a subscription supports the concept of sessions.
	RequiresSession *bool `pulumi:"requiresSession"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus `pulumi:"status"`
	// The subscription name.
	SubscriptionName *string `pulumi:"subscriptionName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
	// Resource manager type of the resource.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions pulumi.BoolPtrInput
	// Value that indicates whether a subscription has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration pulumi.BoolPtrInput
	// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrInput
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Entity availability status for the topic.
	EntityAvailabilityStatus EntityAvailabilityStatusPtrInput
	// Value that indicates whether the entity description is read-only.
	IsReadOnly pulumi.BoolPtrInput
	// Subscription data center location.
	Location pulumi.StringPtrInput
	// The lock duration time span for the subscription.
	LockDuration pulumi.StringPtrInput
	// Number of maximum deliveries.
	MaxDeliveryCount pulumi.IntPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Value indicating if a subscription supports the concept of sessions.
	RequiresSession pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Enumerates the possible values for the status of a messaging entity.
	Status EntityStatusPtrInput
	// The subscription name.
	SubscriptionName pulumi.StringPtrInput
	// The topic name.
	TopicName pulumi.StringInput
	// Resource manager type of the resource.
	Type pulumi.StringPtrInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

// Last time there was a receive request to this subscription.
func (o SubscriptionOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.AccessedAt }).(pulumi.StringOutput)
}

// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
func (o SubscriptionOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

// Message Count Details.
func (o SubscriptionOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *Subscription) MessageCountDetailsResponseOutput { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

// Exact time the message was created.
func (o SubscriptionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
func (o SubscriptionOutput) DeadLetteringOnFilterEvaluationExceptions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.DeadLetteringOnFilterEvaluationExceptions }).(pulumi.BoolPtrOutput)
}

// Value that indicates whether a subscription has dead letter support when a message expires.
func (o SubscriptionOutput) DeadLetteringOnMessageExpiration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolPtrOutput)
}

// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
func (o SubscriptionOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

// Value that indicates whether server-side batched operations are enabled.
func (o SubscriptionOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

// Entity availability status for the topic.
func (o SubscriptionOutput) EntityAvailabilityStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.EntityAvailabilityStatus }).(pulumi.StringPtrOutput)
}

// Value that indicates whether the entity description is read-only.
func (o SubscriptionOutput) IsReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.IsReadOnly }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o SubscriptionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The lock duration time span for the subscription.
func (o SubscriptionOutput) LockDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.LockDuration }).(pulumi.StringPtrOutput)
}

// Number of maximum deliveries.
func (o SubscriptionOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.IntPtrOutput { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

// Number of messages.
func (o SubscriptionOutput) MessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Subscription) pulumi.Float64Output { return v.MessageCount }).(pulumi.Float64Output)
}

// Resource name
func (o SubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Value indicating if a subscription supports the concept of sessions.
func (o SubscriptionOutput) RequiresSession() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.RequiresSession }).(pulumi.BoolPtrOutput)
}

// Enumerates the possible values for the status of a messaging entity.
func (o SubscriptionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Resource type
func (o SubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o SubscriptionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}