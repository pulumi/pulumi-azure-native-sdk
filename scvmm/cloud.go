// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Clouds resource definition.
// API Version: 2020-06-05-preview.
type Cloud struct {
	pulumi.CustomResourceState

	// Capacity of the cloud.
	CloudCapacity CloudCapacityResponseOutput `pulumi:"cloudCapacity"`
	// Name of the cloud in VMMServer.
	CloudName pulumi.StringOutput `pulumi:"cloudName"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of QoS policies available for the cloud.
	StorageQoSPolicies StorageQoSPolicyResponseArrayOutput `pulumi:"storageQoSPolicies"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique ID of the cloud.
	Uuid pulumi.StringPtrOutput `pulumi:"uuid"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrOutput `pulumi:"vmmServerId"`
}

// NewCloud registers a new resource with the given unique name, arguments, and options.
func NewCloud(ctx *pulumi.Context,
	name string, args *CloudArgs, opts ...pulumi.ResourceOption) (*Cloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:Cloud"),
		},
	})
	opts = append(opts, aliases)
	var resource Cloud
	err := ctx.RegisterResource("azure-native:scvmm:Cloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloud gets an existing Cloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudState, opts ...pulumi.ResourceOption) (*Cloud, error) {
	var resource Cloud
	err := ctx.ReadResource("azure-native:scvmm:Cloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cloud resources.
type cloudState struct {
}

type CloudState struct {
}

func (CloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudState)(nil)).Elem()
}

type cloudArgs struct {
	// Name of the Cloud.
	CloudName *string `pulumi:"cloudName"`
	// The extended location.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Unique ID of the cloud.
	Uuid *string `pulumi:"uuid"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

// The set of arguments for constructing a Cloud resource.
type CloudArgs struct {
	// Name of the Cloud.
	CloudName pulumi.StringPtrInput
	// The extended location.
	ExtendedLocation ExtendedLocationInput
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Unique ID of the cloud.
	Uuid pulumi.StringPtrInput
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrInput
}

func (CloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudArgs)(nil)).Elem()
}

type CloudInput interface {
	pulumi.Input

	ToCloudOutput() CloudOutput
	ToCloudOutputWithContext(ctx context.Context) CloudOutput
}

func (*Cloud) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloud)(nil)).Elem()
}

func (i *Cloud) ToCloudOutput() CloudOutput {
	return i.ToCloudOutputWithContext(context.Background())
}

func (i *Cloud) ToCloudOutputWithContext(ctx context.Context) CloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudOutput)
}

type CloudOutput struct{ *pulumi.OutputState }

func (CloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloud)(nil)).Elem()
}

func (o CloudOutput) ToCloudOutput() CloudOutput {
	return o
}

func (o CloudOutput) ToCloudOutputWithContext(ctx context.Context) CloudOutput {
	return o
}

// Capacity of the cloud.
func (o CloudOutput) CloudCapacity() CloudCapacityResponseOutput {
	return o.ApplyT(func(v *Cloud) CloudCapacityResponseOutput { return v.CloudCapacity }).(CloudCapacityResponseOutput)
}

// Name of the cloud in VMMServer.
func (o CloudOutput) CloudName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.CloudName }).(pulumi.StringOutput)
}

// The extended location.
func (o CloudOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Cloud) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Gets or sets the inventory Item ID for the resource.
func (o CloudOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o CloudOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o CloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o CloudOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of QoS policies available for the cloud.
func (o CloudOutput) StorageQoSPolicies() StorageQoSPolicyResponseArrayOutput {
	return o.ApplyT(func(v *Cloud) StorageQoSPolicyResponseArrayOutput { return v.StorageQoSPolicies }).(StorageQoSPolicyResponseArrayOutput)
}

// The system data.
func (o CloudOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cloud) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o CloudOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o CloudOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of the cloud.
func (o CloudOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o CloudOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cloud) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudOutput{})
}