// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DC Access code in the case of Self Managed Shipping.
func ListOrderDCAccessCode(ctx *pulumi.Context, args *ListOrderDCAccessCodeArgs, opts ...pulumi.InvokeOption) (*ListOrderDCAccessCodeResult, error) {
	var rv ListOrderDCAccessCodeResult
	err := ctx.Invoke("azure-native:databoxedge/v20220301:listOrderDCAccessCode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOrderDCAccessCodeArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// DC Access code in the case of Self Managed Shipping.
type ListOrderDCAccessCodeResult struct {
	AuthCode *string `pulumi:"authCode"`
}

func ListOrderDCAccessCodeOutput(ctx *pulumi.Context, args ListOrderDCAccessCodeOutputArgs, opts ...pulumi.InvokeOption) ListOrderDCAccessCodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOrderDCAccessCodeResult, error) {
			args := v.(ListOrderDCAccessCodeArgs)
			r, err := ListOrderDCAccessCode(ctx, &args, opts...)
			var s ListOrderDCAccessCodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOrderDCAccessCodeResultOutput)
}

type ListOrderDCAccessCodeOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListOrderDCAccessCodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOrderDCAccessCodeArgs)(nil)).Elem()
}

// DC Access code in the case of Self Managed Shipping.
type ListOrderDCAccessCodeResultOutput struct{ *pulumi.OutputState }

func (ListOrderDCAccessCodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOrderDCAccessCodeResult)(nil)).Elem()
}

func (o ListOrderDCAccessCodeResultOutput) ToListOrderDCAccessCodeResultOutput() ListOrderDCAccessCodeResultOutput {
	return o
}

func (o ListOrderDCAccessCodeResultOutput) ToListOrderDCAccessCodeResultOutputWithContext(ctx context.Context) ListOrderDCAccessCodeResultOutput {
	return o
}

func (o ListOrderDCAccessCodeResultOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOrderDCAccessCodeResult) *string { return v.AuthCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOrderDCAccessCodeResultOutput{})
}