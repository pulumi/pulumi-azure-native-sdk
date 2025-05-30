// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alertsmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List all alerts in the issue - this method uses pagination to return all alerts
//
// Uses Azure REST API version 2025-03-01-preview.
func ListIssueAlerts(ctx *pulumi.Context, args *ListIssueAlertsArgs, opts ...pulumi.InvokeOption) (*ListIssueAlertsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListIssueAlertsResult
	err := ctx.Invoke("azure-native:alertsmanagement:listIssueAlerts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIssueAlertsArgs struct {
	// The filter to apply on the operation. For example, to filter by relevance, use "$filter=relevance eq 'Relevant'"
	Filter *string `pulumi:"filter"`
	// The name of the IssueResource
	IssueName string `pulumi:"issueName"`
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// Paged collection of RelatedAlert items
type ListIssueAlertsResult struct {
	// The link to the next page of items
	NextLink *string `pulumi:"nextLink"`
	// The RelatedAlert items on this page
	Value []RelatedAlertResponse `pulumi:"value"`
}

func ListIssueAlertsOutput(ctx *pulumi.Context, args ListIssueAlertsOutputArgs, opts ...pulumi.InvokeOption) ListIssueAlertsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListIssueAlertsResultOutput, error) {
			args := v.(ListIssueAlertsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:alertsmanagement:listIssueAlerts", args, ListIssueAlertsResultOutput{}, options).(ListIssueAlertsResultOutput), nil
		}).(ListIssueAlertsResultOutput)
}

type ListIssueAlertsOutputArgs struct {
	// The filter to apply on the operation. For example, to filter by relevance, use "$filter=relevance eq 'Relevant'"
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The name of the IssueResource
	IssueName pulumi.StringInput `pulumi:"issueName"`
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (ListIssueAlertsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIssueAlertsArgs)(nil)).Elem()
}

// Paged collection of RelatedAlert items
type ListIssueAlertsResultOutput struct{ *pulumi.OutputState }

func (ListIssueAlertsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIssueAlertsResult)(nil)).Elem()
}

func (o ListIssueAlertsResultOutput) ToListIssueAlertsResultOutput() ListIssueAlertsResultOutput {
	return o
}

func (o ListIssueAlertsResultOutput) ToListIssueAlertsResultOutputWithContext(ctx context.Context) ListIssueAlertsResultOutput {
	return o
}

// The link to the next page of items
func (o ListIssueAlertsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIssueAlertsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// The RelatedAlert items on this page
func (o ListIssueAlertsResultOutput) Value() RelatedAlertResponseArrayOutput {
	return o.ApplyT(func(v ListIssueAlertsResult) []RelatedAlertResponse { return v.Value }).(RelatedAlertResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIssueAlertsResultOutput{})
}
