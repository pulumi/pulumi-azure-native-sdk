// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcompute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about the view of a gateway.
//
// Uses Azure REST API version 2024-07-31-preview.
//
// Other available API versions: 2024-03-31-preview, 2024-05-20-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:hybridcompute:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayArgs struct {
	// The name of the Gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes an Arc Gateway.
type LookupGatewayResult struct {
	// Specifies the list of features that are enabled for this Gateway.
	AllowedFeatures []string `pulumi:"allowedFeatures"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The endpoint fqdn for the Gateway.
	GatewayEndpoint string `pulumi:"gatewayEndpoint"`
	// A unique, immutable, identifier for the Gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The type of the Gateway resource.
	GatewayType *string `pulumi:"gatewayType"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGatewayResultOutput, error) {
			args := v.(LookupGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridcompute:getGateway", args, LookupGatewayResultOutput{}, options).(LookupGatewayResultOutput), nil
		}).(LookupGatewayResultOutput)
}

type LookupGatewayOutputArgs struct {
	// The name of the Gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}

// Describes an Arc Gateway.
type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

// Specifies the list of features that are enabled for this Gateway.
func (o LookupGatewayResultOutput) AllowedFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayResult) []string { return v.AllowedFeatures }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o LookupGatewayResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The endpoint fqdn for the Gateway.
func (o LookupGatewayResultOutput) GatewayEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.GatewayEndpoint }).(pulumi.StringOutput)
}

// A unique, immutable, identifier for the Gateway.
func (o LookupGatewayResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The type of the Gateway resource.
func (o LookupGatewayResultOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGatewayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
