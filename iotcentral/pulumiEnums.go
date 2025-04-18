// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotcentral

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The name of the SKU.
type AppSku string

const (
	AppSkuST0 = AppSku("ST0")
	AppSkuST1 = AppSku("ST1")
	AppSkuST2 = AppSku("ST2")
)

func (AppSku) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSku)(nil)).Elem()
}

func (e AppSku) ToAppSkuOutput() AppSkuOutput {
	return pulumi.ToOutput(e).(AppSkuOutput)
}

func (e AppSku) ToAppSkuOutputWithContext(ctx context.Context) AppSkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppSkuOutput)
}

func (e AppSku) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return e.ToAppSkuPtrOutputWithContext(context.Background())
}

func (e AppSku) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return AppSku(e).ToAppSkuOutputWithContext(ctx).ToAppSkuPtrOutputWithContext(ctx)
}

func (e AppSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppSkuOutput struct{ *pulumi.OutputState }

func (AppSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSku)(nil)).Elem()
}

func (o AppSkuOutput) ToAppSkuOutput() AppSkuOutput {
	return o
}

func (o AppSkuOutput) ToAppSkuOutputWithContext(ctx context.Context) AppSkuOutput {
	return o
}

func (o AppSkuOutput) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return o.ToAppSkuPtrOutputWithContext(context.Background())
}

func (o AppSkuOutput) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSku) *AppSku {
		return &v
	}).(AppSkuPtrOutput)
}

func (o AppSkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppSkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppSku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppSkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppSkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppSku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppSkuPtrOutput struct{ *pulumi.OutputState }

func (AppSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSku)(nil)).Elem()
}

func (o AppSkuPtrOutput) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return o
}

func (o AppSkuPtrOutput) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return o
}

func (o AppSkuPtrOutput) Elem() AppSkuOutput {
	return o.ApplyT(func(v *AppSku) AppSku {
		if v != nil {
			return *v
		}
		var ret AppSku
		return ret
	}).(AppSkuOutput)
}

func (o AppSkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppSkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppSku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AppSkuInput is an input type that accepts values of the AppSku enum
// A concrete instance of `AppSkuInput` can be one of the following:
//
//	AppSkuST0
//	AppSkuST1
//	AppSkuST2
type AppSkuInput interface {
	pulumi.Input

	ToAppSkuOutput() AppSkuOutput
	ToAppSkuOutputWithContext(context.Context) AppSkuOutput
}

var appSkuPtrType = reflect.TypeOf((**AppSku)(nil)).Elem()

type AppSkuPtrInput interface {
	pulumi.Input

	ToAppSkuPtrOutput() AppSkuPtrOutput
	ToAppSkuPtrOutputWithContext(context.Context) AppSkuPtrOutput
}

type appSkuPtr string

func AppSkuPtr(v string) AppSkuPtrInput {
	return (*appSkuPtr)(&v)
}

func (*appSkuPtr) ElementType() reflect.Type {
	return appSkuPtrType
}

func (in *appSkuPtr) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return pulumi.ToOutput(in).(AppSkuPtrOutput)
}

func (in *appSkuPtr) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppSkuPtrOutput)
}

// The default network action to apply.
type NetworkAction string

const (
	NetworkActionAllow = NetworkAction("Allow")
	NetworkActionDeny  = NetworkAction("Deny")
)

func (NetworkAction) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAction)(nil)).Elem()
}

func (e NetworkAction) ToNetworkActionOutput() NetworkActionOutput {
	return pulumi.ToOutput(e).(NetworkActionOutput)
}

func (e NetworkAction) ToNetworkActionOutputWithContext(ctx context.Context) NetworkActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkActionOutput)
}

func (e NetworkAction) ToNetworkActionPtrOutput() NetworkActionPtrOutput {
	return e.ToNetworkActionPtrOutputWithContext(context.Background())
}

func (e NetworkAction) ToNetworkActionPtrOutputWithContext(ctx context.Context) NetworkActionPtrOutput {
	return NetworkAction(e).ToNetworkActionOutputWithContext(ctx).ToNetworkActionPtrOutputWithContext(ctx)
}

func (e NetworkAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkActionOutput struct{ *pulumi.OutputState }

func (NetworkActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAction)(nil)).Elem()
}

func (o NetworkActionOutput) ToNetworkActionOutput() NetworkActionOutput {
	return o
}

func (o NetworkActionOutput) ToNetworkActionOutputWithContext(ctx context.Context) NetworkActionOutput {
	return o
}

func (o NetworkActionOutput) ToNetworkActionPtrOutput() NetworkActionPtrOutput {
	return o.ToNetworkActionPtrOutputWithContext(context.Background())
}

func (o NetworkActionOutput) ToNetworkActionPtrOutputWithContext(ctx context.Context) NetworkActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkAction) *NetworkAction {
		return &v
	}).(NetworkActionPtrOutput)
}

func (o NetworkActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkActionPtrOutput struct{ *pulumi.OutputState }

func (NetworkActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAction)(nil)).Elem()
}

