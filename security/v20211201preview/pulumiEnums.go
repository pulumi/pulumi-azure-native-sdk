// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

// The multi cloud resource's cloud name.
type CloudName string

const (
	CloudNameAzure  = CloudName("Azure")
	CloudNameAWS    = CloudName("AWS")
	CloudNameGCP    = CloudName("GCP")
	CloudNameGithub = CloudName("Github")
)

// The type of the environment data.
type EnvironmentType string

const (
	EnvironmentTypeAwsAccount  = EnvironmentType("AwsAccount")
	EnvironmentTypeGcpProject  = EnvironmentType("GcpProject")
	EnvironmentTypeGithubScope = EnvironmentType("GithubScope")
)

// The type of the security offering.
type OfferingType string

const (
	OfferingTypeCspmMonitorAws           = OfferingType("CspmMonitorAws")
	OfferingTypeDefenderForContainersAws = OfferingType("DefenderForContainersAws")
	OfferingTypeDefenderForServersAws    = OfferingType("DefenderForServersAws")
	OfferingTypeInformationProtectionAws = OfferingType("InformationProtectionAws")
	OfferingTypeCspmMonitorGcp           = OfferingType("CspmMonitorGcp")
	OfferingTypeCspmMonitorGithub        = OfferingType("CspmMonitorGithub")
	OfferingTypeDefenderForServersGcp    = OfferingType("DefenderForServersGcp")
	OfferingTypeDefenderForContainersGcp = OfferingType("DefenderForContainersGcp")
)

// The multi cloud account's membership type in the organization
type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       = OrganizationMembershipType("Member")
	OrganizationMembershipTypeOrganization = OrganizationMembershipType("Organization")
)

// The available sub plans
type SubPlan string

const (
	SubPlanP1 = SubPlan("P1")
	SubPlanP2 = SubPlan("P2")
)

// The Vulnerability Assessment solution to be provisioned. Can be either 'TVM' or 'Qualys'
type Type string

const (
	TypeQualys = Type("Qualys")
	TypeTVM    = Type("TVM")
)

func init() {
}