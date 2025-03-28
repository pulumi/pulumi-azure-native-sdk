// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250102

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties with which the alert rule resource was created.
type AlertRuleCreationProperties string

const (
	// The alert rule was created with an action group.
	AlertRuleCreationPropertiesCreatedWithActionGroup = AlertRuleCreationProperties("CreatedWithActionGroup")
	// The alert rule was created with no properties.
	AlertRuleCreationPropertiesNone = AlertRuleCreationProperties("None")
)

func (AlertRuleCreationProperties) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleCreationProperties)(nil)).Elem()
}

func (e AlertRuleCreationProperties) ToAlertRuleCreationPropertiesOutput() AlertRuleCreationPropertiesOutput {
	return pulumi.ToOutput(e).(AlertRuleCreationPropertiesOutput)
}

func (e AlertRuleCreationProperties) ToAlertRuleCreationPropertiesOutputWithContext(ctx context.Context) AlertRuleCreationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AlertRuleCreationPropertiesOutput)
}

func (e AlertRuleCreationProperties) ToAlertRuleCreationPropertiesPtrOutput() AlertRuleCreationPropertiesPtrOutput {
	return e.ToAlertRuleCreationPropertiesPtrOutputWithContext(context.Background())
}

func (e AlertRuleCreationProperties) ToAlertRuleCreationPropertiesPtrOutputWithContext(ctx context.Context) AlertRuleCreationPropertiesPtrOutput {
	return AlertRuleCreationProperties(e).ToAlertRuleCreationPropertiesOutputWithContext(ctx).ToAlertRuleCreationPropertiesPtrOutputWithContext(ctx)
}

func (e AlertRuleCreationProperties) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleCreationProperties) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleCreationProperties) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AlertRuleCreationProperties) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AlertRuleCreationPropertiesOutput struct{ *pulumi.OutputState }

func (AlertRuleCreationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleCreationProperties)(nil)).Elem()
}

func (o AlertRuleCreationPropertiesOutput) ToAlertRuleCreationPropertiesOutput() AlertRuleCreationPropertiesOutput {
	return o
}

func (o AlertRuleCreationPropertiesOutput) ToAlertRuleCreationPropertiesOutputWithContext(ctx context.Context) AlertRuleCreationPropertiesOutput {
	return o
}

func (o AlertRuleCreationPropertiesOutput) ToAlertRuleCreationPropertiesPtrOutput() AlertRuleCreationPropertiesPtrOutput {
	return o.ToAlertRuleCreationPropertiesPtrOutputWithContext(context.Background())
}

func (o AlertRuleCreationPropertiesOutput) ToAlertRuleCreationPropertiesPtrOutputWithContext(ctx context.Context) AlertRuleCreationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertRuleCreationProperties) *AlertRuleCreationProperties {
		return &v
	}).(AlertRuleCreationPropertiesPtrOutput)
}

func (o AlertRuleCreationPropertiesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AlertRuleCreationPropertiesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertRuleCreationProperties) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AlertRuleCreationPropertiesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertRuleCreationPropertiesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertRuleCreationProperties) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AlertRuleCreationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AlertRuleCreationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleCreationProperties)(nil)).Elem()
}

func (o AlertRuleCreationPropertiesPtrOutput) ToAlertRuleCreationPropertiesPtrOutput() AlertRuleCreationPropertiesPtrOutput {
	return o
}

func (o AlertRuleCreationPropertiesPtrOutput) ToAlertRuleCreationPropertiesPtrOutputWithContext(ctx context.Context) AlertRuleCreationPropertiesPtrOutput {
	return o
}

func (o AlertRuleCreationPropertiesPtrOutput) Elem() AlertRuleCreationPropertiesOutput {
	return o.ApplyT(func(v *AlertRuleCreationProperties) AlertRuleCreationProperties {
		if v != nil {
			return *v
		}
		var ret AlertRuleCreationProperties
		return ret
	}).(AlertRuleCreationPropertiesOutput)
}

func (o AlertRuleCreationPropertiesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertRuleCreationPropertiesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AlertRuleCreationProperties) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AlertRuleCreationPropertiesInput is an input type that accepts values of the AlertRuleCreationProperties enum
// A concrete instance of `AlertRuleCreationPropertiesInput` can be one of the following:
//
//	AlertRuleCreationPropertiesCreatedWithActionGroup
//	AlertRuleCreationPropertiesNone
type AlertRuleCreationPropertiesInput interface {
	pulumi.Input

	ToAlertRuleCreationPropertiesOutput() AlertRuleCreationPropertiesOutput
	ToAlertRuleCreationPropertiesOutputWithContext(context.Context) AlertRuleCreationPropertiesOutput
}

var alertRuleCreationPropertiesPtrType = reflect.TypeOf((**AlertRuleCreationProperties)(nil)).Elem()

type AlertRuleCreationPropertiesPtrInput interface {
	pulumi.Input

	ToAlertRuleCreationPropertiesPtrOutput() AlertRuleCreationPropertiesPtrOutput
	ToAlertRuleCreationPropertiesPtrOutputWithContext(context.Context) AlertRuleCreationPropertiesPtrOutput
}

type alertRuleCreationPropertiesPtr string

func AlertRuleCreationPropertiesPtr(v string) AlertRuleCreationPropertiesPtrInput {
	return (*alertRuleCreationPropertiesPtr)(&v)
}

func (*alertRuleCreationPropertiesPtr) ElementType() reflect.Type {
	return alertRuleCreationPropertiesPtrType
}

func (in *alertRuleCreationPropertiesPtr) ToAlertRuleCreationPropertiesPtrOutput() AlertRuleCreationPropertiesPtrOutput {
	return pulumi.ToOutput(in).(AlertRuleCreationPropertiesPtrOutput)
}

func (in *alertRuleCreationPropertiesPtr) ToAlertRuleCreationPropertiesPtrOutputWithContext(ctx context.Context) AlertRuleCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AlertRuleCreationPropertiesPtrOutput)
}

// The type of a Kusto offering.
type KustoOfferingType string

const (
	// The Azure Data Explorer cluster Kusto offering.
	KustoOfferingTypeAdx = KustoOfferingType("adx")
	// The free Azure Data Explorer cluster Kusto offering.
	KustoOfferingTypeFree = KustoOfferingType("free")
	// The Fabric Real-Time Analytics Kusto offering.
	KustoOfferingTypeFabric = KustoOfferingType("fabric")
)

func (KustoOfferingType) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoOfferingType)(nil)).Elem()
}

func (e KustoOfferingType) ToKustoOfferingTypeOutput() KustoOfferingTypeOutput {
	return pulumi.ToOutput(e).(KustoOfferingTypeOutput)
}

func (e KustoOfferingType) ToKustoOfferingTypeOutputWithContext(ctx context.Context) KustoOfferingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KustoOfferingTypeOutput)
}

func (e KustoOfferingType) ToKustoOfferingTypePtrOutput() KustoOfferingTypePtrOutput {
	return e.ToKustoOfferingTypePtrOutputWithContext(context.Background())
}

func (e KustoOfferingType) ToKustoOfferingTypePtrOutputWithContext(ctx context.Context) KustoOfferingTypePtrOutput {
	return KustoOfferingType(e).ToKustoOfferingTypeOutputWithContext(ctx).ToKustoOfferingTypePtrOutputWithContext(ctx)
}

func (e KustoOfferingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KustoOfferingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KustoOfferingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KustoOfferingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KustoOfferingTypeOutput struct{ *pulumi.OutputState }

func (KustoOfferingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoOfferingType)(nil)).Elem()
}

func (o KustoOfferingTypeOutput) ToKustoOfferingTypeOutput() KustoOfferingTypeOutput {
	return o
}

func (o KustoOfferingTypeOutput) ToKustoOfferingTypeOutputWithContext(ctx context.Context) KustoOfferingTypeOutput {
	return o
}

func (o KustoOfferingTypeOutput) ToKustoOfferingTypePtrOutput() KustoOfferingTypePtrOutput {
	return o.ToKustoOfferingTypePtrOutputWithContext(context.Background())
}

