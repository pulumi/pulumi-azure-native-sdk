// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private endpoint resource.
//
// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
//
// Other available API versions: 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type PrivateEndpoint struct {
	pulumi.CustomResourceState

	// Application security groups in which the private endpoint IP configuration is included.
	ApplicationSecurityGroups ApplicationSecurityGroupResponseArrayOutput `pulumi:"applicationSecurityGroups"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// An array of custom dns configurations.
	CustomDnsConfigs CustomDnsConfigPropertiesFormatResponseArrayOutput `pulumi:"customDnsConfigs"`
	// The custom name of the network interface attached to the private endpoint.
	CustomNetworkInterfaceName pulumi.StringPtrOutput `pulumi:"customNetworkInterfaceName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the load balancer.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// A list of IP configurations of the private endpoint. This will be used to map to the First Party Service's endpoints.
	IpConfigurations PrivateEndpointIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
	ManualPrivateLinkServiceConnections PrivateLinkServiceConnectionResponseArrayOutput `pulumi:"manualPrivateLinkServiceConnections"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of references to the network interfaces created for this private endpoint.
	NetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	// A grouping of information about the connection to the remote resource.
	PrivateLinkServiceConnections PrivateLinkServiceConnectionResponseArrayOutput `pulumi:"privateLinkServiceConnections"`
	// The provisioning state of the private endpoint resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet SubnetResponsePtrOutput `pulumi:"subnet"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpoint(ctx *pulumi.Context,
	name string, args *PrivateEndpointArgs, opts ...pulumi.ResourceOption) (*PrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Subnet != nil {
		args.Subnet = args.Subnet.ToSubnetTypePtrOutput().ApplyT(func(v *SubnetType) *SubnetType { return v.Defaults() }).(SubnetTypePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20231101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network:InterfaceEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpoint
	err := ctx.RegisterResource("azure-native:network:PrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpoint gets an existing PrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointState, opts ...pulumi.ResourceOption) (*PrivateEndpoint, error) {
	var resource PrivateEndpoint
	err := ctx.ReadResource("azure-native:network:PrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpoint resources.
type privateEndpointState struct {
}

type PrivateEndpointState struct {
}

func (PrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointState)(nil)).Elem()
}

type privateEndpointArgs struct {
	// Application security groups in which the private endpoint IP configuration is included.
	ApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"applicationSecurityGroups"`
	// An array of custom dns configurations.
	CustomDnsConfigs []CustomDnsConfigPropertiesFormat `pulumi:"customDnsConfigs"`
	// The custom name of the network interface attached to the private endpoint.
	CustomNetworkInterfaceName *string `pulumi:"customNetworkInterfaceName"`
	// The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IP configurations of the private endpoint. This will be used to map to the First Party Service's endpoints.
	IpConfigurations []PrivateEndpointIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection `pulumi:"manualPrivateLinkServiceConnections"`
	// The name of the private endpoint.
	PrivateEndpointName *string `pulumi:"privateEndpointName"`
	// A grouping of information about the connection to the remote resource.
	PrivateLinkServiceConnections []PrivateLinkServiceConnection `pulumi:"privateLinkServiceConnections"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet *SubnetType `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateEndpoint resource.
type PrivateEndpointArgs struct {
	// Application security groups in which the private endpoint IP configuration is included.
	ApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput
	// An array of custom dns configurations.
	CustomDnsConfigs CustomDnsConfigPropertiesFormatArrayInput
	// The custom name of the network interface attached to the private endpoint.
	CustomNetworkInterfaceName pulumi.StringPtrInput
	// The extended location of the load balancer.
	ExtendedLocation ExtendedLocationPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// A list of IP configurations of the private endpoint. This will be used to map to the First Party Service's endpoints.
	IpConfigurations PrivateEndpointIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
	ManualPrivateLinkServiceConnections PrivateLinkServiceConnectionArrayInput
	// The name of the private endpoint.
	PrivateEndpointName pulumi.StringPtrInput
	// A grouping of information about the connection to the remote resource.
	PrivateLinkServiceConnections PrivateLinkServiceConnectionArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The ID of the subnet from which the private IP will be allocated.
	Subnet SubnetTypePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointArgs)(nil)).Elem()
}

type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput
}

func (*PrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *PrivateEndpoint) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i *PrivateEndpoint) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

// Application security groups in which the private endpoint IP configuration is included.
func (o PrivateEndpointOutput) ApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v *PrivateEndpoint) ApplicationSecurityGroupResponseArrayOutput {
		return v.ApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

// The Azure API version of the resource.
func (o PrivateEndpointOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// An array of custom dns configurations.
func (o PrivateEndpointOutput) CustomDnsConfigs() CustomDnsConfigPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *PrivateEndpoint) CustomDnsConfigPropertiesFormatResponseArrayOutput { return v.CustomDnsConfigs }).(CustomDnsConfigPropertiesFormatResponseArrayOutput)
}

// The custom name of the network interface attached to the private endpoint.
func (o PrivateEndpointOutput) CustomNetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringPtrOutput { return v.CustomNetworkInterfaceName }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o PrivateEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the load balancer.
func (o PrivateEndpointOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// A list of IP configurations of the private endpoint. This will be used to map to the First Party Service's endpoints.
func (o PrivateEndpointOutput) IpConfigurations() PrivateEndpointIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpointIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(PrivateEndpointIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o PrivateEndpointOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
func (o PrivateEndpointOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateLinkServiceConnectionResponseArrayOutput {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

// Resource name.
func (o PrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of references to the network interfaces created for this private endpoint.
func (o PrivateEndpointOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *PrivateEndpoint) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// A grouping of information about the connection to the remote resource.
func (o PrivateEndpointOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateLinkServiceConnectionResponseArrayOutput {
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

// The provisioning state of the private endpoint resource.
func (o PrivateEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID of the subnet from which the private IP will be allocated.
func (o PrivateEndpointOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) SubnetResponsePtrOutput { return v.Subnet }).(SubnetResponsePtrOutput)
}

// Resource tags.
func (o PrivateEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o PrivateEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
}
