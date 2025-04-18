// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Microsoft.AwsConnector resource
//
// Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2024-12-01.
type EfsMountTarget struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties EfsMountTargetPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEfsMountTarget registers a new resource with the given unique name, arguments, and options.
func NewEfsMountTarget(ctx *pulumi.Context,
	name string, args *EfsMountTargetArgs, opts ...pulumi.ResourceOption) (*EfsMountTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:EfsMountTarget"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EfsMountTarget
	err := ctx.RegisterResource("azure-native:awsconnector:EfsMountTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEfsMountTarget gets an existing EfsMountTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEfsMountTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EfsMountTargetState, opts ...pulumi.ResourceOption) (*EfsMountTarget, error) {
	var resource EfsMountTarget
	err := ctx.ReadResource("azure-native:awsconnector:EfsMountTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EfsMountTarget resources.
type efsMountTargetState struct {
}

type EfsMountTargetState struct {
}

func (EfsMountTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*efsMountTargetState)(nil)).Elem()
}

type efsMountTargetArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of EfsMountTarget
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *EfsMountTargetProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EfsMountTarget resource.
type EfsMountTargetArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of EfsMountTarget
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties EfsMountTargetPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (EfsMountTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*efsMountTargetArgs)(nil)).Elem()
}

type EfsMountTargetInput interface {
	pulumi.Input

	ToEfsMountTargetOutput() EfsMountTargetOutput
	ToEfsMountTargetOutputWithContext(ctx context.Context) EfsMountTargetOutput
}

func (*EfsMountTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**EfsMountTarget)(nil)).Elem()
}

func (i *EfsMountTarget) ToEfsMountTargetOutput() EfsMountTargetOutput {
	return i.ToEfsMountTargetOutputWithContext(context.Background())
}

func (i *EfsMountTarget) ToEfsMountTargetOutputWithContext(ctx context.Context) EfsMountTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsMountTargetOutput)
}

type EfsMountTargetOutput struct{ *pulumi.OutputState }

func (EfsMountTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EfsMountTarget)(nil)).Elem()
}

func (o EfsMountTargetOutput) ToEfsMountTargetOutput() EfsMountTargetOutput {
	return o
}

func (o EfsMountTargetOutput) ToEfsMountTargetOutputWithContext(ctx context.Context) EfsMountTargetOutput {
	return o
}

// The Azure API version of the resource.
func (o EfsMountTargetOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EfsMountTarget) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o EfsMountTargetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EfsMountTarget) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o EfsMountTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EfsMountTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o EfsMountTargetOutput) Properties() EfsMountTargetPropertiesResponseOutput {
	return o.ApplyT(func(v *EfsMountTarget) EfsMountTargetPropertiesResponseOutput { return v.Properties }).(EfsMountTargetPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EfsMountTargetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EfsMountTarget) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o EfsMountTargetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EfsMountTarget) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EfsMountTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EfsMountTarget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EfsMountTargetOutput{})
}
