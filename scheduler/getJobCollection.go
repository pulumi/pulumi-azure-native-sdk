// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a job collection.
//
// Uses Azure REST API version 2016-03-01.
func LookupJobCollection(ctx *pulumi.Context, args *LookupJobCollectionArgs, opts ...pulumi.InvokeOption) (*LookupJobCollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobCollectionResult
	err := ctx.Invoke("azure-native:scheduler:getJobCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCollectionArgs struct {
	// The job collection name.
	JobCollectionName string `pulumi:"jobCollectionName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobCollectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets the job collection resource identifier.
	Id string `pulumi:"id"`
	// Gets or sets the storage account location.
	Location *string `pulumi:"location"`
	// Gets or sets the job collection resource name.
	Name *string `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesResponse `pulumi:"properties"`
	// Gets or sets the tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the job collection resource type.
	Type string `pulumi:"type"`
}

func LookupJobCollectionOutput(ctx *pulumi.Context, args LookupJobCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupJobCollectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupJobCollectionResultOutput, error) {
			args := v.(LookupJobCollectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:scheduler:getJobCollection", args, LookupJobCollectionResultOutput{}, options).(LookupJobCollectionResultOutput), nil
		}).(LookupJobCollectionResultOutput)
}

type LookupJobCollectionOutputArgs struct {
	// The job collection name.
	JobCollectionName pulumi.StringInput `pulumi:"jobCollectionName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCollectionArgs)(nil)).Elem()
}

type LookupJobCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupJobCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobCollectionResult)(nil)).Elem()
}

func (o LookupJobCollectionResultOutput) ToLookupJobCollectionResultOutput() LookupJobCollectionResultOutput {
	return o
}

func (o LookupJobCollectionResultOutput) ToLookupJobCollectionResultOutputWithContext(ctx context.Context) LookupJobCollectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupJobCollectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets the job collection resource identifier.
func (o LookupJobCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the storage account location.
func (o LookupJobCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the job collection resource name.
func (o LookupJobCollectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Gets or sets the job collection properties.
func (o LookupJobCollectionResultOutput) Properties() JobCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) JobCollectionPropertiesResponse { return v.Properties }).(JobCollectionPropertiesResponseOutput)
}

// Gets or sets the tags.
func (o LookupJobCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the job collection resource type.
func (o LookupJobCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobCollectionResultOutput{})
}
