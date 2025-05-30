// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package professionalservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// properties for creation professionalService
type ProfessionalServiceCreationProperties struct {
	// Whether the ProfessionalService subscription will auto renew upon term end.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The billing period eg P1M,P1Y for monthly,yearly respectively
	BillingPeriod *string `pulumi:"billingPeriod"`
	// The offer id.
	OfferId *string `pulumi:"offerId"`
	// The publisher id.
	PublisherId *string `pulumi:"publisherId"`
	// The quote id which the ProfessionalService will be purchase with.
	QuoteId *string `pulumi:"quoteId"`
	// The plan id.
	SkuId *string `pulumi:"skuId"`
	// The store front which initiates the purchase.
	StoreFront *string `pulumi:"storeFront"`
	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit *string `pulumi:"termUnit"`
}

// ProfessionalServiceCreationPropertiesInput is an input type that accepts ProfessionalServiceCreationPropertiesArgs and ProfessionalServiceCreationPropertiesOutput values.
// You can construct a concrete instance of `ProfessionalServiceCreationPropertiesInput` via:
//
//	ProfessionalServiceCreationPropertiesArgs{...}
type ProfessionalServiceCreationPropertiesInput interface {
	pulumi.Input

	ToProfessionalServiceCreationPropertiesOutput() ProfessionalServiceCreationPropertiesOutput
	ToProfessionalServiceCreationPropertiesOutputWithContext(context.Context) ProfessionalServiceCreationPropertiesOutput
}

// properties for creation professionalService
type ProfessionalServiceCreationPropertiesArgs struct {
	// Whether the ProfessionalService subscription will auto renew upon term end.
	AutoRenew pulumi.BoolPtrInput `pulumi:"autoRenew"`
	// The billing period eg P1M,P1Y for monthly,yearly respectively
	BillingPeriod pulumi.StringPtrInput `pulumi:"billingPeriod"`
	// The offer id.
	OfferId pulumi.StringPtrInput `pulumi:"offerId"`
	// The publisher id.
	PublisherId pulumi.StringPtrInput `pulumi:"publisherId"`
	// The quote id which the ProfessionalService will be purchase with.
	QuoteId pulumi.StringPtrInput `pulumi:"quoteId"`
	// The plan id.
	SkuId pulumi.StringPtrInput `pulumi:"skuId"`
	// The store front which initiates the purchase.
	StoreFront pulumi.StringPtrInput `pulumi:"storeFront"`
	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit pulumi.StringPtrInput `pulumi:"termUnit"`
}

func (ProfessionalServiceCreationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfessionalServiceCreationProperties)(nil)).Elem()
}

func (i ProfessionalServiceCreationPropertiesArgs) ToProfessionalServiceCreationPropertiesOutput() ProfessionalServiceCreationPropertiesOutput {
	return i.ToProfessionalServiceCreationPropertiesOutputWithContext(context.Background())
}

func (i ProfessionalServiceCreationPropertiesArgs) ToProfessionalServiceCreationPropertiesOutputWithContext(ctx context.Context) ProfessionalServiceCreationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfessionalServiceCreationPropertiesOutput)
}

func (i ProfessionalServiceCreationPropertiesArgs) ToProfessionalServiceCreationPropertiesPtrOutput() ProfessionalServiceCreationPropertiesPtrOutput {
	return i.ToProfessionalServiceCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i ProfessionalServiceCreationPropertiesArgs) ToProfessionalServiceCreationPropertiesPtrOutputWithContext(ctx context.Context) ProfessionalServiceCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfessionalServiceCreationPropertiesOutput).ToProfessionalServiceCreationPropertiesPtrOutputWithContext(ctx)
}

// ProfessionalServiceCreationPropertiesPtrInput is an input type that accepts ProfessionalServiceCreationPropertiesArgs, ProfessionalServiceCreationPropertiesPtr and ProfessionalServiceCreationPropertiesPtrOutput values.
// You can construct a concrete instance of `ProfessionalServiceCreationPropertiesPtrInput` via:
//
//	        ProfessionalServiceCreationPropertiesArgs{...}
//
//	or:
//
//	        nil
type ProfessionalServiceCreationPropertiesPtrInput interface {
	pulumi.Input

	ToProfessionalServiceCreationPropertiesPtrOutput() ProfessionalServiceCreationPropertiesPtrOutput
	ToProfessionalServiceCreationPropertiesPtrOutputWithContext(context.Context) ProfessionalServiceCreationPropertiesPtrOutput
}

type professionalServiceCreationPropertiesPtrType ProfessionalServiceCreationPropertiesArgs

