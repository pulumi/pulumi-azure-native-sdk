// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
type AccountType string

const (
	AccountType_Standard_LRS   = AccountType("Standard_LRS")
	AccountType_Standard_ZRS   = AccountType("Standard_ZRS")
	AccountType_Standard_GRS   = AccountType("Standard_GRS")
	AccountType_Standard_RAGRS = AccountType("Standard_RAGRS")
	AccountType_Premium_LRS    = AccountType("Premium_LRS")
)

func (AccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountType)(nil)).Elem()
}

func (e AccountType) ToAccountTypeOutput() AccountTypeOutput {
	return pulumi.ToOutput(e).(AccountTypeOutput)
}

func (e AccountType) ToAccountTypeOutputWithContext(ctx context.Context) AccountTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccountTypeOutput)
}

func (e AccountType) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return e.ToAccountTypePtrOutputWithContext(context.Background())
}

func (e AccountType) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return AccountType(e).ToAccountTypeOutputWithContext(ctx).ToAccountTypePtrOutputWithContext(ctx)
}

func (e AccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccountTypeOutput struct{ *pulumi.OutputState }

func (AccountTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountType)(nil)).Elem()
}

func (o AccountTypeOutput) ToAccountTypeOutput() AccountTypeOutput {
	return o
}

func (o AccountTypeOutput) ToAccountTypeOutputWithContext(ctx context.Context) AccountTypeOutput {
	return o
}

func (o AccountTypeOutput) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return o.ToAccountTypePtrOutputWithContext(context.Background())
}

func (o AccountTypeOutput) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountType) *AccountType {
		return &v
	}).(AccountTypePtrOutput)
}

func (o AccountTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccountTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccountType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccountTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccountTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccountType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccountTypePtrOutput struct{ *pulumi.OutputState }

func (AccountTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountType)(nil)).Elem()
}

func (o AccountTypePtrOutput) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return o
}

func (o AccountTypePtrOutput) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return o
}

func (o AccountTypePtrOutput) Elem() AccountTypeOutput {
	return o.ApplyT(func(v *AccountType) AccountType {
		if v != nil {
			return *v
		}
		var ret AccountType
		return ret
	}).(AccountTypeOutput)
}

func (o AccountTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccountTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccountType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccountTypeInput is an input type that accepts AccountTypeArgs and AccountTypeOutput values.
// You can construct a concrete instance of `AccountTypeInput` via:
//
//	AccountTypeArgs{...}
type AccountTypeInput interface {
	pulumi.Input

	ToAccountTypeOutput() AccountTypeOutput
	ToAccountTypeOutputWithContext(context.Context) AccountTypeOutput
}

var accountTypePtrType = reflect.TypeOf((**AccountType)(nil)).Elem()

type AccountTypePtrInput interface {
	pulumi.Input

	ToAccountTypePtrOutput() AccountTypePtrOutput
	ToAccountTypePtrOutputWithContext(context.Context) AccountTypePtrOutput
}

type accountTypePtr string

func AccountTypePtr(v string) AccountTypePtrInput {
	return (*accountTypePtr)(&v)
}

func (*accountTypePtr) ElementType() reflect.Type {
	return accountTypePtrType
}

func (in *accountTypePtr) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return pulumi.ToOutput(in).(AccountTypePtrOutput)
}

func (in *accountTypePtr) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccountTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountTypeOutput{})
	pulumi.RegisterOutputType(AccountTypePtrOutput{})
}