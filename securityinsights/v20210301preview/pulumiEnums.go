// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Alert detail
type AlertDetail string

const (
	// Alert display name
	AlertDetailDisplayName = AlertDetail("DisplayName")
	// Alert severity
	AlertDetailSeverity = AlertDetail("Severity")
)

// The kind of the alert rule
type AlertRuleKind string

const (
	AlertRuleKindScheduled                         = AlertRuleKind("Scheduled")
	AlertRuleKindMicrosoftSecurityIncidentCreation = AlertRuleKind("MicrosoftSecurityIncidentCreation")
	AlertRuleKindFusion                            = AlertRuleKind("Fusion")
	AlertRuleKindMLBehaviorAnalytics               = AlertRuleKind("MLBehaviorAnalytics")
	AlertRuleKindThreatIntelligence                = AlertRuleKind("ThreatIntelligence")
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
	AttackTacticPreAttack           = AttackTactic("PreAttack")
)

// type of connectivity
type ConnectivityType string

const (
	ConnectivityTypeIsConnectedQuery = ConnectivityType("IsConnectedQuery")
)

// Content type.
type ContentType string

const (
	ContentTypeAnalyticRule = ContentType("AnalyticRule")
	ContentTypeWorkbook     = ContentType("Workbook")
)

// The type of identity that last modified the resource.
type CreatedByType string

const (
	CreatedByTypeUser            = CreatedByType("User")
	CreatedByTypeApplication     = CreatedByType("Application")
	CreatedByTypeManagedIdentity = CreatedByType("ManagedIdentity")
	CreatedByTypeKey             = CreatedByType("Key")
)

// the entity query kind
type CustomEntityQueryKind string

const (
	CustomEntityQueryKindActivity = CustomEntityQueryKind("Activity")
)

// The data connector kind
type DataConnectorKind string

const (
	DataConnectorKindAzureActiveDirectory                      = DataConnectorKind("AzureActiveDirectory")
	DataConnectorKindAzureSecurityCenter                       = DataConnectorKind("AzureSecurityCenter")
	DataConnectorKindMicrosoftCloudAppSecurity                 = DataConnectorKind("MicrosoftCloudAppSecurity")
	DataConnectorKindThreatIntelligence                        = DataConnectorKind("ThreatIntelligence")
	DataConnectorKindThreatIntelligenceTaxii                   = DataConnectorKind("ThreatIntelligenceTaxii")
	DataConnectorKindOffice365                                 = DataConnectorKind("Office365")
	DataConnectorKindOfficeATP                                 = DataConnectorKind("OfficeATP")
	DataConnectorKindAmazonWebServicesCloudTrail               = DataConnectorKind("AmazonWebServicesCloudTrail")
	DataConnectorKindAzureAdvancedThreatProtection             = DataConnectorKind("AzureAdvancedThreatProtection")
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection = DataConnectorKind("MicrosoftDefenderAdvancedThreatProtection")
	DataConnectorKindDynamics365                               = DataConnectorKind("Dynamics365")
	DataConnectorKindMicrosoftThreatProtection                 = DataConnectorKind("MicrosoftThreatProtection")
	DataConnectorKindMicrosoftThreatIntelligence               = DataConnectorKind("MicrosoftThreatIntelligence")
	DataConnectorKindGenericUI                                 = DataConnectorKind("GenericUI")
)

// Describe whether this data type connection is enabled or not.
type DataTypeState string

const (
	DataTypeStateEnabled  = DataTypeState("Enabled")
	DataTypeStateDisabled = DataTypeState("Disabled")
)

// The V3 type of the mapped entity
type EntityMappingType string

const (
	// User account entity type
	EntityMappingTypeAccount = EntityMappingType("Account")
	// Host entity type
	EntityMappingTypeHost = EntityMappingType("Host")
	// IP address entity type
	EntityMappingTypeIP = EntityMappingType("IP")
	// Malware entity type
	EntityMappingTypeMalware = EntityMappingType("Malware")
	// System file entity type
	EntityMappingTypeFile = EntityMappingType("File")
	// Process entity type
	EntityMappingTypeProcess = EntityMappingType("Process")
	// Cloud app entity type
	EntityMappingTypeCloudApplication = EntityMappingType("CloudApplication")
	// DNS entity type
	EntityMappingTypeDNS = EntityMappingType("DNS")
	// Azure resource entity type
	EntityMappingTypeAzureResource = EntityMappingType("AzureResource")
	// File-hash entity type
	EntityMappingTypeFileHash = EntityMappingType("FileHash")
	// Registry key entity type
	EntityMappingTypeRegistryKey = EntityMappingType("RegistryKey")
	// Registry value entity type
	EntityMappingTypeRegistryValue = EntityMappingType("RegistryValue")
	// Security group entity type
	EntityMappingTypeSecurityGroup = EntityMappingType("SecurityGroup")
	// URL entity type
	EntityMappingTypeURL = EntityMappingType("URL")
	// Mailbox entity type
	EntityMappingTypeMailbox = EntityMappingType("Mailbox")
	// Mail cluster entity type
	EntityMappingTypeMailCluster = EntityMappingType("MailCluster")
	// Mail message entity type
	EntityMappingTypeMailMessage = EntityMappingType("MailMessage")
	// Submission mail entity type
	EntityMappingTypeSubmissionMail = EntityMappingType("SubmissionMail")
)

