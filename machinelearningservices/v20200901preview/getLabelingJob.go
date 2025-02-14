// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a labeling job by id.
func LookupLabelingJob(ctx *pulumi.Context, args *LookupLabelingJobArgs, opts ...pulumi.InvokeOption) (*LookupLabelingJobResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLabelingJobResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:getLabelingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabelingJobArgs struct {
	// Boolean value to indicate whether to include JobInstructions in response.
	IncludeJobInstructions *bool `pulumi:"includeJobInstructions"`
	// Boolean value to indicate whether to include LabelCategories in response.
	IncludeLabelCategories *bool `pulumi:"includeLabelCategories"`
	// Name and identifier for LabelingJob.
	LabelingJobId string `pulumi:"labelingJobId"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Machine Learning labeling job object wrapped into ARM resource envelope.
type LookupLabelingJobResult struct {
	// The resource URL of the entity (not URL encoded).
	Id string `pulumi:"id"`
	// The name of the resource entity.
	Name string `pulumi:"name"`
	// Definition of a labeling job.
	Properties LabelingJobPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource provider and type.
	Type string `pulumi:"type"`
}

func LookupLabelingJobOutput(ctx *pulumi.Context, args LookupLabelingJobOutputArgs, opts ...pulumi.InvokeOption) LookupLabelingJobResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLabelingJobResultOutput, error) {
			args := v.(LookupLabelingJobArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices/v20200901preview:getLabelingJob", args, LookupLabelingJobResultOutput{}, options).(LookupLabelingJobResultOutput), nil
		}).(LookupLabelingJobResultOutput)
}

type LookupLabelingJobOutputArgs struct {
	// Boolean value to indicate whether to include JobInstructions in response.
	IncludeJobInstructions pulumi.BoolPtrInput `pulumi:"includeJobInstructions"`
	// Boolean value to indicate whether to include LabelCategories in response.
	IncludeLabelCategories pulumi.BoolPtrInput `pulumi:"includeLabelCategories"`
	// Name and identifier for LabelingJob.
	LabelingJobId pulumi.StringInput `pulumi:"labelingJobId"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupLabelingJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabelingJobArgs)(nil)).Elem()
}

// Machine Learning labeling job object wrapped into ARM resource envelope.
type LookupLabelingJobResultOutput struct{ *pulumi.OutputState }

func (LookupLabelingJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabelingJobResult)(nil)).Elem()
}

func (o LookupLabelingJobResultOutput) ToLookupLabelingJobResultOutput() LookupLabelingJobResultOutput {
	return o
}

func (o LookupLabelingJobResultOutput) ToLookupLabelingJobResultOutputWithContext(ctx context.Context) LookupLabelingJobResultOutput {
	return o
}

// The resource URL of the entity (not URL encoded).
func (o LookupLabelingJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource entity.
func (o LookupLabelingJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Definition of a labeling job.
func (o LookupLabelingJobResultOutput) Properties() LabelingJobPropertiesResponseOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) LabelingJobPropertiesResponse { return v.Properties }).(LabelingJobPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupLabelingJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource provider and type.
func (o LookupLabelingJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabelingJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabelingJobResultOutput{})
}
