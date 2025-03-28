// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240821preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of the extended location.
type ExtendedLocationType string

const (
	// Azure Edge Zones location type
	ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")
	// Azure Custom Locations type
	ExtendedLocationTypeCustomLocation = ExtendedLocationType("CustomLocation")
)

func (ExtendedLocationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationType)(nil)).Elem()
}

func (e ExtendedLocationType) ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypeOutput)
}

func (e ExtendedLocationType) ToExtendedLocationTypeOutputWithContext(ctx context.Context) ExtendedLocationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypeOutput)
}

func (e ExtendedLocationType) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return e.ToExtendedLocationTypePtrOutputWithContext(context.Background())
}

func (e ExtendedLocationType) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return ExtendedLocationType(e).ToExtendedLocationTypeOutputWithContext(ctx).ToExtendedLocationTypePtrOutputWithContext(ctx)
}

func (e ExtendedLocationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypeOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationType)(nil)).Elem()
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput {
	return o
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypeOutputWithContext(ctx context.Context) ExtendedLocationTypeOutput {
	return o
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return o.ToExtendedLocationTypePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationType) *ExtendedLocationType {
		return &v
	}).(ExtendedLocationTypePtrOutput)
}

func (o ExtendedLocationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationType)(nil)).Elem()
}

func (o ExtendedLocationTypePtrOutput) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return o
}

func (o ExtendedLocationTypePtrOutput) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return o
}

func (o ExtendedLocationTypePtrOutput) Elem() ExtendedLocationTypeOutput {
	return o.ApplyT(func(v *ExtendedLocationType) ExtendedLocationType {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationType
		return ret
	}).(ExtendedLocationTypeOutput)
}

func (o ExtendedLocationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ExtendedLocationTypeInput is an input type that accepts values of the ExtendedLocationType enum
// A concrete instance of `ExtendedLocationTypeInput` can be one of the following:
//
//	ExtendedLocationTypeEdgeZone
//	ExtendedLocationTypeCustomLocation
type ExtendedLocationTypeInput interface {
	pulumi.Input

	ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput
	ToExtendedLocationTypeOutputWithContext(context.Context) ExtendedLocationTypeOutput
}

var extendedLocationTypePtrType = reflect.TypeOf((**ExtendedLocationType)(nil)).Elem()

type ExtendedLocationTypePtrInput interface {
	pulumi.Input

	ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput
	ToExtendedLocationTypePtrOutputWithContext(context.Context) ExtendedLocationTypePtrOutput
}

type extendedLocationTypePtr string

func ExtendedLocationTypePtr(v string) ExtendedLocationTypePtrInput {
	return (*extendedLocationTypePtr)(&v)
}

func (*extendedLocationTypePtr) ElementType() reflect.Type {
	return extendedLocationTypePtrType
}

func (in *extendedLocationTypePtr) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypePtrOutput)
}

func (in *extendedLocationTypePtr) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypePtrOutput)
}

// Type specifies the type of the Kubernetes secret object, e.g. "Opaque" or"kubernetes.io/tls". The controller must have permission to create secrets of the specified type.
type KubernetesSecretType string

const (
	// Opaque is the default secret type.
	KubernetesSecretTypeOpaque = KubernetesSecretType("Opaque")
	// The kubernetes.io/tls secret type is for storing a certificate and its associated key that are typically used for TLS.
	KubernetesSecretTypeTls = KubernetesSecretType("kubernetes.io/tls")
)

func (KubernetesSecretType) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesSecretType)(nil)).Elem()
}

func (e KubernetesSecretType) ToKubernetesSecretTypeOutput() KubernetesSecretTypeOutput {
	return pulumi.ToOutput(e).(KubernetesSecretTypeOutput)
}

func (e KubernetesSecretType) ToKubernetesSecretTypeOutputWithContext(ctx context.Context) KubernetesSecretTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KubernetesSecretTypeOutput)
}

func (e KubernetesSecretType) ToKubernetesSecretTypePtrOutput() KubernetesSecretTypePtrOutput {
	return e.ToKubernetesSecretTypePtrOutputWithContext(context.Background())
}

