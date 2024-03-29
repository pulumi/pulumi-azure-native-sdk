// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The contact detail class.
type ContactDetail struct {
	// The e-mail address of the contact.
	Email *string `pulumi:"email"`
	// The phone number of the contact.
	Phone *string `pulumi:"phone"`
	// The role of the contact.
	Role *string `pulumi:"role"`
}

// ContactDetailInput is an input type that accepts ContactDetailArgs and ContactDetailOutput values.
// You can construct a concrete instance of `ContactDetailInput` via:
//
//	ContactDetailArgs{...}
type ContactDetailInput interface {
	pulumi.Input

	ToContactDetailOutput() ContactDetailOutput
	ToContactDetailOutputWithContext(context.Context) ContactDetailOutput
}

// The contact detail class.
type ContactDetailArgs struct {
	// The e-mail address of the contact.
	Email pulumi.StringPtrInput `pulumi:"email"`
	// The phone number of the contact.
	Phone pulumi.StringPtrInput `pulumi:"phone"`
	// The role of the contact.
	Role pulumi.StringPtrInput `pulumi:"role"`
}

func (ContactDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetail)(nil)).Elem()
}

func (i ContactDetailArgs) ToContactDetailOutput() ContactDetailOutput {
	return i.ToContactDetailOutputWithContext(context.Background())
}

func (i ContactDetailArgs) ToContactDetailOutputWithContext(ctx context.Context) ContactDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailOutput)
}

// ContactDetailArrayInput is an input type that accepts ContactDetailArray and ContactDetailArrayOutput values.
// You can construct a concrete instance of `ContactDetailArrayInput` via:
//
//	ContactDetailArray{ ContactDetailArgs{...} }
type ContactDetailArrayInput interface {
	pulumi.Input

	ToContactDetailArrayOutput() ContactDetailArrayOutput
	ToContactDetailArrayOutputWithContext(context.Context) ContactDetailArrayOutput
}

type ContactDetailArray []ContactDetailInput

func (ContactDetailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetail)(nil)).Elem()
}

func (i ContactDetailArray) ToContactDetailArrayOutput() ContactDetailArrayOutput {
	return i.ToContactDetailArrayOutputWithContext(context.Background())
}

func (i ContactDetailArray) ToContactDetailArrayOutputWithContext(ctx context.Context) ContactDetailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailArrayOutput)
}

// The contact detail class.
type ContactDetailOutput struct{ *pulumi.OutputState }

func (ContactDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetail)(nil)).Elem()
}

func (o ContactDetailOutput) ToContactDetailOutput() ContactDetailOutput {
	return o
}

func (o ContactDetailOutput) ToContactDetailOutputWithContext(ctx context.Context) ContactDetailOutput {
	return o
}

// The e-mail address of the contact.
func (o ContactDetailOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetail) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The phone number of the contact.
func (o ContactDetailOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetail) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

// The role of the contact.
func (o ContactDetailOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetail) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type ContactDetailArrayOutput struct{ *pulumi.OutputState }

func (ContactDetailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetail)(nil)).Elem()
}

func (o ContactDetailArrayOutput) ToContactDetailArrayOutput() ContactDetailArrayOutput {
	return o
}

func (o ContactDetailArrayOutput) ToContactDetailArrayOutputWithContext(ctx context.Context) ContactDetailArrayOutput {
	return o
}

func (o ContactDetailArrayOutput) Index(i pulumi.IntInput) ContactDetailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactDetail {
		return vs[0].([]ContactDetail)[vs[1].(int)]
	}).(ContactDetailOutput)
}

// The contact detail class.
type ContactDetailResponse struct {
	// The e-mail address of the contact.
	Email *string `pulumi:"email"`
	// The phone number of the contact.
	Phone *string `pulumi:"phone"`
	// The role of the contact.
	Role *string `pulumi:"role"`
}

// The contact detail class.
type ContactDetailResponseOutput struct{ *pulumi.OutputState }

func (ContactDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailResponse)(nil)).Elem()
}

func (o ContactDetailResponseOutput) ToContactDetailResponseOutput() ContactDetailResponseOutput {
	return o
}

func (o ContactDetailResponseOutput) ToContactDetailResponseOutputWithContext(ctx context.Context) ContactDetailResponseOutput {
	return o
}

// The e-mail address of the contact.
func (o ContactDetailResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The phone number of the contact.
func (o ContactDetailResponseOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailResponse) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

// The role of the contact.
func (o ContactDetailResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type ContactDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ContactDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContactDetailResponse)(nil)).Elem()
}

func (o ContactDetailResponseArrayOutput) ToContactDetailResponseArrayOutput() ContactDetailResponseArrayOutput {
	return o
}

func (o ContactDetailResponseArrayOutput) ToContactDetailResponseArrayOutputWithContext(ctx context.Context) ContactDetailResponseArrayOutput {
	return o
}

func (o ContactDetailResponseArrayOutput) Index(i pulumi.IntInput) ContactDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContactDetailResponse {
		return vs[0].([]ContactDetailResponse)[vs[1].(int)]
	}).(ContactDetailResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ContactDetailOutput{})
	pulumi.RegisterOutputType(ContactDetailArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailResponseOutput{})
	pulumi.RegisterOutputType(ContactDetailResponseArrayOutput{})
}
