// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azureactivedirectory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The name of the SKU for the tenant.
type B2CResourceSKUName string

const (
	// Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users (MAU) billing model.
	B2CResourceSKUNameStandard = B2CResourceSKUName("Standard")
	// Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications based billing.
	B2CResourceSKUNamePremiumP1 = B2CResourceSKUName("PremiumP1")
	// Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications based billing.
	B2CResourceSKUNamePremiumP2 = B2CResourceSKUName("PremiumP2")
)

func (B2CResourceSKUName) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKUName)(nil)).Elem()
}

func (e B2CResourceSKUName) ToB2CResourceSKUNameOutput() B2CResourceSKUNameOutput {
	return pulumi.ToOutput(e).(B2CResourceSKUNameOutput)
}

func (e B2CResourceSKUName) ToB2CResourceSKUNameOutputWithContext(ctx context.Context) B2CResourceSKUNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(B2CResourceSKUNameOutput)
}

func (e B2CResourceSKUName) ToB2CResourceSKUNamePtrOutput() B2CResourceSKUNamePtrOutput {
	return e.ToB2CResourceSKUNamePtrOutputWithContext(context.Background())
}

func (e B2CResourceSKUName) ToB2CResourceSKUNamePtrOutputWithContext(ctx context.Context) B2CResourceSKUNamePtrOutput {
	return B2CResourceSKUName(e).ToB2CResourceSKUNameOutputWithContext(ctx).ToB2CResourceSKUNamePtrOutputWithContext(ctx)
}

func (e B2CResourceSKUName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e B2CResourceSKUName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e B2CResourceSKUName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e B2CResourceSKUName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type B2CResourceSKUNameOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKUName)(nil)).Elem()
}

func (o B2CResourceSKUNameOutput) ToB2CResourceSKUNameOutput() B2CResourceSKUNameOutput {
	return o
}

func (o B2CResourceSKUNameOutput) ToB2CResourceSKUNameOutputWithContext(ctx context.Context) B2CResourceSKUNameOutput {
	return o
}

func (o B2CResourceSKUNameOutput) ToB2CResourceSKUNamePtrOutput() B2CResourceSKUNamePtrOutput {
	return o.ToB2CResourceSKUNamePtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUNameOutput) ToB2CResourceSKUNamePtrOutputWithContext(ctx context.Context) B2CResourceSKUNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2CResourceSKUName) *B2CResourceSKUName {
		return &v
	}).(B2CResourceSKUNamePtrOutput)
}

func (o B2CResourceSKUNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o B2CResourceSKUNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e B2CResourceSKUName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o B2CResourceSKUNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e B2CResourceSKUName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type B2CResourceSKUNamePtrOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CResourceSKUName)(nil)).Elem()
}

func (o B2CResourceSKUNamePtrOutput) ToB2CResourceSKUNamePtrOutput() B2CResourceSKUNamePtrOutput {
	return o
}

func (o B2CResourceSKUNamePtrOutput) ToB2CResourceSKUNamePtrOutputWithContext(ctx context.Context) B2CResourceSKUNamePtrOutput {
	return o
}

func (o B2CResourceSKUNamePtrOutput) Elem() B2CResourceSKUNameOutput {
	return o.ApplyT(func(v *B2CResourceSKUName) B2CResourceSKUName {
		if v != nil {
			return *v
		}
		var ret B2CResourceSKUName
		return ret
	}).(B2CResourceSKUNameOutput)
}

func (o B2CResourceSKUNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *B2CResourceSKUName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// B2CResourceSKUNameInput is an input type that accepts values of the B2CResourceSKUName enum
// A concrete instance of `B2CResourceSKUNameInput` can be one of the following:
//
//	B2CResourceSKUNameStandard
//	B2CResourceSKUNamePremiumP1
//	B2CResourceSKUNamePremiumP2
type B2CResourceSKUNameInput interface {
	pulumi.Input

	ToB2CResourceSKUNameOutput() B2CResourceSKUNameOutput
	ToB2CResourceSKUNameOutputWithContext(context.Context) B2CResourceSKUNameOutput
}

var b2cresourceSKUNamePtrType = reflect.TypeOf((**B2CResourceSKUName)(nil)).Elem()

type B2CResourceSKUNamePtrInput interface {
	pulumi.Input

	ToB2CResourceSKUNamePtrOutput() B2CResourceSKUNamePtrOutput
	ToB2CResourceSKUNamePtrOutputWithContext(context.Context) B2CResourceSKUNamePtrOutput
}

type b2cresourceSKUNamePtr string

func B2CResourceSKUNamePtr(v string) B2CResourceSKUNamePtrInput {
	return (*b2cresourceSKUNamePtr)(&v)
}

func (*b2cresourceSKUNamePtr) ElementType() reflect.Type {
	return b2cresourceSKUNamePtrType
}

func (in *b2cresourceSKUNamePtr) ToB2CResourceSKUNamePtrOutput() B2CResourceSKUNamePtrOutput {
	return pulumi.ToOutput(in).(B2CResourceSKUNamePtrOutput)
}

func (in *b2cresourceSKUNamePtr) ToB2CResourceSKUNamePtrOutputWithContext(ctx context.Context) B2CResourceSKUNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(B2CResourceSKUNamePtrOutput)
}

