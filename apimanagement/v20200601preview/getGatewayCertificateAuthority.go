// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gateway certificate authority details.
func LookupGatewayCertificateAuthority(ctx *pulumi.Context, args *LookupGatewayCertificateAuthorityArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCertificateAuthorityResult, error) {
	var rv LookupGatewayCertificateAuthorityResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getGatewayCertificateAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayCertificateAuthorityArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId string `pulumi:"certificateId"`
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Gateway certificate authority details.
type LookupGatewayCertificateAuthorityResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Determines whether certificate authority is trusted.
	IsTrusted *bool `pulumi:"isTrusted"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}

func LookupGatewayCertificateAuthorityOutput(ctx *pulumi.Context, args LookupGatewayCertificateAuthorityOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCertificateAuthorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCertificateAuthorityResult, error) {
			args := v.(LookupGatewayCertificateAuthorityArgs)
			r, err := LookupGatewayCertificateAuthority(ctx, &args, opts...)
			var s LookupGatewayCertificateAuthorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCertificateAuthorityResultOutput)
}

type LookupGatewayCertificateAuthorityOutputArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId pulumi.StringInput `pulumi:"certificateId"`
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayCertificateAuthorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCertificateAuthorityArgs)(nil)).Elem()
}

// Gateway certificate authority details.
type LookupGatewayCertificateAuthorityResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCertificateAuthorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCertificateAuthorityResult)(nil)).Elem()
}

func (o LookupGatewayCertificateAuthorityResultOutput) ToLookupGatewayCertificateAuthorityResultOutput() LookupGatewayCertificateAuthorityResultOutput {
	return o
}

func (o LookupGatewayCertificateAuthorityResultOutput) ToLookupGatewayCertificateAuthorityResultOutputWithContext(ctx context.Context) LookupGatewayCertificateAuthorityResultOutput {
	return o
}

// Resource ID.
func (o LookupGatewayCertificateAuthorityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) string { return v.Id }).(pulumi.StringOutput)
}

// Determines whether certificate authority is trusted.
func (o LookupGatewayCertificateAuthorityResultOutput) IsTrusted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) *bool { return v.IsTrusted }).(pulumi.BoolPtrOutput)
}

// Resource name.
func (o LookupGatewayCertificateAuthorityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type for API Management resource.
func (o LookupGatewayCertificateAuthorityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCertificateAuthorityResultOutput{})
}
