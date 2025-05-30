// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App session pool.
//
// Uses Azure REST API version 2024-10-02-preview.
//
// Other available API versions: 2024-02-02-preview, 2024-08-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupContainerAppsSessionPool(ctx *pulumi.Context, args *LookupContainerAppsSessionPoolArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppsSessionPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerAppsSessionPoolResult
	err := ctx.Invoke("azure-native:app:getContainerAppsSessionPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerAppsSessionPoolArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the session pool.
	SessionPoolName string `pulumi:"sessionPoolName"`
}

// Container App session pool.
type LookupContainerAppsSessionPoolResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The container type of the sessions.
	ContainerType *string `pulumi:"containerType"`
	// The custom container configuration if the containerType is CustomContainer.
	CustomContainerTemplate *CustomContainerTemplateResponse `pulumi:"customContainerTemplate"`
	// The pool configuration if the poolManagementType is dynamic.
	DynamicPoolConfiguration *DynamicPoolConfigurationResponse `pulumi:"dynamicPoolConfiguration"`
	// Resource ID of the session pool's environment.
	EnvironmentId *string `pulumi:"environmentId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Managed identities needed by a session pool to interact with other Azure services to not maintain any secrets or credentials in code.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Optional settings for a Managed Identity that is assigned to the Session pool.
	ManagedIdentitySettings []ManagedIdentitySettingResponse `pulumi:"managedIdentitySettings"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The number of nodes the session pool is using.
	NodeCount int `pulumi:"nodeCount"`
	// The endpoint to manage the pool.
	PoolManagementEndpoint string `pulumi:"poolManagementEndpoint"`
	// The pool management type of the session pool.
	PoolManagementType *string `pulumi:"poolManagementType"`
	// Provisioning state of the session pool.
	ProvisioningState string `pulumi:"provisioningState"`
	// The scale configuration of the session pool.
	ScaleConfiguration *ScaleConfigurationResponse `pulumi:"scaleConfiguration"`
	// The secrets of the session pool.
	Secrets []SessionPoolSecretResponse `pulumi:"secrets"`
	// The network configuration of the sessions in the session pool.
	SessionNetworkConfiguration *SessionNetworkConfigurationResponse `pulumi:"sessionNetworkConfiguration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupContainerAppsSessionPoolOutput(ctx *pulumi.Context, args LookupContainerAppsSessionPoolOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppsSessionPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupContainerAppsSessionPoolResultOutput, error) {
			args := v.(LookupContainerAppsSessionPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:app:getContainerAppsSessionPool", args, LookupContainerAppsSessionPoolResultOutput{}, options).(LookupContainerAppsSessionPoolResultOutput), nil
		}).(LookupContainerAppsSessionPoolResultOutput)
}

type LookupContainerAppsSessionPoolOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the session pool.
	SessionPoolName pulumi.StringInput `pulumi:"sessionPoolName"`
}

func (LookupContainerAppsSessionPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsSessionPoolArgs)(nil)).Elem()
}

// Container App session pool.
type LookupContainerAppsSessionPoolResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppsSessionPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsSessionPoolResult)(nil)).Elem()
}

func (o LookupContainerAppsSessionPoolResultOutput) ToLookupContainerAppsSessionPoolResultOutput() LookupContainerAppsSessionPoolResultOutput {
	return o
}

func (o LookupContainerAppsSessionPoolResultOutput) ToLookupContainerAppsSessionPoolResultOutputWithContext(ctx context.Context) LookupContainerAppsSessionPoolResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupContainerAppsSessionPoolResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The container type of the sessions.
func (o LookupContainerAppsSessionPoolResultOutput) ContainerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *string { return v.ContainerType }).(pulumi.StringPtrOutput)
}

// The custom container configuration if the containerType is CustomContainer.
func (o LookupContainerAppsSessionPoolResultOutput) CustomContainerTemplate() CustomContainerTemplateResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *CustomContainerTemplateResponse {
		return v.CustomContainerTemplate
	}).(CustomContainerTemplateResponsePtrOutput)
}

// The pool configuration if the poolManagementType is dynamic.
func (o LookupContainerAppsSessionPoolResultOutput) DynamicPoolConfiguration() DynamicPoolConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *DynamicPoolConfigurationResponse {
		return v.DynamicPoolConfiguration
	}).(DynamicPoolConfigurationResponsePtrOutput)
}

// Resource ID of the session pool's environment.
func (o LookupContainerAppsSessionPoolResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupContainerAppsSessionPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed identities needed by a session pool to interact with other Azure services to not maintain any secrets or credentials in code.
func (o LookupContainerAppsSessionPoolResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupContainerAppsSessionPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Optional settings for a Managed Identity that is assigned to the Session pool.
func (o LookupContainerAppsSessionPoolResultOutput) ManagedIdentitySettings() ManagedIdentitySettingResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) []ManagedIdentitySettingResponse {
		return v.ManagedIdentitySettings
	}).(ManagedIdentitySettingResponseArrayOutput)
}

// The name of the resource
func (o LookupContainerAppsSessionPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes the session pool is using.
func (o LookupContainerAppsSessionPoolResultOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) int { return v.NodeCount }).(pulumi.IntOutput)
}

// The endpoint to manage the pool.
func (o LookupContainerAppsSessionPoolResultOutput) PoolManagementEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.PoolManagementEndpoint }).(pulumi.StringOutput)
}

// The pool management type of the session pool.
func (o LookupContainerAppsSessionPoolResultOutput) PoolManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *string { return v.PoolManagementType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the session pool.
func (o LookupContainerAppsSessionPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The scale configuration of the session pool.
func (o LookupContainerAppsSessionPoolResultOutput) ScaleConfiguration() ScaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *ScaleConfigurationResponse { return v.ScaleConfiguration }).(ScaleConfigurationResponsePtrOutput)
}

// The secrets of the session pool.
func (o LookupContainerAppsSessionPoolResultOutput) Secrets() SessionPoolSecretResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) []SessionPoolSecretResponse { return v.Secrets }).(SessionPoolSecretResponseArrayOutput)
}

// The network configuration of the sessions in the session pool.
func (o LookupContainerAppsSessionPoolResultOutput) SessionNetworkConfiguration() SessionNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) *SessionNetworkConfigurationResponse {
		return v.SessionNetworkConfiguration
	}).(SessionNetworkConfigurationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupContainerAppsSessionPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupContainerAppsSessionPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContainerAppsSessionPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSessionPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppsSessionPoolResultOutput{})
}
