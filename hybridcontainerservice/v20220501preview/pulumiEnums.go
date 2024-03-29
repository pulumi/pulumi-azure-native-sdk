// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Indicates whether the Arc agents on the provisioned clusters be upgraded automatically to the latest version. Defaults to Enabled.
type AutoUpgradeOptions string

const (
	AutoUpgradeOptionsEnabled  = AutoUpgradeOptions("Enabled")
	AutoUpgradeOptionsDisabled = AutoUpgradeOptions("Disabled")
)

func (AutoUpgradeOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUpgradeOptions)(nil)).Elem()
}

func (e AutoUpgradeOptions) ToAutoUpgradeOptionsOutput() AutoUpgradeOptionsOutput {
	return pulumi.ToOutput(e).(AutoUpgradeOptionsOutput)
}

func (e AutoUpgradeOptions) ToAutoUpgradeOptionsOutputWithContext(ctx context.Context) AutoUpgradeOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoUpgradeOptionsOutput)
}

func (e AutoUpgradeOptions) ToAutoUpgradeOptionsPtrOutput() AutoUpgradeOptionsPtrOutput {
	return e.ToAutoUpgradeOptionsPtrOutputWithContext(context.Background())
}

func (e AutoUpgradeOptions) ToAutoUpgradeOptionsPtrOutputWithContext(ctx context.Context) AutoUpgradeOptionsPtrOutput {
	return AutoUpgradeOptions(e).ToAutoUpgradeOptionsOutputWithContext(ctx).ToAutoUpgradeOptionsPtrOutputWithContext(ctx)
}

func (e AutoUpgradeOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoUpgradeOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoUpgradeOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoUpgradeOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoUpgradeOptionsOutput struct{ *pulumi.OutputState }

func (AutoUpgradeOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUpgradeOptions)(nil)).Elem()
}

func (o AutoUpgradeOptionsOutput) ToAutoUpgradeOptionsOutput() AutoUpgradeOptionsOutput {
	return o
}

func (o AutoUpgradeOptionsOutput) ToAutoUpgradeOptionsOutputWithContext(ctx context.Context) AutoUpgradeOptionsOutput {
	return o
}

func (o AutoUpgradeOptionsOutput) ToAutoUpgradeOptionsPtrOutput() AutoUpgradeOptionsPtrOutput {
	return o.ToAutoUpgradeOptionsPtrOutputWithContext(context.Background())
}

func (o AutoUpgradeOptionsOutput) ToAutoUpgradeOptionsPtrOutputWithContext(ctx context.Context) AutoUpgradeOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoUpgradeOptions) *AutoUpgradeOptions {
		return &v
	}).(AutoUpgradeOptionsPtrOutput)
}

func (o AutoUpgradeOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoUpgradeOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoUpgradeOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoUpgradeOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoUpgradeOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoUpgradeOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoUpgradeOptionsPtrOutput struct{ *pulumi.OutputState }

func (AutoUpgradeOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoUpgradeOptions)(nil)).Elem()
}

func (o AutoUpgradeOptionsPtrOutput) ToAutoUpgradeOptionsPtrOutput() AutoUpgradeOptionsPtrOutput {
	return o
}

func (o AutoUpgradeOptionsPtrOutput) ToAutoUpgradeOptionsPtrOutputWithContext(ctx context.Context) AutoUpgradeOptionsPtrOutput {
	return o
}

func (o AutoUpgradeOptionsPtrOutput) Elem() AutoUpgradeOptionsOutput {
	return o.ApplyT(func(v *AutoUpgradeOptions) AutoUpgradeOptions {
		if v != nil {
			return *v
		}
		var ret AutoUpgradeOptions
		return ret
	}).(AutoUpgradeOptionsOutput)
}

