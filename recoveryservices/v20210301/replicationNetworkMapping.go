// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network Mapping model. Ideally it should have been possible to inherit this class from prev version in InheritedModels as long as there is no difference in structure or method signature. Since there were no base Models for certain fields and methods viz NetworkMappingProperties and Load with required return type, the class has been introduced in its entirety with references to base models to facilitate extensions in subsequent versions.
type ReplicationNetworkMapping struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The Network Mapping Properties.
	Properties NetworkMappingPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationNetworkMapping registers a new resource with the given unique name, arguments, and options.
func NewReplicationNetworkMapping(ctx *pulumi.Context,
	name string, args *ReplicationNetworkMappingArgs, opts ...pulumi.ResourceOption) (*ReplicationNetworkMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.NetworkName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230201:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230401:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230601:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230801:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20240101:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20240201:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20240401:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20241001:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationNetworkMapping"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReplicationNetworkMapping
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210301:ReplicationNetworkMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationNetworkMapping gets an existing ReplicationNetworkMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationNetworkMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationNetworkMappingState, opts ...pulumi.ResourceOption) (*ReplicationNetworkMapping, error) {
	var resource ReplicationNetworkMapping
	err := ctx.ReadResource("azure-native:recoveryservices/v20210301:ReplicationNetworkMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationNetworkMapping resources.
type replicationNetworkMappingState struct {
}

type ReplicationNetworkMappingState struct {
}

func (ReplicationNetworkMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationNetworkMappingState)(nil)).Elem()
}

type replicationNetworkMappingArgs struct {
	// Primary fabric name.
	FabricName string `pulumi:"fabricName"`
	// Network mapping name.
	NetworkMappingName *string `pulumi:"networkMappingName"`
	// Primary network name.
	NetworkName string `pulumi:"networkName"`
	// Input properties for creating network mapping.
	Properties *CreateNetworkMappingInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationNetworkMapping resource.
type ReplicationNetworkMappingArgs struct {
	// Primary fabric name.
	FabricName pulumi.StringInput
	// Network mapping name.
	NetworkMappingName pulumi.StringPtrInput
	// Primary network name.
	NetworkName pulumi.StringInput
	// Input properties for creating network mapping.
	Properties CreateNetworkMappingInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationNetworkMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationNetworkMappingArgs)(nil)).Elem()
}

type ReplicationNetworkMappingInput interface {
	pulumi.Input

	ToReplicationNetworkMappingOutput() ReplicationNetworkMappingOutput
	ToReplicationNetworkMappingOutputWithContext(ctx context.Context) ReplicationNetworkMappingOutput
}

func (*ReplicationNetworkMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationNetworkMapping)(nil)).Elem()
}

func (i *ReplicationNetworkMapping) ToReplicationNetworkMappingOutput() ReplicationNetworkMappingOutput {
	return i.ToReplicationNetworkMappingOutputWithContext(context.Background())
}

func (i *ReplicationNetworkMapping) ToReplicationNetworkMappingOutputWithContext(ctx context.Context) ReplicationNetworkMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationNetworkMappingOutput)
}

type ReplicationNetworkMappingOutput struct{ *pulumi.OutputState }

func (ReplicationNetworkMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationNetworkMapping)(nil)).Elem()
}

func (o ReplicationNetworkMappingOutput) ToReplicationNetworkMappingOutput() ReplicationNetworkMappingOutput {
	return o
}

func (o ReplicationNetworkMappingOutput) ToReplicationNetworkMappingOutputWithContext(ctx context.Context) ReplicationNetworkMappingOutput {
	return o
}

// Resource Location
func (o ReplicationNetworkMappingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationNetworkMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Network Mapping Properties.
func (o ReplicationNetworkMappingOutput) Properties() NetworkMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) NetworkMappingPropertiesResponseOutput { return v.Properties }).(NetworkMappingPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationNetworkMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationNetworkMappingOutput{})
}