// The type of the query's source entity
type EntityType string

const (
	// Entity represents account in the system.
	EntityTypeAccount = EntityType("Account")
	// Entity represents host in the system.
	EntityTypeHost = EntityType("Host")
	// Entity represents file in the system.
	EntityTypeFile = EntityType("File")
	// Entity represents azure resource in the system.
	EntityTypeAzureResource = EntityType("AzureResource")
	// Entity represents cloud application in the system.
	EntityTypeCloudApplication = EntityType("CloudApplication")
	// Entity represents dns in the system.
	EntityTypeDNS = EntityType("DNS")
	// Entity represents file hash in the system.
	EntityTypeFileHash = EntityType("FileHash")
	// Entity represents ip in the system.
	EntityTypeIP = EntityType("IP")
	// Entity represents malware in the system.
	EntityTypeMalware = EntityType("Malware")
	// Entity represents process in the system.
	EntityTypeProcess = EntityType("Process")
	// Entity represents registry key in the system.
	EntityTypeRegistryKey = EntityType("RegistryKey")
	// Entity represents registry value in the system.
	EntityTypeRegistryValue = EntityType("RegistryValue")
	// Entity represents security group in the system.
	EntityTypeSecurityGroup = EntityType("SecurityGroup")
	// Entity represents url in the system.
	EntityTypeURL = EntityType("URL")
	// Entity represents IoT device in the system.
	EntityTypeIoTDevice = EntityType("IoTDevice")
	// Entity represents security alert in the system.
	EntityTypeSecurityAlert = EntityType("SecurityAlert")
	// Entity represents HuntingBookmark in the system.
	EntityTypeHuntingBookmark = EntityType("HuntingBookmark")
	// Entity represents mail cluster in the system.
	EntityTypeMailCluster = EntityType("MailCluster")
	// Entity represents mail message in the system.
	EntityTypeMailMessage = EntityType("MailMessage")
	// Entity represents mailbox in the system.
	EntityTypeMailbox = EntityType("Mailbox")
	// Entity represents submission mail in the system.
	EntityTypeSubmissionMail = EntityType("SubmissionMail")
)

// The event grouping aggregation kinds
type EventGroupingAggregationKind string

