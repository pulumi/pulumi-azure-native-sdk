// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Result of the Get delivery attributes operation.
// API Version: 2021-10-15-preview.
func GetTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetTopicEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid:getTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTopicEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}

// Result of the Get delivery attributes operation.
type GetTopicEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}

func GetTopicEventSubscriptionDeliveryAttributesOutput(ctx *pulumi.Context, args GetTopicEventSubscriptionDeliveryAttributesOutputArgs, opts ...pulumi.InvokeOption) GetTopicEventSubscriptionDeliveryAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTopicEventSubscriptionDeliveryAttributesResult, error) {
			args := v.(GetTopicEventSubscriptionDeliveryAttributesArgs)
			r, err := GetTopicEventSubscriptionDeliveryAttributes(ctx, &args, opts...)
			var s GetTopicEventSubscriptionDeliveryAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTopicEventSubscriptionDeliveryAttributesResultOutput)
}

type GetTopicEventSubscriptionDeliveryAttributesOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (GetTopicEventSubscriptionDeliveryAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicEventSubscriptionDeliveryAttributesArgs)(nil)).Elem()
}

// Result of the Get delivery attributes operation.
type GetTopicEventSubscriptionDeliveryAttributesResultOutput struct{ *pulumi.OutputState }

func (GetTopicEventSubscriptionDeliveryAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicEventSubscriptionDeliveryAttributesResult)(nil)).Elem()
}

func (o GetTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetTopicEventSubscriptionDeliveryAttributesResultOutput() GetTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetTopicEventSubscriptionDeliveryAttributesResultOutput) ToGetTopicEventSubscriptionDeliveryAttributesResultOutputWithContext(ctx context.Context) GetTopicEventSubscriptionDeliveryAttributesResultOutput {
	return o
}

func (o GetTopicEventSubscriptionDeliveryAttributesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetTopicEventSubscriptionDeliveryAttributesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTopicEventSubscriptionDeliveryAttributesResultOutput{})
}