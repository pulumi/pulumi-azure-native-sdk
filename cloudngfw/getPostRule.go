// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudngfw

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a PostRulesResource
//
// Uses Azure REST API version 2025-02-06-preview.
//
// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPostRule(ctx *pulumi.Context, args *LookupPostRuleArgs, opts ...pulumi.InvokeOption) (*LookupPostRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPostRuleResult
	err := ctx.Invoke("azure-native:cloudngfw:getPostRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPostRuleArgs struct {
	// GlobalRulestack resource name
	GlobalRulestackName string `pulumi:"globalRulestackName"`
	// Post Rule priority
	Priority string `pulumi:"priority"`
}

// PostRulestack rule list
type LookupPostRuleResult struct {
	// rule action
	ActionType *string `pulumi:"actionType"`
	// array of rule applications
	Applications []string `pulumi:"applications"`
	// rule comment
	AuditComment *string `pulumi:"auditComment"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// rule category
	Category *CategoryResponse `pulumi:"category"`
	// enable or disable decryption
	DecryptionRuleType *string `pulumi:"decryptionRuleType"`
	// rule description
	Description *string `pulumi:"description"`
	// destination address
	Destination *DestinationAddrResponse `pulumi:"destination"`
	// enable or disable logging
	EnableLogging *string `pulumi:"enableLogging"`
	// etag info
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// inbound Inspection Certificate
	InboundInspectionCertificate *string `pulumi:"inboundInspectionCertificate"`
	// The name of the resource
	Name string `pulumi:"name"`
	// cidr should not be 'any'
	NegateDestination *string `pulumi:"negateDestination"`
	// cidr should not be 'any'
	NegateSource *string `pulumi:"negateSource"`
	Priority     int     `pulumi:"priority"`
	// any, application-default, TCP:number, UDP:number
	Protocol *string `pulumi:"protocol"`
	// prot port list
	ProtocolPortList []string `pulumi:"protocolPortList"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// rule name
	RuleName string `pulumi:"ruleName"`
	// state of this rule
	RuleState *string `pulumi:"ruleState"`
	// source address
	Source *SourceAddrResponse `pulumi:"source"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// tag for rule
	Tags []TagInfoResponse `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPostRuleResult
func (val *LookupPostRuleResult) Defaults() *LookupPostRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Protocol == nil {
		protocol_ := "application-default"
		tmp.Protocol = &protocol_
	}
	return &tmp
}
func LookupPostRuleOutput(ctx *pulumi.Context, args LookupPostRuleOutputArgs, opts ...pulumi.InvokeOption) LookupPostRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPostRuleResultOutput, error) {
			args := v.(LookupPostRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cloudngfw:getPostRule", args, LookupPostRuleResultOutput{}, options).(LookupPostRuleResultOutput), nil
		}).(LookupPostRuleResultOutput)
}

type LookupPostRuleOutputArgs struct {
	// GlobalRulestack resource name
	GlobalRulestackName pulumi.StringInput `pulumi:"globalRulestackName"`
	// Post Rule priority
	Priority pulumi.StringInput `pulumi:"priority"`
}

func (LookupPostRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPostRuleArgs)(nil)).Elem()
}

// PostRulestack rule list
type LookupPostRuleResultOutput struct{ *pulumi.OutputState }

func (LookupPostRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPostRuleResult)(nil)).Elem()
}

func (o LookupPostRuleResultOutput) ToLookupPostRuleResultOutput() LookupPostRuleResultOutput {
	return o
}

func (o LookupPostRuleResultOutput) ToLookupPostRuleResultOutputWithContext(ctx context.Context) LookupPostRuleResultOutput {
	return o
}

// rule action
func (o LookupPostRuleResultOutput) ActionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.ActionType }).(pulumi.StringPtrOutput)
}

// array of rule applications
func (o LookupPostRuleResultOutput) Applications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPostRuleResult) []string { return v.Applications }).(pulumi.StringArrayOutput)
}

// rule comment
func (o LookupPostRuleResultOutput) AuditComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.AuditComment }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupPostRuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostRuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// rule category
func (o LookupPostRuleResultOutput) Category() CategoryResponsePtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *CategoryResponse { return v.Category }).(CategoryResponsePtrOutput)
}

// enable or disable decryption
func (o LookupPostRuleResultOutput) DecryptionRuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.DecryptionRuleType }).(pulumi.StringPtrOutput)
}

// rule description
func (o LookupPostRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// destination address
func (o LookupPostRuleResultOutput) Destination() DestinationAddrResponsePtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *DestinationAddrResponse { return v.Destination }).(DestinationAddrResponsePtrOutput)
}

// enable or disable logging
func (o LookupPostRuleResultOutput) EnableLogging() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.EnableLogging }).(pulumi.StringPtrOutput)
}

// etag info
func (o LookupPostRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPostRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// inbound Inspection Certificate
func (o LookupPostRuleResultOutput) InboundInspectionCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.InboundInspectionCertificate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupPostRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// cidr should not be 'any'
func (o LookupPostRuleResultOutput) NegateDestination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.NegateDestination }).(pulumi.StringPtrOutput)
}

// cidr should not be 'any'
func (o LookupPostRuleResultOutput) NegateSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.NegateSource }).(pulumi.StringPtrOutput)
}

func (o LookupPostRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPostRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

// any, application-default, TCP:number, UDP:number
func (o LookupPostRuleResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// prot port list
func (o LookupPostRuleResultOutput) ProtocolPortList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPostRuleResult) []string { return v.ProtocolPortList }).(pulumi.StringArrayOutput)
}

// Provisioning state of the resource.
func (o LookupPostRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// rule name
func (o LookupPostRuleResultOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostRuleResult) string { return v.RuleName }).(pulumi.StringOutput)
}

// state of this rule
func (o LookupPostRuleResultOutput) RuleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *string { return v.RuleState }).(pulumi.StringPtrOutput)
}

// source address
func (o LookupPostRuleResultOutput) Source() SourceAddrResponsePtrOutput {
	return o.ApplyT(func(v LookupPostRuleResult) *SourceAddrResponse { return v.Source }).(SourceAddrResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPostRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPostRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// tag for rule
func (o LookupPostRuleResultOutput) Tags() TagInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupPostRuleResult) []TagInfoResponse { return v.Tags }).(TagInfoResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPostRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPostRuleResultOutput{})
}
