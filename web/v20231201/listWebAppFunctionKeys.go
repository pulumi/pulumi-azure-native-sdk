// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get function keys for a function in a web site, or a deployment slot.
func ListWebAppFunctionKeys(ctx *pulumi.Context, args *ListWebAppFunctionKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppFunctionKeysResult
	err := ctx.Invoke("azure-native:web/v20231201:listWebAppFunctionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionKeysArgs struct {
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListWebAppFunctionKeysResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppFunctionKeysOutput(ctx *pulumi.Context, args ListWebAppFunctionKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebAppFunctionKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWebAppFunctionKeysResultOutput, error) {
			args := v.(ListWebAppFunctionKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web/v20231201:listWebAppFunctionKeys", args, ListWebAppFunctionKeysResultOutput{}, options).(ListWebAppFunctionKeysResultOutput), nil
		}).(ListWebAppFunctionKeysResultOutput)
}

type ListWebAppFunctionKeysOutputArgs struct {
	// Function name.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppFunctionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionKeysArgs)(nil)).Elem()
}

// String dictionary resource.
type ListWebAppFunctionKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebAppFunctionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionKeysResult)(nil)).Elem()
}

func (o ListWebAppFunctionKeysResultOutput) ToListWebAppFunctionKeysResultOutput() ListWebAppFunctionKeysResultOutput {
	return o
}

func (o ListWebAppFunctionKeysResultOutput) ToListWebAppFunctionKeysResultOutputWithContext(ctx context.Context) ListWebAppFunctionKeysResultOutput {
	return o
}

// Resource Id.
func (o ListWebAppFunctionKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppFunctionKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppFunctionKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListWebAppFunctionKeysResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListWebAppFunctionKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppFunctionKeysResultOutput{})
}