func ProfessionalServiceCreationPropertiesPtr(v *ProfessionalServiceCreationPropertiesArgs) ProfessionalServiceCreationPropertiesPtrInput {
	return (*professionalServiceCreationPropertiesPtrType)(v)
}

func (*professionalServiceCreationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfessionalServiceCreationProperties)(nil)).Elem()
}

func (i *professionalServiceCreationPropertiesPtrType) ToProfessionalServiceCreationPropertiesPtrOutput() ProfessionalServiceCreationPropertiesPtrOutput {
	return i.ToProfessionalServiceCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i *professionalServiceCreationPropertiesPtrType) ToProfessionalServiceCreationPropertiesPtrOutputWithContext(ctx context.Context) ProfessionalServiceCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfessionalServiceCreationPropertiesPtrOutput)
}

// properties for creation professionalService
type ProfessionalServiceCreationPropertiesOutput struct{ *pulumi.OutputState }

func (ProfessionalServiceCreationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfessionalServiceCreationProperties)(nil)).Elem()
}

func (o ProfessionalServiceCreationPropertiesOutput) ToProfessionalServiceCreationPropertiesOutput() ProfessionalServiceCreationPropertiesOutput {
	return o
}

func (o ProfessionalServiceCreationPropertiesOutput) ToProfessionalServiceCreationPropertiesOutputWithContext(ctx context.Context) ProfessionalServiceCreationPropertiesOutput {
	return o
}

func (o ProfessionalServiceCreationPropertiesOutput) ToProfessionalServiceCreationPropertiesPtrOutput() ProfessionalServiceCreationPropertiesPtrOutput {
	return o.ToProfessionalServiceCreationPropertiesPtrOutputWithContext(context.Background())
}

func (o ProfessionalServiceCreationPropertiesOutput) ToProfessionalServiceCreationPropertiesPtrOutputWithContext(ctx context.Context) ProfessionalServiceCreationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfessionalServiceCreationProperties) *ProfessionalServiceCreationProperties {
		return &v
	}).(ProfessionalServiceCreationPropertiesPtrOutput)
}

// Whether the ProfessionalService subscription will auto renew upon term end.
func (o ProfessionalServiceCreationPropertiesOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The billing period eg P1M,P1Y for monthly,yearly respectively
func (o ProfessionalServiceCreationPropertiesOutput) BillingPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.BillingPeriod }).(pulumi.StringPtrOutput)
}

// The offer id.
func (o ProfessionalServiceCreationPropertiesOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

// The publisher id.
func (o ProfessionalServiceCreationPropertiesOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

// The quote id which the ProfessionalService will be purchase with.
func (o ProfessionalServiceCreationPropertiesOutput) QuoteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.QuoteId }).(pulumi.StringPtrOutput)
}

// The plan id.
func (o ProfessionalServiceCreationPropertiesOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.SkuId }).(pulumi.StringPtrOutput)
}

// The store front which initiates the purchase.
func (o ProfessionalServiceCreationPropertiesOutput) StoreFront() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.StoreFront }).(pulumi.StringPtrOutput)
}

// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
func (o ProfessionalServiceCreationPropertiesOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceCreationProperties) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type ProfessionalServiceCreationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ProfessionalServiceCreationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfessionalServiceCreationProperties)(nil)).Elem()
}

func (o ProfessionalServiceCreationPropertiesPtrOutput) ToProfessionalServiceCreationPropertiesPtrOutput() ProfessionalServiceCreationPropertiesPtrOutput {
	return o
}

func (o ProfessionalServiceCreationPropertiesPtrOutput) ToProfessionalServiceCreationPropertiesPtrOutputWithContext(ctx context.Context) ProfessionalServiceCreationPropertiesPtrOutput {
	return o
}

func (o ProfessionalServiceCreationPropertiesPtrOutput) Elem() ProfessionalServiceCreationPropertiesOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) ProfessionalServiceCreationProperties {
		if v != nil {
			return *v
		}
		var ret ProfessionalServiceCreationProperties
		return ret
	}).(ProfessionalServiceCreationPropertiesOutput)
}

// Whether the ProfessionalService subscription will auto renew upon term end.
func (o ProfessionalServiceCreationPropertiesPtrOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AutoRenew
	}).(pulumi.BoolPtrOutput)
}

// The billing period eg P1M,P1Y for monthly,yearly respectively
func (o ProfessionalServiceCreationPropertiesPtrOutput) BillingPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.BillingPeriod
	}).(pulumi.StringPtrOutput)
}

