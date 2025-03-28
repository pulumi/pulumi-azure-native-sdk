// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response of a list operation.
func ListVMHost(ctx *pulumi.Context, args *ListVMHostArgs, opts ...pulumi.InvokeOption) (*ListVMHostResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListVMHostResult
	err := ctx.Invoke("azure-native:elastic/v20250115preview:listVMHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVMHostArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response of a list operation.
type ListVMHostResult struct {
	// Link to the next Vm resource Id, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []VMResourcesResponse `pulumi:"value"`
}

func ListVMHostOutput(ctx *pulumi.Context, args ListVMHostOutputArgs, opts ...pulumi.InvokeOption) ListVMHostResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListVMHostResultOutput, error) {
			args := v.(ListVMHostArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:elastic/v20250115preview:listVMHost", args, ListVMHostResultOutput{}, options).(ListVMHostResultOutput), nil
		}).(ListVMHostResultOutput)
}

type ListVMHostOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListVMHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVMHostArgs)(nil)).Elem()
}

// Response of a list operation.
type ListVMHostResultOutput struct{ *pulumi.OutputState }

func (ListVMHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVMHostResult)(nil)).Elem()
}

func (o ListVMHostResultOutput) ToListVMHostResultOutput() ListVMHostResultOutput {
	return o
}

func (o ListVMHostResultOutput) ToListVMHostResultOutputWithContext(ctx context.Context) ListVMHostResultOutput {
	return o
}

// Link to the next Vm resource Id, if any.
func (o ListVMHostResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVMHostResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Results of a list operation.
func (o ListVMHostResultOutput) Value() VMResourcesResponseArrayOutput {
	return o.ApplyT(func(v ListVMHostResult) []VMResourcesResponse { return v.Value }).(VMResourcesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVMHostResultOutput{})
}
