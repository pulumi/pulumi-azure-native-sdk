// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an image source that is an installation ISO. Currently only supports Red Hat Enterprise Linux 7.2-7.5 ISO's.
type ImageTemplateIsoSource struct {
	// SHA256 Checksum of the ISO image.
	Sha256Checksum string `pulumi:"sha256Checksum"`
	// URI to get the ISO image. This URI has to be accessible to the resource provider at the time of the image template creation.
	SourceURI string `pulumi:"sourceURI"`
	// Specifies the type of source image you want to start with.
	// Expected value is 'ISO'.
	Type string `pulumi:"type"`
}

// Describes an image source that is an installation ISO. Currently only supports Red Hat Enterprise Linux 7.2-7.5 ISO's.
type ImageTemplateIsoSourceResponse struct {
	// SHA256 Checksum of the ISO image.
	Sha256Checksum string `pulumi:"sha256Checksum"`
	// URI to get the ISO image. This URI has to be accessible to the resource provider at the time of the image template creation.
	SourceURI string `pulumi:"sourceURI"`
	// Specifies the type of source image you want to start with.
	// Expected value is 'ISO'.
	Type string `pulumi:"type"`
}

type ImageTemplateLastRunStatusResponse struct {
	// End time of the last run (UTC)
	EndTime *string `pulumi:"endTime"`
	// Verbose information about the last run state
	Message *string `pulumi:"message"`
	// State of the last run
	RunState *string `pulumi:"runState"`
	// Sub-state of the last run
	RunSubState *string `pulumi:"runSubState"`
	// Start time of the last run (UTC)
	StartTime *string `pulumi:"startTime"`
}

type ImageTemplateLastRunStatusResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateLastRunStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponseOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponseOutput {
	return o
}

// End time of the last run (UTC)
func (o ImageTemplateLastRunStatusResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Verbose information about the last run state
func (o ImageTemplateLastRunStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// State of the last run
func (o ImageTemplateLastRunStatusResponseOutput) RunState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.RunState }).(pulumi.StringPtrOutput)
}

// Sub-state of the last run
func (o ImageTemplateLastRunStatusResponseOutput) RunSubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.RunSubState }).(pulumi.StringPtrOutput)
}

// Start time of the last run (UTC)
func (o ImageTemplateLastRunStatusResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Distribute as a Managed Disk Image.
type ImageTemplateManagedImageDistributor struct {
	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]string `pulumi:"artifactTags"`
	// Resource Id of the Managed Disk Image
	ImageId string `pulumi:"imageId"`
	// Azure location for the image, should match if image already exists
	Location string `pulumi:"location"`
	// The name to be used for the associated RunOutput.
	RunOutputName string `pulumi:"runOutputName"`
	// Type of distribution.
	// Expected value is 'ManagedImage'.
	Type string `pulumi:"type"`
}

// Distribute as a Managed Disk Image.
type ImageTemplateManagedImageDistributorResponse struct {
	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]string `pulumi:"artifactTags"`
	// Resource Id of the Managed Disk Image
	ImageId string `pulumi:"imageId"`
	// Azure location for the image, should match if image already exists
	Location string `pulumi:"location"`
	// The name to be used for the associated RunOutput.
	RunOutputName string `pulumi:"runOutputName"`
	// Type of distribution.
	// Expected value is 'ManagedImage'.
	Type string `pulumi:"type"`
}

// Describes an image source that is a managed image in customer subscription.
type ImageTemplateManagedImageSource struct {
	// ARM resource id of the managed image in customer subscription
	ImageId string `pulumi:"imageId"`
	// Specifies the type of source image you want to start with.
	// Expected value is 'ManagedImage'.
	Type string `pulumi:"type"`
}

// Describes an image source that is a managed image in customer subscription.
type ImageTemplateManagedImageSourceResponse struct {
	// ARM resource id of the managed image in customer subscription
	ImageId string `pulumi:"imageId"`
	// Specifies the type of source image you want to start with.
	// Expected value is 'ManagedImage'.
	Type string `pulumi:"type"`
}

