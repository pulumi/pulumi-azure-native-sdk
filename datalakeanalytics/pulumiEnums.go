// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datalakeanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of AAD object the object identifier refers to.
type AADObjectType string

const (
	AADObjectTypeUser             = AADObjectType("User")
	AADObjectTypeGroup            = AADObjectType("Group")
	AADObjectTypeServicePrincipal = AADObjectType("ServicePrincipal")
)

// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
type FirewallAllowAzureIpsState string

const (
	FirewallAllowAzureIpsStateEnabled  = FirewallAllowAzureIpsState("Enabled")
	FirewallAllowAzureIpsStateDisabled = FirewallAllowAzureIpsState("Disabled")
)

func (FirewallAllowAzureIpsState) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallAllowAzureIpsState)(nil)).Elem()
}

func (e FirewallAllowAzureIpsState) ToFirewallAllowAzureIpsStateOutput() FirewallAllowAzureIpsStateOutput {
	return pulumi.ToOutput(e).(FirewallAllowAzureIpsStateOutput)
}

func (e FirewallAllowAzureIpsState) ToFirewallAllowAzureIpsStateOutputWithContext(ctx context.Context) FirewallAllowAzureIpsStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallAllowAzureIpsStateOutput)
}

func (e FirewallAllowAzureIpsState) ToFirewallAllowAzureIpsStatePtrOutput() FirewallAllowAzureIpsStatePtrOutput {
	return e.ToFirewallAllowAzureIpsStatePtrOutputWithContext(context.Background())
}

func (e FirewallAllowAzureIpsState) ToFirewallAllowAzureIpsStatePtrOutputWithContext(ctx context.Context) FirewallAllowAzureIpsStatePtrOutput {
	return FirewallAllowAzureIpsState(e).ToFirewallAllowAzureIpsStateOutputWithContext(ctx).ToFirewallAllowAzureIpsStatePtrOutputWithContext(ctx)
}

func (e FirewallAllowAzureIpsState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallAllowAzureIpsState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallAllowAzureIpsState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallAllowAzureIpsState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallAllowAzureIpsStateOutput struct{ *pulumi.OutputState }

func (FirewallAllowAzureIpsStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallAllowAzureIpsState)(nil)).Elem()
}

func (o FirewallAllowAzureIpsStateOutput) ToFirewallAllowAzureIpsStateOutput() FirewallAllowAzureIpsStateOutput {
	return o
}

func (o FirewallAllowAzureIpsStateOutput) ToFirewallAllowAzureIpsStateOutputWithContext(ctx context.Context) FirewallAllowAzureIpsStateOutput {
	return o
}

func (o FirewallAllowAzureIpsStateOutput) ToFirewallAllowAzureIpsStatePtrOutput() FirewallAllowAzureIpsStatePtrOutput {
	return o.ToFirewallAllowAzureIpsStatePtrOutputWithContext(context.Background())
}

func (o FirewallAllowAzureIpsStateOutput) ToFirewallAllowAzureIpsStatePtrOutputWithContext(ctx context.Context) FirewallAllowAzureIpsStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallAllowAzureIpsState) *FirewallAllowAzureIpsState {
		return &v
	}).(FirewallAllowAzureIpsStatePtrOutput)
}

func (o FirewallAllowAzureIpsStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallAllowAzureIpsStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallAllowAzureIpsState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallAllowAzureIpsStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallAllowAzureIpsStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallAllowAzureIpsState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallAllowAzureIpsStatePtrOutput struct{ *pulumi.OutputState }

func (FirewallAllowAzureIpsStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAllowAzureIpsState)(nil)).Elem()
}

func (o FirewallAllowAzureIpsStatePtrOutput) ToFirewallAllowAzureIpsStatePtrOutput() FirewallAllowAzureIpsStatePtrOutput {
	return o
}

func (o FirewallAllowAzureIpsStatePtrOutput) ToFirewallAllowAzureIpsStatePtrOutputWithContext(ctx context.Context) FirewallAllowAzureIpsStatePtrOutput {
	return o
}

func (o FirewallAllowAzureIpsStatePtrOutput) Elem() FirewallAllowAzureIpsStateOutput {
	return o.ApplyT(func(v *FirewallAllowAzureIpsState) FirewallAllowAzureIpsState {
		if v != nil {
			return *v
		}
		var ret FirewallAllowAzureIpsState
		return ret
	}).(FirewallAllowAzureIpsStateOutput)
}

func (o FirewallAllowAzureIpsStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallAllowAzureIpsStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallAllowAzureIpsState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FirewallAllowAzureIpsStateInput is an input type that accepts FirewallAllowAzureIpsStateArgs and FirewallAllowAzureIpsStateOutput values.
// You can construct a concrete instance of `FirewallAllowAzureIpsStateInput` via:
//
//	FirewallAllowAzureIpsStateArgs{...}
type FirewallAllowAzureIpsStateInput interface {
	pulumi.Input

	ToFirewallAllowAzureIpsStateOutput() FirewallAllowAzureIpsStateOutput
	ToFirewallAllowAzureIpsStateOutputWithContext(context.Context) FirewallAllowAzureIpsStateOutput
}

var firewallAllowAzureIpsStatePtrType = reflect.TypeOf((**FirewallAllowAzureIpsState)(nil)).Elem()

type FirewallAllowAzureIpsStatePtrInput interface {
	pulumi.Input

	ToFirewallAllowAzureIpsStatePtrOutput() FirewallAllowAzureIpsStatePtrOutput
	ToFirewallAllowAzureIpsStatePtrOutputWithContext(context.Context) FirewallAllowAzureIpsStatePtrOutput
}

type firewallAllowAzureIpsStatePtr string

func FirewallAllowAzureIpsStatePtr(v string) FirewallAllowAzureIpsStatePtrInput {
	return (*firewallAllowAzureIpsStatePtr)(&v)
}

func (*firewallAllowAzureIpsStatePtr) ElementType() reflect.Type {
	return firewallAllowAzureIpsStatePtrType
}

func (in *firewallAllowAzureIpsStatePtr) ToFirewallAllowAzureIpsStatePtrOutput() FirewallAllowAzureIpsStatePtrOutput {
	return pulumi.ToOutput(in).(FirewallAllowAzureIpsStatePtrOutput)
}

func (in *firewallAllowAzureIpsStatePtr) ToFirewallAllowAzureIpsStatePtrOutputWithContext(ctx context.Context) FirewallAllowAzureIpsStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallAllowAzureIpsStatePtrOutput)
}

// The current state of the IP address firewall for this account.
type FirewallState string

const (
	FirewallStateEnabled  = FirewallState("Enabled")
	FirewallStateDisabled = FirewallState("Disabled")
)

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallState)(nil)).Elem()
}

func (e FirewallState) ToFirewallStateOutput() FirewallStateOutput {
	return pulumi.ToOutput(e).(FirewallStateOutput)
}

func (e FirewallState) ToFirewallStateOutputWithContext(ctx context.Context) FirewallStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallStateOutput)
}

func (e FirewallState) ToFirewallStatePtrOutput() FirewallStatePtrOutput {
	return e.ToFirewallStatePtrOutputWithContext(context.Background())
}

func (e FirewallState) ToFirewallStatePtrOutputWithContext(ctx context.Context) FirewallStatePtrOutput {
	return FirewallState(e).ToFirewallStateOutputWithContext(ctx).ToFirewallStatePtrOutputWithContext(ctx)
}

func (e FirewallState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallStateOutput struct{ *pulumi.OutputState }

func (FirewallStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallState)(nil)).Elem()
}

func (o FirewallStateOutput) ToFirewallStateOutput() FirewallStateOutput {
	return o
}

func (o FirewallStateOutput) ToFirewallStateOutputWithContext(ctx context.Context) FirewallStateOutput {
	return o
}

func (o FirewallStateOutput) ToFirewallStatePtrOutput() FirewallStatePtrOutput {
	return o.ToFirewallStatePtrOutputWithContext(context.Background())
}

func (o FirewallStateOutput) ToFirewallStatePtrOutputWithContext(ctx context.Context) FirewallStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallState) *FirewallState {
		return &v
	}).(FirewallStatePtrOutput)
}

func (o FirewallStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallStatePtrOutput struct{ *pulumi.OutputState }

func (FirewallStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallState)(nil)).Elem()
}

func (o FirewallStatePtrOutput) ToFirewallStatePtrOutput() FirewallStatePtrOutput {
	return o
}

func (o FirewallStatePtrOutput) ToFirewallStatePtrOutputWithContext(ctx context.Context) FirewallStatePtrOutput {
	return o
}

func (o FirewallStatePtrOutput) Elem() FirewallStateOutput {
	return o.ApplyT(func(v *FirewallState) FirewallState {
		if v != nil {
			return *v
		}
		var ret FirewallState
		return ret
	}).(FirewallStateOutput)
}

func (o FirewallStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FirewallStateInput is an input type that accepts FirewallStateArgs and FirewallStateOutput values.
// You can construct a concrete instance of `FirewallStateInput` via:
//
//	FirewallStateArgs{...}
type FirewallStateInput interface {
	pulumi.Input

	ToFirewallStateOutput() FirewallStateOutput
	ToFirewallStateOutputWithContext(context.Context) FirewallStateOutput
}

var firewallStatePtrType = reflect.TypeOf((**FirewallState)(nil)).Elem()

type FirewallStatePtrInput interface {
	pulumi.Input

	ToFirewallStatePtrOutput() FirewallStatePtrOutput
	ToFirewallStatePtrOutputWithContext(context.Context) FirewallStatePtrOutput
}

type firewallStatePtr string

func FirewallStatePtr(v string) FirewallStatePtrInput {
	return (*firewallStatePtr)(&v)
}

func (*firewallStatePtr) ElementType() reflect.Type {
	return firewallStatePtrType
}

func (in *firewallStatePtr) ToFirewallStatePtrOutput() FirewallStatePtrOutput {
	return pulumi.ToOutput(in).(FirewallStatePtrOutput)
}

func (in *firewallStatePtr) ToFirewallStatePtrOutputWithContext(ctx context.Context) FirewallStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallStatePtrOutput)
}

