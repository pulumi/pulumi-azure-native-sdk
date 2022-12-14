// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in List or Get Event Hub operation
// API Version: 2017-04-01.
func LookupEventHub(ctx *pulumi.Context, args *LookupEventHubArgs, opts ...pulumi.InvokeOption) (*LookupEventHubResult, error) {
	var rv LookupEventHubResult
	err := ctx.Invoke("azure-native:eventhub:getEventHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubArgs struct {
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in List or Get Event Hub operation
type LookupEventHubResult struct {
	// Properties of capture description
	CaptureDescription *CaptureDescriptionResponse `pulumi:"captureDescription"`
	// Exact time the Event Hub was created.
	CreatedAt string `pulumi:"createdAt"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *float64 `pulumi:"messageRetentionInDays"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *float64 `pulumi:"partitionCount"`
	// Current number of shards on the Event Hub.
	PartitionIds []string `pulumi:"partitionIds"`
	// Enumerates the possible values for the status of the Event Hub.
	Status *string `pulumi:"status"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupEventHubOutput(ctx *pulumi.Context, args LookupEventHubOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubResult, error) {
			args := v.(LookupEventHubArgs)
			r, err := LookupEventHub(ctx, &args, opts...)
			var s LookupEventHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubResultOutput)
}

type LookupEventHubOutputArgs struct {
	// The Event Hub name
	EventHubName pulumi.StringInput `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubArgs)(nil)).Elem()
}

// Single item in List or Get Event Hub operation
type LookupEventHubResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubResult)(nil)).Elem()
}

func (o LookupEventHubResultOutput) ToLookupEventHubResultOutput() LookupEventHubResultOutput {
	return o
}

func (o LookupEventHubResultOutput) ToLookupEventHubResultOutputWithContext(ctx context.Context) LookupEventHubResultOutput {
	return o
}

// Properties of capture description
func (o LookupEventHubResultOutput) CaptureDescription() CaptureDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *CaptureDescriptionResponse { return v.CaptureDescription }).(CaptureDescriptionResponsePtrOutput)
}

// Exact time the Event Hub was created.
func (o LookupEventHubResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEventHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.Id }).(pulumi.StringOutput)
}

// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
func (o LookupEventHubResultOutput) MessageRetentionInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *float64 { return v.MessageRetentionInDays }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o LookupEventHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
func (o LookupEventHubResultOutput) PartitionCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *float64 { return v.PartitionCount }).(pulumi.Float64PtrOutput)
}

// Current number of shards on the Event Hub.
func (o LookupEventHubResultOutput) PartitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventHubResult) []string { return v.PartitionIds }).(pulumi.StringArrayOutput)
}

// Enumerates the possible values for the status of the Event Hub.
func (o LookupEventHubResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEventHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.Type }).(pulumi.StringOutput)
}

// The exact time the message was updated.
func (o LookupEventHubResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubResultOutput{})
}
