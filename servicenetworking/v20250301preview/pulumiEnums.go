// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Association Type
type AssociationType string

const (
	// Association of Type Subnet
	AssociationTypeSubnets = AssociationType("subnets")
)

func (AssociationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationType)(nil)).Elem()
}

func (e AssociationType) ToAssociationTypeOutput() AssociationTypeOutput {
	return pulumi.ToOutput(e).(AssociationTypeOutput)
}

func (e AssociationType) ToAssociationTypeOutputWithContext(ctx context.Context) AssociationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssociationTypeOutput)
}

func (e AssociationType) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return e.ToAssociationTypePtrOutputWithContext(context.Background())
}

func (e AssociationType) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return AssociationType(e).ToAssociationTypeOutputWithContext(ctx).ToAssociationTypePtrOutputWithContext(ctx)
}

func (e AssociationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssociationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssociationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssociationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssociationTypeOutput struct{ *pulumi.OutputState }

func (AssociationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationType)(nil)).Elem()
}

func (o AssociationTypeOutput) ToAssociationTypeOutput() AssociationTypeOutput {
	return o
}

func (o AssociationTypeOutput) ToAssociationTypeOutputWithContext(ctx context.Context) AssociationTypeOutput {
	return o
}

func (o AssociationTypeOutput) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return o.ToAssociationTypePtrOutputWithContext(context.Background())
}

func (o AssociationTypeOutput) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssociationType) *AssociationType {
		return &v
	}).(AssociationTypePtrOutput)
}

func (o AssociationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssociationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssociationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssociationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssociationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssociationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssociationTypePtrOutput struct{ *pulumi.OutputState }

func (AssociationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationType)(nil)).Elem()
}

func (o AssociationTypePtrOutput) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return o
}

func (o AssociationTypePtrOutput) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return o
}

func (o AssociationTypePtrOutput) Elem() AssociationTypeOutput {
	return o.ApplyT(func(v *AssociationType) AssociationType {
		if v != nil {
			return *v
		}
		var ret AssociationType
		return ret
	}).(AssociationTypeOutput)
}

func (o AssociationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssociationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssociationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssociationTypeInput is an input type that accepts values of the AssociationType enum
// A concrete instance of `AssociationTypeInput` can be one of the following:
//
//	AssociationTypeSubnets
type AssociationTypeInput interface {
	pulumi.Input

	ToAssociationTypeOutput() AssociationTypeOutput
	ToAssociationTypeOutputWithContext(context.Context) AssociationTypeOutput
}

var associationTypePtrType = reflect.TypeOf((**AssociationType)(nil)).Elem()

type AssociationTypePtrInput interface {
	pulumi.Input

	ToAssociationTypePtrOutput() AssociationTypePtrOutput
	ToAssociationTypePtrOutputWithContext(context.Context) AssociationTypePtrOutput
}

type associationTypePtr string

func AssociationTypePtr(v string) AssociationTypePtrInput {
	return (*associationTypePtr)(&v)
}

func (*associationTypePtr) ElementType() reflect.Type {
	return associationTypePtrType
}

func (in *associationTypePtr) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return pulumi.ToOutput(in).(AssociationTypePtrOutput)
}

func (in *associationTypePtr) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssociationTypePtrOutput)
}

// Action of the Rule
type IpAccessRuleAction string

const (
	// Allow Source Ip Prefixes
	IpAccessRuleActionAllow = IpAccessRuleAction("allow")
	// Deny Source Ip Prefixes
	IpAccessRuleActionDeny = IpAccessRuleAction("deny")
)

func (IpAccessRuleAction) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAccessRuleAction)(nil)).Elem()
}

func (e IpAccessRuleAction) ToIpAccessRuleActionOutput() IpAccessRuleActionOutput {
	return pulumi.ToOutput(e).(IpAccessRuleActionOutput)
}

func (e IpAccessRuleAction) ToIpAccessRuleActionOutputWithContext(ctx context.Context) IpAccessRuleActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpAccessRuleActionOutput)
}

func (e IpAccessRuleAction) ToIpAccessRuleActionPtrOutput() IpAccessRuleActionPtrOutput {
	return e.ToIpAccessRuleActionPtrOutputWithContext(context.Background())
}

