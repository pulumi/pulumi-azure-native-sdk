// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of this DNS zone (Public or Private).
type ZoneType string

const (
	ZoneTypePublic  = ZoneType("Public")
	ZoneTypePrivate = ZoneType("Private")
)

func (ZoneType) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneType)(nil)).Elem()
}

func (e ZoneType) ToZoneTypeOutput() ZoneTypeOutput {
	return pulumi.ToOutput(e).(ZoneTypeOutput)
}

func (e ZoneType) ToZoneTypeOutputWithContext(ctx context.Context) ZoneTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ZoneTypeOutput)
}

func (e ZoneType) ToZoneTypePtrOutput() ZoneTypePtrOutput {
	return e.ToZoneTypePtrOutputWithContext(context.Background())
}

func (e ZoneType) ToZoneTypePtrOutputWithContext(ctx context.Context) ZoneTypePtrOutput {
	return ZoneType(e).ToZoneTypeOutputWithContext(ctx).ToZoneTypePtrOutputWithContext(ctx)
}

func (e ZoneType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ZoneType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ZoneTypeOutput struct{ *pulumi.OutputState }

func (ZoneTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneType)(nil)).Elem()
}

func (o ZoneTypeOutput) ToZoneTypeOutput() ZoneTypeOutput {
	return o
}

func (o ZoneTypeOutput) ToZoneTypeOutputWithContext(ctx context.Context) ZoneTypeOutput {
	return o
}

func (o ZoneTypeOutput) ToZoneTypePtrOutput() ZoneTypePtrOutput {
	return o.ToZoneTypePtrOutputWithContext(context.Background())
}

func (o ZoneTypeOutput) ToZoneTypePtrOutputWithContext(ctx context.Context) ZoneTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZoneType) *ZoneType {
		return &v
	}).(ZoneTypePtrOutput)
}

func (o ZoneTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ZoneTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZoneType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ZoneTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZoneTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZoneType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ZoneTypePtrOutput struct{ *pulumi.OutputState }

func (ZoneTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneType)(nil)).Elem()
}

func (o ZoneTypePtrOutput) ToZoneTypePtrOutput() ZoneTypePtrOutput {
	return o
}

func (o ZoneTypePtrOutput) ToZoneTypePtrOutputWithContext(ctx context.Context) ZoneTypePtrOutput {
	return o
}

func (o ZoneTypePtrOutput) Elem() ZoneTypeOutput {
	return o.ApplyT(func(v *ZoneType) ZoneType {
		if v != nil {
			return *v
		}
		var ret ZoneType
		return ret
	}).(ZoneTypeOutput)
}

func (o ZoneTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZoneTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ZoneType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ZoneTypeInput is an input type that accepts ZoneTypeArgs and ZoneTypeOutput values.
// You can construct a concrete instance of `ZoneTypeInput` via:
//
//	ZoneTypeArgs{...}
type ZoneTypeInput interface {
	pulumi.Input

	ToZoneTypeOutput() ZoneTypeOutput
	ToZoneTypeOutputWithContext(context.Context) ZoneTypeOutput
}

var zoneTypePtrType = reflect.TypeOf((**ZoneType)(nil)).Elem()

type ZoneTypePtrInput interface {
	pulumi.Input

	ToZoneTypePtrOutput() ZoneTypePtrOutput
	ToZoneTypePtrOutputWithContext(context.Context) ZoneTypePtrOutput
}

type zoneTypePtr string

func ZoneTypePtr(v string) ZoneTypePtrInput {
	return (*zoneTypePtr)(&v)
}

func (*zoneTypePtr) ElementType() reflect.Type {
	return zoneTypePtrType
}

func (in *zoneTypePtr) ToZoneTypePtrOutput() ZoneTypePtrOutput {
	return pulumi.ToOutput(in).(ZoneTypePtrOutput)
}

func (in *zoneTypePtr) ToZoneTypePtrOutputWithContext(ctx context.Context) ZoneTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ZoneTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ZoneTypeOutput{})
	pulumi.RegisterOutputType(ZoneTypePtrOutput{})
}