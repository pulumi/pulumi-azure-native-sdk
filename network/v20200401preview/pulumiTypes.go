// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP configuration.
type InboundEndpointIPConfiguration struct {
	// Private IP address of the IP configuration.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// Private IP address allocation method.
	PrivateIpAllocationMethod *string `pulumi:"privateIpAllocationMethod"`
	// The reference to the subnet bound to the IP configuration.
	Subnet *SubResource `pulumi:"subnet"`
}

// Defaults sets the appropriate defaults for InboundEndpointIPConfiguration
func (val *InboundEndpointIPConfiguration) Defaults() *InboundEndpointIPConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateIpAllocationMethod) {
		privateIpAllocationMethod_ := "Dynamic"
		tmp.PrivateIpAllocationMethod = &privateIpAllocationMethod_
	}
	return &tmp
}

// InboundEndpointIPConfigurationInput is an input type that accepts InboundEndpointIPConfigurationArgs and InboundEndpointIPConfigurationOutput values.
// You can construct a concrete instance of `InboundEndpointIPConfigurationInput` via:
//
//	InboundEndpointIPConfigurationArgs{...}
type InboundEndpointIPConfigurationInput interface {
	pulumi.Input

	ToInboundEndpointIPConfigurationOutput() InboundEndpointIPConfigurationOutput
	ToInboundEndpointIPConfigurationOutputWithContext(context.Context) InboundEndpointIPConfigurationOutput
}

// IP configuration.
type InboundEndpointIPConfigurationArgs struct {
	// Private IP address of the IP configuration.
	PrivateIpAddress pulumi.StringPtrInput `pulumi:"privateIpAddress"`
	// Private IP address allocation method.
	PrivateIpAllocationMethod pulumi.StringPtrInput `pulumi:"privateIpAllocationMethod"`
	// The reference to the subnet bound to the IP configuration.
	Subnet SubResourcePtrInput `pulumi:"subnet"`
}

// Defaults sets the appropriate defaults for InboundEndpointIPConfigurationArgs
func (val *InboundEndpointIPConfigurationArgs) Defaults() *InboundEndpointIPConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateIpAllocationMethod) {
		tmp.PrivateIpAllocationMethod = pulumi.StringPtr("Dynamic")
	}
	return &tmp
}
func (InboundEndpointIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointIPConfiguration)(nil)).Elem()
}

func (i InboundEndpointIPConfigurationArgs) ToInboundEndpointIPConfigurationOutput() InboundEndpointIPConfigurationOutput {
	return i.ToInboundEndpointIPConfigurationOutputWithContext(context.Background())
}

func (i InboundEndpointIPConfigurationArgs) ToInboundEndpointIPConfigurationOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundEndpointIPConfigurationOutput)
}

// InboundEndpointIPConfigurationArrayInput is an input type that accepts InboundEndpointIPConfigurationArray and InboundEndpointIPConfigurationArrayOutput values.
// You can construct a concrete instance of `InboundEndpointIPConfigurationArrayInput` via:
//
//	InboundEndpointIPConfigurationArray{ InboundEndpointIPConfigurationArgs{...} }
type InboundEndpointIPConfigurationArrayInput interface {
	pulumi.Input

	ToInboundEndpointIPConfigurationArrayOutput() InboundEndpointIPConfigurationArrayOutput
	ToInboundEndpointIPConfigurationArrayOutputWithContext(context.Context) InboundEndpointIPConfigurationArrayOutput
}

type InboundEndpointIPConfigurationArray []InboundEndpointIPConfigurationInput

func (InboundEndpointIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundEndpointIPConfiguration)(nil)).Elem()
}

func (i InboundEndpointIPConfigurationArray) ToInboundEndpointIPConfigurationArrayOutput() InboundEndpointIPConfigurationArrayOutput {
	return i.ToInboundEndpointIPConfigurationArrayOutputWithContext(context.Background())
}

func (i InboundEndpointIPConfigurationArray) ToInboundEndpointIPConfigurationArrayOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundEndpointIPConfigurationArrayOutput)
}