func (e KubernetesSecretType) ToKubernetesSecretTypePtrOutputWithContext(ctx context.Context) KubernetesSecretTypePtrOutput {
	return KubernetesSecretType(e).ToKubernetesSecretTypeOutputWithContext(ctx).ToKubernetesSecretTypePtrOutputWithContext(ctx)
}

func (e KubernetesSecretType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KubernetesSecretType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KubernetesSecretType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KubernetesSecretType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KubernetesSecretTypeOutput struct{ *pulumi.OutputState }

func (KubernetesSecretTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesSecretType)(nil)).Elem()
}

func (o KubernetesSecretTypeOutput) ToKubernetesSecretTypeOutput() KubernetesSecretTypeOutput {
	return o
}

func (o KubernetesSecretTypeOutput) ToKubernetesSecretTypeOutputWithContext(ctx context.Context) KubernetesSecretTypeOutput {
	return o
}

func (o KubernetesSecretTypeOutput) ToKubernetesSecretTypePtrOutput() KubernetesSecretTypePtrOutput {
	return o.ToKubernetesSecretTypePtrOutputWithContext(context.Background())
}

func (o KubernetesSecretTypeOutput) ToKubernetesSecretTypePtrOutputWithContext(ctx context.Context) KubernetesSecretTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesSecretType) *KubernetesSecretType {
		return &v
	}).(KubernetesSecretTypePtrOutput)
}

func (o KubernetesSecretTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KubernetesSecretTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KubernetesSecretType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KubernetesSecretTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KubernetesSecretTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KubernetesSecretType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KubernetesSecretTypePtrOutput struct{ *pulumi.OutputState }

func (KubernetesSecretTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesSecretType)(nil)).Elem()
}

func (o KubernetesSecretTypePtrOutput) ToKubernetesSecretTypePtrOutput() KubernetesSecretTypePtrOutput {
	return o
}

func (o KubernetesSecretTypePtrOutput) ToKubernetesSecretTypePtrOutputWithContext(ctx context.Context) KubernetesSecretTypePtrOutput {
	return o
}

func (o KubernetesSecretTypePtrOutput) Elem() KubernetesSecretTypeOutput {
	return o.ApplyT(func(v *KubernetesSecretType) KubernetesSecretType {
		if v != nil {
			return *v
		}
		var ret KubernetesSecretType
		return ret
	}).(KubernetesSecretTypeOutput)
}

func (o KubernetesSecretTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KubernetesSecretTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KubernetesSecretType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KubernetesSecretTypeInput is an input type that accepts values of the KubernetesSecretType enum
// A concrete instance of `KubernetesSecretTypeInput` can be one of the following:
//
//	KubernetesSecretTypeOpaque
//	KubernetesSecretTypeTls
type KubernetesSecretTypeInput interface {
	pulumi.Input

	ToKubernetesSecretTypeOutput() KubernetesSecretTypeOutput
	ToKubernetesSecretTypeOutputWithContext(context.Context) KubernetesSecretTypeOutput
}

var kubernetesSecretTypePtrType = reflect.TypeOf((**KubernetesSecretType)(nil)).Elem()

type KubernetesSecretTypePtrInput interface {
	pulumi.Input

	ToKubernetesSecretTypePtrOutput() KubernetesSecretTypePtrOutput
	ToKubernetesSecretTypePtrOutputWithContext(context.Context) KubernetesSecretTypePtrOutput
}

type kubernetesSecretTypePtr string

func KubernetesSecretTypePtr(v string) KubernetesSecretTypePtrInput {
	return (*kubernetesSecretTypePtr)(&v)
}

func (*kubernetesSecretTypePtr) ElementType() reflect.Type {
	return kubernetesSecretTypePtrType
}

func (in *kubernetesSecretTypePtr) ToKubernetesSecretTypePtrOutput() KubernetesSecretTypePtrOutput {
	return pulumi.ToOutput(in).(KubernetesSecretTypePtrOutput)
}

func (in *kubernetesSecretTypePtr) ToKubernetesSecretTypePtrOutputWithContext(ctx context.Context) KubernetesSecretTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KubernetesSecretTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtendedLocationTypeOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypePtrOutput{})
	pulumi.RegisterOutputType(KubernetesSecretTypeOutput{})
	pulumi.RegisterOutputType(KubernetesSecretTypePtrOutput{})
}
