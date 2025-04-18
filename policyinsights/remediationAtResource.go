// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package policyinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The remediation definition.
//
// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2021-10-01.
//
// Other available API versions: 2021-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native policyinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type RemediationAtResource struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The remediation correlation Id. Can be used to find events related to the remediation in the activity log.
	CorrelationId pulumi.StringOutput `pulumi:"correlationId"`
	// The time at which the remediation was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryResponseOutput `pulumi:"deploymentStatus"`
	// The remediation failure threshold settings
	FailureThreshold RemediationPropertiesResponseFailureThresholdPtrOutput `pulumi:"failureThreshold"`
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersResponsePtrOutput `pulumi:"filters"`
	// The time at which the remediation was last updated.
	LastUpdatedOn pulumi.StringOutput `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name pulumi.StringOutput `pulumi:"name"`
	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments pulumi.IntPtrOutput `pulumi:"parallelDeployments"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrOutput `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrOutput `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation. This refers to the entire remediation task, not individual deployments. Allowed values are Evaluating, Canceled, Cancelling, Failed, Complete, or Succeeded.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount pulumi.IntPtrOutput `pulumi:"resourceCount"`
	// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
	ResourceDiscoveryMode pulumi.StringPtrOutput `pulumi:"resourceDiscoveryMode"`
	// The remediation status message. Provides additional details regarding the state of the remediation.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the remediation.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRemediationAtResource registers a new resource with the given unique name, arguments, and options.
func NewRemediationAtResource(ctx *pulumi.Context,
	name string, args *RemediationAtResourceArgs, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:policyinsights/v20180701preview:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20190701:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20211001:RemediationAtResource"),
		},
		{
			Type: pulumi.String("azure-native:policyinsights/v20241001:RemediationAtResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RemediationAtResource
	err := ctx.RegisterResource("azure-native:policyinsights:RemediationAtResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationAtResource gets an existing RemediationAtResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationAtResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtResourceState, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	var resource RemediationAtResource
	err := ctx.ReadResource("azure-native:policyinsights:RemediationAtResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationAtResource resources.
type remediationAtResourceState struct {
}

type RemediationAtResourceState struct {
}

func (RemediationAtResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceState)(nil)).Elem()
}

type remediationAtResourceArgs struct {
	// The remediation failure threshold settings
	FailureThreshold *RemediationPropertiesFailureThreshold `pulumi:"failureThreshold"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFilters `pulumi:"filters"`
	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments *int `pulumi:"parallelDeployments"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The name of the remediation.
	RemediationName *string `pulumi:"remediationName"`
	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount *int `pulumi:"resourceCount"`
	// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
	ResourceDiscoveryMode *string `pulumi:"resourceDiscoveryMode"`
	// Resource ID.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a RemediationAtResource resource.
type RemediationAtResourceArgs struct {
	// The remediation failure threshold settings
	FailureThreshold RemediationPropertiesFailureThresholdPtrInput
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersPtrInput
	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
	ParallelDeployments pulumi.IntPtrInput
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrInput
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The name of the remediation.
	RemediationName pulumi.StringPtrInput
	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
	ResourceCount pulumi.IntPtrInput
	// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
	ResourceDiscoveryMode pulumi.StringPtrInput
	// Resource ID.
	ResourceId pulumi.StringInput
}

func (RemediationAtResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceArgs)(nil)).Elem()
}

type RemediationAtResourceInput interface {
	pulumi.Input

	ToRemediationAtResourceOutput() RemediationAtResourceOutput
	ToRemediationAtResourceOutputWithContext(ctx context.Context) RemediationAtResourceOutput
}

func (*RemediationAtResource) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtResource)(nil)).Elem()
}

func (i *RemediationAtResource) ToRemediationAtResourceOutput() RemediationAtResourceOutput {
	return i.ToRemediationAtResourceOutputWithContext(context.Background())
}

func (i *RemediationAtResource) ToRemediationAtResourceOutputWithContext(ctx context.Context) RemediationAtResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationAtResourceOutput)
}

type RemediationAtResourceOutput struct{ *pulumi.OutputState }

func (RemediationAtResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationAtResource)(nil)).Elem()
}

func (o RemediationAtResourceOutput) ToRemediationAtResourceOutput() RemediationAtResourceOutput {
	return o
}

func (o RemediationAtResourceOutput) ToRemediationAtResourceOutputWithContext(ctx context.Context) RemediationAtResourceOutput {
	return o
}

// The Azure API version of the resource.
func (o RemediationAtResourceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The remediation correlation Id. Can be used to find events related to the remediation in the activity log.
func (o RemediationAtResourceOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.CorrelationId }).(pulumi.StringOutput)
}

// The time at which the remediation was created.
func (o RemediationAtResourceOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// The deployment status summary for all deployments created by the remediation.
func (o RemediationAtResourceOutput) DeploymentStatus() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v *RemediationAtResource) RemediationDeploymentSummaryResponseOutput { return v.DeploymentStatus }).(RemediationDeploymentSummaryResponseOutput)
}

// The remediation failure threshold settings
func (o RemediationAtResourceOutput) FailureThreshold() RemediationPropertiesResponseFailureThresholdPtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) RemediationPropertiesResponseFailureThresholdPtrOutput {
		return v.FailureThreshold
	}).(RemediationPropertiesResponseFailureThresholdPtrOutput)
}

// The filters that will be applied to determine which resources to remediate.
func (o RemediationAtResourceOutput) Filters() RemediationFiltersResponsePtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) RemediationFiltersResponsePtrOutput { return v.Filters }).(RemediationFiltersResponsePtrOutput)
}

// The time at which the remediation was last updated.
func (o RemediationAtResourceOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

// The name of the remediation.
func (o RemediationAtResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation. If not provided, the default parallel deployments value is used.
func (o RemediationAtResourceOutput) ParallelDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.IntPtrOutput { return v.ParallelDeployments }).(pulumi.IntPtrOutput)
}

// The resource ID of the policy assignment that should be remediated.
func (o RemediationAtResourceOutput) PolicyAssignmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringPtrOutput { return v.PolicyAssignmentId }).(pulumi.StringPtrOutput)
}

// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
func (o RemediationAtResourceOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringPtrOutput { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

// The status of the remediation. This refers to the entire remediation task, not individual deployments. Allowed values are Evaluating, Canceled, Cancelling, Failed, Complete, or Succeeded.
func (o RemediationAtResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource count is used.
func (o RemediationAtResourceOutput) ResourceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.IntPtrOutput { return v.ResourceCount }).(pulumi.IntPtrOutput)
}

// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
func (o RemediationAtResourceOutput) ResourceDiscoveryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringPtrOutput { return v.ResourceDiscoveryMode }).(pulumi.StringPtrOutput)
}

// The remediation status message. Provides additional details regarding the state of the remediation.
func (o RemediationAtResourceOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RemediationAtResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RemediationAtResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the remediation.
func (o RemediationAtResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationAtResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemediationAtResourceOutput{})
}
