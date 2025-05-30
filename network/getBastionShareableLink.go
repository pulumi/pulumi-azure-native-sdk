// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Return the Bastion Shareable Links for all the VMs specified in the request.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func GetBastionShareableLink(ctx *pulumi.Context, args *GetBastionShareableLinkArgs, opts ...pulumi.InvokeOption) (*GetBastionShareableLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetBastionShareableLinkResult
	err := ctx.Invoke("azure-native:network:getBastionShareableLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBastionShareableLinkArgs struct {
	// The name of the Bastion Host.
	BastionHostName string `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of VM references.
	Vms []BastionShareableLink `pulumi:"vms"`
}

// Response for all the Bastion Shareable Link endpoints.
type GetBastionShareableLinkResult struct {
	// The URL to get the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// List of Bastion Shareable Links for the request.
	Value []BastionShareableLinkResponse `pulumi:"value"`
}

func GetBastionShareableLinkOutput(ctx *pulumi.Context, args GetBastionShareableLinkOutputArgs, opts ...pulumi.InvokeOption) GetBastionShareableLinkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBastionShareableLinkResultOutput, error) {
			args := v.(GetBastionShareableLinkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getBastionShareableLink", args, GetBastionShareableLinkResultOutput{}, options).(GetBastionShareableLinkResultOutput), nil
		}).(GetBastionShareableLinkResultOutput)
}

type GetBastionShareableLinkOutputArgs struct {
	// The name of the Bastion Host.
	BastionHostName pulumi.StringInput `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// List of VM references.
	Vms BastionShareableLinkArrayInput `pulumi:"vms"`
}

func (GetBastionShareableLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionShareableLinkArgs)(nil)).Elem()
}

// Response for all the Bastion Shareable Link endpoints.
type GetBastionShareableLinkResultOutput struct{ *pulumi.OutputState }

func (GetBastionShareableLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionShareableLinkResult)(nil)).Elem()
}

func (o GetBastionShareableLinkResultOutput) ToGetBastionShareableLinkResultOutput() GetBastionShareableLinkResultOutput {
	return o
}

func (o GetBastionShareableLinkResultOutput) ToGetBastionShareableLinkResultOutputWithContext(ctx context.Context) GetBastionShareableLinkResultOutput {
	return o
}

// The URL to get the next set of results.
func (o GetBastionShareableLinkResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBastionShareableLinkResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of Bastion Shareable Links for the request.
func (o GetBastionShareableLinkResultOutput) Value() BastionShareableLinkResponseArrayOutput {
	return o.ApplyT(func(v GetBastionShareableLinkResult) []BastionShareableLinkResponse { return v.Value }).(BastionShareableLinkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBastionShareableLinkResultOutput{})
}
