// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataSet in a share
//
// Uses Azure REST API version 2021-08-01.
func LookupKustoClusterDataSet(ctx *pulumi.Context, args *LookupKustoClusterDataSetArgs, opts ...pulumi.InvokeOption) (*LookupKustoClusterDataSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKustoClusterDataSetResult
	err := ctx.Invoke("azure-native:datashare:getKustoClusterDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoClusterDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// A kusto cluster data set.
type LookupKustoClusterDataSetResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'KustoCluster'.
	Kind string `pulumi:"kind"`
	// Resource id of the kusto cluster.
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	// Location of the kusto cluster.
	Location string `pulumi:"location"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Provisioning state of the kusto cluster data set.
	ProvisioningState string `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupKustoClusterDataSetOutput(ctx *pulumi.Context, args LookupKustoClusterDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupKustoClusterDataSetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKustoClusterDataSetResultOutput, error) {
			args := v.(LookupKustoClusterDataSetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datashare:getKustoClusterDataSet", args, LookupKustoClusterDataSetResultOutput{}, options).(LookupKustoClusterDataSetResultOutput), nil
		}).(LookupKustoClusterDataSetResultOutput)
}

type LookupKustoClusterDataSetOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName pulumi.StringInput `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupKustoClusterDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetArgs)(nil)).Elem()
}

// A kusto cluster data set.
type LookupKustoClusterDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupKustoClusterDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetResult)(nil)).Elem()
}

func (o LookupKustoClusterDataSetResultOutput) ToLookupKustoClusterDataSetResultOutput() LookupKustoClusterDataSetResultOutput {
	return o
}

func (o LookupKustoClusterDataSetResultOutput) ToLookupKustoClusterDataSetResultOutputWithContext(ctx context.Context) LookupKustoClusterDataSetResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupKustoClusterDataSetResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Unique id for identifying a data set resource
func (o LookupKustoClusterDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupKustoClusterDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'KustoCluster'.
func (o LookupKustoClusterDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource id of the kusto cluster.
func (o LookupKustoClusterDataSetResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

// Location of the kusto cluster.
func (o LookupKustoClusterDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupKustoClusterDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the kusto cluster data set.
func (o LookupKustoClusterDataSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupKustoClusterDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupKustoClusterDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoClusterDataSetResultOutput{})
}