// IP configuration.
type InboundEndpointIPConfigurationOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointIPConfiguration)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationOutput) ToInboundEndpointIPConfigurationOutput() InboundEndpointIPConfigurationOutput {
	return o
}

func (o InboundEndpointIPConfigurationOutput) ToInboundEndpointIPConfigurationOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationOutput {
	return o
}

// Private IP address of the IP configuration.
func (o InboundEndpointIPConfigurationOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfiguration) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// Private IP address allocation method.
func (o InboundEndpointIPConfigurationOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfiguration) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

// The reference to the subnet bound to the IP configuration.
func (o InboundEndpointIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type InboundEndpointIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundEndpointIPConfiguration)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationArrayOutput) ToInboundEndpointIPConfigurationArrayOutput() InboundEndpointIPConfigurationArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationArrayOutput) ToInboundEndpointIPConfigurationArrayOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationArrayOutput) Index(i pulumi.IntInput) InboundEndpointIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundEndpointIPConfiguration {
		return vs[0].([]InboundEndpointIPConfiguration)[vs[1].(int)]
	}).(InboundEndpointIPConfigurationOutput)
}

// IP configuration.
type InboundEndpointIPConfigurationResponse struct {
	// Private IP address of the IP configuration.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// Private IP address allocation method.
	PrivateIpAllocationMethod *string `pulumi:"privateIpAllocationMethod"`
	// The reference to the subnet bound to the IP configuration.
	Subnet *SubResourceResponse `pulumi:"subnet"`
}

// Defaults sets the appropriate defaults for InboundEndpointIPConfigurationResponse
func (val *InboundEndpointIPConfigurationResponse) Defaults() *InboundEndpointIPConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateIpAllocationMethod) {
		privateIpAllocationMethod_ := "Dynamic"
		tmp.PrivateIpAllocationMethod = &privateIpAllocationMethod_
	}
	return &tmp
}

// IP configuration.
type InboundEndpointIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointIPConfigurationResponse)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationResponseOutput) ToInboundEndpointIPConfigurationResponseOutput() InboundEndpointIPConfigurationResponseOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseOutput) ToInboundEndpointIPConfigurationResponseOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationResponseOutput {
	return o
}

// Private IP address of the IP configuration.
func (o InboundEndpointIPConfigurationResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfigurationResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// Private IP address allocation method.
func (o InboundEndpointIPConfigurationResponseOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfigurationResponse) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

// The reference to the subnet bound to the IP configuration.
func (o InboundEndpointIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type InboundEndpointIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundEndpointIPConfigurationResponse)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationResponseArrayOutput) ToInboundEndpointIPConfigurationResponseArrayOutput() InboundEndpointIPConfigurationResponseArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseArrayOutput) ToInboundEndpointIPConfigurationResponseArrayOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationResponseArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) InboundEndpointIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundEndpointIPConfigurationResponse {
		return vs[0].([]InboundEndpointIPConfigurationResponse)[vs[1].(int)]
	}).(InboundEndpointIPConfigurationResponseOutput)
}

// Reference to another ARM resource.
type SubResource struct {
	// Resource ID.
	Id *string `pulumi:"id"`
}

// SubResourceInput is an input type that accepts SubResourceArgs and SubResourceOutput values.
// You can construct a concrete instance of `SubResourceInput` via:
//
//	SubResourceArgs{...}
type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

// Reference to another ARM resource.
type SubResourceArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}

// SubResourcePtrInput is an input type that accepts SubResourceArgs, SubResourcePtr and SubResourcePtrOutput values.
// You can construct a concrete instance of `SubResourcePtrInput` via:
//
//	        SubResourceArgs{...}
//
//	or:
//
//	        nil
type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}

// SubResourceArrayInput is an input type that accepts SubResourceArray and SubResourceArrayOutput values.
// You can construct a concrete instance of `SubResourceArrayInput` via:
//
//	SubResourceArray{ SubResourceArgs{...} }
type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

// Reference to another ARM resource.
type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

// Resource ID.
func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

// Resource ID.
func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

