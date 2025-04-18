// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get notification channel.
//
// Uses Azure REST API version 2018-09-15.
func LookupNotificationChannel(ctx *pulumi.Context, args *LookupNotificationChannelArgs, opts ...pulumi.InvokeOption) (*LookupNotificationChannelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationChannelResult
	err := ctx.Invoke("azure-native:devtestlab:getNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationChannelArgs struct {
	// Specify the $expand query. Example: 'properties($select=webHookUrl)'
	Expand *string `pulumi:"expand"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the notification channel.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A notification.
type LookupNotificationChannelResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The creation date of the notification channel.
	CreatedDate string `pulumi:"createdDate"`
	// Description of notification.
	Description *string `pulumi:"description"`
	// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
	EmailRecipient *string `pulumi:"emailRecipient"`
	// The list of event for which this notification is enabled.
	Events []EventResponse `pulumi:"events"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The locale to use when sending a notification (fallback for unsupported languages is EN).
	NotificationLocale *string `pulumi:"notificationLocale"`
	// The provisioning status of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
	// The webhook URL to send notifications to.
	WebHookUrl *string `pulumi:"webHookUrl"`
}

func LookupNotificationChannelOutput(ctx *pulumi.Context, args LookupNotificationChannelOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationChannelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNotificationChannelResultOutput, error) {
			args := v.(LookupNotificationChannelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devtestlab:getNotificationChannel", args, LookupNotificationChannelResultOutput{}, options).(LookupNotificationChannelResultOutput), nil
		}).(LookupNotificationChannelResultOutput)
}

type LookupNotificationChannelOutputArgs struct {
	// Specify the $expand query. Example: 'properties($select=webHookUrl)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the notification channel.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelArgs)(nil)).Elem()
}

// A notification.
type LookupNotificationChannelResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelResult)(nil)).Elem()
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutput() LookupNotificationChannelResultOutput {
	return o
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutputWithContext(ctx context.Context) LookupNotificationChannelResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupNotificationChannelResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The creation date of the notification channel.
func (o LookupNotificationChannelResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// Description of notification.
func (o LookupNotificationChannelResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
func (o LookupNotificationChannelResultOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.EmailRecipient }).(pulumi.StringPtrOutput)
}

// The list of event for which this notification is enabled.
func (o LookupNotificationChannelResultOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) []EventResponse { return v.Events }).(EventResponseArrayOutput)
}

// The identifier of the resource.
func (o LookupNotificationChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource.
func (o LookupNotificationChannelResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupNotificationChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The locale to use when sending a notification (fallback for unsupported languages is EN).
func (o LookupNotificationChannelResultOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.NotificationLocale }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o LookupNotificationChannelResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o LookupNotificationChannelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupNotificationChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupNotificationChannelResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// The webhook URL to send notifications to.
func (o LookupNotificationChannelResultOutput) WebHookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.WebHookUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationChannelResultOutput{})
}
