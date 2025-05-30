// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredatatransfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the destination endpoint ports for the specified flow and stream ID.
//
// Uses Azure REST API version 2024-09-27.
//
// Other available API versions: 2025-03-01-preview, 2025-04-11-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuredatatransfer [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func GetFlowDestinationEndpointPorts(ctx *pulumi.Context, args *GetFlowDestinationEndpointPortsArgs, opts ...pulumi.InvokeOption) (*GetFlowDestinationEndpointPortsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetFlowDestinationEndpointPortsResult
	err := ctx.Invoke("azure-native:azuredatatransfer:getFlowDestinationEndpointPorts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFlowDestinationEndpointPortsArgs struct {
	// The name for the connection that is to be requested.
	ConnectionName string `pulumi:"connectionName"`
	// The name for the flow that is to be onboarded.
	FlowName string `pulumi:"flowName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of destination endpoint ports for the flow stream
type GetFlowDestinationEndpointPortsResult struct {
	// The destination endpoint port for the flow stream
	Ports []float64 `pulumi:"ports"`
}

func GetFlowDestinationEndpointPortsOutput(ctx *pulumi.Context, args GetFlowDestinationEndpointPortsOutputArgs, opts ...pulumi.InvokeOption) GetFlowDestinationEndpointPortsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFlowDestinationEndpointPortsResultOutput, error) {
			args := v.(GetFlowDestinationEndpointPortsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azuredatatransfer:getFlowDestinationEndpointPorts", args, GetFlowDestinationEndpointPortsResultOutput{}, options).(GetFlowDestinationEndpointPortsResultOutput), nil
		}).(GetFlowDestinationEndpointPortsResultOutput)
}

type GetFlowDestinationEndpointPortsOutputArgs struct {
	// The name for the connection that is to be requested.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name for the flow that is to be onboarded.
	FlowName pulumi.StringInput `pulumi:"flowName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetFlowDestinationEndpointPortsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowDestinationEndpointPortsArgs)(nil)).Elem()
}

// List of destination endpoint ports for the flow stream
type GetFlowDestinationEndpointPortsResultOutput struct{ *pulumi.OutputState }

func (GetFlowDestinationEndpointPortsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowDestinationEndpointPortsResult)(nil)).Elem()
}

func (o GetFlowDestinationEndpointPortsResultOutput) ToGetFlowDestinationEndpointPortsResultOutput() GetFlowDestinationEndpointPortsResultOutput {
	return o
}

func (o GetFlowDestinationEndpointPortsResultOutput) ToGetFlowDestinationEndpointPortsResultOutputWithContext(ctx context.Context) GetFlowDestinationEndpointPortsResultOutput {
	return o
}

// The destination endpoint port for the flow stream
func (o GetFlowDestinationEndpointPortsResultOutput) Ports() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v GetFlowDestinationEndpointPortsResult) []float64 { return v.Ports }).(pulumi.Float64ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlowDestinationEndpointPortsResultOutput{})
}