// Reference to another ARM resource.
type SubResourceResponse struct {
	// Resource ID.
	Id *string `pulumi:"id"`
}

// Reference to another ARM resource.
type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

// Resource ID.
func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

// Resource ID.
func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// Describes a server to forward the DNS queries to.
type TargetDnsServer struct {
	// DNS server IP address.
	IpAddress *string `pulumi:"ipAddress"`
	// DNS server port.
	Port *int `pulumi:"port"`
}

// Defaults sets the appropriate defaults for TargetDnsServer
func (val *TargetDnsServer) Defaults() *TargetDnsServer {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		port_ := 53
		tmp.Port = &port_
	}
	return &tmp
}

// TargetDnsServerInput is an input type that accepts TargetDnsServerArgs and TargetDnsServerOutput values.
// You can construct a concrete instance of `TargetDnsServerInput` via:
//
//	TargetDnsServerArgs{...}
type TargetDnsServerInput interface {
	pulumi.Input

	ToTargetDnsServerOutput() TargetDnsServerOutput
	ToTargetDnsServerOutputWithContext(context.Context) TargetDnsServerOutput
}

// Describes a server to forward the DNS queries to.
type TargetDnsServerArgs struct {
	// DNS server IP address.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// DNS server port.
	Port pulumi.IntPtrInput `pulumi:"port"`
}

// Defaults sets the appropriate defaults for TargetDnsServerArgs
func (val *TargetDnsServerArgs) Defaults() *TargetDnsServerArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		tmp.Port = pulumi.IntPtr(53)
	}
	return &tmp
}
func (TargetDnsServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServer)(nil)).Elem()
}

func (i TargetDnsServerArgs) ToTargetDnsServerOutput() TargetDnsServerOutput {
	return i.ToTargetDnsServerOutputWithContext(context.Background())
}

func (i TargetDnsServerArgs) ToTargetDnsServerOutputWithContext(ctx context.Context) TargetDnsServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetDnsServerOutput)
}

// TargetDnsServerArrayInput is an input type that accepts TargetDnsServerArray and TargetDnsServerArrayOutput values.
// You can construct a concrete instance of `TargetDnsServerArrayInput` via:
//
//	TargetDnsServerArray{ TargetDnsServerArgs{...} }
type TargetDnsServerArrayInput interface {
	pulumi.Input

	ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput
	ToTargetDnsServerArrayOutputWithContext(context.Context) TargetDnsServerArrayOutput
}

type TargetDnsServerArray []TargetDnsServerInput

func (TargetDnsServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServer)(nil)).Elem()
}

func (i TargetDnsServerArray) ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput {
	return i.ToTargetDnsServerArrayOutputWithContext(context.Background())
}

func (i TargetDnsServerArray) ToTargetDnsServerArrayOutputWithContext(ctx context.Context) TargetDnsServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetDnsServerArrayOutput)
}

// Describes a server to forward the DNS queries to.
type TargetDnsServerOutput struct{ *pulumi.OutputState }

func (TargetDnsServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServer)(nil)).Elem()
}

func (o TargetDnsServerOutput) ToTargetDnsServerOutput() TargetDnsServerOutput {
	return o
}

func (o TargetDnsServerOutput) ToTargetDnsServerOutputWithContext(ctx context.Context) TargetDnsServerOutput {
	return o
}

// DNS server IP address.
func (o TargetDnsServerOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetDnsServer) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// DNS server port.
func (o TargetDnsServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetDnsServer) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type TargetDnsServerArrayOutput struct{ *pulumi.OutputState }

func (TargetDnsServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServer)(nil)).Elem()
}

func (o TargetDnsServerArrayOutput) ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput {
	return o
}

func (o TargetDnsServerArrayOutput) ToTargetDnsServerArrayOutputWithContext(ctx context.Context) TargetDnsServerArrayOutput {
	return o
}

func (o TargetDnsServerArrayOutput) Index(i pulumi.IntInput) TargetDnsServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetDnsServer {
		return vs[0].([]TargetDnsServer)[vs[1].(int)]
	}).(TargetDnsServerOutput)
}

