// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about the SKU of the IoT Central application.
type AppSkuInfo struct {
	// The name of the SKU.
	Name string `pulumi:"name"`
}

// AppSkuInfoInput is an input type that accepts AppSkuInfoArgs and AppSkuInfoOutput values.
// You can construct a concrete instance of `AppSkuInfoInput` via:
//
//	AppSkuInfoArgs{...}
type AppSkuInfoInput interface {
	pulumi.Input

	ToAppSkuInfoOutput() AppSkuInfoOutput
	ToAppSkuInfoOutputWithContext(context.Context) AppSkuInfoOutput
}

// Information about the SKU of the IoT Central application.
type AppSkuInfoArgs struct {
	// The name of the SKU.
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return i.ToAppSkuInfoOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput)
}

// Information about the SKU of the IoT Central application.
type AppSkuInfoOutput struct{ *pulumi.OutputState }

func (AppSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return o
}

// The name of the SKU.
func (o AppSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

// Information about the SKU of the IoT Central application.
type AppSkuInfoResponse struct {
	// The name of the SKU.
	Name string `pulumi:"name"`
}

// Information about the SKU of the IoT Central application.
type AppSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return o
}

// The name of the SKU.
func (o AppSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Managed service identity (either system assigned, or none)
type SystemAssignedServiceIdentity struct {
	// Type of managed service identity (either system assigned, or none).
	Type string `pulumi:"type"`
}

// SystemAssignedServiceIdentityInput is an input type that accepts SystemAssignedServiceIdentityArgs and SystemAssignedServiceIdentityOutput values.
// You can construct a concrete instance of `SystemAssignedServiceIdentityInput` via:
//
//	SystemAssignedServiceIdentityArgs{...}
type SystemAssignedServiceIdentityInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput
	ToSystemAssignedServiceIdentityOutputWithContext(context.Context) SystemAssignedServiceIdentityOutput
}

// Managed service identity (either system assigned, or none)
type SystemAssignedServiceIdentityArgs struct {
	// Type of managed service identity (either system assigned, or none).
	Type pulumi.StringInput `pulumi:"type"`
}

func (SystemAssignedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return i.ToSystemAssignedServiceIdentityOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput)
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput).ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx)
}

// SystemAssignedServiceIdentityPtrInput is an input type that accepts SystemAssignedServiceIdentityArgs, SystemAssignedServiceIdentityPtr and SystemAssignedServiceIdentityPtrOutput values.
// You can construct a concrete instance of `SystemAssignedServiceIdentityPtrInput` via:
//
//	        SystemAssignedServiceIdentityArgs{...}
//
//	or:
//
//	        nil
type SystemAssignedServiceIdentityPtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput
	ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Context) SystemAssignedServiceIdentityPtrOutput
}

type systemAssignedServiceIdentityPtrType SystemAssignedServiceIdentityArgs

func SystemAssignedServiceIdentityPtr(v *SystemAssignedServiceIdentityArgs) SystemAssignedServiceIdentityPtrInput {
	return (*systemAssignedServiceIdentityPtrType)(v)
}

func (*systemAssignedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityPtrOutput)
}

// Managed service identity (either system assigned, or none)
type SystemAssignedServiceIdentityOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentity) *SystemAssignedServiceIdentity {
		return &v
	}).(SystemAssignedServiceIdentityPtrOutput)
}

// Type of managed service identity (either system assigned, or none).
func (o SystemAssignedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) Elem() SystemAssignedServiceIdentityOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) SystemAssignedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentity
		return ret
	}).(SystemAssignedServiceIdentityOutput)
}

// Type of managed service identity (either system assigned, or none).
func (o SystemAssignedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// Managed service identity (either system assigned, or none)
type SystemAssignedServiceIdentityResponse struct {
	// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantId string `pulumi:"tenantId"`
	// Type of managed service identity (either system assigned, or none).
	Type string `pulumi:"type"`
}

// Managed service identity (either system assigned, or none)
type SystemAssignedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutput() SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponseOutput {
	return o
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o SystemAssignedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o SystemAssignedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// Type of managed service identity (either system assigned, or none).
func (o SystemAssignedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Elem() SystemAssignedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) SystemAssignedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityResponse
		return ret
	}).(SystemAssignedServiceIdentityResponseOutput)
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o SystemAssignedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o SystemAssignedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Type of managed service identity (either system assigned, or none).
func (o SystemAssignedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSkuInfoOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponsePtrOutput{})
}