// The tier of the tenant.
type B2CResourceSKUTier string

const (
	// The SKU tier used for all Azure AD B2C tenants.
	B2CResourceSKUTierA0 = B2CResourceSKUTier("A0")
)

func (B2CResourceSKUTier) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKUTier)(nil)).Elem()
}

func (e B2CResourceSKUTier) ToB2CResourceSKUTierOutput() B2CResourceSKUTierOutput {
	return pulumi.ToOutput(e).(B2CResourceSKUTierOutput)
}

func (e B2CResourceSKUTier) ToB2CResourceSKUTierOutputWithContext(ctx context.Context) B2CResourceSKUTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(B2CResourceSKUTierOutput)
}

func (e B2CResourceSKUTier) ToB2CResourceSKUTierPtrOutput() B2CResourceSKUTierPtrOutput {
	return e.ToB2CResourceSKUTierPtrOutputWithContext(context.Background())
}

func (e B2CResourceSKUTier) ToB2CResourceSKUTierPtrOutputWithContext(ctx context.Context) B2CResourceSKUTierPtrOutput {
	return B2CResourceSKUTier(e).ToB2CResourceSKUTierOutputWithContext(ctx).ToB2CResourceSKUTierPtrOutputWithContext(ctx)
}

func (e B2CResourceSKUTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e B2CResourceSKUTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e B2CResourceSKUTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e B2CResourceSKUTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type B2CResourceSKUTierOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKUTier)(nil)).Elem()
}

func (o B2CResourceSKUTierOutput) ToB2CResourceSKUTierOutput() B2CResourceSKUTierOutput {
	return o
}

func (o B2CResourceSKUTierOutput) ToB2CResourceSKUTierOutputWithContext(ctx context.Context) B2CResourceSKUTierOutput {
	return o
}

func (o B2CResourceSKUTierOutput) ToB2CResourceSKUTierPtrOutput() B2CResourceSKUTierPtrOutput {
	return o.ToB2CResourceSKUTierPtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUTierOutput) ToB2CResourceSKUTierPtrOutputWithContext(ctx context.Context) B2CResourceSKUTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2CResourceSKUTier) *B2CResourceSKUTier {
		return &v
	}).(B2CResourceSKUTierPtrOutput)
}

func (o B2CResourceSKUTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o B2CResourceSKUTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e B2CResourceSKUTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o B2CResourceSKUTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e B2CResourceSKUTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type B2CResourceSKUTierPtrOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CResourceSKUTier)(nil)).Elem()
}

func (o B2CResourceSKUTierPtrOutput) ToB2CResourceSKUTierPtrOutput() B2CResourceSKUTierPtrOutput {
	return o
}

func (o B2CResourceSKUTierPtrOutput) ToB2CResourceSKUTierPtrOutputWithContext(ctx context.Context) B2CResourceSKUTierPtrOutput {
	return o
}

func (o B2CResourceSKUTierPtrOutput) Elem() B2CResourceSKUTierOutput {
	return o.ApplyT(func(v *B2CResourceSKUTier) B2CResourceSKUTier {
		if v != nil {
			return *v
		}
		var ret B2CResourceSKUTier
		return ret
	}).(B2CResourceSKUTierOutput)
}

