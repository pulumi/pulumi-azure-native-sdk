// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a specific governance rule for the requested scope by ruleId
//
// Uses Azure REST API version 2022-01-01-preview.
func LookupGovernanceRule(ctx *pulumi.Context, args *LookupGovernanceRuleArgs, opts ...pulumi.InvokeOption) (*LookupGovernanceRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGovernanceRuleResult
	err := ctx.Invoke("azure-native:security:getGovernanceRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGovernanceRuleArgs struct {
	// The governance rule key - unique key for the standard governance rule (GUID)
	RuleId string `pulumi:"ruleId"`
	// The scope of the Governance rules. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
	Scope string `pulumi:"scope"`
}

// Governance rule over a given scope
type LookupGovernanceRuleResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Description of the governance rule
	Description *string `pulumi:"description"`
	// Display name of the governance rule
	DisplayName string `pulumi:"displayName"`
	// Excluded scopes, filter out the descendants of the scope (on management scopes)
	ExcludedScopes []string `pulumi:"excludedScopes"`
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification *GovernanceRuleEmailNotificationResponse `pulumi:"governanceEmailNotification"`
	// Resource Id
	Id string `pulumi:"id"`
	// Defines whether the rule is management scope rule (master connector as a single scope or management scope)
	IncludeMemberScopes *bool `pulumi:"includeMemberScopes"`
	// Defines whether the rule is active/inactive
	IsDisabled *bool `pulumi:"isDisabled"`
	// Defines whether there is a grace period on the governance rule
	IsGracePeriod *bool `pulumi:"isGracePeriod"`
	// The governance rule metadata
	Metadata *GovernanceRuleMetadataResponse `pulumi:"metadata"`
	// Resource name
	Name string `pulumi:"name"`
	// The owner source for the governance rule - e.g. Manually by user@contoso.com - see example
	OwnerSource GovernanceRuleOwnerSourceResponse `pulumi:"ownerSource"`
	// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
	RemediationTimeframe *string `pulumi:"remediationTimeframe"`
	// The governance rule priority, priority to the lower number. Rules with the same priority on the same scope will not be allowed
	RulePriority int `pulumi:"rulePriority"`
	// The rule type of the governance rule, defines the source of the rule e.g. Integrated
	RuleType string `pulumi:"ruleType"`
	// The governance rule source, what the rule affects, e.g. Assessments
	SourceResourceType string `pulumi:"sourceResourceType"`
	// The tenantId (GUID)
	TenantId string `pulumi:"tenantId"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupGovernanceRuleOutput(ctx *pulumi.Context, args LookupGovernanceRuleOutputArgs, opts ...pulumi.InvokeOption) LookupGovernanceRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGovernanceRuleResultOutput, error) {
			args := v.(LookupGovernanceRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:security:getGovernanceRule", args, LookupGovernanceRuleResultOutput{}, options).(LookupGovernanceRuleResultOutput), nil
		}).(LookupGovernanceRuleResultOutput)
}

type LookupGovernanceRuleOutputArgs struct {
	// The governance rule key - unique key for the standard governance rule (GUID)
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// The scope of the Governance rules. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupGovernanceRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceRuleArgs)(nil)).Elem()
}

// Governance rule over a given scope
type LookupGovernanceRuleResultOutput struct{ *pulumi.OutputState }

func (LookupGovernanceRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceRuleResult)(nil)).Elem()
}

func (o LookupGovernanceRuleResultOutput) ToLookupGovernanceRuleResultOutput() LookupGovernanceRuleResultOutput {
	return o
}

func (o LookupGovernanceRuleResultOutput) ToLookupGovernanceRuleResultOutputWithContext(ctx context.Context) LookupGovernanceRuleResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupGovernanceRuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Description of the governance rule
func (o LookupGovernanceRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of the governance rule
func (o LookupGovernanceRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Excluded scopes, filter out the descendants of the scope (on management scopes)
func (o LookupGovernanceRuleResultOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) []string { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
func (o LookupGovernanceRuleResultOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *GovernanceRuleEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

// Resource Id
func (o LookupGovernanceRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Defines whether the rule is management scope rule (master connector as a single scope or management scope)
func (o LookupGovernanceRuleResultOutput) IncludeMemberScopes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *bool { return v.IncludeMemberScopes }).(pulumi.BoolPtrOutput)
}

// Defines whether the rule is active/inactive
func (o LookupGovernanceRuleResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// Defines whether there is a grace period on the governance rule
func (o LookupGovernanceRuleResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

// The governance rule metadata
func (o LookupGovernanceRuleResultOutput) Metadata() GovernanceRuleMetadataResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *GovernanceRuleMetadataResponse { return v.Metadata }).(GovernanceRuleMetadataResponsePtrOutput)
}

// Resource name
func (o LookupGovernanceRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The owner source for the governance rule - e.g. Manually by user@contoso.com - see example
func (o LookupGovernanceRuleResultOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) GovernanceRuleOwnerSourceResponse { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
func (o LookupGovernanceRuleResultOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *string { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

// The governance rule priority, priority to the lower number. Rules with the same priority on the same scope will not be allowed
func (o LookupGovernanceRuleResultOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) int { return v.RulePriority }).(pulumi.IntOutput)
}

// The rule type of the governance rule, defines the source of the rule e.g. Integrated
func (o LookupGovernanceRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

// The governance rule source, what the rule affects, e.g. Assessments
func (o LookupGovernanceRuleResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

// The tenantId (GUID)
func (o LookupGovernanceRuleResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type
func (o LookupGovernanceRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGovernanceRuleResultOutput{})
}