// Describes an image source from [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
type ImageTemplatePlatformImageSource struct {
	// Image offer from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Offer *string `pulumi:"offer"`
	// Image Publisher in [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Publisher *string `pulumi:"publisher"`
	// Image sku from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Sku *string `pulumi:"sku"`
	// Specifies the type of source image you want to start with.
	// Expected value is 'PlatformImage'.
	Type string `pulumi:"type"`
	// Image version from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Version *string `pulumi:"version"`
}

// Describes an image source from [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
type ImageTemplatePlatformImageSourceResponse struct {
	// Image offer from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Offer *string `pulumi:"offer"`
	// Image Publisher in [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Publisher *string `pulumi:"publisher"`
	// Image sku from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Sku *string `pulumi:"sku"`
	// Specifies the type of source image you want to start with.
	// Expected value is 'PlatformImage'.
	Type string `pulumi:"type"`
	// Image version from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
	Version *string `pulumi:"version"`
}

// Runs the specified PowerShell on the VM (Windows). Corresponds to Packer powershell provisioner. Exactly one of 'script' or 'inline' can be specified.
type ImageTemplatePowerShellCustomizer struct {
	// Array of PowerShell commands to execute
	Inline []string `pulumi:"inline"`
	// Friendly Name to provide context on what this customization step does
	Name *string `pulumi:"name"`
	// The PowerShell script to be run for customizing. It can be a github link, SAS URI for Azure Storage, etc
	Script *string `pulumi:"script"`
	// The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	// Expected value is 'PowerShell'.
	Type string `pulumi:"type"`
	// Valid exit codes for the PowerShell script. [Default: 0]
	ValidExitCodes []int `pulumi:"validExitCodes"`
}

// Runs the specified PowerShell on the VM (Windows). Corresponds to Packer powershell provisioner. Exactly one of 'script' or 'inline' can be specified.
type ImageTemplatePowerShellCustomizerResponse struct {
	// Array of PowerShell commands to execute
	Inline []string `pulumi:"inline"`
	// Friendly Name to provide context on what this customization step does
	Name *string `pulumi:"name"`
	// The PowerShell script to be run for customizing. It can be a github link, SAS URI for Azure Storage, etc
	Script *string `pulumi:"script"`
	// The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	// Expected value is 'PowerShell'.
	Type string `pulumi:"type"`
	// Valid exit codes for the PowerShell script. [Default: 0]
	ValidExitCodes []int `pulumi:"validExitCodes"`
}

// Reboots a VM and waits for it to come back online (Windows). Corresponds to Packer windows-restart provisioner
type ImageTemplateRestartCustomizer struct {
	// Friendly Name to provide context on what this customization step does
	Name *string `pulumi:"name"`
	// Command to check if restart succeeded [Default: '']
	RestartCheckCommand *string `pulumi:"restartCheckCommand"`
	// Command to execute the restart [Default: 'shutdown /r /f /t 0 /c "packer restart"']
	RestartCommand *string `pulumi:"restartCommand"`
	// Restart timeout specified as a string of magnitude and unit, e.g. '5m' (5 minutes) or '2h' (2 hours) [Default: '5m']
	RestartTimeout *string `pulumi:"restartTimeout"`
	// The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	// Expected value is 'WindowsRestart'.
	Type string `pulumi:"type"`
}

// Reboots a VM and waits for it to come back online (Windows). Corresponds to Packer windows-restart provisioner
type ImageTemplateRestartCustomizerResponse struct {
	// Friendly Name to provide context on what this customization step does
	Name *string `pulumi:"name"`
	// Command to check if restart succeeded [Default: '']
	RestartCheckCommand *string `pulumi:"restartCheckCommand"`
	// Command to execute the restart [Default: 'shutdown /r /f /t 0 /c "packer restart"']
	RestartCommand *string `pulumi:"restartCommand"`
	// Restart timeout specified as a string of magnitude and unit, e.g. '5m' (5 minutes) or '2h' (2 hours) [Default: '5m']
	RestartTimeout *string `pulumi:"restartTimeout"`
	// The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	// Expected value is 'WindowsRestart'.
	Type string `pulumi:"type"`
}

