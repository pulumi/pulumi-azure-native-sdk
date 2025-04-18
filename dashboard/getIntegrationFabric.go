// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dashboard

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration fabric resource type.
//
// Uses Azure REST API version 2024-10-01.
//
// Other available API versions: 2023-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dashboard [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupIntegrationFabric(ctx *pulumi.Context, args *LookupIntegrationFabricArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationFabricResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationFabricResult
	err := ctx.Invoke("azure-native:dashboard:getIntegrationFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationFabricArgs struct {
	// The integration fabric name of Azure Managed Grafana.
	IntegrationFabricName string `pulumi:"integrationFabricName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The integration fabric resource type.
type LookupIntegrationFabricResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name       string                              `pulumi:"name"`
	Properties IntegrationFabricPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIntegrationFabricOutput(ctx *pulumi.Context, args LookupIntegrationFabricOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationFabricResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIntegrationFabricResultOutput, error) {
			args := v.(LookupIntegrationFabricArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:dashboard:getIntegrationFabric", args, LookupIntegrationFabricResultOutput{}, options).(LookupIntegrationFabricResultOutput), nil
		}).(LookupIntegrationFabricResultOutput)
}

type LookupIntegrationFabricOutputArgs struct {
	// The integration fabric name of Azure Managed Grafana.
	IntegrationFabricName pulumi.StringInput `pulumi:"integrationFabricName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIntegrationFabricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationFabricArgs)(nil)).Elem()
}

// The integration fabric resource type.
type LookupIntegrationFabricResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationFabricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationFabricResult)(nil)).Elem()
}

func (o LookupIntegrationFabricResultOutput) ToLookupIntegrationFabricResultOutput() LookupIntegrationFabricResultOutput {
	return o
}

func (o LookupIntegrationFabricResultOutput) ToLookupIntegrationFabricResultOutputWithContext(ctx context.Context) LookupIntegrationFabricResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupIntegrationFabricResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIntegrationFabricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupIntegrationFabricResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIntegrationFabricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationFabricResultOutput) Properties() IntegrationFabricPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) IntegrationFabricPropertiesResponse { return v.Properties }).(IntegrationFabricPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIntegrationFabricResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupIntegrationFabricResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIntegrationFabricResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationFabricResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationFabricResultOutput{})
}
