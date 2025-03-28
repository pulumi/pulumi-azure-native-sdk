// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Service Fabric cluster resource created or in the process of being created in the specified resource group.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20201201preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupClusterArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The cluster resource
type LookupClusterResult struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures []string `pulumi:"addOnFeatures"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicyResponse `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// The Service Fabric runtime versions available for this cluster.
	AvailableClusterVersions []ClusterVersionDetailsResponse `pulumi:"availableClusterVersions"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectoryResponse `pulumi:"azureActiveDirectory"`
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate *CertificateDescriptionResponse `pulumi:"certificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames *ServerCertificateCommonNamesResponse `pulumi:"certificateCommonNames"`
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames []ClientCertificateCommonNameResponse `pulumi:"clientCertificateCommonNames"`
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints []ClientCertificateThumbprintResponse `pulumi:"clientCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	// A service generated unique identifier for the cluster resource.
	ClusterId string `pulumi:"clusterId"`
	// The current state of the cluster.
	//
	//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
	//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
	//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
	//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
	//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
	//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
	//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
	//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
	//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
	//   - Ready - Indicates that the cluster is in a stable state.
	ClusterState string `pulumi:"clusterState"`
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfigResponse `pulumi:"diagnosticsStorageAccountConfig"`
	// Azure resource etag.
	Etag string `pulumi:"etag"`
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled *bool `pulumi:"eventStoreServiceEnabled"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescriptionResponse `pulumi:"fabricSettings"`
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// Azure resource location.
	Location string `pulumi:"location"`
	// The http management endpoint of the cluster.
	ManagementEndpoint string `pulumi:"managementEndpoint"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// The list of node types in the cluster.
	NodeTypes []NodeTypeDescriptionResponse `pulumi:"nodeTypes"`
	// The provisioning state of the cluster resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel *string `pulumi:"reliabilityLevel"`
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate *CertificateDescriptionResponse `pulumi:"reverseProxyCertificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames *ServerCertificateCommonNamesResponse `pulumi:"reverseProxyCertificateCommonNames"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
	// The policy to use when upgrading the cluster.
	UpgradeDescription *ClusterUpgradePolicyResponse `pulumi:"upgradeDescription"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	//
	//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
	//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage *string `pulumi:"vmImage"`
}

// Defaults sets the appropriate defaults for LookupClusterResult
func (val *LookupClusterResult) Defaults() *LookupClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.UpgradeDescription = tmp.UpgradeDescription.Defaults()

	return &tmp
}
func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterResultOutput, error) {
			args := v.(LookupClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:servicefabric/v20201201preview:getCluster", args, LookupClusterResultOutput{}, options).(LookupClusterResultOutput), nil
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// The cluster resource
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

// The list of add-on features to enable in the cluster.
func (o LookupClusterResultOutput) AddOnFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AddOnFeatures }).(pulumi.StringArrayOutput)
}

// The policy used to clean up unused versions.
func (o LookupClusterResultOutput) ApplicationTypeVersionsCleanupPolicy() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ApplicationTypeVersionsCleanupPolicyResponse {
		return v.ApplicationTypeVersionsCleanupPolicy
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

// The Service Fabric runtime versions available for this cluster.
func (o LookupClusterResultOutput) AvailableClusterVersions() ClusterVersionDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClusterVersionDetailsResponse { return v.AvailableClusterVersions }).(ClusterVersionDetailsResponseArrayOutput)
}

// The AAD authentication settings of the cluster.
func (o LookupClusterResultOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
func (o LookupClusterResultOutput) Certificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *CertificateDescriptionResponse { return v.Certificate }).(CertificateDescriptionResponsePtrOutput)
}

// Describes a list of server certificates referenced by common name that are used to secure the cluster.
func (o LookupClusterResultOutput) CertificateCommonNames() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ServerCertificateCommonNamesResponse { return v.CertificateCommonNames }).(ServerCertificateCommonNamesResponsePtrOutput)
}

// The list of client certificates referenced by common name that are allowed to manage the cluster.
func (o LookupClusterResultOutput) ClientCertificateCommonNames() ClientCertificateCommonNameResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClientCertificateCommonNameResponse {
		return v.ClientCertificateCommonNames
	}).(ClientCertificateCommonNameResponseArrayOutput)
}

// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
func (o LookupClusterResultOutput) ClientCertificateThumbprints() ClientCertificateThumbprintResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClientCertificateThumbprintResponse {
		return v.ClientCertificateThumbprints
	}).(ClientCertificateThumbprintResponseArrayOutput)
}

// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
func (o LookupClusterResultOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
func (o LookupClusterResultOutput) ClusterEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterEndpoint }).(pulumi.StringOutput)
}

// A service generated unique identifier for the cluster resource.
func (o LookupClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The current state of the cluster.
//
//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
//   - Ready - Indicates that the cluster is in a stable state.
func (o LookupClusterResultOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterState }).(pulumi.StringOutput)
}

// The storage account information for storing Service Fabric diagnostic logs.
func (o LookupClusterResultOutput) DiagnosticsStorageAccountConfig() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *DiagnosticsStorageAccountConfigResponse {
		return v.DiagnosticsStorageAccountConfig
	}).(DiagnosticsStorageAccountConfigResponsePtrOutput)
}

// Azure resource etag.
func (o LookupClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Indicates if the event store service is enabled.
func (o LookupClusterResultOutput) EventStoreServiceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EventStoreServiceEnabled }).(pulumi.BoolPtrOutput)
}

// The list of custom fabric settings to configure the cluster.
func (o LookupClusterResultOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []SettingsSectionDescriptionResponse { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

// Azure resource identifier.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure resource location.
func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The http management endpoint of the cluster.
func (o LookupClusterResultOutput) ManagementEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ManagementEndpoint }).(pulumi.StringOutput)
}

// Azure resource name.
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of node types in the cluster.
func (o LookupClusterResultOutput) NodeTypes() NodeTypeDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []NodeTypeDescriptionResponse { return v.NodeTypes }).(NodeTypeDescriptionResponseArrayOutput)
}

// The provisioning state of the cluster resource.
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
//
//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
//   - Silver - Run the System services with a target replica set count of 5.
//   - Gold - Run the System services with a target replica set count of 7.
//   - Platinum - Run the System services with a target replica set count of 9.
func (o LookupClusterResultOutput) ReliabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ReliabilityLevel }).(pulumi.StringPtrOutput)
}

// The server certificate used by reverse proxy.
func (o LookupClusterResultOutput) ReverseProxyCertificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *CertificateDescriptionResponse { return v.ReverseProxyCertificate }).(CertificateDescriptionResponsePtrOutput)
}

// Describes a list of server certificates referenced by common name that are used to secure the cluster.
func (o LookupClusterResultOutput) ReverseProxyCertificateCommonNames() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ServerCertificateCommonNamesResponse {
		return v.ReverseProxyCertificateCommonNames
	}).(ServerCertificateCommonNamesResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy to use when upgrading the cluster.
func (o LookupClusterResultOutput) UpgradeDescription() ClusterUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterUpgradePolicyResponse { return v.UpgradeDescription }).(ClusterUpgradePolicyResponsePtrOutput)
}

// The upgrade mode of the cluster when new Service Fabric runtime version is available.
//
//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
func (o LookupClusterResultOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
func (o LookupClusterResultOutput) VmImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.VmImage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
