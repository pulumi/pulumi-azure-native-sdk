// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security GovernanceRule over a given scope
type SecurityConnectorGovernanceRule struct {
	pulumi.CustomResourceState

	// description of the governanceRule
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// display name of the governanceRule
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification GovernanceRuleEmailNotificationResponsePtrOutput `pulumi:"governanceEmailNotification"`
	// Defines whether the rule is active/inactive
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// Defines whether there is a grace period on the governance rule
	IsGracePeriod pulumi.BoolPtrOutput `pulumi:"isGracePeriod"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The Owner source for the governance rule - e.g. Manually by user@contoso.com - see example
	OwnerSource GovernanceRuleOwnerSourceResponseOutput `pulumi:"ownerSource"`
	// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
	RemediationTimeframe pulumi.StringPtrOutput `pulumi:"remediationTimeframe"`
	// The governance rule priority, priority to the lower number. Rules with the same priority on the same subscription will not be allowed
	RulePriority pulumi.IntOutput `pulumi:"rulePriority"`
	// The rule type of the governance rule, defines the source of the rule e.g. Integrated
	RuleType pulumi.StringOutput `pulumi:"ruleType"`
	// The governance rule source, what the rule affects, e.g. Assessments
	SourceResourceType pulumi.StringOutput `pulumi:"sourceResourceType"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityConnectorGovernanceRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityConnectorGovernanceRule(ctx *pulumi.Context,
	name string, args *SecurityConnectorGovernanceRuleArgs, opts ...pulumi.ResourceOption) (*SecurityConnectorGovernanceRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.OwnerSource == nil {
		return nil, errors.New("invalid value for required argument 'OwnerSource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RulePriority == nil {
		return nil, errors.New("invalid value for required argument 'RulePriority'")
	}
	if args.RuleType == nil {
		return nil, errors.New("invalid value for required argument 'RuleType'")
	}
	if args.SecurityConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'SecurityConnectorName'")
	}
	if args.SourceResourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceResourceType'")
	}
	var resource SecurityConnectorGovernanceRule
	err := ctx.RegisterResource("azure-native:security/v20220101preview:SecurityConnectorGovernanceRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityConnectorGovernanceRule gets an existing SecurityConnectorGovernanceRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityConnectorGovernanceRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConnectorGovernanceRuleState, opts ...pulumi.ResourceOption) (*SecurityConnectorGovernanceRule, error) {
	var resource SecurityConnectorGovernanceRule
	err := ctx.ReadResource("azure-native:security/v20220101preview:SecurityConnectorGovernanceRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityConnectorGovernanceRule resources.
type securityConnectorGovernanceRuleState struct {
}

type SecurityConnectorGovernanceRuleState struct {
}

func (SecurityConnectorGovernanceRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorGovernanceRuleState)(nil)).Elem()
}

type securityConnectorGovernanceRuleArgs struct {
	// description of the governanceRule
	Description *string `pulumi:"description"`
	// display name of the governanceRule
	DisplayName string `pulumi:"displayName"`
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification *GovernanceRuleEmailNotification `pulumi:"governanceEmailNotification"`
	// Defines whether the rule is active/inactive
	IsDisabled *bool `pulumi:"isDisabled"`
	// Defines whether there is a grace period on the governance rule
	IsGracePeriod *bool `pulumi:"isGracePeriod"`
	// The Owner source for the governance rule - e.g. Manually by user@contoso.com - see example
	OwnerSource GovernanceRuleOwnerSource `pulumi:"ownerSource"`
	// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
	RemediationTimeframe *string `pulumi:"remediationTimeframe"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security GovernanceRule key - unique key for the standard GovernanceRule
	RuleId *string `pulumi:"ruleId"`
	// The governance rule priority, priority to the lower number. Rules with the same priority on the same subscription will not be allowed
	RulePriority int `pulumi:"rulePriority"`
	// The rule type of the governance rule, defines the source of the rule e.g. Integrated
	RuleType string `pulumi:"ruleType"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
	// The governance rule source, what the rule affects, e.g. Assessments
	SourceResourceType string `pulumi:"sourceResourceType"`
}

// The set of arguments for constructing a SecurityConnectorGovernanceRule resource.
type SecurityConnectorGovernanceRuleArgs struct {
	// description of the governanceRule
	Description pulumi.StringPtrInput
	// display name of the governanceRule
	DisplayName pulumi.StringInput
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification GovernanceRuleEmailNotificationPtrInput
	// Defines whether the rule is active/inactive
	IsDisabled pulumi.BoolPtrInput
	// Defines whether there is a grace period on the governance rule
	IsGracePeriod pulumi.BoolPtrInput
	// The Owner source for the governance rule - e.g. Manually by user@contoso.com - see example
	OwnerSource GovernanceRuleOwnerSourceInput
	// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
	RemediationTimeframe pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The security GovernanceRule key - unique key for the standard GovernanceRule
	RuleId pulumi.StringPtrInput
	// The governance rule priority, priority to the lower number. Rules with the same priority on the same subscription will not be allowed
	RulePriority pulumi.IntInput
	// The rule type of the governance rule, defines the source of the rule e.g. Integrated
	RuleType pulumi.StringInput
	// The security connector name.
	SecurityConnectorName pulumi.StringInput
	// The governance rule source, what the rule affects, e.g. Assessments
	SourceResourceType pulumi.StringInput
}

func (SecurityConnectorGovernanceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorGovernanceRuleArgs)(nil)).Elem()
}

