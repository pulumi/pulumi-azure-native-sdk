// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A data set mapping data transfer object.
//
// Deprecated: Version 2019-11-01 will be removed in v2 of the provider.
type DataSetMapping struct {
	pulumi.CustomResourceState

	// Kind of data set mapping.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewDataSetMapping(ctx *pulumi.Context,
	name string, args *DataSetMappingArgs, opts ...pulumi.ResourceOption) (*DataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:DataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource DataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20191101:DataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSetMapping gets an existing DataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSetMappingState, opts ...pulumi.ResourceOption) (*DataSetMapping, error) {
	var resource DataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20191101:DataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSetMapping resources.
type dataSetMappingState struct {
}

type DataSetMappingState struct {
}

func (DataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetMappingState)(nil)).Elem()
}

type dataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// Kind of data set mapping.
	Kind string `pulumi:"kind"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// The set of arguments for constructing a DataSetMapping resource.
type DataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// Kind of data set mapping.
	Kind pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName pulumi.StringInput
}

func (DataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetMappingArgs)(nil)).Elem()
}

type DataSetMappingInput interface {
	pulumi.Input

	ToDataSetMappingOutput() DataSetMappingOutput
	ToDataSetMappingOutputWithContext(ctx context.Context) DataSetMappingOutput
}

func (*DataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSetMapping)(nil)).Elem()
}

func (i *DataSetMapping) ToDataSetMappingOutput() DataSetMappingOutput {
	return i.ToDataSetMappingOutputWithContext(context.Background())
}

func (i *DataSetMapping) ToDataSetMappingOutputWithContext(ctx context.Context) DataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSetMappingOutput)
}

type DataSetMappingOutput struct{ *pulumi.OutputState }

func (DataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSetMapping)(nil)).Elem()
}

func (o DataSetMappingOutput) ToDataSetMappingOutput() DataSetMappingOutput {
	return o
}

func (o DataSetMappingOutput) ToDataSetMappingOutputWithContext(ctx context.Context) DataSetMappingOutput {
	return o
}

// Kind of data set mapping.
func (o DataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o DataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Type of the azure resource
func (o DataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSetMappingOutput{})
}