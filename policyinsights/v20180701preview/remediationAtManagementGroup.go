// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The remediation definition.
//
// Deprecated: Version 2018-07-01-preview will be removed in v2 of the provider.
type RemediationAtManagementGroup struct {
	pulumi.CustomResourceState

	// The time at which the remediation was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryResponsePtrOutput `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersResponsePtrOutput `pulumi:"filters"`
	// The time at which the remediation was last updated.
	LastUpdatedOn pulumi.StringOutput `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrOutput `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrOutput `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the remediation.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRemediationAtManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewRemediationAtManagementGroup(ctx *pulumi.Context,
	name string, args *RemediationAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*RemediationAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.ManagementGroupsNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupsNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights:RemediationAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20211001:RemediationAtManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtManagementGroup
	err := ctx.RegisterResource("azure-native:policyinsights/v20180701preview:RemediationAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationAtManagementGroup gets an existing RemediationAtManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtManagementGroupState, opts ...pulumi.ResourceOption) (*RemediationAtManagementGroup, error) {
	var resource RemediationAtManagementGroup
	err := ctx.ReadResource("azure-native:policyinsights/v20180701preview:RemediationAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationAtManagementGroup resources.
type remediationAtManagementGroupState struct {
}

type RemediationAtManagementGroupState struct {
}

func (RemediationAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtManagementGroupState)(nil)).Elem()
}

type remediationAtManagementGroupArgs struct {
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus *RemediationDeploymentSummary `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFilters `pulumi:"filters"`
	// Management group ID.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The name of the remediation.
	RemediationName *string `pulumi:"remediationName"`
}

// The set of arguments for constructing a RemediationAtManagementGroup resource.
type RemediationAtManagementGroupArgs struct {
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryPtrInput
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersPtrInput
	// Management group ID.
	ManagementGroupId pulumi.StringInput
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace pulumi.StringInput
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrInput
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The name of the remediation.
	RemediationName pulumi.StringPtrInput
}

func (RemediationAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtManagementGroupArgs)(nil)).Elem()
}

type RemediationAtManagementGroupInput interface {
	pulumi.Input

	ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput
	ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput
}

func (*RemediationAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtManagementGroup)(nil)).Elem()
}

func (i *RemediationAtManagementGroup) ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput {
	return i.ToRemediationAtManagementGroupOutputWithContext(context.Background())
}

func (i *RemediationAtManagementGroup) ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtManagementGroupOutput)
}

type RemediationAtManagementGroupOutput struct{ *pulumi.OutputState }

func (RemediationAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtManagementGroup)(nil)).Elem()
}

func (o RemediationAtManagementGroupOutput) ToRemediationAtManagementGroupOutput() RemediationAtManagementGroupOutput {
	return o
}

func (o RemediationAtManagementGroupOutput) ToRemediationAtManagementGroupOutputWithContext(ctx context.Context) RemediationAtManagementGroupOutput {
	return o
}

// The time at which the remediation was created.
func (o RemediationAtManagementGroupOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// The deployment status summary for all deployments created by the remediation.
func (o RemediationAtManagementGroupOutput) DeploymentStatus() RemediationDeploymentSummaryResponsePtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) RemediationDeploymentSummaryResponsePtrOutput {
		return v.DeploymentStatus
	}).(RemediationDeploymentSummaryResponsePtrOutput)
}

// The filters that will be applied to determine which resources to remediate.
func (o RemediationAtManagementGroupOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) RemediationFiltersResponsePtrOutput { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

// The time at which the remediation was last updated.
func (o RemediationAtManagementGroupOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

// The name of the remediation.
func (o RemediationAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the policy assignment that should be remediated.
func (o RemediationAtManagementGroupOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
func (o RemediationAtManagementGroupOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the remediation.
func (o RemediationAtManagementGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the remediation.
func (o RemediationAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemediationAtManagementGroupOutput{})
}