// The offer id.
func (o ProfessionalServiceCreationPropertiesPtrOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.OfferId
	}).(pulumi.StringPtrOutput)
}

// The publisher id.
func (o ProfessionalServiceCreationPropertiesPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublisherId
	}).(pulumi.StringPtrOutput)
}

// The quote id which the ProfessionalService will be purchase with.
func (o ProfessionalServiceCreationPropertiesPtrOutput) QuoteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.QuoteId
	}).(pulumi.StringPtrOutput)
}

// The plan id.
func (o ProfessionalServiceCreationPropertiesPtrOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SkuId
	}).(pulumi.StringPtrOutput)
}

// The store front which initiates the purchase.
func (o ProfessionalServiceCreationPropertiesPtrOutput) StoreFront() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.StoreFront
	}).(pulumi.StringPtrOutput)
}

// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
func (o ProfessionalServiceCreationPropertiesPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServiceCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

// The current Term object.
type ProfessionalServicePropertiesResponseTerm struct {
	// The end date of the current term
	EndDate *string `pulumi:"endDate"`
	// The start date of the current term
	StartDate *string `pulumi:"startDate"`
	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit *string `pulumi:"termUnit"`
}

// The current Term object.
type ProfessionalServicePropertiesResponseTermOutput struct{ *pulumi.OutputState }

func (ProfessionalServicePropertiesResponseTermOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfessionalServicePropertiesResponseTerm)(nil)).Elem()
}

func (o ProfessionalServicePropertiesResponseTermOutput) ToProfessionalServicePropertiesResponseTermOutput() ProfessionalServicePropertiesResponseTermOutput {
	return o
}

func (o ProfessionalServicePropertiesResponseTermOutput) ToProfessionalServicePropertiesResponseTermOutputWithContext(ctx context.Context) ProfessionalServicePropertiesResponseTermOutput {
	return o
}

// The end date of the current term
func (o ProfessionalServicePropertiesResponseTermOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServicePropertiesResponseTerm) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// The start date of the current term
func (o ProfessionalServicePropertiesResponseTermOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServicePropertiesResponseTerm) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
func (o ProfessionalServicePropertiesResponseTermOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServicePropertiesResponseTerm) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type ProfessionalServicePropertiesResponseTermPtrOutput struct{ *pulumi.OutputState }

func (ProfessionalServicePropertiesResponseTermPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfessionalServicePropertiesResponseTerm)(nil)).Elem()
}

func (o ProfessionalServicePropertiesResponseTermPtrOutput) ToProfessionalServicePropertiesResponseTermPtrOutput() ProfessionalServicePropertiesResponseTermPtrOutput {
	return o
}

func (o ProfessionalServicePropertiesResponseTermPtrOutput) ToProfessionalServicePropertiesResponseTermPtrOutputWithContext(ctx context.Context) ProfessionalServicePropertiesResponseTermPtrOutput {
	return o
}

func (o ProfessionalServicePropertiesResponseTermPtrOutput) Elem() ProfessionalServicePropertiesResponseTermOutput {
	return o.ApplyT(func(v *ProfessionalServicePropertiesResponseTerm) ProfessionalServicePropertiesResponseTerm {
		if v != nil {
			return *v
		}
		var ret ProfessionalServicePropertiesResponseTerm
		return ret
	}).(ProfessionalServicePropertiesResponseTermOutput)
}

// The end date of the current term
func (o ProfessionalServicePropertiesResponseTermPtrOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServicePropertiesResponseTerm) *string {
		if v == nil {
			return nil
		}
		return v.EndDate
	}).(pulumi.StringPtrOutput)
}

// The start date of the current term
func (o ProfessionalServicePropertiesResponseTermPtrOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServicePropertiesResponseTerm) *string {
		if v == nil {
			return nil
		}
		return v.StartDate
	}).(pulumi.StringPtrOutput)
}

// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
func (o ProfessionalServicePropertiesResponseTermPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfessionalServicePropertiesResponseTerm) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

