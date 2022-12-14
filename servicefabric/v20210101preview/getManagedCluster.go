// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The manged cluster resource
func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20210101preview:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The manged cluster resource
type LookupManagedClusterResult struct {
	// List of add-on features to enable on the cluster.
	AddonFeatures []string `pulumi:"addonFeatures"`
	// VM admin user password.
	AdminPassword *string `pulumi:"adminPassword"`
	// VM admin user name.
	AdminUserName string `pulumi:"adminUserName"`
	// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
	AllowRdpAccess *bool `pulumi:"allowRdpAccess"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicyResponse `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectoryResponse `pulumi:"azureActiveDirectory"`
	// The port used for client connections to the cluster.
	ClientConnectionPort *int `pulumi:"clientConnectionPort"`
	// Client certificates that are allowed to manage the cluster.
	Clients []ClientCertificateResponse `pulumi:"clients"`
	// List of thumbprints of the cluster certificates.
	ClusterCertificateThumbprints []string `pulumi:"clusterCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// A service generated unique identifier for the cluster resource.
	ClusterId string `pulumi:"clusterId"`
	// The current state of the cluster.
	ClusterState string `pulumi:"clusterState"`
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0.
	ClusterUpgradeCadence *string `pulumi:"clusterUpgradeCadence"`
	// The cluster dns name.
	DnsName string `pulumi:"dnsName"`
	// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
	EnableAutoOSUpgrade *bool `pulumi:"enableAutoOSUpgrade"`
	// Azure resource etag.
	Etag string `pulumi:"etag"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescriptionResponse `pulumi:"fabricSettings"`
	// The fully qualified domain name associated with the public load balancer of the cluster.
	Fqdn string `pulumi:"fqdn"`
	// The port used for HTTP connections to the cluster.
	HttpGatewayConnectionPort *int `pulumi:"httpGatewayConnectionPort"`
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// The IPv4 address associated with the public load balancer of the cluster.
	Ipv4Address string `pulumi:"ipv4Address"`
	// Load balancing rules that are applied to the public load balancer of the cluster.
	LoadBalancingRules []LoadBalancingRuleResponse `pulumi:"loadBalancingRules"`
	// Azure resource location.
	Location string `pulumi:"location"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// Custom Network Security Rules that are applied to the virtual network of the cluster.
	NetworkSecurityRules []NetworkSecurityRuleResponse `pulumi:"networkSecurityRules"`
	// The provisioning state of the managed cluster resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku of the managed cluster
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupManagedClusterResult
func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientConnectionPort) {
		clientConnectionPort_ := 19000
		tmp.ClientConnectionPort = &clientConnectionPort_
	}
	if isZero(tmp.HttpGatewayConnectionPort) {
		httpGatewayConnectionPort_ := 19080
		tmp.HttpGatewayConnectionPort = &httpGatewayConnectionPort_
	}
	return &tmp
}

func LookupManagedClusterOutput(ctx *pulumi.Context, args LookupManagedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupManagedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedClusterResult, error) {
			args := v.(LookupManagedClusterArgs)
			r, err := LookupManagedCluster(ctx, &args, opts...)
			var s LookupManagedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedClusterResultOutput)
}

type LookupManagedClusterOutputArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterArgs)(nil)).Elem()
}

// The manged cluster resource
type LookupManagedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupManagedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterResult)(nil)).Elem()
}

func (o LookupManagedClusterResultOutput) ToLookupManagedClusterResultOutput() LookupManagedClusterResultOutput {
	return o
}

func (o LookupManagedClusterResultOutput) ToLookupManagedClusterResultOutputWithContext(ctx context.Context) LookupManagedClusterResultOutput {
	return o
}

// List of add-on features to enable on the cluster.
func (o LookupManagedClusterResultOutput) AddonFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []string { return v.AddonFeatures }).(pulumi.StringArrayOutput)
}

// VM admin user password.
func (o LookupManagedClusterResultOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

// VM admin user name.
func (o LookupManagedClusterResultOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.AdminUserName }).(pulumi.StringOutput)
}

// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
func (o LookupManagedClusterResultOutput) AllowRdpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.AllowRdpAccess }).(pulumi.BoolPtrOutput)
}

// The policy used to clean up unused versions.
func (o LookupManagedClusterResultOutput) ApplicationTypeVersionsCleanupPolicy() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ApplicationTypeVersionsCleanupPolicyResponse {
		return v.ApplicationTypeVersionsCleanupPolicy
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

// The AAD authentication settings of the cluster.
func (o LookupManagedClusterResultOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

// The port used for client connections to the cluster.
func (o LookupManagedClusterResultOutput) ClientConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *int { return v.ClientConnectionPort }).(pulumi.IntPtrOutput)
}

// Client certificates that are allowed to manage the cluster.
func (o LookupManagedClusterResultOutput) Clients() ClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ClientCertificateResponse { return v.Clients }).(ClientCertificateResponseArrayOutput)
}

// List of thumbprints of the cluster certificates.
func (o LookupManagedClusterResultOutput) ClusterCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []string { return v.ClusterCertificateThumbprints }).(pulumi.StringArrayOutput)
}

// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
func (o LookupManagedClusterResultOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

// A service generated unique identifier for the cluster resource.
func (o LookupManagedClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The current state of the cluster.
func (o LookupManagedClusterResultOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterState }).(pulumi.StringOutput)
}

// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0.
func (o LookupManagedClusterResultOutput) ClusterUpgradeCadence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ClusterUpgradeCadence }).(pulumi.StringPtrOutput)
}

// The cluster dns name.
func (o LookupManagedClusterResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.DnsName }).(pulumi.StringOutput)
}

// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
func (o LookupManagedClusterResultOutput) EnableAutoOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableAutoOSUpgrade }).(pulumi.BoolPtrOutput)
}

// Azure resource etag.
func (o LookupManagedClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The list of custom fabric settings to configure the cluster.
func (o LookupManagedClusterResultOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []SettingsSectionDescriptionResponse { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

// The fully qualified domain name associated with the public load balancer of the cluster.
func (o LookupManagedClusterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// The port used for HTTP connections to the cluster.
func (o LookupManagedClusterResultOutput) HttpGatewayConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *int { return v.HttpGatewayConnectionPort }).(pulumi.IntPtrOutput)
}

// Azure resource identifier.
func (o LookupManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IPv4 address associated with the public load balancer of the cluster.
func (o LookupManagedClusterResultOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Ipv4Address }).(pulumi.StringOutput)
}

// Load balancing rules that are applied to the public load balancer of the cluster.
func (o LookupManagedClusterResultOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []LoadBalancingRuleResponse { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

// Azure resource location.
func (o LookupManagedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name.
func (o LookupManagedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Custom Network Security Rules that are applied to the virtual network of the cluster.
func (o LookupManagedClusterResultOutput) NetworkSecurityRules() NetworkSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []NetworkSecurityRuleResponse { return v.NetworkSecurityRules }).(NetworkSecurityRuleResponseArrayOutput)
}

// The provisioning state of the managed cluster resource.
func (o LookupManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku of the managed cluster
func (o LookupManagedClusterResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupManagedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupManagedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupManagedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterResultOutput{})
}
