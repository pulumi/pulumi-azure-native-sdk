// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagemover

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Job Definition resource.
//
// Uses Azure REST API version 2024-07-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
//
// Other available API versions: 2023-03-01, 2023-07-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagemover [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type JobDefinition struct {
	pulumi.CustomResourceState

	// Name of the Agent to assign for new Job Runs of this Job Definition.
	AgentName pulumi.StringPtrOutput `pulumi:"agentName"`
	// Fully qualified resource id of the Agent to assign for new Job Runs of this Job Definition.
	AgentResourceId pulumi.StringOutput `pulumi:"agentResourceId"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Strategy to use for copy.
	CopyMode pulumi.StringOutput `pulumi:"copyMode"`
	// A description for the Job Definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Job Run in a non-terminal state, if exists.
	LatestJobRunName pulumi.StringOutput `pulumi:"latestJobRunName"`
	// The fully qualified resource ID of the Job Run in a non-terminal state, if exists.
	LatestJobRunResourceId pulumi.StringOutput `pulumi:"latestJobRunResourceId"`
	// The current status of the Job Run in a non-terminal state, if exists.
	LatestJobRunStatus pulumi.StringOutput `pulumi:"latestJobRunStatus"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of this resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The name of the source Endpoint.
	SourceName pulumi.StringOutput `pulumi:"sourceName"`
	// Fully qualified resource ID of the source Endpoint.
	SourceResourceId pulumi.StringOutput `pulumi:"sourceResourceId"`
	// The subpath to use when reading from the source Endpoint.
	SourceSubpath pulumi.StringPtrOutput `pulumi:"sourceSubpath"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The name of the target Endpoint.
	TargetName pulumi.StringOutput `pulumi:"targetName"`
	// Fully qualified resource ID of the target Endpoint.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// The subpath to use when writing to the target Endpoint.
	TargetSubpath pulumi.StringPtrOutput `pulumi:"targetSubpath"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CopyMode == nil {
		return nil, errors.New("invalid value for required argument 'CopyMode'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceName == nil {
		return nil, errors.New("invalid value for required argument 'SourceName'")
	}
	if args.StorageMoverName == nil {
		return nil, errors.New("invalid value for required argument 'StorageMoverName'")
	}
	if args.TargetName == nil {
		return nil, errors.New("invalid value for required argument 'TargetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagemover/v20220701preview:JobDefinition"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20230301:JobDefinition"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20230701preview:JobDefinition"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20231001:JobDefinition"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20240701:JobDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobDefinition
	err := ctx.RegisterResource("azure-native:storagemover:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("azure-native:storagemover:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
}

type JobDefinitionState struct {
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	// Name of the Agent to assign for new Job Runs of this Job Definition.
	AgentName *string `pulumi:"agentName"`
	// Strategy to use for copy.
	CopyMode string `pulumi:"copyMode"`
	// A description for the Job Definition.
	Description *string `pulumi:"description"`
	// The name of the Job Definition resource.
	JobDefinitionName *string `pulumi:"jobDefinitionName"`
	// The name of the Project resource.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the source Endpoint.
	SourceName string `pulumi:"sourceName"`
	// The subpath to use when reading from the source Endpoint.
	SourceSubpath *string `pulumi:"sourceSubpath"`
	// The name of the Storage Mover resource.
	StorageMoverName string `pulumi:"storageMoverName"`
	// The name of the target Endpoint.
	TargetName string `pulumi:"targetName"`
	// The subpath to use when writing to the target Endpoint.
	TargetSubpath *string `pulumi:"targetSubpath"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// Name of the Agent to assign for new Job Runs of this Job Definition.
	AgentName pulumi.StringPtrInput
	// Strategy to use for copy.
	CopyMode pulumi.StringInput
	// A description for the Job Definition.
	Description pulumi.StringPtrInput
	// The name of the Job Definition resource.
	JobDefinitionName pulumi.StringPtrInput
	// The name of the Project resource.
	ProjectName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the source Endpoint.
	SourceName pulumi.StringInput
	// The subpath to use when reading from the source Endpoint.
	SourceSubpath pulumi.StringPtrInput
	// The name of the Storage Mover resource.
	StorageMoverName pulumi.StringInput
	// The name of the target Endpoint.
	TargetName pulumi.StringInput
	// The subpath to use when writing to the target Endpoint.
	TargetSubpath pulumi.StringPtrInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

// Name of the Agent to assign for new Job Runs of this Job Definition.
func (o JobDefinitionOutput) AgentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.AgentName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource id of the Agent to assign for new Job Runs of this Job Definition.
func (o JobDefinitionOutput) AgentResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.AgentResourceId }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o JobDefinitionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Strategy to use for copy.
func (o JobDefinitionOutput) CopyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.CopyMode }).(pulumi.StringOutput)
}

// A description for the Job Definition.
func (o JobDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the Job Run in a non-terminal state, if exists.
func (o JobDefinitionOutput) LatestJobRunName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.LatestJobRunName }).(pulumi.StringOutput)
}

// The fully qualified resource ID of the Job Run in a non-terminal state, if exists.
func (o JobDefinitionOutput) LatestJobRunResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.LatestJobRunResourceId }).(pulumi.StringOutput)
}

// The current status of the Job Run in a non-terminal state, if exists.
func (o JobDefinitionOutput) LatestJobRunStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.LatestJobRunStatus }).(pulumi.StringOutput)
}

// The name of the resource
func (o JobDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of this resource.
func (o JobDefinitionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The name of the source Endpoint.
func (o JobDefinitionOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.SourceName }).(pulumi.StringOutput)
}

// Fully qualified resource ID of the source Endpoint.
func (o JobDefinitionOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.SourceResourceId }).(pulumi.StringOutput)
}

// The subpath to use when reading from the source Endpoint.
func (o JobDefinitionOutput) SourceSubpath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.SourceSubpath }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o JobDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *JobDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The name of the target Endpoint.
func (o JobDefinitionOutput) TargetName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.TargetName }).(pulumi.StringOutput)
}

// Fully qualified resource ID of the target Endpoint.
func (o JobDefinitionOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.TargetResourceId }).(pulumi.StringOutput)
}

// The subpath to use when writing to the target Endpoint.
func (o JobDefinitionOutput) TargetSubpath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.TargetSubpath }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o JobDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionOutput{})
}
