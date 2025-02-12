// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Direction that specifies whether the access rules is inbound/outbound.
type AccessRuleDirection string

const (
	AccessRuleDirectionInbound  = AccessRuleDirection("Inbound")
	AccessRuleDirectionOutbound = AccessRuleDirection("Outbound")
)

func (AccessRuleDirection) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRuleDirection)(nil)).Elem()
}

func (e AccessRuleDirection) ToAccessRuleDirectionOutput() AccessRuleDirectionOutput {
	return pulumi.ToOutput(e).(AccessRuleDirectionOutput)
}

func (e AccessRuleDirection) ToAccessRuleDirectionOutputWithContext(ctx context.Context) AccessRuleDirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessRuleDirectionOutput)
}

func (e AccessRuleDirection) ToAccessRuleDirectionPtrOutput() AccessRuleDirectionPtrOutput {
	return e.ToAccessRuleDirectionPtrOutputWithContext(context.Background())
}

func (e AccessRuleDirection) ToAccessRuleDirectionPtrOutputWithContext(ctx context.Context) AccessRuleDirectionPtrOutput {
	return AccessRuleDirection(e).ToAccessRuleDirectionOutputWithContext(ctx).ToAccessRuleDirectionPtrOutputWithContext(ctx)
}

func (e AccessRuleDirection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRuleDirection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRuleDirection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessRuleDirection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessRuleDirectionOutput struct{ *pulumi.OutputState }

func (AccessRuleDirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRuleDirection)(nil)).Elem()
}

func (o AccessRuleDirectionOutput) ToAccessRuleDirectionOutput() AccessRuleDirectionOutput {
	return o
}

func (o AccessRuleDirectionOutput) ToAccessRuleDirectionOutputWithContext(ctx context.Context) AccessRuleDirectionOutput {
	return o
}

func (o AccessRuleDirectionOutput) ToAccessRuleDirectionPtrOutput() AccessRuleDirectionPtrOutput {
	return o.ToAccessRuleDirectionPtrOutputWithContext(context.Background())
}

func (o AccessRuleDirectionOutput) ToAccessRuleDirectionPtrOutputWithContext(ctx context.Context) AccessRuleDirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRuleDirection) *AccessRuleDirection {
		return &v
	}).(AccessRuleDirectionPtrOutput)
}

func (o AccessRuleDirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessRuleDirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRuleDirection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessRuleDirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRuleDirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRuleDirection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessRuleDirectionPtrOutput struct{ *pulumi.OutputState }

func (AccessRuleDirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRuleDirection)(nil)).Elem()
}

func (o AccessRuleDirectionPtrOutput) ToAccessRuleDirectionPtrOutput() AccessRuleDirectionPtrOutput {
	return o
}

func (o AccessRuleDirectionPtrOutput) ToAccessRuleDirectionPtrOutputWithContext(ctx context.Context) AccessRuleDirectionPtrOutput {
	return o
}

func (o AccessRuleDirectionPtrOutput) Elem() AccessRuleDirectionOutput {
	return o.ApplyT(func(v *AccessRuleDirection) AccessRuleDirection {
		if v != nil {
			return *v
		}
		var ret AccessRuleDirection
		return ret
	}).(AccessRuleDirectionOutput)
}

func (o AccessRuleDirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRuleDirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessRuleDirection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessRuleDirectionInput is an input type that accepts values of the AccessRuleDirection enum
// A concrete instance of `AccessRuleDirectionInput` can be one of the following:
//
//	AccessRuleDirectionInbound
//	AccessRuleDirectionOutbound
type AccessRuleDirectionInput interface {
	pulumi.Input

	ToAccessRuleDirectionOutput() AccessRuleDirectionOutput
	ToAccessRuleDirectionOutputWithContext(context.Context) AccessRuleDirectionOutput
}

var accessRuleDirectionPtrType = reflect.TypeOf((**AccessRuleDirection)(nil)).Elem()

type AccessRuleDirectionPtrInput interface {
	pulumi.Input

	ToAccessRuleDirectionPtrOutput() AccessRuleDirectionPtrOutput
	ToAccessRuleDirectionPtrOutputWithContext(context.Context) AccessRuleDirectionPtrOutput
}

type accessRuleDirectionPtr string

func AccessRuleDirectionPtr(v string) AccessRuleDirectionPtrInput {
	return (*accessRuleDirectionPtr)(&v)
}

func (*accessRuleDirectionPtr) ElementType() reflect.Type {
	return accessRuleDirectionPtrType
}

func (in *accessRuleDirectionPtr) ToAccessRuleDirectionPtrOutput() AccessRuleDirectionPtrOutput {
	return pulumi.ToOutput(in).(AccessRuleDirectionPtrOutput)
}

func (in *accessRuleDirectionPtr) ToAccessRuleDirectionPtrOutputWithContext(ctx context.Context) AccessRuleDirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessRuleDirectionPtrOutput)
}

// Access mode on the association.
type AssociationAccessMode string

