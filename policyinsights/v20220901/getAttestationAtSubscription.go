// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An attestation resource.
func LookupAttestationAtSubscription(ctx *pulumi.Context, args *LookupAttestationAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAttestationAtSubscriptionResult, error) {
	var rv LookupAttestationAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights/v20220901:getAttestationAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationAtSubscriptionArgs struct {
	// The name of the attestation.
	AttestationName string `pulumi:"attestationName"`
}

// An attestation resource.
type LookupAttestationAtSubscriptionResult struct {
	// The time the evidence was assessed
	AssessmentDate *string `pulumi:"assessmentDate"`
	// Comments describing why this attestation was created.
	Comments *string `pulumi:"comments"`
	// The compliance state that should be set on the resource.
	ComplianceState *string `pulumi:"complianceState"`
	// The evidence supporting the compliance state set in this attestation.
	Evidence []AttestationEvidenceResponse `pulumi:"evidence"`
	// The time the compliance state should expire.
	ExpiresOn *string `pulumi:"expiresOn"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The time the compliance state was last changed in this attestation.
	LastComplianceStateChangeAt string `pulumi:"lastComplianceStateChangeAt"`
	// Additional metadata for this attestation
	Metadata interface{} `pulumi:"metadata"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object ID.
	Owner *string `pulumi:"owner"`
	// The resource ID of the policy assignment that the attestation is setting the state for.
	PolicyAssignmentId string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy assignment assigns a policy set definition the attestation can choose a definition within the set definition with this property or omit this and set the state for the entire set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The status of the attestation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAttestationAtSubscriptionOutput(ctx *pulumi.Context, args LookupAttestationAtSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAttestationAtSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttestationAtSubscriptionResult, error) {
			args := v.(LookupAttestationAtSubscriptionArgs)
			r, err := LookupAttestationAtSubscription(ctx, &args, opts...)
			var s LookupAttestationAtSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttestationAtSubscriptionResultOutput)
}

type LookupAttestationAtSubscriptionOutputArgs struct {
	// The name of the attestation.
	AttestationName pulumi.StringInput `pulumi:"attestationName"`
}

func (LookupAttestationAtSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtSubscriptionArgs)(nil)).Elem()
}

// An attestation resource.
type LookupAttestationAtSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAttestationAtSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationAtSubscriptionResult)(nil)).Elem()
}

func (o LookupAttestationAtSubscriptionResultOutput) ToLookupAttestationAtSubscriptionResultOutput() LookupAttestationAtSubscriptionResultOutput {
	return o
}

func (o LookupAttestationAtSubscriptionResultOutput) ToLookupAttestationAtSubscriptionResultOutputWithContext(ctx context.Context) LookupAttestationAtSubscriptionResultOutput {
	return o
}

// The time the evidence was assessed
func (o LookupAttestationAtSubscriptionResultOutput) AssessmentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.AssessmentDate }).(pulumi.StringPtrOutput)
}

// Comments describing why this attestation was created.
func (o LookupAttestationAtSubscriptionResultOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

// The compliance state that should be set on the resource.
func (o LookupAttestationAtSubscriptionResultOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

// The evidence supporting the compliance state set in this attestation.
func (o LookupAttestationAtSubscriptionResultOutput) Evidence() AttestationEvidenceResponseArrayOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) []AttestationEvidenceResponse { return v.Evidence }).(AttestationEvidenceResponseArrayOutput)
}

// The time the compliance state should expire.
func (o LookupAttestationAtSubscriptionResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAttestationAtSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time the compliance state was last changed in this attestation.
func (o LookupAttestationAtSubscriptionResultOutput) LastComplianceStateChangeAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.LastComplianceStateChangeAt }).(pulumi.StringOutput)
}

// Additional metadata for this attestation
func (o LookupAttestationAtSubscriptionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the resource
func (o LookupAttestationAtSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object ID.
func (o LookupAttestationAtSubscriptionResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

// The resource ID of the policy assignment that the attestation is setting the state for.
func (o LookupAttestationAtSubscriptionResultOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy assignment assigns a policy set definition the attestation can choose a definition within the set definition with this property or omit this and set the state for the entire set definition.
func (o LookupAttestationAtSubscriptionResultOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the attestation.
func (o LookupAttestationAtSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAttestationAtSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAttestationAtSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationAtSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttestationAtSubscriptionResultOutput{})
}