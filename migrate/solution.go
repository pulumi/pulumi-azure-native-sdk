// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Solution REST Resource.
// API Version: 2018-09-01-preview.
type Solution struct {
	pulumi.CustomResourceState

	// Gets or sets the ETAG for optimistic concurrency control.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets the name of this REST resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the properties of the solution.
	Properties SolutionPropertiesResponseOutput `pulumi:"properties"`
	// Gets the type of this REST resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSolution registers a new resource with the given unique name, arguments, and options.
func NewSolution(ctx *pulumi.Context,
	name string, args *SolutionArgs, opts ...pulumi.ResourceOption) (*Solution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrateProjectName == nil {
		return nil, errors.New("invalid value for required argument 'MigrateProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20180901preview:Solution"),
		},
	})
	opts = append(opts, aliases)
	var resource Solution
	err := ctx.RegisterResource("azure-native:migrate:Solution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSolution gets an existing Solution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SolutionState, opts ...pulumi.ResourceOption) (*Solution, error) {
	var resource Solution
	err := ctx.ReadResource("azure-native:migrate:Solution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Solution resources.
type solutionState struct {
}

type SolutionState struct {
}

func (SolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*solutionState)(nil)).Elem()
}

type solutionArgs struct {
	// Name of the Azure Migrate project.
	MigrateProjectName string `pulumi:"migrateProjectName"`
	// Gets or sets the properties of the solution.
	Properties *SolutionProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that migrate project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Unique name of a migration solution within a migrate project.
	SolutionName *string `pulumi:"solutionName"`
}

// The set of arguments for constructing a Solution resource.
type SolutionArgs struct {
	// Name of the Azure Migrate project.
	MigrateProjectName pulumi.StringInput
	// Gets or sets the properties of the solution.
	Properties SolutionPropertiesPtrInput
	// Name of the Azure Resource Group that migrate project is part of.
	ResourceGroupName pulumi.StringInput
	// Unique name of a migration solution within a migrate project.
	SolutionName pulumi.StringPtrInput
}

func (SolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*solutionArgs)(nil)).Elem()
}

type SolutionInput interface {
	pulumi.Input

	ToSolutionOutput() SolutionOutput
	ToSolutionOutputWithContext(ctx context.Context) SolutionOutput
}

func (*Solution) ElementType() reflect.Type {
	return reflect.TypeOf((**Solution)(nil)).Elem()
}

func (i *Solution) ToSolutionOutput() SolutionOutput {
	return i.ToSolutionOutputWithContext(context.Background())
}

func (i *Solution) ToSolutionOutputWithContext(ctx context.Context) SolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionOutput)
}

type SolutionOutput struct{ *pulumi.OutputState }

func (SolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Solution)(nil)).Elem()
}

func (o SolutionOutput) ToSolutionOutput() SolutionOutput {
	return o
}

func (o SolutionOutput) ToSolutionOutputWithContext(ctx context.Context) SolutionOutput {
	return o
}

// Gets or sets the ETAG for optimistic concurrency control.
func (o SolutionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Gets the name of this REST resource.
func (o SolutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the properties of the solution.
func (o SolutionOutput) Properties() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *Solution) SolutionPropertiesResponseOutput { return v.Properties }).(SolutionPropertiesResponseOutput)
}

// Gets the type of this REST resource.
func (o SolutionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SolutionOutput{})
}