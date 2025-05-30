// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthdataaiservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DeidService
//
// Uses Azure REST API version 2024-09-20.
//
// Other available API versions: 2024-02-28-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthdataaiservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupDeidService(ctx *pulumi.Context, args *LookupDeidServiceArgs, opts ...pulumi.InvokeOption) (*LookupDeidServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDeidServiceResult
	err := ctx.Invoke("azure-native:healthdataaiservices:getDeidService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeidServiceArgs struct {
	// The name of the deid service
	DeidServiceName string `pulumi:"deidServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A HealthDataAIServicesProviderHub resource
type LookupDeidServiceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties DeidServicePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDeidServiceOutput(ctx *pulumi.Context, args LookupDeidServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDeidServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDeidServiceResultOutput, error) {
			args := v.(LookupDeidServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:healthdataaiservices:getDeidService", args, LookupDeidServiceResultOutput{}, options).(LookupDeidServiceResultOutput), nil
		}).(LookupDeidServiceResultOutput)
}

type LookupDeidServiceOutputArgs struct {
	// The name of the deid service
	DeidServiceName pulumi.StringInput `pulumi:"deidServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDeidServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeidServiceArgs)(nil)).Elem()
}

// A HealthDataAIServicesProviderHub resource
type LookupDeidServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDeidServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeidServiceResult)(nil)).Elem()
}

func (o LookupDeidServiceResultOutput) ToLookupDeidServiceResultOutput() LookupDeidServiceResultOutput {
	return o
}

func (o LookupDeidServiceResultOutput) ToLookupDeidServiceResultOutputWithContext(ctx context.Context) LookupDeidServiceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDeidServiceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDeidServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o LookupDeidServiceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupDeidServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDeidServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupDeidServiceResultOutput) Properties() DeidServicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) DeidServicePropertiesResponse { return v.Properties }).(DeidServicePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDeidServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDeidServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDeidServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeidServiceResultOutput{})
}
