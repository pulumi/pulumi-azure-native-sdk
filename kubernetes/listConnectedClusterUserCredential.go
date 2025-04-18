// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets cluster user credentials of the connected cluster with a specified resource group and name.
//
// Uses Azure REST API version 2024-02-01-preview.
//
// Other available API versions: 2021-10-01, 2022-05-01-preview, 2022-10-01-preview, 2023-11-01-preview, 2024-01-01, 2024-06-01-preview, 2024-07-01-preview, 2024-07-15-preview, 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kubernetes [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListConnectedClusterUserCredential(ctx *pulumi.Context, args *ListConnectedClusterUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListConnectedClusterUserCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConnectedClusterUserCredentialResult
	err := ctx.Invoke("azure-native:kubernetes:listConnectedClusterUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedClusterUserCredentialArgs struct {
	// The mode of client authentication.
	AuthenticationMethod string `pulumi:"authenticationMethod"`
	// Boolean value to indicate whether the request is for client side proxy or not
	ClientProxy bool `pulumi:"clientProxy"`
	// The name of the Kubernetes cluster on which get is called.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The list of credential result response.
type ListConnectedClusterUserCredentialResult struct {
	// Contains the REP (rendezvous endpoint) and “Sender” access token.
	HybridConnectionConfig HybridConnectionConfigResponse `pulumi:"hybridConnectionConfig"`
	// Base64-encoded Kubernetes configuration file.
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListConnectedClusterUserCredentialOutput(ctx *pulumi.Context, args ListConnectedClusterUserCredentialOutputArgs, opts ...pulumi.InvokeOption) ListConnectedClusterUserCredentialResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListConnectedClusterUserCredentialResultOutput, error) {
			args := v.(ListConnectedClusterUserCredentialArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:kubernetes:listConnectedClusterUserCredential", args, ListConnectedClusterUserCredentialResultOutput{}, options).(ListConnectedClusterUserCredentialResultOutput), nil
		}).(ListConnectedClusterUserCredentialResultOutput)
}

type ListConnectedClusterUserCredentialOutputArgs struct {
	// The mode of client authentication.
	AuthenticationMethod pulumi.StringInput `pulumi:"authenticationMethod"`
	// Boolean value to indicate whether the request is for client side proxy or not
	ClientProxy pulumi.BoolInput `pulumi:"clientProxy"`
	// The name of the Kubernetes cluster on which get is called.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedClusterUserCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialArgs)(nil)).Elem()
}

// The list of credential result response.
type ListConnectedClusterUserCredentialResultOutput struct{ *pulumi.OutputState }

func (ListConnectedClusterUserCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialResult)(nil)).Elem()
}

func (o ListConnectedClusterUserCredentialResultOutput) ToListConnectedClusterUserCredentialResultOutput() ListConnectedClusterUserCredentialResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialResultOutput) ToListConnectedClusterUserCredentialResultOutputWithContext(ctx context.Context) ListConnectedClusterUserCredentialResultOutput {
	return o
}

// Contains the REP (rendezvous endpoint) and “Sender” access token.
func (o ListConnectedClusterUserCredentialResultOutput) HybridConnectionConfig() HybridConnectionConfigResponseOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialResult) HybridConnectionConfigResponse {
		return v.HybridConnectionConfig
	}).(HybridConnectionConfigResponseOutput)
}

// Base64-encoded Kubernetes configuration file.
func (o ListConnectedClusterUserCredentialResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedClusterUserCredentialResultOutput{})
}
