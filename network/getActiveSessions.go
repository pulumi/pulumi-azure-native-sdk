// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the list of currently active sessions on the Bastion.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func GetActiveSessions(ctx *pulumi.Context, args *GetActiveSessionsArgs, opts ...pulumi.InvokeOption) (*GetActiveSessionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetActiveSessionsResult
	err := ctx.Invoke("azure-native:network:getActiveSessions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetActiveSessionsArgs struct {
	// The name of the Bastion Host.
	BastionHostName string `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for GetActiveSessions.
type GetActiveSessionsResult struct {
	// The URL to get the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// List of active sessions on the bastion.
	Value []BastionActiveSessionResponse `pulumi:"value"`
}

func GetActiveSessionsOutput(ctx *pulumi.Context, args GetActiveSessionsOutputArgs, opts ...pulumi.InvokeOption) GetActiveSessionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetActiveSessionsResultOutput, error) {
			args := v.(GetActiveSessionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getActiveSessions", args, GetActiveSessionsResultOutput{}, options).(GetActiveSessionsResultOutput), nil
		}).(GetActiveSessionsResultOutput)
}

type GetActiveSessionsOutputArgs struct {
	// The name of the Bastion Host.
	BastionHostName pulumi.StringInput `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetActiveSessionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveSessionsArgs)(nil)).Elem()
}

// Response for GetActiveSessions.
type GetActiveSessionsResultOutput struct{ *pulumi.OutputState }

func (GetActiveSessionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveSessionsResult)(nil)).Elem()
}

func (o GetActiveSessionsResultOutput) ToGetActiveSessionsResultOutput() GetActiveSessionsResultOutput {
	return o
}

func (o GetActiveSessionsResultOutput) ToGetActiveSessionsResultOutputWithContext(ctx context.Context) GetActiveSessionsResultOutput {
	return o
}

// The URL to get the next set of results.
func (o GetActiveSessionsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActiveSessionsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of active sessions on the bastion.
func (o GetActiveSessionsResultOutput) Value() BastionActiveSessionResponseArrayOutput {
	return o.ApplyT(func(v GetActiveSessionsResult) []BastionActiveSessionResponse { return v.Value }).(BastionActiveSessionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActiveSessionsResultOutput{})
}
