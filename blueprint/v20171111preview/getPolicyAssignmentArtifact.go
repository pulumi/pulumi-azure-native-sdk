// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Blueprint artifact applies Policy assignments.
func LookupPolicyAssignmentArtifact(ctx *pulumi.Context, args *LookupPolicyAssignmentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentArtifactResult, error) {
	var rv LookupPolicyAssignmentArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getPolicyAssignmentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArtifactArgs struct {
	// name of the artifact.
	ArtifactName string `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// Blueprint artifact applies Policy assignments.
type LookupPolicyAssignmentArtifactResult struct {
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []string `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'policyAssignment'.
	Kind string `pulumi:"kind"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Parameter values for the policy definition.
	Parameters map[string]ParameterValueBaseResponse `pulumi:"parameters"`
	// Azure resource ID of the policy definition.
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// Name of the resource group placeholder to which the policy will be assigned.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

func LookupPolicyAssignmentArtifactOutput(ctx *pulumi.Context, args LookupPolicyAssignmentArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyAssignmentArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyAssignmentArtifactResult, error) {
			args := v.(LookupPolicyAssignmentArtifactArgs)
			r, err := LookupPolicyAssignmentArtifact(ctx, &args, opts...)
			var s LookupPolicyAssignmentArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyAssignmentArtifactResultOutput)
}

type LookupPolicyAssignmentArtifactOutputArgs struct {
	// name of the artifact.
	ArtifactName pulumi.StringInput `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
}

func (LookupPolicyAssignmentArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArtifactArgs)(nil)).Elem()
}

// Blueprint artifact applies Policy assignments.
type LookupPolicyAssignmentArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyAssignmentArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArtifactResult)(nil)).Elem()
}

func (o LookupPolicyAssignmentArtifactResultOutput) ToLookupPolicyAssignmentArtifactResultOutput() LookupPolicyAssignmentArtifactResultOutput {
	return o
}

func (o LookupPolicyAssignmentArtifactResultOutput) ToLookupPolicyAssignmentArtifactResultOutputWithContext(ctx context.Context) LookupPolicyAssignmentArtifactResultOutput {
	return o
}

// Artifacts which need to be deployed before the specified artifact.
func (o LookupPolicyAssignmentArtifactResultOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

// Multi-line explain this resource.
func (o LookupPolicyAssignmentArtifactResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o LookupPolicyAssignmentArtifactResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupPolicyAssignmentArtifactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the kind of Blueprint artifact.
// Expected value is 'policyAssignment'.
func (o LookupPolicyAssignmentArtifactResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of this resource.
func (o LookupPolicyAssignmentArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

// Parameter values for the policy definition.
func (o LookupPolicyAssignmentArtifactResultOutput) Parameters() ParameterValueBaseResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) map[string]ParameterValueBaseResponse {
		return v.Parameters
	}).(ParameterValueBaseResponseMapOutput)
}

// Azure resource ID of the policy definition.
func (o LookupPolicyAssignmentArtifactResultOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

// Name of the resource group placeholder to which the policy will be assigned.
func (o LookupPolicyAssignmentArtifactResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Type of this resource.
func (o LookupPolicyAssignmentArtifactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentArtifactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyAssignmentArtifactResultOutput{})
}
