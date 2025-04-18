// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get specified Update
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2022-12-15-preview, 2023-02-01, 2023-03-01, 2023-06-01, 2023-08-01, 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupUpdate(ctx *pulumi.Context, args *LookupUpdateArgs, opts ...pulumi.InvokeOption) (*LookupUpdateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUpdateResult
	err := ctx.Invoke("azure-native:azurestackhci:getUpdate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUpdateArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Update
	UpdateName string `pulumi:"updateName"`
}

// Update details
type LookupUpdateResult struct {
	// Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
	AdditionalProperties *string `pulumi:"additionalProperties"`
	// Indicates the way the update content can be downloaded.
	AvailabilityType *string `pulumi:"availabilityType"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Description of the update.
	Description *string `pulumi:"description"`
	// Display name of the Update
	DisplayName *string `pulumi:"displayName"`
	// Last time the package-specific checks were run.
	HealthCheckDate *string `pulumi:"healthCheckDate"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Date that the update was installed.
	InstalledDate *string `pulumi:"installedDate"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Minimum Sbe Version of the update.
	MinSbeVersionRequired *string `pulumi:"minSbeVersionRequired"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Brief message with instructions for updates of AvailabilityType Notify.
	NotifyMessage *string `pulumi:"notifyMessage"`
	// Path where the update package is available.
	PackagePath *string `pulumi:"packagePath"`
	// Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
	PackageSizeInMb *float64 `pulumi:"packageSizeInMb"`
	// Customer-visible type of the update.
	PackageType *string `pulumi:"packageType"`
	// If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
	Prerequisites []UpdatePrerequisiteResponse `pulumi:"prerequisites"`
	// Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
	ProgressPercentage *float64 `pulumi:"progressPercentage"`
	// Provisioning state of the Updates proxy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Publisher of the update package.
	Publisher *string `pulumi:"publisher"`
	// Link to release notes for the update.
	ReleaseLink *string `pulumi:"releaseLink"`
	// State of the update as it relates to this stamp.
	State *string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Version of the update.
	Version *string `pulumi:"version"`
}

func LookupUpdateOutput(ctx *pulumi.Context, args LookupUpdateOutputArgs, opts ...pulumi.InvokeOption) LookupUpdateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupUpdateResultOutput, error) {
			args := v.(LookupUpdateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azurestackhci:getUpdate", args, LookupUpdateResultOutput{}, options).(LookupUpdateResultOutput), nil
		}).(LookupUpdateResultOutput)
}

type LookupUpdateOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Update
	UpdateName pulumi.StringInput `pulumi:"updateName"`
}

func (LookupUpdateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateArgs)(nil)).Elem()
}

// Update details
type LookupUpdateResultOutput struct{ *pulumi.OutputState }

func (LookupUpdateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateResult)(nil)).Elem()
}

func (o LookupUpdateResultOutput) ToLookupUpdateResultOutput() LookupUpdateResultOutput {
	return o
}

func (o LookupUpdateResultOutput) ToLookupUpdateResultOutputWithContext(ctx context.Context) LookupUpdateResultOutput {
	return o
}

// Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
func (o LookupUpdateResultOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

// Indicates the way the update content can be downloaded.
func (o LookupUpdateResultOutput) AvailabilityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.AvailabilityType }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupUpdateResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Description of the update.
func (o LookupUpdateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of the Update
func (o LookupUpdateResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Last time the package-specific checks were run.
func (o LookupUpdateResultOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupUpdateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Date that the update was installed.
func (o LookupUpdateResultOutput) InstalledDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.InstalledDate }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupUpdateResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Minimum Sbe Version of the update.
func (o LookupUpdateResultOutput) MinSbeVersionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.MinSbeVersionRequired }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupUpdateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Brief message with instructions for updates of AvailabilityType Notify.
func (o LookupUpdateResultOutput) NotifyMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.NotifyMessage }).(pulumi.StringPtrOutput)
}

// Path where the update package is available.
func (o LookupUpdateResultOutput) PackagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.PackagePath }).(pulumi.StringPtrOutput)
}

// Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
func (o LookupUpdateResultOutput) PackageSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *float64 { return v.PackageSizeInMb }).(pulumi.Float64PtrOutput)
}

// Customer-visible type of the update.
func (o LookupUpdateResultOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.PackageType }).(pulumi.StringPtrOutput)
}

// If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
func (o LookupUpdateResultOutput) Prerequisites() UpdatePrerequisiteResponseArrayOutput {
	return o.ApplyT(func(v LookupUpdateResult) []UpdatePrerequisiteResponse { return v.Prerequisites }).(UpdatePrerequisiteResponseArrayOutput)
}

// Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
func (o LookupUpdateResultOutput) ProgressPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *float64 { return v.ProgressPercentage }).(pulumi.Float64PtrOutput)
}

// Provisioning state of the Updates proxy resource.
func (o LookupUpdateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Publisher of the update package.
func (o LookupUpdateResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Link to release notes for the update.
func (o LookupUpdateResultOutput) ReleaseLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.ReleaseLink }).(pulumi.StringPtrOutput)
}

// State of the update as it relates to this stamp.
func (o LookupUpdateResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupUpdateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUpdateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupUpdateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of the update.
func (o LookupUpdateResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUpdateResultOutput{})
}
