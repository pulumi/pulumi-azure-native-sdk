// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
type Output struct {
	pulumi.CustomResourceState

	// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
	Datasource pulumi.AnyOutput `pulumi:"datasource"`
	// Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
	Diagnostics DiagnosticsResponseOutput `pulumi:"diagnostics"`
	// The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
	Serialization pulumi.AnyOutput `pulumi:"serialization"`
	// The size window to constrain a Stream Analytics output to.
	SizeWindow pulumi.Float64PtrOutput `pulumi:"sizeWindow"`
	// The time frame for filtering Stream Analytics job outputs.
	TimeWindow pulumi.StringPtrOutput `pulumi:"timeWindow"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOutput registers a new resource with the given unique name, arguments, and options.
func NewOutput(ctx *pulumi.Context,
	name string, args *OutputArgs, opts ...pulumi.ResourceOption) (*Output, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobName == nil {
		return nil, errors.New("invalid value for required argument 'JobName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:streamanalytics:Output"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20160301:Output"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20170401preview:Output"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20211001preview:Output"),
		},
	})
	opts = append(opts, aliases)
	var resource Output
	err := ctx.RegisterResource("azure-native:streamanalytics/v20200301:Output", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutput gets an existing Output resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputState, opts ...pulumi.ResourceOption) (*Output, error) {
	var resource Output
	err := ctx.ReadResource("azure-native:streamanalytics/v20200301:Output", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Output resources.
type outputState struct {
}

type OutputState struct {
}

func (OutputState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputState)(nil)).Elem()
}

type outputArgs struct {
	// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
	Datasource interface{} `pulumi:"datasource"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// Resource name
	Name *string `pulumi:"name"`
	// The name of the output.
	OutputName *string `pulumi:"outputName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
	Serialization interface{} `pulumi:"serialization"`
	// The size window to constrain a Stream Analytics output to.
	SizeWindow *float64 `pulumi:"sizeWindow"`
	// The time frame for filtering Stream Analytics job outputs.
	TimeWindow *string `pulumi:"timeWindow"`
}

// The set of arguments for constructing a Output resource.
type OutputArgs struct {
	// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
	Datasource pulumi.Input
	// The name of the streaming job.
	JobName pulumi.StringInput
	// Resource name
	Name pulumi.StringPtrInput
	// The name of the output.
	OutputName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
	Serialization pulumi.Input
	// The size window to constrain a Stream Analytics output to.
	SizeWindow pulumi.Float64PtrInput
	// The time frame for filtering Stream Analytics job outputs.
	TimeWindow pulumi.StringPtrInput
}

func (OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputArgs)(nil)).Elem()
}

type OutputInput interface {
	pulumi.Input

	ToOutputOutput() OutputOutput
	ToOutputOutputWithContext(ctx context.Context) OutputOutput
}

func (*Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Output)(nil)).Elem()
}

func (i *Output) ToOutputOutput() OutputOutput {
	return i.ToOutputOutputWithContext(context.Background())
}

func (i *Output) ToOutputOutputWithContext(ctx context.Context) OutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputOutput)
}

type OutputOutput struct{ *pulumi.OutputState }

func (OutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Output)(nil)).Elem()
}

func (o OutputOutput) ToOutputOutput() OutputOutput {
	return o
}

func (o OutputOutput) ToOutputOutputWithContext(ctx context.Context) OutputOutput {
	return o
}

// Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
func (o OutputOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v *Output) pulumi.AnyOutput { return v.Datasource }).(pulumi.AnyOutput)
}

// Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
func (o OutputOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v *Output) DiagnosticsResponseOutput { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

// The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
func (o OutputOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Output) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name
func (o OutputOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Output) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
func (o OutputOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v *Output) pulumi.AnyOutput { return v.Serialization }).(pulumi.AnyOutput)
}

// The size window to constrain a Stream Analytics output to.
func (o OutputOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Output) pulumi.Float64PtrOutput { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

// The time frame for filtering Stream Analytics job outputs.
func (o OutputOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Output) pulumi.StringPtrOutput { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

// Resource type
func (o OutputOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Output) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OutputOutput{})
}