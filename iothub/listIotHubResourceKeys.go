// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iothub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the security metadata for an IoT hub. For more information, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-security.
//
// Uses Azure REST API version 2023-06-30.
//
// Other available API versions: 2016-02-03, 2017-01-19, 2017-07-01, 2018-01-22, 2018-04-01, 2018-12-01-preview, 2019-03-22, 2019-03-22-preview, 2019-07-01-preview, 2019-11-04, 2020-03-01, 2020-04-01, 2020-06-15, 2020-07-10-preview, 2020-08-01, 2020-08-31, 2020-08-31-preview, 2021-02-01-preview, 2021-03-03-preview, 2021-03-31, 2021-07-01, 2021-07-01-preview, 2021-07-02, 2021-07-02-preview, 2022-04-30-preview, 2022-11-15-preview, 2023-06-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iothub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListIotHubResourceKeys(ctx *pulumi.Context, args *ListIotHubResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListIotHubResourceKeysResult
	err := ctx.Invoke("azure-native:iothub:listIotHubResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The list of shared access policies with a next link.
type ListIotHubResourceKeysResult struct {
	// The next link.
	NextLink string `pulumi:"nextLink"`
	// The list of shared access policies.
	Value []SharedAccessSignatureAuthorizationRuleResponse `pulumi:"value"`
}

func ListIotHubResourceKeysOutput(ctx *pulumi.Context, args ListIotHubResourceKeysOutputArgs, opts ...pulumi.InvokeOption) ListIotHubResourceKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListIotHubResourceKeysResultOutput, error) {
			args := v.(ListIotHubResourceKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:iothub:listIotHubResourceKeys", args, ListIotHubResourceKeysResultOutput{}, options).(ListIotHubResourceKeysResultOutput), nil
		}).(ListIotHubResourceKeysResultOutput)
}

type ListIotHubResourceKeysOutputArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListIotHubResourceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysArgs)(nil)).Elem()
}

// The list of shared access policies with a next link.
type ListIotHubResourceKeysResultOutput struct{ *pulumi.OutputState }

func (ListIotHubResourceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysResult)(nil)).Elem()
}

func (o ListIotHubResourceKeysResultOutput) ToListIotHubResourceKeysResultOutput() ListIotHubResourceKeysResultOutput {
	return o
}

func (o ListIotHubResourceKeysResultOutput) ToListIotHubResourceKeysResultOutputWithContext(ctx context.Context) ListIotHubResourceKeysResultOutput {
	return o
}

// The next link.
func (o ListIotHubResourceKeysResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// The list of shared access policies.
func (o ListIotHubResourceKeysResultOutput) Value() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysResult) []SharedAccessSignatureAuthorizationRuleResponse { return v.Value }).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotHubResourceKeysResultOutput{})
}
