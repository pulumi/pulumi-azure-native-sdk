// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Blueprint artifact applies Policy assignments.
type PolicyAssignmentArtifact struct {
	pulumi.CustomResourceState

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayOutput `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'policyAssignment'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameter values for the policy definition.
	Parameters ParameterValueBaseResponseMapOutput `pulumi:"parameters"`
	// Azure resource ID of the policy definition.
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// Name of the resource group placeholder to which the policy will be assigned.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicyAssignmentArtifact registers a new resource with the given unique name, arguments, and options.
func NewPolicyAssignmentArtifact(ctx *pulumi.Context,
	name string, args *PolicyAssignmentArtifactArgs, opts ...pulumi.ResourceOption) (*PolicyAssignmentArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ManagementGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupName'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.PolicyDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitionId'")
	}
	args.Kind = pulumi.String("policyAssignment")
	var resource PolicyAssignmentArtifact
	err := ctx.RegisterResource("azure-native:blueprint/v20171111preview:PolicyAssignmentArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAssignmentArtifact gets an existing PolicyAssignmentArtifact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAssignmentArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentArtifactState, opts ...pulumi.ResourceOption) (*PolicyAssignmentArtifact, error) {
	var resource PolicyAssignmentArtifact
	err := ctx.ReadResource("azure-native:blueprint/v20171111preview:PolicyAssignmentArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAssignmentArtifact resources.
type policyAssignmentArtifactState struct {
}

type PolicyAssignmentArtifactState struct {
}

func (PolicyAssignmentArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArtifactState)(nil)).Elem()
}

type policyAssignmentArtifactArgs struct {
	// name of the artifact.
	ArtifactName *string `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []string `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'policyAssignment'.
	Kind string `pulumi:"kind"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
	// Parameter values for the policy definition.
	Parameters map[string]ParameterValueBase `pulumi:"parameters"`
	// Azure resource ID of the policy definition.
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// Name of the resource group placeholder to which the policy will be assigned.
	ResourceGroup *string `pulumi:"resourceGroup"`
}

// The set of arguments for constructing a PolicyAssignmentArtifact resource.
type PolicyAssignmentArtifactArgs struct {
	// name of the artifact.
	ArtifactName pulumi.StringPtrInput
	// name of the blueprint.
	BlueprintName pulumi.StringInput
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Specifies the kind of Blueprint artifact.
	// Expected value is 'policyAssignment'.
	Kind pulumi.StringInput
	// ManagementGroup where blueprint stores.
	ManagementGroupName pulumi.StringInput
	// Parameter values for the policy definition.
	Parameters ParameterValueBaseMapInput
	// Azure resource ID of the policy definition.
	PolicyDefinitionId pulumi.StringInput
	// Name of the resource group placeholder to which the policy will be assigned.
	ResourceGroup pulumi.StringPtrInput
}

func (PolicyAssignmentArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArtifactArgs)(nil)).Elem()
}

type PolicyAssignmentArtifactInput interface {
	pulumi.Input

	ToPolicyAssignmentArtifactOutput() PolicyAssignmentArtifactOutput
	ToPolicyAssignmentArtifactOutputWithContext(ctx context.Context) PolicyAssignmentArtifactOutput
}

func (*PolicyAssignmentArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentArtifact)(nil)).Elem()
}

func (i *PolicyAssignmentArtifact) ToPolicyAssignmentArtifactOutput() PolicyAssignmentArtifactOutput {
	return i.ToPolicyAssignmentArtifactOutputWithContext(context.Background())
}

func (i *PolicyAssignmentArtifact) ToPolicyAssignmentArtifactOutputWithContext(ctx context.Context) PolicyAssignmentArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentArtifactOutput)
}

type PolicyAssignmentArtifactOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentArtifact)(nil)).Elem()
}

func (o PolicyAssignmentArtifactOutput) ToPolicyAssignmentArtifactOutput() PolicyAssignmentArtifactOutput {
	return o
}

func (o PolicyAssignmentArtifactOutput) ToPolicyAssignmentArtifactOutputWithContext(ctx context.Context) PolicyAssignmentArtifactOutput {
	return o
}

// Artifacts which need to be deployed before the specified artifact.
func (o PolicyAssignmentArtifactOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringArrayOutput { return v.DependsOn }).(pulumi.StringArrayOutput)
}

// Multi-line explain this resource.
func (o PolicyAssignmentArtifactOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o PolicyAssignmentArtifactOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Specifies the kind of Blueprint artifact.
// Expected value is 'policyAssignment'.
func (o PolicyAssignmentArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of this resource.
func (o PolicyAssignmentArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parameter values for the policy definition.
func (o PolicyAssignmentArtifactOutput) Parameters() ParameterValueBaseResponseMapOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) ParameterValueBaseResponseMapOutput { return v.Parameters }).(ParameterValueBaseResponseMapOutput)
}

// Azure resource ID of the policy definition.
func (o PolicyAssignmentArtifactOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

// Name of the resource group placeholder to which the policy will be assigned.
func (o PolicyAssignmentArtifactOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Type of this resource.
func (o PolicyAssignmentArtifactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAssignmentArtifactOutput{})
}