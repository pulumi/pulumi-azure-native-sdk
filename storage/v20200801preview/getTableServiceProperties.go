// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of a storage account’s Table service.
//
// Deprecated: Version 2020-08-01-preview will be removed in v2 of the provider.
func LookupTableServiceProperties(ctx *pulumi.Context, args *LookupTableServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupTableServicePropertiesResult, error) {
	var rv LookupTableServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:getTableServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableServiceName  string `pulumi:"tableServiceName"`
}

// The properties of a storage account’s Table service.
type LookupTableServicePropertiesResult struct {
	Cors *CorsRulesResponse `pulumi:"cors"`
	Id   string             `pulumi:"id"`
	Name string             `pulumi:"name"`
	Type string             `pulumi:"type"`
}

func LookupTableServicePropertiesOutput(ctx *pulumi.Context, args LookupTableServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupTableServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableServicePropertiesResult, error) {
			args := v.(LookupTableServicePropertiesArgs)
			r, err := LookupTableServiceProperties(ctx, &args, opts...)
			var s LookupTableServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableServicePropertiesResultOutput)
}

type LookupTableServicePropertiesOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableServiceName  pulumi.StringInput `pulumi:"tableServiceName"`
}

func (LookupTableServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableServicePropertiesArgs)(nil)).Elem()
}

// The properties of a storage account’s Table service.
type LookupTableServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupTableServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableServicePropertiesResult)(nil)).Elem()
}

func (o LookupTableServicePropertiesResultOutput) ToLookupTableServicePropertiesResultOutput() LookupTableServicePropertiesResultOutput {
	return o
}

func (o LookupTableServicePropertiesResultOutput) ToLookupTableServicePropertiesResultOutputWithContext(ctx context.Context) LookupTableServicePropertiesResultOutput {
	return o
}

func (o LookupTableServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o LookupTableServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTableServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTableServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableServicePropertiesResultOutput{})
}