// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of language extensions that can run within KQL queries.
//
// Uses Azure REST API version 2022-12-29.
//
// Other available API versions: 2022-07-07, 2023-05-02, 2023-08-15, 2024-04-13.
func ListClusterLanguageExtensions(ctx *pulumi.Context, args *ListClusterLanguageExtensionsArgs, opts ...pulumi.InvokeOption) (*ListClusterLanguageExtensionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListClusterLanguageExtensionsResult
	err := ctx.Invoke("azure-native:kusto:listClusterLanguageExtensions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterLanguageExtensionsArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The list of language extension objects.
type ListClusterLanguageExtensionsResult struct {
	// The list of language extensions.
	Value []LanguageExtensionResponse `pulumi:"value"`
}

func ListClusterLanguageExtensionsOutput(ctx *pulumi.Context, args ListClusterLanguageExtensionsOutputArgs, opts ...pulumi.InvokeOption) ListClusterLanguageExtensionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListClusterLanguageExtensionsResultOutput, error) {
			args := v.(ListClusterLanguageExtensionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:kusto:listClusterLanguageExtensions", args, ListClusterLanguageExtensionsResultOutput{}, options).(ListClusterLanguageExtensionsResultOutput), nil
		}).(ListClusterLanguageExtensionsResultOutput)
}

type ListClusterLanguageExtensionsOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterLanguageExtensionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterLanguageExtensionsArgs)(nil)).Elem()
}

// The list of language extension objects.
type ListClusterLanguageExtensionsResultOutput struct{ *pulumi.OutputState }

func (ListClusterLanguageExtensionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterLanguageExtensionsResult)(nil)).Elem()
}

func (o ListClusterLanguageExtensionsResultOutput) ToListClusterLanguageExtensionsResultOutput() ListClusterLanguageExtensionsResultOutput {
	return o
}

func (o ListClusterLanguageExtensionsResultOutput) ToListClusterLanguageExtensionsResultOutputWithContext(ctx context.Context) ListClusterLanguageExtensionsResultOutput {
	return o
}

// The list of language extensions.
func (o ListClusterLanguageExtensionsResultOutput) Value() LanguageExtensionResponseArrayOutput {
	return o.ApplyT(func(v ListClusterLanguageExtensionsResult) []LanguageExtensionResponse { return v.Value }).(LanguageExtensionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterLanguageExtensionsResultOutput{})
}
