// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The incident management service type
type IncidentManagementService string

const (
	IncidentManagementServiceIcm = IncidentManagementService("Icm")
)

func (IncidentManagementService) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentManagementService)(nil)).Elem()
}

func (e IncidentManagementService) ToIncidentManagementServiceOutput() IncidentManagementServiceOutput {
	return pulumi.ToOutput(e).(IncidentManagementServiceOutput)
}

func (e IncidentManagementService) ToIncidentManagementServiceOutputWithContext(ctx context.Context) IncidentManagementServiceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentManagementServiceOutput)
}

func (e IncidentManagementService) ToIncidentManagementServicePtrOutput() IncidentManagementServicePtrOutput {
	return e.ToIncidentManagementServicePtrOutputWithContext(context.Background())
}

func (e IncidentManagementService) ToIncidentManagementServicePtrOutputWithContext(ctx context.Context) IncidentManagementServicePtrOutput {
	return IncidentManagementService(e).ToIncidentManagementServiceOutputWithContext(ctx).ToIncidentManagementServicePtrOutputWithContext(ctx)
}

func (e IncidentManagementService) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentManagementService) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentManagementService) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentManagementService) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentManagementServiceOutput struct{ *pulumi.OutputState }

func (IncidentManagementServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentManagementService)(nil)).Elem()
}

func (o IncidentManagementServiceOutput) ToIncidentManagementServiceOutput() IncidentManagementServiceOutput {
	return o
}

func (o IncidentManagementServiceOutput) ToIncidentManagementServiceOutputWithContext(ctx context.Context) IncidentManagementServiceOutput {
	return o
}

func (o IncidentManagementServiceOutput) ToIncidentManagementServicePtrOutput() IncidentManagementServicePtrOutput {
	return o.ToIncidentManagementServicePtrOutputWithContext(context.Background())
}

func (o IncidentManagementServiceOutput) ToIncidentManagementServicePtrOutputWithContext(ctx context.Context) IncidentManagementServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentManagementService) *IncidentManagementService {
		return &v
	}).(IncidentManagementServicePtrOutput)
}

func (o IncidentManagementServiceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentManagementServiceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentManagementService) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentManagementServiceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentManagementServiceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentManagementService) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentManagementServicePtrOutput struct{ *pulumi.OutputState }

func (IncidentManagementServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentManagementService)(nil)).Elem()
}

func (o IncidentManagementServicePtrOutput) ToIncidentManagementServicePtrOutput() IncidentManagementServicePtrOutput {
	return o
}

func (o IncidentManagementServicePtrOutput) ToIncidentManagementServicePtrOutputWithContext(ctx context.Context) IncidentManagementServicePtrOutput {
	return o
}

func (o IncidentManagementServicePtrOutput) Elem() IncidentManagementServiceOutput {
	return o.ApplyT(func(v *IncidentManagementService) IncidentManagementService {
		if v != nil {
			return *v
		}
		var ret IncidentManagementService
		return ret
	}).(IncidentManagementServiceOutput)
}

func (o IncidentManagementServicePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentManagementServicePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentManagementService) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IncidentManagementServiceInput is an input type that accepts values of the IncidentManagementService enum
// A concrete instance of `IncidentManagementServiceInput` can be one of the following:
//
//	IncidentManagementServiceIcm
type IncidentManagementServiceInput interface {
	pulumi.Input

	ToIncidentManagementServiceOutput() IncidentManagementServiceOutput
	ToIncidentManagementServiceOutputWithContext(context.Context) IncidentManagementServiceOutput
}

var incidentManagementServicePtrType = reflect.TypeOf((**IncidentManagementService)(nil)).Elem()

type IncidentManagementServicePtrInput interface {
	pulumi.Input

	ToIncidentManagementServicePtrOutput() IncidentManagementServicePtrOutput
	ToIncidentManagementServicePtrOutputWithContext(context.Context) IncidentManagementServicePtrOutput
}

type incidentManagementServicePtr string

func IncidentManagementServicePtr(v string) IncidentManagementServicePtrInput {
	return (*incidentManagementServicePtr)(&v)
}

func (*incidentManagementServicePtr) ElementType() reflect.Type {
	return incidentManagementServicePtrType
}

func (in *incidentManagementServicePtr) ToIncidentManagementServicePtrOutput() IncidentManagementServicePtrOutput {
	return pulumi.ToOutput(in).(IncidentManagementServicePtrOutput)
}

func (in *incidentManagementServicePtr) ToIncidentManagementServicePtrOutputWithContext(ctx context.Context) IncidentManagementServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentManagementServicePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentManagementServiceOutput{})
	pulumi.RegisterOutputType(IncidentManagementServicePtrOutput{})
}
