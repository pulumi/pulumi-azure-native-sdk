// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an entity query.
//
// Uses Azure REST API version 2025-01-01-preview.
func LookupActivityCustomEntityQuery(ctx *pulumi.Context, args *LookupActivityCustomEntityQueryArgs, opts ...pulumi.InvokeOption) (*LookupActivityCustomEntityQueryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupActivityCustomEntityQueryResult
	err := ctx.Invoke("azure-native:securityinsights:getActivityCustomEntityQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActivityCustomEntityQueryArgs struct {
	// entity query ID
	EntityQueryId string `pulumi:"entityQueryId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Activity entity query.
type LookupActivityCustomEntityQueryResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The entity query content to display in timeline
	Content *string `pulumi:"content"`
	// The time the activity was created
	CreatedTimeUtc string `pulumi:"createdTimeUtc"`
	// The entity query description
	Description *string `pulumi:"description"`
	// Determines whether this activity is enabled or disabled.
	Enabled *bool `pulumi:"enabled"`
	// The query applied only to entities matching to all filters
	EntitiesFilter map[string][]string `pulumi:"entitiesFilter"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The type of the query's source entity
	InputEntityType *string `pulumi:"inputEntityType"`
	// The kind of the entity query
	// Expected value is 'Activity'.
	Kind string `pulumi:"kind"`
	// The last time the activity was updated
	LastModifiedTimeUtc string `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Activity query definitions
	QueryDefinitions *ActivityEntityQueriesPropertiesResponseQueryDefinitions `pulumi:"queryDefinitions"`
	// List of the fields of the source entity that are required to run the query
	RequiredInputFieldsSets [][]string `pulumi:"requiredInputFieldsSets"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The template id this activity was created from
	TemplateName *string `pulumi:"templateName"`
	// The entity query title
	Title *string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupActivityCustomEntityQueryOutput(ctx *pulumi.Context, args LookupActivityCustomEntityQueryOutputArgs, opts ...pulumi.InvokeOption) LookupActivityCustomEntityQueryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupActivityCustomEntityQueryResultOutput, error) {
			args := v.(LookupActivityCustomEntityQueryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:securityinsights:getActivityCustomEntityQuery", args, LookupActivityCustomEntityQueryResultOutput{}, options).(LookupActivityCustomEntityQueryResultOutput), nil
		}).(LookupActivityCustomEntityQueryResultOutput)
}

type LookupActivityCustomEntityQueryOutputArgs struct {
	// entity query ID
	EntityQueryId pulumi.StringInput `pulumi:"entityQueryId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupActivityCustomEntityQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityCustomEntityQueryArgs)(nil)).Elem()
}

// Represents Activity entity query.
type LookupActivityCustomEntityQueryResultOutput struct{ *pulumi.OutputState }

func (LookupActivityCustomEntityQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityCustomEntityQueryResult)(nil)).Elem()
}

func (o LookupActivityCustomEntityQueryResultOutput) ToLookupActivityCustomEntityQueryResultOutput() LookupActivityCustomEntityQueryResultOutput {
	return o
}

func (o LookupActivityCustomEntityQueryResultOutput) ToLookupActivityCustomEntityQueryResultOutputWithContext(ctx context.Context) LookupActivityCustomEntityQueryResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupActivityCustomEntityQueryResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The entity query content to display in timeline
func (o LookupActivityCustomEntityQueryResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// The time the activity was created
func (o LookupActivityCustomEntityQueryResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The entity query description
func (o LookupActivityCustomEntityQueryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Determines whether this activity is enabled or disabled.
func (o LookupActivityCustomEntityQueryResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The query applied only to entities matching to all filters
func (o LookupActivityCustomEntityQueryResultOutput) EntitiesFilter() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) map[string][]string { return v.EntitiesFilter }).(pulumi.StringArrayMapOutput)
}

// Etag of the azure resource
func (o LookupActivityCustomEntityQueryResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupActivityCustomEntityQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The type of the query's source entity
func (o LookupActivityCustomEntityQueryResultOutput) InputEntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.InputEntityType }).(pulumi.StringPtrOutput)
}

// The kind of the entity query
// Expected value is 'Activity'.
func (o LookupActivityCustomEntityQueryResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The last time the activity was updated
func (o LookupActivityCustomEntityQueryResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupActivityCustomEntityQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Activity query definitions
func (o LookupActivityCustomEntityQueryResultOutput) QueryDefinitions() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *ActivityEntityQueriesPropertiesResponseQueryDefinitions {
		return v.QueryDefinitions
	}).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput)
}

// List of the fields of the source entity that are required to run the query
func (o LookupActivityCustomEntityQueryResultOutput) RequiredInputFieldsSets() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) [][]string { return v.RequiredInputFieldsSets }).(pulumi.StringArrayArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupActivityCustomEntityQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The template id this activity was created from
func (o LookupActivityCustomEntityQueryResultOutput) TemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.TemplateName }).(pulumi.StringPtrOutput)
}

// The entity query title
func (o LookupActivityCustomEntityQueryResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupActivityCustomEntityQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityCustomEntityQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActivityCustomEntityQueryResultOutput{})
}
