// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Application Insights component linked storage accounts
func LookupComponentLinkedStorageAccount(ctx *pulumi.Context, args *LookupComponentLinkedStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupComponentLinkedStorageAccountResult, error) {
	var rv LookupComponentLinkedStorageAccountResult
	err := ctx.Invoke("azure-native:insights/v20200301preview:getComponentLinkedStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentLinkedStorageAccountArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
	StorageType       string `pulumi:"storageType"`
}

// An Application Insights component linked storage accounts
type LookupComponentLinkedStorageAccountResult struct {
	Id                   string  `pulumi:"id"`
	LinkedStorageAccount *string `pulumi:"linkedStorageAccount"`
	Name                 string  `pulumi:"name"`
	Type                 string  `pulumi:"type"`
}

func LookupComponentLinkedStorageAccountOutput(ctx *pulumi.Context, args LookupComponentLinkedStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupComponentLinkedStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentLinkedStorageAccountResult, error) {
			args := v.(LookupComponentLinkedStorageAccountArgs)
			r, err := LookupComponentLinkedStorageAccount(ctx, &args, opts...)
			var s LookupComponentLinkedStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentLinkedStorageAccountResultOutput)
}

type LookupComponentLinkedStorageAccountOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
	StorageType       pulumi.StringInput `pulumi:"storageType"`
}

func (LookupComponentLinkedStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentLinkedStorageAccountArgs)(nil)).Elem()
}

// An Application Insights component linked storage accounts
type LookupComponentLinkedStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupComponentLinkedStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentLinkedStorageAccountResult)(nil)).Elem()
}

func (o LookupComponentLinkedStorageAccountResultOutput) ToLookupComponentLinkedStorageAccountResultOutput() LookupComponentLinkedStorageAccountResultOutput {
	return o
}

func (o LookupComponentLinkedStorageAccountResultOutput) ToLookupComponentLinkedStorageAccountResultOutputWithContext(ctx context.Context) LookupComponentLinkedStorageAccountResultOutput {
	return o
}

func (o LookupComponentLinkedStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentLinkedStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComponentLinkedStorageAccountResultOutput) LinkedStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentLinkedStorageAccountResult) *string { return v.LinkedStorageAccount }).(pulumi.StringPtrOutput)
}

func (o LookupComponentLinkedStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentLinkedStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComponentLinkedStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentLinkedStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentLinkedStorageAccountResultOutput{})
}