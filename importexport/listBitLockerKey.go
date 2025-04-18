// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package importexport

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the BitLocker Keys for all drives in the specified job.
//
// Uses Azure REST API version 2021-01-01.
func ListBitLockerKey(ctx *pulumi.Context, args *ListBitLockerKeyArgs, opts ...pulumi.InvokeOption) (*ListBitLockerKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListBitLockerKeyResult
	err := ctx.Invoke("azure-native:importexport:listBitLockerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBitLockerKeyArgs struct {
	// The name of the import/export job.
	JobName string `pulumi:"jobName"`
	// The resource group name uniquely identifies the resource group within the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// GetBitLockerKeys response
type ListBitLockerKeyResult struct {
	// drive status
	Value []DriveBitLockerKeyResponse `pulumi:"value"`
}

func ListBitLockerKeyOutput(ctx *pulumi.Context, args ListBitLockerKeyOutputArgs, opts ...pulumi.InvokeOption) ListBitLockerKeyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListBitLockerKeyResultOutput, error) {
			args := v.(ListBitLockerKeyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:importexport:listBitLockerKey", args, ListBitLockerKeyResultOutput{}, options).(ListBitLockerKeyResultOutput), nil
		}).(ListBitLockerKeyResultOutput)
}

type ListBitLockerKeyOutputArgs struct {
	// The name of the import/export job.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The resource group name uniquely identifies the resource group within the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBitLockerKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBitLockerKeyArgs)(nil)).Elem()
}

// GetBitLockerKeys response
type ListBitLockerKeyResultOutput struct{ *pulumi.OutputState }

func (ListBitLockerKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBitLockerKeyResult)(nil)).Elem()
}

func (o ListBitLockerKeyResultOutput) ToListBitLockerKeyResultOutput() ListBitLockerKeyResultOutput {
	return o
}

func (o ListBitLockerKeyResultOutput) ToListBitLockerKeyResultOutputWithContext(ctx context.Context) ListBitLockerKeyResultOutput {
	return o
}

// drive status
func (o ListBitLockerKeyResultOutput) Value() DriveBitLockerKeyResponseArrayOutput {
	return o.ApplyT(func(v ListBitLockerKeyResult) []DriveBitLockerKeyResponse { return v.Value }).(DriveBitLockerKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBitLockerKeyResultOutput{})
}
