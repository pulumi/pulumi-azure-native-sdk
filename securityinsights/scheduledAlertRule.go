// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents scheduled alert rule.
// API Version: 2020-01-01.
type ScheduledAlertRule struct {
	pulumi.CustomResourceState

	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringPtrOutput `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the alert rule
	// Expected value is 'Scheduled'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert rule has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The query that creates alerts for this rule.
	Query pulumi.StringOutput `pulumi:"query"`
	// The frequency (in ISO 8601 duration format) for this alert rule to run.
	QueryFrequency pulumi.StringOutput `pulumi:"queryFrequency"`
	// The period (in ISO 8601 duration format) that this alert rule looks at.
	QueryPeriod pulumi.StringOutput `pulumi:"queryPeriod"`
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration pulumi.StringOutput `pulumi:"suppressionDuration"`
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled pulumi.BoolOutput `pulumi:"suppressionEnabled"`
	// The tactics of the alert rule
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// The operation against the threshold that triggers alert rule.
	TriggerOperator pulumi.StringOutput `pulumi:"triggerOperator"`
	// The threshold triggers this alert rule.
	TriggerThreshold pulumi.IntOutput `pulumi:"triggerThreshold"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScheduledAlertRule registers a new resource with the given unique name, arguments, and options.
func NewScheduledAlertRule(ctx *pulumi.Context,
	name string, args *ScheduledAlertRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.QueryFrequency == nil {
		return nil, errors.New("invalid value for required argument 'QueryFrequency'")
	}
	if args.QueryPeriod == nil {
		return nil, errors.New("invalid value for required argument 'QueryPeriod'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.SuppressionDuration == nil {
		return nil, errors.New("invalid value for required argument 'SuppressionDuration'")
	}
	if args.SuppressionEnabled == nil {
		return nil, errors.New("invalid value for required argument 'SuppressionEnabled'")
	}
	if args.TriggerOperator == nil {
		return nil, errors.New("invalid value for required argument 'TriggerOperator'")
	}
	if args.TriggerThreshold == nil {
		return nil, errors.New("invalid value for required argument 'TriggerThreshold'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Scheduled")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ScheduledAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ScheduledAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights:ScheduledAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledAlertRule gets an existing ScheduledAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledAlertRuleState, opts ...pulumi.ResourceOption) (*ScheduledAlertRule, error) {
	var resource ScheduledAlertRule
	err := ctx.ReadResource("azure-native:securityinsights:ScheduledAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledAlertRule resources.
type scheduledAlertRuleState struct {
}

type ScheduledAlertRuleState struct {
}

func (ScheduledAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledAlertRuleState)(nil)).Elem()
}

type scheduledAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName *string `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description *string `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// The kind of the alert rule
	// Expected value is 'Scheduled'.
	Kind string `pulumi:"kind"`
	// The query that creates alerts for this rule.
	Query string `pulumi:"query"`
	// The frequency (in ISO 8601 duration format) for this alert rule to run.
	QueryFrequency string `pulumi:"queryFrequency"`
	// The period (in ISO 8601 duration format) that this alert rule looks at.
	QueryPeriod string `pulumi:"queryPeriod"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// The severity for alerts created by this alert rule.
	Severity string `pulumi:"severity"`
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration string `pulumi:"suppressionDuration"`
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled bool `pulumi:"suppressionEnabled"`
	// The tactics of the alert rule
	Tactics []string `pulumi:"tactics"`
	// The operation against the threshold that triggers alert rule.
	TriggerOperator TriggerOperator `pulumi:"triggerOperator"`
	// The threshold triggers this alert rule.
	TriggerThreshold int `pulumi:"triggerThreshold"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ScheduledAlertRule resource.
type ScheduledAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringPtrInput
	// The description of the alert rule.
	Description pulumi.StringPtrInput
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringInput
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolInput
	// The kind of the alert rule
	// Expected value is 'Scheduled'.
	Kind pulumi.StringInput
	// The query that creates alerts for this rule.
	Query pulumi.StringInput
	// The frequency (in ISO 8601 duration format) for this alert rule to run.
	QueryFrequency pulumi.StringInput
	// The period (in ISO 8601 duration format) that this alert rule looks at.
	QueryPeriod pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringInput
	// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
	SuppressionDuration pulumi.StringInput
	// Determines whether the suppression for this alert rule is enabled or disabled.
	SuppressionEnabled pulumi.BoolInput
	// The tactics of the alert rule
	Tactics pulumi.StringArrayInput
	// The operation against the threshold that triggers alert rule.
	TriggerOperator TriggerOperatorInput
	// The threshold triggers this alert rule.
	TriggerThreshold pulumi.IntInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ScheduledAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledAlertRuleArgs)(nil)).Elem()
}

type ScheduledAlertRuleInput interface {
	pulumi.Input

	ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput
	ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput
}

func (*ScheduledAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAlertRule)(nil)).Elem()
}

func (i *ScheduledAlertRule) ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput {
	return i.ToScheduledAlertRuleOutputWithContext(context.Background())
}

func (i *ScheduledAlertRule) ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledAlertRuleOutput)
}

type ScheduledAlertRuleOutput struct{ *pulumi.OutputState }

func (ScheduledAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAlertRule)(nil)).Elem()
}

func (o ScheduledAlertRuleOutput) ToScheduledAlertRuleOutput() ScheduledAlertRuleOutput {
	return o
}

func (o ScheduledAlertRuleOutput) ToScheduledAlertRuleOutputWithContext(ctx context.Context) ScheduledAlertRuleOutput {
	return o
}

// The Name of the alert rule template used to create this rule.
func (o ScheduledAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

// The description of the alert rule.
func (o ScheduledAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for alerts created by this alert rule.
func (o ScheduledAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o ScheduledAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o ScheduledAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the alert rule
// Expected value is 'Scheduled'.
func (o ScheduledAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert rule has been modified.
func (o ScheduledAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// Azure resource name
func (o ScheduledAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The query that creates alerts for this rule.
func (o ScheduledAlertRuleOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// The frequency (in ISO 8601 duration format) for this alert rule to run.
func (o ScheduledAlertRuleOutput) QueryFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.QueryFrequency }).(pulumi.StringOutput)
}

// The period (in ISO 8601 duration format) that this alert rule looks at.
func (o ScheduledAlertRuleOutput) QueryPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.QueryPeriod }).(pulumi.StringOutput)
}

// The severity for alerts created by this alert rule.
func (o ScheduledAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
func (o ScheduledAlertRuleOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.SuppressionDuration }).(pulumi.StringOutput)
}

// Determines whether the suppression for this alert rule is enabled or disabled.
func (o ScheduledAlertRuleOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.BoolOutput { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

// The tactics of the alert rule
func (o ScheduledAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The operation against the threshold that triggers alert rule.
func (o ScheduledAlertRuleOutput) TriggerOperator() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.TriggerOperator }).(pulumi.StringOutput)
}

// The threshold triggers this alert rule.
func (o ScheduledAlertRuleOutput) TriggerThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.IntOutput { return v.TriggerThreshold }).(pulumi.IntOutput)
}

// Azure resource type
func (o ScheduledAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledAlertRuleOutput{})
}