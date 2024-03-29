// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Interval value in minutes used to create LogAnalytics call rate logs.
type IntervalInMins string

const (
	IntervalInMinsThreeMins  = IntervalInMins("ThreeMins")
	IntervalInMinsFiveMins   = IntervalInMins("FiveMins")
	IntervalInMinsThirtyMins = IntervalInMins("ThirtyMins")
	IntervalInMinsSixtyMins  = IntervalInMins("SixtyMins")
)

func (IntervalInMins) ElementType() reflect.Type {
	return reflect.TypeOf((*IntervalInMins)(nil)).Elem()
}

func (e IntervalInMins) ToIntervalInMinsOutput() IntervalInMinsOutput {
	return pulumi.ToOutput(e).(IntervalInMinsOutput)
}

func (e IntervalInMins) ToIntervalInMinsOutputWithContext(ctx context.Context) IntervalInMinsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntervalInMinsOutput)
}

func (e IntervalInMins) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return e.ToIntervalInMinsPtrOutputWithContext(context.Background())
}

func (e IntervalInMins) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return IntervalInMins(e).ToIntervalInMinsOutputWithContext(ctx).ToIntervalInMinsPtrOutputWithContext(ctx)
}

func (e IntervalInMins) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntervalInMins) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntervalInMins) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntervalInMins) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntervalInMinsOutput struct{ *pulumi.OutputState }

func (IntervalInMinsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntervalInMins)(nil)).Elem()
}

func (o IntervalInMinsOutput) ToIntervalInMinsOutput() IntervalInMinsOutput {
	return o
}

func (o IntervalInMinsOutput) ToIntervalInMinsOutputWithContext(ctx context.Context) IntervalInMinsOutput {
	return o
}

func (o IntervalInMinsOutput) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return o.ToIntervalInMinsPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntervalInMins) *IntervalInMins {
		return &v
	}).(IntervalInMinsPtrOutput)
}

func (o IntervalInMinsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntervalInMins) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntervalInMinsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntervalInMins) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntervalInMinsPtrOutput struct{ *pulumi.OutputState }

func (IntervalInMinsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntervalInMins)(nil)).Elem()
}

func (o IntervalInMinsPtrOutput) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return o
}

func (o IntervalInMinsPtrOutput) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return o
}

func (o IntervalInMinsPtrOutput) Elem() IntervalInMinsOutput {
	return o.ApplyT(func(v *IntervalInMins) IntervalInMins {
		if v != nil {
			return *v
		}
		var ret IntervalInMins
		return ret
	}).(IntervalInMinsOutput)
}

func (o IntervalInMinsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntervalInMins) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IntervalInMinsInput is an input type that accepts values of the IntervalInMins enum
// A concrete instance of `IntervalInMinsInput` can be one of the following:
//
//	IntervalInMinsThreeMins
//	IntervalInMinsFiveMins
//	IntervalInMinsThirtyMins
//	IntervalInMinsSixtyMins
type IntervalInMinsInput interface {
	pulumi.Input

	ToIntervalInMinsOutput() IntervalInMinsOutput
	ToIntervalInMinsOutputWithContext(context.Context) IntervalInMinsOutput
}

var intervalInMinsPtrType = reflect.TypeOf((**IntervalInMins)(nil)).Elem()

type IntervalInMinsPtrInput interface {
	pulumi.Input

	ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput
	ToIntervalInMinsPtrOutputWithContext(context.Context) IntervalInMinsPtrOutput
}

type intervalInMinsPtr string

func IntervalInMinsPtr(v string) IntervalInMinsPtrInput {
	return (*intervalInMinsPtr)(&v)
}

func (*intervalInMinsPtr) ElementType() reflect.Type {
	return intervalInMinsPtrType
}

func (in *intervalInMinsPtr) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return pulumi.ToOutput(in).(IntervalInMinsPtrOutput)
}

func (in *intervalInMinsPtr) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntervalInMinsPtrOutput)
}

// The level code.
type StatusLevelTypes string

const (
	StatusLevelTypesInfo    = StatusLevelTypes("Info")
	StatusLevelTypesWarning = StatusLevelTypes("Warning")
	StatusLevelTypesError   = StatusLevelTypes("Error")
)

func (StatusLevelTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusLevelTypes)(nil)).Elem()
}

func (e StatusLevelTypes) ToStatusLevelTypesOutput() StatusLevelTypesOutput {
	return pulumi.ToOutput(e).(StatusLevelTypesOutput)
}

func (e StatusLevelTypes) ToStatusLevelTypesOutputWithContext(ctx context.Context) StatusLevelTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusLevelTypesOutput)
}

func (e StatusLevelTypes) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return e.ToStatusLevelTypesPtrOutputWithContext(context.Background())
}

func (e StatusLevelTypes) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return StatusLevelTypes(e).ToStatusLevelTypesOutputWithContext(ctx).ToStatusLevelTypesPtrOutputWithContext(ctx)
}

func (e StatusLevelTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusLevelTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusLevelTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusLevelTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusLevelTypesOutput struct{ *pulumi.OutputState }

func (StatusLevelTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusLevelTypes)(nil)).Elem()
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesOutput() StatusLevelTypesOutput {
	return o
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesOutputWithContext(ctx context.Context) StatusLevelTypesOutput {
	return o
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return o.ToStatusLevelTypesPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusLevelTypes) *StatusLevelTypes {
		return &v
	}).(StatusLevelTypesPtrOutput)
}

func (o StatusLevelTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusLevelTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusLevelTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusLevelTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusLevelTypesPtrOutput struct{ *pulumi.OutputState }

func (StatusLevelTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusLevelTypes)(nil)).Elem()
}

func (o StatusLevelTypesPtrOutput) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return o
}

func (o StatusLevelTypesPtrOutput) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return o
}

func (o StatusLevelTypesPtrOutput) Elem() StatusLevelTypesOutput {
	return o.ApplyT(func(v *StatusLevelTypes) StatusLevelTypes {
		if v != nil {
			return *v
		}
		var ret StatusLevelTypes
		return ret
	}).(StatusLevelTypesOutput)
}

func (o StatusLevelTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusLevelTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StatusLevelTypesInput is an input type that accepts values of the StatusLevelTypes enum
// A concrete instance of `StatusLevelTypesInput` can be one of the following:
//
//	StatusLevelTypesInfo
//	StatusLevelTypesWarning
//	StatusLevelTypesError
type StatusLevelTypesInput interface {
	pulumi.Input

	ToStatusLevelTypesOutput() StatusLevelTypesOutput
	ToStatusLevelTypesOutputWithContext(context.Context) StatusLevelTypesOutput
}

var statusLevelTypesPtrType = reflect.TypeOf((**StatusLevelTypes)(nil)).Elem()

type StatusLevelTypesPtrInput interface {
	pulumi.Input

	ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput
	ToStatusLevelTypesPtrOutputWithContext(context.Context) StatusLevelTypesPtrOutput
}

type statusLevelTypesPtr string

func StatusLevelTypesPtr(v string) StatusLevelTypesPtrInput {
	return (*statusLevelTypesPtr)(&v)
}

func (*statusLevelTypesPtr) ElementType() reflect.Type {
	return statusLevelTypesPtrType
}

func (in *statusLevelTypesPtr) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return pulumi.ToOutput(in).(StatusLevelTypesPtrOutput)
}

func (in *statusLevelTypesPtr) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusLevelTypesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IntervalInMinsOutput{})
	pulumi.RegisterOutputType(IntervalInMinsPtrOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesPtrOutput{})
}
