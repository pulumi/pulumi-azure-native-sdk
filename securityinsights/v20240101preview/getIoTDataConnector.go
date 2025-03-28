// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data connector.
func LookupIoTDataConnector(ctx *pulumi.Context, args *LookupIoTDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupIoTDataConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIoTDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20240101preview:getIoTDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIoTDataConnectorArgs struct {
	// Connector ID
	DataConnectorId string `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents IoT data connector.
type LookupIoTDataConnectorResult struct {
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The kind of the data connector
	// Expected value is 'IOT'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The subscription id to connect to, and get the data from.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIoTDataConnectorOutput(ctx *pulumi.Context, args LookupIoTDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupIoTDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIoTDataConnectorResultOutput, error) {
			args := v.(LookupIoTDataConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:securityinsights/v20240101preview:getIoTDataConnector", args, LookupIoTDataConnectorResultOutput{}, options).(LookupIoTDataConnectorResultOutput), nil
		}).(LookupIoTDataConnectorResultOutput)
}

type LookupIoTDataConnectorOutputArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringInput `pulumi:"dataConnectorId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIoTDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTDataConnectorArgs)(nil)).Elem()
}

// Represents IoT data connector.
type LookupIoTDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupIoTDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTDataConnectorResult)(nil)).Elem()
}

func (o LookupIoTDataConnectorResultOutput) ToLookupIoTDataConnectorResultOutput() LookupIoTDataConnectorResultOutput {
	return o
}

func (o LookupIoTDataConnectorResultOutput) ToLookupIoTDataConnectorResultOutputWithContext(ctx context.Context) LookupIoTDataConnectorResultOutput {
	return o
}

// The available data types for the connector.
func (o LookupIoTDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupIoTDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupIoTDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the data connector
// Expected value is 'IOT'.
func (o LookupIoTDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIoTDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The subscription id to connect to, and get the data from.
func (o LookupIoTDataConnectorResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIoTDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIoTDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoTDataConnectorResultOutput{})
}