func (e IpAccessRuleAction) ToIpAccessRuleActionPtrOutputWithContext(ctx context.Context) IpAccessRuleActionPtrOutput {
	return IpAccessRuleAction(e).ToIpAccessRuleActionOutputWithContext(ctx).ToIpAccessRuleActionPtrOutputWithContext(ctx)
}

func (e IpAccessRuleAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAccessRuleAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAccessRuleAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpAccessRuleAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpAccessRuleActionOutput struct{ *pulumi.OutputState }

func (IpAccessRuleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAccessRuleAction)(nil)).Elem()
}

func (o IpAccessRuleActionOutput) ToIpAccessRuleActionOutput() IpAccessRuleActionOutput {
	return o
}

func (o IpAccessRuleActionOutput) ToIpAccessRuleActionOutputWithContext(ctx context.Context) IpAccessRuleActionOutput {
	return o
}

func (o IpAccessRuleActionOutput) ToIpAccessRuleActionPtrOutput() IpAccessRuleActionPtrOutput {
	return o.ToIpAccessRuleActionPtrOutputWithContext(context.Background())
}

func (o IpAccessRuleActionOutput) ToIpAccessRuleActionPtrOutputWithContext(ctx context.Context) IpAccessRuleActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAccessRuleAction) *IpAccessRuleAction {
		return &v
	}).(IpAccessRuleActionPtrOutput)
}

func (o IpAccessRuleActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpAccessRuleActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAccessRuleAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpAccessRuleActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAccessRuleActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAccessRuleAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpAccessRuleActionPtrOutput struct{ *pulumi.OutputState }

func (IpAccessRuleActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAccessRuleAction)(nil)).Elem()
}

func (o IpAccessRuleActionPtrOutput) ToIpAccessRuleActionPtrOutput() IpAccessRuleActionPtrOutput {
	return o
}

func (o IpAccessRuleActionPtrOutput) ToIpAccessRuleActionPtrOutputWithContext(ctx context.Context) IpAccessRuleActionPtrOutput {
	return o
}

func (o IpAccessRuleActionPtrOutput) Elem() IpAccessRuleActionOutput {
	return o.ApplyT(func(v *IpAccessRuleAction) IpAccessRuleAction {
		if v != nil {
			return *v
		}
		var ret IpAccessRuleAction
		return ret
	}).(IpAccessRuleActionOutput)
}

func (o IpAccessRuleActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAccessRuleActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpAccessRuleAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IpAccessRuleActionInput is an input type that accepts values of the IpAccessRuleAction enum
// A concrete instance of `IpAccessRuleActionInput` can be one of the following:
//
//	IpAccessRuleActionAllow
//	IpAccessRuleActionDeny
type IpAccessRuleActionInput interface {
	pulumi.Input

	ToIpAccessRuleActionOutput() IpAccessRuleActionOutput
	ToIpAccessRuleActionOutputWithContext(context.Context) IpAccessRuleActionOutput
}

var ipAccessRuleActionPtrType = reflect.TypeOf((**IpAccessRuleAction)(nil)).Elem()

type IpAccessRuleActionPtrInput interface {
	pulumi.Input

	ToIpAccessRuleActionPtrOutput() IpAccessRuleActionPtrOutput
	ToIpAccessRuleActionPtrOutputWithContext(context.Context) IpAccessRuleActionPtrOutput
}

type ipAccessRuleActionPtr string

func IpAccessRuleActionPtr(v string) IpAccessRuleActionPtrInput {
	return (*ipAccessRuleActionPtr)(&v)
}

func (*ipAccessRuleActionPtr) ElementType() reflect.Type {
	return ipAccessRuleActionPtrType
}

func (in *ipAccessRuleActionPtr) ToIpAccessRuleActionPtrOutput() IpAccessRuleActionPtrOutput {
	return pulumi.ToOutput(in).(IpAccessRuleActionPtrOutput)
}

func (in *ipAccessRuleActionPtr) ToIpAccessRuleActionPtrOutputWithContext(ctx context.Context) IpAccessRuleActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpAccessRuleActionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssociationTypeOutput{})
	pulumi.RegisterOutputType(AssociationTypePtrOutput{})
	pulumi.RegisterOutputType(IpAccessRuleActionOutput{})
	pulumi.RegisterOutputType(IpAccessRuleActionPtrOutput{})
}