// Describes a server to forward the DNS queries to.
type TargetDnsServerResponse struct {
	// DNS server IP address.
	IpAddress *string `pulumi:"ipAddress"`
	// DNS server port.
	Port *int `pulumi:"port"`
}

// Defaults sets the appropriate defaults for TargetDnsServerResponse
func (val *TargetDnsServerResponse) Defaults() *TargetDnsServerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		port_ := 53
		tmp.Port = &port_
	}
	return &tmp
}

// Describes a server to forward the DNS queries to.
type TargetDnsServerResponseOutput struct{ *pulumi.OutputState }

func (TargetDnsServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServerResponse)(nil)).Elem()
}

func (o TargetDnsServerResponseOutput) ToTargetDnsServerResponseOutput() TargetDnsServerResponseOutput {
	return o
}

func (o TargetDnsServerResponseOutput) ToTargetDnsServerResponseOutputWithContext(ctx context.Context) TargetDnsServerResponseOutput {
	return o
}

// DNS server IP address.
func (o TargetDnsServerResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetDnsServerResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// DNS server port.
func (o TargetDnsServerResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetDnsServerResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type TargetDnsServerResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetDnsServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServerResponse)(nil)).Elem()
}

func (o TargetDnsServerResponseArrayOutput) ToTargetDnsServerResponseArrayOutput() TargetDnsServerResponseArrayOutput {
	return o
}

func (o TargetDnsServerResponseArrayOutput) ToTargetDnsServerResponseArrayOutputWithContext(ctx context.Context) TargetDnsServerResponseArrayOutput {
	return o
}

func (o TargetDnsServerResponseArrayOutput) Index(i pulumi.IntInput) TargetDnsServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetDnsServerResponse {
		return vs[0].([]TargetDnsServerResponse)[vs[1].(int)]
	}).(TargetDnsServerResponseOutput)
}

// Reference to DNS forwarding ruleset and associated virtual network link.
type VirtualNetworkDnsForwardingRulesetResponse struct {
	// DNS Forwarding Ruleset Resource ID.
	Id *string `pulumi:"id"`
	// The reference to the virtual network link.
	VirtualNetworkLink *SubResourceResponse `pulumi:"virtualNetworkLink"`
}

// Reference to DNS forwarding ruleset and associated virtual network link.
type VirtualNetworkDnsForwardingRulesetResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsForwardingRulesetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkDnsForwardingRulesetResponse)(nil)).Elem()
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) ToVirtualNetworkDnsForwardingRulesetResponseOutput() VirtualNetworkDnsForwardingRulesetResponseOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) ToVirtualNetworkDnsForwardingRulesetResponseOutputWithContext(ctx context.Context) VirtualNetworkDnsForwardingRulesetResponseOutput {
	return o
}

// DNS Forwarding Ruleset Resource ID.
func (o VirtualNetworkDnsForwardingRulesetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkDnsForwardingRulesetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The reference to the virtual network link.
func (o VirtualNetworkDnsForwardingRulesetResponseOutput) VirtualNetworkLink() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkDnsForwardingRulesetResponse) *SubResourceResponse { return v.VirtualNetworkLink }).(SubResourceResponsePtrOutput)
}

type VirtualNetworkDnsForwardingRulesetResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkDnsForwardingRulesetResponse)(nil)).Elem()
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ToVirtualNetworkDnsForwardingRulesetResponseArrayOutput() VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ToVirtualNetworkDnsForwardingRulesetResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkDnsForwardingRulesetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkDnsForwardingRulesetResponse {
		return vs[0].([]VirtualNetworkDnsForwardingRulesetResponse)[vs[1].(int)]
	}).(VirtualNetworkDnsForwardingRulesetResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationOutput{})
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TargetDnsServerOutput{})
	pulumi.RegisterOutputType(TargetDnsServerArrayOutput{})
	pulumi.RegisterOutputType(TargetDnsServerResponseOutput{})
	pulumi.RegisterOutputType(TargetDnsServerResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsForwardingRulesetResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsForwardingRulesetResponseArrayOutput{})
}