func (o AutoUpgradeOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoUpgradeOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoUpgradeOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AutoUpgradeOptionsInput is an input type that accepts values of the AutoUpgradeOptions enum
// A concrete instance of `AutoUpgradeOptionsInput` can be one of the following:
//
//	AutoUpgradeOptionsEnabled
//	AutoUpgradeOptionsDisabled
type AutoUpgradeOptionsInput interface {
	pulumi.Input

	ToAutoUpgradeOptionsOutput() AutoUpgradeOptionsOutput
	ToAutoUpgradeOptionsOutputWithContext(context.Context) AutoUpgradeOptionsOutput
}

var autoUpgradeOptionsPtrType = reflect.TypeOf((**AutoUpgradeOptions)(nil)).Elem()

type AutoUpgradeOptionsPtrInput interface {
	pulumi.Input

	ToAutoUpgradeOptionsPtrOutput() AutoUpgradeOptionsPtrOutput
	ToAutoUpgradeOptionsPtrOutputWithContext(context.Context) AutoUpgradeOptionsPtrOutput
}

type autoUpgradeOptionsPtr string

func AutoUpgradeOptionsPtr(v string) AutoUpgradeOptionsPtrInput {
	return (*autoUpgradeOptionsPtr)(&v)
}

func (*autoUpgradeOptionsPtr) ElementType() reflect.Type {
	return autoUpgradeOptionsPtrType
}

func (in *autoUpgradeOptionsPtr) ToAutoUpgradeOptionsPtrOutput() AutoUpgradeOptionsPtrOutput {
	return pulumi.ToOutput(in).(AutoUpgradeOptionsPtrOutput)
}

func (in *autoUpgradeOptionsPtr) ToAutoUpgradeOptionsPtrOutputWithContext(ctx context.Context) AutoUpgradeOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoUpgradeOptionsPtrOutput)
}

// LicenseType - The licenseType to use for Windows VMs. Windows_Server is used to enable Azure Hybrid User Benefits for Windows VMs. Possible values include: 'None', 'Windows_Server'
type LicenseType string

const (
	LicenseType_Windows_Server = LicenseType("Windows_Server")
	LicenseTypeNone            = LicenseType("None")
)

func (LicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*LicenseType)(nil)).Elem()
}

func (e LicenseType) ToLicenseTypeOutput() LicenseTypeOutput {
	return pulumi.ToOutput(e).(LicenseTypeOutput)
}

func (e LicenseType) ToLicenseTypeOutputWithContext(ctx context.Context) LicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LicenseTypeOutput)
}

func (e LicenseType) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return e.ToLicenseTypePtrOutputWithContext(context.Background())
}

func (e LicenseType) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return LicenseType(e).ToLicenseTypeOutputWithContext(ctx).ToLicenseTypePtrOutputWithContext(ctx)
}

func (e LicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LicenseTypeOutput struct{ *pulumi.OutputState }

func (LicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LicenseType)(nil)).Elem()
}

func (o LicenseTypeOutput) ToLicenseTypeOutput() LicenseTypeOutput {
	return o
}

func (o LicenseTypeOutput) ToLicenseTypeOutputWithContext(ctx context.Context) LicenseTypeOutput {
	return o
}

func (o LicenseTypeOutput) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return o.ToLicenseTypePtrOutputWithContext(context.Background())
}

func (o LicenseTypeOutput) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LicenseType) *LicenseType {
		return &v
	}).(LicenseTypePtrOutput)
}

func (o LicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LicenseTypePtrOutput struct{ *pulumi.OutputState }

func (LicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseType)(nil)).Elem()
}

func (o LicenseTypePtrOutput) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return o
}

func (o LicenseTypePtrOutput) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return o
}

func (o LicenseTypePtrOutput) Elem() LicenseTypeOutput {
	return o.ApplyT(func(v *LicenseType) LicenseType {
		if v != nil {
			return *v
		}
		var ret LicenseType
		return ret
	}).(LicenseTypeOutput)
}

func (o LicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// LicenseTypeInput is an input type that accepts values of the LicenseType enum
// A concrete instance of `LicenseTypeInput` can be one of the following:
//
//	LicenseType_Windows_Server
//	LicenseTypeNone
type LicenseTypeInput interface {
	pulumi.Input

	ToLicenseTypeOutput() LicenseTypeOutput
	ToLicenseTypeOutputWithContext(context.Context) LicenseTypeOutput
}

var licenseTypePtrType = reflect.TypeOf((**LicenseType)(nil)).Elem()

type LicenseTypePtrInput interface {
	pulumi.Input

	ToLicenseTypePtrOutput() LicenseTypePtrOutput
	ToLicenseTypePtrOutputWithContext(context.Context) LicenseTypePtrOutput
}

type licenseTypePtr string

func LicenseTypePtr(v string) LicenseTypePtrInput {
	return (*licenseTypePtr)(&v)
}

func (*licenseTypePtr) ElementType() reflect.Type {
	return licenseTypePtrType
}

func (in *licenseTypePtr) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return pulumi.ToOutput(in).(LicenseTypePtrOutput)
}

func (in *licenseTypePtr) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LicenseTypePtrOutput)
}

// LoadBalancerSku - The load balancer sku for the provisioned cluster. Possible values: 'unstacked-haproxy', 'stacked-kube-vip', 'stacked-metallb', 'unmanaged'. The default is 'unmanaged'.
type LoadBalancerSku string

