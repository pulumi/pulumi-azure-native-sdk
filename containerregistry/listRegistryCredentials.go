// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the login credentials for the specified container registry.
//
// Uses Azure REST API version 2024-11-01-preview.
//
// Other available API versions: 2019-12-01-preview, 2020-11-01-preview, 2021-06-01-preview, 2021-08-01-preview, 2021-09-01, 2021-12-01-preview, 2022-02-01-preview, 2022-12-01, 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2025-03-01-preview, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListRegistryCredentials(ctx *pulumi.Context, args *ListRegistryCredentialsArgs, opts ...pulumi.InvokeOption) (*ListRegistryCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListRegistryCredentialsResult
	err := ctx.Invoke("azure-native:containerregistry:listRegistryCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryCredentialsArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The response from the ListCredentials operation.
type ListRegistryCredentialsResult struct {
	// The list of passwords for a container registry.
	Passwords []RegistryPasswordResponse `pulumi:"passwords"`
	// The username for a container registry.
	Username *string `pulumi:"username"`
}

func ListRegistryCredentialsOutput(ctx *pulumi.Context, args ListRegistryCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListRegistryCredentialsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListRegistryCredentialsResultOutput, error) {
			args := v.(ListRegistryCredentialsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:containerregistry:listRegistryCredentials", args, ListRegistryCredentialsResultOutput{}, options).(ListRegistryCredentialsResultOutput), nil
		}).(ListRegistryCredentialsResultOutput)
}

type ListRegistryCredentialsOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRegistryCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryCredentialsArgs)(nil)).Elem()
}

// The response from the ListCredentials operation.
type ListRegistryCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListRegistryCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryCredentialsResult)(nil)).Elem()
}

func (o ListRegistryCredentialsResultOutput) ToListRegistryCredentialsResultOutput() ListRegistryCredentialsResultOutput {
	return o
}

func (o ListRegistryCredentialsResultOutput) ToListRegistryCredentialsResultOutputWithContext(ctx context.Context) ListRegistryCredentialsResultOutput {
	return o
}

// The list of passwords for a container registry.
func (o ListRegistryCredentialsResultOutput) Passwords() RegistryPasswordResponseArrayOutput {
	return o.ApplyT(func(v ListRegistryCredentialsResult) []RegistryPasswordResponse { return v.Passwords }).(RegistryPasswordResponseArrayOutput)
}

// The username for a container registry.
func (o ListRegistryCredentialsResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRegistryCredentialsResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryCredentialsResultOutput{})
}
