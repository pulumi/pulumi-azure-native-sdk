// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Site properties
type SiteProperties struct {
	// AddressResource ArmId of Site resource
	AddressResourceId *string `pulumi:"addressResourceId"`
	// Description of Site resource
	Description *string `pulumi:"description"`
	// displayName of Site resource
	DisplayName *string `pulumi:"displayName"`
}

// SitePropertiesInput is an input type that accepts SitePropertiesArgs and SitePropertiesOutput values.
// You can construct a concrete instance of `SitePropertiesInput` via:
//
//	SitePropertiesArgs{...}
type SitePropertiesInput interface {
	pulumi.Input

	ToSitePropertiesOutput() SitePropertiesOutput
	ToSitePropertiesOutputWithContext(context.Context) SitePropertiesOutput
}

// Site properties
type SitePropertiesArgs struct {
	// AddressResource ArmId of Site resource
	AddressResourceId pulumi.StringPtrInput `pulumi:"addressResourceId"`
	// Description of Site resource
	Description pulumi.StringPtrInput `pulumi:"description"`
	// displayName of Site resource
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
}

func (SitePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProperties)(nil)).Elem()
}

func (i SitePropertiesArgs) ToSitePropertiesOutput() SitePropertiesOutput {
	return i.ToSitePropertiesOutputWithContext(context.Background())
}

func (i SitePropertiesArgs) ToSitePropertiesOutputWithContext(ctx context.Context) SitePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitePropertiesOutput)
}

func (i SitePropertiesArgs) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return i.ToSitePropertiesPtrOutputWithContext(context.Background())
}

func (i SitePropertiesArgs) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitePropertiesOutput).ToSitePropertiesPtrOutputWithContext(ctx)
}

// SitePropertiesPtrInput is an input type that accepts SitePropertiesArgs, SitePropertiesPtr and SitePropertiesPtrOutput values.
// You can construct a concrete instance of `SitePropertiesPtrInput` via:
//
//	        SitePropertiesArgs{...}
//
//	or:
//
//	        nil
type SitePropertiesPtrInput interface {
	pulumi.Input

	ToSitePropertiesPtrOutput() SitePropertiesPtrOutput
	ToSitePropertiesPtrOutputWithContext(context.Context) SitePropertiesPtrOutput
}

type sitePropertiesPtrType SitePropertiesArgs

func SitePropertiesPtr(v *SitePropertiesArgs) SitePropertiesPtrInput {
	return (*sitePropertiesPtrType)(v)
}

func (*sitePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProperties)(nil)).Elem()
}

func (i *sitePropertiesPtrType) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return i.ToSitePropertiesPtrOutputWithContext(context.Background())
}

func (i *sitePropertiesPtrType) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitePropertiesPtrOutput)
}

// Site properties
type SitePropertiesOutput struct{ *pulumi.OutputState }

func (SitePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProperties)(nil)).Elem()
}

func (o SitePropertiesOutput) ToSitePropertiesOutput() SitePropertiesOutput {
	return o
}

func (o SitePropertiesOutput) ToSitePropertiesOutputWithContext(ctx context.Context) SitePropertiesOutput {
	return o
}

func (o SitePropertiesOutput) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return o.ToSitePropertiesPtrOutputWithContext(context.Background())
}

func (o SitePropertiesOutput) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteProperties) *SiteProperties {
		return &v
	}).(SitePropertiesPtrOutput)
}

// AddressResource ArmId of Site resource
func (o SitePropertiesOutput) AddressResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProperties) *string { return v.AddressResourceId }).(pulumi.StringPtrOutput)
}

// Description of Site resource
func (o SitePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// displayName of Site resource
func (o SitePropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type SitePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SitePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProperties)(nil)).Elem()
}

func (o SitePropertiesPtrOutput) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return o
}

func (o SitePropertiesPtrOutput) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return o
}

func (o SitePropertiesPtrOutput) Elem() SitePropertiesOutput {
	return o.ApplyT(func(v *SiteProperties) SiteProperties {
		if v != nil {
			return *v
		}
		var ret SiteProperties
		return ret
	}).(SitePropertiesOutput)
}

// AddressResource ArmId of Site resource
func (o SitePropertiesPtrOutput) AddressResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *string {
		if v == nil {
			return nil
		}
		return v.AddressResourceId
	}).(pulumi.StringPtrOutput)
}

// Description of Site resource
func (o SitePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// displayName of Site resource
func (o SitePropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// Site properties
type SitePropertiesResponse struct {
	// AddressResource ArmId of Site resource
	AddressResourceId *string `pulumi:"addressResourceId"`
	// Description of Site resource
	Description *string `pulumi:"description"`
	// displayName of Site resource
	DisplayName *string `pulumi:"displayName"`
	// Provisioning state of last operation
	ProvisioningState string `pulumi:"provisioningState"`
}

// Site properties
type SitePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SitePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SitePropertiesResponse)(nil)).Elem()
}

func (o SitePropertiesResponseOutput) ToSitePropertiesResponseOutput() SitePropertiesResponseOutput {
	return o
}

func (o SitePropertiesResponseOutput) ToSitePropertiesResponseOutputWithContext(ctx context.Context) SitePropertiesResponseOutput {
	return o
}

// AddressResource ArmId of Site resource
func (o SitePropertiesResponseOutput) AddressResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *string { return v.AddressResourceId }).(pulumi.StringPtrOutput)
}

// Description of Site resource
func (o SitePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// displayName of Site resource
func (o SitePropertiesResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Provisioning state of last operation
func (o SitePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SitePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(SitePropertiesOutput{})
	pulumi.RegisterOutputType(SitePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SitePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
