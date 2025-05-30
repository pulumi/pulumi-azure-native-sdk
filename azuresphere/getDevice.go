// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuresphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Device. Use '.unassigned' or '.default' for the device group and product names when a device does not belong to a device group and product.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2022-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuresphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDeviceResult
	err := ctx.Invoke("azure-native:azuresphere:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceArgs struct {
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Name of device group.
	DeviceGroupName string `pulumi:"deviceGroupName"`
	// Device name
	DeviceName string `pulumi:"deviceName"`
	// Name of product.
	ProductName string `pulumi:"productName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An device resource belonging to a device group resource.
type LookupDeviceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// SKU of the chip
	ChipSku string `pulumi:"chipSku"`
	// Device ID
	DeviceId *string `pulumi:"deviceId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// OS version available for installation when update requested
	LastAvailableOsVersion string `pulumi:"lastAvailableOsVersion"`
	// OS version running on device when update requested
	LastInstalledOsVersion string `pulumi:"lastInstalledOsVersion"`
	// Time when update requested and new OS version available
	LastOsUpdateUtc string `pulumi:"lastOsUpdateUtc"`
	// Time when update was last requested
	LastUpdateRequestUtc string `pulumi:"lastUpdateRequestUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDeviceOutput(ctx *pulumi.Context, args LookupDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDeviceResultOutput, error) {
			args := v.(LookupDeviceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azuresphere:getDevice", args, LookupDeviceResultOutput{}, options).(LookupDeviceResultOutput), nil
		}).(LookupDeviceResultOutput)
}

type LookupDeviceOutputArgs struct {
	// Name of catalog
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// Name of device group.
	DeviceGroupName pulumi.StringInput `pulumi:"deviceGroupName"`
	// Device name
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// Name of product.
	ProductName pulumi.StringInput `pulumi:"productName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceArgs)(nil)).Elem()
}

// An device resource belonging to a device group resource.
type LookupDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceResult)(nil)).Elem()
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutput() LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutputWithContext(ctx context.Context) LookupDeviceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDeviceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// SKU of the chip
func (o LookupDeviceResultOutput) ChipSku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ChipSku }).(pulumi.StringOutput)
}

// Device ID
func (o LookupDeviceResultOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.DeviceId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

// OS version available for installation when update requested
func (o LookupDeviceResultOutput) LastAvailableOsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.LastAvailableOsVersion }).(pulumi.StringOutput)
}

// OS version running on device when update requested
func (o LookupDeviceResultOutput) LastInstalledOsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.LastInstalledOsVersion }).(pulumi.StringOutput)
}

// Time when update requested and new OS version available
func (o LookupDeviceResultOutput) LastOsUpdateUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.LastOsUpdateUtc }).(pulumi.StringOutput)
}

// Time when update was last requested
func (o LookupDeviceResultOutput) LastUpdateRequestUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.LastUpdateRequestUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupDeviceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDeviceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeviceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDeviceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceResultOutput{})
}
