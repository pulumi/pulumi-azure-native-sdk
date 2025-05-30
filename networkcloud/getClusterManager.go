// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the properties of the provided cluster manager.
//
// Uses Azure REST API version 2025-02-01.
//
// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupClusterManager(ctx *pulumi.Context, args *LookupClusterManagerArgs, opts ...pulumi.InvokeOption) (*LookupClusterManagerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterManagerResult
	err := ctx.Invoke("azure-native:networkcloud:getClusterManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterManagerArgs struct {
	// The name of the cluster manager.
	ClusterManagerName string `pulumi:"clusterManagerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupClusterManagerResult struct {
	// The resource ID of the Log Analytics workspace that is used for the logs collection.
	AnalyticsWorkspaceId *string `pulumi:"analyticsWorkspaceId"`
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The list of the cluster versions the manager supports. It is used as input in clusterVersion property of a cluster resource.
	ClusterVersions []ClusterAvailableVersionResponse `pulumi:"clusterVersions"`
	// The detailed status that provides additional information about the cluster manager.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// Resource ETag.
	Etag string `pulumi:"etag"`
	// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
	FabricControllerId string `pulumi:"fabricControllerId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The identity of the cluster manager.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfigurationResponse `pulumi:"managedResourceGroupConfiguration"`
	// The extended location (custom location) that represents the cluster manager's control plane location. This extended location is used when creating cluster and rack manifest resources.
	ManagerExtendedLocation ExtendedLocationResponse `pulumi:"managerExtendedLocation"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the cluster manager.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
	VmSize *string `pulumi:"vmSize"`
}

func LookupClusterManagerOutput(ctx *pulumi.Context, args LookupClusterManagerOutputArgs, opts ...pulumi.InvokeOption) LookupClusterManagerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterManagerResultOutput, error) {
			args := v.(LookupClusterManagerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:networkcloud:getClusterManager", args, LookupClusterManagerResultOutput{}, options).(LookupClusterManagerResultOutput), nil
		}).(LookupClusterManagerResultOutput)
}

type LookupClusterManagerOutputArgs struct {
	// The name of the cluster manager.
	ClusterManagerName pulumi.StringInput `pulumi:"clusterManagerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterManagerArgs)(nil)).Elem()
}

type LookupClusterManagerResultOutput struct{ *pulumi.OutputState }

func (LookupClusterManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterManagerResult)(nil)).Elem()
}

func (o LookupClusterManagerResultOutput) ToLookupClusterManagerResultOutput() LookupClusterManagerResultOutput {
	return o
}

func (o LookupClusterManagerResultOutput) ToLookupClusterManagerResultOutputWithContext(ctx context.Context) LookupClusterManagerResultOutput {
	return o
}

// The resource ID of the Log Analytics workspace that is used for the logs collection.
func (o LookupClusterManagerResultOutput) AnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) *string { return v.AnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
func (o LookupClusterManagerResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o LookupClusterManagerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The list of the cluster versions the manager supports. It is used as input in clusterVersion property of a cluster resource.
func (o LookupClusterManagerResultOutput) ClusterVersions() ClusterAvailableVersionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) []ClusterAvailableVersionResponse { return v.ClusterVersions }).(ClusterAvailableVersionResponseArrayOutput)
}

// The detailed status that provides additional information about the cluster manager.
func (o LookupClusterManagerResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupClusterManagerResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// Resource ETag.
func (o LookupClusterManagerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
func (o LookupClusterManagerResultOutput) FabricControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.FabricControllerId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupClusterManagerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the cluster manager.
func (o LookupClusterManagerResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupClusterManagerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The configuration of the managed resource group associated with the resource.
func (o LookupClusterManagerResultOutput) ManagedResourceGroupConfiguration() ManagedResourceGroupConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) *ManagedResourceGroupConfigurationResponse {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedResourceGroupConfigurationResponsePtrOutput)
}

// The extended location (custom location) that represents the cluster manager's control plane location. This extended location is used when creating cluster and rack manifest resources.
func (o LookupClusterManagerResultOutput) ManagerExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) ExtendedLocationResponse { return v.ManagerExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The name of the resource
func (o LookupClusterManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the cluster manager.
func (o LookupClusterManagerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupClusterManagerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterManagerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterManagerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) string { return v.Type }).(pulumi.StringOutput)
}

// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
func (o LookupClusterManagerResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterManagerResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterManagerResultOutput{})
}
