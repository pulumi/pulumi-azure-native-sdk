// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Managed Network Peering Policy resource
type ManagedNetworkPeeringPolicy struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the properties of a Managed Network Policy
	Properties ManagedNetworkPeeringPolicyPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedNetworkPeeringPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedNetworkPeeringPolicy(ctx *pulumi.Context,
	name string, args *ManagedNetworkPeeringPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedNetworkPeeringPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetwork:ManagedNetworkPeeringPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedNetworkPeeringPolicy
	err := ctx.RegisterResource("azure-native:managednetwork/v20190601preview:ManagedNetworkPeeringPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedNetworkPeeringPolicy gets an existing ManagedNetworkPeeringPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedNetworkPeeringPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkPeeringPolicyState, opts ...pulumi.ResourceOption) (*ManagedNetworkPeeringPolicy, error) {
	var resource ManagedNetworkPeeringPolicy
	err := ctx.ReadResource("azure-native:managednetwork/v20190601preview:ManagedNetworkPeeringPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedNetworkPeeringPolicy resources.
type managedNetworkPeeringPolicyState struct {
}

type ManagedNetworkPeeringPolicyState struct {
}

func (ManagedNetworkPeeringPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkPeeringPolicyState)(nil)).Elem()
}

type managedNetworkPeeringPolicyArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Managed Network.
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	// The name of the Managed Network Peering Policy.
	ManagedNetworkPeeringPolicyName *string `pulumi:"managedNetworkPeeringPolicyName"`
	// Gets or sets the properties of a Managed Network Policy
	Properties *ManagedNetworkPeeringPolicyProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedNetworkPeeringPolicy resource.
type ManagedNetworkPeeringPolicyArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Managed Network.
	ManagedNetworkName pulumi.StringInput
	// The name of the Managed Network Peering Policy.
	ManagedNetworkPeeringPolicyName pulumi.StringPtrInput
	// Gets or sets the properties of a Managed Network Policy
	Properties ManagedNetworkPeeringPolicyPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ManagedNetworkPeeringPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkPeeringPolicyArgs)(nil)).Elem()
}

type ManagedNetworkPeeringPolicyInput interface {
	pulumi.Input

	ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput
	ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput
}

func (*ManagedNetworkPeeringPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicy)(nil)).Elem()
}

func (i *ManagedNetworkPeeringPolicy) ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput {
	return i.ToManagedNetworkPeeringPolicyOutputWithContext(context.Background())
}

func (i *ManagedNetworkPeeringPolicy) ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkPeeringPolicyOutput)
}

type ManagedNetworkPeeringPolicyOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicy)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyOutput) ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyOutput) ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput {
	return o
}

// The geo-location where the resource lives
func (o ManagedNetworkPeeringPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ManagedNetworkPeeringPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the properties of a Managed Network Policy
func (o ManagedNetworkPeeringPolicyOutput) Properties() ManagedNetworkPeeringPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) ManagedNetworkPeeringPolicyPropertiesResponseOutput {
		return v.Properties
	}).(ManagedNetworkPeeringPolicyPropertiesResponseOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o ManagedNetworkPeeringPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyOutput{})
}