// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of File services in storage account.
//
// Deprecated: Version 2020-08-01-preview will be removed in v2 of the provider.
func LookupFileServiceProperties(ctx *pulumi.Context, args *LookupFileServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupFileServicePropertiesResult, error) {
	var rv LookupFileServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:getFileServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	FileServicesName  string `pulumi:"fileServicesName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of File services in storage account.
type LookupFileServicePropertiesResult struct {
	Cors                       *CorsRulesResponse             `pulumi:"cors"`
	Id                         string                         `pulumi:"id"`
	Name                       string                         `pulumi:"name"`
	ProtocolSettings           *ProtocolSettingsResponse      `pulumi:"protocolSettings"`
	ShareDeleteRetentionPolicy *DeleteRetentionPolicyResponse `pulumi:"shareDeleteRetentionPolicy"`
	Sku                        SkuResponse                    `pulumi:"sku"`
	Type                       string                         `pulumi:"type"`
}

func LookupFileServicePropertiesOutput(ctx *pulumi.Context, args LookupFileServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupFileServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileServicePropertiesResult, error) {
			args := v.(LookupFileServicePropertiesArgs)
			r, err := LookupFileServiceProperties(ctx, &args, opts...)
			var s LookupFileServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileServicePropertiesResultOutput)
}

type LookupFileServicePropertiesOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	FileServicesName  pulumi.StringInput `pulumi:"fileServicesName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFileServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileServicePropertiesArgs)(nil)).Elem()
}

// The properties of File services in storage account.
type LookupFileServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupFileServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileServicePropertiesResult)(nil)).Elem()
}

func (o LookupFileServicePropertiesResultOutput) ToLookupFileServicePropertiesResultOutput() LookupFileServicePropertiesResultOutput {
	return o
}

func (o LookupFileServicePropertiesResultOutput) ToLookupFileServicePropertiesResultOutputWithContext(ctx context.Context) LookupFileServicePropertiesResultOutput {
	return o
}

func (o LookupFileServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o LookupFileServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileServicePropertiesResultOutput) ProtocolSettings() ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) *ProtocolSettingsResponse { return v.ProtocolSettings }).(ProtocolSettingsResponsePtrOutput)
}

func (o LookupFileServicePropertiesResultOutput) ShareDeleteRetentionPolicy() DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) *DeleteRetentionPolicyResponse {
		return v.ShareDeleteRetentionPolicy
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

func (o LookupFileServicePropertiesResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupFileServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileServicePropertiesResultOutput{})
}