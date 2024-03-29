// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The network access type for accessing Log Analytics query.
type PublicNetworkAccessType string

const (
	// Enables connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Disables public connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func (PublicNetworkAccessType) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessType)(nil)).Elem()
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessTypeOutput)
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypeOutputWithContext(ctx context.Context) PublicNetworkAccessTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessTypeOutput)
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return e.ToPublicNetworkAccessTypePtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return PublicNetworkAccessType(e).ToPublicNetworkAccessTypeOutputWithContext(ctx).ToPublicNetworkAccessTypePtrOutputWithContext(ctx)
}

func (e PublicNetworkAccessType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessTypeOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessType)(nil)).Elem()
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput {
	return o
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypeOutputWithContext(ctx context.Context) PublicNetworkAccessTypeOutput {
	return o
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return o.ToPublicNetworkAccessTypePtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccessType) *PublicNetworkAccessType {
		return &v
	}).(PublicNetworkAccessTypePtrOutput)
}

func (o PublicNetworkAccessTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessTypePtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccessType)(nil)).Elem()
}

func (o PublicNetworkAccessTypePtrOutput) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return o
}

func (o PublicNetworkAccessTypePtrOutput) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return o
}

func (o PublicNetworkAccessTypePtrOutput) Elem() PublicNetworkAccessTypeOutput {
	return o.ApplyT(func(v *PublicNetworkAccessType) PublicNetworkAccessType {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccessType
		return ret
	}).(PublicNetworkAccessTypeOutput)
}

func (o PublicNetworkAccessTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccessType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PublicNetworkAccessTypeInput is an input type that accepts values of the PublicNetworkAccessType enum
// A concrete instance of `PublicNetworkAccessTypeInput` can be one of the following:
//
//	PublicNetworkAccessTypeEnabled
//	PublicNetworkAccessTypeDisabled
type PublicNetworkAccessTypeInput interface {
	pulumi.Input

	ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput
	ToPublicNetworkAccessTypeOutputWithContext(context.Context) PublicNetworkAccessTypeOutput
}

var publicNetworkAccessTypePtrType = reflect.TypeOf((**PublicNetworkAccessType)(nil)).Elem()

type PublicNetworkAccessTypePtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput
	ToPublicNetworkAccessTypePtrOutputWithContext(context.Context) PublicNetworkAccessTypePtrOutput
}

type publicNetworkAccessTypePtr string

func PublicNetworkAccessTypePtr(v string) PublicNetworkAccessTypePtrInput {
	return (*publicNetworkAccessTypePtr)(&v)
}

func (*publicNetworkAccessTypePtr) ElementType() reflect.Type {
	return publicNetworkAccessTypePtrType
}

func (in *publicNetworkAccessTypePtr) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessTypePtrOutput)
}

func (in *publicNetworkAccessTypePtr) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessTypePtrOutput)
}

// The provisioning state of the workspace.
type WorkspaceEntityStatus string

const (
	WorkspaceEntityStatusCreating            = WorkspaceEntityStatus("Creating")
	WorkspaceEntityStatusSucceeded           = WorkspaceEntityStatus("Succeeded")
	WorkspaceEntityStatusFailed              = WorkspaceEntityStatus("Failed")
	WorkspaceEntityStatusCanceled            = WorkspaceEntityStatus("Canceled")
	WorkspaceEntityStatusDeleting            = WorkspaceEntityStatus("Deleting")
	WorkspaceEntityStatusProvisioningAccount = WorkspaceEntityStatus("ProvisioningAccount")
	WorkspaceEntityStatusUpdating            = WorkspaceEntityStatus("Updating")
)

func (WorkspaceEntityStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEntityStatus)(nil)).Elem()
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusOutput() WorkspaceEntityStatusOutput {
	return pulumi.ToOutput(e).(WorkspaceEntityStatusOutput)
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusOutputWithContext(ctx context.Context) WorkspaceEntityStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkspaceEntityStatusOutput)
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return e.ToWorkspaceEntityStatusPtrOutputWithContext(context.Background())
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return WorkspaceEntityStatus(e).ToWorkspaceEntityStatusOutputWithContext(ctx).ToWorkspaceEntityStatusPtrOutputWithContext(ctx)
}

