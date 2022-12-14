// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy Contract details.
func LookupApiPolicy(ctx *pulumi.Context, args *LookupApiPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiPolicyResult, error) {
	var rv LookupApiPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:getApiPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Policy Export Format.
	Format *string `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Policy Contract details.
type LookupApiPolicyResult struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// Defaults sets the appropriate defaults for LookupApiPolicyResult
func (val *LookupApiPolicyResult) Defaults() *LookupApiPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Format) {
		format_ := "xml"
		tmp.Format = &format_
	}
	return &tmp
}

func LookupApiPolicyOutput(ctx *pulumi.Context, args LookupApiPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupApiPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiPolicyResult, error) {
			args := v.(LookupApiPolicyArgs)
			r, err := LookupApiPolicy(ctx, &args, opts...)
			var s LookupApiPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiPolicyResultOutput)
}

type LookupApiPolicyOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Policy Export Format.
	Format pulumi.StringPtrInput `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId pulumi.StringInput `pulumi:"policyId"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPolicyArgs)(nil)).Elem()
}

// Policy Contract details.
type LookupApiPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupApiPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPolicyResult)(nil)).Elem()
}

func (o LookupApiPolicyResultOutput) ToLookupApiPolicyResultOutput() LookupApiPolicyResultOutput {
	return o
}

func (o LookupApiPolicyResultOutput) ToLookupApiPolicyResultOutputWithContext(ctx context.Context) LookupApiPolicyResultOutput {
	return o
}

// Format of the policyContent.
func (o LookupApiPolicyResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupApiPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupApiPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type for API Management resource.
func (o LookupApiPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o LookupApiPolicyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPolicyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiPolicyResultOutput{})
}
