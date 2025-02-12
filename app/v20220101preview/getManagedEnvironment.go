// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the properties of a Managed Environment used to host container apps.
func LookupManagedEnvironment(ctx *pulumi.Context, args *LookupManagedEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupManagedEnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedEnvironmentResult
	err := ctx.Invoke("azure-native:app/v20220101preview:getManagedEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedEnvironmentArgs struct {
	// Name of the Environment.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An environment for hosting container apps
type LookupManagedEnvironmentResult struct {
	// Cluster configuration which enables the log daemon to export app logs to configured destination.
	AppLogsConfiguration *AppLogsConfigurationResponse `pulumi:"appLogsConfiguration"`
	// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
	DaprAIInstrumentationKey *string `pulumi:"daprAIInstrumentationKey"`
	// Default Domain Name for the cluster
	DefaultDomain string `pulumi:"defaultDomain"`
	// Any errors that occurred during deployment or deployment validation
	DeploymentErrors string `pulumi:"deploymentErrors"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the Environment.
	ProvisioningState string `pulumi:"provisioningState"`
	// Static IP of the Environment
	StaticIp string `pulumi:"staticIp"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Vnet configuration for the environment
	VnetConfiguration *VnetConfigurationResponse `pulumi:"vnetConfiguration"`
}

func LookupManagedEnvironmentOutput(ctx *pulumi.Context, args LookupManagedEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupManagedEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupManagedEnvironmentResultOutput, error) {
			args := v.(LookupManagedEnvironmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:app/v20220101preview:getManagedEnvironment", args, LookupManagedEnvironmentResultOutput{}, options).(LookupManagedEnvironmentResultOutput), nil
		}).(LookupManagedEnvironmentResultOutput)
}

type LookupManagedEnvironmentOutputArgs struct {
	// Name of the Environment.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentArgs)(nil)).Elem()
}

// An environment for hosting container apps
type LookupManagedEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupManagedEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentResult)(nil)).Elem()
}

func (o LookupManagedEnvironmentResultOutput) ToLookupManagedEnvironmentResultOutput() LookupManagedEnvironmentResultOutput {
	return o
}

func (o LookupManagedEnvironmentResultOutput) ToLookupManagedEnvironmentResultOutputWithContext(ctx context.Context) LookupManagedEnvironmentResultOutput {
	return o
}

// Cluster configuration which enables the log daemon to export app logs to configured destination.
func (o LookupManagedEnvironmentResultOutput) AppLogsConfiguration() AppLogsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *AppLogsConfigurationResponse { return v.AppLogsConfiguration }).(AppLogsConfigurationResponsePtrOutput)
}

// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
func (o LookupManagedEnvironmentResultOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

// Default Domain Name for the cluster
func (o LookupManagedEnvironmentResultOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Any errors that occurred during deployment or deployment validation
func (o LookupManagedEnvironmentResultOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.DeploymentErrors }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupManagedEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Environment.
func (o LookupManagedEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Static IP of the Environment
func (o LookupManagedEnvironmentResultOutput) StaticIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.StaticIp }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedEnvironmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupManagedEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Vnet configuration for the environment
func (o LookupManagedEnvironmentResultOutput) VnetConfiguration() VnetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *VnetConfigurationResponse { return v.VnetConfiguration }).(VnetConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedEnvironmentResultOutput{})
}
