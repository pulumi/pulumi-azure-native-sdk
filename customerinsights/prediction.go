// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The prediction resource format.
//
// Uses Azure REST API version 2017-04-26. In version 2.x of the Azure Native provider, it used API version 2017-04-26.
type Prediction struct {
	pulumi.CustomResourceState

	// Whether do auto analyze.
	AutoAnalyze pulumi.BoolOutput `pulumi:"autoAnalyze"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Description of the prediction.
	Description pulumi.StringMapOutput `pulumi:"description"`
	// Display name of the prediction.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// The prediction grades.
	Grades PredictionResponseGradesArrayOutput `pulumi:"grades"`
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes pulumi.StringArrayOutput `pulumi:"involvedInteractionTypes"`
	// KPI types involved in the prediction.
	InvolvedKpiTypes pulumi.StringArrayOutput `pulumi:"involvedKpiTypes"`
	// Relationships involved in the prediction.
	InvolvedRelationships pulumi.StringArrayOutput `pulumi:"involvedRelationships"`
	// Definition of the link mapping of prediction.
	Mappings PredictionResponseMappingsOutput `pulumi:"mappings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Negative outcome expression.
	NegativeOutcomeExpression pulumi.StringOutput `pulumi:"negativeOutcomeExpression"`
	// Positive outcome expression.
	PositiveOutcomeExpression pulumi.StringOutput `pulumi:"positiveOutcomeExpression"`
	// Name of the prediction.
	PredictionName pulumi.StringPtrOutput `pulumi:"predictionName"`
	// Primary profile type.
	PrimaryProfileType pulumi.StringOutput `pulumi:"primaryProfileType"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Scope expression.
	ScopeExpression pulumi.StringOutput `pulumi:"scopeExpression"`
	// Score label.
	ScoreLabel pulumi.StringOutput `pulumi:"scoreLabel"`
	// System generated entities.
	SystemGeneratedEntities PredictionResponseSystemGeneratedEntitiesOutput `pulumi:"systemGeneratedEntities"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrediction registers a new resource with the given unique name, arguments, and options.
func NewPrediction(ctx *pulumi.Context,
	name string, args *PredictionArgs, opts ...pulumi.ResourceOption) (*Prediction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoAnalyze == nil {
		return nil, errors.New("invalid value for required argument 'AutoAnalyze'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.Mappings == nil {
		return nil, errors.New("invalid value for required argument 'Mappings'")
	}
	if args.NegativeOutcomeExpression == nil {
		return nil, errors.New("invalid value for required argument 'NegativeOutcomeExpression'")
	}
	if args.PositiveOutcomeExpression == nil {
		return nil, errors.New("invalid value for required argument 'PositiveOutcomeExpression'")
	}
	if args.PrimaryProfileType == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryProfileType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScopeExpression == nil {
		return nil, errors.New("invalid value for required argument 'ScopeExpression'")
	}
	if args.ScoreLabel == nil {
		return nil, errors.New("invalid value for required argument 'ScoreLabel'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Prediction"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Prediction
	err := ctx.RegisterResource("azure-native:customerinsights:Prediction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrediction gets an existing Prediction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrediction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PredictionState, opts ...pulumi.ResourceOption) (*Prediction, error) {
	var resource Prediction
	err := ctx.ReadResource("azure-native:customerinsights:Prediction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Prediction resources.
type predictionState struct {
}

type PredictionState struct {
}

func (PredictionState) ElementType() reflect.Type {
	return reflect.TypeOf((*predictionState)(nil)).Elem()
}

type predictionArgs struct {
	// Whether do auto analyze.
	AutoAnalyze bool `pulumi:"autoAnalyze"`
	// Description of the prediction.
	Description map[string]string `pulumi:"description"`
	// Display name of the prediction.
	DisplayName map[string]string `pulumi:"displayName"`
	// The prediction grades.
	Grades []PredictionGrades `pulumi:"grades"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes []string `pulumi:"involvedInteractionTypes"`
	// KPI types involved in the prediction.
	InvolvedKpiTypes []string `pulumi:"involvedKpiTypes"`
	// Relationships involved in the prediction.
	InvolvedRelationships []string `pulumi:"involvedRelationships"`
	// Definition of the link mapping of prediction.
	Mappings PredictionMappings `pulumi:"mappings"`
	// Negative outcome expression.
	NegativeOutcomeExpression string `pulumi:"negativeOutcomeExpression"`
	// Positive outcome expression.
	PositiveOutcomeExpression string `pulumi:"positiveOutcomeExpression"`
	// Name of the prediction.
	PredictionName *string `pulumi:"predictionName"`
	// Primary profile type.
	PrimaryProfileType string `pulumi:"primaryProfileType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Scope expression.
	ScopeExpression string `pulumi:"scopeExpression"`
	// Score label.
	ScoreLabel string `pulumi:"scoreLabel"`
}

// The set of arguments for constructing a Prediction resource.
type PredictionArgs struct {
	// Whether do auto analyze.
	AutoAnalyze pulumi.BoolInput
	// Description of the prediction.
	Description pulumi.StringMapInput
	// Display name of the prediction.
	DisplayName pulumi.StringMapInput
	// The prediction grades.
	Grades PredictionGradesArrayInput
	// The name of the hub.
	HubName pulumi.StringInput
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes pulumi.StringArrayInput
	// KPI types involved in the prediction.
	InvolvedKpiTypes pulumi.StringArrayInput
	// Relationships involved in the prediction.
	InvolvedRelationships pulumi.StringArrayInput
	// Definition of the link mapping of prediction.
	Mappings PredictionMappingsInput
	// Negative outcome expression.
	NegativeOutcomeExpression pulumi.StringInput
	// Positive outcome expression.
	PositiveOutcomeExpression pulumi.StringInput
	// Name of the prediction.
	PredictionName pulumi.StringPtrInput
	// Primary profile type.
	PrimaryProfileType pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Scope expression.
	ScopeExpression pulumi.StringInput
	// Score label.
	ScoreLabel pulumi.StringInput
}

func (PredictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*predictionArgs)(nil)).Elem()
}

type PredictionInput interface {
	pulumi.Input

	ToPredictionOutput() PredictionOutput
	ToPredictionOutputWithContext(ctx context.Context) PredictionOutput
}

func (*Prediction) ElementType() reflect.Type {
	return reflect.TypeOf((**Prediction)(nil)).Elem()
}

func (i *Prediction) ToPredictionOutput() PredictionOutput {
	return i.ToPredictionOutputWithContext(context.Background())
}

func (i *Prediction) ToPredictionOutputWithContext(ctx context.Context) PredictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictionOutput)
}

type PredictionOutput struct{ *pulumi.OutputState }

func (PredictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Prediction)(nil)).Elem()
}

func (o PredictionOutput) ToPredictionOutput() PredictionOutput {
	return o
}

func (o PredictionOutput) ToPredictionOutputWithContext(ctx context.Context) PredictionOutput {
	return o
}

// Whether do auto analyze.
func (o PredictionOutput) AutoAnalyze() pulumi.BoolOutput {
	return o.ApplyT(func(v *Prediction) pulumi.BoolOutput { return v.AutoAnalyze }).(pulumi.BoolOutput)
}

// The Azure API version of the resource.
func (o PredictionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Description of the prediction.
func (o PredictionOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

// Display name of the prediction.
func (o PredictionOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

// The prediction grades.
func (o PredictionOutput) Grades() PredictionResponseGradesArrayOutput {
	return o.ApplyT(func(v *Prediction) PredictionResponseGradesArrayOutput { return v.Grades }).(PredictionResponseGradesArrayOutput)
}

// Interaction types involved in the prediction.
func (o PredictionOutput) InvolvedInteractionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringArrayOutput { return v.InvolvedInteractionTypes }).(pulumi.StringArrayOutput)
}

// KPI types involved in the prediction.
func (o PredictionOutput) InvolvedKpiTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringArrayOutput { return v.InvolvedKpiTypes }).(pulumi.StringArrayOutput)
}

