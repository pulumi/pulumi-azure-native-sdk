// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151106

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Table.
//
// Deprecated: Version 2015-11-06 will be removed in v2 of the provider.
func LookupDatabaseAccountTable(ctx *pulumi.Context, args *LookupDatabaseAccountTableArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountTableResult, error) {
	var rv LookupDatabaseAccountTableResult
	err := ctx.Invoke("azure-native:documentdb/v20151106:getDatabaseAccountTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}

// An Azure Cosmos DB Table.
type LookupDatabaseAccountTableResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupDatabaseAccountTableOutput(ctx *pulumi.Context, args LookupDatabaseAccountTableOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountTableResult, error) {
			args := v.(LookupDatabaseAccountTableArgs)
			r, err := LookupDatabaseAccountTable(ctx, &args, opts...)
			var s LookupDatabaseAccountTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountTableResultOutput)
}

type LookupDatabaseAccountTableOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableName         pulumi.StringInput `pulumi:"tableName"`
}

func (LookupDatabaseAccountTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountTableArgs)(nil)).Elem()
}

// An Azure Cosmos DB Table.
type LookupDatabaseAccountTableResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountTableResult)(nil)).Elem()
}

func (o LookupDatabaseAccountTableResultOutput) ToLookupDatabaseAccountTableResultOutput() LookupDatabaseAccountTableResultOutput {
	return o
}

func (o LookupDatabaseAccountTableResultOutput) ToLookupDatabaseAccountTableResultOutputWithContext(ctx context.Context) LookupDatabaseAccountTableResultOutput {
	return o
}

func (o LookupDatabaseAccountTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountTableResultOutput{})
}