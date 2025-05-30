// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package management

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves details about given subscription which is associated with the management group.
//
// Uses Azure REST API version 2023-04-01.
//
// Other available API versions: 2021-04-01, 2024-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native management [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupManagementGroupSubscription(ctx *pulumi.Context, args *LookupManagementGroupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementGroupSubscriptionResult
	err := ctx.Invoke("azure-native:management:getManagementGroupSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupSubscriptionArgs struct {
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
	// Subscription ID.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// The details of subscription under management group.
type LookupManagementGroupSubscriptionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The friendly name of the subscription.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified ID for the subscription.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/subscriptions/0000000-0000-0000-0000-000000000001
	Id string `pulumi:"id"`
	// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name string `pulumi:"name"`
	// The ID of the parent management group.
	Parent *DescendantParentGroupInfoResponse `pulumi:"parent"`
	// The state of the subscription.
	State *string `pulumi:"state"`
	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant *string `pulumi:"tenant"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
	Type string `pulumi:"type"`
}

func LookupManagementGroupSubscriptionOutput(ctx *pulumi.Context, args LookupManagementGroupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupSubscriptionResultOutput, error) {
			args := v.(LookupManagementGroupSubscriptionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:management:getManagementGroupSubscription", args, LookupManagementGroupSubscriptionResultOutput{}, options).(LookupManagementGroupSubscriptionResultOutput), nil
		}).(LookupManagementGroupSubscriptionResultOutput)
}

type LookupManagementGroupSubscriptionOutputArgs struct {
	// Management Group ID.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// Subscription ID.
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupManagementGroupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupSubscriptionArgs)(nil)).Elem()
}

// The details of subscription under management group.
type LookupManagementGroupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupSubscriptionResult)(nil)).Elem()
}

func (o LookupManagementGroupSubscriptionResultOutput) ToLookupManagementGroupSubscriptionResultOutput() LookupManagementGroupSubscriptionResultOutput {
	return o
}

func (o LookupManagementGroupSubscriptionResultOutput) ToLookupManagementGroupSubscriptionResultOutputWithContext(ctx context.Context) LookupManagementGroupSubscriptionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupManagementGroupSubscriptionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The friendly name of the subscription.
func (o LookupManagementGroupSubscriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the subscription.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/subscriptions/0000000-0000-0000-0000-000000000001
func (o LookupManagementGroupSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
func (o LookupManagementGroupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the parent management group.
func (o LookupManagementGroupSubscriptionResultOutput) Parent() DescendantParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *DescendantParentGroupInfoResponse { return v.Parent }).(DescendantParentGroupInfoResponsePtrOutput)
}

// The state of the subscription.
func (o LookupManagementGroupSubscriptionResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
func (o LookupManagementGroupSubscriptionResultOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *string { return v.Tenant }).(pulumi.StringPtrOutput)
}

// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
func (o LookupManagementGroupSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupSubscriptionResultOutput{})
}
