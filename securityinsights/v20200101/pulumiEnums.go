// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The alert rule kind
type AlertRuleKind string

const (
	AlertRuleKindScheduled                         = AlertRuleKind("Scheduled")
	AlertRuleKindMicrosoftSecurityIncidentCreation = AlertRuleKind("MicrosoftSecurityIncidentCreation")
	AlertRuleKindFusion                            = AlertRuleKind("Fusion")
)

// The severity for alerts created by this alert rule.
type AlertSeverity string

const (
	// High severity
	AlertSeverityHigh = AlertSeverity("High")
	// Medium severity
	AlertSeverityMedium = AlertSeverity("Medium")
	// Low severity
	AlertSeverityLow = AlertSeverity("Low")
	// Informational severity
	AlertSeverityInformational = AlertSeverity("Informational")
)

// The severity for alerts created by this alert rule.
type AttackTactic string

const (
	AttackTacticInitialAccess       = AttackTactic("InitialAccess")
	AttackTacticExecution           = AttackTactic("Execution")
	AttackTacticPersistence         = AttackTactic("Persistence")
	AttackTacticPrivilegeEscalation = AttackTactic("PrivilegeEscalation")
	AttackTacticDefenseEvasion      = AttackTactic("DefenseEvasion")
	AttackTacticCredentialAccess    = AttackTactic("CredentialAccess")
	AttackTacticDiscovery           = AttackTactic("Discovery")
	AttackTacticLateralMovement     = AttackTactic("LateralMovement")
	AttackTacticCollection          = AttackTactic("Collection")
	AttackTacticExfiltration        = AttackTactic("Exfiltration")
	AttackTacticCommandAndControl   = AttackTactic("CommandAndControl")
	AttackTacticImpact              = AttackTactic("Impact")
)

// The data connector kind
type DataConnectorKind string

const (
	DataConnectorKindAzureActiveDirectory                      = DataConnectorKind("AzureActiveDirectory")
	DataConnectorKindAzureSecurityCenter                       = DataConnectorKind("AzureSecurityCenter")
	DataConnectorKindMicrosoftCloudAppSecurity                 = DataConnectorKind("MicrosoftCloudAppSecurity")
	DataConnectorKindThreatIntelligence                        = DataConnectorKind("ThreatIntelligence")
	DataConnectorKindOffice365                                 = DataConnectorKind("Office365")
	DataConnectorKindAmazonWebServicesCloudTrail               = DataConnectorKind("AmazonWebServicesCloudTrail")
	DataConnectorKindAzureAdvancedThreatProtection             = DataConnectorKind("AzureAdvancedThreatProtection")
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection = DataConnectorKind("MicrosoftDefenderAdvancedThreatProtection")
)

// Describe whether this data type connection is enabled or not.
type DataTypeState string

const (
	DataTypeStateEnabled  = DataTypeState("Enabled")
	DataTypeStateDisabled = DataTypeState("Disabled")
)

// The reason the incident was closed
type IncidentClassification string

const (
	// Incident classification was undetermined
	IncidentClassificationUndetermined = IncidentClassification("Undetermined")
	// Incident was true positive
	IncidentClassificationTruePositive = IncidentClassification("TruePositive")
	// Incident was benign positive
	IncidentClassificationBenignPositive = IncidentClassification("BenignPositive")
	// Incident was false positive
	IncidentClassificationFalsePositive = IncidentClassification("FalsePositive")
)

// The classification reason the incident was closed with
type IncidentClassificationReason string

const (
	// Classification reason was suspicious activity
	IncidentClassificationReasonSuspiciousActivity = IncidentClassificationReason("SuspiciousActivity")
	// Classification reason was suspicious but expected
	IncidentClassificationReasonSuspiciousButExpected = IncidentClassificationReason("SuspiciousButExpected")
	// Classification reason was incorrect alert logic
	IncidentClassificationReasonIncorrectAlertLogic = IncidentClassificationReason("IncorrectAlertLogic")
	// Classification reason was inaccurate data
	IncidentClassificationReasonInaccurateData = IncidentClassificationReason("InaccurateData")
)

// The severity of the incident
type IncidentSeverity string

const (
	// High severity
	IncidentSeverityHigh = IncidentSeverity("High")
	// Medium severity
	IncidentSeverityMedium = IncidentSeverity("Medium")
	// Low severity
	IncidentSeverityLow = IncidentSeverity("Low")
	// Informational severity
	IncidentSeverityInformational = IncidentSeverity("Informational")
)

// The status of the incident
type IncidentStatus string

