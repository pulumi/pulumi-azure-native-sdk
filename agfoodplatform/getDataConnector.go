// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package agfoodplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get specific Data Connector resource by DataConnectorName.
//
// Uses Azure REST API version 2023-06-01-preview.
func LookupDataConnector(ctx *pulumi.Context, args *LookupDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataConnectorResult
	err := ctx.Invoke("azure-native:agfoodplatform:getDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataConnectorArgs struct {
	// Connector name.
	DataConnectorName string `pulumi:"dataConnectorName"`
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName string `pulumi:"dataManagerForAgricultureResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// DataConnector Model.
type LookupDataConnectorResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The ETag value to implement optimistic concurrency.
	ETag string `pulumi:"eTag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// DataConnector Properties.
	Properties DataConnectorPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDataConnectorOutput(ctx *pulumi.Context, args LookupDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataConnectorResultOutput, error) {
			args := v.(LookupDataConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:agfoodplatform:getDataConnector", args, LookupDataConnectorResultOutput{}, options).(LookupDataConnectorResultOutput), nil
		}).(LookupDataConnectorResultOutput)
}

type LookupDataConnectorOutputArgs struct {
	// Connector name.
	DataConnectorName pulumi.StringInput `pulumi:"dataConnectorName"`
	// DataManagerForAgriculture resource name.
	DataManagerForAgricultureResourceName pulumi.StringInput `pulumi:"dataManagerForAgricultureResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataConnectorArgs)(nil)).Elem()
}

// DataConnector Model.
type LookupDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataConnectorResult)(nil)).Elem()
}

func (o LookupDataConnectorResultOutput) ToLookupDataConnectorResultOutput() LookupDataConnectorResultOutput {
	return o
}

func (o LookupDataConnectorResultOutput) ToLookupDataConnectorResultOutputWithContext(ctx context.Context) LookupDataConnectorResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDataConnectorResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The ETag value to implement optimistic concurrency.
func (o LookupDataConnectorResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// DataConnector Properties.
func (o LookupDataConnectorResultOutput) Properties() DataConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) DataConnectorPropertiesResponse { return v.Properties }).(DataConnectorPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataConnectorResultOutput{})
}
