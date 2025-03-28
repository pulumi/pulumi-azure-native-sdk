// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Threat Intelligence alert rule.
type ThreatIntelligenceAlertRule struct {
	pulumi.CustomResourceState

	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringOutput `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the alert rule
	// Expected value is 'ThreatIntelligence'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The severity for alerts created by this alert rule.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// The sub-techniques of the alert rule
	SubTechniques pulumi.StringArrayOutput `pulumi:"subTechniques"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tactics of the alert rule
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// The techniques of the alert rule
	Techniques pulumi.StringArrayOutput `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewThreatIntelligenceAlertRule registers a new resource with the given unique name, arguments, and options.
func NewThreatIntelligenceAlertRule(ctx *pulumi.Context,
	name string, args *ThreatIntelligenceAlertRuleArgs, opts ...pulumi.ResourceOption) (*ThreatIntelligenceAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertRuleTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'AlertRuleTemplateName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("ThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240901:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250301:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:ThreatIntelligenceAlertRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ThreatIntelligenceAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20240401preview:ThreatIntelligenceAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThreatIntelligenceAlertRule gets an existing ThreatIntelligenceAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreatIntelligenceAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelligenceAlertRuleState, opts ...pulumi.ResourceOption) (*ThreatIntelligenceAlertRule, error) {
	var resource ThreatIntelligenceAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20240401preview:ThreatIntelligenceAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThreatIntelligenceAlertRule resources.
type threatIntelligenceAlertRuleState struct {
}

type ThreatIntelligenceAlertRuleState struct {
}

func (ThreatIntelligenceAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceAlertRuleState)(nil)).Elem()
}

type threatIntelligenceAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName string `pulumi:"alertRuleTemplateName"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// The kind of the alert rule
	// Expected value is 'ThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ThreatIntelligenceAlertRule resource.
type ThreatIntelligenceAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringInput
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolInput
	// The kind of the alert rule
	// Expected value is 'ThreatIntelligence'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ThreatIntelligenceAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceAlertRuleArgs)(nil)).Elem()
}

type ThreatIntelligenceAlertRuleInput interface {
	pulumi.Input

	ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput
	ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput
}

func (*ThreatIntelligenceAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceAlertRule)(nil)).Elem()
}

func (i *ThreatIntelligenceAlertRule) ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput {
	return i.ToThreatIntelligenceAlertRuleOutputWithContext(context.Background())
}

func (i *ThreatIntelligenceAlertRule) ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceAlertRuleOutput)
}

type ThreatIntelligenceAlertRuleOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceAlertRule)(nil)).Elem()
}

func (o ThreatIntelligenceAlertRuleOutput) ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput {
	return o
}

func (o ThreatIntelligenceAlertRuleOutput) ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput {
	return o
}

// The Name of the alert rule template used to create this rule.
func (o ThreatIntelligenceAlertRuleOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

// The description of the alert rule.
func (o ThreatIntelligenceAlertRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name for alerts created by this alert rule.
func (o ThreatIntelligenceAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o ThreatIntelligenceAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o ThreatIntelligenceAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the alert rule
// Expected value is 'ThreatIntelligence'.
func (o ThreatIntelligenceAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert has been modified.
func (o ThreatIntelligenceAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o ThreatIntelligenceAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The severity for alerts created by this alert rule.
func (o ThreatIntelligenceAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// The sub-techniques of the alert rule
func (o ThreatIntelligenceAlertRuleOutput) SubTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringArrayOutput { return v.SubTechniques }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ThreatIntelligenceAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the alert rule
func (o ThreatIntelligenceAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the alert rule
func (o ThreatIntelligenceAlertRuleOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ThreatIntelligenceAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelligenceAlertRuleOutput{})
}
