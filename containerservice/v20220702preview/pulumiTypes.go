// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220702preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The credential result response.
type FleetCredentialResultResponse struct {
	// The name of the credential.
	Name string `pulumi:"name"`
	// Base64-encoded Kubernetes configuration file.
	Value string `pulumi:"value"`
}

// The credential result response.
type FleetCredentialResultResponseOutput struct{ *pulumi.OutputState }

func (FleetCredentialResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FleetCredentialResultResponse)(nil)).Elem()
}

func (o FleetCredentialResultResponseOutput) ToFleetCredentialResultResponseOutput() FleetCredentialResultResponseOutput {
	return o
}

func (o FleetCredentialResultResponseOutput) ToFleetCredentialResultResponseOutputWithContext(ctx context.Context) FleetCredentialResultResponseOutput {
	return o
}

// The name of the credential.
func (o FleetCredentialResultResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v FleetCredentialResultResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Base64-encoded Kubernetes configuration file.
func (o FleetCredentialResultResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v FleetCredentialResultResponse) string { return v.Value }).(pulumi.StringOutput)
}

type FleetCredentialResultResponseArrayOutput struct{ *pulumi.OutputState }

func (FleetCredentialResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FleetCredentialResultResponse)(nil)).Elem()
}

func (o FleetCredentialResultResponseArrayOutput) ToFleetCredentialResultResponseArrayOutput() FleetCredentialResultResponseArrayOutput {
	return o
}

func (o FleetCredentialResultResponseArrayOutput) ToFleetCredentialResultResponseArrayOutputWithContext(ctx context.Context) FleetCredentialResultResponseArrayOutput {
	return o
}

func (o FleetCredentialResultResponseArrayOutput) Index(i pulumi.IntInput) FleetCredentialResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FleetCredentialResultResponse {
		return vs[0].([]FleetCredentialResultResponse)[vs[1].(int)]
	}).(FleetCredentialResultResponseOutput)
}

// The FleetHubProfile configures the fleet hub.
type FleetHubProfile struct {
	// DNS prefix used to create the FQDN for the Fleet hub.
	DnsPrefix *string `pulumi:"dnsPrefix"`
}

// FleetHubProfileInput is an input type that accepts FleetHubProfileArgs and FleetHubProfileOutput values.
// You can construct a concrete instance of `FleetHubProfileInput` via:
//
//	FleetHubProfileArgs{...}
type FleetHubProfileInput interface {
	pulumi.Input

	ToFleetHubProfileOutput() FleetHubProfileOutput
	ToFleetHubProfileOutputWithContext(context.Context) FleetHubProfileOutput
}

// The FleetHubProfile configures the fleet hub.
type FleetHubProfileArgs struct {
	// DNS prefix used to create the FQDN for the Fleet hub.
	DnsPrefix pulumi.StringPtrInput `pulumi:"dnsPrefix"`
}

func (FleetHubProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FleetHubProfile)(nil)).Elem()
}

func (i FleetHubProfileArgs) ToFleetHubProfileOutput() FleetHubProfileOutput {
	return i.ToFleetHubProfileOutputWithContext(context.Background())
}

func (i FleetHubProfileArgs) ToFleetHubProfileOutputWithContext(ctx context.Context) FleetHubProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetHubProfileOutput)
}

func (i FleetHubProfileArgs) ToFleetHubProfilePtrOutput() FleetHubProfilePtrOutput {
	return i.ToFleetHubProfilePtrOutputWithContext(context.Background())
}

func (i FleetHubProfileArgs) ToFleetHubProfilePtrOutputWithContext(ctx context.Context) FleetHubProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetHubProfileOutput).ToFleetHubProfilePtrOutputWithContext(ctx)
}

// FleetHubProfilePtrInput is an input type that accepts FleetHubProfileArgs, FleetHubProfilePtr and FleetHubProfilePtrOutput values.
// You can construct a concrete instance of `FleetHubProfilePtrInput` via:
//
//	        FleetHubProfileArgs{...}
//
//	or:
//
//	        nil
type FleetHubProfilePtrInput interface {
	pulumi.Input

	ToFleetHubProfilePtrOutput() FleetHubProfilePtrOutput
	ToFleetHubProfilePtrOutputWithContext(context.Context) FleetHubProfilePtrOutput
}

type fleetHubProfilePtrType FleetHubProfileArgs

func FleetHubProfilePtr(v *FleetHubProfileArgs) FleetHubProfilePtrInput {
	return (*fleetHubProfilePtrType)(v)
}

func (*fleetHubProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetHubProfile)(nil)).Elem()
}

func (i *fleetHubProfilePtrType) ToFleetHubProfilePtrOutput() FleetHubProfilePtrOutput {
	return i.ToFleetHubProfilePtrOutputWithContext(context.Background())
}

func (i *fleetHubProfilePtrType) ToFleetHubProfilePtrOutputWithContext(ctx context.Context) FleetHubProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetHubProfilePtrOutput)
}

