// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of the provided storage appliance.
//
// Uses Azure REST API version 2025-02-01.
//
// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupStorageAppliance(ctx *pulumi.Context, args *LookupStorageApplianceArgs, opts ...pulumi.InvokeOption) (*LookupStorageApplianceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageApplianceResult
	err := ctx.Invoke("azure-native:networkcloud:getStorageAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageApplianceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the storage appliance.
	StorageApplianceName string `pulumi:"storageApplianceName"`
}

type LookupStorageApplianceResult struct {
	// The credentials of the administrative interface on this storage appliance.
	AdministratorCredentials AdministrativeCredentialsResponse `pulumi:"administratorCredentials"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The total capacity of the storage appliance. Measured in GiB.
	Capacity float64 `pulumi:"capacity"`
	// The amount of storage consumed.
	CapacityUsed float64 `pulumi:"capacityUsed"`
	// The resource ID of the cluster this storage appliance is associated with. Measured in GiB.
	ClusterId string `pulumi:"clusterId"`
	// The detailed status of the storage appliance.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// Resource ETag.
	Etag string `pulumi:"etag"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The endpoint for the management interface of the storage appliance.
	ManagementIpv4Address string `pulumi:"managementIpv4Address"`
	// The manufacturer of the storage appliance.
	Manufacturer string `pulumi:"manufacturer"`
	// The model of the storage appliance.
	Model string `pulumi:"model"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the storage appliance.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource ID of the rack where this storage appliance resides.
	RackId string `pulumi:"rackId"`
	// The slot the storage appliance is in the rack based on the BOM configuration.
	RackSlot float64 `pulumi:"rackSlot"`
	// The indicator of whether the storage appliance supports remote vendor management.
	RemoteVendorManagementFeature string `pulumi:"remoteVendorManagementFeature"`
	// The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
	RemoteVendorManagementStatus string `pulumi:"remoteVendorManagementStatus"`
	// The list of statuses that represent secret rotation activity.
	SecretRotationStatus []SecretRotationStatusResponse `pulumi:"secretRotationStatus"`
	// The serial number for the storage appliance.
	SerialNumber string `pulumi:"serialNumber"`
	// The SKU for the storage appliance.
	StorageApplianceSkuId string `pulumi:"storageApplianceSkuId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The version of the storage appliance.
	Version string `pulumi:"version"`
}

func LookupStorageApplianceOutput(ctx *pulumi.Context, args LookupStorageApplianceOutputArgs, opts ...pulumi.InvokeOption) LookupStorageApplianceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStorageApplianceResultOutput, error) {
			args := v.(LookupStorageApplianceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:networkcloud:getStorageAppliance", args, LookupStorageApplianceResultOutput{}, options).(LookupStorageApplianceResultOutput), nil
		}).(LookupStorageApplianceResultOutput)
}

type LookupStorageApplianceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the storage appliance.
	StorageApplianceName pulumi.StringInput `pulumi:"storageApplianceName"`
}

func (LookupStorageApplianceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageApplianceArgs)(nil)).Elem()
}

type LookupStorageApplianceResultOutput struct{ *pulumi.OutputState }

func (LookupStorageApplianceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageApplianceResult)(nil)).Elem()
}

func (o LookupStorageApplianceResultOutput) ToLookupStorageApplianceResultOutput() LookupStorageApplianceResultOutput {
	return o
}

func (o LookupStorageApplianceResultOutput) ToLookupStorageApplianceResultOutputWithContext(ctx context.Context) LookupStorageApplianceResultOutput {
	return o
}

// The credentials of the administrative interface on this storage appliance.
func (o LookupStorageApplianceResultOutput) AdministratorCredentials() AdministrativeCredentialsResponseOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) AdministrativeCredentialsResponse {
		return v.AdministratorCredentials
	}).(AdministrativeCredentialsResponseOutput)
}

// The Azure API version of the resource.
func (o LookupStorageApplianceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The total capacity of the storage appliance. Measured in GiB.
func (o LookupStorageApplianceResultOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStorageApplianceResult) float64 { return v.Capacity }).(pulumi.Float64Output)
}

// The amount of storage consumed.
func (o LookupStorageApplianceResultOutput) CapacityUsed() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStorageApplianceResult) float64 { return v.CapacityUsed }).(pulumi.Float64Output)
}

// The resource ID of the cluster this storage appliance is associated with. Measured in GiB.
func (o LookupStorageApplianceResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The detailed status of the storage appliance.
func (o LookupStorageApplianceResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupStorageApplianceResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// Resource ETag.
func (o LookupStorageApplianceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupStorageApplianceResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupStorageApplianceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupStorageApplianceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The endpoint for the management interface of the storage appliance.
func (o LookupStorageApplianceResultOutput) ManagementIpv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.ManagementIpv4Address }).(pulumi.StringOutput)
}

// The manufacturer of the storage appliance.
func (o LookupStorageApplianceResultOutput) Manufacturer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Manufacturer }).(pulumi.StringOutput)
}

// The model of the storage appliance.
func (o LookupStorageApplianceResultOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Model }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStorageApplianceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the storage appliance.
func (o LookupStorageApplianceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource ID of the rack where this storage appliance resides.
func (o LookupStorageApplianceResultOutput) RackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.RackId }).(pulumi.StringOutput)
}

// The slot the storage appliance is in the rack based on the BOM configuration.
func (o LookupStorageApplianceResultOutput) RackSlot() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStorageApplianceResult) float64 { return v.RackSlot }).(pulumi.Float64Output)
}

// The indicator of whether the storage appliance supports remote vendor management.
func (o LookupStorageApplianceResultOutput) RemoteVendorManagementFeature() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.RemoteVendorManagementFeature }).(pulumi.StringOutput)
}

// The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
func (o LookupStorageApplianceResultOutput) RemoteVendorManagementStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.RemoteVendorManagementStatus }).(pulumi.StringOutput)
}

// The list of statuses that represent secret rotation activity.
func (o LookupStorageApplianceResultOutput) SecretRotationStatus() SecretRotationStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) []SecretRotationStatusResponse { return v.SecretRotationStatus }).(SecretRotationStatusResponseArrayOutput)
}

// The serial number for the storage appliance.
func (o LookupStorageApplianceResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// The SKU for the storage appliance.
func (o LookupStorageApplianceResultOutput) StorageApplianceSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.StorageApplianceSkuId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStorageApplianceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupStorageApplianceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStorageApplianceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The version of the storage appliance.
func (o LookupStorageApplianceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageApplianceResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageApplianceResultOutput{})
}
