// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX DNS Zone
type WorkloadNetworkDnsZone struct {
	pulumi.CustomResourceState

	// Display name of the DNS Zone.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// DNS Server IP array of the DNS Zone.
	DnsServerIps pulumi.StringArrayOutput `pulumi:"dnsServerIps"`
	// Number of DNS Services using the DNS zone.
	DnsServices pulumi.Float64PtrOutput `pulumi:"dnsServices"`
	// Domain names of the DNS Zone.
	Domain pulumi.StringArrayOutput `pulumi:"domain"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// Source IP of the DNS Zone.
	SourceIp pulumi.StringPtrOutput `pulumi:"sourceIp"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkDnsZone registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDnsZoneArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDnsZone"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDnsZone
	err := ctx.RegisterResource("azure-native:avs/v20210601:WorkloadNetworkDnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkDnsZone gets an existing WorkloadNetworkDnsZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsZoneState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
	var resource WorkloadNetworkDnsZone
	err := ctx.ReadResource("azure-native:avs/v20210601:WorkloadNetworkDnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkDnsZone resources.
type workloadNetworkDnsZoneState struct {
}

type WorkloadNetworkDnsZoneState struct {
}

func (WorkloadNetworkDnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsZoneState)(nil)).Elem()
}

type workloadNetworkDnsZoneArgs struct {
	// Display name of the DNS Zone.
	DisplayName *string `pulumi:"displayName"`
	// DNS Server IP array of the DNS Zone.
	DnsServerIps []string `pulumi:"dnsServerIps"`
	// Number of DNS Services using the DNS zone.
	DnsServices *float64 `pulumi:"dnsServices"`
	// NSX DNS Zone identifier. Generally the same as the DNS Zone's display name
	DnsZoneId *string `pulumi:"dnsZoneId"`
	// Domain names of the DNS Zone.
	Domain []string `pulumi:"domain"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// Source IP of the DNS Zone.
	SourceIp *string `pulumi:"sourceIp"`
}

// The set of arguments for constructing a WorkloadNetworkDnsZone resource.
type WorkloadNetworkDnsZoneArgs struct {
	// Display name of the DNS Zone.
	DisplayName pulumi.StringPtrInput
	// DNS Server IP array of the DNS Zone.
	DnsServerIps pulumi.StringArrayInput
	// Number of DNS Services using the DNS zone.
	DnsServices pulumi.Float64PtrInput
	// NSX DNS Zone identifier. Generally the same as the DNS Zone's display name
	DnsZoneId pulumi.StringPtrInput
	// Domain names of the DNS Zone.
	Domain pulumi.StringArrayInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// Source IP of the DNS Zone.
	SourceIp pulumi.StringPtrInput
}

func (WorkloadNetworkDnsZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsZoneArgs)(nil)).Elem()
}

type WorkloadNetworkDnsZoneInput interface {
	pulumi.Input

	ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput
	ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput
}

func (*WorkloadNetworkDnsZone) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsZone)(nil)).Elem()
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return i.ToWorkloadNetworkDnsZoneOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsZoneOutput)
}

type WorkloadNetworkDnsZoneOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsZone)(nil)).Elem()
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return o
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return o
}

// Display name of the DNS Zone.
func (o WorkloadNetworkDnsZoneOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// DNS Server IP array of the DNS Zone.
func (o WorkloadNetworkDnsZoneOutput) DnsServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringArrayOutput { return v.DnsServerIps }).(pulumi.StringArrayOutput)
}

// Number of DNS Services using the DNS zone.
func (o WorkloadNetworkDnsZoneOutput) DnsServices() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.Float64PtrOutput { return v.DnsServices }).(pulumi.Float64PtrOutput)
}

// Domain names of the DNS Zone.
func (o WorkloadNetworkDnsZoneOutput) Domain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringArrayOutput { return v.Domain }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o WorkloadNetworkDnsZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o WorkloadNetworkDnsZoneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkDnsZoneOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// Source IP of the DNS Zone.
func (o WorkloadNetworkDnsZoneOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringPtrOutput { return v.SourceIp }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WorkloadNetworkDnsZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsZone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDnsZoneOutput{})
}