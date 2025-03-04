// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an authorization rule for a NotificationHub by name.
func LookupNotificationHubAuthorizationRule(ctx *pulumi.Context, args *LookupNotificationHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:getNotificationHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubAuthorizationRuleArgs struct {
	// authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The notification hub name.
	NotificationHubName string `pulumi:"notificationHubName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a Namespace AuthorizationRules.
type LookupNotificationHubAuthorizationRuleResult struct {
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

func LookupNotificationHubAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNotificationHubAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubAuthorizationRuleResultOutput, error) {
			args := v.(LookupNotificationHubAuthorizationRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:notificationhubs/v20170401:getNotificationHubAuthorizationRule", args, LookupNotificationHubAuthorizationRuleResultOutput{}, options).(LookupNotificationHubAuthorizationRuleResultOutput), nil
		}).(LookupNotificationHubAuthorizationRuleResultOutput)
}

type LookupNotificationHubAuthorizationRuleOutputArgs struct {
	// authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The notification hub name.
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubAuthorizationRuleArgs)(nil)).Elem()
}

// Description of a Namespace AuthorizationRules.
type LookupNotificationHubAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToLookupNotificationHubAuthorizationRuleResultOutput() LookupNotificationHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToLookupNotificationHubAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNotificationHubAuthorizationRuleResultOutput {
	return o
}

// A string that describes the claim type
func (o LookupNotificationHubAuthorizationRuleResultOutput) ClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.ClaimType }).(pulumi.StringOutput)
}

// A string that describes the claim value
func (o LookupNotificationHubAuthorizationRuleResultOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.ClaimValue }).(pulumi.StringOutput)
}

// The created time for this rule
func (o LookupNotificationHubAuthorizationRuleResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupNotificationHubAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o LookupNotificationHubAuthorizationRuleResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Resource location
func (o LookupNotificationHubAuthorizationRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The last modified time for this rule
func (o LookupNotificationHubAuthorizationRuleResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// Resource name
func (o LookupNotificationHubAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o LookupNotificationHubAuthorizationRuleResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The revision number for the rule
func (o LookupNotificationHubAuthorizationRuleResultOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) int { return v.Revision }).(pulumi.IntOutput)
}

// The rights associated with the rule.
func (o LookupNotificationHubAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o LookupNotificationHubAuthorizationRuleResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

// The sku of the created namespace
func (o LookupNotificationHubAuthorizationRuleResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o LookupNotificationHubAuthorizationRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupNotificationHubAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubAuthorizationRuleResultOutput{})
}
