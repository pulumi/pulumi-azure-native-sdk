// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App Secrets Collection ARM resource.
//
// Uses Azure REST API version 2024-03-01.
//
// Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListContainerAppSecrets(ctx *pulumi.Context, args *ListContainerAppSecretsArgs, opts ...pulumi.InvokeOption) (*ListContainerAppSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListContainerAppSecretsResult
	err := ctx.Invoke("azure-native:app:listContainerAppSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContainerAppSecretsArgs struct {
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Container App Secrets Collection ARM resource.
type ListContainerAppSecretsResult struct {
	// Collection of resources.
	Value []ContainerAppSecretResponse `pulumi:"value"`
}

func ListContainerAppSecretsOutput(ctx *pulumi.Context, args ListContainerAppSecretsOutputArgs, opts ...pulumi.InvokeOption) ListContainerAppSecretsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListContainerAppSecretsResultOutput, error) {
			args := v.(ListContainerAppSecretsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:app:listContainerAppSecrets", args, ListContainerAppSecretsResultOutput{}, options).(ListContainerAppSecretsResultOutput), nil
		}).(ListContainerAppSecretsResultOutput)
}

type ListContainerAppSecretsOutputArgs struct {
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListContainerAppSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppSecretsArgs)(nil)).Elem()
}

// Container App Secrets Collection ARM resource.
type ListContainerAppSecretsResultOutput struct{ *pulumi.OutputState }

func (ListContainerAppSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppSecretsResult)(nil)).Elem()
}

func (o ListContainerAppSecretsResultOutput) ToListContainerAppSecretsResultOutput() ListContainerAppSecretsResultOutput {
	return o
}

func (o ListContainerAppSecretsResultOutput) ToListContainerAppSecretsResultOutputWithContext(ctx context.Context) ListContainerAppSecretsResultOutput {
	return o
}

// Collection of resources.
func (o ListContainerAppSecretsResultOutput) Value() ContainerAppSecretResponseArrayOutput {
	return o.ApplyT(func(v ListContainerAppSecretsResult) []ContainerAppSecretResponse { return v.Value }).(ContainerAppSecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListContainerAppSecretsResultOutput{})
}
