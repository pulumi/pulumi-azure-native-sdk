// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160930

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The orchestrator to use to manage container service cluster resources. Valid values are Swarm, DCOS, and Custom.
type ContainerServiceOchestratorTypes string

const (
	ContainerServiceOchestratorTypesSwarm      = ContainerServiceOchestratorTypes("Swarm")
	ContainerServiceOchestratorTypesDCOS       = ContainerServiceOchestratorTypes("DCOS")
	ContainerServiceOchestratorTypesCustom     = ContainerServiceOchestratorTypes("Custom")
	ContainerServiceOchestratorTypesKubernetes = ContainerServiceOchestratorTypes("Kubernetes")
)

func (ContainerServiceOchestratorTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOchestratorTypes)(nil)).Elem()
}

func (e ContainerServiceOchestratorTypes) ToContainerServiceOchestratorTypesOutput() ContainerServiceOchestratorTypesOutput {
	return pulumi.ToOutput(e).(ContainerServiceOchestratorTypesOutput)
}

func (e ContainerServiceOchestratorTypes) ToContainerServiceOchestratorTypesOutputWithContext(ctx context.Context) ContainerServiceOchestratorTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerServiceOchestratorTypesOutput)
}

func (e ContainerServiceOchestratorTypes) ToContainerServiceOchestratorTypesPtrOutput() ContainerServiceOchestratorTypesPtrOutput {
	return e.ToContainerServiceOchestratorTypesPtrOutputWithContext(context.Background())
}

func (e ContainerServiceOchestratorTypes) ToContainerServiceOchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOchestratorTypesPtrOutput {
	return ContainerServiceOchestratorTypes(e).ToContainerServiceOchestratorTypesOutputWithContext(ctx).ToContainerServiceOchestratorTypesPtrOutputWithContext(ctx)
}

func (e ContainerServiceOchestratorTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceOchestratorTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerServiceOchestratorTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerServiceOchestratorTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerServiceOchestratorTypesOutput struct{ *pulumi.OutputState }

func (ContainerServiceOchestratorTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOchestratorTypes)(nil)).Elem()
}

func (o ContainerServiceOchestratorTypesOutput) ToContainerServiceOchestratorTypesOutput() ContainerServiceOchestratorTypesOutput {
	return o
}

func (o ContainerServiceOchestratorTypesOutput) ToContainerServiceOchestratorTypesOutputWithContext(ctx context.Context) ContainerServiceOchestratorTypesOutput {
	return o
}

func (o ContainerServiceOchestratorTypesOutput) ToContainerServiceOchestratorTypesPtrOutput() ContainerServiceOchestratorTypesPtrOutput {
	return o.ToContainerServiceOchestratorTypesPtrOutputWithContext(context.Background())
}

func (o ContainerServiceOchestratorTypesOutput) ToContainerServiceOchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOchestratorTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceOchestratorTypes) *ContainerServiceOchestratorTypes {
		return &v
	}).(ContainerServiceOchestratorTypesPtrOutput)
}

func (o ContainerServiceOchestratorTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerServiceOchestratorTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceOchestratorTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerServiceOchestratorTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceOchestratorTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerServiceOchestratorTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceOchestratorTypesPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceOchestratorTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceOchestratorTypes)(nil)).Elem()
}

func (o ContainerServiceOchestratorTypesPtrOutput) ToContainerServiceOchestratorTypesPtrOutput() ContainerServiceOchestratorTypesPtrOutput {
	return o
}

func (o ContainerServiceOchestratorTypesPtrOutput) ToContainerServiceOchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOchestratorTypesPtrOutput {
	return o
}

func (o ContainerServiceOchestratorTypesPtrOutput) Elem() ContainerServiceOchestratorTypesOutput {
	return o.ApplyT(func(v *ContainerServiceOchestratorTypes) ContainerServiceOchestratorTypes {
		if v != nil {
			return *v
		}
		var ret ContainerServiceOchestratorTypes
		return ret
	}).(ContainerServiceOchestratorTypesOutput)
}

func (o ContainerServiceOchestratorTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerServiceOchestratorTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerServiceOchestratorTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ContainerServiceOchestratorTypesInput is an input type that accepts ContainerServiceOchestratorTypesArgs and ContainerServiceOchestratorTypesOutput values.
// You can construct a concrete instance of `ContainerServiceOchestratorTypesInput` via:
//
//	ContainerServiceOchestratorTypesArgs{...}
type ContainerServiceOchestratorTypesInput interface {
	pulumi.Input

	ToContainerServiceOchestratorTypesOutput() ContainerServiceOchestratorTypesOutput
	ToContainerServiceOchestratorTypesOutputWithContext(context.Context) ContainerServiceOchestratorTypesOutput
}

var containerServiceOchestratorTypesPtrType = reflect.TypeOf((**ContainerServiceOchestratorTypes)(nil)).Elem()

type ContainerServiceOchestratorTypesPtrInput interface {
	pulumi.Input

	ToContainerServiceOchestratorTypesPtrOutput() ContainerServiceOchestratorTypesPtrOutput
	ToContainerServiceOchestratorTypesPtrOutputWithContext(context.Context) ContainerServiceOchestratorTypesPtrOutput
}

type containerServiceOchestratorTypesPtr string

func ContainerServiceOchestratorTypesPtr(v string) ContainerServiceOchestratorTypesPtrInput {
	return (*containerServiceOchestratorTypesPtr)(&v)
}

func (*containerServiceOchestratorTypesPtr) ElementType() reflect.Type {
	return containerServiceOchestratorTypesPtrType
}

func (in *containerServiceOchestratorTypesPtr) ToContainerServiceOchestratorTypesPtrOutput() ContainerServiceOchestratorTypesPtrOutput {
	return pulumi.ToOutput(in).(ContainerServiceOchestratorTypesPtrOutput)
}

func (in *containerServiceOchestratorTypesPtr) ToContainerServiceOchestratorTypesPtrOutputWithContext(ctx context.Context) ContainerServiceOchestratorTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerServiceOchestratorTypesPtrOutput)
}

// Size of agent VMs.
type ContainerServiceVMSizeTypes string

