// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets cluster user credential of the managed cluster with a specified resource group and name.
func ListManagedClusterUserCredentials(ctx *pulumi.Context, args *ListManagedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterUserCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20190601:listManagedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterUserCredentialsArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// The list of credential result response.
type ListManagedClusterUserCredentialsResult struct {
	// Base64-encoded Kubernetes configuration file.
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterUserCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListManagedClusterUserCredentialsResultOutput, error) {
			args := v.(ListManagedClusterUserCredentialsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:containerservice/v20190601:listManagedClusterUserCredentials", args, ListManagedClusterUserCredentialsResultOutput{}, options).(ListManagedClusterUserCredentialsResultOutput), nil
		}).(ListManagedClusterUserCredentialsResultOutput)
}

type ListManagedClusterUserCredentialsOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListManagedClusterUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsArgs)(nil)).Elem()
}

// The list of credential result response.
type ListManagedClusterUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutput() ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterUserCredentialsResultOutput {
	return o
}

// Base64-encoded Kubernetes configuration file.
func (o ListManagedClusterUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterUserCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterUserCredentialsResultOutput{})
}