const (
	LoadBalancerSku_Unstacked_Haproxy = LoadBalancerSku("unstacked-haproxy")
	LoadBalancerSku_Stacked_Kube_Vip  = LoadBalancerSku("stacked-kube-vip")
	LoadBalancerSku_Stacked_Metallb   = LoadBalancerSku("stacked-metallb")
	LoadBalancerSkuUnmanaged          = LoadBalancerSku("unmanaged")
)

func (LoadBalancerSku) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSku)(nil)).Elem()
}

func (e LoadBalancerSku) ToLoadBalancerSkuOutput() LoadBalancerSkuOutput {
	return pulumi.ToOutput(e).(LoadBalancerSkuOutput)
}

func (e LoadBalancerSku) ToLoadBalancerSkuOutputWithContext(ctx context.Context) LoadBalancerSkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerSkuOutput)
}

func (e LoadBalancerSku) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return e.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSku) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return LoadBalancerSku(e).ToLoadBalancerSkuOutputWithContext(ctx).ToLoadBalancerSkuPtrOutputWithContext(ctx)
}

func (e LoadBalancerSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerSkuOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSku)(nil)).Elem()
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuOutput() LoadBalancerSkuOutput {
	return o
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuOutputWithContext(ctx context.Context) LoadBalancerSkuOutput {
	return o
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return o.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerSku) *LoadBalancerSku {
		return &v
	}).(LoadBalancerSkuPtrOutput)
}

func (o LoadBalancerSkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerSkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSku)(nil)).Elem()
}

func (o LoadBalancerSkuPtrOutput) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return o
}

func (o LoadBalancerSkuPtrOutput) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return o
}

func (o LoadBalancerSkuPtrOutput) Elem() LoadBalancerSkuOutput {
	return o.ApplyT(func(v *LoadBalancerSku) LoadBalancerSku {
		if v != nil {
			return *v
		}
		var ret LoadBalancerSku
		return ret
	}).(LoadBalancerSkuOutput)
}

func (o LoadBalancerSkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerSku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// LoadBalancerSkuInput is an input type that accepts values of the LoadBalancerSku enum
// A concrete instance of `LoadBalancerSkuInput` can be one of the following:
//
//	LoadBalancerSku_Unstacked_Haproxy
//	LoadBalancerSku_Stacked_Kube_Vip
//	LoadBalancerSku_Stacked_Metallb
//	LoadBalancerSkuUnmanaged
type LoadBalancerSkuInput interface {
	pulumi.Input

	ToLoadBalancerSkuOutput() LoadBalancerSkuOutput
	ToLoadBalancerSkuOutputWithContext(context.Context) LoadBalancerSkuOutput
}

var loadBalancerSkuPtrType = reflect.TypeOf((**LoadBalancerSku)(nil)).Elem()

type LoadBalancerSkuPtrInput interface {
	pulumi.Input

	ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput
	ToLoadBalancerSkuPtrOutputWithContext(context.Context) LoadBalancerSkuPtrOutput
}

type loadBalancerSkuPtr string

func LoadBalancerSkuPtr(v string) LoadBalancerSkuPtrInput {
	return (*loadBalancerSkuPtr)(&v)
}

func (*loadBalancerSkuPtr) ElementType() reflect.Type {
	return loadBalancerSkuPtrType
}

func (in *loadBalancerSkuPtr) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerSkuPtrOutput)
}

func (in *loadBalancerSkuPtr) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerSkuPtrOutput)
}

// Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
type Mode string

const (
	ModeSystem = Mode("System")
	ModeLB     = Mode("LB")
	ModeUser   = Mode("User")
)

func (Mode) ElementType() reflect.Type {
	return reflect.TypeOf((*Mode)(nil)).Elem()
}

func (e Mode) ToModeOutput() ModeOutput {
	return pulumi.ToOutput(e).(ModeOutput)
}

func (e Mode) ToModeOutputWithContext(ctx context.Context) ModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ModeOutput)
}

func (e Mode) ToModePtrOutput() ModePtrOutput {
	return e.ToModePtrOutputWithContext(context.Background())
}

func (e Mode) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return Mode(e).ToModeOutputWithContext(ctx).ToModePtrOutputWithContext(ctx)
}

func (e Mode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Mode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Mode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Mode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ModeOutput struct{ *pulumi.OutputState }

func (ModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Mode)(nil)).Elem()
}

func (o ModeOutput) ToModeOutput() ModeOutput {
	return o
}

