// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
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
	case "azure-native:security:APICollection":
		r = &APICollection{}
	case "azure-native:security:APICollectionByAzureApiManagementService":
		r = &APICollectionByAzureApiManagementService{}
	case "azure-native:security:AdvancedThreatProtection":
		r = &AdvancedThreatProtection{}
	case "azure-native:security:AlertsSuppressionRule":
		r = &AlertsSuppressionRule{}
	case "azure-native:security:Application":
		r = &Application{}
	case "azure-native:security:Assessment":
		r = &Assessment{}
	case "azure-native:security:AssessmentMetadataInSubscription":
		r = &AssessmentMetadataInSubscription{}
	case "azure-native:security:AssessmentsMetadataSubscription":
		r = &AssessmentsMetadataSubscription{}
	case "azure-native:security:Assignment":
		r = &Assignment{}
	case "azure-native:security:Automation":
		r = &Automation{}
	case "azure-native:security:AzureServersSetting":
		r = &AzureServersSetting{}
	case "azure-native:security:Connector":
		r = &Connector{}
	case "azure-native:security:CustomAssessmentAutomation":
		r = &CustomAssessmentAutomation{}
	case "azure-native:security:CustomEntityStoreAssignment":
		r = &CustomEntityStoreAssignment{}
	case "azure-native:security:CustomRecommendation":
		r = &CustomRecommendation{}
	case "azure-native:security:DefenderForStorage":
		r = &DefenderForStorage{}
	case "azure-native:security:DevOpsConfiguration":
		r = &DevOpsConfiguration{}
	case "azure-native:security:DeviceSecurityGroup":
		r = &DeviceSecurityGroup{}
	case "azure-native:security:GovernanceAssignment":
		r = &GovernanceAssignment{}
	case "azure-native:security:GovernanceRule":
		r = &GovernanceRule{}
	case "azure-native:security:IotSecuritySolution":
		r = &IotSecuritySolution{}
	case "azure-native:security:JitNetworkAccessPolicy":
		r = &JitNetworkAccessPolicy{}
	case "azure-native:security:Pricing":
		r = &Pricing{}
	case "azure-native:security:SecurityConnector":
		r = &SecurityConnector{}
	case "azure-native:security:SecurityConnectorApplication":
		r = &SecurityConnectorApplication{}
	case "azure-native:security:SecurityContact":
		r = &SecurityContact{}
	case "azure-native:security:SecurityOperator":
		r = &SecurityOperator{}
	case "azure-native:security:SecurityStandard":
		r = &SecurityStandard{}
	case "azure-native:security:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:security:SqlVulnerabilityAssessmentBaselineRule":
		r = &SqlVulnerabilityAssessmentBaselineRule{}
	case "azure-native:security:Standard":
		r = &Standard{}
	case "azure-native:security:StandardAssignment":
		r = &StandardAssignment{}
	case "azure-native:security:WorkspaceSetting":
		r = &WorkspaceSetting{}
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
		"security",
		&module{version},
	)
}
