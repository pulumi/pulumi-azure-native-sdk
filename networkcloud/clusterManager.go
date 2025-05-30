// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2025-02-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.
//
// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ClusterManager struct {
	pulumi.CustomResourceState

	// The resource ID of the Log Analytics workspace that is used for the logs collection.
	AnalyticsWorkspaceId pulumi.StringPtrOutput `pulumi:"analyticsWorkspaceId"`
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The list of the cluster versions the manager supports. It is used as input in clusterVersion property of a cluster resource.
	ClusterVersions ClusterAvailableVersionResponseArrayOutput `pulumi:"clusterVersions"`
	// The detailed status that provides additional information about the cluster manager.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// Resource ETag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
	FabricControllerId pulumi.StringOutput `pulumi:"fabricControllerId"`
	// The identity of the cluster manager.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationResponsePtrOutput `pulumi:"managedResourceGroupConfiguration"`
	// The extended location (custom location) that represents the cluster manager's control plane location. This extended location is used when creating cluster and rack manifest resources.
	ManagerExtendedLocation ExtendedLocationResponseOutput `pulumi:"managerExtendedLocation"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the cluster manager.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
}

// NewClusterManager registers a new resource with the given unique name, arguments, and options.
func NewClusterManager(ctx *pulumi.Context,
	name string, args *ClusterManagerArgs, opts ...pulumi.ResourceOption) (*ClusterManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricControllerId == nil {
		return nil, errors.New("invalid value for required argument 'FabricControllerId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:ClusterManager"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20231001preview:ClusterManager"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20240601preview:ClusterManager"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20240701:ClusterManager"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20241001preview:ClusterManager"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20250201:ClusterManager"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterManager
	err := ctx.RegisterResource("azure-native:networkcloud:ClusterManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterManager gets an existing ClusterManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterManagerState, opts ...pulumi.ResourceOption) (*ClusterManager, error) {
	var resource ClusterManager
	err := ctx.ReadResource("azure-native:networkcloud:ClusterManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterManager resources.
type clusterManagerState struct {
}

type ClusterManagerState struct {
}

func (ClusterManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterManagerState)(nil)).Elem()
}

type clusterManagerArgs struct {
	// The resource ID of the Log Analytics workspace that is used for the logs collection.
	AnalyticsWorkspaceId *string `pulumi:"analyticsWorkspaceId"`
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The name of the cluster manager.
	ClusterManagerName *string `pulumi:"clusterManagerName"`
	// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
	FabricControllerId string `pulumi:"fabricControllerId"`
	// The identity of the cluster manager.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration `pulumi:"managedResourceGroupConfiguration"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
	VmSize *string `pulumi:"vmSize"`
}

// The set of arguments for constructing a ClusterManager resource.
type ClusterManagerArgs struct {
	// The resource ID of the Log Analytics workspace that is used for the logs collection.
	AnalyticsWorkspaceId pulumi.StringPtrInput
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
	AvailabilityZones pulumi.StringArrayInput
	// The name of the cluster manager.
	ClusterManagerName pulumi.StringPtrInput
	// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
	FabricControllerId pulumi.StringInput
	// The identity of the cluster manager.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
	VmSize pulumi.StringPtrInput
}

func (ClusterManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterManagerArgs)(nil)).Elem()
}

type ClusterManagerInput interface {
	pulumi.Input

	ToClusterManagerOutput() ClusterManagerOutput
	ToClusterManagerOutputWithContext(ctx context.Context) ClusterManagerOutput
}

func (*ClusterManager) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterManager)(nil)).Elem()
}

func (i *ClusterManager) ToClusterManagerOutput() ClusterManagerOutput {
	return i.ToClusterManagerOutputWithContext(context.Background())
}

func (i *ClusterManager) ToClusterManagerOutputWithContext(ctx context.Context) ClusterManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterManagerOutput)
}

type ClusterManagerOutput struct{ *pulumi.OutputState }

func (ClusterManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterManager)(nil)).Elem()
}

func (o ClusterManagerOutput) ToClusterManagerOutput() ClusterManagerOutput {
	return o
}

func (o ClusterManagerOutput) ToClusterManagerOutputWithContext(ctx context.Context) ClusterManagerOutput {
	return o
}

// The resource ID of the Log Analytics workspace that is used for the logs collection.
func (o ClusterManagerOutput) AnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringPtrOutput { return v.AnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The Azure availability zones within the region that will be used to support the cluster manager resource.
func (o ClusterManagerOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o ClusterManagerOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The list of the cluster versions the manager supports. It is used as input in clusterVersion property of a cluster resource.
func (o ClusterManagerOutput) ClusterVersions() ClusterAvailableVersionResponseArrayOutput {
	return o.ApplyT(func(v *ClusterManager) ClusterAvailableVersionResponseArrayOutput { return v.ClusterVersions }).(ClusterAvailableVersionResponseArrayOutput)
}

// The detailed status that provides additional information about the cluster manager.
func (o ClusterManagerOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o ClusterManagerOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// Resource ETag.
func (o ClusterManagerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The resource ID of the fabric controller that has one to one mapping with the cluster manager.
func (o ClusterManagerOutput) FabricControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.FabricControllerId }).(pulumi.StringOutput)
}

// The identity of the cluster manager.
func (o ClusterManagerOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ClusterManager) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ClusterManagerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The configuration of the managed resource group associated with the resource.
func (o ClusterManagerOutput) ManagedResourceGroupConfiguration() ManagedResourceGroupConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ClusterManager) ManagedResourceGroupConfigurationResponsePtrOutput {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedResourceGroupConfigurationResponsePtrOutput)
}

// The extended location (custom location) that represents the cluster manager's control plane location. This extended location is used when creating cluster and rack manifest resources.
func (o ClusterManagerOutput) ManagerExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ClusterManager) ExtendedLocationResponseOutput { return v.ManagerExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The name of the resource
func (o ClusterManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the cluster manager.
func (o ClusterManagerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ClusterManagerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ClusterManager) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterManagerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Field deprecated, this value will no longer influence the cluster manager allocation process and will be removed in a future version. The size of the Azure virtual machines to use for hosting the cluster manager resource.
func (o ClusterManagerOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterManager) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterManagerOutput{})
}
