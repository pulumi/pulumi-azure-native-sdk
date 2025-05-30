// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an API Management gateway resource description.
//
// Uses Azure REST API version 2024-06-01-preview.
//
// Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupApiGateway(ctx *pulumi.Context, args *LookupApiGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApiGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiGatewayResult
	err := ctx.Invoke("azure-native:apimanagement:getApiGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiGatewayArgs struct {
	// The name of the API Management gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A single API Management gateway resource in List or Get response.
type LookupApiGatewayResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Information regarding how the gateway should integrate with backend systems.
	Backend *BackendConfigurationResponse `pulumi:"backend"`
	// Information regarding the Configuration API of the API Management gateway. This is only applicable for API gateway with Standard SKU.
	ConfigurationApi *GatewayConfigurationApiResponse `pulumi:"configurationApi"`
	// Creation UTC date of the API Management gateway.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedAtUtc string `pulumi:"createdAtUtc"`
	// ETag of the resource.
	Etag string `pulumi:"etag"`
	// Information regarding how the gateway should be exposed.
	Frontend *FrontendConfigurationResponse `pulumi:"frontend"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The current provisioning state of the API Management gateway which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	ProvisioningState string `pulumi:"provisioningState"`
	// SKU properties of the API Management gateway.
	Sku ApiManagementGatewaySkuPropertiesResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The provisioning state of the API Management gateway, which is targeted by the long running operation started on the gateway.
	TargetProvisioningState string `pulumi:"targetProvisioningState"`
	// Resource type for API Management resource is set to Microsoft.ApiManagement.
	Type string `pulumi:"type"`
	// The type of VPN in which API Management gateway needs to be configured in.
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

func LookupApiGatewayOutput(ctx *pulumi.Context, args LookupApiGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupApiGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApiGatewayResultOutput, error) {
			args := v.(LookupApiGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getApiGateway", args, LookupApiGatewayResultOutput{}, options).(LookupApiGatewayResultOutput), nil
		}).(LookupApiGatewayResultOutput)
}

type LookupApiGatewayOutputArgs struct {
	// The name of the API Management gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApiGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayArgs)(nil)).Elem()
}

// A single API Management gateway resource in List or Get response.
type LookupApiGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupApiGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayResult)(nil)).Elem()
}

func (o LookupApiGatewayResultOutput) ToLookupApiGatewayResultOutput() LookupApiGatewayResultOutput {
	return o
}

func (o LookupApiGatewayResultOutput) ToLookupApiGatewayResultOutputWithContext(ctx context.Context) LookupApiGatewayResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupApiGatewayResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Information regarding how the gateway should integrate with backend systems.
func (o LookupApiGatewayResultOutput) Backend() BackendConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *BackendConfigurationResponse { return v.Backend }).(BackendConfigurationResponsePtrOutput)
}

// Information regarding the Configuration API of the API Management gateway. This is only applicable for API gateway with Standard SKU.
func (o LookupApiGatewayResultOutput) ConfigurationApi() GatewayConfigurationApiResponsePtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *GatewayConfigurationApiResponse { return v.ConfigurationApi }).(GatewayConfigurationApiResponsePtrOutput)
}

// Creation UTC date of the API Management gateway.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupApiGatewayResultOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

// ETag of the resource.
func (o LookupApiGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Information regarding how the gateway should be exposed.
func (o LookupApiGatewayResultOutput) Frontend() FrontendConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *FrontendConfigurationResponse { return v.Frontend }).(FrontendConfigurationResponsePtrOutput)
}

// Resource ID.
func (o LookupApiGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupApiGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupApiGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the API Management gateway which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
func (o LookupApiGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SKU properties of the API Management gateway.
func (o LookupApiGatewayResultOutput) Sku() ApiManagementGatewaySkuPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) ApiManagementGatewaySkuPropertiesResponse { return v.Sku }).(ApiManagementGatewaySkuPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApiGatewayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupApiGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The provisioning state of the API Management gateway, which is targeted by the long running operation started on the gateway.
func (o LookupApiGatewayResultOutput) TargetProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.TargetProvisioningState }).(pulumi.StringOutput)
}

// Resource type for API Management resource is set to Microsoft.ApiManagement.
func (o LookupApiGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// The type of VPN in which API Management gateway needs to be configured in.
func (o LookupApiGatewayResultOutput) VirtualNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *string { return v.VirtualNetworkType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiGatewayResultOutput{})
}
