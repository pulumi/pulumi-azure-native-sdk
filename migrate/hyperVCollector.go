// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.
type HyperVCollector struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput               `pulumi:"azureApiVersion"`
	ETag            pulumi.StringPtrOutput            `pulumi:"eTag"`
	Name            pulumi.StringOutput               `pulumi:"name"`
	Properties      CollectorPropertiesResponseOutput `pulumi:"properties"`
	Type            pulumi.StringOutput               `pulumi:"type"`
}

// NewHyperVCollector registers a new resource with the given unique name, arguments, and options.
func NewHyperVCollector(ctx *pulumi.Context,
	name string, args *HyperVCollectorArgs, opts ...pulumi.ResourceOption) (*HyperVCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230315:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230315:HypervCollectorsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230401preview:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230401preview:HypervCollectorsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230501preview:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230501preview:HypervCollectorsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230909preview:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230909preview:HypervCollectorsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20240101preview:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20240101preview:HypervCollectorsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20240303preview:HyperVCollector"),
		},
		{
			Type: pulumi.String("azure-native:migrate:HypervCollectorsOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HyperVCollector
	err := ctx.RegisterResource("azure-native:migrate:HyperVCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHyperVCollector gets an existing HyperVCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHyperVCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HyperVCollectorState, opts ...pulumi.ResourceOption) (*HyperVCollector, error) {
	var resource HyperVCollector
	err := ctx.ReadResource("azure-native:migrate:HyperVCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HyperVCollector resources.
type hyperVCollectorState struct {
}

type HyperVCollectorState struct {
}

func (HyperVCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVCollectorState)(nil)).Elem()
}

type hyperVCollectorArgs struct {
	ETag *string `pulumi:"eTag"`
	// Unique name of a Hyper-V collector within a project.
	HyperVCollectorName *string `pulumi:"hyperVCollectorName"`
	// Name of the Azure Migrate project.
	ProjectName string               `pulumi:"projectName"`
	Properties  *CollectorProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a HyperVCollector resource.
type HyperVCollectorArgs struct {
	ETag pulumi.StringPtrInput
	// Unique name of a Hyper-V collector within a project.
	HyperVCollectorName pulumi.StringPtrInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	Properties  CollectorPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
}

func (HyperVCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVCollectorArgs)(nil)).Elem()
}

type HyperVCollectorInput interface {
	pulumi.Input

	ToHyperVCollectorOutput() HyperVCollectorOutput
	ToHyperVCollectorOutputWithContext(ctx context.Context) HyperVCollectorOutput
}

func (*HyperVCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**HyperVCollector)(nil)).Elem()
}

func (i *HyperVCollector) ToHyperVCollectorOutput() HyperVCollectorOutput {
	return i.ToHyperVCollectorOutputWithContext(context.Background())
}

func (i *HyperVCollector) ToHyperVCollectorOutputWithContext(ctx context.Context) HyperVCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVCollectorOutput)
}

type HyperVCollectorOutput struct{ *pulumi.OutputState }

func (HyperVCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HyperVCollector)(nil)).Elem()
}

func (o HyperVCollectorOutput) ToHyperVCollectorOutput() HyperVCollectorOutput {
	return o
}

func (o HyperVCollectorOutput) ToHyperVCollectorOutputWithContext(ctx context.Context) HyperVCollectorOutput {
	return o
}

// The Azure API version of the resource.
func (o HyperVCollectorOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *HyperVCollector) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

func (o HyperVCollectorOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HyperVCollector) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o HyperVCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HyperVCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HyperVCollectorOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *HyperVCollector) CollectorPropertiesResponseOutput { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o HyperVCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HyperVCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HyperVCollectorOutput{})
}
