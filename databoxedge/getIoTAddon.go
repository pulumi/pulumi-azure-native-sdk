// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a specific addon by name.
//
// Uses Azure REST API version 2023-07-01.
func LookupIoTAddon(ctx *pulumi.Context, args *LookupIoTAddonArgs, opts ...pulumi.InvokeOption) (*LookupIoTAddonResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIoTAddonResult
	err := ctx.Invoke("azure-native:databoxedge:getIoTAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIoTAddonArgs struct {
	// The addon name.
	AddonName string `pulumi:"addonName"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The role name.
	RoleName string `pulumi:"roleName"`
}

// IoT Addon.
type LookupIoTAddonResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Host OS supported by the IoT addon.
	HostPlatform string `pulumi:"hostPlatform"`
	// Platform where the runtime is hosted.
	HostPlatformType string `pulumi:"hostPlatformType"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// IoT device metadata to which appliance needs to be connected.
	IoTDeviceDetails IoTDeviceInfoResponse `pulumi:"ioTDeviceDetails"`
	// IoT edge device to which the IoT Addon needs to be configured.
	IoTEdgeDeviceDetails IoTDeviceInfoResponse `pulumi:"ioTEdgeDeviceDetails"`
	// Addon type.
	// Expected value is 'IotEdge'.
	Kind string `pulumi:"kind"`
	// The object name.
	Name string `pulumi:"name"`
	// Addon Provisioning State
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of Addon
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// Version of IoT running on the appliance.
	Version string `pulumi:"version"`
}

func LookupIoTAddonOutput(ctx *pulumi.Context, args LookupIoTAddonOutputArgs, opts ...pulumi.InvokeOption) LookupIoTAddonResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIoTAddonResultOutput, error) {
			args := v.(LookupIoTAddonArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:databoxedge:getIoTAddon", args, LookupIoTAddonResultOutput{}, options).(LookupIoTAddonResultOutput), nil
		}).(LookupIoTAddonResultOutput)
}

type LookupIoTAddonOutputArgs struct {
	// The addon name.
	AddonName pulumi.StringInput `pulumi:"addonName"`
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The role name.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (LookupIoTAddonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTAddonArgs)(nil)).Elem()
}

// IoT Addon.
type LookupIoTAddonResultOutput struct{ *pulumi.OutputState }

func (LookupIoTAddonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTAddonResult)(nil)).Elem()
}

func (o LookupIoTAddonResultOutput) ToLookupIoTAddonResultOutput() LookupIoTAddonResultOutput {
	return o
}

func (o LookupIoTAddonResultOutput) ToLookupIoTAddonResultOutputWithContext(ctx context.Context) LookupIoTAddonResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupIoTAddonResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Host OS supported by the IoT addon.
func (o LookupIoTAddonResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

// Platform where the runtime is hosted.
func (o LookupIoTAddonResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupIoTAddonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Id }).(pulumi.StringOutput)
}

// IoT device metadata to which appliance needs to be connected.
func (o LookupIoTAddonResultOutput) IoTDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) IoTDeviceInfoResponse { return v.IoTDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

// IoT edge device to which the IoT Addon needs to be configured.
func (o LookupIoTAddonResultOutput) IoTEdgeDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) IoTDeviceInfoResponse { return v.IoTEdgeDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

// Addon type.
// Expected value is 'IotEdge'.
func (o LookupIoTAddonResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o LookupIoTAddonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Name }).(pulumi.StringOutput)
}

// Addon Provisioning State
func (o LookupIoTAddonResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Addon
func (o LookupIoTAddonResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupIoTAddonResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of IoT running on the appliance.
func (o LookupIoTAddonResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTAddonResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoTAddonResultOutput{})
}
