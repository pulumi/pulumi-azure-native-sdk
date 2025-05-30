// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredatatransfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets pipeline resource.
//
// Uses Azure REST API version 2024-09-27.
//
// Other available API versions: 2023-10-11-preview, 2024-01-25, 2024-05-07, 2024-09-11, 2025-03-01-preview, 2025-04-11-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuredatatransfer [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPipelineResult
	err := ctx.Invoke("azure-native:azuredatatransfer:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	// The name for the pipeline that is to be requested.
	PipelineName string `pulumi:"pipelineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The pipeline resource definition.
type LookupPipelineResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of pipeline
	Properties PipelinePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPipelineResultOutput, error) {
			args := v.(LookupPipelineArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azuredatatransfer:getPipeline", args, LookupPipelineResultOutput{}, options).(LookupPipelineResultOutput), nil
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	// The name for the pipeline that is to be requested.
	PipelineName pulumi.StringInput `pulumi:"pipelineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}

// The pipeline resource definition.
type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupPipelineResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPipelineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of pipeline
func (o LookupPipelineResultOutput) Properties() PipelinePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) PipelinePropertiesResponse { return v.Properties }).(PipelinePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPipelineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPipelineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
