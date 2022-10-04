// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a connected cluster.
func LookupConnectedCluster(ctx *pulumi.Context, args *LookupConnectedClusterArgs, opts ...pulumi.InvokeOption) (*LookupConnectedClusterResult, error) {
	var rv LookupConnectedClusterResult
	err := ctx.Invoke("azure-native:kubernetes/v20221001preview:getConnectedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedClusterArgs struct {
	// The name of the Kubernetes cluster on which get is called.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a connected cluster.
type LookupConnectedClusterResult struct {
	// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate string `pulumi:"agentPublicKeyCertificate"`
	// Version of the agent running on the connected cluster resource
	AgentVersion string `pulumi:"agentVersion"`
	// Indicates whether Azure Hybrid Benefit is opted in
	AzureHybridBenefit *string `pulumi:"azureHybridBenefit"`
	// Represents the connectivity status of the connected cluster.
	ConnectivityStatus string `pulumi:"connectivityStatus"`
	// The Kubernetes distribution running on this connected cluster.
	Distribution *string `pulumi:"distribution"`
	// The Kubernetes distribution version on this connected cluster.
	DistributionVersion *string `pulumi:"distributionVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the connected cluster.
	Identity ConnectedClusterIdentityResponse `pulumi:"identity"`
	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure *string `pulumi:"infrastructure"`
	// The Kubernetes version of the connected cluster resource
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// Time representing the last instance when heart beat was received from the cluster
	LastConnectivityTime string `pulumi:"lastConnectivityTime"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Expiration time of the managed identity certificate
	ManagedIdentityCertificateExpirationTime string `pulumi:"managedIdentityCertificateExpirationTime"`
	// More properties related to the Connected Cluster
	MiscellaneousProperties map[string]string `pulumi:"miscellaneousProperties"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Connected cluster offering
	Offering string `pulumi:"offering"`
	// The resource id of the private link scope this connected cluster is assigned to, if any.
	PrivateLinkScopeResourceId *string `pulumi:"privateLinkScopeResourceId"`
	// Property which describes the state of private link on a connected cluster resource.
	PrivateLinkState *string `pulumi:"privateLinkState"`
	// Provisioning state of the connected cluster resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Number of CPU cores present in the connected cluster resource
	TotalCoreCount int `pulumi:"totalCoreCount"`
	// Number of nodes present in the connected cluster resource
	TotalNodeCount int `pulumi:"totalNodeCount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupConnectedClusterResult
func (val *LookupConnectedClusterResult) Defaults() *LookupConnectedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AzureHybridBenefit) {
		azureHybridBenefit_ := "NotApplicable"
		tmp.AzureHybridBenefit = &azureHybridBenefit_
	}
	tmp.Identity = *tmp.Identity.Defaults()

	if isZero(tmp.PrivateLinkState) {
		privateLinkState_ := "Disabled"
		tmp.PrivateLinkState = &privateLinkState_
	}
	return &tmp
}

func LookupConnectedClusterOutput(ctx *pulumi.Context, args LookupConnectedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedClusterResult, error) {
			args := v.(LookupConnectedClusterArgs)
			r, err := LookupConnectedCluster(ctx, &args, opts...)
			var s LookupConnectedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedClusterResultOutput)
}

type LookupConnectedClusterOutputArgs struct {
	// The name of the Kubernetes cluster on which get is called.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedClusterArgs)(nil)).Elem()
}

// Represents a connected cluster.
type LookupConnectedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedClusterResult)(nil)).Elem()
}

func (o LookupConnectedClusterResultOutput) ToLookupConnectedClusterResultOutput() LookupConnectedClusterResultOutput {
	return o
}

func (o LookupConnectedClusterResultOutput) ToLookupConnectedClusterResultOutputWithContext(ctx context.Context) LookupConnectedClusterResultOutput {
	return o
}

// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
func (o LookupConnectedClusterResultOutput) AgentPublicKeyCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.AgentPublicKeyCertificate }).(pulumi.StringOutput)
}

// Version of the agent running on the connected cluster resource
func (o LookupConnectedClusterResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

// Indicates whether Azure Hybrid Benefit is opted in
func (o LookupConnectedClusterResultOutput) AzureHybridBenefit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.AzureHybridBenefit }).(pulumi.StringPtrOutput)
}

// Represents the connectivity status of the connected cluster.
func (o LookupConnectedClusterResultOutput) ConnectivityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.ConnectivityStatus }).(pulumi.StringOutput)
}

// The Kubernetes distribution running on this connected cluster.
func (o LookupConnectedClusterResultOutput) Distribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.Distribution }).(pulumi.StringPtrOutput)
}

// The Kubernetes distribution version on this connected cluster.
func (o LookupConnectedClusterResultOutput) DistributionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.DistributionVersion }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the connected cluster.
func (o LookupConnectedClusterResultOutput) Identity() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) ConnectedClusterIdentityResponse { return v.Identity }).(ConnectedClusterIdentityResponseOutput)
}

// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
func (o LookupConnectedClusterResultOutput) Infrastructure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.Infrastructure }).(pulumi.StringPtrOutput)
}

// The Kubernetes version of the connected cluster resource
func (o LookupConnectedClusterResultOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// Time representing the last instance when heart beat was received from the cluster
func (o LookupConnectedClusterResultOutput) LastConnectivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.LastConnectivityTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupConnectedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// Expiration time of the managed identity certificate
func (o LookupConnectedClusterResultOutput) ManagedIdentityCertificateExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.ManagedIdentityCertificateExpirationTime }).(pulumi.StringOutput)
}

// More properties related to the Connected Cluster
func (o LookupConnectedClusterResultOutput) MiscellaneousProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) map[string]string { return v.MiscellaneousProperties }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o LookupConnectedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Connected cluster offering
func (o LookupConnectedClusterResultOutput) Offering() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Offering }).(pulumi.StringOutput)
}

// The resource id of the private link scope this connected cluster is assigned to, if any.
func (o LookupConnectedClusterResultOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

// Property which describes the state of private link on a connected cluster resource.
func (o LookupConnectedClusterResultOutput) PrivateLinkState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.PrivateLinkState }).(pulumi.StringPtrOutput)
}

// Provisioning state of the connected cluster resource.
func (o LookupConnectedClusterResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LookupConnectedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupConnectedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Number of CPU cores present in the connected cluster resource
func (o LookupConnectedClusterResultOutput) TotalCoreCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) int { return v.TotalCoreCount }).(pulumi.IntOutput)
}

// Number of nodes present in the connected cluster resource
func (o LookupConnectedClusterResultOutput) TotalNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) int { return v.TotalNodeCount }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedClusterResultOutput{})
}