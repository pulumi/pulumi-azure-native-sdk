// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A streaming job object, containing all information associated with the named streaming job.
type StreamingJob struct {
	pulumi.CustomResourceState

	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrOutput `pulumi:"compatibilityLevel"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrOutput `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrOutput `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionResponseArrayOutput `pulumi:"functions"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputResponseArrayOutput `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState pulumi.StringOutput `pulumi:"jobState"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime pulumi.StringOutput `pulumi:"lastOutputEventTime"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrOutput `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrOutput `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrOutput `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputResponseArrayOutput `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationResponsePtrOutput `pulumi:"transformation"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingJob registers a new resource with the given unique name, arguments, and options.
func NewStreamingJob(ctx *pulumi.Context,
	name string, args *StreamingJobArgs, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:streamanalytics:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20170401preview:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20200301:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20211001preview:StreamingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingJob
	err := ctx.RegisterResource("azure-native:streamanalytics/v20160301:StreamingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingJob gets an existing StreamingJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingJobState, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	var resource StreamingJob
	err := ctx.ReadResource("azure-native:streamanalytics/v20160301:StreamingJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingJob resources.
type streamingJobState struct {
}

type StreamingJobState struct {
}

func (StreamingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobState)(nil)).Elem()
}

type streamingJobArgs struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionType `pulumi:"functions"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputType `pulumi:"inputs"`
	// The name of the streaming job.
	JobName *string `pulumi:"jobName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputType `pulumi:"outputs"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *Transformation `pulumi:"transformation"`
}

// The set of arguments for constructing a StreamingJob resource.
type StreamingJobArgs struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrInput
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrInput
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionTypeArrayInput
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputTypeArrayInput
	// The name of the streaming job.
	JobName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrInput
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrInput
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputTypeArrayInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationPtrInput
}

func (StreamingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobArgs)(nil)).Elem()
}

type StreamingJobInput interface {
	pulumi.Input

	ToStreamingJobOutput() StreamingJobOutput
	ToStreamingJobOutputWithContext(ctx context.Context) StreamingJobOutput
}

func (*StreamingJob) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJob)(nil)).Elem()
}

func (i *StreamingJob) ToStreamingJobOutput() StreamingJobOutput {
	return i.ToStreamingJobOutputWithContext(context.Background())
}

func (i *StreamingJob) ToStreamingJobOutputWithContext(ctx context.Context) StreamingJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingJobOutput)
}

type StreamingJobOutput struct{ *pulumi.OutputState }

func (StreamingJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJob)(nil)).Elem()
}

func (o StreamingJobOutput) ToStreamingJobOutput() StreamingJobOutput {
	return o
}

func (o StreamingJobOutput) ToStreamingJobOutputWithContext(ctx context.Context) StreamingJobOutput {
	return o
}

// Controls certain runtime behaviors of the streaming job.
func (o StreamingJobOutput) CompatibilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.CompatibilityLevel }).(pulumi.StringPtrOutput)
}

// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
func (o StreamingJobOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
func (o StreamingJobOutput) DataLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.DataLocale }).(pulumi.StringPtrOutput)
}

// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
func (o StreamingJobOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
func (o StreamingJobOutput) EventsLateArrivalMaxDelayInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.IntPtrOutput { return v.EventsLateArrivalMaxDelayInSeconds }).(pulumi.IntPtrOutput)
}

// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
func (o StreamingJobOutput) EventsOutOfOrderMaxDelayInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.IntPtrOutput { return v.EventsOutOfOrderMaxDelayInSeconds }).(pulumi.IntPtrOutput)
}

// Indicates the policy to apply to events that arrive out of order in the input event stream.
func (o StreamingJobOutput) EventsOutOfOrderPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.EventsOutOfOrderPolicy }).(pulumi.StringPtrOutput)
}

// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
func (o StreamingJobOutput) Functions() FunctionResponseArrayOutput {
	return o.ApplyT(func(v *StreamingJob) FunctionResponseArrayOutput { return v.Functions }).(FunctionResponseArrayOutput)
}

// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
func (o StreamingJobOutput) Inputs() InputResponseArrayOutput {
	return o.ApplyT(func(v *StreamingJob) InputResponseArrayOutput { return v.Inputs }).(InputResponseArrayOutput)
}

// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
func (o StreamingJobOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

// Describes the state of the streaming job.
func (o StreamingJobOutput) JobState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.JobState }).(pulumi.StringOutput)
}

// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
func (o StreamingJobOutput) LastOutputEventTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.LastOutputEventTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o StreamingJobOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o StreamingJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
func (o StreamingJobOutput) OutputErrorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.OutputErrorPolicy }).(pulumi.StringPtrOutput)
}

// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
func (o StreamingJobOutput) OutputStartMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.OutputStartMode }).(pulumi.StringPtrOutput)
}

// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
func (o StreamingJobOutput) OutputStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringPtrOutput { return v.OutputStartTime }).(pulumi.StringPtrOutput)
}

// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
func (o StreamingJobOutput) Outputs() OutputResponseArrayOutput {
	return o.ApplyT(func(v *StreamingJob) OutputResponseArrayOutput { return v.Outputs }).(OutputResponseArrayOutput)
}

// Describes the provisioning status of the streaming job.
func (o StreamingJobOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
func (o StreamingJobOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *StreamingJob) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags.
func (o StreamingJobOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
func (o StreamingJobOutput) Transformation() TransformationResponsePtrOutput {
	return o.ApplyT(func(v *StreamingJob) TransformationResponsePtrOutput { return v.Transformation }).(TransformationResponsePtrOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o StreamingJobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingJob) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingJobOutput{})
}