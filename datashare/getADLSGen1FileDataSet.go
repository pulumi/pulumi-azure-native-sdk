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
func LookupADLSGen1FileDataSet(ctx *pulumi.Context, args *LookupADLSGen1FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen1FileDataSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupADLSGen1FileDataSetResult
	err := ctx.Invoke("azure-native:datashare:getADLSGen1FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen1FileDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// An ADLS Gen 1 file data set.
type LookupADLSGen1FileDataSetResult struct {
	// The ADLS account name.
	AccountName string `pulumi:"accountName"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// The file name in the ADLS account.
	FileName string `pulumi:"fileName"`
	// The folder path within the ADLS account.
	FolderPath string `pulumi:"folderPath"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'AdlsGen1File'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Resource group of ADLS account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Subscription id of ADLS account.
	SubscriptionId string `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupADLSGen1FileDataSetOutput(ctx *pulumi.Context, args LookupADLSGen1FileDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen1FileDataSetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupADLSGen1FileDataSetResultOutput, error) {
			args := v.(LookupADLSGen1FileDataSetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datashare:getADLSGen1FileDataSet", args, LookupADLSGen1FileDataSetResultOutput{}, options).(LookupADLSGen1FileDataSetResultOutput), nil
		}).(LookupADLSGen1FileDataSetResultOutput)
}

type LookupADLSGen1FileDataSetOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName pulumi.StringInput `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen1FileDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen1FileDataSetArgs)(nil)).Elem()
}

// An ADLS Gen 1 file data set.
type LookupADLSGen1FileDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen1FileDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen1FileDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen1FileDataSetResultOutput) ToLookupADLSGen1FileDataSetResultOutput() LookupADLSGen1FileDataSetResultOutput {
	return o
}

func (o LookupADLSGen1FileDataSetResultOutput) ToLookupADLSGen1FileDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen1FileDataSetResultOutput {
	return o
}

// The ADLS account name.
func (o LookupADLSGen1FileDataSetResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupADLSGen1FileDataSetResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Unique id for identifying a data set resource
func (o LookupADLSGen1FileDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// The file name in the ADLS account.
func (o LookupADLSGen1FileDataSetResultOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.FileName }).(pulumi.StringOutput)
}

// The folder path within the ADLS account.
func (o LookupADLSGen1FileDataSetResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupADLSGen1FileDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'AdlsGen1File'.
func (o LookupADLSGen1FileDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupADLSGen1FileDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource group of ADLS account.
func (o LookupADLSGen1FileDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Subscription id of ADLS account.
func (o LookupADLSGen1FileDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupADLSGen1FileDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupADLSGen1FileDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen1FileDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen1FileDataSetResultOutput{})
}
