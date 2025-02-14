// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an authorization rule for a namespace by name.
func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	// Authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a Namespace AuthorizationRules.
type LookupNamespaceAuthorizationRuleResult struct {
	// A string that describes the claim type
	ClaimType string `pulumi:"claimType"`
	// A string that describes the claim value
	ClaimValue string `pulumi:"claimValue"`
	// The created time for this rule
	CreatedTime string `pulumi:"createdTime"`
	// Resource Id
	Id string `pulumi:"id"`
	// A string that describes the authorization rule.
	KeyName string `pulumi:"keyName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The last modified time for this rule
	ModifiedTime string `pulumi:"modifiedTime"`
	// Resource name
	Name string `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey string `pulumi:"primaryKey"`
	// The revision number for the rule
	Revision int `pulumi:"revision"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey string `pulumi:"secondaryKey"`
	// The sku of the created namespace
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupNamespaceAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNamespaceAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNamespaceAuthorizationRuleResultOutput, error) {
			args := v.(LookupNamespaceAuthorizationRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:notificationhubs/v20170401:getNamespaceAuthorizationRule", args, LookupNamespaceAuthorizationRuleResultOutput{}, options).(LookupNamespaceAuthorizationRuleResultOutput), nil
		}).(LookupNamespaceAuthorizationRuleResultOutput)
}

type LookupNamespaceAuthorizationRuleOutputArgs struct {
	// Authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleArgs)(nil)).Elem()
}

// Description of a Namespace AuthorizationRules.
type LookupNamespaceAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutput() LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

// A string that describes the claim type
func (o LookupNamespaceAuthorizationRuleResultOutput) ClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.ClaimType }).(pulumi.StringOutput)
}

// A string that describes the claim value
func (o LookupNamespaceAuthorizationRuleResultOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.ClaimValue }).(pulumi.StringOutput)
}

// The created time for this rule
func (o LookupNamespaceAuthorizationRuleResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupNamespaceAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o LookupNamespaceAuthorizationRuleResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Resource location
func (o LookupNamespaceAuthorizationRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The last modified time for this rule
func (o LookupNamespaceAuthorizationRuleResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// Resource name
func (o LookupNamespaceAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o LookupNamespaceAuthorizationRuleResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The revision number for the rule
func (o LookupNamespaceAuthorizationRuleResultOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) int { return v.Revision }).(pulumi.IntOutput)
}

// The rights associated with the rule.
func (o LookupNamespaceAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o LookupNamespaceAuthorizationRuleResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

// The sku of the created namespace
func (o LookupNamespaceAuthorizationRuleResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o LookupNamespaceAuthorizationRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupNamespaceAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceAuthorizationRuleResultOutput{})
}