func (e WorkspaceEntityStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceEntityStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceEntityStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkspaceEntityStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkspaceEntityStatusOutput struct{ *pulumi.OutputState }

func (WorkspaceEntityStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEntityStatus)(nil)).Elem()
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusOutput() WorkspaceEntityStatusOutput {
	return o
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusOutputWithContext(ctx context.Context) WorkspaceEntityStatusOutput {
	return o
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return o.ToWorkspaceEntityStatusPtrOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceEntityStatus) *WorkspaceEntityStatus {
		return &v
	}).(WorkspaceEntityStatusPtrOutput)
}

func (o WorkspaceEntityStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceEntityStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkspaceEntityStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceEntityStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspaceEntityStatusPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceEntityStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceEntityStatus)(nil)).Elem()
}

func (o WorkspaceEntityStatusPtrOutput) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return o
}

func (o WorkspaceEntityStatusPtrOutput) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return o
}

func (o WorkspaceEntityStatusPtrOutput) Elem() WorkspaceEntityStatusOutput {
	return o.ApplyT(func(v *WorkspaceEntityStatus) WorkspaceEntityStatus {
		if v != nil {
			return *v
		}
		var ret WorkspaceEntityStatus
		return ret
	}).(WorkspaceEntityStatusOutput)
}

func (o WorkspaceEntityStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkspaceEntityStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkspaceEntityStatusInput is an input type that accepts values of the WorkspaceEntityStatus enum
// A concrete instance of `WorkspaceEntityStatusInput` can be one of the following:
//
//	WorkspaceEntityStatusCreating
//	WorkspaceEntityStatusSucceeded
//	WorkspaceEntityStatusFailed
//	WorkspaceEntityStatusCanceled
//	WorkspaceEntityStatusDeleting
//	WorkspaceEntityStatusProvisioningAccount
//	WorkspaceEntityStatusUpdating
type WorkspaceEntityStatusInput interface {
	pulumi.Input

	ToWorkspaceEntityStatusOutput() WorkspaceEntityStatusOutput
	ToWorkspaceEntityStatusOutputWithContext(context.Context) WorkspaceEntityStatusOutput
}

var workspaceEntityStatusPtrType = reflect.TypeOf((**WorkspaceEntityStatus)(nil)).Elem()

type WorkspaceEntityStatusPtrInput interface {
	pulumi.Input

	ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput
	ToWorkspaceEntityStatusPtrOutputWithContext(context.Context) WorkspaceEntityStatusPtrOutput
}

type workspaceEntityStatusPtr string

func WorkspaceEntityStatusPtr(v string) WorkspaceEntityStatusPtrInput {
	return (*workspaceEntityStatusPtr)(&v)
}

func (*workspaceEntityStatusPtr) ElementType() reflect.Type {
	return workspaceEntityStatusPtrType
}

func (in *workspaceEntityStatusPtr) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return pulumi.ToOutput(in).(WorkspaceEntityStatusPtrOutput)
}

func (in *workspaceEntityStatusPtr) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkspaceEntityStatusPtrOutput)
}

// The name of the SKU.
type WorkspaceSkuNameEnum string

const (
	WorkspaceSkuNameEnumFree                = WorkspaceSkuNameEnum("Free")
	WorkspaceSkuNameEnumStandard            = WorkspaceSkuNameEnum("Standard")
	WorkspaceSkuNameEnumPremium             = WorkspaceSkuNameEnum("Premium")
	WorkspaceSkuNameEnumPerNode             = WorkspaceSkuNameEnum("PerNode")
	WorkspaceSkuNameEnumPerGB2018           = WorkspaceSkuNameEnum("PerGB2018")
	WorkspaceSkuNameEnumStandalone          = WorkspaceSkuNameEnum("Standalone")
	WorkspaceSkuNameEnumCapacityReservation = WorkspaceSkuNameEnum("CapacityReservation")
	WorkspaceSkuNameEnumLACluster           = WorkspaceSkuNameEnum("LACluster")
)

