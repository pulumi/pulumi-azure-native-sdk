// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Action that should be applied.
type ActionType string

const (
	ActionTypeAddActionGroups       = ActionType("AddActionGroups")
	ActionTypeRemoveAllActionGroups = ActionType("RemoveAllActionGroups")
	ActionTypeCorrelateAlerts       = ActionType("CorrelateAlerts")
)

// Days of week.
type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

func (DaysOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DaysOfWeek)(nil)).Elem()
}

func (e DaysOfWeek) ToDaysOfWeekOutput() DaysOfWeekOutput {
	return pulumi.ToOutput(e).(DaysOfWeekOutput)
}

func (e DaysOfWeek) ToDaysOfWeekOutputWithContext(ctx context.Context) DaysOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DaysOfWeekOutput)
}

func (e DaysOfWeek) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return e.ToDaysOfWeekPtrOutputWithContext(context.Background())
}

func (e DaysOfWeek) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return DaysOfWeek(e).ToDaysOfWeekOutputWithContext(ctx).ToDaysOfWeekPtrOutputWithContext(ctx)
}

func (e DaysOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DaysOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DaysOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DaysOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DaysOfWeekOutput struct{ *pulumi.OutputState }

func (DaysOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekOutput) ToDaysOfWeekOutput() DaysOfWeekOutput {
	return o
}

func (o DaysOfWeekOutput) ToDaysOfWeekOutputWithContext(ctx context.Context) DaysOfWeekOutput {
	return o
}

func (o DaysOfWeekOutput) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return o.ToDaysOfWeekPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DaysOfWeek) *DaysOfWeek {
		return &v
	}).(DaysOfWeekPtrOutput)
}

func (o DaysOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DaysOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DaysOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DaysOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DaysOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DaysOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekPtrOutput) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return o
}

func (o DaysOfWeekPtrOutput) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return o
}

func (o DaysOfWeekPtrOutput) Elem() DaysOfWeekOutput {
	return o.ApplyT(func(v *DaysOfWeek) DaysOfWeek {
		if v != nil {
			return *v
		}
		var ret DaysOfWeek
		return ret
	}).(DaysOfWeekOutput)
}

func (o DaysOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DaysOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DaysOfWeekInput is an input type that accepts values of the DaysOfWeek enum
// A concrete instance of `DaysOfWeekInput` can be one of the following:
//
//	DaysOfWeekSunday
//	DaysOfWeekMonday
//	DaysOfWeekTuesday
//	DaysOfWeekWednesday
//	DaysOfWeekThursday
//	DaysOfWeekFriday
//	DaysOfWeekSaturday
type DaysOfWeekInput interface {
	pulumi.Input

	ToDaysOfWeekOutput() DaysOfWeekOutput
	ToDaysOfWeekOutputWithContext(context.Context) DaysOfWeekOutput
}

var daysOfWeekPtrType = reflect.TypeOf((**DaysOfWeek)(nil)).Elem()

type DaysOfWeekPtrInput interface {
	pulumi.Input

	ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput
	ToDaysOfWeekPtrOutputWithContext(context.Context) DaysOfWeekPtrOutput
}

type daysOfWeekPtr string

func DaysOfWeekPtr(v string) DaysOfWeekPtrInput {
	return (*daysOfWeekPtr)(&v)
}

func (*daysOfWeekPtr) ElementType() reflect.Type {
	return daysOfWeekPtrType
}

func (in *daysOfWeekPtr) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DaysOfWeekPtrOutput)
}

func (in *daysOfWeekPtr) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DaysOfWeekPtrOutput)
}

// Field for a given condition.
type Field string

const (
	FieldSeverity            = Field("Severity")
	FieldMonitorService      = Field("MonitorService")
	FieldMonitorCondition    = Field("MonitorCondition")
	FieldSignalType          = Field("SignalType")
	FieldTargetResourceType  = Field("TargetResourceType")
	FieldTargetResource      = Field("TargetResource")
	FieldTargetResourceGroup = Field("TargetResourceGroup")
	FieldAlertRuleId         = Field("AlertRuleId")
	FieldAlertRuleName       = Field("AlertRuleName")
	FieldDescription         = Field("Description")
	FieldAlertContext        = Field("AlertContext")
)

func (Field) ElementType() reflect.Type {
	return reflect.TypeOf((*Field)(nil)).Elem()
}

