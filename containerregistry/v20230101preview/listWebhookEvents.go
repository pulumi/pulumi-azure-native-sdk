// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The result of a request to list events for a webhook.
func ListWebhookEvents(ctx *pulumi.Context, args *ListWebhookEventsArgs, opts ...pulumi.InvokeOption) (*ListWebhookEventsResult, error) {
	var rv ListWebhookEventsResult
	err := ctx.Invoke("azure-native:containerregistry/v20230101preview:listWebhookEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebhookEventsArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName string `pulumi:"webhookName"`
}

// The result of a request to list events for a webhook.
type ListWebhookEventsResult struct {
	// The URI that can be used to request the next list of events.
	NextLink *string `pulumi:"nextLink"`
	// The list of events. Since this list may be incomplete, the nextLink field should be used to request the next list of events.
	Value []EventResponse `pulumi:"value"`
}

func ListWebhookEventsOutput(ctx *pulumi.Context, args ListWebhookEventsOutputArgs, opts ...pulumi.InvokeOption) ListWebhookEventsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebhookEventsResult, error) {
			args := v.(ListWebhookEventsArgs)
			r, err := ListWebhookEvents(ctx, &args, opts...)
			var s ListWebhookEventsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebhookEventsResultOutput)
}

type ListWebhookEventsOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the webhook.
	WebhookName pulumi.StringInput `pulumi:"webhookName"`
}

func (ListWebhookEventsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebhookEventsArgs)(nil)).Elem()
}

// The result of a request to list events for a webhook.
type ListWebhookEventsResultOutput struct{ *pulumi.OutputState }

func (ListWebhookEventsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebhookEventsResult)(nil)).Elem()
}

func (o ListWebhookEventsResultOutput) ToListWebhookEventsResultOutput() ListWebhookEventsResultOutput {
	return o
}

func (o ListWebhookEventsResultOutput) ToListWebhookEventsResultOutputWithContext(ctx context.Context) ListWebhookEventsResultOutput {
	return o
}

// The URI that can be used to request the next list of events.
func (o ListWebhookEventsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebhookEventsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The list of events. Since this list may be incomplete, the nextLink field should be used to request the next list of events.
func (o ListWebhookEventsResultOutput) Value() EventResponseArrayOutput {
	return o.ApplyT(func(v ListWebhookEventsResult) []EventResponse { return v.Value }).(EventResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebhookEventsResultOutput{})
}