func (o B2CResourceSKUTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *B2CResourceSKUTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// B2CResourceSKUTierInput is an input type that accepts values of the B2CResourceSKUTier enum
// A concrete instance of `B2CResourceSKUTierInput` can be one of the following:
//
//	B2CResourceSKUTierA0
type B2CResourceSKUTierInput interface {
	pulumi.Input

	ToB2CResourceSKUTierOutput() B2CResourceSKUTierOutput
	ToB2CResourceSKUTierOutputWithContext(context.Context) B2CResourceSKUTierOutput
}

var b2cresourceSKUTierPtrType = reflect.TypeOf((**B2CResourceSKUTier)(nil)).Elem()

type B2CResourceSKUTierPtrInput interface {
	pulumi.Input

	ToB2CResourceSKUTierPtrOutput() B2CResourceSKUTierPtrOutput
	ToB2CResourceSKUTierPtrOutputWithContext(context.Context) B2CResourceSKUTierPtrOutput
}

type b2cresourceSKUTierPtr string

func B2CResourceSKUTierPtr(v string) B2CResourceSKUTierPtrInput {
	return (*b2cresourceSKUTierPtr)(&v)
}

func (*b2cresourceSKUTierPtr) ElementType() reflect.Type {
	return b2cresourceSKUTierPtrType
}

func (in *b2cresourceSKUTierPtr) ToB2CResourceSKUTierPtrOutput() B2CResourceSKUTierPtrOutput {
	return pulumi.ToOutput(in).(B2CResourceSKUTierPtrOutput)
}

func (in *b2cresourceSKUTierPtr) ToB2CResourceSKUTierPtrOutputWithContext(ctx context.Context) B2CResourceSKUTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(B2CResourceSKUTierPtrOutput)
}

// The name of the SKU for the tenant.
type CIAMResourceSKUName string

const (
	CIAMResourceSKUNameStandard  = CIAMResourceSKUName("Standard")
	CIAMResourceSKUNamePremiumP1 = CIAMResourceSKUName("PremiumP1")
	CIAMResourceSKUNamePremiumP2 = CIAMResourceSKUName("PremiumP2")
)

func (CIAMResourceSKUName) ElementType() reflect.Type {
	return reflect.TypeOf((*CIAMResourceSKUName)(nil)).Elem()
}

func (e CIAMResourceSKUName) ToCIAMResourceSKUNameOutput() CIAMResourceSKUNameOutput {
	return pulumi.ToOutput(e).(CIAMResourceSKUNameOutput)
}

func (e CIAMResourceSKUName) ToCIAMResourceSKUNameOutputWithContext(ctx context.Context) CIAMResourceSKUNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CIAMResourceSKUNameOutput)
}

func (e CIAMResourceSKUName) ToCIAMResourceSKUNamePtrOutput() CIAMResourceSKUNamePtrOutput {
	return e.ToCIAMResourceSKUNamePtrOutputWithContext(context.Background())
}

func (e CIAMResourceSKUName) ToCIAMResourceSKUNamePtrOutputWithContext(ctx context.Context) CIAMResourceSKUNamePtrOutput {
	return CIAMResourceSKUName(e).ToCIAMResourceSKUNameOutputWithContext(ctx).ToCIAMResourceSKUNamePtrOutputWithContext(ctx)
}

func (e CIAMResourceSKUName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CIAMResourceSKUName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CIAMResourceSKUName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CIAMResourceSKUName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CIAMResourceSKUNameOutput struct{ *pulumi.OutputState }

func (CIAMResourceSKUNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CIAMResourceSKUName)(nil)).Elem()
}

func (o CIAMResourceSKUNameOutput) ToCIAMResourceSKUNameOutput() CIAMResourceSKUNameOutput {
	return o
}

func (o CIAMResourceSKUNameOutput) ToCIAMResourceSKUNameOutputWithContext(ctx context.Context) CIAMResourceSKUNameOutput {
	return o
}

func (o CIAMResourceSKUNameOutput) ToCIAMResourceSKUNamePtrOutput() CIAMResourceSKUNamePtrOutput {
	return o.ToCIAMResourceSKUNamePtrOutputWithContext(context.Background())
}

func (o CIAMResourceSKUNameOutput) ToCIAMResourceSKUNamePtrOutputWithContext(ctx context.Context) CIAMResourceSKUNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CIAMResourceSKUName) *CIAMResourceSKUName {
		return &v
	}).(CIAMResourceSKUNamePtrOutput)
}

func (o CIAMResourceSKUNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CIAMResourceSKUNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CIAMResourceSKUName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CIAMResourceSKUNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CIAMResourceSKUNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CIAMResourceSKUName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CIAMResourceSKUNamePtrOutput struct{ *pulumi.OutputState }

func (CIAMResourceSKUNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CIAMResourceSKUName)(nil)).Elem()
}