func (o ModeOutput) ToModeOutputWithContext(ctx context.Context) ModeOutput {
	return o
}

func (o ModeOutput) ToModePtrOutput() ModePtrOutput {
	return o.ToModePtrOutputWithContext(context.Background())
}

func (o ModeOutput) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Mode) *Mode {
		return &v
	}).(ModePtrOutput)
}

func (o ModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Mode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Mode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ModePtrOutput struct{ *pulumi.OutputState }

func (ModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mode)(nil)).Elem()
}

func (o ModePtrOutput) ToModePtrOutput() ModePtrOutput {
	return o
}

func (o ModePtrOutput) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return o
}

func (o ModePtrOutput) Elem() ModeOutput {
	return o.ApplyT(func(v *Mode) Mode {
		if v != nil {
			return *v
		}
		var ret Mode
		return ret
	}).(ModeOutput)
}

func (o ModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Mode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ModeInput is an input type that accepts values of the Mode enum
// A concrete instance of `ModeInput` can be one of the following:
//
//	ModeSystem
//	ModeLB
//	ModeUser
type ModeInput interface {
	pulumi.Input

	ToModeOutput() ModeOutput
	ToModeOutputWithContext(context.Context) ModeOutput
}

var modePtrType = reflect.TypeOf((**Mode)(nil)).Elem()

type ModePtrInput interface {
	pulumi.Input

	ToModePtrOutput() ModePtrOutput
	ToModePtrOutputWithContext(context.Context) ModePtrOutput
}

type modePtr string

func ModePtr(v string) ModePtrInput {
	return (*modePtr)(&v)
}

func (*modePtr) ElementType() reflect.Type {
	return modePtrType
}

func (in *modePtr) ToModePtrOutput() ModePtrOutput {
	return pulumi.ToOutput(in).(ModePtrOutput)
}

func (in *modePtr) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ModePtrOutput)
}

// NetworkPolicy - Network policy used for building Kubernetes network. Possible values include: 'calico', 'flannel'. Default is 'calico'
type NetworkPolicy string

const (
	NetworkPolicyCalico  = NetworkPolicy("calico")
	NetworkPolicyFlannel = NetworkPolicy("flannel")
)

func (NetworkPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicy)(nil)).Elem()
}

func (e NetworkPolicy) ToNetworkPolicyOutput() NetworkPolicyOutput {
	return pulumi.ToOutput(e).(NetworkPolicyOutput)
}

func (e NetworkPolicy) ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkPolicyOutput)
}

func (e NetworkPolicy) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return e.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (e NetworkPolicy) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return NetworkPolicy(e).ToNetworkPolicyOutputWithContext(ctx).ToNetworkPolicyPtrOutputWithContext(ctx)
}

func (e NetworkPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkPolicyOutput struct{ *pulumi.OutputState }

func (NetworkPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicy)(nil)).Elem()
}

func (o NetworkPolicyOutput) ToNetworkPolicyOutput() NetworkPolicyOutput {
	return o
}

func (o NetworkPolicyOutput) ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput {
	return o
}

func (o NetworkPolicyOutput) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return o.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPolicy) *NetworkPolicy {
		return &v
	}).(NetworkPolicyPtrOutput)
}

func (o NetworkPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkPolicyPtrOutput struct{ *pulumi.OutputState }

func (NetworkPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicy)(nil)).Elem()
}

func (o NetworkPolicyPtrOutput) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return o
}

func (o NetworkPolicyPtrOutput) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return o
}

func (o NetworkPolicyPtrOutput) Elem() NetworkPolicyOutput {
	return o.ApplyT(func(v *NetworkPolicy) NetworkPolicy {
		if v != nil {
			return *v
		}
		var ret NetworkPolicy
		return ret
	}).(NetworkPolicyOutput)
}

func (o NetworkPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkPolicyInput is an input type that accepts values of the NetworkPolicy enum
// A concrete instance of `NetworkPolicyInput` can be one of the following:
//
//	NetworkPolicyCalico
//	NetworkPolicyFlannel
type NetworkPolicyInput interface {
	pulumi.Input

	ToNetworkPolicyOutput() NetworkPolicyOutput
	ToNetworkPolicyOutputWithContext(context.Context) NetworkPolicyOutput
}

var networkPolicyPtrType = reflect.TypeOf((**NetworkPolicy)(nil)).Elem()

type NetworkPolicyPtrInput interface {
	pulumi.Input

	ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput
	ToNetworkPolicyPtrOutputWithContext(context.Context) NetworkPolicyPtrOutput
}

type networkPolicyPtr string

func NetworkPolicyPtr(v string) NetworkPolicyPtrInput {
	return (*networkPolicyPtr)(&v)
}

func (*networkPolicyPtr) ElementType() reflect.Type {
	return networkPolicyPtrType
}

func (in *networkPolicyPtr) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return pulumi.ToOutput(in).(NetworkPolicyPtrOutput)
}

