// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A NetworkInterface in a resource group
//
// Deprecated: Version 2015-05-01-preview will be removed in v2 of the provider.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// Gets or sets DNS Settings in  NetworkInterface
	DnsSettings NetworkInterfaceDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// Gets or sets whether IPForwarding is enabled on the NIC
	EnableIPForwarding pulumi.BoolPtrOutput `pulumi:"enableIPForwarding"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets or sets list of IPConfigurations of the NetworkInterface
	IpConfigurations NetworkInterfaceIpConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets the MAC Address of the network interface
	MacAddress pulumi.StringPtrOutput `pulumi:"macAddress"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the reference of the NetworkSecurityGroup resource
	NetworkSecurityGroup SubResourceResponsePtrOutput `pulumi:"networkSecurityGroup"`
	// Gets whether this is a primary NIC on a virtual machine
	Primary pulumi.BoolPtrOutput `pulumi:"primary"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets resource guid property of the network interface resource
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the reference of a VirtualMachine
	VirtualMachine SubResourceResponsePtrOutput `pulumi:"virtualMachine"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkInterface"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkInterface
	err := ctx.RegisterResource("azure-native:network/v20150501preview:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("azure-native:network/v20150501preview:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
}

type NetworkInterfaceState struct {
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// Gets or sets DNS Settings in  NetworkInterface
	DnsSettings *NetworkInterfaceDnsSettings `pulumi:"dnsSettings"`
	// Gets or sets whether IPForwarding is enabled on the NIC
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// Gets or sets list of IPConfigurations of the NetworkInterface
	IpConfigurations []NetworkInterfaceIpConfiguration `pulumi:"ipConfigurations"`
	// Resource location
	Location *string `pulumi:"location"`
	// Gets the MAC Address of the network interface
	MacAddress *string `pulumi:"macAddress"`
	// The name of the network interface.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// Gets or sets the reference of the NetworkSecurityGroup resource
	NetworkSecurityGroup *SubResource `pulumi:"networkSecurityGroup"`
	// Gets whether this is a primary NIC on a virtual machine
	Primary *bool `pulumi:"primary"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets resource guid property of the network interface resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the reference of a VirtualMachine
	VirtualMachine *SubResource `pulumi:"virtualMachine"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Gets or sets DNS Settings in  NetworkInterface
	DnsSettings NetworkInterfaceDnsSettingsPtrInput
	// Gets or sets whether IPForwarding is enabled on the NIC
	EnableIPForwarding pulumi.BoolPtrInput
	// Gets or sets list of IPConfigurations of the NetworkInterface
	IpConfigurations NetworkInterfaceIpConfigurationArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// Gets the MAC Address of the network interface
	MacAddress pulumi.StringPtrInput
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringPtrInput
	// Gets or sets the reference of the NetworkSecurityGroup resource
	NetworkSecurityGroup SubResourcePtrInput
	// Gets whether this is a primary NIC on a virtual machine
	Primary pulumi.BoolPtrInput
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets resource guid property of the network interface resource
	ResourceGuid pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Gets or sets the reference of a VirtualMachine
	VirtualMachine SubResourcePtrInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

// Gets or sets DNS Settings in  NetworkInterface
func (o NetworkInterfaceOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceDnsSettingsResponsePtrOutput { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

// Gets or sets whether IPForwarding is enabled on the NIC
func (o NetworkInterfaceOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated
func (o NetworkInterfaceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Gets or sets list of IPConfigurations of the NetworkInterface
func (o NetworkInterfaceOutput) IpConfigurations() NetworkInterfaceIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceIpConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(NetworkInterfaceIpConfigurationResponseArrayOutput)
}

// Resource location
func (o NetworkInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets the MAC Address of the network interface
func (o NetworkInterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// Resource name
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the reference of the NetworkSecurityGroup resource
func (o NetworkInterfaceOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponsePtrOutput { return v.NetworkSecurityGroup }).(SubResourceResponsePtrOutput)
}

// Gets whether this is a primary NIC on a virtual machine
func (o NetworkInterfaceOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.Primary }).(pulumi.BoolPtrOutput)
}

// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets resource guid property of the network interface resource
func (o NetworkInterfaceOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the reference of a VirtualMachine
func (o NetworkInterfaceOutput) VirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponsePtrOutput { return v.VirtualMachine }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}