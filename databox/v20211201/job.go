// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Job Resource.
type Job struct {
	pulumi.CustomResourceState

	// Reason for cancellation.
	CancellationReason pulumi.StringOutput `pulumi:"cancellationReason"`
	// Delivery Info of Job.
	DeliveryInfo JobDeliveryInfoResponsePtrOutput `pulumi:"deliveryInfo"`
	// Delivery type of Job.
	DeliveryType pulumi.StringPtrOutput `pulumi:"deliveryType"`
	// Details of a job run. This field will only be sent for expand details filter.
	Details pulumi.AnyOutput `pulumi:"details"`
	// Top level error for the job.
	Error CloudErrorResponseOutput `pulumi:"error"`
	// Msi identity of the resource
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Describes whether the job is cancellable or not.
	IsCancellable pulumi.BoolOutput `pulumi:"isCancellable"`
	// Flag to indicate cancellation of scheduled job.
	IsCancellableWithoutFee pulumi.BoolOutput `pulumi:"isCancellableWithoutFee"`
	// Describes whether the job is deletable or not.
	IsDeletable pulumi.BoolOutput `pulumi:"isDeletable"`
	// Is Prepare To Ship Enabled on this job
	IsPrepareToShipEnabled pulumi.BoolOutput `pulumi:"isPrepareToShipEnabled"`
	// Describes whether the shipping address is editable or not.
	IsShippingAddressEditable pulumi.BoolOutput `pulumi:"isShippingAddressEditable"`
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sku type.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Time at which the job was started in UTC ISO 8601 format.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Name of the stage which is in progress.
	Status pulumi.StringOutput `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the data transfer.
	TransferType pulumi.StringOutput `pulumi:"transferType"`
	// Type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.TransferType == nil {
		return nil, errors.New("invalid value for required argument 'TransferType'")
	}
	if isZero(args.DeliveryType) {
		args.DeliveryType = pulumi.StringPtr("NonScheduled")
	}
	if args.Identity != nil {
		args.Identity = args.Identity.ToResourceIdentityPtrOutput().ApplyT(func(v *ResourceIdentity) *ResourceIdentity { return v.Defaults() }).(ResourceIdentityPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databox:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20180101:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20190901:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20200401:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20201101:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210301:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210501:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20210801preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:databox/v20220201:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:databox/v20211201:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:databox/v20211201:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Delivery Info of Job.
	DeliveryInfo *JobDeliveryInfo `pulumi:"deliveryInfo"`
	// Delivery type of Job.
	DeliveryType *string `pulumi:"deliveryType"`
	// Details of a job run. This field will only be sent for expand details filter.
	Details interface{} `pulumi:"details"`
	// Msi identity of the resource
	Identity *ResourceIdentity `pulumi:"identity"`
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	JobName *string `pulumi:"jobName"`
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location *string `pulumi:"location"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku type.
	Sku Sku `pulumi:"sku"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags map[string]string `pulumi:"tags"`
	// Type of the data transfer.
	TransferType string `pulumi:"transferType"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Delivery Info of Job.
	DeliveryInfo JobDeliveryInfoPtrInput
	// Delivery type of Job.
	DeliveryType pulumi.StringPtrInput
	// Details of a job run. This field will only be sent for expand details filter.
	Details pulumi.Input
	// Msi identity of the resource
	Identity ResourceIdentityPtrInput
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	JobName pulumi.StringPtrInput
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location pulumi.StringPtrInput
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput
	// The sku type.
	Sku SkuInput
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags pulumi.StringMapInput
	// Type of the data transfer.
	TransferType pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// Reason for cancellation.
func (o JobOutput) CancellationReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.CancellationReason }).(pulumi.StringOutput)
}

// Delivery Info of Job.
func (o JobOutput) DeliveryInfo() JobDeliveryInfoResponsePtrOutput {
	return o.ApplyT(func(v *Job) JobDeliveryInfoResponsePtrOutput { return v.DeliveryInfo }).(JobDeliveryInfoResponsePtrOutput)
}

// Delivery type of Job.
func (o JobOutput) DeliveryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.DeliveryType }).(pulumi.StringPtrOutput)
}

// Details of a job run. This field will only be sent for expand details filter.
func (o JobOutput) Details() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.Details }).(pulumi.AnyOutput)
}

// Top level error for the job.
func (o JobOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v *Job) CloudErrorResponseOutput { return v.Error }).(CloudErrorResponseOutput)
}

// Msi identity of the resource
func (o JobOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Job) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Describes whether the job is cancellable or not.
func (o JobOutput) IsCancellable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsCancellable }).(pulumi.BoolOutput)
}

// Flag to indicate cancellation of scheduled job.
func (o JobOutput) IsCancellableWithoutFee() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsCancellableWithoutFee }).(pulumi.BoolOutput)
}

// Describes whether the job is deletable or not.
func (o JobOutput) IsDeletable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsDeletable }).(pulumi.BoolOutput)
}

// Is Prepare To Ship Enabled on this job
func (o JobOutput) IsPrepareToShipEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsPrepareToShipEnabled }).(pulumi.BoolOutput)
}

// Describes whether the shipping address is editable or not.
func (o JobOutput) IsShippingAddressEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.IsShippingAddressEditable }).(pulumi.BoolOutput)
}

// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
func (o JobOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the object.
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The sku type.
func (o JobOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Job) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Time at which the job was started in UTC ISO 8601 format.
func (o JobOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Name of the stage which is in progress.
func (o JobOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o JobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Job) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
func (o JobOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the data transfer.
func (o JobOutput) TransferType() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.TransferType }).(pulumi.StringOutput)
}

// Type of the object.
func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}