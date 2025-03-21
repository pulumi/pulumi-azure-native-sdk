// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get installed Solution details by Solution id.
func LookupSolution(ctx *pulumi.Context, args *LookupSolutionArgs, opts ...pulumi.InvokeOption) (*LookupSolutionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSolutionResult
	err := ctx.Invoke("azure-native:agfoodplatform/v20210901preview:getSolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSolutionArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName string `pulumi:"farmBeatsResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Solution Id of the solution.
	SolutionId string `pulumi:"solutionId"`
}

// Solution resource.
type LookupSolutionResult struct {
	// The ETag value to implement optimistic concurrency.
	ETag string `pulumi:"eTag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Solution resource properties.
	Properties SolutionPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSolutionOutput(ctx *pulumi.Context, args LookupSolutionOutputArgs, opts ...pulumi.InvokeOption) LookupSolutionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSolutionResultOutput, error) {
			args := v.(LookupSolutionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:agfoodplatform/v20210901preview:getSolution", args, LookupSolutionResultOutput{}, options).(LookupSolutionResultOutput), nil
		}).(LookupSolutionResultOutput)
}

type LookupSolutionOutputArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName pulumi.StringInput `pulumi:"farmBeatsResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Solution Id of the solution.
	SolutionId pulumi.StringInput `pulumi:"solutionId"`
}

func (LookupSolutionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSolutionArgs)(nil)).Elem()
}

// Solution resource.
type LookupSolutionResultOutput struct{ *pulumi.OutputState }

func (LookupSolutionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSolutionResult)(nil)).Elem()
}

func (o LookupSolutionResultOutput) ToLookupSolutionResultOutput() LookupSolutionResultOutput {
	return o
}

func (o LookupSolutionResultOutput) ToLookupSolutionResultOutputWithContext(ctx context.Context) LookupSolutionResultOutput {
	return o
}

// The ETag value to implement optimistic concurrency.
func (o LookupSolutionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSolutionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSolutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Solution resource properties.
func (o LookupSolutionResultOutput) Properties() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupSolutionResult) SolutionPropertiesResponse { return v.Properties }).(SolutionPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSolutionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSolutionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSolutionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSolutionResultOutput{})
}
