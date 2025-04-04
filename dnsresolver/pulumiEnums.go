// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnsresolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of action to take.
type ActionType string

const (
	ActionTypeAllow = ActionType("Allow")
	ActionTypeAlert = ActionType("Alert")
	ActionTypeBlock = ActionType("Block")
)

func (ActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (e ActionType) ToActionTypeOutput() ActionTypeOutput {
	return pulumi.ToOutput(e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypePtrOutput() ActionTypePtrOutput {
	return e.ToActionTypePtrOutputWithContext(context.Background())
}

func (e ActionType) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return ActionType(e).ToActionTypeOutputWithContext(ctx).ToActionTypePtrOutputWithContext(ctx)
}

func (e ActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionTypeOutput struct{ *pulumi.OutputState }

func (ActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (o ActionTypeOutput) ToActionTypeOutput() ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o.ToActionTypePtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionType) *ActionType {
		return &v
	}).(ActionTypePtrOutput)
}

func (o ActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionTypePtrOutput struct{ *pulumi.OutputState }

func (ActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionType)(nil)).Elem()
}

func (o ActionTypePtrOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) Elem() ActionTypeOutput {
	return o.ApplyT(func(v *ActionType) ActionType {
		if v != nil {
			return *v
		}
		var ret ActionType
		return ret
	}).(ActionTypeOutput)
}

func (o ActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ActionTypeInput is an input type that accepts values of the ActionType enum
// A concrete instance of `ActionTypeInput` can be one of the following:
//
//	ActionTypeAllow
//	ActionTypeAlert
//	ActionTypeBlock
type ActionTypeInput interface {
	pulumi.Input

	ToActionTypeOutput() ActionTypeOutput
	ToActionTypeOutputWithContext(context.Context) ActionTypeOutput
}

var actionTypePtrType = reflect.TypeOf((**ActionType)(nil)).Elem()

type ActionTypePtrInput interface {
	pulumi.Input

	ToActionTypePtrOutput() ActionTypePtrOutput
	ToActionTypePtrOutputWithContext(context.Context) ActionTypePtrOutput
}

type actionTypePtr string

func ActionTypePtr(v string) ActionTypePtrInput {
	return (*actionTypePtr)(&v)
}

func (*actionTypePtr) ElementType() reflect.Type {
	return actionTypePtrType
}

func (in *actionTypePtr) ToActionTypePtrOutput() ActionTypePtrOutput {
	return pulumi.ToOutput(in).(ActionTypePtrOutput)
}

func (in *actionTypePtr) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionTypePtrOutput)
}

// The response code for block actions.
type BlockResponseCode string

const (
	BlockResponseCodeSERVFAIL = BlockResponseCode("SERVFAIL")
)

func (BlockResponseCode) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockResponseCode)(nil)).Elem()
}

func (e BlockResponseCode) ToBlockResponseCodeOutput() BlockResponseCodeOutput {
	return pulumi.ToOutput(e).(BlockResponseCodeOutput)
}

func (e BlockResponseCode) ToBlockResponseCodeOutputWithContext(ctx context.Context) BlockResponseCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlockResponseCodeOutput)
}

func (e BlockResponseCode) ToBlockResponseCodePtrOutput() BlockResponseCodePtrOutput {
	return e.ToBlockResponseCodePtrOutputWithContext(context.Background())
}

func (e BlockResponseCode) ToBlockResponseCodePtrOutputWithContext(ctx context.Context) BlockResponseCodePtrOutput {
	return BlockResponseCode(e).ToBlockResponseCodeOutputWithContext(ctx).ToBlockResponseCodePtrOutputWithContext(ctx)
}

func (e BlockResponseCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlockResponseCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlockResponseCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlockResponseCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlockResponseCodeOutput struct{ *pulumi.OutputState }

func (BlockResponseCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockResponseCode)(nil)).Elem()
}

func (o BlockResponseCodeOutput) ToBlockResponseCodeOutput() BlockResponseCodeOutput {
	return o
}

func (o BlockResponseCodeOutput) ToBlockResponseCodeOutputWithContext(ctx context.Context) BlockResponseCodeOutput {
	return o
}

func (o BlockResponseCodeOutput) ToBlockResponseCodePtrOutput() BlockResponseCodePtrOutput {
	return o.ToBlockResponseCodePtrOutputWithContext(context.Background())
}

func (o BlockResponseCodeOutput) ToBlockResponseCodePtrOutputWithContext(ctx context.Context) BlockResponseCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlockResponseCode) *BlockResponseCode {
		return &v
	}).(BlockResponseCodePtrOutput)
}