func (o KustoOfferingTypeOutput) ToKustoOfferingTypePtrOutputWithContext(ctx context.Context) KustoOfferingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KustoOfferingType) *KustoOfferingType {
		return &v
	}).(KustoOfferingTypePtrOutput)
}

func (o KustoOfferingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KustoOfferingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KustoOfferingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KustoOfferingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KustoOfferingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KustoOfferingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KustoOfferingTypePtrOutput struct{ *pulumi.OutputState }

func (KustoOfferingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoOfferingType)(nil)).Elem()
}

func (o KustoOfferingTypePtrOutput) ToKustoOfferingTypePtrOutput() KustoOfferingTypePtrOutput {
	return o
}

func (o KustoOfferingTypePtrOutput) ToKustoOfferingTypePtrOutputWithContext(ctx context.Context) KustoOfferingTypePtrOutput {
	return o
}

func (o KustoOfferingTypePtrOutput) Elem() KustoOfferingTypeOutput {
	return o.ApplyT(func(v *KustoOfferingType) KustoOfferingType {
		if v != nil {
			return *v
		}
		var ret KustoOfferingType
		return ret
	}).(KustoOfferingTypeOutput)
}

func (o KustoOfferingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KustoOfferingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KustoOfferingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KustoOfferingTypeInput is an input type that accepts values of the KustoOfferingType enum
// A concrete instance of `KustoOfferingTypeInput` can be one of the following:
//
//	KustoOfferingTypeAdx
//	KustoOfferingTypeFree
//	KustoOfferingTypeFabric
type KustoOfferingTypeInput interface {
	pulumi.Input

	ToKustoOfferingTypeOutput() KustoOfferingTypeOutput
	ToKustoOfferingTypeOutputWithContext(context.Context) KustoOfferingTypeOutput
}

var kustoOfferingTypePtrType = reflect.TypeOf((**KustoOfferingType)(nil)).Elem()

type KustoOfferingTypePtrInput interface {
	pulumi.Input

	ToKustoOfferingTypePtrOutput() KustoOfferingTypePtrOutput
	ToKustoOfferingTypePtrOutputWithContext(context.Context) KustoOfferingTypePtrOutput
}

type kustoOfferingTypePtr string

func KustoOfferingTypePtr(v string) KustoOfferingTypePtrInput {
	return (*kustoOfferingTypePtr)(&v)
}

func (*kustoOfferingTypePtr) ElementType() reflect.Type {
	return kustoOfferingTypePtrType
}

func (in *kustoOfferingTypePtr) ToKustoOfferingTypePtrOutput() KustoOfferingTypePtrOutput {
	return pulumi.ToOutput(in).(KustoOfferingTypePtrOutput)
}

func (in *kustoOfferingTypePtr) ToKustoOfferingTypePtrOutputWithContext(ctx context.Context) KustoOfferingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KustoOfferingTypePtrOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ManagedServiceIdentityTypeInput is an input type that accepts values of the ManagedServiceIdentityType enum
// A concrete instance of `ManagedServiceIdentityTypeInput` can be one of the following:
//
//	ManagedServiceIdentityTypeNone
//	ManagedServiceIdentityTypeSystemAssigned
//	ManagedServiceIdentityTypeUserAssigned
//	ManagedServiceIdentityType_SystemAssigned_UserAssigned
type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

// The type of authentication to use when connecting to a target.
type TargetAuthenticationType string

const (
	// The Azure Active Directory authentication.
	TargetAuthenticationTypeAad = TargetAuthenticationType("Aad")
	// The SQL password authentication.
	TargetAuthenticationTypeSql = TargetAuthenticationType("Sql")
)

func (TargetAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetAuthenticationType)(nil)).Elem()
}

func (e TargetAuthenticationType) ToTargetAuthenticationTypeOutput() TargetAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(TargetAuthenticationTypeOutput)
}

func (e TargetAuthenticationType) ToTargetAuthenticationTypeOutputWithContext(ctx context.Context) TargetAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TargetAuthenticationTypeOutput)
}

