// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a job.
//
// Uses Azure REST API version 2016-03-01.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:scheduler:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The job collection name.
	JobCollectionName string `pulumi:"jobCollectionName"`
	// The job name.
	JobName string `pulumi:"jobName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets the job resource identifier.
	Id string `pulumi:"id"`
	// Gets the job resource name.
	Name string `pulumi:"name"`
	// Gets or sets the job properties.
	Properties JobPropertiesResponse `pulumi:"properties"`
	// Gets the job resource type.
	Type string `pulumi:"type"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupJobResultOutput, error) {
			args := v.(LookupJobArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:scheduler:getJob", args, LookupJobResultOutput{}, options).(LookupJobResultOutput), nil
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	// The job collection name.
	JobCollectionName pulumi.StringInput `pulumi:"jobCollectionName"`
	// The job name.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupJobResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets the job resource identifier.
func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets the job resource name.
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the job properties.
func (o LookupJobResultOutput) Properties() JobPropertiesResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobPropertiesResponse { return v.Properties }).(JobPropertiesResponseOutput)
}

// Gets the job resource type.
func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
