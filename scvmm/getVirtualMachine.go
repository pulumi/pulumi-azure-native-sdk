// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements VirtualMachine GET method.
//
// Uses Azure REST API version 2023-04-01-preview.
//
// Other available API versions: 2022-05-21-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:scvmm:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the VirtualMachine.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The VirtualMachines resource definition.
type LookupVirtualMachineResult struct {
	// Availability Sets in vm.
	AvailabilitySets []VirtualMachinePropertiesResponseAvailabilitySets `pulumi:"availabilitySets"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Type of checkpoint supported for the vm.
	CheckpointType *string `pulumi:"checkpointType"`
	// Checkpoints in the vm.
	Checkpoints []CheckpointResponse `pulumi:"checkpoints"`
	// ARM Id of the cloud resource to use for deploying the vm.
	CloudId *string `pulumi:"cloudId"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Gets or sets the generation for the vm.
	Generation *int `pulumi:"generation"`
	// Guest agent status properties.
	GuestAgentProfile *GuestAgentProfileResponse `pulumi:"guestAgentProfile"`
	// Hardware properties.
	HardwareProfile *HardwareProfileResponse `pulumi:"hardwareProfile"`
	// Resource Id
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Last restored checkpoint in the vm.
	LastRestoredVMCheckpoint CheckpointResponse `pulumi:"lastRestoredVMCheckpoint"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Network properties.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// OS properties.
	OsProfile *OsProfileResponse `pulumi:"osProfile"`
	// Gets the power state of the virtual machine.
	PowerState string `pulumi:"powerState"`
	// Gets or sets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Storage properties.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// ARM Id of the template resource to use for deploying the vm.
	TemplateId *string `pulumi:"templateId"`
	// Resource Type
	Type string `pulumi:"type"`
	// Unique ID of the virtual machine.
	Uuid *string `pulumi:"uuid"`
	// VMName is the name of VM on the SCVMM server.
	VmName *string `pulumi:"vmName"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

func LookupVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResultOutput, error) {
			args := v.(LookupVirtualMachineArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:scvmm:getVirtualMachine", args, LookupVirtualMachineResultOutput{}, options).(LookupVirtualMachineResultOutput), nil
		}).(LookupVirtualMachineResultOutput)
}

type LookupVirtualMachineOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the VirtualMachine.
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}

// The VirtualMachines resource definition.
type LookupVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutput() LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineResultOutput {
	return o
}

// Availability Sets in vm.
func (o LookupVirtualMachineResultOutput) AvailabilitySets() VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualMachinePropertiesResponseAvailabilitySets {
		return v.AvailabilitySets
	}).(VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput)
}

// The Azure API version of the resource.
func (o LookupVirtualMachineResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Type of checkpoint supported for the vm.
func (o LookupVirtualMachineResultOutput) CheckpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CheckpointType }).(pulumi.StringPtrOutput)
}

// Checkpoints in the vm.
func (o LookupVirtualMachineResultOutput) Checkpoints() CheckpointResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []CheckpointResponse { return v.Checkpoints }).(CheckpointResponseArrayOutput)
}

// ARM Id of the cloud resource to use for deploying the vm.
func (o LookupVirtualMachineResultOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CloudId }).(pulumi.StringPtrOutput)
}

// The extended location.
func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Gets or sets the generation for the vm.
func (o LookupVirtualMachineResultOutput) Generation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.Generation }).(pulumi.IntPtrOutput)
}

// Guest agent status properties.
func (o LookupVirtualMachineResultOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *GuestAgentProfileResponse { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

// Hardware properties.
func (o LookupVirtualMachineResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// Resource Id
func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupVirtualMachineResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Gets or sets the inventory Item ID for the resource.
func (o LookupVirtualMachineResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Last restored checkpoint in the vm.
func (o LookupVirtualMachineResultOutput) LastRestoredVMCheckpoint() CheckpointResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) CheckpointResponse { return v.LastRestoredVMCheckpoint }).(CheckpointResponseOutput)
}

// Gets or sets the location.
func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network properties.
func (o LookupVirtualMachineResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// OS properties.
func (o LookupVirtualMachineResultOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *OsProfileResponse { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

// Gets the power state of the virtual machine.
func (o LookupVirtualMachineResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.PowerState }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Storage properties.
func (o LookupVirtualMachineResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// The system data.
func (o LookupVirtualMachineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// ARM Id of the template resource to use for deploying the vm.
func (o LookupVirtualMachineResultOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.TemplateId }).(pulumi.StringPtrOutput)
}

// Resource Type
func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of the virtual machine.
func (o LookupVirtualMachineResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

// VMName is the name of VM on the SCVMM server.
func (o LookupVirtualMachineResultOutput) VmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VmName }).(pulumi.StringPtrOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o LookupVirtualMachineResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
