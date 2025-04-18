// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongocluster

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The mode to create a mongo cluster.
type CreateMode string

const (
	// Create a new mongo cluster.
	CreateModeDefault = CreateMode("Default")
	// Create a mongo cluster from a restore point-in-time.
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	// Create a replica cluster in distinct geographic region from the source cluster.
	CreateModeGeoReplica = CreateMode("GeoReplica")
	// Create a replica cluster in the same geographic region as the source cluster.
	CreateModeReplica = CreateMode("Replica")
)

func (CreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (e CreateMode) ToCreateModeOutput() CreateModeOutput {
	return pulumi.ToOutput(e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModePtrOutput() CreateModePtrOutput {
	return e.ToCreateModePtrOutputWithContext(context.Background())
}

func (e CreateMode) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return CreateMode(e).ToCreateModeOutputWithContext(ctx).ToCreateModePtrOutputWithContext(ctx)
}

func (e CreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CreateModeOutput struct{ *pulumi.OutputState }

func (CreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (o CreateModeOutput) ToCreateModeOutput() CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o.ToCreateModePtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateMode) *CreateMode {
		return &v
	}).(CreateModePtrOutput)
}

func (o CreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CreateModePtrOutput struct{ *pulumi.OutputState }

func (CreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateMode)(nil)).Elem()
}

func (o CreateModePtrOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) Elem() CreateModeOutput {
	return o.ApplyT(func(v *CreateMode) CreateMode {
		if v != nil {
			return *v
		}
		var ret CreateMode
		return ret
	}).(CreateModeOutput)
}

func (o CreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CreateModeInput is an input type that accepts values of the CreateMode enum
// A concrete instance of `CreateModeInput` can be one of the following:
//
//	CreateModeDefault
//	CreateModePointInTimeRestore
//	CreateModeGeoReplica
//	CreateModeReplica
type CreateModeInput interface {
	pulumi.Input

	ToCreateModeOutput() CreateModeOutput
	ToCreateModeOutputWithContext(context.Context) CreateModeOutput
}

var createModePtrType = reflect.TypeOf((**CreateMode)(nil)).Elem()

type CreateModePtrInput interface {
	pulumi.Input

	ToCreateModePtrOutput() CreateModePtrOutput
	ToCreateModePtrOutputWithContext(context.Context) CreateModePtrOutput
}

type createModePtr string

func CreateModePtr(v string) CreateModePtrInput {
	return (*createModePtr)(&v)
}

func (*createModePtr) ElementType() reflect.Type {
	return createModePtrType
}

func (in *createModePtr) ToCreateModePtrOutput() CreateModePtrOutput {
	return pulumi.ToOutput(in).(CreateModePtrOutput)
}

func (in *createModePtr) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CreateModePtrOutput)
}

// The target high availability mode requested for the cluster.
type HighAvailabilityMode string

const (
	// High availability mode is disabled. This mode is can see availability impact during faults or maintenance and is not recommended for production.
	HighAvailabilityModeDisabled = HighAvailabilityMode("Disabled")
	// High availability mode is enabled, where each server in a shard is placed in the same availability zone.
	HighAvailabilityModeSameZone = HighAvailabilityMode("SameZone")
	// High availability mode is enabled and preferences ZoneRedundant if availability zones capacity is available in the region, otherwise falls-back to provisioning with SameZone.
	HighAvailabilityModeZoneRedundantPreferred = HighAvailabilityMode("ZoneRedundantPreferred")
)

func (HighAvailabilityMode) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityMode)(nil)).Elem()
}

func (e HighAvailabilityMode) ToHighAvailabilityModeOutput() HighAvailabilityModeOutput {
	return pulumi.ToOutput(e).(HighAvailabilityModeOutput)
}

func (e HighAvailabilityMode) ToHighAvailabilityModeOutputWithContext(ctx context.Context) HighAvailabilityModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HighAvailabilityModeOutput)
}

func (e HighAvailabilityMode) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return e.ToHighAvailabilityModePtrOutputWithContext(context.Background())
}