func (in *networkPolicyPtr) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkPolicyPtrOutput)
}

// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
type OsType string

const (
	OsTypeLinux   = OsType("Linux")
	OsTypeWindows = OsType("Windows")
)

func (OsType) ElementType() reflect.Type {
	return reflect.TypeOf((*OsType)(nil)).Elem()
}

func (e OsType) ToOsTypeOutput() OsTypeOutput {
	return pulumi.ToOutput(e).(OsTypeOutput)
}

func (e OsType) ToOsTypeOutputWithContext(ctx context.Context) OsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OsTypeOutput)
}

func (e OsType) ToOsTypePtrOutput() OsTypePtrOutput {
	return e.ToOsTypePtrOutputWithContext(context.Background())
}

func (e OsType) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return OsType(e).ToOsTypeOutputWithContext(ctx).ToOsTypePtrOutputWithContext(ctx)
}

func (e OsType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OsType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OsType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OsType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OsTypeOutput struct{ *pulumi.OutputState }

func (OsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsType)(nil)).Elem()
}

func (o OsTypeOutput) ToOsTypeOutput() OsTypeOutput {
	return o
}

func (o OsTypeOutput) ToOsTypeOutputWithContext(ctx context.Context) OsTypeOutput {
	return o
}

func (o OsTypeOutput) ToOsTypePtrOutput() OsTypePtrOutput {
	return o.ToOsTypePtrOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsType) *OsType {
		return &v
	}).(OsTypePtrOutput)
}

func (o OsTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OsType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OsTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OsType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OsTypePtrOutput struct{ *pulumi.OutputState }

func (OsTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsType)(nil)).Elem()
}

func (o OsTypePtrOutput) ToOsTypePtrOutput() OsTypePtrOutput {
	return o
}

func (o OsTypePtrOutput) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return o
}

func (o OsTypePtrOutput) Elem() OsTypeOutput {
	return o.ApplyT(func(v *OsType) OsType {
		if v != nil {
			return *v
		}
		var ret OsType
		return ret
	}).(OsTypeOutput)
}

func (o OsTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OsTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OsType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OsTypeInput is an input type that accepts values of the OsType enum
// A concrete instance of `OsTypeInput` can be one of the following:
//
//	OsTypeLinux
//	OsTypeWindows
type OsTypeInput interface {
	pulumi.Input

	ToOsTypeOutput() OsTypeOutput
	ToOsTypeOutputWithContext(context.Context) OsTypeOutput
}

var osTypePtrType = reflect.TypeOf((**OsType)(nil)).Elem()

type OsTypePtrInput interface {
	pulumi.Input

	ToOsTypePtrOutput() OsTypePtrOutput
	ToOsTypePtrOutputWithContext(context.Context) OsTypePtrOutput
}

type osTypePtr string

func OsTypePtr(v string) OsTypePtrInput {
	return (*osTypePtr)(&v)
}

func (*osTypePtr) ElementType() reflect.Type {
	return osTypePtrType
}

func (in *osTypePtr) ToOsTypePtrOutput() OsTypePtrOutput {
	return pulumi.ToOutput(in).(OsTypePtrOutput)
}

func (in *osTypePtr) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OsTypePtrOutput)
}

// The type of identity used for the provisioned cluster. The type SystemAssigned, includes a system created identity. The type None means no identity is assigned to the provisioned cluster.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts values of the ResourceIdentityType enum
// A concrete instance of `ResourceIdentityTypeInput` can be one of the following:
//
//	ResourceIdentityTypeNone
//	ResourceIdentityTypeSystemAssigned
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoUpgradeOptionsOutput{})
	pulumi.RegisterOutputType(AutoUpgradeOptionsPtrOutput{})
	pulumi.RegisterOutputType(LicenseTypeOutput{})
	pulumi.RegisterOutputType(LicenseTypePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuPtrOutput{})
	pulumi.RegisterOutputType(ModeOutput{})
	pulumi.RegisterOutputType(ModePtrOutput{})
	pulumi.RegisterOutputType(NetworkPolicyOutput{})
	pulumi.RegisterOutputType(NetworkPolicyPtrOutput{})
	pulumi.RegisterOutputType(OsTypeOutput{})
	pulumi.RegisterOutputType(OsTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
