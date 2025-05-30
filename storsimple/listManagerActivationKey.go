// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the activation key of the manager.
//
// Uses Azure REST API version 2017-06-01.
func ListManagerActivationKey(ctx *pulumi.Context, args *ListManagerActivationKeyArgs, opts ...pulumi.InvokeOption) (*ListManagerActivationKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagerActivationKeyResult
	err := ctx.Invoke("azure-native:storsimple:listManagerActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagerActivationKeyArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The key.
type ListManagerActivationKeyResult struct {
	// The activation key for the device.
	ActivationKey string `pulumi:"activationKey"`
}

func ListManagerActivationKeyOutput(ctx *pulumi.Context, args ListManagerActivationKeyOutputArgs, opts ...pulumi.InvokeOption) ListManagerActivationKeyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListManagerActivationKeyResultOutput, error) {
			args := v.(ListManagerActivationKeyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:storsimple:listManagerActivationKey", args, ListManagerActivationKeyResultOutput{}, options).(ListManagerActivationKeyResultOutput), nil
		}).(ListManagerActivationKeyResultOutput)
}

type ListManagerActivationKeyOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListManagerActivationKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerActivationKeyArgs)(nil)).Elem()
}

// The key.
type ListManagerActivationKeyResultOutput struct{ *pulumi.OutputState }

func (ListManagerActivationKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerActivationKeyResult)(nil)).Elem()
}

func (o ListManagerActivationKeyResultOutput) ToListManagerActivationKeyResultOutput() ListManagerActivationKeyResultOutput {
	return o
}

func (o ListManagerActivationKeyResultOutput) ToListManagerActivationKeyResultOutputWithContext(ctx context.Context) ListManagerActivationKeyResultOutput {
	return o
}

// The activation key for the device.
func (o ListManagerActivationKeyResultOutput) ActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerActivationKeyResult) string { return v.ActivationKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagerActivationKeyResultOutput{})
}
