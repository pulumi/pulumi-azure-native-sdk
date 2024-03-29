// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// If set to "False", it allows the descendants of this scope to override the pricing configuration set on this scope (allows setting inherited="False"). If set to "True", it prevents overrides and forces this pricing configuration on all the descendants of this scope. This field is only available for subscription-level pricing.
type Enforce string

const (
	// Allows the descendants of this scope to override the pricing configuration set on this scope (allows setting inherited="False")
	EnforceFalse = Enforce("False")
	// Prevents overrides and forces the current scope's pricing configuration to all descendants
	EnforceTrue = Enforce("True")
)

func (Enforce) ElementType() reflect.Type {
	return reflect.TypeOf((*Enforce)(nil)).Elem()
}

func (e Enforce) ToEnforceOutput() EnforceOutput {
	return pulumi.ToOutput(e).(EnforceOutput)
}

func (e Enforce) ToEnforceOutputWithContext(ctx context.Context) EnforceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnforceOutput)
}

func (e Enforce) ToEnforcePtrOutput() EnforcePtrOutput {
	return e.ToEnforcePtrOutputWithContext(context.Background())
}

func (e Enforce) ToEnforcePtrOutputWithContext(ctx context.Context) EnforcePtrOutput {
	return Enforce(e).ToEnforceOutputWithContext(ctx).ToEnforcePtrOutputWithContext(ctx)
}

func (e Enforce) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Enforce) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Enforce) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Enforce) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnforceOutput struct{ *pulumi.OutputState }

func (EnforceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Enforce)(nil)).Elem()
}

func (o EnforceOutput) ToEnforceOutput() EnforceOutput {
	return o
}

func (o EnforceOutput) ToEnforceOutputWithContext(ctx context.Context) EnforceOutput {
	return o
}

func (o EnforceOutput) ToEnforcePtrOutput() EnforcePtrOutput {
	return o.ToEnforcePtrOutputWithContext(context.Background())
}

func (o EnforceOutput) ToEnforcePtrOutputWithContext(ctx context.Context) EnforcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Enforce) *Enforce {
		return &v
	}).(EnforcePtrOutput)
}

func (o EnforceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnforceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Enforce) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnforceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnforceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Enforce) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnforcePtrOutput struct{ *pulumi.OutputState }

func (EnforcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Enforce)(nil)).Elem()
}

func (o EnforcePtrOutput) ToEnforcePtrOutput() EnforcePtrOutput {
	return o
}

func (o EnforcePtrOutput) ToEnforcePtrOutputWithContext(ctx context.Context) EnforcePtrOutput {
	return o
}

func (o EnforcePtrOutput) Elem() EnforceOutput {
	return o.ApplyT(func(v *Enforce) Enforce {
		if v != nil {
			return *v
		}
		var ret Enforce
		return ret
	}).(EnforceOutput)
}

func (o EnforcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnforcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Enforce) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnforceInput is an input type that accepts values of the Enforce enum
// A concrete instance of `EnforceInput` can be one of the following:
//
//	EnforceFalse
//	EnforceTrue
type EnforceInput interface {
	pulumi.Input

	ToEnforceOutput() EnforceOutput
	ToEnforceOutputWithContext(context.Context) EnforceOutput
}

var enforcePtrType = reflect.TypeOf((**Enforce)(nil)).Elem()

type EnforcePtrInput interface {
	pulumi.Input

	ToEnforcePtrOutput() EnforcePtrOutput
	ToEnforcePtrOutputWithContext(context.Context) EnforcePtrOutput
}

type enforcePtr string

func EnforcePtr(v string) EnforcePtrInput {
	return (*enforcePtr)(&v)
}

func (*enforcePtr) ElementType() reflect.Type {
	return enforcePtrType
}

func (in *enforcePtr) ToEnforcePtrOutput() EnforcePtrOutput {
	return pulumi.ToOutput(in).(EnforcePtrOutput)
}

func (in *enforcePtr) ToEnforcePtrOutputWithContext(ctx context.Context) EnforcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnforcePtrOutput)
}

