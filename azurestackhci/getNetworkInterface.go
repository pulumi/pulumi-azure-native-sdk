// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network interface
//
// Uses Azure REST API version 2025-02-01-preview.
//
// Other available API versions: 2022-12-15-preview, 2023-07-01-preview, 2023-09-01-preview, 2024-01-01, 2024-02-01-preview, 2024-05-01-preview, 2024-07-15-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:azurestackhci:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkInterfaceArgs struct {
	// Name of the network interface
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The network interface resource definition.
type LookupNetworkInterfaceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Boolean indicating whether this is a existing local network interface or if one should be created.
	CreateFromLocal *bool `pulumi:"createFromLocal"`
	// DNS Settings for the interface
	DnsSettings *InterfaceDNSSettingsResponse `pulumi:"dnsSettings"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations []IPConfigurationResponse `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// MacAddress - The MAC address of the network interface.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the resource
	Name string `pulumi:"name"`
	// NetworkSecurityGroup - Network Security Group attached to the network interface.
	NetworkSecurityGroup *NetworkSecurityGroupArmReferenceResponse `pulumi:"networkSecurityGroup"`
	// Provisioning state of the network interface.
	ProvisioningState string `pulumi:"provisioningState"`
	// The observed state of network interfaces
	Status NetworkInterfaceStatusResponse `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupNetworkInterfaceResult
func (val *LookupNetworkInterfaceResult) Defaults() *LookupNetworkInterfaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.CreateFromLocal == nil {
		createFromLocal_ := false
		tmp.CreateFromLocal = &createFromLocal_
	}
	return &tmp
}
func LookupNetworkInterfaceOutput(ctx *pulumi.Context, args LookupNetworkInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceResultOutput, error) {
			args := v.(LookupNetworkInterfaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azurestackhci:getNetworkInterface", args, LookupNetworkInterfaceResultOutput{}, options).(LookupNetworkInterfaceResultOutput), nil
		}).(LookupNetworkInterfaceResultOutput)
}

type LookupNetworkInterfaceOutputArgs struct {
	// Name of the network interface
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceArgs)(nil)).Elem()
}

// The network interface resource definition.
type LookupNetworkInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutput() LookupNetworkInterfaceResultOutput {
	return o
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupNetworkInterfaceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Boolean indicating whether this is a existing local network interface or if one should be created.
func (o LookupNetworkInterfaceResultOutput) CreateFromLocal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.CreateFromLocal }).(pulumi.BoolPtrOutput)
}

// DNS Settings for the interface
func (o LookupNetworkInterfaceResultOutput) DnsSettings() InterfaceDNSSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *InterfaceDNSSettingsResponse { return v.DnsSettings }).(InterfaceDNSSettingsResponsePtrOutput)
}

// The extendedLocation of the resource.
func (o LookupNetworkInterfaceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNetworkInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// IPConfigurations - A list of IPConfigurations of the network interface.
func (o LookupNetworkInterfaceResultOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// MacAddress - The MAC address of the network interface.
func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// NetworkSecurityGroup - Network Security Group attached to the network interface.
func (o LookupNetworkInterfaceResultOutput) NetworkSecurityGroup() NetworkSecurityGroupArmReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkSecurityGroupArmReferenceResponse {
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupArmReferenceResponsePtrOutput)
}

// Provisioning state of the network interface.
func (o LookupNetworkInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of network interfaces
func (o LookupNetworkInterfaceResultOutput) Status() NetworkInterfaceStatusResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) NetworkInterfaceStatusResponse { return v.Status }).(NetworkInterfaceStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceResultOutput{})
}