// professionalService properties
type ProfessionalServiceResourceResponseProperties struct {
	// Whether the ProfessionalService subscription will auto renew upon term end.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The billing period eg P1M,P1Y for monthly,yearly respectively
	BillingPeriod *string `pulumi:"billingPeriod"`
	// The created date of this resource.
	Created string `pulumi:"created"`
	// Whether the current term is a Free Trial term
	IsFreeTrial *bool `pulumi:"isFreeTrial"`
	// The last modifier date if this resource.
	LastModified *string `pulumi:"lastModified"`
	// The offer id.
	OfferId *string `pulumi:"offerId"`
	// The metadata about the ProfessionalService subscription such as the AzureSubscriptionId and ResourceUri.
	PaymentChannelMetadata map[string]string `pulumi:"paymentChannelMetadata"`
	// The Payment channel for the ProfessionalServiceSubscription.
	PaymentChannelType *string `pulumi:"paymentChannelType"`
	// The publisher id.
	PublisherId *string `pulumi:"publisherId"`
	// The quote id which the ProfessionalService will be purchase with.
	QuoteId *string `pulumi:"quoteId"`
	// The plan id.
	SkuId *string `pulumi:"skuId"`
	// The ProfessionalService Subscription Status.
	Status *string `pulumi:"status"`
	// The store front which initiates the purchase.
	StoreFront *string `pulumi:"storeFront"`
	// The current Term object.
	Term *ProfessionalServicePropertiesResponseTerm `pulumi:"term"`
	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit *string `pulumi:"termUnit"`
}

// professionalService properties
type ProfessionalServiceResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ProfessionalServiceResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfessionalServiceResourceResponseProperties)(nil)).Elem()
}

func (o ProfessionalServiceResourceResponsePropertiesOutput) ToProfessionalServiceResourceResponsePropertiesOutput() ProfessionalServiceResourceResponsePropertiesOutput {
	return o
}

func (o ProfessionalServiceResourceResponsePropertiesOutput) ToProfessionalServiceResourceResponsePropertiesOutputWithContext(ctx context.Context) ProfessionalServiceResourceResponsePropertiesOutput {
	return o
}

// Whether the ProfessionalService subscription will auto renew upon term end.
func (o ProfessionalServiceResourceResponsePropertiesOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The billing period eg P1M,P1Y for monthly,yearly respectively
func (o ProfessionalServiceResourceResponsePropertiesOutput) BillingPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.BillingPeriod }).(pulumi.StringPtrOutput)
}

// The created date of this resource.
func (o ProfessionalServiceResourceResponsePropertiesOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) string { return v.Created }).(pulumi.StringOutput)
}

// Whether the current term is a Free Trial term
func (o ProfessionalServiceResourceResponsePropertiesOutput) IsFreeTrial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *bool { return v.IsFreeTrial }).(pulumi.BoolPtrOutput)
}

// The last modifier date if this resource.
func (o ProfessionalServiceResourceResponsePropertiesOutput) LastModified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.LastModified }).(pulumi.StringPtrOutput)
}

// The offer id.
func (o ProfessionalServiceResourceResponsePropertiesOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

// The metadata about the ProfessionalService subscription such as the AzureSubscriptionId and ResourceUri.
func (o ProfessionalServiceResourceResponsePropertiesOutput) PaymentChannelMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) map[string]string {
		return v.PaymentChannelMetadata
	}).(pulumi.StringMapOutput)
}

// The Payment channel for the ProfessionalServiceSubscription.
func (o ProfessionalServiceResourceResponsePropertiesOutput) PaymentChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.PaymentChannelType }).(pulumi.StringPtrOutput)
}

// The publisher id.
func (o ProfessionalServiceResourceResponsePropertiesOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

// The quote id which the ProfessionalService will be purchase with.
func (o ProfessionalServiceResourceResponsePropertiesOutput) QuoteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.QuoteId }).(pulumi.StringPtrOutput)
}

// The plan id.
func (o ProfessionalServiceResourceResponsePropertiesOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.SkuId }).(pulumi.StringPtrOutput)
}

// The ProfessionalService Subscription Status.
func (o ProfessionalServiceResourceResponsePropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The store front which initiates the purchase.
func (o ProfessionalServiceResourceResponsePropertiesOutput) StoreFront() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.StoreFront }).(pulumi.StringPtrOutput)
}

// The current Term object.
func (o ProfessionalServiceResourceResponsePropertiesOutput) Term() ProfessionalServicePropertiesResponseTermPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *ProfessionalServicePropertiesResponseTerm {
		return v.Term
	}).(ProfessionalServicePropertiesResponseTermPtrOutput)
}

// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
func (o ProfessionalServiceResourceResponsePropertiesOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfessionalServiceResourceResponseProperties) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfessionalServiceCreationPropertiesOutput{})
	pulumi.RegisterOutputType(ProfessionalServiceCreationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProfessionalServicePropertiesResponseTermOutput{})
	pulumi.RegisterOutputType(ProfessionalServicePropertiesResponseTermPtrOutput{})
	pulumi.RegisterOutputType(ProfessionalServiceResourceResponsePropertiesOutput{})
}
