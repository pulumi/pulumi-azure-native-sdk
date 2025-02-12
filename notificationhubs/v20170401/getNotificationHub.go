// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the notification hubs associated with a namespace.
func LookupNotificationHub(ctx *pulumi.Context, args *LookupNotificationHubArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationHubResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:getNotificationHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubArgs struct {
	// The namespace name.
	NamespaceName string `pulumi:"namespaceName"`
	// The notification hub name.
	NotificationHubName string `pulumi:"notificationHubName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a NotificationHub Resource.
type LookupNotificationHubResult struct {
	// The AdmCredential of the created NotificationHub
	AdmCredential *AdmCredentialResponse `pulumi:"admCredential"`
	// The ApnsCredential of the created NotificationHub
	ApnsCredential *ApnsCredentialResponse `pulumi:"apnsCredential"`
	// The AuthorizationRules of the created NotificationHub
	AuthorizationRules []SharedAccessAuthorizationRulePropertiesResponse `pulumi:"authorizationRules"`
	// The BaiduCredential of the created NotificationHub
	BaiduCredential *BaiduCredentialResponse `pulumi:"baiduCredential"`
	// The GcmCredential of the created NotificationHub
	GcmCredential *GcmCredentialResponse `pulumi:"gcmCredential"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// The MpnsCredential of the created NotificationHub
	MpnsCredential *MpnsCredentialResponse `pulumi:"mpnsCredential"`
	// Resource name
	Name string `pulumi:"name"`
	// The RegistrationTtl of the created NotificationHub
	RegistrationTtl *string `pulumi:"registrationTtl"`
	// The sku of the created namespace
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// The WnsCredential of the created NotificationHub
	WnsCredential *WnsCredentialResponse `pulumi:"wnsCredential"`
}

func LookupNotificationHubOutput(ctx *pulumi.Context, args LookupNotificationHubOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubResultOutput, error) {
			args := v.(LookupNotificationHubArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:notificationhubs/v20170401:getNotificationHub", args, LookupNotificationHubResultOutput{}, options).(LookupNotificationHubResultOutput), nil
		}).(LookupNotificationHubResultOutput)
}

type LookupNotificationHubOutputArgs struct {
	// The namespace name.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The notification hub name.
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubArgs)(nil)).Elem()
}

// Description of a NotificationHub Resource.
type LookupNotificationHubResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubResult)(nil)).Elem()
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutput() LookupNotificationHubResultOutput {
	return o
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutputWithContext(ctx context.Context) LookupNotificationHubResultOutput {
	return o
}

// The AdmCredential of the created NotificationHub
func (o LookupNotificationHubResultOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *AdmCredentialResponse { return v.AdmCredential }).(AdmCredentialResponsePtrOutput)
}

// The ApnsCredential of the created NotificationHub
func (o LookupNotificationHubResultOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *ApnsCredentialResponse { return v.ApnsCredential }).(ApnsCredentialResponsePtrOutput)
}

// The AuthorizationRules of the created NotificationHub
func (o LookupNotificationHubResultOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) []SharedAccessAuthorizationRulePropertiesResponse {
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

// The BaiduCredential of the created NotificationHub
func (o LookupNotificationHubResultOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *BaiduCredentialResponse { return v.BaiduCredential }).(BaiduCredentialResponsePtrOutput)
}

// The GcmCredential of the created NotificationHub
func (o LookupNotificationHubResultOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *GcmCredentialResponse { return v.GcmCredential }).(GcmCredentialResponsePtrOutput)
}

// Resource Id
func (o LookupNotificationHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupNotificationHubResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The MpnsCredential of the created NotificationHub
func (o LookupNotificationHubResultOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *MpnsCredentialResponse { return v.MpnsCredential }).(MpnsCredentialResponsePtrOutput)
}

// Resource name
func (o LookupNotificationHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// The RegistrationTtl of the created NotificationHub
func (o LookupNotificationHubResultOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.RegistrationTtl }).(pulumi.StringPtrOutput)
}

// The sku of the created namespace
func (o LookupNotificationHubResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o LookupNotificationHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupNotificationHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Type }).(pulumi.StringOutput)
}

// The WnsCredential of the created NotificationHub
func (o LookupNotificationHubResultOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *WnsCredentialResponse { return v.WnsCredential }).(WnsCredentialResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubResultOutput{})
}
