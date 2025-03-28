// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified container group in the specified subscription and resource group. The operation returns the properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
func LookupContainerGroup(ctx *pulumi.Context, args *LookupContainerGroupArgs, opts ...pulumi.InvokeOption) (*LookupContainerGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerGroupResult
	err := ctx.Invoke("azure-native:containerinstance/v20210301:getContainerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName string `pulumi:"containerGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A container group.
type LookupContainerGroupResult struct {
	// The containers within the container group.
	Containers []ContainerResponse `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnosticsResponse `pulumi:"diagnostics"`
	// The DNS config information for a container group.
	DnsConfig *DnsConfigurationResponse `pulumi:"dnsConfig"`
	// The encryption properties for a container group.
	EncryptionProperties *EncryptionPropertiesResponse `pulumi:"encryptionProperties"`
	// The resource id.
	Id string `pulumi:"id"`
	// The identity of the container group, if configured.
	Identity *ContainerGroupIdentityResponse `pulumi:"identity"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredentialResponse `pulumi:"imageRegistryCredentials"`
	// The init containers for a container group.
	InitContainers []InitContainerDefinitionResponse `pulumi:"initContainers"`
	// The instance view of the container group. Only valid in response.
	InstanceView ContainerGroupResponseInstanceView `pulumi:"instanceView"`
	// The IP address type of the container group.
	IpAddress *IpAddressResponse `pulumi:"ipAddress"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The network profile information for a container group.
	NetworkProfile *ContainerGroupNetworkProfileResponse `pulumi:"networkProfile"`
	// The operating system type required by the containers in the container group.
	OsType string `pulumi:"osType"`
	// The provisioning state of the container group. This only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The SKU for a container group.
	Sku *string `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes []VolumeResponse `pulumi:"volumes"`
}

func LookupContainerGroupOutput(ctx *pulumi.Context, args LookupContainerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupContainerGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupContainerGroupResultOutput, error) {
			args := v.(LookupContainerGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:containerinstance/v20210301:getContainerGroup", args, LookupContainerGroupResultOutput{}, options).(LookupContainerGroupResultOutput), nil
		}).(LookupContainerGroupResultOutput)
}

type LookupContainerGroupOutputArgs struct {
	// The name of the container group.
	ContainerGroupName pulumi.StringInput `pulumi:"containerGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupArgs)(nil)).Elem()
}

// A container group.
type LookupContainerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupContainerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupResult)(nil)).Elem()
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutput() LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutputWithContext(ctx context.Context) LookupContainerGroupResultOutput {
	return o
}

// The containers within the container group.
func (o LookupContainerGroupResultOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

// The diagnostic information for a container group.
func (o LookupContainerGroupResultOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupDiagnosticsResponse { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

// The DNS config information for a container group.
func (o LookupContainerGroupResultOutput) DnsConfig() DnsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *DnsConfigurationResponse { return v.DnsConfig }).(DnsConfigurationResponsePtrOutput)
}

// The encryption properties for a container group.
func (o LookupContainerGroupResultOutput) EncryptionProperties() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *EncryptionPropertiesResponse { return v.EncryptionProperties }).(EncryptionPropertiesResponsePtrOutput)
}

// The resource id.
func (o LookupContainerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the container group, if configured.
func (o LookupContainerGroupResultOutput) Identity() ContainerGroupIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupIdentityResponse { return v.Identity }).(ContainerGroupIdentityResponsePtrOutput)
}

// The image registry credentials by which the container group is created from.
func (o LookupContainerGroupResultOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ImageRegistryCredentialResponse {
		return v.ImageRegistryCredentials
	}).(ImageRegistryCredentialResponseArrayOutput)
}

// The init containers for a container group.
func (o LookupContainerGroupResultOutput) InitContainers() InitContainerDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []InitContainerDefinitionResponse { return v.InitContainers }).(InitContainerDefinitionResponseArrayOutput)
}

// The instance view of the container group. Only valid in response.
func (o LookupContainerGroupResultOutput) InstanceView() ContainerGroupResponseInstanceViewOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) ContainerGroupResponseInstanceView { return v.InstanceView }).(ContainerGroupResponseInstanceViewOutput)
}

// The IP address type of the container group.
func (o LookupContainerGroupResultOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *IpAddressResponse { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

// The resource location.
func (o LookupContainerGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupContainerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network profile information for a container group.
func (o LookupContainerGroupResultOutput) NetworkProfile() ContainerGroupNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupNetworkProfileResponse { return v.NetworkProfile }).(ContainerGroupNetworkProfileResponsePtrOutput)
}

// The operating system type required by the containers in the container group.
func (o LookupContainerGroupResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.OsType }).(pulumi.StringOutput)
}

// The provisioning state of the container group. This only appears in the response.
func (o LookupContainerGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Restart policy for all containers within the container group.
// - `Always` Always restart
// - `OnFailure` Restart on failure
// - `Never` Never restart
func (o LookupContainerGroupResultOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

// The SKU for a container group.
func (o LookupContainerGroupResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o LookupContainerGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupContainerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The list of volumes that can be mounted by containers in this container group.
func (o LookupContainerGroupResultOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []VolumeResponse { return v.Volumes }).(VolumeResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerGroupResultOutput{})
}
