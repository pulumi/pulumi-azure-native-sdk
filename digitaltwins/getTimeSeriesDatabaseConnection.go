// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the description of an existing time series database connection.
//
// Uses Azure REST API version 2023-01-31.
func LookupTimeSeriesDatabaseConnection(ctx *pulumi.Context, args *LookupTimeSeriesDatabaseConnectionArgs, opts ...pulumi.InvokeOption) (*LookupTimeSeriesDatabaseConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTimeSeriesDatabaseConnectionResult
	err := ctx.Invoke("azure-native:digitaltwins:getTimeSeriesDatabaseConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTimeSeriesDatabaseConnectionArgs struct {
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName string `pulumi:"resourceName"`
	// Name of time series database connection.
	TimeSeriesDatabaseConnectionName string `pulumi:"timeSeriesDatabaseConnectionName"`
}

// Describes a time series database connection resource.
type LookupTimeSeriesDatabaseConnectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Extension resource name.
	Name string `pulumi:"name"`
	// Properties of a specific time series database connection.
	Properties AzureDataExplorerConnectionPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupTimeSeriesDatabaseConnectionResult
func (val *LookupTimeSeriesDatabaseConnectionResult) Defaults() *LookupTimeSeriesDatabaseConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupTimeSeriesDatabaseConnectionOutput(ctx *pulumi.Context, args LookupTimeSeriesDatabaseConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupTimeSeriesDatabaseConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTimeSeriesDatabaseConnectionResultOutput, error) {
			args := v.(LookupTimeSeriesDatabaseConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:digitaltwins:getTimeSeriesDatabaseConnection", args, LookupTimeSeriesDatabaseConnectionResultOutput{}, options).(LookupTimeSeriesDatabaseConnectionResultOutput), nil
		}).(LookupTimeSeriesDatabaseConnectionResultOutput)
}

type LookupTimeSeriesDatabaseConnectionOutputArgs struct {
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// Name of time series database connection.
	TimeSeriesDatabaseConnectionName pulumi.StringInput `pulumi:"timeSeriesDatabaseConnectionName"`
}

func (LookupTimeSeriesDatabaseConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeSeriesDatabaseConnectionArgs)(nil)).Elem()
}

// Describes a time series database connection resource.
type LookupTimeSeriesDatabaseConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupTimeSeriesDatabaseConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeSeriesDatabaseConnectionResult)(nil)).Elem()
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) ToLookupTimeSeriesDatabaseConnectionResultOutput() LookupTimeSeriesDatabaseConnectionResultOutput {
	return o
}

func (o LookupTimeSeriesDatabaseConnectionResultOutput) ToLookupTimeSeriesDatabaseConnectionResultOutputWithContext(ctx context.Context) LookupTimeSeriesDatabaseConnectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupTimeSeriesDatabaseConnectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupTimeSeriesDatabaseConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Extension resource name.
func (o LookupTimeSeriesDatabaseConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of a specific time series database connection.
func (o LookupTimeSeriesDatabaseConnectionResultOutput) Properties() AzureDataExplorerConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) AzureDataExplorerConnectionPropertiesResponse {
		return v.Properties
	}).(AzureDataExplorerConnectionPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupTimeSeriesDatabaseConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o LookupTimeSeriesDatabaseConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeSeriesDatabaseConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTimeSeriesDatabaseConnectionResultOutput{})
}
