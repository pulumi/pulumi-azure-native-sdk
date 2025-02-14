// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a VirtualHub.
func LookupVirtualHub(ctx *pulumi.Context, args *LookupVirtualHubArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHubResult
	err := ctx.Invoke("azure-native:network/v20200401:getVirtualHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubArgs struct {
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// VirtualHub Resource.
type LookupVirtualHubResult struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall *SubResourceResponse `pulumi:"azureFirewall"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway *SubResourceResponse `pulumi:"expressRouteGateway"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway *SubResourceResponse `pulumi:"p2SVpnGateway"`
	// The provisioning state of the virtual hub resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The routeTable associated with this virtual hub.
	RouteTable *VirtualHubRouteTableResponse `pulumi:"routeTable"`
	// The securityPartnerProvider associated with this VirtualHub.
	SecurityPartnerProvider *SubResourceResponse `pulumi:"securityPartnerProvider"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s []VirtualHubRouteTableV2Response `pulumi:"virtualHubRouteTableV2s"`
	// List of all vnet connections with this VirtualHub.
	VirtualNetworkConnections []HubVirtualNetworkConnectionResponse `pulumi:"virtualNetworkConnections"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan *SubResourceResponse `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway *SubResourceResponse `pulumi:"vpnGateway"`
}

func LookupVirtualHubOutput(ctx *pulumi.Context, args LookupVirtualHubOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubResultOutput, error) {
			args := v.(LookupVirtualHubArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network/v20200401:getVirtualHub", args, LookupVirtualHubResultOutput{}, options).(LookupVirtualHubResultOutput), nil
		}).(LookupVirtualHubResultOutput)
}

type LookupVirtualHubOutputArgs struct {
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubArgs)(nil)).Elem()
}

// VirtualHub Resource.
type LookupVirtualHubResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubResult)(nil)).Elem()
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutput() LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutputWithContext(ctx context.Context) LookupVirtualHubResultOutput {
	return o
}

// Address-prefix for this VirtualHub.
func (o LookupVirtualHubResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// The azureFirewall associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) AzureFirewall() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.AzureFirewall }).(SubResourceResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualHubResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The expressRouteGateway associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) ExpressRouteGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.ExpressRouteGateway }).(SubResourceResponsePtrOutput)
}

// Resource ID.
func (o LookupVirtualHubResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupVirtualHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupVirtualHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// The P2SVpnGateway associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) P2SVpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.P2SVpnGateway }).(SubResourceResponsePtrOutput)
}

// The provisioning state of the virtual hub resource.
func (o LookupVirtualHubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The routeTable associated with this virtual hub.
func (o LookupVirtualHubResultOutput) RouteTable() VirtualHubRouteTableResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *VirtualHubRouteTableResponse { return v.RouteTable }).(VirtualHubRouteTableResponsePtrOutput)
}

// The securityPartnerProvider associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) SecurityPartnerProvider() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.SecurityPartnerProvider }).(SubResourceResponsePtrOutput)
}

// The Security Provider name.
func (o LookupVirtualHubResultOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

// The sku of this VirtualHub.
func (o LookupVirtualHubResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupVirtualHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVirtualHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of all virtual hub route table v2s associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) VirtualHubRouteTableV2s() VirtualHubRouteTableV2ResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []VirtualHubRouteTableV2Response { return v.VirtualHubRouteTableV2s }).(VirtualHubRouteTableV2ResponseArrayOutput)
}

// List of all vnet connections with this VirtualHub.
func (o LookupVirtualHubResultOutput) VirtualNetworkConnections() HubVirtualNetworkConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []HubVirtualNetworkConnectionResponse {
		return v.VirtualNetworkConnections
	}).(HubVirtualNetworkConnectionResponseArrayOutput)
}

// The VirtualWAN to which the VirtualHub belongs.
func (o LookupVirtualHubResultOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

// The VpnGateway associated with this VirtualHub.
func (o LookupVirtualHubResultOutput) VpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VpnGateway }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubResultOutput{})
}
