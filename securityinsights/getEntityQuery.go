// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specific entity query.
// API Version: 2021-03-01-preview.
//
// Deprecated: Please use one of the variants: ActivityCustomEntityQuery.
func LookupEntityQuery(ctx *pulumi.Context, args *LookupEntityQueryArgs, opts ...pulumi.InvokeOption) (*LookupEntityQueryResult, error) {
	var rv LookupEntityQueryResult
	err := ctx.Invoke("azure-native:securityinsights:getEntityQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityQueryArgs struct {
	// entity query ID
	EntityQueryId string `pulumi:"entityQueryId"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Specific entity query.
type LookupEntityQueryResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// the entity query kind
	Kind string `pulumi:"kind"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource type
	Type string `pulumi:"type"`
}

func LookupEntityQueryOutput(ctx *pulumi.Context, args LookupEntityQueryOutputArgs, opts ...pulumi.InvokeOption) LookupEntityQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntityQueryResult, error) {
			args := v.(LookupEntityQueryArgs)
			r, err := LookupEntityQuery(ctx, &args, opts...)
			var s LookupEntityQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEntityQueryResultOutput)
}

type LookupEntityQueryOutputArgs struct {
	// entity query ID
	EntityQueryId pulumi.StringInput `pulumi:"entityQueryId"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEntityQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityQueryArgs)(nil)).Elem()
}

// Specific entity query.
type LookupEntityQueryResultOutput struct{ *pulumi.OutputState }

func (LookupEntityQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityQueryResult)(nil)).Elem()
}

func (o LookupEntityQueryResultOutput) ToLookupEntityQueryResultOutput() LookupEntityQueryResultOutput {
	return o
}

func (o LookupEntityQueryResultOutput) ToLookupEntityQueryResultOutputWithContext(ctx context.Context) LookupEntityQueryResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupEntityQueryResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Azure resource Id
func (o LookupEntityQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

// the entity query kind
func (o LookupEntityQueryResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Azure resource name
func (o LookupEntityQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEntityQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource type
func (o LookupEntityQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntityQueryResultOutput{})
}
