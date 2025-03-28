// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Instance dataflowProfile resource
type DataFlowProfile struct {
	pulumi.CustomResourceState

	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties DataFlowProfilePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataFlowProfile registers a new resource with the given unique name, arguments, and options.
func NewDataFlowProfile(ctx *pulumi.Context,
	name string, args *DataFlowProfileArgs, opts ...pulumi.ResourceOption) (*DataFlowProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToDataFlowProfilePropertiesPtrOutput().ApplyT(func(v *DataFlowProfileProperties) *DataFlowProfileProperties { return v.Defaults() }).(DataFlowProfilePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotoperations/v20240815preview:DataFlowProfile"),
		},
		{
			Type: pulumi.String("azure-native:iotoperations/v20240915preview:DataFlowProfile"),
		},
		{
			Type: pulumi.String("azure-native:iotoperations/v20241101:DataFlowProfile"),
		},
		{
			Type: pulumi.String("azure-native:iotoperations/v20250401:DataFlowProfile"),
		},
		{
			Type: pulumi.String("azure-native:iotoperations:DataFlowProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataFlowProfile
	err := ctx.RegisterResource("azure-native:iotoperations/v20240701preview:DataFlowProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataFlowProfile gets an existing DataFlowProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataFlowProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataFlowProfileState, opts ...pulumi.ResourceOption) (*DataFlowProfile, error) {
	var resource DataFlowProfile
	err := ctx.ReadResource("azure-native:iotoperations/v20240701preview:DataFlowProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataFlowProfile resources.
type dataFlowProfileState struct {
}

type DataFlowProfileState struct {
}

func (DataFlowProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowProfileState)(nil)).Elem()
}

type dataFlowProfileArgs struct {
	// Name of Instance dataflowProfile resource
	DataflowProfileName *string `pulumi:"dataflowProfileName"`
	// Edge location of the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Name of instance.
	InstanceName string `pulumi:"instanceName"`
	// The resource-specific properties for this resource.
	Properties *DataFlowProfileProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DataFlowProfile resource.
type DataFlowProfileArgs struct {
	// Name of Instance dataflowProfile resource
	DataflowProfileName pulumi.StringPtrInput
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationInput
	// Name of instance.
	InstanceName pulumi.StringInput
	// The resource-specific properties for this resource.
	Properties DataFlowProfilePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (DataFlowProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowProfileArgs)(nil)).Elem()
}

type DataFlowProfileInput interface {
	pulumi.Input

	ToDataFlowProfileOutput() DataFlowProfileOutput
	ToDataFlowProfileOutputWithContext(ctx context.Context) DataFlowProfileOutput
}

func (*DataFlowProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlowProfile)(nil)).Elem()
}

func (i *DataFlowProfile) ToDataFlowProfileOutput() DataFlowProfileOutput {
	return i.ToDataFlowProfileOutputWithContext(context.Background())
}

func (i *DataFlowProfile) ToDataFlowProfileOutputWithContext(ctx context.Context) DataFlowProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowProfileOutput)
}

type DataFlowProfileOutput struct{ *pulumi.OutputState }

func (DataFlowProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlowProfile)(nil)).Elem()
}

func (o DataFlowProfileOutput) ToDataFlowProfileOutput() DataFlowProfileOutput {
	return o
}

func (o DataFlowProfileOutput) ToDataFlowProfileOutputWithContext(ctx context.Context) DataFlowProfileOutput {
	return o
}

// Edge location of the resource.
func (o DataFlowProfileOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *DataFlowProfile) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The name of the resource
func (o DataFlowProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlowProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o DataFlowProfileOutput) Properties() DataFlowProfilePropertiesResponseOutput {
	return o.ApplyT(func(v *DataFlowProfile) DataFlowProfilePropertiesResponseOutput { return v.Properties }).(DataFlowProfilePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DataFlowProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataFlowProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataFlowProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlowProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataFlowProfileOutput{})
}
