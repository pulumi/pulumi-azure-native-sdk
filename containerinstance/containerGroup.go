// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerinstance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A container group.
//
// Uses Azure REST API version 2024-05-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-01.
//
// Other available API versions: 2023-05-01, 2024-09-01-preview, 2024-10-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerinstance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ContainerGroup struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The properties for confidential container group
	ConfidentialComputeProperties ConfidentialComputePropertiesResponsePtrOutput `pulumi:"confidentialComputeProperties"`
	// The reference container group profile properties.
	ContainerGroupProfile ContainerGroupProfileReferenceDefinitionResponsePtrOutput `pulumi:"containerGroupProfile"`
	// The containers within the container group.
	Containers ContainerResponseArrayOutput `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics ContainerGroupDiagnosticsResponsePtrOutput `pulumi:"diagnostics"`
	// The DNS config information for a container group.
	DnsConfig DnsConfigurationResponsePtrOutput `pulumi:"dnsConfig"`
	// The encryption properties for a container group.
	EncryptionProperties EncryptionPropertiesResponsePtrOutput `pulumi:"encryptionProperties"`
	// extensions used by virtual kubelet
	Extensions DeploymentExtensionSpecResponseArrayOutput `pulumi:"extensions"`
	// The identity of the container group, if configured.
	Identity ContainerGroupIdentityResponsePtrOutput `pulumi:"identity"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials ImageRegistryCredentialResponseArrayOutput `pulumi:"imageRegistryCredentials"`
	// The init containers for a container group.
	InitContainers InitContainerDefinitionResponseArrayOutput `pulumi:"initContainers"`
	// The instance view of the container group. Only valid in response.
	InstanceView ContainerGroupPropertiesResponseInstanceViewOutput `pulumi:"instanceView"`
	// The IP address type of the container group.
	IpAddress IpAddressResponsePtrOutput `pulumi:"ipAddress"`
	// The flag indicating whether the container group is created by standby pool.
	IsCreatedFromStandbyPool pulumi.BoolOutput `pulumi:"isCreatedFromStandbyPool"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The operating system type required by the containers in the container group.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The priority of the container group.
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The provisioning state of the container group. This only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy pulumi.StringPtrOutput `pulumi:"restartPolicy"`
	// The SKU for a container group.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// The reference standby pool profile properties.
	StandbyPoolProfile StandbyPoolProfileDefinitionResponsePtrOutput `pulumi:"standbyPoolProfile"`
	// The subnet resource IDs for a container group.
	SubnetIds ContainerGroupSubnetIdResponseArrayOutput `pulumi:"subnetIds"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes VolumeResponseArrayOutput `pulumi:"volumes"`
	// The zones for the container group.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
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
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.IpAddress != nil {
		args.IpAddress = args.IpAddress.ToIpAddressPtrOutput().ApplyT(func(v *IpAddress) *IpAddress { return v.Defaults() }).(IpAddressPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
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
			Type: pulumi.String("azure-native:containerinstance/v20191201:ContainerGroup"),
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
		{
			Type: pulumi.String("azure-native:containerinstance/v20220901:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20221001preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20230201preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20230501:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20240501preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20240901preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20241001preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20241101preview:ContainerGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContainerGroup
	err := ctx.RegisterResource("azure-native:containerinstance:ContainerGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerinstance:ContainerGroup", name, id, state, &resource, opts...)
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
	// The properties for confidential container group
	ConfidentialComputeProperties *ConfidentialComputeProperties `pulumi:"confidentialComputeProperties"`
	// The name of the container group.
	ContainerGroupName *string `pulumi:"containerGroupName"`
	// The reference container group profile properties.
	ContainerGroupProfile *ContainerGroupProfileReferenceDefinition `pulumi:"containerGroupProfile"`
	// The containers within the container group.
	Containers []Container `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnostics `pulumi:"diagnostics"`
	// The DNS config information for a container group.
	DnsConfig *DnsConfiguration `pulumi:"dnsConfig"`
	// The encryption properties for a container group.
	EncryptionProperties *EncryptionProperties `pulumi:"encryptionProperties"`
	// extensions used by virtual kubelet
	Extensions []DeploymentExtensionSpec `pulumi:"extensions"`
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
	// The operating system type required by the containers in the container group.
	OsType *string `pulumi:"osType"`
	// The priority of the container group.
	Priority *string `pulumi:"priority"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The SKU for a container group.
	Sku *string `pulumi:"sku"`
	// The reference standby pool profile properties.
	StandbyPoolProfile *StandbyPoolProfileDefinition `pulumi:"standbyPoolProfile"`
	// The subnet resource IDs for a container group.
	SubnetIds []ContainerGroupSubnetId `pulumi:"subnetIds"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes []Volume `pulumi:"volumes"`
	// The zones for the container group.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a ContainerGroup resource.
type ContainerGroupArgs struct {
	// The properties for confidential container group
	ConfidentialComputeProperties ConfidentialComputePropertiesPtrInput
	// The name of the container group.
	ContainerGroupName pulumi.StringPtrInput
	// The reference container group profile properties.
	ContainerGroupProfile ContainerGroupProfileReferenceDefinitionPtrInput
	// The containers within the container group.
	Containers ContainerArrayInput
	// The diagnostic information for a container group.
	Diagnostics ContainerGroupDiagnosticsPtrInput
	// The DNS config information for a container group.
	DnsConfig DnsConfigurationPtrInput
	// The encryption properties for a container group.
	EncryptionProperties EncryptionPropertiesPtrInput
	// extensions used by virtual kubelet
	Extensions DeploymentExtensionSpecArrayInput
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
	// The operating system type required by the containers in the container group.
	OsType pulumi.StringPtrInput
	// The priority of the container group.
	Priority pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy pulumi.StringPtrInput
	// The SKU for a container group.
	Sku pulumi.StringPtrInput
	// The reference standby pool profile properties.
	StandbyPoolProfile StandbyPoolProfileDefinitionPtrInput
	// The subnet resource IDs for a container group.
	SubnetIds ContainerGroupSubnetIdArrayInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The list of volumes that can be mounted by containers in this container group.
	Volumes VolumeArrayInput
	// The zones for the container group.
	Zones pulumi.StringArrayInput
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

// The Azure API version of the resource.
func (o ContainerGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The properties for confidential container group
func (o ContainerGroupOutput) ConfidentialComputeProperties() ConfidentialComputePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ConfidentialComputePropertiesResponsePtrOutput {
		return v.ConfidentialComputeProperties
	}).(ConfidentialComputePropertiesResponsePtrOutput)
}

// The reference container group profile properties.
func (o ContainerGroupOutput) ContainerGroupProfile() ContainerGroupProfileReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupProfileReferenceDefinitionResponsePtrOutput {
		return v.ContainerGroupProfile
	}).(ContainerGroupProfileReferenceDefinitionResponsePtrOutput)
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

// extensions used by virtual kubelet
func (o ContainerGroupOutput) Extensions() DeploymentExtensionSpecResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) DeploymentExtensionSpecResponseArrayOutput { return v.Extensions }).(DeploymentExtensionSpecResponseArrayOutput)
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
func (o ContainerGroupOutput) InstanceView() ContainerGroupPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupPropertiesResponseInstanceViewOutput { return v.InstanceView }).(ContainerGroupPropertiesResponseInstanceViewOutput)
}

// The IP address type of the container group.
func (o ContainerGroupOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) IpAddressResponsePtrOutput { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

// The flag indicating whether the container group is created by standby pool.
func (o ContainerGroupOutput) IsCreatedFromStandbyPool() pulumi.BoolOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.BoolOutput { return v.IsCreatedFromStandbyPool }).(pulumi.BoolOutput)
}

// The resource location.
func (o ContainerGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o ContainerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The operating system type required by the containers in the container group.
func (o ContainerGroupOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// The priority of the container group.
func (o ContainerGroupOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.Priority }).(pulumi.StringPtrOutput)
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

// The reference standby pool profile properties.
func (o ContainerGroupOutput) StandbyPoolProfile() StandbyPoolProfileDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) StandbyPoolProfileDefinitionResponsePtrOutput { return v.StandbyPoolProfile }).(StandbyPoolProfileDefinitionResponsePtrOutput)
}

// The subnet resource IDs for a container group.
func (o ContainerGroupOutput) SubnetIds() ContainerGroupSubnetIdResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupSubnetIdResponseArrayOutput { return v.SubnetIds }).(ContainerGroupSubnetIdResponseArrayOutput)
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

// The zones for the container group.
func (o ContainerGroupOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerGroupOutput{})
}
