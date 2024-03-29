// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220831preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sku family of the Cloud HSM Cluster
type CloudHsmClusterSkuFamily string

const (
	CloudHsmClusterSkuFamilyB = CloudHsmClusterSkuFamily("B")
)

func (CloudHsmClusterSkuFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudHsmClusterSkuFamily)(nil)).Elem()
}

func (e CloudHsmClusterSkuFamily) ToCloudHsmClusterSkuFamilyOutput() CloudHsmClusterSkuFamilyOutput {
	return pulumi.ToOutput(e).(CloudHsmClusterSkuFamilyOutput)
}

func (e CloudHsmClusterSkuFamily) ToCloudHsmClusterSkuFamilyOutputWithContext(ctx context.Context) CloudHsmClusterSkuFamilyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CloudHsmClusterSkuFamilyOutput)
}

func (e CloudHsmClusterSkuFamily) ToCloudHsmClusterSkuFamilyPtrOutput() CloudHsmClusterSkuFamilyPtrOutput {
	return e.ToCloudHsmClusterSkuFamilyPtrOutputWithContext(context.Background())
}

func (e CloudHsmClusterSkuFamily) ToCloudHsmClusterSkuFamilyPtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuFamilyPtrOutput {
	return CloudHsmClusterSkuFamily(e).ToCloudHsmClusterSkuFamilyOutputWithContext(ctx).ToCloudHsmClusterSkuFamilyPtrOutputWithContext(ctx)
}

func (e CloudHsmClusterSkuFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudHsmClusterSkuFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudHsmClusterSkuFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CloudHsmClusterSkuFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CloudHsmClusterSkuFamilyOutput struct{ *pulumi.OutputState }

func (CloudHsmClusterSkuFamilyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudHsmClusterSkuFamily)(nil)).Elem()
}

func (o CloudHsmClusterSkuFamilyOutput) ToCloudHsmClusterSkuFamilyOutput() CloudHsmClusterSkuFamilyOutput {
	return o
}

func (o CloudHsmClusterSkuFamilyOutput) ToCloudHsmClusterSkuFamilyOutputWithContext(ctx context.Context) CloudHsmClusterSkuFamilyOutput {
	return o
}

func (o CloudHsmClusterSkuFamilyOutput) ToCloudHsmClusterSkuFamilyPtrOutput() CloudHsmClusterSkuFamilyPtrOutput {
	return o.ToCloudHsmClusterSkuFamilyPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuFamilyOutput) ToCloudHsmClusterSkuFamilyPtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuFamilyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudHsmClusterSkuFamily) *CloudHsmClusterSkuFamily {
		return &v
	}).(CloudHsmClusterSkuFamilyPtrOutput)
}

func (o CloudHsmClusterSkuFamilyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuFamilyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudHsmClusterSkuFamily) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CloudHsmClusterSkuFamilyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuFamilyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudHsmClusterSkuFamily) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CloudHsmClusterSkuFamilyPtrOutput struct{ *pulumi.OutputState }

func (CloudHsmClusterSkuFamilyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudHsmClusterSkuFamily)(nil)).Elem()
}

func (o CloudHsmClusterSkuFamilyPtrOutput) ToCloudHsmClusterSkuFamilyPtrOutput() CloudHsmClusterSkuFamilyPtrOutput {
	return o
}

func (o CloudHsmClusterSkuFamilyPtrOutput) ToCloudHsmClusterSkuFamilyPtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuFamilyPtrOutput {
	return o
}

func (o CloudHsmClusterSkuFamilyPtrOutput) Elem() CloudHsmClusterSkuFamilyOutput {
	return o.ApplyT(func(v *CloudHsmClusterSkuFamily) CloudHsmClusterSkuFamily {
		if v != nil {
			return *v
		}
		var ret CloudHsmClusterSkuFamily
		return ret
	}).(CloudHsmClusterSkuFamilyOutput)
}

func (o CloudHsmClusterSkuFamilyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuFamilyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CloudHsmClusterSkuFamily) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CloudHsmClusterSkuFamilyInput is an input type that accepts values of the CloudHsmClusterSkuFamily enum
// A concrete instance of `CloudHsmClusterSkuFamilyInput` can be one of the following:
//
//	CloudHsmClusterSkuFamilyB
type CloudHsmClusterSkuFamilyInput interface {
	pulumi.Input

	ToCloudHsmClusterSkuFamilyOutput() CloudHsmClusterSkuFamilyOutput
	ToCloudHsmClusterSkuFamilyOutputWithContext(context.Context) CloudHsmClusterSkuFamilyOutput
}

var cloudHsmClusterSkuFamilyPtrType = reflect.TypeOf((**CloudHsmClusterSkuFamily)(nil)).Elem()

type CloudHsmClusterSkuFamilyPtrInput interface {
	pulumi.Input

	ToCloudHsmClusterSkuFamilyPtrOutput() CloudHsmClusterSkuFamilyPtrOutput
	ToCloudHsmClusterSkuFamilyPtrOutputWithContext(context.Context) CloudHsmClusterSkuFamilyPtrOutput
}

type cloudHsmClusterSkuFamilyPtr string

func CloudHsmClusterSkuFamilyPtr(v string) CloudHsmClusterSkuFamilyPtrInput {
	return (*cloudHsmClusterSkuFamilyPtr)(&v)
}

func (*cloudHsmClusterSkuFamilyPtr) ElementType() reflect.Type {
	return cloudHsmClusterSkuFamilyPtrType
}

