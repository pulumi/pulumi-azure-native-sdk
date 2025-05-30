// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An action group resource.
//
// Uses Azure REST API version 2024-10-01-preview.
//
// Other available API versions: 2018-03-01, 2022-06-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ActionGroup struct {
	pulumi.CustomResourceState

	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers ArmRoleReceiverResponseArrayOutput `pulumi:"armRoleReceivers"`
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayOutput `pulumi:"automationRunbookReceivers"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverResponseArrayOutput `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverResponseArrayOutput `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverResponseArrayOutput `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The list of event hub receivers that are part of this action group.
	EventHubReceivers EventHubReceiverResponseArrayOutput `pulumi:"eventHubReceivers"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringOutput `pulumi:"groupShortName"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The list of incident receivers that are part of this action group.
	IncidentReceivers IncidentReceiverResponseArrayOutput `pulumi:"incidentReceivers"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverResponseArrayOutput `pulumi:"itsmReceivers"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverResponseArrayOutput `pulumi:"logicAppReceivers"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverResponseArrayOutput `pulumi:"smsReceivers"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverResponseArrayOutput `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverResponseArrayOutput `pulumi:"webhookReceivers"`
}

// NewActionGroup registers a new resource with the given unique name, arguments, and options.
func NewActionGroup(ctx *pulumi.Context,
	name string, args *ActionGroupArgs, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupShortName == nil {
		return nil, errors.New("invalid value for required argument 'GroupShortName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Enabled == nil {
		args.Enabled = pulumi.Bool(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20230101:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20230901preview:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20241001preview:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20170401:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20180301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20180901:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20190301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20190601:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20210901:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20220401:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20220601:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20230101:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20230901preview:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20241001preview:ActionGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ActionGroup
	err := ctx.RegisterResource("azure-native:monitor:ActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionGroup gets an existing ActionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionGroupState, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	var resource ActionGroup
	err := ctx.ReadResource("azure-native:monitor:ActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionGroup resources.
type actionGroupState struct {
}

type ActionGroupState struct {
}

func (ActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupState)(nil)).Elem()
}

type actionGroupArgs struct {
	// The name of the action group.
	ActionGroupName *string `pulumi:"actionGroupName"`
	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers []ArmRoleReceiver `pulumi:"armRoleReceivers"`
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers []AutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers []AzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers []AzureFunctionReceiver `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers []EmailReceiver `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled bool `pulumi:"enabled"`
	// The list of event hub receivers that are part of this action group.
	EventHubReceivers []EventHubReceiver `pulumi:"eventHubReceivers"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName string `pulumi:"groupShortName"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The list of incident receivers that are part of this action group.
	IncidentReceivers []IncidentReceiver `pulumi:"incidentReceivers"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers []ItsmReceiver `pulumi:"itsmReceivers"`
	// Resource location
	Location *string `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers []LogicAppReceiver `pulumi:"logicAppReceivers"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers []SmsReceiver `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers []VoiceReceiver `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers []WebhookReceiver `pulumi:"webhookReceivers"`
}

// The set of arguments for constructing a ActionGroup resource.
type ActionGroupArgs struct {
	// The name of the action group.
	ActionGroupName pulumi.StringPtrInput
	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers ArmRoleReceiverArrayInput
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverArrayInput
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverArrayInput
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverArrayInput
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverArrayInput
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolInput
	// The list of event hub receivers that are part of this action group.
	EventHubReceivers EventHubReceiverArrayInput
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringInput
	// Managed service identity (system assigned and/or user assigned identities)
	Identity ManagedServiceIdentityPtrInput
	// The list of incident receivers that are part of this action group.
	IncidentReceivers IncidentReceiverArrayInput
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverArrayInput
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverArrayInput
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupArgs)(nil)).Elem()
}

type ActionGroupInput interface {
	pulumi.Input

	ToActionGroupOutput() ActionGroupOutput
	ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput
}

func (*ActionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroup)(nil)).Elem()
}

func (i *ActionGroup) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i *ActionGroup) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}

type ActionGroupOutput struct{ *pulumi.OutputState }

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroup)(nil)).Elem()
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
func (o ActionGroupOutput) ArmRoleReceivers() ArmRoleReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) ArmRoleReceiverResponseArrayOutput { return v.ArmRoleReceivers }).(ArmRoleReceiverResponseArrayOutput)
}

// The list of AutomationRunbook receivers that are part of this action group.
func (o ActionGroupOutput) AutomationRunbookReceivers() AutomationRunbookReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) AutomationRunbookReceiverResponseArrayOutput { return v.AutomationRunbookReceivers }).(AutomationRunbookReceiverResponseArrayOutput)
}

// The Azure API version of the resource.
func (o ActionGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The list of AzureAppPush receivers that are part of this action group.
func (o ActionGroupOutput) AzureAppPushReceivers() AzureAppPushReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) AzureAppPushReceiverResponseArrayOutput { return v.AzureAppPushReceivers }).(AzureAppPushReceiverResponseArrayOutput)
}

// The list of azure function receivers that are part of this action group.
func (o ActionGroupOutput) AzureFunctionReceivers() AzureFunctionReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) AzureFunctionReceiverResponseArrayOutput { return v.AzureFunctionReceivers }).(AzureFunctionReceiverResponseArrayOutput)
}

// The list of email receivers that are part of this action group.
func (o ActionGroupOutput) EmailReceivers() EmailReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) EmailReceiverResponseArrayOutput { return v.EmailReceivers }).(EmailReceiverResponseArrayOutput)
}

// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
func (o ActionGroupOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The list of event hub receivers that are part of this action group.
func (o ActionGroupOutput) EventHubReceivers() EventHubReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) EventHubReceiverResponseArrayOutput { return v.EventHubReceivers }).(EventHubReceiverResponseArrayOutput)
}

// The short name of the action group. This will be used in SMS messages.
func (o ActionGroupOutput) GroupShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.GroupShortName }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o ActionGroupOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ActionGroup) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The list of incident receivers that are part of this action group.
func (o ActionGroupOutput) IncidentReceivers() IncidentReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) IncidentReceiverResponseArrayOutput { return v.IncidentReceivers }).(IncidentReceiverResponseArrayOutput)
}

// The list of ITSM receivers that are part of this action group.
func (o ActionGroupOutput) ItsmReceivers() ItsmReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) ItsmReceiverResponseArrayOutput { return v.ItsmReceivers }).(ItsmReceiverResponseArrayOutput)
}

// Resource location
func (o ActionGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The list of logic app receivers that are part of this action group.
func (o ActionGroupOutput) LogicAppReceivers() LogicAppReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) LogicAppReceiverResponseArrayOutput { return v.LogicAppReceivers }).(LogicAppReceiverResponseArrayOutput)
}

// Azure resource name
func (o ActionGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of SMS receivers that are part of this action group.
func (o ActionGroupOutput) SmsReceivers() SmsReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) SmsReceiverResponseArrayOutput { return v.SmsReceivers }).(SmsReceiverResponseArrayOutput)
}

// Resource tags
func (o ActionGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o ActionGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of voice receivers that are part of this action group.
func (o ActionGroupOutput) VoiceReceivers() VoiceReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) VoiceReceiverResponseArrayOutput { return v.VoiceReceivers }).(VoiceReceiverResponseArrayOutput)
}

// The list of webhook receivers that are part of this action group.
func (o ActionGroupOutput) WebhookReceivers() WebhookReceiverResponseArrayOutput {
	return o.ApplyT(func(v *ActionGroup) WebhookReceiverResponseArrayOutput { return v.WebhookReceivers }).(WebhookReceiverResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionGroupOutput{})
}
