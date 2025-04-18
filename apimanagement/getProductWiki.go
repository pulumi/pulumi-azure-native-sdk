// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Wiki for a Product specified by its identifier.
//
// Uses Azure REST API version 2022-09-01-preview.
//
// Other available API versions: 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupProductWiki(ctx *pulumi.Context, args *LookupProductWikiArgs, opts ...pulumi.InvokeOption) (*LookupProductWikiResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProductWikiResult
	err := ctx.Invoke("azure-native:apimanagement:getProductWiki", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductWikiArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Wiki properties
type LookupProductWikiResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Collection wiki documents included into this wiki.
	Documents []WikiDocumentationContractResponse `pulumi:"documents"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupProductWikiOutput(ctx *pulumi.Context, args LookupProductWikiOutputArgs, opts ...pulumi.InvokeOption) LookupProductWikiResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProductWikiResultOutput, error) {
			args := v.(LookupProductWikiArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getProductWiki", args, LookupProductWikiResultOutput{}, options).(LookupProductWikiResultOutput), nil
		}).(LookupProductWikiResultOutput)
}

type LookupProductWikiOutputArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupProductWikiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductWikiArgs)(nil)).Elem()
}

// Wiki properties
type LookupProductWikiResultOutput struct{ *pulumi.OutputState }

func (LookupProductWikiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductWikiResult)(nil)).Elem()
}

func (o LookupProductWikiResultOutput) ToLookupProductWikiResultOutput() LookupProductWikiResultOutput {
	return o
}

func (o LookupProductWikiResultOutput) ToLookupProductWikiResultOutputWithContext(ctx context.Context) LookupProductWikiResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupProductWikiResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductWikiResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Collection wiki documents included into this wiki.
func (o LookupProductWikiResultOutput) Documents() WikiDocumentationContractResponseArrayOutput {
	return o.ApplyT(func(v LookupProductWikiResult) []WikiDocumentationContractResponse { return v.Documents }).(WikiDocumentationContractResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupProductWikiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductWikiResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupProductWikiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductWikiResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupProductWikiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductWikiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProductWikiResultOutput{})
}
