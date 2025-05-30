// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Diagnostic specified by its identifier.
//
// Uses Azure REST API version 2024-06-01-preview.
//
// Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupWorkspaceDiagnostic(ctx *pulumi.Context, args *LookupWorkspaceDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceDiagnosticResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceDiagnosticArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Diagnostic details.
type LookupWorkspaceDiagnosticResult struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Large Language Models diagnostic settings
	LargeLanguageModel *LLMDiagnosticSettingsResponse `pulumi:"largeLanguageModel"`
	// Log the ClientIP. Default is false.
	LogClientIp *bool `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId string `pulumi:"loggerId"`
	// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
	Metrics *bool `pulumi:"metrics"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The format of the Operation Name for Application Insights telemetries. Default is Name.
	OperationNameFormat *string `pulumi:"operationNameFormat"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettingsResponse `pulumi:"sampling"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity *string `pulumi:"verbosity"`
}

func LookupWorkspaceDiagnosticOutput(ctx *pulumi.Context, args LookupWorkspaceDiagnosticOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceDiagnosticResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceDiagnosticResultOutput, error) {
			args := v.(LookupWorkspaceDiagnosticArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getWorkspaceDiagnostic", args, LookupWorkspaceDiagnosticResultOutput{}, options).(LookupWorkspaceDiagnosticResultOutput), nil
		}).(LookupWorkspaceDiagnosticResultOutput)
}

type LookupWorkspaceDiagnosticOutputArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput `pulumi:"diagnosticId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceDiagnosticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceDiagnosticArgs)(nil)).Elem()
}

// Diagnostic details.
type LookupWorkspaceDiagnosticResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceDiagnosticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceDiagnosticResult)(nil)).Elem()
}

func (o LookupWorkspaceDiagnosticResultOutput) ToLookupWorkspaceDiagnosticResultOutput() LookupWorkspaceDiagnosticResultOutput {
	return o
}

func (o LookupWorkspaceDiagnosticResultOutput) ToLookupWorkspaceDiagnosticResultOutputWithContext(ctx context.Context) LookupWorkspaceDiagnosticResultOutput {
	return o
}

// Specifies for what type of messages sampling settings should not apply.
func (o LookupWorkspaceDiagnosticResultOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *string { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupWorkspaceDiagnosticResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
func (o LookupWorkspaceDiagnosticResultOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
func (o LookupWorkspaceDiagnosticResultOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Sets correlation protocol to use for Application Insights diagnostics.
func (o LookupWorkspaceDiagnosticResultOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *string { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

// Large Language Models diagnostic settings
func (o LookupWorkspaceDiagnosticResultOutput) LargeLanguageModel() LLMDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *LLMDiagnosticSettingsResponse { return v.LargeLanguageModel }).(LLMDiagnosticSettingsResponsePtrOutput)
}

// Log the ClientIP. Default is false.
func (o LookupWorkspaceDiagnosticResultOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *bool { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

// Resource Id of a target logger.
func (o LookupWorkspaceDiagnosticResultOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) string { return v.LoggerId }).(pulumi.StringOutput)
}

// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
func (o LookupWorkspaceDiagnosticResultOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *bool { return v.Metrics }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o LookupWorkspaceDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

// The format of the Operation Name for Application Insights telemetries. Default is Name.
func (o LookupWorkspaceDiagnosticResultOutput) OperationNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *string { return v.OperationNameFormat }).(pulumi.StringPtrOutput)
}

// Sampling settings for Diagnostic.
func (o LookupWorkspaceDiagnosticResultOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *SamplingSettingsResponse { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

// The verbosity level applied to traces emitted by trace policies.
func (o LookupWorkspaceDiagnosticResultOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceDiagnosticResult) *string { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceDiagnosticResultOutput{})
}