// Distribute via Shared Image Gallery.
type ImageTemplateSharedImageDistributor struct {
	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]string `pulumi:"artifactTags"`
	// Resource Id of the Shared Image Gallery image
	GalleryImageId     string   `pulumi:"galleryImageId"`
	ReplicationRegions []string `pulumi:"replicationRegions"`
	// The name to be used for the associated RunOutput.
	RunOutputName string `pulumi:"runOutputName"`
	// Type of distribution.
	// Expected value is 'SharedImage'.
	Type string `pulumi:"type"`
}

// Distribute via Shared Image Gallery.
type ImageTemplateSharedImageDistributorResponse struct {
	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]string `pulumi:"artifactTags"`
	// Resource Id of the Shared Image Gallery image
	GalleryImageId     string   `pulumi:"galleryImageId"`
	ReplicationRegions []string `pulumi:"replicationRegions"`
	// The name to be used for the associated RunOutput.
	RunOutputName string `pulumi:"runOutputName"`
	// Type of distribution.
	// Expected value is 'SharedImage'.
	Type string `pulumi:"type"`
}

// Runs a shell script during the customization phase (Linux). Corresponds to Packer shell provisioner. Exactly one of 'script' or 'inline' can be specified.
type ImageTemplateShellCustomizer struct {
	// Array of shell commands to execute
	Inline []string `pulumi:"inline"`
	// Friendly Name to provide context on what this customization step does
	Name *string `pulumi:"name"`
	// The shell script to be run for customizing. It can be a github link, SAS URI for Azure Storage, etc
	Script *string `pulumi:"script"`
	// The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	// Expected value is 'Shell'.
	Type string `pulumi:"type"`
}

// Runs a shell script during the customization phase (Linux). Corresponds to Packer shell provisioner. Exactly one of 'script' or 'inline' can be specified.
type ImageTemplateShellCustomizerResponse struct {
	// Array of shell commands to execute
	Inline []string `pulumi:"inline"`
	// Friendly Name to provide context on what this customization step does
	Name *string `pulumi:"name"`
	// The shell script to be run for customizing. It can be a github link, SAS URI for Azure Storage, etc
	Script *string `pulumi:"script"`
	// The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	// Expected value is 'Shell'.
	Type string `pulumi:"type"`
}

// Distribute via VHD in a storage account.
type ImageTemplateVhdDistributor struct {
	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]string `pulumi:"artifactTags"`
	// The name to be used for the associated RunOutput.
	RunOutputName string `pulumi:"runOutputName"`
	// Type of distribution.
	// Expected value is 'VHD'.
	Type string `pulumi:"type"`
}

// Distribute via VHD in a storage account.
type ImageTemplateVhdDistributorResponse struct {
	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]string `pulumi:"artifactTags"`
	// The name to be used for the associated RunOutput.
	RunOutputName string `pulumi:"runOutputName"`
	// Type of distribution.
	// Expected value is 'VHD'.
	Type string `pulumi:"type"`
}

type ProvisioningErrorResponse struct {
	// Verbose error message about the provisioning failure
	Message *string `pulumi:"message"`
	// Error code of the provisioning failure
	ProvisioningErrorCode *string `pulumi:"provisioningErrorCode"`
}

type ProvisioningErrorResponseOutput struct{ *pulumi.OutputState }

func (ProvisioningErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningErrorResponse)(nil)).Elem()
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput {
	return o
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponseOutputWithContext(ctx context.Context) ProvisioningErrorResponseOutput {
	return o
}

// Verbose error message about the provisioning failure
func (o ProvisioningErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// Error code of the provisioning failure
func (o ProvisioningErrorResponseOutput) ProvisioningErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.ProvisioningErrorCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponseOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
}