func (o BlockResponseCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlockResponseCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlockResponseCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlockResponseCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlockResponseCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlockResponseCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlockResponseCodePtrOutput struct{ *pulumi.OutputState }

func (BlockResponseCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockResponseCode)(nil)).Elem()
}

func (o BlockResponseCodePtrOutput) ToBlockResponseCodePtrOutput() BlockResponseCodePtrOutput {
	return o
}

func (o BlockResponseCodePtrOutput) ToBlockResponseCodePtrOutputWithContext(ctx context.Context) BlockResponseCodePtrOutput {
	return o
}

func (o BlockResponseCodePtrOutput) Elem() BlockResponseCodeOutput {
	return o.ApplyT(func(v *BlockResponseCode) BlockResponseCode {
		if v != nil {
			return *v
		}
		var ret BlockResponseCode
		return ret
	}).(BlockResponseCodeOutput)
}

func (o BlockResponseCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlockResponseCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlockResponseCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BlockResponseCodeInput is an input type that accepts values of the BlockResponseCode enum
// A concrete instance of `BlockResponseCodeInput` can be one of the following:
//
//	BlockResponseCodeSERVFAIL
type BlockResponseCodeInput interface {
	pulumi.Input

	ToBlockResponseCodeOutput() BlockResponseCodeOutput
	ToBlockResponseCodeOutputWithContext(context.Context) BlockResponseCodeOutput
}

var blockResponseCodePtrType = reflect.TypeOf((**BlockResponseCode)(nil)).Elem()

type BlockResponseCodePtrInput interface {
	pulumi.Input

	ToBlockResponseCodePtrOutput() BlockResponseCodePtrOutput
	ToBlockResponseCodePtrOutputWithContext(context.Context) BlockResponseCodePtrOutput
}

type blockResponseCodePtr string

func BlockResponseCodePtr(v string) BlockResponseCodePtrInput {
	return (*blockResponseCodePtr)(&v)
}

func (*blockResponseCodePtr) ElementType() reflect.Type {
	return blockResponseCodePtrType
}

func (in *blockResponseCodePtr) ToBlockResponseCodePtrOutput() BlockResponseCodePtrOutput {
	return pulumi.ToOutput(in).(BlockResponseCodePtrOutput)
}

func (in *blockResponseCodePtr) ToBlockResponseCodePtrOutputWithContext(ctx context.Context) BlockResponseCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlockResponseCodePtrOutput)
}

// The state of DNS security rule.
type DnsSecurityRuleStateEnum string

const (
	DnsSecurityRuleStateEnumEnabled  = DnsSecurityRuleStateEnum("Enabled")
	DnsSecurityRuleStateEnumDisabled = DnsSecurityRuleStateEnum("Disabled")
)

func (DnsSecurityRuleStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsSecurityRuleStateEnum)(nil)).Elem()
}

func (e DnsSecurityRuleStateEnum) ToDnsSecurityRuleStateEnumOutput() DnsSecurityRuleStateEnumOutput {
	return pulumi.ToOutput(e).(DnsSecurityRuleStateEnumOutput)
}

func (e DnsSecurityRuleStateEnum) ToDnsSecurityRuleStateEnumOutputWithContext(ctx context.Context) DnsSecurityRuleStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DnsSecurityRuleStateEnumOutput)
}

func (e DnsSecurityRuleStateEnum) ToDnsSecurityRuleStateEnumPtrOutput() DnsSecurityRuleStateEnumPtrOutput {
	return e.ToDnsSecurityRuleStateEnumPtrOutputWithContext(context.Background())
}

func (e DnsSecurityRuleStateEnum) ToDnsSecurityRuleStateEnumPtrOutputWithContext(ctx context.Context) DnsSecurityRuleStateEnumPtrOutput {
	return DnsSecurityRuleStateEnum(e).ToDnsSecurityRuleStateEnumOutputWithContext(ctx).ToDnsSecurityRuleStateEnumPtrOutputWithContext(ctx)
}

func (e DnsSecurityRuleStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DnsSecurityRuleStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DnsSecurityRuleStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DnsSecurityRuleStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DnsSecurityRuleStateEnumOutput struct{ *pulumi.OutputState }

func (DnsSecurityRuleStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsSecurityRuleStateEnum)(nil)).Elem()
}

func (o DnsSecurityRuleStateEnumOutput) ToDnsSecurityRuleStateEnumOutput() DnsSecurityRuleStateEnumOutput {
	return o
}

func (o DnsSecurityRuleStateEnumOutput) ToDnsSecurityRuleStateEnumOutputWithContext(ctx context.Context) DnsSecurityRuleStateEnumOutput {
	return o
}