func (e Field) ToFieldOutput() FieldOutput {
	return pulumi.ToOutput(e).(FieldOutput)
}

func (e Field) ToFieldOutputWithContext(ctx context.Context) FieldOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FieldOutput)
}

func (e Field) ToFieldPtrOutput() FieldPtrOutput {
	return e.ToFieldPtrOutputWithContext(context.Background())
}

func (e Field) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return Field(e).ToFieldOutputWithContext(ctx).ToFieldPtrOutputWithContext(ctx)
}

func (e Field) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Field) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Field) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Field) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FieldOutput struct{ *pulumi.OutputState }

func (FieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Field)(nil)).Elem()
}

func (o FieldOutput) ToFieldOutput() FieldOutput {
	return o
}

func (o FieldOutput) ToFieldOutputWithContext(ctx context.Context) FieldOutput {
	return o
}

func (o FieldOutput) ToFieldPtrOutput() FieldPtrOutput {
	return o.ToFieldPtrOutputWithContext(context.Background())
}

func (o FieldOutput) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Field) *Field {
		return &v
	}).(FieldPtrOutput)
}

func (o FieldOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FieldOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Field) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FieldOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FieldOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Field) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FieldPtrOutput struct{ *pulumi.OutputState }

func (FieldPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Field)(nil)).Elem()
}

func (o FieldPtrOutput) ToFieldPtrOutput() FieldPtrOutput {
	return o
}

func (o FieldPtrOutput) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return o
}

func (o FieldPtrOutput) Elem() FieldOutput {
	return o.ApplyT(func(v *Field) Field {
		if v != nil {
			return *v
		}
		var ret Field
		return ret
	}).(FieldOutput)
}

func (o FieldPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FieldPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Field) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FieldInput is an input type that accepts values of the Field enum
// A concrete instance of `FieldInput` can be one of the following:
//
//	FieldSeverity
//	FieldMonitorService
//	FieldMonitorCondition
//	FieldSignalType
//	FieldTargetResourceType
//	FieldTargetResource
//	FieldTargetResourceGroup
//	FieldAlertRuleId
//	FieldAlertRuleName
//	FieldDescription
//	FieldAlertContext
type FieldInput interface {
	pulumi.Input

	ToFieldOutput() FieldOutput
	ToFieldOutputWithContext(context.Context) FieldOutput
}

var fieldPtrType = reflect.TypeOf((**Field)(nil)).Elem()

type FieldPtrInput interface {
	pulumi.Input

	ToFieldPtrOutput() FieldPtrOutput
	ToFieldPtrOutputWithContext(context.Context) FieldPtrOutput
}

type fieldPtr string

func FieldPtr(v string) FieldPtrInput {
	return (*fieldPtr)(&v)
}

func (*fieldPtr) ElementType() reflect.Type {
	return fieldPtrType
}

func (in *fieldPtr) ToFieldPtrOutput() FieldPtrOutput {
	return pulumi.ToOutput(in).(FieldPtrOutput)
}

func (in *fieldPtr) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FieldPtrOutput)
}

// Indicates how to handle child alerts notifications.
type NotificationsForCorrelatedAlerts string

const (
	NotificationsForCorrelatedAlertsNotifyAlways   = NotificationsForCorrelatedAlerts("NotifyAlways")
	NotificationsForCorrelatedAlertsSuppressAlways = NotificationsForCorrelatedAlerts("SuppressAlways")
)

func (NotificationsForCorrelatedAlerts) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationsForCorrelatedAlerts)(nil)).Elem()
}

func (e NotificationsForCorrelatedAlerts) ToNotificationsForCorrelatedAlertsOutput() NotificationsForCorrelatedAlertsOutput {
	return pulumi.ToOutput(e).(NotificationsForCorrelatedAlertsOutput)
}

func (e NotificationsForCorrelatedAlerts) ToNotificationsForCorrelatedAlertsOutputWithContext(ctx context.Context) NotificationsForCorrelatedAlertsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NotificationsForCorrelatedAlertsOutput)
}

func (e NotificationsForCorrelatedAlerts) ToNotificationsForCorrelatedAlertsPtrOutput() NotificationsForCorrelatedAlertsPtrOutput {
	return e.ToNotificationsForCorrelatedAlertsPtrOutputWithContext(context.Background())
}