// The commitment tier for the next month.
type TierType string

const (
	TierTypeConsumption               = TierType("Consumption")
	TierType_Commitment_100AUHours    = TierType("Commitment_100AUHours")
	TierType_Commitment_500AUHours    = TierType("Commitment_500AUHours")
	TierType_Commitment_1000AUHours   = TierType("Commitment_1000AUHours")
	TierType_Commitment_5000AUHours   = TierType("Commitment_5000AUHours")
	TierType_Commitment_10000AUHours  = TierType("Commitment_10000AUHours")
	TierType_Commitment_50000AUHours  = TierType("Commitment_50000AUHours")
	TierType_Commitment_100000AUHours = TierType("Commitment_100000AUHours")
	TierType_Commitment_500000AUHours = TierType("Commitment_500000AUHours")
)

func (TierType) ElementType() reflect.Type {
	return reflect.TypeOf((*TierType)(nil)).Elem()
}

func (e TierType) ToTierTypeOutput() TierTypeOutput {
	return pulumi.ToOutput(e).(TierTypeOutput)
}

func (e TierType) ToTierTypeOutputWithContext(ctx context.Context) TierTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TierTypeOutput)
}

func (e TierType) ToTierTypePtrOutput() TierTypePtrOutput {
	return e.ToTierTypePtrOutputWithContext(context.Background())
}

func (e TierType) ToTierTypePtrOutputWithContext(ctx context.Context) TierTypePtrOutput {
	return TierType(e).ToTierTypeOutputWithContext(ctx).ToTierTypePtrOutputWithContext(ctx)
}

func (e TierType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TierType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TierType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TierType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TierTypeOutput struct{ *pulumi.OutputState }

func (TierTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TierType)(nil)).Elem()
}

func (o TierTypeOutput) ToTierTypeOutput() TierTypeOutput {
	return o
}

func (o TierTypeOutput) ToTierTypeOutputWithContext(ctx context.Context) TierTypeOutput {
	return o
}

func (o TierTypeOutput) ToTierTypePtrOutput() TierTypePtrOutput {
	return o.ToTierTypePtrOutputWithContext(context.Background())
}

func (o TierTypeOutput) ToTierTypePtrOutputWithContext(ctx context.Context) TierTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TierType) *TierType {
		return &v
	}).(TierTypePtrOutput)
}

func (o TierTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TierTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TierType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TierTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TierTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TierType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TierTypePtrOutput struct{ *pulumi.OutputState }

func (TierTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TierType)(nil)).Elem()
}

func (o TierTypePtrOutput) ToTierTypePtrOutput() TierTypePtrOutput {
	return o
}

func (o TierTypePtrOutput) ToTierTypePtrOutputWithContext(ctx context.Context) TierTypePtrOutput {
	return o
}

func (o TierTypePtrOutput) Elem() TierTypeOutput {
	return o.ApplyT(func(v *TierType) TierType {
		if v != nil {
			return *v
		}
		var ret TierType
		return ret
	}).(TierTypeOutput)
}

func (o TierTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TierTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TierType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TierTypeInput is an input type that accepts TierTypeArgs and TierTypeOutput values.
// You can construct a concrete instance of `TierTypeInput` via:
//
//	TierTypeArgs{...}
type TierTypeInput interface {
	pulumi.Input

	ToTierTypeOutput() TierTypeOutput
	ToTierTypeOutputWithContext(context.Context) TierTypeOutput
}

var tierTypePtrType = reflect.TypeOf((**TierType)(nil)).Elem()

type TierTypePtrInput interface {
	pulumi.Input

	ToTierTypePtrOutput() TierTypePtrOutput
	ToTierTypePtrOutputWithContext(context.Context) TierTypePtrOutput
}

type tierTypePtr string

func TierTypePtr(v string) TierTypePtrInput {
	return (*tierTypePtr)(&v)
}

func (*tierTypePtr) ElementType() reflect.Type {
	return tierTypePtrType
}

func (in *tierTypePtr) ToTierTypePtrOutput() TierTypePtrOutput {
	return pulumi.ToOutput(in).(TierTypePtrOutput)
}

func (in *tierTypePtr) ToTierTypePtrOutputWithContext(ctx context.Context) TierTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TierTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallAllowAzureIpsStateOutput{})
	pulumi.RegisterOutputType(FirewallAllowAzureIpsStatePtrOutput{})
	pulumi.RegisterOutputType(FirewallStateOutput{})
	pulumi.RegisterOutputType(FirewallStatePtrOutput{})
	pulumi.RegisterOutputType(TierTypeOutput{})
	pulumi.RegisterOutputType(TierTypePtrOutput{})
}