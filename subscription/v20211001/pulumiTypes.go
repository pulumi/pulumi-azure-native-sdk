// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Put subscription additional properties.
type PutAliasRequestAdditionalProperties struct {
	// Management group Id for the subscription.
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// Owner Id of the subscription
	SubscriptionOwnerId *string `pulumi:"subscriptionOwnerId"`
	// Tenant Id of the subscription
	SubscriptionTenantId *string `pulumi:"subscriptionTenantId"`
	// Tags for the subscription
	Tags map[string]string `pulumi:"tags"`
}

// PutAliasRequestAdditionalPropertiesInput is an input type that accepts PutAliasRequestAdditionalPropertiesArgs and PutAliasRequestAdditionalPropertiesOutput values.
// You can construct a concrete instance of `PutAliasRequestAdditionalPropertiesInput` via:
//
//	PutAliasRequestAdditionalPropertiesArgs{...}
type PutAliasRequestAdditionalPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestAdditionalPropertiesOutput() PutAliasRequestAdditionalPropertiesOutput
	ToPutAliasRequestAdditionalPropertiesOutputWithContext(context.Context) PutAliasRequestAdditionalPropertiesOutput
}

// Put subscription additional properties.
type PutAliasRequestAdditionalPropertiesArgs struct {
	// Management group Id for the subscription.
	ManagementGroupId pulumi.StringPtrInput `pulumi:"managementGroupId"`
	// Owner Id of the subscription
	SubscriptionOwnerId pulumi.StringPtrInput `pulumi:"subscriptionOwnerId"`
	// Tenant Id of the subscription
	SubscriptionTenantId pulumi.StringPtrInput `pulumi:"subscriptionTenantId"`
	// Tags for the subscription
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (PutAliasRequestAdditionalPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesOutput() PutAliasRequestAdditionalPropertiesOutput {
	return i.ToPutAliasRequestAdditionalPropertiesOutputWithContext(context.Background())
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestAdditionalPropertiesOutput)
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return i.ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Background())
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestAdditionalPropertiesOutput).ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx)
}

// PutAliasRequestAdditionalPropertiesPtrInput is an input type that accepts PutAliasRequestAdditionalPropertiesArgs, PutAliasRequestAdditionalPropertiesPtr and PutAliasRequestAdditionalPropertiesPtrOutput values.
// You can construct a concrete instance of `PutAliasRequestAdditionalPropertiesPtrInput` via:
//
//	        PutAliasRequestAdditionalPropertiesArgs{...}
//
//	or:
//
//	        nil
type PutAliasRequestAdditionalPropertiesPtrInput interface {
	pulumi.Input

	ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput
	ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Context) PutAliasRequestAdditionalPropertiesPtrOutput
}

type putAliasRequestAdditionalPropertiesPtrType PutAliasRequestAdditionalPropertiesArgs

func PutAliasRequestAdditionalPropertiesPtr(v *PutAliasRequestAdditionalPropertiesArgs) PutAliasRequestAdditionalPropertiesPtrInput {
	return (*putAliasRequestAdditionalPropertiesPtrType)(v)
}

func (*putAliasRequestAdditionalPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (i *putAliasRequestAdditionalPropertiesPtrType) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return i.ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Background())
}

func (i *putAliasRequestAdditionalPropertiesPtrType) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

// Put subscription additional properties.
type PutAliasRequestAdditionalPropertiesOutput struct{ *pulumi.OutputState }

func (PutAliasRequestAdditionalPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesOutput() PutAliasRequestAdditionalPropertiesOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Background())
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PutAliasRequestAdditionalProperties) *PutAliasRequestAdditionalProperties {
		return &v
	}).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

// Management group Id for the subscription.
func (o PutAliasRequestAdditionalPropertiesOutput) ManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) *string { return v.ManagementGroupId }).(pulumi.StringPtrOutput)
}

// Owner Id of the subscription
func (o PutAliasRequestAdditionalPropertiesOutput) SubscriptionOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) *string { return v.SubscriptionOwnerId }).(pulumi.StringPtrOutput)
}

// Tenant Id of the subscription
func (o PutAliasRequestAdditionalPropertiesOutput) SubscriptionTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) *string { return v.SubscriptionTenantId }).(pulumi.StringPtrOutput)
}

// Tags for the subscription
func (o PutAliasRequestAdditionalPropertiesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type PutAliasRequestAdditionalPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PutAliasRequestAdditionalPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) Elem() PutAliasRequestAdditionalPropertiesOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) PutAliasRequestAdditionalProperties {
		if v != nil {
			return *v
		}
		var ret PutAliasRequestAdditionalProperties
		return ret
	}).(PutAliasRequestAdditionalPropertiesOutput)
}

