// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes the configuration of a Business Application Agent.
//
// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-04-01-preview.
//
// Other available API versions: 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type BusinessApplicationAgent struct {
	pulumi.CustomResourceState

	AgentSystems AgentSystemResponseArrayOutput `pulumi:"agentSystems"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Describes the configuration of a Business Application Agent.
	Configuration SapAgentConfigurationResponseOutput `pulumi:"configuration"`
	DisplayName   pulumi.StringOutput                 `pulumi:"displayName"`
	// Etag of the azure resource
	Etag                pulumi.StringPtrOutput `pulumi:"etag"`
	LastModifiedTimeUtc pulumi.StringOutput    `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBusinessApplicationAgent registers a new resource with the given unique name, arguments, and options.
func NewBusinessApplicationAgent(ctx *pulumi.Context,
	name string, args *BusinessApplicationAgentArgs, opts ...pulumi.ResourceOption) (*BusinessApplicationAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:BusinessApplicationAgent"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:BusinessApplicationAgent"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:BusinessApplicationAgent"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250401preview:BusinessApplicationAgent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BusinessApplicationAgent
	err := ctx.RegisterResource("azure-native:securityinsights:BusinessApplicationAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBusinessApplicationAgent gets an existing BusinessApplicationAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBusinessApplicationAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BusinessApplicationAgentState, opts ...pulumi.ResourceOption) (*BusinessApplicationAgent, error) {
	var resource BusinessApplicationAgent
	err := ctx.ReadResource("azure-native:securityinsights:BusinessApplicationAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BusinessApplicationAgent resources.
type businessApplicationAgentState struct {
}

type BusinessApplicationAgentState struct {
}

func (BusinessApplicationAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*businessApplicationAgentState)(nil)).Elem()
}

type businessApplicationAgentArgs struct {
	// Business Application Agent Name
	AgentResourceName *string `pulumi:"agentResourceName"`
	// Describes the configuration of a Business Application Agent.
	Configuration SapAgentConfiguration `pulumi:"configuration"`
	DisplayName   string                `pulumi:"displayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a BusinessApplicationAgent resource.
type BusinessApplicationAgentArgs struct {
	// Business Application Agent Name
	AgentResourceName pulumi.StringPtrInput
	// Describes the configuration of a Business Application Agent.
	Configuration SapAgentConfigurationInput
	DisplayName   pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (BusinessApplicationAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*businessApplicationAgentArgs)(nil)).Elem()
}

type BusinessApplicationAgentInput interface {
	pulumi.Input

	ToBusinessApplicationAgentOutput() BusinessApplicationAgentOutput
	ToBusinessApplicationAgentOutputWithContext(ctx context.Context) BusinessApplicationAgentOutput
}

func (*BusinessApplicationAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessApplicationAgent)(nil)).Elem()
}

func (i *BusinessApplicationAgent) ToBusinessApplicationAgentOutput() BusinessApplicationAgentOutput {
	return i.ToBusinessApplicationAgentOutputWithContext(context.Background())
}

func (i *BusinessApplicationAgent) ToBusinessApplicationAgentOutputWithContext(ctx context.Context) BusinessApplicationAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessApplicationAgentOutput)
}

type BusinessApplicationAgentOutput struct{ *pulumi.OutputState }

func (BusinessApplicationAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessApplicationAgent)(nil)).Elem()
}

func (o BusinessApplicationAgentOutput) ToBusinessApplicationAgentOutput() BusinessApplicationAgentOutput {
	return o
}

func (o BusinessApplicationAgentOutput) ToBusinessApplicationAgentOutputWithContext(ctx context.Context) BusinessApplicationAgentOutput {
	return o
}

func (o BusinessApplicationAgentOutput) AgentSystems() AgentSystemResponseArrayOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) AgentSystemResponseArrayOutput { return v.AgentSystems }).(AgentSystemResponseArrayOutput)
}

// The Azure API version of the resource.
func (o BusinessApplicationAgentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Describes the configuration of a Business Application Agent.
func (o BusinessApplicationAgentOutput) Configuration() SapAgentConfigurationResponseOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) SapAgentConfigurationResponseOutput { return v.Configuration }).(SapAgentConfigurationResponseOutput)
}

func (o BusinessApplicationAgentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o BusinessApplicationAgentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BusinessApplicationAgentOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o BusinessApplicationAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BusinessApplicationAgentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BusinessApplicationAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessApplicationAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BusinessApplicationAgentOutput{})
}