const (
	ContainerServiceVMSizeTypes_Standard_A0     = ContainerServiceVMSizeTypes("Standard_A0")
	ContainerServiceVMSizeTypes_Standard_A1     = ContainerServiceVMSizeTypes("Standard_A1")
	ContainerServiceVMSizeTypes_Standard_A2     = ContainerServiceVMSizeTypes("Standard_A2")
	ContainerServiceVMSizeTypes_Standard_A3     = ContainerServiceVMSizeTypes("Standard_A3")
	ContainerServiceVMSizeTypes_Standard_A4     = ContainerServiceVMSizeTypes("Standard_A4")
	ContainerServiceVMSizeTypes_Standard_A5     = ContainerServiceVMSizeTypes("Standard_A5")
	ContainerServiceVMSizeTypes_Standard_A6     = ContainerServiceVMSizeTypes("Standard_A6")
	ContainerServiceVMSizeTypes_Standard_A7     = ContainerServiceVMSizeTypes("Standard_A7")
	ContainerServiceVMSizeTypes_Standard_A8     = ContainerServiceVMSizeTypes("Standard_A8")
	ContainerServiceVMSizeTypes_Standard_A9     = ContainerServiceVMSizeTypes("Standard_A9")
	ContainerServiceVMSizeTypes_Standard_A10    = ContainerServiceVMSizeTypes("Standard_A10")
	ContainerServiceVMSizeTypes_Standard_A11    = ContainerServiceVMSizeTypes("Standard_A11")
	ContainerServiceVMSizeTypes_Standard_D1     = ContainerServiceVMSizeTypes("Standard_D1")
	ContainerServiceVMSizeTypes_Standard_D2     = ContainerServiceVMSizeTypes("Standard_D2")
	ContainerServiceVMSizeTypes_Standard_D3     = ContainerServiceVMSizeTypes("Standard_D3")
	ContainerServiceVMSizeTypes_Standard_D4     = ContainerServiceVMSizeTypes("Standard_D4")
	ContainerServiceVMSizeTypes_Standard_D11    = ContainerServiceVMSizeTypes("Standard_D11")
	ContainerServiceVMSizeTypes_Standard_D12    = ContainerServiceVMSizeTypes("Standard_D12")
	ContainerServiceVMSizeTypes_Standard_D13    = ContainerServiceVMSizeTypes("Standard_D13")
	ContainerServiceVMSizeTypes_Standard_D14    = ContainerServiceVMSizeTypes("Standard_D14")
	ContainerServiceVMSizeTypes_Standard_D1_v2  = ContainerServiceVMSizeTypes("Standard_D1_v2")
	ContainerServiceVMSizeTypes_Standard_D2_v2  = ContainerServiceVMSizeTypes("Standard_D2_v2")
	ContainerServiceVMSizeTypes_Standard_D3_v2  = ContainerServiceVMSizeTypes("Standard_D3_v2")
	ContainerServiceVMSizeTypes_Standard_D4_v2  = ContainerServiceVMSizeTypes("Standard_D4_v2")
	ContainerServiceVMSizeTypes_Standard_D5_v2  = ContainerServiceVMSizeTypes("Standard_D5_v2")
	ContainerServiceVMSizeTypes_Standard_D11_v2 = ContainerServiceVMSizeTypes("Standard_D11_v2")
	ContainerServiceVMSizeTypes_Standard_D12_v2 = ContainerServiceVMSizeTypes("Standard_D12_v2")
	ContainerServiceVMSizeTypes_Standard_D13_v2 = ContainerServiceVMSizeTypes("Standard_D13_v2")
	ContainerServiceVMSizeTypes_Standard_D14_v2 = ContainerServiceVMSizeTypes("Standard_D14_v2")
	ContainerServiceVMSizeTypes_Standard_G1     = ContainerServiceVMSizeTypes("Standard_G1")
	ContainerServiceVMSizeTypes_Standard_G2     = ContainerServiceVMSizeTypes("Standard_G2")
	ContainerServiceVMSizeTypes_Standard_G3     = ContainerServiceVMSizeTypes("Standard_G3")
	ContainerServiceVMSizeTypes_Standard_G4     = ContainerServiceVMSizeTypes("Standard_G4")
	ContainerServiceVMSizeTypes_Standard_G5     = ContainerServiceVMSizeTypes("Standard_G5")
	ContainerServiceVMSizeTypes_Standard_DS1    = ContainerServiceVMSizeTypes("Standard_DS1")
	ContainerServiceVMSizeTypes_Standard_DS2    = ContainerServiceVMSizeTypes("Standard_DS2")
	ContainerServiceVMSizeTypes_Standard_DS3    = ContainerServiceVMSizeTypes("Standard_DS3")
	ContainerServiceVMSizeTypes_Standard_DS4    = ContainerServiceVMSizeTypes("Standard_DS4")
	ContainerServiceVMSizeTypes_Standard_DS11   = ContainerServiceVMSizeTypes("Standard_DS11")
	ContainerServiceVMSizeTypes_Standard_DS12   = ContainerServiceVMSizeTypes("Standard_DS12")
	ContainerServiceVMSizeTypes_Standard_DS13   = ContainerServiceVMSizeTypes("Standard_DS13")
	ContainerServiceVMSizeTypes_Standard_DS14   = ContainerServiceVMSizeTypes("Standard_DS14")
	ContainerServiceVMSizeTypes_Standard_GS1    = ContainerServiceVMSizeTypes("Standard_GS1")
	ContainerServiceVMSizeTypes_Standard_GS2    = ContainerServiceVMSizeTypes("Standard_GS2")
	ContainerServiceVMSizeTypes_Standard_GS3    = ContainerServiceVMSizeTypes("Standard_GS3")
	ContainerServiceVMSizeTypes_Standard_GS4    = ContainerServiceVMSizeTypes("Standard_GS4")
	ContainerServiceVMSizeTypes_Standard_GS5    = ContainerServiceVMSizeTypes("Standard_GS5")
)

func init() {
	pulumi.RegisterOutputType(ContainerServiceOchestratorTypesOutput{})
	pulumi.RegisterOutputType(ContainerServiceOchestratorTypesPtrOutput{})
}