// The FleetHubProfile configures the fleet hub.
type FleetHubProfileOutput struct{ *pulumi.OutputState }

func (FleetHubProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FleetHubProfile)(nil)).Elem()
}

func (o FleetHubProfileOutput) ToFleetHubProfileOutput() FleetHubProfileOutput {
	return o
}

func (o FleetHubProfileOutput) ToFleetHubProfileOutputWithContext(ctx context.Context) FleetHubProfileOutput {
	return o
}

func (o FleetHubProfileOutput) ToFleetHubProfilePtrOutput() FleetHubProfilePtrOutput {
	return o.ToFleetHubProfilePtrOutputWithContext(context.Background())
}

func (o FleetHubProfileOutput) ToFleetHubProfilePtrOutputWithContext(ctx context.Context) FleetHubProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FleetHubProfile) *FleetHubProfile {
		return &v
	}).(FleetHubProfilePtrOutput)
}

// DNS prefix used to create the FQDN for the Fleet hub.
func (o FleetHubProfileOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FleetHubProfile) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

type FleetHubProfilePtrOutput struct{ *pulumi.OutputState }

func (FleetHubProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetHubProfile)(nil)).Elem()
}

func (o FleetHubProfilePtrOutput) ToFleetHubProfilePtrOutput() FleetHubProfilePtrOutput {
	return o
}

func (o FleetHubProfilePtrOutput) ToFleetHubProfilePtrOutputWithContext(ctx context.Context) FleetHubProfilePtrOutput {
	return o
}

func (o FleetHubProfilePtrOutput) Elem() FleetHubProfileOutput {
	return o.ApplyT(func(v *FleetHubProfile) FleetHubProfile {
		if v != nil {
			return *v
		}
		var ret FleetHubProfile
		return ret
	}).(FleetHubProfileOutput)
}

// DNS prefix used to create the FQDN for the Fleet hub.
func (o FleetHubProfilePtrOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetHubProfile) *string {
		if v == nil {
			return nil
		}
		return v.DnsPrefix
	}).(pulumi.StringPtrOutput)
}

// The FleetHubProfile configures the fleet hub.
type FleetHubProfileResponse struct {
	// DNS prefix used to create the FQDN for the Fleet hub.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// The FQDN of the Fleet hub.
	Fqdn string `pulumi:"fqdn"`
	// The Kubernetes version of the Fleet hub.
	KubernetesVersion string `pulumi:"kubernetesVersion"`
}

// The FleetHubProfile configures the fleet hub.
type FleetHubProfileResponseOutput struct{ *pulumi.OutputState }

func (FleetHubProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FleetHubProfileResponse)(nil)).Elem()
}

func (o FleetHubProfileResponseOutput) ToFleetHubProfileResponseOutput() FleetHubProfileResponseOutput {
	return o
}

func (o FleetHubProfileResponseOutput) ToFleetHubProfileResponseOutputWithContext(ctx context.Context) FleetHubProfileResponseOutput {
	return o
}

// DNS prefix used to create the FQDN for the Fleet hub.
func (o FleetHubProfileResponseOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FleetHubProfileResponse) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

// The FQDN of the Fleet hub.
func (o FleetHubProfileResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v FleetHubProfileResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

// The Kubernetes version of the Fleet hub.
func (o FleetHubProfileResponseOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v FleetHubProfileResponse) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

type FleetHubProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (FleetHubProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetHubProfileResponse)(nil)).Elem()
}

func (o FleetHubProfileResponsePtrOutput) ToFleetHubProfileResponsePtrOutput() FleetHubProfileResponsePtrOutput {
	return o
}

func (o FleetHubProfileResponsePtrOutput) ToFleetHubProfileResponsePtrOutputWithContext(ctx context.Context) FleetHubProfileResponsePtrOutput {
	return o
}

func (o FleetHubProfileResponsePtrOutput) Elem() FleetHubProfileResponseOutput {
	return o.ApplyT(func(v *FleetHubProfileResponse) FleetHubProfileResponse {
		if v != nil {
			return *v
		}
		var ret FleetHubProfileResponse
		return ret
	}).(FleetHubProfileResponseOutput)
}

// DNS prefix used to create the FQDN for the Fleet hub.
func (o FleetHubProfileResponsePtrOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetHubProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsPrefix
	}).(pulumi.StringPtrOutput)
}

// The FQDN of the Fleet hub.
func (o FleetHubProfileResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetHubProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Fqdn
	}).(pulumi.StringPtrOutput)
}

// The Kubernetes version of the Fleet hub.
func (o FleetHubProfileResponsePtrOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetHubProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KubernetesVersion
	}).(pulumi.StringPtrOutput)
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

func init() {
	pulumi.RegisterOutputType(FleetCredentialResultResponseOutput{})
	pulumi.RegisterOutputType(FleetCredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(FleetHubProfileOutput{})
	pulumi.RegisterOutputType(FleetHubProfilePtrOutput{})
	pulumi.RegisterOutputType(FleetHubProfileResponseOutput{})
	pulumi.RegisterOutputType(FleetHubProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
