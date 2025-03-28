// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements ExternalNetworks GET method.
//
// Uses Azure REST API version 2023-02-01-preview.
//
// Other available API versions: 2023-06-15.
func LookupExternalNetwork(ctx *pulumi.Context, args *LookupExternalNetworkArgs, opts ...pulumi.InvokeOption) (*LookupExternalNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExternalNetworkResult
	err := ctx.Invoke("azure-native:managednetworkfabric:getExternalNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupExternalNetworkArgs struct {
	// Name of the ExternalNetwork
	ExternalNetworkName string `pulumi:"externalNetworkName"`
	// Name of the L3IsolationDomain
	L3IsolationDomainName string `pulumi:"l3IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the ExternalNetwork item.
type LookupExternalNetworkResult struct {
	// AdministrativeState of the externalNetwork. Example: Enabled | Disabled.
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// List of resources the externalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
	DisabledOnResources []string `pulumi:"disabledOnResources"`
	// ARM resource ID of exportRoutePolicy.
	ExportRoutePolicyId *string `pulumi:"exportRoutePolicyId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// ARM resource ID of importRoutePolicy.
	ImportRoutePolicyId *string `pulumi:"importRoutePolicyId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the networkToNetworkInterconnectId of the resource.
	NetworkToNetworkInterconnectId string `pulumi:"networkToNetworkInterconnectId"`
	// option A properties object
	OptionAProperties *ExternalNetworkPropertiesResponseOptionAProperties `pulumi:"optionAProperties"`
	// option B properties object
	OptionBProperties *OptionBPropertiesResponse `pulumi:"optionBProperties"`
	// Peering option list.
	PeeringOption string `pulumi:"peeringOption"`
	// Gets the provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupExternalNetworkResult
func (val *LookupExternalNetworkResult) Defaults() *LookupExternalNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.OptionAProperties = tmp.OptionAProperties.Defaults()

	return &tmp
}
func LookupExternalNetworkOutput(ctx *pulumi.Context, args LookupExternalNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupExternalNetworkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExternalNetworkResultOutput, error) {
			args := v.(LookupExternalNetworkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:managednetworkfabric:getExternalNetwork", args, LookupExternalNetworkResultOutput{}, options).(LookupExternalNetworkResultOutput), nil
		}).(LookupExternalNetworkResultOutput)
}

type LookupExternalNetworkOutputArgs struct {
	// Name of the ExternalNetwork
	ExternalNetworkName pulumi.StringInput `pulumi:"externalNetworkName"`
	// Name of the L3IsolationDomain
	L3IsolationDomainName pulumi.StringInput `pulumi:"l3IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExternalNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalNetworkArgs)(nil)).Elem()
}

// Defines the ExternalNetwork item.
type LookupExternalNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupExternalNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalNetworkResult)(nil)).Elem()
}

func (o LookupExternalNetworkResultOutput) ToLookupExternalNetworkResultOutput() LookupExternalNetworkResultOutput {
	return o
}

func (o LookupExternalNetworkResultOutput) ToLookupExternalNetworkResultOutputWithContext(ctx context.Context) LookupExternalNetworkResultOutput {
	return o
}

// AdministrativeState of the externalNetwork. Example: Enabled | Disabled.
func (o LookupExternalNetworkResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupExternalNetworkResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// List of resources the externalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o LookupExternalNetworkResultOutput) DisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) []string { return v.DisabledOnResources }).(pulumi.StringArrayOutput)
}

// ARM resource ID of exportRoutePolicy.
func (o LookupExternalNetworkResultOutput) ExportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) *string { return v.ExportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupExternalNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// ARM resource ID of importRoutePolicy.
func (o LookupExternalNetworkResultOutput) ImportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) *string { return v.ImportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupExternalNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the networkToNetworkInterconnectId of the resource.
func (o LookupExternalNetworkResultOutput) NetworkToNetworkInterconnectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.NetworkToNetworkInterconnectId }).(pulumi.StringOutput)
}

// option A properties object
func (o LookupExternalNetworkResultOutput) OptionAProperties() ExternalNetworkPropertiesResponseOptionAPropertiesPtrOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) *ExternalNetworkPropertiesResponseOptionAProperties {
		return v.OptionAProperties
	}).(ExternalNetworkPropertiesResponseOptionAPropertiesPtrOutput)
}

// option B properties object
func (o LookupExternalNetworkResultOutput) OptionBProperties() OptionBPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) *OptionBPropertiesResponse { return v.OptionBProperties }).(OptionBPropertiesResponsePtrOutput)
}

// Peering option list.
func (o LookupExternalNetworkResultOutput) PeeringOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.PeeringOption }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o LookupExternalNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupExternalNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupExternalNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExternalNetworkResultOutput{})
}
