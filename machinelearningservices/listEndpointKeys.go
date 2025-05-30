// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2025-01-01-preview.
//
// Other available API versions: 2024-01-01-preview, 2024-07-01-preview, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListEndpointKeys(ctx *pulumi.Context, args *ListEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListEndpointKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices:listEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEndpointKeysArgs struct {
	// Name of the endpoint resource.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListEndpointKeysResult struct {
	// Dictionary of Keys for the endpoint.
	Keys *AccountApiKeysResponse `pulumi:"keys"`
}

func ListEndpointKeysOutput(ctx *pulumi.Context, args ListEndpointKeysOutputArgs, opts ...pulumi.InvokeOption) ListEndpointKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListEndpointKeysResultOutput, error) {
			args := v.(ListEndpointKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices:listEndpointKeys", args, ListEndpointKeysResultOutput{}, options).(ListEndpointKeysResultOutput), nil
		}).(ListEndpointKeysResultOutput)
}

type ListEndpointKeysOutputArgs struct {
	// Name of the endpoint resource.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListEndpointKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointKeysArgs)(nil)).Elem()
}

type ListEndpointKeysResultOutput struct{ *pulumi.OutputState }

func (ListEndpointKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointKeysResult)(nil)).Elem()
}

func (o ListEndpointKeysResultOutput) ToListEndpointKeysResultOutput() ListEndpointKeysResultOutput {
	return o
}

func (o ListEndpointKeysResultOutput) ToListEndpointKeysResultOutputWithContext(ctx context.Context) ListEndpointKeysResultOutput {
	return o
}

// Dictionary of Keys for the endpoint.
func (o ListEndpointKeysResultOutput) Keys() AccountApiKeysResponsePtrOutput {
	return o.ApplyT(func(v ListEndpointKeysResult) *AccountApiKeysResponse { return v.Keys }).(AccountApiKeysResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEndpointKeysResultOutput{})
}
