// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an Azure Databricks Access Connector.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2023-05-01, 2024-09-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databricks [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupAccessConnector(ctx *pulumi.Context, args *LookupAccessConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAccessConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessConnectorResult
	err := ctx.Invoke("azure-native:databricks:getAccessConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessConnectorArgs struct {
	// The name of the Azure Databricks Access Connector.
	ConnectorName string `pulumi:"connectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about Azure Databricks Access Connector.
type LookupAccessConnectorResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Databricks Access Connector properties
	Properties AccessConnectorPropertiesResponse `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupAccessConnectorOutput(ctx *pulumi.Context, args LookupAccessConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAccessConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAccessConnectorResultOutput, error) {
			args := v.(LookupAccessConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:databricks:getAccessConnector", args, LookupAccessConnectorResultOutput{}, options).(LookupAccessConnectorResultOutput), nil
		}).(LookupAccessConnectorResultOutput)
}

type LookupAccessConnectorOutputArgs struct {
	// The name of the Azure Databricks Access Connector.
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessConnectorArgs)(nil)).Elem()
}

// Information about Azure Databricks Access Connector.
type LookupAccessConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAccessConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessConnectorResult)(nil)).Elem()
}

func (o LookupAccessConnectorResultOutput) ToLookupAccessConnectorResultOutput() LookupAccessConnectorResultOutput {
	return o
}

func (o LookupAccessConnectorResultOutput) ToLookupAccessConnectorResultOutputWithContext(ctx context.Context) LookupAccessConnectorResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAccessConnectorResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAccessConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o LookupAccessConnectorResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupAccessConnectorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAccessConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Databricks Access Connector properties
func (o LookupAccessConnectorResultOutput) Properties() AccessConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) AccessConnectorPropertiesResponse { return v.Properties }).(AccessConnectorPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupAccessConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAccessConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupAccessConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessConnectorResultOutput{})
}
