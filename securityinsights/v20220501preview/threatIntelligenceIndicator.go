// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Threat intelligence information object.
type ThreatIntelligenceIndicator struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the entity.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewThreatIntelligenceIndicator registers a new resource with the given unique name, arguments, and options.
func NewThreatIntelligenceIndicator(ctx *pulumi.Context,
	name string, args *ThreatIntelligenceIndicatorArgs, opts ...pulumi.ResourceOption) (*ThreatIntelligenceIndicator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("indicator")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ThreatIntelligenceIndicator"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ThreatIntelligenceIndicator"),
		},
	})
	opts = append(opts, aliases)
	var resource ThreatIntelligenceIndicator
	err := ctx.RegisterResource("azure-native:securityinsights/v20220501preview:ThreatIntelligenceIndicator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThreatIntelligenceIndicator gets an existing ThreatIntelligenceIndicator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreatIntelligenceIndicator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelligenceIndicatorState, opts ...pulumi.ResourceOption) (*ThreatIntelligenceIndicator, error) {
	var resource ThreatIntelligenceIndicator
	err := ctx.ReadResource("azure-native:securityinsights/v20220501preview:ThreatIntelligenceIndicator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThreatIntelligenceIndicator resources.
type threatIntelligenceIndicatorState struct {
}

type ThreatIntelligenceIndicatorState struct {
}

func (ThreatIntelligenceIndicatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceIndicatorState)(nil)).Elem()
}

