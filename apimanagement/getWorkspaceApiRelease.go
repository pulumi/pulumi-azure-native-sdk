// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the details of an API release.
//
// Uses Azure REST API version 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupWorkspaceApiRelease(ctx *pulumi.Context, args *LookupWorkspaceApiReleaseArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceApiReleaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceApiReleaseResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceApiRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceApiReleaseArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Release identifier within an API. Must be unique in the current API Management service instance.
	ReleaseId string `pulumi:"releaseId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// ApiRelease details.
type LookupWorkspaceApiReleaseResult struct {
	// Identifier of the API the release belongs to.
	ApiId *string `pulumi:"apiId"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
	CreatedDateTime string `pulumi:"createdDateTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Release Notes
	Notes *string `pulumi:"notes"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The time the API release was updated.
	UpdatedDateTime string `pulumi:"updatedDateTime"`
}

func LookupWorkspaceApiReleaseOutput(ctx *pulumi.Context, args LookupWorkspaceApiReleaseOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceApiReleaseResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceApiReleaseResultOutput, error) {
			args := v.(LookupWorkspaceApiReleaseArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getWorkspaceApiRelease", args, LookupWorkspaceApiReleaseResultOutput{}, options).(LookupWorkspaceApiReleaseResultOutput), nil
		}).(LookupWorkspaceApiReleaseResultOutput)
}

type LookupWorkspaceApiReleaseOutputArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Release identifier within an API. Must be unique in the current API Management service instance.
	ReleaseId pulumi.StringInput `pulumi:"releaseId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceApiReleaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiReleaseArgs)(nil)).Elem()
}

// ApiRelease details.
type LookupWorkspaceApiReleaseResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceApiReleaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiReleaseResult)(nil)).Elem()
}

func (o LookupWorkspaceApiReleaseResultOutput) ToLookupWorkspaceApiReleaseResultOutput() LookupWorkspaceApiReleaseResultOutput {
	return o
}

func (o LookupWorkspaceApiReleaseResultOutput) ToLookupWorkspaceApiReleaseResultOutputWithContext(ctx context.Context) LookupWorkspaceApiReleaseResultOutput {
	return o
}

// Identifier of the API the release belongs to.
func (o LookupWorkspaceApiReleaseResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupWorkspaceApiReleaseResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
func (o LookupWorkspaceApiReleaseResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceApiReleaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceApiReleaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Release Notes
func (o LookupWorkspaceApiReleaseResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceApiReleaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) string { return v.Type }).(pulumi.StringOutput)
}

// The time the API release was updated.
func (o LookupWorkspaceApiReleaseResultOutput) UpdatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiReleaseResult) string { return v.UpdatedDateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceApiReleaseResultOutput{})
}
