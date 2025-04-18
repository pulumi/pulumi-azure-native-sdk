// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Live pipeline represents a unique instance of a live topology, used for real-time ingestion, archiving and publishing of content for a unique RTSP camera.
//
// Uses Azure REST API version 2021-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-11-01-preview.
type LivePipeline struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
	BitrateKbps pulumi.IntOutput `pulumi:"bitrateKbps"`
	// An optional description for the pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters ParameterDefinitionResponseArrayOutput `pulumi:"parameters"`
	// Current state of the pipeline (read-only).
	State pulumi.StringOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
	TopologyName pulumi.StringOutput `pulumi:"topologyName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLivePipeline registers a new resource with the given unique name, arguments, and options.
func NewLivePipeline(ctx *pulumi.Context,
	name string, args *LivePipelineArgs, opts ...pulumi.ResourceOption) (*LivePipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.BitrateKbps == nil {
		return nil, errors.New("invalid value for required argument 'BitrateKbps'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopologyName == nil {
		return nil, errors.New("invalid value for required argument 'TopologyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:LivePipeline"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LivePipeline
	err := ctx.RegisterResource("azure-native:videoanalyzer:LivePipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLivePipeline gets an existing LivePipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLivePipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LivePipelineState, opts ...pulumi.ResourceOption) (*LivePipeline, error) {
	var resource LivePipeline
	err := ctx.ReadResource("azure-native:videoanalyzer:LivePipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LivePipeline resources.
type livePipelineState struct {
}

type LivePipelineState struct {
}

func (LivePipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*livePipelineState)(nil)).Elem()
}

type livePipelineArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
	BitrateKbps int `pulumi:"bitrateKbps"`
	// An optional description for the pipeline.
	Description *string `pulumi:"description"`
	// Live pipeline unique identifier.
	LivePipelineName *string `pulumi:"livePipelineName"`
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters []ParameterDefinition `pulumi:"parameters"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
	TopologyName string `pulumi:"topologyName"`
}

// The set of arguments for constructing a LivePipeline resource.
type LivePipelineArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput
	// Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
	BitrateKbps pulumi.IntInput
	// An optional description for the pipeline.
	Description pulumi.StringPtrInput
	// Live pipeline unique identifier.
	LivePipelineName pulumi.StringPtrInput
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters ParameterDefinitionArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
	TopologyName pulumi.StringInput
}

func (LivePipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*livePipelineArgs)(nil)).Elem()
}

type LivePipelineInput interface {
	pulumi.Input

	ToLivePipelineOutput() LivePipelineOutput
	ToLivePipelineOutputWithContext(ctx context.Context) LivePipelineOutput
}

func (*LivePipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**LivePipeline)(nil)).Elem()
}

func (i *LivePipeline) ToLivePipelineOutput() LivePipelineOutput {
	return i.ToLivePipelineOutputWithContext(context.Background())
}

func (i *LivePipeline) ToLivePipelineOutputWithContext(ctx context.Context) LivePipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LivePipelineOutput)
}

type LivePipelineOutput struct{ *pulumi.OutputState }

func (LivePipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LivePipeline)(nil)).Elem()
}

func (o LivePipelineOutput) ToLivePipelineOutput() LivePipelineOutput {
	return o
}

func (o LivePipelineOutput) ToLivePipelineOutputWithContext(ctx context.Context) LivePipelineOutput {
	return o
}

// The Azure API version of the resource.
func (o LivePipelineOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
func (o LivePipelineOutput) BitrateKbps() pulumi.IntOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.IntOutput { return v.BitrateKbps }).(pulumi.IntOutput)
}

// An optional description for the pipeline.
func (o LivePipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LivePipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
func (o LivePipelineOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *LivePipeline) ParameterDefinitionResponseArrayOutput { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

// Current state of the pipeline (read-only).
func (o LivePipelineOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LivePipelineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LivePipeline) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
func (o LivePipelineOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.TopologyName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LivePipelineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LivePipelineOutput{})
}
