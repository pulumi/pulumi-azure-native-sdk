// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Virtual Machine Scale Set.
//
// Deprecated: Version 2019-03-01 will be removed in v2 of the provider.
type VirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities AdditionalCapabilitiesResponsePtrOutput `pulumi:"additionalCapabilities"`
	// Policy for automatic repairs.
	AutomaticRepairsPolicy AutomaticRepairsPolicyResponsePtrOutput `pulumi:"automaticRepairsPolicy"`
	// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrOutput `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision pulumi.BoolPtrOutput `pulumi:"overprovision"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// Fault Domain count for each placement group.
	PlatformFaultDomainCount pulumi.IntPtrOutput `pulumi:"platformFaultDomainCount"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourceResponsePtrOutput `pulumi:"proximityPlacementGroup"`
	// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
	ScaleInPolicy ScaleInPolicyResponsePtrOutput `pulumi:"scaleInPolicy"`
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup pulumi.BoolPtrOutput `pulumi:"singlePlacementGroup"`
	// The virtual machine scale set sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the ID which uniquely identifies a Virtual Machine Scale Set.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// The upgrade policy.
	UpgradePolicy UpgradePolicyResponsePtrOutput `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile VirtualMachineScaleSetVMProfileResponsePtrOutput `pulumi:"virtualMachineProfile"`
	// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
	ZoneBalance pulumi.BoolPtrOutput `pulumi:"zoneBalance"`
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSet
	err := ctx.RegisterResource("azure-native:compute/v20190301:VirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSet gets an existing VirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	var resource VirtualMachineScaleSet
	err := ctx.ReadResource("azure-native:compute/v20190301:VirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSet resources.
type virtualMachineScaleSetState struct {
}

type VirtualMachineScaleSetState struct {
}

func (VirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetState)(nil)).Elem()
}

type virtualMachineScaleSetArgs struct {
	// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities *AdditionalCapabilities `pulumi:"additionalCapabilities"`
	// Policy for automatic repairs.
	AutomaticRepairsPolicy *AutomaticRepairsPolicy `pulumi:"automaticRepairsPolicy"`
	// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
	DoNotRunExtensionsOnOverprovisionedVMs *bool `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentity `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision *bool `pulumi:"overprovision"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *Plan `pulumi:"plan"`
	// Fault Domain count for each placement group.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResource `pulumi:"proximityPlacementGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
	ScaleInPolicy *ScaleInPolicy `pulumi:"scaleInPolicy"`
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// The virtual machine scale set sku.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The upgrade policy.
	UpgradePolicy *UpgradePolicy `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfile `pulumi:"virtualMachineProfile"`
	// The name of the VM scale set to create or update.
	VmScaleSetName *string `pulumi:"vmScaleSetName"`
	// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
	ZoneBalance *bool `pulumi:"zoneBalance"`
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a VirtualMachineScaleSet resource.
type VirtualMachineScaleSetArgs struct {
	// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities AdditionalCapabilitiesPtrInput
	// Policy for automatic repairs.
	AutomaticRepairsPolicy AutomaticRepairsPolicyPtrInput
	// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrInput
	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision pulumi.BoolPtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanPtrInput
	// Fault Domain count for each placement group.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
	ScaleInPolicy ScaleInPolicyPtrInput
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup pulumi.BoolPtrInput
	// The virtual machine scale set sku.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The upgrade policy.
	UpgradePolicy UpgradePolicyPtrInput
	// The virtual machine profile.
	VirtualMachineProfile VirtualMachineScaleSetVMProfilePtrInput
	// The name of the VM scale set to create or update.
	VmScaleSetName pulumi.StringPtrInput
	// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
	ZoneBalance pulumi.BoolPtrInput
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
	Zones pulumi.StringArrayInput
}

func (VirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetArgs)(nil)).Elem()
}

type VirtualMachineScaleSetInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput
	ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput
}

func (*VirtualMachineScaleSet) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSet)(nil)).Elem()
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return i.ToVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOutput)
}

type VirtualMachineScaleSetOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSet)(nil)).Elem()
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return o
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return o
}

// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
func (o VirtualMachineScaleSetOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) AdditionalCapabilitiesResponsePtrOutput {
		return v.AdditionalCapabilities
	}).(AdditionalCapabilitiesResponsePtrOutput)
}

// Policy for automatic repairs.
func (o VirtualMachineScaleSetOutput) AutomaticRepairsPolicy() AutomaticRepairsPolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) AutomaticRepairsPolicyResponsePtrOutput {
		return v.AutomaticRepairsPolicy
	}).(AutomaticRepairsPolicyResponsePtrOutput)
}

// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
func (o VirtualMachineScaleSetOutput) DoNotRunExtensionsOnOverprovisionedVMs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.DoNotRunExtensionsOnOverprovisionedVMs }).(pulumi.BoolPtrOutput)
}

// The identity of the virtual machine scale set, if configured.
func (o VirtualMachineScaleSetOutput) Identity() VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) VirtualMachineScaleSetIdentityResponsePtrOutput { return v.Identity }).(VirtualMachineScaleSetIdentityResponsePtrOutput)
}

// Resource location
func (o VirtualMachineScaleSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o VirtualMachineScaleSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
func (o VirtualMachineScaleSetOutput) Overprovision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.Overprovision }).(pulumi.BoolPtrOutput)
}

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
func (o VirtualMachineScaleSetOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

// Fault Domain count for each placement group.
func (o VirtualMachineScaleSetOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.IntPtrOutput { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o VirtualMachineScaleSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
func (o VirtualMachineScaleSetOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) SubResourceResponsePtrOutput { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
func (o VirtualMachineScaleSetOutput) ScaleInPolicy() ScaleInPolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) ScaleInPolicyResponsePtrOutput { return v.ScaleInPolicy }).(ScaleInPolicyResponsePtrOutput)
}

// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
func (o VirtualMachineScaleSetOutput) SinglePlacementGroup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.SinglePlacementGroup }).(pulumi.BoolPtrOutput)
}

// The virtual machine scale set sku.
func (o VirtualMachineScaleSetOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o VirtualMachineScaleSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o VirtualMachineScaleSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the ID which uniquely identifies a Virtual Machine Scale Set.
func (o VirtualMachineScaleSetOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

// The upgrade policy.
func (o VirtualMachineScaleSetOutput) UpgradePolicy() UpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) UpgradePolicyResponsePtrOutput { return v.UpgradePolicy }).(UpgradePolicyResponsePtrOutput)
}

// The virtual machine profile.
func (o VirtualMachineScaleSetOutput) VirtualMachineProfile() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) VirtualMachineScaleSetVMProfileResponsePtrOutput {
		return v.VirtualMachineProfile
	}).(VirtualMachineScaleSetVMProfileResponsePtrOutput)
}

// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
func (o VirtualMachineScaleSetOutput) ZoneBalance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.ZoneBalance }).(pulumi.BoolPtrOutput)
}

// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
func (o VirtualMachineScaleSetOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetOutput{})
}