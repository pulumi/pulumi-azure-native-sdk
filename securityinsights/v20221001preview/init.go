// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk"
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
	case "azure-native:securityinsights/v20221001preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20221001preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20221001preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20221001preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20221001preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20221001preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20221001preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20221001preview:AnomalySecurityMLAnalyticsSettings":
		r = &AnomalySecurityMLAnalyticsSettings{}
	case "azure-native:securityinsights/v20221001preview:AutomationRule":
		r = &AutomationRule{}
	case "azure-native:securityinsights/v20221001preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20221001preview:AwsS3DataConnector":
		r = &AwsS3DataConnector{}
	case "azure-native:securityinsights/v20221001preview:Bookmark":
		r = &Bookmark{}
	case "azure-native:securityinsights/v20221001preview:BookmarkRelation":
		r = &BookmarkRelation{}
	case "azure-native:securityinsights/v20221001preview:CodelessApiPollingDataConnector":
		r = &CodelessApiPollingDataConnector{}
	case "azure-native:securityinsights/v20221001preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20221001preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20221001preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20221001preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20221001preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20221001preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20221001preview:FileImport":
		r = &FileImport{}
	case "azure-native:securityinsights/v20221001preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20221001preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20221001preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20221001preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20221001preview:IoTDataConnector":
		r = &IoTDataConnector{}
	case "azure-native:securityinsights/v20221001preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20221001preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20221001preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20221001preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20221001preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20221001preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20221001preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20221001preview:NrtAlertRule":
		r = &NrtAlertRule{}
	case "azure-native:securityinsights/v20221001preview:Office365ProjectDataConnector":
		r = &Office365ProjectDataConnector{}
	case "azure-native:securityinsights/v20221001preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20221001preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20221001preview:OfficeIRMDataConnector":
		r = &OfficeIRMDataConnector{}
	case "azure-native:securityinsights/v20221001preview:OfficePowerBIDataConnector":
		r = &OfficePowerBIDataConnector{}
	case "azure-native:securityinsights/v20221001preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20221001preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20221001preview:SecurityMLAnalyticsSetting":
		r = &SecurityMLAnalyticsSetting{}
	case "azure-native:securityinsights/v20221001preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20221001preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20221001preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20221001preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20221001preview:ThreatIntelligenceIndicator":
		r = &ThreatIntelligenceIndicator{}
	case "azure-native:securityinsights/v20221001preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20221001preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20221001preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20221001preview:WatchlistItem":
		r = &WatchlistItem{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := pulumiazurenativesdk.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"securityinsights/v20221001preview",
		&module{version},
	)
}