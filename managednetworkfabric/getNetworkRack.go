// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Network Rack resource details.
//
// Uses Azure REST API version 2023-06-15.
//
// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupNetworkRack(ctx *pulumi.Context, args *LookupNetworkRackArgs, opts ...pulumi.InvokeOption) (*LookupNetworkRackResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkRackResult
	err := ctx.Invoke("azure-native:managednetworkfabric:getNetworkRack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkRackArgs struct {
	// Name of the Network Rack.
	NetworkRackName string `pulumi:"networkRackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Network Rack resource definition.
type LookupNetworkRackResult struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of network device ARM resource IDs.
	NetworkDevices []string `pulumi:"networkDevices"`
	// ARM resource ID of the Network Fabric.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// Network Rack SKU name.
	NetworkRackType *string `pulumi:"networkRackType"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNetworkRackOutput(ctx *pulumi.Context, args LookupNetworkRackOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkRackResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkRackResultOutput, error) {
			args := v.(LookupNetworkRackArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:managednetworkfabric:getNetworkRack", args, LookupNetworkRackResultOutput{}, options).(LookupNetworkRackResultOutput), nil
		}).(LookupNetworkRackResultOutput)
}

type LookupNetworkRackOutputArgs struct {
	// Name of the Network Rack.
	NetworkRackName pulumi.StringInput `pulumi:"networkRackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkRackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkRackArgs)(nil)).Elem()
}

// The Network Rack resource definition.
type LookupNetworkRackResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkRackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkRackResult)(nil)).Elem()
}

func (o LookupNetworkRackResultOutput) ToLookupNetworkRackResultOutput() LookupNetworkRackResultOutput {
	return o
}

func (o LookupNetworkRackResultOutput) ToLookupNetworkRackResultOutputWithContext(ctx context.Context) LookupNetworkRackResultOutput {
	return o
}

// Switch configuration description.
func (o LookupNetworkRackResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupNetworkRackResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNetworkRackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkRackResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNetworkRackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of network device ARM resource IDs.
func (o LookupNetworkRackResultOutput) NetworkDevices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) []string { return v.NetworkDevices }).(pulumi.StringArrayOutput)
}

// ARM resource ID of the Network Fabric.
func (o LookupNetworkRackResultOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// Network Rack SKU name.
func (o LookupNetworkRackResultOutput) NetworkRackType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) *string { return v.NetworkRackType }).(pulumi.StringPtrOutput)
}

// Provisioning state of the resource.
func (o LookupNetworkRackResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkRackResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkRackResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkRackResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRackResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkRackResultOutput{})
}
