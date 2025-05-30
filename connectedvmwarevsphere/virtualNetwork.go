// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the virtualNetwork.
//
// Uses Azure REST API version 2023-12-01. In version 2.x of the Azure Native provider, it used API version 2022-07-15-preview.
//
// Other available API versions: 2022-07-15-preview, 2023-03-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type VirtualNetwork struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the virtual network.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets or sets the vCenter Managed Object name for the virtual network.
	MoName pulumi.StringOutput `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual network.
	MoRefId pulumi.StringPtrOutput `pulumi:"moRefId"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource status information.
	Statuses ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Gets or sets the ARM Id of the vCenter resource in which this template resides.
	VCenterId pulumi.StringPtrOutput `pulumi:"vCenterId"`
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20230301preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20231001:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20231201:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetwork gets an existing VirtualNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:VirtualNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetwork resources.
type virtualNetworkState struct {
}

type VirtualNetworkState struct {
}

func (VirtualNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkState)(nil)).Elem()
}

type virtualNetworkArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the virtual network.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual network.
	MoRefId *string `pulumi:"moRefId"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the ARM Id of the vCenter resource in which this template resides.
	VCenterId *string `pulumi:"vCenterId"`
	// Name of the virtual network resource.
	VirtualNetworkName *string `pulumi:"virtualNetworkName"`
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Gets or sets the inventory Item ID for the virtual network.
	InventoryItemId pulumi.StringPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual network.
	MoRefId pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapInput
	// Gets or sets the ARM Id of the vCenter resource in which this template resides.
	VCenterId pulumi.StringPtrInput
	// Name of the virtual network resource.
	VirtualNetworkName pulumi.StringPtrInput
}

func (VirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkArgs)(nil)).Elem()
}

type VirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkOutput() VirtualNetworkOutput
	ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput
}

func (*VirtualNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetwork)(nil)).Elem()
}

func (i *VirtualNetwork) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return i.ToVirtualNetworkOutputWithContext(context.Background())
}

func (i *VirtualNetwork) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkOutput)
}

type VirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return o
}

// The Azure API version of the resource.
func (o VirtualNetworkOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o VirtualNetworkOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the extended location.
func (o VirtualNetworkOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the inventory Item ID for the virtual network.
func (o VirtualNetworkOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o VirtualNetworkOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o VirtualNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the virtual network.
func (o VirtualNetworkOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual network.
func (o VirtualNetworkOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o VirtualNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state.
func (o VirtualNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o VirtualNetworkOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o VirtualNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o VirtualNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o VirtualNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o VirtualNetworkOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this template resides.
func (o VirtualNetworkOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}
