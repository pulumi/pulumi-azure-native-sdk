// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package peering

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Peering Service
//
// Uses Azure REST API version 2022-10-01. In version 2.x of the Azure Native provider, it used API version 2022-10-01.
type PeeringService struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Log Analytics Workspace Properties
	LogAnalyticsWorkspaceProperties LogAnalyticsWorkspacePropertiesResponsePtrOutput `pulumi:"logAnalyticsWorkspaceProperties"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The location (state/province) of the customer.
	PeeringServiceLocation pulumi.StringPtrOutput `pulumi:"peeringServiceLocation"`
	// The name of the service provider.
	PeeringServiceProvider pulumi.StringPtrOutput `pulumi:"peeringServiceProvider"`
	// The backup peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderBackupPeeringLocation pulumi.StringPtrOutput `pulumi:"providerBackupPeeringLocation"`
	// The primary peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderPrimaryPeeringLocation pulumi.StringPtrOutput `pulumi:"providerPrimaryPeeringLocation"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU that defines the type of the peering service.
	Sku PeeringServiceSkuResponsePtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeeringService registers a new resource with the given unique name, arguments, and options.
func NewPeeringService(ctx *pulumi.Context,
	name string, args *PeeringServiceArgs, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20221001:PeeringService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PeeringService
	err := ctx.RegisterResource("azure-native:peering:PeeringService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringService gets an existing PeeringService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringServiceState, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	var resource PeeringService
	err := ctx.ReadResource("azure-native:peering:PeeringService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringService resources.
type peeringServiceState struct {
}

type PeeringServiceState struct {
}

func (PeeringServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServiceState)(nil)).Elem()
}

type peeringServiceArgs struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The location (state/province) of the customer.
	PeeringServiceLocation *string `pulumi:"peeringServiceLocation"`
	// The name of the peering service.
	PeeringServiceName *string `pulumi:"peeringServiceName"`
	// The name of the service provider.
	PeeringServiceProvider *string `pulumi:"peeringServiceProvider"`
	// The backup peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderBackupPeeringLocation *string `pulumi:"providerBackupPeeringLocation"`
	// The primary peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderPrimaryPeeringLocation *string `pulumi:"providerPrimaryPeeringLocation"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU that defines the type of the peering service.
	Sku *PeeringServiceSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PeeringService resource.
type PeeringServiceArgs struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The location (state/province) of the customer.
	PeeringServiceLocation pulumi.StringPtrInput
	// The name of the peering service.
	PeeringServiceName pulumi.StringPtrInput
	// The name of the service provider.
	PeeringServiceProvider pulumi.StringPtrInput
	// The backup peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderBackupPeeringLocation pulumi.StringPtrInput
	// The primary peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderPrimaryPeeringLocation pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SKU that defines the type of the peering service.
	Sku PeeringServiceSkuPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PeeringServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServiceArgs)(nil)).Elem()
}

type PeeringServiceInput interface {
	pulumi.Input

	ToPeeringServiceOutput() PeeringServiceOutput
	ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput
}

func (*PeeringService) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringService)(nil)).Elem()
}

func (i *PeeringService) ToPeeringServiceOutput() PeeringServiceOutput {
	return i.ToPeeringServiceOutputWithContext(context.Background())
}

func (i *PeeringService) ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceOutput)
}

type PeeringServiceOutput struct{ *pulumi.OutputState }

func (PeeringServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringService)(nil)).Elem()
}

func (o PeeringServiceOutput) ToPeeringServiceOutput() PeeringServiceOutput {
	return o
}

func (o PeeringServiceOutput) ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput {
	return o
}

// The Azure API version of the resource.
func (o PeeringServiceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The location of the resource.
func (o PeeringServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The Log Analytics Workspace Properties
func (o PeeringServiceOutput) LogAnalyticsWorkspaceProperties() LogAnalyticsWorkspacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PeeringService) LogAnalyticsWorkspacePropertiesResponsePtrOutput {
		return v.LogAnalyticsWorkspaceProperties
	}).(LogAnalyticsWorkspacePropertiesResponsePtrOutput)
}

// The name of the resource.
func (o PeeringServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The location (state/province) of the customer.
func (o PeeringServiceOutput) PeeringServiceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringPtrOutput { return v.PeeringServiceLocation }).(pulumi.StringPtrOutput)
}

// The name of the service provider.
func (o PeeringServiceOutput) PeeringServiceProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringPtrOutput { return v.PeeringServiceProvider }).(pulumi.StringPtrOutput)
}

// The backup peering (Microsoft/service provider) location to be used for customer traffic.
func (o PeeringServiceOutput) ProviderBackupPeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringPtrOutput { return v.ProviderBackupPeeringLocation }).(pulumi.StringPtrOutput)
}

// The primary peering (Microsoft/service provider) location to be used for customer traffic.
func (o PeeringServiceOutput) ProviderPrimaryPeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringPtrOutput { return v.ProviderPrimaryPeeringLocation }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o PeeringServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU that defines the type of the peering service.
func (o PeeringServiceOutput) Sku() PeeringServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v *PeeringService) PeeringServiceSkuResponsePtrOutput { return v.Sku }).(PeeringServiceSkuResponsePtrOutput)
}

// The resource tags.
func (o PeeringServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o PeeringServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeeringServiceOutput{})
}