// Management group Id for the subscription.
func (o PutAliasRequestAdditionalPropertiesPtrOutput) ManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagementGroupId
	}).(pulumi.StringPtrOutput)
}

// Owner Id of the subscription
func (o PutAliasRequestAdditionalPropertiesPtrOutput) SubscriptionOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionOwnerId
	}).(pulumi.StringPtrOutput)
}

// Tenant Id of the subscription
func (o PutAliasRequestAdditionalPropertiesPtrOutput) SubscriptionTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionTenantId
	}).(pulumi.StringPtrOutput)
}

// Tags for the subscription
func (o PutAliasRequestAdditionalPropertiesPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

// Put subscription properties.
type PutAliasRequestProperties struct {
	// Put alias request additional properties.
	AdditionalProperties *PutAliasRequestAdditionalProperties `pulumi:"additionalProperties"`
	// Billing scope of the subscription.
	// For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
	// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope *string `pulumi:"billingScope"`
	// The friendly name of the subscription.
	DisplayName *string `pulumi:"displayName"`
	// Reseller Id
	ResellerId *string `pulumi:"resellerId"`
	// This parameter can be used to create alias for existing subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The workload type of the subscription. It can be either Production or DevTest.
	Workload *string `pulumi:"workload"`
}

// PutAliasRequestPropertiesInput is an input type that accepts PutAliasRequestPropertiesArgs and PutAliasRequestPropertiesOutput values.
// You can construct a concrete instance of `PutAliasRequestPropertiesInput` via:
//
//	PutAliasRequestPropertiesArgs{...}
type PutAliasRequestPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput
	ToPutAliasRequestPropertiesOutputWithContext(context.Context) PutAliasRequestPropertiesOutput
}

// Put subscription properties.
type PutAliasRequestPropertiesArgs struct {
	// Put alias request additional properties.
	AdditionalProperties PutAliasRequestAdditionalPropertiesPtrInput `pulumi:"additionalProperties"`
	// Billing scope of the subscription.
	// For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
	// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope pulumi.StringPtrInput `pulumi:"billingScope"`
	// The friendly name of the subscription.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Reseller Id
	ResellerId pulumi.StringPtrInput `pulumi:"resellerId"`
	// This parameter can be used to create alias for existing subscription Id
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	// The workload type of the subscription. It can be either Production or DevTest.
	Workload pulumi.StringPtrInput `pulumi:"workload"`
}

func (PutAliasRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return i.ToPutAliasRequestPropertiesOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput)
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput).ToPutAliasRequestPropertiesPtrOutputWithContext(ctx)
}

// PutAliasRequestPropertiesPtrInput is an input type that accepts PutAliasRequestPropertiesArgs, PutAliasRequestPropertiesPtr and PutAliasRequestPropertiesPtrOutput values.
// You can construct a concrete instance of `PutAliasRequestPropertiesPtrInput` via:
//
//	        PutAliasRequestPropertiesArgs{...}
//
//	or:
//
//	        nil
type PutAliasRequestPropertiesPtrInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput
	ToPutAliasRequestPropertiesPtrOutputWithContext(context.Context) PutAliasRequestPropertiesPtrOutput
}

type putAliasRequestPropertiesPtrType PutAliasRequestPropertiesArgs

func PutAliasRequestPropertiesPtr(v *PutAliasRequestPropertiesArgs) PutAliasRequestPropertiesPtrInput {
	return (*putAliasRequestPropertiesPtrType)(v)
}

func (*putAliasRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesPtrOutput)
}

// Put subscription properties.
type PutAliasRequestPropertiesOutput struct{ *pulumi.OutputState }

func (PutAliasRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return o
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return o
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return o.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PutAliasRequestProperties) *PutAliasRequestProperties {
		return &v
	}).(PutAliasRequestPropertiesPtrOutput)
}

// Put alias request additional properties.
func (o PutAliasRequestPropertiesOutput) AdditionalProperties() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *PutAliasRequestAdditionalProperties { return v.AdditionalProperties }).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

// Billing scope of the subscription.
// For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
func (o PutAliasRequestPropertiesOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.BillingScope }).(pulumi.StringPtrOutput)
}

// The friendly name of the subscription.
func (o PutAliasRequestPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Reseller Id
func (o PutAliasRequestPropertiesOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.ResellerId }).(pulumi.StringPtrOutput)
}

// This parameter can be used to create alias for existing subscription Id
func (o PutAliasRequestPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// The workload type of the subscription. It can be either Production or DevTest.
func (o PutAliasRequestPropertiesOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.Workload }).(pulumi.StringPtrOutput)
}

type PutAliasRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PutAliasRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (o PutAliasRequestPropertiesPtrOutput) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestPropertiesPtrOutput) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestPropertiesPtrOutput) Elem() PutAliasRequestPropertiesOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) PutAliasRequestProperties {
		if v != nil {
			return *v
		}
		var ret PutAliasRequestProperties
		return ret
	}).(PutAliasRequestPropertiesOutput)
}

// Put alias request additional properties.
func (o PutAliasRequestPropertiesPtrOutput) AdditionalProperties() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *PutAliasRequestAdditionalProperties {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

// Billing scope of the subscription.
// For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
func (o PutAliasRequestPropertiesPtrOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.BillingScope
	}).(pulumi.StringPtrOutput)
}

// The friendly name of the subscription.
func (o PutAliasRequestPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// Reseller Id
func (o PutAliasRequestPropertiesPtrOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResellerId
	}).(pulumi.StringPtrOutput)
}

// This parameter can be used to create alias for existing subscription Id
func (o PutAliasRequestPropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

// The workload type of the subscription. It can be either Production or DevTest.
func (o PutAliasRequestPropertiesPtrOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.Workload
	}).(pulumi.StringPtrOutput)
}

// Put subscription creation result properties.
type SubscriptionAliasResponsePropertiesResponse struct {
	// The accept ownership state of the resource.
	AcceptOwnershipState string `pulumi:"acceptOwnershipState"`
	// Url to accept ownership of the subscription.
	AcceptOwnershipUrl string `pulumi:"acceptOwnershipUrl"`
	// Billing scope of the subscription.
	// For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
	// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope *string `pulumi:"billingScope"`
	// Created Time
	CreatedTime *string `pulumi:"createdTime"`
	// The display name of the subscription.
	DisplayName *string `pulumi:"displayName"`
	// The Management Group Id.
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Reseller Id
	ResellerId *string `pulumi:"resellerId"`
	// Newly created subscription Id.
	SubscriptionId string `pulumi:"subscriptionId"`
	// Owner Id of the subscription
	SubscriptionOwnerId *string `pulumi:"subscriptionOwnerId"`
	// Tags for the subscription
	Tags map[string]string `pulumi:"tags"`
	// The workload type of the subscription. It can be either Production or DevTest.
	Workload *string `pulumi:"workload"`
}

// Put subscription creation result properties.
type SubscriptionAliasResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionAliasResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionAliasResponsePropertiesResponse)(nil)).Elem()
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ToSubscriptionAliasResponsePropertiesResponseOutput() SubscriptionAliasResponsePropertiesResponseOutput {
	return o
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ToSubscriptionAliasResponsePropertiesResponseOutputWithContext(ctx context.Context) SubscriptionAliasResponsePropertiesResponseOutput {
	return o
}

// The accept ownership state of the resource.
func (o SubscriptionAliasResponsePropertiesResponseOutput) AcceptOwnershipState() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) string { return v.AcceptOwnershipState }).(pulumi.StringOutput)
}

// Url to accept ownership of the subscription.
func (o SubscriptionAliasResponsePropertiesResponseOutput) AcceptOwnershipUrl() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) string { return v.AcceptOwnershipUrl }).(pulumi.StringOutput)
}

// Billing scope of the subscription.
// For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
func (o SubscriptionAliasResponsePropertiesResponseOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.BillingScope }).(pulumi.StringPtrOutput)
}

// Created Time
func (o SubscriptionAliasResponsePropertiesResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// The display name of the subscription.
func (o SubscriptionAliasResponsePropertiesResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The Management Group Id.
func (o SubscriptionAliasResponsePropertiesResponseOutput) ManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.ManagementGroupId }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o SubscriptionAliasResponsePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Reseller Id
func (o SubscriptionAliasResponsePropertiesResponseOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.ResellerId }).(pulumi.StringPtrOutput)
}

// Newly created subscription Id.
func (o SubscriptionAliasResponsePropertiesResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Owner Id of the subscription
func (o SubscriptionAliasResponsePropertiesResponseOutput) SubscriptionOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.SubscriptionOwnerId }).(pulumi.StringPtrOutput)
}

// Tags for the subscription
func (o SubscriptionAliasResponsePropertiesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The workload type of the subscription. It can be either Production or DevTest.
func (o SubscriptionAliasResponsePropertiesResponseOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.Workload }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
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

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
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

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PutAliasRequestAdditionalPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestAdditionalPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionAliasResponsePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}