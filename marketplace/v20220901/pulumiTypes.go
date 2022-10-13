// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// New plans notification details
type NewNotificationsResponse struct {
	// Gets offer display name
	DisplayName *string `pulumi:"displayName"`
	// Gets or sets the icon url
	Icon *string `pulumi:"icon"`
	// Gets a value indicating whether future plans is enabled.
	IsFuturePlansEnabled *bool `pulumi:"isFuturePlansEnabled"`
	// Gets or sets the notification message id
	MessageCode *float64 `pulumi:"messageCode"`
	// Gets offer id
	OfferId *string `pulumi:"offerId"`
	// Gets or sets removed plans notifications
	Plans []PlanNotificationDetailsResponse `pulumi:"plans"`
}

// New plans notification details
type NewNotificationsResponseOutput struct{ *pulumi.OutputState }

func (NewNotificationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NewNotificationsResponse)(nil)).Elem()
}

func (o NewNotificationsResponseOutput) ToNewNotificationsResponseOutput() NewNotificationsResponseOutput {
	return o
}

func (o NewNotificationsResponseOutput) ToNewNotificationsResponseOutputWithContext(ctx context.Context) NewNotificationsResponseOutput {
	return o
}

// Gets offer display name
func (o NewNotificationsResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Gets or sets the icon url
func (o NewNotificationsResponseOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *string { return v.Icon }).(pulumi.StringPtrOutput)
}

// Gets a value indicating whether future plans is enabled.
func (o NewNotificationsResponseOutput) IsFuturePlansEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *bool { return v.IsFuturePlansEnabled }).(pulumi.BoolPtrOutput)
}

// Gets or sets the notification message id
func (o NewNotificationsResponseOutput) MessageCode() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *float64 { return v.MessageCode }).(pulumi.Float64PtrOutput)
}

// Gets offer id
func (o NewNotificationsResponseOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

// Gets or sets removed plans notifications
func (o NewNotificationsResponseOutput) Plans() PlanNotificationDetailsResponseArrayOutput {
	return o.ApplyT(func(v NewNotificationsResponse) []PlanNotificationDetailsResponse { return v.Plans }).(PlanNotificationDetailsResponseArrayOutput)
}

type NewNotificationsResponseArrayOutput struct{ *pulumi.OutputState }

func (NewNotificationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NewNotificationsResponse)(nil)).Elem()
}

func (o NewNotificationsResponseArrayOutput) ToNewNotificationsResponseArrayOutput() NewNotificationsResponseArrayOutput {
	return o
}

func (o NewNotificationsResponseArrayOutput) ToNewNotificationsResponseArrayOutputWithContext(ctx context.Context) NewNotificationsResponseArrayOutput {
	return o
}

func (o NewNotificationsResponseArrayOutput) Index(i pulumi.IntInput) NewNotificationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NewNotificationsResponse {
		return vs[0].([]NewNotificationsResponse)[vs[1].(int)]
	}).(NewNotificationsResponseOutput)
}

type Plan struct {
	// Plan accessibility
	Accessibility *string `pulumi:"accessibility"`
}

