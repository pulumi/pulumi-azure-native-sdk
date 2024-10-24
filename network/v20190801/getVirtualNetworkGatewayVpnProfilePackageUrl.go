// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets pre-generated VPN profile for P2S client of the virtual network gateway in the specified resource group. The profile needs to be generated first using generateVpnProfile.
func GetVirtualNetworkGatewayVpnProfilePackageUrl(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnProfilePackageUrlArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnProfilePackageUrlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualNetworkGatewayVpnProfilePackageUrlResult
	err := ctx.Invoke("azure-native:network/v20190801:getVirtualNetworkGatewayVpnProfilePackageUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnProfilePackageUrlArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

type GetVirtualNetworkGatewayVpnProfilePackageUrlResult struct {
	Value *string `pulumi:"value"`
}

func GetVirtualNetworkGatewayVpnProfilePackageUrlOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayVpnProfilePackageUrlOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput, error) {
			args := v.(GetVirtualNetworkGatewayVpnProfilePackageUrlArgs)
			opts = utilities.PkgInvokeDefaultOpts(opts)
			var rv GetVirtualNetworkGatewayVpnProfilePackageUrlResult
			secret, err := ctx.InvokePackageRaw("azure-native:network/v20190801:getVirtualNetworkGatewayVpnProfilePackageUrl", args, &rv, "", opts...)
			if err != nil {
				return GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput), nil
			}
			return output, nil
		}).(GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput)
}

type GetVirtualNetworkGatewayVpnProfilePackageUrlOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayVpnProfilePackageUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnProfilePackageUrlArgs)(nil)).Elem()
}

type GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnProfilePackageUrlResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput) ToGetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput() GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput) ToGetVirtualNetworkGatewayVpnProfilePackageUrlResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnProfilePackageUrlResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayVpnProfilePackageUrlResultOutput{})
}