func (e HighAvailabilityMode) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return HighAvailabilityMode(e).ToHighAvailabilityModeOutputWithContext(ctx).ToHighAvailabilityModePtrOutputWithContext(ctx)
}

func (e HighAvailabilityMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HighAvailabilityMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HighAvailabilityMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HighAvailabilityMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HighAvailabilityModeOutput struct{ *pulumi.OutputState }

func (HighAvailabilityModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityMode)(nil)).Elem()
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModeOutput() HighAvailabilityModeOutput {
	return o
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModeOutputWithContext(ctx context.Context) HighAvailabilityModeOutput {
	return o
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return o.ToHighAvailabilityModePtrOutputWithContext(context.Background())
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HighAvailabilityMode) *HighAvailabilityMode {
		return &v
	}).(HighAvailabilityModePtrOutput)
}

func (o HighAvailabilityModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HighAvailabilityModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HighAvailabilityMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HighAvailabilityModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HighAvailabilityModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HighAvailabilityMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HighAvailabilityModePtrOutput struct{ *pulumi.OutputState }

func (HighAvailabilityModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailabilityMode)(nil)).Elem()
}

func (o HighAvailabilityModePtrOutput) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return o
}

func (o HighAvailabilityModePtrOutput) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return o
}

func (o HighAvailabilityModePtrOutput) Elem() HighAvailabilityModeOutput {
	return o.ApplyT(func(v *HighAvailabilityMode) HighAvailabilityMode {
		if v != nil {
			return *v
		}
		var ret HighAvailabilityMode
		return ret
	}).(HighAvailabilityModeOutput)
}

func (o HighAvailabilityModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HighAvailabilityModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HighAvailabilityMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// HighAvailabilityModeInput is an input type that accepts values of the HighAvailabilityMode enum
// A concrete instance of `HighAvailabilityModeInput` can be one of the following:
//
//	HighAvailabilityModeDisabled
//	HighAvailabilityModeSameZone
//	HighAvailabilityModeZoneRedundantPreferred
type HighAvailabilityModeInput interface {
	pulumi.Input

	ToHighAvailabilityModeOutput() HighAvailabilityModeOutput
	ToHighAvailabilityModeOutputWithContext(context.Context) HighAvailabilityModeOutput
}

var highAvailabilityModePtrType = reflect.TypeOf((**HighAvailabilityMode)(nil)).Elem()

type HighAvailabilityModePtrInput interface {
	pulumi.Input

	ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput
	ToHighAvailabilityModePtrOutputWithContext(context.Context) HighAvailabilityModePtrOutput
}

type highAvailabilityModePtr string

func HighAvailabilityModePtr(v string) HighAvailabilityModePtrInput {
	return (*highAvailabilityModePtr)(&v)
}

func (*highAvailabilityModePtr) ElementType() reflect.Type {
	return highAvailabilityModePtrType
}

func (in *highAvailabilityModePtr) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return pulumi.ToOutput(in).(HighAvailabilityModePtrOutput)
}

func (in *highAvailabilityModePtr) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HighAvailabilityModePtrOutput)
}

// Preview features that can be enabled on a mongo cluster.
type PreviewFeature string

const (
	// Enables geo replicas preview feature. The feature must be set at create-time on new cluster to enable linking a geo-replica cluster to it.
	PreviewFeatureGeoReplicas = PreviewFeature("GeoReplicas")
)

func (PreviewFeature) ElementType() reflect.Type {
	return reflect.TypeOf((*PreviewFeature)(nil)).Elem()
}

func (e PreviewFeature) ToPreviewFeatureOutput() PreviewFeatureOutput {
	return pulumi.ToOutput(e).(PreviewFeatureOutput)
}

func (e PreviewFeature) ToPreviewFeatureOutputWithContext(ctx context.Context) PreviewFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PreviewFeatureOutput)
}

func (e PreviewFeature) ToPreviewFeaturePtrOutput() PreviewFeaturePtrOutput {
	return e.ToPreviewFeaturePtrOutputWithContext(context.Background())
}

