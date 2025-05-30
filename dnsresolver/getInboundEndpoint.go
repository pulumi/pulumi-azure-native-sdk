// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnsresolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of an inbound endpoint for a DNS resolver.
//
// Uses Azure REST API version 2023-07-01-preview.
//
// Other available API versions: 2020-04-01-preview, 2022-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupInboundEndpoint(ctx *pulumi.Context, args *LookupInboundEndpointArgs, opts ...pulumi.InvokeOption) (*LookupInboundEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInboundEndpointResult
	err := ctx.Invoke("azure-native:dnsresolver:getInboundEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInboundEndpointArgs struct {
	// The name of the DNS resolver.
	DnsResolverName string `pulumi:"dnsResolverName"`
	// The name of the inbound endpoint for the DNS resolver.
	InboundEndpointName string `pulumi:"inboundEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes an inbound endpoint for a DNS resolver.
type LookupInboundEndpointResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// ETag of the inbound endpoint.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// IP configurations for the inbound endpoint.
	IpConfigurations []IpConfigurationResponse `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resourceGuid property of the inbound endpoint resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupInboundEndpointOutput(ctx *pulumi.Context, args LookupInboundEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupInboundEndpointResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInboundEndpointResultOutput, error) {
			args := v.(LookupInboundEndpointArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:dnsresolver:getInboundEndpoint", args, LookupInboundEndpointResultOutput{}, options).(LookupInboundEndpointResultOutput), nil
		}).(LookupInboundEndpointResultOutput)
}

type LookupInboundEndpointOutputArgs struct {
	// The name of the DNS resolver.
	DnsResolverName pulumi.StringInput `pulumi:"dnsResolverName"`
	// The name of the inbound endpoint for the DNS resolver.
	InboundEndpointName pulumi.StringInput `pulumi:"inboundEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInboundEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundEndpointArgs)(nil)).Elem()
}

// Describes an inbound endpoint for a DNS resolver.
type LookupInboundEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupInboundEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundEndpointResult)(nil)).Elem()
}

func (o LookupInboundEndpointResultOutput) ToLookupInboundEndpointResultOutput() LookupInboundEndpointResultOutput {
	return o
}

func (o LookupInboundEndpointResultOutput) ToLookupInboundEndpointResultOutputWithContext(ctx context.Context) LookupInboundEndpointResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupInboundEndpointResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// ETag of the inbound endpoint.
func (o LookupInboundEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupInboundEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP configurations for the inbound endpoint.
func (o LookupInboundEndpointResultOutput) IpConfigurations() IpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) []IpConfigurationResponse { return v.IpConfigurations }).(IpConfigurationResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupInboundEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupInboundEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupInboundEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the inbound endpoint resource.
func (o LookupInboundEndpointResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupInboundEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupInboundEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupInboundEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundEndpointResultOutput{})
}
