// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
//
// Uses Azure REST API version 2024-02-01-preview.
//
// Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupProviderInstance(ctx *pulumi.Context, args *LookupProviderInstanceArgs, opts ...pulumi.InvokeOption) (*LookupProviderInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProviderInstanceResult
	err := ctx.Invoke("azure-native:workloads:getProviderInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProviderInstanceArgs struct {
	// Name of the SAP monitor resource.
	MonitorName string `pulumi:"monitorName"`
	// Name of the provider instance.
	ProviderInstanceName string `pulumi:"providerInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A provider instance associated with SAP monitor.
type LookupProviderInstanceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Defines the provider instance errors.
	Errors ErrorDetailResponse `pulumi:"errors"`
	// Resource health details
	Health HealthResponse `pulumi:"health"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the provider specific properties.
	ProviderSettings interface{} `pulumi:"providerSettings"`
	// State of provisioning of the provider instance
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupProviderInstanceOutput(ctx *pulumi.Context, args LookupProviderInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupProviderInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProviderInstanceResultOutput, error) {
			args := v.(LookupProviderInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:workloads:getProviderInstance", args, LookupProviderInstanceResultOutput{}, options).(LookupProviderInstanceResultOutput), nil
		}).(LookupProviderInstanceResultOutput)
}

type LookupProviderInstanceOutputArgs struct {
	// Name of the SAP monitor resource.
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// Name of the provider instance.
	ProviderInstanceName pulumi.StringInput `pulumi:"providerInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProviderInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderInstanceArgs)(nil)).Elem()
}

// A provider instance associated with SAP monitor.
type LookupProviderInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupProviderInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderInstanceResult)(nil)).Elem()
}

func (o LookupProviderInstanceResultOutput) ToLookupProviderInstanceResultOutput() LookupProviderInstanceResultOutput {
	return o
}

func (o LookupProviderInstanceResultOutput) ToLookupProviderInstanceResultOutputWithContext(ctx context.Context) LookupProviderInstanceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupProviderInstanceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Defines the provider instance errors.
func (o LookupProviderInstanceResultOutput) Errors() ErrorDetailResponseOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) ErrorDetailResponse { return v.Errors }).(ErrorDetailResponseOutput)
}

// Resource health details
func (o LookupProviderInstanceResultOutput) Health() HealthResponseOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) HealthResponse { return v.Health }).(HealthResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupProviderInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupProviderInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the provider specific properties.
func (o LookupProviderInstanceResultOutput) ProviderSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) interface{} { return v.ProviderSettings }).(pulumi.AnyOutput)
}

// State of provisioning of the provider instance
func (o LookupProviderInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupProviderInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupProviderInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderInstanceResultOutput{})
}
