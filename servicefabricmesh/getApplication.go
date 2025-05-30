// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the information about the application resource with the given name. The information include the description and other properties of the application.
//
// Uses Azure REST API version 2018-09-01-preview.
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	// The identity of the application.
	ApplicationResourceName string `pulumi:"applicationResourceName"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// This type describes an application resource.
type LookupApplicationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Internal - used by Visual Studio to setup the debugging session on the local development environment.
	DebugParams *string `pulumi:"debugParams"`
	// User readable description of the application.
	Description *string `pulumi:"description"`
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics *DiagnosticsDescriptionResponse `pulumi:"diagnostics"`
	// Describes the health state of an application resource.
	HealthState string `pulumi:"healthState"`
	// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Names of the services in the application.
	ServiceNames []string `pulumi:"serviceNames"`
	// Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
	Services []ServiceResourceDescriptionResponse `pulumi:"services"`
	// Status of the application.
	Status string `pulumi:"status"`
	// Gives additional information about the current status of the application.
	StatusDetails string `pulumi:"statusDetails"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
	UnhealthyEvaluation string `pulumi:"unhealthyEvaluation"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApplicationResultOutput, error) {
			args := v.(LookupApplicationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:servicefabricmesh:getApplication", args, LookupApplicationResultOutput{}, options).(LookupApplicationResultOutput), nil
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	// The identity of the application.
	ApplicationResourceName pulumi.StringInput `pulumi:"applicationResourceName"`
	// Azure resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

// This type describes an application resource.
type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupApplicationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Internal - used by Visual Studio to setup the debugging session on the local development environment.
func (o LookupApplicationResultOutput) DebugParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.DebugParams }).(pulumi.StringPtrOutput)
}

// User readable description of the application.
func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Describes the diagnostics definition and usage for an application resource.
func (o LookupApplicationResultOutput) Diagnostics() DiagnosticsDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *DiagnosticsDescriptionResponse { return v.Diagnostics }).(DiagnosticsDescriptionResponsePtrOutput)
}

// Describes the health state of an application resource.
func (o LookupApplicationResultOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.HealthState }).(pulumi.StringOutput)
}

// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupApplicationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the resource.
func (o LookupApplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Names of the services in the application.
func (o LookupApplicationResultOutput) ServiceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.ServiceNames }).(pulumi.StringArrayOutput)
}

// Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
func (o LookupApplicationResultOutput) Services() ServiceResourceDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ServiceResourceDescriptionResponse { return v.Services }).(ServiceResourceDescriptionResponseArrayOutput)
}

// Status of the application.
func (o LookupApplicationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Status }).(pulumi.StringOutput)
}

// Gives additional information about the current status of the application.
func (o LookupApplicationResultOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.StatusDetails }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
func (o LookupApplicationResultOutput) UnhealthyEvaluation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.UnhealthyEvaluation }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
