// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation retrieves the policy set definition in the given management group with the given name.
//
// Uses Azure REST API version 2025-01-01.
//
// Other available API versions: 2020-09-01, 2021-06-01, 2023-04-01, 2024-05-01, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicySetDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicySetDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization:getPolicySetDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionAtManagementGroupArgs struct {
	// Comma-separated list of additional properties to be included in the response. Supported values are 'LatestDefinitionVersion, EffectiveDefinitionVersion'.
	Expand *string `pulumi:"expand"`
	// The ID of the management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The name of the policy set definition to get.
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}

// The policy set definition.
type LookupPolicySetDefinitionAtManagementGroupResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The policy set definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the policy set definition.
	Id string `pulumi:"id"`
	// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The name of the policy set definition.
	Name string `pulumi:"name"`
	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups []PolicyDefinitionGroupResponse `pulumi:"policyDefinitionGroups"`
	// An array of policy definition references.
	PolicyDefinitions []PolicyDefinitionReferenceResponse `pulumi:"policyDefinitions"`
	// The type of policy set definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type string `pulumi:"type"`
	// The policy set definition version in #.#.# format.
	Version *string `pulumi:"version"`
	// A list of available versions for this policy set definition.
	Versions []string `pulumi:"versions"`
}

func LookupPolicySetDefinitionAtManagementGroupOutput(ctx *pulumi.Context, args LookupPolicySetDefinitionAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPolicySetDefinitionAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPolicySetDefinitionAtManagementGroupResultOutput, error) {
			args := v.(LookupPolicySetDefinitionAtManagementGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:authorization:getPolicySetDefinitionAtManagementGroup", args, LookupPolicySetDefinitionAtManagementGroupResultOutput{}, options).(LookupPolicySetDefinitionAtManagementGroupResultOutput), nil
		}).(LookupPolicySetDefinitionAtManagementGroupResultOutput)
}

type LookupPolicySetDefinitionAtManagementGroupOutputArgs struct {
	// Comma-separated list of additional properties to be included in the response. Supported values are 'LatestDefinitionVersion, EffectiveDefinitionVersion'.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The ID of the management group.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// The name of the policy set definition to get.
	PolicySetDefinitionName pulumi.StringInput `pulumi:"policySetDefinitionName"`
}

func (LookupPolicySetDefinitionAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionAtManagementGroupArgs)(nil)).Elem()
}

// The policy set definition.
type LookupPolicySetDefinitionAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPolicySetDefinitionAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionAtManagementGroupResult)(nil)).Elem()
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) ToLookupPolicySetDefinitionAtManagementGroupResultOutput() LookupPolicySetDefinitionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) ToLookupPolicySetDefinitionAtManagementGroupResultOutputWithContext(ctx context.Context) LookupPolicySetDefinitionAtManagementGroupResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The policy set definition description.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy set definition.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy set definition.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The name of the policy set definition.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The policy set definition parameters that can be used in policy definition references.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The metadata describing groups of policy definition references within the policy set definition.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) PolicyDefinitionGroups() PolicyDefinitionGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) []PolicyDefinitionGroupResponse {
		return v.PolicyDefinitionGroups
	}).(PolicyDefinitionGroupResponseArrayOutput)
}

// An array of policy definition references.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) []PolicyDefinitionReferenceResponse {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

// The type of policy set definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policySetDefinitions).
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy set definition version in #.#.# format.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// A list of available versions for this policy set definition.
func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicySetDefinitionAtManagementGroupResultOutput{})
}
