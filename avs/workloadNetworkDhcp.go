// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX DHCP
// API Version: 2020-07-17-preview.
type WorkloadNetworkDhcp struct {
	pulumi.CustomResourceState

	// Type of DHCP: SERVER or RELAY.
	DhcpType pulumi.StringOutput `pulumi:"dhcpType"`
	// Display name of the DHCP entity.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// NSX Segments consuming DHCP.
	Segments pulumi.StringArrayOutput `pulumi:"segments"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkDhcp registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkDhcp(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDhcpArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDhcp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DhcpType == nil {
		return nil, errors.New("invalid value for required argument 'DhcpType'")
	}
	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDhcp"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDhcp
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkDhcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkDhcp gets an existing WorkloadNetworkDhcp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkDhcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDhcpState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDhcp, error) {
	var resource WorkloadNetworkDhcp
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkDhcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkDhcp resources.
type workloadNetworkDhcpState struct {
}

type WorkloadNetworkDhcpState struct {
}

func (WorkloadNetworkDhcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDhcpState)(nil)).Elem()
}

type workloadNetworkDhcpArgs struct {
	// NSX DHCP identifier. Generally the same as the DHCP display name
	DhcpId *string `pulumi:"dhcpId"`
	// Type of DHCP: SERVER or RELAY.
	DhcpType string `pulumi:"dhcpType"`
	// Display name of the DHCP entity.
	DisplayName *string `pulumi:"displayName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
}

// The set of arguments for constructing a WorkloadNetworkDhcp resource.
type WorkloadNetworkDhcpArgs struct {
	// NSX DHCP identifier. Generally the same as the DHCP display name
	DhcpId pulumi.StringPtrInput
	// Type of DHCP: SERVER or RELAY.
	DhcpType pulumi.StringInput
	// Display name of the DHCP entity.
	DisplayName pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
}

func (WorkloadNetworkDhcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDhcpArgs)(nil)).Elem()
}

type WorkloadNetworkDhcpInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput
	ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput
}

func (*WorkloadNetworkDhcp) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDhcp)(nil)).Elem()
}

func (i *WorkloadNetworkDhcp) ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput {
	return i.ToWorkloadNetworkDhcpOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDhcp) ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpOutput)
}

type WorkloadNetworkDhcpOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDhcp)(nil)).Elem()
}

func (o WorkloadNetworkDhcpOutput) ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput {
	return o
}

func (o WorkloadNetworkDhcpOutput) ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput {
	return o
}

// Type of DHCP: SERVER or RELAY.
func (o WorkloadNetworkDhcpOutput) DhcpType() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringOutput { return v.DhcpType }).(pulumi.StringOutput)
}

// Display name of the DHCP entity.
func (o WorkloadNetworkDhcpOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o WorkloadNetworkDhcpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o WorkloadNetworkDhcpOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkDhcpOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// NSX Segments consuming DHCP.
func (o WorkloadNetworkDhcpOutput) Segments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringArrayOutput { return v.Segments }).(pulumi.StringArrayOutput)
}

// Resource type.
func (o WorkloadNetworkDhcpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDhcpOutput{})
}