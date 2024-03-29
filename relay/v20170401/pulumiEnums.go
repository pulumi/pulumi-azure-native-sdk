// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessRights string

const (
	AccessRightsManage = AccessRights("Manage")
	AccessRightsSend   = AccessRights("Send")
	AccessRightsListen = AccessRights("Listen")
)

func (AccessRights) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (e AccessRights) ToAccessRightsOutput() AccessRightsOutput {
	return pulumi.ToOutput(e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return e.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return AccessRights(e).ToAccessRightsOutputWithContext(ctx).ToAccessRightsPtrOutputWithContext(ctx)
}

func (e AccessRights) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessRightsOutput struct{ *pulumi.OutputState }

func (AccessRightsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (o AccessRightsOutput) ToAccessRightsOutput() AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRights) *AccessRights {
		return &v
	}).(AccessRightsPtrOutput)
}

func (o AccessRightsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessRightsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessRightsPtrOutput struct{ *pulumi.OutputState }

func (AccessRightsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRights)(nil)).Elem()
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) Elem() AccessRightsOutput {
	return o.ApplyT(func(v *AccessRights) AccessRights {
		if v != nil {
			return *v
		}
		var ret AccessRights
		return ret
	}).(AccessRightsOutput)
}

func (o AccessRightsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessRights) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessRightsInput is an input type that accepts values of the AccessRights enum
// A concrete instance of `AccessRightsInput` can be one of the following:
//
//	AccessRightsManage
//	AccessRightsSend
//	AccessRightsListen
type AccessRightsInput interface {
	pulumi.Input

	ToAccessRightsOutput() AccessRightsOutput
	ToAccessRightsOutputWithContext(context.Context) AccessRightsOutput
}

var accessRightsPtrType = reflect.TypeOf((**AccessRights)(nil)).Elem()

type AccessRightsPtrInput interface {
	pulumi.Input

	ToAccessRightsPtrOutput() AccessRightsPtrOutput
	ToAccessRightsPtrOutputWithContext(context.Context) AccessRightsPtrOutput
}

type accessRightsPtr string

func AccessRightsPtr(v string) AccessRightsPtrInput {
	return (*accessRightsPtr)(&v)
}

func (*accessRightsPtr) ElementType() reflect.Type {
	return accessRightsPtrType
}

func (in *accessRightsPtr) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return pulumi.ToOutput(in).(AccessRightsPtrOutput)
}

func (in *accessRightsPtr) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessRightsPtrOutput)
}

// AccessRightsArrayInput is an input type that accepts AccessRightsArray and AccessRightsArrayOutput values.
// You can construct a concrete instance of `AccessRightsArrayInput` via:
//
//	AccessRightsArray{ AccessRightsArgs{...} }
type AccessRightsArrayInput interface {
	pulumi.Input

	ToAccessRightsArrayOutput() AccessRightsArrayOutput
	ToAccessRightsArrayOutputWithContext(context.Context) AccessRightsArrayOutput
}

type AccessRightsArray []AccessRights

func (AccessRightsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessRights)(nil)).Elem()
}

func (i AccessRightsArray) ToAccessRightsArrayOutput() AccessRightsArrayOutput {
	return i.ToAccessRightsArrayOutputWithContext(context.Background())
}

func (i AccessRightsArray) ToAccessRightsArrayOutputWithContext(ctx context.Context) AccessRightsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRightsArrayOutput)
}

type AccessRightsArrayOutput struct{ *pulumi.OutputState }

func (AccessRightsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessRights)(nil)).Elem()
}

func (o AccessRightsArrayOutput) ToAccessRightsArrayOutput() AccessRightsArrayOutput {
	return o
}

func (o AccessRightsArrayOutput) ToAccessRightsArrayOutputWithContext(ctx context.Context) AccessRightsArrayOutput {
	return o
}

func (o AccessRightsArrayOutput) Index(i pulumi.IntInput) AccessRightsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessRights {
		return vs[0].([]AccessRights)[vs[1].(int)]
	}).(AccessRightsOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessRightsOutput{})
	pulumi.RegisterOutputType(AccessRightsPtrOutput{})
	pulumi.RegisterOutputType(AccessRightsArrayOutput{})
}