func (e NotificationsForCorrelatedAlerts) ToNotificationsForCorrelatedAlertsPtrOutputWithContext(ctx context.Context) NotificationsForCorrelatedAlertsPtrOutput {
	return NotificationsForCorrelatedAlerts(e).ToNotificationsForCorrelatedAlertsOutputWithContext(ctx).ToNotificationsForCorrelatedAlertsPtrOutputWithContext(ctx)
}

func (e NotificationsForCorrelatedAlerts) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationsForCorrelatedAlerts) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationsForCorrelatedAlerts) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotificationsForCorrelatedAlerts) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NotificationsForCorrelatedAlertsOutput struct{ *pulumi.OutputState }

func (NotificationsForCorrelatedAlertsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationsForCorrelatedAlerts)(nil)).Elem()
}

func (o NotificationsForCorrelatedAlertsOutput) ToNotificationsForCorrelatedAlertsOutput() NotificationsForCorrelatedAlertsOutput {
	return o
}

func (o NotificationsForCorrelatedAlertsOutput) ToNotificationsForCorrelatedAlertsOutputWithContext(ctx context.Context) NotificationsForCorrelatedAlertsOutput {
	return o
}

func (o NotificationsForCorrelatedAlertsOutput) ToNotificationsForCorrelatedAlertsPtrOutput() NotificationsForCorrelatedAlertsPtrOutput {
	return o.ToNotificationsForCorrelatedAlertsPtrOutputWithContext(context.Background())
}

func (o NotificationsForCorrelatedAlertsOutput) ToNotificationsForCorrelatedAlertsPtrOutputWithContext(ctx context.Context) NotificationsForCorrelatedAlertsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationsForCorrelatedAlerts) *NotificationsForCorrelatedAlerts {
		return &v
	}).(NotificationsForCorrelatedAlertsPtrOutput)
}

func (o NotificationsForCorrelatedAlertsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NotificationsForCorrelatedAlertsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationsForCorrelatedAlerts) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NotificationsForCorrelatedAlertsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationsForCorrelatedAlertsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationsForCorrelatedAlerts) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NotificationsForCorrelatedAlertsPtrOutput struct{ *pulumi.OutputState }

func (NotificationsForCorrelatedAlertsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationsForCorrelatedAlerts)(nil)).Elem()
}

func (o NotificationsForCorrelatedAlertsPtrOutput) ToNotificationsForCorrelatedAlertsPtrOutput() NotificationsForCorrelatedAlertsPtrOutput {
	return o
}

func (o NotificationsForCorrelatedAlertsPtrOutput) ToNotificationsForCorrelatedAlertsPtrOutputWithContext(ctx context.Context) NotificationsForCorrelatedAlertsPtrOutput {
	return o
}

func (o NotificationsForCorrelatedAlertsPtrOutput) Elem() NotificationsForCorrelatedAlertsOutput {
	return o.ApplyT(func(v *NotificationsForCorrelatedAlerts) NotificationsForCorrelatedAlerts {
		if v != nil {
			return *v
		}
		var ret NotificationsForCorrelatedAlerts
		return ret
	}).(NotificationsForCorrelatedAlertsOutput)
}

func (o NotificationsForCorrelatedAlertsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationsForCorrelatedAlertsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NotificationsForCorrelatedAlerts) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NotificationsForCorrelatedAlertsInput is an input type that accepts values of the NotificationsForCorrelatedAlerts enum
// A concrete instance of `NotificationsForCorrelatedAlertsInput` can be one of the following:
//
//	NotificationsForCorrelatedAlertsNotifyAlways
//	NotificationsForCorrelatedAlertsSuppressAlways
type NotificationsForCorrelatedAlertsInput interface {
	pulumi.Input

	ToNotificationsForCorrelatedAlertsOutput() NotificationsForCorrelatedAlertsOutput
	ToNotificationsForCorrelatedAlertsOutputWithContext(context.Context) NotificationsForCorrelatedAlertsOutput
}

var notificationsForCorrelatedAlertsPtrType = reflect.TypeOf((**NotificationsForCorrelatedAlerts)(nil)).Elem()

