// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ID of the policy definition version.
//
// Uses Azure REST API version 2025-01-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
//
// Other available API versions: 2023-04-01, 2024-05-01, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type PolicyDefinitionVersion struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The policy definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the policy definition version.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters ParameterDefinitionsValueResponseMapOutput `pulumi:"parameters"`
	// The policy rule.
	PolicyRule pulumi.AnyOutput `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/policyDefinitions/versions).
	Type pulumi.StringOutput `pulumi:"type"`
	// The policy definition version in #.#.# format.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewPolicyDefinitionVersion registers a new resource with the given unique name, arguments, and options.
func NewPolicyDefinitionVersion(ctx *pulumi.Context,
	name string, args *PolicyDefinitionVersionArgs, opts ...pulumi.ResourceOption) (*PolicyDefinitionVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDefinitionName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitionName'")
	}
	if args.Mode == nil {
		args.Mode = pulumi.StringPtr("Indexed")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20230401:PolicyDefinitionVersion"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20240501:PolicyDefinitionVersion"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20250101:PolicyDefinitionVersion"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20250301:PolicyDefinitionVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PolicyDefinitionVersion
	err := ctx.RegisterResource("azure-native:authorization:PolicyDefinitionVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyDefinitionVersion gets an existing PolicyDefinitionVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyDefinitionVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyDefinitionVersionState, opts ...pulumi.ResourceOption) (*PolicyDefinitionVersion, error) {
	var resource PolicyDefinitionVersion
	err := ctx.ReadResource("azure-native:authorization:PolicyDefinitionVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyDefinitionVersion resources.
type policyDefinitionVersionState struct {
}

type PolicyDefinitionVersionState struct {
}

func (PolicyDefinitionVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionVersionState)(nil)).Elem()
}

type policyDefinitionVersionArgs struct {
	// The policy definition description.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode *string `pulumi:"mode"`
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters map[string]ParameterDefinitionsValue `pulumi:"parameters"`
	// The name of the policy definition.
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
	// The policy definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion *string `pulumi:"policyDefinitionVersion"`
	// The policy rule.
	PolicyRule interface{} `pulumi:"policyRule"`
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *string `pulumi:"policyType"`
	// The policy definition version in #.#.# format.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a PolicyDefinitionVersion resource.
type PolicyDefinitionVersionArgs struct {
	// The policy definition description.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrInput
	// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata pulumi.Input
	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode pulumi.StringPtrInput
	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters ParameterDefinitionsValueMapInput
	// The name of the policy definition.
	PolicyDefinitionName pulumi.StringInput
	// The policy definition version.  The format is x.y.z where x is the major version number, y is the minor version number, and z is the patch number
	PolicyDefinitionVersion pulumi.StringPtrInput
	// The policy rule.
	PolicyRule pulumi.Input
	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType pulumi.StringPtrInput
	// The policy definition version in #.#.# format.
	Version pulumi.StringPtrInput
}

func (PolicyDefinitionVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyDefinitionVersionArgs)(nil)).Elem()
}

type PolicyDefinitionVersionInput interface {
	pulumi.Input

	ToPolicyDefinitionVersionOutput() PolicyDefinitionVersionOutput
	ToPolicyDefinitionVersionOutputWithContext(ctx context.Context) PolicyDefinitionVersionOutput
}

func (*PolicyDefinitionVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinitionVersion)(nil)).Elem()
}

func (i *PolicyDefinitionVersion) ToPolicyDefinitionVersionOutput() PolicyDefinitionVersionOutput {
	return i.ToPolicyDefinitionVersionOutputWithContext(context.Background())
}

func (i *PolicyDefinitionVersion) ToPolicyDefinitionVersionOutputWithContext(ctx context.Context) PolicyDefinitionVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionVersionOutput)
}

type PolicyDefinitionVersionOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinitionVersion)(nil)).Elem()
}

func (o PolicyDefinitionVersionOutput) ToPolicyDefinitionVersionOutput() PolicyDefinitionVersionOutput {
	return o
}

func (o PolicyDefinitionVersionOutput) ToPolicyDefinitionVersionOutputWithContext(ctx context.Context) PolicyDefinitionVersionOutput {
	return o
}

// The Azure API version of the resource.
func (o PolicyDefinitionVersionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The policy definition description.
func (o PolicyDefinitionVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the policy definition.
func (o PolicyDefinitionVersionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
func (o PolicyDefinitionVersionOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
func (o PolicyDefinitionVersionOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the policy definition version.
func (o PolicyDefinitionVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
func (o PolicyDefinitionVersionOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) ParameterDefinitionsValueResponseMapOutput { return v.Parameters }).(ParameterDefinitionsValueResponseMapOutput)
}

// The policy rule.
func (o PolicyDefinitionVersionOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.AnyOutput { return v.PolicyRule }).(pulumi.AnyOutput)
}

// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
func (o PolicyDefinitionVersionOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o PolicyDefinitionVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/policyDefinitions/versions).
func (o PolicyDefinitionVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The policy definition version in #.#.# format.
func (o PolicyDefinitionVersionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyDefinitionVersion) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyDefinitionVersionOutput{})
}
