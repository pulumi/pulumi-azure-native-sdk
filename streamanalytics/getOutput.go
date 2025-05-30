// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details about the specified output.
//
// Uses Azure REST API version 2020-03-01.
//
// Other available API versions: 2021-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native streamanalytics [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupOutput(ctx *pulumi.Context, args *LookupOutputArgs, opts ...pulumi.InvokeOption) (*LookupOutputResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOutputResult
	err := ctx.Invoke("azure-native:streamanalytics:getOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutputArgs struct {
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the output.
	OutputName string `pulumi:"outputName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
type LookupOutputResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
	Datasource interface{} `pulumi:"datasource"`
	// Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
	Diagnostics DiagnosticsResponse `pulumi:"diagnostics"`
	// The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
	Serialization interface{} `pulumi:"serialization"`
	// The size window to constrain a Stream Analytics output to.
	SizeWindow *int `pulumi:"sizeWindow"`
	// The time frame for filtering Stream Analytics job outputs.
	TimeWindow *string `pulumi:"timeWindow"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupOutputOutput(ctx *pulumi.Context, args LookupOutputOutputArgs, opts ...pulumi.InvokeOption) LookupOutputResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOutputResultOutput, error) {
			args := v.(LookupOutputArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:streamanalytics:getOutput", args, LookupOutputResultOutput{}, options).(LookupOutputResultOutput), nil
		}).(LookupOutputResultOutput)
}

type LookupOutputOutputArgs struct {
	// The name of the streaming job.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the output.
	OutputName pulumi.StringInput `pulumi:"outputName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOutputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutputArgs)(nil)).Elem()
}

// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
type LookupOutputResultOutput struct{ *pulumi.OutputState }

func (LookupOutputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutputResult)(nil)).Elem()
}

func (o LookupOutputResultOutput) ToLookupOutputResultOutput() LookupOutputResultOutput {
	return o
}

func (o LookupOutputResultOutput) ToLookupOutputResultOutputWithContext(ctx context.Context) LookupOutputResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupOutputResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
func (o LookupOutputResultOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupOutputResult) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

// Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
func (o LookupOutputResultOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v LookupOutputResult) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

// The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
func (o LookupOutputResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupOutputResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupOutputResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutputResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
func (o LookupOutputResultOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupOutputResult) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

// The size window to constrain a Stream Analytics output to.
func (o LookupOutputResultOutput) SizeWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOutputResult) *int { return v.SizeWindow }).(pulumi.IntPtrOutput)
}

// The time frame for filtering Stream Analytics job outputs.
func (o LookupOutputResultOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutputResult) *string { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

// Resource type
func (o LookupOutputResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutputResultOutput{})
}