func (o DnsSecurityRuleStateEnumOutput) ToDnsSecurityRuleStateEnumPtrOutput() DnsSecurityRuleStateEnumPtrOutput {
	return o.ToDnsSecurityRuleStateEnumPtrOutputWithContext(context.Background())
}

func (o DnsSecurityRuleStateEnumOutput) ToDnsSecurityRuleStateEnumPtrOutputWithContext(ctx context.Context) DnsSecurityRuleStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsSecurityRuleStateEnum) *DnsSecurityRuleStateEnum {
		return &v
	}).(DnsSecurityRuleStateEnumPtrOutput)
}

func (o DnsSecurityRuleStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DnsSecurityRuleStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DnsSecurityRuleStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DnsSecurityRuleStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DnsSecurityRuleStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DnsSecurityRuleStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DnsSecurityRuleStateEnumPtrOutput struct{ *pulumi.OutputState }

func (DnsSecurityRuleStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSecurityRuleStateEnum)(nil)).Elem()
}

func (o DnsSecurityRuleStateEnumPtrOutput) ToDnsSecurityRuleStateEnumPtrOutput() DnsSecurityRuleStateEnumPtrOutput {
	return o
}

func (o DnsSecurityRuleStateEnumPtrOutput) ToDnsSecurityRuleStateEnumPtrOutputWithContext(ctx context.Context) DnsSecurityRuleStateEnumPtrOutput {
	return o
}

func (o DnsSecurityRuleStateEnumPtrOutput) Elem() DnsSecurityRuleStateEnumOutput {
	return o.ApplyT(func(v *DnsSecurityRuleStateEnum) DnsSecurityRuleStateEnum {
		if v != nil {
			return *v
		}
		var ret DnsSecurityRuleStateEnum
		return ret
	}).(DnsSecurityRuleStateEnumOutput)
}

func (o DnsSecurityRuleStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DnsSecurityRuleStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DnsSecurityRuleStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DnsSecurityRuleStateEnumInput is an input type that accepts values of the DnsSecurityRuleStateEnum enum
// A concrete instance of `DnsSecurityRuleStateEnumInput` can be one of the following:
//
//	DnsSecurityRuleStateEnumEnabled
//	DnsSecurityRuleStateEnumDisabled
type DnsSecurityRuleStateEnumInput interface {
	pulumi.Input

	ToDnsSecurityRuleStateEnumOutput() DnsSecurityRuleStateEnumOutput
	ToDnsSecurityRuleStateEnumOutputWithContext(context.Context) DnsSecurityRuleStateEnumOutput
}

var dnsSecurityRuleStateEnumPtrType = reflect.TypeOf((**DnsSecurityRuleStateEnum)(nil)).Elem()

type DnsSecurityRuleStateEnumPtrInput interface {
	pulumi.Input

	ToDnsSecurityRuleStateEnumPtrOutput() DnsSecurityRuleStateEnumPtrOutput
	ToDnsSecurityRuleStateEnumPtrOutputWithContext(context.Context) DnsSecurityRuleStateEnumPtrOutput
}

type dnsSecurityRuleStateEnumPtr string

func DnsSecurityRuleStateEnumPtr(v string) DnsSecurityRuleStateEnumPtrInput {
	return (*dnsSecurityRuleStateEnumPtr)(&v)
}

func (*dnsSecurityRuleStateEnumPtr) ElementType() reflect.Type {
	return dnsSecurityRuleStateEnumPtrType
}

func (in *dnsSecurityRuleStateEnumPtr) ToDnsSecurityRuleStateEnumPtrOutput() DnsSecurityRuleStateEnumPtrOutput {
	return pulumi.ToOutput(in).(DnsSecurityRuleStateEnumPtrOutput)
}

func (in *dnsSecurityRuleStateEnumPtr) ToDnsSecurityRuleStateEnumPtrOutputWithContext(ctx context.Context) DnsSecurityRuleStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DnsSecurityRuleStateEnumPtrOutput)
}

// The state of forwarding rule.
type ForwardingRuleStateEnum string

const (
	ForwardingRuleStateEnumEnabled  = ForwardingRuleStateEnum("Enabled")
	ForwardingRuleStateEnumDisabled = ForwardingRuleStateEnum("Disabled")
)

func (ForwardingRuleStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardingRuleStateEnum)(nil)).Elem()
}

func (e ForwardingRuleStateEnum) ToForwardingRuleStateEnumOutput() ForwardingRuleStateEnumOutput {
	return pulumi.ToOutput(e).(ForwardingRuleStateEnumOutput)
}

