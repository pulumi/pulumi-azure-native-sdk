// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilenetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attached data network resource. Must be created in the same location as its parent packet core data plane.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-06-01.
//
// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type AttachedDataNetwork struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The DNS servers to signal to UEs to use for this attached data network. This configuration is mandatory - if you don't want DNS servers, you must provide an empty array.
	DnsAddresses pulumi.StringArrayOutput `pulumi:"dnsAddresses"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The network address and port translation (NAPT) configuration.
	// If this is not specified, the attached data network will use a default NAPT configuration with NAPT enabled.
	NaptConfiguration NaptConfigurationResponsePtrOutput `pulumi:"naptConfiguration"`
	// The provisioning state of the attached data network resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will dynamically assign IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session.
	//  You must define at least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix. If you define both, they must be of the same size.
	UserEquipmentAddressPoolPrefix pulumi.StringArrayOutput `pulumi:"userEquipmentAddressPoolPrefix"`
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will assign static IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session. The static IP address for a specific UE is set in StaticIPConfiguration on the corresponding SIM resource.
	// At least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix must be defined. If both are defined, they must be of the same size.
	UserEquipmentStaticAddressPoolPrefix pulumi.StringArrayOutput `pulumi:"userEquipmentStaticAddressPoolPrefix"`
	// The user plane interface on the data network. For 5G networks, this is the N6 interface. For 4G networks, this is the SGi interface.
	UserPlaneDataInterface InterfacePropertiesResponseOutput `pulumi:"userPlaneDataInterface"`
}

// NewAttachedDataNetwork registers a new resource with the given unique name, arguments, and options.
func NewAttachedDataNetwork(ctx *pulumi.Context,
	name string, args *AttachedDataNetworkArgs, opts ...pulumi.ResourceOption) (*AttachedDataNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsAddresses == nil {
		return nil, errors.New("invalid value for required argument 'DnsAddresses'")
	}
	if args.PacketCoreControlPlaneName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCoreControlPlaneName'")
	}
	if args.PacketCoreDataPlaneName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCoreDataPlaneName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserPlaneDataInterface == nil {
		return nil, errors.New("invalid value for required argument 'UserPlaneDataInterface'")
	}
	if args.NaptConfiguration != nil {
		args.NaptConfiguration = args.NaptConfiguration.ToNaptConfigurationPtrOutput().ApplyT(func(v *NaptConfiguration) *NaptConfiguration { return v.Defaults() }).(NaptConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230601:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230901:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20240201:AttachedDataNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20240401:AttachedDataNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AttachedDataNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork:AttachedDataNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDataNetwork gets an existing AttachedDataNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDataNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDataNetworkState, opts ...pulumi.ResourceOption) (*AttachedDataNetwork, error) {
	var resource AttachedDataNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork:AttachedDataNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDataNetwork resources.
type attachedDataNetworkState struct {
}

type AttachedDataNetworkState struct {
}

func (AttachedDataNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDataNetworkState)(nil)).Elem()
}

type attachedDataNetworkArgs struct {
	// The name of the attached data network.
	AttachedDataNetworkName *string `pulumi:"attachedDataNetworkName"`
	// The DNS servers to signal to UEs to use for this attached data network. This configuration is mandatory - if you don't want DNS servers, you must provide an empty array.
	DnsAddresses []string `pulumi:"dnsAddresses"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The network address and port translation (NAPT) configuration.
	// If this is not specified, the attached data network will use a default NAPT configuration with NAPT enabled.
	NaptConfiguration *NaptConfiguration `pulumi:"naptConfiguration"`
	// The name of the packet core control plane.
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	// The name of the packet core data plane.
	PacketCoreDataPlaneName string `pulumi:"packetCoreDataPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will dynamically assign IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session.
	//  You must define at least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix. If you define both, they must be of the same size.
	UserEquipmentAddressPoolPrefix []string `pulumi:"userEquipmentAddressPoolPrefix"`
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will assign static IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session. The static IP address for a specific UE is set in StaticIPConfiguration on the corresponding SIM resource.
	// At least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix must be defined. If both are defined, they must be of the same size.
	UserEquipmentStaticAddressPoolPrefix []string `pulumi:"userEquipmentStaticAddressPoolPrefix"`
	// The user plane interface on the data network. For 5G networks, this is the N6 interface. For 4G networks, this is the SGi interface.
	UserPlaneDataInterface InterfaceProperties `pulumi:"userPlaneDataInterface"`
}

// The set of arguments for constructing a AttachedDataNetwork resource.
type AttachedDataNetworkArgs struct {
	// The name of the attached data network.
	AttachedDataNetworkName pulumi.StringPtrInput
	// The DNS servers to signal to UEs to use for this attached data network. This configuration is mandatory - if you don't want DNS servers, you must provide an empty array.
	DnsAddresses pulumi.StringArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The network address and port translation (NAPT) configuration.
	// If this is not specified, the attached data network will use a default NAPT configuration with NAPT enabled.
	NaptConfiguration NaptConfigurationPtrInput
	// The name of the packet core control plane.
	PacketCoreControlPlaneName pulumi.StringInput
	// The name of the packet core data plane.
	PacketCoreDataPlaneName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will dynamically assign IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session.
	//  You must define at least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix. If you define both, they must be of the same size.
	UserEquipmentAddressPoolPrefix pulumi.StringArrayInput
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will assign static IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session. The static IP address for a specific UE is set in StaticIPConfiguration on the corresponding SIM resource.
	// At least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix must be defined. If both are defined, they must be of the same size.
	UserEquipmentStaticAddressPoolPrefix pulumi.StringArrayInput
	// The user plane interface on the data network. For 5G networks, this is the N6 interface. For 4G networks, this is the SGi interface.
	UserPlaneDataInterface InterfacePropertiesInput
}

func (AttachedDataNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDataNetworkArgs)(nil)).Elem()
}

type AttachedDataNetworkInput interface {
	pulumi.Input

	ToAttachedDataNetworkOutput() AttachedDataNetworkOutput
	ToAttachedDataNetworkOutputWithContext(ctx context.Context) AttachedDataNetworkOutput
}

func (*AttachedDataNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetwork)(nil)).Elem()
}

func (i *AttachedDataNetwork) ToAttachedDataNetworkOutput() AttachedDataNetworkOutput {
	return i.ToAttachedDataNetworkOutputWithContext(context.Background())
}

func (i *AttachedDataNetwork) ToAttachedDataNetworkOutputWithContext(ctx context.Context) AttachedDataNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDataNetworkOutput)
}

type AttachedDataNetworkOutput struct{ *pulumi.OutputState }

func (AttachedDataNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetwork)(nil)).Elem()
}

func (o AttachedDataNetworkOutput) ToAttachedDataNetworkOutput() AttachedDataNetworkOutput {
	return o
}

func (o AttachedDataNetworkOutput) ToAttachedDataNetworkOutputWithContext(ctx context.Context) AttachedDataNetworkOutput {
	return o
}

// The Azure API version of the resource.
func (o AttachedDataNetworkOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The DNS servers to signal to UEs to use for this attached data network. This configuration is mandatory - if you don't want DNS servers, you must provide an empty array.
func (o AttachedDataNetworkOutput) DnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringArrayOutput { return v.DnsAddresses }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o AttachedDataNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AttachedDataNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network address and port translation (NAPT) configuration.
// If this is not specified, the attached data network will use a default NAPT configuration with NAPT enabled.
func (o AttachedDataNetworkOutput) NaptConfiguration() NaptConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) NaptConfigurationResponsePtrOutput { return v.NaptConfiguration }).(NaptConfigurationResponsePtrOutput)
}

// The provisioning state of the attached data network resource.
func (o AttachedDataNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AttachedDataNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AttachedDataNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AttachedDataNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will dynamically assign IP addresses to UEs.
// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session.
//
//	You must define at least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix. If you define both, they must be of the same size.
func (o AttachedDataNetworkOutput) UserEquipmentAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringArrayOutput { return v.UserEquipmentAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will assign static IP addresses to UEs.
// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session. The static IP address for a specific UE is set in StaticIPConfiguration on the corresponding SIM resource.
// At least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix must be defined. If both are defined, they must be of the same size.
func (o AttachedDataNetworkOutput) UserEquipmentStaticAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) pulumi.StringArrayOutput { return v.UserEquipmentStaticAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

// The user plane interface on the data network. For 5G networks, this is the N6 interface. For 4G networks, this is the SGi interface.
func (o AttachedDataNetworkOutput) UserPlaneDataInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v *AttachedDataNetwork) InterfacePropertiesResponseOutput { return v.UserPlaneDataInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedDataNetworkOutput{})
}
