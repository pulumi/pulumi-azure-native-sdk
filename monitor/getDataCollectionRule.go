// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of ARM tracked top level resource.
//
// Uses Azure REST API version 2022-06-01.
func LookupDataCollectionRule(ctx *pulumi.Context, args *LookupDataCollectionRuleArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataCollectionRuleResult
	err := ctx.Invoke("azure-native:monitor:getDataCollectionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionRuleArgs struct {
	// The name of the data collection rule. The name is case insensitive.
	DataCollectionRuleName string `pulumi:"dataCollectionRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of ARM tracked top level resource.
type LookupDataCollectionRuleResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The resource ID of the data collection endpoint that this rule can be used with.
	DataCollectionEndpointId *string `pulumi:"dataCollectionEndpointId"`
	// The specification of data flows.
	DataFlows []DataFlowResponse `pulumi:"dataFlows"`
	// The specification of data sources.
	// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
	DataSources *DataCollectionRuleResponseDataSources `pulumi:"dataSources"`
	// Description of the data collection rule.
	Description *string `pulumi:"description"`
	// The specification of destinations.
	Destinations *DataCollectionRuleResponseDestinations `pulumi:"destinations"`
	// Resource entity tag (ETag).
	Etag string `pulumi:"etag"`
	// Fully qualified ID of the resource.
	Id string `pulumi:"id"`
	// Managed service identity of the resource.
	Identity *DataCollectionRuleResourceResponseIdentity `pulumi:"identity"`
	// The immutable ID of this data collection rule. This property is READ-ONLY.
	ImmutableId string `pulumi:"immutableId"`
	// The kind of the resource.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives.
	Location string `pulumi:"location"`
	// Metadata about the resource
	Metadata DataCollectionRuleResponseMetadata `pulumi:"metadata"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Declaration of custom streams used in this rule.
	StreamDeclarations map[string]StreamDeclarationResponse `pulumi:"streamDeclarations"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData DataCollectionRuleResourceResponseSystemData `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupDataCollectionRuleOutput(ctx *pulumi.Context, args LookupDataCollectionRuleOutputArgs, opts ...pulumi.InvokeOption) LookupDataCollectionRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataCollectionRuleResultOutput, error) {
			args := v.(LookupDataCollectionRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:monitor:getDataCollectionRule", args, LookupDataCollectionRuleResultOutput{}, options).(LookupDataCollectionRuleResultOutput), nil
		}).(LookupDataCollectionRuleResultOutput)
}

type LookupDataCollectionRuleOutputArgs struct {
	// The name of the data collection rule. The name is case insensitive.
	DataCollectionRuleName pulumi.StringInput `pulumi:"dataCollectionRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataCollectionRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleArgs)(nil)).Elem()
}

// Definition of ARM tracked top level resource.
type LookupDataCollectionRuleResultOutput struct{ *pulumi.OutputState }

func (LookupDataCollectionRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleResult)(nil)).Elem()
}

func (o LookupDataCollectionRuleResultOutput) ToLookupDataCollectionRuleResultOutput() LookupDataCollectionRuleResultOutput {
	return o
}

func (o LookupDataCollectionRuleResultOutput) ToLookupDataCollectionRuleResultOutputWithContext(ctx context.Context) LookupDataCollectionRuleResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDataCollectionRuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The resource ID of the data collection endpoint that this rule can be used with.
func (o LookupDataCollectionRuleResultOutput) DataCollectionEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *string { return v.DataCollectionEndpointId }).(pulumi.StringPtrOutput)
}

// The specification of data flows.
func (o LookupDataCollectionRuleResultOutput) DataFlows() DataFlowResponseArrayOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) []DataFlowResponse { return v.DataFlows }).(DataFlowResponseArrayOutput)
}

// The specification of data sources.
// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
func (o LookupDataCollectionRuleResultOutput) DataSources() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *DataCollectionRuleResponseDataSources { return v.DataSources }).(DataCollectionRuleResponseDataSourcesPtrOutput)
}

// Description of the data collection rule.
func (o LookupDataCollectionRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The specification of destinations.
func (o LookupDataCollectionRuleResultOutput) Destinations() DataCollectionRuleResponseDestinationsPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *DataCollectionRuleResponseDestinations { return v.Destinations }).(DataCollectionRuleResponseDestinationsPtrOutput)
}

// Resource entity tag (ETag).
func (o LookupDataCollectionRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified ID of the resource.
func (o LookupDataCollectionRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity of the resource.
func (o LookupDataCollectionRuleResultOutput) Identity() DataCollectionRuleResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *DataCollectionRuleResourceResponseIdentity { return v.Identity }).(DataCollectionRuleResourceResponseIdentityPtrOutput)
}

// The immutable ID of this data collection rule. This property is READ-ONLY.
func (o LookupDataCollectionRuleResultOutput) ImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.ImmutableId }).(pulumi.StringOutput)
}

// The kind of the resource.
func (o LookupDataCollectionRuleResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives.
func (o LookupDataCollectionRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// Metadata about the resource
func (o LookupDataCollectionRuleResultOutput) Metadata() DataCollectionRuleResponseMetadataOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) DataCollectionRuleResponseMetadata { return v.Metadata }).(DataCollectionRuleResponseMetadataOutput)
}

// The name of the resource.
func (o LookupDataCollectionRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource provisioning state.
func (o LookupDataCollectionRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Declaration of custom streams used in this rule.
func (o LookupDataCollectionRuleResultOutput) StreamDeclarations() StreamDeclarationResponseMapOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) map[string]StreamDeclarationResponse {
		return v.StreamDeclarations
	}).(StreamDeclarationResponseMapOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDataCollectionRuleResultOutput) SystemData() DataCollectionRuleResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) DataCollectionRuleResourceResponseSystemData {
		return v.SystemData
	}).(DataCollectionRuleResourceResponseSystemDataOutput)
}

// Resource tags.
func (o LookupDataCollectionRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupDataCollectionRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCollectionRuleResultOutput{})
}