// PlanInput is an input type that accepts PlanArgs and PlanOutput values.
// You can construct a concrete instance of `PlanInput` via:
//
//	PlanArgs{...}
type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	// Plan accessibility
	Accessibility pulumi.StringPtrInput `pulumi:"accessibility"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

// PlanArrayInput is an input type that accepts PlanArray and PlanArrayOutput values.
// You can construct a concrete instance of `PlanArrayInput` via:
//
//	PlanArray{ PlanArgs{...} }
type PlanArrayInput interface {
	pulumi.Input

	ToPlanArrayOutput() PlanArrayOutput
	ToPlanArrayOutputWithContext(context.Context) PlanArrayOutput
}

type PlanArray []PlanInput

func (PlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil)).Elem()
}

func (i PlanArray) ToPlanArrayOutput() PlanArrayOutput {
	return i.ToPlanArrayOutputWithContext(context.Background())
}

func (i PlanArray) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanArrayOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

// Plan accessibility
func (o PlanOutput) Accessibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Accessibility }).(pulumi.StringPtrOutput)
}

type PlanArrayOutput struct{ *pulumi.OutputState }

func (PlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil)).Elem()
}

func (o PlanArrayOutput) ToPlanArrayOutput() PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) Index(i pulumi.IntInput) PlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Plan {
		return vs[0].([]Plan)[vs[1].(int)]
	}).(PlanOutput)
}

// Plan notification details
type PlanNotificationDetailsResponse struct {
	// Gets or sets the plan display name
	PlanDisplayName *string `pulumi:"planDisplayName"`
	// Gets or sets the plan id
	PlanId *string `pulumi:"planId"`
}

// Plan notification details
type PlanNotificationDetailsResponseOutput struct{ *pulumi.OutputState }

func (PlanNotificationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanNotificationDetailsResponse)(nil)).Elem()
}

func (o PlanNotificationDetailsResponseOutput) ToPlanNotificationDetailsResponseOutput() PlanNotificationDetailsResponseOutput {
	return o
}

func (o PlanNotificationDetailsResponseOutput) ToPlanNotificationDetailsResponseOutputWithContext(ctx context.Context) PlanNotificationDetailsResponseOutput {
	return o
}

// Gets or sets the plan display name
func (o PlanNotificationDetailsResponseOutput) PlanDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanNotificationDetailsResponse) *string { return v.PlanDisplayName }).(pulumi.StringPtrOutput)
}

// Gets or sets the plan id
func (o PlanNotificationDetailsResponseOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanNotificationDetailsResponse) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

type PlanNotificationDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (PlanNotificationDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanNotificationDetailsResponse)(nil)).Elem()
}

func (o PlanNotificationDetailsResponseArrayOutput) ToPlanNotificationDetailsResponseArrayOutput() PlanNotificationDetailsResponseArrayOutput {
	return o
}

func (o PlanNotificationDetailsResponseArrayOutput) ToPlanNotificationDetailsResponseArrayOutputWithContext(ctx context.Context) PlanNotificationDetailsResponseArrayOutput {
	return o
}

func (o PlanNotificationDetailsResponseArrayOutput) Index(i pulumi.IntInput) PlanNotificationDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlanNotificationDetailsResponse {
		return vs[0].([]PlanNotificationDetailsResponse)[vs[1].(int)]
	}).(PlanNotificationDetailsResponseOutput)
}

type PlanResponse struct {
	// Plan accessibility
	Accessibility *string `pulumi:"accessibility"`
	// Alternative stack type
	AltStackReference string `pulumi:"altStackReference"`
	// Friendly name for the plan for display in the marketplace
	PlanDisplayName string `pulumi:"planDisplayName"`
	// Text identifier for this plan
	PlanId string `pulumi:"planId"`
	// Identifier for this plan
	SkuId string `pulumi:"skuId"`
	// Stack type (classic or arm)
	StackType string `pulumi:"stackType"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

// Plan accessibility
func (o PlanResponseOutput) Accessibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Accessibility }).(pulumi.StringPtrOutput)
}

// Alternative stack type
func (o PlanResponseOutput) AltStackReference() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.AltStackReference }).(pulumi.StringOutput)
}

// Friendly name for the plan for display in the marketplace
func (o PlanResponseOutput) PlanDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.PlanDisplayName }).(pulumi.StringOutput)
}

// Text identifier for this plan
func (o PlanResponseOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.PlanId }).(pulumi.StringOutput)
}

// Identifier for this plan
func (o PlanResponseOutput) SkuId() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.SkuId }).(pulumi.StringOutput)
}

// Stack type (classic or arm)
func (o PlanResponseOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.StackType }).(pulumi.StringOutput)
}

type PlanResponseArrayOutput struct{ *pulumi.OutputState }

func (PlanResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanResponse)(nil)).Elem()
}

func (o PlanResponseArrayOutput) ToPlanResponseArrayOutput() PlanResponseArrayOutput {
	return o
}

func (o PlanResponseArrayOutput) ToPlanResponseArrayOutputWithContext(ctx context.Context) PlanResponseArrayOutput {
	return o
}

func (o PlanResponseArrayOutput) Index(i pulumi.IntInput) PlanResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlanResponse {
		return vs[0].([]PlanResponse)[vs[1].(int)]
	}).(PlanResponseOutput)
}

type RuleResponse struct {
	// Rule type
	Type  *string  `pulumi:"type"`
	Value []string `pulumi:"value"`
}

type RuleResponseOutput struct{ *pulumi.OutputState }

func (RuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResponse)(nil)).Elem()
}