func (e TargetAuthenticationType) ToTargetAuthenticationTypePtrOutput() TargetAuthenticationTypePtrOutput {
	return e.ToTargetAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e TargetAuthenticationType) ToTargetAuthenticationTypePtrOutputWithContext(ctx context.Context) TargetAuthenticationTypePtrOutput {
	return TargetAuthenticationType(e).ToTargetAuthenticationTypeOutputWithContext(ctx).ToTargetAuthenticationTypePtrOutputWithContext(ctx)
}

func (e TargetAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TargetAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TargetAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TargetAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TargetAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (TargetAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetAuthenticationType)(nil)).Elem()
}

func (o TargetAuthenticationTypeOutput) ToTargetAuthenticationTypeOutput() TargetAuthenticationTypeOutput {
	return o
}

func (o TargetAuthenticationTypeOutput) ToTargetAuthenticationTypeOutputWithContext(ctx context.Context) TargetAuthenticationTypeOutput {
	return o
}

func (o TargetAuthenticationTypeOutput) ToTargetAuthenticationTypePtrOutput() TargetAuthenticationTypePtrOutput {
	return o.ToTargetAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o TargetAuthenticationTypeOutput) ToTargetAuthenticationTypePtrOutputWithContext(ctx context.Context) TargetAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetAuthenticationType) *TargetAuthenticationType {
		return &v
	}).(TargetAuthenticationTypePtrOutput)
}

func (o TargetAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TargetAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TargetAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TargetAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TargetAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TargetAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TargetAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (TargetAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetAuthenticationType)(nil)).Elem()
}

func (o TargetAuthenticationTypePtrOutput) ToTargetAuthenticationTypePtrOutput() TargetAuthenticationTypePtrOutput {
	return o
}

func (o TargetAuthenticationTypePtrOutput) ToTargetAuthenticationTypePtrOutputWithContext(ctx context.Context) TargetAuthenticationTypePtrOutput {
	return o
}

func (o TargetAuthenticationTypePtrOutput) Elem() TargetAuthenticationTypeOutput {
	return o.ApplyT(func(v *TargetAuthenticationType) TargetAuthenticationType {
		if v != nil {
			return *v
		}
		var ret TargetAuthenticationType
		return ret
	}).(TargetAuthenticationTypeOutput)
}

func (o TargetAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TargetAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TargetAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TargetAuthenticationTypeInput is an input type that accepts values of the TargetAuthenticationType enum
// A concrete instance of `TargetAuthenticationTypeInput` can be one of the following:
//
//	TargetAuthenticationTypeAad
//	TargetAuthenticationTypeSql
type TargetAuthenticationTypeInput interface {
	pulumi.Input

	ToTargetAuthenticationTypeOutput() TargetAuthenticationTypeOutput
	ToTargetAuthenticationTypeOutputWithContext(context.Context) TargetAuthenticationTypeOutput
}

var targetAuthenticationTypePtrType = reflect.TypeOf((**TargetAuthenticationType)(nil)).Elem()

type TargetAuthenticationTypePtrInput interface {
	pulumi.Input

	ToTargetAuthenticationTypePtrOutput() TargetAuthenticationTypePtrOutput
	ToTargetAuthenticationTypePtrOutputWithContext(context.Context) TargetAuthenticationTypePtrOutput
}

type targetAuthenticationTypePtr string

func TargetAuthenticationTypePtr(v string) TargetAuthenticationTypePtrInput {
	return (*targetAuthenticationTypePtr)(&v)
}

func (*targetAuthenticationTypePtr) ElementType() reflect.Type {
	return targetAuthenticationTypePtrType
}

func (in *targetAuthenticationTypePtr) ToTargetAuthenticationTypePtrOutput() TargetAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(TargetAuthenticationTypePtrOutput)
}

func (in *targetAuthenticationTypePtr) ToTargetAuthenticationTypePtrOutputWithContext(ctx context.Context) TargetAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TargetAuthenticationTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleCreationPropertiesOutput{})
	pulumi.RegisterOutputType(AlertRuleCreationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KustoOfferingTypeOutput{})
	pulumi.RegisterOutputType(KustoOfferingTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(TargetAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(TargetAuthenticationTypePtrOutput{})
}
