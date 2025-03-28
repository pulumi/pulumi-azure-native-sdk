// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Mobile network resource.
type MobileNetwork struct {
	pulumi.CustomResourceState

	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrOutput `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrOutput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrOutput `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the mobile network resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The unique public land mobile network identifier for the network. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
	PublicLandMobileNetworkIdentifier PlmnIdResponseOutput `pulumi:"publicLandMobileNetworkIdentifier"`
	// The mobile network resource identifier
	ServiceKey pulumi.StringOutput `pulumi:"serviceKey"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMobileNetwork registers a new resource with the given unique name, arguments, and options.
func NewMobileNetwork(ctx *pulumi.Context,
	name string, args *MobileNetworkArgs, opts ...pulumi.ResourceOption) (*MobileNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicLandMobileNetworkIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'PublicLandMobileNetworkIdentifier'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230601:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230901:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20240201:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20240401:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork:MobileNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MobileNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220401preview:MobileNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMobileNetwork gets an existing MobileNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMobileNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MobileNetworkState, opts ...pulumi.ResourceOption) (*MobileNetwork, error) {
	var resource MobileNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220401preview:MobileNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MobileNetwork resources.
type mobileNetworkState struct {
}

type MobileNetworkState struct {
}

func (MobileNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*mobileNetworkState)(nil)).Elem()
}

type mobileNetworkArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mobile network.
	MobileNetworkName *string `pulumi:"mobileNetworkName"`
	// The unique public land mobile network identifier for the network. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
	PublicLandMobileNetworkIdentifier PlmnId `pulumi:"publicLandMobileNetworkIdentifier"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MobileNetwork resource.
type MobileNetworkArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mobile network.
	MobileNetworkName pulumi.StringPtrInput
	// The unique public land mobile network identifier for the network. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
	PublicLandMobileNetworkIdentifier PlmnIdInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MobileNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mobileNetworkArgs)(nil)).Elem()
}

type MobileNetworkInput interface {
	pulumi.Input

	ToMobileNetworkOutput() MobileNetworkOutput
	ToMobileNetworkOutputWithContext(ctx context.Context) MobileNetworkOutput
}

func (*MobileNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetwork)(nil)).Elem()
}

func (i *MobileNetwork) ToMobileNetworkOutput() MobileNetworkOutput {
	return i.ToMobileNetworkOutputWithContext(context.Background())
}

func (i *MobileNetwork) ToMobileNetworkOutputWithContext(ctx context.Context) MobileNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobileNetworkOutput)
}

type MobileNetworkOutput struct{ *pulumi.OutputState }

func (MobileNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetwork)(nil)).Elem()
}

func (o MobileNetworkOutput) ToMobileNetworkOutput() MobileNetworkOutput {
	return o
}

func (o MobileNetworkOutput) ToMobileNetworkOutputWithContext(ctx context.Context) MobileNetworkOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o MobileNetworkOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o MobileNetworkOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o MobileNetworkOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o MobileNetworkOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o MobileNetworkOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o MobileNetworkOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o MobileNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MobileNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the mobile network resource.
func (o MobileNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique public land mobile network identifier for the network. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
func (o MobileNetworkOutput) PublicLandMobileNetworkIdentifier() PlmnIdResponseOutput {
	return o.ApplyT(func(v *MobileNetwork) PlmnIdResponseOutput { return v.PublicLandMobileNetworkIdentifier }).(PlmnIdResponseOutput)
}

// The mobile network resource identifier
func (o MobileNetworkOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.ServiceKey }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MobileNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MobileNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MobileNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MobileNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MobileNetworkOutput{})
}
