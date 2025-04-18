// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List sensitive environment variables of the Job.
//
// Uses Azure REST API version 2024-05-01-preview.
func ListJobEnvSecrets(ctx *pulumi.Context, args *ListJobEnvSecretsArgs, opts ...pulumi.InvokeOption) (*ListJobEnvSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListJobEnvSecretsResult
	err := ctx.Invoke("azure-native:appplatform:listJobEnvSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListJobEnvSecretsArgs struct {
	// The name of the Job resource.
	JobName string `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Secret environment variable collection.
type ListJobEnvSecretsResult struct {
	// Collection of resources.
	Value []SecretResponse `pulumi:"value"`
}

func ListJobEnvSecretsOutput(ctx *pulumi.Context, args ListJobEnvSecretsOutputArgs, opts ...pulumi.InvokeOption) ListJobEnvSecretsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListJobEnvSecretsResultOutput, error) {
			args := v.(ListJobEnvSecretsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:appplatform:listJobEnvSecrets", args, ListJobEnvSecretsResultOutput{}, options).(ListJobEnvSecretsResultOutput), nil
		}).(ListJobEnvSecretsResultOutput)
}

type ListJobEnvSecretsOutputArgs struct {
	// The name of the Job resource.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListJobEnvSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobEnvSecretsArgs)(nil)).Elem()
}

// Secret environment variable collection.
type ListJobEnvSecretsResultOutput struct{ *pulumi.OutputState }

func (ListJobEnvSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobEnvSecretsResult)(nil)).Elem()
}

func (o ListJobEnvSecretsResultOutput) ToListJobEnvSecretsResultOutput() ListJobEnvSecretsResultOutput {
	return o
}

func (o ListJobEnvSecretsResultOutput) ToListJobEnvSecretsResultOutputWithContext(ctx context.Context) ListJobEnvSecretsResultOutput {
	return o
}

// Collection of resources.
func (o ListJobEnvSecretsResultOutput) Value() SecretResponseArrayOutput {
	return o.ApplyT(func(v ListJobEnvSecretsResult) []SecretResponse { return v.Value }).(SecretResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListJobEnvSecretsResultOutput{})
}