func (o RuleResponseOutput) ToRuleResponseOutput() RuleResponseOutput {
	return o
}

func (o RuleResponseOutput) ToRuleResponseOutputWithContext(ctx context.Context) RuleResponseOutput {
	return o
}

// Rule type
func (o RuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o RuleResponseOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleResponse) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type RuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleResponse)(nil)).Elem()
}

func (o RuleResponseArrayOutput) ToRuleResponseArrayOutput() RuleResponseArrayOutput {
	return o
}

func (o RuleResponseArrayOutput) ToRuleResponseArrayOutputWithContext(ctx context.Context) RuleResponseArrayOutput {
	return o
}

func (o RuleResponseArrayOutput) Index(i pulumi.IntInput) RuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleResponse {
		return vs[0].([]RuleResponse)[vs[1].(int)]
	}).(RuleResponseOutput)
}

// List of stop sell offers and plans notifications.
type StopSellOffersPlansNotificationsListPropertiesResponse struct {
	// The offer display name
	DisplayName string `pulumi:"displayName"`
	// The icon url
	Icon string `pulumi:"icon"`
	// A value indicating whether entire offer is in stop sell or only few of its plans
	IsEntire bool `pulumi:"isEntire"`
	// The notification message code
	MessageCode float64 `pulumi:"messageCode"`
	// The offer id
	OfferId string `pulumi:"offerId"`
	// The list of removed plans notifications
	Plans []PlanNotificationDetailsResponse `pulumi:"plans"`
	// True if the offer has public plans
	PublicContext bool `pulumi:"publicContext"`
	// The subscriptions related to private plans
	SubscriptionsIds []string `pulumi:"subscriptionsIds"`
}

// List of stop sell offers and plans notifications.
type StopSellOffersPlansNotificationsListPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StopSellOffersPlansNotificationsListPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StopSellOffersPlansNotificationsListPropertiesResponse)(nil)).Elem()
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseOutput() StopSellOffersPlansNotificationsListPropertiesResponseOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseOutputWithContext(ctx context.Context) StopSellOffersPlansNotificationsListPropertiesResponseOutput {
	return o
}

// The offer display name
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The icon url
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) string { return v.Icon }).(pulumi.StringOutput)
}

// A value indicating whether entire offer is in stop sell or only few of its plans
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) IsEntire() pulumi.BoolOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) bool { return v.IsEntire }).(pulumi.BoolOutput)
}

// The notification message code
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) MessageCode() pulumi.Float64Output {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) float64 { return v.MessageCode }).(pulumi.Float64Output)
}

// The offer id
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) string { return v.OfferId }).(pulumi.StringOutput)
}

// The list of removed plans notifications
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) Plans() PlanNotificationDetailsResponseArrayOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) []PlanNotificationDetailsResponse {
		return v.Plans
	}).(PlanNotificationDetailsResponseArrayOutput)
}

// True if the offer has public plans
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) PublicContext() pulumi.BoolOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) bool { return v.PublicContext }).(pulumi.BoolOutput)
}

// The subscriptions related to private plans
func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) SubscriptionsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) []string { return v.SubscriptionsIds }).(pulumi.StringArrayOutput)
}

type StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StopSellOffersPlansNotificationsListPropertiesResponse)(nil)).Elem()
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseArrayOutput() StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseArrayOutputWithContext(ctx context.Context) StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) Index(i pulumi.IntInput) StopSellOffersPlansNotificationsListPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StopSellOffersPlansNotificationsListPropertiesResponse {
		return vs[0].([]StopSellOffersPlansNotificationsListPropertiesResponse)[vs[1].(int)]
	}).(StopSellOffersPlansNotificationsListPropertiesResponseOutput)
}

// Read only system data
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC)
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Read only system data
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC)
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NewNotificationsResponseOutput{})
	pulumi.RegisterOutputType(NewNotificationsResponseArrayOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanArrayOutput{})
	pulumi.RegisterOutputType(PlanNotificationDetailsResponseOutput{})
	pulumi.RegisterOutputType(PlanNotificationDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponseArrayOutput{})
	pulumi.RegisterOutputType(RuleResponseOutput{})
	pulumi.RegisterOutputType(RuleResponseArrayOutput{})
	pulumi.RegisterOutputType(StopSellOffersPlansNotificationsListPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}