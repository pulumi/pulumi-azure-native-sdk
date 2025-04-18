// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A user-defined logical grouping of machines.
//
// Uses Azure REST API version 2015-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2015-11-01-preview.
type MachineGroup struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count pulumi.IntPtrOutput `pulumi:"count"`
	// User defined name for the group
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource ETAG.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Type of the machine group
	GroupType pulumi.StringPtrOutput `pulumi:"groupType"`
	// Additional resource type qualifier.
	// Expected value is 'machineGroup'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines MachineReferenceWithHintsResponseArrayOutput `pulumi:"machines"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachineGroup registers a new resource with the given unique name, arguments, and options.
func NewMachineGroup(ctx *pulumi.Context,
	name string, args *MachineGroupArgs, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("machineGroup")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:MachineGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MachineGroup
	err := ctx.RegisterResource("azure-native:operationalinsights:MachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineGroup gets an existing MachineGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineGroupState, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	var resource MachineGroup
	err := ctx.ReadResource("azure-native:operationalinsights:MachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineGroup resources.
type machineGroupState struct {
}

type MachineGroupState struct {
}

func (MachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupState)(nil)).Elem()
}

type machineGroupArgs struct {
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count *int `pulumi:"count"`
	// User defined name for the group
	DisplayName string `pulumi:"displayName"`
	// Type of the machine group
	GroupType *string `pulumi:"groupType"`
	// Additional resource type qualifier.
	// Expected value is 'machineGroup'.
	Kind string `pulumi:"kind"`
	// Machine Group resource name.
	MachineGroupName *string `pulumi:"machineGroupName"`
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines []MachineReferenceWithHints `pulumi:"machines"`
	// Resource group name within the specified subscriptionId.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// OMS workspace containing the resources of interest.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MachineGroup resource.
type MachineGroupArgs struct {
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count pulumi.IntPtrInput
	// User defined name for the group
	DisplayName pulumi.StringInput
	// Type of the machine group
	GroupType pulumi.StringPtrInput
	// Additional resource type qualifier.
	// Expected value is 'machineGroup'.
	Kind pulumi.StringInput
	// Machine Group resource name.
	MachineGroupName pulumi.StringPtrInput
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines MachineReferenceWithHintsArrayInput
	// Resource group name within the specified subscriptionId.
	ResourceGroupName pulumi.StringInput
	// OMS workspace containing the resources of interest.
	WorkspaceName pulumi.StringInput
}

func (MachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupArgs)(nil)).Elem()
}

type MachineGroupInput interface {
	pulumi.Input

	ToMachineGroupOutput() MachineGroupOutput
	ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput
}

func (*MachineGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (i *MachineGroup) ToMachineGroupOutput() MachineGroupOutput {
	return i.ToMachineGroupOutputWithContext(context.Background())
}

func (i *MachineGroup) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupOutput)
}

type MachineGroupOutput struct{ *pulumi.OutputState }

func (MachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (o MachineGroupOutput) ToMachineGroupOutput() MachineGroupOutput {
	return o
}

func (o MachineGroupOutput) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return o
}

// The Azure API version of the resource.
func (o MachineGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
func (o MachineGroupOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

// User defined name for the group
func (o MachineGroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource ETAG.
func (o MachineGroupOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Type of the machine group
func (o MachineGroupOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.GroupType }).(pulumi.StringPtrOutput)
}

// Additional resource type qualifier.
// Expected value is 'machineGroup'.
func (o MachineGroupOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
func (o MachineGroupOutput) Machines() MachineReferenceWithHintsResponseArrayOutput {
	return o.ApplyT(func(v *MachineGroup) MachineReferenceWithHintsResponseArrayOutput { return v.Machines }).(MachineReferenceWithHintsResponseArrayOutput)
}

// Resource name.
func (o MachineGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o MachineGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineGroupOutput{})
}