const (
	EventGroupingAggregationKindSingleAlert    = EventGroupingAggregationKind("SingleAlert")
	EventGroupingAggregationKindAlertPerResult = EventGroupingAggregationKind("AlertPerResult")
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

// The kind of content the metadata is for.
type Kind string

const (
	KindDataConnector         = Kind("DataConnector")
	KindDataType              = Kind("DataType")
	KindWorkbook              = Kind("Workbook")
	KindWorkbookTemplate      = Kind("WorkbookTemplate")
	KindPlaybook              = Kind("Playbook")
	KindPlaybookTemplate      = Kind("PlaybookTemplate")
	KindAnalyticsRuleTemplate = Kind("AnalyticsRuleTemplate")
	KindAnalyticsRule         = Kind("AnalyticsRule")
	KindHuntingQuery          = Kind("HuntingQuery")
	KindInvestigationQuery    = Kind("InvestigationQuery")
	KindParser                = Kind("Parser")
	KindWatchlist             = Kind("Watchlist")
	KindWatchlistTemplate     = Kind("WatchlistTemplate")
	KindSolution              = Kind("Solution")
)

// Grouping matching method. When method is Selected at least one of groupByEntities, groupByAlertDetails, groupByCustomDetails must be provided and not empty.
type MatchingMethod string

const (
	// Grouping alerts into a single incident if all the entities match
	MatchingMethodAllEntities = MatchingMethod("AllEntities")
	// Grouping any alerts triggered by this rule into a single incident
	MatchingMethodAnyAlert = MatchingMethod("AnyAlert")
	// Grouping alerts into a single incident if the selected entities, custom details and alert details match
	MatchingMethodSelected = MatchingMethod("Selected")
)

// The alerts' productName on which the cases will be generated
type MicrosoftSecurityProductName string

const (
	MicrosoftSecurityProductName_Microsoft_Cloud_App_Security                  = MicrosoftSecurityProductName("Microsoft Cloud App Security")
	MicrosoftSecurityProductName_Azure_Security_Center                         = MicrosoftSecurityProductName("Azure Security Center")
	MicrosoftSecurityProductName_Azure_Advanced_Threat_Protection              = MicrosoftSecurityProductName("Azure Advanced Threat Protection")
	MicrosoftSecurityProductName_Azure_Active_Directory_Identity_Protection    = MicrosoftSecurityProductName("Azure Active Directory Identity Protection")
	MicrosoftSecurityProductName_Azure_Security_Center_for_IoT                 = MicrosoftSecurityProductName("Azure Security Center for IoT")
	MicrosoftSecurityProductName_Office_365_Advanced_Threat_Protection         = MicrosoftSecurityProductName("Office 365 Advanced Threat Protection")
	MicrosoftSecurityProductName_Microsoft_Defender_Advanced_Threat_Protection = MicrosoftSecurityProductName("Microsoft Defender Advanced Threat Protection")
)

// Operator used for list of dependencies in criteria array.
type Operator string

const (
	OperatorAND = Operator("AND")
	OperatorOR  = Operator("OR")
)

// Permission provider scope
type PermissionProviderScope string

const (
	PermissionProviderScopeResourceGroup = PermissionProviderScope("ResourceGroup")
	PermissionProviderScopeSubscription  = PermissionProviderScope("Subscription")
	PermissionProviderScopeWorkspace     = PermissionProviderScope("Workspace")
)

// The polling frequency for the TAXII server.
type PollingFrequency string

const (
	// Once a minute
	PollingFrequencyOnceAMinute = PollingFrequency("OnceAMinute")
	// Once an hour
	PollingFrequencyOnceAnHour = PollingFrequency("OnceAnHour")
	// Once a day
	PollingFrequencyOnceADay = PollingFrequency("OnceADay")
)

// Provider name
type ProviderName string

const (
	ProviderName_Microsoft_OperationalInsights_solutions              = ProviderName("Microsoft.OperationalInsights/solutions")
	ProviderName_Microsoft_OperationalInsights_workspaces             = ProviderName("Microsoft.OperationalInsights/workspaces")
	ProviderName_Microsoft_OperationalInsights_workspaces_datasources = ProviderName("Microsoft.OperationalInsights/workspaces/datasources")
	ProviderName_Microsoft_aadiam_diagnosticSettings                  = ProviderName("microsoft.aadiam/diagnosticSettings")
	ProviderName_Microsoft_OperationalInsights_workspaces_sharedKeys  = ProviderName("Microsoft.OperationalInsights/workspaces/sharedKeys")
	ProviderName_Microsoft_Authorization_policyAssignments            = ProviderName("Microsoft.Authorization/policyAssignments")
)

// The repository type of the source control
type RepoType string

const (
	RepoTypeGithub = RepoType("Github")
	RepoTypeDevOps = RepoType("DevOps")
)

// The kind of the setting
type SettingKind string

const (
	SettingKindAnomalies       = SettingKind("Anomalies")
	SettingKindEyesOn          = SettingKind("EyesOn")
	SettingKindEntityAnalytics = SettingKind("EntityAnalytics")
	SettingKindUeba            = SettingKind("Ueba")
)

// The kind of the setting
type SettingType string

const (
	SettingTypeCopyableLabel         = SettingType("CopyableLabel")
	SettingTypeInstructionStepsGroup = SettingType("InstructionStepsGroup")
	SettingTypeInfoMessage           = SettingType("InfoMessage")
)

// The source of the watchlist
type Source string

const (
	Source_Local_file     = Source("Local file")
	Source_Remote_storage = Source("Remote storage")
)

// Source type of the content
type SourceKind string

const (
	SourceKindLocalWorkspace   = SourceKind("LocalWorkspace")
	SourceKindCommunity        = SourceKind("Community")
	SourceKindSolution         = SourceKind("Solution")
	SourceKindSourceRepository = SourceKind("SourceRepository")
)

// Type of support for content item
type SupportTier string

const (
	SupportTierMicrosoft = SupportTier("Microsoft")
	SupportTierPartner   = SupportTier("Partner")
	SupportTierCommunity = SupportTier("Community")
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

// The data source that enriched by ueba.
type UebaDataSources string

const (
	UebaDataSourcesAuditLogs     = UebaDataSources("AuditLogs")
	UebaDataSourcesAzureActivity = UebaDataSources("AzureActivity")
	UebaDataSourcesSecurityEvent = UebaDataSources("SecurityEvent")
	UebaDataSourcesSigninLogs    = UebaDataSources("SigninLogs")
)

func init() {
	pulumi.RegisterOutputType(TriggerOperatorOutput{})
	pulumi.RegisterOutputType(TriggerOperatorPtrOutput{})
}