// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Full endpoint url of an event subscription
func GetSystemTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetSystemTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetSystemTopicEventSubscriptionFullUrlResult, error) {
	var rv GetSystemTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20211015preview:getSystemTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSystemTopicEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SystemTopicName       string `pulumi:"systemTopicName"`
}

// Full endpoint url of an event subscription
type GetSystemTopicEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}

func GetSystemTopicEventSubscriptionFullUrlOutput(ctx *pulumi.Context, args GetSystemTopicEventSubscriptionFullUrlOutputArgs, opts ...pulumi.InvokeOption) GetSystemTopicEventSubscriptionFullUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemTopicEventSubscriptionFullUrlResult, error) {
			args := v.(GetSystemTopicEventSubscriptionFullUrlArgs)
			r, err := GetSystemTopicEventSubscriptionFullUrl(ctx, &args, opts...)
			var s GetSystemTopicEventSubscriptionFullUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemTopicEventSubscriptionFullUrlResultOutput)
}

type GetSystemTopicEventSubscriptionFullUrlOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SystemTopicName       pulumi.StringInput `pulumi:"systemTopicName"`
}

func (GetSystemTopicEventSubscriptionFullUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionFullUrlArgs)(nil)).Elem()
}

// Full endpoint url of an event subscription
type GetSystemTopicEventSubscriptionFullUrlResultOutput struct{ *pulumi.OutputState }

func (GetSystemTopicEventSubscriptionFullUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTopicEventSubscriptionFullUrlResult)(nil)).Elem()
}

func (o GetSystemTopicEventSubscriptionFullUrlResultOutput) ToGetSystemTopicEventSubscriptionFullUrlResultOutput() GetSystemTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionFullUrlResultOutput) ToGetSystemTopicEventSubscriptionFullUrlResultOutputWithContext(ctx context.Context) GetSystemTopicEventSubscriptionFullUrlResultOutput {
	return o
}

func (o GetSystemTopicEventSubscriptionFullUrlResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemTopicEventSubscriptionFullUrlResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemTopicEventSubscriptionFullUrlResultOutput{})
}