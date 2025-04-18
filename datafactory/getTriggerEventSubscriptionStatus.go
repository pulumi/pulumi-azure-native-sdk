// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a trigger's event subscription status.
//
// Uses Azure REST API version 2018-06-01.
func GetTriggerEventSubscriptionStatus(ctx *pulumi.Context, args *GetTriggerEventSubscriptionStatusArgs, opts ...pulumi.InvokeOption) (*GetTriggerEventSubscriptionStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetTriggerEventSubscriptionStatusResult
	err := ctx.Invoke("azure-native:datafactory:getTriggerEventSubscriptionStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTriggerEventSubscriptionStatusArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The trigger name.
	TriggerName string `pulumi:"triggerName"`
}

// Defines the response of a trigger subscription operation.
type GetTriggerEventSubscriptionStatusResult struct {
	// Event Subscription Status.
	Status string `pulumi:"status"`
	// Trigger name.
	TriggerName string `pulumi:"triggerName"`
}

func GetTriggerEventSubscriptionStatusOutput(ctx *pulumi.Context, args GetTriggerEventSubscriptionStatusOutputArgs, opts ...pulumi.InvokeOption) GetTriggerEventSubscriptionStatusResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTriggerEventSubscriptionStatusResultOutput, error) {
			args := v.(GetTriggerEventSubscriptionStatusArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datafactory:getTriggerEventSubscriptionStatus", args, GetTriggerEventSubscriptionStatusResultOutput{}, options).(GetTriggerEventSubscriptionStatusResultOutput), nil
		}).(GetTriggerEventSubscriptionStatusResultOutput)
}

type GetTriggerEventSubscriptionStatusOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The trigger name.
	TriggerName pulumi.StringInput `pulumi:"triggerName"`
}

func (GetTriggerEventSubscriptionStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTriggerEventSubscriptionStatusArgs)(nil)).Elem()
}

// Defines the response of a trigger subscription operation.
type GetTriggerEventSubscriptionStatusResultOutput struct{ *pulumi.OutputState }

func (GetTriggerEventSubscriptionStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTriggerEventSubscriptionStatusResult)(nil)).Elem()
}

func (o GetTriggerEventSubscriptionStatusResultOutput) ToGetTriggerEventSubscriptionStatusResultOutput() GetTriggerEventSubscriptionStatusResultOutput {
	return o
}

func (o GetTriggerEventSubscriptionStatusResultOutput) ToGetTriggerEventSubscriptionStatusResultOutputWithContext(ctx context.Context) GetTriggerEventSubscriptionStatusResultOutput {
	return o
}

// Event Subscription Status.
func (o GetTriggerEventSubscriptionStatusResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetTriggerEventSubscriptionStatusResult) string { return v.Status }).(pulumi.StringOutput)
}

// Trigger name.
func (o GetTriggerEventSubscriptionStatusResultOutput) TriggerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTriggerEventSubscriptionStatusResult) string { return v.TriggerName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTriggerEventSubscriptionStatusResultOutput{})
}