// Indicates whether the extension is enabled.
type IsEnabled string

const (
	// Indicates the extension is enabled
	IsEnabledTrue = IsEnabled("True")
	// Indicates the extension is disabled
	IsEnabledFalse = IsEnabled("False")
)

func (IsEnabled) ElementType() reflect.Type {
	return reflect.TypeOf((*IsEnabled)(nil)).Elem()
}

func (e IsEnabled) ToIsEnabledOutput() IsEnabledOutput {
	return pulumi.ToOutput(e).(IsEnabledOutput)
}

func (e IsEnabled) ToIsEnabledOutputWithContext(ctx context.Context) IsEnabledOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IsEnabledOutput)
}

func (e IsEnabled) ToIsEnabledPtrOutput() IsEnabledPtrOutput {
	return e.ToIsEnabledPtrOutputWithContext(context.Background())
}

func (e IsEnabled) ToIsEnabledPtrOutputWithContext(ctx context.Context) IsEnabledPtrOutput {
	return IsEnabled(e).ToIsEnabledOutputWithContext(ctx).ToIsEnabledPtrOutputWithContext(ctx)
}

func (e IsEnabled) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IsEnabled) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IsEnabled) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IsEnabled) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IsEnabledOutput struct{ *pulumi.OutputState }

func (IsEnabledOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IsEnabled)(nil)).Elem()
}

func (o IsEnabledOutput) ToIsEnabledOutput() IsEnabledOutput {
	return o
}

func (o IsEnabledOutput) ToIsEnabledOutputWithContext(ctx context.Context) IsEnabledOutput {
	return o
}

func (o IsEnabledOutput) ToIsEnabledPtrOutput() IsEnabledPtrOutput {
	return o.ToIsEnabledPtrOutputWithContext(context.Background())
}

func (o IsEnabledOutput) ToIsEnabledPtrOutputWithContext(ctx context.Context) IsEnabledPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IsEnabled) *IsEnabled {
		return &v
	}).(IsEnabledPtrOutput)
}

func (o IsEnabledOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IsEnabledOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IsEnabled) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IsEnabledOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IsEnabledOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IsEnabled) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IsEnabledPtrOutput struct{ *pulumi.OutputState }

func (IsEnabledPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IsEnabled)(nil)).Elem()
}

func (o IsEnabledPtrOutput) ToIsEnabledPtrOutput() IsEnabledPtrOutput {
	return o
}

func (o IsEnabledPtrOutput) ToIsEnabledPtrOutputWithContext(ctx context.Context) IsEnabledPtrOutput {
	return o
}

func (o IsEnabledPtrOutput) Elem() IsEnabledOutput {
	return o.ApplyT(func(v *IsEnabled) IsEnabled {
		if v != nil {
			return *v
		}
		var ret IsEnabled
		return ret
	}).(IsEnabledOutput)
}

func (o IsEnabledPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IsEnabledPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IsEnabled) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IsEnabledInput is an input type that accepts values of the IsEnabled enum
// A concrete instance of `IsEnabledInput` can be one of the following:
//
//	IsEnabledTrue
//	IsEnabledFalse
type IsEnabledInput interface {
	pulumi.Input

	ToIsEnabledOutput() IsEnabledOutput
	ToIsEnabledOutputWithContext(context.Context) IsEnabledOutput
}

var isEnabledPtrType = reflect.TypeOf((**IsEnabled)(nil)).Elem()

type IsEnabledPtrInput interface {
	pulumi.Input

	ToIsEnabledPtrOutput() IsEnabledPtrOutput
	ToIsEnabledPtrOutputWithContext(context.Context) IsEnabledPtrOutput
}

type isEnabledPtr string

func IsEnabledPtr(v string) IsEnabledPtrInput {
	return (*isEnabledPtr)(&v)
}

func (*isEnabledPtr) ElementType() reflect.Type {
	return isEnabledPtrType
}

func (in *isEnabledPtr) ToIsEnabledPtrOutput() IsEnabledPtrOutput {
	return pulumi.ToOutput(in).(IsEnabledPtrOutput)
}