func (in *cloudHsmClusterSkuFamilyPtr) ToCloudHsmClusterSkuFamilyPtrOutput() CloudHsmClusterSkuFamilyPtrOutput {
	return pulumi.ToOutput(in).(CloudHsmClusterSkuFamilyPtrOutput)
}

func (in *cloudHsmClusterSkuFamilyPtr) ToCloudHsmClusterSkuFamilyPtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuFamilyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CloudHsmClusterSkuFamilyPtrOutput)
}

// Sku name of the Cloud HSM Cluster
type CloudHsmClusterSkuName string

const (
	CloudHsmClusterSkuName_Standard_B1  = CloudHsmClusterSkuName("Standard_B1")
	CloudHsmClusterSkuName_Standard_B10 = CloudHsmClusterSkuName("Standard B10")
)

func (CloudHsmClusterSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudHsmClusterSkuName)(nil)).Elem()
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNameOutput() CloudHsmClusterSkuNameOutput {
	return pulumi.ToOutput(e).(CloudHsmClusterSkuNameOutput)
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNameOutputWithContext(ctx context.Context) CloudHsmClusterSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CloudHsmClusterSkuNameOutput)
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return e.ToCloudHsmClusterSkuNamePtrOutputWithContext(context.Background())
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return CloudHsmClusterSkuName(e).ToCloudHsmClusterSkuNameOutputWithContext(ctx).ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx)
}

func (e CloudHsmClusterSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudHsmClusterSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudHsmClusterSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CloudHsmClusterSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CloudHsmClusterSkuNameOutput struct{ *pulumi.OutputState }

func (CloudHsmClusterSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudHsmClusterSkuName)(nil)).Elem()
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNameOutput() CloudHsmClusterSkuNameOutput {
	return o
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNameOutputWithContext(ctx context.Context) CloudHsmClusterSkuNameOutput {
	return o
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return o.ToCloudHsmClusterSkuNamePtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudHsmClusterSkuName) *CloudHsmClusterSkuName {
		return &v
	}).(CloudHsmClusterSkuNamePtrOutput)
}

func (o CloudHsmClusterSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudHsmClusterSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CloudHsmClusterSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudHsmClusterSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CloudHsmClusterSkuNamePtrOutput struct{ *pulumi.OutputState }

func (CloudHsmClusterSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudHsmClusterSkuName)(nil)).Elem()
}

func (o CloudHsmClusterSkuNamePtrOutput) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return o
}

func (o CloudHsmClusterSkuNamePtrOutput) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return o
}

func (o CloudHsmClusterSkuNamePtrOutput) Elem() CloudHsmClusterSkuNameOutput {
	return o.ApplyT(func(v *CloudHsmClusterSkuName) CloudHsmClusterSkuName {
		if v != nil {
			return *v
		}
		var ret CloudHsmClusterSkuName
		return ret
	}).(CloudHsmClusterSkuNameOutput)
}

func (o CloudHsmClusterSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CloudHsmClusterSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CloudHsmClusterSkuNameInput is an input type that accepts values of the CloudHsmClusterSkuName enum
// A concrete instance of `CloudHsmClusterSkuNameInput` can be one of the following:
//
//	CloudHsmClusterSkuName_Standard_B1
//	CloudHsmClusterSkuName_Standard_B10
type CloudHsmClusterSkuNameInput interface {
	pulumi.Input

	ToCloudHsmClusterSkuNameOutput() CloudHsmClusterSkuNameOutput
	ToCloudHsmClusterSkuNameOutputWithContext(context.Context) CloudHsmClusterSkuNameOutput
}

var cloudHsmClusterSkuNamePtrType = reflect.TypeOf((**CloudHsmClusterSkuName)(nil)).Elem()

type CloudHsmClusterSkuNamePtrInput interface {
	pulumi.Input

	ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput
	ToCloudHsmClusterSkuNamePtrOutputWithContext(context.Context) CloudHsmClusterSkuNamePtrOutput
}

type cloudHsmClusterSkuNamePtr string

func CloudHsmClusterSkuNamePtr(v string) CloudHsmClusterSkuNamePtrInput {
	return (*cloudHsmClusterSkuNamePtr)(&v)
}

func (*cloudHsmClusterSkuNamePtr) ElementType() reflect.Type {
	return cloudHsmClusterSkuNamePtrType
}

func (in *cloudHsmClusterSkuNamePtr) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return pulumi.ToOutput(in).(CloudHsmClusterSkuNamePtrOutput)
}

func (in *cloudHsmClusterSkuNamePtr) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CloudHsmClusterSkuNamePtrOutput)
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

// The Cloud HSM Cluster's provisioningState
type ProvisioningState string

const (
	ProvisioningStateProvisioning = ProvisioningState("Provisioning")
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProvisioningStateInput is an input type that accepts values of the ProvisioningState enum
// A concrete instance of `ProvisioningStateInput` can be one of the following:
//
//	ProvisioningStateProvisioning
//	ProvisioningStateSucceeded
//	ProvisioningStateFailed
//	ProvisioningStateDeleting
//	ProvisioningStateCanceled
type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudHsmClusterSkuFamilyOutput{})
	pulumi.RegisterOutputType(CloudHsmClusterSkuFamilyPtrOutput{})
	pulumi.RegisterOutputType(CloudHsmClusterSkuNameOutput{})
	pulumi.RegisterOutputType(CloudHsmClusterSkuNamePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
}
