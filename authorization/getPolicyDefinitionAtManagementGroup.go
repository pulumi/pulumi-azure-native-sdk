// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation retrieves the policy definition in the given management group with the given name.
//
// Uses Azure REST API version 2025-01-01.
//
// Other available API versions: 2020-09-01, 2021-06-01, 2023-04-01, 2024-05-01, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPolicyDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicyDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization:getPolicyDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyDefinitionAtManagementGroupArgs struct {
	// The ID of the management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The name of the policy definition to get.
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}

// The policy definition.
type LookupPolicyDefinitionAtManagementGroupResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The policy definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the policy definition.
	Id string `pulumi:"id"`
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode *string `pulumi:"mode"`
	// The name of the policy definition.
	Name string `pulumi:"name"`
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	// The policy rule.
	PolicyRule interface{} `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type string `pulumi:"type"`
	// The policy definition version in #.#.# format.
	Version *string `pulumi:"version"`
	// A list of available versions for this policy definition.
	Versions []string `pulumi:"versions"`
}

// Defaults sets the appropriate defaults for LookupPolicyDefinitionAtManagementGroupResult
func (val *LookupPolicyDefinitionAtManagementGroupResult) Defaults() *LookupPolicyDefinitionAtManagementGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Mode == nil {
		mode_ := "Indexed"
		tmp.Mode = &mode_
	}
	return &tmp
}
func LookupPolicyDefinitionAtManagementGroupOutput(ctx *pulumi.Context, args LookupPolicyDefinitionAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyDefinitionAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPolicyDefinitionAtManagementGroupResultOutput, error) {
			args := v.(LookupPolicyDefinitionAtManagementGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:authorization:getPolicyDefinitionAtManagementGroup", args, LookupPolicyDefinitionAtManagementGroupResultOutput{}, options).(LookupPolicyDefinitionAtManagementGroupResultOutput), nil
		}).(LookupPolicyDefinitionAtManagementGroupResultOutput)
}

type LookupPolicyDefinitionAtManagementGroupOutputArgs struct {
	// The ID of the management group.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// The name of the policy definition to get.
	PolicyDefinitionName pulumi.StringInput `pulumi:"policyDefinitionName"`
}

func (LookupPolicyDefinitionAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionAtManagementGroupArgs)(nil)).Elem()
}

// The policy definition.
type LookupPolicyDefinitionAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyDefinitionAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionAtManagementGroupResult)(nil)).Elem()
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) ToLookupPolicyDefinitionAtManagementGroupResultOutput() LookupPolicyDefinitionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) ToLookupPolicyDefinitionAtManagementGroupResultOutputWithContext(ctx context.Context) LookupPolicyDefinitionAtManagementGroupResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The policy definition description.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the policy definition.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

// The policy rule.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) interface{} { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions).
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The policy definition version in #.#.# format.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// A list of available versions for this policy definition.
func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyDefinitionAtManagementGroupResultOutput{})
}