func (e PreviewFeature) ToPreviewFeaturePtrOutputWithContext(ctx context.Context) PreviewFeaturePtrOutput {
	return PreviewFeature(e).ToPreviewFeatureOutputWithContext(ctx).ToPreviewFeaturePtrOutputWithContext(ctx)
}

func (e PreviewFeature) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreviewFeature) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreviewFeature) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PreviewFeature) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PreviewFeatureOutput struct{ *pulumi.OutputState }

func (PreviewFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreviewFeature)(nil)).Elem()
}

func (o PreviewFeatureOutput) ToPreviewFeatureOutput() PreviewFeatureOutput {
	return o
}

func (o PreviewFeatureOutput) ToPreviewFeatureOutputWithContext(ctx context.Context) PreviewFeatureOutput {
	return o
}

func (o PreviewFeatureOutput) ToPreviewFeaturePtrOutput() PreviewFeaturePtrOutput {
	return o.ToPreviewFeaturePtrOutputWithContext(context.Background())
}

func (o PreviewFeatureOutput) ToPreviewFeaturePtrOutputWithContext(ctx context.Context) PreviewFeaturePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PreviewFeature) *PreviewFeature {
		return &v
	}).(PreviewFeaturePtrOutput)
}

func (o PreviewFeatureOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PreviewFeatureOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreviewFeature) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PreviewFeatureOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreviewFeatureOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreviewFeature) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PreviewFeaturePtrOutput struct{ *pulumi.OutputState }

func (PreviewFeaturePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreviewFeature)(nil)).Elem()
}

func (o PreviewFeaturePtrOutput) ToPreviewFeaturePtrOutput() PreviewFeaturePtrOutput {
	return o
}

func (o PreviewFeaturePtrOutput) ToPreviewFeaturePtrOutputWithContext(ctx context.Context) PreviewFeaturePtrOutput {
	return o
}

func (o PreviewFeaturePtrOutput) Elem() PreviewFeatureOutput {
	return o.ApplyT(func(v *PreviewFeature) PreviewFeature {
		if v != nil {
			return *v
		}
		var ret PreviewFeature
		return ret
	}).(PreviewFeatureOutput)
}

func (o PreviewFeaturePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreviewFeaturePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PreviewFeature) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PreviewFeatureInput is an input type that accepts values of the PreviewFeature enum
// A concrete instance of `PreviewFeatureInput` can be one of the following:
//
//	PreviewFeatureGeoReplicas
type PreviewFeatureInput interface {
	pulumi.Input

	ToPreviewFeatureOutput() PreviewFeatureOutput
	ToPreviewFeatureOutputWithContext(context.Context) PreviewFeatureOutput
}

var previewFeaturePtrType = reflect.TypeOf((**PreviewFeature)(nil)).Elem()

type PreviewFeaturePtrInput interface {
	pulumi.Input

	ToPreviewFeaturePtrOutput() PreviewFeaturePtrOutput
	ToPreviewFeaturePtrOutputWithContext(context.Context) PreviewFeaturePtrOutput
}

type previewFeaturePtr string

func PreviewFeaturePtr(v string) PreviewFeaturePtrInput {
	return (*previewFeaturePtr)(&v)
}

func (*previewFeaturePtr) ElementType() reflect.Type {
	return previewFeaturePtrType
}

func (in *previewFeaturePtr) ToPreviewFeaturePtrOutput() PreviewFeaturePtrOutput {
	return pulumi.ToOutput(in).(PreviewFeaturePtrOutput)
}

func (in *previewFeaturePtr) ToPreviewFeaturePtrOutputWithContext(ctx context.Context) PreviewFeaturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PreviewFeaturePtrOutput)
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

// Whether or not public endpoint access is allowed for this mongo cluster.
type PublicNetworkAccess string

const (
	// If set, mongo cluster can be accessed through private and public methods.
	PublicNetworkAccessEnabled = PublicNetworkAccess("Enabled")
	// If set, the private endpoints are the exclusive access method.
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

func init() {
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(HighAvailabilityModeOutput{})
	pulumi.RegisterOutputType(HighAvailabilityModePtrOutput{})
	pulumi.RegisterOutputType(PreviewFeatureOutput{})
	pulumi.RegisterOutputType(PreviewFeaturePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
}
