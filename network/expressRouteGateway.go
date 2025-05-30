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

// ExpressRoute gateway resource.
//
// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
//
// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ExpressRouteGateway struct {
	pulumi.CustomResourceState

	// Configures this gateway to accept traffic from non Virtual WAN networks.
	AllowNonVirtualWanTraffic pulumi.BoolPtrOutput `pulumi:"allowNonVirtualWanTraffic"`
	// Configuration for auto scaling.
	AutoScaleConfiguration ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput `pulumi:"autoScaleConfiguration"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of ExpressRoute connections to the ExpressRoute gateway.
	ExpressRouteConnections ExpressRouteConnectionResponseArrayOutput `pulumi:"expressRouteConnections"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the express route gateway resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubIdResponseOutput `pulumi:"virtualHub"`
}

// NewExpressRouteGateway registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteGateway(ctx *pulumi.Context,
	name string, args *ExpressRouteGatewayArgs, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHub == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHub'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20231101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240701:ExpressRouteGateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExpressRouteGateway
	err := ctx.RegisterResource("azure-native:network:ExpressRouteGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteGateway gets an existing ExpressRouteGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteGatewayState, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	var resource ExpressRouteGateway
	err := ctx.ReadResource("azure-native:network:ExpressRouteGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteGateway resources.
type expressRouteGatewayState struct {
}

type ExpressRouteGatewayState struct {
}

func (ExpressRouteGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayState)(nil)).Elem()
}

type expressRouteGatewayArgs struct {
	// Configures this gateway to accept traffic from non Virtual WAN networks.
	AllowNonVirtualWanTraffic *bool `pulumi:"allowNonVirtualWanTraffic"`
	// Configuration for auto scaling.
	AutoScaleConfiguration *ExpressRouteGatewayPropertiesAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	// List of ExpressRoute connections to the ExpressRoute gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	ExpressRouteConnections []ExpressRouteConnectionType `pulumi:"expressRouteConnections"`
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName *string `pulumi:"expressRouteGatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubId `pulumi:"virtualHub"`
}

// The set of arguments for constructing a ExpressRouteGateway resource.
type ExpressRouteGatewayArgs struct {
	// Configures this gateway to accept traffic from non Virtual WAN networks.
	AllowNonVirtualWanTraffic pulumi.BoolPtrInput
	// Configuration for auto scaling.
	AutoScaleConfiguration ExpressRouteGatewayPropertiesAutoScaleConfigurationPtrInput
	// List of ExpressRoute connections to the ExpressRoute gateway.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	ExpressRouteConnections ExpressRouteConnectionTypeArrayInput
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubIdInput
}

func (ExpressRouteGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayArgs)(nil)).Elem()
}

type ExpressRouteGatewayInput interface {
	pulumi.Input

	ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput
	ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput
}

func (*ExpressRouteGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteGateway)(nil)).Elem()
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput {
	return i.ToExpressRouteGatewayOutputWithContext(context.Background())
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayOutput)
}

type ExpressRouteGatewayOutput struct{ *pulumi.OutputState }

func (ExpressRouteGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteGateway)(nil)).Elem()
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput {
	return o
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput {
	return o
}

// Configures this gateway to accept traffic from non Virtual WAN networks.
func (o ExpressRouteGatewayOutput) AllowNonVirtualWanTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.BoolPtrOutput { return v.AllowNonVirtualWanTraffic }).(pulumi.BoolPtrOutput)
}

// Configuration for auto scaling.
func (o ExpressRouteGatewayOutput) AutoScaleConfiguration() ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput {
		return v.AutoScaleConfiguration
	}).(ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput)
}

// The Azure API version of the resource.
func (o ExpressRouteGatewayOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ExpressRouteGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of ExpressRoute connections to the ExpressRoute gateway.
func (o ExpressRouteGatewayOutput) ExpressRouteConnections() ExpressRouteConnectionResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) ExpressRouteConnectionResponseArrayOutput {
		return v.ExpressRouteConnections
	}).(ExpressRouteConnectionResponseArrayOutput)
}

// Resource location.
func (o ExpressRouteGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ExpressRouteGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the express route gateway resource.
func (o ExpressRouteGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o ExpressRouteGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ExpressRouteGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
func (o ExpressRouteGatewayOutput) VirtualHub() VirtualHubIdResponseOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) VirtualHubIdResponseOutput { return v.VirtualHub }).(VirtualHubIdResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteGatewayOutput{})
}