type NotificationsForCorrelatedAlertsPtrInput interface {
	pulumi.Input

	ToNotificationsForCorrelatedAlertsPtrOutput() NotificationsForCorrelatedAlertsPtrOutput
	ToNotificationsForCorrelatedAlertsPtrOutputWithContext(context.Context) NotificationsForCorrelatedAlertsPtrOutput
}

type notificationsForCorrelatedAlertsPtr string

func NotificationsForCorrelatedAlertsPtr(v string) NotificationsForCorrelatedAlertsPtrInput {
	return (*notificationsForCorrelatedAlertsPtr)(&v)
}

func (*notificationsForCorrelatedAlertsPtr) ElementType() reflect.Type {
	return notificationsForCorrelatedAlertsPtrType
}

func (in *notificationsForCorrelatedAlertsPtr) ToNotificationsForCorrelatedAlertsPtrOutput() NotificationsForCorrelatedAlertsPtrOutput {
	return pulumi.ToOutput(in).(NotificationsForCorrelatedAlertsPtrOutput)
}

func (in *notificationsForCorrelatedAlertsPtr) ToNotificationsForCorrelatedAlertsPtrOutputWithContext(ctx context.Context) NotificationsForCorrelatedAlertsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NotificationsForCorrelatedAlertsPtrOutput)
}

// Operator for a given condition.
type Operator string

const (
	OperatorEquals         = Operator("Equals")
	OperatorNotEquals      = Operator("NotEquals")
	OperatorContains       = Operator("Contains")
	OperatorDoesNotContain = Operator("DoesNotContain")
)

func (Operator) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (e Operator) ToOperatorOutput() OperatorOutput {
	return pulumi.ToOutput(e).(OperatorOutput)
}

func (e Operator) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorOutput)
}

func (e Operator) ToOperatorPtrOutput() OperatorPtrOutput {
	return e.ToOperatorPtrOutputWithContext(context.Background())
}

func (e Operator) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return Operator(e).ToOperatorOutputWithContext(ctx).ToOperatorPtrOutputWithContext(ctx)
}

func (e Operator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Operator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorOutput struct{ *pulumi.OutputState }

func (OperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (o OperatorOutput) ToOperatorOutput() OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o.ToOperatorPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Operator) *Operator {
		return &v
	}).(OperatorPtrOutput)
}

func (o OperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorPtrOutput struct{ *pulumi.OutputState }

func (OperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Operator)(nil)).Elem()
}

func (o OperatorPtrOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) Elem() OperatorOutput {
	return o.ApplyT(func(v *Operator) Operator {
		if v != nil {
			return *v
		}
		var ret Operator
		return ret
	}).(OperatorOutput)
}

func (o OperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Operator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OperatorInput is an input type that accepts values of the Operator enum
// A concrete instance of `OperatorInput` can be one of the following:
//
//	OperatorEquals
//	OperatorNotEquals
//	OperatorContains
//	OperatorDoesNotContain
type OperatorInput interface {
	pulumi.Input

	ToOperatorOutput() OperatorOutput
	ToOperatorOutputWithContext(context.Context) OperatorOutput
}

var operatorPtrType = reflect.TypeOf((**Operator)(nil)).Elem()

type OperatorPtrInput interface {
	pulumi.Input

	ToOperatorPtrOutput() OperatorPtrOutput
	ToOperatorPtrOutputWithContext(context.Context) OperatorPtrOutput
}

type operatorPtr string

func OperatorPtr(v string) OperatorPtrInput {
	return (*operatorPtr)(&v)
}

func (*operatorPtr) ElementType() reflect.Type {
	return operatorPtrType
}

func (in *operatorPtr) ToOperatorPtrOutput() OperatorPtrOutput {
	return pulumi.ToOutput(in).(OperatorPtrOutput)
}

func (in *operatorPtr) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorPtrOutput)
}

// Specifies when the recurrence should be applied.
type RecurrenceType string

const (
	RecurrenceTypeDaily   = RecurrenceType("Daily")
	RecurrenceTypeWeekly  = RecurrenceType("Weekly")
	RecurrenceTypeMonthly = RecurrenceType("Monthly")
)

// The type of update that needs to be performed.
type UpdateType string

const (
	UpdateTypeTimeBased = UpdateType("timeBased")
)

func (UpdateType) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateType)(nil)).Elem()
}

func (e UpdateType) ToUpdateTypeOutput() UpdateTypeOutput {
	return pulumi.ToOutput(e).(UpdateTypeOutput)
}

