// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies information about the capacity reservation.
type CapacityReservation struct {
	pulumi.CustomResourceState

	// The Capacity reservation instance view.
	InstanceView CapacityReservationInstanceViewResponseOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the value of fault domain count that Capacity Reservation supports for requested VM size.<br>NOTE: The fault domain count specified for a resource (like virtual machines scale set) must be less than or equal to this value if it deploys using capacity reservation.<br><br>Minimum api-version: 2022-08-01.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The date time when the capacity reservation was last updated.
	ProvisioningTime pulumi.StringOutput `pulumi:"provisioningTime"`
	// A unique id generated and assigned to the capacity reservation by the platform which does not change throughout the lifetime of the resource.
	ReservationId pulumi.StringOutput `pulumi:"reservationId"`
	// SKU of the resource for which capacity needs be reserved. The SKU name and capacity is required to be set. Currently VM Skus with the capability called 'CapacityReservationSupported' set to true are supported. Refer to List Microsoft.Compute SKUs in a region (https://docs.microsoft.com/rest/api/compute/resourceskus/list) for supported values.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the time at which the Capacity Reservation resource was created.<br><br>Minimum api-version: 2021-11-01.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of all virtual machine resource ids that are associated with the capacity reservation.
	VirtualMachinesAssociated SubResourceReadOnlyResponseArrayOutput `pulumi:"virtualMachinesAssociated"`
	// Availability Zone to use for this capacity reservation. The zone has to be single value and also should be part for the list of zones specified during the capacity reservation group creation. The zone can be assigned only during creation. If not provided, the reservation supports only non-zonal deployments. If provided, enforces VM/VMSS using this capacity reservation to be in same zone.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCapacityReservation registers a new resource with the given unique name, arguments, and options.
func NewCapacityReservation(ctx *pulumi.Context,
	name string, args *CapacityReservationArgs, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityReservationGroupName == nil {
		return nil, errors.New("invalid value for required argument 'CapacityReservationGroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:CapacityReservation"),
		},
	})
	opts = append(opts, aliases)
	var resource CapacityReservation
	err := ctx.RegisterResource("azure-native:compute/v20221101:CapacityReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityReservation gets an existing CapacityReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationState, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	var resource CapacityReservation
	err := ctx.ReadResource("azure-native:compute/v20221101:CapacityReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityReservation resources.
type capacityReservationState struct {
}

type CapacityReservationState struct {
}

func (CapacityReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationState)(nil)).Elem()
}

type capacityReservationArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName string `pulumi:"capacityReservationGroupName"`
	// The name of the capacity reservation.
	CapacityReservationName *string `pulumi:"capacityReservationName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of the resource for which capacity needs be reserved. The SKU name and capacity is required to be set. Currently VM Skus with the capability called 'CapacityReservationSupported' set to true are supported. Refer to List Microsoft.Compute SKUs in a region (https://docs.microsoft.com/rest/api/compute/resourceskus/list) for supported values.
	Sku Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Availability Zone to use for this capacity reservation. The zone has to be single value and also should be part for the list of zones specified during the capacity reservation group creation. The zone can be assigned only during creation. If not provided, the reservation supports only non-zonal deployments. If provided, enforces VM/VMSS using this capacity reservation to be in same zone.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a CapacityReservation resource.
type CapacityReservationArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName pulumi.StringInput
	// The name of the capacity reservation.
	CapacityReservationName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// SKU of the resource for which capacity needs be reserved. The SKU name and capacity is required to be set. Currently VM Skus with the capability called 'CapacityReservationSupported' set to true are supported. Refer to List Microsoft.Compute SKUs in a region (https://docs.microsoft.com/rest/api/compute/resourceskus/list) for supported values.
	Sku SkuInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Availability Zone to use for this capacity reservation. The zone has to be single value and also should be part for the list of zones specified during the capacity reservation group creation. The zone can be assigned only during creation. If not provided, the reservation supports only non-zonal deployments. If provided, enforces VM/VMSS using this capacity reservation to be in same zone.
	Zones pulumi.StringArrayInput
}

func (CapacityReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationArgs)(nil)).Elem()
}

type CapacityReservationInput interface {
	pulumi.Input

	ToCapacityReservationOutput() CapacityReservationOutput
	ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput
}

func (*CapacityReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil)).Elem()
}

func (i *CapacityReservation) ToCapacityReservationOutput() CapacityReservationOutput {
	return i.ToCapacityReservationOutputWithContext(context.Background())
}

func (i *CapacityReservation) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationOutput)
}

type CapacityReservationOutput struct{ *pulumi.OutputState }

func (CapacityReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil)).Elem()
}

func (o CapacityReservationOutput) ToCapacityReservationOutput() CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return o
}

// The Capacity reservation instance view.
func (o CapacityReservationOutput) InstanceView() CapacityReservationInstanceViewResponseOutput {
	return o.ApplyT(func(v *CapacityReservation) CapacityReservationInstanceViewResponseOutput { return v.InstanceView }).(CapacityReservationInstanceViewResponseOutput)
}

// Resource location
func (o CapacityReservationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o CapacityReservationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the value of fault domain count that Capacity Reservation supports for requested VM size.<br>NOTE: The fault domain count specified for a resource (like virtual machines scale set) must be less than or equal to this value if it deploys using capacity reservation.<br><br>Minimum api-version: 2022-08-01.
func (o CapacityReservationOutput) PlatformFaultDomainCount() pulumi.IntOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.IntOutput { return v.PlatformFaultDomainCount }).(pulumi.IntOutput)
}

// The provisioning state, which only appears in the response.
func (o CapacityReservationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The date time when the capacity reservation was last updated.
func (o CapacityReservationOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.ProvisioningTime }).(pulumi.StringOutput)
}

// A unique id generated and assigned to the capacity reservation by the platform which does not change throughout the lifetime of the resource.
func (o CapacityReservationOutput) ReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.ReservationId }).(pulumi.StringOutput)
}

// SKU of the resource for which capacity needs be reserved. The SKU name and capacity is required to be set. Currently VM Skus with the capability called 'CapacityReservationSupported' set to true are supported. Refer to List Microsoft.Compute SKUs in a region (https://docs.microsoft.com/rest/api/compute/resourceskus/list) for supported values.
func (o CapacityReservationOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *CapacityReservation) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Resource tags
func (o CapacityReservationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the time at which the Capacity Reservation resource was created.<br><br>Minimum api-version: 2021-11-01.
func (o CapacityReservationOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o CapacityReservationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of all virtual machine resource ids that are associated with the capacity reservation.
func (o CapacityReservationOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservation) SubResourceReadOnlyResponseArrayOutput {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

// Availability Zone to use for this capacity reservation. The zone has to be single value and also should be part for the list of zones specified during the capacity reservation group creation. The zone can be assigned only during creation. If not provided, the reservation supports only non-zonal deployments. If provided, enforces VM/VMSS using this capacity reservation to be in same zone.
func (o CapacityReservationOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationOutput{})
}