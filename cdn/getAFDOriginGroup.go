// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing origin group within a profile.
//
// Uses Azure REST API version 2024-09-01.
//
// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-04-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupAFDOriginGroup(ctx *pulumi.Context, args *LookupAFDOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupAFDOriginGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAFDOriginGroupResult
	err := ctx.Invoke("azure-native:cdn:getAFDOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDOriginGroupArgs struct {
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName string `pulumi:"originGroupName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from Azure Front Door.
type LookupAFDOriginGroupResult struct {
	// The Azure API version of the resource.
	AzureApiVersion  string `pulumi:"azureApiVersion"`
	DeploymentStatus string `pulumi:"deploymentStatus"`
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings *HealthProbeParametersResponse `pulumi:"healthProbeSettings"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Load balancing settings for a backend pool
	LoadBalancingSettings *LoadBalancingSettingsParametersResponse `pulumi:"loadBalancingSettings"`
	// Resource name.
	Name string `pulumi:"name"`
	// The name of the profile which holds the origin group.
	ProfileName string `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
	SessionAffinityState *string `pulumi:"sessionAffinityState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAFDOriginGroupOutput(ctx *pulumi.Context, args LookupAFDOriginGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAFDOriginGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAFDOriginGroupResultOutput, error) {
			args := v.(LookupAFDOriginGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cdn:getAFDOriginGroup", args, LookupAFDOriginGroupResultOutput{}, options).(LookupAFDOriginGroupResultOutput), nil
		}).(LookupAFDOriginGroupResultOutput)
}

type LookupAFDOriginGroupOutputArgs struct {
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName pulumi.StringInput `pulumi:"originGroupName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDOriginGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginGroupArgs)(nil)).Elem()
}

// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from Azure Front Door.
type LookupAFDOriginGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAFDOriginGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginGroupResult)(nil)).Elem()
}

func (o LookupAFDOriginGroupResultOutput) ToLookupAFDOriginGroupResultOutput() LookupAFDOriginGroupResultOutput {
	return o
}

func (o LookupAFDOriginGroupResultOutput) ToLookupAFDOriginGroupResultOutputWithContext(ctx context.Context) LookupAFDOriginGroupResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAFDOriginGroupResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

func (o LookupAFDOriginGroupResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Health probe settings to the origin that is used to determine the health of the origin.
func (o LookupAFDOriginGroupResultOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *HealthProbeParametersResponse { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

// Resource ID.
func (o LookupAFDOriginGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Load balancing settings for a backend pool
func (o LookupAFDOriginGroupResultOutput) LoadBalancingSettings() LoadBalancingSettingsParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *LoadBalancingSettingsParametersResponse {
		return v.LoadBalancingSettings
	}).(LoadBalancingSettingsParametersResponsePtrOutput)
}

// Resource name.
func (o LookupAFDOriginGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the profile which holds the origin group.
func (o LookupAFDOriginGroupResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o LookupAFDOriginGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
func (o LookupAFDOriginGroupResultOutput) SessionAffinityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *string { return v.SessionAffinityState }).(pulumi.StringPtrOutput)
}

// Read only system data
func (o LookupAFDOriginGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
func (o LookupAFDOriginGroupResultOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *int {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupAFDOriginGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDOriginGroupResultOutput{})
}
