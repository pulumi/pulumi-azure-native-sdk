// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The metric setting details for the role
//
// Uses Azure REST API version 2023-07-01.
//
// Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupMonitoringConfig(ctx *pulumi.Context, args *LookupMonitoringConfigArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringConfigResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMonitoringConfigResult
	err := ctx.Invoke("azure-native:databoxedge:getMonitoringConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringConfigArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The role name.
	RoleName string `pulumi:"roleName"`
}

// The metric setting details for the role
type LookupMonitoringConfigResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The metrics configuration details
	MetricConfigurations []MetricConfigurationResponse `pulumi:"metricConfigurations"`
	// The object name.
	Name string `pulumi:"name"`
	// Metadata pertaining to creation and last modification of MonitoringConfiguration
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupMonitoringConfigOutput(ctx *pulumi.Context, args LookupMonitoringConfigOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMonitoringConfigResultOutput, error) {
			args := v.(LookupMonitoringConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:databoxedge:getMonitoringConfig", args, LookupMonitoringConfigResultOutput{}, options).(LookupMonitoringConfigResultOutput), nil
		}).(LookupMonitoringConfigResultOutput)
}

type LookupMonitoringConfigOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The role name.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (LookupMonitoringConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringConfigArgs)(nil)).Elem()
}

// The metric setting details for the role
type LookupMonitoringConfigResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringConfigResult)(nil)).Elem()
}

func (o LookupMonitoringConfigResultOutput) ToLookupMonitoringConfigResultOutput() LookupMonitoringConfigResultOutput {
	return o
}

func (o LookupMonitoringConfigResultOutput) ToLookupMonitoringConfigResultOutputWithContext(ctx context.Context) LookupMonitoringConfigResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMonitoringConfigResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupMonitoringConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The metrics configuration details
func (o LookupMonitoringConfigResultOutput) MetricConfigurations() MetricConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) []MetricConfigurationResponse { return v.MetricConfigurations }).(MetricConfigurationResponseArrayOutput)
}

// The object name.
func (o LookupMonitoringConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of MonitoringConfiguration
func (o LookupMonitoringConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupMonitoringConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringConfigResultOutput{})
}
