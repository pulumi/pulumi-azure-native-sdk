// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The workflow type.
// API Version: 2019-05-01.
type Workflow struct {
	pulumi.CustomResourceState

	// The access control configuration.
	AccessControl FlowAccessControlConfigurationResponsePtrOutput `pulumi:"accessControl"`
	// Gets the access endpoint.
	AccessEndpoint pulumi.StringOutput `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The definition.
	Definition pulumi.AnyOutput `pulumi:"definition"`
	// The endpoints configuration.
	EndpointsConfiguration FlowEndpointsConfigurationResponsePtrOutput `pulumi:"endpointsConfiguration"`
	// Managed service identity properties.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The integration account.
	IntegrationAccount ResourceReferenceResponsePtrOutput `pulumi:"integrationAccount"`
	// The integration service environment.
	IntegrationServiceEnvironment ResourceReferenceResponsePtrOutput `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters.
	Parameters WorkflowParameterResponseMapOutput `pulumi:"parameters"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:logic:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:logic:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
}

type WorkflowState struct {
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// The access control configuration.
	AccessControl *FlowAccessControlConfiguration `pulumi:"accessControl"`
	// The definition.
	Definition interface{} `pulumi:"definition"`
	// The endpoints configuration.
	EndpointsConfiguration *FlowEndpointsConfiguration `pulumi:"endpointsConfiguration"`
	// Managed service identity properties.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The integration account.
	IntegrationAccount *ResourceReference `pulumi:"integrationAccount"`
	// The integration service environment.
	IntegrationServiceEnvironment *ResourceReference `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The parameters.
	Parameters map[string]WorkflowParameter `pulumi:"parameters"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The workflow name.
	WorkflowName *string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// The access control configuration.
	AccessControl FlowAccessControlConfigurationPtrInput
	// The definition.
	Definition pulumi.Input
	// The endpoints configuration.
	EndpointsConfiguration FlowEndpointsConfigurationPtrInput
	// Managed service identity properties.
	Identity ManagedServiceIdentityPtrInput
	// The integration account.
	IntegrationAccount ResourceReferencePtrInput
	// The integration service environment.
	IntegrationServiceEnvironment ResourceReferencePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The parameters.
	Parameters WorkflowParameterMapInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The state.
	State pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The workflow name.
	WorkflowName pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

// The access control configuration.
func (o WorkflowOutput) AccessControl() FlowAccessControlConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) FlowAccessControlConfigurationResponsePtrOutput { return v.AccessControl }).(FlowAccessControlConfigurationResponsePtrOutput)
}

// Gets the access endpoint.
func (o WorkflowOutput) AccessEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.AccessEndpoint }).(pulumi.StringOutput)
}

// Gets the changed time.
func (o WorkflowOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// Gets the created time.
func (o WorkflowOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The definition.
func (o WorkflowOutput) Definition() pulumi.AnyOutput {
	return o.ApplyT(func(v *Workflow) pulumi.AnyOutput { return v.Definition }).(pulumi.AnyOutput)
}

// The endpoints configuration.
func (o WorkflowOutput) EndpointsConfiguration() FlowEndpointsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) FlowEndpointsConfigurationResponsePtrOutput { return v.EndpointsConfiguration }).(FlowEndpointsConfigurationResponsePtrOutput)
}

// Managed service identity properties.
func (o WorkflowOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The integration account.
func (o WorkflowOutput) IntegrationAccount() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ResourceReferenceResponsePtrOutput { return v.IntegrationAccount }).(ResourceReferenceResponsePtrOutput)
}

// The integration service environment.
func (o WorkflowOutput) IntegrationServiceEnvironment() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ResourceReferenceResponsePtrOutput { return v.IntegrationServiceEnvironment }).(ResourceReferenceResponsePtrOutput)
}

// The resource location.
func (o WorkflowOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o WorkflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameters.
func (o WorkflowOutput) Parameters() WorkflowParameterResponseMapOutput {
	return o.ApplyT(func(v *Workflow) WorkflowParameterResponseMapOutput { return v.Parameters }).(WorkflowParameterResponseMapOutput)
}

// Gets the provisioning state.
func (o WorkflowOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku.
func (o WorkflowOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Workflow) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// The state.
func (o WorkflowOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o WorkflowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o WorkflowOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets the version.
func (o WorkflowOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}