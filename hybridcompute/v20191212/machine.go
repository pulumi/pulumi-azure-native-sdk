// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191212

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a hybrid machine.
//
// Deprecated: Version 2019-12-12 will be removed in v2 of the provider.
type Machine struct {
	pulumi.CustomResourceState

	// The hybrid machine agent full version.
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey pulumi.StringPtrOutput `pulumi:"clientPublicKey"`
	// Specifies the hybrid machine display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Details about the error state.
	ErrorDetails ErrorDetailResponseArrayOutput `pulumi:"errorDetails"`
	// Machine Extensions information
	Extensions MachineExtensionInstanceViewResponseArrayOutput `pulumi:"extensions"`
	Identity   MachineResponseIdentityPtrOutput                `pulumi:"identity"`
	// The time of the last status change.
	LastStatusChange pulumi.StringOutput `pulumi:"lastStatusChange"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata pertaining to the geographic location of the resource.
	LocationData LocationDataResponsePtrOutput `pulumi:"locationData"`
	// Specifies the hybrid machine FQDN.
	MachineFqdn pulumi.StringOutput `pulumi:"machineFqdn"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operating System running on the hybrid machine.
	OsName pulumi.StringOutput `pulumi:"osName"`
	// Specifies the operating system settings for the hybrid machine.
	OsProfile MachinePropertiesResponseOsProfilePtrOutput `pulumi:"osProfile"`
	// The version of Operating System running on the hybrid machine.
	OsVersion pulumi.StringOutput `pulumi:"osVersion"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The status of the hybrid machine agent.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the hybrid machine unique ID.
	VmId pulumi.StringPtrOutput `pulumi:"vmId"`
}

// NewMachine registers a new resource with the given unique name, arguments, and options.
func NewMachine(ctx *pulumi.Context,
	name string, args *MachineArgs, opts ...pulumi.ResourceOption) (*Machine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190318preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190802preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:Machine"),
		},
	})
	opts = append(opts, aliases)
	var resource Machine
	err := ctx.RegisterResource("azure-native:hybridcompute/v20191212:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachine gets an existing Machine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azure-native:hybridcompute/v20191212:Machine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Machine resources.
type machineState struct {
}

type MachineState struct {
}

func (MachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineState)(nil)).Elem()
}

type machineArgs struct {
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey *string `pulumi:"clientPublicKey"`
	// Machine Extensions information
	Extensions []MachineExtensionInstanceView `pulumi:"extensions"`
	Identity   *MachineIdentity               `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Metadata pertaining to the geographic location of the resource.
	LocationData *LocationData `pulumi:"locationData"`
	// The name of the hybrid machine.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the hybrid machine unique ID.
	VmId *string `pulumi:"vmId"`
}

// The set of arguments for constructing a Machine resource.
type MachineArgs struct {
	// Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey pulumi.StringPtrInput
	// Machine Extensions information
	Extensions MachineExtensionInstanceViewArrayInput
	Identity   MachineIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Metadata pertaining to the geographic location of the resource.
	LocationData LocationDataPtrInput
	// The name of the hybrid machine.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Specifies the hybrid machine unique ID.
	VmId pulumi.StringPtrInput
}

func (MachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineArgs)(nil)).Elem()
}

type MachineInput interface {
	pulumi.Input

	ToMachineOutput() MachineOutput
	ToMachineOutputWithContext(ctx context.Context) MachineOutput
}

func (*Machine) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (i *Machine) ToMachineOutput() MachineOutput {
	return i.ToMachineOutputWithContext(context.Background())
}

func (i *Machine) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineOutput)
}

type MachineOutput struct{ *pulumi.OutputState }

func (MachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (o MachineOutput) ToMachineOutput() MachineOutput {
	return o
}

func (o MachineOutput) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return o
}

// The hybrid machine agent full version.
func (o MachineOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

// Public Key that the client provides to be used during initial resource onboarding
func (o MachineOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

// Specifies the hybrid machine display name.
func (o MachineOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Details about the error state.
func (o MachineOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *Machine) ErrorDetailResponseArrayOutput { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

// Machine Extensions information
func (o MachineOutput) Extensions() MachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v *Machine) MachineExtensionInstanceViewResponseArrayOutput { return v.Extensions }).(MachineExtensionInstanceViewResponseArrayOutput)
}

func (o MachineOutput) Identity() MachineResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Machine) MachineResponseIdentityPtrOutput { return v.Identity }).(MachineResponseIdentityPtrOutput)
}

// The time of the last status change.
func (o MachineOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.LastStatusChange }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o MachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Metadata pertaining to the geographic location of the resource.
func (o MachineOutput) LocationData() LocationDataResponsePtrOutput {
	return o.ApplyT(func(v *Machine) LocationDataResponsePtrOutput { return v.LocationData }).(LocationDataResponsePtrOutput)
}

// Specifies the hybrid machine FQDN.
func (o MachineOutput) MachineFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.MachineFqdn }).(pulumi.StringOutput)
}

// The name of the resource
func (o MachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Operating System running on the hybrid machine.
func (o MachineOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

// Specifies the operating system settings for the hybrid machine.
func (o MachineOutput) OsProfile() MachinePropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v *Machine) MachinePropertiesResponseOsProfilePtrOutput { return v.OsProfile }).(MachinePropertiesResponseOsProfilePtrOutput)
}

// The version of Operating System running on the hybrid machine.
func (o MachineOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
}

// The provisioning state, which only appears in the response.
func (o MachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The status of the hybrid machine agent.
func (o MachineOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o MachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the hybrid machine unique ID.
func (o MachineOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.VmId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineOutput{})
}