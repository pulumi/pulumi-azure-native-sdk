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
// Other available API versions: 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListConnectionAllModels(ctx *pulumi.Context, args *ListConnectionAllModelsArgs, opts ...pulumi.InvokeOption) (*ListConnectionAllModelsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConnectionAllModelsResult
	err := ctx.Invoke("azure-native:machinelearningservices:listConnectionAllModels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionAllModelsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListConnectionAllModelsResult struct {
	// The link to the next page constructed using the continuationToken.  If null, there are no additional pages.
	NextLink *string `pulumi:"nextLink"`
	// List of models.
	Value []EndpointModelPropertiesResponse `pulumi:"value"`
}

func ListConnectionAllModelsOutput(ctx *pulumi.Context, args ListConnectionAllModelsOutputArgs, opts ...pulumi.InvokeOption) ListConnectionAllModelsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListConnectionAllModelsResultOutput, error) {
			args := v.(ListConnectionAllModelsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices:listConnectionAllModels", args, ListConnectionAllModelsResultOutput{}, options).(ListConnectionAllModelsResultOutput), nil
		}).(ListConnectionAllModelsResultOutput)
}

type ListConnectionAllModelsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListConnectionAllModelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionAllModelsArgs)(nil)).Elem()
}

type ListConnectionAllModelsResultOutput struct{ *pulumi.OutputState }

func (ListConnectionAllModelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionAllModelsResult)(nil)).Elem()
}

func (o ListConnectionAllModelsResultOutput) ToListConnectionAllModelsResultOutput() ListConnectionAllModelsResultOutput {
	return o
}

func (o ListConnectionAllModelsResultOutput) ToListConnectionAllModelsResultOutputWithContext(ctx context.Context) ListConnectionAllModelsResultOutput {
	return o
}

// The link to the next page constructed using the continuationToken.  If null, there are no additional pages.
func (o ListConnectionAllModelsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConnectionAllModelsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of models.
func (o ListConnectionAllModelsResultOutput) Value() EndpointModelPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ListConnectionAllModelsResult) []EndpointModelPropertiesResponse { return v.Value }).(EndpointModelPropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectionAllModelsResultOutput{})
}