func (e UpdateType) ToUpdateTypeOutputWithContext(ctx context.Context) UpdateTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UpdateTypeOutput)
}

func (e UpdateType) ToUpdateTypePtrOutput() UpdateTypePtrOutput {
	return e.ToUpdateTypePtrOutputWithContext(context.Background())
}

func (e UpdateType) ToUpdateTypePtrOutputWithContext(ctx context.Context) UpdateTypePtrOutput {
	return UpdateType(e).ToUpdateTypeOutputWithContext(ctx).ToUpdateTypePtrOutputWithContext(ctx)
}

func (e UpdateType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpdateType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpdateType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UpdateType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UpdateTypeOutput struct{ *pulumi.OutputState }

func (UpdateTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateType)(nil)).Elem()
}

func (o UpdateTypeOutput) ToUpdateTypeOutput() UpdateTypeOutput {
	return o
}

func (o UpdateTypeOutput) ToUpdateTypeOutputWithContext(ctx context.Context) UpdateTypeOutput {
	return o
}

func (o UpdateTypeOutput) ToUpdateTypePtrOutput() UpdateTypePtrOutput {
	return o.ToUpdateTypePtrOutputWithContext(context.Background())
}

func (o UpdateTypeOutput) ToUpdateTypePtrOutputWithContext(ctx context.Context) UpdateTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpdateType) *UpdateType {
		return &v
	}).(UpdateTypePtrOutput)
}

func (o UpdateTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UpdateTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpdateType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UpdateTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpdateTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpdateType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UpdateTypePtrOutput struct{ *pulumi.OutputState }

func (UpdateTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateType)(nil)).Elem()
}

func (o UpdateTypePtrOutput) ToUpdateTypePtrOutput() UpdateTypePtrOutput {
	return o
}

func (o UpdateTypePtrOutput) ToUpdateTypePtrOutputWithContext(ctx context.Context) UpdateTypePtrOutput {
	return o
}

func (o UpdateTypePtrOutput) Elem() UpdateTypeOutput {
	return o.ApplyT(func(v *UpdateType) UpdateType {
		if v != nil {
			return *v
		}
		var ret UpdateType
		return ret
	}).(UpdateTypeOutput)
}

func (o UpdateTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpdateTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UpdateType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// UpdateTypeInput is an input type that accepts values of the UpdateType enum
// A concrete instance of `UpdateTypeInput` can be one of the following:
//
//	UpdateTypeTimeBased
type UpdateTypeInput interface {
	pulumi.Input

	ToUpdateTypeOutput() UpdateTypeOutput
	ToUpdateTypeOutputWithContext(context.Context) UpdateTypeOutput
}

var updateTypePtrType = reflect.TypeOf((**UpdateType)(nil)).Elem()

type UpdateTypePtrInput interface {
	pulumi.Input

	ToUpdateTypePtrOutput() UpdateTypePtrOutput
	ToUpdateTypePtrOutputWithContext(context.Context) UpdateTypePtrOutput
}

type updateTypePtr string

func UpdateTypePtr(v string) UpdateTypePtrInput {
	return (*updateTypePtr)(&v)
}

func (*updateTypePtr) ElementType() reflect.Type {
	return updateTypePtrType
}

func (in *updateTypePtr) ToUpdateTypePtrOutput() UpdateTypePtrOutput {
	return pulumi.ToOutput(in).(UpdateTypePtrOutput)
}

func (in *updateTypePtr) ToUpdateTypePtrOutputWithContext(ctx context.Context) UpdateTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UpdateTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DaysOfWeekOutput{})
	pulumi.RegisterOutputType(DaysOfWeekPtrOutput{})
	pulumi.RegisterOutputType(FieldOutput{})
	pulumi.RegisterOutputType(FieldPtrOutput{})
	pulumi.RegisterOutputType(NotificationsForCorrelatedAlertsOutput{})
	pulumi.RegisterOutputType(NotificationsForCorrelatedAlertsPtrOutput{})
	pulumi.RegisterOutputType(OperatorOutput{})
	pulumi.RegisterOutputType(OperatorPtrOutput{})
	pulumi.RegisterOutputType(UpdateTypeOutput{})
	pulumi.RegisterOutputType(UpdateTypePtrOutput{})
}
