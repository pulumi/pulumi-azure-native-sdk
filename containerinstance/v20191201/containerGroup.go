// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A container group.
//
// Deprecated: Version 2019-12-01 will be removed in v2 of the provider.
type ContainerGroup struct {
	pulumi.CustomResourceState

	// The containers within the container group.
	Containers ContainerResponseArrayOutput `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics ContainerGroupDiagnosticsResponsePtrOutput `pulumi:"diagnostics"`
	// The DNS config information for a container group.
	DnsConfig DnsConfigurationResponsePtrOutput `pulumi:"dnsConfig"`
	// The encryption properties for a container group.
	EncryptionProperties EncryptionPropertiesResponsePtrOutput `pulumi:"encryptionProperties"`
	// The identity of the container group, if configured.
	Identity ContainerGroupIdentityResponsePtrOutput `pulumi:"identity"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials ImageRegistryCredentialResponseArrayOutput `pulumi:"imageRegistryCredentials"`
	// The init containers for a container group.
	InitContainers InitContainerDefinitionResponseArrayOutput `pulumi:"initContainers"`
	// The instance view of the container group. Only valid in response.
	InstanceView ContainerGroupResponseInstanceViewOutput `pulumi:"instanceView"`
	// The IP address type of the container group.
	IpAddress IpAddressResponsePtrOutput `pulumi:"ipAddress"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network profile information for a container group.
	NetworkProfile ContainerGroupNetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// The operating system type required by the containers in the container group.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The provisioning state of the container group. This only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy pulumi.StringPtrOutput `pulumi:"restartPolicy"`
	// The SKU for a container group.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes VolumeResponseArrayOutput `pulumi:"volumes"`
}

// NewContainerGroup registers a new resource with the given unique name, arguments, and options.
func NewContainerGroup(ctx *pulumi.Context,
	name string, args *ContainerGroupArgs, opts ...pulumi.ResourceOption) (*ContainerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Containers == nil {
		return nil, errors.New("invalid value for required argument 'Containers'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerinstance:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20170801preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20171001preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20171201preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180201preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180401:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180601:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180901:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20181001:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20201101:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20210301:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20210701:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20210901:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20211001:ContainerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerGroup
	err := ctx.RegisterResource("azure-native:containerinstance/v20191201:ContainerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerGroup gets an existing ContainerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerGroupState, opts ...pulumi.ResourceOption) (*ContainerGroup, error) {
	var resource ContainerGroup
	err := ctx.ReadResource("azure-native:containerinstance/v20191201:ContainerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerGroup resources.
type containerGroupState struct {
}

type ContainerGroupState struct {
}

func (ContainerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupState)(nil)).Elem()
}

type containerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName *string `pulumi:"containerGroupName"`
	// The containers within the container group.
	Containers []Container `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnostics `pulumi:"diagnostics"`
	// The DNS config information for a container group.
	DnsConfig *DnsConfiguration `pulumi:"dnsConfig"`
	// The encryption properties for a container group.
	EncryptionProperties *EncryptionProperties `pulumi:"encryptionProperties"`
	// The identity of the container group, if configured.
	Identity *ContainerGroupIdentity `pulumi:"identity"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The init containers for a container group.
	InitContainers []InitContainerDefinition `pulumi:"initContainers"`
	// The IP address type of the container group.
	IpAddress *IpAddress `pulumi:"ipAddress"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The network profile information for a container group.
	NetworkProfile *ContainerGroupNetworkProfile `pulumi:"networkProfile"`
	// The operating system type required by the containers in the container group.
	OsType string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The SKU for a container group.
	Sku *string `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes []Volume `pulumi:"volumes"`
}

// The set of arguments for constructing a ContainerGroup resource.
type ContainerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName pulumi.StringPtrInput
	// The containers within the container group.
	Containers ContainerArrayInput
	// The diagnostic information for a container group.
	Diagnostics ContainerGroupDiagnosticsPtrInput
	// The DNS config information for a container group.
	DnsConfig DnsConfigurationPtrInput
	// The encryption properties for a container group.
	EncryptionProperties EncryptionPropertiesPtrInput
	// The identity of the container group, if configured.
	Identity ContainerGroupIdentityPtrInput
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials ImageRegistryCredentialArrayInput
	// The init containers for a container group.
	InitContainers InitContainerDefinitionArrayInput
	// The IP address type of the container group.
	IpAddress IpAddressPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The network profile information for a container group.
	NetworkProfile ContainerGroupNetworkProfilePtrInput
	// The operating system type required by the containers in the container group.
	OsType pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy pulumi.StringPtrInput
	// The SKU for a container group.
	Sku pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The list of volumes that can be mounted by containers in this container group.
	Volumes VolumeArrayInput
}

func (ContainerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupArgs)(nil)).Elem()
}

type ContainerGroupInput interface {
	pulumi.Input

	ToContainerGroupOutput() ContainerGroupOutput
	ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput
}

func (*ContainerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroup)(nil)).Elem()
}

func (i *ContainerGroup) ToContainerGroupOutput() ContainerGroupOutput {
	return i.ToContainerGroupOutputWithContext(context.Background())
}

func (i *ContainerGroup) ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupOutput)
}

type ContainerGroupOutput struct{ *pulumi.OutputState }

func (ContainerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroup)(nil)).Elem()
}

func (o ContainerGroupOutput) ToContainerGroupOutput() ContainerGroupOutput {
	return o
}

func (o ContainerGroupOutput) ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput {
	return o
}

// The containers within the container group.
func (o ContainerGroupOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerResponseArrayOutput { return v.Containers }).(ContainerResponseArrayOutput)
}

// The diagnostic information for a container group.
func (o ContainerGroupOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupDiagnosticsResponsePtrOutput { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

// The DNS config information for a container group.
func (o ContainerGroupOutput) DnsConfig() DnsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) DnsConfigurationResponsePtrOutput { return v.DnsConfig }).(DnsConfigurationResponsePtrOutput)
}

// The encryption properties for a container group.
func (o ContainerGroupOutput) EncryptionProperties() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) EncryptionPropertiesResponsePtrOutput { return v.EncryptionProperties }).(EncryptionPropertiesResponsePtrOutput)
}

// The identity of the container group, if configured.
func (o ContainerGroupOutput) Identity() ContainerGroupIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupIdentityResponsePtrOutput { return v.Identity }).(ContainerGroupIdentityResponsePtrOutput)
}

// The image registry credentials by which the container group is created from.
func (o ContainerGroupOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) ImageRegistryCredentialResponseArrayOutput { return v.ImageRegistryCredentials }).(ImageRegistryCredentialResponseArrayOutput)
}

// The init containers for a container group.
func (o ContainerGroupOutput) InitContainers() InitContainerDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) InitContainerDefinitionResponseArrayOutput { return v.InitContainers }).(InitContainerDefinitionResponseArrayOutput)
}

// The instance view of the container group. Only valid in response.
func (o ContainerGroupOutput) InstanceView() ContainerGroupResponseInstanceViewOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupResponseInstanceViewOutput { return v.InstanceView }).(ContainerGroupResponseInstanceViewOutput)
}

// The IP address type of the container group.
func (o ContainerGroupOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) IpAddressResponsePtrOutput { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

// The resource location.
func (o ContainerGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o ContainerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network profile information for a container group.
func (o ContainerGroupOutput) NetworkProfile() ContainerGroupNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(ContainerGroupNetworkProfileResponsePtrOutput)
}

// The operating system type required by the containers in the container group.
func (o ContainerGroupOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// The provisioning state of the container group. This only appears in the response.
func (o ContainerGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Restart policy for all containers within the container group.
// - `Always` Always restart
// - `OnFailure` Restart on failure
// - `Never` Never restart
func (o ContainerGroupOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

// The SKU for a container group.
func (o ContainerGroupOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o ContainerGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o ContainerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of volumes that can be mounted by containers in this container group.
func (o ContainerGroupOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) VolumeResponseArrayOutput { return v.Volumes }).(VolumeResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerGroupOutput{})
}