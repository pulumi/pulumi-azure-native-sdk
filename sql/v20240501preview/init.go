// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240501preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:sql/v20240501preview:BackupShortTermRetentionPolicy":
		r = &BackupShortTermRetentionPolicy{}
	case "azure-native:sql/v20240501preview:DataMaskingPolicy":
		r = &DataMaskingPolicy{}
	case "azure-native:sql/v20240501preview:Database":
		r = &Database{}
	case "azure-native:sql/v20240501preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20240501preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20240501preview:DatabaseSecurityAlertPolicy":
		r = &DatabaseSecurityAlertPolicy{}
	case "azure-native:sql/v20240501preview:DatabaseSqlVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseSqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20240501preview:DatabaseVulnerabilityAssessment":
		r = &DatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20240501preview:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20240501preview:DistributedAvailabilityGroup":
		r = &DistributedAvailabilityGroup{}
	case "azure-native:sql/v20240501preview:ElasticPool":
		r = &ElasticPool{}
	case "azure-native:sql/v20240501preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20240501preview:ExtendedDatabaseBlobAuditingPolicy":
		r = &ExtendedDatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20240501preview:ExtendedServerBlobAuditingPolicy":
		r = &ExtendedServerBlobAuditingPolicy{}
	case "azure-native:sql/v20240501preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20240501preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20240501preview:GeoBackupPolicy":
		r = &GeoBackupPolicy{}
	case "azure-native:sql/v20240501preview:IPv6FirewallRule":
		r = &IPv6FirewallRule{}
	case "azure-native:sql/v20240501preview:InstanceFailoverGroup":
		r = &InstanceFailoverGroup{}
	case "azure-native:sql/v20240501preview:InstancePool":
		r = &InstancePool{}
	case "azure-native:sql/v20240501preview:Job":
		r = &Job{}
	case "azure-native:sql/v20240501preview:JobAgent":
		r = &JobAgent{}
	case "azure-native:sql/v20240501preview:JobCredential":
		r = &JobCredential{}
	case "azure-native:sql/v20240501preview:JobPrivateEndpoint":
		r = &JobPrivateEndpoint{}
	case "azure-native:sql/v20240501preview:JobStep":
		r = &JobStep{}
	case "azure-native:sql/v20240501preview:JobTargetGroup":
		r = &JobTargetGroup{}
	case "azure-native:sql/v20240501preview:LongTermRetentionPolicy":
		r = &LongTermRetentionPolicy{}
	case "azure-native:sql/v20240501preview:ManagedDatabase":
		r = &ManagedDatabase{}
	case "azure-native:sql/v20240501preview:ManagedDatabaseSensitivityLabel":
		r = &ManagedDatabaseSensitivityLabel{}
	case "azure-native:sql/v20240501preview:ManagedDatabaseVulnerabilityAssessment":
		r = &ManagedDatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20240501preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
		r = &ManagedDatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20240501preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20240501preview:ManagedInstanceAdministrator":
		r = &ManagedInstanceAdministrator{}
	case "azure-native:sql/v20240501preview:ManagedInstanceAzureADOnlyAuthentication":
		r = &ManagedInstanceAzureADOnlyAuthentication{}
	case "azure-native:sql/v20240501preview:ManagedInstanceKey":
		r = &ManagedInstanceKey{}
	case "azure-native:sql/v20240501preview:ManagedInstanceLongTermRetentionPolicy":
		r = &ManagedInstanceLongTermRetentionPolicy{}
	case "azure-native:sql/v20240501preview:ManagedInstancePrivateEndpointConnection":
		r = &ManagedInstancePrivateEndpointConnection{}
	case "azure-native:sql/v20240501preview:ManagedInstanceVulnerabilityAssessment":
		r = &ManagedInstanceVulnerabilityAssessment{}
	case "azure-native:sql/v20240501preview:ManagedServerDnsAlias":
		r = &ManagedServerDnsAlias{}
	case "azure-native:sql/v20240501preview:OutboundFirewallRule":
		r = &OutboundFirewallRule{}
	case "azure-native:sql/v20240501preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:sql/v20240501preview:ReplicationLink":
		r = &ReplicationLink{}
	case "azure-native:sql/v20240501preview:SensitivityLabel":
		r = &SensitivityLabel{}
	case "azure-native:sql/v20240501preview:Server":
		r = &Server{}
	case "azure-native:sql/v20240501preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20240501preview:ServerAzureADAdministrator":
		r = &ServerAzureADAdministrator{}
	case "azure-native:sql/v20240501preview:ServerAzureADOnlyAuthentication":
		r = &ServerAzureADOnlyAuthentication{}
	case "azure-native:sql/v20240501preview:ServerBlobAuditingPolicy":
		r = &ServerBlobAuditingPolicy{}
	case "azure-native:sql/v20240501preview:ServerDnsAlias":
		r = &ServerDnsAlias{}
	case "azure-native:sql/v20240501preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20240501preview:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure-native:sql/v20240501preview:ServerTrustCertificate":
		r = &ServerTrustCertificate{}
	case "azure-native:sql/v20240501preview:ServerTrustGroup":
		r = &ServerTrustGroup{}
	case "azure-native:sql/v20240501preview:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:sql/v20240501preview:SqlVulnerabilityAssessmentRuleBaseline":
		r = &SqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20240501preview:SqlVulnerabilityAssessmentsSetting":
		r = &SqlVulnerabilityAssessmentsSetting{}
	case "azure-native:sql/v20240501preview:StartStopManagedInstanceSchedule":
		r = &StartStopManagedInstanceSchedule{}
	case "azure-native:sql/v20240501preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20240501preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20240501preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20240501preview:TransparentDataEncryption":
		r = &TransparentDataEncryption{}
	case "azure-native:sql/v20240501preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	case "azure-native:sql/v20240501preview:WorkloadClassifier":
		r = &WorkloadClassifier{}
	case "azure-native:sql/v20240501preview:WorkloadGroup":
		r = &WorkloadGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"sql/v20240501preview",
		&module{version},
	)
}
