// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240930preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Profile type of the certificate.
type ProfileType string

const (
	// Used for signing files which are distributed publicly.
	ProfileTypePublicTrust = ProfileType("PublicTrust")
	// Used for signing files which are distributed internally within organization or group boundary.
	ProfileTypePrivateTrust = ProfileType("PrivateTrust")
	// Used for signing CI policy files.
	ProfileTypePrivateTrustCIPolicy = ProfileType("PrivateTrustCIPolicy")
	// Used for signing files which are run in secure vbs enclave.
	ProfileTypeVBSEnclave = ProfileType("VBSEnclave")
	// Used for signing files for testing purpose.
	ProfileTypePublicTrustTest = ProfileType("PublicTrustTest")
)

func (ProfileType) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileType)(nil)).Elem()
}

func (e ProfileType) ToProfileTypeOutput() ProfileTypeOutput {
	return pulumi.ToOutput(e).(ProfileTypeOutput)
}

func (e ProfileType) ToProfileTypeOutputWithContext(ctx context.Context) ProfileTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProfileTypeOutput)
}

func (e ProfileType) ToProfileTypePtrOutput() ProfileTypePtrOutput {
	return e.ToProfileTypePtrOutputWithContext(context.Background())
}

func (e ProfileType) ToProfileTypePtrOutputWithContext(ctx context.Context) ProfileTypePtrOutput {
	return ProfileType(e).ToProfileTypeOutputWithContext(ctx).ToProfileTypePtrOutputWithContext(ctx)
}

func (e ProfileType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProfileType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProfileType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProfileType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProfileTypeOutput struct{ *pulumi.OutputState }

func (ProfileTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileType)(nil)).Elem()
}

func (o ProfileTypeOutput) ToProfileTypeOutput() ProfileTypeOutput {
	return o
}

func (o ProfileTypeOutput) ToProfileTypeOutputWithContext(ctx context.Context) ProfileTypeOutput {
	return o
}

func (o ProfileTypeOutput) ToProfileTypePtrOutput() ProfileTypePtrOutput {
	return o.ToProfileTypePtrOutputWithContext(context.Background())
}

func (o ProfileTypeOutput) ToProfileTypePtrOutputWithContext(ctx context.Context) ProfileTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileType) *ProfileType {
		return &v
	}).(ProfileTypePtrOutput)
}

func (o ProfileTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProfileTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProfileType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProfileTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProfileTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProfileType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProfileTypePtrOutput struct{ *pulumi.OutputState }

func (ProfileTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileType)(nil)).Elem()
}

func (o ProfileTypePtrOutput) ToProfileTypePtrOutput() ProfileTypePtrOutput {
	return o
}

func (o ProfileTypePtrOutput) ToProfileTypePtrOutputWithContext(ctx context.Context) ProfileTypePtrOutput {
	return o
}

func (o ProfileTypePtrOutput) Elem() ProfileTypeOutput {
	return o.ApplyT(func(v *ProfileType) ProfileType {
		if v != nil {
			return *v
		}
		var ret ProfileType
		return ret
	}).(ProfileTypeOutput)
}

func (o ProfileTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProfileTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProfileType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProfileTypeInput is an input type that accepts values of the ProfileType enum
// A concrete instance of `ProfileTypeInput` can be one of the following:
//
//	ProfileTypePublicTrust
//	ProfileTypePrivateTrust
//	ProfileTypePrivateTrustCIPolicy
//	ProfileTypeVBSEnclave
//	ProfileTypePublicTrustTest
type ProfileTypeInput interface {
	pulumi.Input

	ToProfileTypeOutput() ProfileTypeOutput
	ToProfileTypeOutputWithContext(context.Context) ProfileTypeOutput
}

var profileTypePtrType = reflect.TypeOf((**ProfileType)(nil)).Elem()

type ProfileTypePtrInput interface {
	pulumi.Input

	ToProfileTypePtrOutput() ProfileTypePtrOutput
	ToProfileTypePtrOutputWithContext(context.Context) ProfileTypePtrOutput
}

type profileTypePtr string

func ProfileTypePtr(v string) ProfileTypePtrInput {
	return (*profileTypePtr)(&v)
}

func (*profileTypePtr) ElementType() reflect.Type {
	return profileTypePtrType
}

func (in *profileTypePtr) ToProfileTypePtrOutput() ProfileTypePtrOutput {
	return pulumi.ToOutput(in).(ProfileTypePtrOutput)
}

func (in *profileTypePtr) ToProfileTypePtrOutputWithContext(ctx context.Context) ProfileTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProfileTypePtrOutput)
}

// Name of the SKU.
type SkuName string

const (
	// Basic sku.
	SkuNameBasic = SkuName("Basic")
	// Premium sku.
	SkuNamePremium = SkuName("Premium")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SkuNameInput is an input type that accepts values of the SkuName enum
// A concrete instance of `SkuNameInput` can be one of the following:
//
//	SkuNameBasic
//	SkuNamePremium
type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileTypeOutput{})
	pulumi.RegisterOutputType(ProfileTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
