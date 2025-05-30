// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webpubsub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified private endpoint connection
//
// Uses Azure REST API version 2024-03-01.
//
// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native webpubsub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupWebPubSubPrivateEndpointConnection(ctx *pulumi.Context, args *LookupWebPubSubPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubPrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebPubSubPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:webpubsub:getWebPubSubPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubPrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection associated with the Azure resource.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A private endpoint connection to an azure resource
type LookupWebPubSubPrivateEndpointConnectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Group IDs
	GroupIds []string `pulumi:"groupIds"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Private endpoint
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Connection state of the private endpoint connection
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWebPubSubPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupWebPubSubPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubPrivateEndpointConnectionResultOutput, error) {
			args := v.(LookupWebPubSubPrivateEndpointConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:webpubsub:getWebPubSubPrivateEndpointConnection", args, LookupWebPubSubPrivateEndpointConnectionResultOutput{}, options).(LookupWebPubSubPrivateEndpointConnectionResultOutput), nil
		}).(LookupWebPubSubPrivateEndpointConnectionResultOutput)
}

type LookupWebPubSubPrivateEndpointConnectionOutputArgs struct {
	// The name of the private endpoint connection associated with the Azure resource.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubPrivateEndpointConnectionArgs)(nil)).Elem()
}

// A private endpoint connection to an azure resource
type LookupWebPubSubPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) ToLookupWebPubSubPrivateEndpointConnectionResultOutput() LookupWebPubSubPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) ToLookupWebPubSubPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupWebPubSubPrivateEndpointConnectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Group IDs
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// Connection state of the private endpoint connection
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// Provisioning state of the resource.
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubPrivateEndpointConnectionResultOutput{})
}