func (e ForwardingRuleStateEnum) ToForwardingRuleStateEnumOutputWithContext(ctx context.Context) ForwardingRuleStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ForwardingRuleStateEnumOutput)
}

func (e ForwardingRuleStateEnum) ToForwardingRuleStateEnumPtrOutput() ForwardingRuleStateEnumPtrOutput {
	return e.ToForwardingRuleStateEnumPtrOutputWithContext(context.Background())
}

func (e ForwardingRuleStateEnum) ToForwardingRuleStateEnumPtrOutputWithContext(ctx context.Context) ForwardingRuleStateEnumPtrOutput {
	return ForwardingRuleStateEnum(e).ToForwardingRuleStateEnumOutputWithContext(ctx).ToForwardingRuleStateEnumPtrOutputWithContext(ctx)
}

func (e ForwardingRuleStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ForwardingRuleStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ForwardingRuleStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ForwardingRuleStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ForwardingRuleStateEnumOutput struct{ *pulumi.OutputState }

func (ForwardingRuleStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardingRuleStateEnum)(nil)).Elem()
}

func (o ForwardingRuleStateEnumOutput) ToForwardingRuleStateEnumOutput() ForwardingRuleStateEnumOutput {
	return o
}

func (o ForwardingRuleStateEnumOutput) ToForwardingRuleStateEnumOutputWithContext(ctx context.Context) ForwardingRuleStateEnumOutput {
	return o
}

func (o ForwardingRuleStateEnumOutput) ToForwardingRuleStateEnumPtrOutput() ForwardingRuleStateEnumPtrOutput {
	return o.ToForwardingRuleStateEnumPtrOutputWithContext(context.Background())
}

func (o ForwardingRuleStateEnumOutput) ToForwardingRuleStateEnumPtrOutputWithContext(ctx context.Context) ForwardingRuleStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ForwardingRuleStateEnum) *ForwardingRuleStateEnum {
		return &v
	}).(ForwardingRuleStateEnumPtrOutput)
}

func (o ForwardingRuleStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ForwardingRuleStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ForwardingRuleStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ForwardingRuleStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ForwardingRuleStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ForwardingRuleStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ForwardingRuleStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ForwardingRuleStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardingRuleStateEnum)(nil)).Elem()
}

func (o ForwardingRuleStateEnumPtrOutput) ToForwardingRuleStateEnumPtrOutput() ForwardingRuleStateEnumPtrOutput {
	return o
}

func (o ForwardingRuleStateEnumPtrOutput) ToForwardingRuleStateEnumPtrOutputWithContext(ctx context.Context) ForwardingRuleStateEnumPtrOutput {
	return o
}

func (o ForwardingRuleStateEnumPtrOutput) Elem() ForwardingRuleStateEnumOutput {
	return o.ApplyT(func(v *ForwardingRuleStateEnum) ForwardingRuleStateEnum {
		if v != nil {
			return *v
		}
		var ret ForwardingRuleStateEnum
		return ret
	}).(ForwardingRuleStateEnumOutput)
}

func (o ForwardingRuleStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ForwardingRuleStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ForwardingRuleStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ForwardingRuleStateEnumInput is an input type that accepts values of the ForwardingRuleStateEnum enum
// A concrete instance of `ForwardingRuleStateEnumInput` can be one of the following:
//
//	ForwardingRuleStateEnumEnabled
//	ForwardingRuleStateEnumDisabled
type ForwardingRuleStateEnumInput interface {
	pulumi.Input

	ToForwardingRuleStateEnumOutput() ForwardingRuleStateEnumOutput
	ToForwardingRuleStateEnumOutputWithContext(context.Context) ForwardingRuleStateEnumOutput
}

var forwardingRuleStateEnumPtrType = reflect.TypeOf((**ForwardingRuleStateEnum)(nil)).Elem()

type ForwardingRuleStateEnumPtrInput interface {
	pulumi.Input

	ToForwardingRuleStateEnumPtrOutput() ForwardingRuleStateEnumPtrOutput
	ToForwardingRuleStateEnumPtrOutputWithContext(context.Context) ForwardingRuleStateEnumPtrOutput
}

type forwardingRuleStateEnumPtr string

func ForwardingRuleStateEnumPtr(v string) ForwardingRuleStateEnumPtrInput {
	return (*forwardingRuleStateEnumPtr)(&v)
}

func (*forwardingRuleStateEnumPtr) ElementType() reflect.Type {
	return forwardingRuleStateEnumPtrType
}

func (in *forwardingRuleStateEnumPtr) ToForwardingRuleStateEnumPtrOutput() ForwardingRuleStateEnumPtrOutput {
	return pulumi.ToOutput(in).(ForwardingRuleStateEnumPtrOutput)
}

func (in *forwardingRuleStateEnumPtr) ToForwardingRuleStateEnumPtrOutputWithContext(ctx context.Context) ForwardingRuleStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ForwardingRuleStateEnumPtrOutput)
}

