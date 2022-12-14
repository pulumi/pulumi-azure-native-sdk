// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

// The owner type for the governance rule owner source
type GovernanceRuleOwnerSourceType string

const (
	// The rule source type defined using resource tag
	GovernanceRuleOwnerSourceTypeByTag = GovernanceRuleOwnerSourceType("ByTag")
	// The rule source type defined manually
	GovernanceRuleOwnerSourceTypeManually = GovernanceRuleOwnerSourceType("Manually")
)

// The governance rule source, what the rule affects, e.g. Assessments
type GovernanceRuleSourceResourceType string

const (
	// The source of the governance rule is assessments
	GovernanceRuleSourceResourceTypeAssessments = GovernanceRuleSourceResourceType("Assessments")
)

// The rule type of the governance rule, defines the source of the rule e.g. Integrated
type GovernanceRuleType string

const (
	// The source of the rule type definition is integrated
	GovernanceRuleTypeIntegrated = GovernanceRuleType("Integrated")
	// The source of the rule type definition is ServiceNow
	GovernanceRuleTypeServiceNow = GovernanceRuleType("ServiceNow")
)

func init() {
}
