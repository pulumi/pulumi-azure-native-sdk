// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcontainerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the user credentials of the provisioned cluster (can only be used within private network)
//
// Uses Azure REST API version 2024-01-01.
//
// Other available API versions: 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcontainerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListProvisionedClusterInstanceUserKubeconfig(ctx *pulumi.Context, args *ListProvisionedClusterInstanceUserKubeconfigArgs, opts ...pulumi.InvokeOption) (*ListProvisionedClusterInstanceUserKubeconfigResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListProvisionedClusterInstanceUserKubeconfigResult
	err := ctx.Invoke("azure-native:hybridcontainerservice:listProvisionedClusterInstanceUserKubeconfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProvisionedClusterInstanceUserKubeconfigArgs struct {
	// The fully qualified Azure Resource Manager identifier of the connected cluster resource.
	ConnectedClusterResourceUri string `pulumi:"connectedClusterResourceUri"`
}

// The list kubeconfig result response.
type ListProvisionedClusterInstanceUserKubeconfigResult struct {
	Error *ListCredentialResponseResponseError `pulumi:"error"`
	// Operation Id
	Id string `pulumi:"id"`
	// Operation Name
	Name       string                                   `pulumi:"name"`
	Properties ListCredentialResponseResponseProperties `pulumi:"properties"`
	// ARM Resource Id of the provisioned cluster instance
	ResourceId string `pulumi:"resourceId"`
	// Provisioning state of the resource
	Status string `pulumi:"status"`
}

func ListProvisionedClusterInstanceUserKubeconfigOutput(ctx *pulumi.Context, args ListProvisionedClusterInstanceUserKubeconfigOutputArgs, opts ...pulumi.InvokeOption) ListProvisionedClusterInstanceUserKubeconfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListProvisionedClusterInstanceUserKubeconfigResultOutput, error) {
			args := v.(ListProvisionedClusterInstanceUserKubeconfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridcontainerservice:listProvisionedClusterInstanceUserKubeconfig", args, ListProvisionedClusterInstanceUserKubeconfigResultOutput{}, options).(ListProvisionedClusterInstanceUserKubeconfigResultOutput), nil
		}).(ListProvisionedClusterInstanceUserKubeconfigResultOutput)
}

type ListProvisionedClusterInstanceUserKubeconfigOutputArgs struct {
	// The fully qualified Azure Resource Manager identifier of the connected cluster resource.
	ConnectedClusterResourceUri pulumi.StringInput `pulumi:"connectedClusterResourceUri"`
}

func (ListProvisionedClusterInstanceUserKubeconfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProvisionedClusterInstanceUserKubeconfigArgs)(nil)).Elem()
}

// The list kubeconfig result response.
type ListProvisionedClusterInstanceUserKubeconfigResultOutput struct{ *pulumi.OutputState }

func (ListProvisionedClusterInstanceUserKubeconfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProvisionedClusterInstanceUserKubeconfigResult)(nil)).Elem()
}

func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) ToListProvisionedClusterInstanceUserKubeconfigResultOutput() ListProvisionedClusterInstanceUserKubeconfigResultOutput {
	return o
}

func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) ToListProvisionedClusterInstanceUserKubeconfigResultOutputWithContext(ctx context.Context) ListProvisionedClusterInstanceUserKubeconfigResultOutput {
	return o
}

func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) Error() ListCredentialResponseResponseErrorPtrOutput {
	return o.ApplyT(func(v ListProvisionedClusterInstanceUserKubeconfigResult) *ListCredentialResponseResponseError {
		return v.Error
	}).(ListCredentialResponseResponseErrorPtrOutput)
}

// Operation Id
func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListProvisionedClusterInstanceUserKubeconfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// Operation Name
func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListProvisionedClusterInstanceUserKubeconfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) Properties() ListCredentialResponseResponsePropertiesOutput {
	return o.ApplyT(func(v ListProvisionedClusterInstanceUserKubeconfigResult) ListCredentialResponseResponseProperties {
		return v.Properties
	}).(ListCredentialResponseResponsePropertiesOutput)
}

// ARM Resource Id of the provisioned cluster instance
func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ListProvisionedClusterInstanceUserKubeconfigResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// Provisioning state of the resource
func (o ListProvisionedClusterInstanceUserKubeconfigResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListProvisionedClusterInstanceUserKubeconfigResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProvisionedClusterInstanceUserKubeconfigResultOutput{})
}
