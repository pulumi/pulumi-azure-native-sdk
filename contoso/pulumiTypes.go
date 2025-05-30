// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package contoso

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Employee properties
type EmployeeProperties struct {
	// Age of employee
	Age *int `pulumi:"age"`
	// City of employee
	City *string `pulumi:"city"`
	// Profile of employee
	Profile *string `pulumi:"profile"`
}

// EmployeePropertiesInput is an input type that accepts EmployeePropertiesArgs and EmployeePropertiesOutput values.
// You can construct a concrete instance of `EmployeePropertiesInput` via:
//
//	EmployeePropertiesArgs{...}
type EmployeePropertiesInput interface {
	pulumi.Input

	ToEmployeePropertiesOutput() EmployeePropertiesOutput
	ToEmployeePropertiesOutputWithContext(context.Context) EmployeePropertiesOutput
}

// Employee properties
type EmployeePropertiesArgs struct {
	// Age of employee
	Age pulumi.IntPtrInput `pulumi:"age"`
	// City of employee
	City pulumi.StringPtrInput `pulumi:"city"`
	// Profile of employee
	Profile pulumi.StringPtrInput `pulumi:"profile"`
}

func (EmployeePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmployeeProperties)(nil)).Elem()
}

func (i EmployeePropertiesArgs) ToEmployeePropertiesOutput() EmployeePropertiesOutput {
	return i.ToEmployeePropertiesOutputWithContext(context.Background())
}

func (i EmployeePropertiesArgs) ToEmployeePropertiesOutputWithContext(ctx context.Context) EmployeePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmployeePropertiesOutput)
}

func (i EmployeePropertiesArgs) ToEmployeePropertiesPtrOutput() EmployeePropertiesPtrOutput {
	return i.ToEmployeePropertiesPtrOutputWithContext(context.Background())
}

func (i EmployeePropertiesArgs) ToEmployeePropertiesPtrOutputWithContext(ctx context.Context) EmployeePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmployeePropertiesOutput).ToEmployeePropertiesPtrOutputWithContext(ctx)
}

// EmployeePropertiesPtrInput is an input type that accepts EmployeePropertiesArgs, EmployeePropertiesPtr and EmployeePropertiesPtrOutput values.
// You can construct a concrete instance of `EmployeePropertiesPtrInput` via:
//
//	        EmployeePropertiesArgs{...}
//
//	or:
//
//	        nil
type EmployeePropertiesPtrInput interface {
	pulumi.Input

	ToEmployeePropertiesPtrOutput() EmployeePropertiesPtrOutput
	ToEmployeePropertiesPtrOutputWithContext(context.Context) EmployeePropertiesPtrOutput
}

type employeePropertiesPtrType EmployeePropertiesArgs

func EmployeePropertiesPtr(v *EmployeePropertiesArgs) EmployeePropertiesPtrInput {
	return (*employeePropertiesPtrType)(v)
}

func (*employeePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmployeeProperties)(nil)).Elem()
}

func (i *employeePropertiesPtrType) ToEmployeePropertiesPtrOutput() EmployeePropertiesPtrOutput {
	return i.ToEmployeePropertiesPtrOutputWithContext(context.Background())
}

func (i *employeePropertiesPtrType) ToEmployeePropertiesPtrOutputWithContext(ctx context.Context) EmployeePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmployeePropertiesPtrOutput)
}

// Employee properties
type EmployeePropertiesOutput struct{ *pulumi.OutputState }

func (EmployeePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmployeeProperties)(nil)).Elem()
}

func (o EmployeePropertiesOutput) ToEmployeePropertiesOutput() EmployeePropertiesOutput {
	return o
}

func (o EmployeePropertiesOutput) ToEmployeePropertiesOutputWithContext(ctx context.Context) EmployeePropertiesOutput {
	return o
}

func (o EmployeePropertiesOutput) ToEmployeePropertiesPtrOutput() EmployeePropertiesPtrOutput {
	return o.ToEmployeePropertiesPtrOutputWithContext(context.Background())
}

func (o EmployeePropertiesOutput) ToEmployeePropertiesPtrOutputWithContext(ctx context.Context) EmployeePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EmployeeProperties) *EmployeeProperties {
		return &v
	}).(EmployeePropertiesPtrOutput)
}

// Age of employee
func (o EmployeePropertiesOutput) Age() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EmployeeProperties) *int { return v.Age }).(pulumi.IntPtrOutput)
}

// City of employee
func (o EmployeePropertiesOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmployeeProperties) *string { return v.City }).(pulumi.StringPtrOutput)
}

// Profile of employee
func (o EmployeePropertiesOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmployeeProperties) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

type EmployeePropertiesPtrOutput struct{ *pulumi.OutputState }

func (EmployeePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmployeeProperties)(nil)).Elem()
}

func (o EmployeePropertiesPtrOutput) ToEmployeePropertiesPtrOutput() EmployeePropertiesPtrOutput {
	return o
}

func (o EmployeePropertiesPtrOutput) ToEmployeePropertiesPtrOutputWithContext(ctx context.Context) EmployeePropertiesPtrOutput {
	return o
}

func (o EmployeePropertiesPtrOutput) Elem() EmployeePropertiesOutput {
	return o.ApplyT(func(v *EmployeeProperties) EmployeeProperties {
		if v != nil {
			return *v
		}
		var ret EmployeeProperties
		return ret
	}).(EmployeePropertiesOutput)
}

// Age of employee
func (o EmployeePropertiesPtrOutput) Age() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EmployeeProperties) *int {
		if v == nil {
			return nil
		}
		return v.Age
	}).(pulumi.IntPtrOutput)
}

// City of employee
func (o EmployeePropertiesPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmployeeProperties) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

// Profile of employee
func (o EmployeePropertiesPtrOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmployeeProperties) *string {
		if v == nil {
			return nil
		}
		return v.Profile
	}).(pulumi.StringPtrOutput)
}

// Employee properties
type EmployeePropertiesResponse struct {
	// Age of employee
	Age *int `pulumi:"age"`
	// City of employee
	City *string `pulumi:"city"`
	// Profile of employee
	Profile *string `pulumi:"profile"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
}

// Employee properties
type EmployeePropertiesResponseOutput struct{ *pulumi.OutputState }

func (EmployeePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmployeePropertiesResponse)(nil)).Elem()
}

func (o EmployeePropertiesResponseOutput) ToEmployeePropertiesResponseOutput() EmployeePropertiesResponseOutput {
	return o
}

func (o EmployeePropertiesResponseOutput) ToEmployeePropertiesResponseOutputWithContext(ctx context.Context) EmployeePropertiesResponseOutput {
	return o
}

// Age of employee
func (o EmployeePropertiesResponseOutput) Age() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EmployeePropertiesResponse) *int { return v.Age }).(pulumi.IntPtrOutput)
}

// City of employee
func (o EmployeePropertiesResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmployeePropertiesResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

// Profile of employee
func (o EmployeePropertiesResponseOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmployeePropertiesResponse) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

// The status of the last operation.
func (o EmployeePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EmployeePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(EmployeePropertiesOutput{})
	pulumi.RegisterOutputType(EmployeePropertiesPtrOutput{})
	pulumi.RegisterOutputType(EmployeePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
