// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240413

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a managed private endpoint.
type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	// The groupId in which the managed private endpoint is created.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARM resource ID of the resource for which the managed private endpoint is created.
	PrivateLinkResourceId pulumi.StringOutput `pulumi:"privateLinkResourceId"`
	// The region of the resource to which the managed private endpoint is created.
	PrivateLinkResourceRegion pulumi.StringPtrOutput `pulumi:"privateLinkResourceRegion"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The user request message.
	RequestMessage pulumi.StringPtrOutput `pulumi:"requestMessage"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.PrivateLinkResourceId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:kusto:ManagedPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure-native:kusto/v20240413:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrivateEndpoint gets an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure-native:kusto/v20240413:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
type managedPrivateEndpointState struct {
}

type ManagedPrivateEndpointState struct {
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The groupId in which the managed private endpoint is created.
	GroupId string `pulumi:"groupId"`
	// The name of the managed private endpoint.
	ManagedPrivateEndpointName *string `pulumi:"managedPrivateEndpointName"`
	// The ARM resource ID of the resource for which the managed private endpoint is created.
	PrivateLinkResourceId string `pulumi:"privateLinkResourceId"`
	// The region of the resource to which the managed private endpoint is created.
	PrivateLinkResourceRegion *string `pulumi:"privateLinkResourceRegion"`
	// The user request message.
	RequestMessage *string `pulumi:"requestMessage"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedPrivateEndpoint resource.
type ManagedPrivateEndpointArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The groupId in which the managed private endpoint is created.
	GroupId pulumi.StringInput
	// The name of the managed private endpoint.
	ManagedPrivateEndpointName pulumi.StringPtrInput
	// The ARM resource ID of the resource for which the managed private endpoint is created.
	PrivateLinkResourceId pulumi.StringInput
	// The region of the resource to which the managed private endpoint is created.
	PrivateLinkResourceRegion pulumi.StringPtrInput
	// The user request message.
	RequestMessage pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

// The groupId in which the managed private endpoint is created.
func (o ManagedPrivateEndpointOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the resource
func (o ManagedPrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARM resource ID of the resource for which the managed private endpoint is created.
func (o ManagedPrivateEndpointOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

// The region of the resource to which the managed private endpoint is created.
func (o ManagedPrivateEndpointOutput) PrivateLinkResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringPtrOutput { return v.PrivateLinkResourceRegion }).(pulumi.StringPtrOutput)
}

// The provisioned state of the resource.
func (o ManagedPrivateEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The user request message.
func (o ManagedPrivateEndpointOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringPtrOutput { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagedPrivateEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedPrivateEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
}
