// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a specific live pipeline by name. If a live pipeline with that name has been previously created, the call will return the JSON representation of that instance.
//
// Uses Azure REST API version 2021-11-01-preview.
func LookupLivePipeline(ctx *pulumi.Context, args *LookupLivePipelineArgs, opts ...pulumi.InvokeOption) (*LookupLivePipelineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLivePipelineResult
	err := ctx.Invoke("azure-native:videoanalyzer:getLivePipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLivePipelineArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// Live pipeline unique identifier.
	LivePipelineName string `pulumi:"livePipelineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Live pipeline represents a unique instance of a live topology, used for real-time ingestion, archiving and publishing of content for a unique RTSP camera.
type LookupLivePipelineResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
	BitrateKbps int `pulumi:"bitrateKbps"`
	// An optional description for the pipeline.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters []ParameterDefinitionResponse `pulumi:"parameters"`
	// Current state of the pipeline (read-only).
	State string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
	TopologyName string `pulumi:"topologyName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLivePipelineOutput(ctx *pulumi.Context, args LookupLivePipelineOutputArgs, opts ...pulumi.InvokeOption) LookupLivePipelineResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLivePipelineResultOutput, error) {
			args := v.(LookupLivePipelineArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:videoanalyzer:getLivePipeline", args, LookupLivePipelineResultOutput{}, options).(LookupLivePipelineResultOutput), nil
		}).(LookupLivePipelineResultOutput)
}

type LookupLivePipelineOutputArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Live pipeline unique identifier.
	LivePipelineName pulumi.StringInput `pulumi:"livePipelineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLivePipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLivePipelineArgs)(nil)).Elem()
}

// Live pipeline represents a unique instance of a live topology, used for real-time ingestion, archiving and publishing of content for a unique RTSP camera.
type LookupLivePipelineResultOutput struct{ *pulumi.OutputState }

func (LookupLivePipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLivePipelineResult)(nil)).Elem()
}

func (o LookupLivePipelineResultOutput) ToLookupLivePipelineResultOutput() LookupLivePipelineResultOutput {
	return o
}

func (o LookupLivePipelineResultOutput) ToLookupLivePipelineResultOutputWithContext(ctx context.Context) LookupLivePipelineResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupLivePipelineResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Maximum bitrate capacity in Kbps reserved for the live pipeline. The allowed range is from 500 to 3000 Kbps in increments of 100 Kbps. If the RTSP camera exceeds this capacity, then the service will disconnect temporarily from the camera. It will retry to re-establish connection (with exponential backoff), checking to see if the camera bitrate is now below the reserved capacity. Doing so will ensure that one 'noisy neighbor' does not affect other live pipelines in your account.
func (o LookupLivePipelineResultOutput) BitrateKbps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) int { return v.BitrateKbps }).(pulumi.IntOutput)
}

// An optional description for the pipeline.
func (o LookupLivePipelineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLivePipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLivePipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
func (o LookupLivePipelineResultOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) []ParameterDefinitionResponse { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

// Current state of the pipeline (read-only).
func (o LookupLivePipelineResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLivePipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The reference to an existing pipeline topology defined for real-time content processing. When activated, this live pipeline will process content according to the pipeline topology definition.
func (o LookupLivePipelineResultOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.TopologyName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLivePipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLivePipelineResultOutput{})
}
