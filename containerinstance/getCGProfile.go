// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerinstance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the properties of the specified container group profile.
//
// Uses Azure REST API version 2024-11-01-preview.
func LookupCGProfile(ctx *pulumi.Context, args *LookupCGProfileArgs, opts ...pulumi.InvokeOption) (*LookupCGProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCGProfileResult
	err := ctx.Invoke("azure-native:containerinstance:getCGProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCGProfileArgs struct {
	// ContainerGroupProfile name.
	ContainerGroupProfileName string `pulumi:"containerGroupProfileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// container group profile object
type LookupCGProfileResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The properties for confidential container group
	ConfidentialComputeProperties *ConfidentialComputePropertiesResponse `pulumi:"confidentialComputeProperties"`
	// The containers within the container group.
	Containers []ContainerResponse `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnosticsResponse `pulumi:"diagnostics"`
	// The encryption properties for a container group.
	EncryptionProperties *EncryptionPropertiesResponse `pulumi:"encryptionProperties"`
	// extensions used by virtual kubelet
	Extensions []DeploymentExtensionSpecResponse `pulumi:"extensions"`
	// The resource id.
	Id string `pulumi:"id"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredentialResponse `pulumi:"imageRegistryCredentials"`
	// The init containers for a container group.
	InitContainers []InitContainerDefinitionResponse `pulumi:"initContainers"`
	// The IP address type of the container group.
	IpAddress *IpAddressResponse `pulumi:"ipAddress"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The operating system type required by the containers in the container group.
	OsType *string `pulumi:"osType"`
	// The priority of the container group.
	Priority *string `pulumi:"priority"`
	// Registered revisions are calculated at request time based off the records in the table logs.
	RegisteredRevisions []float64 `pulumi:"registeredRevisions"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *string `pulumi:"restartPolicy"`
	// Container group profile current revision number
	Revision *float64 `pulumi:"revision"`
	// The container security properties.
	SecurityContext *SecurityContextDefinitionResponse `pulumi:"securityContext"`
	// Shutdown grace period for containers in a container group.
	ShutdownGracePeriod *string `pulumi:"shutdownGracePeriod"`
	// The SKU for a container group.
	Sku *string `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Post completion time to live for containers of a CG
	TimeToLive *string `pulumi:"timeToLive"`
	// The resource type.
	Type string `pulumi:"type"`
	// Gets or sets Krypton use property.
	UseKrypton *bool `pulumi:"useKrypton"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes []VolumeResponse `pulumi:"volumes"`
	// The zones for the container group.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupCGProfileResult
func (val *LookupCGProfileResult) Defaults() *LookupCGProfileResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.IpAddress = tmp.IpAddress.Defaults()

	return &tmp
}
func LookupCGProfileOutput(ctx *pulumi.Context, args LookupCGProfileOutputArgs, opts ...pulumi.InvokeOption) LookupCGProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCGProfileResultOutput, error) {
			args := v.(LookupCGProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:containerinstance:getCGProfile", args, LookupCGProfileResultOutput{}, options).(LookupCGProfileResultOutput), nil
		}).(LookupCGProfileResultOutput)
}

type LookupCGProfileOutputArgs struct {
	// ContainerGroupProfile name.
	ContainerGroupProfileName pulumi.StringInput `pulumi:"containerGroupProfileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCGProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCGProfileArgs)(nil)).Elem()
}

// container group profile object
type LookupCGProfileResultOutput struct{ *pulumi.OutputState }

func (LookupCGProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCGProfileResult)(nil)).Elem()
}

func (o LookupCGProfileResultOutput) ToLookupCGProfileResultOutput() LookupCGProfileResultOutput {
	return o
}

func (o LookupCGProfileResultOutput) ToLookupCGProfileResultOutputWithContext(ctx context.Context) LookupCGProfileResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupCGProfileResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCGProfileResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The properties for confidential container group
func (o LookupCGProfileResultOutput) ConfidentialComputeProperties() ConfidentialComputePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *ConfidentialComputePropertiesResponse {
		return v.ConfidentialComputeProperties
	}).(ConfidentialComputePropertiesResponsePtrOutput)
}

// The containers within the container group.
func (o LookupCGProfileResultOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

// The diagnostic information for a container group.
func (o LookupCGProfileResultOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *ContainerGroupDiagnosticsResponse { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

// The encryption properties for a container group.
func (o LookupCGProfileResultOutput) EncryptionProperties() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *EncryptionPropertiesResponse { return v.EncryptionProperties }).(EncryptionPropertiesResponsePtrOutput)
}

// extensions used by virtual kubelet
func (o LookupCGProfileResultOutput) Extensions() DeploymentExtensionSpecResponseArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []DeploymentExtensionSpecResponse { return v.Extensions }).(DeploymentExtensionSpecResponseArrayOutput)
}

// The resource id.
func (o LookupCGProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCGProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The image registry credentials by which the container group is created from.
func (o LookupCGProfileResultOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []ImageRegistryCredentialResponse { return v.ImageRegistryCredentials }).(ImageRegistryCredentialResponseArrayOutput)
}

// The init containers for a container group.
func (o LookupCGProfileResultOutput) InitContainers() InitContainerDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []InitContainerDefinitionResponse { return v.InitContainers }).(InitContainerDefinitionResponseArrayOutput)
}

// The IP address type of the container group.
func (o LookupCGProfileResultOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *IpAddressResponse { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

// The resource location.
func (o LookupCGProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupCGProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCGProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The operating system type required by the containers in the container group.
func (o LookupCGProfileResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// The priority of the container group.
func (o LookupCGProfileResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// Registered revisions are calculated at request time based off the records in the table logs.
func (o LookupCGProfileResultOutput) RegisteredRevisions() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []float64 { return v.RegisteredRevisions }).(pulumi.Float64ArrayOutput)
}

// Restart policy for all containers within the container group.
// - `Always` Always restart
// - `OnFailure` Restart on failure
// - `Never` Never restart
func (o LookupCGProfileResultOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

// Container group profile current revision number
func (o LookupCGProfileResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

// The container security properties.
func (o LookupCGProfileResultOutput) SecurityContext() SecurityContextDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *SecurityContextDefinitionResponse { return v.SecurityContext }).(SecurityContextDefinitionResponsePtrOutput)
}

// Shutdown grace period for containers in a container group.
func (o LookupCGProfileResultOutput) ShutdownGracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.ShutdownGracePeriod }).(pulumi.StringPtrOutput)
}

// The SKU for a container group.
func (o LookupCGProfileResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCGProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCGProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o LookupCGProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCGProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Post completion time to live for containers of a CG
func (o LookupCGProfileResultOutput) TimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *string { return v.TimeToLive }).(pulumi.StringPtrOutput)
}

// The resource type.
func (o LookupCGProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCGProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets Krypton use property.
func (o LookupCGProfileResultOutput) UseKrypton() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCGProfileResult) *bool { return v.UseKrypton }).(pulumi.BoolPtrOutput)
}

// The list of volumes that can be mounted by containers in this container group.
func (o LookupCGProfileResultOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []VolumeResponse { return v.Volumes }).(VolumeResponseArrayOutput)
}

// The zones for the container group.
func (o LookupCGProfileResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCGProfileResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCGProfileResultOutput{})
}
