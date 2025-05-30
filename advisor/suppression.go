// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package advisor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
//
// Uses Azure REST API version 2023-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-01-01.
//
// Other available API versions: 2023-01-01, 2024-11-18-preview, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native advisor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Suppression struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Gets or sets the expiration time stamp.
	ExpirationTimeStamp pulumi.StringOutput `pulumi:"expirationTimeStamp"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The GUID of the suppression.
	SuppressionId pulumi.StringPtrOutput `pulumi:"suppressionId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The duration for which the suppression is valid.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSuppression registers a new resource with the given unique name, arguments, and options.
func NewSuppression(ctx *pulumi.Context,
	name string, args *SuppressionArgs, opts ...pulumi.ResourceOption) (*Suppression, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecommendationId == nil {
		return nil, errors.New("invalid value for required argument 'RecommendationId'")
	}
	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:advisor/v20160712preview:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20170331:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20170419:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20200101:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20220901:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20221001:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20230101:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20230901preview:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20241118preview:Suppression"),
		},
		{
			Type: pulumi.String("azure-native:advisor/v20250101:Suppression"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Suppression
	err := ctx.RegisterResource("azure-native:advisor:Suppression", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSuppression gets an existing Suppression resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSuppression(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuppressionState, opts ...pulumi.ResourceOption) (*Suppression, error) {
	var resource Suppression
	err := ctx.ReadResource("azure-native:advisor:Suppression", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Suppression resources.
type suppressionState struct {
}

type SuppressionState struct {
}

func (SuppressionState) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionState)(nil)).Elem()
}

type suppressionArgs struct {
	// The name of the suppression.
	Name *string `pulumi:"name"`
	// The recommendation ID.
	RecommendationId string `pulumi:"recommendationId"`
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri string `pulumi:"resourceUri"`
	// The GUID of the suppression.
	SuppressionId *string `pulumi:"suppressionId"`
	// The duration for which the suppression is valid.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Suppression resource.
type SuppressionArgs struct {
	// The name of the suppression.
	Name pulumi.StringPtrInput
	// The recommendation ID.
	RecommendationId pulumi.StringInput
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri pulumi.StringInput
	// The GUID of the suppression.
	SuppressionId pulumi.StringPtrInput
	// The duration for which the suppression is valid.
	Ttl pulumi.StringPtrInput
}

func (SuppressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionArgs)(nil)).Elem()
}

type SuppressionInput interface {
	pulumi.Input

	ToSuppressionOutput() SuppressionOutput
	ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput
}

func (*Suppression) ElementType() reflect.Type {
	return reflect.TypeOf((**Suppression)(nil)).Elem()
}

func (i *Suppression) ToSuppressionOutput() SuppressionOutput {
	return i.ToSuppressionOutputWithContext(context.Background())
}

func (i *Suppression) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionOutput)
}

type SuppressionOutput struct{ *pulumi.OutputState }

func (SuppressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Suppression)(nil)).Elem()
}

func (o SuppressionOutput) ToSuppressionOutput() SuppressionOutput {
	return o
}

func (o SuppressionOutput) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return o
}

// The Azure API version of the resource.
func (o SuppressionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Suppression) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the expiration time stamp.
func (o SuppressionOutput) ExpirationTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Suppression) pulumi.StringOutput { return v.ExpirationTimeStamp }).(pulumi.StringOutput)
}

// The name of the resource
func (o SuppressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Suppression) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The GUID of the suppression.
func (o SuppressionOutput) SuppressionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Suppression) pulumi.StringPtrOutput { return v.SuppressionId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SuppressionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Suppression) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The duration for which the suppression is valid.
func (o SuppressionOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Suppression) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SuppressionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Suppression) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SuppressionOutput{})
}
