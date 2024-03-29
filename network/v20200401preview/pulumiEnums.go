// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private IP address allocation method.
type IpAllocationMethod string

const (
	IpAllocationMethodStatic  = IpAllocationMethod("Static")
	IpAllocationMethodDynamic = IpAllocationMethod("Dynamic")
)

func (IpAllocationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocationMethod)(nil)).Elem()
}

func (e IpAllocationMethod) ToIpAllocationMethodOutput() IpAllocationMethodOutput {
	return pulumi.ToOutput(e).(IpAllocationMethodOutput)
}

func (e IpAllocationMethod) ToIpAllocationMethodOutputWithContext(ctx context.Context) IpAllocationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpAllocationMethodOutput)
}

func (e IpAllocationMethod) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return e.ToIpAllocationMethodPtrOutputWithContext(context.Background())
}

func (e IpAllocationMethod) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return IpAllocationMethod(e).ToIpAllocationMethodOutputWithContext(ctx).ToIpAllocationMethodPtrOutputWithContext(ctx)
}

func (e IpAllocationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAllocationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAllocationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpAllocationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpAllocationMethodOutput struct{ *pulumi.OutputState }

func (IpAllocationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocationMethod)(nil)).Elem()
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodOutput() IpAllocationMethodOutput {
	return o
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodOutputWithContext(ctx context.Context) IpAllocationMethodOutput {
	return o
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return o.ToIpAllocationMethodPtrOutputWithContext(context.Background())
}

func (o IpAllocationMethodOutput) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAllocationMethod) *IpAllocationMethod {
		return &v
	}).(IpAllocationMethodPtrOutput)
}

func (o IpAllocationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpAllocationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAllocationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpAllocationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAllocationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAllocationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpAllocationMethodPtrOutput struct{ *pulumi.OutputState }

func (IpAllocationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocationMethod)(nil)).Elem()
}

func (o IpAllocationMethodPtrOutput) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return o
}

func (o IpAllocationMethodPtrOutput) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return o
}

func (o IpAllocationMethodPtrOutput) Elem() IpAllocationMethodOutput {
	return o.ApplyT(func(v *IpAllocationMethod) IpAllocationMethod {
		if v != nil {
			return *v
		}
		var ret IpAllocationMethod
		return ret
	}).(IpAllocationMethodOutput)
}

func (o IpAllocationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAllocationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpAllocationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IpAllocationMethodInput is an input type that accepts values of the IpAllocationMethod enum
// A concrete instance of `IpAllocationMethodInput` can be one of the following:
//
//	IpAllocationMethodStatic
//	IpAllocationMethodDynamic
type IpAllocationMethodInput interface {
	pulumi.Input

	ToIpAllocationMethodOutput() IpAllocationMethodOutput
	ToIpAllocationMethodOutputWithContext(context.Context) IpAllocationMethodOutput
}

var ipAllocationMethodPtrType = reflect.TypeOf((**IpAllocationMethod)(nil)).Elem()

type IpAllocationMethodPtrInput interface {
	pulumi.Input

	ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput
	ToIpAllocationMethodPtrOutputWithContext(context.Context) IpAllocationMethodPtrOutput
}

type ipAllocationMethodPtr string

func IpAllocationMethodPtr(v string) IpAllocationMethodPtrInput {
	return (*ipAllocationMethodPtr)(&v)
}

func (*ipAllocationMethodPtr) ElementType() reflect.Type {
	return ipAllocationMethodPtrType
}

func (in *ipAllocationMethodPtr) ToIpAllocationMethodPtrOutput() IpAllocationMethodPtrOutput {
	return pulumi.ToOutput(in).(IpAllocationMethodPtrOutput)
}

func (in *ipAllocationMethodPtr) ToIpAllocationMethodPtrOutputWithContext(ctx context.Context) IpAllocationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpAllocationMethodPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IpAllocationMethodOutput{})
	pulumi.RegisterOutputType(IpAllocationMethodPtrOutput{})
}