// Private IP address allocation method.
type IpAllocationMethod string

const (
	IpAllocationMethodStatic  = IpAllocationMethod("Static")
	IpAllocationMethodDynamic = IpAllocationMethod("Dynamic")
)

func (IpAllocationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocationMethod)(nil)).Elem()
}

func (e IpAllocationMethod) ToIpAllocationMethodOutput() IpAllocationMethodOutput {
	return pulumi.ToOutput(e).(IpAllocationMethodOutput)
}

func (e IpAllocationMethod) ToIpAllocationMethodOutputWithContext(ctx context.Context) IpAllocationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpAllocationMethodOutput)
}

func (e IpAllocationMethod) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return e.ToIpAllocationMethodPtrOutputWithContext(context.Background())
}

func (e IpAllocationMethod) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return IpAllocationMethod(e).ToIpAllocationMethodOutputWithContext(ctx).ToIpAllocationMethodPtrOutputWithContext(ctx)
}

func (e IpAllocationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAllocationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAllocationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpAllocationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpAllocationMethodOutput struct{ *pulumi.OutputState }

func (IpAllocationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocationMethod)(nil)).Elem()
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodOutput() IpAllocationMethodOutput {
	return o
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodOutputWithContext(ctx context.Context) IpAllocationMethodOutput {
	return o
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return o.ToIpAllocationMethodPtrOutputWithContext(context.Background())
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAllocationMethod) *IpAllocationMethod {
		return &v
	}).(IpAllocationMethodPtrOutput)
}

func (o IpAllocationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpAllocationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAllocationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpAllocationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAllocationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAllocationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpAllocationMethodPtrOutput struct{ *pulumi.OutputState }

func (IpAllocationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocationMethod)(nil)).Elem()
}

func (o IpAllocationMethodPtrOutput) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return o
}

func (o IpAllocationMethodPtrOutput) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return o
}

func (o IpAllocationMethodPtrOutput) Elem() IpAllocationMethodOutput {
	return o.ApplyT(func(v *IpAllocationMethod) IpAllocationMethod {
		if v != nil {
			return *v
		}
		var ret IpAllocationMethod
		return ret
	}).(IpAllocationMethodOutput)
}

func (o IpAllocationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAllocationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpAllocationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IpAllocationMethodInput is an input type that accepts values of the IpAllocationMethod enum
// A concrete instance of `IpAllocationMethodInput` can be one of the following:
//
//	IpAllocationMethodStatic
//	IpAllocationMethodDynamic
type IpAllocationMethodInput interface {
	pulumi.Input

	ToIpAllocationMethodOutput() IpAllocationMethodOutput
	ToIpAllocationMethodOutputWithContext(context.Context) IpAllocationMethodOutput
}

var ipAllocationMethodPtrType = reflect.TypeOf((**IpAllocationMethod)(nil)).Elem()

type IpAllocationMethodPtrInput interface {
	pulumi.Input

	ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput
	ToIpAllocationMethodPtrOutputWithContext(context.Context) IpAllocationMethodPtrOutput
}

type ipAllocationMethodPtr string

func IpAllocationMethodPtr(v string) IpAllocationMethodPtrInput {
	return (*ipAllocationMethodPtr)(&v)
}

func (*ipAllocationMethodPtr) ElementType() reflect.Type {
	return ipAllocationMethodPtrType
}

func (in *ipAllocationMethodPtr) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return pulumi.ToOutput(in).(IpAllocationMethodPtrOutput)
}

func (in *ipAllocationMethodPtr) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpAllocationMethodPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionTypeOutput{})
	pulumi.RegisterOutputType(ActionTypePtrOutput{})
	pulumi.RegisterOutputType(BlockResponseCodeOutput{})
	pulumi.RegisterOutputType(BlockResponseCodePtrOutput{})
	pulumi.RegisterOutputType(DnsSecurityRuleStateEnumOutput{})
	pulumi.RegisterOutputType(DnsSecurityRuleStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ForwardingRuleStateEnumOutput{})
	pulumi.RegisterOutputType(ForwardingRuleStateEnumPtrOutput{})
	pulumi.RegisterOutputType(IpAllocationMethodOutput{})
	pulumi.RegisterOutputType(IpAllocationMethodPtrOutput{})
}
