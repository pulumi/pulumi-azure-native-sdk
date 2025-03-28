// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupMTPDataConnector(ctx *pulumi.Context, args *LookupMTPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMTPDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMTPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230601preview:getMTPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMTPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents MTP (Microsoft Threat Protection) data connector.
type LookupMTPDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes MTPDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The available filtered providers for the connector.
	FilteredProviders *MtpFilteredProvidersResponse `pulumi:"filteredProviders"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatProtection'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMTPDataConnectorOutput(ctx *pulumi.Context, args LookupMTPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMTPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMTPDataConnectorResultOutput, error) {
			args := v.(LookupMTPDataConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:securityinsights/v20230601preview:getMTPDataConnector", args, LookupMTPDataConnectorResultOutput{}, options).(LookupMTPDataConnectorResultOutput), nil
		}).(LookupMTPDataConnectorResultOutput)
}

type LookupMTPDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMTPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMTPDataConnectorArgs)(nil)).Elem()
}

// Represents MTP (Microsoft Threat Protection) data connector.
type LookupMTPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMTPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMTPDataConnectorResult)(nil)).Elem()
}

func (o LookupMTPDataConnectorResultOutput) ToLookupMTPDataConnectorResultOutput() LookupMTPDataConnectorResultOutput {
	return o
}

func (o LookupMTPDataConnectorResultOutput) ToLookupMTPDataConnectorResultOutputWithContext(ctx context.Context) LookupMTPDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupMTPDataConnectorResultOutput) DataTypes() MTPDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) MTPDataConnectorDataTypesResponse { return v.DataTypes }).(MTPDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o LookupMTPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The available filtered providers for the connector.
func (o LookupMTPDataConnectorResultOutput) FilteredProviders() MtpFilteredProvidersResponsePtrOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) *MtpFilteredProvidersResponse { return v.FilteredProviders }).(MtpFilteredProvidersResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMTPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftThreatProtection'.
func (o LookupMTPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMTPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMTPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o LookupMTPDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMTPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMTPDataConnectorResultOutput{})
}
