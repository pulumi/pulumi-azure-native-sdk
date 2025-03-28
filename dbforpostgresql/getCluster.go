// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a cluster such as compute and storage configuration and cluster lifecycle metadata such as cluster creation date and time.
//
// Uses Azure REST API version 2022-11-08.
//
// Other available API versions: 2023-03-02-preview.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:dbforpostgresql:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupClusterArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a cluster.
type LookupClusterResult struct {
	// The administrator's login name of the servers in the cluster.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The Citus extension version on all cluster servers.
	CitusVersion *string `pulumi:"citusVersion"`
	// If public access is enabled on coordinator.
	CoordinatorEnablePublicIpAccess *bool `pulumi:"coordinatorEnablePublicIpAccess"`
	// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
	CoordinatorServerEdition *string `pulumi:"coordinatorServerEdition"`
	// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorStorageQuotaInMb *int `pulumi:"coordinatorStorageQuotaInMb"`
	// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	CoordinatorVCores *int `pulumi:"coordinatorVCores"`
	// The earliest restore point time (ISO8601 format) for the cluster.
	EarliestRestoreTime string `pulumi:"earliestRestoreTime"`
	// If high availability (HA) is enabled or not for the cluster.
	EnableHa *bool `pulumi:"enableHa"`
	// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
	EnableShardsOnCoordinator *bool `pulumi:"enableShardsOnCoordinator"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maintenance window of a cluster.
	MaintenanceWindow *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
	NodeCount *int `pulumi:"nodeCount"`
	// If public access is enabled on worker nodes.
	NodeEnablePublicIpAccess *bool `pulumi:"nodeEnablePublicIpAccess"`
	// The edition of a node server (default: MemoryOptimized).
	NodeServerEdition *string `pulumi:"nodeServerEdition"`
	// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeStorageQuotaInMb *int `pulumi:"nodeStorageQuotaInMb"`
	// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
	NodeVCores *int `pulumi:"nodeVCores"`
	// Date and time in UTC (ISO8601 format) for cluster restore.
	PointInTimeUTC *string `pulumi:"pointInTimeUTC"`
	// The major PostgreSQL version on all cluster servers.
	PostgresqlVersion *string `pulumi:"postgresqlVersion"`
	// Preferred primary availability zone (AZ) for all cluster servers.
	PreferredPrimaryZone *string `pulumi:"preferredPrimaryZone"`
	// The private endpoint connections for a cluster.
	PrivateEndpointConnections []SimplePrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the cluster
	ProvisioningState string `pulumi:"provisioningState"`
	// The array of read replica clusters.
	ReadReplicas []string `pulumi:"readReplicas"`
	// The list of server names in the cluster
	ServerNames []ServerNameItemResponse `pulumi:"serverNames"`
	// The Azure region of source cluster for read replica clusters.
	SourceLocation *string `pulumi:"sourceLocation"`
	// The resource id of source cluster for read replica clusters.
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// A state of a cluster/server that is visible to user.
	State string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupClusterResult
func (val *LookupClusterResult) Defaults() *LookupClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.MaintenanceWindow = tmp.MaintenanceWindow.Defaults()

	return &tmp
}
func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterResultOutput, error) {
			args := v.(LookupClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:dbforpostgresql:getCluster", args, LookupClusterResultOutput{}, options).(LookupClusterResultOutput), nil
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// Represents a cluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// The administrator's login name of the servers in the cluster.
func (o LookupClusterResultOutput) AdministratorLogin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.AdministratorLogin }).(pulumi.StringOutput)
}

// The Citus extension version on all cluster servers.
func (o LookupClusterResultOutput) CitusVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CitusVersion }).(pulumi.StringPtrOutput)
}

// If public access is enabled on coordinator.
func (o LookupClusterResultOutput) CoordinatorEnablePublicIpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.CoordinatorEnablePublicIpAccess }).(pulumi.BoolPtrOutput)
}

// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
func (o LookupClusterResultOutput) CoordinatorServerEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CoordinatorServerEdition }).(pulumi.StringPtrOutput)
}

// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o LookupClusterResultOutput) CoordinatorStorageQuotaInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.CoordinatorStorageQuotaInMb }).(pulumi.IntPtrOutput)
}

// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o LookupClusterResultOutput) CoordinatorVCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.CoordinatorVCores }).(pulumi.IntPtrOutput)
}

// The earliest restore point time (ISO8601 format) for the cluster.
func (o LookupClusterResultOutput) EarliestRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.EarliestRestoreTime }).(pulumi.StringOutput)
}

// If high availability (HA) is enabled or not for the cluster.
func (o LookupClusterResultOutput) EnableHa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableHa }).(pulumi.BoolPtrOutput)
}

// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
func (o LookupClusterResultOutput) EnableShardsOnCoordinator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableShardsOnCoordinator }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maintenance window of a cluster.
func (o LookupClusterResultOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *MaintenanceWindowResponse { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The name of the resource
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
func (o LookupClusterResultOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

// If public access is enabled on worker nodes.
func (o LookupClusterResultOutput) NodeEnablePublicIpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.NodeEnablePublicIpAccess }).(pulumi.BoolPtrOutput)
}

// The edition of a node server (default: MemoryOptimized).
func (o LookupClusterResultOutput) NodeServerEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.NodeServerEdition }).(pulumi.StringPtrOutput)
}

// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o LookupClusterResultOutput) NodeStorageQuotaInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.NodeStorageQuotaInMb }).(pulumi.IntPtrOutput)
}

// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
func (o LookupClusterResultOutput) NodeVCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.NodeVCores }).(pulumi.IntPtrOutput)
}

// Date and time in UTC (ISO8601 format) for cluster restore.
func (o LookupClusterResultOutput) PointInTimeUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PointInTimeUTC }).(pulumi.StringPtrOutput)
}

// The major PostgreSQL version on all cluster servers.
func (o LookupClusterResultOutput) PostgresqlVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PostgresqlVersion }).(pulumi.StringPtrOutput)
}

// Preferred primary availability zone (AZ) for all cluster servers.
func (o LookupClusterResultOutput) PreferredPrimaryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PreferredPrimaryZone }).(pulumi.StringPtrOutput)
}

// The private endpoint connections for a cluster.
func (o LookupClusterResultOutput) PrivateEndpointConnections() SimplePrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []SimplePrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(SimplePrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the cluster
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The array of read replica clusters.
func (o LookupClusterResultOutput) ReadReplicas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.ReadReplicas }).(pulumi.StringArrayOutput)
}

// The list of server names in the cluster
func (o LookupClusterResultOutput) ServerNames() ServerNameItemResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ServerNameItemResponse { return v.ServerNames }).(ServerNameItemResponseArrayOutput)
}

// The Azure region of source cluster for read replica clusters.
func (o LookupClusterResultOutput) SourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.SourceLocation }).(pulumi.StringPtrOutput)
}

// The resource id of source cluster for read replica clusters.
func (o LookupClusterResultOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

// A state of a cluster/server that is visible to user.
func (o LookupClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