// Relationships involved in the prediction.
func (o PredictionOutput) InvolvedRelationships() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringArrayOutput { return v.InvolvedRelationships }).(pulumi.StringArrayOutput)
}

// Definition of the link mapping of prediction.
func (o PredictionOutput) Mappings() PredictionResponseMappingsOutput {
	return o.ApplyT(func(v *Prediction) PredictionResponseMappingsOutput { return v.Mappings }).(PredictionResponseMappingsOutput)
}

// Resource name.
func (o PredictionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Negative outcome expression.
func (o PredictionOutput) NegativeOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.NegativeOutcomeExpression }).(pulumi.StringOutput)
}

// Positive outcome expression.
func (o PredictionOutput) PositiveOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.PositiveOutcomeExpression }).(pulumi.StringOutput)
}

// Name of the prediction.
func (o PredictionOutput) PredictionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringPtrOutput { return v.PredictionName }).(pulumi.StringPtrOutput)
}

// Primary profile type.
func (o PredictionOutput) PrimaryProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.PrimaryProfileType }).(pulumi.StringOutput)
}

// Provisioning state.
func (o PredictionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Scope expression.
func (o PredictionOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.ScopeExpression }).(pulumi.StringOutput)
}

// Score label.
func (o PredictionOutput) ScoreLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.ScoreLabel }).(pulumi.StringOutput)
}

// System generated entities.
func (o PredictionOutput) SystemGeneratedEntities() PredictionResponseSystemGeneratedEntitiesOutput {
	return o.ApplyT(func(v *Prediction) PredictionResponseSystemGeneratedEntitiesOutput { return v.SystemGeneratedEntities }).(PredictionResponseSystemGeneratedEntitiesOutput)
}

// The hub name.
func (o PredictionOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o PredictionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Prediction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PredictionOutput{})
}
