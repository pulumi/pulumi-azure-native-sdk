// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security assessment metadata
type AssessmentsMetadataSubscription struct {
	pulumi.CustomResourceState

	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringOutput      `pulumi:"assessmentType"`
	Categories     pulumi.StringArrayOutput `pulumi:"categories"`
	// Human readable description of the assessment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrOutput `pulumi:"implementationEffort"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure resource ID of the policy definition that turns this assessment calculation on
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// True if this assessment is in preview release status
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrOutput `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity pulumi.StringOutput      `pulumi:"severity"`
	Threats  pulumi.StringArrayOutput `pulumi:"threats"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The user impact of the assessment
	UserImpact pulumi.StringPtrOutput `pulumi:"userImpact"`
}

// NewAssessmentsMetadataSubscription registers a new resource with the given unique name, arguments, and options.
func NewAssessmentsMetadataSubscription(ctx *pulumi.Context,
	name string, args *AssessmentsMetadataSubscriptionArgs, opts ...pulumi.ResourceOption) (*AssessmentsMetadataSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentType == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentType'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:AssessmentsMetadataSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210601:AssessmentsMetadataSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource AssessmentsMetadataSubscription
	err := ctx.RegisterResource("azure-native:security/v20190101preview:AssessmentsMetadataSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentsMetadataSubscription gets an existing AssessmentsMetadataSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentsMetadataSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentsMetadataSubscriptionState, opts ...pulumi.ResourceOption) (*AssessmentsMetadataSubscription, error) {
	var resource AssessmentsMetadataSubscription
	err := ctx.ReadResource("azure-native:security/v20190101preview:AssessmentsMetadataSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentsMetadataSubscription resources.
type assessmentsMetadataSubscriptionState struct {
}

type AssessmentsMetadataSubscriptionState struct {
}

func (AssessmentsMetadataSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentsMetadataSubscriptionState)(nil)).Elem()
}

type assessmentsMetadataSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName *string `pulumi:"assessmentMetadataName"`
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType string   `pulumi:"assessmentType"`
	Categories     []string `pulumi:"categories"`
	// Human readable description of the assessment
	Description *string `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName string `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort *string `pulumi:"implementationEffort"`
	// True if this assessment is in preview release status
	Preview *bool `pulumi:"preview"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription *string `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity string   `pulumi:"severity"`
	Threats  []string `pulumi:"threats"`
	// The user impact of the assessment
	UserImpact *string `pulumi:"userImpact"`
}

// The set of arguments for constructing a AssessmentsMetadataSubscription resource.
type AssessmentsMetadataSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName pulumi.StringPtrInput
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringInput
	Categories     pulumi.StringArrayInput
	// Human readable description of the assessment
	Description pulumi.StringPtrInput
	// User friendly display name of the assessment
	DisplayName pulumi.StringInput
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrInput
	// True if this assessment is in preview release status
	Preview pulumi.BoolPtrInput
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrInput
	// The severity level of the assessment
	Severity pulumi.StringInput
	Threats  pulumi.StringArrayInput
	// The user impact of the assessment
	UserImpact pulumi.StringPtrInput
}

func (AssessmentsMetadataSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentsMetadataSubscriptionArgs)(nil)).Elem()
}

type AssessmentsMetadataSubscriptionInput interface {
	pulumi.Input

	ToAssessmentsMetadataSubscriptionOutput() AssessmentsMetadataSubscriptionOutput
	ToAssessmentsMetadataSubscriptionOutputWithContext(ctx context.Context) AssessmentsMetadataSubscriptionOutput
}

func (*AssessmentsMetadataSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentsMetadataSubscription)(nil)).Elem()
}

func (i *AssessmentsMetadataSubscription) ToAssessmentsMetadataSubscriptionOutput() AssessmentsMetadataSubscriptionOutput {
	return i.ToAssessmentsMetadataSubscriptionOutputWithContext(context.Background())
}

func (i *AssessmentsMetadataSubscription) ToAssessmentsMetadataSubscriptionOutputWithContext(ctx context.Context) AssessmentsMetadataSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentsMetadataSubscriptionOutput)
}

type AssessmentsMetadataSubscriptionOutput struct{ *pulumi.OutputState }

func (AssessmentsMetadataSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentsMetadataSubscription)(nil)).Elem()
}

func (o AssessmentsMetadataSubscriptionOutput) ToAssessmentsMetadataSubscriptionOutput() AssessmentsMetadataSubscriptionOutput {
	return o
}

func (o AssessmentsMetadataSubscriptionOutput) ToAssessmentsMetadataSubscriptionOutputWithContext(ctx context.Context) AssessmentsMetadataSubscriptionOutput {
	return o
}

// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
func (o AssessmentsMetadataSubscriptionOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringOutput { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o AssessmentsMetadataSubscriptionOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringArrayOutput { return v.Categories }).(pulumi.StringArrayOutput)
}

// Human readable description of the assessment
func (o AssessmentsMetadataSubscriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User friendly display name of the assessment
func (o AssessmentsMetadataSubscriptionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The implementation effort required to remediate this assessment
func (o AssessmentsMetadataSubscriptionOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringPtrOutput { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

// Resource name
func (o AssessmentsMetadataSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure resource ID of the policy definition that turns this assessment calculation on
func (o AssessmentsMetadataSubscriptionOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringOutput { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

// True if this assessment is in preview release status
func (o AssessmentsMetadataSubscriptionOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.BoolPtrOutput { return v.Preview }).(pulumi.BoolPtrOutput)
}

// Human readable description of what you should do to mitigate this security issue
func (o AssessmentsMetadataSubscriptionOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringPtrOutput { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

// The severity level of the assessment
func (o AssessmentsMetadataSubscriptionOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o AssessmentsMetadataSubscriptionOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringArrayOutput { return v.Threats }).(pulumi.StringArrayOutput)
}

// Resource type
func (o AssessmentsMetadataSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The user impact of the assessment
func (o AssessmentsMetadataSubscriptionOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentsMetadataSubscription) pulumi.StringPtrOutput { return v.UserImpact }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentsMetadataSubscriptionOutput{})
}