// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment stack object.
//
// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2022-08-01-preview.
//
// Other available API versions: 2022-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type DeploymentStackAtManagementGroup struct {
	pulumi.CustomResourceState

	// Defines the behavior of resources that are no longer managed after the Deployment stack is updated or deleted.
	ActionOnUnmanage ActionOnUnmanageResponseOutput `pulumi:"actionOnUnmanage"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The correlation id of the last Deployment stack upsert or delete operation. It is in GUID format and is used for tracing.
	CorrelationId pulumi.StringOutput `pulumi:"correlationId"`
	// The debug setting of the deployment.
	DebugSetting DeploymentStacksDebugSettingResponsePtrOutput `pulumi:"debugSetting"`
	// An array of resources that were deleted during the most recent Deployment stack update. Deleted means that the resource was removed from the template and relevant deletion operations were specified.
	DeletedResources ResourceReferenceResponseArrayOutput `pulumi:"deletedResources"`
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettingsResponseOutput `pulumi:"denySettings"`
	// The resourceId of the deployment resource created by the deployment stack.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope pulumi.StringPtrOutput `pulumi:"deploymentScope"`
	// Deployment stack description. Max length of 4096 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An array of resources that were detached during the most recent Deployment stack update. Detached means that the resource was removed from the template, but no relevant deletion operations were specified. So, the resource still exists while no longer being associated with the stack.
	DetachedResources ResourceReferenceResponseArrayOutput `pulumi:"detachedResources"`
	// The duration of the last successful Deployment stack update.
	Duration pulumi.StringOutput `pulumi:"duration"`
	// The error detail.
	Error ErrorDetailResponsePtrOutput `pulumi:"error"`
	// An array of resources that failed to reach goal state during the most recent update. Each resourceId is accompanied by an error message.
	FailedResources ResourceReferenceExtendedResponseArrayOutput `pulumi:"failedResources"`
	// The location of the Deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The outputs of the deployment resource created by the deployment stack.
	Outputs pulumi.AnyOutput `pulumi:"outputs"`
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both.
	Parameters DeploymentParameterResponseMapOutput `pulumi:"parameters"`
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink DeploymentStacksParametersLinkResponsePtrOutput `pulumi:"parametersLink"`
	// State of the deployment stack.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// An array of resources currently managed by the deployment stack.
	Resources ManagedResourceReferenceResponseArrayOutput `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Deployment stack resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeploymentStackAtManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewDeploymentStackAtManagementGroup(ctx *pulumi.Context,
	name string, args *DeploymentStackAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*DeploymentStackAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionOnUnmanage == nil {
		return nil, errors.New("invalid value for required argument 'ActionOnUnmanage'")
	}
	if args.DenySettings == nil {
		return nil, errors.New("invalid value for required argument 'DenySettings'")
	}
	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20220801preview:DeploymentStackAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20240301:DeploymentStackAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DeploymentStackAtManagementGroup
	err := ctx.RegisterResource("azure-native:resources:DeploymentStackAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentStackAtManagementGroup gets an existing DeploymentStackAtManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentStackAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentStackAtManagementGroupState, opts ...pulumi.ResourceOption) (*DeploymentStackAtManagementGroup, error) {
	var resource DeploymentStackAtManagementGroup
	err := ctx.ReadResource("azure-native:resources:DeploymentStackAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentStackAtManagementGroup resources.
type deploymentStackAtManagementGroupState struct {
}

type DeploymentStackAtManagementGroupState struct {
}

func (DeploymentStackAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentStackAtManagementGroupState)(nil)).Elem()
}

type deploymentStackAtManagementGroupArgs struct {
	// Defines the behavior of resources that are no longer managed after the Deployment stack is updated or deleted.
	ActionOnUnmanage ActionOnUnmanage `pulumi:"actionOnUnmanage"`
	// Flag to bypass service errors that indicate the stack resource list is not correctly synchronized.
	BypassStackOutOfSyncError *bool `pulumi:"bypassStackOutOfSyncError"`
	// The debug setting of the deployment.
	DebugSetting *DeploymentStacksDebugSetting `pulumi:"debugSetting"`
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettings `pulumi:"denySettings"`
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope *string `pulumi:"deploymentScope"`
	// Name of the deployment stack.
	DeploymentStackName *string `pulumi:"deploymentStackName"`
	// Deployment stack description. Max length of 4096 characters.
	Description *string `pulumi:"description"`
	// The location of the Deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location *string `pulumi:"location"`
	// Management Group id.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both.
	Parameters map[string]DeploymentParameter `pulumi:"parameters"`
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink *DeploymentStacksParametersLink `pulumi:"parametersLink"`
	// Deployment stack resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template interface{} `pulumi:"template"`
	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *DeploymentStacksTemplateLink `pulumi:"templateLink"`
}

// The set of arguments for constructing a DeploymentStackAtManagementGroup resource.
type DeploymentStackAtManagementGroupArgs struct {
	// Defines the behavior of resources that are no longer managed after the Deployment stack is updated or deleted.
	ActionOnUnmanage ActionOnUnmanageInput
	// Flag to bypass service errors that indicate the stack resource list is not correctly synchronized.
	BypassStackOutOfSyncError pulumi.BoolPtrInput
	// The debug setting of the deployment.
	DebugSetting DeploymentStacksDebugSettingPtrInput
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettingsInput
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope pulumi.StringPtrInput
	// Name of the deployment stack.
	DeploymentStackName pulumi.StringPtrInput
	// Deployment stack description. Max length of 4096 characters.
	Description pulumi.StringPtrInput
	// The location of the Deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location pulumi.StringPtrInput
	// Management Group id.
	ManagementGroupId pulumi.StringInput
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both.
	Parameters DeploymentParameterMapInput
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink DeploymentStacksParametersLinkPtrInput
	// Deployment stack resource tags.
	Tags pulumi.StringMapInput
	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template pulumi.Input
	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink DeploymentStacksTemplateLinkPtrInput
}

func (DeploymentStackAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentStackAtManagementGroupArgs)(nil)).Elem()
}

