// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX DNS Zone
//
// Uses Azure REST API version 2022-05-01.
//
// Other available API versions: 2023-03-01, 2023-09-01.
func LookupWorkloadNetworkDnsZone(ctx *pulumi.Context, args *LookupWorkloadNetworkDnsZoneArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDnsZoneResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkloadNetworkDnsZoneResult
	err := ctx.Invoke("azure-native:avs:getWorkloadNetworkDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkDnsZoneArgs struct {
	// NSX DNS Zone identifier. Generally the same as the DNS Zone's display name
	DnsZoneId string `pulumi:"dnsZoneId"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// NSX DNS Zone
type LookupWorkloadNetworkDnsZoneResult struct {
	// Display name of the DNS Zone.
	DisplayName *string `pulumi:"displayName"`
	// DNS Server IP array of the DNS Zone.
	DnsServerIps []string `pulumi:"dnsServerIps"`
	// Number of DNS Services using the DNS zone.
	DnsServices *float64 `pulumi:"dnsServices"`
	// Domain names of the DNS Zone.
	Domain []string `pulumi:"domain"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// Source IP of the DNS Zone.
	SourceIp *string `pulumi:"sourceIp"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWorkloadNetworkDnsZoneOutput(ctx *pulumi.Context, args LookupWorkloadNetworkDnsZoneOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkDnsZoneResultOutput, error) {
			args := v.(LookupWorkloadNetworkDnsZoneArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:avs:getWorkloadNetworkDnsZone", args, LookupWorkloadNetworkDnsZoneResultOutput{}, options).(LookupWorkloadNetworkDnsZoneResultOutput), nil
		}).(LookupWorkloadNetworkDnsZoneResultOutput)
}

type LookupWorkloadNetworkDnsZoneOutputArgs struct {
	// NSX DNS Zone identifier. Generally the same as the DNS Zone's display name
	DnsZoneId pulumi.StringInput `pulumi:"dnsZoneId"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWorkloadNetworkDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDnsZoneArgs)(nil)).Elem()
}

// NSX DNS Zone
type LookupWorkloadNetworkDnsZoneResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkDnsZoneResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) ToLookupWorkloadNetworkDnsZoneResultOutput() LookupWorkloadNetworkDnsZoneResultOutput {
	return o
}

func (o LookupWorkloadNetworkDnsZoneResultOutput) ToLookupWorkloadNetworkDnsZoneResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkDnsZoneResultOutput {
	return o
}

// Display name of the DNS Zone.
func (o LookupWorkloadNetworkDnsZoneResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// DNS Server IP array of the DNS Zone.
func (o LookupWorkloadNetworkDnsZoneResultOutput) DnsServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) []string { return v.DnsServerIps }).(pulumi.StringArrayOutput)
}

// Number of DNS Services using the DNS zone.
func (o LookupWorkloadNetworkDnsZoneResultOutput) DnsServices() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *float64 { return v.DnsServices }).(pulumi.Float64PtrOutput)
}

// Domain names of the DNS Zone.
func (o LookupWorkloadNetworkDnsZoneResultOutput) Domain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) []string { return v.Domain }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupWorkloadNetworkDnsZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupWorkloadNetworkDnsZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o LookupWorkloadNetworkDnsZoneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o LookupWorkloadNetworkDnsZoneResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

// Source IP of the DNS Zone.
func (o LookupWorkloadNetworkDnsZoneResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWorkloadNetworkDnsZoneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkDnsZoneResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkDnsZoneResultOutput{})
}
