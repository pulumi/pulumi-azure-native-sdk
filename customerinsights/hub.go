// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hub resource.
// API Version: 2017-04-26.
type Hub struct {
	pulumi.CustomResourceState

	// API endpoint URL of the hub.
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// Billing settings of the hub.
	HubBillingInfo HubBillingInfoFormatResponsePtrOutput `pulumi:"hubBillingInfo"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the hub.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
	TenantFeatures pulumi.IntPtrOutput `pulumi:"tenantFeatures"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Web endpoint URL of the hub.
	WebEndpoint pulumi.StringOutput `pulumi:"webEndpoint"`
}

// NewHub registers a new resource with the given unique name, arguments, and options.
func NewHub(ctx *pulumi.Context,
	name string, args *HubArgs, opts ...pulumi.ResourceOption) (*Hub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:Hub"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Hub"),
		},
	})
	opts = append(opts, aliases)
	var resource Hub
	err := ctx.RegisterResource("azure-native:customerinsights:Hub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHub gets an existing Hub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubState, opts ...pulumi.ResourceOption) (*Hub, error) {
	var resource Hub
	err := ctx.ReadResource("azure-native:customerinsights:Hub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hub resources.
type hubState struct {
}

type HubState struct {
}

func (HubState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubState)(nil)).Elem()
}

type hubArgs struct {
	// Billing settings of the hub.
	HubBillingInfo *HubBillingInfoFormat `pulumi:"hubBillingInfo"`
	// The name of the Hub.
	HubName *string `pulumi:"hubName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
	TenantFeatures *int `pulumi:"tenantFeatures"`
}

// The set of arguments for constructing a Hub resource.
type HubArgs struct {
	// Billing settings of the hub.
	HubBillingInfo HubBillingInfoFormatPtrInput
	// The name of the Hub.
	HubName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
	TenantFeatures pulumi.IntPtrInput
}

func (HubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubArgs)(nil)).Elem()
}

type HubInput interface {
	pulumi.Input

	ToHubOutput() HubOutput
	ToHubOutputWithContext(ctx context.Context) HubOutput
}

func (*Hub) ElementType() reflect.Type {
	return reflect.TypeOf((**Hub)(nil)).Elem()
}

func (i *Hub) ToHubOutput() HubOutput {
	return i.ToHubOutputWithContext(context.Background())
}

func (i *Hub) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubOutput)
}

type HubOutput struct{ *pulumi.OutputState }

func (HubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hub)(nil)).Elem()
}

func (o HubOutput) ToHubOutput() HubOutput {
	return o
}

func (o HubOutput) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return o
}

// API endpoint URL of the hub.
func (o HubOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// Billing settings of the hub.
func (o HubOutput) HubBillingInfo() HubBillingInfoFormatResponsePtrOutput {
	return o.ApplyT(func(v *Hub) HubBillingInfoFormatResponsePtrOutput { return v.HubBillingInfo }).(HubBillingInfoFormatResponsePtrOutput)
}

// Resource location.
func (o HubOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o HubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the hub.
func (o HubOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o HubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The bit flags for enabled hub features. Bit 0 is set to 1 indicates graph is enabled, or disabled if set to 0. Bit 1 is set to 1 indicates the hub is disabled, or enabled if set to 0.
func (o HubOutput) TenantFeatures() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Hub) pulumi.IntPtrOutput { return v.TenantFeatures }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o HubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Web endpoint URL of the hub.
func (o HubOutput) WebEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Hub) pulumi.StringOutput { return v.WebEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HubOutput{})
}