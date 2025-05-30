// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package education

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The amount.
type Amount struct {
	// The type of currency being used for the value.
	Currency *string `pulumi:"currency"`
	// Amount value.
	Value *float64 `pulumi:"value"`
}

// AmountInput is an input type that accepts AmountArgs and AmountOutput values.
// You can construct a concrete instance of `AmountInput` via:
//
//	AmountArgs{...}
type AmountInput interface {
	pulumi.Input

	ToAmountOutput() AmountOutput
	ToAmountOutputWithContext(context.Context) AmountOutput
}

// The amount.
type AmountArgs struct {
	// The type of currency being used for the value.
	Currency pulumi.StringPtrInput `pulumi:"currency"`
	// Amount value.
	Value pulumi.Float64PtrInput `pulumi:"value"`
}

func (AmountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Amount)(nil)).Elem()
}

func (i AmountArgs) ToAmountOutput() AmountOutput {
	return i.ToAmountOutputWithContext(context.Background())
}

func (i AmountArgs) ToAmountOutputWithContext(ctx context.Context) AmountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmountOutput)
}

// The amount.
type AmountOutput struct{ *pulumi.OutputState }

func (AmountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Amount)(nil)).Elem()
}

func (o AmountOutput) ToAmountOutput() AmountOutput {
	return o
}

func (o AmountOutput) ToAmountOutputWithContext(ctx context.Context) AmountOutput {
	return o
}

// The type of currency being used for the value.
func (o AmountOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Amount) *string { return v.Currency }).(pulumi.StringPtrOutput)
}

// Amount value.
func (o AmountOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Amount) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

// The amount.
type AmountResponse struct {
	// The type of currency being used for the value.
	Currency *string `pulumi:"currency"`
	// Amount value.
	Value *float64 `pulumi:"value"`
}

// The amount.
type AmountResponseOutput struct{ *pulumi.OutputState }

func (AmountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmountResponse)(nil)).Elem()
}

func (o AmountResponseOutput) ToAmountResponseOutput() AmountResponseOutput {
	return o
}

func (o AmountResponseOutput) ToAmountResponseOutputWithContext(ctx context.Context) AmountResponseOutput {
	return o
}

// The type of currency being used for the value.
func (o AmountResponseOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmountResponse) *string { return v.Currency }).(pulumi.StringPtrOutput)
}

// Amount value.
func (o AmountResponseOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AmountResponse) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
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
	pulumi.RegisterOutputType(AmountOutput{})
	pulumi.RegisterOutputType(AmountResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
