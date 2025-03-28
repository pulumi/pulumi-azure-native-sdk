// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation retrieves a list of routes the virtual network gateway is advertising to the specified peer.
func GetVirtualNetworkGatewayAdvertisedRoutes(ctx *pulumi.Context, args *GetVirtualNetworkGatewayAdvertisedRoutesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayAdvertisedRoutesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualNetworkGatewayAdvertisedRoutesResult
	err := ctx.Invoke("azure-native:network/v20230901:getVirtualNetworkGatewayAdvertisedRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayAdvertisedRoutesArgs struct {
	// The IP address of the peer.
	Peer string `pulumi:"peer"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

// List of virtual network gateway routes.
type GetVirtualNetworkGatewayAdvertisedRoutesResult struct {
	// List of gateway routes.
	Value []GatewayRouteResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayAdvertisedRoutesOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayAdvertisedRoutesOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayAdvertisedRoutesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayAdvertisedRoutesResultOutput, error) {
			args := v.(GetVirtualNetworkGatewayAdvertisedRoutesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network/v20230901:getVirtualNetworkGatewayAdvertisedRoutes", args, GetVirtualNetworkGatewayAdvertisedRoutesResultOutput{}, options).(GetVirtualNetworkGatewayAdvertisedRoutesResultOutput), nil
		}).(GetVirtualNetworkGatewayAdvertisedRoutesResultOutput)
}

type GetVirtualNetworkGatewayAdvertisedRoutesOutputArgs struct {
	// The IP address of the peer.
	Peer pulumi.StringInput `pulumi:"peer"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayAdvertisedRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayAdvertisedRoutesArgs)(nil)).Elem()
}

// List of virtual network gateway routes.
type GetVirtualNetworkGatewayAdvertisedRoutesResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayAdvertisedRoutesResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) ToGetVirtualNetworkGatewayAdvertisedRoutesResultOutput() GetVirtualNetworkGatewayAdvertisedRoutesResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) ToGetVirtualNetworkGatewayAdvertisedRoutesResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayAdvertisedRoutesResultOutput {
	return o
}

// List of gateway routes.
func (o GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) Value() GatewayRouteResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayAdvertisedRoutesResult) []GatewayRouteResponse { return v.Value }).(GatewayRouteResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayAdvertisedRoutesResultOutput{})
}