func (o NetworkActionPtrOutput) ToNetworkActionPtrOutput() NetworkActionPtrOutput {
	return o
}

func (o NetworkActionPtrOutput) ToNetworkActionPtrOutputWithContext(ctx context.Context) NetworkActionPtrOutput {
	return o
}

func (o NetworkActionPtrOutput) Elem() NetworkActionOutput {
	return o.ApplyT(func(v *NetworkAction) NetworkAction {
		if v != nil {
			return *v
		}
		var ret NetworkAction
		return ret
	}).(NetworkActionOutput)
}

func (o NetworkActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkActionInput is an input type that accepts values of the NetworkAction enum
// A concrete instance of `NetworkActionInput` can be one of the following:
//
//	NetworkActionAllow
//	NetworkActionDeny
type NetworkActionInput interface {
	pulumi.Input

	ToNetworkActionOutput() NetworkActionOutput
	ToNetworkActionOutputWithContext(context.Context) NetworkActionOutput
}

var networkActionPtrType = reflect.TypeOf((**NetworkAction)(nil)).Elem()

type NetworkActionPtrInput interface {
	pulumi.Input

	ToNetworkActionPtrOutput() NetworkActionPtrOutput
	ToNetworkActionPtrOutputWithContext(context.Context) NetworkActionPtrOutput
}

type networkActionPtr string

func NetworkActionPtr(v string) NetworkActionPtrInput {
	return (*networkActionPtr)(&v)
}

func (*networkActionPtr) ElementType() reflect.Type {
	return networkActionPtrType
}

func (in *networkActionPtr) ToNetworkActionPtrOutput() NetworkActionPtrOutput {
	return pulumi.ToOutput(in).(NetworkActionPtrOutput)
}

func (in *networkActionPtr) ToNetworkActionPtrOutputWithContext(ctx context.Context) NetworkActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkActionPtrOutput)
}

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PrivateEndpointServiceConnectionStatusInput is an input type that accepts values of the PrivateEndpointServiceConnectionStatus enum
// A concrete instance of `PrivateEndpointServiceConnectionStatusInput` can be one of the following:
//
//	PrivateEndpointServiceConnectionStatusPending
//	PrivateEndpointServiceConnectionStatusApproved
//	PrivateEndpointServiceConnectionStatusRejected
type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

// Whether requests from the public network are allowed.
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PublicNetworkAccessInput is an input type that accepts values of the PublicNetworkAccess enum
// A concrete instance of `PublicNetworkAccessInput` can be one of the following:
//
//	PublicNetworkAccessEnabled
//	PublicNetworkAccessDisabled
type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
}

// Type of managed service identity (either system assigned, or none).
type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           = SystemAssignedServiceIdentityType("None")
	SystemAssignedServiceIdentityTypeSystemAssigned = SystemAssignedServiceIdentityType("SystemAssigned")
)

func (SystemAssignedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(SystemAssignedServiceIdentityTypeOutput)
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemAssignedServiceIdentityTypeOutput)
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return e.ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return SystemAssignedServiceIdentityType(e).ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx).ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e SystemAssignedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemAssignedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemAssignedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemAssignedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemAssignedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypeOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return o.ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentityType) *SystemAssignedServiceIdentityType {
		return &v
	}).(SystemAssignedServiceIdentityTypePtrOutput)
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemAssignedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemAssignedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypePtrOutput) Elem() SystemAssignedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityType) SystemAssignedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityType
		return ret
	}).(SystemAssignedServiceIdentityTypeOutput)
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemAssignedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SystemAssignedServiceIdentityTypeInput is an input type that accepts values of the SystemAssignedServiceIdentityType enum
// A concrete instance of `SystemAssignedServiceIdentityTypeInput` can be one of the following:
//
//	SystemAssignedServiceIdentityTypeNone
//	SystemAssignedServiceIdentityTypeSystemAssigned
type SystemAssignedServiceIdentityTypeInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput
	ToSystemAssignedServiceIdentityTypeOutputWithContext(context.Context) SystemAssignedServiceIdentityTypeOutput
}

var systemAssignedServiceIdentityTypePtrType = reflect.TypeOf((**SystemAssignedServiceIdentityType)(nil)).Elem()

type SystemAssignedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput
	ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Context) SystemAssignedServiceIdentityTypePtrOutput
}

type systemAssignedServiceIdentityTypePtr string

func SystemAssignedServiceIdentityTypePtr(v string) SystemAssignedServiceIdentityTypePtrInput {
	return (*systemAssignedServiceIdentityTypePtr)(&v)
}

func (*systemAssignedServiceIdentityTypePtr) ElementType() reflect.Type {
	return systemAssignedServiceIdentityTypePtrType
}

func (in *systemAssignedServiceIdentityTypePtr) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(SystemAssignedServiceIdentityTypePtrOutput)
}

func (in *systemAssignedServiceIdentityTypePtr) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemAssignedServiceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSkuOutput{})
	pulumi.RegisterOutputType(AppSkuPtrOutput{})
	pulumi.RegisterOutputType(NetworkActionOutput{})
	pulumi.RegisterOutputType(NetworkActionPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityTypePtrOutput{})
}