type threatIntelligenceIndicatorArgs struct {
	// Confidence of threat intelligence entity
	Confidence *int `pulumi:"confidence"`
	// Created by
	Created *string `pulumi:"created"`
	// Created by reference of threat intelligence entity
	CreatedByRef *string `pulumi:"createdByRef"`
	// Is threat intelligence entity defanged
	Defanged *bool `pulumi:"defanged"`
	// Description of a threat intelligence entity
	Description *string `pulumi:"description"`
	// Display name of a threat intelligence entity
	DisplayName *string `pulumi:"displayName"`
	// Extensions map
	Extensions interface{} `pulumi:"extensions"`
	// External ID of threat intelligence entity
	ExternalId *string `pulumi:"externalId"`
	// External last updated time in UTC
	ExternalLastUpdatedTimeUtc *string `pulumi:"externalLastUpdatedTimeUtc"`
	// External References
	ExternalReferences []ThreatIntelligenceExternalReference `pulumi:"externalReferences"`
	// Granular Markings
	GranularMarkings []ThreatIntelligenceGranularMarkingModel `pulumi:"granularMarkings"`
	// Indicator types of threat intelligence entities
	IndicatorTypes []string `pulumi:"indicatorTypes"`
	// Kill chain phases
	KillChainPhases []ThreatIntelligenceKillChainPhase `pulumi:"killChainPhases"`
	// The kind of the threat intelligence entity
	// Expected value is 'indicator'.
	Kind string `pulumi:"kind"`
	// Labels  of threat intelligence entity
	Labels []string `pulumi:"labels"`
	// Language of threat intelligence entity
	Language *string `pulumi:"language"`
	// Last updated time in UTC
	LastUpdatedTimeUtc *string `pulumi:"lastUpdatedTimeUtc"`
	// Modified by
	Modified *string `pulumi:"modified"`
	// Threat intelligence indicator name field.
	Name *string `pulumi:"name"`
	// Threat intelligence entity object marking references
	ObjectMarkingRefs []string `pulumi:"objectMarkingRefs"`
	// Parsed patterns
	ParsedPattern []ThreatIntelligenceParsedPattern `pulumi:"parsedPattern"`
	// Pattern of a threat intelligence entity
	Pattern *string `pulumi:"pattern"`
	// Pattern type of a threat intelligence entity
	PatternType *string `pulumi:"patternType"`
	// Pattern version of a threat intelligence entity
	PatternVersion *string `pulumi:"patternVersion"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Is threat intelligence entity revoked
	Revoked *bool `pulumi:"revoked"`
	// Source of a threat intelligence entity
	Source *string `pulumi:"source"`
	// List of tags
	ThreatIntelligenceTags []string `pulumi:"threatIntelligenceTags"`
	// Threat types
	ThreatTypes []string `pulumi:"threatTypes"`
	// Valid from
	ValidFrom *string `pulumi:"validFrom"`
	// Valid until
	ValidUntil *string `pulumi:"validUntil"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ThreatIntelligenceIndicator resource.
type ThreatIntelligenceIndicatorArgs struct {
	// Confidence of threat intelligence entity
	Confidence pulumi.IntPtrInput
	// Created by
	Created pulumi.StringPtrInput
	// Created by reference of threat intelligence entity
	CreatedByRef pulumi.StringPtrInput
	// Is threat intelligence entity defanged
	Defanged pulumi.BoolPtrInput
	// Description of a threat intelligence entity
	Description pulumi.StringPtrInput
	// Display name of a threat intelligence entity
	DisplayName pulumi.StringPtrInput
	// Extensions map
	Extensions pulumi.Input
	// External ID of threat intelligence entity
	ExternalId pulumi.StringPtrInput
	// External last updated time in UTC
	ExternalLastUpdatedTimeUtc pulumi.StringPtrInput
	// External References
	ExternalReferences ThreatIntelligenceExternalReferenceArrayInput
	// Granular Markings
	GranularMarkings ThreatIntelligenceGranularMarkingModelArrayInput
	// Indicator types of threat intelligence entities
	IndicatorTypes pulumi.StringArrayInput
	// Kill chain phases
	KillChainPhases ThreatIntelligenceKillChainPhaseArrayInput
	// The kind of the threat intelligence entity
	// Expected value is 'indicator'.
	Kind pulumi.StringInput
	// Labels  of threat intelligence entity
	Labels pulumi.StringArrayInput
	// Language of threat intelligence entity
	Language pulumi.StringPtrInput
	// Last updated time in UTC
	LastUpdatedTimeUtc pulumi.StringPtrInput
	// Modified by
	Modified pulumi.StringPtrInput
	// Threat intelligence indicator name field.
	Name pulumi.StringPtrInput
	// Threat intelligence entity object marking references
	ObjectMarkingRefs pulumi.StringArrayInput
	// Parsed patterns
	ParsedPattern ThreatIntelligenceParsedPatternArrayInput
	// Pattern of a threat intelligence entity
	Pattern pulumi.StringPtrInput
	// Pattern type of a threat intelligence entity
	PatternType pulumi.StringPtrInput
	// Pattern version of a threat intelligence entity
	PatternVersion pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Is threat intelligence entity revoked
	Revoked pulumi.BoolPtrInput
	// Source of a threat intelligence entity
	Source pulumi.StringPtrInput
	// List of tags
	ThreatIntelligenceTags pulumi.StringArrayInput
	// Threat types
	ThreatTypes pulumi.StringArrayInput
	// Valid from
	ValidFrom pulumi.StringPtrInput
	// Valid until
	ValidUntil pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ThreatIntelligenceIndicatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceIndicatorArgs)(nil)).Elem()
}

type ThreatIntelligenceIndicatorInput interface {
	pulumi.Input

	ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput
	ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput
}

func (*ThreatIntelligenceIndicator) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceIndicator)(nil)).Elem()
}

func (i *ThreatIntelligenceIndicator) ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput {
	return i.ToThreatIntelligenceIndicatorOutputWithContext(context.Background())
}

func (i *ThreatIntelligenceIndicator) ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceIndicatorOutput)
}

type ThreatIntelligenceIndicatorOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceIndicatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceIndicator)(nil)).Elem()
}

func (o ThreatIntelligenceIndicatorOutput) ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput {
	return o
}

func (o ThreatIntelligenceIndicatorOutput) ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput {
	return o
}

// Etag of the azure resource
func (o ThreatIntelligenceIndicatorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the entity.
func (o ThreatIntelligenceIndicatorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o ThreatIntelligenceIndicatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ThreatIntelligenceIndicatorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ThreatIntelligenceIndicatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceIndicator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelligenceIndicatorOutput{})
}