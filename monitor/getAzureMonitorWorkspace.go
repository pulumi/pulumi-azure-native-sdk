// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specific Azure Monitor workspace
//
// Uses Azure REST API version 2023-10-01-preview.
//
// Other available API versions: 2023-04-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupAzureMonitorWorkspace(ctx *pulumi.Context, args *LookupAzureMonitorWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupAzureMonitorWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAzureMonitorWorkspaceResult
	err := ctx.Invoke("azure-native:monitor:getAzureMonitorWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureMonitorWorkspaceArgs struct {
	// The name of the Azure Monitor workspace. The name is case insensitive.
	AzureMonitorWorkspaceName string `pulumi:"azureMonitorWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Monitor Workspace definition.
type LookupAzureMonitorWorkspaceResult struct {
	// The immutable ID of the Azure Monitor workspace. This property is read-only.
	AccountId string `pulumi:"accountId"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The Data Collection Rule and Endpoint used for ingestion by default.
	DefaultIngestionSettings IngestionSettingsResponse `pulumi:"defaultIngestionSettings"`
	// Resource entity tag (ETag)
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Information about metrics for the Azure Monitor workspace
	Metrics *MetricsResponse `pulumi:"metrics"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state of the Azure Monitor workspace. Set to Succeeded if everything is healthy.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets or sets allow or disallow public network access to workspace
	PublicNetworkAccess string `pulumi:"publicNetworkAccess"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAzureMonitorWorkspaceOutput(ctx *pulumi.Context, args LookupAzureMonitorWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupAzureMonitorWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAzureMonitorWorkspaceResultOutput, error) {
			args := v.(LookupAzureMonitorWorkspaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:monitor:getAzureMonitorWorkspace", args, LookupAzureMonitorWorkspaceResultOutput{}, options).(LookupAzureMonitorWorkspaceResultOutput), nil
		}).(LookupAzureMonitorWorkspaceResultOutput)
}

type LookupAzureMonitorWorkspaceOutputArgs struct {
	// The name of the Azure Monitor workspace. The name is case insensitive.
	AzureMonitorWorkspaceName pulumi.StringInput `pulumi:"azureMonitorWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureMonitorWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureMonitorWorkspaceArgs)(nil)).Elem()
}

// An Azure Monitor Workspace definition.
type LookupAzureMonitorWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupAzureMonitorWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureMonitorWorkspaceResult)(nil)).Elem()
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToLookupAzureMonitorWorkspaceResultOutput() LookupAzureMonitorWorkspaceResultOutput {
	return o
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToLookupAzureMonitorWorkspaceResultOutputWithContext(ctx context.Context) LookupAzureMonitorWorkspaceResultOutput {
	return o
}

// The immutable ID of the Azure Monitor workspace. This property is read-only.
func (o LookupAzureMonitorWorkspaceResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupAzureMonitorWorkspaceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The Data Collection Rule and Endpoint used for ingestion by default.
func (o LookupAzureMonitorWorkspaceResultOutput) DefaultIngestionSettings() IngestionSettingsResponseOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) IngestionSettingsResponse { return v.DefaultIngestionSettings }).(IngestionSettingsResponseOutput)
}

// Resource entity tag (ETag)
func (o LookupAzureMonitorWorkspaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupAzureMonitorWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupAzureMonitorWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Information about metrics for the Azure Monitor workspace
func (o LookupAzureMonitorWorkspaceResultOutput) Metrics() MetricsResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) *MetricsResponse { return v.Metrics }).(MetricsResponsePtrOutput)
}

// The name of the resource
func (o LookupAzureMonitorWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections.
func (o LookupAzureMonitorWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state of the Azure Monitor workspace. Set to Succeeded if everything is healthy.
func (o LookupAzureMonitorWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets allow or disallow public network access to workspace
func (o LookupAzureMonitorWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAzureMonitorWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAzureMonitorWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAzureMonitorWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureMonitorWorkspaceResultOutput{})
}
