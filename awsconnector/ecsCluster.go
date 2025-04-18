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
type EcsCluster struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties EcsClusterPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEcsCluster registers a new resource with the given unique name, arguments, and options.
func NewEcsCluster(ctx *pulumi.Context,
	name string, args *EcsClusterArgs, opts ...pulumi.ResourceOption) (*EcsCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:EcsCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EcsCluster
	err := ctx.RegisterResource("azure-native:awsconnector:EcsCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsCluster gets an existing EcsCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsClusterState, opts ...pulumi.ResourceOption) (*EcsCluster, error) {
	var resource EcsCluster
	err := ctx.ReadResource("azure-native:awsconnector:EcsCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsCluster resources.
type ecsClusterState struct {
}

type EcsClusterState struct {
}

func (EcsClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsClusterState)(nil)).Elem()
}

type ecsClusterArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of EcsCluster
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *EcsClusterProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EcsCluster resource.
type EcsClusterArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of EcsCluster
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties EcsClusterPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (EcsClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsClusterArgs)(nil)).Elem()
}

type EcsClusterInput interface {
	pulumi.Input

	ToEcsClusterOutput() EcsClusterOutput
	ToEcsClusterOutputWithContext(ctx context.Context) EcsClusterOutput
}

func (*EcsCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsCluster)(nil)).Elem()
}

func (i *EcsCluster) ToEcsClusterOutput() EcsClusterOutput {
	return i.ToEcsClusterOutputWithContext(context.Background())
}

func (i *EcsCluster) ToEcsClusterOutputWithContext(ctx context.Context) EcsClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsClusterOutput)
}

type EcsClusterOutput struct{ *pulumi.OutputState }

func (EcsClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsCluster)(nil)).Elem()
}

func (o EcsClusterOutput) ToEcsClusterOutput() EcsClusterOutput {
	return o
}

func (o EcsClusterOutput) ToEcsClusterOutputWithContext(ctx context.Context) EcsClusterOutput {
	return o
}

// The Azure API version of the resource.
func (o EcsClusterOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsCluster) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o EcsClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o EcsClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o EcsClusterOutput) Properties() EcsClusterPropertiesResponseOutput {
	return o.ApplyT(func(v *EcsCluster) EcsClusterPropertiesResponseOutput { return v.Properties }).(EcsClusterPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EcsClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EcsCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o EcsClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EcsCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EcsClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EcsClusterOutput{})
}
