// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The IP Filter Action
type IPAction string

const (
	IPActionAccept = IPAction("Accept")
	IPActionReject = IPAction("Reject")
)

func (IPAction) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAction)(nil)).Elem()
}

func (e IPAction) ToIPActionOutput() IPActionOutput {
	return pulumi.ToOutput(e).(IPActionOutput)
}

func (e IPAction) ToIPActionOutputWithContext(ctx context.Context) IPActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPActionOutput)
}

func (e IPAction) ToIPActionPtrOutput() IPActionPtrOutput {
	return e.ToIPActionPtrOutputWithContext(context.Background())
}

func (e IPAction) ToIPActionPtrOutputWithContext(ctx context.Context) IPActionPtrOutput {
	return IPAction(e).ToIPActionOutputWithContext(ctx).ToIPActionPtrOutputWithContext(ctx)
}

func (e IPAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPActionOutput struct{ *pulumi.OutputState }

func (IPActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAction)(nil)).Elem()
}

func (o IPActionOutput) ToIPActionOutput() IPActionOutput {
	return o
}

func (o IPActionOutput) ToIPActionOutputWithContext(ctx context.Context) IPActionOutput {
	return o
}

func (o IPActionOutput) ToIPActionPtrOutput() IPActionPtrOutput {
	return o.ToIPActionPtrOutputWithContext(context.Background())
}

func (o IPActionOutput) ToIPActionPtrOutputWithContext(ctx context.Context) IPActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAction) *IPAction {
		return &v
	}).(IPActionPtrOutput)
}

func (o IPActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPActionPtrOutput struct{ *pulumi.OutputState }

func (IPActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAction)(nil)).Elem()
}

func (o IPActionPtrOutput) ToIPActionPtrOutput() IPActionPtrOutput {
	return o
}

func (o IPActionPtrOutput) ToIPActionPtrOutputWithContext(ctx context.Context) IPActionPtrOutput {
	return o
}

func (o IPActionPtrOutput) Elem() IPActionOutput {
	return o.ApplyT(func(v *IPAction) IPAction {
		if v != nil {
			return *v
		}
		var ret IPAction
		return ret
	}).(IPActionOutput)
}

func (o IPActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IPActionInput is an input type that accepts values of the IPAction enum
// A concrete instance of `IPActionInput` can be one of the following:
//
//	IPActionAccept
//	IPActionReject
type IPActionInput interface {
	pulumi.Input

	ToIPActionOutput() IPActionOutput
	ToIPActionOutputWithContext(context.Context) IPActionOutput
}

var ipactionPtrType = reflect.TypeOf((**IPAction)(nil)).Elem()

type IPActionPtrInput interface {
	pulumi.Input

	ToIPActionPtrOutput() IPActionPtrOutput
	ToIPActionPtrOutputWithContext(context.Context) IPActionPtrOutput
}

type ipactionPtr string

func IPActionPtr(v string) IPActionPtrInput {
	return (*ipactionPtr)(&v)
}

func (*ipactionPtr) ElementType() reflect.Type {
	return ipactionPtrType
}

func (in *ipactionPtr) ToIPActionPtrOutput() IPActionPtrOutput {
	return pulumi.ToOutput(in).(IPActionPtrOutput)
}

func (in *ipactionPtr) ToIPActionPtrOutputWithContext(ctx context.Context) IPActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPActionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IPActionOutput{})
	pulumi.RegisterOutputType(IPActionPtrOutput{})
}
