// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databasewatcher

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Concrete proxy resource types can be created by aliasing this type using a specific property type.
//
// Uses Azure REST API version 2024-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-07-19-preview.
//
// Other available API versions: 2024-07-19-preview, 2025-01-02. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databasewatcher [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type AlertRuleResource struct {
	pulumi.CustomResourceState

	// The resource ID of the alert rule resource.
	AlertRuleResourceId pulumi.StringOutput `pulumi:"alertRuleResourceId"`
	// The template ID associated with alert rule resource.
	AlertRuleTemplateId pulumi.StringOutput `pulumi:"alertRuleTemplateId"`
	// The alert rule template version.
	AlertRuleTemplateVersion pulumi.StringOutput `pulumi:"alertRuleTemplateVersion"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The properties with which the alert rule resource was created.
	CreatedWithProperties pulumi.StringOutput `pulumi:"createdWithProperties"`
	// The creation time of the alert rule resource.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the alert rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertRuleResource registers a new resource with the given unique name, arguments, and options.
func NewAlertRuleResource(ctx *pulumi.Context,
	name string, args *AlertRuleResourceArgs, opts ...pulumi.ResourceOption) (*AlertRuleResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertRuleResourceId == nil {
		return nil, errors.New("invalid value for required argument 'AlertRuleResourceId'")
	}
	if args.AlertRuleTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'AlertRuleTemplateId'")
	}
	if args.AlertRuleTemplateVersion == nil {
		return nil, errors.New("invalid value for required argument 'AlertRuleTemplateVersion'")
	}
	if args.CreatedWithProperties == nil {
		return nil, errors.New("invalid value for required argument 'CreatedWithProperties'")
	}
	if args.CreationTime == nil {
		return nil, errors.New("invalid value for required argument 'CreationTime'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WatcherName == nil {
		return nil, errors.New("invalid value for required argument 'WatcherName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databasewatcher/v20240719preview:AlertRuleResource"),
		},
		{
			Type: pulumi.String("azure-native:databasewatcher/v20241001preview:AlertRuleResource"),
		},
		{
			Type: pulumi.String("azure-native:databasewatcher/v20250102:AlertRuleResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AlertRuleResource
	err := ctx.RegisterResource("azure-native:databasewatcher:AlertRuleResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRuleResource gets an existing AlertRuleResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRuleResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleResourceState, opts ...pulumi.ResourceOption) (*AlertRuleResource, error) {
	var resource AlertRuleResource
	err := ctx.ReadResource("azure-native:databasewatcher:AlertRuleResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRuleResource resources.
type alertRuleResourceState struct {
}

type AlertRuleResourceState struct {
}

func (AlertRuleResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleResourceState)(nil)).Elem()
}

type alertRuleResourceArgs struct {
	// The resource ID of the alert rule resource.
	AlertRuleResourceId string `pulumi:"alertRuleResourceId"`
	// The alert rule proxy resource name.
	AlertRuleResourceName *string `pulumi:"alertRuleResourceName"`
	// The template ID associated with alert rule resource.
	AlertRuleTemplateId string `pulumi:"alertRuleTemplateId"`
	// The alert rule template version.
	AlertRuleTemplateVersion string `pulumi:"alertRuleTemplateVersion"`
	// The properties with which the alert rule resource was created.
	CreatedWithProperties string `pulumi:"createdWithProperties"`
	// The creation time of the alert rule resource.
	CreationTime string `pulumi:"creationTime"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The database watcher name.
	WatcherName string `pulumi:"watcherName"`
}

// The set of arguments for constructing a AlertRuleResource resource.
type AlertRuleResourceArgs struct {
	// The resource ID of the alert rule resource.
	AlertRuleResourceId pulumi.StringInput
	// The alert rule proxy resource name.
	AlertRuleResourceName pulumi.StringPtrInput
	// The template ID associated with alert rule resource.
	AlertRuleTemplateId pulumi.StringInput
	// The alert rule template version.
	AlertRuleTemplateVersion pulumi.StringInput
	// The properties with which the alert rule resource was created.
	CreatedWithProperties pulumi.StringInput
	// The creation time of the alert rule resource.
	CreationTime pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The database watcher name.
	WatcherName pulumi.StringInput
}

func (AlertRuleResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleResourceArgs)(nil)).Elem()
}

type AlertRuleResourceInput interface {
	pulumi.Input

	ToAlertRuleResourceOutput() AlertRuleResourceOutput
	ToAlertRuleResourceOutputWithContext(ctx context.Context) AlertRuleResourceOutput
}

func (*AlertRuleResource) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleResource)(nil)).Elem()
}

func (i *AlertRuleResource) ToAlertRuleResourceOutput() AlertRuleResourceOutput {
	return i.ToAlertRuleResourceOutputWithContext(context.Background())
}

func (i *AlertRuleResource) ToAlertRuleResourceOutputWithContext(ctx context.Context) AlertRuleResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleResourceOutput)
}

type AlertRuleResourceOutput struct{ *pulumi.OutputState }

func (AlertRuleResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleResource)(nil)).Elem()
}

func (o AlertRuleResourceOutput) ToAlertRuleResourceOutput() AlertRuleResourceOutput {
	return o
}

func (o AlertRuleResourceOutput) ToAlertRuleResourceOutputWithContext(ctx context.Context) AlertRuleResourceOutput {
	return o
}

// The resource ID of the alert rule resource.
func (o AlertRuleResourceOutput) AlertRuleResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.AlertRuleResourceId }).(pulumi.StringOutput)
}

// The template ID associated with alert rule resource.
func (o AlertRuleResourceOutput) AlertRuleTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.AlertRuleTemplateId }).(pulumi.StringOutput)
}

// The alert rule template version.
func (o AlertRuleResourceOutput) AlertRuleTemplateVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.AlertRuleTemplateVersion }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o AlertRuleResourceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The properties with which the alert rule resource was created.
func (o AlertRuleResourceOutput) CreatedWithProperties() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.CreatedWithProperties }).(pulumi.StringOutput)
}

// The creation time of the alert rule resource.
func (o AlertRuleResourceOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o AlertRuleResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the alert rule resource.
func (o AlertRuleResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AlertRuleResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AlertRuleResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AlertRuleResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRuleResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleResourceOutput{})
}
