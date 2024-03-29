// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The authentication type.
type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
)

// The application client type
type ClientType string

const (
	ClientTypeNone       = ClientType("none")
	ClientTypeDotnet     = ClientType("dotnet")
	ClientTypeJava       = ClientType("java")
	ClientTypePython     = ClientType("python")
	ClientTypeGo         = ClientType("go")
	ClientTypePhp        = ClientType("php")
	ClientTypeRuby       = ClientType("ruby")
	ClientTypeDjango     = ClientType("django")
	ClientTypeNodejs     = ClientType("nodejs")
	ClientTypeSpringBoot = ClientType("springBoot")
)

func (ClientType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientType)(nil)).Elem()
}

func (e ClientType) ToClientTypeOutput() ClientTypeOutput {
	return pulumi.ToOutput(e).(ClientTypeOutput)
}

func (e ClientType) ToClientTypeOutputWithContext(ctx context.Context) ClientTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClientTypeOutput)
}

func (e ClientType) ToClientTypePtrOutput() ClientTypePtrOutput {
	return e.ToClientTypePtrOutputWithContext(context.Background())
}

func (e ClientType) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return ClientType(e).ToClientTypeOutputWithContext(ctx).ToClientTypePtrOutputWithContext(ctx)
}

func (e ClientType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClientType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClientTypeOutput struct{ *pulumi.OutputState }

func (ClientTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientType)(nil)).Elem()
}

func (o ClientTypeOutput) ToClientTypeOutput() ClientTypeOutput {
	return o
}

func (o ClientTypeOutput) ToClientTypeOutputWithContext(ctx context.Context) ClientTypeOutput {
	return o
}

func (o ClientTypeOutput) ToClientTypePtrOutput() ClientTypePtrOutput {
	return o.ToClientTypePtrOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientType) *ClientType {
		return &v
	}).(ClientTypePtrOutput)
}

func (o ClientTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClientTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClientTypePtrOutput struct{ *pulumi.OutputState }

func (ClientTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientType)(nil)).Elem()
}

func (o ClientTypePtrOutput) ToClientTypePtrOutput() ClientTypePtrOutput {
	return o
}

func (o ClientTypePtrOutput) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return o
}

func (o ClientTypePtrOutput) Elem() ClientTypeOutput {
	return o.ApplyT(func(v *ClientType) ClientType {
		if v != nil {
			return *v
		}
		var ret ClientType
		return ret
	}).(ClientTypeOutput)
}

func (o ClientTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClientType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClientTypeInput is an input type that accepts values of the ClientType enum
// A concrete instance of `ClientTypeInput` can be one of the following:
//
//	ClientTypeNone
//	ClientTypeDotnet
//	ClientTypeJava
//	ClientTypePython
//	ClientTypeGo
//	ClientTypePhp
//	ClientTypeRuby
//	ClientTypeDjango
//	ClientTypeNodejs
//	ClientTypeSpringBoot
type ClientTypeInput interface {
	pulumi.Input

	ToClientTypeOutput() ClientTypeOutput
	ToClientTypeOutputWithContext(context.Context) ClientTypeOutput
}

var clientTypePtrType = reflect.TypeOf((**ClientType)(nil)).Elem()

type ClientTypePtrInput interface {
	pulumi.Input

	ToClientTypePtrOutput() ClientTypePtrOutput
	ToClientTypePtrOutputWithContext(context.Context) ClientTypePtrOutput
}

type clientTypePtr string

func ClientTypePtr(v string) ClientTypePtrInput {
	return (*clientTypePtr)(&v)
}

func (*clientTypePtr) ElementType() reflect.Type {
	return clientTypePtrType
}

func (in *clientTypePtr) ToClientTypePtrOutput() ClientTypePtrOutput {
	return pulumi.ToOutput(in).(ClientTypePtrOutput)
}

func (in *clientTypePtr) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClientTypePtrOutput)
}

// Type of VNet solution.
type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func (VNetSolutionType) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionType)(nil)).Elem()
}

func (e VNetSolutionType) ToVNetSolutionTypeOutput() VNetSolutionTypeOutput {
	return pulumi.ToOutput(e).(VNetSolutionTypeOutput)
}

func (e VNetSolutionType) ToVNetSolutionTypeOutputWithContext(ctx context.Context) VNetSolutionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VNetSolutionTypeOutput)
}

func (e VNetSolutionType) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return e.ToVNetSolutionTypePtrOutputWithContext(context.Background())
}

func (e VNetSolutionType) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return VNetSolutionType(e).ToVNetSolutionTypeOutputWithContext(ctx).ToVNetSolutionTypePtrOutputWithContext(ctx)
}

func (e VNetSolutionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VNetSolutionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VNetSolutionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VNetSolutionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VNetSolutionTypeOutput struct{ *pulumi.OutputState }

func (VNetSolutionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionType)(nil)).Elem()
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypeOutput() VNetSolutionTypeOutput {
	return o
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypeOutputWithContext(ctx context.Context) VNetSolutionTypeOutput {
	return o
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return o.ToVNetSolutionTypePtrOutputWithContext(context.Background())
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VNetSolutionType) *VNetSolutionType {
		return &v
	}).(VNetSolutionTypePtrOutput)
}

func (o VNetSolutionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VNetSolutionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VNetSolutionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VNetSolutionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VNetSolutionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VNetSolutionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VNetSolutionTypePtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolutionType)(nil)).Elem()
}

func (o VNetSolutionTypePtrOutput) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return o
}

func (o VNetSolutionTypePtrOutput) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return o
}

func (o VNetSolutionTypePtrOutput) Elem() VNetSolutionTypeOutput {
	return o.ApplyT(func(v *VNetSolutionType) VNetSolutionType {
		if v != nil {
			return *v
		}
		var ret VNetSolutionType
		return ret
	}).(VNetSolutionTypeOutput)
}

func (o VNetSolutionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VNetSolutionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VNetSolutionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// VNetSolutionTypeInput is an input type that accepts values of the VNetSolutionType enum
// A concrete instance of `VNetSolutionTypeInput` can be one of the following:
//
//	VNetSolutionTypeServiceEndpoint
//	VNetSolutionTypePrivateLink
type VNetSolutionTypeInput interface {
	pulumi.Input

	ToVNetSolutionTypeOutput() VNetSolutionTypeOutput
	ToVNetSolutionTypeOutputWithContext(context.Context) VNetSolutionTypeOutput
}

var vnetSolutionTypePtrType = reflect.TypeOf((**VNetSolutionType)(nil)).Elem()

type VNetSolutionTypePtrInput interface {
	pulumi.Input

	ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput
	ToVNetSolutionTypePtrOutputWithContext(context.Context) VNetSolutionTypePtrOutput
}

type vnetSolutionTypePtr string

func VNetSolutionTypePtr(v string) VNetSolutionTypePtrInput {
	return (*vnetSolutionTypePtr)(&v)
}

func (*vnetSolutionTypePtr) ElementType() reflect.Type {
	return vnetSolutionTypePtrType
}

func (in *vnetSolutionTypePtr) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return pulumi.ToOutput(in).(VNetSolutionTypePtrOutput)
}

func (in *vnetSolutionTypePtr) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VNetSolutionTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientTypeOutput{})
	pulumi.RegisterOutputType(ClientTypePtrOutput{})
	pulumi.RegisterOutputType(VNetSolutionTypeOutput{})
	pulumi.RegisterOutputType(VNetSolutionTypePtrOutput{})
}