const (
	AssociationAccessModeLearning = AssociationAccessMode("Learning")
	AssociationAccessModeEnforced = AssociationAccessMode("Enforced")
	AssociationAccessModeAudit    = AssociationAccessMode("Audit")
)

func (AssociationAccessMode) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationAccessMode)(nil)).Elem()
}

func (e AssociationAccessMode) ToAssociationAccessModeOutput() AssociationAccessModeOutput {
	return pulumi.ToOutput(e).(AssociationAccessModeOutput)
}

func (e AssociationAccessMode) ToAssociationAccessModeOutputWithContext(ctx context.Context) AssociationAccessModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssociationAccessModeOutput)
}

func (e AssociationAccessMode) ToAssociationAccessModePtrOutput() AssociationAccessModePtrOutput {
	return e.ToAssociationAccessModePtrOutputWithContext(context.Background())
}

func (e AssociationAccessMode) ToAssociationAccessModePtrOutputWithContext(ctx context.Context) AssociationAccessModePtrOutput {
	return AssociationAccessMode(e).ToAssociationAccessModeOutputWithContext(ctx).ToAssociationAccessModePtrOutputWithContext(ctx)
}

func (e AssociationAccessMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssociationAccessMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssociationAccessMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssociationAccessMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssociationAccessModeOutput struct{ *pulumi.OutputState }

func (AssociationAccessModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationAccessMode)(nil)).Elem()
}

func (o AssociationAccessModeOutput) ToAssociationAccessModeOutput() AssociationAccessModeOutput {
	return o
}

func (o AssociationAccessModeOutput) ToAssociationAccessModeOutputWithContext(ctx context.Context) AssociationAccessModeOutput {
	return o
}

func (o AssociationAccessModeOutput) ToAssociationAccessModePtrOutput() AssociationAccessModePtrOutput {
	return o.ToAssociationAccessModePtrOutputWithContext(context.Background())
}

func (o AssociationAccessModeOutput) ToAssociationAccessModePtrOutputWithContext(ctx context.Context) AssociationAccessModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssociationAccessMode) *AssociationAccessMode {
		return &v
	}).(AssociationAccessModePtrOutput)
}

func (o AssociationAccessModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssociationAccessModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssociationAccessMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssociationAccessModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssociationAccessModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssociationAccessMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssociationAccessModePtrOutput struct{ *pulumi.OutputState }

func (AssociationAccessModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationAccessMode)(nil)).Elem()
}

func (o AssociationAccessModePtrOutput) ToAssociationAccessModePtrOutput() AssociationAccessModePtrOutput {
	return o
}

func (o AssociationAccessModePtrOutput) ToAssociationAccessModePtrOutputWithContext(ctx context.Context) AssociationAccessModePtrOutput {
	return o
}

func (o AssociationAccessModePtrOutput) Elem() AssociationAccessModeOutput {
	return o.ApplyT(func(v *AssociationAccessMode) AssociationAccessMode {
		if v != nil {
			return *v
		}
		var ret AssociationAccessMode
		return ret
	}).(AssociationAccessModeOutput)
}

func (o AssociationAccessModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssociationAccessModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssociationAccessMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AssociationAccessModeInput is an input type that accepts values of the AssociationAccessMode enum
// A concrete instance of `AssociationAccessModeInput` can be one of the following:
//
//	AssociationAccessModeLearning
//	AssociationAccessModeEnforced
//	AssociationAccessModeAudit
type AssociationAccessModeInput interface {
	pulumi.Input

	ToAssociationAccessModeOutput() AssociationAccessModeOutput
	ToAssociationAccessModeOutputWithContext(context.Context) AssociationAccessModeOutput
}

var associationAccessModePtrType = reflect.TypeOf((**AssociationAccessMode)(nil)).Elem()

type AssociationAccessModePtrInput interface {
	pulumi.Input

	ToAssociationAccessModePtrOutput() AssociationAccessModePtrOutput
	ToAssociationAccessModePtrOutputWithContext(context.Context) AssociationAccessModePtrOutput
}

type associationAccessModePtr string

func AssociationAccessModePtr(v string) AssociationAccessModePtrInput {
	return (*associationAccessModePtr)(&v)
}

func (*associationAccessModePtr) ElementType() reflect.Type {
	return associationAccessModePtrType
}

func (in *associationAccessModePtr) ToAssociationAccessModePtrOutput() AssociationAccessModePtrOutput {
	return pulumi.ToOutput(in).(AssociationAccessModePtrOutput)
}

func (in *associationAccessModePtr) ToAssociationAccessModePtrOutputWithContext(ctx context.Context) AssociationAccessModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssociationAccessModePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessRuleDirectionOutput{})
	pulumi.RegisterOutputType(AccessRuleDirectionPtrOutput{})
	pulumi.RegisterOutputType(AssociationAccessModeOutput{})
	pulumi.RegisterOutputType(AssociationAccessModePtrOutput{})
}
