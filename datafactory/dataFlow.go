// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data flow resource type.
//
// Uses Azure REST API version 2018-06-01. In version 2.x of the Azure Native provider, it used API version 2018-06-01.
type DataFlow struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Data flow properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataFlow registers a new resource with the given unique name, arguments, and options.
func NewDataFlow(ctx *pulumi.Context,
	name string, args *DataFlowArgs, opts ...pulumi.ResourceOption) (*DataFlow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:DataFlow"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataFlow
	err := ctx.RegisterResource("azure-native:datafactory:DataFlow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataFlow gets an existing DataFlow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataFlowState, opts ...pulumi.ResourceOption) (*DataFlow, error) {
	var resource DataFlow
	err := ctx.ReadResource("azure-native:datafactory:DataFlow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataFlow resources.
type dataFlowState struct {
}

type DataFlowState struct {
}

func (DataFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowState)(nil)).Elem()
}

type dataFlowArgs struct {
	// The data flow name.
	DataFlowName *string `pulumi:"dataFlowName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Data flow properties.
	Properties interface{} `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DataFlow resource.
type DataFlowArgs struct {
	// The data flow name.
	DataFlowName pulumi.StringPtrInput
	// The factory name.
	FactoryName pulumi.StringInput
	// Data flow properties.
	Properties pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (DataFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowArgs)(nil)).Elem()
}

type DataFlowInput interface {
	pulumi.Input

	ToDataFlowOutput() DataFlowOutput
	ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput
}

func (*DataFlow) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlow)(nil)).Elem()
}

func (i *DataFlow) ToDataFlowOutput() DataFlowOutput {
	return i.ToDataFlowOutputWithContext(context.Background())
}

func (i *DataFlow) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowOutput)
}

type DataFlowOutput struct{ *pulumi.OutputState }

func (DataFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlow)(nil)).Elem()
}

func (o DataFlowOutput) ToDataFlowOutput() DataFlowOutput {
	return o
}

func (o DataFlowOutput) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return o
}

// The Azure API version of the resource.
func (o DataFlowOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Etag identifies change in the resource.
func (o DataFlowOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The resource name.
func (o DataFlowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Data flow properties.
func (o DataFlowOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o DataFlowOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataFlowOutput{})
}