func (o CIAMResourceSKUNamePtrOutput) ToCIAMResourceSKUNamePtrOutput() CIAMResourceSKUNamePtrOutput {
	return o
}

func (o CIAMResourceSKUNamePtrOutput) ToCIAMResourceSKUNamePtrOutputWithContext(ctx context.Context) CIAMResourceSKUNamePtrOutput {
	return o
}

func (o CIAMResourceSKUNamePtrOutput) Elem() CIAMResourceSKUNameOutput {
	return o.ApplyT(func(v *CIAMResourceSKUName) CIAMResourceSKUName {
		if v != nil {
			return *v
		}
		var ret CIAMResourceSKUName
		return ret
	}).(CIAMResourceSKUNameOutput)
}

func (o CIAMResourceSKUNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CIAMResourceSKUNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CIAMResourceSKUName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CIAMResourceSKUNameInput is an input type that accepts values of the CIAMResourceSKUName enum
// A concrete instance of `CIAMResourceSKUNameInput` can be one of the following:
//
//	CIAMResourceSKUNameStandard
//	CIAMResourceSKUNamePremiumP1
//	CIAMResourceSKUNamePremiumP2
type CIAMResourceSKUNameInput interface {
	pulumi.Input

	ToCIAMResourceSKUNameOutput() CIAMResourceSKUNameOutput
	ToCIAMResourceSKUNameOutputWithContext(context.Context) CIAMResourceSKUNameOutput
}

var ciamresourceSKUNamePtrType = reflect.TypeOf((**CIAMResourceSKUName)(nil)).Elem()

type CIAMResourceSKUNamePtrInput interface {
	pulumi.Input

	ToCIAMResourceSKUNamePtrOutput() CIAMResourceSKUNamePtrOutput
	ToCIAMResourceSKUNamePtrOutputWithContext(context.Context) CIAMResourceSKUNamePtrOutput
}

type ciamresourceSKUNamePtr string

func CIAMResourceSKUNamePtr(v string) CIAMResourceSKUNamePtrInput {
	return (*ciamresourceSKUNamePtr)(&v)
}

func (*ciamresourceSKUNamePtr) ElementType() reflect.Type {
	return ciamresourceSKUNamePtrType
}

func (in *ciamresourceSKUNamePtr) ToCIAMResourceSKUNamePtrOutput() CIAMResourceSKUNamePtrOutput {
	return pulumi.ToOutput(in).(CIAMResourceSKUNamePtrOutput)
}

func (in *ciamresourceSKUNamePtr) ToCIAMResourceSKUNamePtrOutputWithContext(ctx context.Context) CIAMResourceSKUNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CIAMResourceSKUNamePtrOutput)
}

// The tier of the tenant.
type CIAMResourceSKUTier string

const (
	// The SKU tier used for all Azure AD for customers tenants.
	CIAMResourceSKUTierA0 = CIAMResourceSKUTier("A0")
)

func (CIAMResourceSKUTier) ElementType() reflect.Type {
	return reflect.TypeOf((*CIAMResourceSKUTier)(nil)).Elem()
}

func (e CIAMResourceSKUTier) ToCIAMResourceSKUTierOutput() CIAMResourceSKUTierOutput {
	return pulumi.ToOutput(e).(CIAMResourceSKUTierOutput)
}

func (e CIAMResourceSKUTier) ToCIAMResourceSKUTierOutputWithContext(ctx context.Context) CIAMResourceSKUTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CIAMResourceSKUTierOutput)
}

func (e CIAMResourceSKUTier) ToCIAMResourceSKUTierPtrOutput() CIAMResourceSKUTierPtrOutput {
	return e.ToCIAMResourceSKUTierPtrOutputWithContext(context.Background())
}

func (e CIAMResourceSKUTier) ToCIAMResourceSKUTierPtrOutputWithContext(ctx context.Context) CIAMResourceSKUTierPtrOutput {
	return CIAMResourceSKUTier(e).ToCIAMResourceSKUTierOutputWithContext(ctx).ToCIAMResourceSKUTierPtrOutputWithContext(ctx)
}

func (e CIAMResourceSKUTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CIAMResourceSKUTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CIAMResourceSKUTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CIAMResourceSKUTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CIAMResourceSKUTierOutput struct{ *pulumi.OutputState }

func (CIAMResourceSKUTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CIAMResourceSKUTier)(nil)).Elem()
}

