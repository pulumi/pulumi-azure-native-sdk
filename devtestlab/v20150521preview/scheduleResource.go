// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A schedule.
//
// Deprecated: Version 2015-05-21-preview will be removed in v2 of the provider.
type ScheduleResource struct {
	pulumi.CustomResourceState

	// The daily recurrence of the schedule.
	DailyRecurrence DayDetailsResponsePtrOutput `pulumi:"dailyRecurrence"`
	// The hourly recurrence of the schedule.
	HourlyRecurrence HourDetailsResponsePtrOutput `pulumi:"hourlyRecurrence"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The status of the schedule.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The task type of the schedule.
	TaskType pulumi.StringPtrOutput `pulumi:"taskType"`
	// The time zone id.
	TimeZoneId pulumi.StringPtrOutput `pulumi:"timeZoneId"`
	// The type of the resource.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The weekly recurrence of the schedule.
	WeeklyRecurrence WeekDetailsResponsePtrOutput `pulumi:"weeklyRecurrence"`
}

// NewScheduleResource registers a new resource with the given unique name, arguments, and options.
func NewScheduleResource(ctx *pulumi.Context,
	name string, args *ScheduleResourceArgs, opts ...pulumi.ResourceOption) (*ScheduleResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:ScheduleResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:ScheduleResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ScheduleResource"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduleResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:ScheduleResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduleResource gets an existing ScheduleResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduleResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleResourceState, opts ...pulumi.ResourceOption) (*ScheduleResource, error) {
	var resource ScheduleResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:ScheduleResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduleResource resources.
type scheduleResourceState struct {
}

type ScheduleResourceState struct {
}

func (ScheduleResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleResourceState)(nil)).Elem()
}

type scheduleResourceArgs struct {
	// The daily recurrence of the schedule.
	DailyRecurrence *DayDetails `pulumi:"dailyRecurrence"`
	// The hourly recurrence of the schedule.
	HourlyRecurrence *HourDetails `pulumi:"hourlyRecurrence"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the schedule.
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The task type of the schedule.
	TaskType *string `pulumi:"taskType"`
	// The time zone id.
	TimeZoneId *string `pulumi:"timeZoneId"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The weekly recurrence of the schedule.
	WeeklyRecurrence *WeekDetails `pulumi:"weeklyRecurrence"`
}

// The set of arguments for constructing a ScheduleResource resource.
type ScheduleResourceArgs struct {
	// The daily recurrence of the schedule.
	DailyRecurrence DayDetailsPtrInput
	// The hourly recurrence of the schedule.
	HourlyRecurrence HourDetailsPtrInput
	// The identifier of the resource.
	Id pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The status of the schedule.
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The task type of the schedule.
	TaskType pulumi.StringPtrInput
	// The time zone id.
	TimeZoneId pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The weekly recurrence of the schedule.
	WeeklyRecurrence WeekDetailsPtrInput
}

func (ScheduleResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleResourceArgs)(nil)).Elem()
}

type ScheduleResourceInput interface {
	pulumi.Input

	ToScheduleResourceOutput() ScheduleResourceOutput
	ToScheduleResourceOutputWithContext(ctx context.Context) ScheduleResourceOutput
}

func (*ScheduleResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResource)(nil)).Elem()
}

func (i *ScheduleResource) ToScheduleResourceOutput() ScheduleResourceOutput {
	return i.ToScheduleResourceOutputWithContext(context.Background())
}

func (i *ScheduleResource) ToScheduleResourceOutputWithContext(ctx context.Context) ScheduleResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResourceOutput)
}

type ScheduleResourceOutput struct{ *pulumi.OutputState }

func (ScheduleResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResource)(nil)).Elem()
}

func (o ScheduleResourceOutput) ToScheduleResourceOutput() ScheduleResourceOutput {
	return o
}

func (o ScheduleResourceOutput) ToScheduleResourceOutputWithContext(ctx context.Context) ScheduleResourceOutput {
	return o
}

// The daily recurrence of the schedule.
func (o ScheduleResourceOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResource) DayDetailsResponsePtrOutput { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

// The hourly recurrence of the schedule.
func (o ScheduleResourceOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResource) HourDetailsResponsePtrOutput { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

// The location of the resource.
func (o ScheduleResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ScheduleResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o ScheduleResourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The status of the schedule.
func (o ScheduleResourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o ScheduleResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The task type of the schedule.
func (o ScheduleResourceOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.TaskType }).(pulumi.StringPtrOutput)
}

// The time zone id.
func (o ScheduleResourceOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o ScheduleResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The weekly recurrence of the schedule.
func (o ScheduleResourceOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResource) WeekDetailsResponsePtrOutput { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleResourceOutput{})
}