// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package frontdoor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the properties of an Experiment
//
// Uses Azure REST API version 2019-11-01.
func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExperimentResult
	err := ctx.Invoke("azure-native:frontdoor:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	// The Experiment identifier associated with the Experiment
	ExperimentName string `pulumi:"experimentName"`
	// The Profile identifier associated with the Tenant and Partner
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the properties of an Experiment
type LookupExperimentResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The description of the details or intents of the Experiment
	Description *string `pulumi:"description"`
	// The state of the Experiment
	EnabledState *string `pulumi:"enabledState"`
	// The endpoint A of an experiment
	EndpointA *EndpointResponse `pulumi:"endpointA"`
	// The endpoint B of an experiment
	EndpointB *EndpointResponse `pulumi:"endpointB"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource status.
	ResourceState string `pulumi:"resourceState"`
	// The uri to the Script used in the Experiment
	ScriptFileUri string `pulumi:"scriptFileUri"`
	// The description of Experiment status from the server side
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupExperimentOutput(ctx *pulumi.Context, args LookupExperimentOutputArgs, opts ...pulumi.InvokeOption) LookupExperimentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExperimentResultOutput, error) {
			args := v.(LookupExperimentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:frontdoor:getExperiment", args, LookupExperimentResultOutput{}, options).(LookupExperimentResultOutput), nil
		}).(LookupExperimentResultOutput)
}

type LookupExperimentOutputArgs struct {
	// The Experiment identifier associated with the Experiment
	ExperimentName pulumi.StringInput `pulumi:"experimentName"`
	// The Profile identifier associated with the Tenant and Partner
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExperimentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentArgs)(nil)).Elem()
}

// Defines the properties of an Experiment
type LookupExperimentResultOutput struct{ *pulumi.OutputState }

func (LookupExperimentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentResult)(nil)).Elem()
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutput() LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutputWithContext(ctx context.Context) LookupExperimentResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupExperimentResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The description of the details or intents of the Experiment
func (o LookupExperimentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The state of the Experiment
func (o LookupExperimentResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// The endpoint A of an experiment
func (o LookupExperimentResultOutput) EndpointA() EndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *EndpointResponse { return v.EndpointA }).(EndpointResponsePtrOutput)
}

// The endpoint B of an experiment
func (o LookupExperimentResultOutput) EndpointB() EndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *EndpointResponse { return v.EndpointB }).(EndpointResponsePtrOutput)
}

// Resource ID.
func (o LookupExperimentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupExperimentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupExperimentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource status.
func (o LookupExperimentResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// The uri to the Script used in the Experiment
func (o LookupExperimentResultOutput) ScriptFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.ScriptFileUri }).(pulumi.StringOutput)
}

// The description of Experiment status from the server side
func (o LookupExperimentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupExperimentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExperimentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupExperimentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExperimentResultOutput{})
}
