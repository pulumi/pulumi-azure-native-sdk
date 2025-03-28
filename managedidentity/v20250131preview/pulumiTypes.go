// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250131preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Object for defining the allowed identifiers of external identities. Either 'subject' or 'claimsMatchingExpression' must be defined, but not both.
type FederatedIdentityCredentialPropertiesClaimsMatchingExpression struct {
	// Specifies the version of the flexible fic language used in the expression.
	LanguageVersion int `pulumi:"languageVersion"`
	// Wildcard-based expression for matching incoming subject claims.
	Value string `pulumi:"value"`
}

// FederatedIdentityCredentialPropertiesClaimsMatchingExpressionInput is an input type that accepts FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs and FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput values.
// You can construct a concrete instance of `FederatedIdentityCredentialPropertiesClaimsMatchingExpressionInput` via:
//
//	FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs{...}
type FederatedIdentityCredentialPropertiesClaimsMatchingExpressionInput interface {
	pulumi.Input

	ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput
	ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutputWithContext(context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput
}

// Object for defining the allowed identifiers of external identities. Either 'subject' or 'claimsMatchingExpression' must be defined, but not both.
type FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs struct {
	// Specifies the version of the flexible fic language used in the expression.
	LanguageVersion pulumi.IntInput `pulumi:"languageVersion"`
	// Wildcard-based expression for matching incoming subject claims.
	Value pulumi.StringInput `pulumi:"value"`
}

func (FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedIdentityCredentialPropertiesClaimsMatchingExpression)(nil)).Elem()
}

func (i FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput {
	return i.ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutputWithContext(context.Background())
}

func (i FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput)
}

func (i FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return i.ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(context.Background())
}

func (i FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput).ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(ctx)
}

// FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrInput is an input type that accepts FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs, FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtr and FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput values.
// You can construct a concrete instance of `FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrInput` via:
//
//	        FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs{...}
//
//	or:
//
//	        nil
type FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrInput interface {
	pulumi.Input

	ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput
	ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput
}

type federatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrType FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs

func FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtr(v *FederatedIdentityCredentialPropertiesClaimsMatchingExpressionArgs) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrInput {
	return (*federatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrType)(v)
}

func (*federatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedIdentityCredentialPropertiesClaimsMatchingExpression)(nil)).Elem()
}

func (i *federatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrType) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return i.ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(context.Background())
}

func (i *federatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrType) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput)
}

// Object for defining the allowed identifiers of external identities. Either 'subject' or 'claimsMatchingExpression' must be defined, but not both.
type FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput struct{ *pulumi.OutputState }

func (FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedIdentityCredentialPropertiesClaimsMatchingExpression)(nil)).Elem()
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return o.ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(context.Background())
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FederatedIdentityCredentialPropertiesClaimsMatchingExpression) *FederatedIdentityCredentialPropertiesClaimsMatchingExpression {
		return &v
	}).(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput)
}

// Specifies the version of the flexible fic language used in the expression.
func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) LanguageVersion() pulumi.IntOutput {
	return o.ApplyT(func(v FederatedIdentityCredentialPropertiesClaimsMatchingExpression) int { return v.LanguageVersion }).(pulumi.IntOutput)
}

// Wildcard-based expression for matching incoming subject claims.
func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v FederatedIdentityCredentialPropertiesClaimsMatchingExpression) string { return v.Value }).(pulumi.StringOutput)
}

type FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput struct{ *pulumi.OutputState }

func (FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedIdentityCredentialPropertiesClaimsMatchingExpression)(nil)).Elem()
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput) ToFederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput) Elem() FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput {
	return o.ApplyT(func(v *FederatedIdentityCredentialPropertiesClaimsMatchingExpression) FederatedIdentityCredentialPropertiesClaimsMatchingExpression {
		if v != nil {
			return *v
		}
		var ret FederatedIdentityCredentialPropertiesClaimsMatchingExpression
		return ret
	}).(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput)
}

// Specifies the version of the flexible fic language used in the expression.
func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput) LanguageVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FederatedIdentityCredentialPropertiesClaimsMatchingExpression) *int {
		if v == nil {
			return nil
		}
		return &v.LanguageVersion
	}).(pulumi.IntPtrOutput)
}

// Wildcard-based expression for matching incoming subject claims.
func (o FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedIdentityCredentialPropertiesClaimsMatchingExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

// Object for defining the allowed identifiers of external identities. Either 'subject' or 'claimsMatchingExpression' must be defined, but not both.
type FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression struct {
	// Specifies the version of the flexible fic language used in the expression.
	LanguageVersion int `pulumi:"languageVersion"`
	// Wildcard-based expression for matching incoming subject claims.
	Value string `pulumi:"value"`
}

// Object for defining the allowed identifiers of external identities. Either 'subject' or 'claimsMatchingExpression' must be defined, but not both.
type FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput struct{ *pulumi.OutputState }

func (FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression)(nil)).Elem()
}

func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput) ToFederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput() FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput) ToFederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput {
	return o
}

// Specifies the version of the flexible fic language used in the expression.
func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput) LanguageVersion() pulumi.IntOutput {
	return o.ApplyT(func(v FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression) int {
		return v.LanguageVersion
	}).(pulumi.IntOutput)
}

// Wildcard-based expression for matching incoming subject claims.
func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression) string { return v.Value }).(pulumi.StringOutput)
}

type FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput struct{ *pulumi.OutputState }

func (FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression)(nil)).Elem()
}

func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput) ToFederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput() FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput) ToFederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutputWithContext(ctx context.Context) FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput {
	return o
}

func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput) Elem() FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput {
	return o.ApplyT(func(v *FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression) FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression {
		if v != nil {
			return *v
		}
		var ret FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression
		return ret
	}).(FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput)
}

// Specifies the version of the flexible fic language used in the expression.
func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput) LanguageVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression) *int {
		if v == nil {
			return nil
		}
		return &v.LanguageVersion
	}).(pulumi.IntPtrOutput)
}

// Wildcard-based expression for matching incoming subject claims.
func (o FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Value
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
	pulumi.RegisterOutputType(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionOutput{})
	pulumi.RegisterOutputType(FederatedIdentityCredentialPropertiesClaimsMatchingExpressionPtrOutput{})
	pulumi.RegisterOutputType(FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionOutput{})
	pulumi.RegisterOutputType(FederatedIdentityCredentialPropertiesResponseClaimsMatchingExpressionPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
