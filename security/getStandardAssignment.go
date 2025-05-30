// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation retrieves a single standard assignment, given its name and the scope it was created at.
//
// Uses Azure REST API version 2024-08-01.
func LookupStandardAssignment(ctx *pulumi.Context, args *LookupStandardAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupStandardAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStandardAssignmentResult
	err := ctx.Invoke("azure-native:security:getStandardAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStandardAssignmentArgs struct {
	// The identifier of the resource.
	ResourceId string `pulumi:"resourceId"`
	// The standard assignments assignment key - unique key for the standard assignment
	StandardAssignmentName string `pulumi:"standardAssignmentName"`
}

// Security Assignment on a resource group over a given scope
type LookupStandardAssignmentResult struct {
	// Standard item with key as applied to this standard assignment over the given scope
	AssignedStandard *AssignedStandardItemResponse `pulumi:"assignedStandard"`
	// Additional data about assignment that has Attest effect
	AttestationData *StandardAssignmentPropertiesResponseAttestationData `pulumi:"attestationData"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Description of the standardAssignment
	Description *string `pulumi:"description"`
	// Display name of the standardAssignment
	DisplayName *string `pulumi:"displayName"`
	// Expected effect of this assignment (Audit/Exempt/Attest)
	Effect *string `pulumi:"effect"`
	// Excluded scopes, filter out the descendants of the scope (on management scopes)
	ExcludedScopes []string `pulumi:"excludedScopes"`
	// Additional data about assignment that has Exempt effect
	ExemptionData *StandardAssignmentPropertiesResponseExemptionData `pulumi:"exemptionData"`
	// Expiration date of this assignment as a full ISO date
	ExpiresOn *string `pulumi:"expiresOn"`
	// Resource Id
	Id string `pulumi:"id"`
	// The standard assignment metadata.
	Metadata *StandardAssignmentMetadataResponse `pulumi:"metadata"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupStandardAssignmentOutput(ctx *pulumi.Context, args LookupStandardAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupStandardAssignmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStandardAssignmentResultOutput, error) {
			args := v.(LookupStandardAssignmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:security:getStandardAssignment", args, LookupStandardAssignmentResultOutput{}, options).(LookupStandardAssignmentResultOutput), nil
		}).(LookupStandardAssignmentResultOutput)
}

type LookupStandardAssignmentOutputArgs struct {
	// The identifier of the resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The standard assignments assignment key - unique key for the standard assignment
	StandardAssignmentName pulumi.StringInput `pulumi:"standardAssignmentName"`
}

func (LookupStandardAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardAssignmentArgs)(nil)).Elem()
}

// Security Assignment on a resource group over a given scope
type LookupStandardAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupStandardAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardAssignmentResult)(nil)).Elem()
}

func (o LookupStandardAssignmentResultOutput) ToLookupStandardAssignmentResultOutput() LookupStandardAssignmentResultOutput {
	return o
}

func (o LookupStandardAssignmentResultOutput) ToLookupStandardAssignmentResultOutputWithContext(ctx context.Context) LookupStandardAssignmentResultOutput {
	return o
}

// Standard item with key as applied to this standard assignment over the given scope
func (o LookupStandardAssignmentResultOutput) AssignedStandard() AssignedStandardItemResponsePtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *AssignedStandardItemResponse { return v.AssignedStandard }).(AssignedStandardItemResponsePtrOutput)
}

// Additional data about assignment that has Attest effect
func (o LookupStandardAssignmentResultOutput) AttestationData() StandardAssignmentPropertiesResponseAttestationDataPtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *StandardAssignmentPropertiesResponseAttestationData {
		return v.AttestationData
	}).(StandardAssignmentPropertiesResponseAttestationDataPtrOutput)
}

// The Azure API version of the resource.
func (o LookupStandardAssignmentResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Description of the standardAssignment
func (o LookupStandardAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of the standardAssignment
func (o LookupStandardAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Expected effect of this assignment (Audit/Exempt/Attest)
func (o LookupStandardAssignmentResultOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *string { return v.Effect }).(pulumi.StringPtrOutput)
}

// Excluded scopes, filter out the descendants of the scope (on management scopes)
func (o LookupStandardAssignmentResultOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) []string { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

// Additional data about assignment that has Exempt effect
func (o LookupStandardAssignmentResultOutput) ExemptionData() StandardAssignmentPropertiesResponseExemptionDataPtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *StandardAssignmentPropertiesResponseExemptionData {
		return v.ExemptionData
	}).(StandardAssignmentPropertiesResponseExemptionDataPtrOutput)
}

// Expiration date of this assignment as a full ISO date
func (o LookupStandardAssignmentResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupStandardAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The standard assignment metadata.
func (o LookupStandardAssignmentResultOutput) Metadata() StandardAssignmentMetadataResponsePtrOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) *StandardAssignmentMetadataResponse { return v.Metadata }).(StandardAssignmentMetadataResponsePtrOutput)
}

// Resource name
func (o LookupStandardAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o LookupStandardAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStandardAssignmentResultOutput{})
}