type SecurityConnectorGovernanceRuleInput interface {
	pulumi.Input

	ToSecurityConnectorGovernanceRuleOutput() SecurityConnectorGovernanceRuleOutput
	ToSecurityConnectorGovernanceRuleOutputWithContext(ctx context.Context) SecurityConnectorGovernanceRuleOutput
}

func (*SecurityConnectorGovernanceRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorGovernanceRule)(nil)).Elem()
}

func (i *SecurityConnectorGovernanceRule) ToSecurityConnectorGovernanceRuleOutput() SecurityConnectorGovernanceRuleOutput {
	return i.ToSecurityConnectorGovernanceRuleOutputWithContext(context.Background())
}

func (i *SecurityConnectorGovernanceRule) ToSecurityConnectorGovernanceRuleOutputWithContext(ctx context.Context) SecurityConnectorGovernanceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorGovernanceRuleOutput)
}

type SecurityConnectorGovernanceRuleOutput struct{ *pulumi.OutputState }

func (SecurityConnectorGovernanceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorGovernanceRule)(nil)).Elem()
}

func (o SecurityConnectorGovernanceRuleOutput) ToSecurityConnectorGovernanceRuleOutput() SecurityConnectorGovernanceRuleOutput {
	return o
}

func (o SecurityConnectorGovernanceRuleOutput) ToSecurityConnectorGovernanceRuleOutputWithContext(ctx context.Context) SecurityConnectorGovernanceRuleOutput {
	return o
}

// description of the governanceRule
func (o SecurityConnectorGovernanceRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// display name of the governanceRule
func (o SecurityConnectorGovernanceRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
func (o SecurityConnectorGovernanceRuleOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) GovernanceRuleEmailNotificationResponsePtrOutput {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

// Defines whether the rule is active/inactive
func (o SecurityConnectorGovernanceRuleOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// Defines whether there is a grace period on the governance rule
func (o SecurityConnectorGovernanceRuleOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.BoolPtrOutput { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

// Resource name
func (o SecurityConnectorGovernanceRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Owner source for the governance rule - e.g. Manually by user@contoso.com - see example
func (o SecurityConnectorGovernanceRuleOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) GovernanceRuleOwnerSourceResponseOutput { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
func (o SecurityConnectorGovernanceRuleOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringPtrOutput { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

// The governance rule priority, priority to the lower number. Rules with the same priority on the same subscription will not be allowed
func (o SecurityConnectorGovernanceRuleOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.IntOutput { return v.RulePriority }).(pulumi.IntOutput)
}

// The rule type of the governance rule, defines the source of the rule e.g. Integrated
func (o SecurityConnectorGovernanceRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

// The governance rule source, what the rule affects, e.g. Assessments
func (o SecurityConnectorGovernanceRuleOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.SourceResourceType }).(pulumi.StringOutput)
}

// Resource type
func (o SecurityConnectorGovernanceRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityConnectorGovernanceRuleOutput{})
}