func (in *isEnabledPtr) ToIsEnabledPtrOutputWithContext(ctx context.Context) IsEnabledPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IsEnabledPtrOutput)
}

// Indicates whether the Defender plan is enabled on the selected scope. Microsoft Defender for Cloud is provided in two pricing tiers: free and standard. The standard tier offers advanced security capabilities, while the free tier offers basic security features.
type PricingTier string

const (
	// Get free Microsoft Defender for Cloud experience with basic security features
	PricingTierFree = PricingTier("Free")
	// Get the standard Microsoft Defender for Cloud experience with advanced security features
	PricingTierStandard = PricingTier("Standard")
)

func (PricingTier) ElementType() reflect.Type {
	return reflect.TypeOf((*PricingTier)(nil)).Elem()
}

func (e PricingTier) ToPricingTierOutput() PricingTierOutput {
	return pulumi.ToOutput(e).(PricingTierOutput)
}

func (e PricingTier) ToPricingTierOutputWithContext(ctx context.Context) PricingTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PricingTierOutput)
}

func (e PricingTier) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return e.ToPricingTierPtrOutputWithContext(context.Background())
}

func (e PricingTier) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return PricingTier(e).ToPricingTierOutputWithContext(ctx).ToPricingTierPtrOutputWithContext(ctx)
}

func (e PricingTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PricingTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PricingTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PricingTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PricingTierOutput struct{ *pulumi.OutputState }

func (PricingTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PricingTier)(nil)).Elem()
}

func (o PricingTierOutput) ToPricingTierOutput() PricingTierOutput {
	return o
}

func (o PricingTierOutput) ToPricingTierOutputWithContext(ctx context.Context) PricingTierOutput {
	return o
}

func (o PricingTierOutput) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return o.ToPricingTierPtrOutputWithContext(context.Background())
}

func (o PricingTierOutput) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PricingTier) *PricingTier {
		return &v
	}).(PricingTierPtrOutput)
}

func (o PricingTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PricingTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PricingTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PricingTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PricingTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PricingTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PricingTierPtrOutput struct{ *pulumi.OutputState }

func (PricingTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PricingTier)(nil)).Elem()
}

func (o PricingTierPtrOutput) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return o
}

func (o PricingTierPtrOutput) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return o
}

func (o PricingTierPtrOutput) Elem() PricingTierOutput {
	return o.ApplyT(func(v *PricingTier) PricingTier {
		if v != nil {
			return *v
		}
		var ret PricingTier
		return ret
	}).(PricingTierOutput)
}

func (o PricingTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PricingTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PricingTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PricingTierInput is an input type that accepts values of the PricingTier enum
// A concrete instance of `PricingTierInput` can be one of the following:
//
//	PricingTierFree
//	PricingTierStandard
type PricingTierInput interface {
	pulumi.Input

	ToPricingTierOutput() PricingTierOutput
	ToPricingTierOutputWithContext(context.Context) PricingTierOutput
}

var pricingTierPtrType = reflect.TypeOf((**PricingTier)(nil)).Elem()

type PricingTierPtrInput interface {
	pulumi.Input

	ToPricingTierPtrOutput() PricingTierPtrOutput
	ToPricingTierPtrOutputWithContext(context.Context) PricingTierPtrOutput
}

type pricingTierPtr string

func PricingTierPtr(v string) PricingTierPtrInput {
	return (*pricingTierPtr)(&v)
}

func (*pricingTierPtr) ElementType() reflect.Type {
	return pricingTierPtrType
}

func (in *pricingTierPtr) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return pulumi.ToOutput(in).(PricingTierPtrOutput)
}

func (in *pricingTierPtr) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PricingTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnforceOutput{})
	pulumi.RegisterOutputType(EnforcePtrOutput{})
	pulumi.RegisterOutputType(IsEnabledOutput{})
	pulumi.RegisterOutputType(IsEnabledPtrOutput{})
	pulumi.RegisterOutputType(PricingTierOutput{})
	pulumi.RegisterOutputType(PricingTierPtrOutput{})
}
