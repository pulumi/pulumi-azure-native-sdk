// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details about the specified function.
func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFunctionResult
	err := ctx.Invoke("azure-native:streamanalytics/v20160301:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionArgs struct {
	// The name of the function.
	FunctionName string `pulumi:"functionName"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A function object, containing all information associated with the named function. All functions are contained under a streaming job.
type LookupFunctionResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with a function.
	Properties ScalarFunctionPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupFunctionOutput(ctx *pulumi.Context, args LookupFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFunctionResultOutput, error) {
			args := v.(LookupFunctionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:streamanalytics/v20160301:getFunction", args, LookupFunctionResultOutput{}, options).(LookupFunctionResultOutput), nil
		}).(LookupFunctionResultOutput)
}

type LookupFunctionOutputArgs struct {
	// The name of the function.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// The name of the streaming job.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionArgs)(nil)).Elem()
}

// A function object, containing all information associated with the named function. All functions are contained under a streaming job.
type LookupFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionResult)(nil)).Elem()
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutput() LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutputWithContext(ctx context.Context) LookupFunctionResultOutput {
	return o
}

// Resource Id
func (o LookupFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupFunctionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The properties that are associated with a function.
func (o LookupFunctionResultOutput) Properties() ScalarFunctionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupFunctionResult) ScalarFunctionPropertiesResponse { return v.Properties }).(ScalarFunctionPropertiesResponseOutput)
}

// Resource type
func (o LookupFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionResultOutput{})
}
