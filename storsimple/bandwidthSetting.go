// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The bandwidth setting.
// API Version: 2017-06-01.
type BandwidthSetting struct {
	pulumi.CustomResourceState

	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schedules.
	Schedules BandwidthScheduleResponseArrayOutput `pulumi:"schedules"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The number of volumes that uses the bandwidth setting.
	VolumeCount pulumi.IntOutput `pulumi:"volumeCount"`
}

// NewBandwidthSetting registers a new resource with the given unique name, arguments, and options.
func NewBandwidthSetting(ctx *pulumi.Context,
	name string, args *BandwidthSettingArgs, opts ...pulumi.ResourceOption) (*BandwidthSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Schedules == nil {
		return nil, errors.New("invalid value for required argument 'Schedules'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:BandwidthSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource BandwidthSetting
	err := ctx.RegisterResource("azure-native:storsimple:BandwidthSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthSetting gets an existing BandwidthSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthSettingState, opts ...pulumi.ResourceOption) (*BandwidthSetting, error) {
	var resource BandwidthSetting
	err := ctx.ReadResource("azure-native:storsimple:BandwidthSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthSetting resources.
type bandwidthSettingState struct {
}

type BandwidthSettingState struct {
}

func (BandwidthSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthSettingState)(nil)).Elem()
}

type bandwidthSettingArgs struct {
	// The bandwidth setting name.
	BandwidthSettingName *string `pulumi:"bandwidthSettingName"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *Kind `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The schedules.
	Schedules []BandwidthSchedule `pulumi:"schedules"`
}

// The set of arguments for constructing a BandwidthSetting resource.
type BandwidthSettingArgs struct {
	// The bandwidth setting name.
	BandwidthSettingName pulumi.StringPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind KindPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The schedules.
	Schedules BandwidthScheduleArrayInput
}

func (BandwidthSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthSettingArgs)(nil)).Elem()
}

type BandwidthSettingInput interface {
	pulumi.Input

	ToBandwidthSettingOutput() BandwidthSettingOutput
	ToBandwidthSettingOutputWithContext(ctx context.Context) BandwidthSettingOutput
}

func (*BandwidthSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthSetting)(nil)).Elem()
}

func (i *BandwidthSetting) ToBandwidthSettingOutput() BandwidthSettingOutput {
	return i.ToBandwidthSettingOutputWithContext(context.Background())
}

func (i *BandwidthSetting) ToBandwidthSettingOutputWithContext(ctx context.Context) BandwidthSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthSettingOutput)
}

type BandwidthSettingOutput struct{ *pulumi.OutputState }

func (BandwidthSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthSetting)(nil)).Elem()
}

func (o BandwidthSettingOutput) ToBandwidthSettingOutput() BandwidthSettingOutput {
	return o
}

func (o BandwidthSettingOutput) ToBandwidthSettingOutputWithContext(ctx context.Context) BandwidthSettingOutput {
	return o
}

// The Kind of the object. Currently only Series8000 is supported
func (o BandwidthSettingOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthSetting) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the object.
func (o BandwidthSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The schedules.
func (o BandwidthSettingOutput) Schedules() BandwidthScheduleResponseArrayOutput {
	return o.ApplyT(func(v *BandwidthSetting) BandwidthScheduleResponseArrayOutput { return v.Schedules }).(BandwidthScheduleResponseArrayOutput)
}

// The hierarchical type of the object.
func (o BandwidthSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The number of volumes that uses the bandwidth setting.
func (o BandwidthSettingOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *BandwidthSetting) pulumi.IntOutput { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BandwidthSettingOutput{})
}