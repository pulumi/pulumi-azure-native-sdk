// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements virtual machine template GET method.
//
// Uses Azure REST API version 2023-12-01.
//
// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupVirtualMachineTemplate(ctx *pulumi.Context, args *LookupVirtualMachineTemplateArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineTemplateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineTemplateResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getVirtualMachineTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineTemplateArgs struct {
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the virtual machine template resource.
	VirtualMachineTemplateName string `pulumi:"virtualMachineTemplateName"`
}

// Define the virtualMachineTemplate.
type LookupVirtualMachineTemplateResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// Gets or sets the disks the template.
	Disks []VirtualDiskResponse `pulumi:"disks"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Firmware type
	FirmwareType string `pulumi:"firmwareType"`
	// Gets or sets the folder path of the template.
	FolderPath string `pulumi:"folderPath"`
	// Gets or sets the Id.
	Id string `pulumi:"id"`
	// Gets or sets the inventory Item ID for the virtual machine template.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Gets or sets memory size in MBs for the template.
	MemorySizeMB int `pulumi:"memorySizeMB"`
	// Gets or sets the vCenter Managed Object name for the virtual machine template.
	MoName string `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual machine
	// template.
	MoRefId *string `pulumi:"moRefId"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Gets or sets the network interfaces of the template.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// Gets or sets the number of vCPUs for the template.
	NumCPUs int `pulumi:"numCPUs"`
	// Gets or sets the number of cores per socket for the template.
	// Defaults to 1 if unspecified.
	NumCoresPerSocket int `pulumi:"numCoresPerSocket"`
	// Gets or sets os name.
	OsName string `pulumi:"osName"`
	// Gets or sets the type of the os.
	OsType string `pulumi:"osType"`
	// Gets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource status information.
	Statuses []ResourceStatusResponse `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the current version of VMware Tools.
	ToolsVersion string `pulumi:"toolsVersion"`
	// Gets or sets the current version status of VMware Tools installed in the guest operating system.
	ToolsVersionStatus string `pulumi:"toolsVersionStatus"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid string `pulumi:"uuid"`
	// Gets or sets the ARM Id of the vCenter resource in which this template resides.
	VCenterId *string `pulumi:"vCenterId"`
}

func LookupVirtualMachineTemplateOutput(ctx *pulumi.Context, args LookupVirtualMachineTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineTemplateResultOutput, error) {
			args := v.(LookupVirtualMachineTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:connectedvmwarevsphere:getVirtualMachineTemplate", args, LookupVirtualMachineTemplateResultOutput{}, options).(LookupVirtualMachineTemplateResultOutput), nil
		}).(LookupVirtualMachineTemplateResultOutput)
}

type LookupVirtualMachineTemplateOutputArgs struct {
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the virtual machine template resource.
	VirtualMachineTemplateName pulumi.StringInput `pulumi:"virtualMachineTemplateName"`
}

func (LookupVirtualMachineTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineTemplateArgs)(nil)).Elem()
}

// Define the virtualMachineTemplate.
type LookupVirtualMachineTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineTemplateResult)(nil)).Elem()
}

func (o LookupVirtualMachineTemplateResultOutput) ToLookupVirtualMachineTemplateResultOutput() LookupVirtualMachineTemplateResultOutput {
	return o
}

func (o LookupVirtualMachineTemplateResultOutput) ToLookupVirtualMachineTemplateResultOutputWithContext(ctx context.Context) LookupVirtualMachineTemplateResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupVirtualMachineTemplateResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupVirtualMachineTemplateResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the disks the template.
func (o LookupVirtualMachineTemplateResultOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []VirtualDiskResponse { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

// Gets or sets the extended location.
func (o LookupVirtualMachineTemplateResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Firmware type
func (o LookupVirtualMachineTemplateResultOutput) FirmwareType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.FirmwareType }).(pulumi.StringOutput)
}

// Gets or sets the folder path of the template.
func (o LookupVirtualMachineTemplateResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

// Gets or sets the Id.
func (o LookupVirtualMachineTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the inventory Item ID for the virtual machine template.
func (o LookupVirtualMachineTemplateResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupVirtualMachineTemplateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupVirtualMachineTemplateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets memory size in MBs for the template.
func (o LookupVirtualMachineTemplateResultOutput) MemorySizeMB() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.MemorySizeMB }).(pulumi.IntOutput)
}

// Gets or sets the vCenter Managed Object name for the virtual machine template.
func (o LookupVirtualMachineTemplateResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual machine
// template.
func (o LookupVirtualMachineTemplateResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o LookupVirtualMachineTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the network interfaces of the template.
func (o LookupVirtualMachineTemplateResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// Gets or sets the number of vCPUs for the template.
func (o LookupVirtualMachineTemplateResultOutput) NumCPUs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.NumCPUs }).(pulumi.IntOutput)
}

// Gets or sets the number of cores per socket for the template.
// Defaults to 1 if unspecified.
func (o LookupVirtualMachineTemplateResultOutput) NumCoresPerSocket() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) int { return v.NumCoresPerSocket }).(pulumi.IntOutput)
}

// Gets or sets os name.
func (o LookupVirtualMachineTemplateResultOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.OsName }).(pulumi.StringOutput)
}

// Gets or sets the type of the os.
func (o LookupVirtualMachineTemplateResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.OsType }).(pulumi.StringOutput)
}

// Gets the provisioning state.
func (o LookupVirtualMachineTemplateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupVirtualMachineTemplateResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o LookupVirtualMachineTemplateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o LookupVirtualMachineTemplateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the current version of VMware Tools.
func (o LookupVirtualMachineTemplateResultOutput) ToolsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ToolsVersion }).(pulumi.StringOutput)
}

// Gets or sets the current version status of VMware Tools installed in the guest operating system.
func (o LookupVirtualMachineTemplateResultOutput) ToolsVersionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.ToolsVersionStatus }).(pulumi.StringOutput)
}

// Gets or sets the type of the resource.
func (o LookupVirtualMachineTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupVirtualMachineTemplateResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this template resides.
func (o LookupVirtualMachineTemplateResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineTemplateResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineTemplateResultOutput{})
}