const (
	// An active incident which isn't being handled currently
	IncidentStatusNew = IncidentStatus("New")
	// An active incident which is being handled
	IncidentStatusActive = IncidentStatus("Active")
	// A non-active incident
	IncidentStatusClosed = IncidentStatus("Closed")
)

// The alerts' productName on which the cases will be generated
type MicrosoftSecurityProductName string

const (
	MicrosoftSecurityProductName_Microsoft_Cloud_App_Security               = MicrosoftSecurityProductName("Microsoft Cloud App Security")
	MicrosoftSecurityProductName_Azure_Security_Center                      = MicrosoftSecurityProductName("Azure Security Center")
	MicrosoftSecurityProductName_Azure_Advanced_Threat_Protection           = MicrosoftSecurityProductName("Azure Advanced Threat Protection")
	MicrosoftSecurityProductName_Azure_Active_Directory_Identity_Protection = MicrosoftSecurityProductName("Azure Active Directory Identity Protection")
	MicrosoftSecurityProductName_Azure_Security_Center_for_IoT              = MicrosoftSecurityProductName("Azure Security Center for IoT")
)

// The operation against the threshold that triggers alert rule.
type TriggerOperator string

const (
	TriggerOperatorGreaterThan = TriggerOperator("GreaterThan")
	TriggerOperatorLessThan    = TriggerOperator("LessThan")
	TriggerOperatorEqual       = TriggerOperator("Equal")
	TriggerOperatorNotEqual    = TriggerOperator("NotEqual")
)

func (TriggerOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperator)(nil)).Elem()
}

func (e TriggerOperator) ToTriggerOperatorOutput() TriggerOperatorOutput {
	return pulumi.ToOutput(e).(TriggerOperatorOutput)
}

func (e TriggerOperator) ToTriggerOperatorOutputWithContext(ctx context.Context) TriggerOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerOperatorOutput)
}

func (e TriggerOperator) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return e.ToTriggerOperatorPtrOutputWithContext(context.Background())
}

func (e TriggerOperator) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return TriggerOperator(e).ToTriggerOperatorOutputWithContext(ctx).ToTriggerOperatorPtrOutputWithContext(ctx)
}

func (e TriggerOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerOperatorOutput struct{ *pulumi.OutputState }

func (TriggerOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperator)(nil)).Elem()
}

func (o TriggerOperatorOutput) ToTriggerOperatorOutput() TriggerOperatorOutput {
	return o
}

func (o TriggerOperatorOutput) ToTriggerOperatorOutputWithContext(ctx context.Context) TriggerOperatorOutput {
	return o
}

func (o TriggerOperatorOutput) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return o.ToTriggerOperatorPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerOperator) *TriggerOperator {
		return &v
	}).(TriggerOperatorPtrOutput)
}

func (o TriggerOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerOperatorPtrOutput struct{ *pulumi.OutputState }

func (TriggerOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerOperator)(nil)).Elem()
}

func (o TriggerOperatorPtrOutput) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return o
}

func (o TriggerOperatorPtrOutput) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return o
}

func (o TriggerOperatorPtrOutput) Elem() TriggerOperatorOutput {
	return o.ApplyT(func(v *TriggerOperator) TriggerOperator {
		if v != nil {
			return *v
		}
		var ret TriggerOperator
		return ret
	}).(TriggerOperatorOutput)
}

func (o TriggerOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TriggerOperatorInput is an input type that accepts TriggerOperatorArgs and TriggerOperatorOutput values.
// You can construct a concrete instance of `TriggerOperatorInput` via:
//
//	TriggerOperatorArgs{...}
type TriggerOperatorInput interface {
	pulumi.Input

	ToTriggerOperatorOutput() TriggerOperatorOutput
	ToTriggerOperatorOutputWithContext(context.Context) TriggerOperatorOutput
}

var triggerOperatorPtrType = reflect.TypeOf((**TriggerOperator)(nil)).Elem()

type TriggerOperatorPtrInput interface {
	pulumi.Input

	ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput
	ToTriggerOperatorPtrOutputWithContext(context.Context) TriggerOperatorPtrOutput
}

type triggerOperatorPtr string

func TriggerOperatorPtr(v string) TriggerOperatorPtrInput {
	return (*triggerOperatorPtr)(&v)
}

func (*triggerOperatorPtr) ElementType() reflect.Type {
	return triggerOperatorPtrType
}

func (in *triggerOperatorPtr) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return pulumi.ToOutput(in).(TriggerOperatorPtrOutput)
}

func (in *triggerOperatorPtr) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerOperatorPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerOperatorOutput{})
	pulumi.RegisterOutputType(TriggerOperatorPtrOutput{})
}