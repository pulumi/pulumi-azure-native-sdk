// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the properties of a Managed Environment used to host container apps.
//
// Uses Azure REST API version 2024-03-01.
//
// Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupManagedEnvironment(ctx *pulumi.Context, args *LookupManagedEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupManagedEnvironmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedEnvironmentResult
	err := ctx.Invoke("azure-native:app:getManagedEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedEnvironmentArgs struct {
	// Name of the Environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An environment for hosting container apps
type LookupManagedEnvironmentResult struct {
	// Cluster configuration which enables the log daemon to export app logs to configured destination.
	AppLogsConfiguration *AppLogsConfigurationResponse `pulumi:"appLogsConfiguration"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Custom domain configuration for the environment
	CustomDomainConfiguration *CustomDomainConfigurationResponse `pulumi:"customDomainConfiguration"`
	// Application Insights connection string used by Dapr to export Service to Service communication telemetry
	DaprAIConnectionString *string `pulumi:"daprAIConnectionString"`
	// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
	DaprAIInstrumentationKey *string `pulumi:"daprAIInstrumentationKey"`
	// The configuration of Dapr component.
	DaprConfiguration *DaprConfigurationResponse `pulumi:"daprConfiguration"`
	// Default Domain Name for the cluster
	DefaultDomain string `pulumi:"defaultDomain"`
	// Any errors that occurred during deployment or deployment validation
	DeploymentErrors string `pulumi:"deploymentErrors"`
	// The endpoint of the eventstream of the Environment.
	EventStreamEndpoint string `pulumi:"eventStreamEndpoint"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Name of the platform-managed resource group created for the Managed Environment to host infrastructure resources. If a subnet ID is provided, this resource group will be created in the same subscription as the subnet.
	InfrastructureResourceGroup *string `pulumi:"infrastructureResourceGroup"`
	// The configuration of Keda component.
	KedaConfiguration *KedaConfigurationResponse `pulumi:"kedaConfiguration"`
	// Kind of the Environment.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Peer authentication settings for the Managed Environment
	PeerAuthentication *ManagedEnvironmentResponsePeerAuthentication `pulumi:"peerAuthentication"`
	// Peer traffic settings for the Managed Environment
	PeerTrafficConfiguration *ManagedEnvironmentResponsePeerTrafficConfiguration `pulumi:"peerTrafficConfiguration"`
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
	// Workload profiles configured for the Managed Environment.
	WorkloadProfiles []WorkloadProfileResponse `pulumi:"workloadProfiles"`
	// Whether or not this Managed Environment is zone-redundant.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

func LookupManagedEnvironmentOutput(ctx *pulumi.Context, args LookupManagedEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupManagedEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupManagedEnvironmentResultOutput, error) {
			args := v.(LookupManagedEnvironmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:app:getManagedEnvironment", args, LookupManagedEnvironmentResultOutput{}, options).(LookupManagedEnvironmentResultOutput), nil
		}).(LookupManagedEnvironmentResultOutput)
}

type LookupManagedEnvironmentOutputArgs struct {
	// Name of the Environment.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
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

// The Azure API version of the resource.
func (o LookupManagedEnvironmentResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Custom domain configuration for the environment
func (o LookupManagedEnvironmentResultOutput) CustomDomainConfiguration() CustomDomainConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *CustomDomainConfigurationResponse {
		return v.CustomDomainConfiguration
	}).(CustomDomainConfigurationResponsePtrOutput)
}

// Application Insights connection string used by Dapr to export Service to Service communication telemetry
func (o LookupManagedEnvironmentResultOutput) DaprAIConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.DaprAIConnectionString }).(pulumi.StringPtrOutput)
}

// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
func (o LookupManagedEnvironmentResultOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

// The configuration of Dapr component.
func (o LookupManagedEnvironmentResultOutput) DaprConfiguration() DaprConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *DaprConfigurationResponse { return v.DaprConfiguration }).(DaprConfigurationResponsePtrOutput)
}

// Default Domain Name for the cluster
func (o LookupManagedEnvironmentResultOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Any errors that occurred during deployment or deployment validation
func (o LookupManagedEnvironmentResultOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.DeploymentErrors }).(pulumi.StringOutput)
}

// The endpoint of the eventstream of the Environment.
func (o LookupManagedEnvironmentResultOutput) EventStreamEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.EventStreamEndpoint }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the platform-managed resource group created for the Managed Environment to host infrastructure resources. If a subnet ID is provided, this resource group will be created in the same subscription as the subnet.
func (o LookupManagedEnvironmentResultOutput) InfrastructureResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.InfrastructureResourceGroup }).(pulumi.StringPtrOutput)
}

// The configuration of Keda component.
func (o LookupManagedEnvironmentResultOutput) KedaConfiguration() KedaConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *KedaConfigurationResponse { return v.KedaConfiguration }).(KedaConfigurationResponsePtrOutput)
}

// Kind of the Environment.
func (o LookupManagedEnvironmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupManagedEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Peer authentication settings for the Managed Environment
func (o LookupManagedEnvironmentResultOutput) PeerAuthentication() ManagedEnvironmentResponsePeerAuthenticationPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *ManagedEnvironmentResponsePeerAuthentication {
		return v.PeerAuthentication
	}).(ManagedEnvironmentResponsePeerAuthenticationPtrOutput)
}

// Peer traffic settings for the Managed Environment
func (o LookupManagedEnvironmentResultOutput) PeerTrafficConfiguration() ManagedEnvironmentResponsePeerTrafficConfigurationPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *ManagedEnvironmentResponsePeerTrafficConfiguration {
		return v.PeerTrafficConfiguration
	}).(ManagedEnvironmentResponsePeerTrafficConfigurationPtrOutput)
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

// Workload profiles configured for the Managed Environment.
func (o LookupManagedEnvironmentResultOutput) WorkloadProfiles() WorkloadProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) []WorkloadProfileResponse { return v.WorkloadProfiles }).(WorkloadProfileResponseArrayOutput)
}

// Whether or not this Managed Environment is zone-redundant.
func (o LookupManagedEnvironmentResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedEnvironmentResultOutput{})
}
