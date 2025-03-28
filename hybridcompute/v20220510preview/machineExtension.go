// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220510preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Machine Extension.
type MachineExtension struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes Machine Extension Properties.
	Properties MachineExtensionPropertiesResponseOutput `pulumi:"properties"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachineExtension registers a new resource with the given unique name, arguments, and options.
func NewMachineExtension(ctx *pulumi.Context,
	name string, args *MachineExtensionArgs, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190802preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20191212:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220811preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221110:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221227:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221227preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20230315preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20230620preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20231003preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240331preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240520preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240710:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240731preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240910preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20241110preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20250113:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute:MachineExtension"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MachineExtension
	err := ctx.RegisterResource("azure-native:hybridcompute/v20220510preview:MachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineExtension gets an existing MachineExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineExtensionState, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	var resource MachineExtension
	err := ctx.ReadResource("azure-native:hybridcompute/v20220510preview:MachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineExtension resources.
type machineExtensionState struct {
}

type MachineExtensionState struct {
}

func (MachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionState)(nil)).Elem()
}

type machineExtensionArgs struct {
	// The name of the machine extension.
	ExtensionName *string `pulumi:"extensionName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the machine where the extension should be created or updated.
	MachineName string `pulumi:"machineName"`
	// Describes Machine Extension Properties.
	Properties *MachineExtensionProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MachineExtension resource.
type MachineExtensionArgs struct {
	// The name of the machine extension.
	ExtensionName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the machine where the extension should be created or updated.
	MachineName pulumi.StringInput
	// Describes Machine Extension Properties.
	Properties MachineExtensionPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionArgs)(nil)).Elem()
}

type MachineExtensionInput interface {
	pulumi.Input

	ToMachineExtensionOutput() MachineExtensionOutput
	ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput
}

func (*MachineExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (i *MachineExtension) ToMachineExtensionOutput() MachineExtensionOutput {
	return i.ToMachineExtensionOutputWithContext(context.Background())
}

func (i *MachineExtension) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionOutput)
}

type MachineExtensionOutput struct{ *pulumi.OutputState }

func (MachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (o MachineExtensionOutput) ToMachineExtensionOutput() MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return o
}

// The geo-location where the resource lives
func (o MachineExtensionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MachineExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes Machine Extension Properties.
func (o MachineExtensionOutput) Properties() MachineExtensionPropertiesResponseOutput {
	return o.ApplyT(func(v *MachineExtension) MachineExtensionPropertiesResponseOutput { return v.Properties }).(MachineExtensionPropertiesResponseOutput)
}

// The system meta data relating to this resource.
func (o MachineExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachineExtension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MachineExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MachineExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineExtensionOutput{})
}
