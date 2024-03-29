// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

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
	case "azure-native:securityinsights/v20230601preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20230601preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20230601preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20230601preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20230601preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20230601preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20230601preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20230601preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20230601preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20230601preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20230601preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20230601preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20230601preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20230601preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20230601preview:ContentPackage":
		r = &ContentPackage{}
	case "azure-native:securityinsights/v20230601preview:ContentTemplate":
		r = &ContentTemplate{}
	case "azure-native:securityinsights/v20230601preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20230601preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20230601preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20230601preview:FileImport":
		r = &FileImport{}
	case "azure-native:securityinsights/v20230601preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20230601preview:GCPDataConnector":
		r = &GCPDataConnector{}
	case "azure-native:securityinsights/v20230601preview:Hunt":
		r = &Hunt{}
	case "azure-native:securityinsights/v20230601preview:HuntComment":
		r = &HuntComment{}
	case "azure-native:securityinsights/v20230601preview:HuntRelation":
		r = &HuntRelation{}
	case "azure-native:securityinsights/v20230601preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20230601preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20230601preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20230601preview:IncidentTask":
		r = &IncidentTask{}
	case "azure-native:securityinsights/v20230601preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20230601preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20230601preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20230601preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20230601preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20230601preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20230601preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20230601preview:MicrosoftPurviewInformationProtectionDataConnector":
		r = &MicrosoftPurviewInformationProtectionDataConnector{}
	case "azure-native:securityinsights/v20230601preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20230601preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20230601preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20230601preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20230601preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20230601preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20230601preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20230601preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20230601preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20230601preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20230601preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20230601preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20230601preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20230601preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20230601preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20230601preview:WatchlistItem":
		r = &WatchlistItem{}
	case "azure-native:securityinsights/v20230601preview:WorkspaceManagerAssignment":
		r = &WorkspaceManagerAssignment{}
	case "azure-native:securityinsights/v20230601preview:WorkspaceManagerConfiguration":
		r = &WorkspaceManagerConfiguration{}
	case "azure-native:securityinsights/v20230601preview:WorkspaceManagerGroup":
		r = &WorkspaceManagerGroup{}
	case "azure-native:securityinsights/v20230601preview:WorkspaceManagerMember":
		r = &WorkspaceManagerMember{}
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
		"securityinsights/v20230601preview",
		&module{version},
	)
}
