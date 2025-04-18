// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package orbital

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of L2 Connections attached to an ground station.
//
// Uses Azure REST API version 2024-03-01-preview.
//
// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListGroundStationL2Connections(ctx *pulumi.Context, args *ListGroundStationL2ConnectionsArgs, opts ...pulumi.InvokeOption) (*ListGroundStationL2ConnectionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListGroundStationL2ConnectionsResult
	err := ctx.Invoke("azure-native:orbital:listGroundStationL2Connections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGroundStationL2ConnectionsArgs struct {
	// Ground Station name.
	GroundStationName string `pulumi:"groundStationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for an API service call that lists the resource IDs of resources associated with another resource.
type ListGroundStationL2ConnectionsResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// A list of Azure Resource IDs.
	Value []ResourceIdListResultResponseValue `pulumi:"value"`
}

func ListGroundStationL2ConnectionsOutput(ctx *pulumi.Context, args ListGroundStationL2ConnectionsOutputArgs, opts ...pulumi.InvokeOption) ListGroundStationL2ConnectionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListGroundStationL2ConnectionsResultOutput, error) {
			args := v.(ListGroundStationL2ConnectionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:orbital:listGroundStationL2Connections", args, ListGroundStationL2ConnectionsResultOutput{}, options).(ListGroundStationL2ConnectionsResultOutput), nil
		}).(ListGroundStationL2ConnectionsResultOutput)
}

type ListGroundStationL2ConnectionsOutputArgs struct {
	// Ground Station name.
	GroundStationName pulumi.StringInput `pulumi:"groundStationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListGroundStationL2ConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGroundStationL2ConnectionsArgs)(nil)).Elem()
}

// Response for an API service call that lists the resource IDs of resources associated with another resource.
type ListGroundStationL2ConnectionsResultOutput struct{ *pulumi.OutputState }

func (ListGroundStationL2ConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGroundStationL2ConnectionsResult)(nil)).Elem()
}

func (o ListGroundStationL2ConnectionsResultOutput) ToListGroundStationL2ConnectionsResultOutput() ListGroundStationL2ConnectionsResultOutput {
	return o
}

func (o ListGroundStationL2ConnectionsResultOutput) ToListGroundStationL2ConnectionsResultOutputWithContext(ctx context.Context) ListGroundStationL2ConnectionsResultOutput {
	return o
}

// The URL to get the next set of results.
func (o ListGroundStationL2ConnectionsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListGroundStationL2ConnectionsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// A list of Azure Resource IDs.
func (o ListGroundStationL2ConnectionsResultOutput) Value() ResourceIdListResultResponseValueArrayOutput {
	return o.ApplyT(func(v ListGroundStationL2ConnectionsResult) []ResourceIdListResultResponseValue { return v.Value }).(ResourceIdListResultResponseValueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGroundStationL2ConnectionsResultOutput{})
}
