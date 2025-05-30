// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codesigning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// SKU of the trusted signing account.
type AccountSku struct {
	// Name of the SKU.
	Name string `pulumi:"name"`
}

// AccountSkuInput is an input type that accepts AccountSkuArgs and AccountSkuOutput values.
// You can construct a concrete instance of `AccountSkuInput` via:
//
//	AccountSkuArgs{...}
type AccountSkuInput interface {
	pulumi.Input

	ToAccountSkuOutput() AccountSkuOutput
	ToAccountSkuOutputWithContext(context.Context) AccountSkuOutput
}

// SKU of the trusted signing account.
type AccountSkuArgs struct {
	// Name of the SKU.
	Name pulumi.StringInput `pulumi:"name"`
}

func (AccountSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountSku)(nil)).Elem()
}

func (i AccountSkuArgs) ToAccountSkuOutput() AccountSkuOutput {
	return i.ToAccountSkuOutputWithContext(context.Background())
}

func (i AccountSkuArgs) ToAccountSkuOutputWithContext(ctx context.Context) AccountSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountSkuOutput)
}

func (i AccountSkuArgs) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return i.ToAccountSkuPtrOutputWithContext(context.Background())
}

func (i AccountSkuArgs) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountSkuOutput).ToAccountSkuPtrOutputWithContext(ctx)
}

// AccountSkuPtrInput is an input type that accepts AccountSkuArgs, AccountSkuPtr and AccountSkuPtrOutput values.
// You can construct a concrete instance of `AccountSkuPtrInput` via:
//
//	        AccountSkuArgs{...}
//
//	or:
//
//	        nil
type AccountSkuPtrInput interface {
	pulumi.Input

	ToAccountSkuPtrOutput() AccountSkuPtrOutput
	ToAccountSkuPtrOutputWithContext(context.Context) AccountSkuPtrOutput
}

type accountSkuPtrType AccountSkuArgs

func AccountSkuPtr(v *AccountSkuArgs) AccountSkuPtrInput {
	return (*accountSkuPtrType)(v)
}

func (*accountSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountSku)(nil)).Elem()
}

func (i *accountSkuPtrType) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return i.ToAccountSkuPtrOutputWithContext(context.Background())
}

func (i *accountSkuPtrType) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountSkuPtrOutput)
}

// SKU of the trusted signing account.
type AccountSkuOutput struct{ *pulumi.OutputState }

func (AccountSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountSku)(nil)).Elem()
}

func (o AccountSkuOutput) ToAccountSkuOutput() AccountSkuOutput {
	return o
}

func (o AccountSkuOutput) ToAccountSkuOutputWithContext(ctx context.Context) AccountSkuOutput {
	return o
}

func (o AccountSkuOutput) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return o.ToAccountSkuPtrOutputWithContext(context.Background())
}

func (o AccountSkuOutput) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountSku) *AccountSku {
		return &v
	}).(AccountSkuPtrOutput)
}

// Name of the SKU.
func (o AccountSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AccountSku) string { return v.Name }).(pulumi.StringOutput)
}

type AccountSkuPtrOutput struct{ *pulumi.OutputState }

func (AccountSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountSku)(nil)).Elem()
}

func (o AccountSkuPtrOutput) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return o
}

func (o AccountSkuPtrOutput) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return o
}

func (o AccountSkuPtrOutput) Elem() AccountSkuOutput {
	return o.ApplyT(func(v *AccountSku) AccountSku {
		if v != nil {
			return *v
		}
		var ret AccountSku
		return ret
	}).(AccountSkuOutput)
}

// Name of the SKU.
func (o AccountSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// SKU of the trusted signing account.
type AccountSkuResponse struct {
	// Name of the SKU.
	Name string `pulumi:"name"`
}

// SKU of the trusted signing account.
type AccountSkuResponseOutput struct{ *pulumi.OutputState }

func (AccountSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountSkuResponse)(nil)).Elem()
}

func (o AccountSkuResponseOutput) ToAccountSkuResponseOutput() AccountSkuResponseOutput {
	return o
}

func (o AccountSkuResponseOutput) ToAccountSkuResponseOutputWithContext(ctx context.Context) AccountSkuResponseOutput {
	return o
}

// Name of the SKU.
func (o AccountSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AccountSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type AccountSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountSkuResponse)(nil)).Elem()
}

func (o AccountSkuResponsePtrOutput) ToAccountSkuResponsePtrOutput() AccountSkuResponsePtrOutput {
	return o
}

func (o AccountSkuResponsePtrOutput) ToAccountSkuResponsePtrOutputWithContext(ctx context.Context) AccountSkuResponsePtrOutput {
	return o
}

func (o AccountSkuResponsePtrOutput) Elem() AccountSkuResponseOutput {
	return o.ApplyT(func(v *AccountSkuResponse) AccountSkuResponse {
		if v != nil {
			return *v
		}
		var ret AccountSkuResponse
		return ret
	}).(AccountSkuResponseOutput)
}

// Name of the SKU.
func (o AccountSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(AccountSkuOutput{})
	pulumi.RegisterOutputType(AccountSkuPtrOutput{})
	pulumi.RegisterOutputType(AccountSkuResponseOutput{})
	pulumi.RegisterOutputType(AccountSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