type DeploymentStackAtManagementGroupInput interface {
	pulumi.Input

	ToDeploymentStackAtManagementGroupOutput() DeploymentStackAtManagementGroupOutput
	ToDeploymentStackAtManagementGroupOutputWithContext(ctx context.Context) DeploymentStackAtManagementGroupOutput
}

func (*DeploymentStackAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentStackAtManagementGroup)(nil)).Elem()
}

func (i *DeploymentStackAtManagementGroup) ToDeploymentStackAtManagementGroupOutput() DeploymentStackAtManagementGroupOutput {
	return i.ToDeploymentStackAtManagementGroupOutputWithContext(context.Background())
}

func (i *DeploymentStackAtManagementGroup) ToDeploymentStackAtManagementGroupOutputWithContext(ctx context.Context) DeploymentStackAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentStackAtManagementGroupOutput)
}

type DeploymentStackAtManagementGroupOutput struct{ *pulumi.OutputState }

func (DeploymentStackAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentStackAtManagementGroup)(nil)).Elem()
}

func (o DeploymentStackAtManagementGroupOutput) ToDeploymentStackAtManagementGroupOutput() DeploymentStackAtManagementGroupOutput {
	return o
}

func (o DeploymentStackAtManagementGroupOutput) ToDeploymentStackAtManagementGroupOutputWithContext(ctx context.Context) DeploymentStackAtManagementGroupOutput {
	return o
}

// Defines the behavior of resources that are no longer managed after the Deployment stack is updated or deleted.
func (o DeploymentStackAtManagementGroupOutput) ActionOnUnmanage() ActionOnUnmanageResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ActionOnUnmanageResponseOutput { return v.ActionOnUnmanage }).(ActionOnUnmanageResponseOutput)
}

// The Azure API version of the resource.
func (o DeploymentStackAtManagementGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The correlation id of the last Deployment stack upsert or delete operation. It is in GUID format and is used for tracing.
func (o DeploymentStackAtManagementGroupOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.CorrelationId }).(pulumi.StringOutput)
}

// The debug setting of the deployment.
func (o DeploymentStackAtManagementGroupOutput) DebugSetting() DeploymentStacksDebugSettingResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DeploymentStacksDebugSettingResponsePtrOutput {
		return v.DebugSetting
	}).(DeploymentStacksDebugSettingResponsePtrOutput)
}

// An array of resources that were deleted during the most recent Deployment stack update. Deleted means that the resource was removed from the template and relevant deletion operations were specified.
func (o DeploymentStackAtManagementGroupOutput) DeletedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ResourceReferenceResponseArrayOutput {
		return v.DeletedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// Defines how resources deployed by the stack are locked.
func (o DeploymentStackAtManagementGroupOutput) DenySettings() DenySettingsResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DenySettingsResponseOutput { return v.DenySettings }).(DenySettingsResponseOutput)
}

// The resourceId of the deployment resource created by the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
func (o DeploymentStackAtManagementGroupOutput) DeploymentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringPtrOutput { return v.DeploymentScope }).(pulumi.StringPtrOutput)
}

// Deployment stack description. Max length of 4096 characters.
func (o DeploymentStackAtManagementGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of resources that were detached during the most recent Deployment stack update. Detached means that the resource was removed from the template, but no relevant deletion operations were specified. So, the resource still exists while no longer being associated with the stack.
func (o DeploymentStackAtManagementGroupOutput) DetachedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ResourceReferenceResponseArrayOutput {
		return v.DetachedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// The duration of the last successful Deployment stack update.
func (o DeploymentStackAtManagementGroupOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

// The error detail.
func (o DeploymentStackAtManagementGroupOutput) Error() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ErrorDetailResponsePtrOutput { return v.Error }).(ErrorDetailResponsePtrOutput)
}

// An array of resources that failed to reach goal state during the most recent update. Each resourceId is accompanied by an error message.
func (o DeploymentStackAtManagementGroupOutput) FailedResources() ResourceReferenceExtendedResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ResourceReferenceExtendedResponseArrayOutput {
		return v.FailedResources
	}).(ResourceReferenceExtendedResponseArrayOutput)
}

// The location of the Deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
func (o DeploymentStackAtManagementGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of this resource.
func (o DeploymentStackAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The outputs of the deployment resource created by the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.AnyOutput { return v.Outputs }).(pulumi.AnyOutput)
}

// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both.
func (o DeploymentStackAtManagementGroupOutput) Parameters() DeploymentParameterResponseMapOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DeploymentParameterResponseMapOutput { return v.Parameters }).(DeploymentParameterResponseMapOutput)
}

// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
func (o DeploymentStackAtManagementGroupOutput) ParametersLink() DeploymentStacksParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) DeploymentStacksParametersLinkResponsePtrOutput {
		return v.ParametersLink
	}).(DeploymentStacksParametersLinkResponsePtrOutput)
}

// State of the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of resources currently managed by the deployment stack.
func (o DeploymentStackAtManagementGroupOutput) Resources() ManagedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) ManagedResourceReferenceResponseArrayOutput {
		return v.Resources
	}).(ManagedResourceReferenceResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DeploymentStackAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Deployment stack resource tags.
func (o DeploymentStackAtManagementGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o DeploymentStackAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentStackAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentStackAtManagementGroupOutput{})
}
