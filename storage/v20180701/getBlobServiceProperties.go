// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of a storage account’s Blob service.
//
// Deprecated: Version 2018-07-01 will be removed in v2 of the provider.
func LookupBlobServiceProperties(ctx *pulumi.Context, args *LookupBlobServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupBlobServicePropertiesResult, error) {
	var rv LookupBlobServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20180701:getBlobServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	BlobServicesName  string `pulumi:"blobServicesName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of a storage account’s Blob service.
type LookupBlobServicePropertiesResult struct {
	Cors                  *CorsRulesResponse             `pulumi:"cors"`
	DefaultServiceVersion *string                        `pulumi:"defaultServiceVersion"`
	DeleteRetentionPolicy *DeleteRetentionPolicyResponse `pulumi:"deleteRetentionPolicy"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	Type                  string                         `pulumi:"type"`
}

func LookupBlobServicePropertiesOutput(ctx *pulumi.Context, args LookupBlobServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupBlobServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobServicePropertiesResult, error) {
			args := v.(LookupBlobServicePropertiesArgs)
			r, err := LookupBlobServiceProperties(ctx, &args, opts...)
			var s LookupBlobServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobServicePropertiesResultOutput)
}

type LookupBlobServicePropertiesOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	BlobServicesName  pulumi.StringInput `pulumi:"blobServicesName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobServicePropertiesArgs)(nil)).Elem()
}

// The properties of a storage account’s Blob service.
type LookupBlobServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupBlobServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobServicePropertiesResult)(nil)).Elem()
}

func (o LookupBlobServicePropertiesResultOutput) ToLookupBlobServicePropertiesResultOutput() LookupBlobServicePropertiesResultOutput {
	return o
}

func (o LookupBlobServicePropertiesResultOutput) ToLookupBlobServicePropertiesResultOutputWithContext(ctx context.Context) LookupBlobServicePropertiesResultOutput {
	return o
}

func (o LookupBlobServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) DefaultServiceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *string { return v.DefaultServiceVersion }).(pulumi.StringPtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) DeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) *DeleteRetentionPolicyResponse {
		return v.DeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobServicePropertiesResultOutput{})
}