func (o CIAMResourceSKUTierOutput) ToCIAMResourceSKUTierOutput() CIAMResourceSKUTierOutput {
	return o
}

func (o CIAMResourceSKUTierOutput) ToCIAMResourceSKUTierOutputWithContext(ctx context.Context) CIAMResourceSKUTierOutput {
	return o
}

func (o CIAMResourceSKUTierOutput) ToCIAMResourceSKUTierPtrOutput() CIAMResourceSKUTierPtrOutput {
	return o.ToCIAMResourceSKUTierPtrOutputWithContext(context.Background())
}

func (o CIAMResourceSKUTierOutput) ToCIAMResourceSKUTierPtrOutputWithContext(ctx context.Context) CIAMResourceSKUTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CIAMResourceSKUTier) *CIAMResourceSKUTier {
		return &v
	}).(CIAMResourceSKUTierPtrOutput)
}

func (o CIAMResourceSKUTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CIAMResourceSKUTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CIAMResourceSKUTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CIAMResourceSKUTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CIAMResourceSKUTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CIAMResourceSKUTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CIAMResourceSKUTierPtrOutput struct{ *pulumi.OutputState }

func (CIAMResourceSKUTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CIAMResourceSKUTier)(nil)).Elem()
}

func (o CIAMResourceSKUTierPtrOutput) ToCIAMResourceSKUTierPtrOutput() CIAMResourceSKUTierPtrOutput {
	return o
}

func (o CIAMResourceSKUTierPtrOutput) ToCIAMResourceSKUTierPtrOutputWithContext(ctx context.Context) CIAMResourceSKUTierPtrOutput {
	return o
}

func (o CIAMResourceSKUTierPtrOutput) Elem() CIAMResourceSKUTierOutput {
	return o.ApplyT(func(v *CIAMResourceSKUTier) CIAMResourceSKUTier {
		if v != nil {
			return *v
		}
		var ret CIAMResourceSKUTier
		return ret
	}).(CIAMResourceSKUTierOutput)
}

func (o CIAMResourceSKUTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CIAMResourceSKUTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CIAMResourceSKUTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CIAMResourceSKUTierInput is an input type that accepts values of the CIAMResourceSKUTier enum
// A concrete instance of `CIAMResourceSKUTierInput` can be one of the following:
//
//	CIAMResourceSKUTierA0
type CIAMResourceSKUTierInput interface {
	pulumi.Input

	ToCIAMResourceSKUTierOutput() CIAMResourceSKUTierOutput
	ToCIAMResourceSKUTierOutputWithContext(context.Context) CIAMResourceSKUTierOutput
}

var ciamresourceSKUTierPtrType = reflect.TypeOf((**CIAMResourceSKUTier)(nil)).Elem()

type CIAMResourceSKUTierPtrInput interface {
	pulumi.Input

	ToCIAMResourceSKUTierPtrOutput() CIAMResourceSKUTierPtrOutput
	ToCIAMResourceSKUTierPtrOutputWithContext(context.Context) CIAMResourceSKUTierPtrOutput
}

type ciamresourceSKUTierPtr string

func CIAMResourceSKUTierPtr(v string) CIAMResourceSKUTierPtrInput {
	return (*ciamresourceSKUTierPtr)(&v)
}

func (*ciamresourceSKUTierPtr) ElementType() reflect.Type {
	return ciamresourceSKUTierPtrType
}

func (in *ciamresourceSKUTierPtr) ToCIAMResourceSKUTierPtrOutput() CIAMResourceSKUTierPtrOutput {
	return pulumi.ToOutput(in).(CIAMResourceSKUTierPtrOutput)
}

func (in *ciamresourceSKUTierPtr) ToCIAMResourceSKUTierPtrOutputWithContext(ctx context.Context) CIAMResourceSKUTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CIAMResourceSKUTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(B2CResourceSKUNameOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUNamePtrOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUTierOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUTierPtrOutput{})
	pulumi.RegisterOutputType(CIAMResourceSKUNameOutput{})
	pulumi.RegisterOutputType(CIAMResourceSKUNamePtrOutput{})
	pulumi.RegisterOutputType(CIAMResourceSKUTierOutput{})
	pulumi.RegisterOutputType(CIAMResourceSKUTierPtrOutput{})
}
