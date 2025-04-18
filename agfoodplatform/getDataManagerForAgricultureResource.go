// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package agfoodplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get DataManagerForAgriculture resource.
//
// Uses Azure REST API version 2023-06-01-preview.
func LookupDataManagerForAgricultureResource(ctx *pulumi.Context, args *LookupDataManagerForAgricultureResourceArgs, opts ...pulumi.InvokeOption) (*LookupDataManagerForAgricultureResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataManagerForAgricultureResourceResult
	err := ctx.Invoke("azure-native:agfoodplatform:getDataManagerForAgricultureResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataManagerForAgricultureResourceArgs struct {
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName string `pulumi:"dataManagerForAgricultureResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data Manager For Agriculture ARM Resource.
type LookupDataManagerForAgricultureResourceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Uri of the Data Manager For Agriculture instance.
	InstanceUri string `pulumi:"instanceUri"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Private endpoints.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Data Manager For Agriculture instance provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Sensor integration request model.
	SensorIntegration *SensorIntegrationResponse `pulumi:"sensorIntegration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDataManagerForAgricultureResourceOutput(ctx *pulumi.Context, args LookupDataManagerForAgricultureResourceOutputArgs, opts ...pulumi.InvokeOption) LookupDataManagerForAgricultureResourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataManagerForAgricultureResourceResultOutput, error) {
			args := v.(LookupDataManagerForAgricultureResourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:agfoodplatform:getDataManagerForAgricultureResource", args, LookupDataManagerForAgricultureResourceResultOutput{}, options).(LookupDataManagerForAgricultureResourceResultOutput), nil
		}).(LookupDataManagerForAgricultureResourceResultOutput)
}

type LookupDataManagerForAgricultureResourceOutputArgs struct {
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName pulumi.StringInput `pulumi:"dataManagerForAgricultureResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataManagerForAgricultureResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataManagerForAgricultureResourceArgs)(nil)).Elem()
}

// Data Manager For Agriculture ARM Resource.
type LookupDataManagerForAgricultureResourceResultOutput struct{ *pulumi.OutputState }

func (LookupDataManagerForAgricultureResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataManagerForAgricultureResourceResult)(nil)).Elem()
}

func (o LookupDataManagerForAgricultureResourceResultOutput) ToLookupDataManagerForAgricultureResourceResultOutput() LookupDataManagerForAgricultureResourceResultOutput {
	return o
}

func (o LookupDataManagerForAgricultureResourceResultOutput) ToLookupDataManagerForAgricultureResourceResultOutputWithContext(ctx context.Context) LookupDataManagerForAgricultureResourceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDataManagerForAgricultureResourceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDataManagerForAgricultureResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupDataManagerForAgricultureResourceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Uri of the Data Manager For Agriculture instance.
func (o LookupDataManagerForAgricultureResourceResultOutput) InstanceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.InstanceUri }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDataManagerForAgricultureResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataManagerForAgricultureResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoints.
func (o LookupDataManagerForAgricultureResourceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Data Manager For Agriculture instance provisioning state.
func (o LookupDataManagerForAgricultureResourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
func (o LookupDataManagerForAgricultureResourceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Sensor integration request model.
func (o LookupDataManagerForAgricultureResourceResultOutput) SensorIntegration() SensorIntegrationResponsePtrOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) *SensorIntegrationResponse {
		return v.SensorIntegration
	}).(SensorIntegrationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataManagerForAgricultureResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDataManagerForAgricultureResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataManagerForAgricultureResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerForAgricultureResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataManagerForAgricultureResourceResultOutput{})
}
