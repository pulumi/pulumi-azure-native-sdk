// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recommendationsservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Modeling resource details.
// API Version: 2022-02-01.
type Modeling struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Modeling resource properties.
	Properties ModelingResourceResponsePropertiesOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewModeling registers a new resource with the given unique name, arguments, and options.
func NewModeling(ctx *pulumi.Context,
	name string, args *ModelingArgs, opts ...pulumi.ResourceOption) (*Modeling, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recommendationsservice/v20220201:Modeling"),
		},
		{
			Type: pulumi.String("azure-native:recommendationsservice/v20220301preview:Modeling"),
		},
	})
	opts = append(opts, aliases)
	var resource Modeling
	err := ctx.RegisterResource("azure-native:recommendationsservice:Modeling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModeling gets an existing Modeling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModeling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelingState, opts ...pulumi.ResourceOption) (*Modeling, error) {
	var resource Modeling
	err := ctx.ReadResource("azure-native:recommendationsservice:Modeling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Modeling resources.
type modelingState struct {
}

type ModelingState struct {
}

func (ModelingState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelingState)(nil)).Elem()
}

type modelingArgs struct {
	// The name of the RecommendationsService Account resource.
	AccountName string `pulumi:"accountName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Modeling resource.
	ModelingName *string `pulumi:"modelingName"`
	// Modeling resource properties.
	Properties *ModelingResourceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Modeling resource.
type ModelingArgs struct {
	// The name of the RecommendationsService Account resource.
	AccountName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Modeling resource.
	ModelingName pulumi.StringPtrInput
	// Modeling resource properties.
	Properties ModelingResourcePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ModelingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelingArgs)(nil)).Elem()
}

type ModelingInput interface {
	pulumi.Input

	ToModelingOutput() ModelingOutput
	ToModelingOutputWithContext(ctx context.Context) ModelingOutput
}

func (*Modeling) ElementType() reflect.Type {
	return reflect.TypeOf((**Modeling)(nil)).Elem()
}

func (i *Modeling) ToModelingOutput() ModelingOutput {
	return i.ToModelingOutputWithContext(context.Background())
}

func (i *Modeling) ToModelingOutputWithContext(ctx context.Context) ModelingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingOutput)
}

type ModelingOutput struct{ *pulumi.OutputState }

func (ModelingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Modeling)(nil)).Elem()
}

func (o ModelingOutput) ToModelingOutput() ModelingOutput {
	return o
}

func (o ModelingOutput) ToModelingOutputWithContext(ctx context.Context) ModelingOutput {
	return o
}

// The geo-location where the resource lives
func (o ModelingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ModelingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Modeling resource properties.
func (o ModelingOutput) Properties() ModelingResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *Modeling) ModelingResourceResponsePropertiesOutput { return v.Properties }).(ModelingResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ModelingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Modeling) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ModelingOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ModelingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelingOutput{})
}