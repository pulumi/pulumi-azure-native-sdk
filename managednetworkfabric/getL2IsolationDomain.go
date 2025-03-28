// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements L2 Isolation Domain GET method.
//
// Uses Azure REST API version 2023-02-01-preview.
//
// Other available API versions: 2023-06-15.
func LookupL2IsolationDomain(ctx *pulumi.Context, args *LookupL2IsolationDomainArgs, opts ...pulumi.InvokeOption) (*LookupL2IsolationDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupL2IsolationDomainResult
	err := ctx.Invoke("azure-native:managednetworkfabric:getL2IsolationDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupL2IsolationDomainArgs struct {
	// Name of the L2 Isolation Domain
	L2IsolationDomainName string `pulumi:"l2IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The L2IsolationDomain resource definition.
type LookupL2IsolationDomainResult struct {
	// state. Example: Enabled | Disabled. It indicates administrative state of the isolationDomain, whether it is enabled or disabled. If enabled, the configuration is applied on the devices. If disabled, the configuration is removed from the devices
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// List of resources the L2 Isolation Domain is disabled on. Can be either entire NetworkFabric or NetworkRack.
	DisabledOnResources []string `pulumi:"disabledOnResources"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// maximum transmission unit. Default value is 1500.
	Mtu *int `pulumi:"mtu"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Network Fabric ARM resource id.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// Gets the provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// vlanId. Example: 501.
	VlanId int `pulumi:"vlanId"`
}

func LookupL2IsolationDomainOutput(ctx *pulumi.Context, args LookupL2IsolationDomainOutputArgs, opts ...pulumi.InvokeOption) LookupL2IsolationDomainResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupL2IsolationDomainResultOutput, error) {
			args := v.(LookupL2IsolationDomainArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:managednetworkfabric:getL2IsolationDomain", args, LookupL2IsolationDomainResultOutput{}, options).(LookupL2IsolationDomainResultOutput), nil
		}).(LookupL2IsolationDomainResultOutput)
}

type LookupL2IsolationDomainOutputArgs struct {
	// Name of the L2 Isolation Domain
	L2IsolationDomainName pulumi.StringInput `pulumi:"l2IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupL2IsolationDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2IsolationDomainArgs)(nil)).Elem()
}

// The L2IsolationDomain resource definition.
type LookupL2IsolationDomainResultOutput struct{ *pulumi.OutputState }

func (LookupL2IsolationDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2IsolationDomainResult)(nil)).Elem()
}

func (o LookupL2IsolationDomainResultOutput) ToLookupL2IsolationDomainResultOutput() LookupL2IsolationDomainResultOutput {
	return o
}

func (o LookupL2IsolationDomainResultOutput) ToLookupL2IsolationDomainResultOutputWithContext(ctx context.Context) LookupL2IsolationDomainResultOutput {
	return o
}

// state. Example: Enabled | Disabled. It indicates administrative state of the isolationDomain, whether it is enabled or disabled. If enabled, the configuration is applied on the devices. If disabled, the configuration is removed from the devices
func (o LookupL2IsolationDomainResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupL2IsolationDomainResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// List of resources the L2 Isolation Domain is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o LookupL2IsolationDomainResultOutput) DisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) []string { return v.DisabledOnResources }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupL2IsolationDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupL2IsolationDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

// maximum transmission unit. Default value is 1500.
func (o LookupL2IsolationDomainResultOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) *int { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupL2IsolationDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network Fabric ARM resource id.
func (o LookupL2IsolationDomainResultOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o LookupL2IsolationDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupL2IsolationDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupL2IsolationDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupL2IsolationDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

// vlanId. Example: 501.
func (o LookupL2IsolationDomainResultOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) int { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupL2IsolationDomainResultOutput{})
}
