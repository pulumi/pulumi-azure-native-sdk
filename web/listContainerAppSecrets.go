// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App Secrets Collection ARM resource.
//
// Uses Azure REST API version 2023-01-01.
//
// Other available API versions: 2023-12-01.
func ListContainerAppSecrets(ctx *pulumi.Context, args *ListContainerAppSecretsArgs, opts ...pulumi.InvokeOption) (*ListContainerAppSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListContainerAppSecretsResult
	err := ctx.Invoke("azure-native:web:listContainerAppSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContainerAppSecretsArgs struct {
	// Name of the Container App.
	Name string `pulumi:"name"`
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
			return ctx.InvokeOutput("azure-native:web:listContainerAppSecrets", args, ListContainerAppSecretsResultOutput{}, options).(ListContainerAppSecretsResultOutput), nil
		}).(ListContainerAppSecretsResultOutput)
}

type ListContainerAppSecretsOutputArgs struct {
	// Name of the Container App.
	Name pulumi.StringInput `pulumi:"name"`
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
