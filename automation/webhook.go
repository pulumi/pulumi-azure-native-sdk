// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the webhook type.
//
// Uses Azure REST API version 2023-05-15-preview. In version 2.x of the Azure Native provider, it used API version 2015-10-31.
//
// Other available API versions: 2015-10-31, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Webhook struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the expiry time.
	ExpiryTime pulumi.StringPtrOutput `pulumi:"expiryTime"`
	// Gets or sets the value of the enabled flag of the webhook.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Gets or sets the last invoked time.
	LastInvokedTime pulumi.StringPtrOutput `pulumi:"lastInvokedTime"`
	// Details of the user who last modified the Webhook
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the parameters of the job that is created when the webhook calls the runbook it is associated with.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Gets or sets the name of the hybrid worker group the webhook job will run on.
	RunOn pulumi.StringPtrOutput `pulumi:"runOn"`
	// Gets or sets the runbook the webhook is associated with.
	Runbook RunbookAssociationPropertyResponsePtrOutput `pulumi:"runbook"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the webhook uri.
	Uri pulumi.StringPtrOutput `pulumi:"uri"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20151031:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20241023:Webhook"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Webhook
	err := ctx.RegisterResource("azure-native:automation:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("azure-native:automation:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
}

type WebhookState struct {
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the expiry time.
	ExpiryTime *string `pulumi:"expiryTime"`
	// Gets or sets the value of the enabled flag of webhook.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Gets or sets the name of the webhook.
	Name string `pulumi:"name"`
	// Gets or sets the parameters of the job.
	Parameters map[string]string `pulumi:"parameters"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the name of the hybrid worker group the webhook job will run on.
	RunOn *string `pulumi:"runOn"`
	// Gets or sets the runbook.
	Runbook *RunbookAssociationProperty `pulumi:"runbook"`
	// Gets or sets the uri.
	Uri *string `pulumi:"uri"`
	// The webhook name.
	WebhookName *string `pulumi:"webhookName"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the expiry time.
	ExpiryTime pulumi.StringPtrInput
	// Gets or sets the value of the enabled flag of webhook.
	IsEnabled pulumi.BoolPtrInput
	// Gets or sets the name of the webhook.
	Name pulumi.StringInput
	// Gets or sets the parameters of the job.
	Parameters pulumi.StringMapInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the name of the hybrid worker group the webhook job will run on.
	RunOn pulumi.StringPtrInput
	// Gets or sets the runbook.
	Runbook RunbookAssociationPropertyPtrInput
	// Gets or sets the uri.
	Uri pulumi.StringPtrInput
	// The webhook name.
	WebhookName pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

// The Azure API version of the resource.
func (o WebhookOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the creation time.
func (o WebhookOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o WebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the expiry time.
func (o WebhookOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the value of the enabled flag of the webhook.
func (o WebhookOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last invoked time.
func (o WebhookOutput) LastInvokedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.LastInvokedTime }).(pulumi.StringPtrOutput)
}

// Details of the user who last modified the Webhook
func (o WebhookOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// Gets or sets the last modified time.
func (o WebhookOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the parameters of the job that is created when the webhook calls the runbook it is associated with.
func (o WebhookOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// Gets or sets the name of the hybrid worker group the webhook job will run on.
func (o WebhookOutput) RunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.RunOn }).(pulumi.StringPtrOutput)
}

// Gets or sets the runbook the webhook is associated with.
func (o WebhookOutput) Runbook() RunbookAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *Webhook) RunbookAssociationPropertyResponsePtrOutput { return v.Runbook }).(RunbookAssociationPropertyResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WebhookOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Webhook) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WebhookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the webhook uri.
func (o WebhookOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebhookOutput{})
}
