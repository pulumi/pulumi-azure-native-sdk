// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves details of this Network Packet Broker.
//
// Uses Azure REST API version 2023-06-15.
func LookupNetworkPacketBroker(ctx *pulumi.Context, args *LookupNetworkPacketBrokerArgs, opts ...pulumi.InvokeOption) (*LookupNetworkPacketBrokerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkPacketBrokerResult
	err := ctx.Invoke("azure-native:managednetworkfabric:getNetworkPacketBroker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkPacketBrokerArgs struct {
	// Name of the Network Packet Broker.
	NetworkPacketBrokerName string `pulumi:"networkPacketBrokerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The NetworkPacketBroker resource definition.
type LookupNetworkPacketBrokerResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of neighbor group IDs configured on NPB.
	NeighborGroupIds []string `pulumi:"neighborGroupIds"`
	// List of ARM resource IDs of Network Devices [NPB].
	NetworkDeviceIds []string `pulumi:"networkDeviceIds"`
	// ARM resource ID of the Network Fabric.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// List of network Tap IDs configured on NPB.
	NetworkTapIds []string `pulumi:"networkTapIds"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of network interfaces across NPB devices that are used to mirror source traffic.
	SourceInterfaceIds []string `pulumi:"sourceInterfaceIds"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNetworkPacketBrokerOutput(ctx *pulumi.Context, args LookupNetworkPacketBrokerOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkPacketBrokerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkPacketBrokerResultOutput, error) {
			args := v.(LookupNetworkPacketBrokerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:managednetworkfabric:getNetworkPacketBroker", args, LookupNetworkPacketBrokerResultOutput{}, options).(LookupNetworkPacketBrokerResultOutput), nil
		}).(LookupNetworkPacketBrokerResultOutput)
}

type LookupNetworkPacketBrokerOutputArgs struct {
	// Name of the Network Packet Broker.
	NetworkPacketBrokerName pulumi.StringInput `pulumi:"networkPacketBrokerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkPacketBrokerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkPacketBrokerArgs)(nil)).Elem()
}

// The NetworkPacketBroker resource definition.
type LookupNetworkPacketBrokerResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkPacketBrokerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkPacketBrokerResult)(nil)).Elem()
}

func (o LookupNetworkPacketBrokerResultOutput) ToLookupNetworkPacketBrokerResultOutput() LookupNetworkPacketBrokerResultOutput {
	return o
}

func (o LookupNetworkPacketBrokerResultOutput) ToLookupNetworkPacketBrokerResultOutputWithContext(ctx context.Context) LookupNetworkPacketBrokerResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupNetworkPacketBrokerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNetworkPacketBrokerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkPacketBrokerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNetworkPacketBrokerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of neighbor group IDs configured on NPB.
func (o LookupNetworkPacketBrokerResultOutput) NeighborGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) []string { return v.NeighborGroupIds }).(pulumi.StringArrayOutput)
}

// List of ARM resource IDs of Network Devices [NPB].
func (o LookupNetworkPacketBrokerResultOutput) NetworkDeviceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) []string { return v.NetworkDeviceIds }).(pulumi.StringArrayOutput)
}

// ARM resource ID of the Network Fabric.
func (o LookupNetworkPacketBrokerResultOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// List of network Tap IDs configured on NPB.
func (o LookupNetworkPacketBrokerResultOutput) NetworkTapIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) []string { return v.NetworkTapIds }).(pulumi.StringArrayOutput)
}

// Provisioning state of the resource.
func (o LookupNetworkPacketBrokerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of network interfaces across NPB devices that are used to mirror source traffic.
func (o LookupNetworkPacketBrokerResultOutput) SourceInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) []string { return v.SourceInterfaceIds }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkPacketBrokerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkPacketBrokerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkPacketBrokerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkPacketBrokerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkPacketBrokerResultOutput{})
}
