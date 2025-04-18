// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dependencymap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Operator for process name filter
type ProcessNameFilterOperator string

const (
	// Operator to include items in the result
	ProcessNameFilterOperatorContains = ProcessNameFilterOperator("contains")
	// Operator to exclude items in the result
	ProcessNameFilterOperatorNotContains = ProcessNameFilterOperator("notContains")
)

// Source type of Discovery Source resource.
type SourceType string

const (
	// OffAzure source type
	SourceTypeOffAzure = SourceType("OffAzure")
)

func (SourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (e SourceType) ToSourceTypeOutput() SourceTypeOutput {
	return pulumi.ToOutput(e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return e.ToSourceTypePtrOutputWithContext(context.Background())
}

func (e SourceType) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return SourceType(e).ToSourceTypeOutputWithContext(ctx).ToSourceTypePtrOutputWithContext(ctx)
}

func (e SourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceTypeOutput struct{ *pulumi.OutputState }

func (SourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (o SourceTypeOutput) ToSourceTypeOutput() SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o.ToSourceTypePtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceType) *SourceType {
		return &v
	}).(SourceTypePtrOutput)
}

func (o SourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceTypePtrOutput struct{ *pulumi.OutputState }

func (SourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceType)(nil)).Elem()
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) Elem() SourceTypeOutput {
	return o.ApplyT(func(v *SourceType) SourceType {
		if v != nil {
			return *v
		}
		var ret SourceType
		return ret
	}).(SourceTypeOutput)
}

func (o SourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SourceTypeInput is an input type that accepts values of the SourceType enum
// A concrete instance of `SourceTypeInput` can be one of the following:
//
//	SourceTypeOffAzure
type SourceTypeInput interface {
	pulumi.Input

	ToSourceTypeOutput() SourceTypeOutput
	ToSourceTypeOutputWithContext(context.Context) SourceTypeOutput
}

var sourceTypePtrType = reflect.TypeOf((**SourceType)(nil)).Elem()

type SourceTypePtrInput interface {
	pulumi.Input

	ToSourceTypePtrOutput() SourceTypePtrOutput
	ToSourceTypePtrOutputWithContext(context.Context) SourceTypePtrOutput
}

type sourceTypePtr string

func SourceTypePtr(v string) SourceTypePtrInput {
	return (*sourceTypePtr)(&v)
}

func (*sourceTypePtr) ElementType() reflect.Type {
	return sourceTypePtrType
}

func (in *sourceTypePtr) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return pulumi.ToOutput(in).(SourceTypePtrOutput)
}

func (in *sourceTypePtr) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SourceTypeOutput{})
	pulumi.RegisterOutputType(SourceTypePtrOutput{})
}