func (WorkspaceSkuNameEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuNameEnum)(nil)).Elem()
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumOutput() WorkspaceSkuNameEnumOutput {
	return pulumi.ToOutput(e).(WorkspaceSkuNameEnumOutput)
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkspaceSkuNameEnumOutput)
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return e.ToWorkspaceSkuNameEnumPtrOutputWithContext(context.Background())
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return WorkspaceSkuNameEnum(e).ToWorkspaceSkuNameEnumOutputWithContext(ctx).ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx)
}

func (e WorkspaceSkuNameEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceSkuNameEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceSkuNameEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkspaceSkuNameEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkspaceSkuNameEnumOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuNameEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuNameEnum)(nil)).Elem()
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumOutput() WorkspaceSkuNameEnumOutput {
	return o
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumOutput {
	return o
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return o.ToWorkspaceSkuNameEnumPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSkuNameEnum) *WorkspaceSkuNameEnum {
		return &v
	}).(WorkspaceSkuNameEnumPtrOutput)
}

func (o WorkspaceSkuNameEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceSkuNameEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkspaceSkuNameEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceSkuNameEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspaceSkuNameEnumPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuNameEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSkuNameEnum)(nil)).Elem()
}

func (o WorkspaceSkuNameEnumPtrOutput) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return o
}

func (o WorkspaceSkuNameEnumPtrOutput) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return o
}

func (o WorkspaceSkuNameEnumPtrOutput) Elem() WorkspaceSkuNameEnumOutput {
	return o.ApplyT(func(v *WorkspaceSkuNameEnum) WorkspaceSkuNameEnum {
		if v != nil {
			return *v
		}
		var ret WorkspaceSkuNameEnum
		return ret
	}).(WorkspaceSkuNameEnumOutput)
}

func (o WorkspaceSkuNameEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkspaceSkuNameEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkspaceSkuNameEnumInput is an input type that accepts values of the WorkspaceSkuNameEnum enum
// A concrete instance of `WorkspaceSkuNameEnumInput` can be one of the following:
//
//	WorkspaceSkuNameEnumFree
//	WorkspaceSkuNameEnumStandard
//	WorkspaceSkuNameEnumPremium
//	WorkspaceSkuNameEnumPerNode
//	WorkspaceSkuNameEnumPerGB2018
//	WorkspaceSkuNameEnumStandalone
//	WorkspaceSkuNameEnumCapacityReservation
//	WorkspaceSkuNameEnumLACluster
type WorkspaceSkuNameEnumInput interface {
	pulumi.Input

	ToWorkspaceSkuNameEnumOutput() WorkspaceSkuNameEnumOutput
	ToWorkspaceSkuNameEnumOutputWithContext(context.Context) WorkspaceSkuNameEnumOutput
}

var workspaceSkuNameEnumPtrType = reflect.TypeOf((**WorkspaceSkuNameEnum)(nil)).Elem()

type WorkspaceSkuNameEnumPtrInput interface {
	pulumi.Input

	ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput
	ToWorkspaceSkuNameEnumPtrOutputWithContext(context.Context) WorkspaceSkuNameEnumPtrOutput
}

type workspaceSkuNameEnumPtr string

func WorkspaceSkuNameEnumPtr(v string) WorkspaceSkuNameEnumPtrInput {
	return (*workspaceSkuNameEnumPtr)(&v)
}

func (*workspaceSkuNameEnumPtr) ElementType() reflect.Type {
	return workspaceSkuNameEnumPtrType
}

func (in *workspaceSkuNameEnumPtr) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return pulumi.ToOutput(in).(WorkspaceSkuNameEnumPtrOutput)
}

func (in *workspaceSkuNameEnumPtr) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkspaceSkuNameEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicNetworkAccessTypeOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessTypePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceEntityStatusOutput{})
	pulumi.RegisterOutputType(WorkspaceEntityStatusPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuNameEnumOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuNameEnumPtrOutput{})
}
